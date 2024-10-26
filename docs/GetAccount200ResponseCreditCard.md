# GetAccount200ResponseCreditCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **string** | The expiration month and year of the credit card. | [optional] 
**LastFour** | Pointer to **string** | The last four digits of the credit card associated with this Account. | [optional] 

## Methods

### NewGetAccount200ResponseCreditCard

`func NewGetAccount200ResponseCreditCard() *GetAccount200ResponseCreditCard`

NewGetAccount200ResponseCreditCard instantiates a new GetAccount200ResponseCreditCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccount200ResponseCreditCardWithDefaults

`func NewGetAccount200ResponseCreditCardWithDefaults() *GetAccount200ResponseCreditCard`

NewGetAccount200ResponseCreditCardWithDefaults instantiates a new GetAccount200ResponseCreditCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *GetAccount200ResponseCreditCard) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GetAccount200ResponseCreditCard) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GetAccount200ResponseCreditCard) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GetAccount200ResponseCreditCard) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetLastFour

`func (o *GetAccount200ResponseCreditCard) GetLastFour() string`

GetLastFour returns the LastFour field if non-nil, zero value otherwise.

### GetLastFourOk

`func (o *GetAccount200ResponseCreditCard) GetLastFourOk() (*string, bool)`

GetLastFourOk returns a tuple with the LastFour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFour

`func (o *GetAccount200ResponseCreditCard) SetLastFour(v string)`

SetLastFour sets LastFour field to given value.

### HasLastFour

`func (o *GetAccount200ResponseCreditCard) HasLastFour() bool`

HasLastFour returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


