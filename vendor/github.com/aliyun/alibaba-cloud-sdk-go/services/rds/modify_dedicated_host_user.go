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

// ModifyDedicatedHostUser invokes the rds.ModifyDedicatedHostUser API synchronously
// api document: https://help.aliyun.com/api/rds/modifydedicatedhostuser.html
func (client *Client) ModifyDedicatedHostUser(request *ModifyDedicatedHostUserRequest) (response *ModifyDedicatedHostUserResponse, err error) {
	response = CreateModifyDedicatedHostUserResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyDedicatedHostUserWithChan invokes the rds.ModifyDedicatedHostUser API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydedicatedhostuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedHostUserWithChan(request *ModifyDedicatedHostUserRequest) (<-chan *ModifyDedicatedHostUserResponse, <-chan error) {
	responseChan := make(chan *ModifyDedicatedHostUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyDedicatedHostUser(request)
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

// ModifyDedicatedHostUserWithCallback invokes the rds.ModifyDedicatedHostUser API asynchronously
// api document: https://help.aliyun.com/api/rds/modifydedicatedhostuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyDedicatedHostUserWithCallback(request *ModifyDedicatedHostUserRequest, callback func(response *ModifyDedicatedHostUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyDedicatedHostUserResponse
		var err error
		defer close(result)
		response, err = client.ModifyDedicatedHostUser(request)
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

// ModifyDedicatedHostUserRequest is the request struct for api ModifyDedicatedHostUser
type ModifyDedicatedHostUserRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DedicatedHostName    string           `position:"Query" name:"DedicatedHostName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OldPassword          string           `position:"Query" name:"OldPassword"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	NewPassword          string           `position:"Query" name:"NewPassword"`
	UserName             string           `position:"Query" name:"UserName"`
}

// ModifyDedicatedHostUserResponse is the response struct for api ModifyDedicatedHostUser
type ModifyDedicatedHostUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyDedicatedHostUserRequest creates a request to invoke ModifyDedicatedHostUser API
func CreateModifyDedicatedHostUserRequest() (request *ModifyDedicatedHostUserRequest) {
	request = &ModifyDedicatedHostUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyDedicatedHostUser", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyDedicatedHostUserResponse creates a response to parse from ModifyDedicatedHostUser response
func CreateModifyDedicatedHostUserResponse() (response *ModifyDedicatedHostUserResponse) {
	response = &ModifyDedicatedHostUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
