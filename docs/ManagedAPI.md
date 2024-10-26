# \ManagedAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteManagedContact**](ManagedAPI.md#DeleteManagedContact) | **Delete** /{apiVersion}/managed/contacts/{contactId} | Delete a managed contact
[**DeleteManagedService**](ManagedAPI.md#DeleteManagedService) | **Delete** /{apiVersion}/managed/services/{serviceId} | Delete a managed service
[**GetManagedContact**](ManagedAPI.md#GetManagedContact) | **Get** /{apiVersion}/managed/contacts/{contactId} | Get a managed contact
[**GetManagedContacts**](ManagedAPI.md#GetManagedContacts) | **Get** /{apiVersion}/managed/contacts | List managed contacts
[**GetManagedCredential**](ManagedAPI.md#GetManagedCredential) | **Get** /{apiVersion}/managed/credentials/{credentialId} | Get a managed credential
[**GetManagedCredentials**](ManagedAPI.md#GetManagedCredentials) | **Get** /{apiVersion}/managed/credentials | List managed credentials
[**GetManagedIssue**](ManagedAPI.md#GetManagedIssue) | **Get** /{apiVersion}/managed/issues/{issueId} | Get a managed issue
[**GetManagedIssues**](ManagedAPI.md#GetManagedIssues) | **Get** /{apiVersion}/managed/issues | List managed issues
[**GetManagedLinodeSetting**](ManagedAPI.md#GetManagedLinodeSetting) | **Get** /{apiVersion}/managed/linode-settings/{linodeId} | Get a Linode&#39;s managed settings
[**GetManagedLinodeSettings**](ManagedAPI.md#GetManagedLinodeSettings) | **Get** /{apiVersion}/managed/linode-settings | List managed Linode settings
[**GetManagedService**](ManagedAPI.md#GetManagedService) | **Get** /{apiVersion}/managed/services/{serviceId} | Get a managed service
[**GetManagedServices**](ManagedAPI.md#GetManagedServices) | **Get** /{apiVersion}/managed/services | List managed services
[**GetManagedSshKey**](ManagedAPI.md#GetManagedSshKey) | **Get** /{apiVersion}/managed/credentials/sshkey | Get a managed SSH key
[**GetManagedStats**](ManagedAPI.md#GetManagedStats) | **Get** /{apiVersion}/managed/stats | List managed stats
[**PostDisableManagedService**](ManagedAPI.md#PostDisableManagedService) | **Post** /{apiVersion}/managed/services/{serviceId}/disable | Disable a managed service
[**PostEnableManagedService**](ManagedAPI.md#PostEnableManagedService) | **Post** /{apiVersion}/managed/services/{serviceId}/enable | Enable a managed service
[**PostManagedContact**](ManagedAPI.md#PostManagedContact) | **Post** /{apiVersion}/managed/contacts | Create a managed contact
[**PostManagedCredential**](ManagedAPI.md#PostManagedCredential) | **Post** /{apiVersion}/managed/credentials | Create a managed credential
[**PostManagedCredentialRevoke**](ManagedAPI.md#PostManagedCredentialRevoke) | **Post** /{apiVersion}/managed/credentials/{credentialId}/revoke | Delete a managed credential
[**PostManagedCredentialUsernamePassword**](ManagedAPI.md#PostManagedCredentialUsernamePassword) | **Post** /{apiVersion}/managed/credentials/{credentialId}/update | Update a managed credential&#39;s username and password
[**PostManagedService**](ManagedAPI.md#PostManagedService) | **Post** /{apiVersion}/managed/services | Create a managed service
[**PutManagedContact**](ManagedAPI.md#PutManagedContact) | **Put** /{apiVersion}/managed/contacts/{contactId} | Update a managed contact
[**PutManagedCredential**](ManagedAPI.md#PutManagedCredential) | **Put** /{apiVersion}/managed/credentials/{credentialId} | Update a managed credential
[**PutManagedLinodeSetting**](ManagedAPI.md#PutManagedLinodeSetting) | **Put** /{apiVersion}/managed/linode-settings/{linodeId} | Update a Linode&#39;s managed settings
[**PutManagedService**](ManagedAPI.md#PutManagedService) | **Put** /{apiVersion}/managed/services/{serviceId} | Update a managed service



## DeleteManagedContact

> map[string]interface{} DeleteManagedContact(ctx, apiVersion, contactId).Execute()

Delete a managed contact



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
	contactId := int32(56) // int32 | The ID of the contact to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.DeleteManagedContact(context.Background(), apiVersion, contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.DeleteManagedContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteManagedContact`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.DeleteManagedContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**contactId** | **int32** | The ID of the contact to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedContactRequest struct via the builder pattern


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


## DeleteManagedService

> map[string]interface{} DeleteManagedService(ctx, apiVersion, serviceId).Execute()

Delete a managed service



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
	serviceId := int32(56) // int32 | The ID of the Managed Service to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.DeleteManagedService(context.Background(), apiVersion, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.DeleteManagedService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteManagedService`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.DeleteManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**serviceId** | **int32** | The ID of the Managed Service to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteManagedServiceRequest struct via the builder pattern


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


## GetManagedContact

> GetManagedContacts200ResponseDataInner GetManagedContact(ctx, apiVersion, contactId).Execute()

Get a managed contact



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
	contactId := int32(56) // int32 | The ID of the contact to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.GetManagedContact(context.Background(), apiVersion, contactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedContact`: GetManagedContacts200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**contactId** | **int32** | The ID of the contact to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetManagedContacts200ResponseDataInner**](GetManagedContacts200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedContacts

> GetManagedContacts200Response GetManagedContacts(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List managed contacts



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
	resp, r, err := apiClient.ManagedAPI.GetManagedContacts(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedContacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedContacts`: GetManagedContacts200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedContacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedContactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetManagedContacts200Response**](GetManagedContacts200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedCredential

> GetManagedCredentials200ResponseDataInner GetManagedCredential(ctx, apiVersion, credentialId).Execute()

Get a managed credential



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
	credentialId := int32(56) // int32 | The ID of the Credential to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.GetManagedCredential(context.Background(), apiVersion, credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedCredential`: GetManagedCredentials200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**credentialId** | **int32** | The ID of the Credential to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetManagedCredentials200ResponseDataInner**](GetManagedCredentials200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedCredentials

> GetManagedCredentials200Response GetManagedCredentials(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List managed credentials



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
	resp, r, err := apiClient.ManagedAPI.GetManagedCredentials(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedCredentials`: GetManagedCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetManagedCredentials200Response**](GetManagedCredentials200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedIssue

> GetManagedIssues200ResponseDataInner GetManagedIssue(ctx, apiVersion, issueId).Execute()

Get a managed issue



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
	issueId := int32(56) // int32 | The Issue to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.GetManagedIssue(context.Background(), apiVersion, issueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedIssue`: GetManagedIssues200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**issueId** | **int32** | The Issue to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetManagedIssues200ResponseDataInner**](GetManagedIssues200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedIssues

> GetManagedIssues200Response GetManagedIssues(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List managed issues



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
	resp, r, err := apiClient.ManagedAPI.GetManagedIssues(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedIssues`: GetManagedIssues200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetManagedIssues200Response**](GetManagedIssues200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedLinodeSetting

> GetManagedLinodeSettings200ResponseDataInner GetManagedLinodeSetting(ctx, apiVersion, linodeId).Execute()

Get a Linode's managed settings



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
	linodeId := int32(56) // int32 | The Linode ID whose settings we are accessing.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.GetManagedLinodeSetting(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedLinodeSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedLinodeSetting`: GetManagedLinodeSettings200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedLinodeSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The Linode ID whose settings we are accessing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedLinodeSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetManagedLinodeSettings200ResponseDataInner**](GetManagedLinodeSettings200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedLinodeSettings

> GetManagedLinodeSettings200Response GetManagedLinodeSettings(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List managed Linode settings



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
	resp, r, err := apiClient.ManagedAPI.GetManagedLinodeSettings(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedLinodeSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedLinodeSettings`: GetManagedLinodeSettings200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedLinodeSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedLinodeSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetManagedLinodeSettings200Response**](GetManagedLinodeSettings200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedService

> GetManagedServices200ResponseDataInner GetManagedService(ctx, apiVersion, serviceId).Execute()

Get a managed service



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
	serviceId := int32(56) // int32 | The ID of the Managed Service to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.GetManagedService(context.Background(), apiVersion, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedService`: GetManagedServices200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**serviceId** | **int32** | The ID of the Managed Service to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetManagedServices200ResponseDataInner**](GetManagedServices200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedServices

> GetManagedServices200Response GetManagedServices(ctx, apiVersion).Execute()

List managed services



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
	resp, r, err := apiClient.ManagedAPI.GetManagedServices(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedServices`: GetManagedServices200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetManagedServices200Response**](GetManagedServices200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedSshKey

> GetManagedSshKey200Response GetManagedSshKey(ctx, apiVersion).Execute()

Get a managed SSH key



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
	resp, r, err := apiClient.ManagedAPI.GetManagedSshKey(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedSshKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedSshKey`: GetManagedSshKey200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedSshKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedSshKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetManagedSshKey200Response**](GetManagedSshKey200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetManagedStats

> GetManagedStats200Response GetManagedStats(ctx, apiVersion).Execute()

List managed stats



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
	resp, r, err := apiClient.ManagedAPI.GetManagedStats(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.GetManagedStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetManagedStats`: GetManagedStats200Response
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.GetManagedStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetManagedStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetManagedStats200Response**](GetManagedStats200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDisableManagedService

> GetManagedServices200ResponseDataInner PostDisableManagedService(ctx, apiVersion, serviceId).Execute()

Disable a managed service



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
	serviceId := int32(56) // int32 | The ID of the Managed Service to disable.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PostDisableManagedService(context.Background(), apiVersion, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PostDisableManagedService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDisableManagedService`: GetManagedServices200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PostDisableManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**serviceId** | **int32** | The ID of the Managed Service to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDisableManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetManagedServices200ResponseDataInner**](GetManagedServices200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnableManagedService

> GetManagedServices200ResponseDataInner PostEnableManagedService(ctx, apiVersion, serviceId).Execute()

Enable a managed service



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
	serviceId := int32(56) // int32 | The ID of the Managed Service to enable.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PostEnableManagedService(context.Background(), apiVersion, serviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PostEnableManagedService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEnableManagedService`: GetManagedServices200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PostEnableManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**serviceId** | **int32** | The ID of the Managed Service to enable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEnableManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetManagedServices200ResponseDataInner**](GetManagedServices200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostManagedContact

> GetManagedContacts200ResponseDataInner PostManagedContact(ctx, apiVersion).GetManagedContacts200ResponseDataInner(getManagedContacts200ResponseDataInner).Execute()

Create a managed contact



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
	getManagedContacts200ResponseDataInner := *openapiclient.NewGetManagedContacts200ResponseDataInner() // GetManagedContacts200ResponseDataInner | Information about the contact to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PostManagedContact(context.Background(), apiVersion).GetManagedContacts200ResponseDataInner(getManagedContacts200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PostManagedContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostManagedContact`: GetManagedContacts200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PostManagedContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostManagedContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **getManagedContacts200ResponseDataInner** | [**GetManagedContacts200ResponseDataInner**](GetManagedContacts200ResponseDataInner.md) | Information about the contact to create. | 

### Return type

[**GetManagedContacts200ResponseDataInner**](GetManagedContacts200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostManagedCredential

> GetManagedCredentials200ResponseDataInner PostManagedCredential(ctx, apiVersion).PostManagedCredentialRequest(postManagedCredentialRequest).Execute()

Create a managed credential



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
	postManagedCredentialRequest := *openapiclient.NewPostManagedCredentialRequest("prod-password-1", "s3cur3P@ssw0rd") // PostManagedCredentialRequest | Information about the Credential to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PostManagedCredential(context.Background(), apiVersion).PostManagedCredentialRequest(postManagedCredentialRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PostManagedCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostManagedCredential`: GetManagedCredentials200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PostManagedCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostManagedCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postManagedCredentialRequest** | [**PostManagedCredentialRequest**](PostManagedCredentialRequest.md) | Information about the Credential to create. | 

### Return type

[**GetManagedCredentials200ResponseDataInner**](GetManagedCredentials200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostManagedCredentialRevoke

> map[string]interface{} PostManagedCredentialRevoke(ctx, apiVersion, credentialId).Execute()

Delete a managed credential



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
	credentialId := int32(56) // int32 | The ID of the Credential to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PostManagedCredentialRevoke(context.Background(), apiVersion, credentialId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PostManagedCredentialRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostManagedCredentialRevoke`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PostManagedCredentialRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**credentialId** | **int32** | The ID of the Credential to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostManagedCredentialRevokeRequest struct via the builder pattern


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


## PostManagedCredentialUsernamePassword

> map[string]interface{} PostManagedCredentialUsernamePassword(ctx, apiVersion, credentialId).PostManagedCredentialUsernamePasswordRequest(postManagedCredentialUsernamePasswordRequest).Execute()

Update a managed credential's username and password



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
	credentialId := int32(56) // int32 | The ID of the Credential to update.
	postManagedCredentialUsernamePasswordRequest := *openapiclient.NewPostManagedCredentialUsernamePasswordRequest("s3cur3P@ssw0rd") // PostManagedCredentialUsernamePasswordRequest | The new username and password to assign to the Managed Credential. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PostManagedCredentialUsernamePassword(context.Background(), apiVersion, credentialId).PostManagedCredentialUsernamePasswordRequest(postManagedCredentialUsernamePasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PostManagedCredentialUsernamePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostManagedCredentialUsernamePassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PostManagedCredentialUsernamePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**credentialId** | **int32** | The ID of the Credential to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostManagedCredentialUsernamePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postManagedCredentialUsernamePasswordRequest** | [**PostManagedCredentialUsernamePasswordRequest**](PostManagedCredentialUsernamePasswordRequest.md) | The new username and password to assign to the Managed Credential. | 

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


## PostManagedService

> GetManagedServices200ResponseDataInner PostManagedService(ctx, apiVersion).PostManagedServiceRequest(postManagedServiceRequest).Execute()

Create a managed service



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
	postManagedServiceRequest := *openapiclient.NewPostManagedServiceRequest("https://example.org", "prod-1", "url", int32(30)) // PostManagedServiceRequest | Information about the service to monitor. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PostManagedService(context.Background(), apiVersion).PostManagedServiceRequest(postManagedServiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PostManagedService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostManagedService`: GetManagedServices200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PostManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postManagedServiceRequest** | [**PostManagedServiceRequest**](PostManagedServiceRequest.md) | Information about the service to monitor. | 

### Return type

[**GetManagedServices200ResponseDataInner**](GetManagedServices200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutManagedContact

> GetManagedContacts200ResponseDataInner PutManagedContact(ctx, apiVersion, contactId).GetManagedContacts200ResponseDataInner(getManagedContacts200ResponseDataInner).Execute()

Update a managed contact



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
	contactId := int32(56) // int32 | The ID of the contact to access.
	getManagedContacts200ResponseDataInner := *openapiclient.NewGetManagedContacts200ResponseDataInner() // GetManagedContacts200ResponseDataInner | The fields to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PutManagedContact(context.Background(), apiVersion, contactId).GetManagedContacts200ResponseDataInner(getManagedContacts200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PutManagedContact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutManagedContact`: GetManagedContacts200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PutManagedContact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**contactId** | **int32** | The ID of the contact to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutManagedContactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getManagedContacts200ResponseDataInner** | [**GetManagedContacts200ResponseDataInner**](GetManagedContacts200ResponseDataInner.md) | The fields to update. | 

### Return type

[**GetManagedContacts200ResponseDataInner**](GetManagedContacts200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutManagedCredential

> GetManagedCredentials200ResponseDataInner PutManagedCredential(ctx, apiVersion, credentialId).GetManagedCredentials200ResponseDataInner(getManagedCredentials200ResponseDataInner).Execute()

Update a managed credential



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
	credentialId := int32(56) // int32 | The ID of the Credential to access.
	getManagedCredentials200ResponseDataInner := *openapiclient.NewGetManagedCredentials200ResponseDataInner() // GetManagedCredentials200ResponseDataInner | The fields to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PutManagedCredential(context.Background(), apiVersion, credentialId).GetManagedCredentials200ResponseDataInner(getManagedCredentials200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PutManagedCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutManagedCredential`: GetManagedCredentials200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PutManagedCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**credentialId** | **int32** | The ID of the Credential to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutManagedCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getManagedCredentials200ResponseDataInner** | [**GetManagedCredentials200ResponseDataInner**](GetManagedCredentials200ResponseDataInner.md) | The fields to update. | 

### Return type

[**GetManagedCredentials200ResponseDataInner**](GetManagedCredentials200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutManagedLinodeSetting

> GetManagedLinodeSettings200ResponseDataInner PutManagedLinodeSetting(ctx, apiVersion, linodeId).GetManagedLinodeSettings200ResponseDataInner(getManagedLinodeSettings200ResponseDataInner).Execute()

Update a Linode's managed settings



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
	linodeId := int32(56) // int32 | The Linode ID whose settings we are accessing.
	getManagedLinodeSettings200ResponseDataInner := *openapiclient.NewGetManagedLinodeSettings200ResponseDataInner() // GetManagedLinodeSettings200ResponseDataInner | The settings to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PutManagedLinodeSetting(context.Background(), apiVersion, linodeId).GetManagedLinodeSettings200ResponseDataInner(getManagedLinodeSettings200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PutManagedLinodeSetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutManagedLinodeSetting`: GetManagedLinodeSettings200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PutManagedLinodeSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The Linode ID whose settings we are accessing. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutManagedLinodeSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getManagedLinodeSettings200ResponseDataInner** | [**GetManagedLinodeSettings200ResponseDataInner**](GetManagedLinodeSettings200ResponseDataInner.md) | The settings to update. | 

### Return type

[**GetManagedLinodeSettings200ResponseDataInner**](GetManagedLinodeSettings200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutManagedService

> GetManagedServices200ResponseDataInner PutManagedService(ctx, apiVersion, serviceId).GetManagedServices200ResponseDataInner(getManagedServices200ResponseDataInner).Execute()

Update a managed service



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
	serviceId := int32(56) // int32 | The ID of the Managed Service to access.
	getManagedServices200ResponseDataInner := *openapiclient.NewGetManagedServices200ResponseDataInner() // GetManagedServices200ResponseDataInner | The fields to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManagedAPI.PutManagedService(context.Background(), apiVersion, serviceId).GetManagedServices200ResponseDataInner(getManagedServices200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManagedAPI.PutManagedService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutManagedService`: GetManagedServices200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ManagedAPI.PutManagedService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**serviceId** | **int32** | The ID of the Managed Service to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutManagedServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **getManagedServices200ResponseDataInner** | [**GetManagedServices200ResponseDataInner**](GetManagedServices200ResponseDataInner.md) | The fields to update. | 

### Return type

[**GetManagedServices200ResponseDataInner**](GetManagedServices200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

