# PostNodeBalancerNodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The private IP Address where this backend can be reached. This _must_ be a private IP address. | 
**ConfigId** | Pointer to **int32** | The NodeBalancer Config ID that this Node belongs to. | [optional] [readonly] 
**Id** | Pointer to **int32** | This node&#39;s unique ID. | [optional] [readonly] 
**Label** | **string** | The label for this node.  This is for display purposes only. | 
**Mode** | Pointer to **string** | The mode this NodeBalancer should use when sending traffic to this backend.  - If set to &#x60;accept&#x60; this backend is accepting traffic. - If set to &#x60;reject&#x60; this backend will not receive traffic. - If set to &#x60;drain&#x60; this backend will not receive _new_ traffic, but connections already pinned to it will continue to be routed to it. - If set to &#x60;backup&#x60;, this backend will only receive traffic if all &#x60;accept&#x60; nodes are down. | [optional] 
**NodebalancerId** | Pointer to **int32** | The NodeBalancer ID that this Node belongs to. | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of this node, based on the configured checks of its NodeBalancer Config. | [optional] [readonly] 
**Weight** | Pointer to **int32** | Used when picking a backend to serve a request and is not pinned to a single backend yet.  Nodes with a higher weight will receive more traffic. | [optional] 

## Methods

### NewPostNodeBalancerNodeRequest

`func NewPostNodeBalancerNodeRequest(address string, label string, ) *PostNodeBalancerNodeRequest`

NewPostNodeBalancerNodeRequest instantiates a new PostNodeBalancerNodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostNodeBalancerNodeRequestWithDefaults

`func NewPostNodeBalancerNodeRequestWithDefaults() *PostNodeBalancerNodeRequest`

NewPostNodeBalancerNodeRequestWithDefaults instantiates a new PostNodeBalancerNodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PostNodeBalancerNodeRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PostNodeBalancerNodeRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PostNodeBalancerNodeRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetConfigId

`func (o *PostNodeBalancerNodeRequest) GetConfigId() int32`

GetConfigId returns the ConfigId field if non-nil, zero value otherwise.

### GetConfigIdOk

`func (o *PostNodeBalancerNodeRequest) GetConfigIdOk() (*int32, bool)`

GetConfigIdOk returns a tuple with the ConfigId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigId

`func (o *PostNodeBalancerNodeRequest) SetConfigId(v int32)`

SetConfigId sets ConfigId field to given value.

### HasConfigId

`func (o *PostNodeBalancerNodeRequest) HasConfigId() bool`

HasConfigId returns a boolean if a field has been set.

### GetId

`func (o *PostNodeBalancerNodeRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostNodeBalancerNodeRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostNodeBalancerNodeRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostNodeBalancerNodeRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PostNodeBalancerNodeRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostNodeBalancerNodeRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostNodeBalancerNodeRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMode

`func (o *PostNodeBalancerNodeRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *PostNodeBalancerNodeRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *PostNodeBalancerNodeRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *PostNodeBalancerNodeRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetNodebalancerId

`func (o *PostNodeBalancerNodeRequest) GetNodebalancerId() int32`

GetNodebalancerId returns the NodebalancerId field if non-nil, zero value otherwise.

### GetNodebalancerIdOk

`func (o *PostNodeBalancerNodeRequest) GetNodebalancerIdOk() (*int32, bool)`

GetNodebalancerIdOk returns a tuple with the NodebalancerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancerId

`func (o *PostNodeBalancerNodeRequest) SetNodebalancerId(v int32)`

SetNodebalancerId sets NodebalancerId field to given value.

### HasNodebalancerId

`func (o *PostNodeBalancerNodeRequest) HasNodebalancerId() bool`

HasNodebalancerId returns a boolean if a field has been set.

### GetStatus

`func (o *PostNodeBalancerNodeRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostNodeBalancerNodeRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostNodeBalancerNodeRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostNodeBalancerNodeRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWeight

`func (o *PostNodeBalancerNodeRequest) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *PostNodeBalancerNodeRequest) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *PostNodeBalancerNodeRequest) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *PostNodeBalancerNodeRequest) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


