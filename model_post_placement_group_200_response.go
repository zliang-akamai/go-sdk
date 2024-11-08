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

// checks if the PostPlacementGroup200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPlacementGroup200Response{}

// PostPlacementGroup200Response struct for PostPlacementGroup200Response
type PostPlacementGroup200Response struct {
	// The placement group's ID. You need to provide it for all operations that affect it.
	Id *int32 `json:"id,omitempty"`
	// Whether all of the compute instances in your placement group are compliant. If `true`, all compute instances meet either the grouped-together or spread-apart model, which you determine through your selected `placement_group_type`. If `false`, a compute instance is out of this compliance. For example, assume you've set `anti-affinity:local` as your `placement_group_type` and your group already has three qualifying compute instances on separate hosts, to support the spread-apart model. If a fourth compute instance is assigned that's on the same host as one of the existing three, the placement group is non-compliant. Enforce compliance in your group by setting a `placement_group_policy`.  > 📘 > > Fixing compliance is not self-service. You need to wait for our assistance to physically move compute instances to make the group compliant again.
	IsCompliant *bool `json:"is_compliant,omitempty"`
	// The unique name set for the placement group. A label has these constraints:  - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (`-`), underscores (`_`) or periods (`.`).
	Label *string `json:"label,omitempty"`
	// An array of compute instances included in the placement group.
	Members []GetPlacementGroups200ResponseDataInnerMembersInner `json:"members,omitempty"`
	// How requests to add future compute instances to your placement group are handled, and whether it remains compliant:  - `strict`. Don't assign a new compute instance if it breaks the grouped-together or spread-apart model set by the `placement_group_type`. Use this to ensure the placement group stays compliant (`is_compliant: true`). - `flexible`. Assign a new compute instance, even if it breaks the grouped-together or spread-apart model set by the `placement_group_type`. This makes the group non-compliant (`is_compliant: false`). You need to wait for Akamai to move the offending compute instance to make it compliant again, once the necessary capacity is available in the region. Offers flexibility to add future compute instances if compliance isn't an immediate concern.  <<LB>>  > 📘 > > In rare cases, non-compliance can occur with a `strict` placement group if Akamai needs to failover or migrate your compute instances for maintenance. Fixing non-compliance for a `strict` placement group is prioritized over a `flexible` group.
	PlacementGroupPolicy *string `json:"placement_group_policy,omitempty"`
	// How compute instances are distributed in your placement group. A `placement_group_type` using anti-affinity (`anti-affinity:local`) places compute instances in separate hosts, but still in the same region. This best supports the spread-apart model for high availability. A `placement_group_type` using affinity places compute instances physically close together, possibly on the same host. This supports the grouped-together model for low-latency.  > 📘 > > Currently, only `anti_affinity:local` is available for `placement_group_type`.
	PlacementGroupType *string `json:"placement_group_type,omitempty"`
	// The [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where the placement group was deployed.
	Region *string `json:"region,omitempty"`
}

// NewPostPlacementGroup200Response instantiates a new PostPlacementGroup200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPlacementGroup200Response() *PostPlacementGroup200Response {
	this := PostPlacementGroup200Response{}
	return &this
}

// NewPostPlacementGroup200ResponseWithDefaults instantiates a new PostPlacementGroup200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPlacementGroup200ResponseWithDefaults() *PostPlacementGroup200Response {
	this := PostPlacementGroup200Response{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostPlacementGroup200Response) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPlacementGroup200Response) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostPlacementGroup200Response) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PostPlacementGroup200Response) SetId(v int32) {
	o.Id = &v
}

// GetIsCompliant returns the IsCompliant field value if set, zero value otherwise.
func (o *PostPlacementGroup200Response) GetIsCompliant() bool {
	if o == nil || IsNil(o.IsCompliant) {
		var ret bool
		return ret
	}
	return *o.IsCompliant
}

// GetIsCompliantOk returns a tuple with the IsCompliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPlacementGroup200Response) GetIsCompliantOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompliant) {
		return nil, false
	}
	return o.IsCompliant, true
}

// HasIsCompliant returns a boolean if a field has been set.
func (o *PostPlacementGroup200Response) HasIsCompliant() bool {
	if o != nil && !IsNil(o.IsCompliant) {
		return true
	}

	return false
}

// SetIsCompliant gets a reference to the given bool and assigns it to the IsCompliant field.
func (o *PostPlacementGroup200Response) SetIsCompliant(v bool) {
	o.IsCompliant = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PostPlacementGroup200Response) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPlacementGroup200Response) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PostPlacementGroup200Response) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PostPlacementGroup200Response) SetLabel(v string) {
	o.Label = &v
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *PostPlacementGroup200Response) GetMembers() []GetPlacementGroups200ResponseDataInnerMembersInner {
	if o == nil || IsNil(o.Members) {
		var ret []GetPlacementGroups200ResponseDataInnerMembersInner
		return ret
	}
	return o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPlacementGroup200Response) GetMembersOk() ([]GetPlacementGroups200ResponseDataInnerMembersInner, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *PostPlacementGroup200Response) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given []GetPlacementGroups200ResponseDataInnerMembersInner and assigns it to the Members field.
func (o *PostPlacementGroup200Response) SetMembers(v []GetPlacementGroups200ResponseDataInnerMembersInner) {
	o.Members = v
}

// GetPlacementGroupPolicy returns the PlacementGroupPolicy field value if set, zero value otherwise.
func (o *PostPlacementGroup200Response) GetPlacementGroupPolicy() string {
	if o == nil || IsNil(o.PlacementGroupPolicy) {
		var ret string
		return ret
	}
	return *o.PlacementGroupPolicy
}

// GetPlacementGroupPolicyOk returns a tuple with the PlacementGroupPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPlacementGroup200Response) GetPlacementGroupPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.PlacementGroupPolicy) {
		return nil, false
	}
	return o.PlacementGroupPolicy, true
}

// HasPlacementGroupPolicy returns a boolean if a field has been set.
func (o *PostPlacementGroup200Response) HasPlacementGroupPolicy() bool {
	if o != nil && !IsNil(o.PlacementGroupPolicy) {
		return true
	}

	return false
}

// SetPlacementGroupPolicy gets a reference to the given string and assigns it to the PlacementGroupPolicy field.
func (o *PostPlacementGroup200Response) SetPlacementGroupPolicy(v string) {
	o.PlacementGroupPolicy = &v
}

// GetPlacementGroupType returns the PlacementGroupType field value if set, zero value otherwise.
func (o *PostPlacementGroup200Response) GetPlacementGroupType() string {
	if o == nil || IsNil(o.PlacementGroupType) {
		var ret string
		return ret
	}
	return *o.PlacementGroupType
}

// GetPlacementGroupTypeOk returns a tuple with the PlacementGroupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPlacementGroup200Response) GetPlacementGroupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PlacementGroupType) {
		return nil, false
	}
	return o.PlacementGroupType, true
}

// HasPlacementGroupType returns a boolean if a field has been set.
func (o *PostPlacementGroup200Response) HasPlacementGroupType() bool {
	if o != nil && !IsNil(o.PlacementGroupType) {
		return true
	}

	return false
}

// SetPlacementGroupType gets a reference to the given string and assigns it to the PlacementGroupType field.
func (o *PostPlacementGroup200Response) SetPlacementGroupType(v string) {
	o.PlacementGroupType = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PostPlacementGroup200Response) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostPlacementGroup200Response) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PostPlacementGroup200Response) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PostPlacementGroup200Response) SetRegion(v string) {
	o.Region = &v
}

func (o PostPlacementGroup200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPlacementGroup200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsCompliant) {
		toSerialize["is_compliant"] = o.IsCompliant
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	if !IsNil(o.PlacementGroupPolicy) {
		toSerialize["placement_group_policy"] = o.PlacementGroupPolicy
	}
	if !IsNil(o.PlacementGroupType) {
		toSerialize["placement_group_type"] = o.PlacementGroupType
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	return toSerialize, nil
}

type NullablePostPlacementGroup200Response struct {
	value *PostPlacementGroup200Response
	isSet bool
}

func (v NullablePostPlacementGroup200Response) Get() *PostPlacementGroup200Response {
	return v.value
}

func (v *NullablePostPlacementGroup200Response) Set(val *PostPlacementGroup200Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPlacementGroup200Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPlacementGroup200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPlacementGroup200Response(val *PostPlacementGroup200Response) *NullablePostPlacementGroup200Response {
	return &NullablePostPlacementGroup200Response{value: val, isSet: true}
}

func (v NullablePostPlacementGroup200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPlacementGroup200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


