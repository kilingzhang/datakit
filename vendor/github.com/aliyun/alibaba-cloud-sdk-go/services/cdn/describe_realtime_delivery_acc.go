package cdn

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

// DescribeRealtimeDeliveryAcc invokes the cdn.DescribeRealtimeDeliveryAcc API synchronously
// api document: https://help.aliyun.com/api/cdn/describerealtimedeliveryacc.html
func (client *Client) DescribeRealtimeDeliveryAcc(request *DescribeRealtimeDeliveryAccRequest) (response *DescribeRealtimeDeliveryAccResponse, err error) {
	response = CreateDescribeRealtimeDeliveryAccResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRealtimeDeliveryAccWithChan invokes the cdn.DescribeRealtimeDeliveryAcc API asynchronously
// api document: https://help.aliyun.com/api/cdn/describerealtimedeliveryacc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRealtimeDeliveryAccWithChan(request *DescribeRealtimeDeliveryAccRequest) (<-chan *DescribeRealtimeDeliveryAccResponse, <-chan error) {
	responseChan := make(chan *DescribeRealtimeDeliveryAccResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRealtimeDeliveryAcc(request)
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

// DescribeRealtimeDeliveryAccWithCallback invokes the cdn.DescribeRealtimeDeliveryAcc API asynchronously
// api document: https://help.aliyun.com/api/cdn/describerealtimedeliveryacc.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRealtimeDeliveryAccWithCallback(request *DescribeRealtimeDeliveryAccRequest, callback func(response *DescribeRealtimeDeliveryAccResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRealtimeDeliveryAccResponse
		var err error
		defer close(result)
		response, err = client.DescribeRealtimeDeliveryAcc(request)
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

// DescribeRealtimeDeliveryAccRequest is the request struct for api DescribeRealtimeDeliveryAcc
type DescribeRealtimeDeliveryAccRequest struct {
	*requests.RpcRequest
	Project   string           `position:"Query" name:"Project"`
	StartTime string           `position:"Query" name:"StartTime"`
	EndTime   string           `position:"Query" name:"EndTime"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	Interval  string           `position:"Query" name:"Interval"`
	LogStore  string           `position:"Query" name:"LogStore"`
}

// DescribeRealtimeDeliveryAccResponse is the response struct for api DescribeRealtimeDeliveryAcc
type DescribeRealtimeDeliveryAccResponse struct {
	*responses.BaseResponse
	RequestId               string                  `json:"RequestId" xml:"RequestId"`
	ReatTimeDeliveryAccData ReatTimeDeliveryAccData `json:"ReatTimeDeliveryAccData" xml:"ReatTimeDeliveryAccData"`
}

// CreateDescribeRealtimeDeliveryAccRequest creates a request to invoke DescribeRealtimeDeliveryAcc API
func CreateDescribeRealtimeDeliveryAccRequest() (request *DescribeRealtimeDeliveryAccRequest) {
	request = &DescribeRealtimeDeliveryAccRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2018-05-10", "DescribeRealtimeDeliveryAcc", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeRealtimeDeliveryAccResponse creates a response to parse from DescribeRealtimeDeliveryAcc response
func CreateDescribeRealtimeDeliveryAccResponse() (response *DescribeRealtimeDeliveryAccResponse) {
	response = &DescribeRealtimeDeliveryAccResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
