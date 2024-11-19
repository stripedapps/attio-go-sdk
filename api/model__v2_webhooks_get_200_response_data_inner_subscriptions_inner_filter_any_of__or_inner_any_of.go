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

// checks if the V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf{}

// V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf struct for V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf
type V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf struct {
	Field string `json:"field"`
	Operator string `json:"operator"`
	Value string `json:"value"`
}

type _V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf

// NewV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf instantiates a new V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf(field string, operator string, value string) *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf {
	this := V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf{}
	this.Field = field
	this.Operator = operator
	this.Value = value
	return &this
}

// NewV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOfWithDefaults instantiates a new V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOfWithDefaults() *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf {
	this := V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf{}
	return &this
}

// GetField returns the Field field value
func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) SetField(v string) {
	o.Field = v
}

// GetOperator returns the Operator field value
func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) SetOperator(v string) {
	o.Operator = v
}

// GetValue returns the Value field value
func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) SetValue(v string) {
	o.Value = v
}

func (o V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field"] = o.Field
	toSerialize["operator"] = o.Operator
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field",
		"operator",
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

	varV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf := _V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf)

	if err != nil {
		return err
	}

	*o = V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf(varV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf)

	return err
}

type NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf struct {
	value *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf
	isSet bool
}

func (v NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) Get() *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf {
	return v.value
}

func (v *NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) Set(val *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf(val *V2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) *NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf {
	return &NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf{value: val, isSet: true}
}

func (v NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2WebhooksGet200ResponseDataInnerSubscriptionsInnerFilterAnyOfOrInnerAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


