package ddoscoo

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

// ListLayer7CustomPorts invokes the ddoscoo.ListLayer7CustomPorts API synchronously
func (client *Client) ListLayer7CustomPorts(request *ListLayer7CustomPortsRequest) (response *ListLayer7CustomPortsResponse, err error) {
	response = CreateListLayer7CustomPortsResponse()
	err = client.DoAction(request, response)
	return
}

// ListLayer7CustomPortsWithChan invokes the ddoscoo.ListLayer7CustomPorts API asynchronously
func (client *Client) ListLayer7CustomPortsWithChan(request *ListLayer7CustomPortsRequest) (<-chan *ListLayer7CustomPortsResponse, <-chan error) {
	responseChan := make(chan *ListLayer7CustomPortsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListLayer7CustomPorts(request)
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

// ListLayer7CustomPortsWithCallback invokes the ddoscoo.ListLayer7CustomPorts API asynchronously
func (client *Client) ListLayer7CustomPortsWithCallback(request *ListLayer7CustomPortsRequest, callback func(response *ListLayer7CustomPortsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListLayer7CustomPortsResponse
		var err error
		defer close(result)
		response, err = client.ListLayer7CustomPorts(request)
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

// ListLayer7CustomPortsRequest is the request struct for api ListLayer7CustomPorts
type ListLayer7CustomPortsRequest struct {
	*requests.RpcRequest
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SourceIp        string `position:"Query" name:"SourceIp"`
	Lang            string `position:"Query" name:"Lang"`
}

// ListLayer7CustomPortsResponse is the response struct for api ListLayer7CustomPorts
type ListLayer7CustomPortsResponse struct {
	*responses.BaseResponse
	RequestId         string             `json:"RequestId" xml:"RequestId"`
	Layer7CustomPorts []Layer7CustomPort `json:"Layer7CustomPorts" xml:"Layer7CustomPorts"`
}

// CreateListLayer7CustomPortsRequest creates a request to invoke ListLayer7CustomPorts API
func CreateListLayer7CustomPortsRequest() (request *ListLayer7CustomPortsRequest) {
	request = &ListLayer7CustomPortsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ddoscoo", "2017-12-28", "ListLayer7CustomPorts", "", "")
	request.Method = requests.POST
	return
}

// CreateListLayer7CustomPortsResponse creates a response to parse from ListLayer7CustomPorts response
func CreateListLayer7CustomPortsResponse() (response *ListLayer7CustomPortsResponse) {
	response = &ListLayer7CustomPortsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
