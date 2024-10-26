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

// checks if the GetServiceTransfers200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetServiceTransfers200ResponseDataInner{}

// GetServiceTransfers200ResponseDataInner An object representing a Service Transfer.
type GetServiceTransfers200ResponseDataInner struct {
	// When this transfer was created.
	Created *time.Time `json:"created,omitempty"`
	Entities *GetServiceTransfers200ResponseDataInnerEntities `json:"entities,omitempty"`
	// When this transfer expires. Transfers will automatically expire 24 hours after creation.
	Expiry *time.Time `json:"expiry,omitempty"`
	// If the requesting account created this transfer.
	IsSender *bool `json:"is_sender,omitempty"`
	// The status of the transfer request.  `accepted`: The transfer has been accepted by another user and is currently in progress. Transfers can take up to 3 hours to complete.  `canceled`: The transfer has been canceled by the sender.  `completed`: The transfer has completed successfully.  `failed`: The transfer has failed after initiation.  `pending`: The transfer is ready to be accepted.  `stale`: The transfer has exceeded its expiration date. It can no longer be accepted or canceled.
	Status *string `json:"status,omitempty"`
	// The token used to identify and accept or cancel this transfer.
	Token *string `json:"token,omitempty"`
	// When this transfer was last updated.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewGetServiceTransfers200ResponseDataInner instantiates a new GetServiceTransfers200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetServiceTransfers200ResponseDataInner() *GetServiceTransfers200ResponseDataInner {
	this := GetServiceTransfers200ResponseDataInner{}
	return &this
}

// NewGetServiceTransfers200ResponseDataInnerWithDefaults instantiates a new GetServiceTransfers200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetServiceTransfers200ResponseDataInnerWithDefaults() *GetServiceTransfers200ResponseDataInner {
	this := GetServiceTransfers200ResponseDataInner{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetServiceTransfers200ResponseDataInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceTransfers200ResponseDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetServiceTransfers200ResponseDataInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetServiceTransfers200ResponseDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEntities returns the Entities field value if set, zero value otherwise.
func (o *GetServiceTransfers200ResponseDataInner) GetEntities() GetServiceTransfers200ResponseDataInnerEntities {
	if o == nil || IsNil(o.Entities) {
		var ret GetServiceTransfers200ResponseDataInnerEntities
		return ret
	}
	return *o.Entities
}

// GetEntitiesOk returns a tuple with the Entities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceTransfers200ResponseDataInner) GetEntitiesOk() (*GetServiceTransfers200ResponseDataInnerEntities, bool) {
	if o == nil || IsNil(o.Entities) {
		return nil, false
	}
	return o.Entities, true
}

// HasEntities returns a boolean if a field has been set.
func (o *GetServiceTransfers200ResponseDataInner) HasEntities() bool {
	if o != nil && !IsNil(o.Entities) {
		return true
	}

	return false
}

// SetEntities gets a reference to the given GetServiceTransfers200ResponseDataInnerEntities and assigns it to the Entities field.
func (o *GetServiceTransfers200ResponseDataInner) SetEntities(v GetServiceTransfers200ResponseDataInnerEntities) {
	o.Entities = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *GetServiceTransfers200ResponseDataInner) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceTransfers200ResponseDataInner) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *GetServiceTransfers200ResponseDataInner) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *GetServiceTransfers200ResponseDataInner) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetIsSender returns the IsSender field value if set, zero value otherwise.
func (o *GetServiceTransfers200ResponseDataInner) GetIsSender() bool {
	if o == nil || IsNil(o.IsSender) {
		var ret bool
		return ret
	}
	return *o.IsSender
}

// GetIsSenderOk returns a tuple with the IsSender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceTransfers200ResponseDataInner) GetIsSenderOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSender) {
		return nil, false
	}
	return o.IsSender, true
}

// HasIsSender returns a boolean if a field has been set.
func (o *GetServiceTransfers200ResponseDataInner) HasIsSender() bool {
	if o != nil && !IsNil(o.IsSender) {
		return true
	}

	return false
}

// SetIsSender gets a reference to the given bool and assigns it to the IsSender field.
func (o *GetServiceTransfers200ResponseDataInner) SetIsSender(v bool) {
	o.IsSender = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetServiceTransfers200ResponseDataInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceTransfers200ResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetServiceTransfers200ResponseDataInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetServiceTransfers200ResponseDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *GetServiceTransfers200ResponseDataInner) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceTransfers200ResponseDataInner) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *GetServiceTransfers200ResponseDataInner) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *GetServiceTransfers200ResponseDataInner) SetToken(v string) {
	o.Token = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *GetServiceTransfers200ResponseDataInner) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetServiceTransfers200ResponseDataInner) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *GetServiceTransfers200ResponseDataInner) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *GetServiceTransfers200ResponseDataInner) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o GetServiceTransfers200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetServiceTransfers200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Entities) {
		toSerialize["entities"] = o.Entities
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.IsSender) {
		toSerialize["is_sender"] = o.IsSender
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	return toSerialize, nil
}

type NullableGetServiceTransfers200ResponseDataInner struct {
	value *GetServiceTransfers200ResponseDataInner
	isSet bool
}

func (v NullableGetServiceTransfers200ResponseDataInner) Get() *GetServiceTransfers200ResponseDataInner {
	return v.value
}

func (v *NullableGetServiceTransfers200ResponseDataInner) Set(val *GetServiceTransfers200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetServiceTransfers200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetServiceTransfers200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetServiceTransfers200ResponseDataInner(val *GetServiceTransfers200ResponseDataInner) *NullableGetServiceTransfers200ResponseDataInner {
	return &NullableGetServiceTransfers200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetServiceTransfers200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetServiceTransfers200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


