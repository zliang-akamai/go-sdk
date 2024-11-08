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

// checks if the PostImportDomainRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostImportDomainRequest{}

// PostImportDomainRequest struct for PostImportDomainRequest
type PostImportDomainRequest struct {
	// The domain to import.
	Domain string `json:"domain"`
	// The remote nameserver that allows zone transfers (AXFR).
	RemoteNameserver string `json:"remote_nameserver"`
}

type _PostImportDomainRequest PostImportDomainRequest

// NewPostImportDomainRequest instantiates a new PostImportDomainRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostImportDomainRequest(domain string, remoteNameserver string) *PostImportDomainRequest {
	this := PostImportDomainRequest{}
	this.Domain = domain
	this.RemoteNameserver = remoteNameserver
	return &this
}

// NewPostImportDomainRequestWithDefaults instantiates a new PostImportDomainRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostImportDomainRequestWithDefaults() *PostImportDomainRequest {
	this := PostImportDomainRequest{}
	return &this
}

// GetDomain returns the Domain field value
func (o *PostImportDomainRequest) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *PostImportDomainRequest) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *PostImportDomainRequest) SetDomain(v string) {
	o.Domain = v
}

// GetRemoteNameserver returns the RemoteNameserver field value
func (o *PostImportDomainRequest) GetRemoteNameserver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteNameserver
}

// GetRemoteNameserverOk returns a tuple with the RemoteNameserver field value
// and a boolean to check if the value has been set.
func (o *PostImportDomainRequest) GetRemoteNameserverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteNameserver, true
}

// SetRemoteNameserver sets field value
func (o *PostImportDomainRequest) SetRemoteNameserver(v string) {
	o.RemoteNameserver = v
}

func (o PostImportDomainRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostImportDomainRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["domain"] = o.Domain
	toSerialize["remote_nameserver"] = o.RemoteNameserver
	return toSerialize, nil
}

func (o *PostImportDomainRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"domain",
		"remote_nameserver",
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

	varPostImportDomainRequest := _PostImportDomainRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostImportDomainRequest)

	if err != nil {
		return err
	}

	*o = PostImportDomainRequest(varPostImportDomainRequest)

	return err
}

type NullablePostImportDomainRequest struct {
	value *PostImportDomainRequest
	isSet bool
}

func (v NullablePostImportDomainRequest) Get() *PostImportDomainRequest {
	return v.value
}

func (v *NullablePostImportDomainRequest) Set(val *PostImportDomainRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostImportDomainRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostImportDomainRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostImportDomainRequest(val *PostImportDomainRequest) *NullablePostImportDomainRequest {
	return &NullablePostImportDomainRequest{value: val, isSet: true}
}

func (v NullablePostImportDomainRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostImportDomainRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


