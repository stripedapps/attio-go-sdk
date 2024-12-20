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
)

// checks if the V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues{}

type V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesRecordIdInner struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor `json:"created_by_actor"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	// The value type of the record
	Value string `json:"value"`
}

// V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues An object with `attribute_slug` keys, and an array of value objects as the values. Attributes slugs (for example `email_addresses` or `name`) can be used, including custom attribute slugs.
type V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues struct {
	RecordId []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesRecordIdInner `json:"record_id,omitempty"`
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

// NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues instantiates a new V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues() *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues {
	this := V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues{}
	return &this
}

// NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesWithDefaults instantiates a new V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesWithDefaults() *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues {
	this := V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues{}
	return &this
}

// GetEmailAddresses returns the EmailAddresses field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetEmailAddresses() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner {
	if o == nil || IsNil(o.EmailAddresses) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner
		return ret
	}
	return o.EmailAddresses
}

// GetEmailAddressesOk returns a tuple with the EmailAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetEmailAddressesOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner, bool) {
	if o == nil || IsNil(o.EmailAddresses) {
		return nil, false
	}
	return o.EmailAddresses, true
}

// HasEmailAddresses returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasEmailAddresses() bool {
	if o != nil && !IsNil(o.EmailAddresses) {
		return true
	}

	return false
}

// SetEmailAddresses gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner and assigns it to the EmailAddresses field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetEmailAddresses(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner) {
	o.EmailAddresses = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetName() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner {
	if o == nil || IsNil(o.Name) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetNameOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner and assigns it to the Name field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetName(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetDescription() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Description) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetDescriptionOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Description field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetDescription(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Description = v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetAvatarUrl() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.AvatarUrl) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetAvatarUrlOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.AvatarUrl) {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasAvatarUrl() bool {
	if o != nil && !IsNil(o.AvatarUrl) {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the AvatarUrl field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetAvatarUrl(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.AvatarUrl = v
}

// GetJobTitle returns the JobTitle field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetJobTitle() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.JobTitle) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.JobTitle
}

// GetJobTitleOk returns a tuple with the JobTitle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetJobTitleOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.JobTitle) {
		return nil, false
	}
	return o.JobTitle, true
}

// HasJobTitle returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasJobTitle() bool {
	if o != nil && !IsNil(o.JobTitle) {
		return true
	}

	return false
}

// SetJobTitle gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the JobTitle field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetJobTitle(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.JobTitle = v
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetPhoneNumbers() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner {
	if o == nil || IsNil(o.PhoneNumbers) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner
		return ret
	}
	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetPhoneNumbersOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner, bool) {
	if o == nil || IsNil(o.PhoneNumbers) {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasPhoneNumbers() bool {
	if o != nil && !IsNil(o.PhoneNumbers) {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner and assigns it to the PhoneNumbers field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetPhoneNumbers(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner) {
	o.PhoneNumbers = v
}

// GetLinkedin returns the Linkedin field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetLinkedin() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Linkedin) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Linkedin
}

// GetLinkedinOk returns a tuple with the Linkedin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetLinkedinOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Linkedin) {
		return nil, false
	}
	return o.Linkedin, true
}

// HasLinkedin returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasLinkedin() bool {
	if o != nil && !IsNil(o.Linkedin) {
		return true
	}

	return false
}

// SetLinkedin gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Linkedin field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetLinkedin(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Linkedin = v
}

// GetTwitter returns the Twitter field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetTwitter() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Twitter) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetTwitterOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Twitter) {
		return nil, false
	}
	return o.Twitter, true
}

// HasTwitter returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasTwitter() bool {
	if o != nil && !IsNil(o.Twitter) {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Twitter field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetTwitter(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Twitter = v
}

// GetFacebook returns the Facebook field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetFacebook() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Facebook) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Facebook
}

// GetFacebookOk returns a tuple with the Facebook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetFacebookOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Facebook) {
		return nil, false
	}
	return o.Facebook, true
}

// HasFacebook returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasFacebook() bool {
	if o != nil && !IsNil(o.Facebook) {
		return true
	}

	return false
}

// SetFacebook gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Facebook field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetFacebook(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Facebook = v
}

// GetInstagram returns the Instagram field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetInstagram() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Instagram) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Instagram
}

// GetInstagramOk returns a tuple with the Instagram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetInstagramOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Instagram) {
		return nil, false
	}
	return o.Instagram, true
}

// HasInstagram returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasInstagram() bool {
	if o != nil && !IsNil(o.Instagram) {
		return true
	}

	return false
}

// SetInstagram gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Instagram field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetInstagram(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Instagram = v
}

// GetAngellist returns the Angellist field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetAngellist() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner {
	if o == nil || IsNil(o.Angellist) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner
		return ret
	}
	return o.Angellist
}

// GetAngellistOk returns a tuple with the Angellist field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetAngellistOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner, bool) {
	if o == nil || IsNil(o.Angellist) {
		return nil, false
	}
	return o.Angellist, true
}

// HasAngellist returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasAngellist() bool {
	if o != nil && !IsNil(o.Angellist) {
		return true
	}

	return false
}

// SetAngellist gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner and assigns it to the Angellist field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetAngellist(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner) {
	o.Angellist = v
}

// GetPrimaryLocation returns the PrimaryLocation field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetPrimaryLocation() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner {
	if o == nil || IsNil(o.PrimaryLocation) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner
		return ret
	}
	return o.PrimaryLocation
}

// GetPrimaryLocationOk returns a tuple with the PrimaryLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetPrimaryLocationOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner, bool) {
	if o == nil || IsNil(o.PrimaryLocation) {
		return nil, false
	}
	return o.PrimaryLocation, true
}

// HasPrimaryLocation returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasPrimaryLocation() bool {
	if o != nil && !IsNil(o.PrimaryLocation) {
		return true
	}

	return false
}

// SetPrimaryLocation gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner and assigns it to the PrimaryLocation field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetPrimaryLocation(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) {
	o.PrimaryLocation = v
}

// GetTwitterFollowerCount returns the TwitterFollowerCount field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetTwitterFollowerCount() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner {
	if o == nil || IsNil(o.TwitterFollowerCount) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner
		return ret
	}
	return o.TwitterFollowerCount
}

// GetTwitterFollowerCountOk returns a tuple with the TwitterFollowerCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetTwitterFollowerCountOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner, bool) {
	if o == nil || IsNil(o.TwitterFollowerCount) {
		return nil, false
	}
	return o.TwitterFollowerCount, true
}

// HasTwitterFollowerCount returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasTwitterFollowerCount() bool {
	if o != nil && !IsNil(o.TwitterFollowerCount) {
		return true
	}

	return false
}

// SetTwitterFollowerCount gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner and assigns it to the TwitterFollowerCount field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetTwitterFollowerCount(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner) {
	o.TwitterFollowerCount = v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetCompany() []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner {
	if o == nil || IsNil(o.Company) {
		var ret []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner
		return ret
	}
	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) GetCompanyOk() ([]V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner and assigns it to the Company field.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) SetCompany(v []V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner) {
	o.Company = v
}

func (o V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) ToMap() (map[string]interface{}, error) {
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

type NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues struct {
	value *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) Get() *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) Set(val *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues(val *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues {
	return &NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


