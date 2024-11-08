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

// checks if the GetEvents200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEvents200ResponseDataInner{}

// GetEvents200ResponseDataInner A collection of Event objects. An Event is an action taken against an entity related to your Account. For example, booting a Linode would create an Event. The Events returned depends on your grants.
type GetEvents200ResponseDataInner struct {
	// The action that caused this Event. New actions may be added in the future.
	Action *string `json:"action,omitempty"`
	// When this Event was created.
	Created *time.Time `json:"created,omitempty"`
	// The total duration in seconds that it takes for the Event to complete.
	Duration *float32 `json:"duration,omitempty"`
	Entity *GetEvents200ResponseDataInnerEntity `json:"entity,omitempty"`
	// The unique ID of this Event.
	Id *int32 `json:"id,omitempty"`
	// Provides additional information about the event. Additional information may include, but is not limited to, a more detailed representation of events which can help diagnose non-obvious failures.
	Message NullableString `json:"message,omitempty"`
	// A percentage estimating the amount of time remaining for an Event. Returns `null` for notification events.
	PercentComplete *int32 `json:"percent_complete,omitempty"`
	// The rate of completion of the Event. Only some Events will return rate; for example, migration and resize Events.
	Rate *string `json:"rate,omitempty"`
	// If this Event has been read.
	Read *bool `json:"read,omitempty"`
	SecondaryEntity *GetEvents200ResponseDataInnerSecondaryEntity `json:"secondary_entity,omitempty"`
	// If this Event has been seen.
	Seen *bool `json:"seen,omitempty"`
	// The current status of this Event.
	Status *string `json:"status,omitempty"`
	// The estimated time remaining until the completion of this Event. This value is only returned for some in-progress migration events. For all other in-progress events, the `percent_complete` attribute will indicate about how much more work is to be done.
	TimeRemaining NullableString `json:"time_remaining,omitempty"`
	// The username of the User who caused the Event.
	Username NullableString `json:"username,omitempty"`
}

// NewGetEvents200ResponseDataInner instantiates a new GetEvents200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEvents200ResponseDataInner() *GetEvents200ResponseDataInner {
	this := GetEvents200ResponseDataInner{}
	return &this
}

// NewGetEvents200ResponseDataInnerWithDefaults instantiates a new GetEvents200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEvents200ResponseDataInnerWithDefaults() *GetEvents200ResponseDataInner {
	this := GetEvents200ResponseDataInner{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *GetEvents200ResponseDataInner) SetAction(v string) {
	o.Action = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *GetEvents200ResponseDataInner) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetDuration() float32 {
	if o == nil || IsNil(o.Duration) {
		var ret float32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetDurationOk() (*float32, bool) {
	if o == nil || IsNil(o.Duration) {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasDuration() bool {
	if o != nil && !IsNil(o.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given float32 and assigns it to the Duration field.
func (o *GetEvents200ResponseDataInner) SetDuration(v float32) {
	o.Duration = &v
}

// GetEntity returns the Entity field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetEntity() GetEvents200ResponseDataInnerEntity {
	if o == nil || IsNil(o.Entity) {
		var ret GetEvents200ResponseDataInnerEntity
		return ret
	}
	return *o.Entity
}

// GetEntityOk returns a tuple with the Entity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetEntityOk() (*GetEvents200ResponseDataInnerEntity, bool) {
	if o == nil || IsNil(o.Entity) {
		return nil, false
	}
	return o.Entity, true
}

// HasEntity returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasEntity() bool {
	if o != nil && !IsNil(o.Entity) {
		return true
	}

	return false
}

// SetEntity gets a reference to the given GetEvents200ResponseDataInnerEntity and assigns it to the Entity field.
func (o *GetEvents200ResponseDataInner) SetEntity(v GetEvents200ResponseDataInnerEntity) {
	o.Entity = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *GetEvents200ResponseDataInner) SetId(v int32) {
	o.Id = &v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetEvents200ResponseDataInner) GetMessage() string {
	if o == nil || IsNil(o.Message.Get()) {
		var ret string
		return ret
	}
	return *o.Message.Get()
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEvents200ResponseDataInner) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Message.Get(), o.Message.IsSet()
}

// HasMessage returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasMessage() bool {
	if o != nil && o.Message.IsSet() {
		return true
	}

	return false
}

// SetMessage gets a reference to the given NullableString and assigns it to the Message field.
func (o *GetEvents200ResponseDataInner) SetMessage(v string) {
	o.Message.Set(&v)
}
// SetMessageNil sets the value for Message to be an explicit nil
func (o *GetEvents200ResponseDataInner) SetMessageNil() {
	o.Message.Set(nil)
}

// UnsetMessage ensures that no value is present for Message, not even an explicit nil
func (o *GetEvents200ResponseDataInner) UnsetMessage() {
	o.Message.Unset()
}

// GetPercentComplete returns the PercentComplete field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetPercentComplete() int32 {
	if o == nil || IsNil(o.PercentComplete) {
		var ret int32
		return ret
	}
	return *o.PercentComplete
}

// GetPercentCompleteOk returns a tuple with the PercentComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetPercentCompleteOk() (*int32, bool) {
	if o == nil || IsNil(o.PercentComplete) {
		return nil, false
	}
	return o.PercentComplete, true
}

// HasPercentComplete returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasPercentComplete() bool {
	if o != nil && !IsNil(o.PercentComplete) {
		return true
	}

	return false
}

// SetPercentComplete gets a reference to the given int32 and assigns it to the PercentComplete field.
func (o *GetEvents200ResponseDataInner) SetPercentComplete(v int32) {
	o.PercentComplete = &v
}

// GetRate returns the Rate field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetRate() string {
	if o == nil || IsNil(o.Rate) {
		var ret string
		return ret
	}
	return *o.Rate
}

// GetRateOk returns a tuple with the Rate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetRateOk() (*string, bool) {
	if o == nil || IsNil(o.Rate) {
		return nil, false
	}
	return o.Rate, true
}

// HasRate returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasRate() bool {
	if o != nil && !IsNil(o.Rate) {
		return true
	}

	return false
}

// SetRate gets a reference to the given string and assigns it to the Rate field.
func (o *GetEvents200ResponseDataInner) SetRate(v string) {
	o.Rate = &v
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetRead() bool {
	if o == nil || IsNil(o.Read) {
		var ret bool
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetReadOk() (*bool, bool) {
	if o == nil || IsNil(o.Read) {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasRead() bool {
	if o != nil && !IsNil(o.Read) {
		return true
	}

	return false
}

// SetRead gets a reference to the given bool and assigns it to the Read field.
func (o *GetEvents200ResponseDataInner) SetRead(v bool) {
	o.Read = &v
}

// GetSecondaryEntity returns the SecondaryEntity field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetSecondaryEntity() GetEvents200ResponseDataInnerSecondaryEntity {
	if o == nil || IsNil(o.SecondaryEntity) {
		var ret GetEvents200ResponseDataInnerSecondaryEntity
		return ret
	}
	return *o.SecondaryEntity
}

// GetSecondaryEntityOk returns a tuple with the SecondaryEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetSecondaryEntityOk() (*GetEvents200ResponseDataInnerSecondaryEntity, bool) {
	if o == nil || IsNil(o.SecondaryEntity) {
		return nil, false
	}
	return o.SecondaryEntity, true
}

// HasSecondaryEntity returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasSecondaryEntity() bool {
	if o != nil && !IsNil(o.SecondaryEntity) {
		return true
	}

	return false
}

// SetSecondaryEntity gets a reference to the given GetEvents200ResponseDataInnerSecondaryEntity and assigns it to the SecondaryEntity field.
func (o *GetEvents200ResponseDataInner) SetSecondaryEntity(v GetEvents200ResponseDataInnerSecondaryEntity) {
	o.SecondaryEntity = &v
}

// GetSeen returns the Seen field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetSeen() bool {
	if o == nil || IsNil(o.Seen) {
		var ret bool
		return ret
	}
	return *o.Seen
}

// GetSeenOk returns a tuple with the Seen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetSeenOk() (*bool, bool) {
	if o == nil || IsNil(o.Seen) {
		return nil, false
	}
	return o.Seen, true
}

// HasSeen returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasSeen() bool {
	if o != nil && !IsNil(o.Seen) {
		return true
	}

	return false
}

// SetSeen gets a reference to the given bool and assigns it to the Seen field.
func (o *GetEvents200ResponseDataInner) SetSeen(v bool) {
	o.Seen = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetEvents200ResponseDataInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetEvents200ResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetEvents200ResponseDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetTimeRemaining returns the TimeRemaining field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetEvents200ResponseDataInner) GetTimeRemaining() string {
	if o == nil || IsNil(o.TimeRemaining.Get()) {
		var ret string
		return ret
	}
	return *o.TimeRemaining.Get()
}

// GetTimeRemainingOk returns a tuple with the TimeRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEvents200ResponseDataInner) GetTimeRemainingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeRemaining.Get(), o.TimeRemaining.IsSet()
}

// HasTimeRemaining returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasTimeRemaining() bool {
	if o != nil && o.TimeRemaining.IsSet() {
		return true
	}

	return false
}

// SetTimeRemaining gets a reference to the given NullableString and assigns it to the TimeRemaining field.
func (o *GetEvents200ResponseDataInner) SetTimeRemaining(v string) {
	o.TimeRemaining.Set(&v)
}
// SetTimeRemainingNil sets the value for TimeRemaining to be an explicit nil
func (o *GetEvents200ResponseDataInner) SetTimeRemainingNil() {
	o.TimeRemaining.Set(nil)
}

// UnsetTimeRemaining ensures that no value is present for TimeRemaining, not even an explicit nil
func (o *GetEvents200ResponseDataInner) UnsetTimeRemaining() {
	o.TimeRemaining.Unset()
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GetEvents200ResponseDataInner) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GetEvents200ResponseDataInner) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *GetEvents200ResponseDataInner) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *GetEvents200ResponseDataInner) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *GetEvents200ResponseDataInner) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *GetEvents200ResponseDataInner) UnsetUsername() {
	o.Username.Unset()
}

func (o GetEvents200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEvents200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !IsNil(o.Entity) {
		toSerialize["entity"] = o.Entity
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Message.IsSet() {
		toSerialize["message"] = o.Message.Get()
	}
	if !IsNil(o.PercentComplete) {
		toSerialize["percent_complete"] = o.PercentComplete
	}
	if !IsNil(o.Rate) {
		toSerialize["rate"] = o.Rate
	}
	if !IsNil(o.Read) {
		toSerialize["read"] = o.Read
	}
	if !IsNil(o.SecondaryEntity) {
		toSerialize["secondary_entity"] = o.SecondaryEntity
	}
	if !IsNil(o.Seen) {
		toSerialize["seen"] = o.Seen
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.TimeRemaining.IsSet() {
		toSerialize["time_remaining"] = o.TimeRemaining.Get()
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	return toSerialize, nil
}

type NullableGetEvents200ResponseDataInner struct {
	value *GetEvents200ResponseDataInner
	isSet bool
}

func (v NullableGetEvents200ResponseDataInner) Get() *GetEvents200ResponseDataInner {
	return v.value
}

func (v *NullableGetEvents200ResponseDataInner) Set(val *GetEvents200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEvents200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEvents200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEvents200ResponseDataInner(val *GetEvents200ResponseDataInner) *NullableGetEvents200ResponseDataInner {
	return &NullableGetEvents200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetEvents200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEvents200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


