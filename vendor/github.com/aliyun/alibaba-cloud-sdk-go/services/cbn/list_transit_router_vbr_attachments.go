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

// ListTransitRouterVbrAttachments invokes the cbn.ListTransitRouterVbrAttachments API synchronously
func (client *Client) ListTransitRouterVbrAttachments(request *ListTransitRouterVbrAttachmentsRequest) (response *ListTransitRouterVbrAttachmentsResponse, err error) {
	response = CreateListTransitRouterVbrAttachmentsResponse()
	err = client.DoAction(request, response)
	return
}

// ListTransitRouterVbrAttachmentsWithChan invokes the cbn.ListTransitRouterVbrAttachments API asynchronously
func (client *Client) ListTransitRouterVbrAttachmentsWithChan(request *ListTransitRouterVbrAttachmentsRequest) (<-chan *ListTransitRouterVbrAttachmentsResponse, <-chan error) {
	responseChan := make(chan *ListTransitRouterVbrAttachmentsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListTransitRouterVbrAttachments(request)
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

// ListTransitRouterVbrAttachmentsWithCallback invokes the cbn.ListTransitRouterVbrAttachments API asynchronously
func (client *Client) ListTransitRouterVbrAttachmentsWithCallback(request *ListTransitRouterVbrAttachmentsRequest, callback func(response *ListTransitRouterVbrAttachmentsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListTransitRouterVbrAttachmentsResponse
		var err error
		defer close(result)
		response, err = client.ListTransitRouterVbrAttachments(request)
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

// ListTransitRouterVbrAttachmentsRequest is the request struct for api ListTransitRouterVbrAttachments
type ListTransitRouterVbrAttachmentsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId           requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CenId                     string           `position:"Query" name:"CenId"`
	NextToken                 string           `position:"Query" name:"NextToken"`
	ResourceOwnerAccount      string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount              string           `position:"Query" name:"OwnerAccount"`
	OwnerId                   requests.Integer `position:"Query" name:"OwnerId"`
	TransitRouterId           string           `position:"Query" name:"TransitRouterId"`
	ResourceType              string           `position:"Query" name:"ResourceType"`
	TransitRouterAttachmentId string           `position:"Query" name:"TransitRouterAttachmentId"`
	MaxResults                requests.Integer `position:"Query" name:"MaxResults"`
}

// ListTransitRouterVbrAttachmentsResponse is the response struct for api ListTransitRouterVbrAttachments
type ListTransitRouterVbrAttachmentsResponse struct {
	*responses.BaseResponse
	NextToken                string                    `json:"NextToken" xml:"NextToken"`
	RequestId                string                    `json:"RequestId" xml:"RequestId"`
	TotalCount               int                       `json:"TotalCount" xml:"TotalCount"`
	MaxResults               int                       `json:"MaxResults" xml:"MaxResults"`
	TransitRouterAttachments []TransitRouterAttachment `json:"TransitRouterAttachments" xml:"TransitRouterAttachments"`
}

// CreateListTransitRouterVbrAttachmentsRequest creates a request to invoke ListTransitRouterVbrAttachments API
func CreateListTransitRouterVbrAttachmentsRequest() (request *ListTransitRouterVbrAttachmentsRequest) {
	request = &ListTransitRouterVbrAttachmentsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "ListTransitRouterVbrAttachments", "", "")
	request.Method = requests.POST
	return
}

// CreateListTransitRouterVbrAttachmentsResponse creates a response to parse from ListTransitRouterVbrAttachments response
func CreateListTransitRouterVbrAttachmentsResponse() (response *ListTransitRouterVbrAttachmentsResponse) {
	response = &ListTransitRouterVbrAttachmentsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
