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

// checks if the PostObjectStorageSslRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostObjectStorageSslRequest{}

// PostObjectStorageSslRequest Upload a TLS/SSL certificate and private key to be served when you visit your Object Storage bucket via HTTPS.
type PostObjectStorageSslRequest struct {
	// Your Base64 encoded and PEM formatted SSL certificate.  Line breaks must be represented as `\\n` in the string for requests (but not when using the Linode CLI)
	Certificate string `json:"certificate"`
	// The private key associated with this TLS/SSL certificate.  Line breaks must be represented as `\\n` in the string for requests (but not when using the Linode CLI)
	PrivateKey string `json:"private_key"`
}

type _PostObjectStorageSslRequest PostObjectStorageSslRequest

// NewPostObjectStorageSslRequest instantiates a new PostObjectStorageSslRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostObjectStorageSslRequest(certificate string, privateKey string) *PostObjectStorageSslRequest {
	this := PostObjectStorageSslRequest{}
	this.Certificate = certificate
	this.PrivateKey = privateKey
	return &this
}

// NewPostObjectStorageSslRequestWithDefaults instantiates a new PostObjectStorageSslRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostObjectStorageSslRequestWithDefaults() *PostObjectStorageSslRequest {
	this := PostObjectStorageSslRequest{}
	return &this
}

// GetCertificate returns the Certificate field value
func (o *PostObjectStorageSslRequest) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *PostObjectStorageSslRequest) GetCertificateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *PostObjectStorageSslRequest) SetCertificate(v string) {
	o.Certificate = v
}

// GetPrivateKey returns the PrivateKey field value
func (o *PostObjectStorageSslRequest) GetPrivateKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrivateKey
}

// GetPrivateKeyOk returns a tuple with the PrivateKey field value
// and a boolean to check if the value has been set.
func (o *PostObjectStorageSslRequest) GetPrivateKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrivateKey, true
}

// SetPrivateKey sets field value
func (o *PostObjectStorageSslRequest) SetPrivateKey(v string) {
	o.PrivateKey = v
}

func (o PostObjectStorageSslRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostObjectStorageSslRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["certificate"] = o.Certificate
	toSerialize["private_key"] = o.PrivateKey
	return toSerialize, nil
}

func (o *PostObjectStorageSslRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"certificate",
		"private_key",
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

	varPostObjectStorageSslRequest := _PostObjectStorageSslRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostObjectStorageSslRequest)

	if err != nil {
		return err
	}

	*o = PostObjectStorageSslRequest(varPostObjectStorageSslRequest)

	return err
}

type NullablePostObjectStorageSslRequest struct {
	value *PostObjectStorageSslRequest
	isSet bool
}

func (v NullablePostObjectStorageSslRequest) Get() *PostObjectStorageSslRequest {
	return v.value
}

func (v *NullablePostObjectStorageSslRequest) Set(val *PostObjectStorageSslRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostObjectStorageSslRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostObjectStorageSslRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostObjectStorageSslRequest(val *PostObjectStorageSslRequest) *NullablePostObjectStorageSslRequest {
	return &NullablePostObjectStorageSslRequest{value: val, isSet: true}
}

func (v NullablePostObjectStorageSslRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostObjectStorageSslRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

