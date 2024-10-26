# GetUserGrants200ResponseDatabaseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the entity this grant applies to. | [optional] 
**Label** | Pointer to **string** | The current label of the entity this grant applies to, for display purposes. | [optional] [readonly] 
**Permissions** | Pointer to **NullableString** | The level of access this User has to this entity.  If null, this User has no access. | [optional] 

## Methods

### NewGetUserGrants200ResponseDatabaseInner

`func NewGetUserGrants200ResponseDatabaseInner() *GetUserGrants200ResponseDatabaseInner`

NewGetUserGrants200ResponseDatabaseInner instantiates a new GetUserGrants200ResponseDatabaseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserGrants200ResponseDatabaseInnerWithDefaults

`func NewGetUserGrants200ResponseDatabaseInnerWithDefaults() *GetUserGrants200ResponseDatabaseInner`

NewGetUserGrants200ResponseDatabaseInnerWithDefaults instantiates a new GetUserGrants200ResponseDatabaseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetUserGrants200ResponseDatabaseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetUserGrants200ResponseDatabaseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetUserGrants200ResponseDatabaseInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetUserGrants200ResponseDatabaseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetUserGrants200ResponseDatabaseInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetUserGrants200ResponseDatabaseInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetUserGrants200ResponseDatabaseInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetUserGrants200ResponseDatabaseInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPermissions

`func (o *GetUserGrants200ResponseDatabaseInner) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetUserGrants200ResponseDatabaseInner) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetUserGrants200ResponseDatabaseInner) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetUserGrants200ResponseDatabaseInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *GetUserGrants200ResponseDatabaseInner) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *GetUserGrants200ResponseDatabaseInner) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


