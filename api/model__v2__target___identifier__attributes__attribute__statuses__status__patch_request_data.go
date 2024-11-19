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

// checks if the V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData{}

// V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData struct for V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData
type V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData struct {
	// The Title of the status
	Title *string `json:"title,omitempty"`
	// Whether arriving at this status triggers a celebration effect
	CelebrationEnabled *bool `json:"celebration_enabled,omitempty"`
	// Target time for a record to spend in given status expressed as a ISO-8601 duration string
	TargetTimeInStatus NullableString `json:"target_time_in_status,omitempty" validate:"regexp=P(?:(\\\\d+Y)?(\\\\d+M)?(\\\\d+W)?(\\\\d+D)?(?:T(\\\\d+(?:[\\\\.,]\\\\d+)?H)?(\\\\d+(?:[\\\\.,]\\\\d+)?M)?(\\\\d+(?:[\\\\.,]\\\\d+)?S)?)?)"`
	// Whether or not to archive the status. See our [archiving guide](/docs/archiving-vs-deleting) for more information on archiving.
	IsArchived *bool `json:"is_archived,omitempty"`
}

// NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData instantiates a new V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData() *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData {
	this := V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData{}
	var celebrationEnabled bool = false
	this.CelebrationEnabled = &celebrationEnabled
	return &this
}

// NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestDataWithDefaults instantiates a new V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestDataWithDefaults() *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData {
	this := V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData{}
	var celebrationEnabled bool = false
	this.CelebrationEnabled = &celebrationEnabled
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetTitle(v string) {
	o.Title = &v
}

// GetCelebrationEnabled returns the CelebrationEnabled field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetCelebrationEnabled() bool {
	if o == nil || IsNil(o.CelebrationEnabled) {
		var ret bool
		return ret
	}
	return *o.CelebrationEnabled
}

// GetCelebrationEnabledOk returns a tuple with the CelebrationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetCelebrationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CelebrationEnabled) {
		return nil, false
	}
	return o.CelebrationEnabled, true
}

// HasCelebrationEnabled returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) HasCelebrationEnabled() bool {
	if o != nil && !IsNil(o.CelebrationEnabled) {
		return true
	}

	return false
}

// SetCelebrationEnabled gets a reference to the given bool and assigns it to the CelebrationEnabled field.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetCelebrationEnabled(v bool) {
	o.CelebrationEnabled = &v
}

// GetTargetTimeInStatus returns the TargetTimeInStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetTargetTimeInStatus() string {
	if o == nil || IsNil(o.TargetTimeInStatus.Get()) {
		var ret string
		return ret
	}
	return *o.TargetTimeInStatus.Get()
}

// GetTargetTimeInStatusOk returns a tuple with the TargetTimeInStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetTargetTimeInStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetTimeInStatus.Get(), o.TargetTimeInStatus.IsSet()
}

// HasTargetTimeInStatus returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) HasTargetTimeInStatus() bool {
	if o != nil && o.TargetTimeInStatus.IsSet() {
		return true
	}

	return false
}

// SetTargetTimeInStatus gets a reference to the given NullableString and assigns it to the TargetTimeInStatus field.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetTargetTimeInStatus(v string) {
	o.TargetTimeInStatus.Set(&v)
}
// SetTargetTimeInStatusNil sets the value for TargetTimeInStatus to be an explicit nil
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetTargetTimeInStatusNil() {
	o.TargetTimeInStatus.Set(nil)
}

// UnsetTargetTimeInStatus ensures that no value is present for TargetTimeInStatus, not even an explicit nil
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) UnsetTargetTimeInStatus() {
	o.TargetTimeInStatus.Unset()
}

// GetIsArchived returns the IsArchived field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetIsArchived() bool {
	if o == nil || IsNil(o.IsArchived) {
		var ret bool
		return ret
	}
	return *o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) GetIsArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsArchived) {
		return nil, false
	}
	return o.IsArchived, true
}

// HasIsArchived returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) HasIsArchived() bool {
	if o != nil && !IsNil(o.IsArchived) {
		return true
	}

	return false
}

// SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.
func (o *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) SetIsArchived(v bool) {
	o.IsArchived = &v
}

func (o V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.CelebrationEnabled) {
		toSerialize["celebration_enabled"] = o.CelebrationEnabled
	}
	if o.TargetTimeInStatus.IsSet() {
		toSerialize["target_time_in_status"] = o.TargetTimeInStatus.Get()
	}
	if !IsNil(o.IsArchived) {
		toSerialize["is_archived"] = o.IsArchived
	}
	return toSerialize, nil
}

type NullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData struct {
	value *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) Get() *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) Set(val *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData(val *V2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) *NullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData {
	return &NullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesAttributeStatusesStatusPatchRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


