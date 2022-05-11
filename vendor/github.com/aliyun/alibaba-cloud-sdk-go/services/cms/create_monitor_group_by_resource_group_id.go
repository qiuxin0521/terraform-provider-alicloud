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

// CreateMonitorGroupByResourceGroupId invokes the cms.CreateMonitorGroupByResourceGroupId API synchronously
func (client *Client) CreateMonitorGroupByResourceGroupId(request *CreateMonitorGroupByResourceGroupIdRequest) (response *CreateMonitorGroupByResourceGroupIdResponse, err error) {
	response = CreateCreateMonitorGroupByResourceGroupIdResponse()
	err = client.DoAction(request, response)
	return
}

// CreateMonitorGroupByResourceGroupIdWithChan invokes the cms.CreateMonitorGroupByResourceGroupId API asynchronously
func (client *Client) CreateMonitorGroupByResourceGroupIdWithChan(request *CreateMonitorGroupByResourceGroupIdRequest) (<-chan *CreateMonitorGroupByResourceGroupIdResponse, <-chan error) {
	responseChan := make(chan *CreateMonitorGroupByResourceGroupIdResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateMonitorGroupByResourceGroupId(request)
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

// CreateMonitorGroupByResourceGroupIdWithCallback invokes the cms.CreateMonitorGroupByResourceGroupId API asynchronously
func (client *Client) CreateMonitorGroupByResourceGroupIdWithCallback(request *CreateMonitorGroupByResourceGroupIdRequest, callback func(response *CreateMonitorGroupByResourceGroupIdResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateMonitorGroupByResourceGroupIdResponse
		var err error
		defer close(result)
		response, err = client.CreateMonitorGroupByResourceGroupId(request)
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

// CreateMonitorGroupByResourceGroupIdRequest is the request struct for api CreateMonitorGroupByResourceGroupId
type CreateMonitorGroupByResourceGroupIdRequest struct {
	*requests.RpcRequest
	ResourceGroupName    string           `position:"Query" name:"ResourceGroupName"`
	EnableSubscribeEvent requests.Boolean `position:"Query" name:"EnableSubscribeEvent"`
	ResourceGroupId      string           `position:"Query" name:"ResourceGroupId"`
	EnableInstallAgent   requests.Boolean `position:"Query" name:"EnableInstallAgent"`
	ContactGroupList     *[]string        `position:"Query" name:"ContactGroupList"  type:"Repeated"`
}

// CreateMonitorGroupByResourceGroupIdResponse is the response struct for api CreateMonitorGroupByResourceGroupId
type CreateMonitorGroupByResourceGroupIdResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Id        int64  `json:"Id" xml:"Id"`
}

// CreateCreateMonitorGroupByResourceGroupIdRequest creates a request to invoke CreateMonitorGroupByResourceGroupId API
func CreateCreateMonitorGroupByResourceGroupIdRequest() (request *CreateMonitorGroupByResourceGroupIdRequest) {
	request = &CreateMonitorGroupByResourceGroupIdRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "CreateMonitorGroupByResourceGroupId", "Cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateMonitorGroupByResourceGroupIdResponse creates a response to parse from CreateMonitorGroupByResourceGroupId response
func CreateCreateMonitorGroupByResourceGroupIdResponse() (response *CreateMonitorGroupByResourceGroupIdResponse) {
	response = &CreateMonitorGroupByResourceGroupIdResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
