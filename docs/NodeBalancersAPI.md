# \NodeBalancersAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNodeBalancer**](NodeBalancersAPI.md#DeleteNodeBalancer) | **Delete** /{apiVersion}/nodebalancers/{nodeBalancerId} | Delete a NodeBalancer
[**DeleteNodeBalancerConfig**](NodeBalancersAPI.md#DeleteNodeBalancerConfig) | **Delete** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs/{configId} | Delete a config
[**DeleteNodeBalancerConfigNode**](NodeBalancersAPI.md#DeleteNodeBalancerConfigNode) | **Delete** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs/{configId}/nodes/{nodeId} | Delete a NodeBalancer&#39;s node
[**GetNodeBalancer**](NodeBalancersAPI.md#GetNodeBalancer) | **Get** /{apiVersion}/nodebalancers/{nodeBalancerId} | Get a NodeBalancer
[**GetNodeBalancerConfig**](NodeBalancersAPI.md#GetNodeBalancerConfig) | **Get** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs/{configId} | Get a config
[**GetNodeBalancerConfigNodes**](NodeBalancersAPI.md#GetNodeBalancerConfigNodes) | **Get** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs/{configId}/nodes | List nodes
[**GetNodeBalancerConfigs**](NodeBalancersAPI.md#GetNodeBalancerConfigs) | **Get** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs | List configs
[**GetNodeBalancerFirewalls**](NodeBalancersAPI.md#GetNodeBalancerFirewalls) | **Get** /{apiVersion}/nodebalancers/{nodeBalancerId}/firewalls | List NodeBalancer firewalls
[**GetNodeBalancerNode**](NodeBalancersAPI.md#GetNodeBalancerNode) | **Get** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs/{configId}/nodes/{nodeId} | Get a NodeBalancer&#39;s node
[**GetNodeBalancerStats**](NodeBalancersAPI.md#GetNodeBalancerStats) | **Get** /{apiVersion}/nodebalancers/{nodeBalancerId}/stats | Get NodeBalancer statistics
[**GetNodeBalancerTypes**](NodeBalancersAPI.md#GetNodeBalancerTypes) | **Get** /{apiVersion}/nodebalancers/types | List NodeBalancer types
[**GetNodeBalancers**](NodeBalancersAPI.md#GetNodeBalancers) | **Get** /{apiVersion}/nodebalancers | List NodeBalancers
[**PostNodeBalancer**](NodeBalancersAPI.md#PostNodeBalancer) | **Post** /{apiVersion}/nodebalancers | Create a NodeBalancer
[**PostNodeBalancerConfig**](NodeBalancersAPI.md#PostNodeBalancerConfig) | **Post** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs | Create a config
[**PostNodeBalancerNode**](NodeBalancersAPI.md#PostNodeBalancerNode) | **Post** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs/{configId}/nodes | Create a node
[**PostRebuildNodeBalancerConfig**](NodeBalancersAPI.md#PostRebuildNodeBalancerConfig) | **Post** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs/{configId}/rebuild | Rebuild a config
[**PutNodeBalancer**](NodeBalancersAPI.md#PutNodeBalancer) | **Put** /{apiVersion}/nodebalancers/{nodeBalancerId} | Update a NodeBalancer
[**PutNodeBalancerConfig**](NodeBalancersAPI.md#PutNodeBalancerConfig) | **Put** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs/{configId} | Update a config
[**PutNodeBalancerNode**](NodeBalancersAPI.md#PutNodeBalancerNode) | **Put** /{apiVersion}/nodebalancers/{nodeBalancerId}/configs/{configId}/nodes/{nodeId} | Update a node



## DeleteNodeBalancer

> map[string]interface{} DeleteNodeBalancer(ctx, apiVersion, nodeBalancerId).Execute()

Delete a NodeBalancer



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.DeleteNodeBalancer(context.Background(), apiVersion, nodeBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.DeleteNodeBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNodeBalancer`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.DeleteNodeBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeBalancerRequest struct via the builder pattern


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


## DeleteNodeBalancerConfig

> map[string]interface{} DeleteNodeBalancerConfig(ctx, apiVersion, nodeBalancerId, configId).Execute()

Delete a config



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	configId := int32(56) // int32 | The ID of the Config to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.DeleteNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId, configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.DeleteNodeBalancerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNodeBalancerConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.DeleteNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeBalancerConfigRequest struct via the builder pattern


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


## DeleteNodeBalancerConfigNode

> map[string]interface{} DeleteNodeBalancerConfigNode(ctx, apiVersion, nodeBalancerId, configId, nodeId).Execute()

Delete a NodeBalancer's node



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	configId := int32(56) // int32 | The ID of the Config to access.
	nodeId := "nodeId_example" // string | The ID of the Node to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.DeleteNodeBalancerConfigNode(context.Background(), apiVersion, nodeBalancerId, configId, nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.DeleteNodeBalancerConfigNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNodeBalancerConfigNode`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.DeleteNodeBalancerConfigNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access. | 
**nodeId** | **string** | The ID of the Node to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeBalancerConfigNodeRequest struct via the builder pattern


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


## GetNodeBalancer

> NodeBalancer GetNodeBalancer(ctx, apiVersion, nodeBalancerId).Execute()

Get a NodeBalancer



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.GetNodeBalancer(context.Background(), apiVersion, nodeBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.GetNodeBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeBalancer`: NodeBalancer
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.GetNodeBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NodeBalancer**](NodeBalancer.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerConfig

> GetNodeBalancerConfigs200ResponseDataInner GetNodeBalancerConfig(ctx, apiVersion, nodeBalancerId, configId).Execute()

Get a config



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	configId := int32(56) // int32 | The ID of the Config to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.GetNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId, configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.GetNodeBalancerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeBalancerConfig`: GetNodeBalancerConfigs200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.GetNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetNodeBalancerConfigs200ResponseDataInner**](GetNodeBalancerConfigs200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerConfigNodes

> GetNodeBalancerConfigNodes200Response GetNodeBalancerConfigNodes(ctx, apiVersion, nodeBalancerId, configId).Page(page).PageSize(pageSize).Execute()

List nodes



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	configId := int32(56) // int32 | The ID of the NodeBalancer config to access.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.GetNodeBalancerConfigNodes(context.Background(), apiVersion, nodeBalancerId, configId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.GetNodeBalancerConfigNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeBalancerConfigNodes`: GetNodeBalancerConfigNodes200Response
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.GetNodeBalancerConfigNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the NodeBalancer config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerConfigNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetNodeBalancerConfigNodes200Response**](GetNodeBalancerConfigNodes200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerConfigs

> GetNodeBalancerConfigs200Response GetNodeBalancerConfigs(ctx, apiVersion, nodeBalancerId).Page(page).PageSize(pageSize).Execute()

List configs



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.GetNodeBalancerConfigs(context.Background(), apiVersion, nodeBalancerId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.GetNodeBalancerConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeBalancerConfigs`: GetNodeBalancerConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.GetNodeBalancerConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetNodeBalancerConfigs200Response**](GetNodeBalancerConfigs200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerFirewalls

> GetNodeBalancerFirewalls200Response GetNodeBalancerFirewalls(ctx, apiVersion, nodeBalancerId).Execute()

List NodeBalancer firewalls



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.GetNodeBalancerFirewalls(context.Background(), apiVersion, nodeBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.GetNodeBalancerFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeBalancerFirewalls`: GetNodeBalancerFirewalls200Response
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.GetNodeBalancerFirewalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNodeBalancerFirewalls200Response**](GetNodeBalancerFirewalls200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerNode

> PostNodeBalancerRequestConfigsInnerNodesInner GetNodeBalancerNode(ctx, apiVersion, nodeBalancerId, configId, nodeId).Execute()

Get a NodeBalancer's node



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	configId := int32(56) // int32 | The ID of the Config to access.
	nodeId := "nodeId_example" // string | The ID of the Node to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.GetNodeBalancerNode(context.Background(), apiVersion, nodeBalancerId, configId, nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.GetNodeBalancerNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeBalancerNode`: PostNodeBalancerRequestConfigsInnerNodesInner
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.GetNodeBalancerNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access. | 
**nodeId** | **string** | The ID of the Node to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PostNodeBalancerRequestConfigsInnerNodesInner**](PostNodeBalancerRequestConfigsInnerNodesInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerStats

> GetNodeBalancerStats200Response GetNodeBalancerStats(ctx, apiVersion, nodeBalancerId).Execute()

Get NodeBalancer statistics



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.GetNodeBalancerStats(context.Background(), apiVersion, nodeBalancerId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.GetNodeBalancerStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeBalancerStats`: GetNodeBalancerStats200Response
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.GetNodeBalancerStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNodeBalancerStats200Response**](GetNodeBalancerStats200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancerTypes

> GetNodeBalancerTypes200Response GetNodeBalancerTypes(ctx, apiVersion).Execute()

List NodeBalancer types



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.GetNodeBalancerTypes(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.GetNodeBalancerTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeBalancerTypes`: GetNodeBalancerTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.GetNodeBalancerTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancerTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetNodeBalancerTypes200Response**](GetNodeBalancerTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeBalancers

> GetLinodeNodeBalancers200Response GetNodeBalancers(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List NodeBalancers



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
	resp, r, err := apiClient.NodeBalancersAPI.GetNodeBalancers(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.GetNodeBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeBalancers`: GetLinodeNodeBalancers200Response
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.GetNodeBalancers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetLinodeNodeBalancers200Response**](GetLinodeNodeBalancers200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNodeBalancer

> NodeBalancer PostNodeBalancer(ctx, apiVersion).PostNodeBalancerRequest(postNodeBalancerRequest).Execute()

Create a NodeBalancer



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
	postNodeBalancerRequest := *openapiclient.NewPostNodeBalancerRequest("us-east") // PostNodeBalancerRequest | Information about the NodeBalancer to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.PostNodeBalancer(context.Background(), apiVersion).PostNodeBalancerRequest(postNodeBalancerRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.PostNodeBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNodeBalancer`: NodeBalancer
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.PostNodeBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostNodeBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postNodeBalancerRequest** | [**PostNodeBalancerRequest**](PostNodeBalancerRequest.md) | Information about the NodeBalancer to create. | 

### Return type

[**NodeBalancer**](NodeBalancer.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNodeBalancerConfig

> GetNodeBalancerConfigs200ResponseDataInner PostNodeBalancerConfig(ctx, apiVersion, nodeBalancerId).GetNodeBalancerConfigs200ResponseDataInner(getNodeBalancerConfigs200ResponseDataInner).Execute()

Create a config



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	getNodeBalancerConfigs200ResponseDataInner := *openapiclient.NewGetNodeBalancerConfigs200ResponseDataInner() // GetNodeBalancerConfigs200ResponseDataInner | Information about the port to configure. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.PostNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId).GetNodeBalancerConfigs200ResponseDataInner(getNodeBalancerConfigs200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.PostNodeBalancerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNodeBalancerConfig`: GetNodeBalancerConfigs200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.PostNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNodeBalancerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getNodeBalancerConfigs200ResponseDataInner** | [**GetNodeBalancerConfigs200ResponseDataInner**](GetNodeBalancerConfigs200ResponseDataInner.md) | Information about the port to configure. | 

### Return type

[**GetNodeBalancerConfigs200ResponseDataInner**](GetNodeBalancerConfigs200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostNodeBalancerNode

> PostNodeBalancerRequestConfigsInnerNodesInner PostNodeBalancerNode(ctx, apiVersion, nodeBalancerId, configId).PostNodeBalancerNodeRequest(postNodeBalancerNodeRequest).Execute()

Create a node



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	configId := int32(56) // int32 | The ID of the NodeBalancer config to access.
	postNodeBalancerNodeRequest := *openapiclient.NewPostNodeBalancerNodeRequest("192.168.210.120:80", "node54321") // PostNodeBalancerNodeRequest | Information about the Node to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.PostNodeBalancerNode(context.Background(), apiVersion, nodeBalancerId, configId).PostNodeBalancerNodeRequest(postNodeBalancerNodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.PostNodeBalancerNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostNodeBalancerNode`: PostNodeBalancerRequestConfigsInnerNodesInner
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.PostNodeBalancerNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the NodeBalancer config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostNodeBalancerNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postNodeBalancerNodeRequest** | [**PostNodeBalancerNodeRequest**](PostNodeBalancerNodeRequest.md) | Information about the Node to create. | 

### Return type

[**PostNodeBalancerRequestConfigsInnerNodesInner**](PostNodeBalancerRequestConfigsInnerNodesInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRebuildNodeBalancerConfig

> GetNodeBalancerConfigs200ResponseDataInner PostRebuildNodeBalancerConfig(ctx, apiVersion, nodeBalancerId, configId).PostRebuildNodeBalancerConfigRequest(postRebuildNodeBalancerConfigRequest).Execute()

Rebuild a config



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	configId := int32(56) // int32 | The ID of the Config to access.
	postRebuildNodeBalancerConfigRequest := *openapiclient.NewPostRebuildNodeBalancerConfigRequest([]openapiclient.PostRebuildNodeBalancerConfigRequestAllOfNodesInner{*openapiclient.NewPostRebuildNodeBalancerConfigRequestAllOfNodesInner()}) // PostRebuildNodeBalancerConfigRequest | Information about the NodeBalancer Config to rebuild.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.PostRebuildNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId, configId).PostRebuildNodeBalancerConfigRequest(postRebuildNodeBalancerConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.PostRebuildNodeBalancerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRebuildNodeBalancerConfig`: GetNodeBalancerConfigs200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.PostRebuildNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRebuildNodeBalancerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postRebuildNodeBalancerConfigRequest** | [**PostRebuildNodeBalancerConfigRequest**](PostRebuildNodeBalancerConfigRequest.md) | Information about the NodeBalancer Config to rebuild. | 

### Return type

[**GetNodeBalancerConfigs200ResponseDataInner**](GetNodeBalancerConfigs200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNodeBalancer

> NodeBalancer PutNodeBalancer(ctx, apiVersion, nodeBalancerId).NodeBalancer(nodeBalancer).Execute()

Update a NodeBalancer



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	nodeBalancer := *openapiclient.NewNodeBalancer() // NodeBalancer | The information to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.PutNodeBalancer(context.Background(), apiVersion, nodeBalancerId).NodeBalancer(nodeBalancer).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.PutNodeBalancer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNodeBalancer`: NodeBalancer
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.PutNodeBalancer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNodeBalancerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nodeBalancer** | [**NodeBalancer**](NodeBalancer.md) | The information to update. | 

### Return type

[**NodeBalancer**](NodeBalancer.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNodeBalancerConfig

> GetNodeBalancerConfigs200ResponseDataInner PutNodeBalancerConfig(ctx, apiVersion, nodeBalancerId, configId).GetNodeBalancerConfigs200ResponseDataInner(getNodeBalancerConfigs200ResponseDataInner).Execute()

Update a config



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	configId := int32(56) // int32 | The ID of the Config to access.
	getNodeBalancerConfigs200ResponseDataInner := *openapiclient.NewGetNodeBalancerConfigs200ResponseDataInner() // GetNodeBalancerConfigs200ResponseDataInner | The fields to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.PutNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId, configId).GetNodeBalancerConfigs200ResponseDataInner(getNodeBalancerConfigs200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.PutNodeBalancerConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNodeBalancerConfig`: GetNodeBalancerConfigs200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.PutNodeBalancerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNodeBalancerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **getNodeBalancerConfigs200ResponseDataInner** | [**GetNodeBalancerConfigs200ResponseDataInner**](GetNodeBalancerConfigs200ResponseDataInner.md) | The fields to update. | 

### Return type

[**GetNodeBalancerConfigs200ResponseDataInner**](GetNodeBalancerConfigs200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutNodeBalancerNode

> PostNodeBalancerRequestConfigsInnerNodesInner PutNodeBalancerNode(ctx, apiVersion, nodeBalancerId, configId, nodeId).PostNodeBalancerRequestConfigsInnerNodesInner(postNodeBalancerRequestConfigsInnerNodesInner).Execute()

Update a node



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
	nodeBalancerId := int32(56) // int32 | The ID of the NodeBalancer to access.
	configId := int32(56) // int32 | The ID of the Config to access.
	nodeId := "nodeId_example" // string | The ID of the Node to access.
	postNodeBalancerRequestConfigsInnerNodesInner := *openapiclient.NewPostNodeBalancerRequestConfigsInnerNodesInner() // PostNodeBalancerRequestConfigsInnerNodesInner | The fields to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NodeBalancersAPI.PutNodeBalancerNode(context.Background(), apiVersion, nodeBalancerId, configId, nodeId).PostNodeBalancerRequestConfigsInnerNodesInner(postNodeBalancerRequestConfigsInnerNodesInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NodeBalancersAPI.PutNodeBalancerNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutNodeBalancerNode`: PostNodeBalancerRequestConfigsInnerNodesInner
	fmt.Fprintf(os.Stdout, "Response from `NodeBalancersAPI.PutNodeBalancerNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**nodeBalancerId** | **int32** | The ID of the NodeBalancer to access. | 
**configId** | **int32** | The ID of the Config to access. | 
**nodeId** | **string** | The ID of the Node to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutNodeBalancerNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **postNodeBalancerRequestConfigsInnerNodesInner** | [**PostNodeBalancerRequestConfigsInnerNodesInner**](PostNodeBalancerRequestConfigsInnerNodesInner.md) | The fields to update. | 

### Return type

[**PostNodeBalancerRequestConfigsInnerNodesInner**](PostNodeBalancerRequestConfigsInnerNodesInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

