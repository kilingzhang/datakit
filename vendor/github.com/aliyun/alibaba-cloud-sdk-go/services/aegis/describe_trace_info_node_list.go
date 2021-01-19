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

// DescribeTraceInfoNodeList invokes the aegis.DescribeTraceInfoNodeList API synchronously
// api document: https://help.aliyun.com/api/aegis/describetraceinfonodelist.html
func (client *Client) DescribeTraceInfoNodeList(request *DescribeTraceInfoNodeListRequest) (response *DescribeTraceInfoNodeListResponse, err error) {
	response = CreateDescribeTraceInfoNodeListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTraceInfoNodeListWithChan invokes the aegis.DescribeTraceInfoNodeList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describetraceinfonodelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTraceInfoNodeListWithChan(request *DescribeTraceInfoNodeListRequest) (<-chan *DescribeTraceInfoNodeListResponse, <-chan error) {
	responseChan := make(chan *DescribeTraceInfoNodeListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTraceInfoNodeList(request)
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

// DescribeTraceInfoNodeListWithCallback invokes the aegis.DescribeTraceInfoNodeList API asynchronously
// api document: https://help.aliyun.com/api/aegis/describetraceinfonodelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTraceInfoNodeListWithCallback(request *DescribeTraceInfoNodeListRequest, callback func(response *DescribeTraceInfoNodeListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTraceInfoNodeListResponse
		var err error
		defer close(result)
		response, err = client.DescribeTraceInfoNodeList(request)
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

// DescribeTraceInfoNodeListRequest is the request struct for api DescribeTraceInfoNodeList
type DescribeTraceInfoNodeListRequest struct {
	*requests.RpcRequest
	SourceIp  string           `position:"Query" name:"SourceIp"`
	VertexId  string           `position:"Query" name:"VertexId"`
	StartType string           `position:"Query" name:"StartType"`
	PageSize  requests.Integer `position:"Query" name:"PageSize"`
	From      string           `position:"Query" name:"From"`
	Page      requests.Integer `position:"Query" name:"Page"`
	Lang      string           `position:"Query" name:"Lang"`
	Type      string           `position:"Query" name:"Type"`
	Uuid      string           `position:"Query" name:"Uuid"`
}

// DescribeTraceInfoNodeListResponse is the response struct for api DescribeTraceInfoNodeList
type DescribeTraceInfoNodeListResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	NodeListInfo NodeListInfo `json:"NodeListInfo" xml:"NodeListInfo"`
}

// CreateDescribeTraceInfoNodeListRequest creates a request to invoke DescribeTraceInfoNodeList API
func CreateDescribeTraceInfoNodeListRequest() (request *DescribeTraceInfoNodeListRequest) {
	request = &DescribeTraceInfoNodeListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("aegis", "2016-11-11", "DescribeTraceInfoNodeList", "vipaegis", "openAPI")
	return
}

// CreateDescribeTraceInfoNodeListResponse creates a response to parse from DescribeTraceInfoNodeList response
func CreateDescribeTraceInfoNodeListResponse() (response *DescribeTraceInfoNodeListResponse) {
	response = &DescribeTraceInfoNodeListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
