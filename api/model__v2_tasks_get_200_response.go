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

// checks if the V2TasksGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TasksGet200Response{}

// V2TasksGet200Response Success
type V2TasksGet200Response struct {
	Data []Task `json:"data"`
}

type _V2TasksGet200Response V2TasksGet200Response

// NewV2TasksGet200Response instantiates a new V2TasksGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TasksGet200Response(data []Task) *V2TasksGet200Response {
	this := V2TasksGet200Response{}
	this.Data = data
	return &this
}

// NewV2TasksGet200ResponseWithDefaults instantiates a new V2TasksGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TasksGet200ResponseWithDefaults() *V2TasksGet200Response {
	this := V2TasksGet200Response{}
	return &this
}

// GetData returns the Data field value
func (o *V2TasksGet200Response) GetData() []Task {
	if o == nil {
		var ret []Task
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2TasksGet200Response) GetDataOk() ([]Task, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *V2TasksGet200Response) SetData(v []Task) {
	o.Data = v
}

func (o V2TasksGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TasksGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2TasksGet200Response) UnmarshalJSON(data []byte) (err error) {
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

	varV2TasksGet200Response := _V2TasksGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2TasksGet200Response)

	if err != nil {
		return err
	}

	*o = V2TasksGet200Response(varV2TasksGet200Response)

	return err
}

type NullableV2TasksGet200Response struct {
	value *V2TasksGet200Response
	isSet bool
}

func (v NullableV2TasksGet200Response) Get() *V2TasksGet200Response {
	return v.value
}

func (v *NullableV2TasksGet200Response) Set(val *V2TasksGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TasksGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TasksGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TasksGet200Response(val *V2TasksGet200Response) *NullableV2TasksGet200Response {
	return &NullableV2TasksGet200Response{value: val, isSet: true}
}

func (v NullableV2TasksGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TasksGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

