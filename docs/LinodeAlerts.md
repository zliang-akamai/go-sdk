# LinodeAlerts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **int32** | The percentage of CPU usage required to trigger an alert. If the average CPU usage over two hours exceeds this value, we&#39;ll send you an alert. Your Linode&#39;s total CPU capacity is represented as 100%, multiplied by its number of cores.  For example, a two core Linode&#39;s CPU capacity is represented as 200%. If you want to be alerted at 90% of a two core Linode&#39;s CPU capacity, set the alert value to &#x60;180&#x60;.  The default value is 90% multiplied by the number of cores.  If the value is set to &#x60;0&#x60; (zero), the alert is disabled. | [optional] 
**Io** | Pointer to **int32** | The amount of disk IO operation per second required to trigger an alert. If the average disk IO over two hours exceeds this value, we&#39;ll send you an alert. If set to &#x60;0&#x60; (zero), this alert is disabled. | [optional] 
**NetworkIn** | Pointer to **int32** | The amount of incoming traffic, in Mbit/s, required to trigger an alert. If the average incoming traffic over two hours exceeds this value, we&#39;ll send you an alert. If this is set to &#x60;0&#x60; (zero), the alert is disabled. | [optional] 
**NetworkOut** | Pointer to **int32** | The amount of outbound traffic, in Mbit/s, required to trigger an alert. If the average outbound traffic over two hours exceeds this value, we&#39;ll send you an alert. If this is set to &#x60;0&#x60; (zero), the alert is disabled. | [optional] 
**TransferQuota** | Pointer to **int32** | The percentage of network transfer that may be used before an alert is triggered. When this value is exceeded, we&#39;ll alert you. If this is set to &#x60;0&#x60; (zero), the alert is disabled. | [optional] 

## Methods

### NewLinodeAlerts

`func NewLinodeAlerts() *LinodeAlerts`

NewLinodeAlerts instantiates a new LinodeAlerts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeAlertsWithDefaults

`func NewLinodeAlertsWithDefaults() *LinodeAlerts`

NewLinodeAlertsWithDefaults instantiates a new LinodeAlerts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *LinodeAlerts) GetCpu() int32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *LinodeAlerts) GetCpuOk() (*int32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *LinodeAlerts) SetCpu(v int32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *LinodeAlerts) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetIo

`func (o *LinodeAlerts) GetIo() int32`

GetIo returns the Io field if non-nil, zero value otherwise.

### GetIoOk

`func (o *LinodeAlerts) GetIoOk() (*int32, bool)`

GetIoOk returns a tuple with the Io field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIo

`func (o *LinodeAlerts) SetIo(v int32)`

SetIo sets Io field to given value.

### HasIo

`func (o *LinodeAlerts) HasIo() bool`

HasIo returns a boolean if a field has been set.

### GetNetworkIn

`func (o *LinodeAlerts) GetNetworkIn() int32`

GetNetworkIn returns the NetworkIn field if non-nil, zero value otherwise.

### GetNetworkInOk

`func (o *LinodeAlerts) GetNetworkInOk() (*int32, bool)`

GetNetworkInOk returns a tuple with the NetworkIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIn

`func (o *LinodeAlerts) SetNetworkIn(v int32)`

SetNetworkIn sets NetworkIn field to given value.

### HasNetworkIn

`func (o *LinodeAlerts) HasNetworkIn() bool`

HasNetworkIn returns a boolean if a field has been set.

### GetNetworkOut

`func (o *LinodeAlerts) GetNetworkOut() int32`

GetNetworkOut returns the NetworkOut field if non-nil, zero value otherwise.

### GetNetworkOutOk

`func (o *LinodeAlerts) GetNetworkOutOk() (*int32, bool)`

GetNetworkOutOk returns a tuple with the NetworkOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOut

`func (o *LinodeAlerts) SetNetworkOut(v int32)`

SetNetworkOut sets NetworkOut field to given value.

### HasNetworkOut

`func (o *LinodeAlerts) HasNetworkOut() bool`

HasNetworkOut returns a boolean if a field has been set.

### GetTransferQuota

`func (o *LinodeAlerts) GetTransferQuota() int32`

GetTransferQuota returns the TransferQuota field if non-nil, zero value otherwise.

### GetTransferQuotaOk

`func (o *LinodeAlerts) GetTransferQuotaOk() (*int32, bool)`

GetTransferQuotaOk returns a tuple with the TransferQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferQuota

`func (o *LinodeAlerts) SetTransferQuota(v int32)`

SetTransferQuota sets TransferQuota field to given value.

### HasTransferQuota

`func (o *LinodeAlerts) HasTransferQuota() bool`

HasTransferQuota returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


