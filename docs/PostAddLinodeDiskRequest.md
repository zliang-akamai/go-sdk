# PostAddLinodeDiskRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizedKeys** | Pointer to **[]string** | A list of public SSH keys that will be automatically appended to the root user&#39;s &#x60;~/.ssh/authorized_keys&#x60; file when deploying from an Image. | [optional] 
**AuthorizedUsers** | Pointer to **[]string** | A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the root users &#x60;~/.ssh/authorized_keys&#x60; file automatically when deploying from an Image. | [optional] 
**Filesystem** | Pointer to **string** | The file system of the disk. This can be &#x60;raw&#x60;, which indicates no file system, just a raw binary stream, &#x60;swap&#x60; for a Linux swap area, &#x60;ext3&#x60; or &#x60;ext4&#x60; for the ext3 of ext4 journaling file systems for Linux, respectively, or &#x60;initrd&#x60; for uncompressed initrd, ext2 with a maximum size of 32 MB. | [optional] 
**Image** | Pointer to **string** | An Image ID to deploy the Linode Disk from.  Run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation with authentication to view all available Images. Official Linode Images start with &#x60;linode/&#x60;, while your Account&#39;s Images start with &#x60;private/&#x60;. Creating a disk from a Private Image requires &#x60;read_only&#x60; or &#x60;read_write&#x60; permissions for that Image. Run the [Update a user&#39;s grants](https://techdocs.akamai.com/linode-api/reference/put-user-grants) operation to adjust permissions for an Account Image. | [optional] 
**Label** | Pointer to **string** | The name of the disk. This is for display purposes only. | [optional] 
**RootPass** | Pointer to **string** | This sets the root user&#39;s password on a newly created Linode Disk when deploying from an Image.  - __Required__ when creating a Linode Disk from an Image, including when using a StackScript.  - Must meet a password strength score requirement that is calculated internally by the API. If the strength requirement is not met, you will receive a &#x60;Password does not meet strength requirement&#x60; error. | [optional] 
**Size** | **int32** | The size of the Disk in MB.  Images require a minimum size. Run the [Get an image](https://techdocs.akamai.com/linode-api/reference/get-image) operation to view its size. | 
**StackscriptData** | Pointer to **map[string]interface{}** | This field is required only if the StackScript being deployed requires input data from the User for successful completion. See [User Defined Fields (UDFs)](https://www.linode.com/docs/products/tools/stackscripts/guides/write-a-custom-script/#declare-user-defined-fields-udfs) for more details.  This field is required to be valid JSON.  Total length cannot exceed 65,535 characters. | [optional] 
**StackscriptId** | Pointer to **int32** | A StackScript ID that will cause the referenced StackScript to be run during deployment of this Linode. A compatible &#x60;image&#x60; is required to use a StackScript. To get a list of available StackScript and their permitted Images, run [List StackScripts](https://techdocs.akamai.com/linode-api/reference/get-stack-scripts). This field cannot be used when deploying from a Backup or a Private Image. | [optional] 

## Methods

### NewPostAddLinodeDiskRequest

`func NewPostAddLinodeDiskRequest(size int32, ) *PostAddLinodeDiskRequest`

NewPostAddLinodeDiskRequest instantiates a new PostAddLinodeDiskRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAddLinodeDiskRequestWithDefaults

`func NewPostAddLinodeDiskRequestWithDefaults() *PostAddLinodeDiskRequest`

NewPostAddLinodeDiskRequestWithDefaults instantiates a new PostAddLinodeDiskRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedKeys

`func (o *PostAddLinodeDiskRequest) GetAuthorizedKeys() []string`

GetAuthorizedKeys returns the AuthorizedKeys field if non-nil, zero value otherwise.

### GetAuthorizedKeysOk

`func (o *PostAddLinodeDiskRequest) GetAuthorizedKeysOk() (*[]string, bool)`

GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeys

`func (o *PostAddLinodeDiskRequest) SetAuthorizedKeys(v []string)`

SetAuthorizedKeys sets AuthorizedKeys field to given value.

### HasAuthorizedKeys

`func (o *PostAddLinodeDiskRequest) HasAuthorizedKeys() bool`

HasAuthorizedKeys returns a boolean if a field has been set.

### GetAuthorizedUsers

`func (o *PostAddLinodeDiskRequest) GetAuthorizedUsers() []string`

GetAuthorizedUsers returns the AuthorizedUsers field if non-nil, zero value otherwise.

### GetAuthorizedUsersOk

`func (o *PostAddLinodeDiskRequest) GetAuthorizedUsersOk() (*[]string, bool)`

GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUsers

`func (o *PostAddLinodeDiskRequest) SetAuthorizedUsers(v []string)`

SetAuthorizedUsers sets AuthorizedUsers field to given value.

### HasAuthorizedUsers

`func (o *PostAddLinodeDiskRequest) HasAuthorizedUsers() bool`

HasAuthorizedUsers returns a boolean if a field has been set.

### GetFilesystem

`func (o *PostAddLinodeDiskRequest) GetFilesystem() string`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *PostAddLinodeDiskRequest) GetFilesystemOk() (*string, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *PostAddLinodeDiskRequest) SetFilesystem(v string)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *PostAddLinodeDiskRequest) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetImage

`func (o *PostAddLinodeDiskRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PostAddLinodeDiskRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PostAddLinodeDiskRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PostAddLinodeDiskRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLabel

`func (o *PostAddLinodeDiskRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostAddLinodeDiskRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostAddLinodeDiskRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostAddLinodeDiskRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRootPass

`func (o *PostAddLinodeDiskRequest) GetRootPass() string`

GetRootPass returns the RootPass field if non-nil, zero value otherwise.

### GetRootPassOk

`func (o *PostAddLinodeDiskRequest) GetRootPassOk() (*string, bool)`

GetRootPassOk returns a tuple with the RootPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPass

`func (o *PostAddLinodeDiskRequest) SetRootPass(v string)`

SetRootPass sets RootPass field to given value.

### HasRootPass

`func (o *PostAddLinodeDiskRequest) HasRootPass() bool`

HasRootPass returns a boolean if a field has been set.

### GetSize

`func (o *PostAddLinodeDiskRequest) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PostAddLinodeDiskRequest) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PostAddLinodeDiskRequest) SetSize(v int32)`

SetSize sets Size field to given value.


### GetStackscriptData

`func (o *PostAddLinodeDiskRequest) GetStackscriptData() map[string]interface{}`

GetStackscriptData returns the StackscriptData field if non-nil, zero value otherwise.

### GetStackscriptDataOk

`func (o *PostAddLinodeDiskRequest) GetStackscriptDataOk() (*map[string]interface{}, bool)`

GetStackscriptDataOk returns a tuple with the StackscriptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptData

`func (o *PostAddLinodeDiskRequest) SetStackscriptData(v map[string]interface{})`

SetStackscriptData sets StackscriptData field to given value.

### HasStackscriptData

`func (o *PostAddLinodeDiskRequest) HasStackscriptData() bool`

HasStackscriptData returns a boolean if a field has been set.

### GetStackscriptId

`func (o *PostAddLinodeDiskRequest) GetStackscriptId() int32`

GetStackscriptId returns the StackscriptId field if non-nil, zero value otherwise.

### GetStackscriptIdOk

`func (o *PostAddLinodeDiskRequest) GetStackscriptIdOk() (*int32, bool)`

GetStackscriptIdOk returns a tuple with the StackscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptId

`func (o *PostAddLinodeDiskRequest) SetStackscriptId(v int32)`

SetStackscriptId sets StackscriptId field to given value.

### HasStackscriptId

`func (o *PostAddLinodeDiskRequest) HasStackscriptId() bool`

HasStackscriptId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


