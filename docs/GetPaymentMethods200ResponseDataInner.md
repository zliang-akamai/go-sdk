# GetPaymentMethods200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When the Payment Method was added to the Account. | [optional] [readonly] 
**Data** | Pointer to [**GetPaymentMethods200ResponseDataInnerData**](GetPaymentMethods200ResponseDataInnerData.md) |  | [optional] 
**Id** | Pointer to **int32** | The unique ID of this Payment Method. | [optional] 
**IsDefault** | Pointer to **bool** | Whether this Payment Method is the default method for automatically processing service charges. | [optional] 
**Type** | Pointer to **string** | The type of Payment Method. | [optional] 

## Methods

### NewGetPaymentMethods200ResponseDataInner

`func NewGetPaymentMethods200ResponseDataInner() *GetPaymentMethods200ResponseDataInner`

NewGetPaymentMethods200ResponseDataInner instantiates a new GetPaymentMethods200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPaymentMethods200ResponseDataInnerWithDefaults

`func NewGetPaymentMethods200ResponseDataInnerWithDefaults() *GetPaymentMethods200ResponseDataInner`

NewGetPaymentMethods200ResponseDataInnerWithDefaults instantiates a new GetPaymentMethods200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetPaymentMethods200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetPaymentMethods200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetPaymentMethods200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetPaymentMethods200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetData

`func (o *GetPaymentMethods200ResponseDataInner) GetData() GetPaymentMethods200ResponseDataInnerData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetPaymentMethods200ResponseDataInner) GetDataOk() (*GetPaymentMethods200ResponseDataInnerData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetPaymentMethods200ResponseDataInner) SetData(v GetPaymentMethods200ResponseDataInnerData)`

SetData sets Data field to given value.

### HasData

`func (o *GetPaymentMethods200ResponseDataInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *GetPaymentMethods200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPaymentMethods200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPaymentMethods200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetPaymentMethods200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *GetPaymentMethods200ResponseDataInner) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetPaymentMethods200ResponseDataInner) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetPaymentMethods200ResponseDataInner) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *GetPaymentMethods200ResponseDataInner) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetType

`func (o *GetPaymentMethods200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetPaymentMethods200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetPaymentMethods200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetPaymentMethods200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


