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

// checks if the PutLkeClusterRequestControlPlane type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutLkeClusterRequestControlPlane{}

// PutLkeClusterRequestControlPlane Defines settings for the Kubernetes Control Plane. Allows for the enabling of High Availability (HA) and IP-based Access Control List (ACL) for Control Plane Components. Enabling of either of these for LKE is an __irreversible__ change.  When upgrading pre-existing LKE clusters to use the HA Control Plane, the following changes will additionally occur:  - All nodes will be deleted and new nodes will be created to replace them.  - Any local storage (such as `hostPath` volumes) will be erased.  - The upgrade process may take several minutes to complete, as nodes will be replaced on a rolling basis.  When upgrading pre-existing LKE clusters to use the ACL Control Plane, the following changes will additionally occur:  - All control plane access will go through a Cloud Firewall. There will be a period on which the FQDN DNS record needs to be propagated. Due to TTL and DNS caching, it could take several hours for external clients to switch over to the new mappings.
type PutLkeClusterRequestControlPlane struct {
	Acl *GetLkeClusters200ResponseDataInnerControlPlaneAcl `json:"acl,omitempty"`
	// Enables High Availability for the Control Plane Components of the cluster. Defaults to `false`. Enabling High Availability for LKE is an __irreversible__ change.
	HighAvailability *bool `json:"high_availability,omitempty"`
}

// NewPutLkeClusterRequestControlPlane instantiates a new PutLkeClusterRequestControlPlane object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutLkeClusterRequestControlPlane() *PutLkeClusterRequestControlPlane {
	this := PutLkeClusterRequestControlPlane{}
	var highAvailability bool = false
	this.HighAvailability = &highAvailability
	return &this
}

// NewPutLkeClusterRequestControlPlaneWithDefaults instantiates a new PutLkeClusterRequestControlPlane object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutLkeClusterRequestControlPlaneWithDefaults() *PutLkeClusterRequestControlPlane {
	this := PutLkeClusterRequestControlPlane{}
	var highAvailability bool = false
	this.HighAvailability = &highAvailability
	return &this
}

// GetAcl returns the Acl field value if set, zero value otherwise.
func (o *PutLkeClusterRequestControlPlane) GetAcl() GetLkeClusters200ResponseDataInnerControlPlaneAcl {
	if o == nil || IsNil(o.Acl) {
		var ret GetLkeClusters200ResponseDataInnerControlPlaneAcl
		return ret
	}
	return *o.Acl
}

// GetAclOk returns a tuple with the Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutLkeClusterRequestControlPlane) GetAclOk() (*GetLkeClusters200ResponseDataInnerControlPlaneAcl, bool) {
	if o == nil || IsNil(o.Acl) {
		return nil, false
	}
	return o.Acl, true
}

// HasAcl returns a boolean if a field has been set.
func (o *PutLkeClusterRequestControlPlane) HasAcl() bool {
	if o != nil && !IsNil(o.Acl) {
		return true
	}

	return false
}

// SetAcl gets a reference to the given GetLkeClusters200ResponseDataInnerControlPlaneAcl and assigns it to the Acl field.
func (o *PutLkeClusterRequestControlPlane) SetAcl(v GetLkeClusters200ResponseDataInnerControlPlaneAcl) {
	o.Acl = &v
}

// GetHighAvailability returns the HighAvailability field value if set, zero value otherwise.
func (o *PutLkeClusterRequestControlPlane) GetHighAvailability() bool {
	if o == nil || IsNil(o.HighAvailability) {
		var ret bool
		return ret
	}
	return *o.HighAvailability
}

// GetHighAvailabilityOk returns a tuple with the HighAvailability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutLkeClusterRequestControlPlane) GetHighAvailabilityOk() (*bool, bool) {
	if o == nil || IsNil(o.HighAvailability) {
		return nil, false
	}
	return o.HighAvailability, true
}

// HasHighAvailability returns a boolean if a field has been set.
func (o *PutLkeClusterRequestControlPlane) HasHighAvailability() bool {
	if o != nil && !IsNil(o.HighAvailability) {
		return true
	}

	return false
}

// SetHighAvailability gets a reference to the given bool and assigns it to the HighAvailability field.
func (o *PutLkeClusterRequestControlPlane) SetHighAvailability(v bool) {
	o.HighAvailability = &v
}

func (o PutLkeClusterRequestControlPlane) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutLkeClusterRequestControlPlane) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Acl) {
		toSerialize["acl"] = o.Acl
	}
	if !IsNil(o.HighAvailability) {
		toSerialize["high_availability"] = o.HighAvailability
	}
	return toSerialize, nil
}

type NullablePutLkeClusterRequestControlPlane struct {
	value *PutLkeClusterRequestControlPlane
	isSet bool
}

func (v NullablePutLkeClusterRequestControlPlane) Get() *PutLkeClusterRequestControlPlane {
	return v.value
}

func (v *NullablePutLkeClusterRequestControlPlane) Set(val *PutLkeClusterRequestControlPlane) {
	v.value = val
	v.isSet = true
}

func (v NullablePutLkeClusterRequestControlPlane) IsSet() bool {
	return v.isSet
}

func (v *NullablePutLkeClusterRequestControlPlane) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutLkeClusterRequestControlPlane(val *PutLkeClusterRequestControlPlane) *NullablePutLkeClusterRequestControlPlane {
	return &NullablePutLkeClusterRequestControlPlane{value: val, isSet: true}
}

func (v NullablePutLkeClusterRequestControlPlane) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutLkeClusterRequestControlPlane) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

