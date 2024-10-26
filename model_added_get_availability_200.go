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

// checks if the AddedGetAvailability200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddedGetAvailability200{}

// AddedGetAvailability200 struct for AddedGetAvailability200
type AddedGetAvailability200 struct {
	Data []AddedGetAvailability200AllOfData `json:"data,omitempty"`
	// The current [page](https://techdocs.akamai.com/linode-api/reference/pagination).
	Page *int32 `json:"page,omitempty"`
	// The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination).
	Pages *int32 `json:"pages,omitempty"`
	// The total number of results.
	Results *int32 `json:"results,omitempty"`
}

// NewAddedGetAvailability200 instantiates a new AddedGetAvailability200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddedGetAvailability200() *AddedGetAvailability200 {
	this := AddedGetAvailability200{}
	return &this
}

// NewAddedGetAvailability200WithDefaults instantiates a new AddedGetAvailability200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddedGetAvailability200WithDefaults() *AddedGetAvailability200 {
	this := AddedGetAvailability200{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddedGetAvailability200) GetData() []AddedGetAvailability200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret []AddedGetAvailability200AllOfData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetAvailability200) GetDataOk() ([]AddedGetAvailability200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddedGetAvailability200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AddedGetAvailability200AllOfData and assigns it to the Data field.
func (o *AddedGetAvailability200) SetData(v []AddedGetAvailability200AllOfData) {
	o.Data = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AddedGetAvailability200) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetAvailability200) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AddedGetAvailability200) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *AddedGetAvailability200) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *AddedGetAvailability200) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetAvailability200) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *AddedGetAvailability200) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *AddedGetAvailability200) SetPages(v int32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AddedGetAvailability200) GetResults() int32 {
	if o == nil || IsNil(o.Results) {
		var ret int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetAvailability200) GetResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AddedGetAvailability200) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given int32 and assigns it to the Results field.
func (o *AddedGetAvailability200) SetResults(v int32) {
	o.Results = &v
}

func (o AddedGetAvailability200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddedGetAvailability200) ToMap() (map[string]interface{}, error) {
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

type NullableAddedGetAvailability200 struct {
	value *AddedGetAvailability200
	isSet bool
}

func (v NullableAddedGetAvailability200) Get() *AddedGetAvailability200 {
	return v.value
}

func (v *NullableAddedGetAvailability200) Set(val *AddedGetAvailability200) {
	v.value = val
	v.isSet = true
}

func (v NullableAddedGetAvailability200) IsSet() bool {
	return v.isSet
}

func (v *NullableAddedGetAvailability200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddedGetAvailability200(val *AddedGetAvailability200) *NullableAddedGetAvailability200 {
	return &NullableAddedGetAvailability200{value: val, isSet: true}
}

func (v NullableAddedGetAvailability200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddedGetAvailability200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


