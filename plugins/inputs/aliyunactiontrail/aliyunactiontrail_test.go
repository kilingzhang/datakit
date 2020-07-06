package aliyunactiontrail

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/actiontrail"
	"github.com/influxdata/toml"
)

func TestConfig(t *testing.T) {

	var cfg AliyunActiontrail
	// cfg.Actiontrail = []*ActiontrailInstance{
	// 	&ActiontrailInstance{
	// 		Region:     "",
	// 		AccessID:   "",
	// 		AccessKey:  "",
	// 		MetricName: "aliac",
	// 	},
	// 	&ActiontrailInstance{
	// 		Region:     "11",
	// 		AccessID:   "22",
	// 		AccessKey:  "",
	// 		MetricName: "aliac2",
	// 	},
	// }
	if data, err := toml.Marshal(&cfg); err != nil {
		t.Errorf("%s", err)
	} else {
		log.Printf("%s", string(data))
	}
}

func TestActiontrail(t *testing.T) {
	cli, err := actiontrail.NewClientWithAccessKey(`cn-hangzhou`, `LTAI4FkR2SokHHESouUMrkxV`, `ht4jybX3IrhQAUgHrUOTJRrkP8dONJ`)
	if err != nil {
		t.Errorf("%s", err)
	}

	//startTm := time.Now().Truncate(time.Hour).Add(-time.Hour * 22)

	request := actiontrail.CreateLookupEventsRequest()
	request.Scheme = "https"
	//request.StartTime = unixTimeStrISO8601(startTm)
	//request.EndTime = unixTimeStrISO8601(startTm.Add(time.Minute * 30))

	//log.Printf("range: %s - %s", request.StartTime, request.EndTime)

	response, err := cli.LookupEvents(request)
	if err != nil {
		t.Errorf("LookupEvents failed, %s", err)
	}

	fmt.Printf("%s\n", response.String())

	for _, ev := range response.Events {
		tags := map[string]string{}
		fields := map[string]interface{}{}

		if eventType, ok := ev["eventType"].(string); ok {
			tags["eventType"] = eventType
		}

		if acsRegion, ok := ev["acsRegion"].(string); ok {
			tags["region"] = acsRegion
		}

		fields["eventId"] = ev["eventId"]
		fields["eventSource"] = ev["eventSource"]
		fields["serviceName"] = ev["serviceName"]
		if ev["sourceIpAddress"] != nil {
			fields["sourceIpAddress"] = ev["sourceIpAddress"]
		}
		fields["userAgent"] = ev["userAgent"]
		fields["eventVersion"] = ev["eventVersion"]

		if userIdentity, ok := ev["userIdentity"].(map[string]interface{}); ok {
			//userIdentity:map[accountId:50220571 principalId:50220571 type:root-account userName:root]
			fields["accountId"] = userIdentity["accountId"]
			fields["accountType"] = userIdentity["type"]
			fields["userName"] = userIdentity["userName"]
			fields["principalId"] = userIdentity["principalId"]
		}

		if additionalEventData, ok := ev["additionalEventData"].(map[string]interface{}); ok {
			//additionalEventData:map[isMFAChecked:false loginAccount:13626736491]
			fields["loginAccount"] = additionalEventData["loginAccount"]
			fields["isMFAChecked"] = additionalEventData["isMFAChecked"]
		}

		eventTime := ev["eventTime"].(string) //utc
		evtm, err := time.Parse(`2006-01-02T15:04:05Z`, eventTime)
		if err != nil {
			t.Errorf("%s", err)
		}

		fmt.Printf("%s, %s\n", ev["referencedResources"], evtm)

	}
}
