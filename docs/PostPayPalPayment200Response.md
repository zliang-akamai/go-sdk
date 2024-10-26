# PostPayPalPayment200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutToken** | Pointer to **string** | The checkout token generated for this Payment. | [optional] [readonly] 
**PaymentId** | Pointer to **string** | The paypal-generated ID for this Payment. Used when authorizing the Payment in PayPal&#39;s interface. | [optional] 

## Methods

### NewPostPayPalPayment200Response

`func NewPostPayPalPayment200Response() *PostPayPalPayment200Response`

NewPostPayPalPayment200Response instantiates a new PostPayPalPayment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPayPalPayment200ResponseWithDefaults

`func NewPostPayPalPayment200ResponseWithDefaults() *PostPayPalPayment200Response`

NewPostPayPalPayment200ResponseWithDefaults instantiates a new PostPayPalPayment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutToken

`func (o *PostPayPalPayment200Response) GetCheckoutToken() string`

GetCheckoutToken returns the CheckoutToken field if non-nil, zero value otherwise.

### GetCheckoutTokenOk

`func (o *PostPayPalPayment200Response) GetCheckoutTokenOk() (*string, bool)`

GetCheckoutTokenOk returns a tuple with the CheckoutToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutToken

`func (o *PostPayPalPayment200Response) SetCheckoutToken(v string)`

SetCheckoutToken sets CheckoutToken field to given value.

### HasCheckoutToken

`func (o *PostPayPalPayment200Response) HasCheckoutToken() bool`

HasCheckoutToken returns a boolean if a field has been set.

### GetPaymentId

`func (o *PostPayPalPayment200Response) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PostPayPalPayment200Response) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PostPayPalPayment200Response) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *PostPayPalPayment200Response) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


