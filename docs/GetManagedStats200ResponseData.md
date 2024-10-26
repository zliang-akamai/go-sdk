# GetManagedStats200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | CPU usage stats from the last 24 hours. | [optional] 
**Disk** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | Disk usage stats from the last 24 hours. | [optional] 
**NetIn** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | Inbound network traffic stats from the last 24 hours. | [optional] 
**NetOut** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | Outbound network traffic stats from the last 24 hours. | [optional] 
**Swap** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | Swap usage stats from the last 24 hours. | [optional] 

## Methods

### NewGetManagedStats200ResponseData

`func NewGetManagedStats200ResponseData() *GetManagedStats200ResponseData`

NewGetManagedStats200ResponseData instantiates a new GetManagedStats200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedStats200ResponseDataWithDefaults

`func NewGetManagedStats200ResponseDataWithDefaults() *GetManagedStats200ResponseData`

NewGetManagedStats200ResponseDataWithDefaults instantiates a new GetManagedStats200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *GetManagedStats200ResponseData) GetCpu() []StatsAvailableCpuInner`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *GetManagedStats200ResponseData) GetCpuOk() (*[]StatsAvailableCpuInner, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *GetManagedStats200ResponseData) SetCpu(v []StatsAvailableCpuInner)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *GetManagedStats200ResponseData) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisk

`func (o *GetManagedStats200ResponseData) GetDisk() []StatsAvailableCpuInner`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *GetManagedStats200ResponseData) GetDiskOk() (*[]StatsAvailableCpuInner, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *GetManagedStats200ResponseData) SetDisk(v []StatsAvailableCpuInner)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *GetManagedStats200ResponseData) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetNetIn

`func (o *GetManagedStats200ResponseData) GetNetIn() []StatsAvailableCpuInner`

GetNetIn returns the NetIn field if non-nil, zero value otherwise.

### GetNetInOk

`func (o *GetManagedStats200ResponseData) GetNetInOk() (*[]StatsAvailableCpuInner, bool)`

GetNetInOk returns a tuple with the NetIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIn

`func (o *GetManagedStats200ResponseData) SetNetIn(v []StatsAvailableCpuInner)`

SetNetIn sets NetIn field to given value.

### HasNetIn

`func (o *GetManagedStats200ResponseData) HasNetIn() bool`

HasNetIn returns a boolean if a field has been set.

### GetNetOut

`func (o *GetManagedStats200ResponseData) GetNetOut() []StatsAvailableCpuInner`

GetNetOut returns the NetOut field if non-nil, zero value otherwise.

### GetNetOutOk

`func (o *GetManagedStats200ResponseData) GetNetOutOk() (*[]StatsAvailableCpuInner, bool)`

GetNetOutOk returns a tuple with the NetOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetOut

`func (o *GetManagedStats200ResponseData) SetNetOut(v []StatsAvailableCpuInner)`

SetNetOut sets NetOut field to given value.

### HasNetOut

`func (o *GetManagedStats200ResponseData) HasNetOut() bool`

HasNetOut returns a boolean if a field has been set.

### GetSwap

`func (o *GetManagedStats200ResponseData) GetSwap() []StatsAvailableCpuInner`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *GetManagedStats200ResponseData) GetSwapOk() (*[]StatsAvailableCpuInner, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *GetManagedStats200ResponseData) SetSwap(v []StatsAvailableCpuInner)`

SetSwap sets Swap field to given value.

### HasSwap

`func (o *GetManagedStats200ResponseData) HasSwap() bool`

HasSwap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


