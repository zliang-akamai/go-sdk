# AddedPostPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**PostCreditCardRequest**](PostCreditCardRequest.md) |  | 
**IsDefault** | **bool** | Whether this Payment Method is the default method for automatically processing service charges. | 
**Type** | **string** | The type of Payment Method.  Alternative Payment Methods including Google Pay and PayPal can be added using the Cloud Manager. See the [Manage Payment Methods](https://www.linode.com/docs/products/platform/billing/guides/payment-methods/) guide for details and instructions. | 

## Methods

### NewAddedPostPaymentMethod

`func NewAddedPostPaymentMethod(data PostCreditCardRequest, isDefault bool, type_ string, ) *AddedPostPaymentMethod`

NewAddedPostPaymentMethod instantiates a new AddedPostPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedPostPaymentMethodWithDefaults

`func NewAddedPostPaymentMethodWithDefaults() *AddedPostPaymentMethod`

NewAddedPostPaymentMethodWithDefaults instantiates a new AddedPostPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AddedPostPaymentMethod) GetData() PostCreditCardRequest`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddedPostPaymentMethod) GetDataOk() (*PostCreditCardRequest, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddedPostPaymentMethod) SetData(v PostCreditCardRequest)`

SetData sets Data field to given value.


### GetIsDefault

`func (o *AddedPostPaymentMethod) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *AddedPostPaymentMethod) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *AddedPostPaymentMethod) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetType

`func (o *AddedPostPaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddedPostPaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddedPostPaymentMethod) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


