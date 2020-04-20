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

// DeleteCustomMetric invokes the cms.DeleteCustomMetric API synchronously
// api document: https://help.aliyun.com/api/cms/deletecustommetric.html
func (client *Client) DeleteCustomMetric(request *DeleteCustomMetricRequest) (response *DeleteCustomMetricResponse, err error) {
	response = CreateDeleteCustomMetricResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteCustomMetricWithChan invokes the cms.DeleteCustomMetric API asynchronously
// api document: https://help.aliyun.com/api/cms/deletecustommetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCustomMetricWithChan(request *DeleteCustomMetricRequest) (<-chan *DeleteCustomMetricResponse, <-chan error) {
	responseChan := make(chan *DeleteCustomMetricResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteCustomMetric(request)
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

// DeleteCustomMetricWithCallback invokes the cms.DeleteCustomMetric API asynchronously
// api document: https://help.aliyun.com/api/cms/deletecustommetric.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteCustomMetricWithCallback(request *DeleteCustomMetricRequest, callback func(response *DeleteCustomMetricResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteCustomMetricResponse
		var err error
		defer close(result)
		response, err = client.DeleteCustomMetric(request)
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

// DeleteCustomMetricRequest is the request struct for api DeleteCustomMetric
type DeleteCustomMetricRequest struct {
	*requests.RpcRequest
	GroupId    string `position:"Query" name:"GroupId"`
	UUID       string `position:"Query" name:"UUID"`
	MetricName string `position:"Query" name:"MetricName"`
	Md5        string `position:"Query" name:"Md5"`
}

// DeleteCustomMetricResponse is the response struct for api DeleteCustomMetric
type DeleteCustomMetricResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteCustomMetricRequest creates a request to invoke DeleteCustomMetric API
func CreateDeleteCustomMetricRequest() (request *DeleteCustomMetricRequest) {
	request = &DeleteCustomMetricRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DeleteCustomMetric", "cms", "openAPI")
	return
}

// CreateDeleteCustomMetricResponse creates a response to parse from DeleteCustomMetric response
func CreateDeleteCustomMetricResponse() (response *DeleteCustomMetricResponse) {
	response = &DeleteCustomMetricResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
