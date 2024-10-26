# PutObjectStorageKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | A unique string chosen by the API to identify this key. Used as a user name to identify this key when making requests to the S3 API. | [optional] [readonly] 
**Id** | Pointer to **int32** | This Object Storage key&#39;s unique numeric identifier. | [optional] [readonly] 
**Label** | Pointer to **string** | The label given to this key. For display purposes only. | [optional] 
**Limited** | Pointer to **bool** | Whether this Object Storage key limits access to specific buckets and permissions. Returns &#x60;false&#x60; if this key grants full access.  &gt; 📘 &gt; &gt; The &#x60;bucket_access&#x60; array that contains limited Object Storage key settings doesn&#39;t appear in this response. Store this key&#39;s &#x60;id&#x60; from the response and run [Get an Object Storage key](https://techdocs.akamai.com/linode-api/reference/get-object-storage-key) to view these settings. | [optional] [readonly] 
**Regions** | Pointer to [**[]PutObjectStorageKey200ResponseRegionsInner**](PutObjectStorageKey200ResponseRegionsInner.md) | The key can be used in these regions to create new buckets, but it can&#39;t be used to manage content in those buckets. See [Create an Object Storage key](https://techdocs.akamai.com/linode-api/reference/post-object-storage-keys) for more details. | [optional] 
**SecretKey** | Pointer to **string** | This Object Storage key&#39;s secret key. Used as a password to validate this key when making requests to the S3 API. This value is only revealed in a response after creating or modifying a key. | [optional] [readonly] 

## Methods

### NewPutObjectStorageKey200Response

`func NewPutObjectStorageKey200Response() *PutObjectStorageKey200Response`

NewPutObjectStorageKey200Response instantiates a new PutObjectStorageKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectStorageKey200ResponseWithDefaults

`func NewPutObjectStorageKey200ResponseWithDefaults() *PutObjectStorageKey200Response`

NewPutObjectStorageKey200ResponseWithDefaults instantiates a new PutObjectStorageKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *PutObjectStorageKey200Response) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *PutObjectStorageKey200Response) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *PutObjectStorageKey200Response) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *PutObjectStorageKey200Response) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetId

`func (o *PutObjectStorageKey200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutObjectStorageKey200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutObjectStorageKey200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PutObjectStorageKey200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PutObjectStorageKey200Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutObjectStorageKey200Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutObjectStorageKey200Response) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutObjectStorageKey200Response) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLimited

`func (o *PutObjectStorageKey200Response) GetLimited() bool`

GetLimited returns the Limited field if non-nil, zero value otherwise.

### GetLimitedOk

`func (o *PutObjectStorageKey200Response) GetLimitedOk() (*bool, bool)`

GetLimitedOk returns a tuple with the Limited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimited

`func (o *PutObjectStorageKey200Response) SetLimited(v bool)`

SetLimited sets Limited field to given value.

### HasLimited

`func (o *PutObjectStorageKey200Response) HasLimited() bool`

HasLimited returns a boolean if a field has been set.

### GetRegions

`func (o *PutObjectStorageKey200Response) GetRegions() []PutObjectStorageKey200ResponseRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PutObjectStorageKey200Response) GetRegionsOk() (*[]PutObjectStorageKey200ResponseRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PutObjectStorageKey200Response) SetRegions(v []PutObjectStorageKey200ResponseRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PutObjectStorageKey200Response) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSecretKey

`func (o *PutObjectStorageKey200Response) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *PutObjectStorageKey200Response) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *PutObjectStorageKey200Response) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *PutObjectStorageKey200Response) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


