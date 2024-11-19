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

// checks if the V2ListsListPatchRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ListsListPatchRequestData{}

// V2ListsListPatchRequestData struct for V2ListsListPatchRequestData
type V2ListsListPatchRequestData struct {
	// The human-readable name of the list.
	Name *string `json:"name,omitempty"`
	// A unique, human-readable slug to access the list through API calls. Should be formatted in snake case.
	ApiSlug *string `json:"api_slug,omitempty"`
	// The level of access granted to all members of the workspace for this list. Pass `null` to keep the list private and only grant access to specific workspace members.
	WorkspaceAccess NullableString `json:"workspace_access,omitempty"`
	// The level of access granted to specific workspace members for this list. Pass an empty array to grant access to no workspace members.
	WorkspaceMemberAccess []V2ListsPostRequestDataWorkspaceMemberAccessInner `json:"workspace_member_access,omitempty"`
}

// NewV2ListsListPatchRequestData instantiates a new V2ListsListPatchRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ListsListPatchRequestData() *V2ListsListPatchRequestData {
	this := V2ListsListPatchRequestData{}
	return &this
}

// NewV2ListsListPatchRequestDataWithDefaults instantiates a new V2ListsListPatchRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ListsListPatchRequestDataWithDefaults() *V2ListsListPatchRequestData {
	this := V2ListsListPatchRequestData{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2ListsListPatchRequestData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ListsListPatchRequestData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2ListsListPatchRequestData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V2ListsListPatchRequestData) SetName(v string) {
	o.Name = &v
}

// GetApiSlug returns the ApiSlug field value if set, zero value otherwise.
func (o *V2ListsListPatchRequestData) GetApiSlug() string {
	if o == nil || IsNil(o.ApiSlug) {
		var ret string
		return ret
	}
	return *o.ApiSlug
}

// GetApiSlugOk returns a tuple with the ApiSlug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ListsListPatchRequestData) GetApiSlugOk() (*string, bool) {
	if o == nil || IsNil(o.ApiSlug) {
		return nil, false
	}
	return o.ApiSlug, true
}

// HasApiSlug returns a boolean if a field has been set.
func (o *V2ListsListPatchRequestData) HasApiSlug() bool {
	if o != nil && !IsNil(o.ApiSlug) {
		return true
	}

	return false
}

// SetApiSlug gets a reference to the given string and assigns it to the ApiSlug field.
func (o *V2ListsListPatchRequestData) SetApiSlug(v string) {
	o.ApiSlug = &v
}

// GetWorkspaceAccess returns the WorkspaceAccess field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *V2ListsListPatchRequestData) GetWorkspaceAccess() string {
	if o == nil || IsNil(o.WorkspaceAccess.Get()) {
		var ret string
		return ret
	}
	return *o.WorkspaceAccess.Get()
}

// GetWorkspaceAccessOk returns a tuple with the WorkspaceAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ListsListPatchRequestData) GetWorkspaceAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceAccess.Get(), o.WorkspaceAccess.IsSet()
}

// HasWorkspaceAccess returns a boolean if a field has been set.
func (o *V2ListsListPatchRequestData) HasWorkspaceAccess() bool {
	if o != nil && o.WorkspaceAccess.IsSet() {
		return true
	}

	return false
}

// SetWorkspaceAccess gets a reference to the given NullableString and assigns it to the WorkspaceAccess field.
func (o *V2ListsListPatchRequestData) SetWorkspaceAccess(v string) {
	o.WorkspaceAccess.Set(&v)
}
// SetWorkspaceAccessNil sets the value for WorkspaceAccess to be an explicit nil
func (o *V2ListsListPatchRequestData) SetWorkspaceAccessNil() {
	o.WorkspaceAccess.Set(nil)
}

// UnsetWorkspaceAccess ensures that no value is present for WorkspaceAccess, not even an explicit nil
func (o *V2ListsListPatchRequestData) UnsetWorkspaceAccess() {
	o.WorkspaceAccess.Unset()
}

// GetWorkspaceMemberAccess returns the WorkspaceMemberAccess field value if set, zero value otherwise.
func (o *V2ListsListPatchRequestData) GetWorkspaceMemberAccess() []V2ListsPostRequestDataWorkspaceMemberAccessInner {
	if o == nil || IsNil(o.WorkspaceMemberAccess) {
		var ret []V2ListsPostRequestDataWorkspaceMemberAccessInner
		return ret
	}
	return o.WorkspaceMemberAccess
}

// GetWorkspaceMemberAccessOk returns a tuple with the WorkspaceMemberAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ListsListPatchRequestData) GetWorkspaceMemberAccessOk() ([]V2ListsPostRequestDataWorkspaceMemberAccessInner, bool) {
	if o == nil || IsNil(o.WorkspaceMemberAccess) {
		return nil, false
	}
	return o.WorkspaceMemberAccess, true
}

// HasWorkspaceMemberAccess returns a boolean if a field has been set.
func (o *V2ListsListPatchRequestData) HasWorkspaceMemberAccess() bool {
	if o != nil && !IsNil(o.WorkspaceMemberAccess) {
		return true
	}

	return false
}

// SetWorkspaceMemberAccess gets a reference to the given []V2ListsPostRequestDataWorkspaceMemberAccessInner and assigns it to the WorkspaceMemberAccess field.
func (o *V2ListsListPatchRequestData) SetWorkspaceMemberAccess(v []V2ListsPostRequestDataWorkspaceMemberAccessInner) {
	o.WorkspaceMemberAccess = v
}

func (o V2ListsListPatchRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ListsListPatchRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ApiSlug) {
		toSerialize["api_slug"] = o.ApiSlug
	}
	if o.WorkspaceAccess.IsSet() {
		toSerialize["workspace_access"] = o.WorkspaceAccess.Get()
	}
	if !IsNil(o.WorkspaceMemberAccess) {
		toSerialize["workspace_member_access"] = o.WorkspaceMemberAccess
	}
	return toSerialize, nil
}

type NullableV2ListsListPatchRequestData struct {
	value *V2ListsListPatchRequestData
	isSet bool
}

func (v NullableV2ListsListPatchRequestData) Get() *V2ListsListPatchRequestData {
	return v.value
}

func (v *NullableV2ListsListPatchRequestData) Set(val *V2ListsListPatchRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ListsListPatchRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ListsListPatchRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ListsListPatchRequestData(val *V2ListsListPatchRequestData) *NullableV2ListsListPatchRequestData {
	return &NullableV2ListsListPatchRequestData{value: val, isSet: true}
}

func (v NullableV2ListsListPatchRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ListsListPatchRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

