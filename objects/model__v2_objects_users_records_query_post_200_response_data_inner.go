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

// checks if the V2ObjectsUsersRecordsQueryPost200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsUsersRecordsQueryPost200ResponseDataInner{}

// V2ObjectsUsersRecordsQueryPost200ResponseDataInner struct for V2ObjectsUsersRecordsQueryPost200ResponseDataInner
type V2ObjectsUsersRecordsQueryPost200ResponseDataInner struct {
	Id V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerId `json:"id"`
	// When this record was created.
	CreatedAt string `json:"created_at"`
	Values V2ObjectsUsersRecordsQueryPost200ResponseDataInnerValues `json:"values"`
}

type _V2ObjectsUsersRecordsQueryPost200ResponseDataInner V2ObjectsUsersRecordsQueryPost200ResponseDataInner

// NewV2ObjectsUsersRecordsQueryPost200ResponseDataInner instantiates a new V2ObjectsUsersRecordsQueryPost200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsUsersRecordsQueryPost200ResponseDataInner(id V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerId, createdAt string, values V2ObjectsUsersRecordsQueryPost200ResponseDataInnerValues) *V2ObjectsUsersRecordsQueryPost200ResponseDataInner {
	this := V2ObjectsUsersRecordsQueryPost200ResponseDataInner{}
	this.Id = id
	this.CreatedAt = createdAt
	this.Values = values
	return &this
}

// NewV2ObjectsUsersRecordsQueryPost200ResponseDataInnerWithDefaults instantiates a new V2ObjectsUsersRecordsQueryPost200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsUsersRecordsQueryPost200ResponseDataInnerWithDefaults() *V2ObjectsUsersRecordsQueryPost200ResponseDataInner {
	this := V2ObjectsUsersRecordsQueryPost200ResponseDataInner{}
	return &this
}

// GetId returns the Id field value
func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) GetId() V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerId {
	if o == nil {
		var ret V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) GetIdOk() (*V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) SetId(v V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerId) {
	o.Id = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetValues returns the Values field value
func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) GetValues() V2ObjectsUsersRecordsQueryPost200ResponseDataInnerValues {
	if o == nil {
		var ret V2ObjectsUsersRecordsQueryPost200ResponseDataInnerValues
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) GetValuesOk() (*V2ObjectsUsersRecordsQueryPost200ResponseDataInnerValues, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) SetValues(v V2ObjectsUsersRecordsQueryPost200ResponseDataInnerValues) {
	o.Values = v
}

func (o V2ObjectsUsersRecordsQueryPost200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsUsersRecordsQueryPost200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["values"] = o.Values
	return toSerialize, nil
}

func (o *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
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

	varV2ObjectsUsersRecordsQueryPost200ResponseDataInner := _V2ObjectsUsersRecordsQueryPost200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ObjectsUsersRecordsQueryPost200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = V2ObjectsUsersRecordsQueryPost200ResponseDataInner(varV2ObjectsUsersRecordsQueryPost200ResponseDataInner)

	return err
}

type NullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner struct {
	value *V2ObjectsUsersRecordsQueryPost200ResponseDataInner
	isSet bool
}

func (v NullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner) Get() *V2ObjectsUsersRecordsQueryPost200ResponseDataInner {
	return v.value
}

func (v *NullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner) Set(val *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner(val *V2ObjectsUsersRecordsQueryPost200ResponseDataInner) *NullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner {
	return &NullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner{value: val, isSet: true}
}

func (v NullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsUsersRecordsQueryPost200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


