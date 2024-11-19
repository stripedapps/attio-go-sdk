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
	"fmt"
)

// checks if the V2CommentsPostRequestDataAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2CommentsPostRequestDataAnyOf1{}

// V2CommentsPostRequestDataAnyOf1 struct for V2CommentsPostRequestDataAnyOf1
type V2CommentsPostRequestDataAnyOf1 struct {
	// The format that the comment content is provided in. The `plaintext` format uses the line feed character `\\n` to create new lines within the note content. Rich text formatting and links are not supported.
	Format string `json:"format"`
	// The content of the comment itself. Workspace members can be mentioned using their email address, otherwise email addresses will be presented to users as clickable mailto links.
	Content string `json:"content"`
	Author V2CommentsPostRequestDataAnyOfAuthor `json:"author"`
	// `created_at` will default to the current time. However, if you wish to backdate a comment for migration or other purposes, you can override with a custom `created_at` value. Note that dates before 1970 or in the future are not allowed.
	CreatedAt *string `json:"created_at,omitempty"`
	Record V2CommentsPostRequestDataAnyOf1Record `json:"record"`
	AdditionalProperties map[string]interface{}
}

type _V2CommentsPostRequestDataAnyOf1 V2CommentsPostRequestDataAnyOf1

// NewV2CommentsPostRequestDataAnyOf1 instantiates a new V2CommentsPostRequestDataAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2CommentsPostRequestDataAnyOf1(format string, content string, author V2CommentsPostRequestDataAnyOfAuthor, record V2CommentsPostRequestDataAnyOf1Record) *V2CommentsPostRequestDataAnyOf1 {
	this := V2CommentsPostRequestDataAnyOf1{}
	this.Format = format
	this.Content = content
	this.Author = author
	this.Record = record
	return &this
}

// NewV2CommentsPostRequestDataAnyOf1WithDefaults instantiates a new V2CommentsPostRequestDataAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2CommentsPostRequestDataAnyOf1WithDefaults() *V2CommentsPostRequestDataAnyOf1 {
	this := V2CommentsPostRequestDataAnyOf1{}
	return &this
}

// GetFormat returns the Format field value
func (o *V2CommentsPostRequestDataAnyOf1) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf1) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *V2CommentsPostRequestDataAnyOf1) SetFormat(v string) {
	o.Format = v
}

// GetContent returns the Content field value
func (o *V2CommentsPostRequestDataAnyOf1) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf1) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *V2CommentsPostRequestDataAnyOf1) SetContent(v string) {
	o.Content = v
}

// GetAuthor returns the Author field value
func (o *V2CommentsPostRequestDataAnyOf1) GetAuthor() V2CommentsPostRequestDataAnyOfAuthor {
	if o == nil {
		var ret V2CommentsPostRequestDataAnyOfAuthor
		return ret
	}

	return o.Author
}

// GetAuthorOk returns a tuple with the Author field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf1) GetAuthorOk() (*V2CommentsPostRequestDataAnyOfAuthor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Author, true
}

// SetAuthor sets field value
func (o *V2CommentsPostRequestDataAnyOf1) SetAuthor(v V2CommentsPostRequestDataAnyOfAuthor) {
	o.Author = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *V2CommentsPostRequestDataAnyOf1) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf1) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *V2CommentsPostRequestDataAnyOf1) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *V2CommentsPostRequestDataAnyOf1) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetRecord returns the Record field value
func (o *V2CommentsPostRequestDataAnyOf1) GetRecord() V2CommentsPostRequestDataAnyOf1Record {
	if o == nil {
		var ret V2CommentsPostRequestDataAnyOf1Record
		return ret
	}

	return o.Record
}

// GetRecordOk returns a tuple with the Record field value
// and a boolean to check if the value has been set.
func (o *V2CommentsPostRequestDataAnyOf1) GetRecordOk() (*V2CommentsPostRequestDataAnyOf1Record, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Record, true
}

// SetRecord sets field value
func (o *V2CommentsPostRequestDataAnyOf1) SetRecord(v V2CommentsPostRequestDataAnyOf1Record) {
	o.Record = v
}

func (o V2CommentsPostRequestDataAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2CommentsPostRequestDataAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["format"] = o.Format
	toSerialize["content"] = o.Content
	toSerialize["author"] = o.Author
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	toSerialize["record"] = o.Record

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2CommentsPostRequestDataAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"format",
		"content",
		"author",
		"record",
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

	varV2CommentsPostRequestDataAnyOf1 := _V2CommentsPostRequestDataAnyOf1{}

	err = json.Unmarshal(data, &varV2CommentsPostRequestDataAnyOf1)

	if err != nil {
		return err
	}

	*o = V2CommentsPostRequestDataAnyOf1(varV2CommentsPostRequestDataAnyOf1)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "format")
		delete(additionalProperties, "content")
		delete(additionalProperties, "author")
		delete(additionalProperties, "created_at")
		delete(additionalProperties, "record")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2CommentsPostRequestDataAnyOf1 struct {
	value *V2CommentsPostRequestDataAnyOf1
	isSet bool
}

func (v NullableV2CommentsPostRequestDataAnyOf1) Get() *V2CommentsPostRequestDataAnyOf1 {
	return v.value
}

func (v *NullableV2CommentsPostRequestDataAnyOf1) Set(val *V2CommentsPostRequestDataAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableV2CommentsPostRequestDataAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableV2CommentsPostRequestDataAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2CommentsPostRequestDataAnyOf1(val *V2CommentsPostRequestDataAnyOf1) *NullableV2CommentsPostRequestDataAnyOf1 {
	return &NullableV2CommentsPostRequestDataAnyOf1{value: val, isSet: true}
}

func (v NullableV2CommentsPostRequestDataAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2CommentsPostRequestDataAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

