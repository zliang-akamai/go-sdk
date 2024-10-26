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

// checks if the AccountAvailability type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAvailability{}

// AccountAvailability Account Service Availability object.
type AccountAvailability struct {
	// A list of services _available_ to your account in the `region`.
	Available []string `json:"available,omitempty"`
	// The Akamai cloud computing data center (region), represented by a slug value. You can view a full list of regions and their associated slugs with the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation.
	Region *string `json:"region,omitempty"`
	// A list of services _unavailable_ to your account in the `region`.
	Unavailable []string `json:"unavailable,omitempty"`
}

// NewAccountAvailability instantiates a new AccountAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAvailability() *AccountAvailability {
	this := AccountAvailability{}
	return &this
}

// NewAccountAvailabilityWithDefaults instantiates a new AccountAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAvailabilityWithDefaults() *AccountAvailability {
	this := AccountAvailability{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *AccountAvailability) GetAvailable() []string {
	if o == nil || IsNil(o.Available) {
		var ret []string
		return ret
	}
	return o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAvailability) GetAvailableOk() ([]string, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *AccountAvailability) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given []string and assigns it to the Available field.
func (o *AccountAvailability) SetAvailable(v []string) {
	o.Available = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AccountAvailability) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAvailability) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *AccountAvailability) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AccountAvailability) SetRegion(v string) {
	o.Region = &v
}

// GetUnavailable returns the Unavailable field value if set, zero value otherwise.
func (o *AccountAvailability) GetUnavailable() []string {
	if o == nil || IsNil(o.Unavailable) {
		var ret []string
		return ret
	}
	return o.Unavailable
}

// GetUnavailableOk returns a tuple with the Unavailable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAvailability) GetUnavailableOk() ([]string, bool) {
	if o == nil || IsNil(o.Unavailable) {
		return nil, false
	}
	return o.Unavailable, true
}

// HasUnavailable returns a boolean if a field has been set.
func (o *AccountAvailability) HasUnavailable() bool {
	if o != nil && !IsNil(o.Unavailable) {
		return true
	}

	return false
}

// SetUnavailable gets a reference to the given []string and assigns it to the Unavailable field.
func (o *AccountAvailability) SetUnavailable(v []string) {
	o.Unavailable = v
}

func (o AccountAvailability) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAvailability) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Unavailable) {
		toSerialize["unavailable"] = o.Unavailable
	}
	return toSerialize, nil
}

type NullableAccountAvailability struct {
	value *AccountAvailability
	isSet bool
}

func (v NullableAccountAvailability) Get() *AccountAvailability {
	return v.value
}

func (v *NullableAccountAvailability) Set(val *AccountAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAvailability(val *AccountAvailability) *NullableAccountAvailability {
	return &NullableAccountAvailability{value: val, isSet: true}
}

func (v NullableAccountAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

