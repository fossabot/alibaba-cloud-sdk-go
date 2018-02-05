package ecs

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

func (client *Client) DescribeResourcesModification(request *DescribeResourcesModificationRequest) (response *DescribeResourcesModificationResponse, err error) {
	response = CreateDescribeResourcesModificationResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeResourcesModificationWithChan(request *DescribeResourcesModificationRequest) (<-chan *DescribeResourcesModificationResponse, <-chan error) {
	responseChan := make(chan *DescribeResourcesModificationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeResourcesModification(request)
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

func (client *Client) DescribeResourcesModificationWithCallback(request *DescribeResourcesModificationRequest, callback func(response *DescribeResourcesModificationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeResourcesModificationResponse
		var err error
		defer close(result)
		response, err = client.DescribeResourcesModification(request)
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

type DescribeResourcesModificationRequest struct {
	*requests.RpcRequest
}

type DescribeResourcesModificationResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	AvailableZones AvailableZones `json:"AvailableZones" xml:"AvailableZones"`
}

func CreateDescribeResourcesModificationRequest() (request *DescribeResourcesModificationRequest) {
	request = &DescribeResourcesModificationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeResourcesModification", "", "")
	return
}

func CreateDescribeResourcesModificationResponse() (response *DescribeResourcesModificationResponse) {
	response = &DescribeResourcesModificationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
