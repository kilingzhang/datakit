package sas

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

// FixCheckWarnings invokes the sas.FixCheckWarnings API synchronously
// api document: https://help.aliyun.com/api/sas/fixcheckwarnings.html
func (client *Client) FixCheckWarnings(request *FixCheckWarningsRequest) (response *FixCheckWarningsResponse, err error) {
	response = CreateFixCheckWarningsResponse()
	err = client.DoAction(request, response)
	return
}

// FixCheckWarningsWithChan invokes the sas.FixCheckWarnings API asynchronously
// api document: https://help.aliyun.com/api/sas/fixcheckwarnings.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FixCheckWarningsWithChan(request *FixCheckWarningsRequest) (<-chan *FixCheckWarningsResponse, <-chan error) {
	responseChan := make(chan *FixCheckWarningsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.FixCheckWarnings(request)
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

// FixCheckWarningsWithCallback invokes the sas.FixCheckWarnings API asynchronously
// api document: https://help.aliyun.com/api/sas/fixcheckwarnings.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) FixCheckWarningsWithCallback(request *FixCheckWarningsRequest, callback func(response *FixCheckWarningsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *FixCheckWarningsResponse
		var err error
		defer close(result)
		response, err = client.FixCheckWarnings(request)
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

// FixCheckWarningsRequest is the request struct for api FixCheckWarnings
type FixCheckWarningsRequest struct {
	*requests.RpcRequest
	RiskId      requests.Integer `position:"Query" name:"RiskId"`
	CheckParams string           `position:"Query" name:"CheckParams"`
	SourceIp    string           `position:"Query" name:"SourceIp"`
	Lang        string           `position:"Query" name:"Lang"`
	Uuids       string           `position:"Query" name:"Uuids"`
}

// FixCheckWarningsResponse is the response struct for api FixCheckWarnings
type FixCheckWarningsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	BatchId   int64  `json:"BatchId" xml:"BatchId"`
}

// CreateFixCheckWarningsRequest creates a request to invoke FixCheckWarnings API
func CreateFixCheckWarningsRequest() (request *FixCheckWarningsRequest) {
	request = &FixCheckWarningsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Sas", "2018-12-03", "FixCheckWarnings", "sas", "openAPI")
	return
}

// CreateFixCheckWarningsResponse creates a response to parse from FixCheckWarnings response
func CreateFixCheckWarningsResponse() (response *FixCheckWarningsResponse) {
	response = &FixCheckWarningsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
