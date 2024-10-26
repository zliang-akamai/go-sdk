# GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | A time value given in a combined date and time format that represents when the database backup was created. | [optional] 
**Id** | Pointer to **int32** | The ID of the database backup object. | [optional] 
**Label** | Pointer to **string** | The database backup&#39;s label, for display purposes only.  Must include only ASCII letters or numbers. | [optional] 
**Type** | Pointer to **string** | The type of database backup, determined by how the backup was created. | [optional] 

## Methods

### NewGetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner

`func NewGetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner() *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner`

NewGetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner instantiates a new GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasesMysqlInstanceBackups200ResponseAllOfDataInnerWithDefaults

`func NewGetDatabasesMysqlInstanceBackups200ResponseAllOfDataInnerWithDefaults() *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner`

NewGetDatabasesMysqlInstanceBackups200ResponseAllOfDataInnerWithDefaults instantiates a new GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDatabasesMysqlInstanceBackups200ResponseAllOfDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


