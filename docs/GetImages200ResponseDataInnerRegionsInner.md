# GetImages200ResponseDataInnerRegionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | The unique identifier for the core compute region where this image is stored. | [optional] 
**Status** | Pointer to **string** | The status of the image in this &#x60;region&#x60;. Possible values are &#x60;available&#x60;, &#x60;creating&#x60;, &#x60;pending&#x60;, &#x60;pending deletion&#x60;, &#x60;pending replication&#x60;, or &#x60;replicating&#x60;. | [optional] 

## Methods

### NewGetImages200ResponseDataInnerRegionsInner

`func NewGetImages200ResponseDataInnerRegionsInner() *GetImages200ResponseDataInnerRegionsInner`

NewGetImages200ResponseDataInnerRegionsInner instantiates a new GetImages200ResponseDataInnerRegionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImages200ResponseDataInnerRegionsInnerWithDefaults

`func NewGetImages200ResponseDataInnerRegionsInnerWithDefaults() *GetImages200ResponseDataInnerRegionsInner`

NewGetImages200ResponseDataInnerRegionsInnerWithDefaults instantiates a new GetImages200ResponseDataInnerRegionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *GetImages200ResponseDataInnerRegionsInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetImages200ResponseDataInnerRegionsInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetImages200ResponseDataInnerRegionsInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetImages200ResponseDataInnerRegionsInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *GetImages200ResponseDataInnerRegionsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetImages200ResponseDataInnerRegionsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetImages200ResponseDataInnerRegionsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetImages200ResponseDataInnerRegionsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


