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

// checks if the PutDomainRecordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutDomainRecordRequest{}

// PutDomainRecordRequest A Domain Record Update request object.
type PutDomainRecordRequest struct {
	// The name of this Record. For requests, this property's actual usage and whether it is required depends on the type of record this represents:  `A` and `AAAA`: The hostname or FQDN of the Record.  `NS`: The subdomain, if any, to use with the Domain of the Record. Wildcard NS records (`*`) are not supported.  `MX`: The mail subdomain. For example, `sub` for the address `user@sub.example.com` under the `example.com` Domain.  - The left-most subdomain component may be an asterisk (`*`) to designate a wildcard subdomain. - Other subdomain components must only contain letters, digits, and hyphens, start with a letter, end with a letter or digit, and contain less than 64 characters. - Must be an empty string (`\"\"`) for a Null MX Record.  `CNAME`: The hostname. Must be unique. Required.  `TXT`: The hostname.  `SRV`: Unused. Use the `service` property to set the service name for this record.  `CAA`: The subdomain. Omit or enter an empty string (`\"\"`) to apply to the entire Domain.  `PTR`: See our guide on how to [Configure Your Linode for Reverse DNS (rDNS)](https://www.linode.com/docs/guides/configure-rdns/).
	Name *string `json:"name,omitempty"`
	// The port this Record points to. Only valid and required for SRV record requests.
	Port *int32 `json:"port,omitempty"`
	// The priority of the target host for this Record. Lower values are preferred. Only valid for MX and SRV record requests. Required for SRV record requests.  Defaults to `0` for MX record requests. Must be `0` for Null MX records.
	Priority *int32 `json:"priority,omitempty"`
	// The protocol this Record's service communicates with. An underscore (`_`) is prepended automatically to the submitted value for this property. Only valid for SRV record requests.
	Protocol NullableString `json:"protocol,omitempty"`
	// The name of the service. An underscore (`_`) is prepended and a period (`.`) is appended automatically to the submitted value for this property. Only valid and required for SRV record requests.
	Service NullableString `json:"service,omitempty"`
	// The tag portion of a CAA record. Only valid and required for CAA record requests.
	Tag NullableString `json:"tag,omitempty"`
	// The target for this Record. For requests, this property's actual usage and whether it is required depends on the type of record this represents:  `A` and `AAAA`: The IP address. Use `[remote_addr]` to submit the IPv4 address of the request. Required.  `NS`: The name server. Must be a valid domain. Required.  `MX`: The mail server. Must be a valid domain unless creating a Null MX Record. Required.  - Must have less than 254 total characters. - The left-most domain component may be an asterisk (`*`) to designate a wildcard domain. - Other domain components must only contain letters, digits, and hyphens, start with a letter, end with a letter or digit, and contain less than 64 characters. - To create a [Null MX Record](https://datatracker.ietf.org/doc/html/rfc7505), first [remove](https://techdocs.akamai.com/linode-api/reference/delete-domain-record) any additional MX records, then create an MX record with empty strings (`\"\"`) for the `target` and `name`. If a Domain has a Null MX record, new MX records cannot be created.  `CNAME`: The alias. Must be a valid domain. Required.  `TXT`: The value. Required.  `SRV`: The target domain or subdomain. If a subdomain is entered, it is automatically used with the Domain. To configure for a different domain, enter a valid FQDN. For example, the value `www` with a Domain for `example.com` results in a target set to `www.example.com`, whereas the value `sample.com` results in a target set to `sample.com`. Required.  `CAA`: The value. For `issue` or `issuewild` tags, the domain of your certificate issuer. For the `iodef` tag, a contact or submission URL (domain, http, https, or mailto). Requirements depend on the tag for this record:    - `issue`: The domain of your certificate issuer. Must include a valid domain. May include additional parameters separated with semicolons (`;`), for example: `www.example.com; foo=bar`   - `issuewild`: The domain of your wildcard certificate issuer. Must be a valid domain and must not start with an asterisk (`*`).   - `iodef`: Must be either (1) a valid domain, (2) a valid domain prepended with `http://` or `https://`, or (3) a valid email address prepended with `mailto:`.  `PTR`: Required. See our guide on how to [Configure Your Linode for Reverse DNS (rDNS)](https://www.linode.com/docs/guides/configure-rdns/).  With the exception of A, AAAA, and CAA records, this field accepts a trailing period.
	Target *string `json:"target,omitempty"`
	// \"Time to Live\" - the amount of time in seconds that this Domain's records may be cached by resolvers or other domain servers. Valid values are 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200 - any other value will be rounded to the nearest valid value.
	TtlSec *int32 `json:"ttl_sec,omitempty"`
	// The relative weight of this Record used in the case of identical priority. Higher values are preferred. Only valid and required for SRV record requests.
	Weight *int32 `json:"weight,omitempty"`
}

// NewPutDomainRecordRequest instantiates a new PutDomainRecordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutDomainRecordRequest() *PutDomainRecordRequest {
	this := PutDomainRecordRequest{}
	return &this
}

// NewPutDomainRecordRequestWithDefaults instantiates a new PutDomainRecordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutDomainRecordRequestWithDefaults() *PutDomainRecordRequest {
	this := PutDomainRecordRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PutDomainRecordRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutDomainRecordRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PutDomainRecordRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PutDomainRecordRequest) SetName(v string) {
	o.Name = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *PutDomainRecordRequest) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutDomainRecordRequest) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *PutDomainRecordRequest) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *PutDomainRecordRequest) SetPort(v int32) {
	o.Port = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *PutDomainRecordRequest) GetPriority() int32 {
	if o == nil || IsNil(o.Priority) {
		var ret int32
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutDomainRecordRequest) GetPriorityOk() (*int32, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *PutDomainRecordRequest) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given int32 and assigns it to the Priority field.
func (o *PutDomainRecordRequest) SetPriority(v int32) {
	o.Priority = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutDomainRecordRequest) GetProtocol() string {
	if o == nil || IsNil(o.Protocol.Get()) {
		var ret string
		return ret
	}
	return *o.Protocol.Get()
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutDomainRecordRequest) GetProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Protocol.Get(), o.Protocol.IsSet()
}

// HasProtocol returns a boolean if a field has been set.
func (o *PutDomainRecordRequest) HasProtocol() bool {
	if o != nil && o.Protocol.IsSet() {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given NullableString and assigns it to the Protocol field.
func (o *PutDomainRecordRequest) SetProtocol(v string) {
	o.Protocol.Set(&v)
}
// SetProtocolNil sets the value for Protocol to be an explicit nil
func (o *PutDomainRecordRequest) SetProtocolNil() {
	o.Protocol.Set(nil)
}

// UnsetProtocol ensures that no value is present for Protocol, not even an explicit nil
func (o *PutDomainRecordRequest) UnsetProtocol() {
	o.Protocol.Unset()
}

// GetService returns the Service field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutDomainRecordRequest) GetService() string {
	if o == nil || IsNil(o.Service.Get()) {
		var ret string
		return ret
	}
	return *o.Service.Get()
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutDomainRecordRequest) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Service.Get(), o.Service.IsSet()
}

// HasService returns a boolean if a field has been set.
func (o *PutDomainRecordRequest) HasService() bool {
	if o != nil && o.Service.IsSet() {
		return true
	}

	return false
}

// SetService gets a reference to the given NullableString and assigns it to the Service field.
func (o *PutDomainRecordRequest) SetService(v string) {
	o.Service.Set(&v)
}
// SetServiceNil sets the value for Service to be an explicit nil
func (o *PutDomainRecordRequest) SetServiceNil() {
	o.Service.Set(nil)
}

// UnsetService ensures that no value is present for Service, not even an explicit nil
func (o *PutDomainRecordRequest) UnsetService() {
	o.Service.Unset()
}

// GetTag returns the Tag field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PutDomainRecordRequest) GetTag() string {
	if o == nil || IsNil(o.Tag.Get()) {
		var ret string
		return ret
	}
	return *o.Tag.Get()
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PutDomainRecordRequest) GetTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tag.Get(), o.Tag.IsSet()
}

// HasTag returns a boolean if a field has been set.
func (o *PutDomainRecordRequest) HasTag() bool {
	if o != nil && o.Tag.IsSet() {
		return true
	}

	return false
}

// SetTag gets a reference to the given NullableString and assigns it to the Tag field.
func (o *PutDomainRecordRequest) SetTag(v string) {
	o.Tag.Set(&v)
}
// SetTagNil sets the value for Tag to be an explicit nil
func (o *PutDomainRecordRequest) SetTagNil() {
	o.Tag.Set(nil)
}

// UnsetTag ensures that no value is present for Tag, not even an explicit nil
func (o *PutDomainRecordRequest) UnsetTag() {
	o.Tag.Unset()
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *PutDomainRecordRequest) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutDomainRecordRequest) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *PutDomainRecordRequest) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *PutDomainRecordRequest) SetTarget(v string) {
	o.Target = &v
}

// GetTtlSec returns the TtlSec field value if set, zero value otherwise.
func (o *PutDomainRecordRequest) GetTtlSec() int32 {
	if o == nil || IsNil(o.TtlSec) {
		var ret int32
		return ret
	}
	return *o.TtlSec
}

// GetTtlSecOk returns a tuple with the TtlSec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutDomainRecordRequest) GetTtlSecOk() (*int32, bool) {
	if o == nil || IsNil(o.TtlSec) {
		return nil, false
	}
	return o.TtlSec, true
}

// HasTtlSec returns a boolean if a field has been set.
func (o *PutDomainRecordRequest) HasTtlSec() bool {
	if o != nil && !IsNil(o.TtlSec) {
		return true
	}

	return false
}

// SetTtlSec gets a reference to the given int32 and assigns it to the TtlSec field.
func (o *PutDomainRecordRequest) SetTtlSec(v int32) {
	o.TtlSec = &v
}

// GetWeight returns the Weight field value if set, zero value otherwise.
func (o *PutDomainRecordRequest) GetWeight() int32 {
	if o == nil || IsNil(o.Weight) {
		var ret int32
		return ret
	}
	return *o.Weight
}

// GetWeightOk returns a tuple with the Weight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutDomainRecordRequest) GetWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Weight) {
		return nil, false
	}
	return o.Weight, true
}

// HasWeight returns a boolean if a field has been set.
func (o *PutDomainRecordRequest) HasWeight() bool {
	if o != nil && !IsNil(o.Weight) {
		return true
	}

	return false
}

// SetWeight gets a reference to the given int32 and assigns it to the Weight field.
func (o *PutDomainRecordRequest) SetWeight(v int32) {
	o.Weight = &v
}

func (o PutDomainRecordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutDomainRecordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if o.Protocol.IsSet() {
		toSerialize["protocol"] = o.Protocol.Get()
	}
	if o.Service.IsSet() {
		toSerialize["service"] = o.Service.Get()
	}
	if o.Tag.IsSet() {
		toSerialize["tag"] = o.Tag.Get()
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	if !IsNil(o.TtlSec) {
		toSerialize["ttl_sec"] = o.TtlSec
	}
	if !IsNil(o.Weight) {
		toSerialize["weight"] = o.Weight
	}
	return toSerialize, nil
}

type NullablePutDomainRecordRequest struct {
	value *PutDomainRecordRequest
	isSet bool
}

func (v NullablePutDomainRecordRequest) Get() *PutDomainRecordRequest {
	return v.value
}

func (v *NullablePutDomainRecordRequest) Set(val *PutDomainRecordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutDomainRecordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutDomainRecordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutDomainRecordRequest(val *PutDomainRecordRequest) *NullablePutDomainRecordRequest {
	return &NullablePutDomainRecordRequest{value: val, isSet: true}
}

func (v NullablePutDomainRecordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutDomainRecordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

