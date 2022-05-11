package cbn

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

// DescribeChildInstanceRegions invokes the cbn.DescribeChildInstanceRegions API synchronously
func (client *Client) DescribeChildInstanceRegions(request *DescribeChildInstanceRegionsRequest) (response *DescribeChildInstanceRegionsResponse, err error) {
	response = CreateDescribeChildInstanceRegionsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeChildInstanceRegionsWithChan invokes the cbn.DescribeChildInstanceRegions API asynchronously
func (client *Client) DescribeChildInstanceRegionsWithChan(request *DescribeChildInstanceRegionsRequest) (<-chan *DescribeChildInstanceRegionsResponse, <-chan error) {
	responseChan := make(chan *DescribeChildInstanceRegionsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeChildInstanceRegions(request)
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

// DescribeChildInstanceRegionsWithCallback invokes the cbn.DescribeChildInstanceRegions API asynchronously
func (client *Client) DescribeChildInstanceRegionsWithCallback(request *DescribeChildInstanceRegionsRequest, callback func(response *DescribeChildInstanceRegionsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeChildInstanceRegionsResponse
		var err error
		defer close(result)
		response, err = client.DescribeChildInstanceRegions(request)
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

// DescribeChildInstanceRegionsRequest is the request struct for api DescribeChildInstanceRegions
type DescribeChildInstanceRegionsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ProductType          string           `position:"Query" name:"ProductType"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ChildInstanceOwnerId requests.Integer `position:"Query" name:"ChildInstanceOwnerId"`
}

// DescribeChildInstanceRegionsResponse is the response struct for api DescribeChildInstanceRegions
type DescribeChildInstanceRegionsResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	Regions   Regions `json:"Regions" xml:"Regions"`
}

// CreateDescribeChildInstanceRegionsRequest creates a request to invoke DescribeChildInstanceRegions API
func CreateDescribeChildInstanceRegionsRequest() (request *DescribeChildInstanceRegionsRequest) {
	request = &DescribeChildInstanceRegionsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "DescribeChildInstanceRegions", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeChildInstanceRegionsResponse creates a response to parse from DescribeChildInstanceRegions response
func CreateDescribeChildInstanceRegionsResponse() (response *DescribeChildInstanceRegionsResponse) {
	response = &DescribeChildInstanceRegionsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
