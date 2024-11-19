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
	"time"
	"fmt"
)

// checks if the V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3{}

// V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 struct for V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3
type V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor `json:"created_by_actor"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	// A date represents a single calendar year, month and day, independent of timezone. If hours, months, seconds or timezones are provided, they will be trimmed. For example, \"2023\" and \"2023-01\" will be coerced into \"2023-01-01\", and \"2023-01-02\", \"2023-01-02T13:00\", \"2023-01-02T14:00:00\", \"2023-01-02T15:00:00.000000000\", and \"2023-01-02T15:00:00.000000000+02:00\" will all be coerced to \"2023-01-02\". If a timezone is provided that would result in a different calendar date in UTC, the date will be coerced to UTC and then the timezone component will be trimmed. For example, the value \"2023-01-02T23:00:00-10:00\" will be returned as \"2023-01-03\". The maximum date is \"9999-12-31\".
	Value string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, attributeType string, value string) *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3{}
	this.ActiveFrom = activeFrom
	this.ActiveUntil = activeUntil
	this.CreatedByActor = createdByActor
	this.AttributeType = attributeType
	this.Value = value
	return &this
}

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3WithDefaults instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3WithDefaults() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3{}
	return &this
}

// GetActiveFrom returns the ActiveFrom field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetActiveFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActiveFrom
}

// GetActiveFromOk returns a tuple with the ActiveFrom field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetActiveFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveFrom, true
}

// SetActiveFrom sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) SetActiveFrom(v time.Time) {
	o.ActiveFrom = v
}

// GetActiveUntil returns the ActiveUntil field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetActiveUntil() time.Time {
	if o == nil || o.ActiveUntil.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveUntil.Get()
}

// GetActiveUntilOk returns a tuple with the ActiveUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetActiveUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveUntil.Get(), o.ActiveUntil.IsSet()
}

// SetActiveUntil sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) SetActiveUntil(v time.Time) {
	o.ActiveUntil.Set(&v)
}

// GetCreatedByActor returns the CreatedByActor field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor {
	if o == nil {
		var ret V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor) {
	o.CreatedByActor = v
}

// GetAttributeType returns the AttributeType field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) SetAttributeType(v string) {
	o.AttributeType = v
}

// GetValue returns the Value field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) SetValue(v string) {
	o.Value = v
}

func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_from"] = o.ActiveFrom
	toSerialize["active_until"] = o.ActiveUntil.Get()
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["attribute_type"] = o.AttributeType
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active_from",
		"active_until",
		"created_by_actor",
		"attribute_type",
		"value",
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

	varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 := _V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3{}

	err = json.Unmarshal(data, &varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3)

	if err != nil {
		return err
	}

	*o = V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3(varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_from")
		delete(additionalProperties, "active_until")
		delete(additionalProperties, "created_by_actor")
		delete(additionalProperties, "attribute_type")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 struct {
	value *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3
	isSet bool
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) Get() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 {
	return v.value
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) Set(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3 {
	return &NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3{value: val, isSet: true}
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


