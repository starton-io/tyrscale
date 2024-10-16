/*
Tyrscale Manager API

This is the manager service for Tyrscale

API version: 1.0
Contact: support@starton.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tyrscalesdkgo

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type UpstreamsAPI interface {

	/*
	DeleteUpstream Delete a upstream

	Delete a upstream

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param routeUuid Route UUID
	@param uuid Upstream UUID
	@return ApiDeleteUpstreamRequest
	*/
	DeleteUpstream(ctx context.Context, routeUuid string, uuid string) ApiDeleteUpstreamRequest

	// DeleteUpstreamExecute executes the request
	//  @return ResponsesDefaultSuccessResponseUpstreamUpsertRes
	DeleteUpstreamExecute(r ApiDeleteUpstreamRequest) (*ResponsesDefaultSuccessResponseUpstreamUpsertRes, *http.Response, error)

	/*
	ListUpstreams Get list upstreams

	Get list upstreams

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param routeUuid Route UUID
	@return ApiListUpstreamsRequest
	*/
	ListUpstreams(ctx context.Context, routeUuid string) ApiListUpstreamsRequest

	// ListUpstreamsExecute executes the request
	//  @return ResponsesDefaultSuccessResponseListUpstreamRes
	ListUpstreamsExecute(r ApiListUpstreamsRequest) (*ResponsesDefaultSuccessResponseListUpstreamRes, *http.Response, error)

	/*
	UpsertUpstream Create or update a upstream

	Create or update a upstream

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param routeUuid Route UUID
	@return ApiUpsertUpstreamRequest
	*/
	UpsertUpstream(ctx context.Context, routeUuid string) ApiUpsertUpstreamRequest

	// UpsertUpstreamExecute executes the request
	//  @return ResponsesCreatedSuccessResponseUpstreamUpsertRes
	UpsertUpstreamExecute(r ApiUpsertUpstreamRequest) (*ResponsesCreatedSuccessResponseUpstreamUpsertRes, *http.Response, error)
}

// UpstreamsAPIService UpstreamsAPI service
type UpstreamsAPIService service

type ApiDeleteUpstreamRequest struct {
	ctx context.Context
	ApiService UpstreamsAPI
	routeUuid string
	uuid string
}

func (r ApiDeleteUpstreamRequest) Execute() (*ResponsesDefaultSuccessResponseUpstreamUpsertRes, *http.Response, error) {
	return r.ApiService.DeleteUpstreamExecute(r)
}

/*
DeleteUpstream Delete a upstream

Delete a upstream

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param routeUuid Route UUID
 @param uuid Upstream UUID
 @return ApiDeleteUpstreamRequest
*/
func (a *UpstreamsAPIService) DeleteUpstream(ctx context.Context, routeUuid string, uuid string) ApiDeleteUpstreamRequest {
	return ApiDeleteUpstreamRequest{
		ApiService: a,
		ctx: ctx,
		routeUuid: routeUuid,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return ResponsesDefaultSuccessResponseUpstreamUpsertRes
func (a *UpstreamsAPIService) DeleteUpstreamExecute(r ApiDeleteUpstreamRequest) (*ResponsesDefaultSuccessResponseUpstreamUpsertRes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponsesDefaultSuccessResponseUpstreamUpsertRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpstreamsAPIService.DeleteUpstream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/routes/{route_uuid}/upstreams/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"route_uuid"+"}", url.PathEscape(parameterValueToString(r.routeUuid, "routeUuid")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ResponsesBadRequestResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ResponsesInternalServerErrorResponse
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

type ApiListUpstreamsRequest struct {
	ctx context.Context
	ApiService UpstreamsAPI
	routeUuid string
}

func (r ApiListUpstreamsRequest) Execute() (*ResponsesDefaultSuccessResponseListUpstreamRes, *http.Response, error) {
	return r.ApiService.ListUpstreamsExecute(r)
}

/*
ListUpstreams Get list upstreams

Get list upstreams

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param routeUuid Route UUID
 @return ApiListUpstreamsRequest
*/
func (a *UpstreamsAPIService) ListUpstreams(ctx context.Context, routeUuid string) ApiListUpstreamsRequest {
	return ApiListUpstreamsRequest{
		ApiService: a,
		ctx: ctx,
		routeUuid: routeUuid,
	}
}

// Execute executes the request
//  @return ResponsesDefaultSuccessResponseListUpstreamRes
func (a *UpstreamsAPIService) ListUpstreamsExecute(r ApiListUpstreamsRequest) (*ResponsesDefaultSuccessResponseListUpstreamRes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponsesDefaultSuccessResponseListUpstreamRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpstreamsAPIService.ListUpstreams")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/routes/{route_uuid}/upstreams"
	localVarPath = strings.Replace(localVarPath, "{"+"route_uuid"+"}", url.PathEscape(parameterValueToString(r.routeUuid, "routeUuid")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ResponsesBadRequestResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ResponsesInternalServerErrorResponse
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

type ApiUpsertUpstreamRequest struct {
	ctx context.Context
	ApiService UpstreamsAPI
	routeUuid string
	upstream *Upstream
}

// Upstream request
func (r ApiUpsertUpstreamRequest) Upstream(upstream Upstream) ApiUpsertUpstreamRequest {
	r.upstream = &upstream
	return r
}

func (r ApiUpsertUpstreamRequest) Execute() (*ResponsesCreatedSuccessResponseUpstreamUpsertRes, *http.Response, error) {
	return r.ApiService.UpsertUpstreamExecute(r)
}

/*
UpsertUpstream Create or update a upstream

Create or update a upstream

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param routeUuid Route UUID
 @return ApiUpsertUpstreamRequest
*/
func (a *UpstreamsAPIService) UpsertUpstream(ctx context.Context, routeUuid string) ApiUpsertUpstreamRequest {
	return ApiUpsertUpstreamRequest{
		ApiService: a,
		ctx: ctx,
		routeUuid: routeUuid,
	}
}

// Execute executes the request
//  @return ResponsesCreatedSuccessResponseUpstreamUpsertRes
func (a *UpstreamsAPIService) UpsertUpstreamExecute(r ApiUpsertUpstreamRequest) (*ResponsesCreatedSuccessResponseUpstreamUpsertRes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponsesCreatedSuccessResponseUpstreamUpsertRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpstreamsAPIService.UpsertUpstream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/routes/{route_uuid}/upstreams"
	localVarPath = strings.Replace(localVarPath, "{"+"route_uuid"+"}", url.PathEscape(parameterValueToString(r.routeUuid, "routeUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.upstream == nil {
		return localVarReturnValue, nil, reportError("upstream is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.upstream
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ResponsesBadRequestResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ResponsesInternalServerErrorResponse
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
