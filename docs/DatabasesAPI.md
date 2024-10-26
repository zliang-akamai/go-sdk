# \DatabasesAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDatabaseMysqlInstanceBackup**](DatabasesAPI.md#DeleteDatabaseMysqlInstanceBackup) | **Delete** /{apiVersion}/databases/mysql/instances/{instanceId}/backups/{backupId} | Delete a managed MySQL database backup
[**DeleteDatabasePostgreSqlInstanceBackup**](DatabasesAPI.md#DeleteDatabasePostgreSqlInstanceBackup) | **Delete** /{apiVersion}/databases/postgresql/instances/{instanceId}/backups/{backupId} | Delete a managed PostgreSQL database backup
[**DeleteDatabasesMysqlInstance**](DatabasesAPI.md#DeleteDatabasesMysqlInstance) | **Delete** /{apiVersion}/databases/mysql/instances/{instanceId} | Delete a managed MySQL database
[**DeleteDatabasesPostgreSqlInstance**](DatabasesAPI.md#DeleteDatabasesPostgreSqlInstance) | **Delete** /{apiVersion}/databases/postgresql/instances/{instanceId} | Delete a managed PostgreSQL database
[**GetDatabasesEngine**](DatabasesAPI.md#GetDatabasesEngine) | **Get** /{apiVersion}/databases/engines/{engineId} | Get a managed database engine
[**GetDatabasesEngines**](DatabasesAPI.md#GetDatabasesEngines) | **Get** /{apiVersion}/databases/engines | List managed database engines
[**GetDatabasesInstances**](DatabasesAPI.md#GetDatabasesInstances) | **Get** /{apiVersion}/databases/instances | List managed databases
[**GetDatabasesMysqlInstance**](DatabasesAPI.md#GetDatabasesMysqlInstance) | **Get** /{apiVersion}/databases/mysql/instances/{instanceId} | Get a managed MySQL database
[**GetDatabasesMysqlInstanceBackup**](DatabasesAPI.md#GetDatabasesMysqlInstanceBackup) | **Get** /{apiVersion}/databases/mysql/instances/{instanceId}/backups/{backupId} | Get a managed MySQL database backup
[**GetDatabasesMysqlInstanceBackups**](DatabasesAPI.md#GetDatabasesMysqlInstanceBackups) | **Get** /{apiVersion}/databases/mysql/instances/{instanceId}/backups | List managed MySQL database backups
[**GetDatabasesMysqlInstanceCredentials**](DatabasesAPI.md#GetDatabasesMysqlInstanceCredentials) | **Get** /{apiVersion}/databases/mysql/instances/{instanceId}/credentials | Get managed MySQL database credentials
[**GetDatabasesMysqlInstanceSsl**](DatabasesAPI.md#GetDatabasesMysqlInstanceSsl) | **Get** /{apiVersion}/databases/mysql/instances/{instanceId}/ssl | Get a managed MySQL database SSL certificate
[**GetDatabasesMysqlInstances**](DatabasesAPI.md#GetDatabasesMysqlInstances) | **Get** /{apiVersion}/databases/mysql/instances | List managed MySQL databases
[**GetDatabasesPostgreSqlInstance**](DatabasesAPI.md#GetDatabasesPostgreSqlInstance) | **Get** /{apiVersion}/databases/postgresql/instances/{instanceId} | Get a managed PostgreSQL database
[**GetDatabasesPostgreSqlInstanceBackups**](DatabasesAPI.md#GetDatabasesPostgreSqlInstanceBackups) | **Get** /{apiVersion}/databases/postgresql/instances/{instanceId}/backups | List managed PostgreSQL database backups
[**GetDatabasesPostgreSqlInstanceCredentials**](DatabasesAPI.md#GetDatabasesPostgreSqlInstanceCredentials) | **Get** /{apiVersion}/databases/postgresql/instances/{instanceId}/credentials | Get managed PostgreSQL database credentials
[**GetDatabasesPostgreSqlInstances**](DatabasesAPI.md#GetDatabasesPostgreSqlInstances) | **Get** /{apiVersion}/databases/postgresql/instances | List managed PostgreSQL databases
[**GetDatabasesPostgresqlInstanceBackup**](DatabasesAPI.md#GetDatabasesPostgresqlInstanceBackup) | **Get** /{apiVersion}/databases/postgresql/instances/{instanceId}/backups/{backupId} | Get a managed PostgreSQL database backup
[**GetDatabasesPostgresqlInstanceSsl**](DatabasesAPI.md#GetDatabasesPostgresqlInstanceSsl) | **Get** /{apiVersion}/databases/postgresql/instances/{instanceId}/ssl | Get a managed PostgreSQL database SSL certificate
[**GetDatabasesType**](DatabasesAPI.md#GetDatabasesType) | **Get** /{apiVersion}/databases/types/{typeId} | Get a managed database type
[**GetDatabasesTypes**](DatabasesAPI.md#GetDatabasesTypes) | **Get** /{apiVersion}/databases/types | List managed database types
[**PostDatabasesMysqlInstanceBackup**](DatabasesAPI.md#PostDatabasesMysqlInstanceBackup) | **Post** /{apiVersion}/databases/mysql/instances/{instanceId}/backups | Create a managed MySQL database backup snapshot
[**PostDatabasesMysqlInstanceBackupRestore**](DatabasesAPI.md#PostDatabasesMysqlInstanceBackupRestore) | **Post** /{apiVersion}/databases/mysql/instances/{instanceId}/backups/{backupId}/restore | Restore a managed MySQL database backup
[**PostDatabasesMysqlInstanceCredentialsReset**](DatabasesAPI.md#PostDatabasesMysqlInstanceCredentialsReset) | **Post** /{apiVersion}/databases/mysql/instances/{instanceId}/credentials/reset | Reset managed MySQL database credentials
[**PostDatabasesMysqlInstancePatch**](DatabasesAPI.md#PostDatabasesMysqlInstancePatch) | **Post** /{apiVersion}/databases/mysql/instances/{instanceId}/patch | Patch a managed MySQL database
[**PostDatabasesMysqlInstances**](DatabasesAPI.md#PostDatabasesMysqlInstances) | **Post** /{apiVersion}/databases/mysql/instances | Create a managed MySQL database
[**PostDatabasesPostgreSqlInstanceBackup**](DatabasesAPI.md#PostDatabasesPostgreSqlInstanceBackup) | **Post** /{apiVersion}/databases/postgresql/instances/{instanceId}/backups | Create a managed PostgreSQL database backup snapshot
[**PostDatabasesPostgreSqlInstanceBackupRestore**](DatabasesAPI.md#PostDatabasesPostgreSqlInstanceBackupRestore) | **Post** /{apiVersion}/databases/postgresql/instances/{instanceId}/backups/{backupId}/restore | Restore a managed PostgreSQL database backup
[**PostDatabasesPostgreSqlInstanceCredentialsReset**](DatabasesAPI.md#PostDatabasesPostgreSqlInstanceCredentialsReset) | **Post** /{apiVersion}/databases/postgresql/instances/{instanceId}/credentials/reset | Reset managed PostgreSQL database credentials
[**PostDatabasesPostgreSqlInstancePatch**](DatabasesAPI.md#PostDatabasesPostgreSqlInstancePatch) | **Post** /{apiVersion}/databases/postgresql/instances/{instanceId}/patch | Patch a managed PostgreSQL database
[**PostDatabasesPostgreSqlInstances**](DatabasesAPI.md#PostDatabasesPostgreSqlInstances) | **Post** /{apiVersion}/databases/postgresql/instances | Create a managed PostgreSQL database
[**PutDatabasesMysqlInstance**](DatabasesAPI.md#PutDatabasesMysqlInstance) | **Put** /{apiVersion}/databases/mysql/instances/{instanceId} | Update a managed MySQL database
[**PutDatabasesPostgreSqlInstance**](DatabasesAPI.md#PutDatabasesPostgreSqlInstance) | **Put** /{apiVersion}/databases/postgresql/instances/{instanceId} | Update a managed PostgreSQL database



## DeleteDatabaseMysqlInstanceBackup

> map[string]interface{} DeleteDatabaseMysqlInstanceBackup(ctx, apiVersion, instanceId, backupId).Execute()

Delete a managed MySQL database backup



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	backupId := int32(56) // int32 | The ID of the Managed MySQL Database backup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.DeleteDatabaseMysqlInstanceBackup(context.Background(), apiVersion, instanceId, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabaseMysqlInstanceBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDatabaseMysqlInstanceBackup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteDatabaseMysqlInstanceBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 
**backupId** | **int32** | The ID of the Managed MySQL Database backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabaseMysqlInstanceBackupRequest struct via the builder pattern


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


## DeleteDatabasePostgreSqlInstanceBackup

> map[string]interface{} DeleteDatabasePostgreSqlInstanceBackup(ctx, apiVersion, instanceId, backupId).Execute()

Delete a managed PostgreSQL database backup



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	backupId := int32(56) // int32 | The ID of the Managed PostgreSQL Database backup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.DeleteDatabasePostgreSqlInstanceBackup(context.Background(), apiVersion, instanceId, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabasePostgreSqlInstanceBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDatabasePostgreSqlInstanceBackup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteDatabasePostgreSqlInstanceBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 
**backupId** | **int32** | The ID of the Managed PostgreSQL Database backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabasePostgreSqlInstanceBackupRequest struct via the builder pattern


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


## DeleteDatabasesMysqlInstance

> map[string]interface{} DeleteDatabasesMysqlInstance(ctx, apiVersion, instanceId).Execute()

Delete a managed MySQL database



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.DeleteDatabasesMysqlInstance(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabasesMysqlInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDatabasesMysqlInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteDatabasesMysqlInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabasesMysqlInstanceRequest struct via the builder pattern


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


## DeleteDatabasesPostgreSqlInstance

> map[string]interface{} DeleteDatabasesPostgreSqlInstance(ctx, apiVersion, instanceId).Execute()

Delete a managed PostgreSQL database



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.DeleteDatabasesPostgreSqlInstance(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.DeleteDatabasesPostgreSqlInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDatabasesPostgreSqlInstance`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.DeleteDatabasesPostgreSqlInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDatabasesPostgreSqlInstanceRequest struct via the builder pattern


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


## GetDatabasesEngine

> GetDatabasesEngines200ResponseAllOfDataInner GetDatabasesEngine(ctx, apiVersion, engineId).Page(page).PageSize(pageSize).Execute()

Get a managed database engine



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
	engineId := "engineId_example" // string | The ID of the Managed Database engine.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesEngine(context.Background(), apiVersion, engineId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesEngine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesEngine`: GetDatabasesEngines200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesEngine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**engineId** | **string** | The ID of the Managed Database engine. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesEngineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetDatabasesEngines200ResponseAllOfDataInner**](GetDatabasesEngines200ResponseAllOfDataInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesEngines

> GetDatabasesEngines200Response GetDatabasesEngines(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List managed database engines



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
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesEngines(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesEngines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesEngines`: GetDatabasesEngines200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesEngines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesEnginesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetDatabasesEngines200Response**](GetDatabasesEngines200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesInstances

> GetDatabasesInstances200Response GetDatabasesInstances(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List managed databases



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
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesInstances(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesInstances`: GetDatabasesInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetDatabasesInstances200Response**](GetDatabasesInstances200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesMysqlInstance

> GetDatabasesMysqlInstances200ResponseAllOfDataInner GetDatabasesMysqlInstance(ctx, apiVersion, instanceId).Execute()

Get a managed MySQL database



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstance(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesMysqlInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesMysqlInstance`: GetDatabasesMysqlInstances200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesMysqlInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesMysqlInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDatabasesMysqlInstances200ResponseAllOfDataInner**](GetDatabasesMysqlInstances200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesMysqlInstanceBackup

> GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner GetDatabasesMysqlInstanceBackup(ctx, apiVersion, instanceId, backupId).Execute()

Get a managed MySQL database backup



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	backupId := int32(56) // int32 | The ID of the Managed MySQL Database backup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstanceBackup(context.Background(), apiVersion, instanceId, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesMysqlInstanceBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesMysqlInstanceBackup`: GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesMysqlInstanceBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 
**backupId** | **int32** | The ID of the Managed MySQL Database backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesMysqlInstanceBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner**](GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesMysqlInstanceBackups

> GetDatabasesMysqlInstanceBackups200Response GetDatabasesMysqlInstanceBackups(ctx, apiVersion, instanceId).Page(page).PageSize(pageSize).Execute()

List managed MySQL database backups



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstanceBackups(context.Background(), apiVersion, instanceId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesMysqlInstanceBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesMysqlInstanceBackups`: GetDatabasesMysqlInstanceBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesMysqlInstanceBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesMysqlInstanceBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetDatabasesMysqlInstanceBackups200Response**](GetDatabasesMysqlInstanceBackups200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesMysqlInstanceCredentials

> GetDatabasesMysqlInstanceCredentials200Response GetDatabasesMysqlInstanceCredentials(ctx, apiVersion, instanceId).Execute()

Get managed MySQL database credentials



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstanceCredentials(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesMysqlInstanceCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesMysqlInstanceCredentials`: GetDatabasesMysqlInstanceCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesMysqlInstanceCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesMysqlInstanceCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDatabasesMysqlInstanceCredentials200Response**](GetDatabasesMysqlInstanceCredentials200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesMysqlInstanceSsl

> GetDatabasesMysqlInstanceSsl200Response GetDatabasesMysqlInstanceSsl(ctx, apiVersion, instanceId).Execute()

Get a managed MySQL database SSL certificate



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
	instanceId := int32(56) // int32 | The ID of the Managed MySQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstanceSsl(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesMysqlInstanceSsl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesMysqlInstanceSsl`: GetDatabasesMysqlInstanceSsl200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesMysqlInstanceSsl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed MySQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesMysqlInstanceSslRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDatabasesMysqlInstanceSsl200Response**](GetDatabasesMysqlInstanceSsl200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesMysqlInstances

> GetDatabasesMysqlInstances200Response GetDatabasesMysqlInstances(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List managed MySQL databases



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
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstances(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesMysqlInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesMysqlInstances`: GetDatabasesMysqlInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesMysqlInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesMysqlInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetDatabasesMysqlInstances200Response**](GetDatabasesMysqlInstances200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesPostgreSqlInstance

> GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner GetDatabasesPostgreSqlInstance(ctx, apiVersion, instanceId).Execute()

Get a managed PostgreSQL database



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesPostgreSqlInstance(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesPostgreSqlInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesPostgreSqlInstance`: GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesPostgreSqlInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesPostgreSqlInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner**](GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesPostgreSqlInstanceBackups

> GetDatabasesMysqlInstanceBackups200Response GetDatabasesPostgreSqlInstanceBackups(ctx, apiVersion, instanceId).Page(page).PageSize(pageSize).Execute()

List managed PostgreSQL database backups



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesPostgreSqlInstanceBackups(context.Background(), apiVersion, instanceId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesPostgreSqlInstanceBackups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesPostgreSqlInstanceBackups`: GetDatabasesMysqlInstanceBackups200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesPostgreSqlInstanceBackups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesPostgreSqlInstanceBackupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetDatabasesMysqlInstanceBackups200Response**](GetDatabasesMysqlInstanceBackups200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesPostgreSqlInstanceCredentials

> GetDatabasesMysqlInstanceCredentials200Response GetDatabasesPostgreSqlInstanceCredentials(ctx, apiVersion, instanceId).Execute()

Get managed PostgreSQL database credentials



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesPostgreSqlInstanceCredentials(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesPostgreSqlInstanceCredentials``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesPostgreSqlInstanceCredentials`: GetDatabasesMysqlInstanceCredentials200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesPostgreSqlInstanceCredentials`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesPostgreSqlInstanceCredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDatabasesMysqlInstanceCredentials200Response**](GetDatabasesMysqlInstanceCredentials200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesPostgreSqlInstances

> GetDatabasesPostgreSqlInstances200Response GetDatabasesPostgreSqlInstances(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List managed PostgreSQL databases



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
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesPostgreSqlInstances(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesPostgreSqlInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesPostgreSqlInstances`: GetDatabasesPostgreSqlInstances200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesPostgreSqlInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesPostgreSqlInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetDatabasesPostgreSqlInstances200Response**](GetDatabasesPostgreSqlInstances200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesPostgresqlInstanceBackup

> GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner GetDatabasesPostgresqlInstanceBackup(ctx, apiVersion, instanceId, backupId).Execute()

Get a managed PostgreSQL database backup



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	backupId := int32(56) // int32 | The ID of the Managed PostgreSQL Database backup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesPostgresqlInstanceBackup(context.Background(), apiVersion, instanceId, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesPostgresqlInstanceBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesPostgresqlInstanceBackup`: GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesPostgresqlInstanceBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 
**backupId** | **int32** | The ID of the Managed PostgreSQL Database backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesPostgresqlInstanceBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner**](GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesPostgresqlInstanceSsl

> GetDatabasesMysqlInstanceSsl200Response GetDatabasesPostgresqlInstanceSsl(ctx, apiVersion, instanceId).Execute()

Get a managed PostgreSQL database SSL certificate



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesPostgresqlInstanceSsl(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesPostgresqlInstanceSsl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesPostgresqlInstanceSsl`: GetDatabasesMysqlInstanceSsl200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesPostgresqlInstanceSsl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesPostgresqlInstanceSslRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetDatabasesMysqlInstanceSsl200Response**](GetDatabasesMysqlInstanceSsl200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesType

> GetDatabasesTypes200ResponseAllOfDataInner GetDatabasesType(ctx, apiVersion, typeId).Page(page).PageSize(pageSize).Execute()

Get a managed database type



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
	typeId := "typeId_example" // string | The ID of the Managed Database type.
	page := int32(56) // int32 | The page of a collection to return. (optional) (default to 1)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesType(context.Background(), apiVersion, typeId).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesType`: GetDatabasesTypes200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**typeId** | **string** | The ID of the Managed Database type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetDatabasesTypes200ResponseAllOfDataInner**](GetDatabasesTypes200ResponseAllOfDataInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDatabasesTypes

> GetDatabasesTypes200Response GetDatabasesTypes(ctx, apiVersion).Page(page).PageSize(pageSize).Execute()

List managed database types



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
	resp, r, err := apiClient.DatabasesAPI.GetDatabasesTypes(context.Background(), apiVersion).Page(page).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.GetDatabasesTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDatabasesTypes`: GetDatabasesTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.GetDatabasesTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetDatabasesTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | The page of a collection to return. | [default to 1]
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetDatabasesTypes200Response**](GetDatabasesTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatabasesMysqlInstanceBackup

> map[string]interface{} PostDatabasesMysqlInstanceBackup(ctx, apiVersion, instanceId).PostDatabasesMysqlInstanceBackupRequest(postDatabasesMysqlInstanceBackupRequest).Execute()

Create a managed MySQL database backup snapshot



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	postDatabasesMysqlInstanceBackupRequest := *openapiclient.NewPostDatabasesMysqlInstanceBackupRequest("db-snapshot") // PostDatabasesMysqlInstanceBackupRequest | Information about the snapshot backup to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstanceBackup(context.Background(), apiVersion, instanceId).PostDatabasesMysqlInstanceBackupRequest(postDatabasesMysqlInstanceBackupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesMysqlInstanceBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesMysqlInstanceBackup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesMysqlInstanceBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesMysqlInstanceBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postDatabasesMysqlInstanceBackupRequest** | [**PostDatabasesMysqlInstanceBackupRequest**](PostDatabasesMysqlInstanceBackupRequest.md) | Information about the snapshot backup to create. | 

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


## PostDatabasesMysqlInstanceBackupRestore

> map[string]interface{} PostDatabasesMysqlInstanceBackupRestore(ctx, apiVersion, instanceId, backupId).Execute()

Restore a managed MySQL database backup



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
	instanceId := int32(56) // int32 | The ID of the Managed MySQL Database.
	backupId := int32(56) // int32 | The ID of the Managed MySQL Database backup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstanceBackupRestore(context.Background(), apiVersion, instanceId, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesMysqlInstanceBackupRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesMysqlInstanceBackupRestore`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesMysqlInstanceBackupRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed MySQL Database. | 
**backupId** | **int32** | The ID of the Managed MySQL Database backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesMysqlInstanceBackupRestoreRequest struct via the builder pattern


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


## PostDatabasesMysqlInstanceCredentialsReset

> map[string]interface{} PostDatabasesMysqlInstanceCredentialsReset(ctx, apiVersion, instanceId).Execute()

Reset managed MySQL database credentials



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
	instanceId := int32(56) // int32 | The ID of the Managed MySQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstanceCredentialsReset(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesMysqlInstanceCredentialsReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesMysqlInstanceCredentialsReset`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesMysqlInstanceCredentialsReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed MySQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesMysqlInstanceCredentialsResetRequest struct via the builder pattern


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


## PostDatabasesMysqlInstancePatch

> map[string]interface{} PostDatabasesMysqlInstancePatch(ctx, apiVersion, instanceId).Execute()

Patch a managed MySQL database



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstancePatch(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesMysqlInstancePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesMysqlInstancePatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesMysqlInstancePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesMysqlInstancePatchRequest struct via the builder pattern


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


## PostDatabasesMysqlInstances

> GetDatabasesMysqlInstances200ResponseAllOfDataInner PostDatabasesMysqlInstances(ctx, apiVersion).PostDatabasesMysqlInstancesRequest(postDatabasesMysqlInstancesRequest).Execute()

Create a managed MySQL database



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
	postDatabasesMysqlInstancesRequest := *openapiclient.NewPostDatabasesMysqlInstancesRequest("mysql/8.0.26", "example-db", "us-east", "g6-dedicated-2") // PostDatabasesMysqlInstancesRequest | Information about the Managed MySQL Database you are creating.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstances(context.Background(), apiVersion).PostDatabasesMysqlInstancesRequest(postDatabasesMysqlInstancesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesMysqlInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesMysqlInstances`: GetDatabasesMysqlInstances200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesMysqlInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesMysqlInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDatabasesMysqlInstancesRequest** | [**PostDatabasesMysqlInstancesRequest**](PostDatabasesMysqlInstancesRequest.md) | Information about the Managed MySQL Database you are creating. | 

### Return type

[**GetDatabasesMysqlInstances200ResponseAllOfDataInner**](GetDatabasesMysqlInstances200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostDatabasesPostgreSqlInstanceBackup

> map[string]interface{} PostDatabasesPostgreSqlInstanceBackup(ctx, apiVersion, instanceId).PostDatabasesMysqlInstanceBackupRequest(postDatabasesMysqlInstanceBackupRequest).Execute()

Create a managed PostgreSQL database backup snapshot



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	postDatabasesMysqlInstanceBackupRequest := *openapiclient.NewPostDatabasesMysqlInstanceBackupRequest("db-snapshot") // PostDatabasesMysqlInstanceBackupRequest | Information about the snapshot backup to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstanceBackup(context.Background(), apiVersion, instanceId).PostDatabasesMysqlInstanceBackupRequest(postDatabasesMysqlInstanceBackupRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesPostgreSqlInstanceBackup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesPostgreSqlInstanceBackup`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesPostgreSqlInstanceBackup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesPostgreSqlInstanceBackupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **postDatabasesMysqlInstanceBackupRequest** | [**PostDatabasesMysqlInstanceBackupRequest**](PostDatabasesMysqlInstanceBackupRequest.md) | Information about the snapshot backup to create. | 

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


## PostDatabasesPostgreSqlInstanceBackupRestore

> map[string]interface{} PostDatabasesPostgreSqlInstanceBackupRestore(ctx, apiVersion, instanceId, backupId).Execute()

Restore a managed PostgreSQL database backup



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	backupId := int32(56) // int32 | The ID of the Managed PostgreSQL Database backup.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstanceBackupRestore(context.Background(), apiVersion, instanceId, backupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesPostgreSqlInstanceBackupRestore``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesPostgreSqlInstanceBackupRestore`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesPostgreSqlInstanceBackupRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 
**backupId** | **int32** | The ID of the Managed PostgreSQL Database backup. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesPostgreSqlInstanceBackupRestoreRequest struct via the builder pattern


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


## PostDatabasesPostgreSqlInstanceCredentialsReset

> map[string]interface{} PostDatabasesPostgreSqlInstanceCredentialsReset(ctx, apiVersion, instanceId).Execute()

Reset managed PostgreSQL database credentials



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstanceCredentialsReset(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesPostgreSqlInstanceCredentialsReset``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesPostgreSqlInstanceCredentialsReset`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesPostgreSqlInstanceCredentialsReset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesPostgreSqlInstanceCredentialsResetRequest struct via the builder pattern


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


## PostDatabasesPostgreSqlInstancePatch

> map[string]interface{} PostDatabasesPostgreSqlInstancePatch(ctx, apiVersion, instanceId).Execute()

Patch a managed PostgreSQL database



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstancePatch(context.Background(), apiVersion, instanceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesPostgreSqlInstancePatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesPostgreSqlInstancePatch`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesPostgreSqlInstancePatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesPostgreSqlInstancePatchRequest struct via the builder pattern


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


## PostDatabasesPostgreSqlInstances

> GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner PostDatabasesPostgreSqlInstances(ctx, apiVersion).PostDatabasesPostgreSqlInstancesRequest(postDatabasesPostgreSqlInstancesRequest).Execute()

Create a managed PostgreSQL database



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
	postDatabasesPostgreSqlInstancesRequest := *openapiclient.NewPostDatabasesPostgreSqlInstancesRequest("postgresql/13.2", "example-db", "us-east", "g6-dedicated-2") // PostDatabasesPostgreSqlInstancesRequest | Information about the Managed PostgreSQL Database you are creating.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstances(context.Background(), apiVersion).PostDatabasesPostgreSqlInstancesRequest(postDatabasesPostgreSqlInstancesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PostDatabasesPostgreSqlInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostDatabasesPostgreSqlInstances`: GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PostDatabasesPostgreSqlInstances`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostDatabasesPostgreSqlInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postDatabasesPostgreSqlInstancesRequest** | [**PostDatabasesPostgreSqlInstancesRequest**](PostDatabasesPostgreSqlInstancesRequest.md) | Information about the Managed PostgreSQL Database you are creating. | 

### Return type

[**GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner**](GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDatabasesMysqlInstance

> GetDatabasesMysqlInstances200ResponseAllOfDataInner PutDatabasesMysqlInstance(ctx, apiVersion, instanceId).PutDatabasesMysqlInstanceRequest(putDatabasesMysqlInstanceRequest).Execute()

Update a managed MySQL database



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	putDatabasesMysqlInstanceRequest := *openapiclient.NewPutDatabasesMysqlInstanceRequest() // PutDatabasesMysqlInstanceRequest | Updated information for the Managed MySQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PutDatabasesMysqlInstance(context.Background(), apiVersion, instanceId).PutDatabasesMysqlInstanceRequest(putDatabasesMysqlInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PutDatabasesMysqlInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDatabasesMysqlInstance`: GetDatabasesMysqlInstances200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PutDatabasesMysqlInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDatabasesMysqlInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putDatabasesMysqlInstanceRequest** | [**PutDatabasesMysqlInstanceRequest**](PutDatabasesMysqlInstanceRequest.md) | Updated information for the Managed MySQL Database. | 

### Return type

[**GetDatabasesMysqlInstances200ResponseAllOfDataInner**](GetDatabasesMysqlInstances200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDatabasesPostgreSqlInstance

> GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner PutDatabasesPostgreSqlInstance(ctx, apiVersion, instanceId).PutDatabasesPostgreSqlInstanceRequest(putDatabasesPostgreSqlInstanceRequest).Execute()

Update a managed PostgreSQL database



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
	instanceId := int32(56) // int32 | The ID of the Managed PostgreSQL Database.
	putDatabasesPostgreSqlInstanceRequest := *openapiclient.NewPutDatabasesPostgreSqlInstanceRequest() // PutDatabasesPostgreSqlInstanceRequest | Updated information for the Managed PostgreSQL Database.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DatabasesAPI.PutDatabasesPostgreSqlInstance(context.Background(), apiVersion, instanceId).PutDatabasesPostgreSqlInstanceRequest(putDatabasesPostgreSqlInstanceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DatabasesAPI.PutDatabasesPostgreSqlInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutDatabasesPostgreSqlInstance`: GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner
	fmt.Fprintf(os.Stdout, "Response from `DatabasesAPI.PutDatabasesPostgreSqlInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**instanceId** | **int32** | The ID of the Managed PostgreSQL Database. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDatabasesPostgreSqlInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putDatabasesPostgreSqlInstanceRequest** | [**PutDatabasesPostgreSqlInstanceRequest**](PutDatabasesPostgreSqlInstanceRequest.md) | Updated information for the Managed PostgreSQL Database. | 

### Return type

[**GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner**](GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

