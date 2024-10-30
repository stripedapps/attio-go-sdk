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

// checks if the V2ObjectsUsersRecordsQueryPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsUsersRecordsQueryPost200Response{}

// V2ObjectsUsersRecordsQueryPost200Response Success
type V2ObjectsUsersRecordsQueryPost200Response struct {
	Data []V2ObjectsUsersRecordsQueryPost200ResponseDataInner `json:"data"`
}

type _V2ObjectsUsersRecordsQueryPost200Response V2ObjectsUsersRecordsQueryPost200Response

// NewV2ObjectsUsersRecordsQueryPost200Response instantiates a new V2ObjectsUsersRecordsQueryPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsUsersRecordsQueryPost200Response(data []V2ObjectsUsersRecordsQueryPost200ResponseDataInner) *V2ObjectsUsersRecordsQueryPost200Response {
	this := V2ObjectsUsersRecordsQueryPost200Response{}
	this.Data = data
	return &this
}

// NewV2ObjectsUsersRecordsQueryPost200ResponseWithDefaults instantiates a new V2ObjectsUsersRecordsQueryPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsUsersRecordsQueryPost200ResponseWithDefaults() *V2ObjectsUsersRecordsQueryPost200Response {
	this := V2ObjectsUsersRecordsQueryPost200Response{}
	return &this
}

// GetData returns the Data field value
func (o *V2ObjectsUsersRecordsQueryPost200Response) GetData() []V2ObjectsUsersRecordsQueryPost200ResponseDataInner {
	if o == nil {
		var ret []V2ObjectsUsersRecordsQueryPost200ResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsUsersRecordsQueryPost200Response) GetDataOk() ([]V2ObjectsUsersRecordsQueryPost200ResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *V2ObjectsUsersRecordsQueryPost200Response) SetData(v []V2ObjectsUsersRecordsQueryPost200ResponseDataInner) {
	o.Data = v
}

func (o V2ObjectsUsersRecordsQueryPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsUsersRecordsQueryPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2ObjectsUsersRecordsQueryPost200Response) UnmarshalJSON(data []byte) (err error) {
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

	varV2ObjectsUsersRecordsQueryPost200Response := _V2ObjectsUsersRecordsQueryPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ObjectsUsersRecordsQueryPost200Response)

	if err != nil {
		return err
	}

	*o = V2ObjectsUsersRecordsQueryPost200Response(varV2ObjectsUsersRecordsQueryPost200Response)

	return err
}

type NullableV2ObjectsUsersRecordsQueryPost200Response struct {
	value *V2ObjectsUsersRecordsQueryPost200Response
	isSet bool
}

func (v NullableV2ObjectsUsersRecordsQueryPost200Response) Get() *V2ObjectsUsersRecordsQueryPost200Response {
	return v.value
}

func (v *NullableV2ObjectsUsersRecordsQueryPost200Response) Set(val *V2ObjectsUsersRecordsQueryPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsUsersRecordsQueryPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsUsersRecordsQueryPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsUsersRecordsQueryPost200Response(val *V2ObjectsUsersRecordsQueryPost200Response) *NullableV2ObjectsUsersRecordsQueryPost200Response {
	return &NullableV2ObjectsUsersRecordsQueryPost200Response{value: val, isSet: true}
}

func (v NullableV2ObjectsUsersRecordsQueryPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsUsersRecordsQueryPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


