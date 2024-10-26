/*
Linode API

[Read the API documentation](https://techdocs.akamai.com/linode-api/reference/api).

API version: 4.189.2
Contact: support@linode.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PostLinodeInstanceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostLinodeInstanceRequest{}

// PostLinodeInstanceRequest struct for PostLinodeInstanceRequest
type PostLinodeInstanceRequest struct {
	// A list of public SSH keys that will be automatically appended to the root user's `~/.ssh/authorized_keys` file when deploying from an Image.
	AuthorizedKeys []string `json:"authorized_keys,omitempty"`
	// A list of usernames. If the usernames have associated SSH keys, the keys will be appended to the root users `~/.ssh/authorized_keys` file automatically when deploying from an Image.
	AuthorizedUsers []string `json:"authorized_users,omitempty"`
	// This field defaults to `true` if the Linode is created with an Image or from a Backup. If it is deployed from an Image or a Backup and you wish it to remain `offline` after deployment, set this to `false`.
	Booted *bool `json:"booted,omitempty"`
	// Local disk encryption ensures that your data stored on Linodes is secured. Disk encryption protects against unauthorized data access by keeping the data encrypted if the disk is ever removed from the data center, decommissioned, or disposed of. The platform manages the encryption and decryption for you.  By default, encryption is `enabled` on all Linodes. If you opted out of encryption or if the Linode was created prior to local disk encryption support, you can encrypt your data using [Rebuild](https://techdocs.akamai.com/linode-api/reference/post-rebuild-linode-instance).
	DiskEncryption *string `json:"disk_encryption,omitempty"`
	// An Image ID to deploy the Linode Disk from.  Run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation with authentication to view all available Images. Official Linode Images start with `linode/`, while your Account's Images start with `private/`. Creating a disk from a Private Image requires `read_only` or `read_write` permissions for that Image. Run the [Update a user's grants](https://techdocs.akamai.com/linode-api/reference/put-user-grants) operation to adjust permissions for an Account Image.
	Image *string `json:"image,omitempty"`
	Metadata *PostLinodeInstanceRequestAllOfMetadata `json:"metadata,omitempty"`
	// This sets the root user's password on a newly created Linode Disk when deploying from an Image.  - __Required__ when creating a Linode Disk from an Image, including when using a StackScript.  - Must meet a password strength score requirement that is calculated internally by the API. If the strength requirement is not met, you will receive a `Password does not meet strength requirement` error.
	RootPass *string `json:"root_pass,omitempty"`
	// This field is required only if the StackScript being deployed requires input data from the User for successful completion. See [User Defined Fields (UDFs)](https://www.linode.com/docs/products/tools/stackscripts/guides/write-a-custom-script/#declare-user-defined-fields-udfs) for more details.  This field is required to be valid JSON.  Total length cannot exceed 65,535 characters.
	StackscriptData map[string]interface{} `json:"stackscript_data,omitempty"`
	// A StackScript ID that will cause the referenced StackScript to be run during deployment of this Linode. A compatible `image` is required to use a StackScript. To get a list of available StackScript and their permitted Images, run [List StackScripts](https://techdocs.akamai.com/linode-api/reference/get-stack-scripts). This field cannot be used when deploying from a Backup or a Private Image.
	StackscriptId *int32 `json:"stackscript_id,omitempty"`
	// A Backup ID from another Linode's available backups. Your User must have `read_write` access to that Linode, the Backup must have a `status` of `successful`, and the Linode must be deployed to the same `region` as the Backup. Run [List backups](https://techdocs.akamai.com/linode-api/reference/get-backups) for a Linode's available backups.  This field and the `image` field are mutually exclusive.
	BackupId *int32 `json:"backup_id,omitempty"`
	// If this field is set to `true`, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. The cost for the Backup service is dependent on the Type of Linode deployed.  This option is always treated as `true` if the account-wide `backups_enabled` setting is `true`.  See [Get account settings](https://techdocs.akamai.com/linode-api/reference/get-account-settings) for more information.  Backup pricing is included in the response from [List types](https://techdocs.akamai.com/linode-api/reference/get-linode-types)
	BackupsEnabled *bool `json:"backups_enabled,omitempty"`
	// The `id` of the Firewall to attach this Linode to upon creation.
	FirewallId *int32 `json:"firewall_id,omitempty"`
	// The group label for this Linode.
	// Deprecated
	Group *string `json:"group,omitempty"`
	// An array of Network Interfaces to add to this Linode's Configuration Profile. At least one and up to three Interface objects can exist in this array. The position in the array determines which of the Linode's network Interfaces is configured:  - First [0]:  eth0 - Second [1]: eth1 - Third [2]:  eth2  When updating a Linode's Interfaces, _each Interface must be redefined_. An empty `interfaces` array results in a default `public` type Interface configuration only.  If no public Interface is configured, public IP addresses are still assigned to the Linode but will not be usable without manual configuration.  __Note__. Changes to Linode Interface configurations can be enabled by rebooting the Linode.  `vpc` details  See the [VPC documentation](https://www.linode.com/docs/products/networking/vpc/#technical-specifications) guide for its specifications and limitations.  `vlan` details  - Only Next Generation Network (NGN) data centers support VLANs. Run the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation to view the capabilities of data center regions. If a VLAN is attached to your Linode and you attempt to migrate or clone it to a non-NGN data center, the migration or cloning will not initiate. If a Linode cannot be migrated or cloned because of an incompatibility, you will be prompted to select a different data center or contact support. - See the [VLANs Overview](https://www.linode.com/docs/products/networking/vlans/#technical-specifications) guide to view additional specifications and limitations.
	Interfaces []PostLinodeInstanceRequestAllOfInterfacesInner `json:"interfaces,omitempty"`
	// Provides a name for the Linode. If not provided, the API generates one for it.  Linode labels have the following constraints:  - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (`-`), underscores (`_`) or periods (`.`). - Cannot have two hyphens (`--`), underscores (`__`) or periods (`..`) in a row.
	Label *string `json:"label,omitempty" validate:"regexp=^[a-zA-Z]((?!--|__|\\\\.\\\\.)[a-zA-Z0-9-_.])+$"`
	PlacementGroup *PostLinodeInstanceRequestAllOfPlacementGroup `json:"placement_group,omitempty"`
	// If true, the created Linode will have private networking enabled and assigned a private IPv4 address.
	PrivateIp *bool `json:"private_ip,omitempty"`
	// The [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where the Linode will be located.
	Region string `json:"region"`
	// When deploying from an Image, this field is optional, otherwise it is ignored. This is used to set the swap disk size for the newly created Linode.
	SwapSize *int32 `json:"swap_size,omitempty"`
	// Tags to help you organize your content.
	Tags []string `json:"tags,omitempty"`
	// The [Linode type](https://techdocs.akamai.com/linode-api/reference/get-linode-types) of the Linode you are creating.
	Type string `json:"type"`
}

type _PostLinodeInstanceRequest PostLinodeInstanceRequest

// NewPostLinodeInstanceRequest instantiates a new PostLinodeInstanceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostLinodeInstanceRequest(region string, type_ string) *PostLinodeInstanceRequest {
	this := PostLinodeInstanceRequest{}
	var booted bool = true
	this.Booted = &booted
	this.Region = region
	var swapSize int32 = 512
	this.SwapSize = &swapSize
	this.Type = type_
	return &this
}

// NewPostLinodeInstanceRequestWithDefaults instantiates a new PostLinodeInstanceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostLinodeInstanceRequestWithDefaults() *PostLinodeInstanceRequest {
	this := PostLinodeInstanceRequest{}
	var booted bool = true
	this.Booted = &booted
	var swapSize int32 = 512
	this.SwapSize = &swapSize
	return &this
}

// GetAuthorizedKeys returns the AuthorizedKeys field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetAuthorizedKeys() []string {
	if o == nil || IsNil(o.AuthorizedKeys) {
		var ret []string
		return ret
	}
	return o.AuthorizedKeys
}

// GetAuthorizedKeysOk returns a tuple with the AuthorizedKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetAuthorizedKeysOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizedKeys) {
		return nil, false
	}
	return o.AuthorizedKeys, true
}

// HasAuthorizedKeys returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasAuthorizedKeys() bool {
	if o != nil && !IsNil(o.AuthorizedKeys) {
		return true
	}

	return false
}

// SetAuthorizedKeys gets a reference to the given []string and assigns it to the AuthorizedKeys field.
func (o *PostLinodeInstanceRequest) SetAuthorizedKeys(v []string) {
	o.AuthorizedKeys = v
}

// GetAuthorizedUsers returns the AuthorizedUsers field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetAuthorizedUsers() []string {
	if o == nil || IsNil(o.AuthorizedUsers) {
		var ret []string
		return ret
	}
	return o.AuthorizedUsers
}

// GetAuthorizedUsersOk returns a tuple with the AuthorizedUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetAuthorizedUsersOk() ([]string, bool) {
	if o == nil || IsNil(o.AuthorizedUsers) {
		return nil, false
	}
	return o.AuthorizedUsers, true
}

// HasAuthorizedUsers returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasAuthorizedUsers() bool {
	if o != nil && !IsNil(o.AuthorizedUsers) {
		return true
	}

	return false
}

// SetAuthorizedUsers gets a reference to the given []string and assigns it to the AuthorizedUsers field.
func (o *PostLinodeInstanceRequest) SetAuthorizedUsers(v []string) {
	o.AuthorizedUsers = v
}

// GetBooted returns the Booted field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetBooted() bool {
	if o == nil || IsNil(o.Booted) {
		var ret bool
		return ret
	}
	return *o.Booted
}

// GetBootedOk returns a tuple with the Booted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetBootedOk() (*bool, bool) {
	if o == nil || IsNil(o.Booted) {
		return nil, false
	}
	return o.Booted, true
}

// HasBooted returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasBooted() bool {
	if o != nil && !IsNil(o.Booted) {
		return true
	}

	return false
}

// SetBooted gets a reference to the given bool and assigns it to the Booted field.
func (o *PostLinodeInstanceRequest) SetBooted(v bool) {
	o.Booted = &v
}

// GetDiskEncryption returns the DiskEncryption field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetDiskEncryption() string {
	if o == nil || IsNil(o.DiskEncryption) {
		var ret string
		return ret
	}
	return *o.DiskEncryption
}

// GetDiskEncryptionOk returns a tuple with the DiskEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetDiskEncryptionOk() (*string, bool) {
	if o == nil || IsNil(o.DiskEncryption) {
		return nil, false
	}
	return o.DiskEncryption, true
}

// HasDiskEncryption returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasDiskEncryption() bool {
	if o != nil && !IsNil(o.DiskEncryption) {
		return true
	}

	return false
}

// SetDiskEncryption gets a reference to the given string and assigns it to the DiskEncryption field.
func (o *PostLinodeInstanceRequest) SetDiskEncryption(v string) {
	o.DiskEncryption = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *PostLinodeInstanceRequest) SetImage(v string) {
	o.Image = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetMetadata() PostLinodeInstanceRequestAllOfMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret PostLinodeInstanceRequestAllOfMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetMetadataOk() (*PostLinodeInstanceRequestAllOfMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given PostLinodeInstanceRequestAllOfMetadata and assigns it to the Metadata field.
func (o *PostLinodeInstanceRequest) SetMetadata(v PostLinodeInstanceRequestAllOfMetadata) {
	o.Metadata = &v
}

// GetRootPass returns the RootPass field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetRootPass() string {
	if o == nil || IsNil(o.RootPass) {
		var ret string
		return ret
	}
	return *o.RootPass
}

// GetRootPassOk returns a tuple with the RootPass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetRootPassOk() (*string, bool) {
	if o == nil || IsNil(o.RootPass) {
		return nil, false
	}
	return o.RootPass, true
}

// HasRootPass returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasRootPass() bool {
	if o != nil && !IsNil(o.RootPass) {
		return true
	}

	return false
}

// SetRootPass gets a reference to the given string and assigns it to the RootPass field.
func (o *PostLinodeInstanceRequest) SetRootPass(v string) {
	o.RootPass = &v
}

// GetStackscriptData returns the StackscriptData field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetStackscriptData() map[string]interface{} {
	if o == nil || IsNil(o.StackscriptData) {
		var ret map[string]interface{}
		return ret
	}
	return o.StackscriptData
}

// GetStackscriptDataOk returns a tuple with the StackscriptData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetStackscriptDataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.StackscriptData) {
		return map[string]interface{}{}, false
	}
	return o.StackscriptData, true
}

// HasStackscriptData returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasStackscriptData() bool {
	if o != nil && !IsNil(o.StackscriptData) {
		return true
	}

	return false
}

// SetStackscriptData gets a reference to the given map[string]interface{} and assigns it to the StackscriptData field.
func (o *PostLinodeInstanceRequest) SetStackscriptData(v map[string]interface{}) {
	o.StackscriptData = v
}

// GetStackscriptId returns the StackscriptId field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetStackscriptId() int32 {
	if o == nil || IsNil(o.StackscriptId) {
		var ret int32
		return ret
	}
	return *o.StackscriptId
}

// GetStackscriptIdOk returns a tuple with the StackscriptId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetStackscriptIdOk() (*int32, bool) {
	if o == nil || IsNil(o.StackscriptId) {
		return nil, false
	}
	return o.StackscriptId, true
}

// HasStackscriptId returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasStackscriptId() bool {
	if o != nil && !IsNil(o.StackscriptId) {
		return true
	}

	return false
}

// SetStackscriptId gets a reference to the given int32 and assigns it to the StackscriptId field.
func (o *PostLinodeInstanceRequest) SetStackscriptId(v int32) {
	o.StackscriptId = &v
}

// GetBackupId returns the BackupId field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetBackupId() int32 {
	if o == nil || IsNil(o.BackupId) {
		var ret int32
		return ret
	}
	return *o.BackupId
}

// GetBackupIdOk returns a tuple with the BackupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetBackupIdOk() (*int32, bool) {
	if o == nil || IsNil(o.BackupId) {
		return nil, false
	}
	return o.BackupId, true
}

// HasBackupId returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasBackupId() bool {
	if o != nil && !IsNil(o.BackupId) {
		return true
	}

	return false
}

// SetBackupId gets a reference to the given int32 and assigns it to the BackupId field.
func (o *PostLinodeInstanceRequest) SetBackupId(v int32) {
	o.BackupId = &v
}

// GetBackupsEnabled returns the BackupsEnabled field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetBackupsEnabled() bool {
	if o == nil || IsNil(o.BackupsEnabled) {
		var ret bool
		return ret
	}
	return *o.BackupsEnabled
}

// GetBackupsEnabledOk returns a tuple with the BackupsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetBackupsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BackupsEnabled) {
		return nil, false
	}
	return o.BackupsEnabled, true
}

// HasBackupsEnabled returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasBackupsEnabled() bool {
	if o != nil && !IsNil(o.BackupsEnabled) {
		return true
	}

	return false
}

// SetBackupsEnabled gets a reference to the given bool and assigns it to the BackupsEnabled field.
func (o *PostLinodeInstanceRequest) SetBackupsEnabled(v bool) {
	o.BackupsEnabled = &v
}

// GetFirewallId returns the FirewallId field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetFirewallId() int32 {
	if o == nil || IsNil(o.FirewallId) {
		var ret int32
		return ret
	}
	return *o.FirewallId
}

// GetFirewallIdOk returns a tuple with the FirewallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetFirewallIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FirewallId) {
		return nil, false
	}
	return o.FirewallId, true
}

// HasFirewallId returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasFirewallId() bool {
	if o != nil && !IsNil(o.FirewallId) {
		return true
	}

	return false
}

// SetFirewallId gets a reference to the given int32 and assigns it to the FirewallId field.
func (o *PostLinodeInstanceRequest) SetFirewallId(v int32) {
	o.FirewallId = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
// Deprecated
func (o *PostLinodeInstanceRequest) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *PostLinodeInstanceRequest) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
// Deprecated
func (o *PostLinodeInstanceRequest) SetGroup(v string) {
	o.Group = &v
}

// GetInterfaces returns the Interfaces field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetInterfaces() []PostLinodeInstanceRequestAllOfInterfacesInner {
	if o == nil || IsNil(o.Interfaces) {
		var ret []PostLinodeInstanceRequestAllOfInterfacesInner
		return ret
	}
	return o.Interfaces
}

// GetInterfacesOk returns a tuple with the Interfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetInterfacesOk() ([]PostLinodeInstanceRequestAllOfInterfacesInner, bool) {
	if o == nil || IsNil(o.Interfaces) {
		return nil, false
	}
	return o.Interfaces, true
}

// HasInterfaces returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasInterfaces() bool {
	if o != nil && !IsNil(o.Interfaces) {
		return true
	}

	return false
}

// SetInterfaces gets a reference to the given []PostLinodeInstanceRequestAllOfInterfacesInner and assigns it to the Interfaces field.
func (o *PostLinodeInstanceRequest) SetInterfaces(v []PostLinodeInstanceRequestAllOfInterfacesInner) {
	o.Interfaces = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *PostLinodeInstanceRequest) SetLabel(v string) {
	o.Label = &v
}

// GetPlacementGroup returns the PlacementGroup field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetPlacementGroup() PostLinodeInstanceRequestAllOfPlacementGroup {
	if o == nil || IsNil(o.PlacementGroup) {
		var ret PostLinodeInstanceRequestAllOfPlacementGroup
		return ret
	}
	return *o.PlacementGroup
}

// GetPlacementGroupOk returns a tuple with the PlacementGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetPlacementGroupOk() (*PostLinodeInstanceRequestAllOfPlacementGroup, bool) {
	if o == nil || IsNil(o.PlacementGroup) {
		return nil, false
	}
	return o.PlacementGroup, true
}

// HasPlacementGroup returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasPlacementGroup() bool {
	if o != nil && !IsNil(o.PlacementGroup) {
		return true
	}

	return false
}

// SetPlacementGroup gets a reference to the given PostLinodeInstanceRequestAllOfPlacementGroup and assigns it to the PlacementGroup field.
func (o *PostLinodeInstanceRequest) SetPlacementGroup(v PostLinodeInstanceRequestAllOfPlacementGroup) {
	o.PlacementGroup = &v
}

// GetPrivateIp returns the PrivateIp field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetPrivateIp() bool {
	if o == nil || IsNil(o.PrivateIp) {
		var ret bool
		return ret
	}
	return *o.PrivateIp
}

// GetPrivateIpOk returns a tuple with the PrivateIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetPrivateIpOk() (*bool, bool) {
	if o == nil || IsNil(o.PrivateIp) {
		return nil, false
	}
	return o.PrivateIp, true
}

// HasPrivateIp returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasPrivateIp() bool {
	if o != nil && !IsNil(o.PrivateIp) {
		return true
	}

	return false
}

// SetPrivateIp gets a reference to the given bool and assigns it to the PrivateIp field.
func (o *PostLinodeInstanceRequest) SetPrivateIp(v bool) {
	o.PrivateIp = &v
}

// GetRegion returns the Region field value
func (o *PostLinodeInstanceRequest) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *PostLinodeInstanceRequest) SetRegion(v string) {
	o.Region = v
}

// GetSwapSize returns the SwapSize field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetSwapSize() int32 {
	if o == nil || IsNil(o.SwapSize) {
		var ret int32
		return ret
	}
	return *o.SwapSize
}

// GetSwapSizeOk returns a tuple with the SwapSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetSwapSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.SwapSize) {
		return nil, false
	}
	return o.SwapSize, true
}

// HasSwapSize returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasSwapSize() bool {
	if o != nil && !IsNil(o.SwapSize) {
		return true
	}

	return false
}

// SetSwapSize gets a reference to the given int32 and assigns it to the SwapSize field.
func (o *PostLinodeInstanceRequest) SetSwapSize(v int32) {
	o.SwapSize = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PostLinodeInstanceRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PostLinodeInstanceRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *PostLinodeInstanceRequest) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *PostLinodeInstanceRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PostLinodeInstanceRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PostLinodeInstanceRequest) SetType(v string) {
	o.Type = v
}

func (o PostLinodeInstanceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostLinodeInstanceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizedKeys) {
		toSerialize["authorized_keys"] = o.AuthorizedKeys
	}
	if !IsNil(o.AuthorizedUsers) {
		toSerialize["authorized_users"] = o.AuthorizedUsers
	}
	if !IsNil(o.Booted) {
		toSerialize["booted"] = o.Booted
	}
	if !IsNil(o.DiskEncryption) {
		toSerialize["disk_encryption"] = o.DiskEncryption
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.RootPass) {
		toSerialize["root_pass"] = o.RootPass
	}
	if !IsNil(o.StackscriptData) {
		toSerialize["stackscript_data"] = o.StackscriptData
	}
	if !IsNil(o.StackscriptId) {
		toSerialize["stackscript_id"] = o.StackscriptId
	}
	if !IsNil(o.BackupId) {
		toSerialize["backup_id"] = o.BackupId
	}
	if !IsNil(o.BackupsEnabled) {
		toSerialize["backups_enabled"] = o.BackupsEnabled
	}
	if !IsNil(o.FirewallId) {
		toSerialize["firewall_id"] = o.FirewallId
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Interfaces) {
		toSerialize["interfaces"] = o.Interfaces
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.PlacementGroup) {
		toSerialize["placement_group"] = o.PlacementGroup
	}
	if !IsNil(o.PrivateIp) {
		toSerialize["private_ip"] = o.PrivateIp
	}
	toSerialize["region"] = o.Region
	if !IsNil(o.SwapSize) {
		toSerialize["swap_size"] = o.SwapSize
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *PostLinodeInstanceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"region",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPostLinodeInstanceRequest := _PostLinodeInstanceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostLinodeInstanceRequest)

	if err != nil {
		return err
	}

	*o = PostLinodeInstanceRequest(varPostLinodeInstanceRequest)

	return err
}

type NullablePostLinodeInstanceRequest struct {
	value *PostLinodeInstanceRequest
	isSet bool
}

func (v NullablePostLinodeInstanceRequest) Get() *PostLinodeInstanceRequest {
	return v.value
}

func (v *NullablePostLinodeInstanceRequest) Set(val *PostLinodeInstanceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostLinodeInstanceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostLinodeInstanceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostLinodeInstanceRequest(val *PostLinodeInstanceRequest) *NullablePostLinodeInstanceRequest {
	return &NullablePostLinodeInstanceRequest{value: val, isSet: true}
}

func (v NullablePostLinodeInstanceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostLinodeInstanceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

