# PostPayPalPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelUrl** | **string** | The URL to have PayPal redirect to when Payment is canceled. | 
**RedirectUrl** | **string** | The URL to have PayPal redirect to when Payment is approved. | 
**Usd** | **string** | The payment amount in USD. Minimum accepted value of $5 USD. Maximum accepted value of $500 USD or credit card payment limit; whichever value is highest. PayPal&#39;s maximum transaction limit is $10,000 USD. | 

## Methods

### NewPostPayPalPaymentRequest

`func NewPostPayPalPaymentRequest(cancelUrl string, redirectUrl string, usd string, ) *PostPayPalPaymentRequest`

NewPostPayPalPaymentRequest instantiates a new PostPayPalPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPayPalPaymentRequestWithDefaults

`func NewPostPayPalPaymentRequestWithDefaults() *PostPayPalPaymentRequest`

NewPostPayPalPaymentRequestWithDefaults instantiates a new PostPayPalPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelUrl

`func (o *PostPayPalPaymentRequest) GetCancelUrl() string`

GetCancelUrl returns the CancelUrl field if non-nil, zero value otherwise.

### GetCancelUrlOk

`func (o *PostPayPalPaymentRequest) GetCancelUrlOk() (*string, bool)`

GetCancelUrlOk returns a tuple with the CancelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelUrl

`func (o *PostPayPalPaymentRequest) SetCancelUrl(v string)`

SetCancelUrl sets CancelUrl field to given value.


### GetRedirectUrl

`func (o *PostPayPalPaymentRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *PostPayPalPaymentRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *PostPayPalPaymentRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.


### GetUsd

`func (o *PostPayPalPaymentRequest) GetUsd() string`

GetUsd returns the Usd field if non-nil, zero value otherwise.

### GetUsdOk

`func (o *PostPayPalPaymentRequest) GetUsdOk() (*string, bool)`

GetUsdOk returns a tuple with the Usd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsd

`func (o *PostPayPalPaymentRequest) SetUsd(v string)`

SetUsd sets Usd field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


