# GetObjectStorageBucketContent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]GetObjectStorageBucketContent200ResponseDataInner**](GetObjectStorageBucketContent200ResponseDataInner.md) |  | [optional] 
**IsTruncated** | Pointer to **bool** | Designates if there is another page of bucket objects. | [optional] 
**NextMarker** | Pointer to **NullableString** | Returns the value you should pass to the &#x60;marker&#x60; query parameter to get the next page of objects. If there is no next page, &#x60;null&#x60; will be returned. | [optional] 

## Methods

### NewGetObjectStorageBucketContent200Response

`func NewGetObjectStorageBucketContent200Response() *GetObjectStorageBucketContent200Response`

NewGetObjectStorageBucketContent200Response instantiates a new GetObjectStorageBucketContent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageBucketContent200ResponseWithDefaults

`func NewGetObjectStorageBucketContent200ResponseWithDefaults() *GetObjectStorageBucketContent200Response`

NewGetObjectStorageBucketContent200ResponseWithDefaults instantiates a new GetObjectStorageBucketContent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetObjectStorageBucketContent200Response) GetData() []GetObjectStorageBucketContent200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetObjectStorageBucketContent200Response) GetDataOk() (*[]GetObjectStorageBucketContent200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetObjectStorageBucketContent200Response) SetData(v []GetObjectStorageBucketContent200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *GetObjectStorageBucketContent200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetIsTruncated

`func (o *GetObjectStorageBucketContent200Response) GetIsTruncated() bool`

GetIsTruncated returns the IsTruncated field if non-nil, zero value otherwise.

### GetIsTruncatedOk

`func (o *GetObjectStorageBucketContent200Response) GetIsTruncatedOk() (*bool, bool)`

GetIsTruncatedOk returns a tuple with the IsTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTruncated

`func (o *GetObjectStorageBucketContent200Response) SetIsTruncated(v bool)`

SetIsTruncated sets IsTruncated field to given value.

### HasIsTruncated

`func (o *GetObjectStorageBucketContent200Response) HasIsTruncated() bool`

HasIsTruncated returns a boolean if a field has been set.

### GetNextMarker

`func (o *GetObjectStorageBucketContent200Response) GetNextMarker() string`

GetNextMarker returns the NextMarker field if non-nil, zero value otherwise.

### GetNextMarkerOk

`func (o *GetObjectStorageBucketContent200Response) GetNextMarkerOk() (*string, bool)`

GetNextMarkerOk returns a tuple with the NextMarker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextMarker

`func (o *GetObjectStorageBucketContent200Response) SetNextMarker(v string)`

SetNextMarker sets NextMarker field to given value.

### HasNextMarker

`func (o *GetObjectStorageBucketContent200Response) HasNextMarker() bool`

HasNextMarker returns a boolean if a field has been set.

### SetNextMarkerNil

`func (o *GetObjectStorageBucketContent200Response) SetNextMarkerNil(b bool)`

 SetNextMarkerNil sets the value for NextMarker to be an explicit nil

### UnsetNextMarker
`func (o *GetObjectStorageBucketContent200Response) UnsetNextMarker()`

UnsetNextMarker ensures that no value is present for NextMarker, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


