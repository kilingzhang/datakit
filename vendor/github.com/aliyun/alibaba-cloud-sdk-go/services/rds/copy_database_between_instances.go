package rds

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

// CopyDatabaseBetweenInstances invokes the rds.CopyDatabaseBetweenInstances API synchronously
// api document: https://help.aliyun.com/api/rds/copydatabasebetweeninstances.html
func (client *Client) CopyDatabaseBetweenInstances(request *CopyDatabaseBetweenInstancesRequest) (response *CopyDatabaseBetweenInstancesResponse, err error) {
	response = CreateCopyDatabaseBetweenInstancesResponse()
	err = client.DoAction(request, response)
	return
}

// CopyDatabaseBetweenInstancesWithChan invokes the rds.CopyDatabaseBetweenInstances API asynchronously
// api document: https://help.aliyun.com/api/rds/copydatabasebetweeninstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyDatabaseBetweenInstancesWithChan(request *CopyDatabaseBetweenInstancesRequest) (<-chan *CopyDatabaseBetweenInstancesResponse, <-chan error) {
	responseChan := make(chan *CopyDatabaseBetweenInstancesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CopyDatabaseBetweenInstances(request)
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

// CopyDatabaseBetweenInstancesWithCallback invokes the rds.CopyDatabaseBetweenInstances API asynchronously
// api document: https://help.aliyun.com/api/rds/copydatabasebetweeninstances.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CopyDatabaseBetweenInstancesWithCallback(request *CopyDatabaseBetweenInstancesRequest, callback func(response *CopyDatabaseBetweenInstancesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CopyDatabaseBetweenInstancesResponse
		var err error
		defer close(result)
		response, err = client.CopyDatabaseBetweenInstances(request)
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

// CopyDatabaseBetweenInstancesRequest is the request struct for api CopyDatabaseBetweenInstances
type CopyDatabaseBetweenInstancesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	SyncUserPrivilege    string           `position:"Query" name:"SyncUserPrivilege"`
	DbNames              string           `position:"Query" name:"DbNames"`
	TargetDBInstanceId   string           `position:"Query" name:"TargetDBInstanceId"`
	PayType              string           `position:"Query" name:"PayType"`
}

// CopyDatabaseBetweenInstancesResponse is the response struct for api CopyDatabaseBetweenInstances
type CopyDatabaseBetweenInstancesResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
}

// CreateCopyDatabaseBetweenInstancesRequest creates a request to invoke CopyDatabaseBetweenInstances API
func CreateCopyDatabaseBetweenInstancesRequest() (request *CopyDatabaseBetweenInstancesRequest) {
	request = &CopyDatabaseBetweenInstancesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CopyDatabaseBetweenInstances", "rds", "openAPI")
	return
}

// CreateCopyDatabaseBetweenInstancesResponse creates a response to parse from CopyDatabaseBetweenInstances response
func CreateCopyDatabaseBetweenInstancesResponse() (response *CopyDatabaseBetweenInstancesResponse) {
	response = &CopyDatabaseBetweenInstancesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
