# Linode2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**LinodeAlerts**](LinodeAlerts.md) |  | [optional] 
**Backups** | Pointer to [**LinodeBackups**](LinodeBackups.md) |  | [optional] 
**Capabilities** | Pointer to **[]string** | A list of capabilities this compute instance supports. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | When this Linode was created. | [optional] [readonly] 
**DiskEncryption** | Pointer to **NullableString** | Indicates the local disk encryption setting for this Linode. If the Linode is part of an LKE cluster, the value is &#x60;null&#x60;. | [optional] [readonly] [default to "enabled"]
**Group** | Pointer to **string** | The group label for this Linode. | [optional] 
**HasUserData** | Pointer to **bool** | Whether this compute instance was provisioned with &#x60;user_data&#x60; provided via the Metadata service. See the [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance) description for more information on Metadata. | [optional] [readonly] 
**HostUuid** | Pointer to **string** | The Linode&#39;s host machine, as a UUID. | [optional] [readonly] 
**Hypervisor** | Pointer to **string** | The virtualization software powering this Linode. | [optional] [readonly] 
**Id** | Pointer to **int32** | This Linode&#39;s ID which must be provided for all operations impacting this Linode. | [optional] [readonly] 
**Image** | Pointer to **string** | An Image ID to deploy the Linode Disk from.  Run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation with authentication to view all available Images. Official Linode Images start with &#x60;linode/&#x60;, while your Account&#39;s Images start with &#x60;private/&#x60;. Creating a disk from a Private Image requires &#x60;read_only&#x60; or &#x60;read_write&#x60; permissions for that Image. Run the [Update a user&#39;s grants](https://techdocs.akamai.com/linode-api/reference/put-user-grants) operation to adjust permissions for an Account Image. | [optional] 
**Ipv4** | Pointer to **[]string** | This Linode&#39;s IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to [Open a support ticket](https://techdocs.akamai.com/linode-api/reference/post-ticket) to get additional IPv4 addresses.  IPv4 addresses may be reassigned between your Linodes, or shared with other Linodes. See the [networking](https://techdocs.akamai.com/linode-api/reference/post-firewalls) operations for details. | [optional] [readonly] 
**Ipv6** | Pointer to **NullableString** | This Linode&#39;s IPv6 SLAAC address. This address is specific to a Linode, and may not be shared. If the Linode has not been assigned an IPv6 address, the return value will be &#x60;null&#x60;. | [optional] [readonly] 
**Label** | Pointer to **string** | Provides a name for the Linode. If not provided, the API generates one for it.  Linode labels have the following constraints:  - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (&#x60;-&#x60;), underscores (&#x60;_&#x60;) or periods (&#x60;.&#x60;). - Cannot have two hyphens (&#x60;--&#x60;), underscores (&#x60;__&#x60;) or periods (&#x60;..&#x60;) in a row. | [optional] 
**LkeClusterId** | Pointer to **NullableInt32** | The ID of the Kubernetes cluster if the Linode is part of cluster. | [optional] [readonly] 
**PlacementGroup** | Pointer to [**NullableLinode1PlacementGroup**](Linode1PlacementGroup.md) |  | [optional] 
**Region** | Pointer to **string** | The [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where the Linode deployed. A Linode&#39;s region can only be changed by initiating a [cross data center migration](https://techdocs.akamai.com/linode-api/reference/post-migrate-linode-instance). | [optional] [readonly] 
**Specs** | Pointer to [**LinodeSpecs**](LinodeSpecs.md) |  | [optional] 
**Status** | Pointer to **string** | A brief description of this Linode&#39;s current state. This field may change without direct action from you. For example, when a compute instance goes into maintenance mode, its status is &#x60;stopped&#x60;. Status is generally self-explanatory, based on its name.  - &#x60;busy&#x60; indicates you&#39;ve assigned the compute instance to a [placement group](https://techdocs.akamai.com/cloud-computing/docs/work-with-placement-groups), but the compute instance is currently booting. Once the boot completes, the API completes the assignment and updates the compute instance&#39;s &#x60;status&#x60; accordingly. - &#x60;provisioning&#x60; indicates that the API is applying operating system or Marketplace applications on the compute instance. - &#x60;billing_suspension&#x60; indicates that payment is past due on the compute instance, so we&#39;ve suspended its use. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | Tags to help you organize your content. | [optional] 
**Type** | Pointer to **string** | This is the [Linode type](https://techdocs.akamai.com/linode-api/reference/get-linode-types) that this Linode was deployed with. To change a Linode&#39;s type, use [Resize a Linode](https://techdocs.akamai.com/linode-api/reference/post-resize-linode-instance). | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this Linode was last updated. | [optional] [readonly] 
**WatchdogEnabled** | Pointer to **bool** | The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and reboots it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie gives up if there have been more than 5 boot jobs issued within 15 minutes. | [optional] 

## Methods

### NewLinode2

`func NewLinode2() *Linode2`

NewLinode2 instantiates a new Linode2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinode2WithDefaults

`func NewLinode2WithDefaults() *Linode2`

NewLinode2WithDefaults instantiates a new Linode2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *Linode2) GetAlerts() LinodeAlerts`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *Linode2) GetAlertsOk() (*LinodeAlerts, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *Linode2) SetAlerts(v LinodeAlerts)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *Linode2) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetBackups

`func (o *Linode2) GetBackups() LinodeBackups`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *Linode2) GetBackupsOk() (*LinodeBackups, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *Linode2) SetBackups(v LinodeBackups)`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *Linode2) HasBackups() bool`

HasBackups returns a boolean if a field has been set.

### GetCapabilities

`func (o *Linode2) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *Linode2) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *Linode2) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *Linode2) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreated

`func (o *Linode2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Linode2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Linode2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Linode2) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDiskEncryption

`func (o *Linode2) GetDiskEncryption() string`

GetDiskEncryption returns the DiskEncryption field if non-nil, zero value otherwise.

### GetDiskEncryptionOk

`func (o *Linode2) GetDiskEncryptionOk() (*string, bool)`

GetDiskEncryptionOk returns a tuple with the DiskEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskEncryption

`func (o *Linode2) SetDiskEncryption(v string)`

SetDiskEncryption sets DiskEncryption field to given value.

### HasDiskEncryption

`func (o *Linode2) HasDiskEncryption() bool`

HasDiskEncryption returns a boolean if a field has been set.

### SetDiskEncryptionNil

`func (o *Linode2) SetDiskEncryptionNil(b bool)`

 SetDiskEncryptionNil sets the value for DiskEncryption to be an explicit nil

### UnsetDiskEncryption
`func (o *Linode2) UnsetDiskEncryption()`

UnsetDiskEncryption ensures that no value is present for DiskEncryption, not even an explicit nil
### GetGroup

`func (o *Linode2) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Linode2) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Linode2) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Linode2) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHasUserData

`func (o *Linode2) GetHasUserData() bool`

GetHasUserData returns the HasUserData field if non-nil, zero value otherwise.

### GetHasUserDataOk

`func (o *Linode2) GetHasUserDataOk() (*bool, bool)`

GetHasUserDataOk returns a tuple with the HasUserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUserData

`func (o *Linode2) SetHasUserData(v bool)`

SetHasUserData sets HasUserData field to given value.

### HasHasUserData

`func (o *Linode2) HasHasUserData() bool`

HasHasUserData returns a boolean if a field has been set.

### GetHostUuid

`func (o *Linode2) GetHostUuid() string`

GetHostUuid returns the HostUuid field if non-nil, zero value otherwise.

### GetHostUuidOk

`func (o *Linode2) GetHostUuidOk() (*string, bool)`

GetHostUuidOk returns a tuple with the HostUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostUuid

`func (o *Linode2) SetHostUuid(v string)`

SetHostUuid sets HostUuid field to given value.

### HasHostUuid

`func (o *Linode2) HasHostUuid() bool`

HasHostUuid returns a boolean if a field has been set.

### GetHypervisor

`func (o *Linode2) GetHypervisor() string`

GetHypervisor returns the Hypervisor field if non-nil, zero value otherwise.

### GetHypervisorOk

`func (o *Linode2) GetHypervisorOk() (*string, bool)`

GetHypervisorOk returns a tuple with the Hypervisor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHypervisor

`func (o *Linode2) SetHypervisor(v string)`

SetHypervisor sets Hypervisor field to given value.

### HasHypervisor

`func (o *Linode2) HasHypervisor() bool`

HasHypervisor returns a boolean if a field has been set.

### GetId

`func (o *Linode2) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Linode2) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Linode2) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Linode2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImage

`func (o *Linode2) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *Linode2) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *Linode2) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *Linode2) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetIpv4

`func (o *Linode2) GetIpv4() []string`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *Linode2) GetIpv4Ok() (*[]string, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *Linode2) SetIpv4(v []string)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *Linode2) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetIpv6

`func (o *Linode2) GetIpv6() string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *Linode2) GetIpv6Ok() (*string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *Linode2) SetIpv6(v string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *Linode2) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### SetIpv6Nil

`func (o *Linode2) SetIpv6Nil(b bool)`

 SetIpv6Nil sets the value for Ipv6 to be an explicit nil

### UnsetIpv6
`func (o *Linode2) UnsetIpv6()`

UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
### GetLabel

`func (o *Linode2) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Linode2) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Linode2) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Linode2) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLkeClusterId

`func (o *Linode2) GetLkeClusterId() int32`

GetLkeClusterId returns the LkeClusterId field if non-nil, zero value otherwise.

### GetLkeClusterIdOk

`func (o *Linode2) GetLkeClusterIdOk() (*int32, bool)`

GetLkeClusterIdOk returns a tuple with the LkeClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLkeClusterId

`func (o *Linode2) SetLkeClusterId(v int32)`

SetLkeClusterId sets LkeClusterId field to given value.

### HasLkeClusterId

`func (o *Linode2) HasLkeClusterId() bool`

HasLkeClusterId returns a boolean if a field has been set.

### SetLkeClusterIdNil

`func (o *Linode2) SetLkeClusterIdNil(b bool)`

 SetLkeClusterIdNil sets the value for LkeClusterId to be an explicit nil

### UnsetLkeClusterId
`func (o *Linode2) UnsetLkeClusterId()`

UnsetLkeClusterId ensures that no value is present for LkeClusterId, not even an explicit nil
### GetPlacementGroup

`func (o *Linode2) GetPlacementGroup() Linode1PlacementGroup`

GetPlacementGroup returns the PlacementGroup field if non-nil, zero value otherwise.

### GetPlacementGroupOk

`func (o *Linode2) GetPlacementGroupOk() (*Linode1PlacementGroup, bool)`

GetPlacementGroupOk returns a tuple with the PlacementGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroup

`func (o *Linode2) SetPlacementGroup(v Linode1PlacementGroup)`

SetPlacementGroup sets PlacementGroup field to given value.

### HasPlacementGroup

`func (o *Linode2) HasPlacementGroup() bool`

HasPlacementGroup returns a boolean if a field has been set.

### SetPlacementGroupNil

`func (o *Linode2) SetPlacementGroupNil(b bool)`

 SetPlacementGroupNil sets the value for PlacementGroup to be an explicit nil

### UnsetPlacementGroup
`func (o *Linode2) UnsetPlacementGroup()`

UnsetPlacementGroup ensures that no value is present for PlacementGroup, not even an explicit nil
### GetRegion

`func (o *Linode2) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Linode2) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Linode2) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Linode2) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSpecs

`func (o *Linode2) GetSpecs() LinodeSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *Linode2) GetSpecsOk() (*LinodeSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *Linode2) SetSpecs(v LinodeSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *Linode2) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetStatus

`func (o *Linode2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Linode2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Linode2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Linode2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Linode2) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Linode2) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Linode2) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Linode2) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *Linode2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Linode2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Linode2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Linode2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *Linode2) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Linode2) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Linode2) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Linode2) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetWatchdogEnabled

`func (o *Linode2) GetWatchdogEnabled() bool`

GetWatchdogEnabled returns the WatchdogEnabled field if non-nil, zero value otherwise.

### GetWatchdogEnabledOk

`func (o *Linode2) GetWatchdogEnabledOk() (*bool, bool)`

GetWatchdogEnabledOk returns a tuple with the WatchdogEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdogEnabled

`func (o *Linode2) SetWatchdogEnabled(v bool)`

SetWatchdogEnabled sets WatchdogEnabled field to given value.

### HasWatchdogEnabled

`func (o *Linode2) HasWatchdogEnabled() bool`

HasWatchdogEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


