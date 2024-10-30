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

// checks if the V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner{}

// V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner struct for V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner
type V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner struct {
	// A raw text field. Values are limited to 10MB.
	Value string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner

// NewV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner instantiates a new V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner(value string) *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner {
	this := V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner{}
	this.Value = value
	return &this
}

// NewV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInnerWithDefaults instantiates a new V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInnerWithDefaults() *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner {
	this := V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner{}
	return &this
}

// GetValue returns the Value field value
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) SetValue(v string) {
	o.Value = v
}

func (o V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"value",
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

	varV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner := _V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner{}

	err = json.Unmarshal(data, &varV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner)

	if err != nil {
		return err
	}

	*o = V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner(varV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner struct {
	value *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) Get() *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) Set(val *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner(val *V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) *NullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner {
	return &NullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


