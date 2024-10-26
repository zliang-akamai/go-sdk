# GetAccountLogins200ResponseDataInner

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

### NewGetAccountLogins200ResponseDataInner

`func NewGetAccountLogins200ResponseDataInner() *GetAccountLogins200ResponseDataInner`

NewGetAccountLogins200ResponseDataInner instantiates a new GetAccountLogins200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountLogins200ResponseDataInnerWithDefaults

`func NewGetAccountLogins200ResponseDataInnerWithDefaults() *GetAccountLogins200ResponseDataInner`

NewGetAccountLogins200ResponseDataInnerWithDefaults instantiates a new GetAccountLogins200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *GetAccountLogins200ResponseDataInner) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *GetAccountLogins200ResponseDataInner) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *GetAccountLogins200ResponseDataInner) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.

### HasDatetime

`func (o *GetAccountLogins200ResponseDataInner) HasDatetime() bool`

HasDatetime returns a boolean if a field has been set.

### GetId

`func (o *GetAccountLogins200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccountLogins200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccountLogins200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetAccountLogins200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *GetAccountLogins200ResponseDataInner) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetAccountLogins200ResponseDataInner) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetAccountLogins200ResponseDataInner) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetAccountLogins200ResponseDataInner) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetRestricted

`func (o *GetAccountLogins200ResponseDataInner) GetRestricted() bool`

GetRestricted returns the Restricted field if non-nil, zero value otherwise.

### GetRestrictedOk

`func (o *GetAccountLogins200ResponseDataInner) GetRestrictedOk() (*bool, bool)`

GetRestrictedOk returns a tuple with the Restricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestricted

`func (o *GetAccountLogins200ResponseDataInner) SetRestricted(v bool)`

SetRestricted sets Restricted field to given value.

### HasRestricted

`func (o *GetAccountLogins200ResponseDataInner) HasRestricted() bool`

HasRestricted returns a boolean if a field has been set.

### GetStatus

`func (o *GetAccountLogins200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAccountLogins200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAccountLogins200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAccountLogins200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsername

`func (o *GetAccountLogins200ResponseDataInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetAccountLogins200ResponseDataInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetAccountLogins200ResponseDataInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetAccountLogins200ResponseDataInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


