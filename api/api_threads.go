/*
Attio API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
Contact: support@attio.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ThreadsAPIService ThreadsAPI service
type ThreadsAPIService service

type ApiV2ThreadsGetRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	recordId *string
	object *string
	entryId *string
	list *string
	limit *int32
	offset *int32
}

func (r ApiV2ThreadsGetRequest) RecordId(recordId string) ApiV2ThreadsGetRequest {
	r.recordId = &recordId
	return r
}

func (r ApiV2ThreadsGetRequest) Object(object string) ApiV2ThreadsGetRequest {
	r.object = &object
	return r
}

func (r ApiV2ThreadsGetRequest) EntryId(entryId string) ApiV2ThreadsGetRequest {
	r.entryId = &entryId
	return r
}

func (r ApiV2ThreadsGetRequest) List(list string) ApiV2ThreadsGetRequest {
	r.list = &list
	return r
}

func (r ApiV2ThreadsGetRequest) Limit(limit int32) ApiV2ThreadsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiV2ThreadsGetRequest) Offset(offset int32) ApiV2ThreadsGetRequest {
	r.offset = &offset
	return r
}

func (r ApiV2ThreadsGetRequest) Execute() (*V2ThreadsGet200Response, *http.Response, error) {
	return r.ApiService.V2ThreadsGetExecute(r)
}

/*
V2ThreadsGet List threads

List threads of comments on a record or list entry.

To view threads on records, you will need the `object_configuration:read` and `record_permission:read` scopes.

To view threads on list entries, you will need the `list_configuration:read` and `list_entry:read` scopes.

Required scopes: `comment:read`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiV2ThreadsGetRequest
*/
func (a *ThreadsAPIService) V2ThreadsGet(ctx context.Context) ApiV2ThreadsGetRequest {
	return ApiV2ThreadsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return V2ThreadsGet200Response
func (a *ThreadsAPIService) V2ThreadsGetExecute(r ApiV2ThreadsGetRequest) (*V2ThreadsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V2ThreadsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.V2ThreadsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/threads"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.recordId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "record_id", r.recordId, "form", "")
	}
	if r.object != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "object", r.object, "form", "")
	}
	if r.entryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "entry_id", r.entryId, "form", "")
	}
	if r.list != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "list", r.list, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiV2ThreadsThreadIdGetRequest struct {
	ctx context.Context
	ApiService *ThreadsAPIService
	threadId string
}

func (r ApiV2ThreadsThreadIdGetRequest) Execute() (*V2ThreadsThreadIdGet200Response, *http.Response, error) {
	return r.ApiService.V2ThreadsThreadIdGetExecute(r)
}

/*
V2ThreadsThreadIdGet Get a thread

Get all comments in a thread.

To view threads on records, you will need the `object_configuration:read` and `record_permission:read` scopes.

To view threads on list entries, you will need the `list_configuration:read` and `list_entry:read` scopes.

Required scopes: `comment:read`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param threadId
 @return ApiV2ThreadsThreadIdGetRequest
*/
func (a *ThreadsAPIService) V2ThreadsThreadIdGet(ctx context.Context, threadId string) ApiV2ThreadsThreadIdGetRequest {
	return ApiV2ThreadsThreadIdGetRequest{
		ApiService: a,
		ctx: ctx,
		threadId: threadId,
	}
}

// Execute executes the request
//  @return V2ThreadsThreadIdGet200Response
func (a *ThreadsAPIService) V2ThreadsThreadIdGetExecute(r ApiV2ThreadsThreadIdGetRequest) (*V2ThreadsThreadIdGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *V2ThreadsThreadIdGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThreadsAPIService.V2ThreadsThreadIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/threads/{thread_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"thread_id"+"}", url.PathEscape(parameterValueToString(r.threadId, "threadId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v V2ThreadsThreadIdGet404Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
