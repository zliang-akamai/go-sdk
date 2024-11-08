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

// checks if the GetInvoiceItems200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInvoiceItems200Response{}

// GetInvoiceItems200Response struct for GetInvoiceItems200Response
type GetInvoiceItems200Response struct {
	Data []GetInvoiceItems200ResponseDataInner `json:"data,omitempty"`
	// The current [page](https://techdocs.akamai.com/linode-api/reference/pagination).
	Page *int32 `json:"page,omitempty"`
	// The total number of [pages](https://techdocs.akamai.com/linode-api/reference/pagination).
	Pages *int32 `json:"pages,omitempty"`
	// The total number of results.
	Results *int32 `json:"results,omitempty"`
}

// NewGetInvoiceItems200Response instantiates a new GetInvoiceItems200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInvoiceItems200Response() *GetInvoiceItems200Response {
	this := GetInvoiceItems200Response{}
	return &this
}

// NewGetInvoiceItems200ResponseWithDefaults instantiates a new GetInvoiceItems200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInvoiceItems200ResponseWithDefaults() *GetInvoiceItems200Response {
	this := GetInvoiceItems200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetInvoiceItems200Response) GetData() []GetInvoiceItems200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetInvoiceItems200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoiceItems200Response) GetDataOk() ([]GetInvoiceItems200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetInvoiceItems200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetInvoiceItems200ResponseDataInner and assigns it to the Data field.
func (o *GetInvoiceItems200Response) SetData(v []GetInvoiceItems200ResponseDataInner) {
	o.Data = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *GetInvoiceItems200Response) GetPage() int32 {
	if o == nil || IsNil(o.Page) {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoiceItems200Response) GetPageOk() (*int32, bool) {
	if o == nil || IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *GetInvoiceItems200Response) HasPage() bool {
	if o != nil && !IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *GetInvoiceItems200Response) SetPage(v int32) {
	o.Page = &v
}

// GetPages returns the Pages field value if set, zero value otherwise.
func (o *GetInvoiceItems200Response) GetPages() int32 {
	if o == nil || IsNil(o.Pages) {
		var ret int32
		return ret
	}
	return *o.Pages
}

// GetPagesOk returns a tuple with the Pages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoiceItems200Response) GetPagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Pages) {
		return nil, false
	}
	return o.Pages, true
}

// HasPages returns a boolean if a field has been set.
func (o *GetInvoiceItems200Response) HasPages() bool {
	if o != nil && !IsNil(o.Pages) {
		return true
	}

	return false
}

// SetPages gets a reference to the given int32 and assigns it to the Pages field.
func (o *GetInvoiceItems200Response) SetPages(v int32) {
	o.Pages = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *GetInvoiceItems200Response) GetResults() int32 {
	if o == nil || IsNil(o.Results) {
		var ret int32
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetInvoiceItems200Response) GetResultsOk() (*int32, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *GetInvoiceItems200Response) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given int32 and assigns it to the Results field.
func (o *GetInvoiceItems200Response) SetResults(v int32) {
	o.Results = &v
}

func (o GetInvoiceItems200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInvoiceItems200Response) ToMap() (map[string]interface{}, error) {
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

type NullableGetInvoiceItems200Response struct {
	value *GetInvoiceItems200Response
	isSet bool
}

func (v NullableGetInvoiceItems200Response) Get() *GetInvoiceItems200Response {
	return v.value
}

func (v *NullableGetInvoiceItems200Response) Set(val *GetInvoiceItems200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInvoiceItems200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInvoiceItems200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInvoiceItems200Response(val *GetInvoiceItems200Response) *NullableGetInvoiceItems200Response {
	return &NullableGetInvoiceItems200Response{value: val, isSet: true}
}

func (v NullableGetInvoiceItems200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInvoiceItems200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


