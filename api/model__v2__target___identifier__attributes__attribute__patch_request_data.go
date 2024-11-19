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

// checks if the V2TargetIdentifierAttributesAttributePatchRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2TargetIdentifierAttributesAttributePatchRequestData{}

// V2TargetIdentifierAttributesAttributePatchRequestData struct for V2TargetIdentifierAttributesAttributePatchRequestData
type V2TargetIdentifierAttributesAttributePatchRequestData struct {
	// The name of the attribute. The title will be visible across Attio's UI.
	Title *string `json:"title,omitempty"`
	// A text description for the attribute.
	Description *string `json:"description,omitempty"`
	// A unique, human-readable slug to access the attribute through URLs and API calls. Formatted in snake case.
	ApiSlug *string `json:"api_slug,omitempty"`
	// When `is_required` is `true`, new records/entries must have a value for this attribute. If `false`, values may be `null`. This value does not affect existing data and you do not need to backfill `null` values if changing `is_required` from `false` to `true`.
	IsRequired *bool `json:"is_required,omitempty"`
	// Whether or not new values for this attribute must be unique. Uniqueness restrictions are only applied to new data and do not apply retroactively to previously created data.
	IsUnique *bool `json:"is_unique,omitempty"`
	DefaultValue *V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue `json:"default_value,omitempty"`
	Config *V2TargetIdentifierAttributesAttributePatchRequestDataConfig `json:"config,omitempty"`
	// Whether the attribute has been archived or not. See our [archiving guide](/docs/archiving-vs-deleting) for more information on archiving.
	IsArchived *bool `json:"is_archived,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _V2TargetIdentifierAttributesAttributePatchRequestData V2TargetIdentifierAttributesAttributePatchRequestData

// NewV2TargetIdentifierAttributesAttributePatchRequestData instantiates a new V2TargetIdentifierAttributesAttributePatchRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2TargetIdentifierAttributesAttributePatchRequestData() *V2TargetIdentifierAttributesAttributePatchRequestData {
	this := V2TargetIdentifierAttributesAttributePatchRequestData{}
	return &this
}

// NewV2TargetIdentifierAttributesAttributePatchRequestDataWithDefaults instantiates a new V2TargetIdentifierAttributesAttributePatchRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2TargetIdentifierAttributesAttributePatchRequestDataWithDefaults() *V2TargetIdentifierAttributesAttributePatchRequestData {
	this := V2TargetIdentifierAttributesAttributePatchRequestData{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetTitle(v string) {
	o.Title = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetDescription(v string) {
	o.Description = &v
}

// GetApiSlug returns the ApiSlug field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetApiSlug() string {
	if o == nil || IsNil(o.ApiSlug) {
		var ret string
		return ret
	}
	return *o.ApiSlug
}

// GetApiSlugOk returns a tuple with the ApiSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetApiSlugOk() (*string, bool) {
	if o == nil || IsNil(o.ApiSlug) {
		return nil, false
	}
	return o.ApiSlug, true
}

// HasApiSlug returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasApiSlug() bool {
	if o != nil && !IsNil(o.ApiSlug) {
		return true
	}

	return false
}

// SetApiSlug gets a reference to the given string and assigns it to the ApiSlug field.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetApiSlug(v string) {
	o.ApiSlug = &v
}

// GetIsRequired returns the IsRequired field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsRequired() bool {
	if o == nil || IsNil(o.IsRequired) {
		var ret bool
		return ret
	}
	return *o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRequired) {
		return nil, false
	}
	return o.IsRequired, true
}

// HasIsRequired returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasIsRequired() bool {
	if o != nil && !IsNil(o.IsRequired) {
		return true
	}

	return false
}

// SetIsRequired gets a reference to the given bool and assigns it to the IsRequired field.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetIsRequired(v bool) {
	o.IsRequired = &v
}

// GetIsUnique returns the IsUnique field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsUnique() bool {
	if o == nil || IsNil(o.IsUnique) {
		var ret bool
		return ret
	}
	return *o.IsUnique
}

// GetIsUniqueOk returns a tuple with the IsUnique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsUniqueOk() (*bool, bool) {
	if o == nil || IsNil(o.IsUnique) {
		return nil, false
	}
	return o.IsUnique, true
}

// HasIsUnique returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasIsUnique() bool {
	if o != nil && !IsNil(o.IsUnique) {
		return true
	}

	return false
}

// SetIsUnique gets a reference to the given bool and assigns it to the IsUnique field.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetIsUnique(v bool) {
	o.IsUnique = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetDefaultValue() V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue {
	if o == nil || IsNil(o.DefaultValue) {
		var ret V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetDefaultValueOk() (*V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue, bool) {
	if o == nil || IsNil(o.DefaultValue) {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasDefaultValue() bool {
	if o != nil && !IsNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue and assigns it to the DefaultValue field.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetDefaultValue(v V2TargetIdentifierAttributesAttributePatchRequestDataDefaultValue) {
	o.DefaultValue = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetConfig() V2TargetIdentifierAttributesAttributePatchRequestDataConfig {
	if o == nil || IsNil(o.Config) {
		var ret V2TargetIdentifierAttributesAttributePatchRequestDataConfig
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetConfigOk() (*V2TargetIdentifierAttributesAttributePatchRequestDataConfig, bool) {
	if o == nil || IsNil(o.Config) {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given V2TargetIdentifierAttributesAttributePatchRequestDataConfig and assigns it to the Config field.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetConfig(v V2TargetIdentifierAttributesAttributePatchRequestDataConfig) {
	o.Config = &v
}

// GetIsArchived returns the IsArchived field value if set, zero value otherwise.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsArchived() bool {
	if o == nil || IsNil(o.IsArchived) {
		var ret bool
		return ret
	}
	return *o.IsArchived
}

// GetIsArchivedOk returns a tuple with the IsArchived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) GetIsArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsArchived) {
		return nil, false
	}
	return o.IsArchived, true
}

// HasIsArchived returns a boolean if a field has been set.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) HasIsArchived() bool {
	if o != nil && !IsNil(o.IsArchived) {
		return true
	}

	return false
}

// SetIsArchived gets a reference to the given bool and assigns it to the IsArchived field.
func (o *V2TargetIdentifierAttributesAttributePatchRequestData) SetIsArchived(v bool) {
	o.IsArchived = &v
}

func (o V2TargetIdentifierAttributesAttributePatchRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2TargetIdentifierAttributesAttributePatchRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ApiSlug) {
		toSerialize["api_slug"] = o.ApiSlug
	}
	if !IsNil(o.IsRequired) {
		toSerialize["is_required"] = o.IsRequired
	}
	if !IsNil(o.IsUnique) {
		toSerialize["is_unique"] = o.IsUnique
	}
	if !IsNil(o.DefaultValue) {
		toSerialize["default_value"] = o.DefaultValue
	}
	if !IsNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.IsArchived) {
		toSerialize["is_archived"] = o.IsArchived
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2TargetIdentifierAttributesAttributePatchRequestData) UnmarshalJSON(data []byte) (err error) {
	varV2TargetIdentifierAttributesAttributePatchRequestData := _V2TargetIdentifierAttributesAttributePatchRequestData{}

	err = json.Unmarshal(data, &varV2TargetIdentifierAttributesAttributePatchRequestData)

	if err != nil {
		return err
	}

	*o = V2TargetIdentifierAttributesAttributePatchRequestData(varV2TargetIdentifierAttributesAttributePatchRequestData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "title")
		delete(additionalProperties, "description")
		delete(additionalProperties, "api_slug")
		delete(additionalProperties, "is_required")
		delete(additionalProperties, "is_unique")
		delete(additionalProperties, "default_value")
		delete(additionalProperties, "config")
		delete(additionalProperties, "is_archived")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2TargetIdentifierAttributesAttributePatchRequestData struct {
	value *V2TargetIdentifierAttributesAttributePatchRequestData
	isSet bool
}

func (v NullableV2TargetIdentifierAttributesAttributePatchRequestData) Get() *V2TargetIdentifierAttributesAttributePatchRequestData {
	return v.value
}

func (v *NullableV2TargetIdentifierAttributesAttributePatchRequestData) Set(val *V2TargetIdentifierAttributesAttributePatchRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2TargetIdentifierAttributesAttributePatchRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2TargetIdentifierAttributesAttributePatchRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2TargetIdentifierAttributesAttributePatchRequestData(val *V2TargetIdentifierAttributesAttributePatchRequestData) *NullableV2TargetIdentifierAttributesAttributePatchRequestData {
	return &NullableV2TargetIdentifierAttributesAttributePatchRequestData{value: val, isSet: true}
}

func (v NullableV2TargetIdentifierAttributesAttributePatchRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2TargetIdentifierAttributesAttributePatchRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


