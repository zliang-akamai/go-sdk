# GetTaggedObjects200ResponseDataInnerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**LinodeAlerts**](LinodeAlerts.md) |  | [optional] 
**Backups** | Pointer to [**LinodeBackups**](LinodeBackups.md) |  | [optional] 
**Capabilities** | Pointer to **[]string** | A list of capabilities this compute instance supports. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | When this NodeBalancer was created. | [optional] [readonly] 
**DiskEncryption** | Pointer to **NullableString** | Indicates the local disk encryption setting for this Linode. If the Linode is part of an LKE cluster, the value is &#x60;null&#x60;. | [optional] [readonly] [default to "enabled"]
**Group** | Pointer to **string** | The group this Domain belongs to.  This is for display purposes only. | [optional] 
**HasUserData** | Pointer to **bool** | Whether this compute instance was provisioned with &#x60;user_data&#x60; provided via the Metadata service. See the [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance) description for more information on Metadata. | [optional] [readonly] 
**HostUuid** | Pointer to **string** | The Linode&#39;s host machine, as a UUID. | [optional] [readonly] 
**Hypervisor** | Pointer to **string** | The virtualization software powering this Linode. | [optional] [readonly] 
**Id** | Pointer to **int32** | This NodeBalancer&#39;s unique ID. | [optional] [readonly] 
**Image** | Pointer to **string** | An Image ID to deploy the Linode Disk from.  Run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation with authentication to view all available Images. Official Linode Images start with &#x60;linode/&#x60;, while your Account&#39;s Images start with &#x60;private/&#x60;. Creating a disk from a Private Image requires &#x60;read_only&#x60; or &#x60;read_write&#x60; permissions for that Image. Run the [Update a user&#39;s grants](https://techdocs.akamai.com/linode-api/reference/put-user-grants) operation to adjust permissions for an Account Image. | [optional] 
**Ipv4** | Pointer to **string** | This NodeBalancer&#39;s public IPv4 address. | [optional] [readonly] 
**Ipv6** | Pointer to **NullableString** | This NodeBalancer&#39;s public IPv6 address. | [optional] [readonly] 
**Label** | Pointer to **string** | This NodeBalancer&#39;s label. These must be unique on your Account. | [optional] 
**LkeClusterId** | Pointer to **NullableInt32** | The ID of the Kubernetes cluster if the Linode is part of cluster. | [optional] [readonly] 
**PlacementGroup** | Pointer to [**NullableLinode1PlacementGroup**](Linode1PlacementGroup.md) |  | [optional] 
**Region** | Pointer to **string** | The Region where this NodeBalancer is located. NodeBalancers only support backends in the same Region. | [optional] [readonly] 
**Specs** | Pointer to [**LinodeSpecs**](LinodeSpecs.md) |  | [optional] 
**Status** | Pointer to **string** | The current status of the volume.  Can be one of:    - &#x60;creating&#x60;. The volume is being created and is not yet available for use.   - &#x60;active&#x60;. The volume is online and available for use.   - &#x60;resizing&#x60;. The volume is in the process of upgrading its current capacity.   - &#x60;key_rotating&#x60;. The volume is in the process of rotating encryption keys. Requests to resize, delete, or clone a volume fails during encryption key rotation. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | An array of Tags applied to this object.  Tags are for organizational purposes only. | [optional] 
**Type** | Pointer to **string** | This is the [Linode type](https://techdocs.akamai.com/linode-api/reference/get-linode-types) that this Linode was deployed with. To change a Linode&#39;s type, use [Resize a Linode](https://techdocs.akamai.com/linode-api/reference/post-resize-linode-instance). | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this NodeBalancer was last updated. | [optional] [readonly] 
**WatchdogEnabled** | Pointer to **bool** | The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and reboots it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie gives up if there have been more than 5 boot jobs issued within 15 minutes. | [optional] 
**AxfrIps** | Pointer to **[]string** | The list of IPs that may perform a zone transfer for this Domain. The total combined length of all data within this array cannot exceed 1000 characters.  __Note__. This is potentially dangerous, and should be set to an empty list unless you intend to use it. | [optional] 
**Description** | Pointer to **string** | A description for this Domain. This is for display purposes only. | [optional] 
**Domain** | Pointer to **string** | The domain this Domain represents. Domain labels cannot be longer than 63 characters and must conform to [RFC1035](https://tools.ietf.org/html/rfc1035). Domains must be unique on Linode&#39;s platform, including across different Linode accounts; there cannot be two Domains representing the same domain. | [optional] 
**ExpireSec** | Pointer to **int32** | The amount of time in seconds that may pass before this Domain is no longer authoritative.  - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200.  - Any other value is rounded up to the nearest valid value.  - A value of 0 is equivalent to the default value of 1209600. | [optional] [default to 0]
**MasterIps** | Pointer to **[]string** | The IP addresses representing the master DNS for this Domain. At least one value is required for &#x60;type&#x60; slave Domains. The total combined length of all data within this array cannot exceed 1000 characters. | [optional] 
**RefreshSec** | Pointer to **int32** | The amount of time in seconds before this Domain should be refreshed.  - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200.  - Any other value is rounded up to the nearest valid value.  - A value of 0 is equivalent to the default value of 14400. | [optional] [default to 0]
**RetrySec** | Pointer to **int32** | The interval, in seconds, at which a failed refresh should be retried.  - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200.  - Any other value is rounded up to the nearest valid value.  - A value of 0 is equivalent to the default value of 14400. | [optional] [default to 0]
**SoaEmail** | Pointer to **string** | Start of Authority email address. This is required for &#x60;type&#x60; master Domains. | [optional] 
**TtlSec** | Pointer to **int32** | \&quot;Time to Live\&quot; - the amount of time in seconds that this Domain&#39;s records may be cached by resolvers or other domain servers.  - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200. - Any other value is rounded up to the nearest valid value. - A value of 0 is equivalent to the default value of 86400. | [optional] [default to 0]
**Encryption** | Pointer to **string** | Displays if encryption is enabled on this volume. | [optional] [readonly] 
**FilesystemPath** | Pointer to **string** | The full filesystem path for the Volume based on the Volume&#39;s label. Path is /dev/disk/by-id/scsi-0Linode_Volume_ + Volume label. | [optional] [readonly] 
**HardwareType** | Pointer to **string** | The storage type of this Volume. | [optional] [readonly] 
**LinodeId** | Pointer to **NullableInt32** | If a Volume is attached to a specific Linode, the ID of that Linode will be displayed here. | [optional] 
**LinodeLabel** | Pointer to **NullableString** | If a Volume is attached to a specific Linode, the label of that Linode will be displayed here. | [optional] [readonly] 
**Size** | Pointer to **int32** | The Volume&#39;s size, in GiB. | [optional] 
**ClientConnThrottle** | Pointer to **int32** | Throttle connections per second.  Set to 0 (zero) to disable throttling. | [optional] 
**Hostname** | Pointer to **string** | This NodeBalancer&#39;s hostname, beginning with its IP address and ending with _.ip.linodeusercontent.com_. | [optional] [readonly] 
**Transfer** | Pointer to [**NodeBalancerTransfer**](NodeBalancerTransfer.md) |  | [optional] 

## Methods

### NewGetTaggedObjects200ResponseDataInnerData

`func NewGetTaggedObjects200ResponseDataInnerData() *GetTaggedObjects200ResponseDataInnerData`

NewGetTaggedObjects200ResponseDataInnerData instantiates a new GetTaggedObjects200ResponseDataInnerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaggedObjects200ResponseDataInnerDataWithDefaults

`func NewGetTaggedObjects200ResponseDataInnerDataWithDefaults() *GetTaggedObjects200ResponseDataInnerData`

NewGetTaggedObjects200ResponseDataInnerDataWithDefaults instantiates a new GetTaggedObjects200ResponseDataInnerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *GetTaggedObjects200ResponseDataInnerData) GetAlerts() LinodeAlerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetAlertsOk() (*LinodeAlerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *GetTaggedObjects200ResponseDataInnerData) SetAlerts(v LinodeAlerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *GetTaggedObjects200ResponseDataInnerData) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetBackups

`func (o *GetTaggedObjects200ResponseDataInnerData) GetBackups() LinodeBackups`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetBackupsOk() (*LinodeBackups, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *GetTaggedObjects200ResponseDataInnerData) SetBackups(v LinodeBackups)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *GetTaggedObjects200ResponseDataInnerData) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetCapabilities

`func (o *GetTaggedObjects200ResponseDataInnerData) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetTaggedObjects200ResponseDataInnerData) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *GetTaggedObjects200ResponseDataInnerData) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreated

`func (o *GetTaggedObjects200ResponseDataInnerData) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetTaggedObjects200ResponseDataInnerData) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetTaggedObjects200ResponseDataInnerData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *GetTaggedObjects200ResponseDataInnerData) GetDiskEncryption() string`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetDiskEncryptionOk() (*string, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *GetTaggedObjects200ResponseDataInnerData) SetDiskEncryption(v string)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *GetTaggedObjects200ResponseDataInnerData) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### SetDiskEncryptionNil

`func (o *GetTaggedObjects200ResponseDataInnerData) SetDiskEncryptionNil(b bool)`

 SetDiskEncryptionNil sets the value for DiskEncryption to be an explicit nil

### UnsetDiskEncryption
`func (o *GetTaggedObjects200ResponseDataInnerData) UnsetDiskEncryption()`

UnsetDiskEncryption ensures that no value is present for DiskEncryption, not even an explicit nil
### GetGroup

`func (o *GetTaggedObjects200ResponseDataInnerData) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetTaggedObjects200ResponseDataInnerData) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetTaggedObjects200ResponseDataInnerData) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHasUserData

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHasUserData() bool`

GetHasUserData returns the HasUserData field if non-nil, zero value otherwise.

### GetHasUserDataOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHasUserDataOk() (*bool, bool)`

GetHasUserDataOk returns a tuple with the HasUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUserData

`func (o *GetTaggedObjects200ResponseDataInnerData) SetHasUserData(v bool)`

SetHasUserData sets HasUserData field to given value.

### HasHasUserData

`func (o *GetTaggedObjects200ResponseDataInnerData) HasHasUserData() bool`

HasHasUserData returns a boolean if a field has been set.

### GetHostUuid

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *GetTaggedObjects200ResponseDataInnerData) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.

### HasHostUuid

`func (o *GetTaggedObjects200ResponseDataInnerData) HasHostUuid() bool`

HasHostUuid returns a boolean if a field has been set.

### GetHypervisor

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *GetTaggedObjects200ResponseDataInnerData) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *GetTaggedObjects200ResponseDataInnerData) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetId

`func (o *GetTaggedObjects200ResponseDataInnerData) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTaggedObjects200ResponseDataInnerData) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTaggedObjects200ResponseDataInnerData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *GetTaggedObjects200ResponseDataInnerData) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GetTaggedObjects200ResponseDataInnerData) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *GetTaggedObjects200ResponseDataInnerData) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIpv4

`func (o *GetTaggedObjects200ResponseDataInnerData) GetIpv4() string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *GetTaggedObjects200ResponseDataInnerData) GetIpv4Ok() (*string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *GetTaggedObjects200ResponseDataInnerData) SetIpv4(v string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *GetTaggedObjects200ResponseDataInnerData) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *GetTaggedObjects200ResponseDataInnerData) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *GetTaggedObjects200ResponseDataInnerData) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *GetTaggedObjects200ResponseDataInnerData) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *GetTaggedObjects200ResponseDataInnerData) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### SetIpv6Nil

`func (o *GetTaggedObjects200ResponseDataInnerData) SetIpv6Nil(b bool)`

 SetIpv6Nil sets the value for Ipv6 to be an explicit nil

### UnsetIpv6
`func (o *GetTaggedObjects200ResponseDataInnerData) UnsetIpv6()`

UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
### GetLabel

`func (o *GetTaggedObjects200ResponseDataInnerData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetTaggedObjects200ResponseDataInnerData) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetTaggedObjects200ResponseDataInnerData) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLkeClusterId

`func (o *GetTaggedObjects200ResponseDataInnerData) GetLkeClusterId() int32`

GetLkeClusterId returns the LkeClusterId field if non-nil, zero value otherwise.

### GetLkeClusterIdOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetLkeClusterIdOk() (*int32, bool)`

GetLkeClusterIdOk returns a tuple with the LkeClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLkeClusterId

`func (o *GetTaggedObjects200ResponseDataInnerData) SetLkeClusterId(v int32)`

SetLkeClusterId sets LkeClusterId field to given value.

### HasLkeClusterId

`func (o *GetTaggedObjects200ResponseDataInnerData) HasLkeClusterId() bool`

HasLkeClusterId returns a boolean if a field has been set.

### SetLkeClusterIdNil

`func (o *GetTaggedObjects200ResponseDataInnerData) SetLkeClusterIdNil(b bool)`

 SetLkeClusterIdNil sets the value for LkeClusterId to be an explicit nil

### UnsetLkeClusterId
`func (o *GetTaggedObjects200ResponseDataInnerData) UnsetLkeClusterId()`

UnsetLkeClusterId ensures that no value is present for LkeClusterId, not even an explicit nil
### GetPlacementGroup

`func (o *GetTaggedObjects200ResponseDataInnerData) GetPlacementGroup() Linode1PlacementGroup`

GetPlacementGroup returns the PlacementGroup field if non-nil, zero value otherwise.

### GetPlacementGroupOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetPlacementGroupOk() (*Linode1PlacementGroup, bool)`

GetPlacementGroupOk returns a tuple with the PlacementGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroup

`func (o *GetTaggedObjects200ResponseDataInnerData) SetPlacementGroup(v Linode1PlacementGroup)`

SetPlacementGroup sets PlacementGroup field to given value.

### HasPlacementGroup

`func (o *GetTaggedObjects200ResponseDataInnerData) HasPlacementGroup() bool`

HasPlacementGroup returns a boolean if a field has been set.

### SetPlacementGroupNil

`func (o *GetTaggedObjects200ResponseDataInnerData) SetPlacementGroupNil(b bool)`

 SetPlacementGroupNil sets the value for PlacementGroup to be an explicit nil

### UnsetPlacementGroup
`func (o *GetTaggedObjects200ResponseDataInnerData) UnsetPlacementGroup()`

UnsetPlacementGroup ensures that no value is present for PlacementGroup, not even an explicit nil
### GetRegion

`func (o *GetTaggedObjects200ResponseDataInnerData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetTaggedObjects200ResponseDataInnerData) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetTaggedObjects200ResponseDataInnerData) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSpecs

`func (o *GetTaggedObjects200ResponseDataInnerData) GetSpecs() LinodeSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetSpecsOk() (*LinodeSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *GetTaggedObjects200ResponseDataInnerData) SetSpecs(v LinodeSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *GetTaggedObjects200ResponseDataInnerData) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetStatus

`func (o *GetTaggedObjects200ResponseDataInnerData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTaggedObjects200ResponseDataInnerData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTaggedObjects200ResponseDataInnerData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetTaggedObjects200ResponseDataInnerData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetTaggedObjects200ResponseDataInnerData) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetTaggedObjects200ResponseDataInnerData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *GetTaggedObjects200ResponseDataInnerData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTaggedObjects200ResponseDataInnerData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTaggedObjects200ResponseDataInnerData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *GetTaggedObjects200ResponseDataInnerData) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetTaggedObjects200ResponseDataInnerData) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetTaggedObjects200ResponseDataInnerData) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetWatchdogEnabled

`func (o *GetTaggedObjects200ResponseDataInnerData) GetWatchdogEnabled() bool`

GetWatchdogEnabled returns the WatchdogEnabled field if non-nil, zero value otherwise.

### GetWatchdogEnabledOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetWatchdogEnabledOk() (*bool, bool)`

GetWatchdogEnabledOk returns a tuple with the WatchdogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdogEnabled

`func (o *GetTaggedObjects200ResponseDataInnerData) SetWatchdogEnabled(v bool)`

SetWatchdogEnabled sets WatchdogEnabled field to given value.

### HasWatchdogEnabled

`func (o *GetTaggedObjects200ResponseDataInnerData) HasWatchdogEnabled() bool`

HasWatchdogEnabled returns a boolean if a field has been set.

### GetAxfrIps

`func (o *GetTaggedObjects200ResponseDataInnerData) GetAxfrIps() []string`

GetAxfrIps returns the AxfrIps field if non-nil, zero value otherwise.

### GetAxfrIpsOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetAxfrIpsOk() (*[]string, bool)`

GetAxfrIpsOk returns a tuple with the AxfrIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxfrIps

`func (o *GetTaggedObjects200ResponseDataInnerData) SetAxfrIps(v []string)`

SetAxfrIps sets AxfrIps field to given value.

### HasAxfrIps

`func (o *GetTaggedObjects200ResponseDataInnerData) HasAxfrIps() bool`

HasAxfrIps returns a boolean if a field has been set.

### GetDescription

`func (o *GetTaggedObjects200ResponseDataInnerData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTaggedObjects200ResponseDataInnerData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTaggedObjects200ResponseDataInnerData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *GetTaggedObjects200ResponseDataInnerData) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetTaggedObjects200ResponseDataInnerData) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetTaggedObjects200ResponseDataInnerData) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExpireSec

`func (o *GetTaggedObjects200ResponseDataInnerData) GetExpireSec() int32`

GetExpireSec returns the ExpireSec field if non-nil, zero value otherwise.

### GetExpireSecOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetExpireSecOk() (*int32, bool)`

GetExpireSecOk returns a tuple with the ExpireSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireSec

`func (o *GetTaggedObjects200ResponseDataInnerData) SetExpireSec(v int32)`

SetExpireSec sets ExpireSec field to given value.

### HasExpireSec

`func (o *GetTaggedObjects200ResponseDataInnerData) HasExpireSec() bool`

HasExpireSec returns a boolean if a field has been set.

### GetMasterIps

`func (o *GetTaggedObjects200ResponseDataInnerData) GetMasterIps() []string`

GetMasterIps returns the MasterIps field if non-nil, zero value otherwise.

### GetMasterIpsOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetMasterIpsOk() (*[]string, bool)`

GetMasterIpsOk returns a tuple with the MasterIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterIps

`func (o *GetTaggedObjects200ResponseDataInnerData) SetMasterIps(v []string)`

SetMasterIps sets MasterIps field to given value.

### HasMasterIps

`func (o *GetTaggedObjects200ResponseDataInnerData) HasMasterIps() bool`

HasMasterIps returns a boolean if a field has been set.

### GetRefreshSec

`func (o *GetTaggedObjects200ResponseDataInnerData) GetRefreshSec() int32`

GetRefreshSec returns the RefreshSec field if non-nil, zero value otherwise.

### GetRefreshSecOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetRefreshSecOk() (*int32, bool)`

GetRefreshSecOk returns a tuple with the RefreshSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshSec

`func (o *GetTaggedObjects200ResponseDataInnerData) SetRefreshSec(v int32)`

SetRefreshSec sets RefreshSec field to given value.

### HasRefreshSec

`func (o *GetTaggedObjects200ResponseDataInnerData) HasRefreshSec() bool`

HasRefreshSec returns a boolean if a field has been set.

### GetRetrySec

`func (o *GetTaggedObjects200ResponseDataInnerData) GetRetrySec() int32`

GetRetrySec returns the RetrySec field if non-nil, zero value otherwise.

### GetRetrySecOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetRetrySecOk() (*int32, bool)`

GetRetrySecOk returns a tuple with the RetrySec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrySec

`func (o *GetTaggedObjects200ResponseDataInnerData) SetRetrySec(v int32)`

SetRetrySec sets RetrySec field to given value.

### HasRetrySec

`func (o *GetTaggedObjects200ResponseDataInnerData) HasRetrySec() bool`

HasRetrySec returns a boolean if a field has been set.

### GetSoaEmail

`func (o *GetTaggedObjects200ResponseDataInnerData) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *GetTaggedObjects200ResponseDataInnerData) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *GetTaggedObjects200ResponseDataInnerData) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetTtlSec

`func (o *GetTaggedObjects200ResponseDataInnerData) GetTtlSec() int32`

GetTtlSec returns the TtlSec field if non-nil, zero value otherwise.

### GetTtlSecOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetTtlSecOk() (*int32, bool)`

GetTtlSecOk returns a tuple with the TtlSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSec

`func (o *GetTaggedObjects200ResponseDataInnerData) SetTtlSec(v int32)`

SetTtlSec sets TtlSec field to given value.

### HasTtlSec

`func (o *GetTaggedObjects200ResponseDataInnerData) HasTtlSec() bool`

HasTtlSec returns a boolean if a field has been set.

### GetEncryption

`func (o *GetTaggedObjects200ResponseDataInnerData) GetEncryption() string`

GetEncryption returns the Encryption field if non-nil, zero value otherwise.

### GetEncryptionOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetEncryptionOk() (*string, bool)`

GetEncryptionOk returns a tuple with the Encryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryption

`func (o *GetTaggedObjects200ResponseDataInnerData) SetEncryption(v string)`

SetEncryption sets Encryption field to given value.

### HasEncryption

`func (o *GetTaggedObjects200ResponseDataInnerData) HasEncryption() bool`

HasEncryption returns a boolean if a field has been set.

### GetFilesystemPath

`func (o *GetTaggedObjects200ResponseDataInnerData) GetFilesystemPath() string`

GetFilesystemPath returns the FilesystemPath field if non-nil, zero value otherwise.

### GetFilesystemPathOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetFilesystemPathOk() (*string, bool)`

GetFilesystemPathOk returns a tuple with the FilesystemPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystemPath

`func (o *GetTaggedObjects200ResponseDataInnerData) SetFilesystemPath(v string)`

SetFilesystemPath sets FilesystemPath field to given value.

### HasFilesystemPath

`func (o *GetTaggedObjects200ResponseDataInnerData) HasFilesystemPath() bool`

HasFilesystemPath returns a boolean if a field has been set.

### GetHardwareType

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHardwareType() string`

GetHardwareType returns the HardwareType field if non-nil, zero value otherwise.

### GetHardwareTypeOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHardwareTypeOk() (*string, bool)`

GetHardwareTypeOk returns a tuple with the HardwareType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareType

`func (o *GetTaggedObjects200ResponseDataInnerData) SetHardwareType(v string)`

SetHardwareType sets HardwareType field to given value.

### HasHardwareType

`func (o *GetTaggedObjects200ResponseDataInnerData) HasHardwareType() bool`

HasHardwareType returns a boolean if a field has been set.

### GetLinodeId

`func (o *GetTaggedObjects200ResponseDataInnerData) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *GetTaggedObjects200ResponseDataInnerData) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *GetTaggedObjects200ResponseDataInnerData) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### SetLinodeIdNil

`func (o *GetTaggedObjects200ResponseDataInnerData) SetLinodeIdNil(b bool)`

 SetLinodeIdNil sets the value for LinodeId to be an explicit nil

### UnsetLinodeId
`func (o *GetTaggedObjects200ResponseDataInnerData) UnsetLinodeId()`

UnsetLinodeId ensures that no value is present for LinodeId, not even an explicit nil
### GetLinodeLabel

`func (o *GetTaggedObjects200ResponseDataInnerData) GetLinodeLabel() string`

GetLinodeLabel returns the LinodeLabel field if non-nil, zero value otherwise.

### GetLinodeLabelOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetLinodeLabelOk() (*string, bool)`

GetLinodeLabelOk returns a tuple with the LinodeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeLabel

`func (o *GetTaggedObjects200ResponseDataInnerData) SetLinodeLabel(v string)`

SetLinodeLabel sets LinodeLabel field to given value.

### HasLinodeLabel

`func (o *GetTaggedObjects200ResponseDataInnerData) HasLinodeLabel() bool`

HasLinodeLabel returns a boolean if a field has been set.

### SetLinodeLabelNil

`func (o *GetTaggedObjects200ResponseDataInnerData) SetLinodeLabelNil(b bool)`

 SetLinodeLabelNil sets the value for LinodeLabel to be an explicit nil

### UnsetLinodeLabel
`func (o *GetTaggedObjects200ResponseDataInnerData) UnsetLinodeLabel()`

UnsetLinodeLabel ensures that no value is present for LinodeLabel, not even an explicit nil
### GetSize

`func (o *GetTaggedObjects200ResponseDataInnerData) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetTaggedObjects200ResponseDataInnerData) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetTaggedObjects200ResponseDataInnerData) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetClientConnThrottle

`func (o *GetTaggedObjects200ResponseDataInnerData) GetClientConnThrottle() int32`

GetClientConnThrottle returns the ClientConnThrottle field if non-nil, zero value otherwise.

### GetClientConnThrottleOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetClientConnThrottleOk() (*int32, bool)`

GetClientConnThrottleOk returns a tuple with the ClientConnThrottle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnThrottle

`func (o *GetTaggedObjects200ResponseDataInnerData) SetClientConnThrottle(v int32)`

SetClientConnThrottle sets ClientConnThrottle field to given value.

### HasClientConnThrottle

`func (o *GetTaggedObjects200ResponseDataInnerData) HasClientConnThrottle() bool`

HasClientConnThrottle returns a boolean if a field has been set.

### GetHostname

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetTaggedObjects200ResponseDataInnerData) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetTaggedObjects200ResponseDataInnerData) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetTransfer

`func (o *GetTaggedObjects200ResponseDataInnerData) GetTransfer() NodeBalancerTransfer`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *GetTaggedObjects200ResponseDataInnerData) GetTransferOk() (*NodeBalancerTransfer, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *GetTaggedObjects200ResponseDataInnerData) SetTransfer(v NodeBalancerTransfer)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *GetTaggedObjects200ResponseDataInnerData) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


