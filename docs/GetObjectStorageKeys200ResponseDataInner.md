# GetObjectStorageKeys200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | A unique string chosen by the API to identify this key. Used as a user name to identify this key when making requests to the S3 API. | [optional] [readonly] 
**BucketAccess** | Pointer to [**[]GetObjectStorageKeys200ResponseDataInnerBucketAccessInner**](GetObjectStorageKeys200ResponseDataInnerBucketAccessInner.md) | Settings that limit access to specific buckets, each with a specific permission level. | [optional] 
**Id** | Pointer to **int32** | This Object Storage key&#39;s unique ID. | [optional] [readonly] 
**Label** | Pointer to **string** | The label given to this key. For display purposes only. | [optional] 
**Limited** | Pointer to **bool** | Whether this Object Storage key limits access to specific buckets and permissions. Returns &#x60;false&#x60; if this key grants full access. Specific limitations are set in &#x60;bucket_access&#x60;. | [optional] [readonly] 
**Regions** | Pointer to [**[]GetObjectStorageKeys200ResponseDataInnerRegionsInner**](GetObjectStorageKeys200ResponseDataInnerRegionsInner.md) | The key can be used in these regions to create new buckets but it can&#39;t be used to manage content in those buckets. See [Create an Object Storage key](https://techdocs.akamai.com/linode-api/reference/post-object-storage-keys) for more details. | [optional] 
**SecretKey** | Pointer to **string** | This Object Storage key&#39;s secret key. Used as a password to validate this key when making requests to the S3 API. This value is only revealed in a response after creating or modifying a key. | [optional] [readonly] 

## Methods

### NewGetObjectStorageKeys200ResponseDataInner

`func NewGetObjectStorageKeys200ResponseDataInner() *GetObjectStorageKeys200ResponseDataInner`

NewGetObjectStorageKeys200ResponseDataInner instantiates a new GetObjectStorageKeys200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageKeys200ResponseDataInnerWithDefaults

`func NewGetObjectStorageKeys200ResponseDataInnerWithDefaults() *GetObjectStorageKeys200ResponseDataInner`

NewGetObjectStorageKeys200ResponseDataInnerWithDefaults instantiates a new GetObjectStorageKeys200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *GetObjectStorageKeys200ResponseDataInner) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *GetObjectStorageKeys200ResponseDataInner) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *GetObjectStorageKeys200ResponseDataInner) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *GetObjectStorageKeys200ResponseDataInner) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetBucketAccess

`func (o *GetObjectStorageKeys200ResponseDataInner) GetBucketAccess() []GetObjectStorageKeys200ResponseDataInnerBucketAccessInner`

GetBucketAccess returns the BucketAccess field if non-nil, zero value otherwise.

### GetBucketAccessOk

`func (o *GetObjectStorageKeys200ResponseDataInner) GetBucketAccessOk() (*[]GetObjectStorageKeys200ResponseDataInnerBucketAccessInner, bool)`

GetBucketAccessOk returns a tuple with the BucketAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketAccess

`func (o *GetObjectStorageKeys200ResponseDataInner) SetBucketAccess(v []GetObjectStorageKeys200ResponseDataInnerBucketAccessInner)`

SetBucketAccess sets BucketAccess field to given value.

### HasBucketAccess

`func (o *GetObjectStorageKeys200ResponseDataInner) HasBucketAccess() bool`

HasBucketAccess returns a boolean if a field has been set.

### GetId

`func (o *GetObjectStorageKeys200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetObjectStorageKeys200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetObjectStorageKeys200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetObjectStorageKeys200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetObjectStorageKeys200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetObjectStorageKeys200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetObjectStorageKeys200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetObjectStorageKeys200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLimited

`func (o *GetObjectStorageKeys200ResponseDataInner) GetLimited() bool`

GetLimited returns the Limited field if non-nil, zero value otherwise.

### GetLimitedOk

`func (o *GetObjectStorageKeys200ResponseDataInner) GetLimitedOk() (*bool, bool)`

GetLimitedOk returns a tuple with the Limited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimited

`func (o *GetObjectStorageKeys200ResponseDataInner) SetLimited(v bool)`

SetLimited sets Limited field to given value.

### HasLimited

`func (o *GetObjectStorageKeys200ResponseDataInner) HasLimited() bool`

HasLimited returns a boolean if a field has been set.

### GetRegions

`func (o *GetObjectStorageKeys200ResponseDataInner) GetRegions() []GetObjectStorageKeys200ResponseDataInnerRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GetObjectStorageKeys200ResponseDataInner) GetRegionsOk() (*[]GetObjectStorageKeys200ResponseDataInnerRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GetObjectStorageKeys200ResponseDataInner) SetRegions(v []GetObjectStorageKeys200ResponseDataInnerRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *GetObjectStorageKeys200ResponseDataInner) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSecretKey

`func (o *GetObjectStorageKeys200ResponseDataInner) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *GetObjectStorageKeys200ResponseDataInner) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *GetObjectStorageKeys200ResponseDataInner) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *GetObjectStorageKeys200ResponseDataInner) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


