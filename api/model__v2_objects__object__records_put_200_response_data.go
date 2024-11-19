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

// checks if the V2ObjectsObjectRecordsPut200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsObjectRecordsPut200ResponseData{}

// V2ObjectsObjectRecordsPut200ResponseData struct for V2ObjectsObjectRecordsPut200ResponseData
type V2ObjectsObjectRecordsPut200ResponseData struct {
	Id V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId `json:"id"`
	// When this record was created.
	CreatedAt string `json:"created_at"`
	// A record type with an attribute `api_slug` as the key, and an array of value objects as the values.
	Values map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner `json:"values"`
}

type _V2ObjectsObjectRecordsPut200ResponseData V2ObjectsObjectRecordsPut200ResponseData

// NewV2ObjectsObjectRecordsPut200ResponseData instantiates a new V2ObjectsObjectRecordsPut200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsObjectRecordsPut200ResponseData(id V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, createdAt string, values map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner) *V2ObjectsObjectRecordsPut200ResponseData {
	this := V2ObjectsObjectRecordsPut200ResponseData{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Values = values
	return &this
}

// NewV2ObjectsObjectRecordsPut200ResponseDataWithDefaults instantiates a new V2ObjectsObjectRecordsPut200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsObjectRecordsPut200ResponseDataWithDefaults() *V2ObjectsObjectRecordsPut200ResponseData {
	this := V2ObjectsObjectRecordsPut200ResponseData{}
	return &this
}

// GetId returns the Id field value
func (o *V2ObjectsObjectRecordsPut200ResponseData) GetId() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId {
	if o == nil {
		var ret V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsPut200ResponseData) GetIdOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V2ObjectsObjectRecordsPut200ResponseData) SetId(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerId) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *V2ObjectsObjectRecordsPut200ResponseData) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsPut200ResponseData) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *V2ObjectsObjectRecordsPut200ResponseData) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetValues returns the Values field value
func (o *V2ObjectsObjectRecordsPut200ResponseData) GetValues() map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner {
	if o == nil {
		var ret map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsPut200ResponseData) GetValuesOk() (map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner, bool) {
	if o == nil {
		return map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner{}, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *V2ObjectsObjectRecordsPut200ResponseData) SetValues(v map[string][]V2ObjectsObjectRecordsPut200ResponseDataValuesValueInner) {
	o.Values = v
}

func (o V2ObjectsObjectRecordsPut200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsObjectRecordsPut200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *V2ObjectsObjectRecordsPut200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"created_at",
		"values",
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

	varV2ObjectsObjectRecordsPut200ResponseData := _V2ObjectsObjectRecordsPut200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ObjectsObjectRecordsPut200ResponseData)

	if err != nil {
		return err
	}

	*o = V2ObjectsObjectRecordsPut200ResponseData(varV2ObjectsObjectRecordsPut200ResponseData)

	return err
}

type NullableV2ObjectsObjectRecordsPut200ResponseData struct {
	value *V2ObjectsObjectRecordsPut200ResponseData
	isSet bool
}

func (v NullableV2ObjectsObjectRecordsPut200ResponseData) Get() *V2ObjectsObjectRecordsPut200ResponseData {
	return v.value
}

func (v *NullableV2ObjectsObjectRecordsPut200ResponseData) Set(val *V2ObjectsObjectRecordsPut200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsObjectRecordsPut200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsObjectRecordsPut200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsObjectRecordsPut200ResponseData(val *V2ObjectsObjectRecordsPut200ResponseData) *NullableV2ObjectsObjectRecordsPut200ResponseData {
	return &NullableV2ObjectsObjectRecordsPut200ResponseData{value: val, isSet: true}
}

func (v NullableV2ObjectsObjectRecordsPut200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsObjectRecordsPut200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


