# GetPlacementGroups200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The placement group&#39;s ID. You need to provide it for all operations that affect it. | [optional] 
**IsCompliant** | Pointer to **bool** | Whether all of the compute instances in your placement group are compliant. If &#x60;true&#x60;, all compute instances meet either the grouped-together or spread-apart model, which you determine through your selected &#x60;placement_group_type&#x60;. If &#x60;false&#x60;, a compute instance is out of this compliance. For example, assume you&#39;ve set &#x60;anti-affinity:local&#x60; as your &#x60;placement_group_type&#x60; and your group already has three qualifying compute instances on separate hosts, to support the spread-apart model. If a fourth compute instance is assigned that&#39;s on the same host as one of the existing three, the placement group is non-compliant. Enforce compliance in your group by setting a &#x60;placement_group_policy&#x60;.  &gt; 📘 &gt; &gt; Fixing compliance is not self-service. You need to wait for our assistance to physically move compute instances to make the group compliant again. | [optional] 
**Label** | Pointer to **string** | The unique name set for the placement group. A label has these constraints:  - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (&#x60;-&#x60;), underscores (&#x60;_&#x60;) or periods (&#x60;.&#x60;). | [optional] 
**Members** | Pointer to [**[]GetPlacementGroups200ResponseDataInnerMembersInner**](GetPlacementGroups200ResponseDataInnerMembersInner.md) | An array of compute instances included in the placement group. | [optional] 
**Migrations** | Pointer to [**GetPlacementGroups200ResponseDataInnerMigrations**](GetPlacementGroups200ResponseDataInnerMigrations.md) |  | [optional] 
**PlacementGroupPolicy** | Pointer to **string** | How requests to add future compute instances to your placement group are handled, and whether it remains compliant:  - &#x60;strict&#x60;. Don&#39;t assign a new compute instance if it breaks the grouped-together or spread-apart model set by the &#x60;placement_group_type&#x60;. Use this to ensure the placement group stays compliant (&#x60;is_compliant: true&#x60;). - &#x60;flexible&#x60;. Assign a new compute instance, even if it breaks the grouped-together or spread-apart model set by the &#x60;placement_group_type&#x60;. This makes the group non-compliant (&#x60;is_compliant: false&#x60;). You need to wait for Akamai to move the offending compute instance to make it compliant again, once the necessary capacity is available in the region. Offers flexibility to add future compute instances if compliance isn&#39;t an immediate concern.  &lt;&lt;LB&gt;&gt;  &gt; 📘 &gt; &gt; In rare cases, non-compliance can occur with a &#x60;strict&#x60; placement group if Akamai needs to failover or migrate your compute instances for maintenance. Fixing non-compliance for a &#x60;strict&#x60; placement group is prioritized over a &#x60;flexible&#x60; group. | [optional] 
**PlacementGroupType** | Pointer to **string** | How compute instances are distributed in your placement group. A &#x60;placement_group_type&#x60; using anti-affinity (&#x60;anti-affinity:local&#x60;) places compute instances in separate hosts, but still in the same region. This best supports the spread-apart model for high availability. A &#x60;placement_group_type&#x60; using affinity places compute instances physically close together, possibly on the same host. This supports the grouped-together model for low-latency.  &gt; 📘 &gt; &gt; Currently, only &#x60;anti_affinity:local&#x60; is available for &#x60;placement_group_type&#x60;. | [optional] [readonly] 
**Region** | Pointer to **string** | The [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where the placement group was deployed. | [optional] [readonly] 

## Methods

### NewGetPlacementGroups200ResponseDataInner

`func NewGetPlacementGroups200ResponseDataInner() *GetPlacementGroups200ResponseDataInner`

NewGetPlacementGroups200ResponseDataInner instantiates a new GetPlacementGroups200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlacementGroups200ResponseDataInnerWithDefaults

`func NewGetPlacementGroups200ResponseDataInnerWithDefaults() *GetPlacementGroups200ResponseDataInner`

NewGetPlacementGroups200ResponseDataInnerWithDefaults instantiates a new GetPlacementGroups200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetPlacementGroups200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPlacementGroups200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPlacementGroups200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetPlacementGroups200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsCompliant

`func (o *GetPlacementGroups200ResponseDataInner) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *GetPlacementGroups200ResponseDataInner) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *GetPlacementGroups200ResponseDataInner) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.

### HasIsCompliant

`func (o *GetPlacementGroups200ResponseDataInner) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### GetLabel

`func (o *GetPlacementGroups200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetPlacementGroups200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetPlacementGroups200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetPlacementGroups200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMembers

`func (o *GetPlacementGroups200ResponseDataInner) GetMembers() []GetPlacementGroups200ResponseDataInnerMembersInner`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetPlacementGroups200ResponseDataInner) GetMembersOk() (*[]GetPlacementGroups200ResponseDataInnerMembersInner, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetPlacementGroups200ResponseDataInner) SetMembers(v []GetPlacementGroups200ResponseDataInnerMembersInner)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *GetPlacementGroups200ResponseDataInner) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetMigrations

`func (o *GetPlacementGroups200ResponseDataInner) GetMigrations() GetPlacementGroups200ResponseDataInnerMigrations`

GetMigrations returns the Migrations field if non-nil, zero value otherwise.

### GetMigrationsOk

`func (o *GetPlacementGroups200ResponseDataInner) GetMigrationsOk() (*GetPlacementGroups200ResponseDataInnerMigrations, bool)`

GetMigrationsOk returns a tuple with the Migrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrations

`func (o *GetPlacementGroups200ResponseDataInner) SetMigrations(v GetPlacementGroups200ResponseDataInnerMigrations)`

SetMigrations sets Migrations field to given value.

### HasMigrations

`func (o *GetPlacementGroups200ResponseDataInner) HasMigrations() bool`

HasMigrations returns a boolean if a field has been set.

### GetPlacementGroupPolicy

`func (o *GetPlacementGroups200ResponseDataInner) GetPlacementGroupPolicy() string`

GetPlacementGroupPolicy returns the PlacementGroupPolicy field if non-nil, zero value otherwise.

### GetPlacementGroupPolicyOk

`func (o *GetPlacementGroups200ResponseDataInner) GetPlacementGroupPolicyOk() (*string, bool)`

GetPlacementGroupPolicyOk returns a tuple with the PlacementGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroupPolicy

`func (o *GetPlacementGroups200ResponseDataInner) SetPlacementGroupPolicy(v string)`

SetPlacementGroupPolicy sets PlacementGroupPolicy field to given value.

### HasPlacementGroupPolicy

`func (o *GetPlacementGroups200ResponseDataInner) HasPlacementGroupPolicy() bool`

HasPlacementGroupPolicy returns a boolean if a field has been set.

### GetPlacementGroupType

`func (o *GetPlacementGroups200ResponseDataInner) GetPlacementGroupType() string`

GetPlacementGroupType returns the PlacementGroupType field if non-nil, zero value otherwise.

### GetPlacementGroupTypeOk

`func (o *GetPlacementGroups200ResponseDataInner) GetPlacementGroupTypeOk() (*string, bool)`

GetPlacementGroupTypeOk returns a tuple with the PlacementGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroupType

`func (o *GetPlacementGroups200ResponseDataInner) SetPlacementGroupType(v string)`

SetPlacementGroupType sets PlacementGroupType field to given value.

### HasPlacementGroupType

`func (o *GetPlacementGroups200ResponseDataInner) HasPlacementGroupType() bool`

HasPlacementGroupType returns a boolean if a field has been set.

### GetRegion

`func (o *GetPlacementGroups200ResponseDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetPlacementGroups200ResponseDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetPlacementGroups200ResponseDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetPlacementGroups200ResponseDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

