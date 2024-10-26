# PostDatabasesMysqlInstanceBackupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | The label for the Database snapshot backup.  - Can only contain ASCII letters, numbers, and underscores (&#x60;_&#x60;). - Must be unique among other backup labels for this Database. | 
**Target** | Pointer to **string** | The Database cluster target. If the Database is a high availability cluster, choosing &#x60;secondary&#x60; creates a snapshot backup of a replica node. | [optional] [default to "primary"]

## Methods

### NewPostDatabasesMysqlInstanceBackupRequest

`func NewPostDatabasesMysqlInstanceBackupRequest(label string, ) *PostDatabasesMysqlInstanceBackupRequest`

NewPostDatabasesMysqlInstanceBackupRequest instantiates a new PostDatabasesMysqlInstanceBackupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDatabasesMysqlInstanceBackupRequestWithDefaults

`func NewPostDatabasesMysqlInstanceBackupRequestWithDefaults() *PostDatabasesMysqlInstanceBackupRequest`

NewPostDatabasesMysqlInstanceBackupRequestWithDefaults instantiates a new PostDatabasesMysqlInstanceBackupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PostDatabasesMysqlInstanceBackupRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostDatabasesMysqlInstanceBackupRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostDatabasesMysqlInstanceBackupRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTarget

`func (o *PostDatabasesMysqlInstanceBackupRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PostDatabasesMysqlInstanceBackupRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PostDatabasesMysqlInstanceBackupRequest) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PostDatabasesMysqlInstanceBackupRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


