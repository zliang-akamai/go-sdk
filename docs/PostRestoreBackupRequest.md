# PostRestoreBackupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinodeId** | **int32** | The ID of the Linode to restore a Backup to. | 
**Overwrite** | Pointer to **bool** | If True, deletes all Disks and Configs on the target Linode before restoring.  If False, and the Disk image size is larger than the available space on the Linode, an error message indicating insufficient space is returned. | [optional] 

## Methods

### NewPostRestoreBackupRequest

`func NewPostRestoreBackupRequest(linodeId int32, ) *PostRestoreBackupRequest`

NewPostRestoreBackupRequest instantiates a new PostRestoreBackupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRestoreBackupRequestWithDefaults

`func NewPostRestoreBackupRequestWithDefaults() *PostRestoreBackupRequest`

NewPostRestoreBackupRequestWithDefaults instantiates a new PostRestoreBackupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinodeId

`func (o *PostRestoreBackupRequest) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PostRestoreBackupRequest) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PostRestoreBackupRequest) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.


### GetOverwrite

`func (o *PostRestoreBackupRequest) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *PostRestoreBackupRequest) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *PostRestoreBackupRequest) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *PostRestoreBackupRequest) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


