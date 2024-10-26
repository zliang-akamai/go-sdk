# GetDatabasesTypes200ResponseAllOfDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | The compute class category. | [optional] 
**Deprecated** | Pointer to **bool** | Whether this Database plan type has been deprecated and is no longer available. | [optional] 
**Disk** | Pointer to **int32** | The amount of disk space set aside for Databases of this plan type. The value is represented in megabytes. | [optional] 
**Engines** | Pointer to [**GetDatabasesTypes200ResponseAllOfDataInnerEngines**](GetDatabasesTypes200ResponseAllOfDataInnerEngines.md) |  | [optional] 
**Id** | Pointer to **string** | The ID representing the Managed Database node plan type. | [optional] [readonly] 
**Label** | Pointer to **string** | A human-readable string that describes each plan type. For display purposes only. | [optional] [readonly] 
**Memory** | Pointer to **int32** | The amount of RAM allocated to Database created of this plan type. The value is represented in megabytes. | [optional] 
**Vcpus** | Pointer to **int32** | The number of CPUs allocated to databases of this plan type. | [optional] 

## Methods

### NewGetDatabasesTypes200ResponseAllOfDataInner

`func NewGetDatabasesTypes200ResponseAllOfDataInner() *GetDatabasesTypes200ResponseAllOfDataInner`

NewGetDatabasesTypes200ResponseAllOfDataInner instantiates a new GetDatabasesTypes200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasesTypes200ResponseAllOfDataInnerWithDefaults

`func NewGetDatabasesTypes200ResponseAllOfDataInnerWithDefaults() *GetDatabasesTypes200ResponseAllOfDataInner`

NewGetDatabasesTypes200ResponseAllOfDataInnerWithDefaults instantiates a new GetDatabasesTypes200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetDeprecated

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDisk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetEngines

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetEngines() GetDatabasesTypes200ResponseAllOfDataInnerEngines`

GetEngines returns the Engines field if non-nil, zero value otherwise.

### GetEnginesOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetEnginesOk() (*GetDatabasesTypes200ResponseAllOfDataInnerEngines, bool)`

GetEnginesOk returns a tuple with the Engines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngines

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetEngines(v GetDatabasesTypes200ResponseAllOfDataInnerEngines)`

SetEngines sets Engines field to given value.

### HasEngines

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasEngines() bool`

HasEngines returns a boolean if a field has been set.

### GetId

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMemory

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetVcpus

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *GetDatabasesTypes200ResponseAllOfDataInner) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


