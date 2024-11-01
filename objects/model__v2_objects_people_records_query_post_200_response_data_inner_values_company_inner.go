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

// checks if the V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner{}

// V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner struct for V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner
type V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor `json:"created_by_actor"`
	// A slug identifying the object that the referenced record belongs to.
	TargetObject string `json:"target_object"`
	// A UUID to identify the referenced record.
	TargetRecordId string `json:"target_record_id"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner

// NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner instantiates a new V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, targetObject string, targetRecordId string, attributeType string) *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner {
	this := V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner{}
	this.ActiveFrom = activeFrom
	this.ActiveUntil = activeUntil
	this.CreatedByActor = createdByActor
	this.TargetObject = targetObject
	this.TargetRecordId = targetRecordId
	this.AttributeType = attributeType
	return &this
}

// NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInnerWithDefaults instantiates a new V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInnerWithDefaults() *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner {
	this := V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner{}
	return &this
}

// GetActiveFrom returns the ActiveFrom field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetActiveFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActiveFrom
}

// GetActiveFromOk returns a tuple with the ActiveFrom field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetActiveFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveFrom, true
}

// SetActiveFrom sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) SetActiveFrom(v time.Time) {
	o.ActiveFrom = v
}

// GetActiveUntil returns the ActiveUntil field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetActiveUntil() time.Time {
	if o == nil || o.ActiveUntil.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveUntil.Get()
}

// GetActiveUntilOk returns a tuple with the ActiveUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetActiveUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveUntil.Get(), o.ActiveUntil.IsSet()
}

// SetActiveUntil sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) SetActiveUntil(v time.Time) {
	o.ActiveUntil.Set(&v)
}

// GetCreatedByActor returns the CreatedByActor field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetCreatedByActor() V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor {
	if o == nil {
		var ret V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetCreatedByActorOk() (*V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) SetCreatedByActor(v V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor) {
	o.CreatedByActor = v
}

// GetTargetObject returns the TargetObject field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetTargetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetObject
}

// GetTargetObjectOk returns a tuple with the TargetObject field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetTargetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetObject, true
}

// SetTargetObject sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) SetTargetObject(v string) {
	o.TargetObject = v
}

// GetTargetRecordId returns the TargetRecordId field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetTargetRecordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetRecordId
}

// GetTargetRecordIdOk returns a tuple with the TargetRecordId field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetTargetRecordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetRecordId, true
}

// SetTargetRecordId sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) SetTargetRecordId(v string) {
	o.TargetRecordId = v
}

// GetAttributeType returns the AttributeType field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) SetAttributeType(v string) {
	o.AttributeType = v
}

func (o V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_from"] = o.ActiveFrom
	toSerialize["active_until"] = o.ActiveUntil.Get()
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["target_object"] = o.TargetObject
	toSerialize["target_record_id"] = o.TargetRecordId
	toSerialize["attribute_type"] = o.AttributeType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active_from",
		"active_until",
		"created_by_actor",
		"target_object",
		"target_record_id",
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

	varV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner := _V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner{}

	err = json.Unmarshal(data, &varV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner)

	if err != nil {
		return err
	}

	*o = V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner(varV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_from")
		delete(additionalProperties, "active_until")
		delete(additionalProperties, "created_by_actor")
		delete(additionalProperties, "target_object")
		delete(additionalProperties, "target_record_id")
		delete(additionalProperties, "attribute_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner struct {
	value *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) Get() *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) Set(val *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner(val *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner {
	return &NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


