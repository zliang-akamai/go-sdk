# GetFirewallDevices200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this Device was created. | [optional] [readonly] 
**Entity** | Pointer to [**GetFirewallDevices200ResponseDataInnerEntity**](GetFirewallDevices200ResponseDataInnerEntity.md) |  | [optional] 
**Id** | Pointer to **int32** | The Device&#39;s unique ID. | [optional] 
**Updated** | Pointer to **time.Time** | When this Device was last updated. | [optional] [readonly] 

## Methods

### NewGetFirewallDevices200ResponseDataInner

`func NewGetFirewallDevices200ResponseDataInner() *GetFirewallDevices200ResponseDataInner`

NewGetFirewallDevices200ResponseDataInner instantiates a new GetFirewallDevices200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFirewallDevices200ResponseDataInnerWithDefaults

`func NewGetFirewallDevices200ResponseDataInnerWithDefaults() *GetFirewallDevices200ResponseDataInner`

NewGetFirewallDevices200ResponseDataInnerWithDefaults instantiates a new GetFirewallDevices200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetFirewallDevices200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetFirewallDevices200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetFirewallDevices200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetFirewallDevices200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEntity

`func (o *GetFirewallDevices200ResponseDataInner) GetEntity() GetFirewallDevices200ResponseDataInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GetFirewallDevices200ResponseDataInner) GetEntityOk() (*GetFirewallDevices200ResponseDataInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GetFirewallDevices200ResponseDataInner) SetEntity(v GetFirewallDevices200ResponseDataInnerEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GetFirewallDevices200ResponseDataInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetId

`func (o *GetFirewallDevices200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetFirewallDevices200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetFirewallDevices200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetFirewallDevices200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUpdated

`func (o *GetFirewallDevices200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetFirewallDevices200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetFirewallDevices200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetFirewallDevices200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


