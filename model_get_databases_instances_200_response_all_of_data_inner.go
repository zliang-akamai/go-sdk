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

// checks if the GetDatabasesInstances200ResponseAllOfDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDatabasesInstances200ResponseAllOfDataInner{}

// GetDatabasesInstances200ResponseAllOfDataInner A general Managed Database instance object containing properties that are identical for all database types.
type GetDatabasesInstances200ResponseAllOfDataInner struct {
	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.  By default, this is an empty array (`[]`), which blocks all connections (both public and private) to the Managed Database.  If `0.0.0.0/0` is a value in this list, then all IP addresses can access the Managed Database.
	AllowList []string `json:"allow_list,omitempty"`
	// The number of Linode Instance nodes deployed to the Managed Database.  Choosing 3 nodes creates a high availability cluster consisting of 1 primary node and 2 replica nodes.
	ClusterSize *int32 `json:"cluster_size,omitempty"`
	// When this Managed Database was created.
	Created *time.Time `json:"created,omitempty"`
	// Whether the Managed Databases is encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
	// The Managed Database engine type.
	Engine *string `json:"engine,omitempty"`
	Hosts *GetDatabasesInstances200ResponseAllOfDataInnerHosts `json:"hosts,omitempty"`
	// A unique ID that can be used to identify and reference the Managed Database.
	Id *int32 `json:"id,omitempty"`
	// Append this to `https://api.linode.com` to run commands for the Managed Database.
	InstanceUri *string `json:"instance_uri,omitempty"`
	// A unique, user-defined string referring to the Managed Database.
	Label *string `json:"label,omitempty"`
	// The [Region](https://techdocs.akamai.com/linode-api/reference/get-regions) ID for the Managed Database.
	Region *string `json:"region,omitempty"`
	// The operating status of the Managed Database.
	Status *string `json:"status,omitempty"`
	// The Linode Instance type used by the Managed Database for its nodes.
	Type *string `json:"type,omitempty"`
	// When this Managed Database was last updated.
	Updated *time.Time `json:"updated,omitempty"`
	Updates *GetDatabasesInstances200ResponseAllOfDataInnerUpdates `json:"updates,omitempty"`
	// The Managed Database engine version.
	Version *string `json:"version,omitempty"`
}

// NewGetDatabasesInstances200ResponseAllOfDataInner instantiates a new GetDatabasesInstances200ResponseAllOfDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDatabasesInstances200ResponseAllOfDataInner() *GetDatabasesInstances200ResponseAllOfDataInner {
	this := GetDatabasesInstances200ResponseAllOfDataInner{}
	var clusterSize int32 = 1
	this.ClusterSize = &clusterSize
	var encrypted bool = false
	this.Encrypted = &encrypted
	return &this
}

// NewGetDatabasesInstances200ResponseAllOfDataInnerWithDefaults instantiates a new GetDatabasesInstances200ResponseAllOfDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDatabasesInstances200ResponseAllOfDataInnerWithDefaults() *GetDatabasesInstances200ResponseAllOfDataInner {
	this := GetDatabasesInstances200ResponseAllOfDataInner{}
	var clusterSize int32 = 1
	this.ClusterSize = &clusterSize
	var encrypted bool = false
	this.Encrypted = &encrypted
	return &this
}

// GetAllowList returns the AllowList field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetAllowList() []string {
	if o == nil || IsNil(o.AllowList) {
		var ret []string
		return ret
	}
	return o.AllowList
}

// GetAllowListOk returns a tuple with the AllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowList) {
		return nil, false
	}
	return o.AllowList, true
}

// HasAllowList returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasAllowList() bool {
	if o != nil && !IsNil(o.AllowList) {
		return true
	}

	return false
}

// SetAllowList gets a reference to the given []string and assigns it to the AllowList field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetAllowList(v []string) {
	o.AllowList = v
}

// GetClusterSize returns the ClusterSize field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetClusterSize() int32 {
	if o == nil || IsNil(o.ClusterSize) {
		var ret int32
		return ret
	}
	return *o.ClusterSize
}

// GetClusterSizeOk returns a tuple with the ClusterSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetClusterSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ClusterSize) {
		return nil, false
	}
	return o.ClusterSize, true
}

// HasClusterSize returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasClusterSize() bool {
	if o != nil && !IsNil(o.ClusterSize) {
		return true
	}

	return false
}

// SetClusterSize gets a reference to the given int32 and assigns it to the ClusterSize field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetClusterSize(v int32) {
	o.ClusterSize = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetEncrypted() bool {
	if o == nil || IsNil(o.Encrypted) {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.Encrypted) {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasEncrypted() bool {
	if o != nil && !IsNil(o.Encrypted) {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetEngine returns the Engine field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetEngine() string {
	if o == nil || IsNil(o.Engine) {
		var ret string
		return ret
	}
	return *o.Engine
}

// GetEngineOk returns a tuple with the Engine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetEngineOk() (*string, bool) {
	if o == nil || IsNil(o.Engine) {
		return nil, false
	}
	return o.Engine, true
}

// HasEngine returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasEngine() bool {
	if o != nil && !IsNil(o.Engine) {
		return true
	}

	return false
}

// SetEngine gets a reference to the given string and assigns it to the Engine field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetEngine(v string) {
	o.Engine = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetHosts() GetDatabasesInstances200ResponseAllOfDataInnerHosts {
	if o == nil || IsNil(o.Hosts) {
		var ret GetDatabasesInstances200ResponseAllOfDataInnerHosts
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetHostsOk() (*GetDatabasesInstances200ResponseAllOfDataInnerHosts, bool) {
	if o == nil || IsNil(o.Hosts) {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasHosts() bool {
	if o != nil && !IsNil(o.Hosts) {
		return true
	}

	return false
}

// SetHosts gets a reference to the given GetDatabasesInstances200ResponseAllOfDataInnerHosts and assigns it to the Hosts field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetHosts(v GetDatabasesInstances200ResponseAllOfDataInnerHosts) {
	o.Hosts = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetId(v int32) {
	o.Id = &v
}

// GetInstanceUri returns the InstanceUri field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetInstanceUri() string {
	if o == nil || IsNil(o.InstanceUri) {
		var ret string
		return ret
	}
	return *o.InstanceUri
}

// GetInstanceUriOk returns a tuple with the InstanceUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetInstanceUriOk() (*string, bool) {
	if o == nil || IsNil(o.InstanceUri) {
		return nil, false
	}
	return o.InstanceUri, true
}

// HasInstanceUri returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasInstanceUri() bool {
	if o != nil && !IsNil(o.InstanceUri) {
		return true
	}

	return false
}

// SetInstanceUri gets a reference to the given string and assigns it to the InstanceUri field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetInstanceUri(v string) {
	o.InstanceUri = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetRegion(v string) {
	o.Region = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetType(v string) {
	o.Type = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetUpdates returns the Updates field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetUpdates() GetDatabasesInstances200ResponseAllOfDataInnerUpdates {
	if o == nil || IsNil(o.Updates) {
		var ret GetDatabasesInstances200ResponseAllOfDataInnerUpdates
		return ret
	}
	return *o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetUpdatesOk() (*GetDatabasesInstances200ResponseAllOfDataInnerUpdates, bool) {
	if o == nil || IsNil(o.Updates) {
		return nil, false
	}
	return o.Updates, true
}

// HasUpdates returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasUpdates() bool {
	if o != nil && !IsNil(o.Updates) {
		return true
	}

	return false
}

// SetUpdates gets a reference to the given GetDatabasesInstances200ResponseAllOfDataInnerUpdates and assigns it to the Updates field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetUpdates(v GetDatabasesInstances200ResponseAllOfDataInnerUpdates) {
	o.Updates = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetVersion(v string) {
	o.Version = &v
}

func (o GetDatabasesInstances200ResponseAllOfDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDatabasesInstances200ResponseAllOfDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowList) {
		toSerialize["allow_list"] = o.AllowList
	}
	if !IsNil(o.ClusterSize) {
		toSerialize["cluster_size"] = o.ClusterSize
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Encrypted) {
		toSerialize["encrypted"] = o.Encrypted
	}
	if !IsNil(o.Engine) {
		toSerialize["engine"] = o.Engine
	}
	if !IsNil(o.Hosts) {
		toSerialize["hosts"] = o.Hosts
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InstanceUri) {
		toSerialize["instance_uri"] = o.InstanceUri
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.Updates) {
		toSerialize["updates"] = o.Updates
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableGetDatabasesInstances200ResponseAllOfDataInner struct {
	value *GetDatabasesInstances200ResponseAllOfDataInner
	isSet bool
}

func (v NullableGetDatabasesInstances200ResponseAllOfDataInner) Get() *GetDatabasesInstances200ResponseAllOfDataInner {
	return v.value
}

func (v *NullableGetDatabasesInstances200ResponseAllOfDataInner) Set(val *GetDatabasesInstances200ResponseAllOfDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDatabasesInstances200ResponseAllOfDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDatabasesInstances200ResponseAllOfDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDatabasesInstances200ResponseAllOfDataInner(val *GetDatabasesInstances200ResponseAllOfDataInner) *NullableGetDatabasesInstances200ResponseAllOfDataInner {
	return &NullableGetDatabasesInstances200ResponseAllOfDataInner{value: val, isSet: true}
}

func (v NullableGetDatabasesInstances200ResponseAllOfDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDatabasesInstances200ResponseAllOfDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


