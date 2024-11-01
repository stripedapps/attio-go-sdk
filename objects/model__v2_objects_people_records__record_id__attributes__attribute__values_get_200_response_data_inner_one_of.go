/*
Attio Standard Objects

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"fmt"
)

// checks if the V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf{}

// V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf struct for V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf
type V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor `json:"created_by_actor"`
	// The type of the referenced actor. [Read more information on actor types here](/docs/actors).
	ReferencedActorType string `json:"referenced_actor_type"`
	// The ID of the referenced actor.
	ReferencedActorId string `json:"referenced_actor_id"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf

// NewV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf instantiates a new V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, referencedActorType string, referencedActorId string, attributeType string) *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf {
	this := V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf{}
	this.ActiveFrom = activeFrom
	this.ActiveUntil = activeUntil
	this.CreatedByActor = createdByActor
	this.ReferencedActorType = referencedActorType
	this.ReferencedActorId = referencedActorId
	this.AttributeType = attributeType
	return &this
}

// NewV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOfWithDefaults instantiates a new V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOfWithDefaults() *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf {
	this := V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf{}
	return &this
}

// GetActiveFrom returns the ActiveFrom field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetActiveFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActiveFrom
}

// GetActiveFromOk returns a tuple with the ActiveFrom field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetActiveFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveFrom, true
}

// SetActiveFrom sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) SetActiveFrom(v time.Time) {
	o.ActiveFrom = v
}

// GetActiveUntil returns the ActiveUntil field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetActiveUntil() time.Time {
	if o == nil || o.ActiveUntil.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveUntil.Get()
}

// GetActiveUntilOk returns a tuple with the ActiveUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetActiveUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveUntil.Get(), o.ActiveUntil.IsSet()
}

// SetActiveUntil sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) SetActiveUntil(v time.Time) {
	o.ActiveUntil.Set(&v)
}

// GetCreatedByActor returns the CreatedByActor field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetCreatedByActor() V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor {
	if o == nil {
		var ret V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetCreatedByActorOk() (*V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) SetCreatedByActor(v V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor) {
	o.CreatedByActor = v
}

// GetReferencedActorType returns the ReferencedActorType field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetReferencedActorType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferencedActorType
}

// GetReferencedActorTypeOk returns a tuple with the ReferencedActorType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetReferencedActorTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferencedActorType, true
}

// SetReferencedActorType sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) SetReferencedActorType(v string) {
	o.ReferencedActorType = v
}

// GetReferencedActorId returns the ReferencedActorId field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetReferencedActorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferencedActorId
}

// GetReferencedActorIdOk returns a tuple with the ReferencedActorId field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetReferencedActorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferencedActorId, true
}

// SetReferencedActorId sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) SetReferencedActorId(v string) {
	o.ReferencedActorId = v
}

// GetAttributeType returns the AttributeType field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) SetAttributeType(v string) {
	o.AttributeType = v
}

func (o V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_from"] = o.ActiveFrom
	toSerialize["active_until"] = o.ActiveUntil.Get()
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["referenced_actor_type"] = o.ReferencedActorType
	toSerialize["referenced_actor_id"] = o.ReferencedActorId
	toSerialize["attribute_type"] = o.AttributeType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active_from",
		"active_until",
		"created_by_actor",
		"referenced_actor_type",
		"referenced_actor_id",
		"attribute_type",
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

	varV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf := _V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf{}

	err = json.Unmarshal(data, &varV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf)

	if err != nil {
		return err
	}

	*o = V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf(varV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_from")
		delete(additionalProperties, "active_until")
		delete(additionalProperties, "created_by_actor")
		delete(additionalProperties, "referenced_actor_type")
		delete(additionalProperties, "referenced_actor_id")
		delete(additionalProperties, "attribute_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf struct {
	value *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) Get() *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) Set(val *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf(val *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) *NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf {
	return &NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


