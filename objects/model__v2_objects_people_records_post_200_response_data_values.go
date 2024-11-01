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

// checks if the V2ObjectsPeopleRecordsPost200ResponseDataValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsPost200ResponseDataValues{}

// V2ObjectsPeopleRecordsPost200ResponseDataValues An object with `attribute_slug` keys, and an array of value objects as the values. Attributes slugs (for example `email_addresses` or `name`) can be used, including custom attribute slugs.
type V2ObjectsPeopleRecordsPost200ResponseDataValues struct {
	EmailAddresses []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner `json:"email_addresses,omitempty"`
	Name []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner `json:"name,omitempty"`
	Description []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner `json:"description,omitempty"`
	AvatarUrl []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner `json:"avatar_url,omitempty"`
	JobTitle []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner `json:"job_title,omitempty"`
	PhoneNumbers []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner `json:"phone_numbers,omitempty"`
	Linkedin []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner `json:"linkedin,omitempty"`
	Twitter []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner `json:"twitter,omitempty"`
	Facebook []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner `json:"facebook,omitempty"`
	Instagram []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner `json:"instagram,omitempty"`
	Angellist []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner `json:"angellist,omitempty"`
	PrimaryLocation []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner `json:"primary_location,omitempty"`
	TwitterFollowerCount []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner `json:"twitter_follower_count,omitempty"`
	Company []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner `json:"company,omitempty"`
}

// NewV2ObjectsPeopleRecordsPost200ResponseDataValues instantiates a new V2ObjectsPeopleRecordsPost200ResponseDataValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsPost200ResponseDataValues() *V2ObjectsPeopleRecordsPost200ResponseDataValues {
	this := V2ObjectsPeopleRecordsPost200ResponseDataValues{}
	return &this
}

// NewV2ObjectsPeopleRecordsPost200ResponseDataValuesWithDefaults instantiates a new V2ObjectsPeopleRecordsPost200ResponseDataValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsPost200ResponseDataValuesWithDefaults() *V2ObjectsPeopleRecordsPost200ResponseDataValues {
	this := V2ObjectsPeopleRecordsPost200ResponseDataValues{}
	return &this
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetEmailAddresses() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner {
	if o == nil || IsNil(o.EmailAddresses) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner
		return ret
	}
	return o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetEmailAddressesOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner, bool) {
	if o == nil || IsNil(o.EmailAddresses) {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasEmailAddresses() bool {
	if o != nil && !IsNil(o.EmailAddresses) {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner and assigns it to the EmailAddresses field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetEmailAddresses(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) {
	o.EmailAddresses = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetName() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner {
	if o == nil || IsNil(o.Name) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetNameOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner and assigns it to the Name field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetName(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetDescription() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Description) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetDescriptionOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Description field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetDescription(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Description = v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetAvatarUrl() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetAvatarUrlOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the AvatarUrl field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetAvatarUrl(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.AvatarUrl = v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetJobTitle() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.JobTitle) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetJobTitleOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.JobTitle) {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasJobTitle() bool {
	if o != nil && !IsNil(o.JobTitle) {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the JobTitle field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetJobTitle(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.JobTitle = v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetPhoneNumbers() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner {
	if o == nil || IsNil(o.PhoneNumbers) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner
		return ret
	}
	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetPhoneNumbersOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner, bool) {
	if o == nil || IsNil(o.PhoneNumbers) {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasPhoneNumbers() bool {
	if o != nil && !IsNil(o.PhoneNumbers) {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner and assigns it to the PhoneNumbers field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetPhoneNumbers(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner) {
	o.PhoneNumbers = v
}

// GetLinkedin returns the Linkedin field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetLinkedin() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Linkedin) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Linkedin
}

// GetLinkedinOk returns a tuple with the Linkedin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetLinkedinOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Linkedin) {
		return nil, false
	}
	return o.Linkedin, true
}

// HasLinkedin returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasLinkedin() bool {
	if o != nil && !IsNil(o.Linkedin) {
		return true
	}

	return false
}

// SetLinkedin gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Linkedin field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetLinkedin(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Linkedin = v
}

// GetTwitter returns the Twitter field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetTwitter() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Twitter) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetTwitterOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Twitter) {
		return nil, false
	}
	return o.Twitter, true
}

// HasTwitter returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasTwitter() bool {
	if o != nil && !IsNil(o.Twitter) {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Twitter field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetTwitter(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Twitter = v
}

// GetFacebook returns the Facebook field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetFacebook() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Facebook) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Facebook
}

// GetFacebookOk returns a tuple with the Facebook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetFacebookOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Facebook) {
		return nil, false
	}
	return o.Facebook, true
}

// HasFacebook returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasFacebook() bool {
	if o != nil && !IsNil(o.Facebook) {
		return true
	}

	return false
}

// SetFacebook gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Facebook field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetFacebook(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Facebook = v
}

// GetInstagram returns the Instagram field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetInstagram() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Instagram) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Instagram
}

// GetInstagramOk returns a tuple with the Instagram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetInstagramOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Instagram) {
		return nil, false
	}
	return o.Instagram, true
}

// HasInstagram returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasInstagram() bool {
	if o != nil && !IsNil(o.Instagram) {
		return true
	}

	return false
}

// SetInstagram gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Instagram field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetInstagram(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Instagram = v
}

// GetAngellist returns the Angellist field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetAngellist() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Angellist) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Angellist
}

// GetAngellistOk returns a tuple with the Angellist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetAngellistOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Angellist) {
		return nil, false
	}
	return o.Angellist, true
}

// HasAngellist returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasAngellist() bool {
	if o != nil && !IsNil(o.Angellist) {
		return true
	}

	return false
}

// SetAngellist gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Angellist field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetAngellist(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Angellist = v
}

// GetPrimaryLocation returns the PrimaryLocation field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetPrimaryLocation() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner {
	if o == nil || IsNil(o.PrimaryLocation) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner
		return ret
	}
	return o.PrimaryLocation
}

// GetPrimaryLocationOk returns a tuple with the PrimaryLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetPrimaryLocationOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner, bool) {
	if o == nil || IsNil(o.PrimaryLocation) {
		return nil, false
	}
	return o.PrimaryLocation, true
}

// HasPrimaryLocation returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasPrimaryLocation() bool {
	if o != nil && !IsNil(o.PrimaryLocation) {
		return true
	}

	return false
}

// SetPrimaryLocation gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner and assigns it to the PrimaryLocation field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetPrimaryLocation(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) {
	o.PrimaryLocation = v
}

// GetTwitterFollowerCount returns the TwitterFollowerCount field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetTwitterFollowerCount() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner {
	if o == nil || IsNil(o.TwitterFollowerCount) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner
		return ret
	}
	return o.TwitterFollowerCount
}

// GetTwitterFollowerCountOk returns a tuple with the TwitterFollowerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetTwitterFollowerCountOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner, bool) {
	if o == nil || IsNil(o.TwitterFollowerCount) {
		return nil, false
	}
	return o.TwitterFollowerCount, true
}

// HasTwitterFollowerCount returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasTwitterFollowerCount() bool {
	if o != nil && !IsNil(o.TwitterFollowerCount) {
		return true
	}

	return false
}

// SetTwitterFollowerCount gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner and assigns it to the TwitterFollowerCount field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetTwitterFollowerCount(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner) {
	o.TwitterFollowerCount = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetCompany() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner {
	if o == nil || IsNil(o.Company) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner
		return ret
	}
	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) GetCompanyOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner and assigns it to the Company field.
func (o *V2ObjectsPeopleRecordsPost200ResponseDataValues) SetCompany(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) {
	o.Company = v
}

func (o V2ObjectsPeopleRecordsPost200ResponseDataValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsPost200ResponseDataValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EmailAddresses) {
		toSerialize["email_addresses"] = o.EmailAddresses
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.AvatarUrl) {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if !IsNil(o.JobTitle) {
		toSerialize["job_title"] = o.JobTitle
	}
	if !IsNil(o.PhoneNumbers) {
		toSerialize["phone_numbers"] = o.PhoneNumbers
	}
	if !IsNil(o.Linkedin) {
		toSerialize["linkedin"] = o.Linkedin
	}
	if !IsNil(o.Twitter) {
		toSerialize["twitter"] = o.Twitter
	}
	if !IsNil(o.Facebook) {
		toSerialize["facebook"] = o.Facebook
	}
	if !IsNil(o.Instagram) {
		toSerialize["instagram"] = o.Instagram
	}
	if !IsNil(o.Angellist) {
		toSerialize["angellist"] = o.Angellist
	}
	if !IsNil(o.PrimaryLocation) {
		toSerialize["primary_location"] = o.PrimaryLocation
	}
	if !IsNil(o.TwitterFollowerCount) {
		toSerialize["twitter_follower_count"] = o.TwitterFollowerCount
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	return toSerialize, nil
}

type NullableV2ObjectsPeopleRecordsPost200ResponseDataValues struct {
	value *V2ObjectsPeopleRecordsPost200ResponseDataValues
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsPost200ResponseDataValues) Get() *V2ObjectsPeopleRecordsPost200ResponseDataValues {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsPost200ResponseDataValues) Set(val *V2ObjectsPeopleRecordsPost200ResponseDataValues) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsPost200ResponseDataValues) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsPost200ResponseDataValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsPost200ResponseDataValues(val *V2ObjectsPeopleRecordsPost200ResponseDataValues) *NullableV2ObjectsPeopleRecordsPost200ResponseDataValues {
	return &NullableV2ObjectsPeopleRecordsPost200ResponseDataValues{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsPost200ResponseDataValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsPost200ResponseDataValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


