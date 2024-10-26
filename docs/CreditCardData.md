# CreditCardData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardType** | Pointer to **string** | The type of credit card. | [optional] [readonly] 
**Expiry** | Pointer to **string** | The expiration month and year of the credit card. | [optional] [readonly] 
**LastFour** | Pointer to **string** | The last four digits of the credit card number. | [optional] [readonly] 

## Methods

### NewCreditCardData

`func NewCreditCardData() *CreditCardData`

NewCreditCardData instantiates a new CreditCardData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditCardDataWithDefaults

`func NewCreditCardDataWithDefaults() *CreditCardData`

NewCreditCardDataWithDefaults instantiates a new CreditCardData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardType

`func (o *CreditCardData) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *CreditCardData) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *CreditCardData) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *CreditCardData) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetExpiry

`func (o *CreditCardData) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *CreditCardData) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *CreditCardData) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *CreditCardData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetLastFour

`func (o *CreditCardData) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *CreditCardData) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *CreditCardData) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *CreditCardData) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


