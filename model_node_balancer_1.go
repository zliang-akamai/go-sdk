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
	"time"
)

// checks if the NodeBalancer1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeBalancer1{}

// NodeBalancer1 Linode's load balancing solution.  Can handle multiple ports, SSL termination, and any number of backends.  NodeBalancer ports are configured with NodeBalancer Configs, and each config is given one or more NodeBalancer Node that accepts traffic.  The traffic should be routed to the  NodeBalancer's ip address, the NodeBalancer will handle routing individual requests to backends.
type NodeBalancer1 struct {
	// Throttle connections per second.  Set to 0 (zero) to disable throttling.
	ClientConnThrottle *int32 `json:"client_conn_throttle,omitempty"`
	// When this NodeBalancer was created.
	Created *time.Time `json:"created,omitempty"`
	// This NodeBalancer's hostname, beginning with its IP address and ending with _.ip.linodeusercontent.com_.
	Hostname *string `json:"hostname,omitempty"`
	// This NodeBalancer's unique ID.
	Id *int32 `json:"id,omitempty"`
	// This NodeBalancer's public IPv4 address.
	Ipv4 *string `json:"ipv4,omitempty"`
	// This NodeBalancer's public IPv6 address.
	Ipv6 NullableString `json:"ipv6,omitempty"`
	// This NodeBalancer's label. These must be unique on your Account.
	Label *string `json:"label,omitempty" validate:"regexp=[a-zA-Z0-9-_]{3,32}"`
	// The Region where this NodeBalancer is located. NodeBalancers only support backends in the same Region.
	Region *string `json:"region,omitempty"`
	// An array of Tags applied to this object.  Tags are for organizational purposes only.
	Tags []string `json:"tags,omitempty"`
	Transfer *NodeBalancerTransfer `json:"transfer,omitempty"`
	// When this NodeBalancer was last updated.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewNodeBalancer1 instantiates a new NodeBalancer1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeBalancer1() *NodeBalancer1 {
	this := NodeBalancer1{}
	return &this
}

// NewNodeBalancer1WithDefaults instantiates a new NodeBalancer1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeBalancer1WithDefaults() *NodeBalancer1 {
	this := NodeBalancer1{}
	return &this
}

// GetClientConnThrottle returns the ClientConnThrottle field value if set, zero value otherwise.
func (o *NodeBalancer1) GetClientConnThrottle() int32 {
	if o == nil || IsNil(o.ClientConnThrottle) {
		var ret int32
		return ret
	}
	return *o.ClientConnThrottle
}

// GetClientConnThrottleOk returns a tuple with the ClientConnThrottle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetClientConnThrottleOk() (*int32, bool) {
	if o == nil || IsNil(o.ClientConnThrottle) {
		return nil, false
	}
	return o.ClientConnThrottle, true
}

// HasClientConnThrottle returns a boolean if a field has been set.
func (o *NodeBalancer1) HasClientConnThrottle() bool {
	if o != nil && !IsNil(o.ClientConnThrottle) {
		return true
	}

	return false
}

// SetClientConnThrottle gets a reference to the given int32 and assigns it to the ClientConnThrottle field.
func (o *NodeBalancer1) SetClientConnThrottle(v int32) {
	o.ClientConnThrottle = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *NodeBalancer1) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *NodeBalancer1) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *NodeBalancer1) SetCreated(v time.Time) {
	o.Created = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *NodeBalancer1) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *NodeBalancer1) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *NodeBalancer1) SetHostname(v string) {
	o.Hostname = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NodeBalancer1) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *NodeBalancer1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *NodeBalancer1) SetId(v int32) {
	o.Id = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *NodeBalancer1) GetIpv4() string {
	if o == nil || IsNil(o.Ipv4) {
		var ret string
		return ret
	}
	return *o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetIpv4Ok() (*string, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *NodeBalancer1) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given string and assigns it to the Ipv4 field.
func (o *NodeBalancer1) SetIpv4(v string) {
	o.Ipv4 = &v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NodeBalancer1) GetIpv6() string {
	if o == nil || IsNil(o.Ipv6.Get()) {
		var ret string
		return ret
	}
	return *o.Ipv6.Get()
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NodeBalancer1) GetIpv6Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv6.Get(), o.Ipv6.IsSet()
}

// HasIpv6 returns a boolean if a field has been set.
func (o *NodeBalancer1) HasIpv6() bool {
	if o != nil && o.Ipv6.IsSet() {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given NullableString and assigns it to the Ipv6 field.
func (o *NodeBalancer1) SetIpv6(v string) {
	o.Ipv6.Set(&v)
}
// SetIpv6Nil sets the value for Ipv6 to be an explicit nil
func (o *NodeBalancer1) SetIpv6Nil() {
	o.Ipv6.Set(nil)
}

// UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
func (o *NodeBalancer1) UnsetIpv6() {
	o.Ipv6.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *NodeBalancer1) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *NodeBalancer1) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *NodeBalancer1) SetLabel(v string) {
	o.Label = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *NodeBalancer1) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *NodeBalancer1) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *NodeBalancer1) SetRegion(v string) {
	o.Region = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *NodeBalancer1) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *NodeBalancer1) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *NodeBalancer1) SetTags(v []string) {
	o.Tags = v
}

// GetTransfer returns the Transfer field value if set, zero value otherwise.
func (o *NodeBalancer1) GetTransfer() NodeBalancerTransfer {
	if o == nil || IsNil(o.Transfer) {
		var ret NodeBalancerTransfer
		return ret
	}
	return *o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetTransferOk() (*NodeBalancerTransfer, bool) {
	if o == nil || IsNil(o.Transfer) {
		return nil, false
	}
	return o.Transfer, true
}

// HasTransfer returns a boolean if a field has been set.
func (o *NodeBalancer1) HasTransfer() bool {
	if o != nil && !IsNil(o.Transfer) {
		return true
	}

	return false
}

// SetTransfer gets a reference to the given NodeBalancerTransfer and assigns it to the Transfer field.
func (o *NodeBalancer1) SetTransfer(v NodeBalancerTransfer) {
	o.Transfer = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *NodeBalancer1) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeBalancer1) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *NodeBalancer1) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *NodeBalancer1) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o NodeBalancer1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeBalancer1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientConnThrottle) {
		toSerialize["client_conn_throttle"] = o.ClientConnThrottle
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if o.Ipv6.IsSet() {
		toSerialize["ipv6"] = o.Ipv6.Get()
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Transfer) {
		toSerialize["transfer"] = o.Transfer
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	return toSerialize, nil
}

type NullableNodeBalancer1 struct {
	value *NodeBalancer1
	isSet bool
}

func (v NullableNodeBalancer1) Get() *NodeBalancer1 {
	return v.value
}

func (v *NullableNodeBalancer1) Set(val *NodeBalancer1) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeBalancer1) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeBalancer1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeBalancer1(val *NodeBalancer1) *NullableNodeBalancer1 {
	return &NullableNodeBalancer1{value: val, isSet: true}
}

func (v NullableNodeBalancer1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeBalancer1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

