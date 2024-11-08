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

// checks if the AddedGetInvoiceItems200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddedGetInvoiceItems200{}

// AddedGetInvoiceItems200 struct for AddedGetInvoiceItems200
type AddedGetInvoiceItems200 struct {
	Data []GetInvoiceItems200ResponseDataInner `json:"data,omitempty"`
	// The current [page](https://techdocs.akamai.com/linode-api/reference/pagination).
	Page *int32 `json:"page,omitempty"`
	// The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination).
	Pages *int32 `json:"pages,omitempty"`
	// The total number of results.
	Results *int32 `json:"results,omitempty"`
}

// NewAddedGetInvoiceItems200 instantiates a new AddedGetInvoiceItems200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddedGetInvoiceItems200() *AddedGetInvoiceItems200 {
	this := AddedGetInvoiceItems200{}
	return &this
}

// NewAddedGetInvoiceItems200WithDefaults instantiates a new AddedGetInvoiceItems200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddedGetInvoiceItems200WithDefaults() *AddedGetInvoiceItems200 {
	this := AddedGetInvoiceItems200{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddedGetInvoiceItems200) GetData() []GetInvoiceItems200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetInvoiceItems200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetInvoiceItems200) GetDataOk() ([]GetInvoiceItems200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddedGetInvoiceItems200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetInvoiceItems200ResponseDataInner and assigns it to the Data field.
func (o *AddedGetInvoiceItems200) SetData(v []GetInvoiceItems200ResponseDataInner) {
	o.Data = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AddedGetInvoiceItems200) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetInvoiceItems200) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AddedGetInvoiceItems200) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *AddedGetInvoiceItems200) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *AddedGetInvoiceItems200) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetInvoiceItems200) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *AddedGetInvoiceItems200) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *AddedGetInvoiceItems200) SetPages(v int32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AddedGetInvoiceItems200) GetResults() int32 {
	if o == nil || IsNil(o.Results) {
		var ret int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetInvoiceItems200) GetResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AddedGetInvoiceItems200) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given int32 and assigns it to the Results field.
func (o *AddedGetInvoiceItems200) SetResults(v int32) {
	o.Results = &v
}

func (o AddedGetInvoiceItems200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddedGetInvoiceItems200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	return toSerialize, nil
}

type NullableAddedGetInvoiceItems200 struct {
	value *AddedGetInvoiceItems200
	isSet bool
}

func (v NullableAddedGetInvoiceItems200) Get() *AddedGetInvoiceItems200 {
	return v.value
}

func (v *NullableAddedGetInvoiceItems200) Set(val *AddedGetInvoiceItems200) {
	v.value = val
	v.isSet = true
}

func (v NullableAddedGetInvoiceItems200) IsSet() bool {
	return v.isSet
}

func (v *NullableAddedGetInvoiceItems200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddedGetInvoiceItems200(val *AddedGetInvoiceItems200) *NullableAddedGetInvoiceItems200 {
	return &NullableAddedGetInvoiceItems200{value: val, isSet: true}
}

func (v NullableAddedGetInvoiceItems200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddedGetInvoiceItems200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


