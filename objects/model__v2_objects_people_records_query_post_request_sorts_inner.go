/*
Attio Standard Objects

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// V2ObjectsPeopleRecordsQueryPostRequestSortsInner struct for V2ObjectsPeopleRecordsQueryPostRequestSortsInner
type V2ObjectsPeopleRecordsQueryPostRequestSortsInner struct {
	V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf
	V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *V2ObjectsPeopleRecordsQueryPostRequestSortsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf
	err = json.Unmarshal(data, &dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf);
	if err == nil {
		jsonV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf, _ := json.Marshal(dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf)
		if string(jsonV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf) == "{}" { // empty struct
			dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf = nil
		} else {
			return nil // data stored in dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf, return on the first match
		}
	} else {
		dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf = nil
	}

	// try to unmarshal JSON data into V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1
	err = json.Unmarshal(data, &dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1);
	if err == nil {
		jsonV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1, _ := json.Marshal(dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1)
		if string(jsonV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) == "{}" { // empty struct
			dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 = nil
		} else {
			return nil // data stored in dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1, return on the first match
		}
	} else {
		dst.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(V2ObjectsPeopleRecordsQueryPostRequestSortsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *V2ObjectsPeopleRecordsQueryPostRequestSortsInner) MarshalJSON() ([]byte, error) {
	if src.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf != nil {
		return json.Marshal(&src.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf)
	}

	if src.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 != nil {
		return json.Marshal(&src.V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner struct {
	value *V2ObjectsPeopleRecordsQueryPostRequestSortsInner
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner) Get() *V2ObjectsPeopleRecordsQueryPostRequestSortsInner {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner) Set(val *V2ObjectsPeopleRecordsQueryPostRequestSortsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner(val *V2ObjectsPeopleRecordsQueryPostRequestSortsInner) *NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner {
	return &NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


