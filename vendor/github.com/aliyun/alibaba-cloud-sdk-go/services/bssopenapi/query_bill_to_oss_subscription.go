package bssopenapi

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

// QueryBillToOSSSubscription invokes the bssopenapi.QueryBillToOSSSubscription API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/querybilltoosssubscription.html
func (client *Client) QueryBillToOSSSubscription(request *QueryBillToOSSSubscriptionRequest) (response *QueryBillToOSSSubscriptionResponse, err error) {
	response = CreateQueryBillToOSSSubscriptionResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBillToOSSSubscriptionWithChan invokes the bssopenapi.QueryBillToOSSSubscription API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querybilltoosssubscription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryBillToOSSSubscriptionWithChan(request *QueryBillToOSSSubscriptionRequest) (<-chan *QueryBillToOSSSubscriptionResponse, <-chan error) {
	responseChan := make(chan *QueryBillToOSSSubscriptionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryBillToOSSSubscription(request)
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

// QueryBillToOSSSubscriptionWithCallback invokes the bssopenapi.QueryBillToOSSSubscription API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/querybilltoosssubscription.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryBillToOSSSubscriptionWithCallback(request *QueryBillToOSSSubscriptionRequest, callback func(response *QueryBillToOSSSubscriptionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryBillToOSSSubscriptionResponse
		var err error
		defer close(result)
		response, err = client.QueryBillToOSSSubscription(request)
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

// QueryBillToOSSSubscriptionRequest is the request struct for api QueryBillToOSSSubscription
type QueryBillToOSSSubscriptionRequest struct {
	*requests.RpcRequest
}

// QueryBillToOSSSubscriptionResponse is the response struct for api QueryBillToOSSSubscription
type QueryBillToOSSSubscriptionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryBillToOSSSubscriptionRequest creates a request to invoke QueryBillToOSSSubscription API
func CreateQueryBillToOSSSubscriptionRequest() (request *QueryBillToOSSSubscriptionRequest) {
	request = &QueryBillToOSSSubscriptionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryBillToOSSSubscription", "", "")
	return
}

// CreateQueryBillToOSSSubscriptionResponse creates a response to parse from QueryBillToOSSSubscription response
func CreateQueryBillToOSSSubscriptionResponse() (response *QueryBillToOSSSubscriptionResponse) {
	response = &QueryBillToOSSSubscriptionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
