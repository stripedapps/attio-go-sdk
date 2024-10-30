/*
Attio Standard Objects

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V2ObjectsWorkspacesRecordsPutRequestDataValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsWorkspacesRecordsPutRequestDataValues{}

// V2ObjectsWorkspacesRecordsPutRequestDataValues This object's keys should be the slugs or IDs of the attributes you wish to update. Below, you'll find documentation for the value types of each standard workspace attribute. For information on potential custom attributes, refer to our [attribute type docs](/docs/attribute-types).
type V2ObjectsWorkspacesRecordsPutRequestDataValues struct {
	WorkspaceId []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner `json:"workspace_id,omitempty"`
	Name []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner `json:"name,omitempty"`
	Users []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner `json:"users,omitempty"`
	Company []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner `json:"company,omitempty"`
	AvatarUrl []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner `json:"avatar_url,omitempty"`
}

// NewV2ObjectsWorkspacesRecordsPutRequestDataValues instantiates a new V2ObjectsWorkspacesRecordsPutRequestDataValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsWorkspacesRecordsPutRequestDataValues() *V2ObjectsWorkspacesRecordsPutRequestDataValues {
	this := V2ObjectsWorkspacesRecordsPutRequestDataValues{}
	return &this
}

// NewV2ObjectsWorkspacesRecordsPutRequestDataValuesWithDefaults instantiates a new V2ObjectsWorkspacesRecordsPutRequestDataValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsWorkspacesRecordsPutRequestDataValuesWithDefaults() *V2ObjectsWorkspacesRecordsPutRequestDataValues {
	this := V2ObjectsWorkspacesRecordsPutRequestDataValues{}
	return &this
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetWorkspaceId() []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner {
	if o == nil || IsNil(o.WorkspaceId) {
		var ret []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner
		return ret
	}
	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetWorkspaceIdOk() ([]V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.WorkspaceId) {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) HasWorkspaceId() bool {
	if o != nil && !IsNil(o.WorkspaceId) {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner and assigns it to the WorkspaceId field.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) SetWorkspaceId(v []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) {
	o.WorkspaceId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetName() []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner {
	if o == nil || IsNil(o.Name) {
		var ret []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetNameOk() ([]V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner and assigns it to the Name field.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) SetName(v []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) {
	o.Name = v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetUsers() []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner {
	if o == nil || IsNil(o.Users) {
		var ret []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetUsersOk() ([]V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner and assigns it to the Users field.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) SetUsers(v []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner) {
	o.Users = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetCompany() []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner {
	if o == nil || IsNil(o.Company) {
		var ret []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner
		return ret
	}
	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetCompanyOk() ([]V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner and assigns it to the Company field.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) SetCompany(v []V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner) {
	o.Company = v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetAvatarUrl() []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner
		return ret
	}
	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) GetAvatarUrlOk() ([]V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner and assigns it to the AvatarUrl field.
func (o *V2ObjectsWorkspacesRecordsPutRequestDataValues) SetAvatarUrl(v []V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner) {
	o.AvatarUrl = v
}

func (o V2ObjectsWorkspacesRecordsPutRequestDataValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsWorkspacesRecordsPutRequestDataValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WorkspaceId) {
		toSerialize["workspace_id"] = o.WorkspaceId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	return toSerialize, nil
}

type NullableV2ObjectsWorkspacesRecordsPutRequestDataValues struct {
	value *V2ObjectsWorkspacesRecordsPutRequestDataValues
	isSet bool
}

func (v NullableV2ObjectsWorkspacesRecordsPutRequestDataValues) Get() *V2ObjectsWorkspacesRecordsPutRequestDataValues {
	return v.value
}

func (v *NullableV2ObjectsWorkspacesRecordsPutRequestDataValues) Set(val *V2ObjectsWorkspacesRecordsPutRequestDataValues) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsWorkspacesRecordsPutRequestDataValues) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsWorkspacesRecordsPutRequestDataValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsWorkspacesRecordsPutRequestDataValues(val *V2ObjectsWorkspacesRecordsPutRequestDataValues) *NullableV2ObjectsWorkspacesRecordsPutRequestDataValues {
	return &NullableV2ObjectsWorkspacesRecordsPutRequestDataValues{value: val, isSet: true}
}

func (v NullableV2ObjectsWorkspacesRecordsPutRequestDataValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsWorkspacesRecordsPutRequestDataValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


