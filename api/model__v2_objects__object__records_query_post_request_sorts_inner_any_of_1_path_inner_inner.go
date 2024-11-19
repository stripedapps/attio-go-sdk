/*
Attio API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
Contact: support@attio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)


// V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner struct for V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner
type V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner struct {
	value *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner
	isSet bool
}

func (v NullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) Get() *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner {
	return v.value
}

func (v *NullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) Set(val *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner(val *V2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) *NullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner {
	return &NullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner{value: val, isSet: true}
}

func (v NullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsObjectRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


