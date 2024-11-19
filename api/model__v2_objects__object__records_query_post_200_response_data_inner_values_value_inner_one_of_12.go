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

// checks if the V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12{}

// V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 struct for V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12
type V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor `json:"created_by_actor"`
	Status Status `json:"status"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, status Status, attributeType string) *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12{}
	this.ActiveFrom = activeFrom
	this.ActiveUntil = activeUntil
	this.CreatedByActor = createdByActor
	this.Status = status
	this.AttributeType = attributeType
	return &this
}

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12WithDefaults instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12WithDefaults() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12{}
	return &this
}

// GetActiveFrom returns the ActiveFrom field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetActiveFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActiveFrom
}

// GetActiveFromOk returns a tuple with the ActiveFrom field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetActiveFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveFrom, true
}

// SetActiveFrom sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) SetActiveFrom(v time.Time) {
	o.ActiveFrom = v
}

// GetActiveUntil returns the ActiveUntil field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetActiveUntil() time.Time {
	if o == nil || o.ActiveUntil.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveUntil.Get()
}

// GetActiveUntilOk returns a tuple with the ActiveUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetActiveUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveUntil.Get(), o.ActiveUntil.IsSet()
}

// SetActiveUntil sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) SetActiveUntil(v time.Time) {
	o.ActiveUntil.Set(&v)
}

// GetCreatedByActor returns the CreatedByActor field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor {
	if o == nil {
		var ret V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor) {
	o.CreatedByActor = v
}

// GetStatus returns the Status field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetStatus() Status {
	if o == nil {
		var ret Status
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetStatusOk() (*Status, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) SetStatus(v Status) {
	o.Status = v
}

// GetAttributeType returns the AttributeType field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) SetAttributeType(v string) {
	o.AttributeType = v
}

func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) ToMap() (map[string]interface{}, error) {
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

func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) UnmarshalJSON(data []byte) (err error) {
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

	varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 := _V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12{}

	err = json.Unmarshal(data, &varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12)

	if err != nil {
		return err
	}

	*o = V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12(varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12)

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

type NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 struct {
	value *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12
	isSet bool
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) Get() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 {
	return v.value
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) Set(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12 {
	return &NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12{value: val, isSet: true}
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf12) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


