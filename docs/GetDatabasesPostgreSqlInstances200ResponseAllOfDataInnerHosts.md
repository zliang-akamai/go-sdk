# GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Primary** | Pointer to **NullableString** | The primary host for the Managed Database. | [optional] 
**Secondary** | Pointer to **NullableString** | The secondary/private network host for the Managed Database.  A private network host and a private IP can only be used to access a Database Cluster from Linodes in the same data center and will not incur transfer costs.  __Note__. The secondary hostname is publicly viewable and accessible. | [optional] 

## Methods

### NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts

`func NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts() *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts`

NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts instantiates a new GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHostsWithDefaults

`func NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHostsWithDefaults() *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts`

NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHostsWithDefaults instantiates a new GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimary

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) GetPrimary() string`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) GetPrimaryOk() (*string, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) SetPrimary(v string)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### SetPrimaryNil

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) SetPrimaryNil(b bool)`

 SetPrimaryNil sets the value for Primary to be an explicit nil

### UnsetPrimary
`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) UnsetPrimary()`

UnsetPrimary ensures that no value is present for Primary, not even an explicit nil
### GetSecondary

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) GetSecondary() string`

GetSecondary returns the Secondary field if non-nil, zero value otherwise.

### GetSecondaryOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) GetSecondaryOk() (*string, bool)`

GetSecondaryOk returns a tuple with the Secondary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondary

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) SetSecondary(v string)`

SetSecondary sets Secondary field to given value.

### HasSecondary

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) HasSecondary() bool`

HasSecondary returns a boolean if a field has been set.

### SetSecondaryNil

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) SetSecondaryNil(b bool)`

 SetSecondaryNil sets the value for Secondary to be an explicit nil

### UnsetSecondary
`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts) UnsetSecondary()`

UnsetSecondary ensures that no value is present for Secondary, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


