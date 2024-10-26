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
	"time"
)

// checks if the Linode1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Linode1{}

// Linode1 struct for Linode1
type Linode1 struct {
	Alerts *LinodeAlerts `json:"alerts,omitempty"`
	Backups *LinodeBackups `json:"backups,omitempty"`
	// A list of capabilities this compute instance supports.
	Capabilities []string `json:"capabilities,omitempty"`
	// When this Linode was created.
	Created *time.Time `json:"created,omitempty"`
	// Indicates the local disk encryption setting for this Linode. If the Linode is part of an LKE cluster, the value is `null`.
	DiskEncryption NullableString `json:"disk_encryption,omitempty"`
	// The group label for this Linode.
	// Deprecated
	Group *string `json:"group,omitempty"`
	// Whether this compute instance was provisioned with `user_data` provided via the Metadata service. See the [Create a Linode](https://techdocs.akamai.com/linode-api/reference/post-linode-instance) description for more information on Metadata.
	HasUserData *bool `json:"has_user_data,omitempty"`
	// The Linode's host machine, as a UUID.
	HostUuid *string `json:"host_uuid,omitempty"`
	// The virtualization software powering this Linode.
	Hypervisor *string `json:"hypervisor,omitempty"`
	// This Linode's ID which must be provided for all operations impacting this Linode.
	Id *int32 `json:"id,omitempty"`
	// An Image ID to deploy the Linode Disk from.  Run the [List images](https://techdocs.akamai.com/linode-api/reference/get-images) operation with authentication to view all available Images. Official Linode Images start with `linode/`, while your Account's Images start with `private/`. Creating a disk from a Private Image requires `read_only` or `read_write` permissions for that Image. Run the [Update a user's grants](https://techdocs.akamai.com/linode-api/reference/put-user-grants) operation to adjust permissions for an Account Image.
	Image *string `json:"image,omitempty"`
	// This Linode's IPv4 Addresses. Each Linode is assigned a single public IPv4 address upon creation, and may get a single private IPv4 address if needed. You may need to [Open a support ticket](https://techdocs.akamai.com/linode-api/reference/post-ticket) to get additional IPv4 addresses.  IPv4 addresses may be reassigned between your Linodes, or shared with other Linodes. See the [networking](https://techdocs.akamai.com/linode-api/reference/post-firewalls) operations for details.
	Ipv4 []string `json:"ipv4,omitempty"`
	// This Linode's IPv6 SLAAC address. This address is specific to a Linode, and may not be shared. If the Linode has not been assigned an IPv6 address, the return value will be `null`.
	Ipv6 NullableString `json:"ipv6,omitempty"`
	// Provides a name for the Linode. If not provided, the API generates one for it.  Linode labels have the following constraints:  - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (`-`), underscores (`_`) or periods (`.`). - Cannot have two hyphens (`--`), underscores (`__`) or periods (`..`) in a row.
	Label *string `json:"label,omitempty" validate:"regexp=^[a-zA-Z]((?!--|__|\\\\.\\\\.)[a-zA-Z0-9-_.])+$"`
	// The ID of the Kubernetes cluster if the Linode is part of cluster.
	LkeClusterId NullableInt32 `json:"lke_cluster_id,omitempty"`
	PlacementGroup NullableLinode1PlacementGroup `json:"placement_group,omitempty"`
	// The [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where the Linode deployed. A Linode's region can only be changed by initiating a [cross data center migration](https://techdocs.akamai.com/linode-api/reference/post-migrate-linode-instance).
	Region *string `json:"region,omitempty"`
	Specs *LinodeSpecs `json:"specs,omitempty"`
	// A brief description of this Linode's current state. This field may change without direct action from you. For example, when a compute instance goes into maintenance mode, its status is `stopped`. Status is generally self-explanatory, based on its name.  - `busy` indicates you've assigned the compute instance to a [placement group](https://techdocs.akamai.com/cloud-computing/docs/work-with-placement-groups), but the compute instance is currently booting. Once the boot completes, the API completes the assignment and updates the compute instance's `status` accordingly. - `provisioning` indicates that the API is applying operating system or Marketplace applications on the compute instance. - `billing_suspension` indicates that payment is past due on the compute instance, so we've suspended its use.
	Status *string `json:"status,omitempty"`
	// Tags to help you organize your content.
	Tags []string `json:"tags,omitempty"`
	// This is the [Linode type](https://techdocs.akamai.com/linode-api/reference/get-linode-types) that this Linode was deployed with. To change a Linode's type, use [Resize a Linode](https://techdocs.akamai.com/linode-api/reference/post-resize-linode-instance).
	Type *string `json:"type,omitempty"`
	// When this Linode was last updated.
	Updated *time.Time `json:"updated,omitempty"`
	// The watchdog, named Lassie, is a Shutdown Watchdog that monitors your Linode and reboots it if it powers off unexpectedly. It works by issuing a boot job when your Linode powers off without a shutdown job being responsible. To prevent a loop, Lassie gives up if there have been more than 5 boot jobs issued within 15 minutes.
	WatchdogEnabled *bool `json:"watchdog_enabled,omitempty"`
}

// NewLinode1 instantiates a new Linode1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinode1() *Linode1 {
	this := Linode1{}
	return &this
}

// NewLinode1WithDefaults instantiates a new Linode1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinode1WithDefaults() *Linode1 {
	this := Linode1{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *Linode1) GetAlerts() LinodeAlerts {
	if o == nil || IsNil(o.Alerts) {
		var ret LinodeAlerts
		return ret
	}
	return *o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetAlertsOk() (*LinodeAlerts, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *Linode1) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given LinodeAlerts and assigns it to the Alerts field.
func (o *Linode1) SetAlerts(v LinodeAlerts) {
	o.Alerts = &v
}

// GetBackups returns the Backups field value if set, zero value otherwise.
func (o *Linode1) GetBackups() LinodeBackups {
	if o == nil || IsNil(o.Backups) {
		var ret LinodeBackups
		return ret
	}
	return *o.Backups
}

// GetBackupsOk returns a tuple with the Backups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetBackupsOk() (*LinodeBackups, bool) {
	if o == nil || IsNil(o.Backups) {
		return nil, false
	}
	return o.Backups, true
}

// HasBackups returns a boolean if a field has been set.
func (o *Linode1) HasBackups() bool {
	if o != nil && !IsNil(o.Backups) {
		return true
	}

	return false
}

// SetBackups gets a reference to the given LinodeBackups and assigns it to the Backups field.
func (o *Linode1) SetBackups(v LinodeBackups) {
	o.Backups = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *Linode1) GetCapabilities() []string {
	if o == nil || IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *Linode1) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *Linode1) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *Linode1) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *Linode1) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *Linode1) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDiskEncryption returns the DiskEncryption field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Linode1) GetDiskEncryption() string {
	if o == nil || IsNil(o.DiskEncryption.Get()) {
		var ret string
		return ret
	}
	return *o.DiskEncryption.Get()
}

// GetDiskEncryptionOk returns a tuple with the DiskEncryption field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Linode1) GetDiskEncryptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DiskEncryption.Get(), o.DiskEncryption.IsSet()
}

// HasDiskEncryption returns a boolean if a field has been set.
func (o *Linode1) HasDiskEncryption() bool {
	if o != nil && o.DiskEncryption.IsSet() {
		return true
	}

	return false
}

// SetDiskEncryption gets a reference to the given NullableString and assigns it to the DiskEncryption field.
func (o *Linode1) SetDiskEncryption(v string) {
	o.DiskEncryption.Set(&v)
}
// SetDiskEncryptionNil sets the value for DiskEncryption to be an explicit nil
func (o *Linode1) SetDiskEncryptionNil() {
	o.DiskEncryption.Set(nil)
}

// UnsetDiskEncryption ensures that no value is present for DiskEncryption, not even an explicit nil
func (o *Linode1) UnsetDiskEncryption() {
	o.DiskEncryption.Unset()
}

// GetGroup returns the Group field value if set, zero value otherwise.
// Deprecated
func (o *Linode1) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Linode1) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *Linode1) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
// Deprecated
func (o *Linode1) SetGroup(v string) {
	o.Group = &v
}

// GetHasUserData returns the HasUserData field value if set, zero value otherwise.
func (o *Linode1) GetHasUserData() bool {
	if o == nil || IsNil(o.HasUserData) {
		var ret bool
		return ret
	}
	return *o.HasUserData
}

// GetHasUserDataOk returns a tuple with the HasUserData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetHasUserDataOk() (*bool, bool) {
	if o == nil || IsNil(o.HasUserData) {
		return nil, false
	}
	return o.HasUserData, true
}

// HasHasUserData returns a boolean if a field has been set.
func (o *Linode1) HasHasUserData() bool {
	if o != nil && !IsNil(o.HasUserData) {
		return true
	}

	return false
}

// SetHasUserData gets a reference to the given bool and assigns it to the HasUserData field.
func (o *Linode1) SetHasUserData(v bool) {
	o.HasUserData = &v
}

// GetHostUuid returns the HostUuid field value if set, zero value otherwise.
func (o *Linode1) GetHostUuid() string {
	if o == nil || IsNil(o.HostUuid) {
		var ret string
		return ret
	}
	return *o.HostUuid
}

// GetHostUuidOk returns a tuple with the HostUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetHostUuidOk() (*string, bool) {
	if o == nil || IsNil(o.HostUuid) {
		return nil, false
	}
	return o.HostUuid, true
}

// HasHostUuid returns a boolean if a field has been set.
func (o *Linode1) HasHostUuid() bool {
	if o != nil && !IsNil(o.HostUuid) {
		return true
	}

	return false
}

// SetHostUuid gets a reference to the given string and assigns it to the HostUuid field.
func (o *Linode1) SetHostUuid(v string) {
	o.HostUuid = &v
}

// GetHypervisor returns the Hypervisor field value if set, zero value otherwise.
func (o *Linode1) GetHypervisor() string {
	if o == nil || IsNil(o.Hypervisor) {
		var ret string
		return ret
	}
	return *o.Hypervisor
}

// GetHypervisorOk returns a tuple with the Hypervisor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetHypervisorOk() (*string, bool) {
	if o == nil || IsNil(o.Hypervisor) {
		return nil, false
	}
	return o.Hypervisor, true
}

// HasHypervisor returns a boolean if a field has been set.
func (o *Linode1) HasHypervisor() bool {
	if o != nil && !IsNil(o.Hypervisor) {
		return true
	}

	return false
}

// SetHypervisor gets a reference to the given string and assigns it to the Hypervisor field.
func (o *Linode1) SetHypervisor(v string) {
	o.Hypervisor = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Linode1) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Linode1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Linode1) SetId(v int32) {
	o.Id = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *Linode1) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *Linode1) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *Linode1) SetImage(v string) {
	o.Image = &v
}

// GetIpv4 returns the Ipv4 field value if set, zero value otherwise.
func (o *Linode1) GetIpv4() []string {
	if o == nil || IsNil(o.Ipv4) {
		var ret []string
		return ret
	}
	return o.Ipv4
}

// GetIpv4Ok returns a tuple with the Ipv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetIpv4Ok() ([]string, bool) {
	if o == nil || IsNil(o.Ipv4) {
		return nil, false
	}
	return o.Ipv4, true
}

// HasIpv4 returns a boolean if a field has been set.
func (o *Linode1) HasIpv4() bool {
	if o != nil && !IsNil(o.Ipv4) {
		return true
	}

	return false
}

// SetIpv4 gets a reference to the given []string and assigns it to the Ipv4 field.
func (o *Linode1) SetIpv4(v []string) {
	o.Ipv4 = v
}

// GetIpv6 returns the Ipv6 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Linode1) GetIpv6() string {
	if o == nil || IsNil(o.Ipv6.Get()) {
		var ret string
		return ret
	}
	return *o.Ipv6.Get()
}

// GetIpv6Ok returns a tuple with the Ipv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Linode1) GetIpv6Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Ipv6.Get(), o.Ipv6.IsSet()
}

// HasIpv6 returns a boolean if a field has been set.
func (o *Linode1) HasIpv6() bool {
	if o != nil && o.Ipv6.IsSet() {
		return true
	}

	return false
}

// SetIpv6 gets a reference to the given NullableString and assigns it to the Ipv6 field.
func (o *Linode1) SetIpv6(v string) {
	o.Ipv6.Set(&v)
}
// SetIpv6Nil sets the value for Ipv6 to be an explicit nil
func (o *Linode1) SetIpv6Nil() {
	o.Ipv6.Set(nil)
}

// UnsetIpv6 ensures that no value is present for Ipv6, not even an explicit nil
func (o *Linode1) UnsetIpv6() {
	o.Ipv6.Unset()
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *Linode1) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *Linode1) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *Linode1) SetLabel(v string) {
	o.Label = &v
}

// GetLkeClusterId returns the LkeClusterId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Linode1) GetLkeClusterId() int32 {
	if o == nil || IsNil(o.LkeClusterId.Get()) {
		var ret int32
		return ret
	}
	return *o.LkeClusterId.Get()
}

// GetLkeClusterIdOk returns a tuple with the LkeClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Linode1) GetLkeClusterIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LkeClusterId.Get(), o.LkeClusterId.IsSet()
}

// HasLkeClusterId returns a boolean if a field has been set.
func (o *Linode1) HasLkeClusterId() bool {
	if o != nil && o.LkeClusterId.IsSet() {
		return true
	}

	return false
}

// SetLkeClusterId gets a reference to the given NullableInt32 and assigns it to the LkeClusterId field.
func (o *Linode1) SetLkeClusterId(v int32) {
	o.LkeClusterId.Set(&v)
}
// SetLkeClusterIdNil sets the value for LkeClusterId to be an explicit nil
func (o *Linode1) SetLkeClusterIdNil() {
	o.LkeClusterId.Set(nil)
}

// UnsetLkeClusterId ensures that no value is present for LkeClusterId, not even an explicit nil
func (o *Linode1) UnsetLkeClusterId() {
	o.LkeClusterId.Unset()
}

// GetPlacementGroup returns the PlacementGroup field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Linode1) GetPlacementGroup() Linode1PlacementGroup {
	if o == nil || IsNil(o.PlacementGroup.Get()) {
		var ret Linode1PlacementGroup
		return ret
	}
	return *o.PlacementGroup.Get()
}

// GetPlacementGroupOk returns a tuple with the PlacementGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Linode1) GetPlacementGroupOk() (*Linode1PlacementGroup, bool) {
	if o == nil {
		return nil, false
	}
	return o.PlacementGroup.Get(), o.PlacementGroup.IsSet()
}

// HasPlacementGroup returns a boolean if a field has been set.
func (o *Linode1) HasPlacementGroup() bool {
	if o != nil && o.PlacementGroup.IsSet() {
		return true
	}

	return false
}

// SetPlacementGroup gets a reference to the given NullableLinode1PlacementGroup and assigns it to the PlacementGroup field.
func (o *Linode1) SetPlacementGroup(v Linode1PlacementGroup) {
	o.PlacementGroup.Set(&v)
}
// SetPlacementGroupNil sets the value for PlacementGroup to be an explicit nil
func (o *Linode1) SetPlacementGroupNil() {
	o.PlacementGroup.Set(nil)
}

// UnsetPlacementGroup ensures that no value is present for PlacementGroup, not even an explicit nil
func (o *Linode1) UnsetPlacementGroup() {
	o.PlacementGroup.Unset()
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *Linode1) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *Linode1) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *Linode1) SetRegion(v string) {
	o.Region = &v
}

// GetSpecs returns the Specs field value if set, zero value otherwise.
func (o *Linode1) GetSpecs() LinodeSpecs {
	if o == nil || IsNil(o.Specs) {
		var ret LinodeSpecs
		return ret
	}
	return *o.Specs
}

// GetSpecsOk returns a tuple with the Specs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetSpecsOk() (*LinodeSpecs, bool) {
	if o == nil || IsNil(o.Specs) {
		return nil, false
	}
	return o.Specs, true
}

// HasSpecs returns a boolean if a field has been set.
func (o *Linode1) HasSpecs() bool {
	if o != nil && !IsNil(o.Specs) {
		return true
	}

	return false
}

// SetSpecs gets a reference to the given LinodeSpecs and assigns it to the Specs field.
func (o *Linode1) SetSpecs(v LinodeSpecs) {
	o.Specs = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Linode1) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Linode1) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Linode1) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Linode1) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Linode1) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *Linode1) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Linode1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Linode1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Linode1) SetType(v string) {
	o.Type = &v
}

// GetUpdated returns the Updated field value if set, zero value otherwise.
func (o *Linode1) GetUpdated() time.Time {
	if o == nil || IsNil(o.Updated) {
		var ret time.Time
		return ret
	}
	return *o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Updated) {
		return nil, false
	}
	return o.Updated, true
}

// HasUpdated returns a boolean if a field has been set.
func (o *Linode1) HasUpdated() bool {
	if o != nil && !IsNil(o.Updated) {
		return true
	}

	return false
}

// SetUpdated gets a reference to the given time.Time and assigns it to the Updated field.
func (o *Linode1) SetUpdated(v time.Time) {
	o.Updated = &v
}

// GetWatchdogEnabled returns the WatchdogEnabled field value if set, zero value otherwise.
func (o *Linode1) GetWatchdogEnabled() bool {
	if o == nil || IsNil(o.WatchdogEnabled) {
		var ret bool
		return ret
	}
	return *o.WatchdogEnabled
}

// GetWatchdogEnabledOk returns a tuple with the WatchdogEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Linode1) GetWatchdogEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.WatchdogEnabled) {
		return nil, false
	}
	return o.WatchdogEnabled, true
}

// HasWatchdogEnabled returns a boolean if a field has been set.
func (o *Linode1) HasWatchdogEnabled() bool {
	if o != nil && !IsNil(o.WatchdogEnabled) {
		return true
	}

	return false
}

// SetWatchdogEnabled gets a reference to the given bool and assigns it to the WatchdogEnabled field.
func (o *Linode1) SetWatchdogEnabled(v bool) {
	o.WatchdogEnabled = &v
}

func (o Linode1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Linode1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.Backups) {
		toSerialize["backups"] = o.Backups
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if o.DiskEncryption.IsSet() {
		toSerialize["disk_encryption"] = o.DiskEncryption.Get()
	}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.HasUserData) {
		toSerialize["has_user_data"] = o.HasUserData
	}
	if !IsNil(o.HostUuid) {
		toSerialize["host_uuid"] = o.HostUuid
	}
	if !IsNil(o.Hypervisor) {
		toSerialize["hypervisor"] = o.Hypervisor
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Ipv4) {
		toSerialize["ipv4"] = o.Ipv4
	}
	if o.Ipv6.IsSet() {
		toSerialize["ipv6"] = o.Ipv6.Get()
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if o.LkeClusterId.IsSet() {
		toSerialize["lke_cluster_id"] = o.LkeClusterId.Get()
	}
	if o.PlacementGroup.IsSet() {
		toSerialize["placement_group"] = o.PlacementGroup.Get()
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Specs) {
		toSerialize["specs"] = o.Specs
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Updated) {
		toSerialize["updated"] = o.Updated
	}
	if !IsNil(o.WatchdogEnabled) {
		toSerialize["watchdog_enabled"] = o.WatchdogEnabled
	}
	return toSerialize, nil
}

type NullableLinode1 struct {
	value *Linode1
	isSet bool
}

func (v NullableLinode1) Get() *Linode1 {
	return v.value
}

func (v *NullableLinode1) Set(val *Linode1) {
	v.value = val
	v.isSet = true
}

func (v NullableLinode1) IsSet() bool {
	return v.isSet
}

func (v *NullableLinode1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinode1(val *Linode1) *NullableLinode1 {
	return &NullableLinode1{value: val, isSet: true}
}

func (v NullableLinode1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLinode1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


