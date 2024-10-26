# LinodePlacementGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The placement group&#39;s ID. You need to provide it for all operations that affect it. | [optional] 
**Label** | Pointer to **string** | The unique name set for the placement group. A label has these constraints:  - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (&#x60;-&#x60;), underscores (&#x60;_&#x60;) or periods (&#x60;.&#x60;). | [optional] 
**MigratingTo** | Pointer to **string** | The unique identifier for the [placement group](https://techdocs.akamai.com/cloud-computing/docs/work-with-placement-groups) to which this Linode is being migrated. Displayed as &#x60;null&#x60; if the Linode is not being migrated to a new placement group. | [optional] [readonly] 
**PlacementGroupPolicy** | Pointer to **string** | How requests to add future compute instances to your placement group are handled, and whether it remains compliant:  - &#x60;strict&#x60;. Don&#39;t assign a new compute instance if it breaks the grouped-together or spread-apart model set by the &#x60;placement_group_type&#x60;. Use this to ensure the placement group stays compliant (&#x60;is_compliant: true&#x60;). - &#x60;flexible&#x60;. Assign a new compute instance, even if it breaks the grouped-together or spread-apart model set by the &#x60;placement_group_type&#x60;. This makes the group non-compliant (&#x60;is_compliant: false&#x60;). You need to wait for Akamai to move the offending compute instance to make it compliant again, once the necessary capacity is available in the region. Offers flexibility to add future compute instances if compliance isn&#39;t an immediate concern.  &lt;&lt;LB&gt;&gt;  &gt; 📘 &gt; &gt; In rare cases, non-compliance can occur with a &#x60;strict&#x60; placement group if Akamai needs to failover or migrate your compute instances for maintenance. Fixing non-compliance for a &#x60;strict&#x60; placement group is prioritized over a &#x60;flexible&#x60; group. | [optional] 
**PlacementGroupType** | Pointer to **string** | How compute instances are distributed in your placement group. A &#x60;placement_group_type&#x60; using anti-affinity (&#x60;anti-affinity:local&#x60;) places compute instances in separate hosts, but still in the same region. This best supports the spread-apart model for high availability. A &#x60;placement_group_type&#x60; using affinity places compute instances physically close together, possibly on the same host. This supports the grouped-together model for low-latency.  &gt; 📘 &gt; &gt; Currently, only &#x60;anti_affinity:local&#x60; is available for &#x60;placement_group_type&#x60;. | [optional] [readonly] 

## Methods

### NewLinodePlacementGroup

`func NewLinodePlacementGroup() *LinodePlacementGroup`

NewLinodePlacementGroup instantiates a new LinodePlacementGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodePlacementGroupWithDefaults

`func NewLinodePlacementGroupWithDefaults() *LinodePlacementGroup`

NewLinodePlacementGroupWithDefaults instantiates a new LinodePlacementGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinodePlacementGroup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinodePlacementGroup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinodePlacementGroup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LinodePlacementGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *LinodePlacementGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LinodePlacementGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LinodePlacementGroup) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LinodePlacementGroup) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMigratingTo

`func (o *LinodePlacementGroup) GetMigratingTo() string`

GetMigratingTo returns the MigratingTo field if non-nil, zero value otherwise.

### GetMigratingToOk

`func (o *LinodePlacementGroup) GetMigratingToOk() (*string, bool)`

GetMigratingToOk returns a tuple with the MigratingTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigratingTo

`func (o *LinodePlacementGroup) SetMigratingTo(v string)`

SetMigratingTo sets MigratingTo field to given value.

### HasMigratingTo

`func (o *LinodePlacementGroup) HasMigratingTo() bool`

HasMigratingTo returns a boolean if a field has been set.

### GetPlacementGroupPolicy

`func (o *LinodePlacementGroup) GetPlacementGroupPolicy() string`

GetPlacementGroupPolicy returns the PlacementGroupPolicy field if non-nil, zero value otherwise.

### GetPlacementGroupPolicyOk

`func (o *LinodePlacementGroup) GetPlacementGroupPolicyOk() (*string, bool)`

GetPlacementGroupPolicyOk returns a tuple with the PlacementGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroupPolicy

`func (o *LinodePlacementGroup) SetPlacementGroupPolicy(v string)`

SetPlacementGroupPolicy sets PlacementGroupPolicy field to given value.

### HasPlacementGroupPolicy

`func (o *LinodePlacementGroup) HasPlacementGroupPolicy() bool`

HasPlacementGroupPolicy returns a boolean if a field has been set.

### GetPlacementGroupType

`func (o *LinodePlacementGroup) GetPlacementGroupType() string`

GetPlacementGroupType returns the PlacementGroupType field if non-nil, zero value otherwise.

### GetPlacementGroupTypeOk

`func (o *LinodePlacementGroup) GetPlacementGroupTypeOk() (*string, bool)`

GetPlacementGroupTypeOk returns a tuple with the PlacementGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroupType

`func (o *LinodePlacementGroup) SetPlacementGroupType(v string)`

SetPlacementGroupType sets PlacementGroupType field to given value.

### HasPlacementGroupType

`func (o *LinodePlacementGroup) HasPlacementGroupType() bool`

HasPlacementGroupType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


