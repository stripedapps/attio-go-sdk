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

// checks if the V2TargetIdentifierAttributesAttributeOptionsPostRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TargetIdentifierAttributesAttributeOptionsPostRequestData{}

// V2TargetIdentifierAttributesAttributeOptionsPostRequestData struct for V2TargetIdentifierAttributesAttributeOptionsPostRequestData
type V2TargetIdentifierAttributesAttributeOptionsPostRequestData struct {
	// The Title of the select option
	Title string `json:"title"`
}

type _V2TargetIdentifierAttributesAttributeOptionsPostRequestData V2TargetIdentifierAttributesAttributeOptionsPostRequestData

// NewV2TargetIdentifierAttributesAttributeOptionsPostRequestData instantiates a new V2TargetIdentifierAttributesAttributeOptionsPostRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TargetIdentifierAttributesAttributeOptionsPostRequestData(title string) *V2TargetIdentifierAttributesAttributeOptionsPostRequestData {
	this := V2TargetIdentifierAttributesAttributeOptionsPostRequestData{}
	this.Title = title
	return &this
}

// NewV2TargetIdentifierAttributesAttributeOptionsPostRequestDataWithDefaults instantiates a new V2TargetIdentifierAttributesAttributeOptionsPostRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TargetIdentifierAttributesAttributeOptionsPostRequestDataWithDefaults() *V2TargetIdentifierAttributesAttributeOptionsPostRequestData {
	this := V2TargetIdentifierAttributesAttributeOptionsPostRequestData{}
	return &this
}

// GetTitle returns the Title field value
func (o *V2TargetIdentifierAttributesAttributeOptionsPostRequestData) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeOptionsPostRequestData) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *V2TargetIdentifierAttributesAttributeOptionsPostRequestData) SetTitle(v string) {
	o.Title = v
}

func (o V2TargetIdentifierAttributesAttributeOptionsPostRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TargetIdentifierAttributesAttributeOptionsPostRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	return toSerialize, nil
}

func (o *V2TargetIdentifierAttributesAttributeOptionsPostRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
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

	varV2TargetIdentifierAttributesAttributeOptionsPostRequestData := _V2TargetIdentifierAttributesAttributeOptionsPostRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2TargetIdentifierAttributesAttributeOptionsPostRequestData)

	if err != nil {
		return err
	}

	*o = V2TargetIdentifierAttributesAttributeOptionsPostRequestData(varV2TargetIdentifierAttributesAttributeOptionsPostRequestData)

	return err
}

type NullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData struct {
	value *V2TargetIdentifierAttributesAttributeOptionsPostRequestData
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData) Get() *V2TargetIdentifierAttributesAttributeOptionsPostRequestData {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData) Set(val *V2TargetIdentifierAttributesAttributeOptionsPostRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData(val *V2TargetIdentifierAttributesAttributeOptionsPostRequestData) *NullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData {
	return &NullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesAttributeOptionsPostRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

