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

// checks if the V2TargetIdentifierAttributesAttributeOptionsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TargetIdentifierAttributesAttributeOptionsPostRequest{}

// V2TargetIdentifierAttributesAttributeOptionsPostRequest struct for V2TargetIdentifierAttributesAttributeOptionsPostRequest
type V2TargetIdentifierAttributesAttributeOptionsPostRequest struct {
	Data V2TargetIdentifierAttributesAttributeOptionsPostRequestData `json:"data"`
}

type _V2TargetIdentifierAttributesAttributeOptionsPostRequest V2TargetIdentifierAttributesAttributeOptionsPostRequest

// NewV2TargetIdentifierAttributesAttributeOptionsPostRequest instantiates a new V2TargetIdentifierAttributesAttributeOptionsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TargetIdentifierAttributesAttributeOptionsPostRequest(data V2TargetIdentifierAttributesAttributeOptionsPostRequestData) *V2TargetIdentifierAttributesAttributeOptionsPostRequest {
	this := V2TargetIdentifierAttributesAttributeOptionsPostRequest{}
	this.Data = data
	return &this
}

// NewV2TargetIdentifierAttributesAttributeOptionsPostRequestWithDefaults instantiates a new V2TargetIdentifierAttributesAttributeOptionsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TargetIdentifierAttributesAttributeOptionsPostRequestWithDefaults() *V2TargetIdentifierAttributesAttributeOptionsPostRequest {
	this := V2TargetIdentifierAttributesAttributeOptionsPostRequest{}
	return &this
}

// GetData returns the Data field value
func (o *V2TargetIdentifierAttributesAttributeOptionsPostRequest) GetData() V2TargetIdentifierAttributesAttributeOptionsPostRequestData {
	if o == nil {
		var ret V2TargetIdentifierAttributesAttributeOptionsPostRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeOptionsPostRequest) GetDataOk() (*V2TargetIdentifierAttributesAttributeOptionsPostRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *V2TargetIdentifierAttributesAttributeOptionsPostRequest) SetData(v V2TargetIdentifierAttributesAttributeOptionsPostRequestData) {
	o.Data = v
}

func (o V2TargetIdentifierAttributesAttributeOptionsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TargetIdentifierAttributesAttributeOptionsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2TargetIdentifierAttributesAttributeOptionsPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varV2TargetIdentifierAttributesAttributeOptionsPostRequest := _V2TargetIdentifierAttributesAttributeOptionsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2TargetIdentifierAttributesAttributeOptionsPostRequest)

	if err != nil {
		return err
	}

	*o = V2TargetIdentifierAttributesAttributeOptionsPostRequest(varV2TargetIdentifierAttributesAttributeOptionsPostRequest)

	return err
}

type NullableV2TargetIdentifierAttributesAttributeOptionsPostRequest struct {
	value *V2TargetIdentifierAttributesAttributeOptionsPostRequest
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesAttributeOptionsPostRequest) Get() *V2TargetIdentifierAttributesAttributeOptionsPostRequest {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesAttributeOptionsPostRequest) Set(val *V2TargetIdentifierAttributesAttributeOptionsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesAttributeOptionsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesAttributeOptionsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesAttributeOptionsPostRequest(val *V2TargetIdentifierAttributesAttributeOptionsPostRequest) *NullableV2TargetIdentifierAttributesAttributeOptionsPostRequest {
	return &NullableV2TargetIdentifierAttributesAttributeOptionsPostRequest{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesAttributeOptionsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesAttributeOptionsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

