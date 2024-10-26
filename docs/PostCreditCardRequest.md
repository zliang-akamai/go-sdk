# PostCreditCardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardNumber** | **string** | Your credit card number. No spaces or hyphens (&#x60;-&#x60;) allowed. | 
**Cvv** | **string** | CVV (Card Verification Value) of the credit card, typically found on the back of the card. | 
**ExpiryMonth** | **int32** | A value from 1-12 representing the expiration month of your credit card.    - 1 &#x3D; January   - 2 &#x3D; February   - 3 &#x3D; March   - Etc. | 
**ExpiryYear** | **int32** | A four-digit integer representing the expiration year of your credit card.  The combination of &#x60;expiry_month&#x60; and &#x60;expiry_year&#x60; must result in a month/year combination of the current month or in the future. An expiration date set in the past is invalid. | 

## Methods

### NewPostCreditCardRequest

`func NewPostCreditCardRequest(cardNumber string, cvv string, expiryMonth int32, expiryYear int32, ) *PostCreditCardRequest`

NewPostCreditCardRequest instantiates a new PostCreditCardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreditCardRequestWithDefaults

`func NewPostCreditCardRequestWithDefaults() *PostCreditCardRequest`

NewPostCreditCardRequestWithDefaults instantiates a new PostCreditCardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardNumber

`func (o *PostCreditCardRequest) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *PostCreditCardRequest) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *PostCreditCardRequest) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetCvv

`func (o *PostCreditCardRequest) GetCvv() string`

GetCvv returns the Cvv field if non-nil, zero value otherwise.

### GetCvvOk

`func (o *PostCreditCardRequest) GetCvvOk() (*string, bool)`

GetCvvOk returns a tuple with the Cvv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCvv

`func (o *PostCreditCardRequest) SetCvv(v string)`

SetCvv sets Cvv field to given value.


### GetExpiryMonth

`func (o *PostCreditCardRequest) GetExpiryMonth() int32`

GetExpiryMonth returns the ExpiryMonth field if non-nil, zero value otherwise.

### GetExpiryMonthOk

`func (o *PostCreditCardRequest) GetExpiryMonthOk() (*int32, bool)`

GetExpiryMonthOk returns a tuple with the ExpiryMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryMonth

`func (o *PostCreditCardRequest) SetExpiryMonth(v int32)`

SetExpiryMonth sets ExpiryMonth field to given value.


### GetExpiryYear

`func (o *PostCreditCardRequest) GetExpiryYear() int32`

GetExpiryYear returns the ExpiryYear field if non-nil, zero value otherwise.

### GetExpiryYearOk

`func (o *PostCreditCardRequest) GetExpiryYearOk() (*int32, bool)`

GetExpiryYearOk returns a tuple with the ExpiryYear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiryYear

`func (o *PostCreditCardRequest) SetExpiryYear(v int32)`

SetExpiryYear sets ExpiryYear field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


