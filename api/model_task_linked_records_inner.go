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

// checks if the TaskLinkedRecordsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskLinkedRecordsInner{}

// TaskLinkedRecordsInner struct for TaskLinkedRecordsInner
type TaskLinkedRecordsInner struct {
	// The ID of the parent object the task refers to. At present, only `people` and `companies` are supported.
	TargetObjectId string `json:"target_object_id"`
	// The ID of the parent record the task refers to.
	TargetRecordId string `json:"target_record_id"`
}

type _TaskLinkedRecordsInner TaskLinkedRecordsInner

// NewTaskLinkedRecordsInner instantiates a new TaskLinkedRecordsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskLinkedRecordsInner(targetObjectId string, targetRecordId string) *TaskLinkedRecordsInner {
	this := TaskLinkedRecordsInner{}
	this.TargetObjectId = targetObjectId
	this.TargetRecordId = targetRecordId
	return &this
}

// NewTaskLinkedRecordsInnerWithDefaults instantiates a new TaskLinkedRecordsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskLinkedRecordsInnerWithDefaults() *TaskLinkedRecordsInner {
	this := TaskLinkedRecordsInner{}
	return &this
}

// GetTargetObjectId returns the TargetObjectId field value
func (o *TaskLinkedRecordsInner) GetTargetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetObjectId
}

// GetTargetObjectIdOk returns a tuple with the TargetObjectId field value
// and a boolean to check if the value has been set.
func (o *TaskLinkedRecordsInner) GetTargetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetObjectId, true
}

// SetTargetObjectId sets field value
func (o *TaskLinkedRecordsInner) SetTargetObjectId(v string) {
	o.TargetObjectId = v
}

// GetTargetRecordId returns the TargetRecordId field value
func (o *TaskLinkedRecordsInner) GetTargetRecordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetRecordId
}

// GetTargetRecordIdOk returns a tuple with the TargetRecordId field value
// and a boolean to check if the value has been set.
func (o *TaskLinkedRecordsInner) GetTargetRecordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetRecordId, true
}

// SetTargetRecordId sets field value
func (o *TaskLinkedRecordsInner) SetTargetRecordId(v string) {
	o.TargetRecordId = v
}

func (o TaskLinkedRecordsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskLinkedRecordsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target_object_id"] = o.TargetObjectId
	toSerialize["target_record_id"] = o.TargetRecordId
	return toSerialize, nil
}

func (o *TaskLinkedRecordsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"target_object_id",
		"target_record_id",
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

	varTaskLinkedRecordsInner := _TaskLinkedRecordsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTaskLinkedRecordsInner)

	if err != nil {
		return err
	}

	*o = TaskLinkedRecordsInner(varTaskLinkedRecordsInner)

	return err
}

type NullableTaskLinkedRecordsInner struct {
	value *TaskLinkedRecordsInner
	isSet bool
}

func (v NullableTaskLinkedRecordsInner) Get() *TaskLinkedRecordsInner {
	return v.value
}

func (v *NullableTaskLinkedRecordsInner) Set(val *TaskLinkedRecordsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskLinkedRecordsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskLinkedRecordsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskLinkedRecordsInner(val *TaskLinkedRecordsInner) *NullableTaskLinkedRecordsInner {
	return &NullableTaskLinkedRecordsInner{value: val, isSet: true}
}

func (v NullableTaskLinkedRecordsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskLinkedRecordsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


