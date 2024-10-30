/*
Attio Standard Objects

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner{}

// V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner struct for V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner
type V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor `json:"created_by_actor"`
	OriginalEmailAddress string `json:"original_email_address"`
	EmailAddress string `json:"email_address"`
	EmailDomain string `json:"email_domain"`
	EmailRootDomain string `json:"email_root_domain"`
	EmailLocalSpecifier string `json:"email_local_specifier"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
}

type _V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner

// NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner instantiates a new V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, originalEmailAddress string, emailAddress string, emailDomain string, emailRootDomain string, emailLocalSpecifier string, attributeType string) *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner {
	this := V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner{}
	this.ActiveFrom = activeFrom
	this.ActiveUntil = activeUntil
	this.CreatedByActor = createdByActor
	this.OriginalEmailAddress = originalEmailAddress
	this.EmailAddress = emailAddress
	this.EmailDomain = emailDomain
	this.EmailRootDomain = emailRootDomain
	this.EmailLocalSpecifier = emailLocalSpecifier
	this.AttributeType = attributeType
	return &this
}

// NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerWithDefaults instantiates a new V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerWithDefaults() *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner {
	this := V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner{}
	return &this
}

// GetActiveFrom returns the ActiveFrom field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetActiveFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActiveFrom
}

// GetActiveFromOk returns a tuple with the ActiveFrom field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetActiveFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveFrom, true
}

// SetActiveFrom sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) SetActiveFrom(v time.Time) {
	o.ActiveFrom = v
}

// GetActiveUntil returns the ActiveUntil field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetActiveUntil() time.Time {
	if o == nil || o.ActiveUntil.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveUntil.Get()
}

// GetActiveUntilOk returns a tuple with the ActiveUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetActiveUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveUntil.Get(), o.ActiveUntil.IsSet()
}

// SetActiveUntil sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) SetActiveUntil(v time.Time) {
	o.ActiveUntil.Set(&v)
}

// GetCreatedByActor returns the CreatedByActor field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetCreatedByActor() V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor {
	if o == nil {
		var ret V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetCreatedByActorOk() (*V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) SetCreatedByActor(v V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor) {
	o.CreatedByActor = v
}

// GetOriginalEmailAddress returns the OriginalEmailAddress field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetOriginalEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginalEmailAddress
}

// GetOriginalEmailAddressOk returns a tuple with the OriginalEmailAddress field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetOriginalEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalEmailAddress, true
}

// SetOriginalEmailAddress sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) SetOriginalEmailAddress(v string) {
	o.OriginalEmailAddress = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetEmailDomain returns the EmailDomain field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetEmailDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailDomain
}

// GetEmailDomainOk returns a tuple with the EmailDomain field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetEmailDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailDomain, true
}

// SetEmailDomain sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) SetEmailDomain(v string) {
	o.EmailDomain = v
}

// GetEmailRootDomain returns the EmailRootDomain field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetEmailRootDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailRootDomain
}

// GetEmailRootDomainOk returns a tuple with the EmailRootDomain field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetEmailRootDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailRootDomain, true
}

// SetEmailRootDomain sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) SetEmailRootDomain(v string) {
	o.EmailRootDomain = v
}

// GetEmailLocalSpecifier returns the EmailLocalSpecifier field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetEmailLocalSpecifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailLocalSpecifier
}

// GetEmailLocalSpecifierOk returns a tuple with the EmailLocalSpecifier field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetEmailLocalSpecifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailLocalSpecifier, true
}

// SetEmailLocalSpecifier sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) SetEmailLocalSpecifier(v string) {
	o.EmailLocalSpecifier = v
}

// GetAttributeType returns the AttributeType field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) SetAttributeType(v string) {
	o.AttributeType = v
}

func (o V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_from"] = o.ActiveFrom
	toSerialize["active_until"] = o.ActiveUntil.Get()
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["original_email_address"] = o.OriginalEmailAddress
	toSerialize["email_address"] = o.EmailAddress
	toSerialize["email_domain"] = o.EmailDomain
	toSerialize["email_root_domain"] = o.EmailRootDomain
	toSerialize["email_local_specifier"] = o.EmailLocalSpecifier
	toSerialize["attribute_type"] = o.AttributeType
	return toSerialize, nil
}

func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active_from",
		"active_until",
		"created_by_actor",
		"original_email_address",
		"email_address",
		"email_domain",
		"email_root_domain",
		"email_local_specifier",
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

	varV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner := _V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner)

	if err != nil {
		return err
	}

	*o = V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner(varV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner)

	return err
}

type NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner struct {
	value *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) Get() *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) Set(val *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner(val *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner {
	return &NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


