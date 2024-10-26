# \ObjectStorageAPI

All URIs are relative to *https://api.linode.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObjectStorageBucket**](ObjectStorageAPI.md#DeleteObjectStorageBucket) | **Delete** /{apiVersion}/object-storage/buckets/{regionId}/{bucket} | Remove an Object Storage bucket
[**DeleteObjectStorageKey**](ObjectStorageAPI.md#DeleteObjectStorageKey) | **Delete** /{apiVersion}/object-storage/keys/{keyId} | Revoke an Object Storage key
[**DeleteObjectStorageSsl**](ObjectStorageAPI.md#DeleteObjectStorageSsl) | **Delete** /{apiVersion}/object-storage/buckets/{regionId}/{bucket}/ssl | Delete an Object Storage TLS/SSL certificate
[**GetObjectStorageBucket**](ObjectStorageAPI.md#GetObjectStorageBucket) | **Get** /{apiVersion}/object-storage/buckets/{regionId}/{bucket} | Get an Object Storage bucket
[**GetObjectStorageBucketAcl**](ObjectStorageAPI.md#GetObjectStorageBucketAcl) | **Get** /{apiVersion}/object-storage/buckets/{regionId}/{bucket}/object-acl | Get an Object Storage object ACL config
[**GetObjectStorageBucketContent**](ObjectStorageAPI.md#GetObjectStorageBucketContent) | **Get** /{apiVersion}/object-storage/buckets/{regionId}/{bucket}/object-list | List Object Storage bucket contents
[**GetObjectStorageBucketinCluster**](ObjectStorageAPI.md#GetObjectStorageBucketinCluster) | **Get** /{apiVersion}/object-storage/buckets/{regionId} | List Object Storage buckets per region
[**GetObjectStorageBuckets**](ObjectStorageAPI.md#GetObjectStorageBuckets) | **Get** /{apiVersion}/object-storage/buckets | List Object Storage buckets
[**GetObjectStorageCluster**](ObjectStorageAPI.md#GetObjectStorageCluster) | **Get** /{apiVersion}/object-storage/clusters/{clusterId} | Get a cluster
[**GetObjectStorageClusters**](ObjectStorageAPI.md#GetObjectStorageClusters) | **Get** /{apiVersion}/object-storage/clusters | List clusters
[**GetObjectStorageKey**](ObjectStorageAPI.md#GetObjectStorageKey) | **Get** /{apiVersion}/object-storage/keys/{keyId} | Get an Object Storage key
[**GetObjectStorageKeys**](ObjectStorageAPI.md#GetObjectStorageKeys) | **Get** /{apiVersion}/object-storage/keys | List Object Storage keys
[**GetObjectStorageSsl**](ObjectStorageAPI.md#GetObjectStorageSsl) | **Get** /{apiVersion}/object-storage/buckets/{regionId}/{bucket}/ssl | Get an Object Storage TLS/SSL certificate
[**GetObjectStorageTransfer**](ObjectStorageAPI.md#GetObjectStorageTransfer) | **Get** /{apiVersion}/object-storage/transfer | Get Object Storage transfer data
[**GetObjectStorageTypes**](ObjectStorageAPI.md#GetObjectStorageTypes) | **Get** /{apiVersion}/object-storage/types | List Object Storage types
[**PostCancelObjectStorage**](ObjectStorageAPI.md#PostCancelObjectStorage) | **Post** /{apiVersion}/object-storage/cancel | Cancel Object Storage
[**PostObjectStorageBucket**](ObjectStorageAPI.md#PostObjectStorageBucket) | **Post** /{apiVersion}/object-storage/buckets | Create an Object Storage bucket
[**PostObjectStorageBucketAccess**](ObjectStorageAPI.md#PostObjectStorageBucketAccess) | **Post** /{apiVersion}/object-storage/buckets/{regionId}/{bucket}/access | Modify access to an Object Storage bucket
[**PostObjectStorageKeys**](ObjectStorageAPI.md#PostObjectStorageKeys) | **Post** /{apiVersion}/object-storage/keys | Create an Object Storage key
[**PostObjectStorageObjectUrl**](ObjectStorageAPI.md#PostObjectStorageObjectUrl) | **Post** /{apiVersion}/object-storage/buckets/{regionId}/{bucket}/object-url | Create a URL for an object
[**PostObjectStorageSsl**](ObjectStorageAPI.md#PostObjectStorageSsl) | **Post** /{apiVersion}/object-storage/buckets/{regionId}/{bucket}/ssl | Upload an Object Storage TLS/SSL certificate
[**PutObjectStorageBucketAcl**](ObjectStorageAPI.md#PutObjectStorageBucketAcl) | **Put** /{apiVersion}/object-storage/buckets/{regionId}/{bucket}/object-acl | Update an object&#39;s ACL config
[**PutObjectStorageKey**](ObjectStorageAPI.md#PutObjectStorageKey) | **Put** /{apiVersion}/object-storage/keys/{keyId} | Update an Object Storage key
[**PutStorageBucketAccess**](ObjectStorageAPI.md#PutStorageBucketAccess) | **Put** /{apiVersion}/object-storage/buckets/{regionId}/{bucket}/access | Update access to an Object Storage bucket



## DeleteObjectStorageBucket

> map[string]interface{} DeleteObjectStorageBucket(ctx, apiVersion, regionId, bucket).Execute()

Remove an Object Storage bucket



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.DeleteObjectStorageBucket(context.Background(), apiVersion, regionId, bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.DeleteObjectStorageBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjectStorageBucket`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.DeleteObjectStorageBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageBucketRequest struct via the builder pattern


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


## DeleteObjectStorageKey

> map[string]interface{} DeleteObjectStorageKey(ctx, apiVersion, keyId).Execute()

Revoke an Object Storage key



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
	keyId := int32(56) // int32 | The key to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.DeleteObjectStorageKey(context.Background(), apiVersion, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.DeleteObjectStorageKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjectStorageKey`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.DeleteObjectStorageKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**keyId** | **int32** | The key to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageKeyRequest struct via the builder pattern


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


## DeleteObjectStorageSsl

> map[string]interface{} DeleteObjectStorageSsl(ctx, apiVersion, regionId, bucket).Execute()

Delete an Object Storage TLS/SSL certificate



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.DeleteObjectStorageSsl(context.Background(), apiVersion, regionId, bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.DeleteObjectStorageSsl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteObjectStorageSsl`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.DeleteObjectStorageSsl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectStorageSslRequest struct via the builder pattern


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


## GetObjectStorageBucket

> GetObjectStorageBuckets200ResponseDataInner GetObjectStorageBucket(ctx, apiVersion, regionId, bucket).Execute()

Get an Object Storage bucket



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageBucket(context.Background(), apiVersion, regionId, bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageBucket`: GetObjectStorageBuckets200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetObjectStorageBuckets200ResponseDataInner**](GetObjectStorageBuckets200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageBucketAcl

> GetObjectStorageBucketAcl200Response GetObjectStorageBucketAcl(ctx, apiVersion, regionId, bucket).Name(name).Execute()

Get an Object Storage object ACL config



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.
	name := "name_example" // string | The `name` of the object for which to retrieve its Access Control List (ACL). Run the [List Object Storage bucket contents](https://techdocs.akamai.com/linode-api/reference/get-object-storage-bucket-content) operation to access all object names in a bucket.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageBucketAcl(context.Background(), apiVersion, regionId, bucket).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageBucketAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageBucketAcl`: GetObjectStorageBucketAcl200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageBucketAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **name** | **string** | The &#x60;name&#x60; of the object for which to retrieve its Access Control List (ACL). Run the [List Object Storage bucket contents](https://techdocs.akamai.com/linode-api/reference/get-object-storage-bucket-content) operation to access all object names in a bucket. | 

### Return type

[**GetObjectStorageBucketAcl200Response**](GetObjectStorageBucketAcl200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageBucketContent

> GetObjectStorageBucketContent200Response GetObjectStorageBucketContent(ctx, apiVersion, regionId, bucket).Marker(marker).Delimiter(delimiter).Prefix(prefix).PageSize(pageSize).Execute()

List Object Storage bucket contents



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.
	marker := "marker_example" // string | The \"marker\" for this request, which can be used to paginate through large buckets. Its value should be the value of the `next_marker` property returned with the last page. Listing bucket contents _does not_ support arbitrary page access. See the `next_marker` property in the responses section for more details. (optional)
	delimiter := "delimiter_example" // string | The delimiter for object names; if given, object names will be returned up to the first occurrence of this character. This is most commonly used with the `/` character to allow bucket transversal in a manner similar to a filesystem, however any delimiter may be used. Use in conjunction with `prefix` to see object names past the first occurrence of the delimiter. (optional)
	prefix := "prefix_example" // string | Filters objects returned to only those whose name start with the given prefix. Commonly used in conjunction with `delimiter` to allow transversal of bucket contents in a manner similar to a filesystem. (optional)
	pageSize := int32(56) // int32 | The number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageBucketContent(context.Background(), apiVersion, regionId, bucket).Marker(marker).Delimiter(delimiter).Prefix(prefix).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageBucketContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageBucketContent`: GetObjectStorageBucketContent200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageBucketContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **marker** | **string** | The \&quot;marker\&quot; for this request, which can be used to paginate through large buckets. Its value should be the value of the &#x60;next_marker&#x60; property returned with the last page. Listing bucket contents _does not_ support arbitrary page access. See the &#x60;next_marker&#x60; property in the responses section for more details. | 
 **delimiter** | **string** | The delimiter for object names; if given, object names will be returned up to the first occurrence of this character. This is most commonly used with the &#x60;/&#x60; character to allow bucket transversal in a manner similar to a filesystem, however any delimiter may be used. Use in conjunction with &#x60;prefix&#x60; to see object names past the first occurrence of the delimiter. | 
 **prefix** | **string** | Filters objects returned to only those whose name start with the given prefix. Commonly used in conjunction with &#x60;delimiter&#x60; to allow transversal of bucket contents in a manner similar to a filesystem. | 
 **pageSize** | **int32** | The number of items to return per page. | [default to 100]

### Return type

[**GetObjectStorageBucketContent200Response**](GetObjectStorageBucketContent200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageBucketinCluster

> GetObjectStorageBuckets200Response GetObjectStorageBucketinCluster(ctx, apiVersion, regionId).Execute()

List Object Storage buckets per region



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageBucketinCluster(context.Background(), apiVersion, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageBucketinCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageBucketinCluster`: GetObjectStorageBuckets200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageBucketinCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketinClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetObjectStorageBuckets200Response**](GetObjectStorageBuckets200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageBuckets

> GetObjectStorageBuckets200Response GetObjectStorageBuckets(ctx, apiVersion).Execute()

List Object Storage buckets



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
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageBuckets(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageBuckets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageBuckets`: GetObjectStorageBuckets200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageBuckets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageBucketsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetObjectStorageBuckets200Response**](GetObjectStorageBuckets200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageCluster

> GetObjectStorageClusters200ResponseDataInner GetObjectStorageCluster(ctx, apiVersion, clusterId).Execute()

Get a cluster



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
	clusterId := "us-east-1" // string | Identifies a cluster where this bucket lives. For backward compatibility with Object Storage in this API.  > 📘 > > You can use the applicable `regionId`, for example `us-west`, in place of the `clusterId`, for example, `us-west-1`. Run [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) to see all regions.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageCluster(context.Background(), apiVersion, clusterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageCluster``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageCluster`: GetObjectStorageClusters200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageCluster`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**clusterId** | **string** | Identifies a cluster where this bucket lives. For backward compatibility with Object Storage in this API.  &gt; 📘 &gt; &gt; You can use the applicable &#x60;regionId&#x60;, for example &#x60;us-west&#x60;, in place of the &#x60;clusterId&#x60;, for example, &#x60;us-west-1&#x60;. Run [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) to see all regions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageClusterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetObjectStorageClusters200ResponseDataInner**](GetObjectStorageClusters200ResponseDataInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageClusters

> GetObjectStorageClusters200Response GetObjectStorageClusters(ctx, apiVersion).Execute()

List clusters



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
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageClusters(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageClusters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageClusters`: GetObjectStorageClusters200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageClusters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetObjectStorageClusters200Response**](GetObjectStorageClusters200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageKey

> GetObjectStorageKeys200ResponseDataInner GetObjectStorageKey(ctx, apiVersion, keyId).Execute()

Get an Object Storage key



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
	keyId := int32(56) // int32 | The key to look up.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageKey(context.Background(), apiVersion, keyId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageKey`: GetObjectStorageKeys200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**keyId** | **int32** | The key to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetObjectStorageKeys200ResponseDataInner**](GetObjectStorageKeys200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageKeys

> GetObjectStorageKeys200Response GetObjectStorageKeys(ctx, apiVersion).Execute()

List Object Storage keys



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
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageKeys(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageKeys`: GetObjectStorageKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetObjectStorageKeys200Response**](GetObjectStorageKeys200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageSsl

> GetObjectStorageSsl200Response GetObjectStorageSsl(ctx, apiVersion, regionId, bucket).Execute()

Get an Object Storage TLS/SSL certificate



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageSsl(context.Background(), apiVersion, regionId, bucket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageSsl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageSsl`: GetObjectStorageSsl200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageSsl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageSslRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetObjectStorageSsl200Response**](GetObjectStorageSsl200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageTransfer

> GetObjectStorageTransfer200Response GetObjectStorageTransfer(ctx, apiVersion).Execute()

Get Object Storage transfer data



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
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageTransfer(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageTransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageTransfer`: GetObjectStorageTransfer200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetObjectStorageTransfer200Response**](GetObjectStorageTransfer200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectStorageTypes

> GetObjectStorageTypes200Response GetObjectStorageTypes(ctx, apiVersion).Execute()

List Object Storage types



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
	resp, r, err := apiClient.ObjectStorageAPI.GetObjectStorageTypes(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.GetObjectStorageTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectStorageTypes`: GetObjectStorageTypes200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.GetObjectStorageTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectStorageTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetObjectStorageTypes200Response**](GetObjectStorageTypes200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCancelObjectStorage

> map[string]interface{} PostCancelObjectStorage(ctx, apiVersion).Execute()

Cancel Object Storage



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
	resp, r, err := apiClient.ObjectStorageAPI.PostCancelObjectStorage(context.Background(), apiVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.PostCancelObjectStorage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCancelObjectStorage`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.PostCancelObjectStorage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostCancelObjectStorageRequest struct via the builder pattern


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


## PostObjectStorageBucket

> GetObjectStorageBuckets200ResponseDataInner PostObjectStorageBucket(ctx, apiVersion).PostObjectStorageBucketRequest(postObjectStorageBucketRequest).Execute()

Create an Object Storage bucket



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
	postObjectStorageBucketRequest := *openapiclient.NewPostObjectStorageBucketRequest("example-bucket") // PostObjectStorageBucketRequest | Information about the bucket you want to create. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.PostObjectStorageBucket(context.Background(), apiVersion).PostObjectStorageBucketRequest(postObjectStorageBucketRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.PostObjectStorageBucket``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostObjectStorageBucket`: GetObjectStorageBuckets200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.PostObjectStorageBucket`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostObjectStorageBucketRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postObjectStorageBucketRequest** | [**PostObjectStorageBucketRequest**](PostObjectStorageBucketRequest.md) | Information about the bucket you want to create. | 

### Return type

[**GetObjectStorageBuckets200ResponseDataInner**](GetObjectStorageBuckets200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostObjectStorageBucketAccess

> map[string]interface{} PostObjectStorageBucketAccess(ctx, apiVersion, regionId, bucket).PutStorageBucketAccessRequest(putStorageBucketAccessRequest).Execute()

Modify access to an Object Storage bucket



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.
	putStorageBucketAccessRequest := *openapiclient.NewPutStorageBucketAccessRequest() // PutStorageBucketAccessRequest | The changes to make to the bucket's access controls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.PostObjectStorageBucketAccess(context.Background(), apiVersion, regionId, bucket).PutStorageBucketAccessRequest(putStorageBucketAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.PostObjectStorageBucketAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostObjectStorageBucketAccess`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.PostObjectStorageBucketAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostObjectStorageBucketAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putStorageBucketAccessRequest** | [**PutStorageBucketAccessRequest**](PutStorageBucketAccessRequest.md) | The changes to make to the bucket&#39;s access controls. | 

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


## PostObjectStorageKeys

> GetObjectStorageKeys200ResponseDataInner PostObjectStorageKeys(ctx, apiVersion).PostObjectStorageKeysRequest(postObjectStorageKeysRequest).Execute()

Create an Object Storage key



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
	postObjectStorageKeysRequest := *openapiclient.NewPostObjectStorageKeysRequest() // PostObjectStorageKeysRequest | The settings necessary to create a new key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.PostObjectStorageKeys(context.Background(), apiVersion).PostObjectStorageKeysRequest(postObjectStorageKeysRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.PostObjectStorageKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostObjectStorageKeys`: GetObjectStorageKeys200ResponseDataInner
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.PostObjectStorageKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPostObjectStorageKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postObjectStorageKeysRequest** | [**PostObjectStorageKeysRequest**](PostObjectStorageKeysRequest.md) | The settings necessary to create a new key. | 

### Return type

[**GetObjectStorageKeys200ResponseDataInner**](GetObjectStorageKeys200ResponseDataInner.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostObjectStorageObjectUrl

> PostObjectStorageObjectUrl200Response PostObjectStorageObjectUrl(ctx, apiVersion, regionId, bucket).PostObjectStorageObjectUrlRequest(postObjectStorageObjectUrlRequest).Execute()

Create a URL for an object



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.
	postObjectStorageObjectUrlRequest := *openapiclient.NewPostObjectStorageObjectUrlRequest("GET", "example") // PostObjectStorageObjectUrlRequest | Information about the request to sign. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.PostObjectStorageObjectUrl(context.Background(), apiVersion, regionId, bucket).PostObjectStorageObjectUrlRequest(postObjectStorageObjectUrlRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.PostObjectStorageObjectUrl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostObjectStorageObjectUrl`: PostObjectStorageObjectUrl200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.PostObjectStorageObjectUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostObjectStorageObjectUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postObjectStorageObjectUrlRequest** | [**PostObjectStorageObjectUrlRequest**](PostObjectStorageObjectUrlRequest.md) | Information about the request to sign. | 

### Return type

[**PostObjectStorageObjectUrl200Response**](PostObjectStorageObjectUrl200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostObjectStorageSsl

> GetObjectStorageSsl200Response PostObjectStorageSsl(ctx, apiVersion, regionId, bucket).PostObjectStorageSslRequest(postObjectStorageSslRequest).Execute()

Upload an Object Storage TLS/SSL certificate



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.
	postObjectStorageSslRequest := *openapiclient.NewPostObjectStorageSslRequest("-----BEGIN CERTIFICATE-----
CERTIFICATE_INFORMATION
-----END CERTIFICATE-----", "-----BEGIN PRIVATE KEY-----
PRIVATE_KEY_INFORMATION
-----END PRIVATE KEY-----") // PostObjectStorageSslRequest | Upload this TLS/SSL certificate with its corresponding secret key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.PostObjectStorageSsl(context.Background(), apiVersion, regionId, bucket).PostObjectStorageSslRequest(postObjectStorageSslRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.PostObjectStorageSsl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostObjectStorageSsl`: GetObjectStorageSsl200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.PostObjectStorageSsl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostObjectStorageSslRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **postObjectStorageSslRequest** | [**PostObjectStorageSslRequest**](PostObjectStorageSslRequest.md) | Upload this TLS/SSL certificate with its corresponding secret key. | 

### Return type

[**GetObjectStorageSsl200Response**](GetObjectStorageSsl200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutObjectStorageBucketAcl

> GetObjectStorageBucketAcl200Response PutObjectStorageBucketAcl(ctx, apiVersion, regionId, bucket).PutObjectStorageBucketAclRequest(putObjectStorageBucketAclRequest).Execute()

Update an object's ACL config



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.
	putObjectStorageBucketAclRequest := *openapiclient.NewPutObjectStorageBucketAclRequest("public-read", "Name_example") // PutObjectStorageBucketAclRequest | The changes to make to this Object's access controls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.PutObjectStorageBucketAcl(context.Background(), apiVersion, regionId, bucket).PutObjectStorageBucketAclRequest(putObjectStorageBucketAclRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.PutObjectStorageBucketAcl``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutObjectStorageBucketAcl`: GetObjectStorageBucketAcl200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.PutObjectStorageBucketAcl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutObjectStorageBucketAclRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putObjectStorageBucketAclRequest** | [**PutObjectStorageBucketAclRequest**](PutObjectStorageBucketAclRequest.md) | The changes to make to this Object&#39;s access controls. | 

### Return type

[**GetObjectStorageBucketAcl200Response**](GetObjectStorageBucketAcl200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutObjectStorageKey

> PutObjectStorageKey200Response PutObjectStorageKey(ctx, apiVersion, keyId).PutObjectStorageKeyRequest(putObjectStorageKeyRequest).Execute()

Update an Object Storage key



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
	keyId := int32(56) // int32 | The key to look up.
	putObjectStorageKeyRequest := *openapiclient.NewPutObjectStorageKeyRequest() // PutObjectStorageKeyRequest | The fields to update. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.PutObjectStorageKey(context.Background(), apiVersion, keyId).PutObjectStorageKeyRequest(putObjectStorageKeyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.PutObjectStorageKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutObjectStorageKey`: PutObjectStorageKey200Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.PutObjectStorageKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**keyId** | **int32** | The key to look up. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutObjectStorageKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **putObjectStorageKeyRequest** | [**PutObjectStorageKeyRequest**](PutObjectStorageKeyRequest.md) | The fields to update. | 

### Return type

[**PutObjectStorageKey200Response**](PutObjectStorageKey200Response.md)

### Authorization

[personalAccessToken](../README.md#personalAccessToken), [oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutStorageBucketAccess

> map[string]interface{} PutStorageBucketAccess(ctx, apiVersion, regionId, bucket).PutStorageBucketAccessRequest(putStorageBucketAccessRequest).Execute()

Update access to an Object Storage bucket



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
	regionId := "regionId_example" // string | Identifies a region where this bucket lives.  > 📘 > > You can use a `clusterId` in place of `regionId` in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster `id`.
	bucket := "bucket_example" // string | The bucket name.
	putStorageBucketAccessRequest := *openapiclient.NewPutStorageBucketAccessRequest() // PutStorageBucketAccessRequest | The changes to make to the bucket's access controls. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectStorageAPI.PutStorageBucketAccess(context.Background(), apiVersion, regionId, bucket).PutStorageBucketAccessRequest(putStorageBucketAccessRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectStorageAPI.PutStorageBucketAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutStorageBucketAccess`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `ObjectStorageAPI.PutStorageBucketAccess`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiVersion** | **string** | __Enum__ Call either the &#x60;v4&#x60; URL, or &#x60;v4beta&#x60; for operations still in Beta. | [default to &quot;v4&quot;]
**regionId** | **string** | Identifies a region where this bucket lives.  &gt; 📘 &gt; &gt; You can use a &#x60;clusterId&#x60; in place of &#x60;regionId&#x60; in requests for buckets that you created using the legacy version of the API. Run [List clusters](https://techdocs.akamai.com/linode-api/reference/get-object-storage-clusters) to see each cluster &#x60;id&#x60;. | 
**bucket** | **string** | The bucket name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutStorageBucketAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **putStorageBucketAccessRequest** | [**PutStorageBucketAccessRequest**](PutStorageBucketAccessRequest.md) | The changes to make to the bucket&#39;s access controls. | 

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

