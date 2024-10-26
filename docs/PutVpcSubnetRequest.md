# PutVpcSubnetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The VPC Subnet&#39;s label, for display purposes only.  - Must be unique among the VPC&#39;s Subnets. - Can only contain ASCII letters, numbers, and hyphens (&#x60;-&#x60;). You can&#39;t use two consecutive hyphens (&#x60;--&#x60;). | [optional] 

## Methods

### NewPutVpcSubnetRequest

`func NewPutVpcSubnetRequest() *PutVpcSubnetRequest`

NewPutVpcSubnetRequest instantiates a new PutVpcSubnetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutVpcSubnetRequestWithDefaults

`func NewPutVpcSubnetRequestWithDefaults() *PutVpcSubnetRequest`

NewPutVpcSubnetRequestWithDefaults instantiates a new PutVpcSubnetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PutVpcSubnetRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutVpcSubnetRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutVpcSubnetRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutVpcSubnetRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


