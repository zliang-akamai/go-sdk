# GetPayments200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** | When the Payment was made. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID of the Payment. | [optional] [readonly] 
**Usd** | Pointer to **int32** | The amount, in US dollars, of the Payment. | [optional] [readonly] 

## Methods

### NewGetPayments200ResponseDataInner

`func NewGetPayments200ResponseDataInner() *GetPayments200ResponseDataInner`

NewGetPayments200ResponseDataInner instantiates a new GetPayments200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayments200ResponseDataInnerWithDefaults

`func NewGetPayments200ResponseDataInnerWithDefaults() *GetPayments200ResponseDataInner`

NewGetPayments200ResponseDataInnerWithDefaults instantiates a new GetPayments200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetPayments200ResponseDataInner) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetPayments200ResponseDataInner) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetPayments200ResponseDataInner) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetPayments200ResponseDataInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetId

`func (o *GetPayments200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPayments200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPayments200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetPayments200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsd

`func (o *GetPayments200ResponseDataInner) GetUsd() int32`

GetUsd returns the Usd field if non-nil, zero value otherwise.

### GetUsdOk

`func (o *GetPayments200ResponseDataInner) GetUsdOk() (*int32, bool)`

GetUsdOk returns a tuple with the Usd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsd

`func (o *GetPayments200ResponseDataInner) SetUsd(v int32)`

SetUsd sets Usd field to given value.

### HasUsd

`func (o *GetPayments200ResponseDataInner) HasUsd() bool`

HasUsd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


