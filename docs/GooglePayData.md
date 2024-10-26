# GooglePayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardType** | Pointer to **string** | The type of credit card. | [optional] [readonly] 
**Expiry** | Pointer to **string** | The expiration month and year of the credit card. | [optional] [readonly] 
**LastFour** | Pointer to **string** | The last four digits of the credit card number. | [optional] [readonly] 

## Methods

### NewGooglePayData

`func NewGooglePayData() *GooglePayData`

NewGooglePayData instantiates a new GooglePayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGooglePayDataWithDefaults

`func NewGooglePayDataWithDefaults() *GooglePayData`

NewGooglePayDataWithDefaults instantiates a new GooglePayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardType

`func (o *GooglePayData) GetCardType() string`

GetCardType returns the CardType field if non-nil, zero value otherwise.

### GetCardTypeOk

`func (o *GooglePayData) GetCardTypeOk() (*string, bool)`

GetCardTypeOk returns a tuple with the CardType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardType

`func (o *GooglePayData) SetCardType(v string)`

SetCardType sets CardType field to given value.

### HasCardType

`func (o *GooglePayData) HasCardType() bool`

HasCardType returns a boolean if a field has been set.

### GetExpiry

`func (o *GooglePayData) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GooglePayData) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GooglePayData) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GooglePayData) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetLastFour

`func (o *GooglePayData) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *GooglePayData) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *GooglePayData) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *GooglePayData) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


