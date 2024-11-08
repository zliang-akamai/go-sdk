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

// checks if the GetManagedServices200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetManagedServices200ResponseDataInner{}

// GetManagedServices200ResponseDataInner A service that Linode is monitoring as part of your Managed services. If issues are detected with this service, a ManagedIssue will be opened and, optionally, Linode special forces will attempt to resolve the Issue.
type GetManagedServices200ResponseDataInner struct {
	// The URL at which this Service is monitored. URL parameters such as `?no-cache=1` are preserved. URL fragments/anchors such as `#monitor` are __not__ preserved.
	Address *string `json:"address,omitempty"`
	// What to expect to find in the response body for the Service to be considered up.
	Body NullableString `json:"body,omitempty"`
	// The group of ManagedContacts who should be notified or consulted with when an Issue is detected.
	ConsultationGroup *string `json:"consultation_group,omitempty"`
	// When this Managed Service was created.
	Created *time.Time `json:"created,omitempty"`
	// An array of ManagedCredential IDs that should be used when attempting to resolve issues with this Service.
	Credentials []int32 `json:"credentials,omitempty"`
	// This Service's unique ID.
	Id *int32 `json:"id,omitempty"`
	// The label for this Service. This is for display purposes only.
	Label *string `json:"label,omitempty" validate:"regexp=[a-zA-Z0-9-_ \\\\.]{3,64}"`
	// Any information relevant to the Service that Linode special forces should know when attempting to resolve Issues.
	Notes NullableString `json:"notes,omitempty"`
	// The Region in which this Service is located. This is required if address is a private IP, and may not be set otherwise.
	Region *string `json:"region,omitempty"`
	// How this Service is monitored.
	ServiceType *string `json:"service_type,omitempty"`
	// The current status of this Service.
	Status *string `json:"status,omitempty"`
	// How long to wait, in seconds, for a response before considering the Service to be down.
	Timeout *int32 `json:"timeout,omitempty"`
	// When this Managed Service was last updated.
	Updated *time.Time `json:"updated,omitempty"`
}

// NewGetManagedServices200ResponseDataInner instantiates a new GetManagedServices200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetManagedServices200ResponseDataInner() *GetManagedServices200ResponseDataInner {
	this := GetManagedServices200ResponseDataInner{}
	return &this
}

// NewGetManagedServices200ResponseDataInnerWithDefaults instantiates a new GetManagedServices200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetManagedServices200ResponseDataInnerWithDefaults() *GetManagedServices200ResponseDataInner {
	this := GetManagedServices200ResponseDataInner{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GetManagedServices200ResponseDataInner) SetAddress(v string) {
	o.Address = &v
}

// GetBody returns the Body field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetManagedServices200ResponseDataInner) GetBody() string {
	if o == nil || IsNil(o.Body.Get()) {
		var ret string
		return ret
	}
	return *o.Body.Get()
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetManagedServices200ResponseDataInner) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body.Get(), o.Body.IsSet()
}

// HasBody returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasBody() bool {
	if o != nil && o.Body.IsSet() {
		return true
	}

	return false
}

// SetBody gets a reference to the given NullableString and assigns it to the Body field.
func (o *GetManagedServices200ResponseDataInner) SetBody(v string) {
	o.Body.Set(&v)
}
// SetBodyNil sets the value for Body to be an explicit nil
func (o *GetManagedServices200ResponseDataInner) SetBodyNil() {
	o.Body.Set(nil)
}

// UnsetBody ensures that no value is present for Body, not even an explicit nil
func (o *GetManagedServices200ResponseDataInner) UnsetBody() {
	o.Body.Unset()
}

// GetConsultationGroup returns the ConsultationGroup field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetConsultationGroup() string {
	if o == nil || IsNil(o.ConsultationGroup) {
		var ret string
		return ret
	}
	return *o.ConsultationGroup
}

// GetConsultationGroupOk returns a tuple with the ConsultationGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetConsultationGroupOk() (*string, bool) {
	if o == nil || IsNil(o.ConsultationGroup) {
		return nil, false
	}
	return o.ConsultationGroup, true
}

// HasConsultationGroup returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasConsultationGroup() bool {
	if o != nil && !IsNil(o.ConsultationGroup) {
		return true
	}

	return false
}

// SetConsultationGroup gets a reference to the given string and assigns it to the ConsultationGroup field.
func (o *GetManagedServices200ResponseDataInner) SetConsultationGroup(v string) {
	o.ConsultationGroup = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetManagedServices200ResponseDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetCredentials() []int32 {
	if o == nil || IsNil(o.Credentials) {
		var ret []int32
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetCredentialsOk() ([]int32, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []int32 and assigns it to the Credentials field.
func (o *GetManagedServices200ResponseDataInner) SetCredentials(v []int32) {
	o.Credentials = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetManagedServices200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *GetManagedServices200ResponseDataInner) SetLabel(v string) {
	o.Label = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetManagedServices200ResponseDataInner) GetNotes() string {
	if o == nil || IsNil(o.Notes.Get()) {
		var ret string
		return ret
	}
	return *o.Notes.Get()
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetManagedServices200ResponseDataInner) GetNotesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Notes.Get(), o.Notes.IsSet()
}

// HasNotes returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasNotes() bool {
	if o != nil && o.Notes.IsSet() {
		return true
	}

	return false
}

// SetNotes gets a reference to the given NullableString and assigns it to the Notes field.
func (o *GetManagedServices200ResponseDataInner) SetNotes(v string) {
	o.Notes.Set(&v)
}
// SetNotesNil sets the value for Notes to be an explicit nil
func (o *GetManagedServices200ResponseDataInner) SetNotesNil() {
	o.Notes.Set(nil)
}

// UnsetNotes ensures that no value is present for Notes, not even an explicit nil
func (o *GetManagedServices200ResponseDataInner) UnsetNotes() {
	o.Notes.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GetManagedServices200ResponseDataInner) SetRegion(v string) {
	o.Region = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetServiceType() string {
	if o == nil || IsNil(o.ServiceType) {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *GetManagedServices200ResponseDataInner) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetManagedServices200ResponseDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *GetManagedServices200ResponseDataInner) SetTimeout(v int32) {
	o.Timeout = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *GetManagedServices200ResponseDataInner) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetManagedServices200ResponseDataInner) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *GetManagedServices200ResponseDataInner) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *GetManagedServices200ResponseDataInner) SetUpdated(v time.Time) {
	o.Updated = &v
}

func (o GetManagedServices200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetManagedServices200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if o.Body.IsSet() {
		toSerialize["body"] = o.Body.Get()
	}
	if !IsNil(o.ConsultationGroup) {
		toSerialize["consultation_group"] = o.ConsultationGroup
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Credentials) {
		toSerialize["credentials"] = o.Credentials
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.Notes.IsSet() {
		toSerialize["notes"] = o.Notes.Get()
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.ServiceType) {
		toSerialize["service_type"] = o.ServiceType
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	return toSerialize, nil
}

type NullableGetManagedServices200ResponseDataInner struct {
	value *GetManagedServices200ResponseDataInner
	isSet bool
}

func (v NullableGetManagedServices200ResponseDataInner) Get() *GetManagedServices200ResponseDataInner {
	return v.value
}

func (v *NullableGetManagedServices200ResponseDataInner) Set(val *GetManagedServices200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetManagedServices200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetManagedServices200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetManagedServices200ResponseDataInner(val *GetManagedServices200ResponseDataInner) *NullableGetManagedServices200ResponseDataInner {
	return &NullableGetManagedServices200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetManagedServices200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetManagedServices200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


