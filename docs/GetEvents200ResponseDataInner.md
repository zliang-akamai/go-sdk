# GetEvents200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** | The action that caused this Event. New actions may be added in the future. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | When this Event was created. | [optional] [readonly] 
**Duration** | Pointer to **float32** | The total duration in seconds that it takes for the Event to complete. | [optional] [readonly] 
**Entity** | Pointer to [**GetEvents200ResponseDataInnerEntity**](GetEvents200ResponseDataInnerEntity.md) |  | [optional] 
**Id** | Pointer to **int32** | The unique ID of this Event. | [optional] [readonly] 
**Message** | Pointer to **NullableString** | Provides additional information about the event. Additional information may include, but is not limited to, a more detailed representation of events which can help diagnose non-obvious failures. | [optional] 
**PercentComplete** | Pointer to **int32** | A percentage estimating the amount of time remaining for an Event. Returns &#x60;null&#x60; for notification events. | [optional] [readonly] 
**Rate** | Pointer to **string** | The rate of completion of the Event. Only some Events will return rate; for example, migration and resize Events. | [optional] [readonly] 
**Read** | Pointer to **bool** | If this Event has been read. | [optional] [readonly] 
**SecondaryEntity** | Pointer to [**GetEvents200ResponseDataInnerSecondaryEntity**](GetEvents200ResponseDataInnerSecondaryEntity.md) |  | [optional] 
**Seen** | Pointer to **bool** | If this Event has been seen. | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of this Event. | [optional] [readonly] 
**TimeRemaining** | Pointer to **NullableString** | The estimated time remaining until the completion of this Event. This value is only returned for some in-progress migration events. For all other in-progress events, the &#x60;percent_complete&#x60; attribute will indicate about how much more work is to be done. | [optional] [readonly] 
**Username** | Pointer to **NullableString** | The username of the User who caused the Event. | [optional] [readonly] 

## Methods

### NewGetEvents200ResponseDataInner

`func NewGetEvents200ResponseDataInner() *GetEvents200ResponseDataInner`

NewGetEvents200ResponseDataInner instantiates a new GetEvents200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseDataInnerWithDefaults

`func NewGetEvents200ResponseDataInnerWithDefaults() *GetEvents200ResponseDataInner`

NewGetEvents200ResponseDataInnerWithDefaults instantiates a new GetEvents200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *GetEvents200ResponseDataInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetEvents200ResponseDataInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetEvents200ResponseDataInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *GetEvents200ResponseDataInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreated

`func (o *GetEvents200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetEvents200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetEvents200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetEvents200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDuration

`func (o *GetEvents200ResponseDataInner) GetDuration() float32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetEvents200ResponseDataInner) GetDurationOk() (*float32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetEvents200ResponseDataInner) SetDuration(v float32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetEvents200ResponseDataInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEntity

`func (o *GetEvents200ResponseDataInner) GetEntity() GetEvents200ResponseDataInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GetEvents200ResponseDataInner) GetEntityOk() (*GetEvents200ResponseDataInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GetEvents200ResponseDataInner) SetEntity(v GetEvents200ResponseDataInnerEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GetEvents200ResponseDataInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetId

`func (o *GetEvents200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvents200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvents200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetEvents200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *GetEvents200ResponseDataInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetEvents200ResponseDataInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetEvents200ResponseDataInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetEvents200ResponseDataInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GetEvents200ResponseDataInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GetEvents200ResponseDataInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetPercentComplete

`func (o *GetEvents200ResponseDataInner) GetPercentComplete() int32`

GetPercentComplete returns the PercentComplete field if non-nil, zero value otherwise.

### GetPercentCompleteOk

`func (o *GetEvents200ResponseDataInner) GetPercentCompleteOk() (*int32, bool)`

GetPercentCompleteOk returns a tuple with the PercentComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentComplete

`func (o *GetEvents200ResponseDataInner) SetPercentComplete(v int32)`

SetPercentComplete sets PercentComplete field to given value.

### HasPercentComplete

`func (o *GetEvents200ResponseDataInner) HasPercentComplete() bool`

HasPercentComplete returns a boolean if a field has been set.

### GetRate

`func (o *GetEvents200ResponseDataInner) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *GetEvents200ResponseDataInner) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *GetEvents200ResponseDataInner) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *GetEvents200ResponseDataInner) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRead

`func (o *GetEvents200ResponseDataInner) GetRead() bool`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *GetEvents200ResponseDataInner) GetReadOk() (*bool, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *GetEvents200ResponseDataInner) SetRead(v bool)`

SetRead sets Read field to given value.

### HasRead

`func (o *GetEvents200ResponseDataInner) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetSecondaryEntity

`func (o *GetEvents200ResponseDataInner) GetSecondaryEntity() GetEvents200ResponseDataInnerSecondaryEntity`

GetSecondaryEntity returns the SecondaryEntity field if non-nil, zero value otherwise.

### GetSecondaryEntityOk

`func (o *GetEvents200ResponseDataInner) GetSecondaryEntityOk() (*GetEvents200ResponseDataInnerSecondaryEntity, bool)`

GetSecondaryEntityOk returns a tuple with the SecondaryEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEntity

`func (o *GetEvents200ResponseDataInner) SetSecondaryEntity(v GetEvents200ResponseDataInnerSecondaryEntity)`

SetSecondaryEntity sets SecondaryEntity field to given value.

### HasSecondaryEntity

`func (o *GetEvents200ResponseDataInner) HasSecondaryEntity() bool`

HasSecondaryEntity returns a boolean if a field has been set.

### GetSeen

`func (o *GetEvents200ResponseDataInner) GetSeen() bool`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *GetEvents200ResponseDataInner) GetSeenOk() (*bool, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *GetEvents200ResponseDataInner) SetSeen(v bool)`

SetSeen sets Seen field to given value.

### HasSeen

`func (o *GetEvents200ResponseDataInner) HasSeen() bool`

HasSeen returns a boolean if a field has been set.

### GetStatus

`func (o *GetEvents200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetEvents200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetEvents200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetEvents200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeRemaining

`func (o *GetEvents200ResponseDataInner) GetTimeRemaining() string`

GetTimeRemaining returns the TimeRemaining field if non-nil, zero value otherwise.

### GetTimeRemainingOk

`func (o *GetEvents200ResponseDataInner) GetTimeRemainingOk() (*string, bool)`

GetTimeRemainingOk returns a tuple with the TimeRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRemaining

`func (o *GetEvents200ResponseDataInner) SetTimeRemaining(v string)`

SetTimeRemaining sets TimeRemaining field to given value.

### HasTimeRemaining

`func (o *GetEvents200ResponseDataInner) HasTimeRemaining() bool`

HasTimeRemaining returns a boolean if a field has been set.

### SetTimeRemainingNil

`func (o *GetEvents200ResponseDataInner) SetTimeRemainingNil(b bool)`

 SetTimeRemainingNil sets the value for TimeRemaining to be an explicit nil

### UnsetTimeRemaining
`func (o *GetEvents200ResponseDataInner) UnsetTimeRemaining()`

UnsetTimeRemaining ensures that no value is present for TimeRemaining, not even an explicit nil
### GetUsername

`func (o *GetEvents200ResponseDataInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetEvents200ResponseDataInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetEvents200ResponseDataInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetEvents200ResponseDataInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *GetEvents200ResponseDataInner) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *GetEvents200ResponseDataInner) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


