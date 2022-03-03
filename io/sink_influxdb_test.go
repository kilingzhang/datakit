package io

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	_ "github.com/influxdata/influxdb1-client" // this is important because of the bug in go mod
	client "github.com/influxdata/influxdb1-client/v2"
)

func ExampleClient_query() {
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://localhost:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
	}
	defer c.Close()

	q := client.NewQuery("SELECT count(value) FROM cpu_load", "mydb", "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		fmt.Println(response.Results)
	}
}

// go test -v -timeout 30s -run ^TestWrite$ gitlab.jiagouyun.com/cloudcare-tools/datakit/io
func TestWrite(t *testing.T) {
	sampleSize := 1000

	// Make client
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://172.16.239.130:8086",
	})
	if err != nil {
		fmt.Println("Error creating InfluxDB Client: ", err.Error())
		return
	}
	defer c.Close()

	rand.Seed(42)

	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "db0",
		Precision: "ns",
	})
	if err != nil {
		fmt.Println("NewBatchPoints: failed: ", err.Error())
		return
	}

	for i := 0; i < sampleSize; i++ {
		regions := []string{"us-west1", "us-west2", "us-west3", "us-east1"}
		tags := map[string]string{
			"cpu":    "cpu-total",
			"host":   fmt.Sprintf("host%d", rand.Intn(1000)),
			"region": regions[rand.Intn(len(regions))],
		}

		idle := rand.Float64() * 100.0
		fields := map[string]interface{}{
			"idle": idle,
			"busy": 100.0 - idle,
		}

		pt, err := client.NewPoint(
			"cpu_usage",
			tags,
			fields,
			time.Now(),
		)
		if err != nil {
			println("Error:", err.Error())
			continue
		}
		bp.AddPoint(pt)
	}

	err = c.Write(bp)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
