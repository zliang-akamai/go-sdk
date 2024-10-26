# \LinodeStackScriptsAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteStackScript**](LinodeStackScriptsAPI.md#DeleteStackScript) | **Delete** /{apiVersion}/linode/stackscripts/{stackscriptId} | Delete a StackScript
[**GetStackScript**](LinodeStackScriptsAPI.md#GetStackScript) | **Get** /{apiVersion}/linode/stackscripts/{stackscriptId} | Get a StackScript
[**GetStackScripts**](LinodeStackScriptsAPI.md#GetStackScripts) | **Get** /{apiVersion}/linode/stackscripts | List StackScripts
[**PostAddStackScript**](LinodeStackScriptsAPI.md#PostAddStackScript) | **Post** /{apiVersion}/linode/stackscripts | Create a StackScript
[**PutStackScript**](LinodeStackScriptsAPI.md#PutStackScript) | **Put** /{apiVersion}/linode/stackscripts/{stackscriptId} | Update a StackScript



## DeleteStackScript

> map[string]interface{} DeleteStackScript(ctx, apiVersion, stackscriptId).Execute()

Delete a StackScript



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
	stackscriptId := "stackscriptId_example" // string | The ID of the StackScript to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeStackScriptsAPI.DeleteStackScript(context.Background(), apiVersion, stackscriptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeStackScriptsAPI.DeleteStackScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStackScript`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeStackScriptsAPI.DeleteStackScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**stackscriptId** | **string** | The ID of the StackScript to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStackScriptRequest struct via the builder pattern


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


## GetStackScript

> GetStackScripts200ResponseDataInner GetStackScript(ctx, apiVersion, stackscriptId).Execute()

Get a StackScript



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
	stackscriptId := "stackscriptId_example" // string | The ID of the StackScript to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeStackScriptsAPI.GetStackScript(context.Background(), apiVersion, stackscriptId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeStackScriptsAPI.GetStackScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackScript`: GetStackScripts200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeStackScriptsAPI.GetStackScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**stackscriptId** | **string** | The ID of the StackScript to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetStackScripts200ResponseDataInner**](GetStackScripts200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStackScripts

> GetStackScripts200Response GetStackScripts(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List StackScripts



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
	resp, r, err := apiClient.LinodeStackScriptsAPI.GetStackScripts(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeStackScriptsAPI.GetStackScripts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStackScripts`: GetStackScripts200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeStackScriptsAPI.GetStackScripts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetStackScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetStackScripts200Response**](GetStackScripts200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAddStackScript

> GetStackScripts200ResponseDataInner PostAddStackScript(ctx, apiVersion).PostAddStackScriptRequest(postAddStackScriptRequest).Execute()

Create a StackScript



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
	postAddStackScriptRequest := *openapiclient.NewPostAddStackScriptRequest([]string{"Images_example"}, "a-stackscript", "\"#!/bin/bash\"") // PostAddStackScriptRequest | The properties to set for the new StackScript.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeStackScriptsAPI.PostAddStackScript(context.Background(), apiVersion).PostAddStackScriptRequest(postAddStackScriptRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeStackScriptsAPI.PostAddStackScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAddStackScript`: GetStackScripts200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeStackScriptsAPI.PostAddStackScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostAddStackScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postAddStackScriptRequest** | [**PostAddStackScriptRequest**](PostAddStackScriptRequest.md) | The properties to set for the new StackScript. | 

### Return type

[**GetStackScripts200ResponseDataInner**](GetStackScripts200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutStackScript

> GetStackScripts200ResponseDataInner PutStackScript(ctx, apiVersion, stackscriptId).GetStackScripts200ResponseDataInner(getStackScripts200ResponseDataInner).Execute()

Update a StackScript



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
	stackscriptId := "stackscriptId_example" // string | The ID of the StackScript to look up.
	getStackScripts200ResponseDataInner := *openapiclient.NewGetStackScripts200ResponseDataInner() // GetStackScripts200ResponseDataInner | The fields to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeStackScriptsAPI.PutStackScript(context.Background(), apiVersion, stackscriptId).GetStackScripts200ResponseDataInner(getStackScripts200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeStackScriptsAPI.PutStackScript``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutStackScript`: GetStackScripts200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeStackScriptsAPI.PutStackScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**stackscriptId** | **string** | The ID of the StackScript to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutStackScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getStackScripts200ResponseDataInner** | [**GetStackScripts200ResponseDataInner**](GetStackScripts200ResponseDataInner.md) | The fields to update. | 

### Return type

[**GetStackScripts200ResponseDataInner**](GetStackScripts200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

