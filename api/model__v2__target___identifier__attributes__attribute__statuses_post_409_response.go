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

// checks if the V2TargetIdentifierAttributesAttributeStatusesPost409Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TargetIdentifierAttributesAttributeStatusesPost409Response{}

// V2TargetIdentifierAttributesAttributeStatusesPost409Response Conflict
type V2TargetIdentifierAttributesAttributeStatusesPost409Response struct {
	StatusCode float32 `json:"status_code"`
	Type string `json:"type"`
	Code string `json:"code"`
	Message string `json:"message"`
}

type _V2TargetIdentifierAttributesAttributeStatusesPost409Response V2TargetIdentifierAttributesAttributeStatusesPost409Response

// NewV2TargetIdentifierAttributesAttributeStatusesPost409Response instantiates a new V2TargetIdentifierAttributesAttributeStatusesPost409Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TargetIdentifierAttributesAttributeStatusesPost409Response(statusCode float32, type_ string, code string, message string) *V2TargetIdentifierAttributesAttributeStatusesPost409Response {
	this := V2TargetIdentifierAttributesAttributeStatusesPost409Response{}
	this.StatusCode = statusCode
	this.Type = type_
	this.Code = code
	this.Message = message
	return &this
}

// NewV2TargetIdentifierAttributesAttributeStatusesPost409ResponseWithDefaults instantiates a new V2TargetIdentifierAttributesAttributeStatusesPost409Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TargetIdentifierAttributesAttributeStatusesPost409ResponseWithDefaults() *V2TargetIdentifierAttributesAttributeStatusesPost409Response {
	this := V2TargetIdentifierAttributesAttributeStatusesPost409Response{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) GetStatusCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) GetStatusCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) SetStatusCode(v float32) {
	o.StatusCode = v
}

// GetType returns the Type field value
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) SetType(v string) {
	o.Type = v
}

// GetCode returns the Code field value
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) SetCode(v string) {
	o.Code = v
}

// GetMessage returns the Message field value
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) SetMessage(v string) {
	o.Message = v
}

func (o V2TargetIdentifierAttributesAttributeStatusesPost409Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TargetIdentifierAttributesAttributeStatusesPost409Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status_code"] = o.StatusCode
	toSerialize["type"] = o.Type
	toSerialize["code"] = o.Code
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *V2TargetIdentifierAttributesAttributeStatusesPost409Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status_code",
		"type",
		"code",
		"message",
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

	varV2TargetIdentifierAttributesAttributeStatusesPost409Response := _V2TargetIdentifierAttributesAttributeStatusesPost409Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2TargetIdentifierAttributesAttributeStatusesPost409Response)

	if err != nil {
		return err
	}

	*o = V2TargetIdentifierAttributesAttributeStatusesPost409Response(varV2TargetIdentifierAttributesAttributeStatusesPost409Response)

	return err
}

type NullableV2TargetIdentifierAttributesAttributeStatusesPost409Response struct {
	value *V2TargetIdentifierAttributesAttributeStatusesPost409Response
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesAttributeStatusesPost409Response) Get() *V2TargetIdentifierAttributesAttributeStatusesPost409Response {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesAttributeStatusesPost409Response) Set(val *V2TargetIdentifierAttributesAttributeStatusesPost409Response) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesAttributeStatusesPost409Response) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesAttributeStatusesPost409Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesAttributeStatusesPost409Response(val *V2TargetIdentifierAttributesAttributeStatusesPost409Response) *NullableV2TargetIdentifierAttributesAttributeStatusesPost409Response {
	return &NullableV2TargetIdentifierAttributesAttributeStatusesPost409Response{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesAttributeStatusesPost409Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesAttributeStatusesPost409Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

