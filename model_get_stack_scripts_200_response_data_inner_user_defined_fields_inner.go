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
	"bytes"
	"fmt"
)

// checks if the GetStackScripts200ResponseDataInnerUserDefinedFieldsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStackScripts200ResponseDataInnerUserDefinedFieldsInner{}

// GetStackScripts200ResponseDataInnerUserDefinedFieldsInner A custom field defined by the User with a special syntax within a StackScript. Derived from the contents of the script.
type GetStackScripts200ResponseDataInnerUserDefinedFieldsInner struct {
	// The default value.  If not specified, this value will be used.
	Default *string `json:"default,omitempty"`
	// An example value for the field.
	Example string `json:"example"`
	// A human-readable label for the field that will serve as the input prompt for entering the value during deployment.
	Label string `json:"label"`
	// A list of acceptable values for the field in any quantity, combination or order.
	ManyOf *string `json:"manyOf,omitempty"`
	// The name of the field.
	Name string `json:"name"`
	// A list of acceptable single values for the field.
	OneOf *string `json:"oneOf,omitempty"`
}

type _GetStackScripts200ResponseDataInnerUserDefinedFieldsInner GetStackScripts200ResponseDataInnerUserDefinedFieldsInner

// NewGetStackScripts200ResponseDataInnerUserDefinedFieldsInner instantiates a new GetStackScripts200ResponseDataInnerUserDefinedFieldsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStackScripts200ResponseDataInnerUserDefinedFieldsInner(example string, label string, name string) *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner {
	this := GetStackScripts200ResponseDataInnerUserDefinedFieldsInner{}
	this.Example = example
	this.Label = label
	this.Name = name
	return &this
}

// NewGetStackScripts200ResponseDataInnerUserDefinedFieldsInnerWithDefaults instantiates a new GetStackScripts200ResponseDataInnerUserDefinedFieldsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStackScripts200ResponseDataInnerUserDefinedFieldsInnerWithDefaults() *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner {
	this := GetStackScripts200ResponseDataInnerUserDefinedFieldsInner{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetDefault() string {
	if o == nil || IsNil(o.Default) {
		var ret string
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetDefaultOk() (*string, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given string and assigns it to the Default field.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) SetDefault(v string) {
	o.Default = &v
}

// GetExample returns the Example field value
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetExample() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Example
}

// GetExampleOk returns a tuple with the Example field value
// and a boolean to check if the value has been set.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetExampleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Example, true
}

// SetExample sets field value
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) SetExample(v string) {
	o.Example = v
}

// GetLabel returns the Label field value
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) SetLabel(v string) {
	o.Label = v
}

// GetManyOf returns the ManyOf field value if set, zero value otherwise.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetManyOf() string {
	if o == nil || IsNil(o.ManyOf) {
		var ret string
		return ret
	}
	return *o.ManyOf
}

// GetManyOfOk returns a tuple with the ManyOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetManyOfOk() (*string, bool) {
	if o == nil || IsNil(o.ManyOf) {
		return nil, false
	}
	return o.ManyOf, true
}

// HasManyOf returns a boolean if a field has been set.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) HasManyOf() bool {
	if o != nil && !IsNil(o.ManyOf) {
		return true
	}

	return false
}

// SetManyOf gets a reference to the given string and assigns it to the ManyOf field.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) SetManyOf(v string) {
	o.ManyOf = &v
}

// GetName returns the Name field value
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) SetName(v string) {
	o.Name = v
}

// GetOneOf returns the OneOf field value if set, zero value otherwise.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetOneOf() string {
	if o == nil || IsNil(o.OneOf) {
		var ret string
		return ret
	}
	return *o.OneOf
}

// GetOneOfOk returns a tuple with the OneOf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) GetOneOfOk() (*string, bool) {
	if o == nil || IsNil(o.OneOf) {
		return nil, false
	}
	return o.OneOf, true
}

// HasOneOf returns a boolean if a field has been set.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) HasOneOf() bool {
	if o != nil && !IsNil(o.OneOf) {
		return true
	}

	return false
}

// SetOneOf gets a reference to the given string and assigns it to the OneOf field.
func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) SetOneOf(v string) {
	o.OneOf = &v
}

func (o GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	toSerialize["example"] = o.Example
	toSerialize["label"] = o.Label
	if !IsNil(o.ManyOf) {
		toSerialize["manyOf"] = o.ManyOf
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.OneOf) {
		toSerialize["oneOf"] = o.OneOf
	}
	return toSerialize, nil
}

func (o *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"example",
		"label",
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varGetStackScripts200ResponseDataInnerUserDefinedFieldsInner := _GetStackScripts200ResponseDataInnerUserDefinedFieldsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStackScripts200ResponseDataInnerUserDefinedFieldsInner)

	if err != nil {
		return err
	}

	*o = GetStackScripts200ResponseDataInnerUserDefinedFieldsInner(varGetStackScripts200ResponseDataInnerUserDefinedFieldsInner)

	return err
}

type NullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner struct {
	value *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner
	isSet bool
}

func (v NullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner) Get() *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner {
	return v.value
}

func (v *NullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner) Set(val *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner(val *GetStackScripts200ResponseDataInnerUserDefinedFieldsInner) *NullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner {
	return &NullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner{value: val, isSet: true}
}

func (v NullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStackScripts200ResponseDataInnerUserDefinedFieldsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


