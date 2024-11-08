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
)

// checks if the PostRebuildNodeBalancerConfigRequestAllOfNodesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostRebuildNodeBalancerConfigRequestAllOfNodesInner{}

// PostRebuildNodeBalancerConfigRequestAllOfNodesInner NodeBalancer Node request object including ID.
type PostRebuildNodeBalancerConfigRequestAllOfNodesInner struct {
	// The private IP Address where this backend can be reached. This _must_ be a private IP address.
	Address *string `json:"address,omitempty"`
	// The unique ID of the Node to update.
	Id *int32 `json:"id,omitempty"`
	// The label for this node.  This is for display purposes only.
	Label *string `json:"label,omitempty" validate:"regexp=[a-zA-Z0-9-_.]{3,32}"`
	// The mode this NodeBalancer should use when sending traffic to this backend.  - If set to `accept` this backend is accepting traffic. - If set to `reject` this backend will not receive traffic. - If set to `drain` this backend will not receive _new_ traffic, but connections already pinned to it will continue to be routed to it. - If set to `backup`, this backend will only receive traffic if all `accept` nodes are down.
	Mode *string `json:"mode,omitempty"`
	// Used when picking a backend to serve a request and is not pinned to a single backend yet.  Nodes with a higher weight will receive more traffic.
	Weight *int32 `json:"weight,omitempty"`
}

// NewPostRebuildNodeBalancerConfigRequestAllOfNodesInner instantiates a new PostRebuildNodeBalancerConfigRequestAllOfNodesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostRebuildNodeBalancerConfigRequestAllOfNodesInner() *PostRebuildNodeBalancerConfigRequestAllOfNodesInner {
	this := PostRebuildNodeBalancerConfigRequestAllOfNodesInner{}
	return &this
}

// NewPostRebuildNodeBalancerConfigRequestAllOfNodesInnerWithDefaults instantiates a new PostRebuildNodeBalancerConfigRequestAllOfNodesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostRebuildNodeBalancerConfigRequestAllOfNodesInnerWithDefaults() *PostRebuildNodeBalancerConfigRequestAllOfNodesInner {
	this := PostRebuildNodeBalancerConfigRequestAllOfNodesInner{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetAddress(v string) {
	o.Address = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetLabel(v string) {
	o.Label = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetMode() string {
	if o == nil || IsNil(o.Mode) {
		var ret string
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetModeOk() (*string, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given string and assigns it to the Mode field.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetMode(v string) {
	o.Mode = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetWeight(v int32) {
	o.Weight = &v
}

func (o PostRebuildNodeBalancerConfigRequestAllOfNodesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostRebuildNodeBalancerConfigRequestAllOfNodesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner struct {
	value *PostRebuildNodeBalancerConfigRequestAllOfNodesInner
	isSet bool
}

func (v NullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner) Get() *PostRebuildNodeBalancerConfigRequestAllOfNodesInner {
	return v.value
}

func (v *NullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner) Set(val *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner(val *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) *NullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner {
	return &NullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner{value: val, isSet: true}
}

func (v NullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostRebuildNodeBalancerConfigRequestAllOfNodesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


