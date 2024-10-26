# GetObjectStorageTypes200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID representing the Object Storage type. | [optional] [readonly] 
**Label** | Pointer to **string** | The Object Storage type label is for display purposes only. | [optional] [readonly] 
**Price** | Pointer to [**GetObjectStorageTypes200ResponseDataInnerPrice**](GetObjectStorageTypes200ResponseDataInnerPrice.md) |  | [optional] 
**RegionPrices** | Pointer to [**[]GetLkeTypes200ResponseDataInnerRegionPricesInner**](GetLkeTypes200ResponseDataInnerRegionPricesInner.md) |  | [optional] 
**Transfer** | Pointer to **int32** | The monthly outbound transfer amount, in MB. | [optional] [readonly] 

## Methods

### NewGetObjectStorageTypes200ResponseDataInner

`func NewGetObjectStorageTypes200ResponseDataInner() *GetObjectStorageTypes200ResponseDataInner`

NewGetObjectStorageTypes200ResponseDataInner instantiates a new GetObjectStorageTypes200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageTypes200ResponseDataInnerWithDefaults

`func NewGetObjectStorageTypes200ResponseDataInnerWithDefaults() *GetObjectStorageTypes200ResponseDataInner`

NewGetObjectStorageTypes200ResponseDataInnerWithDefaults instantiates a new GetObjectStorageTypes200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetObjectStorageTypes200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetObjectStorageTypes200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetObjectStorageTypes200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetObjectStorageTypes200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetObjectStorageTypes200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetObjectStorageTypes200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetObjectStorageTypes200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetObjectStorageTypes200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPrice

`func (o *GetObjectStorageTypes200ResponseDataInner) GetPrice() GetObjectStorageTypes200ResponseDataInnerPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetObjectStorageTypes200ResponseDataInner) GetPriceOk() (*GetObjectStorageTypes200ResponseDataInnerPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetObjectStorageTypes200ResponseDataInner) SetPrice(v GetObjectStorageTypes200ResponseDataInnerPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetObjectStorageTypes200ResponseDataInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRegionPrices

`func (o *GetObjectStorageTypes200ResponseDataInner) GetRegionPrices() []GetLkeTypes200ResponseDataInnerRegionPricesInner`

GetRegionPrices returns the RegionPrices field if non-nil, zero value otherwise.

### GetRegionPricesOk

`func (o *GetObjectStorageTypes200ResponseDataInner) GetRegionPricesOk() (*[]GetLkeTypes200ResponseDataInnerRegionPricesInner, bool)`

GetRegionPricesOk returns a tuple with the RegionPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionPrices

`func (o *GetObjectStorageTypes200ResponseDataInner) SetRegionPrices(v []GetLkeTypes200ResponseDataInnerRegionPricesInner)`

SetRegionPrices sets RegionPrices field to given value.

### HasRegionPrices

`func (o *GetObjectStorageTypes200ResponseDataInner) HasRegionPrices() bool`

HasRegionPrices returns a boolean if a field has been set.

### GetTransfer

`func (o *GetObjectStorageTypes200ResponseDataInner) GetTransfer() int32`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *GetObjectStorageTypes200ResponseDataInner) GetTransferOk() (*int32, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *GetObjectStorageTypes200ResponseDataInner) SetTransfer(v int32)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *GetObjectStorageTypes200ResponseDataInner) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


