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

// DescribeRecommendInstanceType invokes the ecs.DescribeRecommendInstanceType API synchronously
// api document: https://help.aliyun.com/api/ecs/describerecommendinstancetype.html
func (client *Client) DescribeRecommendInstanceType(request *DescribeRecommendInstanceTypeRequest) (response *DescribeRecommendInstanceTypeResponse, err error) {
	response = CreateDescribeRecommendInstanceTypeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRecommendInstanceTypeWithChan invokes the ecs.DescribeRecommendInstanceType API asynchronously
// api document: https://help.aliyun.com/api/ecs/describerecommendinstancetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecommendInstanceTypeWithChan(request *DescribeRecommendInstanceTypeRequest) (<-chan *DescribeRecommendInstanceTypeResponse, <-chan error) {
	responseChan := make(chan *DescribeRecommendInstanceTypeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRecommendInstanceType(request)
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

// DescribeRecommendInstanceTypeWithCallback invokes the ecs.DescribeRecommendInstanceType API asynchronously
// api document: https://help.aliyun.com/api/ecs/describerecommendinstancetype.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRecommendInstanceTypeWithCallback(request *DescribeRecommendInstanceTypeRequest, callback func(response *DescribeRecommendInstanceTypeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRecommendInstanceTypeResponse
		var err error
		defer close(result)
		response, err = client.DescribeRecommendInstanceType(request)
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

// DescribeRecommendInstanceTypeRequest is the request struct for api DescribeRecommendInstanceType
type DescribeRecommendInstanceTypeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstancePpsRx        requests.Integer `position:"Query" name:"InstancePpsRx"`
	Memory               requests.Float   `position:"Query" name:"Memory"`
	InstancePpsTx        requests.Integer `position:"Query" name:"InstancePpsTx"`
	IoOptimized          string           `position:"Query" name:"IoOptimized"`
	NetworkType          string           `position:"Query" name:"NetworkType"`
	Scene                string           `position:"Query" name:"Scene"`
	InstanceBandwidthTx  requests.Integer `position:"Query" name:"InstanceBandwidthTx"`
	Cores                requests.Integer `position:"Query" name:"Cores"`
	InstanceBandwidthRx  requests.Integer `position:"Query" name:"InstanceBandwidthRx"`
	SystemDiskCategory   string           `position:"Query" name:"SystemDiskCategory"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
	InstanceChargeType   string           `position:"Query" name:"InstanceChargeType"`
	MaxPrice             requests.Float   `position:"Query" name:"MaxPrice"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	InstanceTypeFamily   *[]string        `position:"Query" name:"InstanceTypeFamily"  type:"Repeated"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SpotStrategy         string           `position:"Query" name:"SpotStrategy"`
	PriorityStrategy     string           `position:"Query" name:"PriorityStrategy"`
	InstanceFamilyLevel  string           `position:"Query" name:"InstanceFamilyLevel"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
}

// DescribeRecommendInstanceTypeResponse is the response struct for api DescribeRecommendInstanceType
type DescribeRecommendInstanceTypeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateDescribeRecommendInstanceTypeRequest creates a request to invoke DescribeRecommendInstanceType API
func CreateDescribeRecommendInstanceTypeRequest() (request *DescribeRecommendInstanceTypeRequest) {
	request = &DescribeRecommendInstanceTypeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeRecommendInstanceType", "ecs", "openAPI")
	return
}

// CreateDescribeRecommendInstanceTypeResponse creates a response to parse from DescribeRecommendInstanceType response
func CreateDescribeRecommendInstanceTypeResponse() (response *DescribeRecommendInstanceTypeResponse) {
	response = &DescribeRecommendInstanceTypeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
