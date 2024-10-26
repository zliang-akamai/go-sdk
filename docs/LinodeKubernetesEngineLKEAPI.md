# \LinodeKubernetesEngineLKEAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLkeCluster**](LinodeKubernetesEngineLKEAPI.md#DeleteLkeCluster) | **Delete** /{apiVersion}/lke/clusters/{clusterId} | Delete a Kubernetes cluster
[**DeleteLkeClusterAcl**](LinodeKubernetesEngineLKEAPI.md#DeleteLkeClusterAcl) | **Delete** /{apiVersion}/lke/clusters/{clusterId}/control_plane_acl | Delete the control plane access control list
[**DeleteLkeClusterKubeconfig**](LinodeKubernetesEngineLKEAPI.md#DeleteLkeClusterKubeconfig) | **Delete** /{apiVersion}/lke/clusters/{clusterId}/kubeconfig | Delete a Kubeconfig
[**DeleteLkeClusterNode**](LinodeKubernetesEngineLKEAPI.md#DeleteLkeClusterNode) | **Delete** /{apiVersion}/lke/clusters/{clusterId}/nodes/{nodeId} | Delete a node
[**DeleteLkeNodePool**](LinodeKubernetesEngineLKEAPI.md#DeleteLkeNodePool) | **Delete** /{apiVersion}/lke/clusters/{clusterId}/pools/{poolId} | Delete a node pool
[**DeleteLkeServiceToken**](LinodeKubernetesEngineLKEAPI.md#DeleteLkeServiceToken) | **Delete** /{apiVersion}/lke/clusters/{clusterId}/servicetoken | Delete a service token
[**GetLkeCluster**](LinodeKubernetesEngineLKEAPI.md#GetLkeCluster) | **Get** /{apiVersion}/lke/clusters/{clusterId} | Get a Kubernetes cluster
[**GetLkeClusterAcl**](LinodeKubernetesEngineLKEAPI.md#GetLkeClusterAcl) | **Get** /{apiVersion}/lke/clusters/{clusterId}/control_plane_acl | Get the control plane access control list
[**GetLkeClusterApiEndpoints**](LinodeKubernetesEngineLKEAPI.md#GetLkeClusterApiEndpoints) | **Get** /{apiVersion}/lke/clusters/{clusterId}/api-endpoints | List Kubernetes API endpoints
[**GetLkeClusterDashboard**](LinodeKubernetesEngineLKEAPI.md#GetLkeClusterDashboard) | **Get** /{apiVersion}/lke/clusters/{clusterId}/dashboard | Get a Kubernetes cluster dashboard URL
[**GetLkeClusterKubeconfig**](LinodeKubernetesEngineLKEAPI.md#GetLkeClusterKubeconfig) | **Get** /{apiVersion}/lke/clusters/{clusterId}/kubeconfig | Get a Kubeconfig
[**GetLkeClusterNode**](LinodeKubernetesEngineLKEAPI.md#GetLkeClusterNode) | **Get** /{apiVersion}/lke/clusters/{clusterId}/nodes/{nodeId} | Get a node
[**GetLkeClusterPools**](LinodeKubernetesEngineLKEAPI.md#GetLkeClusterPools) | **Get** /{apiVersion}/lke/clusters/{clusterId}/pools | List node pools
[**GetLkeClusters**](LinodeKubernetesEngineLKEAPI.md#GetLkeClusters) | **Get** /{apiVersion}/lke/clusters | List Kubernetes clusters
[**GetLkeNodePool**](LinodeKubernetesEngineLKEAPI.md#GetLkeNodePool) | **Get** /{apiVersion}/lke/clusters/{clusterId}/pools/{poolId} | Get a node pool
[**GetLkeTypes**](LinodeKubernetesEngineLKEAPI.md#GetLkeTypes) | **Get** /{apiVersion}/lke/types | List Kubernetes types
[**GetLkeVersion**](LinodeKubernetesEngineLKEAPI.md#GetLkeVersion) | **Get** /{apiVersion}/lke/versions/{version} | Get a Kubernetes version
[**GetLkeVersions**](LinodeKubernetesEngineLKEAPI.md#GetLkeVersions) | **Get** /{apiVersion}/lke/versions | List Kubernetes versions
[**PostLkeCluster**](LinodeKubernetesEngineLKEAPI.md#PostLkeCluster) | **Post** /{apiVersion}/lke/clusters | Create a Kubernetes cluster
[**PostLkeClusterNodeRecycle**](LinodeKubernetesEngineLKEAPI.md#PostLkeClusterNodeRecycle) | **Post** /{apiVersion}/lke/clusters/{clusterId}/nodes/{nodeId}/recycle | Recycle a node
[**PostLkeClusterPoolRecycle**](LinodeKubernetesEngineLKEAPI.md#PostLkeClusterPoolRecycle) | **Post** /{apiVersion}/lke/clusters/{clusterId}/pools/{poolId}/recycle | Recycle a node pool
[**PostLkeClusterPools**](LinodeKubernetesEngineLKEAPI.md#PostLkeClusterPools) | **Post** /{apiVersion}/lke/clusters/{clusterId}/pools | Create a node pool
[**PostLkeClusterRecycle**](LinodeKubernetesEngineLKEAPI.md#PostLkeClusterRecycle) | **Post** /{apiVersion}/lke/clusters/{clusterId}/recycle | Recycle cluster nodes
[**PostLkeClusterRegenerate**](LinodeKubernetesEngineLKEAPI.md#PostLkeClusterRegenerate) | **Post** /{apiVersion}/lke/clusters/{clusterId}/regenerate | Regenerate a Kubernetes cluster
[**PutLkeCluster**](LinodeKubernetesEngineLKEAPI.md#PutLkeCluster) | **Put** /{apiVersion}/lke/clusters/{clusterId} | Update a Kubernetes cluster
[**PutLkeClusterAcl**](LinodeKubernetesEngineLKEAPI.md#PutLkeClusterAcl) | **Put** /{apiVersion}/lke/clusters/{clusterId}/control_plane_acl | Update the control plane access control list
[**PutLkeNodePool**](LinodeKubernetesEngineLKEAPI.md#PutLkeNodePool) | **Put** /{apiVersion}/lke/clusters/{clusterId}/pools/{poolId} | Update a node pool



## DeleteLkeCluster

> map[string]interface{} DeleteLkeCluster(ctx, apiVersion, clusterId).Execute()

Delete a Kubernetes cluster



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeCluster(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.DeleteLkeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLkeCluster`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.DeleteLkeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLkeClusterRequest struct via the builder pattern


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


## DeleteLkeClusterAcl

> map[string]interface{} DeleteLkeClusterAcl(ctx, apiVersion, clusterId).Execute()

Delete the control plane access control list



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeClusterAcl(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.DeleteLkeClusterAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLkeClusterAcl`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.DeleteLkeClusterAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLkeClusterAclRequest struct via the builder pattern


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


## DeleteLkeClusterKubeconfig

> map[string]interface{} DeleteLkeClusterKubeconfig(ctx, apiVersion, clusterId).Execute()

Delete a Kubeconfig



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeClusterKubeconfig(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.DeleteLkeClusterKubeconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLkeClusterKubeconfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.DeleteLkeClusterKubeconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLkeClusterKubeconfigRequest struct via the builder pattern


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


## DeleteLkeClusterNode

> map[string]interface{} DeleteLkeClusterNode(ctx, apiVersion, clusterId, nodeId).Execute()

Delete a node



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster containing the Node.
	nodeId := "nodeId_example" // string | The ID of the Node to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeClusterNode(context.Background(), apiVersion, clusterId, nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.DeleteLkeClusterNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLkeClusterNode`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.DeleteLkeClusterNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster containing the Node. | 
**nodeId** | **string** | The ID of the Node to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLkeClusterNodeRequest struct via the builder pattern


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


## DeleteLkeNodePool

> map[string]interface{} DeleteLkeNodePool(ctx, apiVersion, clusterId, poolId).Execute()

Delete a node pool



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
	poolId := int32(56) // int32 | ID of the Pool to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeNodePool(context.Background(), apiVersion, clusterId, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.DeleteLkeNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLkeNodePool`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.DeleteLkeNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 
**poolId** | **int32** | ID of the Pool to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLkeNodePoolRequest struct via the builder pattern


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


## DeleteLkeServiceToken

> map[string]interface{} DeleteLkeServiceToken(ctx, apiVersion, clusterId).Execute()

Delete a service token



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
	clusterId := int32(56) // int32 | ID of the target Kubernetes cluster.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeServiceToken(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.DeleteLkeServiceToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLkeServiceToken`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.DeleteLkeServiceToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the target Kubernetes cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLkeServiceTokenRequest struct via the builder pattern


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


## GetLkeCluster

> GetLkeClusters200ResponseDataInner GetLkeCluster(ctx, apiVersion, clusterId).Execute()

Get a Kubernetes cluster



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeCluster(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeCluster`: GetLkeClusters200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLkeClusters200ResponseDataInner**](GetLkeClusters200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeClusterAcl

> GetLkeClusterAcl200Response GetLkeClusterAcl(ctx, apiVersion, clusterId).Execute()

Get the control plane access control list



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterAcl(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeClusterAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeClusterAcl`: GetLkeClusterAcl200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeClusterAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeClusterAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLkeClusterAcl200Response**](GetLkeClusterAcl200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeClusterApiEndpoints

> GetLkeClusterApiEndpoints200Response GetLkeClusterApiEndpoints(ctx, apiVersion, clusterId).Execute()

List Kubernetes API endpoints



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterApiEndpoints(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeClusterApiEndpoints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeClusterApiEndpoints`: GetLkeClusterApiEndpoints200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeClusterApiEndpoints`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeClusterApiEndpointsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLkeClusterApiEndpoints200Response**](GetLkeClusterApiEndpoints200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeClusterDashboard

> GetLkeClusterDashboard200Response GetLkeClusterDashboard(ctx, apiVersion, clusterId).Execute()

Get a Kubernetes cluster dashboard URL



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterDashboard(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeClusterDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeClusterDashboard`: GetLkeClusterDashboard200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeClusterDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeClusterDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLkeClusterDashboard200Response**](GetLkeClusterDashboard200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeClusterKubeconfig

> GetLkeClusterKubeconfig200Response GetLkeClusterKubeconfig(ctx, apiVersion, clusterId).Execute()

Get a Kubeconfig



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterKubeconfig(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeClusterKubeconfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeClusterKubeconfig`: GetLkeClusterKubeconfig200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeClusterKubeconfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeClusterKubeconfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLkeClusterKubeconfig200Response**](GetLkeClusterKubeconfig200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeClusterNode

> GetLkeClusterNode200Response GetLkeClusterNode(ctx, apiVersion, clusterId, nodeId).Execute()

Get a node



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster containing the Node.
	nodeId := "nodeId_example" // string | The ID of the Node to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterNode(context.Background(), apiVersion, clusterId, nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeClusterNode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeClusterNode`: GetLkeClusterNode200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeClusterNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster containing the Node. | 
**nodeId** | **string** | The ID of the Node to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeClusterNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetLkeClusterNode200Response**](GetLkeClusterNode200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeClusterPools

> GetLkeClusterPools200Response GetLkeClusterPools(ctx, apiVersion, clusterId).Execute()

List node pools



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterPools(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeClusterPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeClusterPools`: GetLkeClusterPools200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeClusterPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeClusterPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLkeClusterPools200Response**](GetLkeClusterPools200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeClusters

> GetLkeClusters200Response GetLkeClusters(ctx, apiVersion).Execute()

List Kubernetes clusters



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
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusters(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeClusters`: GetLkeClusters200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLkeClusters200Response**](GetLkeClusters200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeNodePool

> GetLkeClusterPools200ResponseDataInner GetLkeNodePool(ctx, apiVersion, clusterId, poolId).Execute()

Get a node pool



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
	poolId := int32(56) // int32 | ID of the Pool to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeNodePool(context.Background(), apiVersion, clusterId, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeNodePool`: GetLkeClusterPools200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 
**poolId** | **int32** | ID of the Pool to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetLkeClusterPools200ResponseDataInner**](GetLkeClusterPools200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeTypes

> GetLkeTypes200Response GetLkeTypes(ctx, apiVersion).Execute()

List Kubernetes types



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
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeTypes(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeTypes`: GetLkeTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLkeTypes200Response**](GetLkeTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeVersion

> GetLkeVersions200ResponseDataInner GetLkeVersion(ctx, apiVersion, version).Execute()

Get a Kubernetes version



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
	version := "version_example" // string | The LKE version to view.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeVersion(context.Background(), apiVersion, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeVersion`: GetLkeVersions200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**version** | **string** | The LKE version to view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLkeVersions200ResponseDataInner**](GetLkeVersions200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLkeVersions

> GetLkeVersions200Response GetLkeVersions(ctx, apiVersion).Execute()

List Kubernetes versions



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
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeVersions(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.GetLkeVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLkeVersions`: GetLkeVersions200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.GetLkeVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLkeVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLkeVersions200Response**](GetLkeVersions200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLkeCluster

> PostLkeCluster200Response PostLkeCluster(ctx, apiVersion).PostLkeClusterRequest(postLkeClusterRequest).Execute()

Create a Kubernetes cluster



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
	postLkeClusterRequest := *openapiclient.NewPostLkeClusterRequest("1.27", "lkecluster12345", []openapiclient.PostLkeClusterRequestNodePoolsInner{*openapiclient.NewPostLkeClusterRequestNodePoolsInner(int32(6), "g6-standard-4")}, "us-central") // PostLkeClusterRequest | Configuration for the Kubernetes cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeCluster(context.Background(), apiVersion).PostLkeClusterRequest(postLkeClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.PostLkeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLkeCluster`: PostLkeCluster200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.PostLkeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostLkeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postLkeClusterRequest** | [**PostLkeClusterRequest**](PostLkeClusterRequest.md) | Configuration for the Kubernetes cluster. | 

### Return type

[**PostLkeCluster200Response**](PostLkeCluster200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLkeClusterNodeRecycle

> map[string]interface{} PostLkeClusterNodeRecycle(ctx, apiVersion, clusterId, nodeId).Execute()

Recycle a node



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster containing the Node.
	nodeId := "nodeId_example" // string | ID of the Node to be recycled.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterNodeRecycle(context.Background(), apiVersion, clusterId, nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.PostLkeClusterNodeRecycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLkeClusterNodeRecycle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.PostLkeClusterNodeRecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster containing the Node. | 
**nodeId** | **string** | ID of the Node to be recycled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLkeClusterNodeRecycleRequest struct via the builder pattern


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


## PostLkeClusterPoolRecycle

> map[string]interface{} PostLkeClusterPoolRecycle(ctx, apiVersion, clusterId, poolId).Execute()

Recycle a node pool



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster this Node Pool is attached to.
	poolId := int32(56) // int32 | ID of the Node Pool to be recycled.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterPoolRecycle(context.Background(), apiVersion, clusterId, poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.PostLkeClusterPoolRecycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLkeClusterPoolRecycle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.PostLkeClusterPoolRecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster this Node Pool is attached to. | 
**poolId** | **int32** | ID of the Node Pool to be recycled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLkeClusterPoolRecycleRequest struct via the builder pattern


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


## PostLkeClusterPools

> GetLkeClusterPools200ResponseDataInner PostLkeClusterPools(ctx, apiVersion, clusterId).PostLkeClusterPoolsRequest(postLkeClusterPoolsRequest).Execute()

Create a node pool



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
	postLkeClusterPoolsRequest := *openapiclient.NewPostLkeClusterPoolsRequest(int32(6), "g6-standard-4") // PostLkeClusterPoolsRequest | Configuration for the Node Pool.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterPools(context.Background(), apiVersion, clusterId).PostLkeClusterPoolsRequest(postLkeClusterPoolsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.PostLkeClusterPools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLkeClusterPools`: GetLkeClusterPools200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.PostLkeClusterPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLkeClusterPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postLkeClusterPoolsRequest** | [**PostLkeClusterPoolsRequest**](PostLkeClusterPoolsRequest.md) | Configuration for the Node Pool. | 

### Return type

[**GetLkeClusterPools200ResponseDataInner**](GetLkeClusterPools200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLkeClusterRecycle

> map[string]interface{} PostLkeClusterRecycle(ctx, apiVersion, clusterId).Execute()

Recycle cluster nodes



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster which contains nodes to be recycled.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterRecycle(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.PostLkeClusterRecycle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLkeClusterRecycle`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.PostLkeClusterRecycle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster which contains nodes to be recycled. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLkeClusterRecycleRequest struct via the builder pattern


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


## PostLkeClusterRegenerate

> map[string]interface{} PostLkeClusterRegenerate(ctx, apiVersion, clusterId).PostLkeClusterRegenerateRequest(postLkeClusterRegenerateRequest).Execute()

Regenerate a Kubernetes cluster



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
	clusterId := int32(56) // int32 | ID of the target Kubernetes cluster.
	postLkeClusterRegenerateRequest := *openapiclient.NewPostLkeClusterRegenerateRequest() // PostLkeClusterRegenerateRequest | The Kubernetes Cluster Regenerate request object. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterRegenerate(context.Background(), apiVersion, clusterId).PostLkeClusterRegenerateRequest(postLkeClusterRegenerateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.PostLkeClusterRegenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLkeClusterRegenerate`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.PostLkeClusterRegenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the target Kubernetes cluster. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLkeClusterRegenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postLkeClusterRegenerateRequest** | [**PostLkeClusterRegenerateRequest**](PostLkeClusterRegenerateRequest.md) | The Kubernetes Cluster Regenerate request object. | 

### Return type

**map[string]interface{}**

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLkeCluster

> PostLkeCluster200Response PutLkeCluster(ctx, apiVersion, clusterId).PutLkeClusterRequest(putLkeClusterRequest).Execute()

Update a Kubernetes cluster



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
	putLkeClusterRequest := *openapiclient.NewPutLkeClusterRequest() // PutLkeClusterRequest | The fields to update the Kubernetes cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.PutLkeCluster(context.Background(), apiVersion, clusterId).PutLkeClusterRequest(putLkeClusterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.PutLkeCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLkeCluster`: PostLkeCluster200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.PutLkeCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLkeClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putLkeClusterRequest** | [**PutLkeClusterRequest**](PutLkeClusterRequest.md) | The fields to update the Kubernetes cluster. | 

### Return type

[**PostLkeCluster200Response**](PostLkeCluster200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLkeClusterAcl

> GetLkeClusterAcl200Response PutLkeClusterAcl(ctx, apiVersion, clusterId).PutLkeClusterAclRequest(putLkeClusterAclRequest).Execute()

Update the control plane access control list



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
	putLkeClusterAclRequest := *openapiclient.NewPutLkeClusterAclRequest() // PutLkeClusterAclRequest | The fields to update the cluster. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.PutLkeClusterAcl(context.Background(), apiVersion, clusterId).PutLkeClusterAclRequest(putLkeClusterAclRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.PutLkeClusterAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLkeClusterAcl`: GetLkeClusterAcl200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.PutLkeClusterAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLkeClusterAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putLkeClusterAclRequest** | [**PutLkeClusterAclRequest**](PutLkeClusterAclRequest.md) | The fields to update the cluster. | 

### Return type

[**GetLkeClusterAcl200Response**](GetLkeClusterAcl200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLkeNodePool

> GetLkeClusterPools200ResponseDataInner PutLkeNodePool(ctx, apiVersion, clusterId, poolId).PutLkeNodePoolRequest(putLkeNodePoolRequest).Execute()

Update a node pool



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
	clusterId := int32(56) // int32 | ID of the Kubernetes cluster to look up.
	poolId := int32(56) // int32 | ID of the Pool to look up.
	putLkeNodePoolRequest := *openapiclient.NewPutLkeNodePoolRequest() // PutLkeNodePoolRequest | The fields to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeKubernetesEngineLKEAPI.PutLkeNodePool(context.Background(), apiVersion, clusterId, poolId).PutLkeNodePoolRequest(putLkeNodePoolRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeKubernetesEngineLKEAPI.PutLkeNodePool``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLkeNodePool`: GetLkeClusterPools200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeKubernetesEngineLKEAPI.PutLkeNodePool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **int32** | ID of the Kubernetes cluster to look up. | 
**poolId** | **int32** | ID of the Pool to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLkeNodePoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putLkeNodePoolRequest** | [**PutLkeNodePoolRequest**](PutLkeNodePoolRequest.md) | The fields to update. | 

### Return type

[**GetLkeClusterPools200ResponseDataInner**](GetLkeClusterPools200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

