# GetPaymentMethods200ResponseDataInnerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardType** | Pointer to **string** | The type of credit card. | [optional] [readonly] 
**Expiry** | Pointer to **string** | The expiration month and year of the credit card. | [optional] [readonly] 
**LastFour** | Pointer to **string** | The last four digits of the credit card number. | [optional] [readonly] 
**Email** | Pointer to **string** | The email address associated with your PayPal account. | [optional] [readonly] 
**PaypalId** | Pointer to **string** | PayPal Merchant ID associated with your PayPal account. | [optional] [readonly] 

## Methods

### NewGetPaymentMethods200ResponseDataInnerData

`func NewGetPaymentMethods200ResponseDataInnerData() *GetPaymentMethods200ResponseDataInnerData`

NewGetPaymentMethods200ResponseDataInnerData instantiates a new GetPaymentMethods200ResponseDataInnerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentMethods200ResponseDataInnerDataWithDefaults

`func NewGetPaymentMethods200ResponseDataInnerDataWithDefaults() *GetPaymentMethods200ResponseDataInnerData`

NewGetPaymentMethods200ResponseDataInnerDataWithDefaults instantiates a new GetPaymentMethods200ResponseDataInnerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardType

`func (o *GetPaymentMethods200ResponseDataInnerData) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *GetPaymentMethods200ResponseDataInnerData) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *GetPaymentMethods200ResponseDataInnerData) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *GetPaymentMethods200ResponseDataInnerData) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetExpiry

`func (o *GetPaymentMethods200ResponseDataInnerData) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GetPaymentMethods200ResponseDataInnerData) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GetPaymentMethods200ResponseDataInnerData) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GetPaymentMethods200ResponseDataInnerData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetLastFour

`func (o *GetPaymentMethods200ResponseDataInnerData) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *GetPaymentMethods200ResponseDataInnerData) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *GetPaymentMethods200ResponseDataInnerData) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *GetPaymentMethods200ResponseDataInnerData) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.

### GetEmail

`func (o *GetPaymentMethods200ResponseDataInnerData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetPaymentMethods200ResponseDataInnerData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetPaymentMethods200ResponseDataInnerData) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetPaymentMethods200ResponseDataInnerData) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPaypalId

`func (o *GetPaymentMethods200ResponseDataInnerData) GetPaypalId() string`

GetPaypalId returns the PaypalId field if non-nil, zero value otherwise.

### GetPaypalIdOk

`func (o *GetPaymentMethods200ResponseDataInnerData) GetPaypalIdOk() (*string, bool)`

GetPaypalIdOk returns a tuple with the PaypalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaypalId

`func (o *GetPaymentMethods200ResponseDataInnerData) SetPaypalId(v string)`

SetPaypalId sets PaypalId field to given value.

### HasPaypalId

`func (o *GetPaymentMethods200ResponseDataInnerData) HasPaypalId() bool`

HasPaypalId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


