# PostCloneLinodeInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupsEnabled** | Pointer to **bool** | If this field is set to &#x60;true&#x60;, the created Linode will automatically be enrolled in the Linode Backup service. This will incur an additional charge. Pricing is included in the response from [List types](https://techdocs.akamai.com/linode-api/reference/get-linode-types).  - Can only be included when cloning to a new Linode. | [optional] 
**Configs** | Pointer to **[]int32** | An array of configuration profile IDs.  - If the &#x60;configs&#x60; parameter __is not provided__, then __all configuration profiles and their associated disks will be cloned__ from the source Linode. Any disks specified by the &#x60;disks&#x60; parameter will also be cloned. - __If an empty array is provided__ for the &#x60;configs&#x60; parameter, then __no configuration profiles (nor their associated disks) will be cloned__ from the source Linode. Any disks specified by the &#x60;disks&#x60; parameter will still be cloned. - __If a non-empty array is provided__ for the &#x60;configs&#x60; parameter, then __the configuration profiles specified in the array (and their associated disks) will be cloned__ from the source Linode. Any disks specified by the &#x60;disks&#x60; parameter will also be cloned. | [optional] 
**Disks** | Pointer to **[]int32** | An array of disk IDs.  - If the &#x60;disks&#x60; parameter __is not provided__, then __no extra disks will be cloned__ from the source Linode. All disks associated with the configuration profiles specified by the &#x60;configs&#x60; parameter will still be cloned. - __If an empty array is provided__ for the &#x60;disks&#x60; parameter, then __no extra disks will be cloned__ from the source Linode. All disks associated with the configuration profiles specified by the &#x60;configs&#x60; parameter will still be cloned. - __If a non-empty array is provided__ for the &#x60;disks&#x60; parameter, then __the disks specified in the array will be cloned__ from the source Linode, in addition to any disks associated with the configuration profiles specified by the &#x60;configs&#x60; parameter. | [optional] 
**Group** | Pointer to **string** | A label used to group Linodes for display. Linodes are not required to have a group. | [optional] 
**Label** | Pointer to **string** | The label to assign this Linode when cloning to a new Linode.  - Can only be provided when cloning to a new Linode. - Defaults to &#x60;linode&#x60;. | [optional] 
**LinodeId** | Pointer to **int32** | If an existing Linode is the target for the clone, the ID of that Linode. The existing Linode must have enough resources to accept the clone. | [optional] 
**Metadata** | Pointer to [**PostLinodeInstanceRequestAllOfMetadata**](PostLinodeInstanceRequestAllOfMetadata.md) |  | [optional] 
**PlacementGroup** | Pointer to [**PostCloneLinodeInstanceRequestPlacementGroup**](PostCloneLinodeInstanceRequestPlacementGroup.md) |  | [optional] 
**PrivateIp** | Pointer to **bool** | If true, the created Linode will have private networking enabled and assigned a private IPv4 address.  - Can only be provided when cloning to a new Linode. | [optional] 
**Region** | Pointer to **string** | This is the Region where the Linode will be deployed. To view all available Regions you can deploy to, run [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions).  - Region can only be provided and is required when cloning to a new Linode. | [optional] 
**Type** | Pointer to **string** | A Linode&#39;s Type determines what resources are available to it, including disk space, memory, and virtual cpus. The amounts available to a specific Linode are returned as &#x60;specs&#x60; on the Linode object.  To view all available Linode Types you can deploy with, run [List types](https://techdocs.akamai.com/linode-api/reference/get-linode-types).  - Type can only be provided and is required when cloning to a new Linode. | [optional] 

## Methods

### NewPostCloneLinodeInstanceRequest

`func NewPostCloneLinodeInstanceRequest() *PostCloneLinodeInstanceRequest`

NewPostCloneLinodeInstanceRequest instantiates a new PostCloneLinodeInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCloneLinodeInstanceRequestWithDefaults

`func NewPostCloneLinodeInstanceRequestWithDefaults() *PostCloneLinodeInstanceRequest`

NewPostCloneLinodeInstanceRequestWithDefaults instantiates a new PostCloneLinodeInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupsEnabled

`func (o *PostCloneLinodeInstanceRequest) GetBackupsEnabled() bool`

GetBackupsEnabled returns the BackupsEnabled field if non-nil, zero value otherwise.

### GetBackupsEnabledOk

`func (o *PostCloneLinodeInstanceRequest) GetBackupsEnabledOk() (*bool, bool)`

GetBackupsEnabledOk returns a tuple with the BackupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsEnabled

`func (o *PostCloneLinodeInstanceRequest) SetBackupsEnabled(v bool)`

SetBackupsEnabled sets BackupsEnabled field to given value.

### HasBackupsEnabled

`func (o *PostCloneLinodeInstanceRequest) HasBackupsEnabled() bool`

HasBackupsEnabled returns a boolean if a field has been set.

### GetConfigs

`func (o *PostCloneLinodeInstanceRequest) GetConfigs() []int32`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *PostCloneLinodeInstanceRequest) GetConfigsOk() (*[]int32, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *PostCloneLinodeInstanceRequest) SetConfigs(v []int32)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *PostCloneLinodeInstanceRequest) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetDisks

`func (o *PostCloneLinodeInstanceRequest) GetDisks() []int32`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *PostCloneLinodeInstanceRequest) GetDisksOk() (*[]int32, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *PostCloneLinodeInstanceRequest) SetDisks(v []int32)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *PostCloneLinodeInstanceRequest) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetGroup

`func (o *PostCloneLinodeInstanceRequest) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *PostCloneLinodeInstanceRequest) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *PostCloneLinodeInstanceRequest) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *PostCloneLinodeInstanceRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetLabel

`func (o *PostCloneLinodeInstanceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostCloneLinodeInstanceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostCloneLinodeInstanceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostCloneLinodeInstanceRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLinodeId

`func (o *PostCloneLinodeInstanceRequest) GetLinodeId() int32`

GetLinodeId returns the LinodeId field if non-nil, zero value otherwise.

### GetLinodeIdOk

`func (o *PostCloneLinodeInstanceRequest) GetLinodeIdOk() (*int32, bool)`

GetLinodeIdOk returns a tuple with the LinodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodeId

`func (o *PostCloneLinodeInstanceRequest) SetLinodeId(v int32)`

SetLinodeId sets LinodeId field to given value.

### HasLinodeId

`func (o *PostCloneLinodeInstanceRequest) HasLinodeId() bool`

HasLinodeId returns a boolean if a field has been set.

### GetMetadata

`func (o *PostCloneLinodeInstanceRequest) GetMetadata() PostLinodeInstanceRequestAllOfMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PostCloneLinodeInstanceRequest) GetMetadataOk() (*PostLinodeInstanceRequestAllOfMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PostCloneLinodeInstanceRequest) SetMetadata(v PostLinodeInstanceRequestAllOfMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PostCloneLinodeInstanceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPlacementGroup

`func (o *PostCloneLinodeInstanceRequest) GetPlacementGroup() PostCloneLinodeInstanceRequestPlacementGroup`

GetPlacementGroup returns the PlacementGroup field if non-nil, zero value otherwise.

### GetPlacementGroupOk

`func (o *PostCloneLinodeInstanceRequest) GetPlacementGroupOk() (*PostCloneLinodeInstanceRequestPlacementGroup, bool)`

GetPlacementGroupOk returns a tuple with the PlacementGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroup

`func (o *PostCloneLinodeInstanceRequest) SetPlacementGroup(v PostCloneLinodeInstanceRequestPlacementGroup)`

SetPlacementGroup sets PlacementGroup field to given value.

### HasPlacementGroup

`func (o *PostCloneLinodeInstanceRequest) HasPlacementGroup() bool`

HasPlacementGroup returns a boolean if a field has been set.

### GetPrivateIp

`func (o *PostCloneLinodeInstanceRequest) GetPrivateIp() bool`

GetPrivateIp returns the PrivateIp field if non-nil, zero value otherwise.

### GetPrivateIpOk

`func (o *PostCloneLinodeInstanceRequest) GetPrivateIpOk() (*bool, bool)`

GetPrivateIpOk returns a tuple with the PrivateIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateIp

`func (o *PostCloneLinodeInstanceRequest) SetPrivateIp(v bool)`

SetPrivateIp sets PrivateIp field to given value.

### HasPrivateIp

`func (o *PostCloneLinodeInstanceRequest) HasPrivateIp() bool`

HasPrivateIp returns a boolean if a field has been set.

### GetRegion

`func (o *PostCloneLinodeInstanceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostCloneLinodeInstanceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostCloneLinodeInstanceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostCloneLinodeInstanceRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetType

`func (o *PostCloneLinodeInstanceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostCloneLinodeInstanceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostCloneLinodeInstanceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostCloneLinodeInstanceRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


