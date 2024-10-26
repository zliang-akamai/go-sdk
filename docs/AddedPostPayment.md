# AddedPostPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentMethodId** | Pointer to **int32** | The ID of the Payment Method to apply to the Payment. | [optional] 
**Usd** | Pointer to **string** | The amount in US Dollars of the Payment.  - Can begin with or without &#x60;$&#x60;. - Commas (&#x60;,&#x60;) are not accepted. - Must end with a decimal expression, such as &#x60;.00&#x60; or &#x60;.99&#x60;. - Minimum: &#x60;$5.00&#x60; or the Account balance, whichever is lower. - Maximum: &#x60;$2000.00&#x60; or the Account balance up to &#x60;$50000.00&#x60;, whichever is greater. | [optional] 

## Methods

### NewAddedPostPayment

`func NewAddedPostPayment() *AddedPostPayment`

NewAddedPostPayment instantiates a new AddedPostPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedPostPaymentWithDefaults

`func NewAddedPostPaymentWithDefaults() *AddedPostPayment`

NewAddedPostPaymentWithDefaults instantiates a new AddedPostPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentMethodId

`func (o *AddedPostPayment) GetPaymentMethodId() int32`

GetPaymentMethodId returns the PaymentMethodId field if non-nil, zero value otherwise.

### GetPaymentMethodIdOk

`func (o *AddedPostPayment) GetPaymentMethodIdOk() (*int32, bool)`

GetPaymentMethodIdOk returns a tuple with the PaymentMethodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethodId

`func (o *AddedPostPayment) SetPaymentMethodId(v int32)`

SetPaymentMethodId sets PaymentMethodId field to given value.

### HasPaymentMethodId

`func (o *AddedPostPayment) HasPaymentMethodId() bool`

HasPaymentMethodId returns a boolean if a field has been set.

### GetUsd

`func (o *AddedPostPayment) GetUsd() string`

GetUsd returns the Usd field if non-nil, zero value otherwise.

### GetUsdOk

`func (o *AddedPostPayment) GetUsdOk() (*string, bool)`

GetUsdOk returns a tuple with the Usd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsd

`func (o *AddedPostPayment) SetUsd(v string)`

SetUsd sets Usd field to given value.

### HasUsd

`func (o *AddedPostPayment) HasUsd() bool`

HasUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

