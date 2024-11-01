/*
Attio Standard Objects

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V2ObjectsDealsRecordsPut200ResponseDataValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsDealsRecordsPut200ResponseDataValues{}

// V2ObjectsDealsRecordsPut200ResponseDataValues An object with `attribute_slug` keys, and an array of value objects as the values. Attributes slugs (for example `name` or `stage`) can be used, including custom attribute slugs.
type V2ObjectsDealsRecordsPut200ResponseDataValues struct {
	Name []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner `json:"name,omitempty"`
	Stage []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 `json:"stage,omitempty"`
	Owner []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf `json:"owner,omitempty"`
	Value []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf2 `json:"value,omitempty"`
	AssociatedPeople []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner `json:"associated_people,omitempty"`
	AssociatedCompany []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner `json:"associated_company,omitempty"`
}

// NewV2ObjectsDealsRecordsPut200ResponseDataValues instantiates a new V2ObjectsDealsRecordsPut200ResponseDataValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsDealsRecordsPut200ResponseDataValues() *V2ObjectsDealsRecordsPut200ResponseDataValues {
	this := V2ObjectsDealsRecordsPut200ResponseDataValues{}
	return &this
}

// NewV2ObjectsDealsRecordsPut200ResponseDataValuesWithDefaults instantiates a new V2ObjectsDealsRecordsPut200ResponseDataValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsDealsRecordsPut200ResponseDataValuesWithDefaults() *V2ObjectsDealsRecordsPut200ResponseDataValues {
	this := V2ObjectsDealsRecordsPut200ResponseDataValues{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetName() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Name) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetNameOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Name field.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) SetName(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Name = v
}

// GetStage returns the Stage field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetStage() []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 {
	if o == nil || IsNil(o.Stage) {
		var ret []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6
		return ret
	}
	return o.Stage
}

// GetStageOk returns a tuple with the Stage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetStageOk() ([]V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6, bool) {
	if o == nil || IsNil(o.Stage) {
		return nil, false
	}
	return o.Stage, true
}

// HasStage returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) HasStage() bool {
	if o != nil && !IsNil(o.Stage) {
		return true
	}

	return false
}

// SetStage gets a reference to the given []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6 and assigns it to the Stage field.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) SetStage(v []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6) {
	o.Stage = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetOwner() []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf {
	if o == nil || IsNil(o.Owner) {
		var ret []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf
		return ret
	}
	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetOwnerOk() ([]V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf and assigns it to the Owner field.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) SetOwner(v []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf) {
	o.Owner = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetValue() []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf2 {
	if o == nil || IsNil(o.Value) {
		var ret []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf2
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetValueOk() ([]V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf2, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf2 and assigns it to the Value field.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) SetValue(v []V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf2) {
	o.Value = v
}

// GetAssociatedPeople returns the AssociatedPeople field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetAssociatedPeople() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner {
	if o == nil || IsNil(o.AssociatedPeople) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner
		return ret
	}
	return o.AssociatedPeople
}

// GetAssociatedPeopleOk returns a tuple with the AssociatedPeople field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetAssociatedPeopleOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner, bool) {
	if o == nil || IsNil(o.AssociatedPeople) {
		return nil, false
	}
	return o.AssociatedPeople, true
}

// HasAssociatedPeople returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) HasAssociatedPeople() bool {
	if o != nil && !IsNil(o.AssociatedPeople) {
		return true
	}

	return false
}

// SetAssociatedPeople gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner and assigns it to the AssociatedPeople field.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) SetAssociatedPeople(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) {
	o.AssociatedPeople = v
}

// GetAssociatedCompany returns the AssociatedCompany field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetAssociatedCompany() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner {
	if o == nil || IsNil(o.AssociatedCompany) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner
		return ret
	}
	return o.AssociatedCompany
}

// GetAssociatedCompanyOk returns a tuple with the AssociatedCompany field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) GetAssociatedCompanyOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner, bool) {
	if o == nil || IsNil(o.AssociatedCompany) {
		return nil, false
	}
	return o.AssociatedCompany, true
}

// HasAssociatedCompany returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) HasAssociatedCompany() bool {
	if o != nil && !IsNil(o.AssociatedCompany) {
		return true
	}

	return false
}

// SetAssociatedCompany gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner and assigns it to the AssociatedCompany field.
func (o *V2ObjectsDealsRecordsPut200ResponseDataValues) SetAssociatedCompany(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) {
	o.AssociatedCompany = v
}

func (o V2ObjectsDealsRecordsPut200ResponseDataValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsDealsRecordsPut200ResponseDataValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Stage) {
		toSerialize["stage"] = o.Stage
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.AssociatedPeople) {
		toSerialize["associated_people"] = o.AssociatedPeople
	}
	if !IsNil(o.AssociatedCompany) {
		toSerialize["associated_company"] = o.AssociatedCompany
	}
	return toSerialize, nil
}

type NullableV2ObjectsDealsRecordsPut200ResponseDataValues struct {
	value *V2ObjectsDealsRecordsPut200ResponseDataValues
	isSet bool
}

func (v NullableV2ObjectsDealsRecordsPut200ResponseDataValues) Get() *V2ObjectsDealsRecordsPut200ResponseDataValues {
	return v.value
}

func (v *NullableV2ObjectsDealsRecordsPut200ResponseDataValues) Set(val *V2ObjectsDealsRecordsPut200ResponseDataValues) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsDealsRecordsPut200ResponseDataValues) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsDealsRecordsPut200ResponseDataValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsDealsRecordsPut200ResponseDataValues(val *V2ObjectsDealsRecordsPut200ResponseDataValues) *NullableV2ObjectsDealsRecordsPut200ResponseDataValues {
	return &NullableV2ObjectsDealsRecordsPut200ResponseDataValues{value: val, isSet: true}
}

func (v NullableV2ObjectsDealsRecordsPut200ResponseDataValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsDealsRecordsPut200ResponseDataValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


