# PostRebootLinodeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **int32** | The Linode Config ID to reboot into.  If null or omitted, the last booted config will be used.  If there was no last booted config and this Linode only has one config, it will be used.  If a config cannot be determined, an error will be returned. | [optional] 

## Methods

### NewPostRebootLinodeInstanceRequest

`func NewPostRebootLinodeInstanceRequest() *PostRebootLinodeInstanceRequest`

NewPostRebootLinodeInstanceRequest instantiates a new PostRebootLinodeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRebootLinodeInstanceRequestWithDefaults

`func NewPostRebootLinodeInstanceRequestWithDefaults() *PostRebootLinodeInstanceRequest`

NewPostRebootLinodeInstanceRequestWithDefaults instantiates a new PostRebootLinodeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *PostRebootLinodeInstanceRequest) GetConfigId() int32`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *PostRebootLinodeInstanceRequest) GetConfigIdOk() (*int32, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *PostRebootLinodeInstanceRequest) SetConfigId(v int32)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *PostRebootLinodeInstanceRequest) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


