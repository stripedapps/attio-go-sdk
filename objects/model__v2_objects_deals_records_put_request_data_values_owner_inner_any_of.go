/*
Attio Standard Objects

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// checks if the V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf{}

// V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf struct for V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf
type V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf struct {
	// The type of the referenced actor. Currently, only workspace members can be written into actor reference attributes. [Read more information on actor types here](/docs/actors).
	ReferencedActorType string `json:"referenced_actor_type"`
	// The ID of the referenced Actor.
	ReferencedActorId string `json:"referenced_actor_id"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf

// NewV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf instantiates a new V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf(referencedActorType string, referencedActorId string) *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf {
	this := V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf{}
	this.ReferencedActorType = referencedActorType
	this.ReferencedActorId = referencedActorId
	return &this
}

// NewV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOfWithDefaults instantiates a new V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOfWithDefaults() *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf {
	this := V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf{}
	return &this
}

// GetReferencedActorType returns the ReferencedActorType field value
func (o *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) GetReferencedActorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferencedActorType
}

// GetReferencedActorTypeOk returns a tuple with the ReferencedActorType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) GetReferencedActorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferencedActorType, true
}

// SetReferencedActorType sets field value
func (o *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) SetReferencedActorType(v string) {
	o.ReferencedActorType = v
}

// GetReferencedActorId returns the ReferencedActorId field value
func (o *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) GetReferencedActorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferencedActorId
}

// GetReferencedActorIdOk returns a tuple with the ReferencedActorId field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) GetReferencedActorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferencedActorId, true
}

// SetReferencedActorId sets field value
func (o *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) SetReferencedActorId(v string) {
	o.ReferencedActorId = v
}

func (o V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenced_actor_type"] = o.ReferencedActorType
	toSerialize["referenced_actor_id"] = o.ReferencedActorId

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"referenced_actor_type",
		"referenced_actor_id",
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

	varV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf := _V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf{}

	err = json.Unmarshal(data, &varV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf)

	if err != nil {
		return err
	}

	*o = V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf(varV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "referenced_actor_type")
		delete(additionalProperties, "referenced_actor_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf struct {
	value *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf
	isSet bool
}

func (v NullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) Get() *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf {
	return v.value
}

func (v *NullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) Set(val *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf(val *V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) *NullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf {
	return &NullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf{value: val, isSet: true}
}

func (v NullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


