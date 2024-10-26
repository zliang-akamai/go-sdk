# GetInvoices200ResponseDataInner

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

### NewGetInvoices200ResponseDataInner

`func NewGetInvoices200ResponseDataInner() *GetInvoices200ResponseDataInner`

NewGetInvoices200ResponseDataInner instantiates a new GetInvoices200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoices200ResponseDataInnerWithDefaults

`func NewGetInvoices200ResponseDataInnerWithDefaults() *GetInvoices200ResponseDataInner`

NewGetInvoices200ResponseDataInnerWithDefaults instantiates a new GetInvoices200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingSource

`func (o *GetInvoices200ResponseDataInner) GetBillingSource() string`

GetBillingSource returns the BillingSource field if non-nil, zero value otherwise.

### GetBillingSourceOk

`func (o *GetInvoices200ResponseDataInner) GetBillingSourceOk() (*string, bool)`

GetBillingSourceOk returns a tuple with the BillingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSource

`func (o *GetInvoices200ResponseDataInner) SetBillingSource(v string)`

SetBillingSource sets BillingSource field to given value.

### HasBillingSource

`func (o *GetInvoices200ResponseDataInner) HasBillingSource() bool`

HasBillingSource returns a boolean if a field has been set.

### GetDate

`func (o *GetInvoices200ResponseDataInner) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetInvoices200ResponseDataInner) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetInvoices200ResponseDataInner) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetInvoices200ResponseDataInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *GetInvoices200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInvoices200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInvoices200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetInvoices200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetInvoices200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetInvoices200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetInvoices200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetInvoices200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSubtotal

`func (o *GetInvoices200ResponseDataInner) GetSubtotal() float32`

GetSubtotal returns the Subtotal field if non-nil, zero value otherwise.

### GetSubtotalOk

`func (o *GetInvoices200ResponseDataInner) GetSubtotalOk() (*float32, bool)`

GetSubtotalOk returns a tuple with the Subtotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtotal

`func (o *GetInvoices200ResponseDataInner) SetSubtotal(v float32)`

SetSubtotal sets Subtotal field to given value.

### HasSubtotal

`func (o *GetInvoices200ResponseDataInner) HasSubtotal() bool`

HasSubtotal returns a boolean if a field has been set.

### GetTax

`func (o *GetInvoices200ResponseDataInner) GetTax() float32`

GetTax returns the Tax field if non-nil, zero value otherwise.

### GetTaxOk

`func (o *GetInvoices200ResponseDataInner) GetTaxOk() (*float32, bool)`

GetTaxOk returns a tuple with the Tax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTax

`func (o *GetInvoices200ResponseDataInner) SetTax(v float32)`

SetTax sets Tax field to given value.

### HasTax

`func (o *GetInvoices200ResponseDataInner) HasTax() bool`

HasTax returns a boolean if a field has been set.

### GetTaxSummary

`func (o *GetInvoices200ResponseDataInner) GetTaxSummary() []GetInvoices200ResponseDataInnerTaxSummaryInner`

GetTaxSummary returns the TaxSummary field if non-nil, zero value otherwise.

### GetTaxSummaryOk

`func (o *GetInvoices200ResponseDataInner) GetTaxSummaryOk() (*[]GetInvoices200ResponseDataInnerTaxSummaryInner, bool)`

GetTaxSummaryOk returns a tuple with the TaxSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxSummary

`func (o *GetInvoices200ResponseDataInner) SetTaxSummary(v []GetInvoices200ResponseDataInnerTaxSummaryInner)`

SetTaxSummary sets TaxSummary field to given value.

### HasTaxSummary

`func (o *GetInvoices200ResponseDataInner) HasTaxSummary() bool`

HasTaxSummary returns a boolean if a field has been set.

### GetTotal

`func (o *GetInvoices200ResponseDataInner) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetInvoices200ResponseDataInner) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetInvoices200ResponseDataInner) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetInvoices200ResponseDataInner) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


