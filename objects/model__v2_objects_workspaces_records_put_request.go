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

// checks if the V2ObjectsWorkspacesRecordsPutRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsWorkspacesRecordsPutRequest{}

// V2ObjectsWorkspacesRecordsPutRequest struct for V2ObjectsWorkspacesRecordsPutRequest
type V2ObjectsWorkspacesRecordsPutRequest struct {
	Data V2ObjectsWorkspacesRecordsPutRequestData `json:"data"`
}

type _V2ObjectsWorkspacesRecordsPutRequest V2ObjectsWorkspacesRecordsPutRequest

// NewV2ObjectsWorkspacesRecordsPutRequest instantiates a new V2ObjectsWorkspacesRecordsPutRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsWorkspacesRecordsPutRequest(data V2ObjectsWorkspacesRecordsPutRequestData) *V2ObjectsWorkspacesRecordsPutRequest {
	this := V2ObjectsWorkspacesRecordsPutRequest{}
	this.Data = data
	return &this
}

// NewV2ObjectsWorkspacesRecordsPutRequestWithDefaults instantiates a new V2ObjectsWorkspacesRecordsPutRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsWorkspacesRecordsPutRequestWithDefaults() *V2ObjectsWorkspacesRecordsPutRequest {
	this := V2ObjectsWorkspacesRecordsPutRequest{}
	return &this
}

// GetData returns the Data field value
func (o *V2ObjectsWorkspacesRecordsPutRequest) GetData() V2ObjectsWorkspacesRecordsPutRequestData {
	if o == nil {
		var ret V2ObjectsWorkspacesRecordsPutRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequest) GetDataOk() (*V2ObjectsWorkspacesRecordsPutRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *V2ObjectsWorkspacesRecordsPutRequest) SetData(v V2ObjectsWorkspacesRecordsPutRequestData) {
	o.Data = v
}

func (o V2ObjectsWorkspacesRecordsPutRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsWorkspacesRecordsPutRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2ObjectsWorkspacesRecordsPutRequest) UnmarshalJSON(data []byte) (err error) {
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

	varV2ObjectsWorkspacesRecordsPutRequest := _V2ObjectsWorkspacesRecordsPutRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ObjectsWorkspacesRecordsPutRequest)

	if err != nil {
		return err
	}

	*o = V2ObjectsWorkspacesRecordsPutRequest(varV2ObjectsWorkspacesRecordsPutRequest)

	return err
}

type NullableV2ObjectsWorkspacesRecordsPutRequest struct {
	value *V2ObjectsWorkspacesRecordsPutRequest
	isSet bool
}

func (v NullableV2ObjectsWorkspacesRecordsPutRequest) Get() *V2ObjectsWorkspacesRecordsPutRequest {
	return v.value
}

func (v *NullableV2ObjectsWorkspacesRecordsPutRequest) Set(val *V2ObjectsWorkspacesRecordsPutRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsWorkspacesRecordsPutRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsWorkspacesRecordsPutRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsWorkspacesRecordsPutRequest(val *V2ObjectsWorkspacesRecordsPutRequest) *NullableV2ObjectsWorkspacesRecordsPutRequest {
	return &NullableV2ObjectsWorkspacesRecordsPutRequest{value: val, isSet: true}
}

func (v NullableV2ObjectsWorkspacesRecordsPutRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsWorkspacesRecordsPutRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


