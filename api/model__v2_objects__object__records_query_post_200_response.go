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

// checks if the V2ObjectsObjectRecordsQueryPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsObjectRecordsQueryPost200Response{}

// V2ObjectsObjectRecordsQueryPost200Response Success
type V2ObjectsObjectRecordsQueryPost200Response struct {
	Data []V2ObjectsObjectRecordsQueryPost200ResponseDataInner `json:"data"`
}

type _V2ObjectsObjectRecordsQueryPost200Response V2ObjectsObjectRecordsQueryPost200Response

// NewV2ObjectsObjectRecordsQueryPost200Response instantiates a new V2ObjectsObjectRecordsQueryPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsObjectRecordsQueryPost200Response(data []V2ObjectsObjectRecordsQueryPost200ResponseDataInner) *V2ObjectsObjectRecordsQueryPost200Response {
	this := V2ObjectsObjectRecordsQueryPost200Response{}
	this.Data = data
	return &this
}

// NewV2ObjectsObjectRecordsQueryPost200ResponseWithDefaults instantiates a new V2ObjectsObjectRecordsQueryPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsObjectRecordsQueryPost200ResponseWithDefaults() *V2ObjectsObjectRecordsQueryPost200Response {
	this := V2ObjectsObjectRecordsQueryPost200Response{}
	return &this
}

// GetData returns the Data field value
func (o *V2ObjectsObjectRecordsQueryPost200Response) GetData() []V2ObjectsObjectRecordsQueryPost200ResponseDataInner {
	if o == nil {
		var ret []V2ObjectsObjectRecordsQueryPost200ResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200Response) GetDataOk() ([]V2ObjectsObjectRecordsQueryPost200ResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *V2ObjectsObjectRecordsQueryPost200Response) SetData(v []V2ObjectsObjectRecordsQueryPost200ResponseDataInner) {
	o.Data = v
}

func (o V2ObjectsObjectRecordsQueryPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsObjectRecordsQueryPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2ObjectsObjectRecordsQueryPost200Response) UnmarshalJSON(data []byte) (err error) {
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

	varV2ObjectsObjectRecordsQueryPost200Response := _V2ObjectsObjectRecordsQueryPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ObjectsObjectRecordsQueryPost200Response)

	if err != nil {
		return err
	}

	*o = V2ObjectsObjectRecordsQueryPost200Response(varV2ObjectsObjectRecordsQueryPost200Response)

	return err
}

type NullableV2ObjectsObjectRecordsQueryPost200Response struct {
	value *V2ObjectsObjectRecordsQueryPost200Response
	isSet bool
}

func (v NullableV2ObjectsObjectRecordsQueryPost200Response) Get() *V2ObjectsObjectRecordsQueryPost200Response {
	return v.value
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200Response) Set(val *V2ObjectsObjectRecordsQueryPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsObjectRecordsQueryPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsObjectRecordsQueryPost200Response(val *V2ObjectsObjectRecordsQueryPost200Response) *NullableV2ObjectsObjectRecordsQueryPost200Response {
	return &NullableV2ObjectsObjectRecordsQueryPost200Response{value: val, isSet: true}
}

func (v NullableV2ObjectsObjectRecordsQueryPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


