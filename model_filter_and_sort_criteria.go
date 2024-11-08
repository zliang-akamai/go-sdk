/*
Linode API

[Read the API documentation](https://techdocs.akamai.com/linode-api/reference/api).

API version: 4.189.2
Contact: support@linode.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the FilterAndSortCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FilterAndSortCriteria{}

// FilterAndSortCriteria struct for FilterAndSortCriteria
type FilterAndSortCriteria struct {
	// All conditions need to be true.
	And []map[string]interface{} `json:"+and,omitempty"`
	// The provided string needs to be in the value.
	Contains *string `json:"+contains,omitempty"`
	// The value needs to be greater than the provided number.
	Gt *float32 `json:"+gt,omitempty"`
	// The value needs to be greater than or equal to the provided number.
	Gte *float32 `json:"+gte,omitempty"`
	// The value needs to be less than the provided number.
	Lt *float32 `json:"+lt,omitempty"`
	// The value needs to be less than or equal to the provided number.
	Lte *float32 `json:"+lte,omitempty"`
	// The provided string is left out of the results.
	Neq *string `json:"+neq,omitempty"`
	// At least one condition needs to be true.
	Or []map[string]interface{} `json:"+or,omitempty"`
	// Sort in ascending (`asc`) or descending (`desc`) order. This defaults to `asc`. Requires `+order_by`.
	Order *string `json:"+order,omitempty"`
	// Order results based on the provided attribute. The attribute needs to be filterable.
	OrderBy *string `json:"+order_by,omitempty"`
}

// NewFilterAndSortCriteria instantiates a new FilterAndSortCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilterAndSortCriteria() *FilterAndSortCriteria {
	this := FilterAndSortCriteria{}
	var order string = "asc"
	this.Order = &order
	return &this
}

// NewFilterAndSortCriteriaWithDefaults instantiates a new FilterAndSortCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilterAndSortCriteriaWithDefaults() *FilterAndSortCriteria {
	this := FilterAndSortCriteria{}
	var order string = "asc"
	this.Order = &order
	return &this
}

// GetAnd returns the And field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetAnd() []map[string]interface{} {
	if o == nil || IsNil(o.And) {
		var ret []map[string]interface{}
		return ret
	}
	return o.And
}

// GetAndOk returns a tuple with the And field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetAndOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.And) {
		return nil, false
	}
	return o.And, true
}

// HasAnd returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasAnd() bool {
	if o != nil && !IsNil(o.And) {
		return true
	}

	return false
}

// SetAnd gets a reference to the given []map[string]interface{} and assigns it to the And field.
func (o *FilterAndSortCriteria) SetAnd(v []map[string]interface{}) {
	o.And = v
}

// GetContains returns the Contains field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetContains() string {
	if o == nil || IsNil(o.Contains) {
		var ret string
		return ret
	}
	return *o.Contains
}

// GetContainsOk returns a tuple with the Contains field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetContainsOk() (*string, bool) {
	if o == nil || IsNil(o.Contains) {
		return nil, false
	}
	return o.Contains, true
}

// HasContains returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasContains() bool {
	if o != nil && !IsNil(o.Contains) {
		return true
	}

	return false
}

// SetContains gets a reference to the given string and assigns it to the Contains field.
func (o *FilterAndSortCriteria) SetContains(v string) {
	o.Contains = &v
}

// GetGt returns the Gt field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetGt() float32 {
	if o == nil || IsNil(o.Gt) {
		var ret float32
		return ret
	}
	return *o.Gt
}

// GetGtOk returns a tuple with the Gt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetGtOk() (*float32, bool) {
	if o == nil || IsNil(o.Gt) {
		return nil, false
	}
	return o.Gt, true
}

// HasGt returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasGt() bool {
	if o != nil && !IsNil(o.Gt) {
		return true
	}

	return false
}

// SetGt gets a reference to the given float32 and assigns it to the Gt field.
func (o *FilterAndSortCriteria) SetGt(v float32) {
	o.Gt = &v
}

// GetGte returns the Gte field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetGte() float32 {
	if o == nil || IsNil(o.Gte) {
		var ret float32
		return ret
	}
	return *o.Gte
}

// GetGteOk returns a tuple with the Gte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetGteOk() (*float32, bool) {
	if o == nil || IsNil(o.Gte) {
		return nil, false
	}
	return o.Gte, true
}

// HasGte returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasGte() bool {
	if o != nil && !IsNil(o.Gte) {
		return true
	}

	return false
}

// SetGte gets a reference to the given float32 and assigns it to the Gte field.
func (o *FilterAndSortCriteria) SetGte(v float32) {
	o.Gte = &v
}

// GetLt returns the Lt field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetLt() float32 {
	if o == nil || IsNil(o.Lt) {
		var ret float32
		return ret
	}
	return *o.Lt
}

// GetLtOk returns a tuple with the Lt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetLtOk() (*float32, bool) {
	if o == nil || IsNil(o.Lt) {
		return nil, false
	}
	return o.Lt, true
}

// HasLt returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasLt() bool {
	if o != nil && !IsNil(o.Lt) {
		return true
	}

	return false
}

// SetLt gets a reference to the given float32 and assigns it to the Lt field.
func (o *FilterAndSortCriteria) SetLt(v float32) {
	o.Lt = &v
}

// GetLte returns the Lte field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetLte() float32 {
	if o == nil || IsNil(o.Lte) {
		var ret float32
		return ret
	}
	return *o.Lte
}

// GetLteOk returns a tuple with the Lte field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetLteOk() (*float32, bool) {
	if o == nil || IsNil(o.Lte) {
		return nil, false
	}
	return o.Lte, true
}

// HasLte returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasLte() bool {
	if o != nil && !IsNil(o.Lte) {
		return true
	}

	return false
}

// SetLte gets a reference to the given float32 and assigns it to the Lte field.
func (o *FilterAndSortCriteria) SetLte(v float32) {
	o.Lte = &v
}

// GetNeq returns the Neq field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetNeq() string {
	if o == nil || IsNil(o.Neq) {
		var ret string
		return ret
	}
	return *o.Neq
}

// GetNeqOk returns a tuple with the Neq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetNeqOk() (*string, bool) {
	if o == nil || IsNil(o.Neq) {
		return nil, false
	}
	return o.Neq, true
}

// HasNeq returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasNeq() bool {
	if o != nil && !IsNil(o.Neq) {
		return true
	}

	return false
}

// SetNeq gets a reference to the given string and assigns it to the Neq field.
func (o *FilterAndSortCriteria) SetNeq(v string) {
	o.Neq = &v
}

// GetOr returns the Or field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetOr() []map[string]interface{} {
	if o == nil || IsNil(o.Or) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Or
}

// GetOrOk returns a tuple with the Or field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetOrOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Or) {
		return nil, false
	}
	return o.Or, true
}

// HasOr returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasOr() bool {
	if o != nil && !IsNil(o.Or) {
		return true
	}

	return false
}

// SetOr gets a reference to the given []map[string]interface{} and assigns it to the Or field.
func (o *FilterAndSortCriteria) SetOr(v []map[string]interface{}) {
	o.Or = v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetOrder() string {
	if o == nil || IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetOrderOk() (*string, bool) {
	if o == nil || IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasOrder() bool {
	if o != nil && !IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the Order field.
func (o *FilterAndSortCriteria) SetOrder(v string) {
	o.Order = &v
}

// GetOrderBy returns the OrderBy field value if set, zero value otherwise.
func (o *FilterAndSortCriteria) GetOrderBy() string {
	if o == nil || IsNil(o.OrderBy) {
		var ret string
		return ret
	}
	return *o.OrderBy
}

// GetOrderByOk returns a tuple with the OrderBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilterAndSortCriteria) GetOrderByOk() (*string, bool) {
	if o == nil || IsNil(o.OrderBy) {
		return nil, false
	}
	return o.OrderBy, true
}

// HasOrderBy returns a boolean if a field has been set.
func (o *FilterAndSortCriteria) HasOrderBy() bool {
	if o != nil && !IsNil(o.OrderBy) {
		return true
	}

	return false
}

// SetOrderBy gets a reference to the given string and assigns it to the OrderBy field.
func (o *FilterAndSortCriteria) SetOrderBy(v string) {
	o.OrderBy = &v
}

func (o FilterAndSortCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FilterAndSortCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.And) {
		toSerialize["+and"] = o.And
	}
	if !IsNil(o.Contains) {
		toSerialize["+contains"] = o.Contains
	}
	if !IsNil(o.Gt) {
		toSerialize["+gt"] = o.Gt
	}
	if !IsNil(o.Gte) {
		toSerialize["+gte"] = o.Gte
	}
	if !IsNil(o.Lt) {
		toSerialize["+lt"] = o.Lt
	}
	if !IsNil(o.Lte) {
		toSerialize["+lte"] = o.Lte
	}
	if !IsNil(o.Neq) {
		toSerialize["+neq"] = o.Neq
	}
	if !IsNil(o.Or) {
		toSerialize["+or"] = o.Or
	}
	if !IsNil(o.Order) {
		toSerialize["+order"] = o.Order
	}
	if !IsNil(o.OrderBy) {
		toSerialize["+order_by"] = o.OrderBy
	}
	return toSerialize, nil
}

type NullableFilterAndSortCriteria struct {
	value *FilterAndSortCriteria
	isSet bool
}

func (v NullableFilterAndSortCriteria) Get() *FilterAndSortCriteria {
	return v.value
}

func (v *NullableFilterAndSortCriteria) Set(val *FilterAndSortCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableFilterAndSortCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableFilterAndSortCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilterAndSortCriteria(val *FilterAndSortCriteria) *NullableFilterAndSortCriteria {
	return &NullableFilterAndSortCriteria{value: val, isSet: true}
}

func (v NullableFilterAndSortCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilterAndSortCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


