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

// checks if the GetDevices200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDevices200ResponseDataInner{}

// GetDevices200ResponseDataInner A trusted device object represents an active Remember Me session with [login.linode.com](https://login.linode.com).
type GetDevices200ResponseDataInner struct {
	// When this Remember Me session was started.  This corresponds to the time of login with the \"Remember Me\" box checked.
	Created *time.Time `json:"created,omitempty"`
	// When this TrustedDevice session expires.  Sessions typically last 30 days.
	Expiry *time.Time `json:"expiry,omitempty"`
	// The unique ID for this TrustedDevice.
	Id *int32 `json:"id,omitempty"`
	// The last time this TrustedDevice was successfully used to authenticate to [login.linode.com](https://login.linode.com).
	LastAuthenticated *time.Time `json:"last_authenticated,omitempty"`
	// The last IP Address to successfully authenticate with this TrustedDevice.
	LastRemoteAddr *string `json:"last_remote_addr,omitempty"`
	// The User Agent of the browser that created this TrustedDevice session.
	UserAgent *string `json:"user_agent,omitempty"`
}

// NewGetDevices200ResponseDataInner instantiates a new GetDevices200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDevices200ResponseDataInner() *GetDevices200ResponseDataInner {
	this := GetDevices200ResponseDataInner{}
	return &this
}

// NewGetDevices200ResponseDataInnerWithDefaults instantiates a new GetDevices200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDevices200ResponseDataInnerWithDefaults() *GetDevices200ResponseDataInner {
	this := GetDevices200ResponseDataInner{}
	return &this
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetDevices200ResponseDataInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDevices200ResponseDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetDevices200ResponseDataInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetDevices200ResponseDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *GetDevices200ResponseDataInner) GetExpiry() time.Time {
	if o == nil || IsNil(o.Expiry) {
		var ret time.Time
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDevices200ResponseDataInner) GetExpiryOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Expiry) {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *GetDevices200ResponseDataInner) HasExpiry() bool {
	if o != nil && !IsNil(o.Expiry) {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given time.Time and assigns it to the Expiry field.
func (o *GetDevices200ResponseDataInner) SetExpiry(v time.Time) {
	o.Expiry = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetDevices200ResponseDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDevices200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetDevices200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetDevices200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetLastAuthenticated returns the LastAuthenticated field value if set, zero value otherwise.
func (o *GetDevices200ResponseDataInner) GetLastAuthenticated() time.Time {
	if o == nil || IsNil(o.LastAuthenticated) {
		var ret time.Time
		return ret
	}
	return *o.LastAuthenticated
}

// GetLastAuthenticatedOk returns a tuple with the LastAuthenticated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDevices200ResponseDataInner) GetLastAuthenticatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastAuthenticated) {
		return nil, false
	}
	return o.LastAuthenticated, true
}

// HasLastAuthenticated returns a boolean if a field has been set.
func (o *GetDevices200ResponseDataInner) HasLastAuthenticated() bool {
	if o != nil && !IsNil(o.LastAuthenticated) {
		return true
	}

	return false
}

// SetLastAuthenticated gets a reference to the given time.Time and assigns it to the LastAuthenticated field.
func (o *GetDevices200ResponseDataInner) SetLastAuthenticated(v time.Time) {
	o.LastAuthenticated = &v
}

// GetLastRemoteAddr returns the LastRemoteAddr field value if set, zero value otherwise.
func (o *GetDevices200ResponseDataInner) GetLastRemoteAddr() string {
	if o == nil || IsNil(o.LastRemoteAddr) {
		var ret string
		return ret
	}
	return *o.LastRemoteAddr
}

// GetLastRemoteAddrOk returns a tuple with the LastRemoteAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDevices200ResponseDataInner) GetLastRemoteAddrOk() (*string, bool) {
	if o == nil || IsNil(o.LastRemoteAddr) {
		return nil, false
	}
	return o.LastRemoteAddr, true
}

// HasLastRemoteAddr returns a boolean if a field has been set.
func (o *GetDevices200ResponseDataInner) HasLastRemoteAddr() bool {
	if o != nil && !IsNil(o.LastRemoteAddr) {
		return true
	}

	return false
}

// SetLastRemoteAddr gets a reference to the given string and assigns it to the LastRemoteAddr field.
func (o *GetDevices200ResponseDataInner) SetLastRemoteAddr(v string) {
	o.LastRemoteAddr = &v
}

// GetUserAgent returns the UserAgent field value if set, zero value otherwise.
func (o *GetDevices200ResponseDataInner) GetUserAgent() string {
	if o == nil || IsNil(o.UserAgent) {
		var ret string
		return ret
	}
	return *o.UserAgent
}

// GetUserAgentOk returns a tuple with the UserAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDevices200ResponseDataInner) GetUserAgentOk() (*string, bool) {
	if o == nil || IsNil(o.UserAgent) {
		return nil, false
	}
	return o.UserAgent, true
}

// HasUserAgent returns a boolean if a field has been set.
func (o *GetDevices200ResponseDataInner) HasUserAgent() bool {
	if o != nil && !IsNil(o.UserAgent) {
		return true
	}

	return false
}

// SetUserAgent gets a reference to the given string and assigns it to the UserAgent field.
func (o *GetDevices200ResponseDataInner) SetUserAgent(v string) {
	o.UserAgent = &v
}

func (o GetDevices200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDevices200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Expiry) {
		toSerialize["expiry"] = o.Expiry
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LastAuthenticated) {
		toSerialize["last_authenticated"] = o.LastAuthenticated
	}
	if !IsNil(o.LastRemoteAddr) {
		toSerialize["last_remote_addr"] = o.LastRemoteAddr
	}
	if !IsNil(o.UserAgent) {
		toSerialize["user_agent"] = o.UserAgent
	}
	return toSerialize, nil
}

type NullableGetDevices200ResponseDataInner struct {
	value *GetDevices200ResponseDataInner
	isSet bool
}

func (v NullableGetDevices200ResponseDataInner) Get() *GetDevices200ResponseDataInner {
	return v.value
}

func (v *NullableGetDevices200ResponseDataInner) Set(val *GetDevices200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDevices200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDevices200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDevices200ResponseDataInner(val *GetDevices200ResponseDataInner) *NullableGetDevices200ResponseDataInner {
	return &NullableGetDevices200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetDevices200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDevices200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


