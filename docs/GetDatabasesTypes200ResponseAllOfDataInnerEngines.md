# GetDatabasesTypes200ResponseAllOfDataInnerEngines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mysql** | Pointer to [**[]GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner**](GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner.md) | Pricing details for MySQL Managed Databases. | [optional] 
**Postgresql** | Pointer to [**[]GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner**](GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner.md) | Pricing details for PostgreSQL Managed Databases. | [optional] 

## Methods

### NewGetDatabasesTypes200ResponseAllOfDataInnerEngines

`func NewGetDatabasesTypes200ResponseAllOfDataInnerEngines() *GetDatabasesTypes200ResponseAllOfDataInnerEngines`

NewGetDatabasesTypes200ResponseAllOfDataInnerEngines instantiates a new GetDatabasesTypes200ResponseAllOfDataInnerEngines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesWithDefaults

`func NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesWithDefaults() *GetDatabasesTypes200ResponseAllOfDataInnerEngines`

NewGetDatabasesTypes200ResponseAllOfDataInnerEnginesWithDefaults instantiates a new GetDatabasesTypes200ResponseAllOfDataInnerEngines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMysql

`func (o *GetDatabasesTypes200ResponseAllOfDataInnerEngines) GetMysql() []GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner`

GetMysql returns the Mysql field if non-nil, zero value otherwise.

### GetMysqlOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInnerEngines) GetMysqlOk() (*[]GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner, bool)`

GetMysqlOk returns a tuple with the Mysql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMysql

`func (o *GetDatabasesTypes200ResponseAllOfDataInnerEngines) SetMysql(v []GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner)`

SetMysql sets Mysql field to given value.

### HasMysql

`func (o *GetDatabasesTypes200ResponseAllOfDataInnerEngines) HasMysql() bool`

HasMysql returns a boolean if a field has been set.

### GetPostgresql

`func (o *GetDatabasesTypes200ResponseAllOfDataInnerEngines) GetPostgresql() []GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner`

GetPostgresql returns the Postgresql field if non-nil, zero value otherwise.

### GetPostgresqlOk

`func (o *GetDatabasesTypes200ResponseAllOfDataInnerEngines) GetPostgresqlOk() (*[]GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner, bool)`

GetPostgresqlOk returns a tuple with the Postgresql field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostgresql

`func (o *GetDatabasesTypes200ResponseAllOfDataInnerEngines) SetPostgresql(v []GetDatabasesTypes200ResponseAllOfDataInnerEnginesMysqlInner)`

SetPostgresql sets Postgresql field to given value.

### HasPostgresql

`func (o *GetDatabasesTypes200ResponseAllOfDataInnerEngines) HasPostgresql() bool`

HasPostgresql returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


