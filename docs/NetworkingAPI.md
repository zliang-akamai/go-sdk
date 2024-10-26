# \NetworkingAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFirewall**](NetworkingAPI.md#DeleteFirewall) | **Delete** /{apiVersion}/networking/firewalls/{firewallId} | Delete a firewall
[**DeleteFirewallDevice**](NetworkingAPI.md#DeleteFirewallDevice) | **Delete** /{apiVersion}/networking/firewalls/{firewallId}/devices/{deviceId} | Delete a firewall device
[**DeleteIpv6Range**](NetworkingAPI.md#DeleteIpv6Range) | **Delete** /{apiVersion}/networking/ipv6/ranges/{range} | Delete an IPv6 range
[**GetFirewall**](NetworkingAPI.md#GetFirewall) | **Get** /{apiVersion}/networking/firewalls/{firewallId} | Get a firewall
[**GetFirewallDevice**](NetworkingAPI.md#GetFirewallDevice) | **Get** /{apiVersion}/networking/firewalls/{firewallId}/devices/{deviceId} | Get a firewall device
[**GetFirewallDevices**](NetworkingAPI.md#GetFirewallDevices) | **Get** /{apiVersion}/networking/firewalls/{firewallId}/devices | List firewall devices
[**GetFirewallRuleVersion**](NetworkingAPI.md#GetFirewallRuleVersion) | **Get** /{apiVersion}/networking/firewalls/{firewallId}/history/rules/{version} | Get a firewall rule version
[**GetFirewallRuleVersions**](NetworkingAPI.md#GetFirewallRuleVersions) | **Get** /{apiVersion}/networking/firewalls/{firewallId}/history | List firewall rule versions
[**GetFirewallRules**](NetworkingAPI.md#GetFirewallRules) | **Get** /{apiVersion}/networking/firewalls/{firewallId}/rules | List firewall rules
[**GetFirewalls**](NetworkingAPI.md#GetFirewalls) | **Get** /{apiVersion}/networking/firewalls | List firewalls
[**GetIp**](NetworkingAPI.md#GetIp) | **Get** /{apiVersion}/networking/ips/{address} | Get an IP address
[**GetIps**](NetworkingAPI.md#GetIps) | **Get** /{apiVersion}/networking/ips | List IP addresses
[**GetIpv6Pools**](NetworkingAPI.md#GetIpv6Pools) | **Get** /{apiVersion}/networking/ipv6/pools | List IPv6 pools
[**GetIpv6Range**](NetworkingAPI.md#GetIpv6Range) | **Get** /{apiVersion}/networking/ipv6/ranges/{range} | Get an IPv6 range
[**GetIpv6Ranges**](NetworkingAPI.md#GetIpv6Ranges) | **Get** /{apiVersion}/networking/ipv6/ranges | List IPv6 ranges
[**GetVlans**](NetworkingAPI.md#GetVlans) | **Get** /{apiVersion}/networking/vlans | List VLANs
[**PostAllocateIp**](NetworkingAPI.md#PostAllocateIp) | **Post** /{apiVersion}/networking/ips | Allocate an IP address
[**PostAssignIps**](NetworkingAPI.md#PostAssignIps) | **Post** /{apiVersion}/networking/ips/assign | Assign IP addresses
[**PostAssignIpv4s**](NetworkingAPI.md#PostAssignIpv4s) | **Post** /{apiVersion}/networking/ipv4/assign | Assign IPv4s to Linodes
[**PostFirewallDevice**](NetworkingAPI.md#PostFirewallDevice) | **Post** /{apiVersion}/networking/firewalls/{firewallId}/devices | Create a firewall device
[**PostFirewalls**](NetworkingAPI.md#PostFirewalls) | **Post** /{apiVersion}/networking/firewalls | Create a firewall
[**PostIpv6Range**](NetworkingAPI.md#PostIpv6Range) | **Post** /{apiVersion}/networking/ipv6/ranges | Create an IPv6 range
[**PostShareIps**](NetworkingAPI.md#PostShareIps) | **Post** /{apiVersion}/networking/ips/share | Share IP addresses
[**PostShareIpv4s**](NetworkingAPI.md#PostShareIpv4s) | **Post** /{apiVersion}/networking/ipv4/share | Configure IPv4 sharing
[**PutFirewall**](NetworkingAPI.md#PutFirewall) | **Put** /{apiVersion}/networking/firewalls/{firewallId} | Update a firewall
[**PutFirewallRules**](NetworkingAPI.md#PutFirewallRules) | **Put** /{apiVersion}/networking/firewalls/{firewallId}/rules | Update firewall rules
[**PutIp**](NetworkingAPI.md#PutIp) | **Put** /{apiVersion}/networking/ips/{address} | Update an IP address&#39;s RDNS



## DeleteFirewall

> map[string]interface{} DeleteFirewall(ctx, apiVersion, firewallId).Execute()

Delete a firewall



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.DeleteFirewall(context.Background(), apiVersion, firewallId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.DeleteFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewall`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.DeleteFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallRequest struct via the builder pattern


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


## DeleteFirewallDevice

> map[string]interface{} DeleteFirewallDevice(ctx, apiVersion, firewallId, deviceId).Execute()

Delete a firewall device



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.
	deviceId := int32(56) // int32 | ID of the Firewall Device to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.DeleteFirewallDevice(context.Background(), apiVersion, firewallId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.DeleteFirewallDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFirewallDevice`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.DeleteFirewallDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 
**deviceId** | **int32** | ID of the Firewall Device to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFirewallDeviceRequest struct via the builder pattern


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


## DeleteIpv6Range

> map[string]interface{} DeleteIpv6Range(ctx, apiVersion, range_).Execute()

Delete an IPv6 range



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
	range_ := "range__example" // string | The IPv6 range to access. Corresponds to the `range` property of objects returned from the [List IPv6 ranges](https://techdocs.akamai.com/linode-api/reference/get-ipv6-ranges) operation.  __Note__. Omit the prefix length of the IPv6 range.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.DeleteIpv6Range(context.Background(), apiVersion, range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.DeleteIpv6Range``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIpv6Range`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.DeleteIpv6Range`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**range_** | **string** | The IPv6 range to access. Corresponds to the &#x60;range&#x60; property of objects returned from the [List IPv6 ranges](https://techdocs.akamai.com/linode-api/reference/get-ipv6-ranges) operation.  __Note__. Omit the prefix length of the IPv6 range. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIpv6RangeRequest struct via the builder pattern


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


## GetFirewall

> GetLinodeFirewalls200ResponseDataInner GetFirewall(ctx, apiVersion, firewallId).Execute()

Get a firewall



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.GetFirewall(context.Background(), apiVersion, firewallId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewall`: GetLinodeFirewalls200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLinodeFirewalls200ResponseDataInner**](GetLinodeFirewalls200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallDevice

> GetFirewallDevices200ResponseDataInner GetFirewallDevice(ctx, apiVersion, firewallId, deviceId).Execute()

Get a firewall device



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.
	deviceId := int32(56) // int32 | ID of the Firewall Device to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.GetFirewallDevice(context.Background(), apiVersion, firewallId, deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetFirewallDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallDevice`: GetFirewallDevices200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetFirewallDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 
**deviceId** | **int32** | ID of the Firewall Device to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetFirewallDevices200ResponseDataInner**](GetFirewallDevices200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallDevices

> GetFirewallDevices200Response GetFirewallDevices(ctx, apiVersion, firewallId).Page(page).PageSize(pageSize).Execute()

List firewall devices



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.GetFirewallDevices(context.Background(), apiVersion, firewallId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetFirewallDevices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallDevices`: GetFirewallDevices200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetFirewallDevices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallDevicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetFirewallDevices200Response**](GetFirewallDevices200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRuleVersion

> GetLinodeFirewalls200ResponseDataInner GetFirewallRuleVersion(ctx, apiVersion, firewallId, version).Execute()

Get a firewall rule version



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.
	version := int32(3) // int32 | The firewall rule version to view.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.GetFirewallRuleVersion(context.Background(), apiVersion, firewallId, version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetFirewallRuleVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallRuleVersion`: GetLinodeFirewalls200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetFirewallRuleVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 
**version** | **int32** | The firewall rule version to view. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRuleVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetLinodeFirewalls200ResponseDataInner**](GetLinodeFirewalls200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRuleVersions

> GetLinodeFirewalls200ResponseDataInner GetFirewallRuleVersions(ctx, apiVersion, firewallId).Execute()

List firewall rule versions



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.GetFirewallRuleVersions(context.Background(), apiVersion, firewallId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetFirewallRuleVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallRuleVersions`: GetLinodeFirewalls200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetFirewallRuleVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRuleVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLinodeFirewalls200ResponseDataInner**](GetLinodeFirewalls200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewallRules

> GetLinodeFirewalls200ResponseDataInnerRules GetFirewallRules(ctx, apiVersion, firewallId).Execute()

List firewall rules



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.GetFirewallRules(context.Background(), apiVersion, firewallId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallRules`: GetLinodeFirewalls200ResponseDataInnerRules
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLinodeFirewalls200ResponseDataInnerRules**](GetLinodeFirewalls200ResponseDataInnerRules.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFirewalls

> GetLinodeFirewalls200Response GetFirewalls(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List firewalls



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
	resp, r, err := apiClient.NetworkingAPI.GetFirewalls(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewalls`: GetLinodeFirewalls200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetFirewalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetLinodeFirewalls200Response**](GetLinodeFirewalls200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIp

> GetLinodeIps200ResponseIpv4PublicInner GetIp(ctx, apiVersion, address).Execute()

Get an IP address



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
	address := "address_example" // string | The address to operate on.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.GetIp(context.Background(), apiVersion, address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIp`: GetLinodeIps200ResponseIpv4PublicInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**address** | **string** | The address to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLinodeIps200ResponseIpv4PublicInner**](GetLinodeIps200ResponseIpv4PublicInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIps

> GetIps200Response GetIps(ctx, apiVersion).SkipIpv6Rdns(skipIpv6Rdns).Execute()

List IP addresses



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
	skipIpv6Rdns := true // bool | When `true`, the `rdns` property for any `ipv6` type addresses always returns `null` regardless of whether RDNS data exists for the address. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.GetIps(context.Background(), apiVersion).SkipIpv6Rdns(skipIpv6Rdns).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIps`: GetIps200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **skipIpv6Rdns** | **bool** | When &#x60;true&#x60;, the &#x60;rdns&#x60; property for any &#x60;ipv6&#x60; type addresses always returns &#x60;null&#x60; regardless of whether RDNS data exists for the address. | [default to false]

### Return type

[**GetIps200Response**](GetIps200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpv6Pools

> GetIpv6Pools200Response GetIpv6Pools(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List IPv6 pools



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
	resp, r, err := apiClient.NetworkingAPI.GetIpv6Pools(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetIpv6Pools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpv6Pools`: GetIpv6Pools200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetIpv6Pools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpv6PoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetIpv6Pools200Response**](GetIpv6Pools200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpv6Range

> GetIpv6Range200Response GetIpv6Range(ctx, apiVersion, range_).Execute()

Get an IPv6 range



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
	range_ := "range__example" // string | The IPv6 range to access. Corresponds to the `range` property of objects returned from the [List IPv6 ranges](https://techdocs.akamai.com/linode-api/reference/get-ipv6-ranges) operation.  __Note__. Omit the prefix length of the IPv6 range.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.GetIpv6Range(context.Background(), apiVersion, range_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetIpv6Range``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpv6Range`: GetIpv6Range200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetIpv6Range`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**range_** | **string** | The IPv6 range to access. Corresponds to the &#x60;range&#x60; property of objects returned from the [List IPv6 ranges](https://techdocs.akamai.com/linode-api/reference/get-ipv6-ranges) operation.  __Note__. Omit the prefix length of the IPv6 range. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpv6RangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetIpv6Range200Response**](GetIpv6Range200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIpv6Ranges

> GetIpv6Ranges200Response GetIpv6Ranges(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List IPv6 ranges



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
	resp, r, err := apiClient.NetworkingAPI.GetIpv6Ranges(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetIpv6Ranges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIpv6Ranges`: GetIpv6Ranges200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetIpv6Ranges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetIpv6RangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetIpv6Ranges200Response**](GetIpv6Ranges200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVlans

> GetVlans200Response GetVlans(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List VLANs



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
	resp, r, err := apiClient.NetworkingAPI.GetVlans(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.GetVlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVlans`: GetVlans200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.GetVlans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetVlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetVlans200Response**](GetVlans200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAllocateIp

> GetLinodeIps200ResponseIpv4PublicInner PostAllocateIp(ctx, apiVersion).PostAllocateIpRequest(postAllocateIpRequest).Execute()

Allocate an IP address



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
	postAllocateIpRequest := *openapiclient.NewPostAllocateIpRequest(int32(123), true, "ipv4") // PostAllocateIpRequest | Information about the address you are creating.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PostAllocateIp(context.Background(), apiVersion).PostAllocateIpRequest(postAllocateIpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PostAllocateIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAllocateIp`: GetLinodeIps200ResponseIpv4PublicInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PostAllocateIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostAllocateIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postAllocateIpRequest** | [**PostAllocateIpRequest**](PostAllocateIpRequest.md) | Information about the address you are creating. | 

### Return type

[**GetLinodeIps200ResponseIpv4PublicInner**](GetLinodeIps200ResponseIpv4PublicInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAssignIps

> map[string]interface{} PostAssignIps(ctx, apiVersion).PostAssignIpsRequest(postAssignIpsRequest).Execute()

Assign IP addresses



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
	postAssignIpsRequest := *openapiclient.NewPostAssignIpsRequest([]openapiclient.PostAssignIpsRequestAssignmentsInner{*openapiclient.NewPostAssignIpsRequestAssignmentsInner("192.0.2.1", int32(123))}, "us-east") // PostAssignIpsRequest | Information about what IPv4 address or IPv6 range to assign, and to which Linode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PostAssignIps(context.Background(), apiVersion).PostAssignIpsRequest(postAssignIpsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PostAssignIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAssignIps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PostAssignIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostAssignIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postAssignIpsRequest** | [**PostAssignIpsRequest**](PostAssignIpsRequest.md) | Information about what IPv4 address or IPv6 range to assign, and to which Linode. | 

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


## PostAssignIpv4s

> map[string]interface{} PostAssignIpv4s(ctx, apiVersion).PostAssignIpsRequest(postAssignIpsRequest).Execute()

Assign IPv4s to Linodes



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
	postAssignIpsRequest := *openapiclient.NewPostAssignIpsRequest([]openapiclient.PostAssignIpsRequestAssignmentsInner{*openapiclient.NewPostAssignIpsRequestAssignmentsInner("192.0.2.1", int32(123))}, "us-east") // PostAssignIpsRequest | Information about what IPv4 address to assign, and to which Linode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PostAssignIpv4s(context.Background(), apiVersion).PostAssignIpsRequest(postAssignIpsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PostAssignIpv4s``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAssignIpv4s`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PostAssignIpv4s`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostAssignIpv4sRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postAssignIpsRequest** | [**PostAssignIpsRequest**](PostAssignIpsRequest.md) | Information about what IPv4 address to assign, and to which Linode. | 

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


## PostFirewallDevice

> GetFirewallDevices200ResponseDataInner PostFirewallDevice(ctx, apiVersion, firewallId).PostFirewallDeviceRequest(postFirewallDeviceRequest).Execute()

Create a firewall device



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.
	postFirewallDeviceRequest := *openapiclient.NewPostFirewallDeviceRequest(int32(123), "linode") // PostFirewallDeviceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PostFirewallDevice(context.Background(), apiVersion, firewallId).PostFirewallDeviceRequest(postFirewallDeviceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PostFirewallDevice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFirewallDevice`: GetFirewallDevices200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PostFirewallDevice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostFirewallDeviceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postFirewallDeviceRequest** | [**PostFirewallDeviceRequest**](PostFirewallDeviceRequest.md) |  | 

### Return type

[**GetFirewallDevices200ResponseDataInner**](GetFirewallDevices200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostFirewalls

> GetLinodeFirewalls200ResponseDataInner PostFirewalls(ctx, apiVersion).PostFirewallsRequest(postFirewallsRequest).Execute()

Create a firewall



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
	postFirewallsRequest := *openapiclient.NewPostFirewallsRequest(*openapiclient.NewPostFirewallsRequestAllOfRules()) // PostFirewallsRequest | Creates a Firewall object that can be applied to a service to filter the service's network traffic. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PostFirewalls(context.Background(), apiVersion).PostFirewallsRequest(postFirewallsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PostFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostFirewalls`: GetLinodeFirewalls200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PostFirewalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostFirewallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postFirewallsRequest** | [**PostFirewallsRequest**](PostFirewallsRequest.md) | Creates a Firewall object that can be applied to a service to filter the service&#39;s network traffic. | 

### Return type

[**GetLinodeFirewalls200ResponseDataInner**](GetLinodeFirewalls200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostIpv6Range

> PostIpv6Range200Response PostIpv6Range(ctx, apiVersion).PostIpv6RangeRequest(postIpv6RangeRequest).Execute()

Create an IPv6 range



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
	postIpv6RangeRequest := *openapiclient.NewPostIpv6RangeRequest(int32(123)) // PostIpv6RangeRequest | Information about the IPv6 range to create.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PostIpv6Range(context.Background(), apiVersion).PostIpv6RangeRequest(postIpv6RangeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PostIpv6Range``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostIpv6Range`: PostIpv6Range200Response
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PostIpv6Range`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostIpv6RangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postIpv6RangeRequest** | [**PostIpv6RangeRequest**](PostIpv6RangeRequest.md) | Information about the IPv6 range to create. | 

### Return type

[**PostIpv6Range200Response**](PostIpv6Range200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostShareIps

> map[string]interface{} PostShareIps(ctx, apiVersion).PostShareIpsRequest(postShareIpsRequest).Execute()

Share IP addresses



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
	postShareIpsRequest := *openapiclient.NewPostShareIpsRequest([]string{"Ips_example"}, int32(123)) // PostShareIpsRequest | Information about what IPs to share with which Linode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PostShareIps(context.Background(), apiVersion).PostShareIpsRequest(postShareIpsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PostShareIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostShareIps`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PostShareIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostShareIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postShareIpsRequest** | [**PostShareIpsRequest**](PostShareIpsRequest.md) | Information about what IPs to share with which Linode. | 

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


## PostShareIpv4s

> map[string]interface{} PostShareIpv4s(ctx, apiVersion).PostShareIpsRequest(postShareIpsRequest).Execute()

Configure IPv4 sharing



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
	postShareIpsRequest := *openapiclient.NewPostShareIpsRequest([]string{"Ips_example"}, int32(123)) // PostShareIpsRequest | Information about what IPs to share with which Linode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PostShareIpv4s(context.Background(), apiVersion).PostShareIpsRequest(postShareIpsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PostShareIpv4s``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostShareIpv4s`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PostShareIpv4s`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostShareIpv4sRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postShareIpsRequest** | [**PostShareIpsRequest**](PostShareIpsRequest.md) | Information about what IPs to share with which Linode. | 

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


## PutFirewall

> GetLinodeFirewalls200ResponseDataInner PutFirewall(ctx, apiVersion, firewallId).PutFirewallRequest(putFirewallRequest).Execute()

Update a firewall



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.
	putFirewallRequest := *openapiclient.NewPutFirewallRequest() // PutFirewallRequest | The Firewall information to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PutFirewall(context.Background(), apiVersion, firewallId).PutFirewallRequest(putFirewallRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PutFirewall``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFirewall`: GetLinodeFirewalls200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PutFirewall`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFirewallRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putFirewallRequest** | [**PutFirewallRequest**](PutFirewallRequest.md) | The Firewall information to update. | 

### Return type

[**GetLinodeFirewalls200ResponseDataInner**](GetLinodeFirewalls200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutFirewallRules

> GetLinodeFirewalls200ResponseDataInnerRules PutFirewallRules(ctx, apiVersion, firewallId).PutFirewallRulesRequest(putFirewallRulesRequest).Execute()

Update firewall rules



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
	firewallId := int32(56) // int32 | ID of the Firewall to access.
	putFirewallRulesRequest := *openapiclient.NewPutFirewallRulesRequest() // PutFirewallRulesRequest | The Firewall Rules information to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PutFirewallRules(context.Background(), apiVersion, firewallId).PutFirewallRulesRequest(putFirewallRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PutFirewallRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutFirewallRules`: GetLinodeFirewalls200ResponseDataInnerRules
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PutFirewallRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**firewallId** | **int32** | ID of the Firewall to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutFirewallRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putFirewallRulesRequest** | [**PutFirewallRulesRequest**](PutFirewallRulesRequest.md) | The Firewall Rules information to update. | 

### Return type

[**GetLinodeFirewalls200ResponseDataInnerRules**](GetLinodeFirewalls200ResponseDataInnerRules.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutIp

> GetLinodeIps200ResponseIpv4PublicInner PutIp(ctx, apiVersion, address).PutIpRequest(putIpRequest).Execute()

Update an IP address's RDNS



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
	address := "address_example" // string | The address to operate on.
	putIpRequest := *openapiclient.NewPutIpRequest("test.example.org") // PutIpRequest | The information to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworkingAPI.PutIp(context.Background(), apiVersion, address).PutIpRequest(putIpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworkingAPI.PutIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutIp`: GetLinodeIps200ResponseIpv4PublicInner
	fmt.Fprintf(os.Stdout, "Response from `NetworkingAPI.PutIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**address** | **string** | The address to operate on. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putIpRequest** | [**PutIpRequest**](PutIpRequest.md) | The information to update. | 

### Return type

[**GetLinodeIps200ResponseIpv4PublicInner**](GetLinodeIps200ResponseIpv4PublicInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

