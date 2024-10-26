# GetObjectStorageKeys200ResponseDataInnerBucketAccessInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | The name of the bucket the key can access in the &#x60;region&#x60;. | [optional] 
**Cluster** | Pointer to **string** | Identifies the legacy cluster where this key can be used. The key grants access to each specified &#x60;bucket_name&#x60;, based on the &#x60;permissions&#x60; set. To support backward compatibility, the API generates this value, based on the &#x60;region&#x60; set for a new key. See [Create an Object Storage key](https://techdocs.akamai.com/linode-api/reference/post-object-storage-keys) for more information. | [optional] 
**Permissions** | Pointer to **string** | The level of access the key grants to the &#x60;bucket_name&#x60;. Keys with &#x60;read_write&#x60; access can manage content in the &#x60;bucket_name&#x60;, while &#x60;read_only&#x60; can be used to view content. See [Create an Object Storage key(ref:post-object-storage-keys) for more details. | [optional] 
**Region** | Pointer to **string** | The region where the Object Store &#x60;bucket_name&#x60; resides. | [optional] 

## Methods

### NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInner

`func NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInner() *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner`

NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInner instantiates a new GetObjectStorageKeys200ResponseDataInnerBucketAccessInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInnerWithDefaults

`func NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInnerWithDefaults() *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner`

NewGetObjectStorageKeys200ResponseDataInnerBucketAccessInnerWithDefaults instantiates a new GetObjectStorageKeys200ResponseDataInnerBucketAccessInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetCluster

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetPermissions

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRegion

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetObjectStorageKeys200ResponseDataInnerBucketAccessInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


