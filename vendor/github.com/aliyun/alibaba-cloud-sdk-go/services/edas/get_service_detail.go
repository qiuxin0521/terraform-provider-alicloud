package edas

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

// GetServiceDetail invokes the edas.GetServiceDetail API synchronously
func (client *Client) GetServiceDetail(request *GetServiceDetailRequest) (response *GetServiceDetailResponse, err error) {
	response = CreateGetServiceDetailResponse()
	err = client.DoAction(request, response)
	return
}

// GetServiceDetailWithChan invokes the edas.GetServiceDetail API asynchronously
func (client *Client) GetServiceDetailWithChan(request *GetServiceDetailRequest) (<-chan *GetServiceDetailResponse, <-chan error) {
	responseChan := make(chan *GetServiceDetailResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetServiceDetail(request)
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

// GetServiceDetailWithCallback invokes the edas.GetServiceDetail API asynchronously
func (client *Client) GetServiceDetailWithCallback(request *GetServiceDetailRequest, callback func(response *GetServiceDetailResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetServiceDetailResponse
		var err error
		defer close(result)
		response, err = client.GetServiceDetail(request)
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

// GetServiceDetailRequest is the request struct for api GetServiceDetail
type GetServiceDetailRequest struct {
	*requests.RoaRequest
	RegistryType   string `position:"Query" name:"registryType"`
	ServiceType    string `position:"Query" name:"serviceType"`
	Origin         string `position:"Query" name:"origin"`
	AppId          string `position:"Query" name:"appId"`
	Ip             string `position:"Query" name:"ip"`
	Namespace      string `position:"Query" name:"namespace"`
	ServiceVersion string `position:"Query" name:"serviceVersion"`
	ServiceName    string `position:"Query" name:"serviceName"`
	Source         string `position:"Query" name:"source"`
	Region         string `position:"Query" name:"region"`
	ServiceId      string `position:"Query" name:"serviceId"`
	Group          string `position:"Query" name:"group"`
}

// GetServiceDetailResponse is the response struct for api GetServiceDetail
type GetServiceDetailResponse struct {
	*responses.BaseResponse
	Code    int    `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
	Success bool   `json:"Success" xml:"Success"`
	Data    Data   `json:"Data" xml:"Data"`
}

// CreateGetServiceDetailRequest creates a request to invoke GetServiceDetail API
func CreateGetServiceDetailRequest() (request *GetServiceDetailRequest) {
	request = &GetServiceDetailRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "GetServiceDetail", "/pop/sp/api/mseForOam/getServiceDetail", "Edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetServiceDetailResponse creates a response to parse from GetServiceDetail response
func CreateGetServiceDetailResponse() (response *GetServiceDetailResponse) {
	response = &GetServiceDetailResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
