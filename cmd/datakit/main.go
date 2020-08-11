package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	iowrite "io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kardianos/service"

	"gitlab.jiagouyun.com/cloudcare-tools/cliutils/logger"
	uhttp "gitlab.jiagouyun.com/cloudcare-tools/cliutils/network/http"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/config"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/git"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/io"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/all"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/druid"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/inputs/trace"
	_ "gitlab.jiagouyun.com/cloudcare-tools/datakit/plugins/outputs/all"
	"gitlab.jiagouyun.com/cloudcare-tools/datakit/telegrafwrap"
)

var (
	flagVersion        = flag.Bool("version", false, `show verison info`)
	flagDataWay        = flag.String("dataway", ``, `dataway IP:Port`)
	flagCheckConfigDir = flag.Bool("check-config-dir", false, `check datakit conf.d, list configired and mis-configured collectors`)
	flagInputFilters   = flag.String("input-filter", "", "filter the inputs to enable, separator is :")

	flagListCollectors    = flag.Bool("tree", false, `list vailable collectors`)
	flagListConfigSamples = flag.Bool("config-samples", false, `list all config samples`)
)

var (
	stopCh     chan struct{} = make(chan struct{})
	waitExitCh chan struct{} = make(chan struct{})

	inputFilters = []string{}
	l            *logger.Logger
)

func main() {

	logger.SetStdoutRootLogger(logger.DEBUG, logger.OPT_DEFAULT)
	l = logger.SLogger("main")

	flag.Parse()

	applyFlags()

	loadConfig()

	svcConfig := &service.Config{
		Name: datakit.ServiceName,
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		l.Fatal(err)
		return
	}

	l.Info("starting datakit service")

	if err = s.Run(); err != nil {
		l.Fatal(err)
	}
}

func applyFlags() {

	if *flagVersion {
		fmt.Printf(`
       Version: %s
        Commit: %s
        Branch: %s
 Build At(UTC): %s
Golang Version: %s
      Uploader: %s
`, git.Version, git.Commit, git.Branch, git.BuildAt, git.Golang, git.Uploader)
		os.Exit(0)
	}

	if *flagListCollectors {
		showAllCollectors()
		os.Exit(0)
	}

	if *flagListConfigSamples {
		showAllConfigSamples()
		os.Exit(0)
	}

	if *flagCheckConfigDir {
		config.CheckConfd()
		os.Exit(0)
	}

	if *flagInputFilters != "" {
		inputFilters = strings.Split(":"+strings.TrimSpace(*flagInputFilters)+":", ":")
	}
}

func showAllCollectors() {
	collectors := map[string][]string{}

	for k, v := range inputs.Inputs {
		cat := v().Catalog()
		collectors[cat] = append(collectors[cat], k)
	}

	ndatakit := 0
	for k, vs := range collectors {
		//fmt.Println(k)
		for _, v := range vs {
			//fmt.Printf("  |--[d][% 12s] %s\n", k, v)
			fmt.Printf("[d][% 12s] %s\n", k, v)
			ndatakit++
		}
	}

	nagent := 0
	collectors = map[string][]string{}
	for k, v := range config.TelegrafInputs {
		collectors[v.Catalog] = append(collectors[v.Catalog], k)
	}

	for k, vs := range collectors {
		//fmt.Println(k)
		for _, v := range vs {
			//fmt.Printf("  |--[t] %s\n", v)
			fmt.Printf("[t][% 12s] %s\n", k, v)
			nagent++
		}
	}

	fmt.Println("===================================")
	fmt.Printf("total: %d, datakit: %d, agent: %d\n", ndatakit+nagent, ndatakit, nagent)
}

func showAllConfigSamples() {
	for k, v := range inputs.Inputs {
		sample := v().SampleConfig()
		fmt.Printf("%s\n========= [D] ==========\n%s\n", k, sample)
	}

	for k, v := range config.TelegrafInputs {
		fmt.Printf("%s\n========= [T] ==========\n%s\n", k, v.Sample)
	}
}

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run(s)
	return nil
}

func (p *program) run(s service.Service) {
	__run()
}

func (p *program) Stop(s service.Service) error {
	close(stopCh)

	// We must wait here:
	// On windows, we stop datakit in services.msc, if datakit process do not
	// echo to here, services.msc will complain the datakit process has been
	// exit unexpected
	<-waitExitCh

	return nil
}

func exitDatakit() {
	datakit.Exit.Close()

	l.Info("wait all goroutines exit...")
	datakit.WG.Wait()

	l.Info("closing waitExitCh...")
	close(waitExitCh)
}

func __run() {

	datakit.WG.Add(1)
	go func() {
		defer datakit.WG.Done()
		if err := runTelegraf(); err != nil {
			l.Fatalf("fail to start sub service: %s", err)
		}

		l.Info("telegraf process exit ok")
	}()

	l.Info("datakit start...")
	if err := runDatakit(); err != nil && err != context.Canceled {
		l.Fatalf("datakit abort: %s", err)
	}

	l.Info("datakit start ok. Wait signal or service stop...")

	// NOTE:
	// Actually, the datakit process been managed by system service, no matter on
	// windows/UNIX, datakit should exit via `service-stop' operation, so the signal
	// branch should not reached, but for daily debugging(ctrl-c), we kept the signal
	// exit option.
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-signals:
		if sig == syscall.SIGHUP {
			// TODO: reload configures
		} else {
			l.Infof("get signal %v, wait & exit", sig)
			exitDatakit()
		}
	case <-stopCh:
		l.Infof("service stopping")
		exitDatakit()
	}

	l.Info("datakit exit.")
}

func loadConfig() {
	config.Cfg.InputFilters = inputFilters

	for {
		if err := config.LoadCfg(); err != nil {
			l.Errorf("load config failed: %s", err)
			time.Sleep(time.Second)
		} else {
			break
		}
	}

	l = logger.SLogger("main")
}

func runTelegraf() error {
	telegrafwrap.Svr.Start()
	return nil
}

func runDatakit() error {

	l = logger.SLogger("datakit")
	io.Start()

	// start HTTP server
	datakit.WG.Add(1)
	go func() {
		defer datakit.WG.Done()
		httpStart(config.Cfg.MainCfg.HTTPServerAddr)
		l.Info("HTTPServer goroutine exit")
	}()

	if err := runInputs(); err != nil {
		l.Error("error running inputs: %v", err)
	}

	return nil
}

func runInputs() error {

	for name, ips := range config.Cfg.Inputs {
		for _, input := range ips {
			switch input.(type) {
			case inputs.Input:
				l.Infof("starting input %s ...", name)
				datakit.WG.Add(1)
				go func(i inputs.Input, name string) {
					defer datakit.WG.Done()
					i.Run()
					l.Infof("input %s exited", name)
				}(input, name)
			default:
				l.Warn("ignore input %s", name)
			}
		}
	}
	return nil
}

func httpStart(addr string) {
	router := gin.New()
	gin.DisableConsoleColor()

	l.Info("set gin log to %s", config.Cfg.MainCfg.GinLog)
	f, err := os.Create(config.Cfg.MainCfg.GinLog)
	if err != nil {
		l.Fatalf("create gin log failed: %s", err)
	}

	gin.DefaultWriter = iowrite.MultiWriter(f)
	if config.Cfg.MainCfg.LogLevel != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(uhttp.CORSMiddleware)

	// TODO: need any method?
	// router.Any()

	if _, ok := config.Cfg.Inputs["trace"]; ok {
		l.Info("open route for trace")
		router.POST("/trac", func(c *gin.Context) { trace.Handle(c.Writer, c.Request) })
		router.POST("/v3/segment", func(c *gin.Context) { trace.Handle(c.Writer, c.Request) })
		router.POST("/v3/segments", func(c *gin.Context) { trace.Handle(c.Writer, c.Request) })
		router.POST("/v3/management/reportProperties", func(c *gin.Context) { trace.Handle(c.Writer, c.Request) })
		router.POST("/v3/management/keepAlive", func(c *gin.Context) { trace.Handle(c.Writer, c.Request) })
	}

	if _, ok := config.Cfg.Inputs["druid"]; ok {
		l.Info("open route for druid")
		router.POST("/druid", func(c *gin.Context) { druid.Handle(c.Writer, c.Request) })
	}

	// internal datakit stats API
	router.POST("/stats", func(c *gin.Context) { getInputsStats(c.Writer, c.Request) })
	// ansible api
	router.POST("/ansible", func(c *gin.Context) { AnsibleHander(c.Writer, c.Request) })

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			l.Error(err)
		}
		l.Info("http server exit")
	}()

	<-datakit.Exit.Wait()
	l.Info("stopping http server...")

	if err := srv.Shutdown(context.Background()); err != nil {
		l.Errorf("Failed of http server shutdown, err: %s", err.Error())

	} else {
		l.Info("http server shutdown ok")
	}

	return
}

func AnsibleHander(w http.ResponseWriter, r *http.Request) {
	dataType := r.URL.Query().Get("type")
	body, err := ioutil.ReadAll(r.Body)
	l.Infof("ansible body {}", string(body))
	defer r.Body.Close()

	if err != nil {
		l.Errorf("failed of http parsen body in ansible err:%s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch dataType {
	case "keyevent":
		if err := io.NamedFeed(body, io.KeyEvent, "ansible"); err != nil {
			l.Errorf("failed to io Feed, err: %s", err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)

	case "metric":
		if err := io.NamedFeed(body, io.Metric, "ansible"); err != nil {
			l.Errorf("failed to io Feed, err: %s", err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)

	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

}

func getInputsStats(w http.ResponseWriter, r *http.Request) {
	res, err := io.GetStats("") // get all inputs stats
	if err != nil {
		l.Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	body, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
	return
}

func requestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		path := r.URL.Path
		raw := r.URL.RawQuery
		clientIP := r.RemoteAddr

		targetMux.ServeHTTP(w, r)

		if raw != "" {
			path = path + "?" + raw
		}

		l.Infof(" %15s | %#v\n", clientIP, path)
	})
}
