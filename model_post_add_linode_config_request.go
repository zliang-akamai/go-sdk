/*
Linode API

[Read the API documentation](https://techdocs.akamai.com/linode-api/reference/api).

API version: 4.189.2
Contact: support@linode.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PostAddLinodeConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostAddLinodeConfigRequest{}

// PostAddLinodeConfigRequest struct for PostAddLinodeConfigRequest
type PostAddLinodeConfigRequest struct {
	// Optional field for arbitrary User comments on this Config.
	Comments NullableString `json:"comments,omitempty"`
	Devices GetLinodeConfigs200ResponseDataInnerDevices `json:"devices"`
	Helpers *GetLinodeConfigs200ResponseDataInnerHelpers `json:"helpers,omitempty"`
	// The ID of this Config.
	Id *int32 `json:"id,omitempty"`
	// An array of Network Interfaces to add to this Linode's Configuration Profile. At least one and up to three Interface objects can exist in this array. The position in the array determines which of the Linode's network Interfaces is configured:  - First [0]:  eth0 - Second [1]: eth1 - Third [2]:  eth2  When updating a Linode's Interfaces, _each Interface must be redefined_. An empty `interfaces` array results in a default `public` type Interface configuration only.  If no public Interface is configured, public IP addresses are still assigned to the Linode but will not be usable without manual configuration.  __Note__. Changes to Linode Interface configurations can be enabled by rebooting the Linode.  `vpc` details  See the [VPC documentation](https://www.linode.com/docs/products/networking/vpc/#technical-specifications) guide for its specifications and limitations.  `vlan` details  - Only Next Generation Network (NGN) data centers support VLANs. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to view the capabilities of data center regions. If a VLAN is attached to your Linode and you attempt to migrate or clone it to a non-NGN data center, the migration or cloning will not initiate. If a Linode cannot be migrated or cloned because of an incompatibility, you will be prompted to select a different data center or contact support. - See the [VLANs Overview](https://www.linode.com/docs/products/networking/vlans/#technical-specifications) guide to view additional specifications and limitations.
	Interfaces []PostLinodeInstanceRequestAllOfInterfacesInner `json:"interfaces,omitempty"`
	// A Kernel ID to boot a Linode with. Here are examples of commonly used kernels:  - `linode/latest-64bit` (default): Our latest kernel at the time of instance boot/reboot. - `linode/grub2`: The upstream distribution-supplied kernel that is installed on the primary disk, or a custom kernel if installed. - `linode/direct-disk`: The MBR (Master Boot Record) of the primary disk/root device, used instead of a Linux kernel.  For a complete list of options, run the [List kernels](https://techdocs.akamai.com/linode-api/reference/get-kernels) operation.
	Kernel *string `json:"kernel,omitempty"`
	// The Config's label is for display purposes only.
	Label string `json:"label"`
	// Defaults to the total RAM of the Linode.
	MemoryLimit *int32 `json:"memory_limit,omitempty"`
	// The root device to boot.  - If no value or an invalid value is provided, root device will default to `/dev/sda`. - If the device specified at the root device location is not mounted, the Linode will not boot until a device is mounted.
	RootDevice *string `json:"root_device,omitempty" validate:"regexp=a-z, A-Z, 0-9, \\/, _, -"`
	// Defines the state of your Linode after booting. Defaults to `default`.
	RunLevel *string `json:"run_level,omitempty"`
	// Controls the virtualization mode. Defaults to `paravirt`.  - `paravirt` is suitable for most cases. Linodes running in paravirt mode share some qualities with the host, ultimately making it run faster since there is less transition between it and the host. - `fullvirt` affords more customization, but is slower because 100% of the VM is virtualized.
	VirtMode *string `json:"virt_mode,omitempty"`
}

type _PostAddLinodeConfigRequest PostAddLinodeConfigRequest

// NewPostAddLinodeConfigRequest instantiates a new PostAddLinodeConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostAddLinodeConfigRequest(devices GetLinodeConfigs200ResponseDataInnerDevices, label string) *PostAddLinodeConfigRequest {
	this := PostAddLinodeConfigRequest{}
	this.Devices = devices
	var kernel string = "linode/latest-64bit"
	this.Kernel = &kernel
	this.Label = label
	return &this
}

// NewPostAddLinodeConfigRequestWithDefaults instantiates a new PostAddLinodeConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostAddLinodeConfigRequestWithDefaults() *PostAddLinodeConfigRequest {
	this := PostAddLinodeConfigRequest{}
	var kernel string = "linode/latest-64bit"
	this.Kernel = &kernel
	return &this
}

// GetComments returns the Comments field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostAddLinodeConfigRequest) GetComments() string {
	if o == nil || IsNil(o.Comments.Get()) {
		var ret string
		return ret
	}
	return *o.Comments.Get()
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostAddLinodeConfigRequest) GetCommentsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Comments.Get(), o.Comments.IsSet()
}

// HasComments returns a boolean if a field has been set.
func (o *PostAddLinodeConfigRequest) HasComments() bool {
	if o != nil && o.Comments.IsSet() {
		return true
	}

	return false
}

// SetComments gets a reference to the given NullableString and assigns it to the Comments field.
func (o *PostAddLinodeConfigRequest) SetComments(v string) {
	o.Comments.Set(&v)
}
// SetCommentsNil sets the value for Comments to be an explicit nil
func (o *PostAddLinodeConfigRequest) SetCommentsNil() {
	o.Comments.Set(nil)
}

// UnsetComments ensures that no value is present for Comments, not even an explicit nil
func (o *PostAddLinodeConfigRequest) UnsetComments() {
	o.Comments.Unset()
}

// GetDevices returns the Devices field value
func (o *PostAddLinodeConfigRequest) GetDevices() GetLinodeConfigs200ResponseDataInnerDevices {
	if o == nil {
		var ret GetLinodeConfigs200ResponseDataInnerDevices
		return ret
	}

	return o.Devices
}

// GetDevicesOk returns a tuple with the Devices field value
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetDevicesOk() (*GetLinodeConfigs200ResponseDataInnerDevices, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Devices, true
}

// SetDevices sets field value
func (o *PostAddLinodeConfigRequest) SetDevices(v GetLinodeConfigs200ResponseDataInnerDevices) {
	o.Devices = v
}

// GetHelpers returns the Helpers field value if set, zero value otherwise.
func (o *PostAddLinodeConfigRequest) GetHelpers() GetLinodeConfigs200ResponseDataInnerHelpers {
	if o == nil || IsNil(o.Helpers) {
		var ret GetLinodeConfigs200ResponseDataInnerHelpers
		return ret
	}
	return *o.Helpers
}

// GetHelpersOk returns a tuple with the Helpers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetHelpersOk() (*GetLinodeConfigs200ResponseDataInnerHelpers, bool) {
	if o == nil || IsNil(o.Helpers) {
		return nil, false
	}
	return o.Helpers, true
}

// HasHelpers returns a boolean if a field has been set.
func (o *PostAddLinodeConfigRequest) HasHelpers() bool {
	if o != nil && !IsNil(o.Helpers) {
		return true
	}

	return false
}

// SetHelpers gets a reference to the given GetLinodeConfigs200ResponseDataInnerHelpers and assigns it to the Helpers field.
func (o *PostAddLinodeConfigRequest) SetHelpers(v GetLinodeConfigs200ResponseDataInnerHelpers) {
	o.Helpers = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostAddLinodeConfigRequest) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostAddLinodeConfigRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PostAddLinodeConfigRequest) SetId(v int32) {
	o.Id = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *PostAddLinodeConfigRequest) GetInterfaces() []PostLinodeInstanceRequestAllOfInterfacesInner {
	if o == nil || IsNil(o.Interfaces) {
		var ret []PostLinodeInstanceRequestAllOfInterfacesInner
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetInterfacesOk() ([]PostLinodeInstanceRequestAllOfInterfacesInner, bool) {
	if o == nil || IsNil(o.Interfaces) {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *PostAddLinodeConfigRequest) HasInterfaces() bool {
	if o != nil && !IsNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []PostLinodeInstanceRequestAllOfInterfacesInner and assigns it to the Interfaces field.
func (o *PostAddLinodeConfigRequest) SetInterfaces(v []PostLinodeInstanceRequestAllOfInterfacesInner) {
	o.Interfaces = v
}

// GetKernel returns the Kernel field value if set, zero value otherwise.
func (o *PostAddLinodeConfigRequest) GetKernel() string {
	if o == nil || IsNil(o.Kernel) {
		var ret string
		return ret
	}
	return *o.Kernel
}

// GetKernelOk returns a tuple with the Kernel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetKernelOk() (*string, bool) {
	if o == nil || IsNil(o.Kernel) {
		return nil, false
	}
	return o.Kernel, true
}

// HasKernel returns a boolean if a field has been set.
func (o *PostAddLinodeConfigRequest) HasKernel() bool {
	if o != nil && !IsNil(o.Kernel) {
		return true
	}

	return false
}

// SetKernel gets a reference to the given string and assigns it to the Kernel field.
func (o *PostAddLinodeConfigRequest) SetKernel(v string) {
	o.Kernel = &v
}

// GetLabel returns the Label field value
func (o *PostAddLinodeConfigRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PostAddLinodeConfigRequest) SetLabel(v string) {
	o.Label = v
}

// GetMemoryLimit returns the MemoryLimit field value if set, zero value otherwise.
func (o *PostAddLinodeConfigRequest) GetMemoryLimit() int32 {
	if o == nil || IsNil(o.MemoryLimit) {
		var ret int32
		return ret
	}
	return *o.MemoryLimit
}

// GetMemoryLimitOk returns a tuple with the MemoryLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetMemoryLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.MemoryLimit) {
		return nil, false
	}
	return o.MemoryLimit, true
}

// HasMemoryLimit returns a boolean if a field has been set.
func (o *PostAddLinodeConfigRequest) HasMemoryLimit() bool {
	if o != nil && !IsNil(o.MemoryLimit) {
		return true
	}

	return false
}

// SetMemoryLimit gets a reference to the given int32 and assigns it to the MemoryLimit field.
func (o *PostAddLinodeConfigRequest) SetMemoryLimit(v int32) {
	o.MemoryLimit = &v
}

// GetRootDevice returns the RootDevice field value if set, zero value otherwise.
func (o *PostAddLinodeConfigRequest) GetRootDevice() string {
	if o == nil || IsNil(o.RootDevice) {
		var ret string
		return ret
	}
	return *o.RootDevice
}

// GetRootDeviceOk returns a tuple with the RootDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetRootDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.RootDevice) {
		return nil, false
	}
	return o.RootDevice, true
}

// HasRootDevice returns a boolean if a field has been set.
func (o *PostAddLinodeConfigRequest) HasRootDevice() bool {
	if o != nil && !IsNil(o.RootDevice) {
		return true
	}

	return false
}

// SetRootDevice gets a reference to the given string and assigns it to the RootDevice field.
func (o *PostAddLinodeConfigRequest) SetRootDevice(v string) {
	o.RootDevice = &v
}

// GetRunLevel returns the RunLevel field value if set, zero value otherwise.
func (o *PostAddLinodeConfigRequest) GetRunLevel() string {
	if o == nil || IsNil(o.RunLevel) {
		var ret string
		return ret
	}
	return *o.RunLevel
}

// GetRunLevelOk returns a tuple with the RunLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetRunLevelOk() (*string, bool) {
	if o == nil || IsNil(o.RunLevel) {
		return nil, false
	}
	return o.RunLevel, true
}

// HasRunLevel returns a boolean if a field has been set.
func (o *PostAddLinodeConfigRequest) HasRunLevel() bool {
	if o != nil && !IsNil(o.RunLevel) {
		return true
	}

	return false
}

// SetRunLevel gets a reference to the given string and assigns it to the RunLevel field.
func (o *PostAddLinodeConfigRequest) SetRunLevel(v string) {
	o.RunLevel = &v
}

// GetVirtMode returns the VirtMode field value if set, zero value otherwise.
func (o *PostAddLinodeConfigRequest) GetVirtMode() string {
	if o == nil || IsNil(o.VirtMode) {
		var ret string
		return ret
	}
	return *o.VirtMode
}

// GetVirtModeOk returns a tuple with the VirtMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostAddLinodeConfigRequest) GetVirtModeOk() (*string, bool) {
	if o == nil || IsNil(o.VirtMode) {
		return nil, false
	}
	return o.VirtMode, true
}

// HasVirtMode returns a boolean if a field has been set.
func (o *PostAddLinodeConfigRequest) HasVirtMode() bool {
	if o != nil && !IsNil(o.VirtMode) {
		return true
	}

	return false
}

// SetVirtMode gets a reference to the given string and assigns it to the VirtMode field.
func (o *PostAddLinodeConfigRequest) SetVirtMode(v string) {
	o.VirtMode = &v
}

func (o PostAddLinodeConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostAddLinodeConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Comments.IsSet() {
		toSerialize["comments"] = o.Comments.Get()
	}
	toSerialize["devices"] = o.Devices
	if !IsNil(o.Helpers) {
		toSerialize["helpers"] = o.Helpers
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	if !IsNil(o.Kernel) {
		toSerialize["kernel"] = o.Kernel
	}
	toSerialize["label"] = o.Label
	if !IsNil(o.MemoryLimit) {
		toSerialize["memory_limit"] = o.MemoryLimit
	}
	if !IsNil(o.RootDevice) {
		toSerialize["root_device"] = o.RootDevice
	}
	if !IsNil(o.RunLevel) {
		toSerialize["run_level"] = o.RunLevel
	}
	if !IsNil(o.VirtMode) {
		toSerialize["virt_mode"] = o.VirtMode
	}
	return toSerialize, nil
}

func (o *PostAddLinodeConfigRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"devices",
		"label",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostAddLinodeConfigRequest := _PostAddLinodeConfigRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostAddLinodeConfigRequest)

	if err != nil {
		return err
	}

	*o = PostAddLinodeConfigRequest(varPostAddLinodeConfigRequest)

	return err
}

type NullablePostAddLinodeConfigRequest struct {
	value *PostAddLinodeConfigRequest
	isSet bool
}

func (v NullablePostAddLinodeConfigRequest) Get() *PostAddLinodeConfigRequest {
	return v.value
}

func (v *NullablePostAddLinodeConfigRequest) Set(val *PostAddLinodeConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostAddLinodeConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostAddLinodeConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostAddLinodeConfigRequest(val *PostAddLinodeConfigRequest) *NullablePostAddLinodeConfigRequest {
	return &NullablePostAddLinodeConfigRequest{value: val, isSet: true}
}

func (v NullablePostAddLinodeConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostAddLinodeConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


