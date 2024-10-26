# PostLinodeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizedKeys** | Pointer to **[]string** | A list of public SSH keys that will be automatically appended to the root user&#39;s &#x60;~/.ssh/authorized_keys&#x60; file when deploying from an Image. | [optional] 
**AuthorizedUsers** | Pointer to **[]string** | A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the root users &#x60;~/.ssh/authorized_keys&#x60; file automatically when deploying from an Image. | [optional] 
**Booted** | Pointer to **bool** | This field defaults to &#x60;true&#x60; if the Linode is created with an Image or from a Backup. If it is deployed from an Image or a Backup and you wish it to remain &#x60;offline&#x60; after deployment, set this to &#x60;false&#x60;. | [optional] [default to true]
**DiskEncryption** | Pointer to **string** | Local disk encryption ensures that your data stored on Linodes is secured. Disk encryption protects against unauthorized data access by keeping the data encrypted if the disk is ever removed from the data center, decommissioned, or disposed of. The platform manages the encryption and decryption for you.  By default, encryption is &#x60;enabled&#x60; on all Linodes. If you opted out of encryption or if the Linode was created prior to local disk encryption support, you can encrypt your data using [Rebuild](https://techdocs.akamai.com/linode-api/reference/post-rebuild-linode-instance). | [optional] 
**Image** | Pointer to **string** | An Image ID to deploy the Linode Disk from.  Run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation with authentication to view all available Images. Official Linode Images start with &#x60;linode/&#x60;, while your Account&#39;s Images start with &#x60;private/&#x60;. Creating a disk from a Private Image requires &#x60;read_only&#x60; or &#x60;read_write&#x60; permissions for that Image. Run the [Update a user&#39;s grants](https://techdocs.akamai.com/linode-api/reference/put-user-grants) operation to adjust permissions for an Account Image. | [optional] 
**Metadata** | Pointer to [**PostLinodeInstanceRequestAllOfMetadata**](PostLinodeInstanceRequestAllOfMetadata.md) |  | [optional] 
**RootPass** | Pointer to **string** | This sets the root user&#39;s password on a newly created Linode Disk when deploying from an Image.  - __Required__ when creating a Linode Disk from an Image, including when using a StackScript.  - Must meet a password strength score requirement that is calculated internally by the API. If the strength requirement is not met, you will receive a &#x60;Password does not meet strength requirement&#x60; error. | [optional] 
**StackscriptData** | Pointer to **map[string]interface{}** | This field is required only if the StackScript being deployed requires input data from the User for successful completion. See [User Defined Fields (UDFs)](https://www.linode.com/docs/products/tools/stackscripts/guides/write-a-custom-script/#declare-user-defined-fields-udfs) for more details.  This field is required to be valid JSON.  Total length cannot exceed 65,535 characters. | [optional] 
**StackscriptId** | Pointer to **int32** | A StackScript ID that will cause the referenced StackScript to be run during deployment of this Linode. A compatible &#x60;image&#x60; is required to use a StackScript. To get a list of available StackScript and their permitted Images, run [List StackScripts](https://techdocs.akamai.com/linode-api/reference/get-stack-scripts). This field cannot be used when deploying from a Backup or a Private Image. | [optional] 
**BackupId** | Pointer to **int32** | A Backup ID from another Linode&#39;s available backups. Your User must have &#x60;read_write&#x60; access to that Linode, the Backup must have a &#x60;status&#x60; of &#x60;successful&#x60;, and the Linode must be deployed to the same &#x60;region&#x60; as the Backup. Run [List backups](https://techdocs.akamai.com/linode-api/reference/get-backups) for a Linode&#39;s available backups.  This field and the &#x60;image&#x60; field are mutually exclusive. | [optional] 
**BackupsEnabled** | Pointer to **bool** | If this field is set to &#x60;true&#x60;, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.  This option is always treated as &#x60;true&#x60; if the account-wide &#x60;backups_enabled&#x60; setting is &#x60;true&#x60;.  See [Get account settings](https://techdocs.akamai.com/linode-api/reference/get-account-settings) for more information.  Backup pricing is included in the response from [List types](https://techdocs.akamai.com/linode-api/reference/get-linode-types) | [optional] 
**FirewallId** | Pointer to **int32** | The &#x60;id&#x60; of the Firewall to attach this Linode to upon creation. | [optional] 
**Group** | Pointer to **string** | The group label for this Linode. | [optional] 
**Interfaces** | Pointer to [**[]PostLinodeInstanceRequestAllOfInterfacesInner**](PostLinodeInstanceRequestAllOfInterfacesInner.md) | An array of Network Interfaces to add to this Linode&#39;s Configuration Profile. At least one and up to three Interface objects can exist in this array. The position in the array determines which of the Linode&#39;s network Interfaces is configured:  - First [0]:  eth0 - Second [1]: eth1 - Third [2]:  eth2  When updating a Linode&#39;s Interfaces, _each Interface must be redefined_. An empty &#x60;interfaces&#x60; array results in a default &#x60;public&#x60; type Interface configuration only.  If no public Interface is configured, public IP addresses are still assigned to the Linode but will not be usable without manual configuration.  __Note__. Changes to Linode Interface configurations can be enabled by rebooting the Linode.  &#x60;vpc&#x60; details  See the [VPC documentation](https://www.linode.com/docs/products/networking/vpc/#technical-specifications) guide for its specifications and limitations.  &#x60;vlan&#x60; details  - Only Next Generation Network (NGN) data centers support VLANs. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to view the capabilities of data center regions. If a VLAN is attached to your Linode and you attempt to migrate or clone it to a non-NGN data center, the migration or cloning will not initiate. If a Linode cannot be migrated or cloned because of an incompatibility, you will be prompted to select a different data center or contact support. - See the [VLANs Overview](https://www.linode.com/docs/products/networking/vlans/#technical-specifications) guide to view additional specifications and limitations. | [optional] 
**Label** | Pointer to **string** | Provides a name for the Linode. If not provided, the API generates one for it.  Linode labels have the following constraints:  - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (&#x60;-&#x60;), underscores (&#x60;_&#x60;) or periods (&#x60;.&#x60;). - Cannot have two hyphens (&#x60;--&#x60;), underscores (&#x60;__&#x60;) or periods (&#x60;..&#x60;) in a row. | [optional] 
**PlacementGroup** | Pointer to [**PostLinodeInstanceRequestAllOfPlacementGroup**](PostLinodeInstanceRequestAllOfPlacementGroup.md) |  | [optional] 
**PrivateIp** | Pointer to **bool** | If true, the created Linode will have private networking enabled and assigned a private IPv4 address. | [optional] 
**Region** | **string** | The [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where the Linode will be located. | 
**SwapSize** | Pointer to **int32** | When deploying from an Image, this field is optional, otherwise it is ignored. This is used to set the swap disk size for the newly created Linode. | [optional] [default to 512]
**Tags** | Pointer to **[]string** | Tags to help you organize your content. | [optional] 
**Type** | **string** | The [Linode type](https://techdocs.akamai.com/linode-api/reference/get-linode-types) of the Linode you are creating. | 

## Methods

### NewPostLinodeInstanceRequest

`func NewPostLinodeInstanceRequest(region string, type_ string, ) *PostLinodeInstanceRequest`

NewPostLinodeInstanceRequest instantiates a new PostLinodeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLinodeInstanceRequestWithDefaults

`func NewPostLinodeInstanceRequestWithDefaults() *PostLinodeInstanceRequest`

NewPostLinodeInstanceRequestWithDefaults instantiates a new PostLinodeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizedKeys

`func (o *PostLinodeInstanceRequest) GetAuthorizedKeys() []string`

GetAuthorizedKeys returns the AuthorizedKeys field if non-nil, zero value otherwise.

### GetAuthorizedKeysOk

`func (o *PostLinodeInstanceRequest) GetAuthorizedKeysOk() (*[]string, bool)`

GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedKeys

`func (o *PostLinodeInstanceRequest) SetAuthorizedKeys(v []string)`

SetAuthorizedKeys sets AuthorizedKeys field to given value.

### HasAuthorizedKeys

`func (o *PostLinodeInstanceRequest) HasAuthorizedKeys() bool`

HasAuthorizedKeys returns a boolean if a field has been set.

### GetAuthorizedUsers

`func (o *PostLinodeInstanceRequest) GetAuthorizedUsers() []string`

GetAuthorizedUsers returns the AuthorizedUsers field if non-nil, zero value otherwise.

### GetAuthorizedUsersOk

`func (o *PostLinodeInstanceRequest) GetAuthorizedUsersOk() (*[]string, bool)`

GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedUsers

`func (o *PostLinodeInstanceRequest) SetAuthorizedUsers(v []string)`

SetAuthorizedUsers sets AuthorizedUsers field to given value.

### HasAuthorizedUsers

`func (o *PostLinodeInstanceRequest) HasAuthorizedUsers() bool`

HasAuthorizedUsers returns a boolean if a field has been set.

### GetBooted

`func (o *PostLinodeInstanceRequest) GetBooted() bool`

GetBooted returns the Booted field if non-nil, zero value otherwise.

### GetBootedOk

`func (o *PostLinodeInstanceRequest) GetBootedOk() (*bool, bool)`

GetBootedOk returns a tuple with the Booted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBooted

`func (o *PostLinodeInstanceRequest) SetBooted(v bool)`

SetBooted sets Booted field to given value.

### HasBooted

`func (o *PostLinodeInstanceRequest) HasBooted() bool`

HasBooted returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *PostLinodeInstanceRequest) GetDiskEncryption() string`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *PostLinodeInstanceRequest) GetDiskEncryptionOk() (*string, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *PostLinodeInstanceRequest) SetDiskEncryption(v string)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *PostLinodeInstanceRequest) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### GetImage

`func (o *PostLinodeInstanceRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *PostLinodeInstanceRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *PostLinodeInstanceRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *PostLinodeInstanceRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetMetadata

`func (o *PostLinodeInstanceRequest) GetMetadata() PostLinodeInstanceRequestAllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostLinodeInstanceRequest) GetMetadataOk() (*PostLinodeInstanceRequestAllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostLinodeInstanceRequest) SetMetadata(v PostLinodeInstanceRequestAllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PostLinodeInstanceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRootPass

`func (o *PostLinodeInstanceRequest) GetRootPass() string`

GetRootPass returns the RootPass field if non-nil, zero value otherwise.

### GetRootPassOk

`func (o *PostLinodeInstanceRequest) GetRootPassOk() (*string, bool)`

GetRootPassOk returns a tuple with the RootPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPass

`func (o *PostLinodeInstanceRequest) SetRootPass(v string)`

SetRootPass sets RootPass field to given value.

### HasRootPass

`func (o *PostLinodeInstanceRequest) HasRootPass() bool`

HasRootPass returns a boolean if a field has been set.

### GetStackscriptData

`func (o *PostLinodeInstanceRequest) GetStackscriptData() map[string]interface{}`

GetStackscriptData returns the StackscriptData field if non-nil, zero value otherwise.

### GetStackscriptDataOk

`func (o *PostLinodeInstanceRequest) GetStackscriptDataOk() (*map[string]interface{}, bool)`

GetStackscriptDataOk returns a tuple with the StackscriptData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptData

`func (o *PostLinodeInstanceRequest) SetStackscriptData(v map[string]interface{})`

SetStackscriptData sets StackscriptData field to given value.

### HasStackscriptData

`func (o *PostLinodeInstanceRequest) HasStackscriptData() bool`

HasStackscriptData returns a boolean if a field has been set.

### GetStackscriptId

`func (o *PostLinodeInstanceRequest) GetStackscriptId() int32`

GetStackscriptId returns the StackscriptId field if non-nil, zero value otherwise.

### GetStackscriptIdOk

`func (o *PostLinodeInstanceRequest) GetStackscriptIdOk() (*int32, bool)`

GetStackscriptIdOk returns a tuple with the StackscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscriptId

`func (o *PostLinodeInstanceRequest) SetStackscriptId(v int32)`

SetStackscriptId sets StackscriptId field to given value.

### HasStackscriptId

`func (o *PostLinodeInstanceRequest) HasStackscriptId() bool`

HasStackscriptId returns a boolean if a field has been set.

### GetBackupId

`func (o *PostLinodeInstanceRequest) GetBackupId() int32`

GetBackupId returns the BackupId field if non-nil, zero value otherwise.

### GetBackupIdOk

`func (o *PostLinodeInstanceRequest) GetBackupIdOk() (*int32, bool)`

GetBackupIdOk returns a tuple with the BackupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupId

`func (o *PostLinodeInstanceRequest) SetBackupId(v int32)`

SetBackupId sets BackupId field to given value.

### HasBackupId

`func (o *PostLinodeInstanceRequest) HasBackupId() bool`

HasBackupId returns a boolean if a field has been set.

### GetBackupsEnabled

`func (o *PostLinodeInstanceRequest) GetBackupsEnabled() bool`

GetBackupsEnabled returns the BackupsEnabled field if non-nil, zero value otherwise.

### GetBackupsEnabledOk

`func (o *PostLinodeInstanceRequest) GetBackupsEnabledOk() (*bool, bool)`

GetBackupsEnabledOk returns a tuple with the BackupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsEnabled

`func (o *PostLinodeInstanceRequest) SetBackupsEnabled(v bool)`

SetBackupsEnabled sets BackupsEnabled field to given value.

### HasBackupsEnabled

`func (o *PostLinodeInstanceRequest) HasBackupsEnabled() bool`

HasBackupsEnabled returns a boolean if a field has been set.

### GetFirewallId

`func (o *PostLinodeInstanceRequest) GetFirewallId() int32`

GetFirewallId returns the FirewallId field if non-nil, zero value otherwise.

### GetFirewallIdOk

`func (o *PostLinodeInstanceRequest) GetFirewallIdOk() (*int32, bool)`

GetFirewallIdOk returns a tuple with the FirewallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallId

`func (o *PostLinodeInstanceRequest) SetFirewallId(v int32)`

SetFirewallId sets FirewallId field to given value.

### HasFirewallId

`func (o *PostLinodeInstanceRequest) HasFirewallId() bool`

HasFirewallId returns a boolean if a field has been set.

### GetGroup

`func (o *PostLinodeInstanceRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PostLinodeInstanceRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PostLinodeInstanceRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PostLinodeInstanceRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInterfaces

`func (o *PostLinodeInstanceRequest) GetInterfaces() []PostLinodeInstanceRequestAllOfInterfacesInner`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *PostLinodeInstanceRequest) GetInterfacesOk() (*[]PostLinodeInstanceRequestAllOfInterfacesInner, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *PostLinodeInstanceRequest) SetInterfaces(v []PostLinodeInstanceRequestAllOfInterfacesInner)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *PostLinodeInstanceRequest) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetLabel

`func (o *PostLinodeInstanceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostLinodeInstanceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostLinodeInstanceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostLinodeInstanceRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPlacementGroup

`func (o *PostLinodeInstanceRequest) GetPlacementGroup() PostLinodeInstanceRequestAllOfPlacementGroup`

GetPlacementGroup returns the PlacementGroup field if non-nil, zero value otherwise.

### GetPlacementGroupOk

`func (o *PostLinodeInstanceRequest) GetPlacementGroupOk() (*PostLinodeInstanceRequestAllOfPlacementGroup, bool)`

GetPlacementGroupOk returns a tuple with the PlacementGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroup

`func (o *PostLinodeInstanceRequest) SetPlacementGroup(v PostLinodeInstanceRequestAllOfPlacementGroup)`

SetPlacementGroup sets PlacementGroup field to given value.

### HasPlacementGroup

`func (o *PostLinodeInstanceRequest) HasPlacementGroup() bool`

HasPlacementGroup returns a boolean if a field has been set.

### GetPrivateIp

`func (o *PostLinodeInstanceRequest) GetPrivateIp() bool`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *PostLinodeInstanceRequest) GetPrivateIpOk() (*bool, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *PostLinodeInstanceRequest) SetPrivateIp(v bool)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *PostLinodeInstanceRequest) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetRegion

`func (o *PostLinodeInstanceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostLinodeInstanceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostLinodeInstanceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetSwapSize

`func (o *PostLinodeInstanceRequest) GetSwapSize() int32`

GetSwapSize returns the SwapSize field if non-nil, zero value otherwise.

### GetSwapSizeOk

`func (o *PostLinodeInstanceRequest) GetSwapSizeOk() (*int32, bool)`

GetSwapSizeOk returns a tuple with the SwapSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwapSize

`func (o *PostLinodeInstanceRequest) SetSwapSize(v int32)`

SetSwapSize sets SwapSize field to given value.

### HasSwapSize

`func (o *PostLinodeInstanceRequest) HasSwapSize() bool`

HasSwapSize returns a boolean if a field has been set.

### GetTags

`func (o *PostLinodeInstanceRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostLinodeInstanceRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostLinodeInstanceRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostLinodeInstanceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *PostLinodeInstanceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostLinodeInstanceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostLinodeInstanceRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


