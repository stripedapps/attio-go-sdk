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

// checks if the List type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &List{}

// List struct for List
type List struct {
	Id ListId `json:"id"`
	// A human-readable slug for use in URLs and responses.
	ApiSlug string `json:"api_slug"`
	// The name of the list, as viewed in the UI.
	Name string `json:"name"`
	// A UUID or slug to identify the allowed object type for records added to this list. All new Lists are expected to have exactly one parent object. However, some legacy lists may have multiple allowed parents so the return type of this field is an array.
	ParentObject []string `json:"parent_object"`
	// The level of access granted to all members of the workspace for this list. `null` values represent a private list that only grants access to specific workspace members via the `workspace_member_access` property.
	WorkspaceAccess NullableString `json:"workspace_access"`
	// The level of access granted to specific workspace members for this list. An empty array represents a list that has granted access to no workspace members.
	WorkspaceMemberAccess []V2ListsPostRequestDataWorkspaceMemberAccessInner `json:"workspace_member_access"`
	CreatedByActor ListCreatedByActor `json:"created_by_actor"`
	// When the list was created.
	CreatedAt string `json:"created_at"`
}

type _List List

// NewList instantiates a new List object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewList(id ListId, apiSlug string, name string, parentObject []string, workspaceAccess NullableString, workspaceMemberAccess []V2ListsPostRequestDataWorkspaceMemberAccessInner, createdByActor ListCreatedByActor, createdAt string) *List {
	this := List{}
	this.Id = id
	this.ApiSlug = apiSlug
	this.Name = name
	this.ParentObject = parentObject
	this.WorkspaceAccess = workspaceAccess
	this.WorkspaceMemberAccess = workspaceMemberAccess
	this.CreatedByActor = createdByActor
	this.CreatedAt = createdAt
	return &this
}

// NewListWithDefaults instantiates a new List object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWithDefaults() *List {
	this := List{}
	return &this
}

// GetId returns the Id field value
func (o *List) GetId() ListId {
	if o == nil {
		var ret ListId
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *List) GetIdOk() (*ListId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *List) SetId(v ListId) {
	o.Id = v
}

// GetApiSlug returns the ApiSlug field value
func (o *List) GetApiSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiSlug
}

// GetApiSlugOk returns a tuple with the ApiSlug field value
// and a boolean to check if the value has been set.
func (o *List) GetApiSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiSlug, true
}

// SetApiSlug sets field value
func (o *List) SetApiSlug(v string) {
	o.ApiSlug = v
}

// GetName returns the Name field value
func (o *List) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *List) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *List) SetName(v string) {
	o.Name = v
}

// GetParentObject returns the ParentObject field value
func (o *List) GetParentObject() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ParentObject
}

// GetParentObjectOk returns a tuple with the ParentObject field value
// and a boolean to check if the value has been set.
func (o *List) GetParentObjectOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ParentObject, true
}

// SetParentObject sets field value
func (o *List) SetParentObject(v []string) {
	o.ParentObject = v
}

// GetWorkspaceAccess returns the WorkspaceAccess field value
// If the value is explicit nil, the zero value for string will be returned
func (o *List) GetWorkspaceAccess() string {
	if o == nil || o.WorkspaceAccess.Get() == nil {
		var ret string
		return ret
	}

	return *o.WorkspaceAccess.Get()
}

// GetWorkspaceAccessOk returns a tuple with the WorkspaceAccess field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *List) GetWorkspaceAccessOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceAccess.Get(), o.WorkspaceAccess.IsSet()
}

// SetWorkspaceAccess sets field value
func (o *List) SetWorkspaceAccess(v string) {
	o.WorkspaceAccess.Set(&v)
}

// GetWorkspaceMemberAccess returns the WorkspaceMemberAccess field value
func (o *List) GetWorkspaceMemberAccess() []V2ListsPostRequestDataWorkspaceMemberAccessInner {
	if o == nil {
		var ret []V2ListsPostRequestDataWorkspaceMemberAccessInner
		return ret
	}

	return o.WorkspaceMemberAccess
}

// GetWorkspaceMemberAccessOk returns a tuple with the WorkspaceMemberAccess field value
// and a boolean to check if the value has been set.
func (o *List) GetWorkspaceMemberAccessOk() ([]V2ListsPostRequestDataWorkspaceMemberAccessInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceMemberAccess, true
}

// SetWorkspaceMemberAccess sets field value
func (o *List) SetWorkspaceMemberAccess(v []V2ListsPostRequestDataWorkspaceMemberAccessInner) {
	o.WorkspaceMemberAccess = v
}

// GetCreatedByActor returns the CreatedByActor field value
func (o *List) GetCreatedByActor() ListCreatedByActor {
	if o == nil {
		var ret ListCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *List) GetCreatedByActorOk() (*ListCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *List) SetCreatedByActor(v ListCreatedByActor) {
	o.CreatedByActor = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *List) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *List) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *List) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o List) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o List) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["api_slug"] = o.ApiSlug
	toSerialize["name"] = o.Name
	toSerialize["parent_object"] = o.ParentObject
	toSerialize["workspace_access"] = o.WorkspaceAccess.Get()
	toSerialize["workspace_member_access"] = o.WorkspaceMemberAccess
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *List) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"api_slug",
		"name",
		"parent_object",
		"workspace_access",
		"workspace_member_access",
		"created_by_actor",
		"created_at",
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

	varList := _List{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varList)

	if err != nil {
		return err
	}

	*o = List(varList)

	return err
}

type NullableList struct {
	value *List
	isSet bool
}

func (v NullableList) Get() *List {
	return v.value
}

func (v *NullableList) Set(val *List) {
	v.value = val
	v.isSet = true
}

func (v NullableList) IsSet() bool {
	return v.isSet
}

func (v *NullableList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableList(val *List) *NullableList {
	return &NullableList{value: val, isSet: true}
}

func (v NullableList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


