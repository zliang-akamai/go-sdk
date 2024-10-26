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

// checks if the GetIpv6Ranges200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIpv6Ranges200Response{}

// GetIpv6Ranges200Response struct for GetIpv6Ranges200Response
type GetIpv6Ranges200Response struct {
	Data []GetLinodeIps200ResponseIpv6GlobalInner `json:"data,omitempty"`
	// The current [page](https://techdocs.akamai.com/linode-api/reference/pagination).
	Page *int32 `json:"page,omitempty"`
	// The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination).
	Pages *int32 `json:"pages,omitempty"`
	// The total number of results.
	Results *int32 `json:"results,omitempty"`
}

// NewGetIpv6Ranges200Response instantiates a new GetIpv6Ranges200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIpv6Ranges200Response() *GetIpv6Ranges200Response {
	this := GetIpv6Ranges200Response{}
	return &this
}

// NewGetIpv6Ranges200ResponseWithDefaults instantiates a new GetIpv6Ranges200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIpv6Ranges200ResponseWithDefaults() *GetIpv6Ranges200Response {
	this := GetIpv6Ranges200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetIpv6Ranges200Response) GetData() []GetLinodeIps200ResponseIpv6GlobalInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetLinodeIps200ResponseIpv6GlobalInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpv6Ranges200Response) GetDataOk() ([]GetLinodeIps200ResponseIpv6GlobalInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetIpv6Ranges200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetLinodeIps200ResponseIpv6GlobalInner and assigns it to the Data field.
func (o *GetIpv6Ranges200Response) SetData(v []GetLinodeIps200ResponseIpv6GlobalInner) {
	o.Data = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetIpv6Ranges200Response) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpv6Ranges200Response) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetIpv6Ranges200Response) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetIpv6Ranges200Response) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *GetIpv6Ranges200Response) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpv6Ranges200Response) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *GetIpv6Ranges200Response) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *GetIpv6Ranges200Response) SetPages(v int32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetIpv6Ranges200Response) GetResults() int32 {
	if o == nil || IsNil(o.Results) {
		var ret int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetIpv6Ranges200Response) GetResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetIpv6Ranges200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given int32 and assigns it to the Results field.
func (o *GetIpv6Ranges200Response) SetResults(v int32) {
	o.Results = &v
}

func (o GetIpv6Ranges200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIpv6Ranges200Response) ToMap() (map[string]interface{}, error) {
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

type NullableGetIpv6Ranges200Response struct {
	value *GetIpv6Ranges200Response
	isSet bool
}

func (v NullableGetIpv6Ranges200Response) Get() *GetIpv6Ranges200Response {
	return v.value
}

func (v *NullableGetIpv6Ranges200Response) Set(val *GetIpv6Ranges200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIpv6Ranges200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIpv6Ranges200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIpv6Ranges200Response(val *GetIpv6Ranges200Response) *NullableGetIpv6Ranges200Response {
	return &NullableGetIpv6Ranges200Response{value: val, isSet: true}
}

func (v NullableGetIpv6Ranges200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIpv6Ranges200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


