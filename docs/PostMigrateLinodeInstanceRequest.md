# PostMigrateLinodeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlacementGroup** | Pointer to [**PostMigrateLinodeInstanceRequestPlacementGroup**](PostMigrateLinodeInstanceRequestPlacementGroup.md) |  | [optional] 
**Region** | Pointer to **string** | The region to which the Linode will be migrated. Must be a valid region slug. A list of regions can be viewed by running the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation. A cross data center migration will cancel a pending migration that has not yet been initiated. A cross data center migration will initiate a &#x60;linode_migrate_datacenter_create&#x60; event. | [optional] 
**Type** | Pointer to **string** | Type of migration used in moving to a new host or Linode type.  &#x60;warm&#x60;: the Linode will not power down until the migration is complete. Warm migrations are not available for DC migrations.  &#x60;cold&#x60;: the Linode will be powered down and migrated. When the migration is complete, the Linode will be powered on. | [optional] [default to "cold"]
**Upgrade** | Pointer to **bool** | When initiating a cross DC migration, setting this value to true will also ensure that the Linode is upgraded to the latest generation of hardware that corresponds to your Linode&#39;s Type, if any free upgrades are available for it. If no free upgrades are available, and this value is set to true, then the endpoint will return a 400 error code and the migration will not be performed. If the data center set in the &#x60;region&#x60; field does not allow upgrades, then the endpoint will return a 400 error code and the migration will not be performed. | [optional] [default to false]

## Methods

### NewPostMigrateLinodeInstanceRequest

`func NewPostMigrateLinodeInstanceRequest() *PostMigrateLinodeInstanceRequest`

NewPostMigrateLinodeInstanceRequest instantiates a new PostMigrateLinodeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostMigrateLinodeInstanceRequestWithDefaults

`func NewPostMigrateLinodeInstanceRequestWithDefaults() *PostMigrateLinodeInstanceRequest`

NewPostMigrateLinodeInstanceRequestWithDefaults instantiates a new PostMigrateLinodeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacementGroup

`func (o *PostMigrateLinodeInstanceRequest) GetPlacementGroup() PostMigrateLinodeInstanceRequestPlacementGroup`

GetPlacementGroup returns the PlacementGroup field if non-nil, zero value otherwise.

### GetPlacementGroupOk

`func (o *PostMigrateLinodeInstanceRequest) GetPlacementGroupOk() (*PostMigrateLinodeInstanceRequestPlacementGroup, bool)`

GetPlacementGroupOk returns a tuple with the PlacementGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroup

`func (o *PostMigrateLinodeInstanceRequest) SetPlacementGroup(v PostMigrateLinodeInstanceRequestPlacementGroup)`

SetPlacementGroup sets PlacementGroup field to given value.

### HasPlacementGroup

`func (o *PostMigrateLinodeInstanceRequest) HasPlacementGroup() bool`

HasPlacementGroup returns a boolean if a field has been set.

### GetRegion

`func (o *PostMigrateLinodeInstanceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostMigrateLinodeInstanceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostMigrateLinodeInstanceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostMigrateLinodeInstanceRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *PostMigrateLinodeInstanceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostMigrateLinodeInstanceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostMigrateLinodeInstanceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostMigrateLinodeInstanceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpgrade

`func (o *PostMigrateLinodeInstanceRequest) GetUpgrade() bool`

GetUpgrade returns the Upgrade field if non-nil, zero value otherwise.

### GetUpgradeOk

`func (o *PostMigrateLinodeInstanceRequest) GetUpgradeOk() (*bool, bool)`

GetUpgradeOk returns a tuple with the Upgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgrade

`func (o *PostMigrateLinodeInstanceRequest) SetUpgrade(v bool)`

SetUpgrade sets Upgrade field to given value.

### HasUpgrade

`func (o *PostMigrateLinodeInstanceRequest) HasUpgrade() bool`

HasUpgrade returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


