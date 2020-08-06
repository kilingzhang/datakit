package ecs

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

// AcceptInquiredSystemEvent invokes the ecs.AcceptInquiredSystemEvent API synchronously
// api document: https://help.aliyun.com/api/ecs/acceptinquiredsystemevent.html
func (client *Client) AcceptInquiredSystemEvent(request *AcceptInquiredSystemEventRequest) (response *AcceptInquiredSystemEventResponse, err error) {
	response = CreateAcceptInquiredSystemEventResponse()
	err = client.DoAction(request, response)
	return
}

// AcceptInquiredSystemEventWithChan invokes the ecs.AcceptInquiredSystemEvent API asynchronously
// api document: https://help.aliyun.com/api/ecs/acceptinquiredsystemevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AcceptInquiredSystemEventWithChan(request *AcceptInquiredSystemEventRequest) (<-chan *AcceptInquiredSystemEventResponse, <-chan error) {
	responseChan := make(chan *AcceptInquiredSystemEventResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AcceptInquiredSystemEvent(request)
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

// AcceptInquiredSystemEventWithCallback invokes the ecs.AcceptInquiredSystemEvent API asynchronously
// api document: https://help.aliyun.com/api/ecs/acceptinquiredsystemevent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AcceptInquiredSystemEventWithCallback(request *AcceptInquiredSystemEventRequest, callback func(response *AcceptInquiredSystemEventResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AcceptInquiredSystemEventResponse
		var err error
		defer close(result)
		response, err = client.AcceptInquiredSystemEvent(request)
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

// AcceptInquiredSystemEventRequest is the request struct for api AcceptInquiredSystemEvent
type AcceptInquiredSystemEventRequest struct {
	*requests.RpcRequest
	EventId              string           `position:"Query" name:"EventId"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// AcceptInquiredSystemEventResponse is the response struct for api AcceptInquiredSystemEvent
type AcceptInquiredSystemEventResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAcceptInquiredSystemEventRequest creates a request to invoke AcceptInquiredSystemEvent API
func CreateAcceptInquiredSystemEventRequest() (request *AcceptInquiredSystemEventRequest) {
	request = &AcceptInquiredSystemEventRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AcceptInquiredSystemEvent", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAcceptInquiredSystemEventResponse creates a response to parse from AcceptInquiredSystemEvent response
func CreateAcceptInquiredSystemEventResponse() (response *AcceptInquiredSystemEventResponse) {
	response = &AcceptInquiredSystemEventResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
