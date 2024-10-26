# GetNetworkTransferPrices200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID representing the network transfer price. | [optional] [readonly] 
**Label** | Pointer to **string** | The network transfer price label is for display purposes only. | [optional] [readonly] 
**Price** | Pointer to [**GetNetworkTransferPrices200ResponseDataInnerPrice**](GetNetworkTransferPrices200ResponseDataInnerPrice.md) |  | [optional] 
**RegionPrices** | Pointer to [**[]GetNetworkTransferPrices200ResponseDataInnerRegionPricesInner**](GetNetworkTransferPrices200ResponseDataInnerRegionPricesInner.md) |  | [optional] 
**Transfer** | Pointer to **int32** | The monthly outbound transfer amount, in MB. | [optional] [readonly] 

## Methods

### NewGetNetworkTransferPrices200ResponseDataInner

`func NewGetNetworkTransferPrices200ResponseDataInner() *GetNetworkTransferPrices200ResponseDataInner`

NewGetNetworkTransferPrices200ResponseDataInner instantiates a new GetNetworkTransferPrices200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkTransferPrices200ResponseDataInnerWithDefaults

`func NewGetNetworkTransferPrices200ResponseDataInnerWithDefaults() *GetNetworkTransferPrices200ResponseDataInner`

NewGetNetworkTransferPrices200ResponseDataInnerWithDefaults instantiates a new GetNetworkTransferPrices200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkTransferPrices200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkTransferPrices200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetNetworkTransferPrices200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetNetworkTransferPrices200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPrice

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetPrice() GetNetworkTransferPrices200ResponseDataInnerPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetPriceOk() (*GetNetworkTransferPrices200ResponseDataInnerPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetNetworkTransferPrices200ResponseDataInner) SetPrice(v GetNetworkTransferPrices200ResponseDataInnerPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetNetworkTransferPrices200ResponseDataInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRegionPrices

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetRegionPrices() []GetNetworkTransferPrices200ResponseDataInnerRegionPricesInner`

GetRegionPrices returns the RegionPrices field if non-nil, zero value otherwise.

### GetRegionPricesOk

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetRegionPricesOk() (*[]GetNetworkTransferPrices200ResponseDataInnerRegionPricesInner, bool)`

GetRegionPricesOk returns a tuple with the RegionPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionPrices

`func (o *GetNetworkTransferPrices200ResponseDataInner) SetRegionPrices(v []GetNetworkTransferPrices200ResponseDataInnerRegionPricesInner)`

SetRegionPrices sets RegionPrices field to given value.

### HasRegionPrices

`func (o *GetNetworkTransferPrices200ResponseDataInner) HasRegionPrices() bool`

HasRegionPrices returns a boolean if a field has been set.

### GetTransfer

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetTransfer() int32`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *GetNetworkTransferPrices200ResponseDataInner) GetTransferOk() (*int32, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *GetNetworkTransferPrices200ResponseDataInner) SetTransfer(v int32)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *GetNetworkTransferPrices200ResponseDataInner) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


