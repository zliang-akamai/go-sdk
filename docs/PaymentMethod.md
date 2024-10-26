# PaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When the Payment Method was added to the Account. | [optional] [readonly] 
**Data** | Pointer to [**GetPaymentMethods200ResponseDataInnerData**](GetPaymentMethods200ResponseDataInnerData.md) |  | [optional] 
**Id** | Pointer to **int32** | The unique ID of this Payment Method. | [optional] 
**IsDefault** | Pointer to **bool** | Whether this Payment Method is the default method for automatically processing service charges. | [optional] 
**Type** | Pointer to **string** | The type of Payment Method. | [optional] 

## Methods

### NewPaymentMethod

`func NewPaymentMethod() *PaymentMethod`

NewPaymentMethod instantiates a new PaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentMethodWithDefaults

`func NewPaymentMethodWithDefaults() *PaymentMethod`

NewPaymentMethodWithDefaults instantiates a new PaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PaymentMethod) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PaymentMethod) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PaymentMethod) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PaymentMethod) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *PaymentMethod) GetData() GetPaymentMethods200ResponseDataInnerData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentMethod) GetDataOk() (*GetPaymentMethods200ResponseDataInnerData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentMethod) SetData(v GetPaymentMethods200ResponseDataInnerData)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentMethod) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *PaymentMethod) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentMethod) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentMethod) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentMethod) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *PaymentMethod) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *PaymentMethod) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *PaymentMethod) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *PaymentMethod) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetType

`func (o *PaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


