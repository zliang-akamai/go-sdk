# PaypalExecute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayerId** | **string** | The PayerID returned by PayPal during the transaction authorization process. | 
**PaymentId** | **string** | The PaymentID returned from [Stage a PayPal payment](https://techdocs.akamai.com/linode-api/reference/post-pay-pal-payment) that has been approved with PayPal. | 

## Methods

### NewPaypalExecute

`func NewPaypalExecute(payerId string, paymentId string, ) *PaypalExecute`

NewPaypalExecute instantiates a new PaypalExecute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalExecuteWithDefaults

`func NewPaypalExecuteWithDefaults() *PaypalExecute`

NewPaypalExecuteWithDefaults instantiates a new PaypalExecute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayerId

`func (o *PaypalExecute) GetPayerId() string`

GetPayerId returns the PayerId field if non-nil, zero value otherwise.

### GetPayerIdOk

`func (o *PaypalExecute) GetPayerIdOk() (*string, bool)`

GetPayerIdOk returns a tuple with the PayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerId

`func (o *PaypalExecute) SetPayerId(v string)`

SetPayerId sets PayerId field to given value.


### GetPaymentId

`func (o *PaypalExecute) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *PaypalExecute) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *PaypalExecute) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


