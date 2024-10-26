# PostResizeLinodeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAutoDiskResize** | Pointer to **bool** | Automatically resize disks when resizing a Linode. When resizing down to a smaller plan your Linode&#39;s data must fit within the smaller disk size. | [optional] [default to true]
**MigrationType** | Pointer to **string** | Type of migration used in moving to a new host or Linode type.  &#x60;warm&#x60;: the Linode will not power down until the migration is complete. Warm migrations are not available for DC migrations.  &#x60;cold&#x60;: the Linode will be powered down and migrated. When the migration is complete, the Linode will be powered on. | [optional] [default to "cold"]
**Type** | **string** | The ID representing the Linode Type. | 

## Methods

### NewPostResizeLinodeInstanceRequest

`func NewPostResizeLinodeInstanceRequest(type_ string, ) *PostResizeLinodeInstanceRequest`

NewPostResizeLinodeInstanceRequest instantiates a new PostResizeLinodeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostResizeLinodeInstanceRequestWithDefaults

`func NewPostResizeLinodeInstanceRequestWithDefaults() *PostResizeLinodeInstanceRequest`

NewPostResizeLinodeInstanceRequestWithDefaults instantiates a new PostResizeLinodeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAutoDiskResize

`func (o *PostResizeLinodeInstanceRequest) GetAllowAutoDiskResize() bool`

GetAllowAutoDiskResize returns the AllowAutoDiskResize field if non-nil, zero value otherwise.

### GetAllowAutoDiskResizeOk

`func (o *PostResizeLinodeInstanceRequest) GetAllowAutoDiskResizeOk() (*bool, bool)`

GetAllowAutoDiskResizeOk returns a tuple with the AllowAutoDiskResize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAutoDiskResize

`func (o *PostResizeLinodeInstanceRequest) SetAllowAutoDiskResize(v bool)`

SetAllowAutoDiskResize sets AllowAutoDiskResize field to given value.

### HasAllowAutoDiskResize

`func (o *PostResizeLinodeInstanceRequest) HasAllowAutoDiskResize() bool`

HasAllowAutoDiskResize returns a boolean if a field has been set.

### GetMigrationType

`func (o *PostResizeLinodeInstanceRequest) GetMigrationType() string`

GetMigrationType returns the MigrationType field if non-nil, zero value otherwise.

### GetMigrationTypeOk

`func (o *PostResizeLinodeInstanceRequest) GetMigrationTypeOk() (*string, bool)`

GetMigrationTypeOk returns a tuple with the MigrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrationType

`func (o *PostResizeLinodeInstanceRequest) SetMigrationType(v string)`

SetMigrationType sets MigrationType field to given value.

### HasMigrationType

`func (o *PostResizeLinodeInstanceRequest) HasMigrationType() bool`

HasMigrationType returns a boolean if a field has been set.

### GetType

`func (o *PostResizeLinodeInstanceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostResizeLinodeInstanceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostResizeLinodeInstanceRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


