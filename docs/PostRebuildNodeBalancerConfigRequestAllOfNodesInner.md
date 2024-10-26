# PostRebuildNodeBalancerConfigRequestAllOfNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The private IP Address where this backend can be reached. This _must_ be a private IP address. | [optional] 
**Id** | Pointer to **int32** | The unique ID of the Node to update. | [optional] 
**Label** | Pointer to **string** | The label for this node.  This is for display purposes only. | [optional] 
**Mode** | Pointer to **string** | The mode this NodeBalancer should use when sending traffic to this backend.  - If set to &#x60;accept&#x60; this backend is accepting traffic. - If set to &#x60;reject&#x60; this backend will not receive traffic. - If set to &#x60;drain&#x60; this backend will not receive _new_ traffic, but connections already pinned to it will continue to be routed to it. - If set to &#x60;backup&#x60;, this backend will only receive traffic if all &#x60;accept&#x60; nodes are down. | [optional] 
**Weight** | Pointer to **int32** | Used when picking a backend to serve a request and is not pinned to a single backend yet.  Nodes with a higher weight will receive more traffic. | [optional] 

## Methods

### NewPostRebuildNodeBalancerConfigRequestAllOfNodesInner

`func NewPostRebuildNodeBalancerConfigRequestAllOfNodesInner() *PostRebuildNodeBalancerConfigRequestAllOfNodesInner`

NewPostRebuildNodeBalancerConfigRequestAllOfNodesInner instantiates a new PostRebuildNodeBalancerConfigRequestAllOfNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostRebuildNodeBalancerConfigRequestAllOfNodesInnerWithDefaults

`func NewPostRebuildNodeBalancerConfigRequestAllOfNodesInnerWithDefaults() *PostRebuildNodeBalancerConfigRequestAllOfNodesInner`

NewPostRebuildNodeBalancerConfigRequestAllOfNodesInnerWithDefaults instantiates a new PostRebuildNodeBalancerConfigRequestAllOfNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetId

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMode

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetWeight

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PostRebuildNodeBalancerConfigRequestAllOfNodesInner) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


