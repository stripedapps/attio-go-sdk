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

// checks if the V2ListsPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ListsPostRequest{}

// V2ListsPostRequest struct for V2ListsPostRequest
type V2ListsPostRequest struct {
	Data V2ListsPostRequestData `json:"data"`
}

type _V2ListsPostRequest V2ListsPostRequest

// NewV2ListsPostRequest instantiates a new V2ListsPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ListsPostRequest(data V2ListsPostRequestData) *V2ListsPostRequest {
	this := V2ListsPostRequest{}
	this.Data = data
	return &this
}

// NewV2ListsPostRequestWithDefaults instantiates a new V2ListsPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ListsPostRequestWithDefaults() *V2ListsPostRequest {
	this := V2ListsPostRequest{}
	return &this
}

// GetData returns the Data field value
func (o *V2ListsPostRequest) GetData() V2ListsPostRequestData {
	if o == nil {
		var ret V2ListsPostRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2ListsPostRequest) GetDataOk() (*V2ListsPostRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *V2ListsPostRequest) SetData(v V2ListsPostRequestData) {
	o.Data = v
}

func (o V2ListsPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ListsPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2ListsPostRequest) UnmarshalJSON(data []byte) (err error) {
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

	varV2ListsPostRequest := _V2ListsPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ListsPostRequest)

	if err != nil {
		return err
	}

	*o = V2ListsPostRequest(varV2ListsPostRequest)

	return err
}

type NullableV2ListsPostRequest struct {
	value *V2ListsPostRequest
	isSet bool
}

func (v NullableV2ListsPostRequest) Get() *V2ListsPostRequest {
	return v.value
}

func (v *NullableV2ListsPostRequest) Set(val *V2ListsPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ListsPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ListsPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ListsPostRequest(val *V2ListsPostRequest) *NullableV2ListsPostRequest {
	return &NullableV2ListsPostRequest{value: val, isSet: true}
}

func (v NullableV2ListsPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ListsPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


