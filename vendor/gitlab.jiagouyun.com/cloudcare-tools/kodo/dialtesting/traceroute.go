package dialtesting

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

// max config
const MAX_TIMEOUT = 30 * time.Second
const MAX_HOPS = 60
const MAX_RETRY = 3

// traceroute option
type TracerouteOption struct {
	Hops    int
	Retry   int
	Timeout time.Duration
}

// response for sent packet, may be failed response when timeout
type Response struct {
	From         net.IP
	ResponseTime time.Duration

	fail bool
}

// each retry response
type RouteItem struct {
	IP           string        `json:"ip"`
	ResponseTime time.Duration `json:"response_time"`
}

// route summary for each hop
type Route struct {
	Total   int           `json:"total"`
	Failed  int           `json:"failed"`
	Loss    float64       `json:"loss"`
	AvgCost time.Duration `json:"avg_cost"`
	MinCost time.Duration `json:"min_cost"`
	MaxCost time.Duration `json:"max_cost"`
	StdCost time.Duration `json:"std_cost"`
	Items   []*RouteItem  `json:"items"`
}

// sent packet
type Packet struct {
	ID  int
	Dst net.IP

	startTime time.Time
}

// traceroute specified host with max hops and timeout
type Traceroute struct {
	Host    string
	Hops    int
	Retry   int
	Timeout time.Duration

	routes     []*Route
	sentPacket *Packet
	response   chan *Response
	stopCh     chan interface{}
	id         uint32
	mu         sync.Mutex
}

// init config: hops, retry, timeout should not be greater than the max value.
func (t *Traceroute) init() {
	if t.Hops <= 0 {
		t.Hops = 30
	} else if t.Hops > MAX_HOPS {
		t.Hops = MAX_HOPS
	}

	if t.Retry <= 0 {
		t.Retry = 3
	} else if t.Retry > MAX_RETRY {
		t.Retry = MAX_RETRY
	}

	if t.Timeout <= 0 {
		t.Timeout = 10 * time.Second
	} else if t.Timeout > MAX_TIMEOUT {
		t.Timeout = MAX_TIMEOUT
	}

	t.routes = make([]*Route, 0)

	t.response = make(chan *Response)
	t.stopCh = make(chan interface{})

	t.id = t.getRandomID()
}

// getRandomID generate random id, max 60000
func (t *Traceroute) getRandomID() uint32 {
	rand.Seed(time.Now().UnixNano())
	return uint32(rand.Intn(60000))
}

func (t *Traceroute) Run() error {
	ips, err := net.LookupIP(t.Host)
	if err != nil {
		return err
	}

	t.init()

	if len(ips) == 0 {
		return fmt.Errorf("invalid host: %s", t.Host)
	}
	ip := ips[0]

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := t.startTrace(ip); err != nil {
			fmt.Println("start trace error: ", err)
		} else {
			fmt.Println("start trace end")
		}
	}()

	go func() {
		defer wg.Done()
		if err := t.listenICMP(); err != nil {
			fmt.Println("listen icmp error: ", err)
		} else {
			fmt.Println("listen icmp end")
		}
	}()
	wg.Wait()
	return nil
}

func (t *Traceroute) startTrace(ip net.IP) error {
	var icmpResponse *Response

	defer close(t.stopCh)

	for i := 1; i <= t.Hops; i++ {
		isReply := false
		routeItems := []*RouteItem{}
		responseTimes := []float64{}
		var minCost, maxCost time.Duration
		var failed int
		for j := 0; j < t.Retry; j++ {
			if err := t.sendICMP(ip, i); err != nil {
				return err
			}
			icmpResponse = <-t.response
			t.mu.Lock()
			t.sentPacket = nil
			t.mu.Unlock()
			routeItem := &RouteItem{
				IP:           icmpResponse.From.String(),
				ResponseTime: icmpResponse.ResponseTime,
			}

			if icmpResponse.fail {
				routeItem.IP = "*"
				failed++
			} else {
				if icmpResponse.From.String() == ip.String() {
					isReply = true
				}

				if icmpResponse.ResponseTime > 0 {
					if minCost == 0 || minCost > icmpResponse.ResponseTime {
						minCost = icmpResponse.ResponseTime
					}

					if maxCost == 0 || maxCost < icmpResponse.ResponseTime {
						maxCost = icmpResponse.ResponseTime
					}

					responseTimes = append(responseTimes, float64(icmpResponse.ResponseTime))
				}

			}

			routeItems = append(routeItems, routeItem)
		}

		loss, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(failed)*100/float64(t.Retry)), 64)

		route := &Route{
			Total:   t.Retry,
			Failed:  failed,
			Loss:    loss,
			MinCost: minCost,
			MaxCost: maxCost,
			StdCost: time.Duration(std(responseTimes)),
			Items:   routeItems,
		}
		t.routes = append(t.routes, route)

		if isReply {
			return nil
		}

	}

	return nil
}

func (t *Traceroute) dealPacket(from *net.IPAddr, data []byte) {
	packetRecvTime := time.Now()
	t.mu.Lock()
	defer t.mu.Unlock()
	packet := t.sentPacket

	if packet == nil {
		return
	}

	if packetRecvTime.Sub(packet.startTime) > t.Timeout {
		t.response <- &Response{fail: true}
		return
	}

	if len(data) == 0 {
		return
	}

	if from.IP.To4() == nil {
		return
	}

	msg, err := icmp.ParseMessage(1, data)

	if err != nil {
		fmt.Println(err)
		return
	}

	if msg.Type == ipv4.ICMPTypeEchoReply {
		echo := msg.Body.(*icmp.Echo)

		if echo.ID != packet.ID {
			return
		}

	} else {
		icmpData := t.getReplyData(msg)
		if len(icmpData) < ipv4.HeaderLen {
			return
		}

		var packetID int

		func() {
			switch icmpData[0] >> 4 {
			case ipv4.Version:
				header, err := ipv4.ParseHeader(icmpData)
				if err != nil {
					return
				}
				packetID = header.ID
			case ipv6.Version:
				header, err := ipv6.ParseHeader(icmpData)
				if err != nil {
					return
				}

				packetID = header.FlowLabel
			}

		}()
		if packetID != packet.ID {
			return
		}
	}

	t.response <- &Response{From: from.IP, ResponseTime: packetRecvTime.Sub(packet.startTime)}
}

func (t *Traceroute) listenICMP() error {
	var addr *net.IPAddr
	conn, err := net.ListenIP("ip4:icmp", addr)

	if err != nil {
		return err
	}

	defer conn.Close()

	for {
		select {
		case <-t.stopCh:
			return nil
		default:
		}

		buf := make([]byte, 1500)

		conn.SetDeadline(time.Now().Add(10 * time.Second))

		n, from, _ := conn.ReadFromIP(buf)

		go t.dealPacket(from, buf[:n])

	}

}

func (t *Traceroute) getReplyData(msg *icmp.Message) []byte {
	switch b := msg.Body.(type) {
	case *icmp.TimeExceeded:
		return b.Data
	case *icmp.DstUnreach:
		return b.Data
	case *icmp.ParamProb:
		return b.Data
	}

	return nil
}

func (t *Traceroute) sendICMP(ip net.IP, ttl int) error {
	if ip.To4() == nil {
		return fmt.Errorf("support ip version 4 only")
	}
	id := uint16(atomic.AddUint32(&t.id, 1))

	dst := net.ParseIP(ip.String())
	echoBody := &icmp.Echo{
		ID:  int(id),
		Seq: int(id),
	}
	msg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Body: echoBody,
	}

	p, err := msg.Marshal(nil)
	if err != nil {
		return err
	}

	ipHeader := &ipv4.Header{
		Version:  ipv4.Version,
		Len:      ipv4.HeaderLen,
		TotalLen: ipv4.HeaderLen + len(p),
		TOS:      16,
		ID:       int(id),
		Dst:      dst,
		Protocol: 1,
		TTL:      ttl,
	}

	buf, err := ipHeader.Marshal()

	if err != nil {
		return err
	}

	buf = append(buf, p...)

	conn, err := net.ListenIP("ip4:ip", nil)

	if err != nil {
		return err
	}
	defer conn.Close()

	raw, err := conn.SyscallConn()
	if err != nil {
		return err
	}

	_ = raw.Control(func(fd uintptr) {
		err = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.IP_HDRINCL, 1)
	})

	if err != nil {
		return err
	}

	t.mu.Lock()
	t.sentPacket = &Packet{ID: echoBody.ID, Dst: ipHeader.Dst, startTime: time.Now()}
	t.mu.Unlock()

	_, err = conn.WriteToIP(buf, &net.IPAddr{IP: dst})

	if err != nil {
		return err
	}

	return nil
}

func TracerouteIP(ip string, opt *TracerouteOption) (routes []*Route, err error) {
	if opt == nil {
		opt = &TracerouteOption{
			Hops:    30,
			Retry:   3,
			Timeout: 10 * time.Second,
		}
	}

	traceroute := Traceroute{
		Host:    ip,
		Hops:    opt.Hops,
		Retry:   opt.Retry,
		Timeout: opt.Timeout,
	}

	err = traceroute.Run()

	if err != nil {
		return
	}

	routes = traceroute.routes

	return
}

func mean(v []float64) float64 {
	var res float64 = 0
	var n int = len(v)
	for i := 0; i < n; i++ {
		res += v[i]
	}
	return res / float64(n)
}

func variance(v []float64) float64 {
	var res float64 = 0
	var m = mean(v)
	var n int = len(v)
	for i := 0; i < n; i++ {
		res += (v[i] - m) * (v[i] - m)
	}
	if n <= 1 {
		return 0
	}
	return res / float64(n-1)
}
func std(v []float64) float64 {
	if len(v) == 0 {
		return 0
	}
	return math.Sqrt(variance(v))
}
