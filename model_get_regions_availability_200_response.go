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

// checks if the GetRegionsAvailability200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRegionsAvailability200Response{}

// GetRegionsAvailability200Response struct for GetRegionsAvailability200Response
type GetRegionsAvailability200Response struct {
	// The current [page](https://techdocs.akamai.com/linode-api/reference/pagination).
	Page *int32 `json:"page,omitempty"`
	// The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination).
	Pages *int32 `json:"pages,omitempty"`
	// The total number of results.
	Results *int32 `json:"results,omitempty"`
	Data []GetRegionsAvailability200ResponseAllOfDataInner `json:"data,omitempty"`
}

// NewGetRegionsAvailability200Response instantiates a new GetRegionsAvailability200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRegionsAvailability200Response() *GetRegionsAvailability200Response {
	this := GetRegionsAvailability200Response{}
	return &this
}

// NewGetRegionsAvailability200ResponseWithDefaults instantiates a new GetRegionsAvailability200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRegionsAvailability200ResponseWithDefaults() *GetRegionsAvailability200Response {
	this := GetRegionsAvailability200Response{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetRegionsAvailability200Response) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionsAvailability200Response) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetRegionsAvailability200Response) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetRegionsAvailability200Response) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *GetRegionsAvailability200Response) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionsAvailability200Response) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *GetRegionsAvailability200Response) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *GetRegionsAvailability200Response) SetPages(v int32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetRegionsAvailability200Response) GetResults() int32 {
	if o == nil || IsNil(o.Results) {
		var ret int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionsAvailability200Response) GetResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetRegionsAvailability200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given int32 and assigns it to the Results field.
func (o *GetRegionsAvailability200Response) SetResults(v int32) {
	o.Results = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRegionsAvailability200Response) GetData() []GetRegionsAvailability200ResponseAllOfDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetRegionsAvailability200ResponseAllOfDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRegionsAvailability200Response) GetDataOk() ([]GetRegionsAvailability200ResponseAllOfDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRegionsAvailability200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetRegionsAvailability200ResponseAllOfDataInner and assigns it to the Data field.
func (o *GetRegionsAvailability200Response) SetData(v []GetRegionsAvailability200ResponseAllOfDataInner) {
	o.Data = v
}

func (o GetRegionsAvailability200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRegionsAvailability200Response) ToMap() (map[string]interface{}, error) {
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

type NullableGetRegionsAvailability200Response struct {
	value *GetRegionsAvailability200Response
	isSet bool
}

func (v NullableGetRegionsAvailability200Response) Get() *GetRegionsAvailability200Response {
	return v.value
}

func (v *NullableGetRegionsAvailability200Response) Set(val *GetRegionsAvailability200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRegionsAvailability200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRegionsAvailability200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRegionsAvailability200Response(val *GetRegionsAvailability200Response) *NullableGetRegionsAvailability200Response {
	return &NullableGetRegionsAvailability200Response{value: val, isSet: true}
}

func (v NullableGetRegionsAvailability200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRegionsAvailability200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

