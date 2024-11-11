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

// checks if the V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1{}

// V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 Sort by path
type V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 struct {
	// The direction to sort the results by.
	Direction string `json:"direction"`
	// You may use the `path` property to traverse record reference attributes and parent records on list entries. `path` accepts an array of tuples where the first element of each tuple is the slug or ID of a list/object, and the second element is the slug or ID of an attribute on that list/object. The first element of the first tuple must correspond to the list or object that you are querying. For example, if you wanted to sort by the name of the parent record (a company) on a list with the slug \"sales\", you would pass the value `[['sales', 'parent_record'], ['companies', 'name']]`.
	Path [][]V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner `json:"path"`
	// Which field on the value to sort by e.g. \"last_name\" on a name value.
	Field *string `json:"field,omitempty"`
}

type _V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1

// NewV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 instantiates a new V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1(direction string, path [][]V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 {
	this := V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1{}
	this.Direction = direction
	this.Path = path
	return &this
}

// NewV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1WithDefaults instantiates a new V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1WithDefaults() *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 {
	this := V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1{}
	return &this
}

// GetDirection returns the Direction field value
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) GetDirection() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) GetDirectionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Direction, true
}

// SetDirection sets field value
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) SetDirection(v string) {
	o.Direction = v
}

// GetPath returns the Path field value
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) GetPath() [][]V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner {
	if o == nil {
		var ret [][]V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) GetPathOk() ([][]V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path, true
}

// SetPath sets field value
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) SetPath(v [][]V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner) {
	o.Path = v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) SetField(v string) {
	o.Field = &v
}

func (o V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["direction"] = o.Direction
	toSerialize["path"] = o.Path
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	return toSerialize, nil
}

func (o *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"direction",
		"path",
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

	varV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 := _V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1)

	if err != nil {
		return err
	}

	*o = V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1(varV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1)

	return err
}

type NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 struct {
	value *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) Get() *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) Set(val *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1(val *V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) *NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1 {
	return &NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


