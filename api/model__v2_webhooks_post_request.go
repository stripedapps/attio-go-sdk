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

// checks if the V2WebhooksPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2WebhooksPostRequest{}

// V2WebhooksPostRequest struct for V2WebhooksPostRequest
type V2WebhooksPostRequest struct {
	Data V2WebhooksPostRequestData `json:"data"`
}

type _V2WebhooksPostRequest V2WebhooksPostRequest

// NewV2WebhooksPostRequest instantiates a new V2WebhooksPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2WebhooksPostRequest(data V2WebhooksPostRequestData) *V2WebhooksPostRequest {
	this := V2WebhooksPostRequest{}
	this.Data = data
	return &this
}

// NewV2WebhooksPostRequestWithDefaults instantiates a new V2WebhooksPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2WebhooksPostRequestWithDefaults() *V2WebhooksPostRequest {
	this := V2WebhooksPostRequest{}
	return &this
}

// GetData returns the Data field value
func (o *V2WebhooksPostRequest) GetData() V2WebhooksPostRequestData {
	if o == nil {
		var ret V2WebhooksPostRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksPostRequest) GetDataOk() (*V2WebhooksPostRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *V2WebhooksPostRequest) SetData(v V2WebhooksPostRequestData) {
	o.Data = v
}

func (o V2WebhooksPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2WebhooksPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *V2WebhooksPostRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varV2WebhooksPostRequest := _V2WebhooksPostRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2WebhooksPostRequest)

	if err != nil {
		return err
	}

	*o = V2WebhooksPostRequest(varV2WebhooksPostRequest)

	return err
}

type NullableV2WebhooksPostRequest struct {
	value *V2WebhooksPostRequest
	isSet bool
}

func (v NullableV2WebhooksPostRequest) Get() *V2WebhooksPostRequest {
	return v.value
}

func (v *NullableV2WebhooksPostRequest) Set(val *V2WebhooksPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2WebhooksPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2WebhooksPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2WebhooksPostRequest(val *V2WebhooksPostRequest) *NullableV2WebhooksPostRequest {
	return &NullableV2WebhooksPostRequest{value: val, isSet: true}
}

func (v NullableV2WebhooksPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2WebhooksPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

