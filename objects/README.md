# Go API client for openapi

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.0
- Package version: 1.0.0
- Generator version: 7.8.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/stripedapps/attio-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.attio.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*CompaniesAPI* | [**V2ObjectsCompaniesRecordsPost**](docs/CompaniesAPI.md#v2objectscompaniesrecordspost) | **Post** /v2/objects/companies/records | Create a company record
*CompaniesAPI* | [**V2ObjectsCompaniesRecordsPut**](docs/CompaniesAPI.md#v2objectscompaniesrecordsput) | **Put** /v2/objects/companies/records | Assert a company record
*CompaniesAPI* | [**V2ObjectsCompaniesRecordsQueryPost**](docs/CompaniesAPI.md#v2objectscompaniesrecordsquerypost) | **Post** /v2/objects/companies/records/query | List company records
*CompaniesAPI* | [**V2ObjectsCompaniesRecordsRecordIdAttributesAttributeValuesGet**](docs/CompaniesAPI.md#v2objectscompaniesrecordsrecordidattributesattributevaluesget) | **Get** /v2/objects/companies/records/{record_id}/attributes/{attribute}/values | List company record attribute values
*CompaniesAPI* | [**V2ObjectsCompaniesRecordsRecordIdDelete**](docs/CompaniesAPI.md#v2objectscompaniesrecordsrecordiddelete) | **Delete** /v2/objects/companies/records/{record_id} | Delete a company record
*CompaniesAPI* | [**V2ObjectsCompaniesRecordsRecordIdEntriesGet**](docs/CompaniesAPI.md#v2objectscompaniesrecordsrecordidentriesget) | **Get** /v2/objects/companies/records/{record_id}/entries | List company record entries
*CompaniesAPI* | [**V2ObjectsCompaniesRecordsRecordIdGet**](docs/CompaniesAPI.md#v2objectscompaniesrecordsrecordidget) | **Get** /v2/objects/companies/records/{record_id} | Get a company record
*CompaniesAPI* | [**V2ObjectsCompaniesRecordsRecordIdPatch**](docs/CompaniesAPI.md#v2objectscompaniesrecordsrecordidpatch) | **Patch** /v2/objects/companies/records/{record_id} | Update a company record
*DealsAPI* | [**V2ObjectsDealsRecordsPost**](docs/DealsAPI.md#v2objectsdealsrecordspost) | **Post** /v2/objects/deals/records | Create a deal record
*DealsAPI* | [**V2ObjectsDealsRecordsPut**](docs/DealsAPI.md#v2objectsdealsrecordsput) | **Put** /v2/objects/deals/records | Assert a deal record
*DealsAPI* | [**V2ObjectsDealsRecordsQueryPost**](docs/DealsAPI.md#v2objectsdealsrecordsquerypost) | **Post** /v2/objects/deals/records/query | List deal records
*DealsAPI* | [**V2ObjectsDealsRecordsRecordIdAttributesAttributeValuesGet**](docs/DealsAPI.md#v2objectsdealsrecordsrecordidattributesattributevaluesget) | **Get** /v2/objects/deals/records/{record_id}/attributes/{attribute}/values | List deal record attribute values
*DealsAPI* | [**V2ObjectsDealsRecordsRecordIdDelete**](docs/DealsAPI.md#v2objectsdealsrecordsrecordiddelete) | **Delete** /v2/objects/deals/records/{record_id} | Delete a deal record
*DealsAPI* | [**V2ObjectsDealsRecordsRecordIdEntriesGet**](docs/DealsAPI.md#v2objectsdealsrecordsrecordidentriesget) | **Get** /v2/objects/deals/records/{record_id}/entries | List deal record entries
*DealsAPI* | [**V2ObjectsDealsRecordsRecordIdGet**](docs/DealsAPI.md#v2objectsdealsrecordsrecordidget) | **Get** /v2/objects/deals/records/{record_id} | Get a deal record
*DealsAPI* | [**V2ObjectsDealsRecordsRecordIdPatch**](docs/DealsAPI.md#v2objectsdealsrecordsrecordidpatch) | **Patch** /v2/objects/deals/records/{record_id} | Update a deal record
*PeopleAPI* | [**V2ObjectsPeopleRecordsPost**](docs/PeopleAPI.md#v2objectspeoplerecordspost) | **Post** /v2/objects/people/records | Create a person Record
*PeopleAPI* | [**V2ObjectsPeopleRecordsPut**](docs/PeopleAPI.md#v2objectspeoplerecordsput) | **Put** /v2/objects/people/records | Assert a person Record
*PeopleAPI* | [**V2ObjectsPeopleRecordsQueryPost**](docs/PeopleAPI.md#v2objectspeoplerecordsquerypost) | **Post** /v2/objects/people/records/query | List person records 
*PeopleAPI* | [**V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet**](docs/PeopleAPI.md#v2objectspeoplerecordsrecordidattributesattributevaluesget) | **Get** /v2/objects/people/records/{record_id}/attributes/{attribute}/values | List person record attribute values
*PeopleAPI* | [**V2ObjectsPeopleRecordsRecordIdDelete**](docs/PeopleAPI.md#v2objectspeoplerecordsrecordiddelete) | **Delete** /v2/objects/people/records/{record_id} | Delete a person Record
*PeopleAPI* | [**V2ObjectsPeopleRecordsRecordIdEntriesGet**](docs/PeopleAPI.md#v2objectspeoplerecordsrecordidentriesget) | **Get** /v2/objects/people/records/{record_id}/entries | List person record entries
*PeopleAPI* | [**V2ObjectsPeopleRecordsRecordIdGet**](docs/PeopleAPI.md#v2objectspeoplerecordsrecordidget) | **Get** /v2/objects/people/records/{record_id} | Get a person Record
*PeopleAPI* | [**V2ObjectsPeopleRecordsRecordIdPatch**](docs/PeopleAPI.md#v2objectspeoplerecordsrecordidpatch) | **Patch** /v2/objects/people/records/{record_id} | Update a person Record
*UsersAPI* | [**V2ObjectsUsersRecordsPost**](docs/UsersAPI.md#v2objectsusersrecordspost) | **Post** /v2/objects/users/records | Create a user record
*UsersAPI* | [**V2ObjectsUsersRecordsPut**](docs/UsersAPI.md#v2objectsusersrecordsput) | **Put** /v2/objects/users/records | Assert a user record
*UsersAPI* | [**V2ObjectsUsersRecordsQueryPost**](docs/UsersAPI.md#v2objectsusersrecordsquerypost) | **Post** /v2/objects/users/records/query | List user records
*UsersAPI* | [**V2ObjectsUsersRecordsRecordIdAttributesAttributeValuesGet**](docs/UsersAPI.md#v2objectsusersrecordsrecordidattributesattributevaluesget) | **Get** /v2/objects/users/records/{record_id}/attributes/{attribute}/values | List user record attribute values
*UsersAPI* | [**V2ObjectsUsersRecordsRecordIdDelete**](docs/UsersAPI.md#v2objectsusersrecordsrecordiddelete) | **Delete** /v2/objects/users/records/{record_id} | Delete a user record
*UsersAPI* | [**V2ObjectsUsersRecordsRecordIdEntriesGet**](docs/UsersAPI.md#v2objectsusersrecordsrecordidentriesget) | **Get** /v2/objects/users/records/{record_id}/entries | List user record entries
*UsersAPI* | [**V2ObjectsUsersRecordsRecordIdGet**](docs/UsersAPI.md#v2objectsusersrecordsrecordidget) | **Get** /v2/objects/users/records/{record_id} | Get a user record
*UsersAPI* | [**V2ObjectsUsersRecordsRecordIdPatch**](docs/UsersAPI.md#v2objectsusersrecordsrecordidpatch) | **Patch** /v2/objects/users/records/{record_id} | Update a user Record
*WorkspacesAPI* | [**V2ObjectsWorkspacesRecordsPost**](docs/WorkspacesAPI.md#v2objectsworkspacesrecordspost) | **Post** /v2/objects/workspaces/records | Create a workspace record
*WorkspacesAPI* | [**V2ObjectsWorkspacesRecordsPut**](docs/WorkspacesAPI.md#v2objectsworkspacesrecordsput) | **Put** /v2/objects/workspaces/records | Assert a workspace record
*WorkspacesAPI* | [**V2ObjectsWorkspacesRecordsQueryPost**](docs/WorkspacesAPI.md#v2objectsworkspacesrecordsquerypost) | **Post** /v2/objects/workspaces/records/query | List workspace records
*WorkspacesAPI* | [**V2ObjectsWorkspacesRecordsRecordIdAttributesAttributeValuesGet**](docs/WorkspacesAPI.md#v2objectsworkspacesrecordsrecordidattributesattributevaluesget) | **Get** /v2/objects/workspaces/records/{record_id}/attributes/{attribute}/values | List workspace record attribute values
*WorkspacesAPI* | [**V2ObjectsWorkspacesRecordsRecordIdDelete**](docs/WorkspacesAPI.md#v2objectsworkspacesrecordsrecordiddelete) | **Delete** /v2/objects/workspaces/records/{record_id} | Delete a workspace record
*WorkspacesAPI* | [**V2ObjectsWorkspacesRecordsRecordIdEntriesGet**](docs/WorkspacesAPI.md#v2objectsworkspacesrecordsrecordidentriesget) | **Get** /v2/objects/workspaces/records/{record_id}/entries | List workspace record entries
*WorkspacesAPI* | [**V2ObjectsWorkspacesRecordsRecordIdGet**](docs/WorkspacesAPI.md#v2objectsworkspacesrecordsrecordidget) | **Get** /v2/objects/workspaces/records/{record_id} | Get a workspace record
*WorkspacesAPI* | [**V2ObjectsWorkspacesRecordsRecordIdPatch**](docs/WorkspacesAPI.md#v2objectsworkspacesrecordsrecordidpatch) | **Patch** /v2/objects/workspaces/records/{record_id} | Update a workspace Record


## Documentation For Models

 - [SelectOption](docs/SelectOption.md)
 - [SelectOptionId](docs/SelectOptionId.md)
 - [Status](docs/Status.md)
 - [StatusId](docs/StatusId.md)
 - [V2ObjectsCompaniesRecordsPost200Response](docs/V2ObjectsCompaniesRecordsPost200Response.md)
 - [V2ObjectsCompaniesRecordsPost200ResponseData](docs/V2ObjectsCompaniesRecordsPost200ResponseData.md)
 - [V2ObjectsCompaniesRecordsPost200ResponseDataValues](docs/V2ObjectsCompaniesRecordsPost200ResponseDataValues.md)
 - [V2ObjectsCompaniesRecordsPut200Response](docs/V2ObjectsCompaniesRecordsPut200Response.md)
 - [V2ObjectsCompaniesRecordsPutRequest](docs/V2ObjectsCompaniesRecordsPutRequest.md)
 - [V2ObjectsCompaniesRecordsPutRequestData](docs/V2ObjectsCompaniesRecordsPutRequestData.md)
 - [V2ObjectsCompaniesRecordsPutRequestDataValues](docs/V2ObjectsCompaniesRecordsPutRequestDataValues.md)
 - [V2ObjectsCompaniesRecordsPutRequestDataValuesCategoriesInner](docs/V2ObjectsCompaniesRecordsPutRequestDataValuesCategoriesInner.md)
 - [V2ObjectsCompaniesRecordsPutRequestDataValuesFoundationDateInner](docs/V2ObjectsCompaniesRecordsPutRequestDataValuesFoundationDateInner.md)
 - [V2ObjectsCompaniesRecordsQueryPost200Response](docs/V2ObjectsCompaniesRecordsQueryPost200Response.md)
 - [V2ObjectsCompaniesRecordsQueryPost200ResponseDataInner](docs/V2ObjectsCompaniesRecordsQueryPost200ResponseDataInner.md)
 - [V2ObjectsCompaniesRecordsQueryPost200ResponseDataInnerValues](docs/V2ObjectsCompaniesRecordsQueryPost200ResponseDataInnerValues.md)
 - [V2ObjectsCompaniesRecordsQueryPostRequest](docs/V2ObjectsCompaniesRecordsQueryPostRequest.md)
 - [V2ObjectsCompaniesRecordsRecordIdAttributesAttributeValuesGet200Response](docs/V2ObjectsCompaniesRecordsRecordIdAttributesAttributeValuesGet200Response.md)
 - [V2ObjectsCompaniesRecordsRecordIdEntriesGet200Response](docs/V2ObjectsCompaniesRecordsRecordIdEntriesGet200Response.md)
 - [V2ObjectsCompaniesRecordsRecordIdEntriesGet200ResponseDataInner](docs/V2ObjectsCompaniesRecordsRecordIdEntriesGet200ResponseDataInner.md)
 - [V2ObjectsDealsRecordsPut200Response](docs/V2ObjectsDealsRecordsPut200Response.md)
 - [V2ObjectsDealsRecordsPut200ResponseData](docs/V2ObjectsDealsRecordsPut200ResponseData.md)
 - [V2ObjectsDealsRecordsPut200ResponseDataValues](docs/V2ObjectsDealsRecordsPut200ResponseDataValues.md)
 - [V2ObjectsDealsRecordsPutRequest](docs/V2ObjectsDealsRecordsPutRequest.md)
 - [V2ObjectsDealsRecordsPutRequestData](docs/V2ObjectsDealsRecordsPutRequestData.md)
 - [V2ObjectsDealsRecordsPutRequestDataValues](docs/V2ObjectsDealsRecordsPutRequestDataValues.md)
 - [V2ObjectsDealsRecordsPutRequestDataValuesOwnerInner](docs/V2ObjectsDealsRecordsPutRequestDataValuesOwnerInner.md)
 - [V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf](docs/V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf.md)
 - [V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf1](docs/V2ObjectsDealsRecordsPutRequestDataValuesOwnerInnerAnyOf1.md)
 - [V2ObjectsDealsRecordsPutRequestDataValuesStageInner](docs/V2ObjectsDealsRecordsPutRequestDataValuesStageInner.md)
 - [V2ObjectsDealsRecordsPutRequestDataValuesValueInner](docs/V2ObjectsDealsRecordsPutRequestDataValuesValueInner.md)
 - [V2ObjectsDealsRecordsQueryPost200Response](docs/V2ObjectsDealsRecordsQueryPost200Response.md)
 - [V2ObjectsDealsRecordsQueryPostRequest](docs/V2ObjectsDealsRecordsQueryPostRequest.md)
 - [V2ObjectsPeopleRecordsPost200Response](docs/V2ObjectsPeopleRecordsPost200Response.md)
 - [V2ObjectsPeopleRecordsPost200ResponseData](docs/V2ObjectsPeopleRecordsPost200ResponseData.md)
 - [V2ObjectsPeopleRecordsPost200ResponseDataValues](docs/V2ObjectsPeopleRecordsPost200ResponseDataValues.md)
 - [V2ObjectsPeopleRecordsPost400Response](docs/V2ObjectsPeopleRecordsPost400Response.md)
 - [V2ObjectsPeopleRecordsPut200Response](docs/V2ObjectsPeopleRecordsPut200Response.md)
 - [V2ObjectsPeopleRecordsPut400Response](docs/V2ObjectsPeopleRecordsPut400Response.md)
 - [V2ObjectsPeopleRecordsPut404Response](docs/V2ObjectsPeopleRecordsPut404Response.md)
 - [V2ObjectsPeopleRecordsPutRequest](docs/V2ObjectsPeopleRecordsPutRequest.md)
 - [V2ObjectsPeopleRecordsPutRequestData](docs/V2ObjectsPeopleRecordsPutRequestData.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValues](docs/V2ObjectsPeopleRecordsPutRequestDataValues.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner](docs/V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInner.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf](docs/V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1](docs/V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInner](docs/V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInner.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf](docs/V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1](docs/V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf1.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2](docs/V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf2.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3](docs/V2ObjectsPeopleRecordsPutRequestDataValuesCompanyInnerAnyOf1SlugOrIdOfMatchingAttributeInnerAnyOf3.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner](docs/V2ObjectsPeopleRecordsPutRequestDataValuesDescriptionInner.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesEmailAddressesInner](docs/V2ObjectsPeopleRecordsPutRequestDataValuesEmailAddressesInner.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesNameInner](docs/V2ObjectsPeopleRecordsPutRequestDataValuesNameInner.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesPhoneNumbersInner](docs/V2ObjectsPeopleRecordsPutRequestDataValuesPhoneNumbersInner.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesPrimaryLocationInner](docs/V2ObjectsPeopleRecordsPutRequestDataValuesPrimaryLocationInner.md)
 - [V2ObjectsPeopleRecordsPutRequestDataValuesTwitterFollowerCountInner](docs/V2ObjectsPeopleRecordsPutRequestDataValuesTwitterFollowerCountInner.md)
 - [V2ObjectsPeopleRecordsQueryPost200Response](docs/V2ObjectsPeopleRecordsQueryPost200Response.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInner](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInner.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerId](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerId.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValues.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesCompanyInner.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesDescriptionInner.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInner.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesEmailAddressesInnerCreatedByActor.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesNameInner.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPhoneNumbersInner.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesPrimaryLocationInner.md)
 - [V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner](docs/V2ObjectsPeopleRecordsQueryPost200ResponseDataInnerValuesTwitterFollowerCountInner.md)
 - [V2ObjectsPeopleRecordsQueryPost400Response](docs/V2ObjectsPeopleRecordsQueryPost400Response.md)
 - [V2ObjectsPeopleRecordsQueryPost404Response](docs/V2ObjectsPeopleRecordsQueryPost404Response.md)
 - [V2ObjectsPeopleRecordsQueryPostRequest](docs/V2ObjectsPeopleRecordsQueryPostRequest.md)
 - [V2ObjectsPeopleRecordsQueryPostRequestSortsInner](docs/V2ObjectsPeopleRecordsQueryPostRequestSortsInner.md)
 - [V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf](docs/V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf.md)
 - [V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1](docs/V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1.md)
 - [V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner](docs/V2ObjectsPeopleRecordsQueryPostRequestSortsInnerAnyOf1PathInnerInner.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200Response](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200Response.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInner](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInner.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf1](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf1.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf2](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf2.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf3](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf3.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf4](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf4.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf5](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf5.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf6.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf7](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf7.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf8](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf8.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf9](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet200ResponseDataInnerOneOf9.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet400Response](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet400Response.md)
 - [V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet404Response](docs/V2ObjectsPeopleRecordsRecordIdAttributesAttributeValuesGet404Response.md)
 - [V2ObjectsPeopleRecordsRecordIdEntriesGet200Response](docs/V2ObjectsPeopleRecordsRecordIdEntriesGet200Response.md)
 - [V2ObjectsPeopleRecordsRecordIdEntriesGet200ResponseDataInner](docs/V2ObjectsPeopleRecordsRecordIdEntriesGet200ResponseDataInner.md)
 - [V2ObjectsPeopleRecordsRecordIdGet404Response](docs/V2ObjectsPeopleRecordsRecordIdGet404Response.md)
 - [V2ObjectsPeopleRecordsRecordIdPatch400Response](docs/V2ObjectsPeopleRecordsRecordIdPatch400Response.md)
 - [V2ObjectsUsersRecordsPut200Response](docs/V2ObjectsUsersRecordsPut200Response.md)
 - [V2ObjectsUsersRecordsPutRequest](docs/V2ObjectsUsersRecordsPutRequest.md)
 - [V2ObjectsUsersRecordsPutRequestData](docs/V2ObjectsUsersRecordsPutRequestData.md)
 - [V2ObjectsUsersRecordsPutRequestDataValues](docs/V2ObjectsUsersRecordsPutRequestDataValues.md)
 - [V2ObjectsUsersRecordsQueryPost200Response](docs/V2ObjectsUsersRecordsQueryPost200Response.md)
 - [V2ObjectsUsersRecordsQueryPost200ResponseDataInner](docs/V2ObjectsUsersRecordsQueryPost200ResponseDataInner.md)
 - [V2ObjectsUsersRecordsQueryPost200ResponseDataInnerValues](docs/V2ObjectsUsersRecordsQueryPost200ResponseDataInnerValues.md)
 - [V2ObjectsUsersRecordsQueryPostRequest](docs/V2ObjectsUsersRecordsQueryPostRequest.md)
 - [V2ObjectsUsersRecordsRecordIdAttributesAttributeValuesGet200Response](docs/V2ObjectsUsersRecordsRecordIdAttributesAttributeValuesGet200Response.md)
 - [V2ObjectsWorkspacesRecordsPost200Response](docs/V2ObjectsWorkspacesRecordsPost200Response.md)
 - [V2ObjectsWorkspacesRecordsPost200ResponseData](docs/V2ObjectsWorkspacesRecordsPost200ResponseData.md)
 - [V2ObjectsWorkspacesRecordsPost200ResponseDataValues](docs/V2ObjectsWorkspacesRecordsPost200ResponseDataValues.md)
 - [V2ObjectsWorkspacesRecordsPut200Response](docs/V2ObjectsWorkspacesRecordsPut200Response.md)
 - [V2ObjectsWorkspacesRecordsPut200ResponseData](docs/V2ObjectsWorkspacesRecordsPut200ResponseData.md)
 - [V2ObjectsWorkspacesRecordsPut200ResponseDataValues](docs/V2ObjectsWorkspacesRecordsPut200ResponseDataValues.md)
 - [V2ObjectsWorkspacesRecordsPutRequest](docs/V2ObjectsWorkspacesRecordsPutRequest.md)
 - [V2ObjectsWorkspacesRecordsPutRequestData](docs/V2ObjectsWorkspacesRecordsPutRequestData.md)
 - [V2ObjectsWorkspacesRecordsPutRequestDataValues](docs/V2ObjectsWorkspacesRecordsPutRequestDataValues.md)
 - [V2ObjectsWorkspacesRecordsQueryPost200Response](docs/V2ObjectsWorkspacesRecordsQueryPost200Response.md)
 - [V2ObjectsWorkspacesRecordsQueryPostRequest](docs/V2ObjectsWorkspacesRecordsQueryPostRequest.md)
 - [V2ObjectsWorkspacesRecordsRecordIdAttributesAttributeValuesGet200Response](docs/V2ObjectsWorkspacesRecordsRecordIdAttributesAttributeValuesGet200Response.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### oauth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://app.attio.com/authorize
- **Scopes**: 
 - **user_management:read**: View workspace members.
 - **user_management:read-write**: View workspace members.
 - **record_permission:read**: View, and optionally write, records.
 - **record_permission:read-write**: View, and optionally write, records.
 - **object_configuration:read**: View, and optionally write, the configuration and attributes of objects.
 - **object_configuration:read-write**: View, and optionally write, the configuration and attributes of objects.
 - **list_entry:read**: View, and optionally write, the entries in a list.
 - **list_entry:read-write**: View, and optionally write, the entries in a list.
 - **list_configuration:read**: View, and optionally write, the configuration and attributes of lists.
 - **list_configuration:read-write**: View, and optionally write, the configuration and attributes of lists.
 - **public_collection:read**: View, and optionally write, both the settings and information within public collections.
 - **public_collection:read-write**: View, and optionally write, both the settings and information within public collections.
 - **private_collection:read**: View, and optionally modify, both the settings and information of all collections within the workspace, regardless of their access settings.
 - **private_collection:read-write**: View, and optionally modify, both the settings and information of all collections within the workspace, regardless of their access settings.
 - **comment:read**: View comments (and threads), and optionally write comments.
 - **comment:read-write**: View comments (and threads), and optionally write comments.
 - **task:read**: View, and optionally write, tasks.
 - **task:read-write**: View, and optionally write, tasks.
 - **note:read**: View, and optionally write, notes.
 - **note:read-write**: View, and optionally write, notes.
 - **webhook:read**: View, and optionally manage, webhooks.
 - **webhook:read-write**: View, and optionally manage, webhooks.

Example

```go
auth := context.WithValue(context.Background(), openapi.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, openapi.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



