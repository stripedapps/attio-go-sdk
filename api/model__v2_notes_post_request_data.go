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

// checks if the V2NotesPostRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2NotesPostRequestData{}

// V2NotesPostRequestData struct for V2NotesPostRequestData
type V2NotesPostRequestData struct {
	// The ID or slug of the parent object the note belongs to.
	ParentObject string `json:"parent_object"`
	// The ID of the parent record the note belongs to.
	ParentRecordId string `json:"parent_record_id"`
	// The note title. The title is plaintext only and has no formatting.
	Title string `json:"title"`
	// The format of the note content to be created. The `plaintext` format uses the line feed character `\\n` to create new lines within the note content. Rich text formatting, links and @references are not supported.
	Format string `json:"format"`
	// The representation of the note content in the specified format.
	Content string `json:"content"`
	// `created_at` will default to the current time. However, if you wish to backdate a note for migration or other purposes, you can override with a custom `created_at` value. Note that dates before 1970 or in the future are not allowed.
	CreatedAt *string `json:"created_at,omitempty"`
}

type _V2NotesPostRequestData V2NotesPostRequestData

// NewV2NotesPostRequestData instantiates a new V2NotesPostRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2NotesPostRequestData(parentObject string, parentRecordId string, title string, format string, content string) *V2NotesPostRequestData {
	this := V2NotesPostRequestData{}
	this.ParentObject = parentObject
	this.ParentRecordId = parentRecordId
	this.Title = title
	this.Format = format
	this.Content = content
	return &this
}

// NewV2NotesPostRequestDataWithDefaults instantiates a new V2NotesPostRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2NotesPostRequestDataWithDefaults() *V2NotesPostRequestData {
	this := V2NotesPostRequestData{}
	return &this
}

// GetParentObject returns the ParentObject field value
func (o *V2NotesPostRequestData) GetParentObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentObject
}

// GetParentObjectOk returns a tuple with the ParentObject field value
// and a boolean to check if the value has been set.
func (o *V2NotesPostRequestData) GetParentObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentObject, true
}

// SetParentObject sets field value
func (o *V2NotesPostRequestData) SetParentObject(v string) {
	o.ParentObject = v
}

// GetParentRecordId returns the ParentRecordId field value
func (o *V2NotesPostRequestData) GetParentRecordId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParentRecordId
}

// GetParentRecordIdOk returns a tuple with the ParentRecordId field value
// and a boolean to check if the value has been set.
func (o *V2NotesPostRequestData) GetParentRecordIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParentRecordId, true
}

// SetParentRecordId sets field value
func (o *V2NotesPostRequestData) SetParentRecordId(v string) {
	o.ParentRecordId = v
}

// GetTitle returns the Title field value
func (o *V2NotesPostRequestData) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *V2NotesPostRequestData) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *V2NotesPostRequestData) SetTitle(v string) {
	o.Title = v
}

// GetFormat returns the Format field value
func (o *V2NotesPostRequestData) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *V2NotesPostRequestData) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *V2NotesPostRequestData) SetFormat(v string) {
	o.Format = v
}

// GetContent returns the Content field value
func (o *V2NotesPostRequestData) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *V2NotesPostRequestData) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *V2NotesPostRequestData) SetContent(v string) {
	o.Content = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *V2NotesPostRequestData) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2NotesPostRequestData) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *V2NotesPostRequestData) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *V2NotesPostRequestData) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

func (o V2NotesPostRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2NotesPostRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["parent_object"] = o.ParentObject
	toSerialize["parent_record_id"] = o.ParentRecordId
	toSerialize["title"] = o.Title
	toSerialize["format"] = o.Format
	toSerialize["content"] = o.Content
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

func (o *V2NotesPostRequestData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"parent_object",
		"parent_record_id",
		"title",
		"format",
		"content",
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

	varV2NotesPostRequestData := _V2NotesPostRequestData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV2NotesPostRequestData)

	if err != nil {
		return err
	}

	*o = V2NotesPostRequestData(varV2NotesPostRequestData)

	return err
}

type NullableV2NotesPostRequestData struct {
	value *V2NotesPostRequestData
	isSet bool
}

func (v NullableV2NotesPostRequestData) Get() *V2NotesPostRequestData {
	return v.value
}

func (v *NullableV2NotesPostRequestData) Set(val *V2NotesPostRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2NotesPostRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2NotesPostRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2NotesPostRequestData(val *V2NotesPostRequestData) *NullableV2NotesPostRequestData {
	return &NullableV2NotesPostRequestData{value: val, isSet: true}
}

func (v NullableV2NotesPostRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2NotesPostRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


