# GetRegionsAvailability200ResponseAllOfDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Whether the compute instance type is available in the region. | [optional] 
**Plan** | Pointer to **string** | The compute instance [Type](https://techdocs.akamai.com/linode-api/reference/get-linode-types) ID. | [optional] 
**Region** | Pointer to **string** | The [Region](https://techdocs.akamai.com/linode-api/reference/get-regions) ID. | [optional] 

## Methods

### NewGetRegionsAvailability200ResponseAllOfDataInner

`func NewGetRegionsAvailability200ResponseAllOfDataInner() *GetRegionsAvailability200ResponseAllOfDataInner`

NewGetRegionsAvailability200ResponseAllOfDataInner instantiates a new GetRegionsAvailability200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegionsAvailability200ResponseAllOfDataInnerWithDefaults

`func NewGetRegionsAvailability200ResponseAllOfDataInnerWithDefaults() *GetRegionsAvailability200ResponseAllOfDataInner`

NewGetRegionsAvailability200ResponseAllOfDataInnerWithDefaults instantiates a new GetRegionsAvailability200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetPlan

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) GetPlan() string`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) GetPlanOk() (*string, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) SetPlan(v string)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRegion

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetRegionsAvailability200ResponseAllOfDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


