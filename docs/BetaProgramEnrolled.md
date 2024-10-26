# BetaProgramEnrolled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Additional details regarding the Beta Program. | [optional] [readonly] 
**Ended** | Pointer to **NullableTime** | The date-time that the Beta Program ended.  &#x60;null&#x60; indicates that the Beta Program is ongoing. | [optional] [readonly] 
**Enrolled** | Pointer to **time.Time** | The date-time of Account enrollment to the Beta Program. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier of the Beta Program. | [optional] 
**Label** | Pointer to **string** | The name of the Beta Program. | [optional] [readonly] 
**Started** | Pointer to **time.Time** | The start date-time of the Beta Program. | [optional] [readonly] 

## Methods

### NewBetaProgramEnrolled

`func NewBetaProgramEnrolled() *BetaProgramEnrolled`

NewBetaProgramEnrolled instantiates a new BetaProgramEnrolled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaProgramEnrolledWithDefaults

`func NewBetaProgramEnrolledWithDefaults() *BetaProgramEnrolled`

NewBetaProgramEnrolledWithDefaults instantiates a new BetaProgramEnrolled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *BetaProgramEnrolled) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BetaProgramEnrolled) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BetaProgramEnrolled) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BetaProgramEnrolled) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *BetaProgramEnrolled) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *BetaProgramEnrolled) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnded

`func (o *BetaProgramEnrolled) GetEnded() time.Time`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *BetaProgramEnrolled) GetEndedOk() (*time.Time, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *BetaProgramEnrolled) SetEnded(v time.Time)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *BetaProgramEnrolled) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### SetEndedNil

`func (o *BetaProgramEnrolled) SetEndedNil(b bool)`

 SetEndedNil sets the value for Ended to be an explicit nil

### UnsetEnded
`func (o *BetaProgramEnrolled) UnsetEnded()`

UnsetEnded ensures that no value is present for Ended, not even an explicit nil
### GetEnrolled

`func (o *BetaProgramEnrolled) GetEnrolled() time.Time`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *BetaProgramEnrolled) GetEnrolledOk() (*time.Time, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *BetaProgramEnrolled) SetEnrolled(v time.Time)`

SetEnrolled sets Enrolled field to given value.

### HasEnrolled

`func (o *BetaProgramEnrolled) HasEnrolled() bool`

HasEnrolled returns a boolean if a field has been set.

### GetId

`func (o *BetaProgramEnrolled) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BetaProgramEnrolled) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BetaProgramEnrolled) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BetaProgramEnrolled) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *BetaProgramEnrolled) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *BetaProgramEnrolled) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *BetaProgramEnrolled) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *BetaProgramEnrolled) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStarted

`func (o *BetaProgramEnrolled) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *BetaProgramEnrolled) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *BetaProgramEnrolled) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *BetaProgramEnrolled) HasStarted() bool`

HasStarted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


