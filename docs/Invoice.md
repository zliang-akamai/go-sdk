# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingSource** | Pointer to **string** | &#x60;akamai&#x60;: This Invoice was generated according to the terms of an agreement between the customer and Akamai.  &#x60;linode&#x60;: This Invoice was generated according to the default terms, prices, and discounts. | [optional] [readonly] 
**Date** | Pointer to **time.Time** | When this Invoice was generated. | [optional] [readonly] 
**Id** | Pointer to **int32** | The Invoice&#39;s unique ID. | [optional] [readonly] 
**Label** | Pointer to **string** | The Invoice&#39;s display label. | [optional] [readonly] 
**Subtotal** | Pointer to **float32** | The amount of the Invoice before taxes in US Dollars. | [optional] [readonly] 
**Tax** | Pointer to **float32** | The amount of tax levied on the Invoice in US Dollars. | [optional] [readonly] 
**TaxSummary** | Pointer to [**[]GetInvoices200ResponseDataInnerTaxSummaryInner**](GetInvoices200ResponseDataInnerTaxSummaryInner.md) | The amount of tax broken down into subtotals by source. | [optional] [readonly] 
**Total** | Pointer to **float32** | The amount of the Invoice after taxes in US Dollars. | [optional] [readonly] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingSource

`func (o *Invoice) GetBillingSource() string`

GetBillingSource returns the BillingSource field if non-nil, zero value otherwise.

### GetBillingSourceOk

`func (o *Invoice) GetBillingSourceOk() (*string, bool)`

GetBillingSourceOk returns a tuple with the BillingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSource

`func (o *Invoice) SetBillingSource(v string)`

SetBillingSource sets BillingSource field to given value.

### HasBillingSource

`func (o *Invoice) HasBillingSource() bool`

HasBillingSource returns a boolean if a field has been set.

### GetDate

`func (o *Invoice) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Invoice) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Invoice) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *Invoice) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *Invoice) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *Invoice) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Invoice) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Invoice) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Invoice) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSubtotal

`func (o *Invoice) GetSubtotal() float32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *Invoice) GetSubtotalOk() (*float32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *Invoice) SetSubtotal(v float32)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *Invoice) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTax

`func (o *Invoice) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *Invoice) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *Invoice) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *Invoice) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxSummary

`func (o *Invoice) GetTaxSummary() []GetInvoices200ResponseDataInnerTaxSummaryInner`

GetTaxSummary returns the TaxSummary field if non-nil, zero value otherwise.

### GetTaxSummaryOk

`func (o *Invoice) GetTaxSummaryOk() (*[]GetInvoices200ResponseDataInnerTaxSummaryInner, bool)`

GetTaxSummaryOk returns a tuple with the TaxSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSummary

`func (o *Invoice) SetTaxSummary(v []GetInvoices200ResponseDataInnerTaxSummaryInner)`

SetTaxSummary sets TaxSummary field to given value.

### HasTaxSummary

`func (o *Invoice) HasTaxSummary() bool`

HasTaxSummary returns a boolean if a field has been set.

### GetTotal

`func (o *Invoice) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Invoice) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Invoice) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Invoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


