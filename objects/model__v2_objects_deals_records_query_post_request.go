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

// checks if the V2ObjectsDealsRecordsQueryPostRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsDealsRecordsQueryPostRequest{}

// V2ObjectsDealsRecordsQueryPostRequest struct for V2ObjectsDealsRecordsQueryPostRequest
type V2ObjectsDealsRecordsQueryPostRequest struct {
	// An object used to filter results to a subset of records. See the [full guide to filtering and sorting here](/docs/filtering-and-sorting).
	Filter map[string]interface{} `json:"filter,omitempty"`
	// An object used to sort results. See the [full guide to filtering and sorting here](/docs/filtering-and-sorting).
	Sorts []V2ObjectsPeopleRecordsQueryPostRequestSortsInner `json:"sorts,omitempty"`
	// The maximum number of results to return. Defaults to 500. See the [full guide to pagination here](/docs/pagination).
	Limit *float32 `json:"limit,omitempty"`
	// The number of results to skip over before returning. Defaults to 0. See the [full guide to pagination here](/docs/pagination).
	Offset *float32 `json:"offset,omitempty"`
}

// NewV2ObjectsDealsRecordsQueryPostRequest instantiates a new V2ObjectsDealsRecordsQueryPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsDealsRecordsQueryPostRequest() *V2ObjectsDealsRecordsQueryPostRequest {
	this := V2ObjectsDealsRecordsQueryPostRequest{}
	return &this
}

// NewV2ObjectsDealsRecordsQueryPostRequestWithDefaults instantiates a new V2ObjectsDealsRecordsQueryPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsDealsRecordsQueryPostRequestWithDefaults() *V2ObjectsDealsRecordsQueryPostRequest {
	this := V2ObjectsDealsRecordsQueryPostRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsQueryPostRequest) GetFilter() map[string]interface{} {
	if o == nil || IsNil(o.Filter) {
		var ret map[string]interface{}
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsQueryPostRequest) GetFilterOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Filter) {
		return map[string]interface{}{}, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsQueryPostRequest) HasFilter() bool {
	if o != nil && !IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given map[string]interface{} and assigns it to the Filter field.
func (o *V2ObjectsDealsRecordsQueryPostRequest) SetFilter(v map[string]interface{}) {
	o.Filter = v
}

// GetSorts returns the Sorts field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsQueryPostRequest) GetSorts() []V2ObjectsPeopleRecordsQueryPostRequestSortsInner {
	if o == nil || IsNil(o.Sorts) {
		var ret []V2ObjectsPeopleRecordsQueryPostRequestSortsInner
		return ret
	}
	return o.Sorts
}

// GetSortsOk returns a tuple with the Sorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsQueryPostRequest) GetSortsOk() ([]V2ObjectsPeopleRecordsQueryPostRequestSortsInner, bool) {
	if o == nil || IsNil(o.Sorts) {
		return nil, false
	}
	return o.Sorts, true
}

// HasSorts returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsQueryPostRequest) HasSorts() bool {
	if o != nil && !IsNil(o.Sorts) {
		return true
	}

	return false
}

// SetSorts gets a reference to the given []V2ObjectsPeopleRecordsQueryPostRequestSortsInner and assigns it to the Sorts field.
func (o *V2ObjectsDealsRecordsQueryPostRequest) SetSorts(v []V2ObjectsPeopleRecordsQueryPostRequestSortsInner) {
	o.Sorts = v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsQueryPostRequest) GetLimit() float32 {
	if o == nil || IsNil(o.Limit) {
		var ret float32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsQueryPostRequest) GetLimitOk() (*float32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsQueryPostRequest) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given float32 and assigns it to the Limit field.
func (o *V2ObjectsDealsRecordsQueryPostRequest) SetLimit(v float32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *V2ObjectsDealsRecordsQueryPostRequest) GetOffset() float32 {
	if o == nil || IsNil(o.Offset) {
		var ret float32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsDealsRecordsQueryPostRequest) GetOffsetOk() (*float32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *V2ObjectsDealsRecordsQueryPostRequest) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given float32 and assigns it to the Offset field.
func (o *V2ObjectsDealsRecordsQueryPostRequest) SetOffset(v float32) {
	o.Offset = &v
}

func (o V2ObjectsDealsRecordsQueryPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsDealsRecordsQueryPostRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filter) {
		toSerialize["filter"] = o.Filter
	}
	if !IsNil(o.Sorts) {
		toSerialize["sorts"] = o.Sorts
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableV2ObjectsDealsRecordsQueryPostRequest struct {
	value *V2ObjectsDealsRecordsQueryPostRequest
	isSet bool
}

func (v NullableV2ObjectsDealsRecordsQueryPostRequest) Get() *V2ObjectsDealsRecordsQueryPostRequest {
	return v.value
}

func (v *NullableV2ObjectsDealsRecordsQueryPostRequest) Set(val *V2ObjectsDealsRecordsQueryPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsDealsRecordsQueryPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsDealsRecordsQueryPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsDealsRecordsQueryPostRequest(val *V2ObjectsDealsRecordsQueryPostRequest) *NullableV2ObjectsDealsRecordsQueryPostRequest {
	return &NullableV2ObjectsDealsRecordsQueryPostRequest{value: val, isSet: true}
}

func (v NullableV2ObjectsDealsRecordsQueryPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsDealsRecordsQueryPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


