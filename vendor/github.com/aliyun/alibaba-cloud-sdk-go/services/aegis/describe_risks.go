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

// DescribeRisks invokes the aegis.DescribeRisks API synchronously
// api document: https://help.aliyun.com/api/aegis/describerisks.html
func (client *Client) DescribeRisks(request *DescribeRisksRequest) (response *DescribeRisksResponse, err error) {
	response = CreateDescribeRisksResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeRisksWithChan invokes the aegis.DescribeRisks API asynchronously
// api document: https://help.aliyun.com/api/aegis/describerisks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRisksWithChan(request *DescribeRisksRequest) (<-chan *DescribeRisksResponse, <-chan error) {
	responseChan := make(chan *DescribeRisksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRisks(request)
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

// DescribeRisksWithCallback invokes the aegis.DescribeRisks API asynchronously
// api document: https://help.aliyun.com/api/aegis/describerisks.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeRisksWithCallback(request *DescribeRisksRequest, callback func(response *DescribeRisksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRisksResponse
		var err error
		defer close(result)
		response, err = client.DescribeRisks(request)
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

// DescribeRisksRequest is the request struct for api DescribeRisks
type DescribeRisksRequest struct {
	*requests.RpcRequest
	RiskName string           `position:"Query" name:"RiskName"`
	SourceIp string           `position:"Query" name:"SourceIp"`
	Limit    requests.Integer `position:"Query" name:"Limit"`
	Lang     string           `position:"Query" name:"Lang"`
	RiskId   requests.Integer `position:"Query" name:"RiskId"`
}

// DescribeRisksResponse is the response struct for api DescribeRisks
type DescribeRisksResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	Risks      []Risk `json:"Risks" xml:"Risks"`
}

// CreateDescribeRisksRequest creates a request to invoke DescribeRisks API
func CreateDescribeRisksRequest() (request *DescribeRisksRequest) {
	request = &DescribeRisksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeRisks", "vipaegis", "openAPI")
	return
}

// CreateDescribeRisksResponse creates a response to parse from DescribeRisks response
func CreateDescribeRisksResponse() (response *DescribeRisksResponse) {
	response = &DescribeRisksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
