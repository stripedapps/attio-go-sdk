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

// checks if the V2ObjectsObjectRecordsPutRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsObjectRecordsPutRequestData{}

// V2ObjectsObjectRecordsPutRequestData struct for V2ObjectsObjectRecordsPutRequestData
type V2ObjectsObjectRecordsPutRequestData struct {
	// An object with an attribute `api_slug` or `attribute_id` as the key, and a single value (for single-select attributes), or an array of values (for single or multi-select attributes) as the values. For complete documentation on values for all attribute types, please see our [attribute type docs](/docs/attribute-types).
	Values map[string][]interface{} `json:"values"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsObjectRecordsPutRequestData V2ObjectsObjectRecordsPutRequestData

// NewV2ObjectsObjectRecordsPutRequestData instantiates a new V2ObjectsObjectRecordsPutRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsObjectRecordsPutRequestData(values map[string][]interface{}) *V2ObjectsObjectRecordsPutRequestData {
	this := V2ObjectsObjectRecordsPutRequestData{}
	this.Values = values
	return &this
}

// NewV2ObjectsObjectRecordsPutRequestDataWithDefaults instantiates a new V2ObjectsObjectRecordsPutRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsObjectRecordsPutRequestDataWithDefaults() *V2ObjectsObjectRecordsPutRequestData {
	this := V2ObjectsObjectRecordsPutRequestData{}
	return &this
}

// GetValues returns the Values field value
func (o *V2ObjectsObjectRecordsPutRequestData) GetValues() map[string][]interface{} {
	if o == nil {
		var ret map[string][]interface{}
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsPutRequestData) GetValuesOk() (map[string][]interface{}, bool) {
	if o == nil {
		return map[string][]interface{}{}, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *V2ObjectsObjectRecordsPutRequestData) SetValues(v map[string][]interface{}) {
	o.Values = v
}

func (o V2ObjectsObjectRecordsPutRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsObjectRecordsPutRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["values"] = o.Values

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsObjectRecordsPutRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"values",
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

	varV2ObjectsObjectRecordsPutRequestData := _V2ObjectsObjectRecordsPutRequestData{}

	err = json.Unmarshal(data, &varV2ObjectsObjectRecordsPutRequestData)

	if err != nil {
		return err
	}

	*o = V2ObjectsObjectRecordsPutRequestData(varV2ObjectsObjectRecordsPutRequestData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsObjectRecordsPutRequestData struct {
	value *V2ObjectsObjectRecordsPutRequestData
	isSet bool
}

func (v NullableV2ObjectsObjectRecordsPutRequestData) Get() *V2ObjectsObjectRecordsPutRequestData {
	return v.value
}

func (v *NullableV2ObjectsObjectRecordsPutRequestData) Set(val *V2ObjectsObjectRecordsPutRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsObjectRecordsPutRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsObjectRecordsPutRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsObjectRecordsPutRequestData(val *V2ObjectsObjectRecordsPutRequestData) *NullableV2ObjectsObjectRecordsPutRequestData {
	return &NullableV2ObjectsObjectRecordsPutRequestData{value: val, isSet: true}
}

func (v NullableV2ObjectsObjectRecordsPutRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsObjectRecordsPutRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

