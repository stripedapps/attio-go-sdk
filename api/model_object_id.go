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

// checks if the ObjectId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectId{}

// ObjectId struct for ObjectId
type ObjectId struct {
	// A UUID to identify the workspace this object belongs to.
	WorkspaceId string `json:"workspace_id"`
	// A UUID to identify the object.
	ObjectId string `json:"object_id"`
}

type _ObjectId ObjectId

// NewObjectId instantiates a new ObjectId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectId(workspaceId string, objectId string) *ObjectId {
	this := ObjectId{}
	this.WorkspaceId = workspaceId
	this.ObjectId = objectId
	return &this
}

// NewObjectIdWithDefaults instantiates a new ObjectId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectIdWithDefaults() *ObjectId {
	this := ObjectId{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *ObjectId) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *ObjectId) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *ObjectId) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetObjectId returns the ObjectId field value
func (o *ObjectId) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *ObjectId) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *ObjectId) SetObjectId(v string) {
	o.ObjectId = v
}

func (o ObjectId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["workspace_id"] = o.WorkspaceId
	toSerialize["object_id"] = o.ObjectId
	return toSerialize, nil
}

func (o *ObjectId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"workspace_id",
		"object_id",
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

	varObjectId := _ObjectId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varObjectId)

	if err != nil {
		return err
	}

	*o = ObjectId(varObjectId)

	return err
}

type NullableObjectId struct {
	value *ObjectId
	isSet bool
}

func (v NullableObjectId) Get() *ObjectId {
	return v.value
}

func (v *NullableObjectId) Set(val *ObjectId) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectId) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectId(val *ObjectId) *NullableObjectId {
	return &NullableObjectId{value: val, isSet: true}
}

func (v NullableObjectId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

