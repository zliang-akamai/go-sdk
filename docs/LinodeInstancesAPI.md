# \LinodeInstancesAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDisk**](LinodeInstancesAPI.md#DeleteDisk) | **Delete** /{apiVersion}/linode/instances/{linodeId}/disks/{diskId} | Delete a disk
[**DeleteLinodeConfig**](LinodeInstancesAPI.md#DeleteLinodeConfig) | **Delete** /{apiVersion}/linode/instances/{linodeId}/configs/{configId} | Delete a config profile
[**DeleteLinodeConfigInterface**](LinodeInstancesAPI.md#DeleteLinodeConfigInterface) | **Delete** /{apiVersion}/linode/instances/{linodeId}/configs/{configId}/interfaces/{interfaceId} | Delete a configuration profile interface
[**DeleteLinodeInstance**](LinodeInstancesAPI.md#DeleteLinodeInstance) | **Delete** /{apiVersion}/linode/instances/{linodeId} | Delete a Linode
[**DeleteLinodeIp**](LinodeInstancesAPI.md#DeleteLinodeIp) | **Delete** /{apiVersion}/linode/instances/{linodeId}/ips/{address} | Delete an IPv4 address
[**GetBackup**](LinodeInstancesAPI.md#GetBackup) | **Get** /{apiVersion}/linode/instances/{linodeId}/backups/{backupId} | Get a backup
[**GetBackups**](LinodeInstancesAPI.md#GetBackups) | **Get** /{apiVersion}/linode/instances/{linodeId}/backups | List backups
[**GetKernel**](LinodeInstancesAPI.md#GetKernel) | **Get** /{apiVersion}/linode/kernels/{kernelId} | Get a kernel
[**GetKernels**](LinodeInstancesAPI.md#GetKernels) | **Get** /{apiVersion}/linode/kernels | List kernels
[**GetLinodeConfig**](LinodeInstancesAPI.md#GetLinodeConfig) | **Get** /{apiVersion}/linode/instances/{linodeId}/configs/{configId} | Get a config profile
[**GetLinodeConfigInterface**](LinodeInstancesAPI.md#GetLinodeConfigInterface) | **Get** /{apiVersion}/linode/instances/{linodeId}/configs/{configId}/interfaces/{interfaceId} | Get a configuration profile interface
[**GetLinodeConfigInterfaces**](LinodeInstancesAPI.md#GetLinodeConfigInterfaces) | **Get** /{apiVersion}/linode/instances/{linodeId}/configs/{configId}/interfaces | List configuration profile interfaces
[**GetLinodeConfigs**](LinodeInstancesAPI.md#GetLinodeConfigs) | **Get** /{apiVersion}/linode/instances/{linodeId}/configs | List config profiles
[**GetLinodeDisk**](LinodeInstancesAPI.md#GetLinodeDisk) | **Get** /{apiVersion}/linode/instances/{linodeId}/disks/{diskId} | Get a disk
[**GetLinodeDisks**](LinodeInstancesAPI.md#GetLinodeDisks) | **Get** /{apiVersion}/linode/instances/{linodeId}/disks | List disks
[**GetLinodeFirewalls**](LinodeInstancesAPI.md#GetLinodeFirewalls) | **Get** /{apiVersion}/linode/instances/{linodeId}/firewalls | List a Linode&#39;s firewalls
[**GetLinodeInstance**](LinodeInstancesAPI.md#GetLinodeInstance) | **Get** /{apiVersion}/linode/instances/{linodeId} | Get a Linode
[**GetLinodeInstances**](LinodeInstancesAPI.md#GetLinodeInstances) | **Get** /{apiVersion}/linode/instances | List Linodes
[**GetLinodeIp**](LinodeInstancesAPI.md#GetLinodeIp) | **Get** /{apiVersion}/linode/instances/{linodeId}/ips/{address} | Get a Linode&#39;s IP address
[**GetLinodeIps**](LinodeInstancesAPI.md#GetLinodeIps) | **Get** /{apiVersion}/linode/instances/{linodeId}/ips | Get networking information
[**GetLinodeNodeBalancers**](LinodeInstancesAPI.md#GetLinodeNodeBalancers) | **Get** /{apiVersion}/linode/instances/{linodeId}/nodebalancers | List Linode NodeBalancers
[**GetLinodeStats**](LinodeInstancesAPI.md#GetLinodeStats) | **Get** /{apiVersion}/linode/instances/{linodeId}/stats | Get Linode statistics
[**GetLinodeStatsByYearMonth**](LinodeInstancesAPI.md#GetLinodeStatsByYearMonth) | **Get** /{apiVersion}/linode/instances/{linodeId}/stats/{year}/{month} | Get monthly statistics
[**GetLinodeTransfer**](LinodeInstancesAPI.md#GetLinodeTransfer) | **Get** /{apiVersion}/linode/instances/{linodeId}/transfer | Get a network transfer
[**GetLinodeTransferByYearMonth**](LinodeInstancesAPI.md#GetLinodeTransferByYearMonth) | **Get** /{apiVersion}/linode/instances/{linodeId}/transfer/{year}/{month} | Get monthly network transfer stats
[**GetLinodeType**](LinodeInstancesAPI.md#GetLinodeType) | **Get** /{apiVersion}/linode/types/{typeId} | Get a type
[**GetLinodeTypes**](LinodeInstancesAPI.md#GetLinodeTypes) | **Get** /{apiVersion}/linode/types | List types
[**GetLinodeVolumes**](LinodeInstancesAPI.md#GetLinodeVolumes) | **Get** /{apiVersion}/linode/instances/{linodeId}/volumes | List a Linode&#39;s volumes
[**PostAddLinodeConfig**](LinodeInstancesAPI.md#PostAddLinodeConfig) | **Post** /{apiVersion}/linode/instances/{linodeId}/configs | Create a config profile
[**PostAddLinodeDisk**](LinodeInstancesAPI.md#PostAddLinodeDisk) | **Post** /{apiVersion}/linode/instances/{linodeId}/disks | Create a disk
[**PostAddLinodeIp**](LinodeInstancesAPI.md#PostAddLinodeIp) | **Post** /{apiVersion}/linode/instances/{linodeId}/ips | Allocate an IPv4 address
[**PostApplyFirewalls**](LinodeInstancesAPI.md#PostApplyFirewalls) | **Post** /{apiVersion}/linode/instances/{linodeId}/firewalls/apply | Apply a Linode&#39;s firewalls
[**PostBootLinodeInstance**](LinodeInstancesAPI.md#PostBootLinodeInstance) | **Post** /{apiVersion}/linode/instances/{linodeId}/boot | Boot a Linode
[**PostCancelBackups**](LinodeInstancesAPI.md#PostCancelBackups) | **Post** /{apiVersion}/linode/instances/{linodeId}/backups/cancel | Cancel backups
[**PostCloneLinodeDisk**](LinodeInstancesAPI.md#PostCloneLinodeDisk) | **Post** /{apiVersion}/linode/instances/{linodeId}/disks/{diskId}/clone | Clone a disk
[**PostCloneLinodeInstance**](LinodeInstancesAPI.md#PostCloneLinodeInstance) | **Post** /{apiVersion}/linode/instances/{linodeId}/clone | Clone a Linode
[**PostEnableBackups**](LinodeInstancesAPI.md#PostEnableBackups) | **Post** /{apiVersion}/linode/instances/{linodeId}/backups/enable | Enable backups
[**PostLinodeConfigInterface**](LinodeInstancesAPI.md#PostLinodeConfigInterface) | **Post** /{apiVersion}/linode/instances/{linodeId}/configs/{configId}/interfaces | Add a configuration profile interface
[**PostLinodeConfigInterfaces**](LinodeInstancesAPI.md#PostLinodeConfigInterfaces) | **Post** /{apiVersion}/linode/instances/{linodeId}/configs/{configId}/interfaces/order | Reorder configuration profile interfaces
[**PostLinodeInstance**](LinodeInstancesAPI.md#PostLinodeInstance) | **Post** /{apiVersion}/linode/instances | Create a Linode
[**PostMigrateLinodeInstance**](LinodeInstancesAPI.md#PostMigrateLinodeInstance) | **Post** /{apiVersion}/linode/instances/{linodeId}/migrate | Initiate a DC migration/pending host migration
[**PostMutateLinodeInstance**](LinodeInstancesAPI.md#PostMutateLinodeInstance) | **Post** /{apiVersion}/linode/instances/{linodeId}/mutate | Upgrade a Linode
[**PostRebootLinodeInstance**](LinodeInstancesAPI.md#PostRebootLinodeInstance) | **Post** /{apiVersion}/linode/instances/{linodeId}/reboot | Reboot a Linode
[**PostRebuildLinodeInstance**](LinodeInstancesAPI.md#PostRebuildLinodeInstance) | **Post** /{apiVersion}/linode/instances/{linodeId}/rebuild | Rebuild a Linode
[**PostRescueLinodeInstance**](LinodeInstancesAPI.md#PostRescueLinodeInstance) | **Post** /{apiVersion}/linode/instances/{linodeId}/rescue | Boot a Linode into rescue mode
[**PostResetDiskPassword**](LinodeInstancesAPI.md#PostResetDiskPassword) | **Post** /{apiVersion}/linode/instances/{linodeId}/disks/{diskId}/password | Reset a disk root password
[**PostResetLinodePassword**](LinodeInstancesAPI.md#PostResetLinodePassword) | **Post** /{apiVersion}/linode/instances/{linodeId}/password | Reset a Linode&#39;s root password
[**PostResizeDisk**](LinodeInstancesAPI.md#PostResizeDisk) | **Post** /{apiVersion}/linode/instances/{linodeId}/disks/{diskId}/resize | Resize a disk
[**PostResizeLinodeInstance**](LinodeInstancesAPI.md#PostResizeLinodeInstance) | **Post** /{apiVersion}/linode/instances/{linodeId}/resize | Resize a Linode
[**PostRestoreBackup**](LinodeInstancesAPI.md#PostRestoreBackup) | **Post** /{apiVersion}/linode/instances/{linodeId}/backups/{backupId}/restore | Restore a backup
[**PostShutdownLinodeInstance**](LinodeInstancesAPI.md#PostShutdownLinodeInstance) | **Post** /{apiVersion}/linode/instances/{linodeId}/shutdown | Shut down a Linode
[**PostSnapshot**](LinodeInstancesAPI.md#PostSnapshot) | **Post** /{apiVersion}/linode/instances/{linodeId}/backups | Create a snapshot
[**PutDisk**](LinodeInstancesAPI.md#PutDisk) | **Put** /{apiVersion}/linode/instances/{linodeId}/disks/{diskId} | Update a disk
[**PutLinodeConfig**](LinodeInstancesAPI.md#PutLinodeConfig) | **Put** /{apiVersion}/linode/instances/{linodeId}/configs/{configId} | Update a config profile
[**PutLinodeConfigInterface**](LinodeInstancesAPI.md#PutLinodeConfigInterface) | **Put** /{apiVersion}/linode/instances/{linodeId}/configs/{configId}/interfaces/{interfaceId} | Update a configuration profile interface
[**PutLinodeInstance**](LinodeInstancesAPI.md#PutLinodeInstance) | **Put** /{apiVersion}/linode/instances/{linodeId} | Update a Linode
[**PutLinodeIp**](LinodeInstancesAPI.md#PutLinodeIp) | **Put** /{apiVersion}/linode/instances/{linodeId}/ips/{address} | Update an IP address&#39;s RDNS for a Linode



## DeleteDisk

> map[string]interface{} DeleteDisk(ctx, apiVersion, linodeId, diskId).Execute()

Delete a disk



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	diskId := int32(56) // int32 | ID of the Disk to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.DeleteDisk(context.Background(), apiVersion, linodeId, diskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.DeleteDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDisk`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.DeleteDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiskRequest struct via the builder pattern


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


## DeleteLinodeConfig

> map[string]interface{} DeleteLinodeConfig(ctx, apiVersion, linodeId, configId).Execute()

Delete a config profile



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
	linodeId := int32(56) // int32 | The `id` of the Linode.
	configId := int32(56) // int32 | The `id` of the Configuration Profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.DeleteLinodeConfig(context.Background(), apiVersion, linodeId, configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.DeleteLinodeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLinodeConfig`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.DeleteLinodeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The &#x60;id&#x60; of the Linode. | 
**configId** | **int32** | The &#x60;id&#x60; of the Configuration Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinodeConfigRequest struct via the builder pattern


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


## DeleteLinodeConfigInterface

> map[string]interface{} DeleteLinodeConfigInterface(ctx, apiVersion, linodeId, configId, interfaceId).Execute()

Delete a configuration profile interface



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
	linodeId := int32(56) // int32 | The `id` of the Linode.
	configId := int32(56) // int32 | The `id` of the Configuration Profile.
	interfaceId := int32(56) // int32 | The `id` of the Linode Configuration Profile Interface.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.DeleteLinodeConfigInterface(context.Background(), apiVersion, linodeId, configId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.DeleteLinodeConfigInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLinodeConfigInterface`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.DeleteLinodeConfigInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The &#x60;id&#x60; of the Linode. | 
**configId** | **int32** | The &#x60;id&#x60; of the Configuration Profile. | 
**interfaceId** | **int32** | The &#x60;id&#x60; of the Linode Configuration Profile Interface. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinodeConfigInterfaceRequest struct via the builder pattern


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


## DeleteLinodeInstance

> map[string]interface{} DeleteLinodeInstance(ctx, apiVersion, linodeId).Execute()

Delete a Linode



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.DeleteLinodeInstance(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.DeleteLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLinodeInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.DeleteLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinodeInstanceRequest struct via the builder pattern


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


## DeleteLinodeIp

> map[string]interface{} DeleteLinodeIp(ctx, apiVersion, linodeId, address).Execute()

Delete an IPv4 address



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
	linodeId := int32(56) // int32 | The ID of the Linode.
	address := "address_example" // string | The IP address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.DeleteLinodeIp(context.Background(), apiVersion, linodeId, address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.DeleteLinodeIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteLinodeIp`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.DeleteLinodeIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode. | 
**address** | **string** | The IP address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinodeIpRequest struct via the builder pattern


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


## GetBackup

> GetBackups200ResponseSnapshotCurrent GetBackup(ctx, apiVersion, linodeId, backupId).Execute()

Get a backup



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
	linodeId := int32(56) // int32 | The ID of the Linode the Backup belongs to.
	backupId := int32(56) // int32 | The ID of the Backup to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetBackup(context.Background(), apiVersion, linodeId, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackup`: GetBackups200ResponseSnapshotCurrent
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode the Backup belongs to. | 
**backupId** | **int32** | The ID of the Backup to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetBackups200ResponseSnapshotCurrent**](GetBackups200ResponseSnapshotCurrent.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBackups

> GetBackups200Response GetBackups(ctx, apiVersion, linodeId).Execute()

List backups



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
	linodeId := int32(56) // int32 | The ID of the Linode the backups belong to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetBackups(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBackups`: GetBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode the backups belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetBackups200Response**](GetBackups200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKernel

> GetKernels200ResponseDataInner GetKernel(ctx, apiVersion, kernelId).Execute()

Get a kernel



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
	kernelId := "kernelId_example" // string | ID of the Kernel to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetKernel(context.Background(), apiVersion, kernelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetKernel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKernel`: GetKernels200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetKernel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**kernelId** | **string** | ID of the Kernel to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetKernelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetKernels200ResponseDataInner**](GetKernels200ResponseDataInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKernels

> GetKernels200Response GetKernels(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List kernels



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
	resp, r, err := apiClient.LinodeInstancesAPI.GetKernels(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetKernels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetKernels`: GetKernels200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetKernels`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetKernelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetKernels200Response**](GetKernels200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeConfig

> GetLinodeConfigs200ResponseDataInner GetLinodeConfig(ctx, apiVersion, linodeId, configId).Execute()

Get a config profile



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
	linodeId := int32(56) // int32 | The `id` of the Linode.
	configId := int32(56) // int32 | The `id` of the Configuration Profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeConfig(context.Background(), apiVersion, linodeId, configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeConfig`: GetLinodeConfigs200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The &#x60;id&#x60; of the Linode. | 
**configId** | **int32** | The &#x60;id&#x60; of the Configuration Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetLinodeConfigs200ResponseDataInner**](GetLinodeConfigs200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeConfigInterface

> PostLinodeInstanceRequestAllOfInterfacesInner GetLinodeConfigInterface(ctx, apiVersion, linodeId, configId, interfaceId).Execute()

Get a configuration profile interface



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
	linodeId := int32(56) // int32 | The `id` of the Linode.
	configId := int32(56) // int32 | The `id` of the Configuration Profile.
	interfaceId := int32(56) // int32 | The `id` of the Linode Configuration Profile Interface.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeConfigInterface(context.Background(), apiVersion, linodeId, configId, interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeConfigInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeConfigInterface`: PostLinodeInstanceRequestAllOfInterfacesInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeConfigInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The &#x60;id&#x60; of the Linode. | 
**configId** | **int32** | The &#x60;id&#x60; of the Configuration Profile. | 
**interfaceId** | **int32** | The &#x60;id&#x60; of the Linode Configuration Profile Interface. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeConfigInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**PostLinodeInstanceRequestAllOfInterfacesInner**](PostLinodeInstanceRequestAllOfInterfacesInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeConfigInterfaces

> []PostLinodeInstanceRequestAllOfInterfacesInner GetLinodeConfigInterfaces(ctx, apiVersion, linodeId, configId).Execute()

List configuration profile interfaces



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
	linodeId := int32(56) // int32 | The `id` of the Linode.
	configId := int32(56) // int32 | The `id` of the Configuration Profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeConfigInterfaces(context.Background(), apiVersion, linodeId, configId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeConfigInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeConfigInterfaces`: []PostLinodeInstanceRequestAllOfInterfacesInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeConfigInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The &#x60;id&#x60; of the Linode. | 
**configId** | **int32** | The &#x60;id&#x60; of the Configuration Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeConfigInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]PostLinodeInstanceRequestAllOfInterfacesInner**](PostLinodeInstanceRequestAllOfInterfacesInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeConfigs

> GetLinodeConfigs200Response GetLinodeConfigs(ctx, apiVersion, linodeId).Page(page).PageSize(pageSize).Execute()

List config profiles



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
	linodeId := int32(56) // int32 | ID of the Linode to look up Configuration profiles for.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeConfigs(context.Background(), apiVersion, linodeId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeConfigs`: GetLinodeConfigs200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeConfigs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up Configuration profiles for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetLinodeConfigs200Response**](GetLinodeConfigs200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeDisk

> GetLinodeDisks200ResponseDataInner GetLinodeDisk(ctx, apiVersion, linodeId, diskId).Execute()

Get a disk



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	diskId := int32(56) // int32 | ID of the Disk to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeDisk(context.Background(), apiVersion, linodeId, diskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeDisk`: GetLinodeDisks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetLinodeDisks200ResponseDataInner**](GetLinodeDisks200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeDisks

> GetLinodeDisks200Response GetLinodeDisks(ctx, apiVersion, linodeId).Page(page).PageSize(pageSize).Execute()

List disks



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeDisks(context.Background(), apiVersion, linodeId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeDisks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeDisks`: GetLinodeDisks200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeDisks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeDisksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetLinodeDisks200Response**](GetLinodeDisks200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeFirewalls

> GetLinodeFirewalls200Response GetLinodeFirewalls(ctx, apiVersion, linodeId).Page(page).PageSize(pageSize).Execute()

List a Linode's firewalls



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
	linodeId := int32(56) // int32 | ID of the Linode to access.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeFirewalls(context.Background(), apiVersion, linodeId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeFirewalls`: GetLinodeFirewalls200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeFirewalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to access. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeFirewallsRequest struct via the builder pattern


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


## GetLinodeInstance

> Linode GetLinodeInstance(ctx, apiVersion, linodeId).Execute()

Get a Linode



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeInstance(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeInstance`: Linode
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Linode**](Linode.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeInstances

> GetLinodeInstances200Response GetLinodeInstances(ctx, apiVersion).XFilter(xFilter).Page(page).PageSize(pageSize).Execute()

List Linodes



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
	xFilter := openapiclient.get_linode_instances_X_Filter_parameter{FilterAndSortCriteria: openapiclient.NewFilterAndSortCriteria()} // GetLinodeInstancesXFilterParameter | Specifies a JSON object to filter down the results. See [Filtering and sorting](filtering-and-sorting) for details. (optional)
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeInstances(context.Background(), apiVersion).XFilter(xFilter).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeInstances`: GetLinodeInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xFilter** | [**GetLinodeInstancesXFilterParameter**](GetLinodeInstancesXFilterParameter.md) | Specifies a JSON object to filter down the results. See [Filtering and sorting](filtering-and-sorting) for details. | 
 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetLinodeInstances200Response**](GetLinodeInstances200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeIp

> GetLinodeIps200ResponseIpv4PublicInner GetLinodeIp(ctx, apiVersion, linodeId, address).Execute()

Get a Linode's IP address



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
	linodeId := int32(56) // int32 | The ID of the Linode.
	address := "address_example" // string | The IP address.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeIp(context.Background(), apiVersion, linodeId, address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeIp`: GetLinodeIps200ResponseIpv4PublicInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode. | 
**address** | **string** | The IP address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeIpRequest struct via the builder pattern


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


## GetLinodeIps

> GetLinodeIps200Response GetLinodeIps(ctx, apiVersion, linodeId).Execute()

Get networking information



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeIps(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeIps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeIps`: GetLinodeIps200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeIps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeIpsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLinodeIps200Response**](GetLinodeIps200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeNodeBalancers

> GetLinodeNodeBalancers200Response GetLinodeNodeBalancers(ctx, apiVersion, linodeId).Execute()

List Linode NodeBalancers



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeNodeBalancers(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeNodeBalancers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeNodeBalancers`: GetLinodeNodeBalancers200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeNodeBalancers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeNodeBalancersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetLinodeStats

> GetLinodeStats200Response GetLinodeStats(ctx, apiVersion, linodeId).Execute()

Get Linode statistics



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeStats(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeStats`: GetLinodeStats200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeStats`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeStatsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLinodeStats200Response**](GetLinodeStats200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeStatsByYearMonth

> GetLinodeStats200Response GetLinodeStatsByYearMonth(ctx, apiVersion, linodeId, year, month).Execute()

Get monthly statistics



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	year := int32(56) // int32 | Numeric value representing the year to look up.
	month := int32(56) // int32 | Numeric value representing the month to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeStatsByYearMonth(context.Background(), apiVersion, linodeId, year, month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeStatsByYearMonth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeStatsByYearMonth`: GetLinodeStats200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeStatsByYearMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 
**year** | **int32** | Numeric value representing the year to look up. | 
**month** | **int32** | Numeric value representing the month to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeStatsByYearMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetLinodeStats200Response**](GetLinodeStats200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeTransfer

> GetLinodeTransfer200Response GetLinodeTransfer(ctx, apiVersion, linodeId).Execute()

Get a network transfer



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeTransfer(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeTransfer`: GetLinodeTransfer200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLinodeTransfer200Response**](GetLinodeTransfer200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeTransferByYearMonth

> GetLinodeTransferByYearMonth200Response GetLinodeTransferByYearMonth(ctx, apiVersion, linodeId, year, month).Execute()

Get monthly network transfer stats



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	year := int32(56) // int32 | Numeric value representing the year to look up.
	month := int32(56) // int32 | Numeric value representing the month to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeTransferByYearMonth(context.Background(), apiVersion, linodeId, year, month).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeTransferByYearMonth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeTransferByYearMonth`: GetLinodeTransferByYearMonth200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeTransferByYearMonth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 
**year** | **int32** | Numeric value representing the year to look up. | 
**month** | **int32** | Numeric value representing the month to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeTransferByYearMonthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**GetLinodeTransferByYearMonth200Response**](GetLinodeTransferByYearMonth200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeType

> GetLinodeTypes200ResponseDataInner GetLinodeType(ctx, apiVersion, typeId).Execute()

Get a type



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
	typeId := "typeId_example" // string | The ID of the Linode Type to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeType(context.Background(), apiVersion, typeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeType`: GetLinodeTypes200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**typeId** | **string** | The ID of the Linode Type to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetLinodeTypes200ResponseDataInner**](GetLinodeTypes200ResponseDataInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeTypes

> GetLinodeTypes200Response GetLinodeTypes(ctx, apiVersion).Execute()

List types



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
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeTypes(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeTypes`: GetLinodeTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetLinodeTypes200Response**](GetLinodeTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinodeVolumes

> GetLinodeVolumes200Response GetLinodeVolumes(ctx, apiVersion, linodeId).Page(page).PageSize(pageSize).Execute()

List a Linode's volumes



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.GetLinodeVolumes(context.Background(), apiVersion, linodeId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.GetLinodeVolumes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinodeVolumes`: GetLinodeVolumes200Response
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.GetLinodeVolumes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinodeVolumesRequest struct via the builder pattern


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


## PostAddLinodeConfig

> GetLinodeConfigs200ResponseDataInner PostAddLinodeConfig(ctx, apiVersion, linodeId).PostAddLinodeConfigRequest(postAddLinodeConfigRequest).Execute()

Create a config profile



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
	linodeId := int32(56) // int32 | ID of the Linode to look up Configuration profiles for.
	postAddLinodeConfigRequest := *openapiclient.NewPostAddLinodeConfigRequest(*openapiclient.NewGetLinodeConfigs200ResponseDataInnerDevices(), "My Config") // PostAddLinodeConfigRequest | The parameters to set when creating the Configuration profile. This determines which kernel, devices, how much memory, etc. a Linode boots with.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostAddLinodeConfig(context.Background(), apiVersion, linodeId).PostAddLinodeConfigRequest(postAddLinodeConfigRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostAddLinodeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAddLinodeConfig`: GetLinodeConfigs200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostAddLinodeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up Configuration profiles for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAddLinodeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postAddLinodeConfigRequest** | [**PostAddLinodeConfigRequest**](PostAddLinodeConfigRequest.md) | The parameters to set when creating the Configuration profile. This determines which kernel, devices, how much memory, etc. a Linode boots with. | 

### Return type

[**GetLinodeConfigs200ResponseDataInner**](GetLinodeConfigs200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAddLinodeDisk

> GetLinodeDisks200ResponseDataInner PostAddLinodeDisk(ctx, apiVersion, linodeId).PostAddLinodeDiskRequest(postAddLinodeDiskRequest).Execute()

Create a disk



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	postAddLinodeDiskRequest := *openapiclient.NewPostAddLinodeDiskRequest(int32(48640)) // PostAddLinodeDiskRequest | The parameters to set when creating the Disk.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostAddLinodeDisk(context.Background(), apiVersion, linodeId).PostAddLinodeDiskRequest(postAddLinodeDiskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostAddLinodeDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAddLinodeDisk`: GetLinodeDisks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostAddLinodeDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAddLinodeDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postAddLinodeDiskRequest** | [**PostAddLinodeDiskRequest**](PostAddLinodeDiskRequest.md) | The parameters to set when creating the Disk. | 

### Return type

[**GetLinodeDisks200ResponseDataInner**](GetLinodeDisks200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAddLinodeIp

> GetLinodeIps200ResponseIpv4PublicInner PostAddLinodeIp(ctx, apiVersion, linodeId).PostAddLinodeIpRequest(postAddLinodeIpRequest).Execute()

Allocate an IPv4 address



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	postAddLinodeIpRequest := *openapiclient.NewPostAddLinodeIpRequest(true, "ipv4") // PostAddLinodeIpRequest | Information about the address you are creating.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostAddLinodeIp(context.Background(), apiVersion, linodeId).PostAddLinodeIpRequest(postAddLinodeIpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostAddLinodeIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostAddLinodeIp`: GetLinodeIps200ResponseIpv4PublicInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostAddLinodeIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostAddLinodeIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postAddLinodeIpRequest** | [**PostAddLinodeIpRequest**](PostAddLinodeIpRequest.md) | Information about the address you are creating. | 

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


## PostApplyFirewalls

> map[string]interface{} PostApplyFirewalls(ctx, apiVersion, linodeId).Execute()

Apply a Linode's firewalls



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
	linodeId := int32(56) // int32 | The ID of the Linode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostApplyFirewalls(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostApplyFirewalls``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostApplyFirewalls`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostApplyFirewalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostApplyFirewallsRequest struct via the builder pattern


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


## PostBootLinodeInstance

> map[string]interface{} PostBootLinodeInstance(ctx, apiVersion, linodeId).PostBootLinodeInstanceRequest(postBootLinodeInstanceRequest).Execute()

Boot a Linode



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
	linodeId := int32(56) // int32 | The ID of the Linode to boot.
	postBootLinodeInstanceRequest := *openapiclient.NewPostBootLinodeInstanceRequest() // PostBootLinodeInstanceRequest | Optional configuration to boot into (see above). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostBootLinodeInstance(context.Background(), apiVersion, linodeId).PostBootLinodeInstanceRequest(postBootLinodeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostBootLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostBootLinodeInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostBootLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode to boot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostBootLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postBootLinodeInstanceRequest** | [**PostBootLinodeInstanceRequest**](PostBootLinodeInstanceRequest.md) | Optional configuration to boot into (see above). | 

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


## PostCancelBackups

> map[string]interface{} PostCancelBackups(ctx, apiVersion, linodeId).Execute()

Cancel backups



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
	linodeId := int32(56) // int32 | The ID of the Linode to cancel backup service for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostCancelBackups(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostCancelBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCancelBackups`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostCancelBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode to cancel backup service for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCancelBackupsRequest struct via the builder pattern


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


## PostCloneLinodeDisk

> GetLinodeDisks200ResponseDataInner PostCloneLinodeDisk(ctx, apiVersion, linodeId, diskId).Execute()

Clone a disk



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	diskId := int32(56) // int32 | ID of the Disk to clone.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostCloneLinodeDisk(context.Background(), apiVersion, linodeId, diskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostCloneLinodeDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloneLinodeDisk`: GetLinodeDisks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostCloneLinodeDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloneLinodeDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetLinodeDisks200ResponseDataInner**](GetLinodeDisks200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCloneLinodeInstance

> Linode1 PostCloneLinodeInstance(ctx, apiVersion, linodeId).PostCloneLinodeInstanceRequest(postCloneLinodeInstanceRequest).Execute()

Clone a Linode



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
	linodeId := int32(56) // int32 | ID of the Linode to clone.
	postCloneLinodeInstanceRequest := *openapiclient.NewPostCloneLinodeInstanceRequest() // PostCloneLinodeInstanceRequest | The requested state your Linode will be cloned into.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostCloneLinodeInstance(context.Background(), apiVersion, linodeId).PostCloneLinodeInstanceRequest(postCloneLinodeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostCloneLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCloneLinodeInstance`: Linode1
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostCloneLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to clone. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostCloneLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postCloneLinodeInstanceRequest** | [**PostCloneLinodeInstanceRequest**](PostCloneLinodeInstanceRequest.md) | The requested state your Linode will be cloned into. | 

### Return type

[**Linode1**](Linode1.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostEnableBackups

> map[string]interface{} PostEnableBackups(ctx, apiVersion, linodeId).Execute()

Enable backups



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
	linodeId := int32(56) // int32 | The ID of the Linode to enable backup service for.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostEnableBackups(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostEnableBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostEnableBackups`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostEnableBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode to enable backup service for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostEnableBackupsRequest struct via the builder pattern


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


## PostLinodeConfigInterface

> PostLinodeInstanceRequestAllOfInterfacesInner PostLinodeConfigInterface(ctx, apiVersion, linodeId, configId).PostLinodeConfigInterfaceRequest(postLinodeConfigInterfaceRequest).Execute()

Add a configuration profile interface



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
	linodeId := int32(56) // int32 | The `id` of the Linode.
	configId := int32(56) // int32 | The `id` of the Configuration Profile.
	postLinodeConfigInterfaceRequest := *openapiclient.NewPostLinodeConfigInterfaceRequest("vlan") // PostLinodeConfigInterfaceRequest | The Interface to add to the Configuration Profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostLinodeConfigInterface(context.Background(), apiVersion, linodeId, configId).PostLinodeConfigInterfaceRequest(postLinodeConfigInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostLinodeConfigInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLinodeConfigInterface`: PostLinodeInstanceRequestAllOfInterfacesInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostLinodeConfigInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The &#x60;id&#x60; of the Linode. | 
**configId** | **int32** | The &#x60;id&#x60; of the Configuration Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLinodeConfigInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postLinodeConfigInterfaceRequest** | [**PostLinodeConfigInterfaceRequest**](PostLinodeConfigInterfaceRequest.md) | The Interface to add to the Configuration Profile. | 

### Return type

[**PostLinodeInstanceRequestAllOfInterfacesInner**](PostLinodeInstanceRequestAllOfInterfacesInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostLinodeConfigInterfaces

> map[string]interface{} PostLinodeConfigInterfaces(ctx, apiVersion, linodeId, configId).PostLinodeConfigInterfacesRequest(postLinodeConfigInterfacesRequest).Execute()

Reorder configuration profile interfaces



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
	linodeId := int32(56) // int32 | The `id` of the Linode.
	configId := int32(56) // int32 | The `id` of the Configuration Profile.
	postLinodeConfigInterfacesRequest := *openapiclient.NewPostLinodeConfigInterfacesRequest([]int32{int32(101)}) // PostLinodeConfigInterfacesRequest | The desired Interface order for the Configuration Profile.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostLinodeConfigInterfaces(context.Background(), apiVersion, linodeId, configId).PostLinodeConfigInterfacesRequest(postLinodeConfigInterfacesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostLinodeConfigInterfaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLinodeConfigInterfaces`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostLinodeConfigInterfaces`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The &#x60;id&#x60; of the Linode. | 
**configId** | **int32** | The &#x60;id&#x60; of the Configuration Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostLinodeConfigInterfacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postLinodeConfigInterfacesRequest** | [**PostLinodeConfigInterfacesRequest**](PostLinodeConfigInterfacesRequest.md) | The desired Interface order for the Configuration Profile. | 

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


## PostLinodeInstance

> Linode1 PostLinodeInstance(ctx, apiVersion).PostLinodeInstanceRequest(postLinodeInstanceRequest).Execute()

Create a Linode



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
	postLinodeInstanceRequest := *openapiclient.NewPostLinodeInstanceRequest("us-east", "g6-standard-2") // PostLinodeInstanceRequest | The requested initial state of a new Linode.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostLinodeInstance(context.Background(), apiVersion).PostLinodeInstanceRequest(postLinodeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostLinodeInstance`: Linode1
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postLinodeInstanceRequest** | [**PostLinodeInstanceRequest**](PostLinodeInstanceRequest.md) | The requested initial state of a new Linode. | 

### Return type

[**Linode1**](Linode1.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostMigrateLinodeInstance

> map[string]interface{} PostMigrateLinodeInstance(ctx, apiVersion, linodeId).PostMigrateLinodeInstanceRequest(postMigrateLinodeInstanceRequest).Execute()

Initiate a DC migration/pending host migration



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
	linodeId := int32(56) // int32 | ID of the Linode to migrate.
	postMigrateLinodeInstanceRequest := *openapiclient.NewPostMigrateLinodeInstanceRequest() // PostMigrateLinodeInstanceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostMigrateLinodeInstance(context.Background(), apiVersion, linodeId).PostMigrateLinodeInstanceRequest(postMigrateLinodeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostMigrateLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMigrateLinodeInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostMigrateLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to migrate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMigrateLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postMigrateLinodeInstanceRequest** | [**PostMigrateLinodeInstanceRequest**](PostMigrateLinodeInstanceRequest.md) |  | 

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


## PostMutateLinodeInstance

> map[string]interface{} PostMutateLinodeInstance(ctx, apiVersion, linodeId).PostMutateLinodeInstanceRequest(postMutateLinodeInstanceRequest).Execute()

Upgrade a Linode



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
	linodeId := int32(56) // int32 | ID of the Linode to mutate.
	postMutateLinodeInstanceRequest := *openapiclient.NewPostMutateLinodeInstanceRequest() // PostMutateLinodeInstanceRequest | Whether to automatically resize disks or not. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostMutateLinodeInstance(context.Background(), apiVersion, linodeId).PostMutateLinodeInstanceRequest(postMutateLinodeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostMutateLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostMutateLinodeInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostMutateLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to mutate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostMutateLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postMutateLinodeInstanceRequest** | [**PostMutateLinodeInstanceRequest**](PostMutateLinodeInstanceRequest.md) | Whether to automatically resize disks or not. | 

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


## PostRebootLinodeInstance

> map[string]interface{} PostRebootLinodeInstance(ctx, apiVersion, linodeId).PostRebootLinodeInstanceRequest(postRebootLinodeInstanceRequest).Execute()

Reboot a Linode



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
	linodeId := int32(56) // int32 | ID of the linode to reboot.
	postRebootLinodeInstanceRequest := *openapiclient.NewPostRebootLinodeInstanceRequest() // PostRebootLinodeInstanceRequest | Optional reboot parameters. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostRebootLinodeInstance(context.Background(), apiVersion, linodeId).PostRebootLinodeInstanceRequest(postRebootLinodeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostRebootLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRebootLinodeInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostRebootLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the linode to reboot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRebootLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postRebootLinodeInstanceRequest** | [**PostRebootLinodeInstanceRequest**](PostRebootLinodeInstanceRequest.md) | Optional reboot parameters. | 

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


## PostRebuildLinodeInstance

> Linode1 PostRebuildLinodeInstance(ctx, apiVersion, linodeId).PostRebuildLinodeInstanceRequest(postRebuildLinodeInstanceRequest).Execute()

Rebuild a Linode



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
	linodeId := int32(56) // int32 | ID of the Linode to rebuild.
	postRebuildLinodeInstanceRequest := *openapiclient.NewPostRebuildLinodeInstanceRequest("linode/debian9", "aComplexP@ssword") // PostRebuildLinodeInstanceRequest | The requested state your Linode will be rebuilt into.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostRebuildLinodeInstance(context.Background(), apiVersion, linodeId).PostRebuildLinodeInstanceRequest(postRebuildLinodeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostRebuildLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRebuildLinodeInstance`: Linode1
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostRebuildLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to rebuild. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRebuildLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postRebuildLinodeInstanceRequest** | [**PostRebuildLinodeInstanceRequest**](PostRebuildLinodeInstanceRequest.md) | The requested state your Linode will be rebuilt into. | 

### Return type

[**Linode1**](Linode1.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRescueLinodeInstance

> map[string]interface{} PostRescueLinodeInstance(ctx, apiVersion, linodeId).PostRescueLinodeInstanceRequest(postRescueLinodeInstanceRequest).Execute()

Boot a Linode into rescue mode



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
	linodeId := int32(56) // int32 | ID of the Linode to rescue.
	postRescueLinodeInstanceRequest := *openapiclient.NewPostRescueLinodeInstanceRequest() // PostRescueLinodeInstanceRequest | Optional object of devices to be mounted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostRescueLinodeInstance(context.Background(), apiVersion, linodeId).PostRescueLinodeInstanceRequest(postRescueLinodeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostRescueLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRescueLinodeInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostRescueLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to rescue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRescueLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postRescueLinodeInstanceRequest** | [**PostRescueLinodeInstanceRequest**](PostRescueLinodeInstanceRequest.md) | Optional object of devices to be mounted. | 

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


## PostResetDiskPassword

> map[string]interface{} PostResetDiskPassword(ctx, apiVersion, linodeId, diskId).PostResetDiskPasswordRequest(postResetDiskPasswordRequest).Execute()

Reset a disk root password



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	diskId := int32(56) // int32 | ID of the Disk to look up.
	postResetDiskPasswordRequest := *openapiclient.NewPostResetDiskPasswordRequest("another@complex^Password123") // PostResetDiskPasswordRequest | The new password.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostResetDiskPassword(context.Background(), apiVersion, linodeId, diskId).PostResetDiskPasswordRequest(postResetDiskPasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostResetDiskPassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResetDiskPassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostResetDiskPassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostResetDiskPasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postResetDiskPasswordRequest** | [**PostResetDiskPasswordRequest**](PostResetDiskPasswordRequest.md) | The new password. | 

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


## PostResetLinodePassword

> map[string]interface{} PostResetLinodePassword(ctx, apiVersion, linodeId).PostResetLinodePasswordRequest(postResetLinodePasswordRequest).Execute()

Reset a Linode's root password



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
	linodeId := int32(56) // int32 | ID of the Linode for which to reset its root password.
	postResetLinodePasswordRequest := *openapiclient.NewPostResetLinodePasswordRequest("a$eCure4assw0rd!43v51") // PostResetLinodePasswordRequest | This Linode's new root password. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostResetLinodePassword(context.Background(), apiVersion, linodeId).PostResetLinodePasswordRequest(postResetLinodePasswordRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostResetLinodePassword``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResetLinodePassword`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostResetLinodePassword`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode for which to reset its root password. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostResetLinodePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postResetLinodePasswordRequest** | [**PostResetLinodePasswordRequest**](PostResetLinodePasswordRequest.md) | This Linode&#39;s new root password. | 

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


## PostResizeDisk

> map[string]interface{} PostResizeDisk(ctx, apiVersion, linodeId, diskId).PostResizeDiskRequest(postResizeDiskRequest).Execute()

Resize a disk



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	diskId := int32(56) // int32 | ID of the Disk to look up.
	postResizeDiskRequest := *openapiclient.NewPostResizeDiskRequest(int32(2048)) // PostResizeDiskRequest | The new size of the Disk.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostResizeDisk(context.Background(), apiVersion, linodeId, diskId).PostResizeDiskRequest(postResizeDiskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostResizeDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResizeDisk`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostResizeDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostResizeDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postResizeDiskRequest** | [**PostResizeDiskRequest**](PostResizeDiskRequest.md) | The new size of the Disk. | 

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


## PostResizeLinodeInstance

> map[string]interface{} PostResizeLinodeInstance(ctx, apiVersion, linodeId).PostResizeLinodeInstanceRequest(postResizeLinodeInstanceRequest).Execute()

Resize a Linode



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
	linodeId := int32(56) // int32 | ID of the Linode to resize.
	postResizeLinodeInstanceRequest := *openapiclient.NewPostResizeLinodeInstanceRequest("g6-standard-2") // PostResizeLinodeInstanceRequest | The Type your current Linode will resize to, and whether to attempt to automatically resize the Linode's disks.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostResizeLinodeInstance(context.Background(), apiVersion, linodeId).PostResizeLinodeInstanceRequest(postResizeLinodeInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostResizeLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostResizeLinodeInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostResizeLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to resize. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostResizeLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postResizeLinodeInstanceRequest** | [**PostResizeLinodeInstanceRequest**](PostResizeLinodeInstanceRequest.md) | The Type your current Linode will resize to, and whether to attempt to automatically resize the Linode&#39;s disks. | 

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


## PostRestoreBackup

> map[string]interface{} PostRestoreBackup(ctx, apiVersion, linodeId, backupId).PostRestoreBackupRequest(postRestoreBackupRequest).Execute()

Restore a backup



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
	linodeId := int32(56) // int32 | The ID of the Linode that the Backup belongs to.
	backupId := int32(56) // int32 | The ID of the Backup to restore.
	postRestoreBackupRequest := *openapiclient.NewPostRestoreBackupRequest(int32(234)) // PostRestoreBackupRequest | Parameters to provide when restoring the Backup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostRestoreBackup(context.Background(), apiVersion, linodeId, backupId).PostRestoreBackupRequest(postRestoreBackupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostRestoreBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostRestoreBackup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostRestoreBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode that the Backup belongs to. | 
**backupId** | **int32** | The ID of the Backup to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostRestoreBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postRestoreBackupRequest** | [**PostRestoreBackupRequest**](PostRestoreBackupRequest.md) | Parameters to provide when restoring the Backup. | 

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


## PostShutdownLinodeInstance

> map[string]interface{} PostShutdownLinodeInstance(ctx, apiVersion, linodeId).Execute()

Shut down a Linode



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
	linodeId := int32(56) // int32 | ID of the Linode to shutdown.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostShutdownLinodeInstance(context.Background(), apiVersion, linodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostShutdownLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostShutdownLinodeInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostShutdownLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to shutdown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostShutdownLinodeInstanceRequest struct via the builder pattern


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


## PostSnapshot

> GetBackups200ResponseSnapshotCurrent PostSnapshot(ctx, apiVersion, linodeId).PostSnapshotRequest(postSnapshotRequest).Execute()

Create a snapshot



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
	linodeId := int32(56) // int32 | The ID of the Linode the backups belong to.
	postSnapshotRequest := *openapiclient.NewPostSnapshotRequest("SnapshotLabel") // PostSnapshotRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PostSnapshot(context.Background(), apiVersion, linodeId).PostSnapshotRequest(postSnapshotRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PostSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostSnapshot`: GetBackups200ResponseSnapshotCurrent
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PostSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode the backups belong to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postSnapshotRequest** | [**PostSnapshotRequest**](PostSnapshotRequest.md) |  | 

### Return type

[**GetBackups200ResponseSnapshotCurrent**](GetBackups200ResponseSnapshotCurrent.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDisk

> GetLinodeDisks200ResponseDataInner PutDisk(ctx, apiVersion, linodeId, diskId).PutDiskRequest(putDiskRequest).Execute()

Update a disk



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	diskId := int32(56) // int32 | ID of the Disk to look up.
	putDiskRequest := *openapiclient.NewPutDiskRequest() // PutDiskRequest | Updates the parameters of a single Disk.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PutDisk(context.Background(), apiVersion, linodeId, diskId).PutDiskRequest(putDiskRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PutDisk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDisk`: GetLinodeDisks200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PutDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 
**diskId** | **int32** | ID of the Disk to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putDiskRequest** | [**PutDiskRequest**](PutDiskRequest.md) | Updates the parameters of a single Disk. | 

### Return type

[**GetLinodeDisks200ResponseDataInner**](GetLinodeDisks200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLinodeConfig

> GetLinodeConfigs200ResponseDataInner PutLinodeConfig(ctx, apiVersion, linodeId, configId).GetLinodeConfigs200ResponseDataInner(getLinodeConfigs200ResponseDataInner).Execute()

Update a config profile



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
	linodeId := int32(56) // int32 | The `id` of the Linode.
	configId := int32(56) // int32 | The `id` of the Configuration Profile.
	getLinodeConfigs200ResponseDataInner := *openapiclient.NewGetLinodeConfigs200ResponseDataInner() // GetLinodeConfigs200ResponseDataInner | The Configuration profile parameters to modify.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PutLinodeConfig(context.Background(), apiVersion, linodeId, configId).GetLinodeConfigs200ResponseDataInner(getLinodeConfigs200ResponseDataInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PutLinodeConfig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLinodeConfig`: GetLinodeConfigs200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PutLinodeConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The &#x60;id&#x60; of the Linode. | 
**configId** | **int32** | The &#x60;id&#x60; of the Configuration Profile. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLinodeConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **getLinodeConfigs200ResponseDataInner** | [**GetLinodeConfigs200ResponseDataInner**](GetLinodeConfigs200ResponseDataInner.md) | The Configuration profile parameters to modify. | 

### Return type

[**GetLinodeConfigs200ResponseDataInner**](GetLinodeConfigs200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLinodeConfigInterface

> PostLinodeInstanceRequestAllOfInterfacesInner PutLinodeConfigInterface(ctx, apiVersion, linodeId, configId, interfaceId).PutLinodeConfigInterfaceRequest(putLinodeConfigInterfaceRequest).Execute()

Update a configuration profile interface



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
	linodeId := int32(56) // int32 | The `id` of the Linode.
	configId := int32(56) // int32 | The `id` of the Configuration Profile.
	interfaceId := int32(56) // int32 | The `id` of the Linode Configuration Profile Interface.
	putLinodeConfigInterfaceRequest := *openapiclient.NewPutLinodeConfigInterfaceRequest() // PutLinodeConfigInterfaceRequest | The updated Interface.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PutLinodeConfigInterface(context.Background(), apiVersion, linodeId, configId, interfaceId).PutLinodeConfigInterfaceRequest(putLinodeConfigInterfaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PutLinodeConfigInterface``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLinodeConfigInterface`: PostLinodeInstanceRequestAllOfInterfacesInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PutLinodeConfigInterface`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The &#x60;id&#x60; of the Linode. | 
**configId** | **int32** | The &#x60;id&#x60; of the Configuration Profile. | 
**interfaceId** | **int32** | The &#x60;id&#x60; of the Linode Configuration Profile Interface. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLinodeConfigInterfaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **putLinodeConfigInterfaceRequest** | [**PutLinodeConfigInterfaceRequest**](PutLinodeConfigInterfaceRequest.md) | The updated Interface. | 

### Return type

[**PostLinodeInstanceRequestAllOfInterfacesInner**](PostLinodeInstanceRequestAllOfInterfacesInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLinodeInstance

> Linode1 PutLinodeInstance(ctx, apiVersion, linodeId).Linode1(linode1).Execute()

Update a Linode



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
	linodeId := int32(56) // int32 | ID of the Linode to look up.
	linode1 := *openapiclient.NewLinode1() // Linode1 | Any field that is not marked as `readOnly` may be updated. Fields that are marked `readOnly` will be ignored. If any updated field fails to pass validation, the Linode will not be updated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PutLinodeInstance(context.Background(), apiVersion, linodeId).Linode1(linode1).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PutLinodeInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLinodeInstance`: Linode1
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PutLinodeInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | ID of the Linode to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLinodeInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **linode1** | [**Linode1**](Linode1.md) | Any field that is not marked as &#x60;readOnly&#x60; may be updated. Fields that are marked &#x60;readOnly&#x60; will be ignored. If any updated field fails to pass validation, the Linode will not be updated. | 

### Return type

[**Linode1**](Linode1.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutLinodeIp

> GetLinodeIps200ResponseIpv4PublicInner PutLinodeIp(ctx, apiVersion, linodeId, address).PutLinodeIpRequest(putLinodeIpRequest).Execute()

Update an IP address's RDNS for a Linode



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
	linodeId := int32(56) // int32 | The ID of the Linode.
	address := "address_example" // string | The IP address.
	putLinodeIpRequest := *openapiclient.NewPutLinodeIpRequest("test.example.org") // PutLinodeIpRequest | The information to update for the IP address. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LinodeInstancesAPI.PutLinodeIp(context.Background(), apiVersion, linodeId, address).PutLinodeIpRequest(putLinodeIpRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LinodeInstancesAPI.PutLinodeIp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutLinodeIp`: GetLinodeIps200ResponseIpv4PublicInner
	fmt.Fprintf(os.Stdout, "Response from `LinodeInstancesAPI.PutLinodeIp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**linodeId** | **int32** | The ID of the Linode. | 
**address** | **string** | The IP address. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutLinodeIpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putLinodeIpRequest** | [**PutLinodeIpRequest**](PutLinodeIpRequest.md) | The information to update for the IP address. | 

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

