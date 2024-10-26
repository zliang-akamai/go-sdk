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

// checks if the GetLkeClusterKubeconfig200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLkeClusterKubeconfig200Response{}

// GetLkeClusterKubeconfig200Response struct for GetLkeClusterKubeconfig200Response
type GetLkeClusterKubeconfig200Response struct {
	// The Base64-encoded Kubeconfig file for this Cluster.
	Kubeconfig *string `json:"kubeconfig,omitempty"`
}

// NewGetLkeClusterKubeconfig200Response instantiates a new GetLkeClusterKubeconfig200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLkeClusterKubeconfig200Response() *GetLkeClusterKubeconfig200Response {
	this := GetLkeClusterKubeconfig200Response{}
	return &this
}

// NewGetLkeClusterKubeconfig200ResponseWithDefaults instantiates a new GetLkeClusterKubeconfig200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLkeClusterKubeconfig200ResponseWithDefaults() *GetLkeClusterKubeconfig200Response {
	this := GetLkeClusterKubeconfig200Response{}
	return &this
}

// GetKubeconfig returns the Kubeconfig field value if set, zero value otherwise.
func (o *GetLkeClusterKubeconfig200Response) GetKubeconfig() string {
	if o == nil || IsNil(o.Kubeconfig) {
		var ret string
		return ret
	}
	return *o.Kubeconfig
}

// GetKubeconfigOk returns a tuple with the Kubeconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLkeClusterKubeconfig200Response) GetKubeconfigOk() (*string, bool) {
	if o == nil || IsNil(o.Kubeconfig) {
		return nil, false
	}
	return o.Kubeconfig, true
}

// HasKubeconfig returns a boolean if a field has been set.
func (o *GetLkeClusterKubeconfig200Response) HasKubeconfig() bool {
	if o != nil && !IsNil(o.Kubeconfig) {
		return true
	}

	return false
}

// SetKubeconfig gets a reference to the given string and assigns it to the Kubeconfig field.
func (o *GetLkeClusterKubeconfig200Response) SetKubeconfig(v string) {
	o.Kubeconfig = &v
}

func (o GetLkeClusterKubeconfig200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLkeClusterKubeconfig200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kubeconfig) {
		toSerialize["kubeconfig"] = o.Kubeconfig
	}
	return toSerialize, nil
}

type NullableGetLkeClusterKubeconfig200Response struct {
	value *GetLkeClusterKubeconfig200Response
	isSet bool
}

func (v NullableGetLkeClusterKubeconfig200Response) Get() *GetLkeClusterKubeconfig200Response {
	return v.value
}

func (v *NullableGetLkeClusterKubeconfig200Response) Set(val *GetLkeClusterKubeconfig200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLkeClusterKubeconfig200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLkeClusterKubeconfig200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLkeClusterKubeconfig200Response(val *GetLkeClusterKubeconfig200Response) *NullableGetLkeClusterKubeconfig200Response {
	return &NullableGetLkeClusterKubeconfig200Response{value: val, isSet: true}
}

func (v NullableGetLkeClusterKubeconfig200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLkeClusterKubeconfig200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

