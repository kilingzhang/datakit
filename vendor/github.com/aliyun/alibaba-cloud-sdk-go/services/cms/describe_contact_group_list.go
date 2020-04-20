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

// DescribeContactGroupList invokes the cms.DescribeContactGroupList API synchronously
// api document: https://help.aliyun.com/api/cms/describecontactgrouplist.html
func (client *Client) DescribeContactGroupList(request *DescribeContactGroupListRequest) (response *DescribeContactGroupListResponse, err error) {
	response = CreateDescribeContactGroupListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeContactGroupListWithChan invokes the cms.DescribeContactGroupList API asynchronously
// api document: https://help.aliyun.com/api/cms/describecontactgrouplist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeContactGroupListWithChan(request *DescribeContactGroupListRequest) (<-chan *DescribeContactGroupListResponse, <-chan error) {
	responseChan := make(chan *DescribeContactGroupListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeContactGroupList(request)
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

// DescribeContactGroupListWithCallback invokes the cms.DescribeContactGroupList API asynchronously
// api document: https://help.aliyun.com/api/cms/describecontactgrouplist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeContactGroupListWithCallback(request *DescribeContactGroupListRequest, callback func(response *DescribeContactGroupListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeContactGroupListResponse
		var err error
		defer close(result)
		response, err = client.DescribeContactGroupList(request)
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

// DescribeContactGroupListRequest is the request struct for api DescribeContactGroupList
type DescribeContactGroupListRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeContactGroupListResponse is the response struct for api DescribeContactGroupList
type DescribeContactGroupListResponse struct {
	*responses.BaseResponse
	Success       bool                                    `json:"Success" xml:"Success"`
	Code          string                                  `json:"Code" xml:"Code"`
	Message       string                                  `json:"Message" xml:"Message"`
	Total         int                                     `json:"Total" xml:"Total"`
	RequestId     string                                  `json:"RequestId" xml:"RequestId"`
	ContactGroups ContactGroupsInDescribeContactGroupList `json:"ContactGroups" xml:"ContactGroups"`
}

// CreateDescribeContactGroupListRequest creates a request to invoke DescribeContactGroupList API
func CreateDescribeContactGroupListRequest() (request *DescribeContactGroupListRequest) {
	request = &DescribeContactGroupListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeContactGroupList", "cms", "openAPI")
	return
}

// CreateDescribeContactGroupListResponse creates a response to parse from DescribeContactGroupList response
func CreateDescribeContactGroupListResponse() (response *DescribeContactGroupListResponse) {
	response = &DescribeContactGroupListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
