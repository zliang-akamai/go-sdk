# GetObjectStorageBucketContent200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Etag** | Pointer to **string** | An MD-5 hash of the object. &#x60;null&#x60; if this object represents a prefix. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time this object was last modified. &#x60;null&#x60; if this object represents a prefix. | [optional] 
**Name** | Pointer to **string** | The name of this object or prefix. | [optional] 
**Owner** | Pointer to **string** | The owner of this object, as a UUID. &#x60;null&#x60; if this object represents a prefix. | [optional] 
**Size** | Pointer to **int32** | The size of this object, in bytes. &#x60;null&#x60; if this object represents a prefix. | [optional] 

## Methods

### NewGetObjectStorageBucketContent200ResponseDataInner

`func NewGetObjectStorageBucketContent200ResponseDataInner() *GetObjectStorageBucketContent200ResponseDataInner`

NewGetObjectStorageBucketContent200ResponseDataInner instantiates a new GetObjectStorageBucketContent200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageBucketContent200ResponseDataInnerWithDefaults

`func NewGetObjectStorageBucketContent200ResponseDataInnerWithDefaults() *GetObjectStorageBucketContent200ResponseDataInner`

NewGetObjectStorageBucketContent200ResponseDataInnerWithDefaults instantiates a new GetObjectStorageBucketContent200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtag

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetEtag() string`

GetEtag returns the Etag field if non-nil, zero value otherwise.

### GetEtagOk

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetEtagOk() (*string, bool)`

GetEtagOk returns a tuple with the Etag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtag

`func (o *GetObjectStorageBucketContent200ResponseDataInner) SetEtag(v string)`

SetEtag sets Etag field to given value.

### HasEtag

`func (o *GetObjectStorageBucketContent200ResponseDataInner) HasEtag() bool`

HasEtag returns a boolean if a field has been set.

### GetLastModified

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *GetObjectStorageBucketContent200ResponseDataInner) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *GetObjectStorageBucketContent200ResponseDataInner) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetName

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetObjectStorageBucketContent200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetObjectStorageBucketContent200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetObjectStorageBucketContent200ResponseDataInner) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetObjectStorageBucketContent200ResponseDataInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetSize

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetObjectStorageBucketContent200ResponseDataInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetObjectStorageBucketContent200ResponseDataInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetObjectStorageBucketContent200ResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


