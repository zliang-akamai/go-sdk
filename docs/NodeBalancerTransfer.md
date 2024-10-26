# NodeBalancerTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **NullableFloat32** | The total outbound transfer, in MB, used for this NodeBalancer this month. | [optional] [readonly] 
**Out** | Pointer to **NullableFloat32** | The total inbound transfer, in MB, used for this NodeBalancer this month. | [optional] [readonly] 
**Total** | Pointer to **NullableFloat32** | The total transfer, in MB, used by this NodeBalancer this month. | [optional] [readonly] 

## Methods

### NewNodeBalancerTransfer

`func NewNodeBalancerTransfer() *NodeBalancerTransfer`

NewNodeBalancerTransfer instantiates a new NodeBalancerTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeBalancerTransferWithDefaults

`func NewNodeBalancerTransferWithDefaults() *NodeBalancerTransfer`

NewNodeBalancerTransferWithDefaults instantiates a new NodeBalancerTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *NodeBalancerTransfer) GetIn() float32`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *NodeBalancerTransfer) GetInOk() (*float32, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *NodeBalancerTransfer) SetIn(v float32)`

SetIn sets In field to given value.

### HasIn

`func (o *NodeBalancerTransfer) HasIn() bool`

HasIn returns a boolean if a field has been set.

### SetInNil

`func (o *NodeBalancerTransfer) SetInNil(b bool)`

 SetInNil sets the value for In to be an explicit nil

### UnsetIn
`func (o *NodeBalancerTransfer) UnsetIn()`

UnsetIn ensures that no value is present for In, not even an explicit nil
### GetOut

`func (o *NodeBalancerTransfer) GetOut() float32`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *NodeBalancerTransfer) GetOutOk() (*float32, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *NodeBalancerTransfer) SetOut(v float32)`

SetOut sets Out field to given value.

### HasOut

`func (o *NodeBalancerTransfer) HasOut() bool`

HasOut returns a boolean if a field has been set.

### SetOutNil

`func (o *NodeBalancerTransfer) SetOutNil(b bool)`

 SetOutNil sets the value for Out to be an explicit nil

### UnsetOut
`func (o *NodeBalancerTransfer) UnsetOut()`

UnsetOut ensures that no value is present for Out, not even an explicit nil
### GetTotal

`func (o *NodeBalancerTransfer) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *NodeBalancerTransfer) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *NodeBalancerTransfer) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *NodeBalancerTransfer) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### SetTotalNil

`func (o *NodeBalancerTransfer) SetTotalNil(b bool)`

 SetTotalNil sets the value for Total to be an explicit nil

### UnsetTotal
`func (o *NodeBalancerTransfer) UnsetTotal()`

UnsetTotal ensures that no value is present for Total, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


