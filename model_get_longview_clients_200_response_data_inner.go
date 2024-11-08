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

// checks if the GetLongviewClients200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLongviewClients200ResponseDataInner{}

// GetLongviewClients200ResponseDataInner A LongviewClient is a single monitor set up to track statistics about one of your servers.
type GetLongviewClients200ResponseDataInner struct {
	// The API key for this Client, used when configuring the Longview Client application on your Linode.  Returns as `[REDACTED]` if you do not have read-write access to this client.
	ApiKey *string `json:"api_key,omitempty"`
	Apps *GetLongviewClients200ResponseDataInnerApps `json:"apps,omitempty"`
	// When this Longview Client was created.
	Created *time.Time `json:"created,omitempty"`
	// This Client's unique ID.
	Id *int32 `json:"id,omitempty"`
	// The install code for this Client, used when configuring the Longview Client application on your Linode.  Returns as `[REDACTED]` if you do not have read-write access to this client.
	InstallCode *string `json:"install_code,omitempty"`
	// This Client's unique label. This is for display purposes only.
	Label *string `json:"label,omitempty" validate:"regexp=[a-zA-Z0-9-_]{3,32}"`
	// When this Longview Client was last updated.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewGetLongviewClients200ResponseDataInner instantiates a new GetLongviewClients200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLongviewClients200ResponseDataInner() *GetLongviewClients200ResponseDataInner {
	this := GetLongviewClients200ResponseDataInner{}
	return &this
}

// NewGetLongviewClients200ResponseDataInnerWithDefaults instantiates a new GetLongviewClients200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLongviewClients200ResponseDataInnerWithDefaults() *GetLongviewClients200ResponseDataInner {
	this := GetLongviewClients200ResponseDataInner{}
	return &this
}

// GetApiKey returns the ApiKey field value if set, zero value otherwise.
func (o *GetLongviewClients200ResponseDataInner) GetApiKey() string {
	if o == nil || IsNil(o.ApiKey) {
		var ret string
		return ret
	}
	return *o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLongviewClients200ResponseDataInner) GetApiKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ApiKey) {
		return nil, false
	}
	return o.ApiKey, true
}

// HasApiKey returns a boolean if a field has been set.
func (o *GetLongviewClients200ResponseDataInner) HasApiKey() bool {
	if o != nil && !IsNil(o.ApiKey) {
		return true
	}

	return false
}

// SetApiKey gets a reference to the given string and assigns it to the ApiKey field.
func (o *GetLongviewClients200ResponseDataInner) SetApiKey(v string) {
	o.ApiKey = &v
}

// GetApps returns the Apps field value if set, zero value otherwise.
func (o *GetLongviewClients200ResponseDataInner) GetApps() GetLongviewClients200ResponseDataInnerApps {
	if o == nil || IsNil(o.Apps) {
		var ret GetLongviewClients200ResponseDataInnerApps
		return ret
	}
	return *o.Apps
}

// GetAppsOk returns a tuple with the Apps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLongviewClients200ResponseDataInner) GetAppsOk() (*GetLongviewClients200ResponseDataInnerApps, bool) {
	if o == nil || IsNil(o.Apps) {
		return nil, false
	}
	return o.Apps, true
}

// HasApps returns a boolean if a field has been set.
func (o *GetLongviewClients200ResponseDataInner) HasApps() bool {
	if o != nil && !IsNil(o.Apps) {
		return true
	}

	return false
}

// SetApps gets a reference to the given GetLongviewClients200ResponseDataInnerApps and assigns it to the Apps field.
func (o *GetLongviewClients200ResponseDataInner) SetApps(v GetLongviewClients200ResponseDataInnerApps) {
	o.Apps = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetLongviewClients200ResponseDataInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLongviewClients200ResponseDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetLongviewClients200ResponseDataInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetLongviewClients200ResponseDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetLongviewClients200ResponseDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLongviewClients200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetLongviewClients200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetLongviewClients200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetInstallCode returns the InstallCode field value if set, zero value otherwise.
func (o *GetLongviewClients200ResponseDataInner) GetInstallCode() string {
	if o == nil || IsNil(o.InstallCode) {
		var ret string
		return ret
	}
	return *o.InstallCode
}

// GetInstallCodeOk returns a tuple with the InstallCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLongviewClients200ResponseDataInner) GetInstallCodeOk() (*string, bool) {
	if o == nil || IsNil(o.InstallCode) {
		return nil, false
	}
	return o.InstallCode, true
}

// HasInstallCode returns a boolean if a field has been set.
func (o *GetLongviewClients200ResponseDataInner) HasInstallCode() bool {
	if o != nil && !IsNil(o.InstallCode) {
		return true
	}

	return false
}

// SetInstallCode gets a reference to the given string and assigns it to the InstallCode field.
func (o *GetLongviewClients200ResponseDataInner) SetInstallCode(v string) {
	o.InstallCode = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetLongviewClients200ResponseDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLongviewClients200ResponseDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetLongviewClients200ResponseDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetLongviewClients200ResponseDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *GetLongviewClients200ResponseDataInner) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLongviewClients200ResponseDataInner) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *GetLongviewClients200ResponseDataInner) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *GetLongviewClients200ResponseDataInner) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o GetLongviewClients200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLongviewClients200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiKey) {
		toSerialize["api_key"] = o.ApiKey
	}
	if !IsNil(o.Apps) {
		toSerialize["apps"] = o.Apps
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InstallCode) {
		toSerialize["install_code"] = o.InstallCode
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	return toSerialize, nil
}

type NullableGetLongviewClients200ResponseDataInner struct {
	value *GetLongviewClients200ResponseDataInner
	isSet bool
}

func (v NullableGetLongviewClients200ResponseDataInner) Get() *GetLongviewClients200ResponseDataInner {
	return v.value
}

func (v *NullableGetLongviewClients200ResponseDataInner) Set(val *GetLongviewClients200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLongviewClients200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLongviewClients200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLongviewClients200ResponseDataInner(val *GetLongviewClients200ResponseDataInner) *NullableGetLongviewClients200ResponseDataInner {
	return &NullableGetLongviewClients200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetLongviewClients200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLongviewClients200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


