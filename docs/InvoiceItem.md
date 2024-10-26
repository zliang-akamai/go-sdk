# InvoiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **float32** | The price, in US dollars, of the Invoice Item. Equal to the unit price multiplied by quantity. | [optional] [readonly] 
**From** | Pointer to **time.Time** | The date the Invoice Item started, based on month. | [optional] [readonly] 
**Label** | Pointer to **string** | The Invoice Item&#39;s display label. | [optional] [readonly] 
**Quantity** | Pointer to **int32** | The quantity of this Item for the specified Invoice. | [optional] [readonly] 
**Region** | Pointer to **NullableString** | The ID of the applicable Region associated with this Invoice Item.  &#x60;null&#x60; if there is no applicable Region. | [optional] [readonly] 
**Tax** | Pointer to **float32** | The amount of tax levied on this Item in US Dollars. | [optional] [readonly] 
**To** | Pointer to **time.Time** | The date the Invoice Item ended, based on month. | [optional] [readonly] 
**Total** | Pointer to **float32** | The price of this Item after taxes in US Dollars. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of service, ether &#x60;hourly&#x60; or &#x60;misc&#x60;. | [optional] [readonly] 
**UnitPrice** | Pointer to **string** | The monthly service fee in US Dollars for this Item. | [optional] [readonly] 

## Methods

### NewInvoiceItem

`func NewInvoiceItem() *InvoiceItem`

NewInvoiceItem instantiates a new InvoiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceItemWithDefaults

`func NewInvoiceItemWithDefaults() *InvoiceItem`

NewInvoiceItemWithDefaults instantiates a new InvoiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InvoiceItem) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceItem) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceItem) SetAmount(v float32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoiceItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetFrom

`func (o *InvoiceItem) GetFrom() time.Time`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *InvoiceItem) GetFromOk() (*time.Time, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *InvoiceItem) SetFrom(v time.Time)`

SetFrom sets From field to given value.

### HasFrom

`func (o *InvoiceItem) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLabel

`func (o *InvoiceItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *InvoiceItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *InvoiceItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *InvoiceItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetQuantity

`func (o *InvoiceItem) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *InvoiceItem) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *InvoiceItem) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *InvoiceItem) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRegion

`func (o *InvoiceItem) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *InvoiceItem) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *InvoiceItem) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *InvoiceItem) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *InvoiceItem) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *InvoiceItem) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetTax

`func (o *InvoiceItem) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *InvoiceItem) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *InvoiceItem) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *InvoiceItem) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTo

`func (o *InvoiceItem) GetTo() time.Time`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InvoiceItem) GetToOk() (*time.Time, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InvoiceItem) SetTo(v time.Time)`

SetTo sets To field to given value.

### HasTo

`func (o *InvoiceItem) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTotal

`func (o *InvoiceItem) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InvoiceItem) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InvoiceItem) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InvoiceItem) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *InvoiceItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnitPrice

`func (o *InvoiceItem) GetUnitPrice() string`

GetUnitPrice returns the UnitPrice field if non-nil, zero value otherwise.

### GetUnitPriceOk

`func (o *InvoiceItem) GetUnitPriceOk() (*string, bool)`

GetUnitPriceOk returns a tuple with the UnitPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPrice

`func (o *InvoiceItem) SetUnitPrice(v string)`

SetUnitPrice sets UnitPrice field to given value.

### HasUnitPrice

`func (o *InvoiceItem) HasUnitPrice() bool`

HasUnitPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


