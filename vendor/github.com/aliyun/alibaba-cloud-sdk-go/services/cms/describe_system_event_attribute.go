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

// DescribeSystemEventAttribute invokes the cms.DescribeSystemEventAttribute API synchronously
// api document: https://help.aliyun.com/api/cms/describesystemeventattribute.html
func (client *Client) DescribeSystemEventAttribute(request *DescribeSystemEventAttributeRequest) (response *DescribeSystemEventAttributeResponse, err error) {
	response = CreateDescribeSystemEventAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSystemEventAttributeWithChan invokes the cms.DescribeSystemEventAttribute API asynchronously
// api document: https://help.aliyun.com/api/cms/describesystemeventattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSystemEventAttributeWithChan(request *DescribeSystemEventAttributeRequest) (<-chan *DescribeSystemEventAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeSystemEventAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSystemEventAttribute(request)
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

// DescribeSystemEventAttributeWithCallback invokes the cms.DescribeSystemEventAttribute API asynchronously
// api document: https://help.aliyun.com/api/cms/describesystemeventattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSystemEventAttributeWithCallback(request *DescribeSystemEventAttributeRequest, callback func(response *DescribeSystemEventAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSystemEventAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeSystemEventAttribute(request)
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

// DescribeSystemEventAttributeRequest is the request struct for api DescribeSystemEventAttribute
type DescribeSystemEventAttributeRequest struct {
	*requests.RpcRequest
	StartTime      string           `position:"Query" name:"StartTime"`
	SearchKeywords string           `position:"Query" name:"SearchKeywords"`
	PageNumber     requests.Integer `position:"Query" name:"PageNumber"`
	PageSize       requests.Integer `position:"Query" name:"PageSize"`
	Product        string           `position:"Query" name:"Product"`
	Level          string           `position:"Query" name:"Level"`
	GroupId        string           `position:"Query" name:"GroupId"`
	EndTime        string           `position:"Query" name:"EndTime"`
	Name           string           `position:"Query" name:"Name"`
	EventType      string           `position:"Query" name:"EventType"`
	Status         string           `position:"Query" name:"Status"`
}

// DescribeSystemEventAttributeResponse is the response struct for api DescribeSystemEventAttribute
type DescribeSystemEventAttributeResponse struct {
	*responses.BaseResponse
	Code         string       `json:"Code" xml:"Code"`
	Message      string       `json:"Message" xml:"Message"`
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Success      string       `json:"Success" xml:"Success"`
	SystemEvents SystemEvents `json:"SystemEvents" xml:"SystemEvents"`
}

// CreateDescribeSystemEventAttributeRequest creates a request to invoke DescribeSystemEventAttribute API
func CreateDescribeSystemEventAttributeRequest() (request *DescribeSystemEventAttributeRequest) {
	request = &DescribeSystemEventAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeSystemEventAttribute", "cms", "openAPI")
	return
}

// CreateDescribeSystemEventAttributeResponse creates a response to parse from DescribeSystemEventAttribute response
func CreateDescribeSystemEventAttributeResponse() (response *DescribeSystemEventAttributeResponse) {
	response = &DescribeSystemEventAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
