# AddedPostPayPalPayment200

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutToken** | Pointer to **string** | The checkout token generated for this Payment. | [optional] [readonly] 
**PaymentId** | Pointer to **string** | The paypal-generated ID for this Payment. Used when authorizing the Payment in PayPal&#39;s interface. | [optional] 

## Methods

### NewAddedPostPayPalPayment200

`func NewAddedPostPayPalPayment200() *AddedPostPayPalPayment200`

NewAddedPostPayPalPayment200 instantiates a new AddedPostPayPalPayment200 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedPostPayPalPayment200WithDefaults

`func NewAddedPostPayPalPayment200WithDefaults() *AddedPostPayPalPayment200`

NewAddedPostPayPalPayment200WithDefaults instantiates a new AddedPostPayPalPayment200 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutToken

`func (o *AddedPostPayPalPayment200) GetCheckoutToken() string`

GetCheckoutToken returns the CheckoutToken field if non-nil, zero value otherwise.

### GetCheckoutTokenOk

`func (o *AddedPostPayPalPayment200) GetCheckoutTokenOk() (*string, bool)`

GetCheckoutTokenOk returns a tuple with the CheckoutToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutToken

`func (o *AddedPostPayPalPayment200) SetCheckoutToken(v string)`

SetCheckoutToken sets CheckoutToken field to given value.

### HasCheckoutToken

`func (o *AddedPostPayPalPayment200) HasCheckoutToken() bool`

HasCheckoutToken returns a boolean if a field has been set.

### GetPaymentId

`func (o *AddedPostPayPalPayment200) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *AddedPostPayPalPayment200) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *AddedPostPayPalPayment200) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *AddedPostPayPalPayment200) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


