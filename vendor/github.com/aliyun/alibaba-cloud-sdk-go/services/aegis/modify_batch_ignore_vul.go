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

// ModifyBatchIgnoreVul invokes the aegis.ModifyBatchIgnoreVul API synchronously
// api document: https://help.aliyun.com/api/aegis/modifybatchignorevul.html
func (client *Client) ModifyBatchIgnoreVul(request *ModifyBatchIgnoreVulRequest) (response *ModifyBatchIgnoreVulResponse, err error) {
	response = CreateModifyBatchIgnoreVulResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyBatchIgnoreVulWithChan invokes the aegis.ModifyBatchIgnoreVul API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifybatchignorevul.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBatchIgnoreVulWithChan(request *ModifyBatchIgnoreVulRequest) (<-chan *ModifyBatchIgnoreVulResponse, <-chan error) {
	responseChan := make(chan *ModifyBatchIgnoreVulResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyBatchIgnoreVul(request)
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

// ModifyBatchIgnoreVulWithCallback invokes the aegis.ModifyBatchIgnoreVul API asynchronously
// api document: https://help.aliyun.com/api/aegis/modifybatchignorevul.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyBatchIgnoreVulWithCallback(request *ModifyBatchIgnoreVulRequest, callback func(response *ModifyBatchIgnoreVulResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyBatchIgnoreVulResponse
		var err error
		defer close(result)
		response, err = client.ModifyBatchIgnoreVul(request)
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

// ModifyBatchIgnoreVulRequest is the request struct for api ModifyBatchIgnoreVul
type ModifyBatchIgnoreVulRequest struct {
	*requests.RpcRequest
	Reason   string `position:"Query" name:"Reason"`
	SourceIp string `position:"Query" name:"SourceIp"`
	Info     string `position:"Query" name:"Info"`
}

// ModifyBatchIgnoreVulResponse is the response struct for api ModifyBatchIgnoreVul
type ModifyBatchIgnoreVulResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyBatchIgnoreVulRequest creates a request to invoke ModifyBatchIgnoreVul API
func CreateModifyBatchIgnoreVulRequest() (request *ModifyBatchIgnoreVulRequest) {
	request = &ModifyBatchIgnoreVulRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "ModifyBatchIgnoreVul", "vipaegis", "openAPI")
	return
}

// CreateModifyBatchIgnoreVulResponse creates a response to parse from ModifyBatchIgnoreVul response
func CreateModifyBatchIgnoreVulResponse() (response *ModifyBatchIgnoreVulResponse) {
	response = &ModifyBatchIgnoreVulResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
