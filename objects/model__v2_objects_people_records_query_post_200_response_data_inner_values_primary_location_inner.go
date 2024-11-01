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
	"fmt"
)

// checks if the V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner{}

// V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner struct for V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner
type V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner struct {
	// The point in time at which this value was made \"active\". `active_from` can be considered roughly analogous to `created_at`.
	ActiveFrom time.Time `json:"active_from"`
	// The point in time at which this value was deactivated. If `null`, the value is active.
	ActiveUntil NullableTime `json:"active_until"`
	CreatedByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor `json:"created_by_actor"`
	// The first line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.
	Line1 string `json:"line_1"`
	// The second line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.
	Line2 string `json:"line_2"`
	// The third line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.
	Line3 string `json:"line_3"`
	// The fourth line of the address. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.
	Line4 string `json:"line_4"`
	// The town, neighborhood or area the location is in.
	Locality string `json:"locality"`
	// The state, county, province or region that the location is in.
	Region string `json:"region"`
	// The postcode or zip code for the location. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.}
	Postcode string `json:"postcode"`
	// The ISO 3166-1 alpha-2 country code for the country this location is in.
	CountryCode string `json:"country_code"`
	// The latitude of the location. Validated by the regular expression `/^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$/`. Note that this value is not currently represented in the UI but will be persisted and readable through API calls.}
	Latitude string `json:"latitude" validate:"regexp=^[-+]?([1-8]?\\\\d(\\\\.\\\\d+)?|90(\\\\.0+)?)$"`
	// The longitude of the location. Validated by the regular expression `/^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$/`
	Longitude string `json:"longitude" validate:"regexp=^[-+]?(180(\\\\.0+)?|((1[0-7]\\\\d)|([1-9]?\\\\d))(\\\\.\\\\d+)?)$"`
	// The attribute type of the value.
	AttributeType string `json:"attribute_type"`
	AdditionalProperties map[string]interface{}
}

type _V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner

// NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner instantiates a new V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, line1 string, line2 string, line3 string, line4 string, locality string, region string, postcode string, countryCode string, latitude string, longitude string, attributeType string) *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner {
	this := V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner{}
	this.ActiveFrom = activeFrom
	this.ActiveUntil = activeUntil
	this.CreatedByActor = createdByActor
	this.Line1 = line1
	this.Line2 = line2
	this.Line3 = line3
	this.Line4 = line4
	this.Locality = locality
	this.Region = region
	this.Postcode = postcode
	this.CountryCode = countryCode
	this.Latitude = latitude
	this.Longitude = longitude
	this.AttributeType = attributeType
	return &this
}

// NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInnerWithDefaults instantiates a new V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInnerWithDefaults() *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner {
	this := V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner{}
	return &this
}

// GetActiveFrom returns the ActiveFrom field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetActiveFrom() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ActiveFrom
}

// GetActiveFromOk returns a tuple with the ActiveFrom field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetActiveFromOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveFrom, true
}

// SetActiveFrom sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetActiveFrom(v time.Time) {
	o.ActiveFrom = v
}

// GetActiveUntil returns the ActiveUntil field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetActiveUntil() time.Time {
	if o == nil || o.ActiveUntil.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.ActiveUntil.Get()
}

// GetActiveUntilOk returns a tuple with the ActiveUntil field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetActiveUntilOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActiveUntil.Get(), o.ActiveUntil.IsSet()
}

// SetActiveUntil sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetActiveUntil(v time.Time) {
	o.ActiveUntil.Set(&v)
}

// GetCreatedByActor returns the CreatedByActor field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetCreatedByActor() V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor {
	if o == nil {
		var ret V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor
		return ret
	}

	return o.CreatedByActor
}

// GetCreatedByActorOk returns a tuple with the CreatedByActor field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetCreatedByActorOk() (*V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedByActor, true
}

// SetCreatedByActor sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetCreatedByActor(v V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor) {
	o.CreatedByActor = v
}

// GetLine1 returns the Line1 field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLine1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Line1
}

// GetLine1Ok returns a tuple with the Line1 field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLine1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line1, true
}

// SetLine1 sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetLine1(v string) {
	o.Line1 = v
}

// GetLine2 returns the Line2 field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLine2() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Line2
}

// GetLine2Ok returns a tuple with the Line2 field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLine2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line2, true
}

// SetLine2 sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetLine2(v string) {
	o.Line2 = v
}

// GetLine3 returns the Line3 field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLine3() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Line3
}

// GetLine3Ok returns a tuple with the Line3 field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLine3Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line3, true
}

// SetLine3 sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetLine3(v string) {
	o.Line3 = v
}

// GetLine4 returns the Line4 field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLine4() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Line4
}

// GetLine4Ok returns a tuple with the Line4 field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLine4Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line4, true
}

// SetLine4 sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetLine4(v string) {
	o.Line4 = v
}

// GetLocality returns the Locality field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLocality() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locality
}

// GetLocalityOk returns a tuple with the Locality field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLocalityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locality, true
}

// SetLocality sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetLocality(v string) {
	o.Locality = v
}

// GetRegion returns the Region field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetRegion(v string) {
	o.Region = v
}

// GetPostcode returns the Postcode field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetPostcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Postcode
}

// GetPostcodeOk returns a tuple with the Postcode field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetPostcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Postcode, true
}

// SetPostcode sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetPostcode(v string) {
	o.Postcode = v
}

// GetCountryCode returns the CountryCode field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetLatitude returns the Latitude field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLatitude() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLatitudeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Latitude, true
}

// SetLatitude sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetLatitude(v string) {
	o.Latitude = v
}

// GetLongitude returns the Longitude field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLongitude() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetLongitudeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Longitude, true
}

// SetLongitude sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetLongitude(v string) {
	o.Longitude = v
}

// GetAttributeType returns the AttributeType field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetAttributeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AttributeType
}

// GetAttributeTypeOk returns a tuple with the AttributeType field value
// and a boolean to check if the value has been set.
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) GetAttributeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AttributeType, true
}

// SetAttributeType sets field value
func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) SetAttributeType(v string) {
	o.AttributeType = v
}

func (o V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active_from"] = o.ActiveFrom
	toSerialize["active_until"] = o.ActiveUntil.Get()
	toSerialize["created_by_actor"] = o.CreatedByActor
	toSerialize["line_1"] = o.Line1
	toSerialize["line_2"] = o.Line2
	toSerialize["line_3"] = o.Line3
	toSerialize["line_4"] = o.Line4
	toSerialize["locality"] = o.Locality
	toSerialize["region"] = o.Region
	toSerialize["postcode"] = o.Postcode
	toSerialize["country_code"] = o.CountryCode
	toSerialize["latitude"] = o.Latitude
	toSerialize["longitude"] = o.Longitude
	toSerialize["attribute_type"] = o.AttributeType

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active_from",
		"active_until",
		"created_by_actor",
		"line_1",
		"line_2",
		"line_3",
		"line_4",
		"locality",
		"region",
		"postcode",
		"country_code",
		"latitude",
		"longitude",
		"attribute_type",
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

	varV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner := _V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner{}

	err = json.Unmarshal(data, &varV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner)

	if err != nil {
		return err
	}

	*o = V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner(varV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "active_from")
		delete(additionalProperties, "active_until")
		delete(additionalProperties, "created_by_actor")
		delete(additionalProperties, "line_1")
		delete(additionalProperties, "line_2")
		delete(additionalProperties, "line_3")
		delete(additionalProperties, "line_4")
		delete(additionalProperties, "locality")
		delete(additionalProperties, "region")
		delete(additionalProperties, "postcode")
		delete(additionalProperties, "country_code")
		delete(additionalProperties, "latitude")
		delete(additionalProperties, "longitude")
		delete(additionalProperties, "attribute_type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner struct {
	value *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner
	isSet bool
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) Get() *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner {
	return v.value
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) Set(val *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) {
	v.value = val
	v.isSet = true
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) IsSet() bool {
	return v.isSet
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner(val *V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner {
	return &NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner{value: val, isSet: true}
}

func (v NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


