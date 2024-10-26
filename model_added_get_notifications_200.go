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

// checks if the AddedGetNotifications200 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddedGetNotifications200{}

// AddedGetNotifications200 struct for AddedGetNotifications200
type AddedGetNotifications200 struct {
	Data []GetNotifications200ResponseDataInner `json:"data,omitempty"`
	// The current [page](https://techdocs.akamai.com/linode-api/reference/pagination).
	Page *int32 `json:"page,omitempty"`
	// The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination).
	Pages *int32 `json:"pages,omitempty"`
	// The total number of results.
	Results *int32 `json:"results,omitempty"`
}

// NewAddedGetNotifications200 instantiates a new AddedGetNotifications200 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddedGetNotifications200() *AddedGetNotifications200 {
	this := AddedGetNotifications200{}
	return &this
}

// NewAddedGetNotifications200WithDefaults instantiates a new AddedGetNotifications200 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddedGetNotifications200WithDefaults() *AddedGetNotifications200 {
	this := AddedGetNotifications200{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AddedGetNotifications200) GetData() []GetNotifications200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetNotifications200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetNotifications200) GetDataOk() ([]GetNotifications200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AddedGetNotifications200) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetNotifications200ResponseDataInner and assigns it to the Data field.
func (o *AddedGetNotifications200) SetData(v []GetNotifications200ResponseDataInner) {
	o.Data = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *AddedGetNotifications200) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetNotifications200) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AddedGetNotifications200) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *AddedGetNotifications200) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *AddedGetNotifications200) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetNotifications200) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *AddedGetNotifications200) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *AddedGetNotifications200) SetPages(v int32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *AddedGetNotifications200) GetResults() int32 {
	if o == nil || IsNil(o.Results) {
		var ret int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddedGetNotifications200) GetResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *AddedGetNotifications200) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given int32 and assigns it to the Results field.
func (o *AddedGetNotifications200) SetResults(v int32) {
	o.Results = &v
}

func (o AddedGetNotifications200) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddedGetNotifications200) ToMap() (map[string]interface{}, error) {
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

type NullableAddedGetNotifications200 struct {
	value *AddedGetNotifications200
	isSet bool
}

func (v NullableAddedGetNotifications200) Get() *AddedGetNotifications200 {
	return v.value
}

func (v *NullableAddedGetNotifications200) Set(val *AddedGetNotifications200) {
	v.value = val
	v.isSet = true
}

func (v NullableAddedGetNotifications200) IsSet() bool {
	return v.isSet
}

func (v *NullableAddedGetNotifications200) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddedGetNotifications200(val *AddedGetNotifications200) *NullableAddedGetNotifications200 {
	return &NullableAddedGetNotifications200{value: val, isSet: true}
}

func (v NullableAddedGetNotifications200) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddedGetNotifications200) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

