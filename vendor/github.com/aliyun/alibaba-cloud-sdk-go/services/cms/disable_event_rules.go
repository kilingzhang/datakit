package cms

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

// DisableEventRules invokes the cms.DisableEventRules API synchronously
// api document: https://help.aliyun.com/api/cms/disableeventrules.html
func (client *Client) DisableEventRules(request *DisableEventRulesRequest) (response *DisableEventRulesResponse, err error) {
	response = CreateDisableEventRulesResponse()
	err = client.DoAction(request, response)
	return
}

// DisableEventRulesWithChan invokes the cms.DisableEventRules API asynchronously
// api document: https://help.aliyun.com/api/cms/disableeventrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableEventRulesWithChan(request *DisableEventRulesRequest) (<-chan *DisableEventRulesResponse, <-chan error) {
	responseChan := make(chan *DisableEventRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DisableEventRules(request)
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

// DisableEventRulesWithCallback invokes the cms.DisableEventRules API asynchronously
// api document: https://help.aliyun.com/api/cms/disableeventrules.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DisableEventRulesWithCallback(request *DisableEventRulesRequest, callback func(response *DisableEventRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DisableEventRulesResponse
		var err error
		defer close(result)
		response, err = client.DisableEventRules(request)
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

// DisableEventRulesRequest is the request struct for api DisableEventRules
type DisableEventRulesRequest struct {
	*requests.RpcRequest
	RuleNames *[]string `position:"Query" name:"RuleNames"  type:"Repeated"`
}

// DisableEventRulesResponse is the response struct for api DisableEventRules
type DisableEventRulesResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDisableEventRulesRequest creates a request to invoke DisableEventRules API
func CreateDisableEventRulesRequest() (request *DisableEventRulesRequest) {
	request = &DisableEventRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DisableEventRules", "cms", "openAPI")
	return
}

// CreateDisableEventRulesResponse creates a response to parse from DisableEventRules response
func CreateDisableEventRulesResponse() (response *DisableEventRulesResponse) {
	response = &DisableEventRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
