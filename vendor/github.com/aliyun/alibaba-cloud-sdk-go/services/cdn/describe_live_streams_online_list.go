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

// DescribeLiveStreamsOnlineList invokes the cdn.DescribeLiveStreamsOnlineList API synchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamsonlinelist.html
func (client *Client) DescribeLiveStreamsOnlineList(request *DescribeLiveStreamsOnlineListRequest) (response *DescribeLiveStreamsOnlineListResponse, err error) {
	response = CreateDescribeLiveStreamsOnlineListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeLiveStreamsOnlineListWithChan invokes the cdn.DescribeLiveStreamsOnlineList API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamsonlinelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamsOnlineListWithChan(request *DescribeLiveStreamsOnlineListRequest) (<-chan *DescribeLiveStreamsOnlineListResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamsOnlineListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamsOnlineList(request)
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

// DescribeLiveStreamsOnlineListWithCallback invokes the cdn.DescribeLiveStreamsOnlineList API asynchronously
// api document: https://help.aliyun.com/api/cdn/describelivestreamsonlinelist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeLiveStreamsOnlineListWithCallback(request *DescribeLiveStreamsOnlineListRequest, callback func(response *DescribeLiveStreamsOnlineListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamsOnlineListResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamsOnlineList(request)
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

// DescribeLiveStreamsOnlineListRequest is the request struct for api DescribeLiveStreamsOnlineList
type DescribeLiveStreamsOnlineListRequest struct {
	*requests.RpcRequest
	PageNum       requests.Integer `position:"Query" name:"PageNum"`
	AppName       string           `position:"Query" name:"AppName"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
	StreamType    string           `position:"Query" name:"StreamType"`
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeLiveStreamsOnlineListResponse is the response struct for api DescribeLiveStreamsOnlineList
type DescribeLiveStreamsOnlineListResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	PageNum    int        `json:"PageNum" xml:"PageNum"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	TotalNum   int        `json:"TotalNum" xml:"TotalNum"`
	TotalPage  int        `json:"TotalPage" xml:"TotalPage"`
	OnlineInfo OnlineInfo `json:"OnlineInfo" xml:"OnlineInfo"`
}

// CreateDescribeLiveStreamsOnlineListRequest creates a request to invoke DescribeLiveStreamsOnlineList API
func CreateDescribeLiveStreamsOnlineListRequest() (request *DescribeLiveStreamsOnlineListRequest) {
	request = &DescribeLiveStreamsOnlineListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamsOnlineList", "", "")
	return
}

// CreateDescribeLiveStreamsOnlineListResponse creates a response to parse from DescribeLiveStreamsOnlineList response
func CreateDescribeLiveStreamsOnlineListResponse() (response *DescribeLiveStreamsOnlineListResponse) {
	response = &DescribeLiveStreamsOnlineListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
