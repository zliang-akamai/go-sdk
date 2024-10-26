# \VPCsAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVpc**](VPCsAPI.md#DeleteVpc) | **Delete** /{apiVersion}/vpcs/{vpcId} | Delete a VPC
[**DeleteVpcSubnet**](VPCsAPI.md#DeleteVpcSubnet) | **Delete** /{apiVersion}/vpcs/{vpcId}/subnets/{vpcSubnetId} | Delete a VPC subnet
[**GetVpc**](VPCsAPI.md#GetVpc) | **Get** /{apiVersion}/vpcs/{vpcId} | Get a VPC
[**GetVpcIps**](VPCsAPI.md#GetVpcIps) | **Get** /{apiVersion}/vpcs/{vpcId}/ips | List a VPC&#39;s IP addresses
[**GetVpcSubnet**](VPCsAPI.md#GetVpcSubnet) | **Get** /{apiVersion}/vpcs/{vpcId}/subnets/{vpcSubnetId} | Get a VPC subnet
[**GetVpcSubnets**](VPCsAPI.md#GetVpcSubnets) | **Get** /{apiVersion}/vpcs/{vpcId}/subnets | List VPC subnets
[**GetVpcs**](VPCsAPI.md#GetVpcs) | **Get** /{apiVersion}/vpcs | List VPCs
[**GetVpcsIps**](VPCsAPI.md#GetVpcsIps) | **Get** /{apiVersion}/vpcs/ips | List VPC IP addresses
[**PostVpc**](VPCsAPI.md#PostVpc) | **Post** /{apiVersion}/vpcs | Create a VPC
[**PostVpcSubnet**](VPCsAPI.md#PostVpcSubnet) | **Post** /{apiVersion}/vpcs/{vpcId}/subnets | Create a VPC subnet
[**PutVpc**](VPCsAPI.md#PutVpc) | **Put** /{apiVersion}/vpcs/{vpcId} | Update a VPC
[**PutVpcSubnet**](VPCsAPI.md#PutVpcSubnet) | **Put** /{apiVersion}/vpcs/{vpcId}/subnets/{vpcSubnetId} | Update a VPC subnet



## DeleteVpc

> map[string]interface{} DeleteVpc(ctx, apiVersion, vpcId).Execute()

Delete a VPC



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	vpcId := int32(56) // int32 | The `id` of the VPC.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.DeleteVpc(context.Background(), apiVersion, vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.DeleteVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVpc`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.DeleteVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**vpcId** | **int32** | The &#x60;id&#x60; of the VPC. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**map[string]interface{}**

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVpcSubnet

> map[string]interface{} DeleteVpcSubnet(ctx, apiVersion, vpcId, vpcSubnetId).Execute()

Delete a VPC subnet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	vpcId := int32(56) // int32 | The `id` of the VPC.
	vpcSubnetId := int32(56) // int32 | The `id` of the VPC Subnet.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.DeleteVpcSubnet(context.Background(), apiVersion, vpcId, vpcSubnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.DeleteVpcSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVpcSubnet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.DeleteVpcSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**vpcId** | **int32** | The &#x60;id&#x60; of the VPC. | 
**vpcSubnetId** | **int32** | The &#x60;id&#x60; of the VPC Subnet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVpcSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**map[string]interface{}**

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpc

> GetVpcs200ResponseAllOfDataInner GetVpc(ctx, apiVersion, vpcId).Execute()

Get a VPC



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	vpcId := int32(56) // int32 | The `id` of the VPC.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.GetVpc(context.Background(), apiVersion, vpcId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.GetVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpc`: GetVpcs200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.GetVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**vpcId** | **int32** | The &#x60;id&#x60; of the VPC. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVpcs200ResponseAllOfDataInner**](GetVpcs200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpcIps

> GetVpcsIps200Response GetVpcIps(ctx, apiVersion, vpcId).Page(page).PageSize(pageSize).Execute()

List a VPC's IP addresses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	vpcId := int32(56) // int32 | The `id` of the VPC.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.GetVpcIps(context.Background(), apiVersion, vpcId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.GetVpcIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpcIps`: GetVpcsIps200Response
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.GetVpcIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**vpcId** | **int32** | The &#x60;id&#x60; of the VPC. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetVpcsIps200Response**](GetVpcsIps200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpcSubnet

> GetVpcs200ResponseAllOfDataInnerSubnetsInner GetVpcSubnet(ctx, apiVersion, vpcId, vpcSubnetId).Execute()

Get a VPC subnet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	vpcId := int32(56) // int32 | The `id` of the VPC.
	vpcSubnetId := int32(56) // int32 | The `id` of the VPC Subnet.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.GetVpcSubnet(context.Background(), apiVersion, vpcId, vpcSubnetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.GetVpcSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpcSubnet`: GetVpcs200ResponseAllOfDataInnerSubnetsInner
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.GetVpcSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**vpcId** | **int32** | The &#x60;id&#x60; of the VPC. | 
**vpcSubnetId** | **int32** | The &#x60;id&#x60; of the VPC Subnet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetVpcs200ResponseAllOfDataInnerSubnetsInner**](GetVpcs200ResponseAllOfDataInnerSubnetsInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpcSubnets

> GetVpcSubnets200Response GetVpcSubnets(ctx, apiVersion, vpcId).Page(page).PageSize(pageSize).Execute()

List VPC subnets



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	vpcId := int32(56) // int32 | The `id` of the VPC.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.GetVpcSubnets(context.Background(), apiVersion, vpcId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.GetVpcSubnets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpcSubnets`: GetVpcSubnets200Response
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.GetVpcSubnets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**vpcId** | **int32** | The &#x60;id&#x60; of the VPC. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcSubnetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetVpcSubnets200Response**](GetVpcSubnets200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpcs

> GetVpcs200Response GetVpcs(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List VPCs



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.GetVpcs(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.GetVpcs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpcs`: GetVpcs200Response
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.GetVpcs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetVpcs200Response**](GetVpcs200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVpcsIps

> GetVpcsIps200Response GetVpcsIps(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List VPC IP addresses



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.GetVpcsIps(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.GetVpcsIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVpcsIps`: GetVpcsIps200Response
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.GetVpcsIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetVpcsIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetVpcsIps200Response**](GetVpcsIps200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVpc

> GetVpcs200ResponseAllOfDataInner PostVpc(ctx, apiVersion).PostVpcRequest(postVpcRequest).Execute()

Create a VPC



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	postVpcRequest := *openapiclient.NewPostVpcRequest("cool-vpc", "us-east") // PostVpcRequest | VPC Create request object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.PostVpc(context.Background(), apiVersion).PostVpcRequest(postVpcRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.PostVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVpc`: GetVpcs200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.PostVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postVpcRequest** | [**PostVpcRequest**](PostVpcRequest.md) | VPC Create request object. | 

### Return type

[**GetVpcs200ResponseAllOfDataInner**](GetVpcs200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVpcSubnet

> GetVpcs200ResponseAllOfDataInnerSubnetsInner PostVpcSubnet(ctx, apiVersion, vpcId).PostVpcSubnetRequest(postVpcSubnetRequest).Execute()

Create a VPC subnet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	vpcId := int32(56) // int32 | The `id` of the VPC.
	postVpcSubnetRequest := *openapiclient.NewPostVpcSubnetRequest("10.0.1.0/24", "cool-vpc-subnet") // PostVpcSubnetRequest | VPC Subnet Create request object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.PostVpcSubnet(context.Background(), apiVersion, vpcId).PostVpcSubnetRequest(postVpcSubnetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.PostVpcSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVpcSubnet`: GetVpcs200ResponseAllOfDataInnerSubnetsInner
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.PostVpcSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**vpcId** | **int32** | The &#x60;id&#x60; of the VPC. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostVpcSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postVpcSubnetRequest** | [**PostVpcSubnetRequest**](PostVpcSubnetRequest.md) | VPC Subnet Create request object. | 

### Return type

[**GetVpcs200ResponseAllOfDataInnerSubnetsInner**](GetVpcs200ResponseAllOfDataInnerSubnetsInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVpc

> GetVpcs200ResponseAllOfDataInner PutVpc(ctx, apiVersion, vpcId).PutVpcRequest(putVpcRequest).Execute()

Update a VPC



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	vpcId := int32(56) // int32 | The `id` of the VPC.
	putVpcRequest := *openapiclient.NewPutVpcRequest() // PutVpcRequest | VPC Update request object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.PutVpc(context.Background(), apiVersion, vpcId).PutVpcRequest(putVpcRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.PutVpc``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutVpc`: GetVpcs200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.PutVpc`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**vpcId** | **int32** | The &#x60;id&#x60; of the VPC. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVpcRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putVpcRequest** | [**PutVpcRequest**](PutVpcRequest.md) | VPC Update request object. | 

### Return type

[**GetVpcs200ResponseAllOfDataInner**](GetVpcs200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVpcSubnet

> GetVpcs200ResponseAllOfDataInnerSubnetsInner PutVpcSubnet(ctx, apiVersion, vpcId, vpcSubnetId).PutVpcSubnetRequest(putVpcSubnetRequest).Execute()

Update a VPC subnet



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	apiVersion := "apiVersion_example" // string | __Enum__ Call either the `v4` URL, or `v4beta` for operations still in Beta. (default to "v4")
	vpcId := int32(56) // int32 | The `id` of the VPC.
	vpcSubnetId := int32(56) // int32 | The `id` of the VPC Subnet.
	putVpcSubnetRequest := *openapiclient.NewPutVpcSubnetRequest() // PutVpcSubnetRequest | VPC Update request object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VPCsAPI.PutVpcSubnet(context.Background(), apiVersion, vpcId, vpcSubnetId).PutVpcSubnetRequest(putVpcSubnetRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VPCsAPI.PutVpcSubnet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutVpcSubnet`: GetVpcs200ResponseAllOfDataInnerSubnetsInner
	fmt.Fprintf(os.Stdout, "Response from `VPCsAPI.PutVpcSubnet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**vpcId** | **int32** | The &#x60;id&#x60; of the VPC. | 
**vpcSubnetId** | **int32** | The &#x60;id&#x60; of the VPC Subnet. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVpcSubnetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putVpcSubnetRequest** | [**PutVpcSubnetRequest**](PutVpcSubnetRequest.md) | VPC Update request object. | 

### Return type

[**GetVpcs200ResponseAllOfDataInnerSubnetsInner**](GetVpcs200ResponseAllOfDataInnerSubnetsInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

