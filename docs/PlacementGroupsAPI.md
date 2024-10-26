# \PlacementGroupsAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePlacementGroup**](PlacementGroupsAPI.md#DeletePlacementGroup) | **Delete** /{apiVersion}/placement/groups/{groupId} | Delete a placement group
[**GetPlacementGroup**](PlacementGroupsAPI.md#GetPlacementGroup) | **Get** /{apiVersion}/placement/groups/{groupId} | Get a placement group
[**GetPlacementGroups**](PlacementGroupsAPI.md#GetPlacementGroups) | **Get** /{apiVersion}/placement/groups | List placement groups
[**PostGroupAddLinode**](PlacementGroupsAPI.md#PostGroupAddLinode) | **Post** /{apiVersion}/placement/groups/{groupId}/assign | Assign a placement group
[**PostGroupUnassign**](PlacementGroupsAPI.md#PostGroupUnassign) | **Post** /{apiVersion}/placement/groups/{groupId}/unassign | Unassign a placement group
[**PostPlacementGroup**](PlacementGroupsAPI.md#PostPlacementGroup) | **Post** /{apiVersion}/placement/groups | Create placement group
[**PutPlacementGroup**](PlacementGroupsAPI.md#PutPlacementGroup) | **Put** /{apiVersion}/placement/groups/{groupId} | Update a placement group



## DeletePlacementGroup

> map[string]interface{} DeletePlacementGroup(ctx, apiVersion, groupId).Execute()

Delete a placement group



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
	groupId := int32(56) // int32 | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the `id` for the applicable placement group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementGroupsAPI.DeletePlacementGroup(context.Background(), apiVersion, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementGroupsAPI.DeletePlacementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePlacementGroup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlacementGroupsAPI.DeletePlacementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**groupId** | **int32** | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the &#x60;id&#x60; for the applicable placement group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlacementGroupRequest struct via the builder pattern


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


## GetPlacementGroup

> GetPlacementGroups200ResponseDataInner GetPlacementGroup(ctx, apiVersion, groupId).Execute()

Get a placement group



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
	groupId := int32(56) // int32 | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the `id` for the applicable placement group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementGroupsAPI.GetPlacementGroup(context.Background(), apiVersion, groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementGroupsAPI.GetPlacementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlacementGroup`: GetPlacementGroups200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `PlacementGroupsAPI.GetPlacementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**groupId** | **int32** | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the &#x60;id&#x60; for the applicable placement group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlacementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPlacementGroups200ResponseDataInner**](GetPlacementGroups200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlacementGroups

> GetPlacementGroups200Response GetPlacementGroups(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List placement groups



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
	resp, r, err := apiClient.PlacementGroupsAPI.GetPlacementGroups(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementGroupsAPI.GetPlacementGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlacementGroups`: GetPlacementGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `PlacementGroupsAPI.GetPlacementGroups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlacementGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetPlacementGroups200Response**](GetPlacementGroups200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGroupAddLinode

> PostPlacementGroup200Response PostGroupAddLinode(ctx, apiVersion, groupId).PostGroupAddLinodeRequest(postGroupAddLinodeRequest).Execute()

Assign a placement group



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
	groupId := int32(56) // int32 | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the `id` for the applicable placement group.
	postGroupAddLinodeRequest := *openapiclient.NewPostGroupAddLinodeRequest() // PostGroupAddLinodeRequest | The compute instances you want to add to your placement group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementGroupsAPI.PostGroupAddLinode(context.Background(), apiVersion, groupId).PostGroupAddLinodeRequest(postGroupAddLinodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementGroupsAPI.PostGroupAddLinode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGroupAddLinode`: PostPlacementGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `PlacementGroupsAPI.PostGroupAddLinode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**groupId** | **int32** | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the &#x60;id&#x60; for the applicable placement group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGroupAddLinodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postGroupAddLinodeRequest** | [**PostGroupAddLinodeRequest**](PostGroupAddLinodeRequest.md) | The compute instances you want to add to your placement group. | 

### Return type

[**PostPlacementGroup200Response**](PostPlacementGroup200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostGroupUnassign

> PostPlacementGroup200Response PostGroupUnassign(ctx, apiVersion, groupId).PostGroupAddLinodeRequest(postGroupAddLinodeRequest).Execute()

Unassign a placement group



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
	groupId := int32(56) // int32 | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the `id` for the applicable placement group.
	postGroupAddLinodeRequest := *openapiclient.NewPostGroupAddLinodeRequest() // PostGroupAddLinodeRequest | The compute instances you want to remove from your placement group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementGroupsAPI.PostGroupUnassign(context.Background(), apiVersion, groupId).PostGroupAddLinodeRequest(postGroupAddLinodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementGroupsAPI.PostGroupUnassign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostGroupUnassign`: PostPlacementGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `PlacementGroupsAPI.PostGroupUnassign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**groupId** | **int32** | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the &#x60;id&#x60; for the applicable placement group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostGroupUnassignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postGroupAddLinodeRequest** | [**PostGroupAddLinodeRequest**](PostGroupAddLinodeRequest.md) | The compute instances you want to remove from your placement group. | 

### Return type

[**PostPlacementGroup200Response**](PostPlacementGroup200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostPlacementGroup

> PostPlacementGroup200Response PostPlacementGroup(ctx, apiVersion).PostPlacementGroupRequest(postPlacementGroupRequest).Execute()

Create placement group



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
	postPlacementGroupRequest := *openapiclient.NewPostPlacementGroupRequest("PG_Miami_failover", "strict", "["anti_affinity:local"]", "us-iad") // PostPlacementGroupRequest | The requested initial state of the new placement group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementGroupsAPI.PostPlacementGroup(context.Background(), apiVersion).PostPlacementGroupRequest(postPlacementGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementGroupsAPI.PostPlacementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostPlacementGroup`: PostPlacementGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `PlacementGroupsAPI.PostPlacementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostPlacementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postPlacementGroupRequest** | [**PostPlacementGroupRequest**](PostPlacementGroupRequest.md) | The requested initial state of the new placement group. | 

### Return type

[**PostPlacementGroup200Response**](PostPlacementGroup200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPlacementGroup

> PostPlacementGroup200Response PutPlacementGroup(ctx, apiVersion, groupId).PutPlacementGroupRequest(putPlacementGroupRequest).Execute()

Update a placement group



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
	groupId := int32(56) // int32 | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the `id` for the applicable placement group.
	putPlacementGroupRequest := *openapiclient.NewPutPlacementGroupRequest() // PutPlacementGroupRequest | Include the `label` parameter to update the name of your placement group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlacementGroupsAPI.PutPlacementGroup(context.Background(), apiVersion, groupId).PutPlacementGroupRequest(putPlacementGroupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlacementGroupsAPI.PutPlacementGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPlacementGroup`: PostPlacementGroup200Response
	fmt.Fprintf(os.Stdout, "Response from `PlacementGroupsAPI.PutPlacementGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**groupId** | **int32** | ID of the placement group to look up. Run the [List placement groups](https://techdocs.akamai.com/linode-api/reference/get-placement-groups) operation and store the &#x60;id&#x60; for the applicable placement group. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPlacementGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putPlacementGroupRequest** | [**PutPlacementGroupRequest**](PutPlacementGroupRequest.md) | Include the &#x60;label&#x60; parameter to update the name of your placement group. | 

### Return type

[**PostPlacementGroup200Response**](PostPlacementGroup200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

