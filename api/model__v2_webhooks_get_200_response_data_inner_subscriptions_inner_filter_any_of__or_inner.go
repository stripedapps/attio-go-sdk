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


// V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner struct for V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner
type V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner struct {
	V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf
	V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1 *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf
	err = json.Unmarshal(data, &dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf);
	if err == nil {
		jsonV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf, _ := json.Marshal(dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf)
		if string(jsonV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) == "{}" { // empty struct
			dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf = nil
		} else {
			return nil // data stored in dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf, return on the first match
		}
	} else {
		dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf = nil
	}

	// try to unmarshal JSON data into V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1
	err = json.Unmarshal(data, &dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1);
	if err == nil {
		jsonV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1, _ := json.Marshal(dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1)
		if string(jsonV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1) == "{}" { // empty struct
			dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1 = nil
		} else {
			return nil // data stored in dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1, return on the first match
		}
	} else {
		dst.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) MarshalJSON() ([]byte, error) {
	if src.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf != nil {
		return json.Marshal(&src.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf)
	}

	if src.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1 != nil {
		return json.Marshal(&src.V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf1)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner struct {
	value *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner
	isSet bool
}

func (v NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) Get() *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner {
	return v.value
}

func (v *NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) Set(val *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner(val *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) *NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner {
	return &NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner{value: val, isSet: true}
}

func (v NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


