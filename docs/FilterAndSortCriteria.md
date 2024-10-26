# FilterAndSortCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**And** | Pointer to **[]map[string]interface{}** | All conditions need to be true. | [optional] 
**Contains** | Pointer to **string** | The provided string needs to be in the value. | [optional] 
**Gt** | Pointer to **float32** | The value needs to be greater than the provided number. | [optional] 
**Gte** | Pointer to **float32** | The value needs to be greater than or equal to the provided number. | [optional] 
**Lt** | Pointer to **float32** | The value needs to be less than the provided number. | [optional] 
**Lte** | Pointer to **float32** | The value needs to be less than or equal to the provided number. | [optional] 
**Neq** | Pointer to **string** | The provided string is left out of the results. | [optional] 
**Or** | Pointer to **[]map[string]interface{}** | At least one condition needs to be true. | [optional] 
**Order** | Pointer to **string** | Sort in ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;) order. This defaults to &#x60;asc&#x60;. Requires &#x60;+order_by&#x60;. | [optional] [default to "asc"]
**OrderBy** | Pointer to **string** | Order results based on the provided attribute. The attribute needs to be filterable. | [optional] 

## Methods

### NewFilterAndSortCriteria

`func NewFilterAndSortCriteria() *FilterAndSortCriteria`

NewFilterAndSortCriteria instantiates a new FilterAndSortCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterAndSortCriteriaWithDefaults

`func NewFilterAndSortCriteriaWithDefaults() *FilterAndSortCriteria`

NewFilterAndSortCriteriaWithDefaults instantiates a new FilterAndSortCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnd

`func (o *FilterAndSortCriteria) GetAnd() []map[string]interface{}`

GetAnd returns the And field if non-nil, zero value otherwise.

### GetAndOk

`func (o *FilterAndSortCriteria) GetAndOk() (*[]map[string]interface{}, bool)`

GetAndOk returns a tuple with the And field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnd

`func (o *FilterAndSortCriteria) SetAnd(v []map[string]interface{})`

SetAnd sets And field to given value.

### HasAnd

`func (o *FilterAndSortCriteria) HasAnd() bool`

HasAnd returns a boolean if a field has been set.

### GetContains

`func (o *FilterAndSortCriteria) GetContains() string`

GetContains returns the Contains field if non-nil, zero value otherwise.

### GetContainsOk

`func (o *FilterAndSortCriteria) GetContainsOk() (*string, bool)`

GetContainsOk returns a tuple with the Contains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContains

`func (o *FilterAndSortCriteria) SetContains(v string)`

SetContains sets Contains field to given value.

### HasContains

`func (o *FilterAndSortCriteria) HasContains() bool`

HasContains returns a boolean if a field has been set.

### GetGt

`func (o *FilterAndSortCriteria) GetGt() float32`

GetGt returns the Gt field if non-nil, zero value otherwise.

### GetGtOk

`func (o *FilterAndSortCriteria) GetGtOk() (*float32, bool)`

GetGtOk returns a tuple with the Gt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGt

`func (o *FilterAndSortCriteria) SetGt(v float32)`

SetGt sets Gt field to given value.

### HasGt

`func (o *FilterAndSortCriteria) HasGt() bool`

HasGt returns a boolean if a field has been set.

### GetGte

`func (o *FilterAndSortCriteria) GetGte() float32`

GetGte returns the Gte field if non-nil, zero value otherwise.

### GetGteOk

`func (o *FilterAndSortCriteria) GetGteOk() (*float32, bool)`

GetGteOk returns a tuple with the Gte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGte

`func (o *FilterAndSortCriteria) SetGte(v float32)`

SetGte sets Gte field to given value.

### HasGte

`func (o *FilterAndSortCriteria) HasGte() bool`

HasGte returns a boolean if a field has been set.

### GetLt

`func (o *FilterAndSortCriteria) GetLt() float32`

GetLt returns the Lt field if non-nil, zero value otherwise.

### GetLtOk

`func (o *FilterAndSortCriteria) GetLtOk() (*float32, bool)`

GetLtOk returns a tuple with the Lt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLt

`func (o *FilterAndSortCriteria) SetLt(v float32)`

SetLt sets Lt field to given value.

### HasLt

`func (o *FilterAndSortCriteria) HasLt() bool`

HasLt returns a boolean if a field has been set.

### GetLte

`func (o *FilterAndSortCriteria) GetLte() float32`

GetLte returns the Lte field if non-nil, zero value otherwise.

### GetLteOk

`func (o *FilterAndSortCriteria) GetLteOk() (*float32, bool)`

GetLteOk returns a tuple with the Lte field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLte

`func (o *FilterAndSortCriteria) SetLte(v float32)`

SetLte sets Lte field to given value.

### HasLte

`func (o *FilterAndSortCriteria) HasLte() bool`

HasLte returns a boolean if a field has been set.

### GetNeq

`func (o *FilterAndSortCriteria) GetNeq() string`

GetNeq returns the Neq field if non-nil, zero value otherwise.

### GetNeqOk

`func (o *FilterAndSortCriteria) GetNeqOk() (*string, bool)`

GetNeqOk returns a tuple with the Neq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeq

`func (o *FilterAndSortCriteria) SetNeq(v string)`

SetNeq sets Neq field to given value.

### HasNeq

`func (o *FilterAndSortCriteria) HasNeq() bool`

HasNeq returns a boolean if a field has been set.

### GetOr

`func (o *FilterAndSortCriteria) GetOr() []map[string]interface{}`

GetOr returns the Or field if non-nil, zero value otherwise.

### GetOrOk

`func (o *FilterAndSortCriteria) GetOrOk() (*[]map[string]interface{}, bool)`

GetOrOk returns a tuple with the Or field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOr

`func (o *FilterAndSortCriteria) SetOr(v []map[string]interface{})`

SetOr sets Or field to given value.

### HasOr

`func (o *FilterAndSortCriteria) HasOr() bool`

HasOr returns a boolean if a field has been set.

### GetOrder

`func (o *FilterAndSortCriteria) GetOrder() string`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *FilterAndSortCriteria) GetOrderOk() (*string, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *FilterAndSortCriteria) SetOrder(v string)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *FilterAndSortCriteria) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetOrderBy

`func (o *FilterAndSortCriteria) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *FilterAndSortCriteria) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *FilterAndSortCriteria) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *FilterAndSortCriteria) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


