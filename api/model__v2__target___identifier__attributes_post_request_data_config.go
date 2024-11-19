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

// checks if the V2TargetIdentifierAttributesPostRequestDataConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TargetIdentifierAttributesPostRequestDataConfig{}

// V2TargetIdentifierAttributesPostRequestDataConfig struct for V2TargetIdentifierAttributesPostRequestDataConfig
type V2TargetIdentifierAttributesPostRequestDataConfig struct {
	Currency *V2TargetIdentifierAttributesPostRequestDataConfigCurrency `json:"currency,omitempty"`
	RecordReference *V2TargetIdentifierAttributesPostRequestDataConfigRecordReference `json:"record_reference,omitempty"`
}

// NewV2TargetIdentifierAttributesPostRequestDataConfig instantiates a new V2TargetIdentifierAttributesPostRequestDataConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TargetIdentifierAttributesPostRequestDataConfig() *V2TargetIdentifierAttributesPostRequestDataConfig {
	this := V2TargetIdentifierAttributesPostRequestDataConfig{}
	return &this
}

// NewV2TargetIdentifierAttributesPostRequestDataConfigWithDefaults instantiates a new V2TargetIdentifierAttributesPostRequestDataConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TargetIdentifierAttributesPostRequestDataConfigWithDefaults() *V2TargetIdentifierAttributesPostRequestDataConfig {
	this := V2TargetIdentifierAttributesPostRequestDataConfig{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesPostRequestDataConfig) GetCurrency() V2TargetIdentifierAttributesPostRequestDataConfigCurrency {
	if o == nil || IsNil(o.Currency) {
		var ret V2TargetIdentifierAttributesPostRequestDataConfigCurrency
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesPostRequestDataConfig) GetCurrencyOk() (*V2TargetIdentifierAttributesPostRequestDataConfigCurrency, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesPostRequestDataConfig) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given V2TargetIdentifierAttributesPostRequestDataConfigCurrency and assigns it to the Currency field.
func (o *V2TargetIdentifierAttributesPostRequestDataConfig) SetCurrency(v V2TargetIdentifierAttributesPostRequestDataConfigCurrency) {
	o.Currency = &v
}

// GetRecordReference returns the RecordReference field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesPostRequestDataConfig) GetRecordReference() V2TargetIdentifierAttributesPostRequestDataConfigRecordReference {
	if o == nil || IsNil(o.RecordReference) {
		var ret V2TargetIdentifierAttributesPostRequestDataConfigRecordReference
		return ret
	}
	return *o.RecordReference
}

// GetRecordReferenceOk returns a tuple with the RecordReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesPostRequestDataConfig) GetRecordReferenceOk() (*V2TargetIdentifierAttributesPostRequestDataConfigRecordReference, bool) {
	if o == nil || IsNil(o.RecordReference) {
		return nil, false
	}
	return o.RecordReference, true
}

// HasRecordReference returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesPostRequestDataConfig) HasRecordReference() bool {
	if o != nil && !IsNil(o.RecordReference) {
		return true
	}

	return false
}

// SetRecordReference gets a reference to the given V2TargetIdentifierAttributesPostRequestDataConfigRecordReference and assigns it to the RecordReference field.
func (o *V2TargetIdentifierAttributesPostRequestDataConfig) SetRecordReference(v V2TargetIdentifierAttributesPostRequestDataConfigRecordReference) {
	o.RecordReference = &v
}

func (o V2TargetIdentifierAttributesPostRequestDataConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TargetIdentifierAttributesPostRequestDataConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.RecordReference) {
		toSerialize["record_reference"] = o.RecordReference
	}
	return toSerialize, nil
}

type NullableV2TargetIdentifierAttributesPostRequestDataConfig struct {
	value *V2TargetIdentifierAttributesPostRequestDataConfig
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesPostRequestDataConfig) Get() *V2TargetIdentifierAttributesPostRequestDataConfig {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesPostRequestDataConfig) Set(val *V2TargetIdentifierAttributesPostRequestDataConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesPostRequestDataConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesPostRequestDataConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesPostRequestDataConfig(val *V2TargetIdentifierAttributesPostRequestDataConfig) *NullableV2TargetIdentifierAttributesPostRequestDataConfig {
	return &NullableV2TargetIdentifierAttributesPostRequestDataConfig{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesPostRequestDataConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesPostRequestDataConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

