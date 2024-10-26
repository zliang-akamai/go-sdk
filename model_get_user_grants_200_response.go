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

// checks if the GetUserGrants200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserGrants200Response{}

// GetUserGrants200Response A structure representing all grants a restricted User has on the Account. Not available for unrestricted users, as they have access to everything without grants. If retrieved from the `/profile/grants` endpoint, entities to which a User has no access will be omitted.
type GetUserGrants200Response struct {
	// The grants this User has for each Database that is owned by this Account.
	Database []GetUserGrants200ResponseDatabaseInner `json:"database,omitempty"`
	// The grants this User has for each Domain that is owned by this Account.
	Domain []GetUserGrants200ResponseDatabaseInner `json:"domain,omitempty"`
	// The grants this User has for each Firewall that is owned by this Account.
	Firewall []GetUserGrants200ResponseDatabaseInner `json:"firewall,omitempty"`
	Global *GetUserGrants200ResponseGlobal `json:"global,omitempty"`
	// The grants this User has for each Image that is owned by this Account.
	Image []GetUserGrants200ResponseDatabaseInner `json:"image,omitempty"`
	// The grants this User has for each Linode that is owned by this Account.
	Linode []GetUserGrants200ResponseDatabaseInner `json:"linode,omitempty"`
	// The grants this User has for each Longview Client that is owned by this Account.
	Longview []GetUserGrants200ResponseDatabaseInner `json:"longview,omitempty"`
	// The grants this User has for each NodeBalancer that is owned by this Account.
	Nodebalancer []GetUserGrants200ResponseDatabaseInner `json:"nodebalancer,omitempty"`
	// The grants this User has for each Placement Group that is owned by this Account.
	PlacementGroup []GetUserGrants200ResponseDatabaseInner `json:"placement_group,omitempty"`
	// The grants this User has for each StackScript that is owned by this Account.
	Stackscript []GetUserGrants200ResponseDatabaseInner `json:"stackscript,omitempty"`
	// The grants this User has for each Block Storage Volume that is owned by this Account.
	Volume []GetUserGrants200ResponseDatabaseInner `json:"volume,omitempty"`
	// The grants this User has for each VPC that is owned by this Account.
	Vpc []GetUserGrants200ResponseDatabaseInner `json:"vpc,omitempty"`
}

// NewGetUserGrants200Response instantiates a new GetUserGrants200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserGrants200Response() *GetUserGrants200Response {
	this := GetUserGrants200Response{}
	return &this
}

// NewGetUserGrants200ResponseWithDefaults instantiates a new GetUserGrants200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserGrants200ResponseWithDefaults() *GetUserGrants200Response {
	this := GetUserGrants200Response{}
	return &this
}

// GetDatabase returns the Database field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetDatabase() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Database) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Database
}

// GetDatabaseOk returns a tuple with the Database field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetDatabaseOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Database) {
		return nil, false
	}
	return o.Database, true
}

// HasDatabase returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasDatabase() bool {
	if o != nil && !IsNil(o.Database) {
		return true
	}

	return false
}

// SetDatabase gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Database field.
func (o *GetUserGrants200Response) SetDatabase(v []GetUserGrants200ResponseDatabaseInner) {
	o.Database = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetDomain() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Domain) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetDomainOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Domain field.
func (o *GetUserGrants200Response) SetDomain(v []GetUserGrants200ResponseDatabaseInner) {
	o.Domain = v
}

// GetFirewall returns the Firewall field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetFirewall() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Firewall) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Firewall
}

// GetFirewallOk returns a tuple with the Firewall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetFirewallOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Firewall) {
		return nil, false
	}
	return o.Firewall, true
}

// HasFirewall returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasFirewall() bool {
	if o != nil && !IsNil(o.Firewall) {
		return true
	}

	return false
}

// SetFirewall gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Firewall field.
func (o *GetUserGrants200Response) SetFirewall(v []GetUserGrants200ResponseDatabaseInner) {
	o.Firewall = v
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetGlobal() GetUserGrants200ResponseGlobal {
	if o == nil || IsNil(o.Global) {
		var ret GetUserGrants200ResponseGlobal
		return ret
	}
	return *o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetGlobalOk() (*GetUserGrants200ResponseGlobal, bool) {
	if o == nil || IsNil(o.Global) {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasGlobal() bool {
	if o != nil && !IsNil(o.Global) {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given GetUserGrants200ResponseGlobal and assigns it to the Global field.
func (o *GetUserGrants200Response) SetGlobal(v GetUserGrants200ResponseGlobal) {
	o.Global = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetImage() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Image) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetImageOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Image field.
func (o *GetUserGrants200Response) SetImage(v []GetUserGrants200ResponseDatabaseInner) {
	o.Image = v
}

// GetLinode returns the Linode field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetLinode() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Linode) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Linode
}

// GetLinodeOk returns a tuple with the Linode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetLinodeOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Linode) {
		return nil, false
	}
	return o.Linode, true
}

// HasLinode returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasLinode() bool {
	if o != nil && !IsNil(o.Linode) {
		return true
	}

	return false
}

// SetLinode gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Linode field.
func (o *GetUserGrants200Response) SetLinode(v []GetUserGrants200ResponseDatabaseInner) {
	o.Linode = v
}

// GetLongview returns the Longview field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetLongview() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Longview) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Longview
}

// GetLongviewOk returns a tuple with the Longview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetLongviewOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Longview) {
		return nil, false
	}
	return o.Longview, true
}

// HasLongview returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasLongview() bool {
	if o != nil && !IsNil(o.Longview) {
		return true
	}

	return false
}

// SetLongview gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Longview field.
func (o *GetUserGrants200Response) SetLongview(v []GetUserGrants200ResponseDatabaseInner) {
	o.Longview = v
}

// GetNodebalancer returns the Nodebalancer field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetNodebalancer() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Nodebalancer) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Nodebalancer
}

// GetNodebalancerOk returns a tuple with the Nodebalancer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetNodebalancerOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Nodebalancer) {
		return nil, false
	}
	return o.Nodebalancer, true
}

// HasNodebalancer returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasNodebalancer() bool {
	if o != nil && !IsNil(o.Nodebalancer) {
		return true
	}

	return false
}

// SetNodebalancer gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Nodebalancer field.
func (o *GetUserGrants200Response) SetNodebalancer(v []GetUserGrants200ResponseDatabaseInner) {
	o.Nodebalancer = v
}

// GetPlacementGroup returns the PlacementGroup field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetPlacementGroup() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.PlacementGroup) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.PlacementGroup
}

// GetPlacementGroupOk returns a tuple with the PlacementGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetPlacementGroupOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.PlacementGroup) {
		return nil, false
	}
	return o.PlacementGroup, true
}

// HasPlacementGroup returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasPlacementGroup() bool {
	if o != nil && !IsNil(o.PlacementGroup) {
		return true
	}

	return false
}

// SetPlacementGroup gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the PlacementGroup field.
func (o *GetUserGrants200Response) SetPlacementGroup(v []GetUserGrants200ResponseDatabaseInner) {
	o.PlacementGroup = v
}

// GetStackscript returns the Stackscript field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetStackscript() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Stackscript) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Stackscript
}

// GetStackscriptOk returns a tuple with the Stackscript field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetStackscriptOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Stackscript) {
		return nil, false
	}
	return o.Stackscript, true
}

// HasStackscript returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasStackscript() bool {
	if o != nil && !IsNil(o.Stackscript) {
		return true
	}

	return false
}

// SetStackscript gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Stackscript field.
func (o *GetUserGrants200Response) SetStackscript(v []GetUserGrants200ResponseDatabaseInner) {
	o.Stackscript = v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetVolume() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Volume) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetVolumeOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Volume field.
func (o *GetUserGrants200Response) SetVolume(v []GetUserGrants200ResponseDatabaseInner) {
	o.Volume = v
}

// GetVpc returns the Vpc field value if set, zero value otherwise.
func (o *GetUserGrants200Response) GetVpc() []GetUserGrants200ResponseDatabaseInner {
	if o == nil || IsNil(o.Vpc) {
		var ret []GetUserGrants200ResponseDatabaseInner
		return ret
	}
	return o.Vpc
}

// GetVpcOk returns a tuple with the Vpc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserGrants200Response) GetVpcOk() ([]GetUserGrants200ResponseDatabaseInner, bool) {
	if o == nil || IsNil(o.Vpc) {
		return nil, false
	}
	return o.Vpc, true
}

// HasVpc returns a boolean if a field has been set.
func (o *GetUserGrants200Response) HasVpc() bool {
	if o != nil && !IsNil(o.Vpc) {
		return true
	}

	return false
}

// SetVpc gets a reference to the given []GetUserGrants200ResponseDatabaseInner and assigns it to the Vpc field.
func (o *GetUserGrants200Response) SetVpc(v []GetUserGrants200ResponseDatabaseInner) {
	o.Vpc = v
}

func (o GetUserGrants200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserGrants200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Database) {
		toSerialize["database"] = o.Database
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Firewall) {
		toSerialize["firewall"] = o.Firewall
	}
	if !IsNil(o.Global) {
		toSerialize["global"] = o.Global
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Linode) {
		toSerialize["linode"] = o.Linode
	}
	if !IsNil(o.Longview) {
		toSerialize["longview"] = o.Longview
	}
	if !IsNil(o.Nodebalancer) {
		toSerialize["nodebalancer"] = o.Nodebalancer
	}
	if !IsNil(o.PlacementGroup) {
		toSerialize["placement_group"] = o.PlacementGroup
	}
	if !IsNil(o.Stackscript) {
		toSerialize["stackscript"] = o.Stackscript
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.Vpc) {
		toSerialize["vpc"] = o.Vpc
	}
	return toSerialize, nil
}

type NullableGetUserGrants200Response struct {
	value *GetUserGrants200Response
	isSet bool
}

func (v NullableGetUserGrants200Response) Get() *GetUserGrants200Response {
	return v.value
}

func (v *NullableGetUserGrants200Response) Set(val *GetUserGrants200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserGrants200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserGrants200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserGrants200Response(val *GetUserGrants200Response) *NullableGetUserGrants200Response {
	return &NullableGetUserGrants200Response{value: val, isSet: true}
}

func (v NullableGetUserGrants200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserGrants200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


