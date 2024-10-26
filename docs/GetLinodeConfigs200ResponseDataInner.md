# GetLinodeConfigs200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **NullableString** | Optional field for arbitrary User comments on this Config. | [optional] 
**Devices** | Pointer to [**GetLinodeConfigs200ResponseDataInnerDevices**](GetLinodeConfigs200ResponseDataInnerDevices.md) |  | [optional] 
**Helpers** | Pointer to [**GetLinodeConfigs200ResponseDataInnerHelpers**](GetLinodeConfigs200ResponseDataInnerHelpers.md) |  | [optional] 
**Id** | Pointer to **int32** | The ID of this Config. | [optional] [readonly] 
**Interfaces** | Pointer to [**[]PostLinodeInstanceRequestAllOfInterfacesInner**](PostLinodeInstanceRequestAllOfInterfacesInner.md) | An array of Network Interfaces to add to this Linode&#39;s Configuration Profile. At least one and up to three Interface objects can exist in this array. The position in the array determines which of the Linode&#39;s network Interfaces is configured:  - First [0]:  eth0 - Second [1]: eth1 - Third [2]:  eth2  When updating a Linode&#39;s Interfaces, _each Interface must be redefined_. An empty &#x60;interfaces&#x60; array results in a default &#x60;public&#x60; type Interface configuration only.  If no public Interface is configured, public IP addresses are still assigned to the Linode but will not be usable without manual configuration.  __Note__. Changes to Linode Interface configurations can be enabled by rebooting the Linode.  &#x60;vpc&#x60; details  See the [VPC documentation](https://www.linode.com/docs/products/networking/vpc/#technical-specifications) guide for its specifications and limitations.  &#x60;vlan&#x60; details  - Only Next Generation Network (NGN) data centers support VLANs. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to view the capabilities of data center regions. If a VLAN is attached to your Linode and you attempt to migrate or clone it to a non-NGN data center, the migration or cloning will not initiate. If a Linode cannot be migrated or cloned because of an incompatibility, you will be prompted to select a different data center or contact support. - See the [VLANs Overview](https://www.linode.com/docs/products/networking/vlans/#technical-specifications) guide to view additional specifications and limitations. | [optional] 
**Kernel** | Pointer to **string** | A Kernel ID to boot a Linode with. Here are examples of commonly used kernels:  - &#x60;linode/latest-64bit&#x60; (default): Our latest kernel at the time of instance boot/reboot. - &#x60;linode/grub2&#x60;: The upstream distribution-supplied kernel that is installed on the primary disk, or a custom kernel if installed. - &#x60;linode/direct-disk&#x60;: The MBR (Master Boot Record) of the primary disk/root device, used instead of a Linux kernel.  For a complete list of options, run the [List kernels](https://techdocs.akamai.com/linode-api/reference/get-kernels) operation. | [optional] [default to "linode/latest-64bit"]
**Label** | Pointer to **string** | The Config&#39;s label is for display purposes only. | [optional] 
**MemoryLimit** | Pointer to **int32** | Defaults to the total RAM of the Linode. | [optional] 
**RootDevice** | Pointer to **string** | The root device to boot.  - If no value or an invalid value is provided, root device will default to &#x60;/dev/sda&#x60;. - If the device specified at the root device location is not mounted, the Linode will not boot until a device is mounted. | [optional] 
**RunLevel** | Pointer to **string** | Defines the state of your Linode after booting. Defaults to &#x60;default&#x60;. | [optional] 
**VirtMode** | Pointer to **string** | Controls the virtualization mode. Defaults to &#x60;paravirt&#x60;.  - &#x60;paravirt&#x60; is suitable for most cases. Linodes running in paravirt mode share some qualities with the host, ultimately making it run faster since there is less transition between it and the host. - &#x60;fullvirt&#x60; affords more customization, but is slower because 100% of the VM is virtualized. | [optional] 

## Methods

### NewGetLinodeConfigs200ResponseDataInner

`func NewGetLinodeConfigs200ResponseDataInner() *GetLinodeConfigs200ResponseDataInner`

NewGetLinodeConfigs200ResponseDataInner instantiates a new GetLinodeConfigs200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeConfigs200ResponseDataInnerWithDefaults

`func NewGetLinodeConfigs200ResponseDataInnerWithDefaults() *GetLinodeConfigs200ResponseDataInner`

NewGetLinodeConfigs200ResponseDataInnerWithDefaults instantiates a new GetLinodeConfigs200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *GetLinodeConfigs200ResponseDataInner) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetLinodeConfigs200ResponseDataInner) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *GetLinodeConfigs200ResponseDataInner) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *GetLinodeConfigs200ResponseDataInner) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *GetLinodeConfigs200ResponseDataInner) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetDevices

`func (o *GetLinodeConfigs200ResponseDataInner) GetDevices() GetLinodeConfigs200ResponseDataInnerDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetDevicesOk() (*GetLinodeConfigs200ResponseDataInnerDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *GetLinodeConfigs200ResponseDataInner) SetDevices(v GetLinodeConfigs200ResponseDataInnerDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *GetLinodeConfigs200ResponseDataInner) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetHelpers

`func (o *GetLinodeConfigs200ResponseDataInner) GetHelpers() GetLinodeConfigs200ResponseDataInnerHelpers`

GetHelpers returns the Helpers field if non-nil, zero value otherwise.

### GetHelpersOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetHelpersOk() (*GetLinodeConfigs200ResponseDataInnerHelpers, bool)`

GetHelpersOk returns a tuple with the Helpers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpers

`func (o *GetLinodeConfigs200ResponseDataInner) SetHelpers(v GetLinodeConfigs200ResponseDataInnerHelpers)`

SetHelpers sets Helpers field to given value.

### HasHelpers

`func (o *GetLinodeConfigs200ResponseDataInner) HasHelpers() bool`

HasHelpers returns a boolean if a field has been set.

### GetId

`func (o *GetLinodeConfigs200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLinodeConfigs200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetLinodeConfigs200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaces

`func (o *GetLinodeConfigs200ResponseDataInner) GetInterfaces() []PostLinodeInstanceRequestAllOfInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetInterfacesOk() (*[]PostLinodeInstanceRequestAllOfInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *GetLinodeConfigs200ResponseDataInner) SetInterfaces(v []PostLinodeInstanceRequestAllOfInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *GetLinodeConfigs200ResponseDataInner) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetKernel

`func (o *GetLinodeConfigs200ResponseDataInner) GetKernel() string`

GetKernel returns the Kernel field if non-nil, zero value otherwise.

### GetKernelOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetKernelOk() (*string, bool)`

GetKernelOk returns a tuple with the Kernel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKernel

`func (o *GetLinodeConfigs200ResponseDataInner) SetKernel(v string)`

SetKernel sets Kernel field to given value.

### HasKernel

`func (o *GetLinodeConfigs200ResponseDataInner) HasKernel() bool`

HasKernel returns a boolean if a field has been set.

### GetLabel

`func (o *GetLinodeConfigs200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetLinodeConfigs200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetLinodeConfigs200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *GetLinodeConfigs200ResponseDataInner) GetMemoryLimit() int32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetMemoryLimitOk() (*int32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *GetLinodeConfigs200ResponseDataInner) SetMemoryLimit(v int32)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *GetLinodeConfigs200ResponseDataInner) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetRootDevice

`func (o *GetLinodeConfigs200ResponseDataInner) GetRootDevice() string`

GetRootDevice returns the RootDevice field if non-nil, zero value otherwise.

### GetRootDeviceOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetRootDeviceOk() (*string, bool)`

GetRootDeviceOk returns a tuple with the RootDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootDevice

`func (o *GetLinodeConfigs200ResponseDataInner) SetRootDevice(v string)`

SetRootDevice sets RootDevice field to given value.

### HasRootDevice

`func (o *GetLinodeConfigs200ResponseDataInner) HasRootDevice() bool`

HasRootDevice returns a boolean if a field has been set.

### GetRunLevel

`func (o *GetLinodeConfigs200ResponseDataInner) GetRunLevel() string`

GetRunLevel returns the RunLevel field if non-nil, zero value otherwise.

### GetRunLevelOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetRunLevelOk() (*string, bool)`

GetRunLevelOk returns a tuple with the RunLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunLevel

`func (o *GetLinodeConfigs200ResponseDataInner) SetRunLevel(v string)`

SetRunLevel sets RunLevel field to given value.

### HasRunLevel

`func (o *GetLinodeConfigs200ResponseDataInner) HasRunLevel() bool`

HasRunLevel returns a boolean if a field has been set.

### GetVirtMode

`func (o *GetLinodeConfigs200ResponseDataInner) GetVirtMode() string`

GetVirtMode returns the VirtMode field if non-nil, zero value otherwise.

### GetVirtModeOk

`func (o *GetLinodeConfigs200ResponseDataInner) GetVirtModeOk() (*string, bool)`

GetVirtModeOk returns a tuple with the VirtMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtMode

`func (o *GetLinodeConfigs200ResponseDataInner) SetVirtMode(v string)`

SetVirtMode sets VirtMode field to given value.

### HasVirtMode

`func (o *GetLinodeConfigs200ResponseDataInner) HasVirtMode() bool`

HasVirtMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


