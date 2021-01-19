package slb

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

// DeleteMasterSlaveServerGroup invokes the slb.DeleteMasterSlaveServerGroup API synchronously
// api document: https://help.aliyun.com/api/slb/deletemasterslaveservergroup.html
func (client *Client) DeleteMasterSlaveServerGroup(request *DeleteMasterSlaveServerGroupRequest) (response *DeleteMasterSlaveServerGroupResponse, err error) {
	response = CreateDeleteMasterSlaveServerGroupResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteMasterSlaveServerGroupWithChan invokes the slb.DeleteMasterSlaveServerGroup API asynchronously
// api document: https://help.aliyun.com/api/slb/deletemasterslaveservergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMasterSlaveServerGroupWithChan(request *DeleteMasterSlaveServerGroupRequest) (<-chan *DeleteMasterSlaveServerGroupResponse, <-chan error) {
	responseChan := make(chan *DeleteMasterSlaveServerGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMasterSlaveServerGroup(request)
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

// DeleteMasterSlaveServerGroupWithCallback invokes the slb.DeleteMasterSlaveServerGroup API asynchronously
// api document: https://help.aliyun.com/api/slb/deletemasterslaveservergroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMasterSlaveServerGroupWithCallback(request *DeleteMasterSlaveServerGroupRequest, callback func(response *DeleteMasterSlaveServerGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMasterSlaveServerGroupResponse
		var err error
		defer close(result)
		response, err = client.DeleteMasterSlaveServerGroup(request)
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

// DeleteMasterSlaveServerGroupRequest is the request struct for api DeleteMasterSlaveServerGroup
type DeleteMasterSlaveServerGroupRequest struct {
	*requests.RpcRequest
	AccessKeyId              string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId          requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount     string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount             string           `position:"Query" name:"OwnerAccount"`
	OwnerId                  requests.Integer `position:"Query" name:"OwnerId"`
	Tags                     string           `position:"Query" name:"Tags"`
	MasterSlaveServerGroupId string           `position:"Query" name:"MasterSlaveServerGroupId"`
}

// DeleteMasterSlaveServerGroupResponse is the response struct for api DeleteMasterSlaveServerGroup
type DeleteMasterSlaveServerGroupResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteMasterSlaveServerGroupRequest creates a request to invoke DeleteMasterSlaveServerGroup API
func CreateDeleteMasterSlaveServerGroupRequest() (request *DeleteMasterSlaveServerGroupRequest) {
	request = &DeleteMasterSlaveServerGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DeleteMasterSlaveServerGroup", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteMasterSlaveServerGroupResponse creates a response to parse from DeleteMasterSlaveServerGroup response
func CreateDeleteMasterSlaveServerGroupResponse() (response *DeleteMasterSlaveServerGroupResponse) {
	response = &DeleteMasterSlaveServerGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
