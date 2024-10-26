# \VolumesAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVolume**](VolumesAPI.md#DeleteVolume) | **Delete** /{apiVersion}/volumes/{volumeId} | Delete a volume
[**GetVolume**](VolumesAPI.md#GetVolume) | **Get** /{apiVersion}/volumes/{volumeId} | Get a volume
[**GetVolumeTypes**](VolumesAPI.md#GetVolumeTypes) | **Get** /{apiVersion}/volumes/types | List volume types
[**GetVolumes**](VolumesAPI.md#GetVolumes) | **Get** /{apiVersion}/volumes | List volumes
[**PostAttachVolume**](VolumesAPI.md#PostAttachVolume) | **Post** /{apiVersion}/volumes/{volumeId}/attach | Attach a volume
[**PostCloneVolume**](VolumesAPI.md#PostCloneVolume) | **Post** /{apiVersion}/volumes/{volumeId}/clone | Clone a volume
[**PostDetachVolume**](VolumesAPI.md#PostDetachVolume) | **Post** /{apiVersion}/volumes/{volumeId}/detach | Detach a volume
[**PostResizeVolume**](VolumesAPI.md#PostResizeVolume) | **Post** /{apiVersion}/volumes/{volumeId}/resize | Resize a volume
[**PostVolume**](VolumesAPI.md#PostVolume) | **Post** /{apiVersion}/volumes | Create a volume
[**PutVolume**](VolumesAPI.md#PutVolume) | **Put** /{apiVersion}/volumes/{volumeId} | Update a volume



## DeleteVolume

> map[string]interface{} DeleteVolume(ctx, apiVersion, volumeId).Execute()

Delete a volume



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
	volumeId := int32(56) // int32 | ID of the Volume to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.DeleteVolume(context.Background(), apiVersion, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.DeleteVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteVolume`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.DeleteVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**volumeId** | **int32** | ID of the Volume to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVolumeRequest struct via the builder pattern


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


## GetVolume

> Volume GetVolume(ctx, apiVersion, volumeId).Page(page).PageSize(pageSize).Execute()

Get a volume



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
	volumeId := int32(56) // int32 | ID of the Volume to look up.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.GetVolume(context.Background(), apiVersion, volumeId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.GetVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**volumeId** | **int32** | ID of the Volume to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**Volume**](Volume.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumeTypes

> GetVolumeTypes200Response GetVolumeTypes(ctx, apiVersion).Execute()

List volume types



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
	resp, r, err := apiClient.VolumesAPI.GetVolumeTypes(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetVolumeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolumeTypes`: GetVolumeTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.GetVolumeTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumeTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetVolumeTypes200Response**](GetVolumeTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVolumes

> GetLinodeVolumes200Response GetVolumes(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List volumes



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
	resp, r, err := apiClient.VolumesAPI.GetVolumes(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.GetVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVolumes`: GetLinodeVolumes200Response
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.GetVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetVolumesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetLinodeVolumes200Response**](GetLinodeVolumes200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAttachVolume

> Volume PostAttachVolume(ctx, apiVersion, volumeId).PostAttachVolumeRequest(postAttachVolumeRequest).Execute()

Attach a volume



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
	volumeId := int32(56) // int32 | ID of the Volume to attach.
	postAttachVolumeRequest := *openapiclient.NewPostAttachVolumeRequest(int32(123)) // PostAttachVolumeRequest | Volume to attach to a Linode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.PostAttachVolume(context.Background(), apiVersion, volumeId).PostAttachVolumeRequest(postAttachVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.PostAttachVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAttachVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.PostAttachVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**volumeId** | **int32** | ID of the Volume to attach. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAttachVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postAttachVolumeRequest** | [**PostAttachVolumeRequest**](PostAttachVolumeRequest.md) | Volume to attach to a Linode. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloneVolume

> Volume PostCloneVolume(ctx, apiVersion, volumeId).PostCloneVolumeRequest(postCloneVolumeRequest).Execute()

Clone a volume



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
	volumeId := int32(56) // int32 | ID of the Volume to clone.
	postCloneVolumeRequest := *openapiclient.NewPostCloneVolumeRequest("my-volume") // PostCloneVolumeRequest | The requested state your Volume will be cloned into.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.PostCloneVolume(context.Background(), apiVersion, volumeId).PostCloneVolumeRequest(postCloneVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.PostCloneVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloneVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.PostCloneVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**volumeId** | **int32** | ID of the Volume to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloneVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postCloneVolumeRequest** | [**PostCloneVolumeRequest**](PostCloneVolumeRequest.md) | The requested state your Volume will be cloned into. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDetachVolume

> map[string]interface{} PostDetachVolume(ctx, apiVersion, volumeId).Execute()

Detach a volume



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
	volumeId := int32(56) // int32 | ID of the Volume to detach.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.PostDetachVolume(context.Background(), apiVersion, volumeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.PostDetachVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDetachVolume`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.PostDetachVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**volumeId** | **int32** | ID of the Volume to detach. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDetachVolumeRequest struct via the builder pattern


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


## PostResizeVolume

> Volume PostResizeVolume(ctx, apiVersion, volumeId).PostResizeVolumeRequest(postResizeVolumeRequest).Execute()

Resize a volume



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
	volumeId := int32(56) // int32 | ID of the Volume to resize.
	postResizeVolumeRequest := *openapiclient.NewPostResizeVolumeRequest(int32(30)) // PostResizeVolumeRequest | The requested size to increase your Volume to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.PostResizeVolume(context.Background(), apiVersion, volumeId).PostResizeVolumeRequest(postResizeVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.PostResizeVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResizeVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.PostResizeVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**volumeId** | **int32** | ID of the Volume to resize. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostResizeVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postResizeVolumeRequest** | [**PostResizeVolumeRequest**](PostResizeVolumeRequest.md) | The requested size to increase your Volume to. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostVolume

> Volume PostVolume(ctx, apiVersion).PostVolumeRequest(postVolumeRequest).Execute()

Create a volume



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
	postVolumeRequest := *openapiclient.NewPostVolumeRequest("my-volume") // PostVolumeRequest | The requested initial state of a new Volume.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.PostVolume(context.Background(), apiVersion).PostVolumeRequest(postVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.PostVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.PostVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postVolumeRequest** | [**PostVolumeRequest**](PostVolumeRequest.md) | The requested initial state of a new Volume. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutVolume

> Volume PutVolume(ctx, apiVersion, volumeId).PutVolumeRequest(putVolumeRequest).Execute()

Update a volume



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
	volumeId := int32(56) // int32 | ID of the Volume to look up.
	putVolumeRequest := *openapiclient.NewPutVolumeRequest() // PutVolumeRequest | If any updated field fails to pass validation, the Volume will not be updated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VolumesAPI.PutVolume(context.Background(), apiVersion, volumeId).PutVolumeRequest(putVolumeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VolumesAPI.PutVolume``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutVolume`: Volume
	fmt.Fprintf(os.Stdout, "Response from `VolumesAPI.PutVolume`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**volumeId** | **int32** | ID of the Volume to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutVolumeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putVolumeRequest** | [**PutVolumeRequest**](PutVolumeRequest.md) | If any updated field fails to pass validation, the Volume will not be updated. | 

### Return type

[**Volume**](Volume.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

