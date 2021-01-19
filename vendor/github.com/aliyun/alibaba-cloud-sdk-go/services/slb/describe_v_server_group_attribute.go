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

// DescribeVServerGroupAttribute invokes the slb.DescribeVServerGroupAttribute API synchronously
// api document: https://help.aliyun.com/api/slb/describevservergroupattribute.html
func (client *Client) DescribeVServerGroupAttribute(request *DescribeVServerGroupAttributeRequest) (response *DescribeVServerGroupAttributeResponse, err error) {
	response = CreateDescribeVServerGroupAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVServerGroupAttributeWithChan invokes the slb.DescribeVServerGroupAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describevservergroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVServerGroupAttributeWithChan(request *DescribeVServerGroupAttributeRequest) (<-chan *DescribeVServerGroupAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeVServerGroupAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVServerGroupAttribute(request)
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

// DescribeVServerGroupAttributeWithCallback invokes the slb.DescribeVServerGroupAttribute API asynchronously
// api document: https://help.aliyun.com/api/slb/describevservergroupattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVServerGroupAttributeWithCallback(request *DescribeVServerGroupAttributeRequest, callback func(response *DescribeVServerGroupAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVServerGroupAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeVServerGroupAttribute(request)
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

// DescribeVServerGroupAttributeRequest is the request struct for api DescribeVServerGroupAttribute
type DescribeVServerGroupAttributeRequest struct {
	*requests.RpcRequest
	AccessKeyId          string           `position:"Query" name:"access_key_id"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VServerGroupId       string           `position:"Query" name:"VServerGroupId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	Tags                 string           `position:"Query" name:"Tags"`
}

// DescribeVServerGroupAttributeResponse is the response struct for api DescribeVServerGroupAttribute
type DescribeVServerGroupAttributeResponse struct {
	*responses.BaseResponse
	RequestId        string                                        `json:"RequestId" xml:"RequestId"`
	VServerGroupId   string                                        `json:"VServerGroupId" xml:"VServerGroupId"`
	VServerGroupName string                                        `json:"VServerGroupName" xml:"VServerGroupName"`
	LoadBalancerId   string                                        `json:"LoadBalancerId" xml:"LoadBalancerId"`
	BackendServers   BackendServersInDescribeVServerGroupAttribute `json:"BackendServers" xml:"BackendServers"`
}

// CreateDescribeVServerGroupAttributeRequest creates a request to invoke DescribeVServerGroupAttribute API
func CreateDescribeVServerGroupAttributeRequest() (request *DescribeVServerGroupAttributeRequest) {
	request = &DescribeVServerGroupAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2014-05-15", "DescribeVServerGroupAttribute", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVServerGroupAttributeResponse creates a response to parse from DescribeVServerGroupAttribute response
func CreateDescribeVServerGroupAttributeResponse() (response *DescribeVServerGroupAttributeResponse) {
	response = &DescribeVServerGroupAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
