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
	"bytes"
	"fmt"
)

// checks if the OutputValueAnyOf18 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputValueAnyOf18{}

// OutputValueAnyOf18 struct for OutputValueAnyOf18
type OutputValueAnyOf18 struct {
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	// A timestamp value represents a single, universal moment in time using an ISO 8601 formatted string. This means that a timestamp consists of a date, a time (with nanosecond precision), and a time zone. Attio will coerce timestamps which do not provide full nanosecond precision and UTC is assumed if no time zone is provided. For example, \"2023\", \"2023-01\", \"2023-01-02\", \"2023-01-02T13:00\", \"2023-01-02T13:00:00\", and \"2023-01-02T13:00:00.000000000\" will all be coerced to \"2023-01-02T13:00:00.000000000Z\". Timestamps are always returned in UTC. For example, writing a timestamp value using the string \"2023-01-02T13:00:00.000000000+02:00\" will result in the value \"2023-01-02T11:00:00.000000000Z\" being returned. The maximum date is \"9999-12-31T23:59:59.999999999Z\".
	Value string `json:"value"`
}

type _OutputValueAnyOf18 OutputValueAnyOf18

// NewOutputValueAnyOf18 instantiates a new OutputValueAnyOf18 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputValueAnyOf18(attributeType string, value string) *OutputValueAnyOf18 {
	this := OutputValueAnyOf18{}
	this.AttributeType = attributeType
	this.Value = value
	return &this
}

// NewOutputValueAnyOf18WithDefaults instantiates a new OutputValueAnyOf18 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputValueAnyOf18WithDefaults() *OutputValueAnyOf18 {
	this := OutputValueAnyOf18{}
	return &this
}

// GetAttributeType returns the AttributeType field value
func (o *OutputValueAnyOf18) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *OutputValueAnyOf18) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *OutputValueAnyOf18) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetValue returns the Value field value
func (o *OutputValueAnyOf18) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OutputValueAnyOf18) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OutputValueAnyOf18) SetValue(v string) {
	o.Value = v
}

func (o OutputValueAnyOf18) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputValueAnyOf18) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attribute_type"] = o.AttributeType
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *OutputValueAnyOf18) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"attribute_type",
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

	varOutputValueAnyOf18 := _OutputValueAnyOf18{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputValueAnyOf18)

	if err != nil {
		return err
	}

	*o = OutputValueAnyOf18(varOutputValueAnyOf18)

	return err
}

type NullableOutputValueAnyOf18 struct {
	value *OutputValueAnyOf18
	isSet bool
}

func (v NullableOutputValueAnyOf18) Get() *OutputValueAnyOf18 {
	return v.value
}

func (v *NullableOutputValueAnyOf18) Set(val *OutputValueAnyOf18) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputValueAnyOf18) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputValueAnyOf18) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputValueAnyOf18(val *OutputValueAnyOf18) *NullableOutputValueAnyOf18 {
	return &NullableOutputValueAnyOf18{value: val, isSet: true}
}

func (v NullableOutputValueAnyOf18) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputValueAnyOf18) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


