# GetLongviewPlan200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientsIncluded** | Pointer to **int32** | The number of Longview Clients that may be created with this Subscription tier. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of this Subscription tier. | [optional] [readonly] 
**Label** | Pointer to **string** | A display name for this Subscription tier. | [optional] [readonly] 
**Price** | Pointer to [**GetLongviewPlan200ResponsePrice**](GetLongviewPlan200ResponsePrice.md) |  | [optional] 

## Methods

### NewGetLongviewPlan200Response

`func NewGetLongviewPlan200Response() *GetLongviewPlan200Response`

NewGetLongviewPlan200Response instantiates a new GetLongviewPlan200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLongviewPlan200ResponseWithDefaults

`func NewGetLongviewPlan200ResponseWithDefaults() *GetLongviewPlan200Response`

NewGetLongviewPlan200ResponseWithDefaults instantiates a new GetLongviewPlan200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientsIncluded

`func (o *GetLongviewPlan200Response) GetClientsIncluded() int32`

GetClientsIncluded returns the ClientsIncluded field if non-nil, zero value otherwise.

### GetClientsIncludedOk

`func (o *GetLongviewPlan200Response) GetClientsIncludedOk() (*int32, bool)`

GetClientsIncludedOk returns a tuple with the ClientsIncluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsIncluded

`func (o *GetLongviewPlan200Response) SetClientsIncluded(v int32)`

SetClientsIncluded sets ClientsIncluded field to given value.

### HasClientsIncluded

`func (o *GetLongviewPlan200Response) HasClientsIncluded() bool`

HasClientsIncluded returns a boolean if a field has been set.

### GetId

`func (o *GetLongviewPlan200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLongviewPlan200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLongviewPlan200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetLongviewPlan200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetLongviewPlan200Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetLongviewPlan200Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetLongviewPlan200Response) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetLongviewPlan200Response) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPrice

`func (o *GetLongviewPlan200Response) GetPrice() GetLongviewPlan200ResponsePrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetLongviewPlan200Response) GetPriceOk() (*GetLongviewPlan200ResponsePrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetLongviewPlan200Response) SetPrice(v GetLongviewPlan200ResponsePrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetLongviewPlan200Response) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


