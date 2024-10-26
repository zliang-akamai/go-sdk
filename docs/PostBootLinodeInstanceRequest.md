# PostBootLinodeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **int32** | The Linode Config ID to boot into. | [optional] 

## Methods

### NewPostBootLinodeInstanceRequest

`func NewPostBootLinodeInstanceRequest() *PostBootLinodeInstanceRequest`

NewPostBootLinodeInstanceRequest instantiates a new PostBootLinodeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostBootLinodeInstanceRequestWithDefaults

`func NewPostBootLinodeInstanceRequestWithDefaults() *PostBootLinodeInstanceRequest`

NewPostBootLinodeInstanceRequestWithDefaults instantiates a new PostBootLinodeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *PostBootLinodeInstanceRequest) GetConfigId() int32`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *PostBootLinodeInstanceRequest) GetConfigIdOk() (*int32, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *PostBootLinodeInstanceRequest) SetConfigId(v int32)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *PostBootLinodeInstanceRequest) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


