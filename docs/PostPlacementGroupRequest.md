# PostPlacementGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A unique name for the placement group. A placement group &#x60;label&#x60; has the following constraints:  - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (&#x60;-&#x60;), underscores (&#x60;_&#x60;), or periods (&#x60;.&#x60;). | 
**PlacementGroupPolicy** | **string** | How requests to add future compute instances to your placement group are handled:  - &#x60;strict&#x60;. Don&#39;t add a compute instance if it breaks the grouped-together or spread-apart model set by your &#x60;placement_group_type&#x60;. For example, with &#x60;anti-affinity:local&#x60; as your &#x60;placement_group_type&#x60;, assume it already has three qualifying compute instances on separate hosts, to support the spread-apart model. If a fourth compute instance is assigned that&#39;s on the same host as one of the existing three, an error is thrown and the system won&#39;t add it. Ensures the placement group stays compliant with your selected model. - &#x60;flexible&#x60;. Add a new compute instance, even if it breaks the grouped-together or spread-apart model set by your &#x60;placement_group_type&#x60;. Breaking the model makes the placement group non-compliant. You need to wait for Akamai to move the offending compute instances to make the group compliant again, once the necessary capacity is available in the region. Offers flexibility to add future compute instances if compliance isn&#39;t an immediate concern.  &gt; 📘 &gt; &gt; Once you create a placement group, you can&#39;t change its &#x60;placement_group_policy&#x60; setting. | 
**PlacementGroupType** | **string** | How compute instances are distributed in your placement group. A &#x60;placement_group_type&#x60; using anti-affinity (&#x60;anti-affinity:local&#x60;) places compute instances in separate hosts, but still in the same region. This best supports the spread-apart model for high availability. A &#x60;placement_group_type&#x60; using affinity places compute instances physically close together, possibly on the same host. This supports the grouped-together model for low-latency.  &gt; 📘 &gt; &gt; Currently, only &#x60;anti_affinity:local&#x60; is available for &#x60;placement_group_type&#x60;. | 
**Region** | **string** | The data center that houses the compute instances you want to add to your placement group. Run the [List Linodes](https://techdocs.akamai.com/linode-api/reference/get-linode-instances) operation to view your compute instances and store the &#x60;region&#x60; identifier. | 

## Methods

### NewPostPlacementGroupRequest

`func NewPostPlacementGroupRequest(label string, placementGroupPolicy string, placementGroupType string, region string, ) *PostPlacementGroupRequest`

NewPostPlacementGroupRequest instantiates a new PostPlacementGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPlacementGroupRequestWithDefaults

`func NewPostPlacementGroupRequestWithDefaults() *PostPlacementGroupRequest`

NewPostPlacementGroupRequestWithDefaults instantiates a new PostPlacementGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PostPlacementGroupRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostPlacementGroupRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostPlacementGroupRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPlacementGroupPolicy

`func (o *PostPlacementGroupRequest) GetPlacementGroupPolicy() string`

GetPlacementGroupPolicy returns the PlacementGroupPolicy field if non-nil, zero value otherwise.

### GetPlacementGroupPolicyOk

`func (o *PostPlacementGroupRequest) GetPlacementGroupPolicyOk() (*string, bool)`

GetPlacementGroupPolicyOk returns a tuple with the PlacementGroupPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroupPolicy

`func (o *PostPlacementGroupRequest) SetPlacementGroupPolicy(v string)`

SetPlacementGroupPolicy sets PlacementGroupPolicy field to given value.


### GetPlacementGroupType

`func (o *PostPlacementGroupRequest) GetPlacementGroupType() string`

GetPlacementGroupType returns the PlacementGroupType field if non-nil, zero value otherwise.

### GetPlacementGroupTypeOk

`func (o *PostPlacementGroupRequest) GetPlacementGroupTypeOk() (*string, bool)`

GetPlacementGroupTypeOk returns a tuple with the PlacementGroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroupType

`func (o *PostPlacementGroupRequest) SetPlacementGroupType(v string)`

SetPlacementGroupType sets PlacementGroupType field to given value.


### GetRegion

`func (o *PostPlacementGroupRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostPlacementGroupRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostPlacementGroupRequest) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


