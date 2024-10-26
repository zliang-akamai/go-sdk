# AddedGetUser200AllOfLastLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginDatetime** | Pointer to **time.Time** | The date and time of this User&#39;s most recent login attempt. | [optional] [readonly] 
**Status** | Pointer to **string** | The result of the most recent login attempt for this User. | [optional] [readonly] 

## Methods

### NewAddedGetUser200AllOfLastLogin

`func NewAddedGetUser200AllOfLastLogin() *AddedGetUser200AllOfLastLogin`

NewAddedGetUser200AllOfLastLogin instantiates a new AddedGetUser200AllOfLastLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedGetUser200AllOfLastLoginWithDefaults

`func NewAddedGetUser200AllOfLastLoginWithDefaults() *AddedGetUser200AllOfLastLogin`

NewAddedGetUser200AllOfLastLoginWithDefaults instantiates a new AddedGetUser200AllOfLastLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginDatetime

`func (o *AddedGetUser200AllOfLastLogin) GetLoginDatetime() time.Time`

GetLoginDatetime returns the LoginDatetime field if non-nil, zero value otherwise.

### GetLoginDatetimeOk

`func (o *AddedGetUser200AllOfLastLogin) GetLoginDatetimeOk() (*time.Time, bool)`

GetLoginDatetimeOk returns a tuple with the LoginDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDatetime

`func (o *AddedGetUser200AllOfLastLogin) SetLoginDatetime(v time.Time)`

SetLoginDatetime sets LoginDatetime field to given value.

### HasLoginDatetime

`func (o *AddedGetUser200AllOfLastLogin) HasLoginDatetime() bool`

HasLoginDatetime returns a boolean if a field has been set.

### GetStatus

`func (o *AddedGetUser200AllOfLastLogin) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddedGetUser200AllOfLastLogin) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddedGetUser200AllOfLastLogin) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddedGetUser200AllOfLastLogin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


