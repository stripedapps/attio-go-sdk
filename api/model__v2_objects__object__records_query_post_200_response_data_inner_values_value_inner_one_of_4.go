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

// checks if the V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4{}

// V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 struct for V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4
type V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor `json:"created_by_actor"`
	Domain string `json:"domain"`
	RootDomain string `json:"root_domain"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, domain string, rootDomain string, attributeType string) *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4{}
	this.ActiveFrom = activeFrom
	this.ActiveUntil = activeUntil
	this.CreatedByActor = createdByActor
	this.Domain = domain
	this.RootDomain = rootDomain
	this.AttributeType = attributeType
	return &this
}

// NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4WithDefaults instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4WithDefaults() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 {
	this := V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4{}
	return &this
}

// GetActiveFrom returns the ActiveFrom field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetActiveFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActiveFrom
}

// GetActiveFromOk returns a tuple with the ActiveFrom field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetActiveFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveFrom, true
}

// SetActiveFrom sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) SetActiveFrom(v time.Time) {
	o.ActiveFrom = v
}

// GetActiveUntil returns the ActiveUntil field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetActiveUntil() time.Time {
	if o == nil || o.ActiveUntil.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveUntil.Get()
}

// GetActiveUntilOk returns a tuple with the ActiveUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetActiveUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveUntil.Get(), o.ActiveUntil.IsSet()
}

// SetActiveUntil sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) SetActiveUntil(v time.Time) {
	o.ActiveUntil.Set(&v)
}

// GetCreatedByActor returns the CreatedByActor field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor {
	if o == nil {
		var ret V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOfCreatedByActor) {
	o.CreatedByActor = v
}

// GetDomain returns the Domain field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) SetDomain(v string) {
	o.Domain = v
}

// GetRootDomain returns the RootDomain field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetRootDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RootDomain
}

// GetRootDomainOk returns a tuple with the RootDomain field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetRootDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RootDomain, true
}

// SetRootDomain sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) SetRootDomain(v string) {
	o.RootDomain = v
}

// GetAttributeType returns the AttributeType field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) SetAttributeType(v string) {
	o.AttributeType = v
}

func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_from"] = o.ActiveFrom
	toSerialize["active_until"] = o.ActiveUntil.Get()
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["domain"] = o.Domain
	toSerialize["root_domain"] = o.RootDomain
	toSerialize["attribute_type"] = o.AttributeType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active_from",
		"active_until",
		"created_by_actor",
		"domain",
		"root_domain",
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

	varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 := _V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4{}

	err = json.Unmarshal(data, &varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4)

	if err != nil {
		return err
	}

	*o = V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4(varV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_from")
		delete(additionalProperties, "active_until")
		delete(additionalProperties, "created_by_actor")
		delete(additionalProperties, "domain")
		delete(additionalProperties, "root_domain")
		delete(additionalProperties, "attribute_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 struct {
	value *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4
	isSet bool
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) Get() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 {
	return v.value
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) Set(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4(val *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4 {
	return &NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4{value: val, isSet: true}
}

func (v NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerOneOf4) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

