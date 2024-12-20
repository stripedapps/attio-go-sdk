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

// checks if the V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6{}

// V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 struct for V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6
type V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor `json:"created_by_actor"`
	Status Status `json:"status"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6

// NewV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 instantiates a new V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, status Status, attributeType string) *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 {
	this := V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6{}
	this.ActiveFrom = activeFrom
	this.ActiveUntil = activeUntil
	this.CreatedByActor = createdByActor
	this.Status = status
	this.AttributeType = attributeType
	return &this
}

// NewV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6WithDefaults instantiates a new V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6WithDefaults() *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 {
	this := V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6{}
	return &this
}

// GetActiveFrom returns the ActiveFrom field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetActiveFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActiveFrom
}

// GetActiveFromOk returns a tuple with the ActiveFrom field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetActiveFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveFrom, true
}

// SetActiveFrom sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) SetActiveFrom(v time.Time) {
	o.ActiveFrom = v
}

// GetActiveUntil returns the ActiveUntil field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetActiveUntil() time.Time {
	if o == nil || o.ActiveUntil.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveUntil.Get()
}

// GetActiveUntilOk returns a tuple with the ActiveUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetActiveUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveUntil.Get(), o.ActiveUntil.IsSet()
}

// SetActiveUntil sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) SetActiveUntil(v time.Time) {
	o.ActiveUntil.Set(&v)
}

// GetCreatedByActor returns the CreatedByActor field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetCreatedByActor() V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor {
	if o == nil {
		var ret V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetCreatedByActorOk() (*V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) SetCreatedByActor(v V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor) {
	o.CreatedByActor = v
}

// GetStatus returns the Status field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetStatus() Status {
	if o == nil {
		var ret Status
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetStatusOk() (*Status, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) SetStatus(v Status) {
	o.Status = v
}

// GetAttributeType returns the AttributeType field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) SetAttributeType(v string) {
	o.AttributeType = v
}

func (o V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_from"] = o.ActiveFrom
	toSerialize["active_until"] = o.ActiveUntil.Get()
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["status"] = o.Status
	toSerialize["attribute_type"] = o.AttributeType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active_from",
		"active_until",
		"created_by_actor",
		"status",
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

	varV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 := _V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6{}

	err = json.Unmarshal(data, &varV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6)

	if err != nil {
		return err
	}

	*o = V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6(varV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_from")
		delete(additionalProperties, "active_until")
		delete(additionalProperties, "created_by_actor")
		delete(additionalProperties, "status")
		delete(additionalProperties, "attribute_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 struct {
	value *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) Get() *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) Set(val *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6(val *V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) *NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 {
	return &NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


