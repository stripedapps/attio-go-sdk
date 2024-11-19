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
)

// checks if the V2TasksTaskIdPatchRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TasksTaskIdPatchRequestData{}

// V2TasksTaskIdPatchRequestData struct for V2TasksTaskIdPatchRequestData
type V2TasksTaskIdPatchRequestData struct {
	// The deadline of the task, in ISO 8601 format.
	DeadlineAt *string `json:"deadline_at,omitempty"`
	// Whether the task has been completed.
	IsCompleted *bool `json:"is_completed,omitempty"`
	// Records linked to the task. Creating record links within task content text is not possible via the API at present.
	LinkedRecords []V2TasksTaskIdPatchRequestDataLinkedRecordsInner `json:"linked_records,omitempty"`
	// Workspace members assigned to this task.
	Assignees []V2TasksTaskIdPatchRequestDataAssigneesInner `json:"assignees,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V2TasksTaskIdPatchRequestData V2TasksTaskIdPatchRequestData

// NewV2TasksTaskIdPatchRequestData instantiates a new V2TasksTaskIdPatchRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TasksTaskIdPatchRequestData() *V2TasksTaskIdPatchRequestData {
	this := V2TasksTaskIdPatchRequestData{}
	return &this
}

// NewV2TasksTaskIdPatchRequestDataWithDefaults instantiates a new V2TasksTaskIdPatchRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TasksTaskIdPatchRequestDataWithDefaults() *V2TasksTaskIdPatchRequestData {
	this := V2TasksTaskIdPatchRequestData{}
	return &this
}

// GetDeadlineAt returns the DeadlineAt field value if set, zero value otherwise.
func (o *V2TasksTaskIdPatchRequestData) GetDeadlineAt() string {
	if o == nil || IsNil(o.DeadlineAt) {
		var ret string
		return ret
	}
	return *o.DeadlineAt
}

// GetDeadlineAtOk returns a tuple with the DeadlineAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TasksTaskIdPatchRequestData) GetDeadlineAtOk() (*string, bool) {
	if o == nil || IsNil(o.DeadlineAt) {
		return nil, false
	}
	return o.DeadlineAt, true
}

// HasDeadlineAt returns a boolean if a field has been set.
func (o *V2TasksTaskIdPatchRequestData) HasDeadlineAt() bool {
	if o != nil && !IsNil(o.DeadlineAt) {
		return true
	}

	return false
}

// SetDeadlineAt gets a reference to the given string and assigns it to the DeadlineAt field.
func (o *V2TasksTaskIdPatchRequestData) SetDeadlineAt(v string) {
	o.DeadlineAt = &v
}

// GetIsCompleted returns the IsCompleted field value if set, zero value otherwise.
func (o *V2TasksTaskIdPatchRequestData) GetIsCompleted() bool {
	if o == nil || IsNil(o.IsCompleted) {
		var ret bool
		return ret
	}
	return *o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TasksTaskIdPatchRequestData) GetIsCompletedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCompleted) {
		return nil, false
	}
	return o.IsCompleted, true
}

// HasIsCompleted returns a boolean if a field has been set.
func (o *V2TasksTaskIdPatchRequestData) HasIsCompleted() bool {
	if o != nil && !IsNil(o.IsCompleted) {
		return true
	}

	return false
}

// SetIsCompleted gets a reference to the given bool and assigns it to the IsCompleted field.
func (o *V2TasksTaskIdPatchRequestData) SetIsCompleted(v bool) {
	o.IsCompleted = &v
}

// GetLinkedRecords returns the LinkedRecords field value if set, zero value otherwise.
func (o *V2TasksTaskIdPatchRequestData) GetLinkedRecords() []V2TasksTaskIdPatchRequestDataLinkedRecordsInner {
	if o == nil || IsNil(o.LinkedRecords) {
		var ret []V2TasksTaskIdPatchRequestDataLinkedRecordsInner
		return ret
	}
	return o.LinkedRecords
}

// GetLinkedRecordsOk returns a tuple with the LinkedRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TasksTaskIdPatchRequestData) GetLinkedRecordsOk() ([]V2TasksTaskIdPatchRequestDataLinkedRecordsInner, bool) {
	if o == nil || IsNil(o.LinkedRecords) {
		return nil, false
	}
	return o.LinkedRecords, true
}

// HasLinkedRecords returns a boolean if a field has been set.
func (o *V2TasksTaskIdPatchRequestData) HasLinkedRecords() bool {
	if o != nil && !IsNil(o.LinkedRecords) {
		return true
	}

	return false
}

// SetLinkedRecords gets a reference to the given []V2TasksTaskIdPatchRequestDataLinkedRecordsInner and assigns it to the LinkedRecords field.
func (o *V2TasksTaskIdPatchRequestData) SetLinkedRecords(v []V2TasksTaskIdPatchRequestDataLinkedRecordsInner) {
	o.LinkedRecords = v
}

// GetAssignees returns the Assignees field value if set, zero value otherwise.
func (o *V2TasksTaskIdPatchRequestData) GetAssignees() []V2TasksTaskIdPatchRequestDataAssigneesInner {
	if o == nil || IsNil(o.Assignees) {
		var ret []V2TasksTaskIdPatchRequestDataAssigneesInner
		return ret
	}
	return o.Assignees
}

// GetAssigneesOk returns a tuple with the Assignees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TasksTaskIdPatchRequestData) GetAssigneesOk() ([]V2TasksTaskIdPatchRequestDataAssigneesInner, bool) {
	if o == nil || IsNil(o.Assignees) {
		return nil, false
	}
	return o.Assignees, true
}

// HasAssignees returns a boolean if a field has been set.
func (o *V2TasksTaskIdPatchRequestData) HasAssignees() bool {
	if o != nil && !IsNil(o.Assignees) {
		return true
	}

	return false
}

// SetAssignees gets a reference to the given []V2TasksTaskIdPatchRequestDataAssigneesInner and assigns it to the Assignees field.
func (o *V2TasksTaskIdPatchRequestData) SetAssignees(v []V2TasksTaskIdPatchRequestDataAssigneesInner) {
	o.Assignees = v
}

func (o V2TasksTaskIdPatchRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TasksTaskIdPatchRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeadlineAt) {
		toSerialize["deadline_at"] = o.DeadlineAt
	}
	if !IsNil(o.IsCompleted) {
		toSerialize["is_completed"] = o.IsCompleted
	}
	if !IsNil(o.LinkedRecords) {
		toSerialize["linked_records"] = o.LinkedRecords
	}
	if !IsNil(o.Assignees) {
		toSerialize["assignees"] = o.Assignees
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2TasksTaskIdPatchRequestData) UnmarshalJSON(data []byte) (err error) {
	varV2TasksTaskIdPatchRequestData := _V2TasksTaskIdPatchRequestData{}

	err = json.Unmarshal(data, &varV2TasksTaskIdPatchRequestData)

	if err != nil {
		return err
	}

	*o = V2TasksTaskIdPatchRequestData(varV2TasksTaskIdPatchRequestData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "deadline_at")
		delete(additionalProperties, "is_completed")
		delete(additionalProperties, "linked_records")
		delete(additionalProperties, "assignees")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2TasksTaskIdPatchRequestData struct {
	value *V2TasksTaskIdPatchRequestData
	isSet bool
}

func (v NullableV2TasksTaskIdPatchRequestData) Get() *V2TasksTaskIdPatchRequestData {
	return v.value
}

func (v *NullableV2TasksTaskIdPatchRequestData) Set(val *V2TasksTaskIdPatchRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TasksTaskIdPatchRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TasksTaskIdPatchRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TasksTaskIdPatchRequestData(val *V2TasksTaskIdPatchRequestData) *NullableV2TasksTaskIdPatchRequestData {
	return &NullableV2TasksTaskIdPatchRequestData{value: val, isSet: true}
}

func (v NullableV2TasksTaskIdPatchRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TasksTaskIdPatchRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

