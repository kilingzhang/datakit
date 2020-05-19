package aegis

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// QueryCrackEvent invokes the aegis.QueryCrackEvent API synchronously
// api document: https://help.aliyun.com/api/aegis/querycrackevent.html
func (client *Client) QueryCrackEvent(request *QueryCrackEventRequest) (response *QueryCrackEventResponse, err error) {
	response = CreateQueryCrackEventResponse()
	err = client.DoAction(request, response)
	return
}

// QueryCrackEventWithChan invokes the aegis.QueryCrackEvent API asynchronously
// api document: https://help.aliyun.com/api/aegis/querycrackevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCrackEventWithChan(request *QueryCrackEventRequest) (<-chan *QueryCrackEventResponse, <-chan error) {
	responseChan := make(chan *QueryCrackEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryCrackEvent(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// QueryCrackEventWithCallback invokes the aegis.QueryCrackEvent API asynchronously
// api document: https://help.aliyun.com/api/aegis/querycrackevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryCrackEventWithCallback(request *QueryCrackEventRequest, callback func(response *QueryCrackEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryCrackEventResponse
		var err error
		defer close(result)
		response, err = client.QueryCrackEvent(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// QueryCrackEventRequest is the request struct for api QueryCrackEvent
type QueryCrackEventRequest struct {
	*requests.RpcRequest
	EndTime     string           `position:"Query" name:"EndTime"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	StartTime   string           `position:"Query" name:"StartTime"`
	Uuid        string           `position:"Query" name:"Uuid"`
	Status      requests.Integer `position:"Query" name:"Status"`
}

// QueryCrackEventResponse is the response struct for api QueryCrackEvent
type QueryCrackEventResponse struct {
	*responses.BaseResponse
	RequestId string `json:"requestId" xml:"requestId"`
	Code      string `json:"Code" xml:"Code"`
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryCrackEventRequest creates a request to invoke QueryCrackEvent API
func CreateQueryCrackEventRequest() (request *QueryCrackEventRequest) {
	request = &QueryCrackEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "QueryCrackEvent", "vipaegis", "openAPI")
	return
}

// CreateQueryCrackEventResponse creates a response to parse from QueryCrackEvent response
func CreateQueryCrackEventResponse() (response *QueryCrackEventResponse) {
	response = &QueryCrackEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
