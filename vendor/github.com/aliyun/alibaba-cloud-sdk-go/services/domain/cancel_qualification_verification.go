package domain

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

// CancelQualificationVerification invokes the domain.CancelQualificationVerification API synchronously
// api document: https://help.aliyun.com/api/domain/cancelqualificationverification.html
func (client *Client) CancelQualificationVerification(request *CancelQualificationVerificationRequest) (response *CancelQualificationVerificationResponse, err error) {
	response = CreateCancelQualificationVerificationResponse()
	err = client.DoAction(request, response)
	return
}

// CancelQualificationVerificationWithChan invokes the domain.CancelQualificationVerification API asynchronously
// api document: https://help.aliyun.com/api/domain/cancelqualificationverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelQualificationVerificationWithChan(request *CancelQualificationVerificationRequest) (<-chan *CancelQualificationVerificationResponse, <-chan error) {
	responseChan := make(chan *CancelQualificationVerificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CancelQualificationVerification(request)
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

// CancelQualificationVerificationWithCallback invokes the domain.CancelQualificationVerification API asynchronously
// api document: https://help.aliyun.com/api/domain/cancelqualificationverification.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CancelQualificationVerificationWithCallback(request *CancelQualificationVerificationRequest, callback func(response *CancelQualificationVerificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CancelQualificationVerificationResponse
		var err error
		defer close(result)
		response, err = client.CancelQualificationVerification(request)
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

// CancelQualificationVerificationRequest is the request struct for api CancelQualificationVerification
type CancelQualificationVerificationRequest struct {
	*requests.RpcRequest
	QualificationType string `position:"Query" name:"QualificationType"`
	InstanceId        string `position:"Query" name:"InstanceId"`
	UserClientIp      string `position:"Query" name:"UserClientIp"`
	Lang              string `position:"Query" name:"Lang"`
}

// CancelQualificationVerificationResponse is the response struct for api CancelQualificationVerification
type CancelQualificationVerificationResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateCancelQualificationVerificationRequest creates a request to invoke CancelQualificationVerification API
func CreateCancelQualificationVerificationRequest() (request *CancelQualificationVerificationRequest) {
	request = &CancelQualificationVerificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Domain", "2018-01-29", "CancelQualificationVerification", "domain", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCancelQualificationVerificationResponse creates a response to parse from CancelQualificationVerification response
func CreateCancelQualificationVerificationResponse() (response *CancelQualificationVerificationResponse) {
	response = &CancelQualificationVerificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
