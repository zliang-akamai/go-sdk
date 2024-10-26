# GetEnrolledBetaPrograms200ResponseAllOfDataInner

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

### NewGetEnrolledBetaPrograms200ResponseAllOfDataInner

`func NewGetEnrolledBetaPrograms200ResponseAllOfDataInner() *GetEnrolledBetaPrograms200ResponseAllOfDataInner`

NewGetEnrolledBetaPrograms200ResponseAllOfDataInner instantiates a new GetEnrolledBetaPrograms200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnrolledBetaPrograms200ResponseAllOfDataInnerWithDefaults

`func NewGetEnrolledBetaPrograms200ResponseAllOfDataInnerWithDefaults() *GetEnrolledBetaPrograms200ResponseAllOfDataInner`

NewGetEnrolledBetaPrograms200ResponseAllOfDataInnerWithDefaults instantiates a new GetEnrolledBetaPrograms200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnded

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetEnded() time.Time`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetEndedOk() (*time.Time, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) SetEnded(v time.Time)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### SetEndedNil

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) SetEndedNil(b bool)`

 SetEndedNil sets the value for Ended to be an explicit nil

### UnsetEnded
`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) UnsetEnded()`

UnsetEnded ensures that no value is present for Ended, not even an explicit nil
### GetEnrolled

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetEnrolled() time.Time`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetEnrolledOk() (*time.Time, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) SetEnrolled(v time.Time)`

SetEnrolled sets Enrolled field to given value.

### HasEnrolled

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) HasEnrolled() bool`

HasEnrolled returns a boolean if a field has been set.

### GetId

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStarted

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetStarted() time.Time`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) GetStartedOk() (*time.Time, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) SetStarted(v time.Time)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *GetEnrolledBetaPrograms200ResponseAllOfDataInner) HasStarted() bool`

HasStarted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


