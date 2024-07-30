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


type RoutesAPI interface {

	/*
	CreateRoute Create a route

	Create a route

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateRouteRequest
	*/
	CreateRoute(ctx context.Context) ApiCreateRouteRequest

	// CreateRouteExecute executes the request
	//  @return ResponsesCreatedSuccessResponseCreateRouteRes
	CreateRouteExecute(r ApiCreateRouteRequest) (*ResponsesCreatedSuccessResponseCreateRouteRes, *http.Response, error)

	/*
	DeleteRoute Delete a route

	Delete a route

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uuid Route UUID
	@return ApiDeleteRouteRequest
	*/
	DeleteRoute(ctx context.Context, uuid string) ApiDeleteRouteRequest

	// DeleteRouteExecute executes the request
	//  @return ResponsesDefaultSuccessResponseWithoutData
	DeleteRouteExecute(r ApiDeleteRouteRequest) (*ResponsesDefaultSuccessResponseWithoutData, *http.Response, error)

	/*
	ListRoutes Get list routes

	Get list routes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListRoutesRequest
	*/
	ListRoutes(ctx context.Context) ApiListRoutesRequest

	// ListRoutesExecute executes the request
	//  @return ResponsesDefaultSuccessResponseListRouteRes
	ListRoutesExecute(r ApiListRoutesRequest) (*ResponsesDefaultSuccessResponseListRouteRes, *http.Response, error)

	/*
	UpdateRoute Update a route

	Update a route

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param uuid UUID
	@return ApiUpdateRouteRequest
	*/
	UpdateRoute(ctx context.Context, uuid string) ApiUpdateRouteRequest

	// UpdateRouteExecute executes the request
	//  @return ResponsesDefaultSuccessResponseWithoutData
	UpdateRouteExecute(r ApiUpdateRouteRequest) (*ResponsesDefaultSuccessResponseWithoutData, *http.Response, error)
}

// RoutesAPIService RoutesAPI service
type RoutesAPIService service

type ApiCreateRouteRequest struct {
	ctx context.Context
	ApiService RoutesAPI
	route *CreateRouteReq
}

// Route request
func (r ApiCreateRouteRequest) Route(route CreateRouteReq) ApiCreateRouteRequest {
	r.route = &route
	return r
}

func (r ApiCreateRouteRequest) Execute() (*ResponsesCreatedSuccessResponseCreateRouteRes, *http.Response, error) {
	return r.ApiService.CreateRouteExecute(r)
}

/*
CreateRoute Create a route

Create a route

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateRouteRequest
*/
func (a *RoutesAPIService) CreateRoute(ctx context.Context) ApiCreateRouteRequest {
	return ApiCreateRouteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResponsesCreatedSuccessResponseCreateRouteRes
func (a *RoutesAPIService) CreateRouteExecute(r ApiCreateRouteRequest) (*ResponsesCreatedSuccessResponseCreateRouteRes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponsesCreatedSuccessResponseCreateRouteRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutesAPIService.CreateRoute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/routes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.route == nil {
		return localVarReturnValue, nil, reportError("route is required and must be specified")
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
	localVarPostBody = r.route
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

type ApiDeleteRouteRequest struct {
	ctx context.Context
	ApiService RoutesAPI
	uuid string
}

func (r ApiDeleteRouteRequest) Execute() (*ResponsesDefaultSuccessResponseWithoutData, *http.Response, error) {
	return r.ApiService.DeleteRouteExecute(r)
}

/*
DeleteRoute Delete a route

Delete a route

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid Route UUID
 @return ApiDeleteRouteRequest
*/
func (a *RoutesAPIService) DeleteRoute(ctx context.Context, uuid string) ApiDeleteRouteRequest {
	return ApiDeleteRouteRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return ResponsesDefaultSuccessResponseWithoutData
func (a *RoutesAPIService) DeleteRouteExecute(r ApiDeleteRouteRequest) (*ResponsesDefaultSuccessResponseWithoutData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponsesDefaultSuccessResponseWithoutData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutesAPIService.DeleteRoute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/routes/{uuid}"
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

type ApiListRoutesRequest struct {
	ctx context.Context
	ApiService RoutesAPI
	host *string
	loadBalancerStrategy *string
	path *string
	uuid *string
}

func (r ApiListRoutesRequest) Host(host string) ApiListRoutesRequest {
	r.host = &host
	return r
}

func (r ApiListRoutesRequest) LoadBalancerStrategy(loadBalancerStrategy string) ApiListRoutesRequest {
	r.loadBalancerStrategy = &loadBalancerStrategy
	return r
}

func (r ApiListRoutesRequest) Path(path string) ApiListRoutesRequest {
	r.path = &path
	return r
}

func (r ApiListRoutesRequest) Uuid(uuid string) ApiListRoutesRequest {
	r.uuid = &uuid
	return r
}

func (r ApiListRoutesRequest) Execute() (*ResponsesDefaultSuccessResponseListRouteRes, *http.Response, error) {
	return r.ApiService.ListRoutesExecute(r)
}

/*
ListRoutes Get list routes

Get list routes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListRoutesRequest
*/
func (a *RoutesAPIService) ListRoutes(ctx context.Context) ApiListRoutesRequest {
	return ApiListRoutesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ResponsesDefaultSuccessResponseListRouteRes
func (a *RoutesAPIService) ListRoutesExecute(r ApiListRoutesRequest) (*ResponsesDefaultSuccessResponseListRouteRes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponsesDefaultSuccessResponseListRouteRes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutesAPIService.ListRoutes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/routes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.host != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "host", r.host, "")
	}
	if r.loadBalancerStrategy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "loadBalancerStrategy", r.loadBalancerStrategy, "")
	}
	if r.path != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "path", r.path, "")
	}
	if r.uuid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "uuid", r.uuid, "")
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

type ApiUpdateRouteRequest struct {
	ctx context.Context
	ApiService RoutesAPI
	uuid string
	route *UpdateRouteReq
}

// Route request
func (r ApiUpdateRouteRequest) Route(route UpdateRouteReq) ApiUpdateRouteRequest {
	r.route = &route
	return r
}

func (r ApiUpdateRouteRequest) Execute() (*ResponsesDefaultSuccessResponseWithoutData, *http.Response, error) {
	return r.ApiService.UpdateRouteExecute(r)
}

/*
UpdateRoute Update a route

Update a route

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param uuid UUID
 @return ApiUpdateRouteRequest
*/
func (a *RoutesAPIService) UpdateRoute(ctx context.Context, uuid string) ApiUpdateRouteRequest {
	return ApiUpdateRouteRequest{
		ApiService: a,
		ctx: ctx,
		uuid: uuid,
	}
}

// Execute executes the request
//  @return ResponsesDefaultSuccessResponseWithoutData
func (a *RoutesAPIService) UpdateRouteExecute(r ApiUpdateRouteRequest) (*ResponsesDefaultSuccessResponseWithoutData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResponsesDefaultSuccessResponseWithoutData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RoutesAPIService.UpdateRoute")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/routes/{uuid}"
	localVarPath = strings.Replace(localVarPath, "{"+"uuid"+"}", url.PathEscape(parameterValueToString(r.uuid, "uuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.route == nil {
		return localVarReturnValue, nil, reportError("route is required and must be specified")
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
	localVarPostBody = r.route
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
