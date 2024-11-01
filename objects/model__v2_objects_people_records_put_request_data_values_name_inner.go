/*
Attio Standard Objects

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V2ObjectsPeopleRecordsPutRequestDataValuesNameInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsPutRequestDataValuesNameInner{}

// V2ObjectsPeopleRecordsPutRequestDataValuesNameInner struct for V2ObjectsPeopleRecordsPutRequestDataValuesNameInner
type V2ObjectsPeopleRecordsPutRequestDataValuesNameInner struct {
	// The first name.
	FirstName *string `json:"first_name,omitempty"`
	// The last name.
	LastName *string `json:"last_name,omitempty"`
	// The full name.
	FullName *string `json:"full_name,omitempty"`
}

// NewV2ObjectsPeopleRecordsPutRequestDataValuesNameInner instantiates a new V2ObjectsPeopleRecordsPutRequestDataValuesNameInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsPutRequestDataValuesNameInner() *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner {
	this := V2ObjectsPeopleRecordsPutRequestDataValuesNameInner{}
	return &this
}

// NewV2ObjectsPeopleRecordsPutRequestDataValuesNameInnerWithDefaults instantiates a new V2ObjectsPeopleRecordsPutRequestDataValuesNameInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsPutRequestDataValuesNameInnerWithDefaults() *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner {
	this := V2ObjectsPeopleRecordsPutRequestDataValuesNameInner{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) SetLastName(v string) {
	o.LastName = &v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) SetFullName(v string) {
	o.FullName = &v
}

func (o V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.FullName) {
		toSerialize["full_name"] = o.FullName
	}
	return toSerialize, nil
}

type NullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner struct {
	value *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner) Get() *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner) Set(val *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner(val *V2ObjectsPeopleRecordsPutRequestDataValuesNameInner) *NullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner {
	return &NullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsPutRequestDataValuesNameInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


