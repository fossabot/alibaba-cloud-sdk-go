package cs

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

// DescribeServiceContainers invokes the cs.DescribeServiceContainers API synchronously
// api document: https://help.aliyun.com/api/cs/describeservicecontainers.html
func (client *Client) DescribeServiceContainers(request *DescribeServiceContainersRequest) (response *DescribeServiceContainersResponse, err error) {
	response = CreateDescribeServiceContainersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeServiceContainersWithChan invokes the cs.DescribeServiceContainers API asynchronously
// api document: https://help.aliyun.com/api/cs/describeservicecontainers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServiceContainersWithChan(request *DescribeServiceContainersRequest) (<-chan *DescribeServiceContainersResponse, <-chan error) {
	responseChan := make(chan *DescribeServiceContainersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeServiceContainers(request)
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

// DescribeServiceContainersWithCallback invokes the cs.DescribeServiceContainers API asynchronously
// api document: https://help.aliyun.com/api/cs/describeservicecontainers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeServiceContainersWithCallback(request *DescribeServiceContainersRequest, callback func(response *DescribeServiceContainersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeServiceContainersResponse
		var err error
		defer close(result)
		response, err = client.DescribeServiceContainers(request)
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

// DescribeServiceContainersRequest is the request struct for api DescribeServiceContainers
type DescribeServiceContainersRequest struct {
	*requests.RoaRequest
	ClusterId string `position:"Path" name:"ClusterId"`
	ServiceId string `position:"Path" name:"ServiceId"`
}

// DescribeServiceContainersResponse is the response struct for api DescribeServiceContainers
type DescribeServiceContainersResponse struct {
	*responses.BaseResponse
}

// CreateDescribeServiceContainersRequest creates a request to invoke DescribeServiceContainers API
func CreateDescribeServiceContainersRequest() (request *DescribeServiceContainersRequest) {
	request = &DescribeServiceContainersRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeServiceContainers", "/clusters/[ClusterId]/services/[ServiceId]/containers", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeServiceContainersResponse creates a response to parse from DescribeServiceContainers response
func CreateDescribeServiceContainersResponse() (response *DescribeServiceContainersResponse) {
	response = &DescribeServiceContainersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
