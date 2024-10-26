# PutVpcRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A written description to help distinguish the VPC. | [optional] [default to ""]
**Label** | Pointer to **string** | The VPC&#39;s label, for display purposes only.  - Needs to be unique among the Account&#39;s VPCs. - Can only contain ASCII letters, numbers, and hyphens (&#x60;-&#x60;). You can&#39;t use two consecutive hyphens (&#x60;--&#x60;). | [optional] 

## Methods

### NewPutVpcRequest

`func NewPutVpcRequest() *PutVpcRequest`

NewPutVpcRequest instantiates a new PutVpcRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutVpcRequestWithDefaults

`func NewPutVpcRequestWithDefaults() *PutVpcRequest`

NewPutVpcRequestWithDefaults instantiates a new PutVpcRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PutVpcRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PutVpcRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PutVpcRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PutVpcRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *PutVpcRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutVpcRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutVpcRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutVpcRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


