# GetManagedLinodeSettings200ResponseDataInnerSsh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to **bool** | If true, Linode special forces may access this Linode over ssh to respond to Issues. | [optional] [default to true]
**Ip** | Pointer to **string** | The IP Linode special forces should use to access this Linode when responding to an Issue.  By default, any of a Linode&#39;s IP addresses can be used for incident response access. | [optional] [default to "any"]
**Port** | Pointer to **NullableInt32** | The port Linode special forces should use to access this Linode over ssh to respond to an Issue.  The default &#x60;null&#x60; value corresponds to port 22. | [optional] 
**User** | Pointer to **NullableString** | The specific user, if any, Linode&#39;s special forces should use when accessing this Linode to respond to an issue.  The default &#x60;null&#x60; value corresponds to the root user. | [optional] 

## Methods

### NewGetManagedLinodeSettings200ResponseDataInnerSsh

`func NewGetManagedLinodeSettings200ResponseDataInnerSsh() *GetManagedLinodeSettings200ResponseDataInnerSsh`

NewGetManagedLinodeSettings200ResponseDataInnerSsh instantiates a new GetManagedLinodeSettings200ResponseDataInnerSsh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedLinodeSettings200ResponseDataInnerSshWithDefaults

`func NewGetManagedLinodeSettings200ResponseDataInnerSshWithDefaults() *GetManagedLinodeSettings200ResponseDataInnerSsh`

NewGetManagedLinodeSettings200ResponseDataInnerSshWithDefaults instantiates a new GetManagedLinodeSettings200ResponseDataInnerSsh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) GetAccess() bool`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) GetAccessOk() (*bool, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) SetAccess(v bool)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetIp

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetPort

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetUser

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *GetManagedLinodeSettings200ResponseDataInnerSsh) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


