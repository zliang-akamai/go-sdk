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

// checks if the GetVpcSubnets200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetVpcSubnets200Response{}

// GetVpcSubnets200Response struct for GetVpcSubnets200Response
type GetVpcSubnets200Response struct {
	// The current [page](https://techdocs.akamai.com/linode-api/reference/pagination).
	Page *int32 `json:"page,omitempty"`
	// The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination).
	Pages *int32 `json:"pages,omitempty"`
	// The total number of results.
	Results *int32 `json:"results,omitempty"`
	Data []GetVpcs200ResponseAllOfDataInnerSubnetsInner `json:"data,omitempty"`
}

// NewGetVpcSubnets200Response instantiates a new GetVpcSubnets200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetVpcSubnets200Response() *GetVpcSubnets200Response {
	this := GetVpcSubnets200Response{}
	return &this
}

// NewGetVpcSubnets200ResponseWithDefaults instantiates a new GetVpcSubnets200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetVpcSubnets200ResponseWithDefaults() *GetVpcSubnets200Response {
	this := GetVpcSubnets200Response{}
	return &this
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetVpcSubnets200Response) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcSubnets200Response) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetVpcSubnets200Response) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetVpcSubnets200Response) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *GetVpcSubnets200Response) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcSubnets200Response) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *GetVpcSubnets200Response) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *GetVpcSubnets200Response) SetPages(v int32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetVpcSubnets200Response) GetResults() int32 {
	if o == nil || IsNil(o.Results) {
		var ret int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcSubnets200Response) GetResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetVpcSubnets200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given int32 and assigns it to the Results field.
func (o *GetVpcSubnets200Response) SetResults(v int32) {
	o.Results = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetVpcSubnets200Response) GetData() []GetVpcs200ResponseAllOfDataInnerSubnetsInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetVpcs200ResponseAllOfDataInnerSubnetsInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetVpcSubnets200Response) GetDataOk() ([]GetVpcs200ResponseAllOfDataInnerSubnetsInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetVpcSubnets200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetVpcs200ResponseAllOfDataInnerSubnetsInner and assigns it to the Data field.
func (o *GetVpcSubnets200Response) SetData(v []GetVpcs200ResponseAllOfDataInnerSubnetsInner) {
	o.Data = v
}

func (o GetVpcSubnets200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetVpcSubnets200Response) ToMap() (map[string]interface{}, error) {
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

type NullableGetVpcSubnets200Response struct {
	value *GetVpcSubnets200Response
	isSet bool
}

func (v NullableGetVpcSubnets200Response) Get() *GetVpcSubnets200Response {
	return v.value
}

func (v *NullableGetVpcSubnets200Response) Set(val *GetVpcSubnets200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetVpcSubnets200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetVpcSubnets200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetVpcSubnets200Response(val *GetVpcSubnets200Response) *NullableGetVpcSubnets200Response {
	return &NullableGetVpcSubnets200Response{value: val, isSet: true}
}

func (v NullableGetVpcSubnets200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetVpcSubnets200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


