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

// checks if the V2TargetIdentifierAttributesGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TargetIdentifierAttributesGet200Response{}

// V2TargetIdentifierAttributesGet200Response Success
type V2TargetIdentifierAttributesGet200Response struct {
	Data []Attribute `json:"data"`
}

type _V2TargetIdentifierAttributesGet200Response V2TargetIdentifierAttributesGet200Response

// NewV2TargetIdentifierAttributesGet200Response instantiates a new V2TargetIdentifierAttributesGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TargetIdentifierAttributesGet200Response(data []Attribute) *V2TargetIdentifierAttributesGet200Response {
	this := V2TargetIdentifierAttributesGet200Response{}
	this.Data = data
	return &this
}

// NewV2TargetIdentifierAttributesGet200ResponseWithDefaults instantiates a new V2TargetIdentifierAttributesGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TargetIdentifierAttributesGet200ResponseWithDefaults() *V2TargetIdentifierAttributesGet200Response {
	this := V2TargetIdentifierAttributesGet200Response{}
	return &this
}

// GetData returns the Data field value
func (o *V2TargetIdentifierAttributesGet200Response) GetData() []Attribute {
	if o == nil {
		var ret []Attribute
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesGet200Response) GetDataOk() ([]Attribute, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *V2TargetIdentifierAttributesGet200Response) SetData(v []Attribute) {
	o.Data = v
}

func (o V2TargetIdentifierAttributesGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TargetIdentifierAttributesGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2TargetIdentifierAttributesGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varV2TargetIdentifierAttributesGet200Response := _V2TargetIdentifierAttributesGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2TargetIdentifierAttributesGet200Response)

	if err != nil {
		return err
	}

	*o = V2TargetIdentifierAttributesGet200Response(varV2TargetIdentifierAttributesGet200Response)

	return err
}

type NullableV2TargetIdentifierAttributesGet200Response struct {
	value *V2TargetIdentifierAttributesGet200Response
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesGet200Response) Get() *V2TargetIdentifierAttributesGet200Response {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesGet200Response) Set(val *V2TargetIdentifierAttributesGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesGet200Response(val *V2TargetIdentifierAttributesGet200Response) *NullableV2TargetIdentifierAttributesGet200Response {
	return &NullableV2TargetIdentifierAttributesGet200Response{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

