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

// checks if the OutputValueAnyOf13 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutputValueAnyOf13{}

// OutputValueAnyOf13 struct for OutputValueAnyOf13
type OutputValueAnyOf13 struct {
	Status Status `json:"status"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
}

type _OutputValueAnyOf13 OutputValueAnyOf13

// NewOutputValueAnyOf13 instantiates a new OutputValueAnyOf13 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutputValueAnyOf13(status Status, attributeType string) *OutputValueAnyOf13 {
	this := OutputValueAnyOf13{}
	this.Status = status
	this.AttributeType = attributeType
	return &this
}

// NewOutputValueAnyOf13WithDefaults instantiates a new OutputValueAnyOf13 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutputValueAnyOf13WithDefaults() *OutputValueAnyOf13 {
	this := OutputValueAnyOf13{}
	return &this
}

// GetStatus returns the Status field value
func (o *OutputValueAnyOf13) GetStatus() Status {
	if o == nil {
		var ret Status
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *OutputValueAnyOf13) GetStatusOk() (*Status, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *OutputValueAnyOf13) SetStatus(v Status) {
	o.Status = v
}

// GetAttributeType returns the AttributeType field value
func (o *OutputValueAnyOf13) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *OutputValueAnyOf13) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *OutputValueAnyOf13) SetAttributeType(v string) {
	o.AttributeType = v
}

func (o OutputValueAnyOf13) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutputValueAnyOf13) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	toSerialize["attribute_type"] = o.AttributeType
	return toSerialize, nil
}

func (o *OutputValueAnyOf13) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
		"attribute_type",
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

	varOutputValueAnyOf13 := _OutputValueAnyOf13{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOutputValueAnyOf13)

	if err != nil {
		return err
	}

	*o = OutputValueAnyOf13(varOutputValueAnyOf13)

	return err
}

type NullableOutputValueAnyOf13 struct {
	value *OutputValueAnyOf13
	isSet bool
}

func (v NullableOutputValueAnyOf13) Get() *OutputValueAnyOf13 {
	return v.value
}

func (v *NullableOutputValueAnyOf13) Set(val *OutputValueAnyOf13) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputValueAnyOf13) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputValueAnyOf13) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputValueAnyOf13(val *OutputValueAnyOf13) *NullableOutputValueAnyOf13 {
	return &NullableOutputValueAnyOf13{value: val, isSet: true}
}

func (v NullableOutputValueAnyOf13) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputValueAnyOf13) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

