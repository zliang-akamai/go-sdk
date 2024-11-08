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

// checks if the GooglePayData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GooglePayData{}

// GooglePayData Google Pay information.
type GooglePayData struct {
	// The type of credit card.
	CardType *string `json:"card_type,omitempty"`
	// The expiration month and year of the credit card.
	Expiry *string `json:"expiry,omitempty"`
	// The last four digits of the credit card number.
	LastFour *string `json:"last_four,omitempty"`
}

// NewGooglePayData instantiates a new GooglePayData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGooglePayData() *GooglePayData {
	this := GooglePayData{}
	return &this
}

// NewGooglePayDataWithDefaults instantiates a new GooglePayData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGooglePayDataWithDefaults() *GooglePayData {
	this := GooglePayData{}
	return &this
}

// GetCardType returns the CardType field value if set, zero value otherwise.
func (o *GooglePayData) GetCardType() string {
	if o == nil || IsNil(o.CardType) {
		var ret string
		return ret
	}
	return *o.CardType
}

// GetCardTypeOk returns a tuple with the CardType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglePayData) GetCardTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CardType) {
		return nil, false
	}
	return o.CardType, true
}

// HasCardType returns a boolean if a field has been set.
func (o *GooglePayData) HasCardType() bool {
	if o != nil && !IsNil(o.CardType) {
		return true
	}

	return false
}

// SetCardType gets a reference to the given string and assigns it to the CardType field.
func (o *GooglePayData) SetCardType(v string) {
	o.CardType = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *GooglePayData) GetExpiry() string {
	if o == nil || IsNil(o.Expiry) {
		var ret string
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglePayData) GetExpiryOk() (*string, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *GooglePayData) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given string and assigns it to the Expiry field.
func (o *GooglePayData) SetExpiry(v string) {
	o.Expiry = &v
}

// GetLastFour returns the LastFour field value if set, zero value otherwise.
func (o *GooglePayData) GetLastFour() string {
	if o == nil || IsNil(o.LastFour) {
		var ret string
		return ret
	}
	return *o.LastFour
}

// GetLastFourOk returns a tuple with the LastFour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GooglePayData) GetLastFourOk() (*string, bool) {
	if o == nil || IsNil(o.LastFour) {
		return nil, false
	}
	return o.LastFour, true
}

// HasLastFour returns a boolean if a field has been set.
func (o *GooglePayData) HasLastFour() bool {
	if o != nil && !IsNil(o.LastFour) {
		return true
	}

	return false
}

// SetLastFour gets a reference to the given string and assigns it to the LastFour field.
func (o *GooglePayData) SetLastFour(v string) {
	o.LastFour = &v
}

func (o GooglePayData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GooglePayData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CardType) {
		toSerialize["card_type"] = o.CardType
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.LastFour) {
		toSerialize["last_four"] = o.LastFour
	}
	return toSerialize, nil
}

type NullableGooglePayData struct {
	value *GooglePayData
	isSet bool
}

func (v NullableGooglePayData) Get() *GooglePayData {
	return v.value
}

func (v *NullableGooglePayData) Set(val *GooglePayData) {
	v.value = val
	v.isSet = true
}

func (v NullableGooglePayData) IsSet() bool {
	return v.isSet
}

func (v *NullableGooglePayData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGooglePayData(val *GooglePayData) *NullableGooglePayData {
	return &NullableGooglePayData{value: val, isSet: true}
}

func (v NullableGooglePayData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGooglePayData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


