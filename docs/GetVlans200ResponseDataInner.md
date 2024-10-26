# GetVlans200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date this VLAN was created. | [optional] [readonly] 
**Label** | Pointer to **string** | The name of this VLAN. | [optional] [readonly] 
**Linodes** | Pointer to **[]int32** | An array of Linode IDs attached to this VLAN. | [optional] [readonly] 
**Region** | Pointer to **string** | This VLAN&#39;s data center region.  __Note__. Currently, a VLAN can only be assigned to a Linode within the same data center region. | [optional] [readonly] 

## Methods

### NewGetVlans200ResponseDataInner

`func NewGetVlans200ResponseDataInner() *GetVlans200ResponseDataInner`

NewGetVlans200ResponseDataInner instantiates a new GetVlans200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVlans200ResponseDataInnerWithDefaults

`func NewGetVlans200ResponseDataInnerWithDefaults() *GetVlans200ResponseDataInner`

NewGetVlans200ResponseDataInnerWithDefaults instantiates a new GetVlans200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetVlans200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetVlans200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetVlans200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetVlans200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLabel

`func (o *GetVlans200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetVlans200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetVlans200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetVlans200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLinodes

`func (o *GetVlans200ResponseDataInner) GetLinodes() []int32`

GetLinodes returns the Linodes field if non-nil, zero value otherwise.

### GetLinodesOk

`func (o *GetVlans200ResponseDataInner) GetLinodesOk() (*[]int32, bool)`

GetLinodesOk returns a tuple with the Linodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodes

`func (o *GetVlans200ResponseDataInner) SetLinodes(v []int32)`

SetLinodes sets Linodes field to given value.

### HasLinodes

`func (o *GetVlans200ResponseDataInner) HasLinodes() bool`

HasLinodes returns a boolean if a field has been set.

### GetRegion

`func (o *GetVlans200ResponseDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetVlans200ResponseDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetVlans200ResponseDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetVlans200ResponseDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


