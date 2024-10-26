# Login

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | Pointer to **time.Time** | When the login was initiated. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID of this login object. | [optional] [readonly] 
**Ip** | Pointer to **string** | The remote IP address that requested the login. | [optional] [readonly] 
**Restricted** | Pointer to **bool** | True if the User that attempted the login was a restricted User, false otherwise. | [optional] [readonly] 
**Status** | Pointer to **string** | Whether the login attempt succeeded or failed. | [optional] [readonly] 
**Username** | Pointer to **string** | The username of the User that attempted the login. | [optional] [readonly] 

## Methods

### NewLogin

`func NewLogin() *Login`

NewLogin instantiates a new Login object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoginWithDefaults

`func NewLoginWithDefaults() *Login`

NewLoginWithDefaults instantiates a new Login object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *Login) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *Login) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *Login) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *Login) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetId

`func (o *Login) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Login) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Login) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Login) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *Login) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *Login) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *Login) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *Login) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetRestricted

`func (o *Login) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *Login) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *Login) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *Login) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetStatus

`func (o *Login) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Login) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Login) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Login) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *Login) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Login) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Login) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Login) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


