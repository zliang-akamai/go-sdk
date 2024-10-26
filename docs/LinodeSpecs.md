# LinodeSpecs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | Pointer to **int32** | The amount of storage space, in MB, this Linode has access to. A typical Linode divides this space between a primary disk with an &#x60;image&#x60; deployed to it, and a swap disk, usually 512 MB. This is the default configuration created when deploying a Linode with an &#x60;image&#x60; through [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance). While this configuration is suitable for 99% of use cases, if you need finer control over your Linode&#39;s disks, see the [List disks](https://techdocs.akamai.com/linode-api/reference/get-linode-disks) operations. | [optional] [readonly] 
**Gpus** | Pointer to **int32** | The number of GPUs this Linode has access to. | [optional] [readonly] 
**Memory** | Pointer to **int32** | The amount of RAM, in MB, this Linode has access to.  Typically, a Linode boots with all of its available RAM, but this can be configured in a config profile. See the [List config profiles](https://techdocs.akamai.com/linode-api/reference/get-linode-configs) operation for more information. | [optional] [readonly] 
**Transfer** | Pointer to **int32** | The amount of network transfer this Linode is allotted each month. | [optional] [readonly] 
**Vcpus** | Pointer to **int32** | The number of VCPUs this Linode has access to. | [optional] [readonly] 

## Methods

### NewLinodeSpecs

`func NewLinodeSpecs() *LinodeSpecs`

NewLinodeSpecs instantiates a new LinodeSpecs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeSpecsWithDefaults

`func NewLinodeSpecsWithDefaults() *LinodeSpecs`

NewLinodeSpecsWithDefaults instantiates a new LinodeSpecs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *LinodeSpecs) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *LinodeSpecs) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *LinodeSpecs) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *LinodeSpecs) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetGpus

`func (o *LinodeSpecs) GetGpus() int32`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *LinodeSpecs) GetGpusOk() (*int32, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *LinodeSpecs) SetGpus(v int32)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *LinodeSpecs) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetMemory

`func (o *LinodeSpecs) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *LinodeSpecs) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *LinodeSpecs) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *LinodeSpecs) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetTransfer

`func (o *LinodeSpecs) GetTransfer() int32`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *LinodeSpecs) GetTransferOk() (*int32, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *LinodeSpecs) SetTransfer(v int32)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *LinodeSpecs) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetVcpus

`func (o *LinodeSpecs) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *LinodeSpecs) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *LinodeSpecs) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *LinodeSpecs) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


