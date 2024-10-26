# PaypalData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The email address associated with your PayPal account. | [optional] [readonly] 
**PaypalId** | Pointer to **string** | PayPal Merchant ID associated with your PayPal account. | [optional] [readonly] 

## Methods

### NewPaypalData

`func NewPaypalData() *PaypalData`

NewPaypalData instantiates a new PaypalData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaypalDataWithDefaults

`func NewPaypalDataWithDefaults() *PaypalData`

NewPaypalDataWithDefaults instantiates a new PaypalData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *PaypalData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *PaypalData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *PaypalData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *PaypalData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPaypalId

`func (o *PaypalData) GetPaypalId() string`

GetPaypalId returns the PaypalId field if non-nil, zero value otherwise.

### GetPaypalIdOk

`func (o *PaypalData) GetPaypalIdOk() (*string, bool)`

GetPaypalIdOk returns a tuple with the PaypalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalId

`func (o *PaypalData) SetPaypalId(v string)`

SetPaypalId sets PaypalId field to given value.

### HasPaypalId

`func (o *PaypalData) HasPaypalId() bool`

HasPaypalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


