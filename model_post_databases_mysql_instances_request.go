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

// checks if the PostDatabasesMysqlInstancesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostDatabasesMysqlInstancesRequest{}

// PostDatabasesMysqlInstancesRequest Managed MySQL Database request object.
type PostDatabasesMysqlInstancesRequest struct {
	// A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.  By default, this is an empty array (`[]`), which blocks all connections (both public and private) to the Managed Database.  If `0.0.0.0/0` is a value in this list, then all IP addresses can access the Managed Database.
	AllowList []string `json:"allow_list,omitempty"`
	// The number of Linode Instance nodes deployed to the Managed Database.  Choosing 3 nodes creates a high availability cluster consisting of 1 primary node and 2 replica nodes.
	ClusterSize *int32 `json:"cluster_size,omitempty"`
	// Whether the Managed Databases is encrypted.
	Encrypted *bool `json:"encrypted,omitempty"`
	// The Managed Database engine in engine/version format.
	Engine string `json:"engine"`
	// A unique, user-defined string referring to the Managed Database.
	Label string `json:"label"`
	// The [Region](https://techdocs.akamai.com/linode-api/reference/get-regions) ID for the Managed Database.
	Region string `json:"region"`
	// The replication method used for the Managed Database.  Defaults to `none` for a single cluster and `semi_synch` for a high availability cluster.  Must be `none` for a single node cluster.  Must be `asynch` or `semi_synch` for a high availability cluster.
	ReplicationType *string `json:"replication_type,omitempty"`
	// Whether to require SSL credentials to establish a connection to the Managed Database.  Run the [Get managed MySQL database credentials](https://techdocs.akamai.com/linode-api/reference/get-databases-mysql-instance-credentials) operation for access information.
	SslConnection *bool `json:"ssl_connection,omitempty"`
	// The Linode Instance type used by the Managed Database for its nodes.
	Type string `json:"type"`
}

type _PostDatabasesMysqlInstancesRequest PostDatabasesMysqlInstancesRequest

// NewPostDatabasesMysqlInstancesRequest instantiates a new PostDatabasesMysqlInstancesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostDatabasesMysqlInstancesRequest(engine string, label string, region string, type_ string) *PostDatabasesMysqlInstancesRequest {
	this := PostDatabasesMysqlInstancesRequest{}
	var clusterSize int32 = 1
	this.ClusterSize = &clusterSize
	var encrypted bool = false
	this.Encrypted = &encrypted
	this.Engine = engine
	this.Label = label
	this.Region = region
	var sslConnection bool = true
	this.SslConnection = &sslConnection
	this.Type = type_
	return &this
}

// NewPostDatabasesMysqlInstancesRequestWithDefaults instantiates a new PostDatabasesMysqlInstancesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostDatabasesMysqlInstancesRequestWithDefaults() *PostDatabasesMysqlInstancesRequest {
	this := PostDatabasesMysqlInstancesRequest{}
	var clusterSize int32 = 1
	this.ClusterSize = &clusterSize
	var encrypted bool = false
	this.Encrypted = &encrypted
	var sslConnection bool = true
	this.SslConnection = &sslConnection
	return &this
}

// GetAllowList returns the AllowList field value if set, zero value otherwise.
func (o *PostDatabasesMysqlInstancesRequest) GetAllowList() []string {
	if o == nil || IsNil(o.AllowList) {
		var ret []string
		return ret
	}
	return o.AllowList
}

// GetAllowListOk returns a tuple with the AllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDatabasesMysqlInstancesRequest) GetAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowList) {
		return nil, false
	}
	return o.AllowList, true
}

// HasAllowList returns a boolean if a field has been set.
func (o *PostDatabasesMysqlInstancesRequest) HasAllowList() bool {
	if o != nil && !IsNil(o.AllowList) {
		return true
	}

	return false
}

// SetAllowList gets a reference to the given []string and assigns it to the AllowList field.
func (o *PostDatabasesMysqlInstancesRequest) SetAllowList(v []string) {
	o.AllowList = v
}

// GetClusterSize returns the ClusterSize field value if set, zero value otherwise.
func (o *PostDatabasesMysqlInstancesRequest) GetClusterSize() int32 {
	if o == nil || IsNil(o.ClusterSize) {
		var ret int32
		return ret
	}
	return *o.ClusterSize
}

// GetClusterSizeOk returns a tuple with the ClusterSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDatabasesMysqlInstancesRequest) GetClusterSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.ClusterSize) {
		return nil, false
	}
	return o.ClusterSize, true
}

// HasClusterSize returns a boolean if a field has been set.
func (o *PostDatabasesMysqlInstancesRequest) HasClusterSize() bool {
	if o != nil && !IsNil(o.ClusterSize) {
		return true
	}

	return false
}

// SetClusterSize gets a reference to the given int32 and assigns it to the ClusterSize field.
func (o *PostDatabasesMysqlInstancesRequest) SetClusterSize(v int32) {
	o.ClusterSize = &v
}

// GetEncrypted returns the Encrypted field value if set, zero value otherwise.
func (o *PostDatabasesMysqlInstancesRequest) GetEncrypted() bool {
	if o == nil || IsNil(o.Encrypted) {
		var ret bool
		return ret
	}
	return *o.Encrypted
}

// GetEncryptedOk returns a tuple with the Encrypted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDatabasesMysqlInstancesRequest) GetEncryptedOk() (*bool, bool) {
	if o == nil || IsNil(o.Encrypted) {
		return nil, false
	}
	return o.Encrypted, true
}

// HasEncrypted returns a boolean if a field has been set.
func (o *PostDatabasesMysqlInstancesRequest) HasEncrypted() bool {
	if o != nil && !IsNil(o.Encrypted) {
		return true
	}

	return false
}

// SetEncrypted gets a reference to the given bool and assigns it to the Encrypted field.
func (o *PostDatabasesMysqlInstancesRequest) SetEncrypted(v bool) {
	o.Encrypted = &v
}

// GetEngine returns the Engine field value
func (o *PostDatabasesMysqlInstancesRequest) GetEngine() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Engine
}

// GetEngineOk returns a tuple with the Engine field value
// and a boolean to check if the value has been set.
func (o *PostDatabasesMysqlInstancesRequest) GetEngineOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Engine, true
}

// SetEngine sets field value
func (o *PostDatabasesMysqlInstancesRequest) SetEngine(v string) {
	o.Engine = v
}

// GetLabel returns the Label field value
func (o *PostDatabasesMysqlInstancesRequest) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *PostDatabasesMysqlInstancesRequest) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *PostDatabasesMysqlInstancesRequest) SetLabel(v string) {
	o.Label = v
}

// GetRegion returns the Region field value
func (o *PostDatabasesMysqlInstancesRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *PostDatabasesMysqlInstancesRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *PostDatabasesMysqlInstancesRequest) SetRegion(v string) {
	o.Region = v
}

// GetReplicationType returns the ReplicationType field value if set, zero value otherwise.
func (o *PostDatabasesMysqlInstancesRequest) GetReplicationType() string {
	if o == nil || IsNil(o.ReplicationType) {
		var ret string
		return ret
	}
	return *o.ReplicationType
}

// GetReplicationTypeOk returns a tuple with the ReplicationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDatabasesMysqlInstancesRequest) GetReplicationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ReplicationType) {
		return nil, false
	}
	return o.ReplicationType, true
}

// HasReplicationType returns a boolean if a field has been set.
func (o *PostDatabasesMysqlInstancesRequest) HasReplicationType() bool {
	if o != nil && !IsNil(o.ReplicationType) {
		return true
	}

	return false
}

// SetReplicationType gets a reference to the given string and assigns it to the ReplicationType field.
func (o *PostDatabasesMysqlInstancesRequest) SetReplicationType(v string) {
	o.ReplicationType = &v
}

// GetSslConnection returns the SslConnection field value if set, zero value otherwise.
func (o *PostDatabasesMysqlInstancesRequest) GetSslConnection() bool {
	if o == nil || IsNil(o.SslConnection) {
		var ret bool
		return ret
	}
	return *o.SslConnection
}

// GetSslConnectionOk returns a tuple with the SslConnection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostDatabasesMysqlInstancesRequest) GetSslConnectionOk() (*bool, bool) {
	if o == nil || IsNil(o.SslConnection) {
		return nil, false
	}
	return o.SslConnection, true
}

// HasSslConnection returns a boolean if a field has been set.
func (o *PostDatabasesMysqlInstancesRequest) HasSslConnection() bool {
	if o != nil && !IsNil(o.SslConnection) {
		return true
	}

	return false
}

// SetSslConnection gets a reference to the given bool and assigns it to the SslConnection field.
func (o *PostDatabasesMysqlInstancesRequest) SetSslConnection(v bool) {
	o.SslConnection = &v
}

// GetType returns the Type field value
func (o *PostDatabasesMysqlInstancesRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostDatabasesMysqlInstancesRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostDatabasesMysqlInstancesRequest) SetType(v string) {
	o.Type = v
}

func (o PostDatabasesMysqlInstancesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostDatabasesMysqlInstancesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowList) {
		toSerialize["allow_list"] = o.AllowList
	}
	if !IsNil(o.ClusterSize) {
		toSerialize["cluster_size"] = o.ClusterSize
	}
	if !IsNil(o.Encrypted) {
		toSerialize["encrypted"] = o.Encrypted
	}
	toSerialize["engine"] = o.Engine
	toSerialize["label"] = o.Label
	toSerialize["region"] = o.Region
	if !IsNil(o.ReplicationType) {
		toSerialize["replication_type"] = o.ReplicationType
	}
	if !IsNil(o.SslConnection) {
		toSerialize["ssl_connection"] = o.SslConnection
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *PostDatabasesMysqlInstancesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"engine",
		"label",
		"region",
		"type",
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

	varPostDatabasesMysqlInstancesRequest := _PostDatabasesMysqlInstancesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostDatabasesMysqlInstancesRequest)

	if err != nil {
		return err
	}

	*o = PostDatabasesMysqlInstancesRequest(varPostDatabasesMysqlInstancesRequest)

	return err
}

type NullablePostDatabasesMysqlInstancesRequest struct {
	value *PostDatabasesMysqlInstancesRequest
	isSet bool
}

func (v NullablePostDatabasesMysqlInstancesRequest) Get() *PostDatabasesMysqlInstancesRequest {
	return v.value
}

func (v *NullablePostDatabasesMysqlInstancesRequest) Set(val *PostDatabasesMysqlInstancesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostDatabasesMysqlInstancesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostDatabasesMysqlInstancesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostDatabasesMysqlInstancesRequest(val *PostDatabasesMysqlInstancesRequest) *NullablePostDatabasesMysqlInstancesRequest {
	return &NullablePostDatabasesMysqlInstancesRequest{value: val, isSet: true}
}

func (v NullablePostDatabasesMysqlInstancesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostDatabasesMysqlInstancesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


