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

// checks if the AddedGetEntityTransfers200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddedGetEntityTransfers200{}

// AddedGetEntityTransfers200 struct for AddedGetEntityTransfers200
type AddedGetEntityTransfers200 struct {
	// The current [page](https://techdocs.akamai.com/linode-api/reference/pagination).
	Page *int32 `json:"page,omitempty"`
	// The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination).
	Pages *int32 `json:"pages,omitempty"`
	// The total number of results.
	Results *int32 `json:"results,omitempty"`
	Data []AddedGetEntityTransfers200AllOfData `json:"data,omitempty"`
}

// NewAddedGetEntityTransfers200 instantiates a new AddedGetEntityTransfers200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddedGetEntityTransfers200() *AddedGetEntityTransfers200 {
	this := AddedGetEntityTransfers200{}
	return &this
}

// NewAddedGetEntityTransfers200WithDefaults instantiates a new AddedGetEntityTransfers200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddedGetEntityTransfers200WithDefaults() *AddedGetEntityTransfers200 {
	this := AddedGetEntityTransfers200{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AddedGetEntityTransfers200) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetEntityTransfers200) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AddedGetEntityTransfers200) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *AddedGetEntityTransfers200) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *AddedGetEntityTransfers200) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetEntityTransfers200) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *AddedGetEntityTransfers200) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *AddedGetEntityTransfers200) SetPages(v int32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AddedGetEntityTransfers200) GetResults() int32 {
	if o == nil || IsNil(o.Results) {
		var ret int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetEntityTransfers200) GetResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AddedGetEntityTransfers200) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given int32 and assigns it to the Results field.
func (o *AddedGetEntityTransfers200) SetResults(v int32) {
	o.Results = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddedGetEntityTransfers200) GetData() []AddedGetEntityTransfers200AllOfData {
	if o == nil || IsNil(o.Data) {
		var ret []AddedGetEntityTransfers200AllOfData
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetEntityTransfers200) GetDataOk() ([]AddedGetEntityTransfers200AllOfData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddedGetEntityTransfers200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []AddedGetEntityTransfers200AllOfData and assigns it to the Data field.
func (o *AddedGetEntityTransfers200) SetData(v []AddedGetEntityTransfers200AllOfData) {
	o.Data = v
}

func (o AddedGetEntityTransfers200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddedGetEntityTransfers200) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Page) {
		toSerialize["page"] = o.Page
	}
	if !IsNil(o.Pages) {
		toSerialize["pages"] = o.Pages
	}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAddedGetEntityTransfers200 struct {
	value *AddedGetEntityTransfers200
	isSet bool
}

func (v NullableAddedGetEntityTransfers200) Get() *AddedGetEntityTransfers200 {
	return v.value
}

func (v *NullableAddedGetEntityTransfers200) Set(val *AddedGetEntityTransfers200) {
	v.value = val
	v.isSet = true
}

func (v NullableAddedGetEntityTransfers200) IsSet() bool {
	return v.isSet
}

func (v *NullableAddedGetEntityTransfers200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddedGetEntityTransfers200(val *AddedGetEntityTransfers200) *NullableAddedGetEntityTransfers200 {
	return &NullableAddedGetEntityTransfers200{value: val, isSet: true}
}

func (v NullableAddedGetEntityTransfers200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddedGetEntityTransfers200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


