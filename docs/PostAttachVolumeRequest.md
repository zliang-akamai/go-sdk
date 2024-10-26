# PostAttachVolumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigId** | Pointer to **int32** | The ID of the Linode Config to include this Volume in. Must belong to the Linode referenced by &#x60;linode_id&#x60;. If not given, the last booted Config will be chosen. | [optional] 
**LinodeId** | **int32** | The ID of the Linode to attach the volume to. | 
**PersistAcrossBoots** | Pointer to **bool** | Defaults to true, if false is provided, the Volume will not be attached to the Linode Config. In this case more than 8 Volumes may be attached to a Linode if a Linode has 16GB of RAM or more. The number of volumes that can be attached is equal to the number of GB of RAM that the Linode has, up to a maximum of 64. &#x60;config_id&#x60; should not be passed if this is set to false and linode_id must be passed. The Linode must be running. | [optional] 

## Methods

### NewPostAttachVolumeRequest

`func NewPostAttachVolumeRequest(linodeId int32, ) *PostAttachVolumeRequest`

NewPostAttachVolumeRequest instantiates a new PostAttachVolumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAttachVolumeRequestWithDefaults

`func NewPostAttachVolumeRequestWithDefaults() *PostAttachVolumeRequest`

NewPostAttachVolumeRequestWithDefaults instantiates a new PostAttachVolumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigId

`func (o *PostAttachVolumeRequest) GetConfigId() int32`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *PostAttachVolumeRequest) GetConfigIdOk() (*int32, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *PostAttachVolumeRequest) SetConfigId(v int32)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *PostAttachVolumeRequest) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetLinodeId

`func (o *PostAttachVolumeRequest) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PostAttachVolumeRequest) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PostAttachVolumeRequest) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.


### GetPersistAcrossBoots

`func (o *PostAttachVolumeRequest) GetPersistAcrossBoots() bool`

GetPersistAcrossBoots returns the PersistAcrossBoots field if non-nil, zero value otherwise.

### GetPersistAcrossBootsOk

`func (o *PostAttachVolumeRequest) GetPersistAcrossBootsOk() (*bool, bool)`

GetPersistAcrossBootsOk returns a tuple with the PersistAcrossBoots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistAcrossBoots

`func (o *PostAttachVolumeRequest) SetPersistAcrossBoots(v bool)`

SetPersistAcrossBoots sets PersistAcrossBoots field to given value.

### HasPersistAcrossBoots

`func (o *PostAttachVolumeRequest) HasPersistAcrossBoots() bool`

HasPersistAcrossBoots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


