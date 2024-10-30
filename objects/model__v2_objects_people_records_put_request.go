/*
Attio Standard Objects

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the V2ObjectsPeopleRecordsPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsPutRequest{}

// V2ObjectsPeopleRecordsPutRequest struct for V2ObjectsPeopleRecordsPutRequest
type V2ObjectsPeopleRecordsPutRequest struct {
	Data V2ObjectsPeopleRecordsPutRequestData `json:"data"`
}

type _V2ObjectsPeopleRecordsPutRequest V2ObjectsPeopleRecordsPutRequest

// NewV2ObjectsPeopleRecordsPutRequest instantiates a new V2ObjectsPeopleRecordsPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsPutRequest(data V2ObjectsPeopleRecordsPutRequestData) *V2ObjectsPeopleRecordsPutRequest {
	this := V2ObjectsPeopleRecordsPutRequest{}
	this.Data = data
	return &this
}

// NewV2ObjectsPeopleRecordsPutRequestWithDefaults instantiates a new V2ObjectsPeopleRecordsPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsPutRequestWithDefaults() *V2ObjectsPeopleRecordsPutRequest {
	this := V2ObjectsPeopleRecordsPutRequest{}
	return &this
}

// GetData returns the Data field value
func (o *V2ObjectsPeopleRecordsPutRequest) GetData() V2ObjectsPeopleRecordsPutRequestData {
	if o == nil {
		var ret V2ObjectsPeopleRecordsPutRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPutRequest) GetDataOk() (*V2ObjectsPeopleRecordsPutRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *V2ObjectsPeopleRecordsPutRequest) SetData(v V2ObjectsPeopleRecordsPutRequestData) {
	o.Data = v
}

func (o V2ObjectsPeopleRecordsPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2ObjectsPeopleRecordsPutRequest) UnmarshalJSON(data []byte) (err error) {
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

	varV2ObjectsPeopleRecordsPutRequest := _V2ObjectsPeopleRecordsPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ObjectsPeopleRecordsPutRequest)

	if err != nil {
		return err
	}

	*o = V2ObjectsPeopleRecordsPutRequest(varV2ObjectsPeopleRecordsPutRequest)

	return err
}

type NullableV2ObjectsPeopleRecordsPutRequest struct {
	value *V2ObjectsPeopleRecordsPutRequest
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsPutRequest) Get() *V2ObjectsPeopleRecordsPutRequest {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsPutRequest) Set(val *V2ObjectsPeopleRecordsPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsPutRequest(val *V2ObjectsPeopleRecordsPutRequest) *NullableV2ObjectsPeopleRecordsPutRequest {
	return &NullableV2ObjectsPeopleRecordsPutRequest{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


