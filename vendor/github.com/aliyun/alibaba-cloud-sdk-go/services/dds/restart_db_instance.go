package dds

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

// RestartDBInstance invokes the dds.RestartDBInstance API synchronously
// api document: https://help.aliyun.com/api/dds/restartdbinstance.html
func (client *Client) RestartDBInstance(request *RestartDBInstanceRequest) (response *RestartDBInstanceResponse, err error) {
	response = CreateRestartDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RestartDBInstanceWithChan invokes the dds.RestartDBInstance API asynchronously
// api document: https://help.aliyun.com/api/dds/restartdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestartDBInstanceWithChan(request *RestartDBInstanceRequest) (<-chan *RestartDBInstanceResponse, <-chan error) {
	responseChan := make(chan *RestartDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestartDBInstance(request)
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

// RestartDBInstanceWithCallback invokes the dds.RestartDBInstance API asynchronously
// api document: https://help.aliyun.com/api/dds/restartdbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestartDBInstanceWithCallback(request *RestartDBInstanceRequest, callback func(response *RestartDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestartDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.RestartDBInstance(request)
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

// RestartDBInstanceRequest is the request struct for api RestartDBInstance
type RestartDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	NodeId               string           `position:"Query" name:"NodeId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RestartDBInstanceResponse is the response struct for api RestartDBInstance
type RestartDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRestartDBInstanceRequest creates a request to invoke RestartDBInstance API
func CreateRestartDBInstanceRequest() (request *RestartDBInstanceRequest) {
	request = &RestartDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dds", "2015-12-01", "RestartDBInstance", "Dds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRestartDBInstanceResponse creates a response to parse from RestartDBInstance response
func CreateRestartDBInstanceResponse() (response *RestartDBInstanceResponse) {
	response = &RestartDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
