/*
Attio Webhook Events

This page contains the schema, and an example, for every type of Webhook that we currently send. Our [using Webhooks guide](/docs/using-webhooks) contains more information about how to use Webhooks.

API version: latest
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ObjectAttributeEventsAPIService ObjectAttributeEventsAPI service
type ObjectAttributeEventsAPIService service

type ApiWebhooksObjectAttributeCreatedRequest struct {
	ctx context.Context
	ApiService *ObjectAttributeEventsAPIService
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiWebhooksObjectAttributeCreatedRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiWebhooksObjectAttributeCreatedRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiWebhooksObjectAttributeCreatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksObjectAttributeCreatedExecute(r)
}

/*
WebhooksObjectAttributeCreated object-attribute.created

This event is fired whenever an object attribute is created (e.g. when defining a new attribute "Rating" on the company object).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksObjectAttributeCreatedRequest
*/
func (a *ObjectAttributeEventsAPIService) WebhooksObjectAttributeCreated(ctx context.Context) ApiWebhooksObjectAttributeCreatedRequest {
	return ApiWebhooksObjectAttributeCreatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ObjectAttributeEventsAPIService) WebhooksObjectAttributeCreatedExecute(r ApiWebhooksObjectAttributeCreatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectAttributeEventsAPIService.WebhooksObjectAttributeCreated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-attribute.created"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uNKNOWNBASETYPE
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiWebhooksObjectAttributeUpdatedRequest struct {
	ctx context.Context
	ApiService *ObjectAttributeEventsAPIService
	uNKNOWNBASETYPE *UNKNOWN_BASE_TYPE
}

func (r ApiWebhooksObjectAttributeUpdatedRequest) UNKNOWNBASETYPE(uNKNOWNBASETYPE UNKNOWN_BASE_TYPE) ApiWebhooksObjectAttributeUpdatedRequest {
	r.uNKNOWNBASETYPE = &uNKNOWNBASETYPE
	return r
}

func (r ApiWebhooksObjectAttributeUpdatedRequest) Execute() (*http.Response, error) {
	return r.ApiService.WebhooksObjectAttributeUpdatedExecute(r)
}

/*
WebhooksObjectAttributeUpdated object-attribute.updated

This event is fired whenever an object attribute is updated (e.g. when renaming the "Rating" attribute to "Score" on the company object).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWebhooksObjectAttributeUpdatedRequest
*/
func (a *ObjectAttributeEventsAPIService) WebhooksObjectAttributeUpdated(ctx context.Context) ApiWebhooksObjectAttributeUpdatedRequest {
	return ApiWebhooksObjectAttributeUpdatedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *ObjectAttributeEventsAPIService) WebhooksObjectAttributeUpdatedExecute(r ApiWebhooksObjectAttributeUpdatedRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ObjectAttributeEventsAPIService.WebhooksObjectAttributeUpdated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/object-attribute.updated"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.uNKNOWNBASETYPE
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
