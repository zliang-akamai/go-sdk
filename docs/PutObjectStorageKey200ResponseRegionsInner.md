# PutObjectStorageKey200ResponseRegionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies each region where you can use the Object Storage key. | [optional] 
**S3Endpoint** | Pointer to **string** | The S3-compatible hostname you can use to access your Object Storage buckets in this region. | [optional] 

## Methods

### NewPutObjectStorageKey200ResponseRegionsInner

`func NewPutObjectStorageKey200ResponseRegionsInner() *PutObjectStorageKey200ResponseRegionsInner`

NewPutObjectStorageKey200ResponseRegionsInner instantiates a new PutObjectStorageKey200ResponseRegionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectStorageKey200ResponseRegionsInnerWithDefaults

`func NewPutObjectStorageKey200ResponseRegionsInnerWithDefaults() *PutObjectStorageKey200ResponseRegionsInner`

NewPutObjectStorageKey200ResponseRegionsInnerWithDefaults instantiates a new PutObjectStorageKey200ResponseRegionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PutObjectStorageKey200ResponseRegionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PutObjectStorageKey200ResponseRegionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PutObjectStorageKey200ResponseRegionsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PutObjectStorageKey200ResponseRegionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetS3Endpoint

`func (o *PutObjectStorageKey200ResponseRegionsInner) GetS3Endpoint() string`

GetS3Endpoint returns the S3Endpoint field if non-nil, zero value otherwise.

### GetS3EndpointOk

`func (o *PutObjectStorageKey200ResponseRegionsInner) GetS3EndpointOk() (*string, bool)`

GetS3EndpointOk returns a tuple with the S3Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Endpoint

`func (o *PutObjectStorageKey200ResponseRegionsInner) SetS3Endpoint(v string)`

SetS3Endpoint sets S3Endpoint field to given value.

### HasS3Endpoint

`func (o *PutObjectStorageKey200ResponseRegionsInner) HasS3Endpoint() bool`

HasS3Endpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


