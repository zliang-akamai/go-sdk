# \LongviewAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLongviewClient**](LongviewAPI.md#DeleteLongviewClient) | **Delete** /{apiVersion}/longview/clients/{clientId} | Delete a Longview client
[**GetLongviewClient**](LongviewAPI.md#GetLongviewClient) | **Get** /{apiVersion}/longview/clients/{clientId} | Get a Longview client
[**GetLongviewClients**](LongviewAPI.md#GetLongviewClients) | **Get** /{apiVersion}/longview/clients | List Longview clients
[**GetLongviewPlan**](LongviewAPI.md#GetLongviewPlan) | **Get** /{apiVersion}/longview/plan | Get a Longview plan
[**GetLongviewSubscription**](LongviewAPI.md#GetLongviewSubscription) | **Get** /{apiVersion}/longview/subscriptions/{subscriptionId} | Get a Longview subscription
[**GetLongviewSubscriptions**](LongviewAPI.md#GetLongviewSubscriptions) | **Get** /{apiVersion}/longview/subscriptions | List Longview subscriptions
[**GetLongviewTypes**](LongviewAPI.md#GetLongviewTypes) | **Get** /{apiVersion}/longview/types | List Longview types
[**PostLongviewClient**](LongviewAPI.md#PostLongviewClient) | **Post** /{apiVersion}/longview/clients | Create a Longview client
[**PutLongviewClient**](LongviewAPI.md#PutLongviewClient) | **Put** /{apiVersion}/longview/clients/{clientId} | Update a Longview client
[**PutLongviewPlan**](LongviewAPI.md#PutLongviewPlan) | **Put** /{apiVersion}/longview/plan | Update a Longview plan



## DeleteLongviewClient

> map[string]interface{} DeleteLongviewClient(ctx, apiVersion, clientId).Execute()

Delete a Longview client



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
	clientId := int32(56) // int32 | The Longview Client ID to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LongviewAPI.DeleteLongviewClient(context.Background(), apiVersion, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.DeleteLongviewClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLongviewClient`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.DeleteLongviewClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clientId** | **int32** | The Longview Client ID to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLongviewClientRequest struct via the builder pattern


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


## GetLongviewClient

> GetLongviewClients200ResponseDataInner GetLongviewClient(ctx, apiVersion, clientId).Execute()

Get a Longview client



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
	clientId := int32(56) // int32 | The Longview Client ID to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LongviewAPI.GetLongviewClient(context.Background(), apiVersion, clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.GetLongviewClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLongviewClient`: GetLongviewClients200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.GetLongviewClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clientId** | **int32** | The Longview Client ID to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLongviewClients200ResponseDataInner**](GetLongviewClients200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewClients

> GetLongviewClients200Response GetLongviewClients(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List Longview clients



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
	resp, r, err := apiClient.LongviewAPI.GetLongviewClients(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.GetLongviewClients``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLongviewClients`: GetLongviewClients200Response
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.GetLongviewClients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewClientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetLongviewClients200Response**](GetLongviewClients200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewPlan

> GetLongviewPlan200Response GetLongviewPlan(ctx, apiVersion).Execute()

Get a Longview plan



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
	resp, r, err := apiClient.LongviewAPI.GetLongviewPlan(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.GetLongviewPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLongviewPlan`: GetLongviewPlan200Response
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.GetLongviewPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLongviewPlan200Response**](GetLongviewPlan200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewSubscription

> GetLongviewPlan200Response GetLongviewSubscription(ctx, apiVersion, subscriptionId).Execute()

Get a Longview subscription



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
	subscriptionId := "subscriptionId_example" // string | The Longview Subscription to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LongviewAPI.GetLongviewSubscription(context.Background(), apiVersion, subscriptionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.GetLongviewSubscription``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLongviewSubscription`: GetLongviewPlan200Response
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.GetLongviewSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**subscriptionId** | **string** | The Longview Subscription to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLongviewPlan200Response**](GetLongviewPlan200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewSubscriptions

> GetLongviewSubscriptions200Response GetLongviewSubscriptions(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List Longview subscriptions



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
	resp, r, err := apiClient.LongviewAPI.GetLongviewSubscriptions(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.GetLongviewSubscriptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLongviewSubscriptions`: GetLongviewSubscriptions200Response
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.GetLongviewSubscriptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewSubscriptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetLongviewSubscriptions200Response**](GetLongviewSubscriptions200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLongviewTypes

> GetLongviewTypes200Response GetLongviewTypes(ctx, apiVersion).Execute()

List Longview types



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
	resp, r, err := apiClient.LongviewAPI.GetLongviewTypes(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.GetLongviewTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLongviewTypes`: GetLongviewTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.GetLongviewTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLongviewTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLongviewTypes200Response**](GetLongviewTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLongviewClient

> GetLongviewClients200ResponseDataInner PostLongviewClient(ctx, apiVersion).GetLongviewClients200ResponseDataInner(getLongviewClients200ResponseDataInner).Execute()

Create a Longview client



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
	getLongviewClients200ResponseDataInner := *openapiclient.NewGetLongviewClients200ResponseDataInner() // GetLongviewClients200ResponseDataInner | Information about the LongviewClient to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LongviewAPI.PostLongviewClient(context.Background(), apiVersion).GetLongviewClients200ResponseDataInner(getLongviewClients200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.PostLongviewClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLongviewClient`: GetLongviewClients200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.PostLongviewClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostLongviewClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getLongviewClients200ResponseDataInner** | [**GetLongviewClients200ResponseDataInner**](GetLongviewClients200ResponseDataInner.md) | Information about the LongviewClient to create. | 

### Return type

[**GetLongviewClients200ResponseDataInner**](GetLongviewClients200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLongviewClient

> GetLongviewClients200ResponseDataInner PutLongviewClient(ctx, apiVersion, clientId).GetLongviewClients200ResponseDataInner(getLongviewClients200ResponseDataInner).Execute()

Update a Longview client



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
	clientId := int32(56) // int32 | The Longview Client ID to access.
	getLongviewClients200ResponseDataInner := *openapiclient.NewGetLongviewClients200ResponseDataInner() // GetLongviewClients200ResponseDataInner | The fields to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LongviewAPI.PutLongviewClient(context.Background(), apiVersion, clientId).GetLongviewClients200ResponseDataInner(getLongviewClients200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.PutLongviewClient``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLongviewClient`: GetLongviewClients200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.PutLongviewClient`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clientId** | **int32** | The Longview Client ID to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLongviewClientRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getLongviewClients200ResponseDataInner** | [**GetLongviewClients200ResponseDataInner**](GetLongviewClients200ResponseDataInner.md) | The fields to update. | 

### Return type

[**GetLongviewClients200ResponseDataInner**](GetLongviewClients200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLongviewPlan

> GetLongviewPlan200Response PutLongviewPlan(ctx, apiVersion).PutLongviewPlanRequest(putLongviewPlanRequest).Execute()

Update a Longview plan



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
	putLongviewPlanRequest := *openapiclient.NewPutLongviewPlanRequest() // PutLongviewPlanRequest | Update your Longview subscription plan.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LongviewAPI.PutLongviewPlan(context.Background(), apiVersion).PutLongviewPlanRequest(putLongviewPlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LongviewAPI.PutLongviewPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLongviewPlan`: GetLongviewPlan200Response
	fmt.Fprintf(os.Stdout, "Response from `LongviewAPI.PutLongviewPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPutLongviewPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putLongviewPlanRequest** | [**PutLongviewPlanRequest**](PutLongviewPlanRequest.md) | Update your Longview subscription plan. | 

### Return type

[**GetLongviewPlan200Response**](GetLongviewPlan200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

