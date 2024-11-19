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

// checks if the V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf{}

// V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf struct for V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf
type V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf struct {
	Type string `json:"type"`
	Template V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfTemplate `json:"template"`
}

type _V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf

// NewV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf instantiates a new V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf(type_ string, template V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfTemplate) *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf {
	this := V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf{}
	this.Type = type_
	this.Template = template
	return &this
}

// NewV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfWithDefaults instantiates a new V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfWithDefaults() *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf {
	this := V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf{}
	return &this
}

// GetType returns the Type field value
func (o *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) SetType(v string) {
	o.Type = v
}

// GetTemplate returns the Template field value
func (o *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) GetTemplate() V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfTemplate {
	if o == nil {
		var ret V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) GetTemplateOk() (*V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) SetTemplate(v V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOfTemplate) {
	o.Template = v
}

func (o V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

func (o *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"template",
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

	varV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf := _V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf)

	if err != nil {
		return err
	}

	*o = V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf(varV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf)

	return err
}

type NullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf struct {
	value *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) Get() *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) Set(val *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf(val *V2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) *NullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf {
	return &NullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesPostRequestDataDefaultValueOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


