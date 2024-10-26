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

// checks if the PostNodeBalancerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostNodeBalancerRequest{}

// PostNodeBalancerRequest struct for PostNodeBalancerRequest
type PostNodeBalancerRequest struct {
	// Throttle connections per second.  Set to 0 (zero) to disable throttling.
	ClientConnThrottle *int32 `json:"client_conn_throttle,omitempty"`
	// The port Configs to create for this NodeBalancer.  Each Config must have a unique port and at least one Node.
	Configs []PostNodeBalancerRequestConfigsInner `json:"configs,omitempty"`
	// The ID of the Firewall to assign to the NodeBalancer.  - A NodeBalancer can have only one Firewall assigned to it. - Firewalls only apply to inbound TCP traffic to NodeBalancers.
	FirewallId *int32 `json:"firewall_id,omitempty"`
	// This NodeBalancer's label. These must be unique on your Account.
	Label *string `json:"label,omitempty" validate:"regexp=[a-zA-Z0-9-_]{3,32}"`
	// The ID of the Region to create this NodeBalancer in.
	Region string `json:"region"`
	// An array of Tags applied to this object. Tags are for organizational purposes only.
	Tags []string `json:"tags,omitempty"`
}

type _PostNodeBalancerRequest PostNodeBalancerRequest

// NewPostNodeBalancerRequest instantiates a new PostNodeBalancerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostNodeBalancerRequest(region string) *PostNodeBalancerRequest {
	this := PostNodeBalancerRequest{}
	this.Region = region
	return &this
}

// NewPostNodeBalancerRequestWithDefaults instantiates a new PostNodeBalancerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostNodeBalancerRequestWithDefaults() *PostNodeBalancerRequest {
	this := PostNodeBalancerRequest{}
	return &this
}

// GetClientConnThrottle returns the ClientConnThrottle field value if set, zero value otherwise.
func (o *PostNodeBalancerRequest) GetClientConnThrottle() int32 {
	if o == nil || IsNil(o.ClientConnThrottle) {
		var ret int32
		return ret
	}
	return *o.ClientConnThrottle
}

// GetClientConnThrottleOk returns a tuple with the ClientConnThrottle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostNodeBalancerRequest) GetClientConnThrottleOk() (*int32, bool) {
	if o == nil || IsNil(o.ClientConnThrottle) {
		return nil, false
	}
	return o.ClientConnThrottle, true
}

// HasClientConnThrottle returns a boolean if a field has been set.
func (o *PostNodeBalancerRequest) HasClientConnThrottle() bool {
	if o != nil && !IsNil(o.ClientConnThrottle) {
		return true
	}

	return false
}

// SetClientConnThrottle gets a reference to the given int32 and assigns it to the ClientConnThrottle field.
func (o *PostNodeBalancerRequest) SetClientConnThrottle(v int32) {
	o.ClientConnThrottle = &v
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *PostNodeBalancerRequest) GetConfigs() []PostNodeBalancerRequestConfigsInner {
	if o == nil || IsNil(o.Configs) {
		var ret []PostNodeBalancerRequestConfigsInner
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostNodeBalancerRequest) GetConfigsOk() ([]PostNodeBalancerRequestConfigsInner, bool) {
	if o == nil || IsNil(o.Configs) {
		return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *PostNodeBalancerRequest) HasConfigs() bool {
	if o != nil && !IsNil(o.Configs) {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given []PostNodeBalancerRequestConfigsInner and assigns it to the Configs field.
func (o *PostNodeBalancerRequest) SetConfigs(v []PostNodeBalancerRequestConfigsInner) {
	o.Configs = v
}

// GetFirewallId returns the FirewallId field value if set, zero value otherwise.
func (o *PostNodeBalancerRequest) GetFirewallId() int32 {
	if o == nil || IsNil(o.FirewallId) {
		var ret int32
		return ret
	}
	return *o.FirewallId
}

// GetFirewallIdOk returns a tuple with the FirewallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostNodeBalancerRequest) GetFirewallIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FirewallId) {
		return nil, false
	}
	return o.FirewallId, true
}

// HasFirewallId returns a boolean if a field has been set.
func (o *PostNodeBalancerRequest) HasFirewallId() bool {
	if o != nil && !IsNil(o.FirewallId) {
		return true
	}

	return false
}

// SetFirewallId gets a reference to the given int32 and assigns it to the FirewallId field.
func (o *PostNodeBalancerRequest) SetFirewallId(v int32) {
	o.FirewallId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PostNodeBalancerRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostNodeBalancerRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PostNodeBalancerRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PostNodeBalancerRequest) SetLabel(v string) {
	o.Label = &v
}

// GetRegion returns the Region field value
func (o *PostNodeBalancerRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *PostNodeBalancerRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *PostNodeBalancerRequest) SetRegion(v string) {
	o.Region = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PostNodeBalancerRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostNodeBalancerRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PostNodeBalancerRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PostNodeBalancerRequest) SetTags(v []string) {
	o.Tags = v
}

func (o PostNodeBalancerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostNodeBalancerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientConnThrottle) {
		toSerialize["client_conn_throttle"] = o.ClientConnThrottle
	}
	if !IsNil(o.Configs) {
		toSerialize["configs"] = o.Configs
	}
	if !IsNil(o.FirewallId) {
		toSerialize["firewall_id"] = o.FirewallId
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	toSerialize["region"] = o.Region
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

func (o *PostNodeBalancerRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
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

	varPostNodeBalancerRequest := _PostNodeBalancerRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostNodeBalancerRequest)

	if err != nil {
		return err
	}

	*o = PostNodeBalancerRequest(varPostNodeBalancerRequest)

	return err
}

type NullablePostNodeBalancerRequest struct {
	value *PostNodeBalancerRequest
	isSet bool
}

func (v NullablePostNodeBalancerRequest) Get() *PostNodeBalancerRequest {
	return v.value
}

func (v *NullablePostNodeBalancerRequest) Set(val *PostNodeBalancerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostNodeBalancerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostNodeBalancerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostNodeBalancerRequest(val *PostNodeBalancerRequest) *NullablePostNodeBalancerRequest {
	return &NullablePostNodeBalancerRequest{value: val, isSet: true}
}

func (v NullablePostNodeBalancerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostNodeBalancerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

