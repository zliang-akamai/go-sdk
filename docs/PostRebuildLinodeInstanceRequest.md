# PostRebuildLinodeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizedKeys** | Pointer to **[]string** | A list of public SSH keys that will be automatically appended to the root user&#39;s &#x60;~/.ssh/authorized_keys&#x60; file when deploying from an Image. | [optional] 
**AuthorizedUsers** | Pointer to **[]string** | A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the root users &#x60;~/.ssh/authorized_keys&#x60; file automatically when deploying from an Image. | [optional] 
**Booted** | Pointer to **bool** | This field defaults to &#x60;true&#x60; if the Linode is created with an Image or from a Backup. If it is deployed from an Image or a Backup and you wish it to remain &#x60;offline&#x60; after deployment, set this to &#x60;false&#x60;. | [optional] [default to true]
**DiskEncryption** | Pointer to **string** | Local disk encryption ensures that your data stored on Linodes is secured. Disk encryption protects against unauthorized data access by keeping the data encrypted if the disk is ever removed from the data center, decommissioned, or disposed of. The platform manages the encryption and decryption for you.  By default, encryption is &#x60;enabled&#x60; on all Linodes. If you opted out of encryption or if the Linode was created prior to local disk encryption support, you can encrypt your data using [Rebuild](https://techdocs.akamai.com/linode-api/reference/post-rebuild-linode-instance). | [optional] 
**Image** | **string** | An Image ID to deploy the Linode Disk from.  Run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation with authentication to view all available Images. Official Linode Images start with &#x60;linode/&#x60;, while your Account&#39;s Images start with &#x60;private/&#x60;. Creating a disk from a Private Image requires &#x60;read_only&#x60; or &#x60;read_write&#x60; permissions for that Image. Run the [Update a user&#39;s grants](https://techdocs.akamai.com/linode-api/reference/put-user-grants) operation to adjust permissions for an Account Image. | 
**Metadata** | Pointer to [**PostLinodeInstanceRequestAllOfMetadata**](PostLinodeInstanceRequestAllOfMetadata.md) |  | [optional] 
**RootPass** | **string** | This sets the root user&#39;s password on a newly created Linode Disk when deploying from an Image.  - __Required__ when creating a Linode Disk from an Image, including when using a StackScript.  - Must meet a password strength score requirement that is calculated internally by the API. If the strength requirement is not met, you will receive a &#x60;Password does not meet strength requirement&#x60; error. | 
**StackscriptData** | Pointer to **map[string]interface{}** | This field is required only if the StackScript being deployed requires input data from the User for successful completion. See [User Defined Fields (UDFs)](https://www.linode.com/docs/products/tools/stackscripts/guides/write-a-custom-script/#declare-user-defined-fields-udfs) for more details.  This field is required to be valid JSON.  Total length cannot exceed 65,535 characters. | [optional] 
**StackscriptId** | Pointer to **int32** | A StackScript ID that will cause the referenced StackScript to be run during deployment of this Linode. A compatible &#x60;image&#x60; is required to use a StackScript. To get a list of available StackScript and their permitted Images, run [List StackScripts](https://techdocs.akamai.com/linode-api/reference/get-stack-scripts). This field cannot be used when deploying from a Backup or a Private Image. | [optional] 
**Type** | Pointer to **string** | The ID of the [Linode type](https://techdocs.akamai.com/linode-api/reference/get-linode-types) to resize to with this request. | [optional] 

## Methods

### NewPostRebuildLinodeInstanceRequest

`func NewPostRebuildLinodeInstanceRequest(image string, rootPass string, ) *PostRebuildLinodeInstanceRequest`

NewPostRebuildLinodeInstanceRequest instantiates a new PostRebuildLinodeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRebuildLinodeInstanceRequestWithDefaults

`func NewPostRebuildLinodeInstanceRequestWithDefaults() *PostRebuildLinodeInstanceRequest`

NewPostRebuildLinodeInstanceRequestWithDefaults instantiates a new PostRebuildLinodeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedKeys

`func (o *PostRebuildLinodeInstanceRequest) GetAuthorizedKeys() []string`

GetAuthorizedKeys returns the AuthorizedKeys field if non-nil, zero value otherwise.

### GetAuthorizedKeysOk

`func (o *PostRebuildLinodeInstanceRequest) GetAuthorizedKeysOk() (*[]string, bool)`

GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeys

`func (o *PostRebuildLinodeInstanceRequest) SetAuthorizedKeys(v []string)`

SetAuthorizedKeys sets AuthorizedKeys field to given value.

### HasAuthorizedKeys

`func (o *PostRebuildLinodeInstanceRequest) HasAuthorizedKeys() bool`

HasAuthorizedKeys returns a boolean if a field has been set.

### GetAuthorizedUsers

`func (o *PostRebuildLinodeInstanceRequest) GetAuthorizedUsers() []string`

GetAuthorizedUsers returns the AuthorizedUsers field if non-nil, zero value otherwise.

### GetAuthorizedUsersOk

`func (o *PostRebuildLinodeInstanceRequest) GetAuthorizedUsersOk() (*[]string, bool)`

GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUsers

`func (o *PostRebuildLinodeInstanceRequest) SetAuthorizedUsers(v []string)`

SetAuthorizedUsers sets AuthorizedUsers field to given value.

### HasAuthorizedUsers

`func (o *PostRebuildLinodeInstanceRequest) HasAuthorizedUsers() bool`

HasAuthorizedUsers returns a boolean if a field has been set.

### GetBooted

`func (o *PostRebuildLinodeInstanceRequest) GetBooted() bool`

GetBooted returns the Booted field if non-nil, zero value otherwise.

### GetBootedOk

`func (o *PostRebuildLinodeInstanceRequest) GetBootedOk() (*bool, bool)`

GetBootedOk returns a tuple with the Booted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooted

`func (o *PostRebuildLinodeInstanceRequest) SetBooted(v bool)`

SetBooted sets Booted field to given value.

### HasBooted

`func (o *PostRebuildLinodeInstanceRequest) HasBooted() bool`

HasBooted returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *PostRebuildLinodeInstanceRequest) GetDiskEncryption() string`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *PostRebuildLinodeInstanceRequest) GetDiskEncryptionOk() (*string, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *PostRebuildLinodeInstanceRequest) SetDiskEncryption(v string)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *PostRebuildLinodeInstanceRequest) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetImage

`func (o *PostRebuildLinodeInstanceRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PostRebuildLinodeInstanceRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PostRebuildLinodeInstanceRequest) SetImage(v string)`

SetImage sets Image field to given value.


### GetMetadata

`func (o *PostRebuildLinodeInstanceRequest) GetMetadata() PostLinodeInstanceRequestAllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostRebuildLinodeInstanceRequest) GetMetadataOk() (*PostLinodeInstanceRequestAllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostRebuildLinodeInstanceRequest) SetMetadata(v PostLinodeInstanceRequestAllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PostRebuildLinodeInstanceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRootPass

`func (o *PostRebuildLinodeInstanceRequest) GetRootPass() string`

GetRootPass returns the RootPass field if non-nil, zero value otherwise.

### GetRootPassOk

`func (o *PostRebuildLinodeInstanceRequest) GetRootPassOk() (*string, bool)`

GetRootPassOk returns a tuple with the RootPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPass

`func (o *PostRebuildLinodeInstanceRequest) SetRootPass(v string)`

SetRootPass sets RootPass field to given value.


### GetStackscriptData

`func (o *PostRebuildLinodeInstanceRequest) GetStackscriptData() map[string]interface{}`

GetStackscriptData returns the StackscriptData field if non-nil, zero value otherwise.

### GetStackscriptDataOk

`func (o *PostRebuildLinodeInstanceRequest) GetStackscriptDataOk() (*map[string]interface{}, bool)`

GetStackscriptDataOk returns a tuple with the StackscriptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptData

`func (o *PostRebuildLinodeInstanceRequest) SetStackscriptData(v map[string]interface{})`

SetStackscriptData sets StackscriptData field to given value.

### HasStackscriptData

`func (o *PostRebuildLinodeInstanceRequest) HasStackscriptData() bool`

HasStackscriptData returns a boolean if a field has been set.

### GetStackscriptId

`func (o *PostRebuildLinodeInstanceRequest) GetStackscriptId() int32`

GetStackscriptId returns the StackscriptId field if non-nil, zero value otherwise.

### GetStackscriptIdOk

`func (o *PostRebuildLinodeInstanceRequest) GetStackscriptIdOk() (*int32, bool)`

GetStackscriptIdOk returns a tuple with the StackscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptId

`func (o *PostRebuildLinodeInstanceRequest) SetStackscriptId(v int32)`

SetStackscriptId sets StackscriptId field to given value.

### HasStackscriptId

`func (o *PostRebuildLinodeInstanceRequest) HasStackscriptId() bool`

HasStackscriptId returns a boolean if a field has been set.

### GetType

`func (o *PostRebuildLinodeInstanceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostRebuildLinodeInstanceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostRebuildLinodeInstanceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostRebuildLinodeInstanceRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


