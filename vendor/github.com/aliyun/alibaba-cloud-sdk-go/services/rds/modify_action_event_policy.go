package rds

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

// ModifyActionEventPolicy invokes the rds.ModifyActionEventPolicy API synchronously
// api document: https://help.aliyun.com/api/rds/modifyactioneventpolicy.html
func (client *Client) ModifyActionEventPolicy(request *ModifyActionEventPolicyRequest) (response *ModifyActionEventPolicyResponse, err error) {
	response = CreateModifyActionEventPolicyResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyActionEventPolicyWithChan invokes the rds.ModifyActionEventPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyactioneventpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyActionEventPolicyWithChan(request *ModifyActionEventPolicyRequest) (<-chan *ModifyActionEventPolicyResponse, <-chan error) {
	responseChan := make(chan *ModifyActionEventPolicyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyActionEventPolicy(request)
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

// ModifyActionEventPolicyWithCallback invokes the rds.ModifyActionEventPolicy API asynchronously
// api document: https://help.aliyun.com/api/rds/modifyactioneventpolicy.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyActionEventPolicyWithCallback(request *ModifyActionEventPolicyRequest, callback func(response *ModifyActionEventPolicyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyActionEventPolicyResponse
		var err error
		defer close(result)
		response, err = client.ModifyActionEventPolicy(request)
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

// ModifyActionEventPolicyRequest is the request struct for api ModifyActionEventPolicy
type ModifyActionEventPolicyRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	EnableEventLog       string           `position:"Query" name:"EnableEventLog"`
}

// ModifyActionEventPolicyResponse is the response struct for api ModifyActionEventPolicy
type ModifyActionEventPolicyResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	RegionId       string `json:"RegionId" xml:"RegionId"`
	EnableEventLog string `json:"EnableEventLog" xml:"EnableEventLog"`
}

// CreateModifyActionEventPolicyRequest creates a request to invoke ModifyActionEventPolicy API
func CreateModifyActionEventPolicyRequest() (request *ModifyActionEventPolicyRequest) {
	request = &ModifyActionEventPolicyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyActionEventPolicy", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyActionEventPolicyResponse creates a response to parse from ModifyActionEventPolicy response
func CreateModifyActionEventPolicyResponse() (response *ModifyActionEventPolicyResponse) {
	response = &ModifyActionEventPolicyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
