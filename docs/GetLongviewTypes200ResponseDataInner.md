# GetLongviewTypes200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID representing the Longview type. | [optional] [readonly] 
**Label** | Pointer to **string** | The Longview type label is for display purposes only. | [optional] [readonly] 
**Price** | Pointer to [**GetLongviewTypes200ResponseDataInnerPrice**](GetLongviewTypes200ResponseDataInnerPrice.md) |  | [optional] 
**RegionPrices** | Pointer to [**[]GetLkeTypes200ResponseDataInnerRegionPricesInner**](GetLkeTypes200ResponseDataInnerRegionPricesInner.md) |  | [optional] 
**Transfer** | Pointer to **int32** | The monthly outbound transfer amount, in MB. | [optional] [readonly] 

## Methods

### NewGetLongviewTypes200ResponseDataInner

`func NewGetLongviewTypes200ResponseDataInner() *GetLongviewTypes200ResponseDataInner`

NewGetLongviewTypes200ResponseDataInner instantiates a new GetLongviewTypes200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLongviewTypes200ResponseDataInnerWithDefaults

`func NewGetLongviewTypes200ResponseDataInnerWithDefaults() *GetLongviewTypes200ResponseDataInner`

NewGetLongviewTypes200ResponseDataInnerWithDefaults instantiates a new GetLongviewTypes200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLongviewTypes200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLongviewTypes200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLongviewTypes200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetLongviewTypes200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetLongviewTypes200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetLongviewTypes200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetLongviewTypes200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetLongviewTypes200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPrice

`func (o *GetLongviewTypes200ResponseDataInner) GetPrice() GetLongviewTypes200ResponseDataInnerPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetLongviewTypes200ResponseDataInner) GetPriceOk() (*GetLongviewTypes200ResponseDataInnerPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetLongviewTypes200ResponseDataInner) SetPrice(v GetLongviewTypes200ResponseDataInnerPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetLongviewTypes200ResponseDataInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRegionPrices

`func (o *GetLongviewTypes200ResponseDataInner) GetRegionPrices() []GetLkeTypes200ResponseDataInnerRegionPricesInner`

GetRegionPrices returns the RegionPrices field if non-nil, zero value otherwise.

### GetRegionPricesOk

`func (o *GetLongviewTypes200ResponseDataInner) GetRegionPricesOk() (*[]GetLkeTypes200ResponseDataInnerRegionPricesInner, bool)`

GetRegionPricesOk returns a tuple with the RegionPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionPrices

`func (o *GetLongviewTypes200ResponseDataInner) SetRegionPrices(v []GetLkeTypes200ResponseDataInnerRegionPricesInner)`

SetRegionPrices sets RegionPrices field to given value.

### HasRegionPrices

`func (o *GetLongviewTypes200ResponseDataInner) HasRegionPrices() bool`

HasRegionPrices returns a boolean if a field has been set.

### GetTransfer

`func (o *GetLongviewTypes200ResponseDataInner) GetTransfer() int32`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *GetLongviewTypes200ResponseDataInner) GetTransferOk() (*int32, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *GetLongviewTypes200ResponseDataInner) SetTransfer(v int32)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *GetLongviewTypes200ResponseDataInner) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


