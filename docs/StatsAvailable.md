# StatsAvailable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | CPU usage stats from the last 24 hours. | [optional] 
**Disk** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | Disk usage stats from the last 24 hours. | [optional] 
**NetIn** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | Inbound network traffic stats from the last 24 hours. | [optional] 
**NetOut** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | Outbound network traffic stats from the last 24 hours. | [optional] 
**Swap** | Pointer to [**[]StatsAvailableCpuInner**](StatsAvailableCpuInner.md) | Swap usage stats from the last 24 hours. | [optional] 

## Methods

### NewStatsAvailable

`func NewStatsAvailable() *StatsAvailable`

NewStatsAvailable instantiates a new StatsAvailable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatsAvailableWithDefaults

`func NewStatsAvailableWithDefaults() *StatsAvailable`

NewStatsAvailableWithDefaults instantiates a new StatsAvailable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *StatsAvailable) GetCpu() []StatsAvailableCpuInner`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *StatsAvailable) GetCpuOk() (*[]StatsAvailableCpuInner, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *StatsAvailable) SetCpu(v []StatsAvailableCpuInner)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *StatsAvailable) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetDisk

`func (o *StatsAvailable) GetDisk() []StatsAvailableCpuInner`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *StatsAvailable) GetDiskOk() (*[]StatsAvailableCpuInner, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *StatsAvailable) SetDisk(v []StatsAvailableCpuInner)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *StatsAvailable) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetNetIn

`func (o *StatsAvailable) GetNetIn() []StatsAvailableCpuInner`

GetNetIn returns the NetIn field if non-nil, zero value otherwise.

### GetNetInOk

`func (o *StatsAvailable) GetNetInOk() (*[]StatsAvailableCpuInner, bool)`

GetNetInOk returns a tuple with the NetIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetIn

`func (o *StatsAvailable) SetNetIn(v []StatsAvailableCpuInner)`

SetNetIn sets NetIn field to given value.

### HasNetIn

`func (o *StatsAvailable) HasNetIn() bool`

HasNetIn returns a boolean if a field has been set.

### GetNetOut

`func (o *StatsAvailable) GetNetOut() []StatsAvailableCpuInner`

GetNetOut returns the NetOut field if non-nil, zero value otherwise.

### GetNetOutOk

`func (o *StatsAvailable) GetNetOutOk() (*[]StatsAvailableCpuInner, bool)`

GetNetOutOk returns a tuple with the NetOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetOut

`func (o *StatsAvailable) SetNetOut(v []StatsAvailableCpuInner)`

SetNetOut sets NetOut field to given value.

### HasNetOut

`func (o *StatsAvailable) HasNetOut() bool`

HasNetOut returns a boolean if a field has been set.

### GetSwap

`func (o *StatsAvailable) GetSwap() []StatsAvailableCpuInner`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *StatsAvailable) GetSwapOk() (*[]StatsAvailableCpuInner, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *StatsAvailable) SetSwap(v []StatsAvailableCpuInner)`

SetSwap sets Swap field to given value.

### HasSwap

`func (o *StatsAvailable) HasSwap() bool`

HasSwap returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


