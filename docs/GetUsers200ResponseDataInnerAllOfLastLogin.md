# GetUsers200ResponseDataInnerAllOfLastLogin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginDatetime** | Pointer to **time.Time** | The date and time of this User&#39;s most recent login attempt. | [optional] [readonly] 
**Status** | Pointer to **string** | The result of the most recent login attempt for this User. | [optional] [readonly] 

## Methods

### NewGetUsers200ResponseDataInnerAllOfLastLogin

`func NewGetUsers200ResponseDataInnerAllOfLastLogin() *GetUsers200ResponseDataInnerAllOfLastLogin`

NewGetUsers200ResponseDataInnerAllOfLastLogin instantiates a new GetUsers200ResponseDataInnerAllOfLastLogin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUsers200ResponseDataInnerAllOfLastLoginWithDefaults

`func NewGetUsers200ResponseDataInnerAllOfLastLoginWithDefaults() *GetUsers200ResponseDataInnerAllOfLastLogin`

NewGetUsers200ResponseDataInnerAllOfLastLoginWithDefaults instantiates a new GetUsers200ResponseDataInnerAllOfLastLogin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginDatetime

`func (o *GetUsers200ResponseDataInnerAllOfLastLogin) GetLoginDatetime() time.Time`

GetLoginDatetime returns the LoginDatetime field if non-nil, zero value otherwise.

### GetLoginDatetimeOk

`func (o *GetUsers200ResponseDataInnerAllOfLastLogin) GetLoginDatetimeOk() (*time.Time, bool)`

GetLoginDatetimeOk returns a tuple with the LoginDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginDatetime

`func (o *GetUsers200ResponseDataInnerAllOfLastLogin) SetLoginDatetime(v time.Time)`

SetLoginDatetime sets LoginDatetime field to given value.

### HasLoginDatetime

`func (o *GetUsers200ResponseDataInnerAllOfLastLogin) HasLoginDatetime() bool`

HasLoginDatetime returns a boolean if a field has been set.

### GetStatus

`func (o *GetUsers200ResponseDataInnerAllOfLastLogin) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetUsers200ResponseDataInnerAllOfLastLogin) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetUsers200ResponseDataInnerAllOfLastLogin) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetUsers200ResponseDataInnerAllOfLastLogin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


