/*
Attio OAuth flow

Welcome to the Attio OAuth flow documentation. All of the OAuth APIs are documented here.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OauthIntrospectPost200ResponseAnyOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OauthIntrospectPost200ResponseAnyOf1{}

// OauthIntrospectPost200ResponseAnyOf1 struct for OauthIntrospectPost200ResponseAnyOf1
type OauthIntrospectPost200ResponseAnyOf1 struct {
	// Whether the token is currently active and usable.
	Active bool `json:"active"`
	// A space-separated list of scopes associated with this token
	Scope string `json:"scope"`
	// The app ID of the OAuth application that requested this token
	ClientId string `json:"client_id"`
	// The type of token, always Bearer for tokens acquired via the OAuth 2.0 flow.
	TokenType string `json:"token_type"`
	// The time at which this token will expire, if set, as a number of seconds since January 1 1970 UTC.
	Exp NullableFloat32 `json:"exp"`
	// The time at which this token was issued, as a number of seconds since January 1 1970 UTC.
	Iat float32 `json:"iat"`
	// Since Bearer tokens grant Workspace-level permissions, this property contains the workspace_id.
	Sub string `json:"sub"`
	// The intended audience for this token, for Bearer tokens this is the same as the client_id.
	Aud string `json:"aud"`
	// The issuer of the token. Always attio.com
	Iss string `json:"iss"`
	// The ID of the workspace member who authorised this token initially, if known
	AuthorizedByWorkspaceMemberId NullableString `json:"authorized_by_workspace_member_id"`
	// The ID of the workspace the token is scoped to.
	WorkspaceId string `json:"workspace_id"`
	// The name of the workspace the token is scoped to.
	WorkspaceName string `json:"workspace_name"`
	// The slug of the workspace the token is scoped to.
	WorkspaceSlug string `json:"workspace_slug"`
	// The logo URL of the workspace the token is scoped to.
	WorkspaceLogoUrl NullableString `json:"workspace_logo_url"`
}

type _OauthIntrospectPost200ResponseAnyOf1 OauthIntrospectPost200ResponseAnyOf1

// NewOauthIntrospectPost200ResponseAnyOf1 instantiates a new OauthIntrospectPost200ResponseAnyOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauthIntrospectPost200ResponseAnyOf1(active bool, scope string, clientId string, tokenType string, exp NullableFloat32, iat float32, sub string, aud string, iss string, authorizedByWorkspaceMemberId NullableString, workspaceId string, workspaceName string, workspaceSlug string, workspaceLogoUrl NullableString) *OauthIntrospectPost200ResponseAnyOf1 {
	this := OauthIntrospectPost200ResponseAnyOf1{}
	this.Active = active
	this.Scope = scope
	this.ClientId = clientId
	this.TokenType = tokenType
	this.Exp = exp
	this.Iat = iat
	this.Sub = sub
	this.Aud = aud
	this.Iss = iss
	this.AuthorizedByWorkspaceMemberId = authorizedByWorkspaceMemberId
	this.WorkspaceId = workspaceId
	this.WorkspaceName = workspaceName
	this.WorkspaceSlug = workspaceSlug
	this.WorkspaceLogoUrl = workspaceLogoUrl
	return &this
}

// NewOauthIntrospectPost200ResponseAnyOf1WithDefaults instantiates a new OauthIntrospectPost200ResponseAnyOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauthIntrospectPost200ResponseAnyOf1WithDefaults() *OauthIntrospectPost200ResponseAnyOf1 {
	this := OauthIntrospectPost200ResponseAnyOf1{}
	return &this
}

// GetActive returns the Active field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetActive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Active
}

// GetActiveOk returns a tuple with the Active field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetActiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Active, true
}

// SetActive sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetActive(v bool) {
	o.Active = v
}

// GetScope returns the Scope field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetScope(v string) {
	o.Scope = v
}

// GetClientId returns the ClientId field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetClientId(v string) {
	o.ClientId = v
}

// GetTokenType returns the TokenType field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetTokenType(v string) {
	o.TokenType = v
}

// GetExp returns the Exp field value
// If the value is explicit nil, the zero value for float32 will be returned
func (o *OauthIntrospectPost200ResponseAnyOf1) GetExp() float32 {
	if o == nil || o.Exp.Get() == nil {
		var ret float32
		return ret
	}

	return *o.Exp.Get()
}

// GetExpOk returns a tuple with the Exp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OauthIntrospectPost200ResponseAnyOf1) GetExpOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Exp.Get(), o.Exp.IsSet()
}

// SetExp sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetExp(v float32) {
	o.Exp.Set(&v)
}

// GetIat returns the Iat field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetIat() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Iat
}

// GetIatOk returns a tuple with the Iat field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetIatOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iat, true
}

// SetIat sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetIat(v float32) {
	o.Iat = v
}

// GetSub returns the Sub field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetSub() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sub
}

// GetSubOk returns a tuple with the Sub field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetSubOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sub, true
}

// SetSub sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetSub(v string) {
	o.Sub = v
}

// GetAud returns the Aud field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetAud() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Aud
}

// GetAudOk returns a tuple with the Aud field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetAudOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aud, true
}

// SetAud sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetAud(v string) {
	o.Aud = v
}

// GetIss returns the Iss field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetIss() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Iss
}

// GetIssOk returns a tuple with the Iss field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetIssOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iss, true
}

// SetIss sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetIss(v string) {
	o.Iss = v
}

// GetAuthorizedByWorkspaceMemberId returns the AuthorizedByWorkspaceMemberId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OauthIntrospectPost200ResponseAnyOf1) GetAuthorizedByWorkspaceMemberId() string {
	if o == nil || o.AuthorizedByWorkspaceMemberId.Get() == nil {
		var ret string
		return ret
	}

	return *o.AuthorizedByWorkspaceMemberId.Get()
}

// GetAuthorizedByWorkspaceMemberIdOk returns a tuple with the AuthorizedByWorkspaceMemberId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OauthIntrospectPost200ResponseAnyOf1) GetAuthorizedByWorkspaceMemberIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizedByWorkspaceMemberId.Get(), o.AuthorizedByWorkspaceMemberId.IsSet()
}

// SetAuthorizedByWorkspaceMemberId sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetAuthorizedByWorkspaceMemberId(v string) {
	o.AuthorizedByWorkspaceMemberId.Set(&v)
}

// GetWorkspaceId returns the WorkspaceId field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetWorkspaceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetWorkspaceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceId, true
}

// SetWorkspaceId sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetWorkspaceId(v string) {
	o.WorkspaceId = v
}

// GetWorkspaceName returns the WorkspaceName field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetWorkspaceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceName
}

// GetWorkspaceNameOk returns a tuple with the WorkspaceName field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetWorkspaceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceName, true
}

// SetWorkspaceName sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetWorkspaceName(v string) {
	o.WorkspaceName = v
}

// GetWorkspaceSlug returns the WorkspaceSlug field value
func (o *OauthIntrospectPost200ResponseAnyOf1) GetWorkspaceSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkspaceSlug
}

// GetWorkspaceSlugOk returns a tuple with the WorkspaceSlug field value
// and a boolean to check if the value has been set.
func (o *OauthIntrospectPost200ResponseAnyOf1) GetWorkspaceSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkspaceSlug, true
}

// SetWorkspaceSlug sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetWorkspaceSlug(v string) {
	o.WorkspaceSlug = v
}

// GetWorkspaceLogoUrl returns the WorkspaceLogoUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OauthIntrospectPost200ResponseAnyOf1) GetWorkspaceLogoUrl() string {
	if o == nil || o.WorkspaceLogoUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.WorkspaceLogoUrl.Get()
}

// GetWorkspaceLogoUrlOk returns a tuple with the WorkspaceLogoUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OauthIntrospectPost200ResponseAnyOf1) GetWorkspaceLogoUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WorkspaceLogoUrl.Get(), o.WorkspaceLogoUrl.IsSet()
}

// SetWorkspaceLogoUrl sets field value
func (o *OauthIntrospectPost200ResponseAnyOf1) SetWorkspaceLogoUrl(v string) {
	o.WorkspaceLogoUrl.Set(&v)
}

func (o OauthIntrospectPost200ResponseAnyOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OauthIntrospectPost200ResponseAnyOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["active"] = o.Active
	toSerialize["scope"] = o.Scope
	toSerialize["client_id"] = o.ClientId
	toSerialize["token_type"] = o.TokenType
	toSerialize["exp"] = o.Exp.Get()
	toSerialize["iat"] = o.Iat
	toSerialize["sub"] = o.Sub
	toSerialize["aud"] = o.Aud
	toSerialize["iss"] = o.Iss
	toSerialize["authorized_by_workspace_member_id"] = o.AuthorizedByWorkspaceMemberId.Get()
	toSerialize["workspace_id"] = o.WorkspaceId
	toSerialize["workspace_name"] = o.WorkspaceName
	toSerialize["workspace_slug"] = o.WorkspaceSlug
	toSerialize["workspace_logo_url"] = o.WorkspaceLogoUrl.Get()
	return toSerialize, nil
}

func (o *OauthIntrospectPost200ResponseAnyOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"active",
		"scope",
		"client_id",
		"token_type",
		"exp",
		"iat",
		"sub",
		"aud",
		"iss",
		"authorized_by_workspace_member_id",
		"workspace_id",
		"workspace_name",
		"workspace_slug",
		"workspace_logo_url",
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

	varOauthIntrospectPost200ResponseAnyOf1 := _OauthIntrospectPost200ResponseAnyOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOauthIntrospectPost200ResponseAnyOf1)

	if err != nil {
		return err
	}

	*o = OauthIntrospectPost200ResponseAnyOf1(varOauthIntrospectPost200ResponseAnyOf1)

	return err
}

type NullableOauthIntrospectPost200ResponseAnyOf1 struct {
	value *OauthIntrospectPost200ResponseAnyOf1
	isSet bool
}

func (v NullableOauthIntrospectPost200ResponseAnyOf1) Get() *OauthIntrospectPost200ResponseAnyOf1 {
	return v.value
}

func (v *NullableOauthIntrospectPost200ResponseAnyOf1) Set(val *OauthIntrospectPost200ResponseAnyOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableOauthIntrospectPost200ResponseAnyOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableOauthIntrospectPost200ResponseAnyOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauthIntrospectPost200ResponseAnyOf1(val *OauthIntrospectPost200ResponseAnyOf1) *NullableOauthIntrospectPost200ResponseAnyOf1 {
	return &NullableOauthIntrospectPost200ResponseAnyOf1{value: val, isSet: true}
}

func (v NullableOauthIntrospectPost200ResponseAnyOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauthIntrospectPost200ResponseAnyOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


