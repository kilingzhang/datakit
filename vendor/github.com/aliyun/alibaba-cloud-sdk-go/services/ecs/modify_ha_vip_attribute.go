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

// ModifyHaVipAttribute invokes the ecs.ModifyHaVipAttribute API synchronously
// api document: https://help.aliyun.com/api/ecs/modifyhavipattribute.html
func (client *Client) ModifyHaVipAttribute(request *ModifyHaVipAttributeRequest) (response *ModifyHaVipAttributeResponse, err error) {
	response = CreateModifyHaVipAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyHaVipAttributeWithChan invokes the ecs.ModifyHaVipAttribute API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyhavipattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHaVipAttributeWithChan(request *ModifyHaVipAttributeRequest) (<-chan *ModifyHaVipAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyHaVipAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyHaVipAttribute(request)
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

// ModifyHaVipAttributeWithCallback invokes the ecs.ModifyHaVipAttribute API asynchronously
// api document: https://help.aliyun.com/api/ecs/modifyhavipattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyHaVipAttributeWithCallback(request *ModifyHaVipAttributeRequest, callback func(response *ModifyHaVipAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyHaVipAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyHaVipAttribute(request)
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

// ModifyHaVipAttributeRequest is the request struct for api ModifyHaVipAttribute
type ModifyHaVipAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	Description          string           `position:"Query" name:"Description"`
	HaVipId              string           `position:"Query" name:"HaVipId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// ModifyHaVipAttributeResponse is the response struct for api ModifyHaVipAttribute
type ModifyHaVipAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyHaVipAttributeRequest creates a request to invoke ModifyHaVipAttribute API
func CreateModifyHaVipAttributeRequest() (request *ModifyHaVipAttributeRequest) {
	request = &ModifyHaVipAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "ModifyHaVipAttribute", "ecs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateModifyHaVipAttributeResponse creates a response to parse from ModifyHaVipAttribute response
func CreateModifyHaVipAttributeResponse() (response *ModifyHaVipAttributeResponse) {
	response = &ModifyHaVipAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
