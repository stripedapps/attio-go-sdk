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


// V2TasksTaskIdPatchRequestDataLinkedRecordsInner struct for V2TasksTaskIdPatchRequestDataLinkedRecordsInner
type V2TasksTaskIdPatchRequestDataLinkedRecordsInner struct {
	V2TasksPostRequestDataLinkedRecordsInnerAnyOf1 *V2TasksPostRequestDataLinkedRecordsInnerAnyOf1
	V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf *V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *V2TasksTaskIdPatchRequestDataLinkedRecordsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into V2TasksPostRequestDataLinkedRecordsInnerAnyOf1
	err = json.Unmarshal(data, &dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1);
	if err == nil {
		jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1, _ := json.Marshal(dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1)
		if string(jsonV2TasksPostRequestDataLinkedRecordsInnerAnyOf1) == "{}" { // empty struct
			dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1 = nil
		} else {
			return nil // data stored in dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1, return on the first match
		}
	} else {
		dst.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1 = nil
	}

	// try to unmarshal JSON data into V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf
	err = json.Unmarshal(data, &dst.V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf);
	if err == nil {
		jsonV2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf, _ := json.Marshal(dst.V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf)
		if string(jsonV2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf) == "{}" { // empty struct
			dst.V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf = nil
		} else {
			return nil // data stored in dst.V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf, return on the first match
		}
	} else {
		dst.V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(V2TasksTaskIdPatchRequestDataLinkedRecordsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *V2TasksTaskIdPatchRequestDataLinkedRecordsInner) MarshalJSON() ([]byte, error) {
	if src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1 != nil {
		return json.Marshal(&src.V2TasksPostRequestDataLinkedRecordsInnerAnyOf1)
	}

	if src.V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf != nil {
		return json.Marshal(&src.V2TasksTaskIdPatchRequestDataLinkedRecordsInnerAnyOf)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner struct {
	value *V2TasksTaskIdPatchRequestDataLinkedRecordsInner
	isSet bool
}

func (v NullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner) Get() *V2TasksTaskIdPatchRequestDataLinkedRecordsInner {
	return v.value
}

func (v *NullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner) Set(val *V2TasksTaskIdPatchRequestDataLinkedRecordsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner(val *V2TasksTaskIdPatchRequestDataLinkedRecordsInner) *NullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner {
	return &NullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner{value: val, isSet: true}
}

func (v NullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TasksTaskIdPatchRequestDataLinkedRecordsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

