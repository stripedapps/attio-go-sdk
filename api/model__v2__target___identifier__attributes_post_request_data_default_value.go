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
	"gopkg.in/validator.v2"
	"fmt"
)

// V2TargetIdentifierAttributesPostRequestDataDefaultValue - The default value for this attribute. Static values are used to directly populate values using their contents. Dynamic values are used to lookup data at the point of creation. For example, you could use a dynamic value to insert a value for the currently logged in user. Which default values are available is dependent on the type of the attribute. Default values are not currently supported on people or company objects.
type V2TargetIdentifierAttributesPostRequestDataDefaultValue struct {
	V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf
	V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1 *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1
	Any *interface{}
}

// V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfAsV2TargetIdentifierAttributesPostRequestDataDefaultValue is a convenience function that returns V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf wrapped in V2TargetIdentifierAttributesPostRequestDataDefaultValue
func V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfAsV2TargetIdentifierAttributesPostRequestDataDefaultValue(v *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) V2TargetIdentifierAttributesPostRequestDataDefaultValue {
	return V2TargetIdentifierAttributesPostRequestDataDefaultValue{
		V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf: v,
	}
}

// V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1AsV2TargetIdentifierAttributesPostRequestDataDefaultValue is a convenience function that returns V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1 wrapped in V2TargetIdentifierAttributesPostRequestDataDefaultValue
func V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1AsV2TargetIdentifierAttributesPostRequestDataDefaultValue(v *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1) V2TargetIdentifierAttributesPostRequestDataDefaultValue {
	return V2TargetIdentifierAttributesPostRequestDataDefaultValue{
		V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1: v,
	}
}

// interface{}AsV2TargetIdentifierAttributesPostRequestDataDefaultValue is a convenience function that returns interface{} wrapped in V2TargetIdentifierAttributesPostRequestDataDefaultValue
func AnyAsV2TargetIdentifierAttributesPostRequestDataDefaultValue(v *interface{}) V2TargetIdentifierAttributesPostRequestDataDefaultValue {
	return V2TargetIdentifierAttributesPostRequestDataDefaultValue{
		Any: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *V2TargetIdentifierAttributesPostRequestDataDefaultValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf
	err = newStrictDecoder(data).Decode(&dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf)
	if err == nil {
		jsonV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf, _ := json.Marshal(dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf)
		if string(jsonV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) == "{}" { // empty struct
			dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf = nil
		} else {
			if err = validator.Validate(dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf); err != nil {
				dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf = nil
	}

	// try to unmarshal data into V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1
	err = newStrictDecoder(data).Decode(&dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1)
	if err == nil {
		jsonV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1, _ := json.Marshal(dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1)
		if string(jsonV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1) == "{}" { // empty struct
			dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1 = nil
		} else {
			if err = validator.Validate(dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1); err != nil {
				dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1 = nil
	}

	// try to unmarshal data into Any
	err = newStrictDecoder(data).Decode(&dst.Any)
	if err == nil {
		jsonAny, _ := json.Marshal(dst.Any)
		if string(jsonAny) == "{}" { // empty struct
			dst.Any = nil
		} else {
			if err = validator.Validate(dst.Any); err != nil {
				dst.Any = nil
			} else {
				match++
			}
		}
	} else {
		dst.Any = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf = nil
		dst.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1 = nil
		dst.Any = nil

		return fmt.Errorf("data matches more than one schema in oneOf(V2TargetIdentifierAttributesPostRequestDataDefaultValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(V2TargetIdentifierAttributesPostRequestDataDefaultValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src V2TargetIdentifierAttributesPostRequestDataDefaultValue) MarshalJSON() ([]byte, error) {
	if src.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf != nil {
		return json.Marshal(&src.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf)
	}

	if src.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1 != nil {
		return json.Marshal(&src.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1)
	}

	if src.Any != nil {
		return json.Marshal(&src.Any)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *V2TargetIdentifierAttributesPostRequestDataDefaultValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf != nil {
		return obj.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf
	}

	if obj.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1 != nil {
		return obj.V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf1
	}

	if obj.Any != nil {
		return obj.Any
	}

	// all schemas are nil
	return nil
}

type NullableV2TargetIdentifierAttributesPostRequestDataDefaultValue struct {
	value *V2TargetIdentifierAttributesPostRequestDataDefaultValue
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesPostRequestDataDefaultValue) Get() *V2TargetIdentifierAttributesPostRequestDataDefaultValue {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesPostRequestDataDefaultValue) Set(val *V2TargetIdentifierAttributesPostRequestDataDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesPostRequestDataDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesPostRequestDataDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesPostRequestDataDefaultValue(val *V2TargetIdentifierAttributesPostRequestDataDefaultValue) *NullableV2TargetIdentifierAttributesPostRequestDataDefaultValue {
	return &NullableV2TargetIdentifierAttributesPostRequestDataDefaultValue{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesPostRequestDataDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesPostRequestDataDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


