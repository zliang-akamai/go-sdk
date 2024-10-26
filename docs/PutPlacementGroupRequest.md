# PutPlacementGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | A unique name for the placement group. A placement group &#x60;label&#x60; has the following constraints:  - It needs to begin and end with an alphanumeric character. - It can only consist of alphanumeric characters, hyphens (&#x60;-&#x60;), underscores (&#x60;_&#x60;), or periods (&#x60;.&#x60;). | [optional] 

## Methods

### NewPutPlacementGroupRequest

`func NewPutPlacementGroupRequest() *PutPlacementGroupRequest`

NewPutPlacementGroupRequest instantiates a new PutPlacementGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutPlacementGroupRequestWithDefaults

`func NewPutPlacementGroupRequestWithDefaults() *PutPlacementGroupRequest`

NewPutPlacementGroupRequestWithDefaults instantiates a new PutPlacementGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PutPlacementGroupRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutPlacementGroupRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutPlacementGroupRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutPlacementGroupRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


