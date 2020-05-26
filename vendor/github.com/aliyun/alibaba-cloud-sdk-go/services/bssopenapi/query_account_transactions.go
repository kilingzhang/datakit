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

// QueryAccountTransactions invokes the bssopenapi.QueryAccountTransactions API synchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryaccounttransactions.html
func (client *Client) QueryAccountTransactions(request *QueryAccountTransactionsRequest) (response *QueryAccountTransactionsResponse, err error) {
	response = CreateQueryAccountTransactionsResponse()
	err = client.DoAction(request, response)
	return
}

// QueryAccountTransactionsWithChan invokes the bssopenapi.QueryAccountTransactions API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryaccounttransactions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAccountTransactionsWithChan(request *QueryAccountTransactionsRequest) (<-chan *QueryAccountTransactionsResponse, <-chan error) {
	responseChan := make(chan *QueryAccountTransactionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryAccountTransactions(request)
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

// QueryAccountTransactionsWithCallback invokes the bssopenapi.QueryAccountTransactions API asynchronously
// api document: https://help.aliyun.com/api/bssopenapi/queryaccounttransactions.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryAccountTransactionsWithCallback(request *QueryAccountTransactionsRequest, callback func(response *QueryAccountTransactionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryAccountTransactionsResponse
		var err error
		defer close(result)
		response, err = client.QueryAccountTransactions(request)
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

// QueryAccountTransactionsRequest is the request struct for api QueryAccountTransactions
type QueryAccountTransactionsRequest struct {
	*requests.RpcRequest
	PageNum              requests.Integer `position:"Query" name:"PageNum"`
	CreateTimeEnd        string           `position:"Query" name:"CreateTimeEnd"`
	RecordID             string           `position:"Query" name:"RecordID"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	TransactionChannelSN string           `position:"Query" name:"TransactionChannelSN"`
	CreateTimeStart      string           `position:"Query" name:"CreateTimeStart"`
	TransactionNumber    string           `position:"Query" name:"TransactionNumber"`
}

// QueryAccountTransactionsResponse is the response struct for api QueryAccountTransactions
type QueryAccountTransactionsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryAccountTransactionsRequest creates a request to invoke QueryAccountTransactions API
func CreateQueryAccountTransactionsRequest() (request *QueryAccountTransactionsRequest) {
	request = &QueryAccountTransactionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("BssOpenApi", "2017-12-14", "QueryAccountTransactions", "", "")
	return
}

// CreateQueryAccountTransactionsResponse creates a response to parse from QueryAccountTransactions response
func CreateQueryAccountTransactionsResponse() (response *QueryAccountTransactionsResponse) {
	response = &QueryAccountTransactionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
