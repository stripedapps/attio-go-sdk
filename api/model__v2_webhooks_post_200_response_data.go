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

// checks if the V2WebhooksPost200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2WebhooksPost200ResponseData{}

// V2WebhooksPost200ResponseData struct for V2WebhooksPost200ResponseData
type V2WebhooksPost200ResponseData struct {
	// URL where the webhook events will be delivered to.
	TargetUrl string `json:"target_url" validate:"regexp=^https:\\/\\/.*"`
	// One or more events the webhook is subscribed to.
	Subscriptions []V2WebhooksGet200ResponseDataInnerSubscriptionsInner `json:"subscriptions"`
	Id V2WebhooksGet200ResponseDataInnerId `json:"id"`
	// The state of the webhook. Webhooks marked as active and degraded will receive events, inactive ones will not. If a webhook remains in the degraded state for 7 days, it will be marked inactive.
	Status string `json:"status"`
	// When the webhook was created.
	CreatedAt string `json:"created_at"`
	// The key which is used to sign the webhook events. This is only shown when setting up the webhook initially.
	Secret string `json:"secret"`
}

type _V2WebhooksPost200ResponseData V2WebhooksPost200ResponseData

// NewV2WebhooksPost200ResponseData instantiates a new V2WebhooksPost200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2WebhooksPost200ResponseData(targetUrl string, subscriptions []V2WebhooksGet200ResponseDataInnerSubscriptionsInner, id V2WebhooksGet200ResponseDataInnerId, status string, createdAt string, secret string) *V2WebhooksPost200ResponseData {
	this := V2WebhooksPost200ResponseData{}
	this.TargetUrl = targetUrl
	this.Subscriptions = subscriptions
	this.Id = id
	this.Status = status
	this.CreatedAt = createdAt
	this.Secret = secret
	return &this
}

// NewV2WebhooksPost200ResponseDataWithDefaults instantiates a new V2WebhooksPost200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2WebhooksPost200ResponseDataWithDefaults() *V2WebhooksPost200ResponseData {
	this := V2WebhooksPost200ResponseData{}
	return &this
}

// GetTargetUrl returns the TargetUrl field value
func (o *V2WebhooksPost200ResponseData) GetTargetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksPost200ResponseData) GetTargetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetUrl, true
}

// SetTargetUrl sets field value
func (o *V2WebhooksPost200ResponseData) SetTargetUrl(v string) {
	o.TargetUrl = v
}

// GetSubscriptions returns the Subscriptions field value
func (o *V2WebhooksPost200ResponseData) GetSubscriptions() []V2WebhooksGet200ResponseDataInnerSubscriptionsInner {
	if o == nil {
		var ret []V2WebhooksGet200ResponseDataInnerSubscriptionsInner
		return ret
	}

	return o.Subscriptions
}

// GetSubscriptionsOk returns a tuple with the Subscriptions field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksPost200ResponseData) GetSubscriptionsOk() ([]V2WebhooksGet200ResponseDataInnerSubscriptionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subscriptions, true
}

// SetSubscriptions sets field value
func (o *V2WebhooksPost200ResponseData) SetSubscriptions(v []V2WebhooksGet200ResponseDataInnerSubscriptionsInner) {
	o.Subscriptions = v
}

// GetId returns the Id field value
func (o *V2WebhooksPost200ResponseData) GetId() V2WebhooksGet200ResponseDataInnerId {
	if o == nil {
		var ret V2WebhooksGet200ResponseDataInnerId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksPost200ResponseData) GetIdOk() (*V2WebhooksGet200ResponseDataInnerId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *V2WebhooksPost200ResponseData) SetId(v V2WebhooksGet200ResponseDataInnerId) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *V2WebhooksPost200ResponseData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksPost200ResponseData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *V2WebhooksPost200ResponseData) SetStatus(v string) {
	o.Status = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *V2WebhooksPost200ResponseData) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksPost200ResponseData) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *V2WebhooksPost200ResponseData) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetSecret returns the Secret field value
func (o *V2WebhooksPost200ResponseData) GetSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value
// and a boolean to check if the value has been set.
func (o *V2WebhooksPost200ResponseData) GetSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Secret, true
}

// SetSecret sets field value
func (o *V2WebhooksPost200ResponseData) SetSecret(v string) {
	o.Secret = v
}

func (o V2WebhooksPost200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2WebhooksPost200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["target_url"] = o.TargetUrl
	toSerialize["subscriptions"] = o.Subscriptions
	toSerialize["id"] = o.Id
	toSerialize["status"] = o.Status
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["secret"] = o.Secret
	return toSerialize, nil
}

func (o *V2WebhooksPost200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"target_url",
		"subscriptions",
		"id",
		"status",
		"created_at",
		"secret",
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

	varV2WebhooksPost200ResponseData := _V2WebhooksPost200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2WebhooksPost200ResponseData)

	if err != nil {
		return err
	}

	*o = V2WebhooksPost200ResponseData(varV2WebhooksPost200ResponseData)

	return err
}

type NullableV2WebhooksPost200ResponseData struct {
	value *V2WebhooksPost200ResponseData
	isSet bool
}

func (v NullableV2WebhooksPost200ResponseData) Get() *V2WebhooksPost200ResponseData {
	return v.value
}

func (v *NullableV2WebhooksPost200ResponseData) Set(val *V2WebhooksPost200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2WebhooksPost200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2WebhooksPost200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2WebhooksPost200ResponseData(val *V2WebhooksPost200ResponseData) *NullableV2WebhooksPost200ResponseData {
	return &NullableV2WebhooksPost200ResponseData{value: val, isSet: true}
}

func (v NullableV2WebhooksPost200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2WebhooksPost200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


