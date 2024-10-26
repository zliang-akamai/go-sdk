# PostObjectStorageKeysRequestBucketAccessInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketName** | Pointer to **string** | The &#x60;label&#x60; set for a bucket that this key grants access to. | [optional] 
**Permissions** | Pointer to **string** | The level of access the key grants to the specified &#x60;bucket_name&#x60;. Keys with &#x60;read_write&#x60; access can manage content in the &#x60;bucket_name&#x60;, while &#x60;read_only&#x60; can be used to view content. See [Create an Object Storage key]((ref:post-object-storage-keys) for more details. | [optional] 
**Region** | Pointer to **string** | The region where the &#x60;bucket_name&#x60; you want the key to access is located.  &gt; 📘 &gt; &gt; If you repeat the same &#x60;region&#x60; across objects, the response includes a single &#x60;s3_endpoint&#x60; for this region. Use it to access all &#x60;bucket_name&#x60; entries. | [optional] 

## Methods

### NewPostObjectStorageKeysRequestBucketAccessInner

`func NewPostObjectStorageKeysRequestBucketAccessInner() *PostObjectStorageKeysRequestBucketAccessInner`

NewPostObjectStorageKeysRequestBucketAccessInner instantiates a new PostObjectStorageKeysRequestBucketAccessInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostObjectStorageKeysRequestBucketAccessInnerWithDefaults

`func NewPostObjectStorageKeysRequestBucketAccessInnerWithDefaults() *PostObjectStorageKeysRequestBucketAccessInner`

NewPostObjectStorageKeysRequestBucketAccessInnerWithDefaults instantiates a new PostObjectStorageKeysRequestBucketAccessInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketName

`func (o *PostObjectStorageKeysRequestBucketAccessInner) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *PostObjectStorageKeysRequestBucketAccessInner) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *PostObjectStorageKeysRequestBucketAccessInner) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *PostObjectStorageKeysRequestBucketAccessInner) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetPermissions

`func (o *PostObjectStorageKeysRequestBucketAccessInner) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PostObjectStorageKeysRequestBucketAccessInner) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PostObjectStorageKeysRequestBucketAccessInner) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PostObjectStorageKeysRequestBucketAccessInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetRegion

`func (o *PostObjectStorageKeysRequestBucketAccessInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostObjectStorageKeysRequestBucketAccessInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostObjectStorageKeysRequestBucketAccessInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostObjectStorageKeysRequestBucketAccessInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


