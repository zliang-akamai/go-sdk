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
	"gopkg.in/validator.v2"
	"fmt"
)

// GetLinodeInstancesXFilterParameter - Specifies the `X-Filter` header JSON object's filtering and sort criteria.
type GetLinodeInstancesXFilterParameter struct {
	FilterAndSortCriteria *FilterAndSortCriteria
	MapmapOfStringAny *map[string]interface{}
}

// FilterAndSortCriteriaAsGetLinodeInstancesXFilterParameter is a convenience function that returns FilterAndSortCriteria wrapped in GetLinodeInstancesXFilterParameter
func FilterAndSortCriteriaAsGetLinodeInstancesXFilterParameter(v *FilterAndSortCriteria) GetLinodeInstancesXFilterParameter {
	return GetLinodeInstancesXFilterParameter{
		FilterAndSortCriteria: v,
	}
}

// map[string]interface{}AsGetLinodeInstancesXFilterParameter is a convenience function that returns map[string]interface{} wrapped in GetLinodeInstancesXFilterParameter
func MapmapOfStringAnyAsGetLinodeInstancesXFilterParameter(v *map[string]interface{}) GetLinodeInstancesXFilterParameter {
	return GetLinodeInstancesXFilterParameter{
		MapmapOfStringAny: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetLinodeInstancesXFilterParameter) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FilterAndSortCriteria
	err = newStrictDecoder(data).Decode(&dst.FilterAndSortCriteria)
	if err == nil {
		jsonFilterAndSortCriteria, _ := json.Marshal(dst.FilterAndSortCriteria)
		if string(jsonFilterAndSortCriteria) == "{}" { // empty struct
			dst.FilterAndSortCriteria = nil
		} else {
			if err = validator.Validate(dst.FilterAndSortCriteria); err != nil {
				dst.FilterAndSortCriteria = nil
			} else {
				match++
			}
		}
	} else {
		dst.FilterAndSortCriteria = nil
	}

	// try to unmarshal data into MapmapOfStringAny
	err = newStrictDecoder(data).Decode(&dst.MapmapOfStringAny)
	if err == nil {
		jsonMapmapOfStringAny, _ := json.Marshal(dst.MapmapOfStringAny)
		if string(jsonMapmapOfStringAny) == "{}" { // empty struct
			dst.MapmapOfStringAny = nil
		} else {
			if err = validator.Validate(dst.MapmapOfStringAny); err != nil {
				dst.MapmapOfStringAny = nil
			} else {
				match++
			}
		}
	} else {
		dst.MapmapOfStringAny = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FilterAndSortCriteria = nil
		dst.MapmapOfStringAny = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetLinodeInstancesXFilterParameter)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetLinodeInstancesXFilterParameter)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetLinodeInstancesXFilterParameter) MarshalJSON() ([]byte, error) {
	if src.FilterAndSortCriteria != nil {
		return json.Marshal(&src.FilterAndSortCriteria)
	}

	if src.MapmapOfStringAny != nil {
		return json.Marshal(&src.MapmapOfStringAny)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetLinodeInstancesXFilterParameter) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.FilterAndSortCriteria != nil {
		return obj.FilterAndSortCriteria
	}

	if obj.MapmapOfStringAny != nil {
		return obj.MapmapOfStringAny
	}

	// all schemas are nil
	return nil
}

type NullableGetLinodeInstancesXFilterParameter struct {
	value *GetLinodeInstancesXFilterParameter
	isSet bool
}

func (v NullableGetLinodeInstancesXFilterParameter) Get() *GetLinodeInstancesXFilterParameter {
	return v.value
}

func (v *NullableGetLinodeInstancesXFilterParameter) Set(val *GetLinodeInstancesXFilterParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLinodeInstancesXFilterParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLinodeInstancesXFilterParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLinodeInstancesXFilterParameter(val *GetLinodeInstancesXFilterParameter) *NullableGetLinodeInstancesXFilterParameter {
	return &NullableGetLinodeInstancesXFilterParameter{value: val, isSet: true}
}

func (v NullableGetLinodeInstancesXFilterParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLinodeInstancesXFilterParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

