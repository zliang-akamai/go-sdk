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

// checks if the AddedGetServiceTransfers200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddedGetServiceTransfers200{}

// AddedGetServiceTransfers200 struct for AddedGetServiceTransfers200
type AddedGetServiceTransfers200 struct {
	Data []GetServiceTransfers200ResponseDataInner `json:"data,omitempty"`
	// The current [page](https://techdocs.akamai.com/linode-api/reference/pagination).
	Page *int32 `json:"page,omitempty"`
	// The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination).
	Pages *int32 `json:"pages,omitempty"`
	// The total number of results.
	Results *int32 `json:"results,omitempty"`
}

// NewAddedGetServiceTransfers200 instantiates a new AddedGetServiceTransfers200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddedGetServiceTransfers200() *AddedGetServiceTransfers200 {
	this := AddedGetServiceTransfers200{}
	return &this
}

// NewAddedGetServiceTransfers200WithDefaults instantiates a new AddedGetServiceTransfers200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddedGetServiceTransfers200WithDefaults() *AddedGetServiceTransfers200 {
	this := AddedGetServiceTransfers200{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddedGetServiceTransfers200) GetData() []GetServiceTransfers200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetServiceTransfers200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetServiceTransfers200) GetDataOk() ([]GetServiceTransfers200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddedGetServiceTransfers200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetServiceTransfers200ResponseDataInner and assigns it to the Data field.
func (o *AddedGetServiceTransfers200) SetData(v []GetServiceTransfers200ResponseDataInner) {
	o.Data = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AddedGetServiceTransfers200) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetServiceTransfers200) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AddedGetServiceTransfers200) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *AddedGetServiceTransfers200) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *AddedGetServiceTransfers200) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetServiceTransfers200) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *AddedGetServiceTransfers200) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *AddedGetServiceTransfers200) SetPages(v int32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AddedGetServiceTransfers200) GetResults() int32 {
	if o == nil || IsNil(o.Results) {
		var ret int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetServiceTransfers200) GetResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AddedGetServiceTransfers200) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given int32 and assigns it to the Results field.
func (o *AddedGetServiceTransfers200) SetResults(v int32) {
	o.Results = &v
}

func (o AddedGetServiceTransfers200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddedGetServiceTransfers200) ToMap() (map[string]interface{}, error) {
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

type NullableAddedGetServiceTransfers200 struct {
	value *AddedGetServiceTransfers200
	isSet bool
}

func (v NullableAddedGetServiceTransfers200) Get() *AddedGetServiceTransfers200 {
	return v.value
}

func (v *NullableAddedGetServiceTransfers200) Set(val *AddedGetServiceTransfers200) {
	v.value = val
	v.isSet = true
}

func (v NullableAddedGetServiceTransfers200) IsSet() bool {
	return v.isSet
}

func (v *NullableAddedGetServiceTransfers200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddedGetServiceTransfers200(val *AddedGetServiceTransfers200) *NullableAddedGetServiceTransfers200 {
	return &NullableAddedGetServiceTransfers200{value: val, isSet: true}
}

func (v NullableAddedGetServiceTransfers200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddedGetServiceTransfers200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


