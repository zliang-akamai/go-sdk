# PostExecutePayPalPaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayerId** | **string** | The PayerID returned by PayPal during the transaction authorization process. | 
**PaymentId** | **string** | The PaymentID returned from [Stage a PayPal payment](https://techdocs.akamai.com/linode-api/reference/post-pay-pal-payment) that has been approved with PayPal. | 

## Methods

### NewPostExecutePayPalPaymentRequest

`func NewPostExecutePayPalPaymentRequest(payerId string, paymentId string, ) *PostExecutePayPalPaymentRequest`

NewPostExecutePayPalPaymentRequest instantiates a new PostExecutePayPalPaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostExecutePayPalPaymentRequestWithDefaults

`func NewPostExecutePayPalPaymentRequestWithDefaults() *PostExecutePayPalPaymentRequest`

NewPostExecutePayPalPaymentRequestWithDefaults instantiates a new PostExecutePayPalPaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayerId

`func (o *PostExecutePayPalPaymentRequest) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *PostExecutePayPalPaymentRequest) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *PostExecutePayPalPaymentRequest) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetPaymentId

`func (o *PostExecutePayPalPaymentRequest) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PostExecutePayPalPaymentRequest) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PostExecutePayPalPaymentRequest) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


