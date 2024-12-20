/*
Attio OAuth flow

Welcome to the Attio OAuth flow documentation. All of the OAuth APIs are documented here.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OauthTokenPost200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OauthTokenPost200Response{}

// OauthTokenPost200Response struct for OauthTokenPost200Response
type OauthTokenPost200Response struct {
	// An access token for this workspace.
	AccessToken string `json:"access_token"`
	// The type of token.
	TokenType string `json:"token_type"`
}

type _OauthTokenPost200Response OauthTokenPost200Response

// NewOauthTokenPost200Response instantiates a new OauthTokenPost200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauthTokenPost200Response(accessToken string, tokenType string) *OauthTokenPost200Response {
	this := OauthTokenPost200Response{}
	this.AccessToken = accessToken
	this.TokenType = tokenType
	return &this
}

// NewOauthTokenPost200ResponseWithDefaults instantiates a new OauthTokenPost200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauthTokenPost200ResponseWithDefaults() *OauthTokenPost200Response {
	this := OauthTokenPost200Response{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *OauthTokenPost200Response) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *OauthTokenPost200Response) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *OauthTokenPost200Response) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetTokenType returns the TokenType field value
func (o *OauthTokenPost200Response) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *OauthTokenPost200Response) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *OauthTokenPost200Response) SetTokenType(v string) {
	o.TokenType = v
}

func (o OauthTokenPost200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OauthTokenPost200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["token_type"] = o.TokenType
	return toSerialize, nil
}

func (o *OauthTokenPost200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
		"token_type",
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

	varOauthTokenPost200Response := _OauthTokenPost200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	//decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOauthTokenPost200Response)

	if err != nil {
		return err
	}

	*o = OauthTokenPost200Response(varOauthTokenPost200Response)

	return err
}

type NullableOauthTokenPost200Response struct {
	value *OauthTokenPost200Response
	isSet bool
}

func (v NullableOauthTokenPost200Response) Get() *OauthTokenPost200Response {
	return v.value
}

func (v *NullableOauthTokenPost200Response) Set(val *OauthTokenPost200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthTokenPost200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthTokenPost200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthTokenPost200Response(val *OauthTokenPost200Response) *NullableOauthTokenPost200Response {
	return &NullableOauthTokenPost200Response{value: val, isSet: true}
}

func (v NullableOauthTokenPost200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthTokenPost200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


