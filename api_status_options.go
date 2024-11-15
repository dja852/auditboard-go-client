/*
AuditBoard Developer Portal API Documentation

External API reference documentation.

API version: 23.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auditboard

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// StatusOptionsAPIService StatusOptionsAPI service
type StatusOptionsAPIService service

type ApiStatusOptionsGetRequest struct {
	ctx context.Context
	ApiService *StatusOptionsAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiStatusOptionsGetRequest) Include(include []string) ApiStatusOptionsGetRequest {
	r.include = &include
	return r
}

func (r ApiStatusOptionsGetRequest) Execute() (*StatusOptionsGet200Response, *http.Response, error) {
	return r.ApiService.StatusOptionsGetExecute(r)
}

/*
StatusOptionsGet Method for StatusOptionsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStatusOptionsGetRequest
*/
func (a *StatusOptionsAPIService) StatusOptionsGet(ctx context.Context) ApiStatusOptionsGetRequest {
	return ApiStatusOptionsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StatusOptionsGet200Response
func (a *StatusOptionsAPIService) StatusOptionsGetExecute(r ApiStatusOptionsGetRequest) (*StatusOptionsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StatusOptionsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusOptionsAPIService.StatusOptionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/status_options"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "csv")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiStatusOptionsPostRequest struct {
	ctx context.Context
	ApiService *StatusOptionsAPIService
	statusOptionsPostRequest *StatusOptionsPostRequest
}

func (r ApiStatusOptionsPostRequest) StatusOptionsPostRequest(statusOptionsPostRequest StatusOptionsPostRequest) ApiStatusOptionsPostRequest {
	r.statusOptionsPostRequest = &statusOptionsPostRequest
	return r
}

func (r ApiStatusOptionsPostRequest) Execute() (*StatusOptionsGet200Response, *http.Response, error) {
	return r.ApiService.StatusOptionsPostExecute(r)
}

/*
StatusOptionsPost Method for StatusOptionsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiStatusOptionsPostRequest
*/
func (a *StatusOptionsAPIService) StatusOptionsPost(ctx context.Context) ApiStatusOptionsPostRequest {
	return ApiStatusOptionsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return StatusOptionsGet200Response
func (a *StatusOptionsAPIService) StatusOptionsPostExecute(r ApiStatusOptionsPostRequest) (*StatusOptionsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StatusOptionsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusOptionsAPIService.StatusOptionsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/status_options"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.statusOptionsPostRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiStatusOptionsStatusOptionIdDeleteRequest struct {
	ctx context.Context
	ApiService *StatusOptionsAPIService
	statusOptionId int64
}

func (r ApiStatusOptionsStatusOptionIdDeleteRequest) Execute() (*StatusOptions, *http.Response, error) {
	return r.ApiService.StatusOptionsStatusOptionIdDeleteExecute(r)
}

/*
StatusOptionsStatusOptionIdDelete Method for StatusOptionsStatusOptionIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param statusOptionId Model id
 @return ApiStatusOptionsStatusOptionIdDeleteRequest
*/
func (a *StatusOptionsAPIService) StatusOptionsStatusOptionIdDelete(ctx context.Context, statusOptionId int64) ApiStatusOptionsStatusOptionIdDeleteRequest {
	return ApiStatusOptionsStatusOptionIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		statusOptionId: statusOptionId,
	}
}

// Execute executes the request
//  @return StatusOptions
func (a *StatusOptionsAPIService) StatusOptionsStatusOptionIdDeleteExecute(r ApiStatusOptionsStatusOptionIdDeleteRequest) (*StatusOptions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StatusOptions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusOptionsAPIService.StatusOptionsStatusOptionIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/status_options/{status_option_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"status_option_id"+"}", url.PathEscape(parameterValueToString(r.statusOptionId, "statusOptionId")), -1)

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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiStatusOptionsStatusOptionIdGetRequest struct {
	ctx context.Context
	ApiService *StatusOptionsAPIService
	statusOptionId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiStatusOptionsStatusOptionIdGetRequest) Include(include []string) ApiStatusOptionsStatusOptionIdGetRequest {
	r.include = &include
	return r
}

func (r ApiStatusOptionsStatusOptionIdGetRequest) Execute() (*StatusOptions, *http.Response, error) {
	return r.ApiService.StatusOptionsStatusOptionIdGetExecute(r)
}

/*
StatusOptionsStatusOptionIdGet Method for StatusOptionsStatusOptionIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param statusOptionId Model id
 @return ApiStatusOptionsStatusOptionIdGetRequest
*/
func (a *StatusOptionsAPIService) StatusOptionsStatusOptionIdGet(ctx context.Context, statusOptionId int64) ApiStatusOptionsStatusOptionIdGetRequest {
	return ApiStatusOptionsStatusOptionIdGetRequest{
		ApiService: a,
		ctx: ctx,
		statusOptionId: statusOptionId,
	}
}

// Execute executes the request
//  @return StatusOptions
func (a *StatusOptionsAPIService) StatusOptionsStatusOptionIdGetExecute(r ApiStatusOptionsStatusOptionIdGetRequest) (*StatusOptions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StatusOptions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusOptionsAPIService.StatusOptionsStatusOptionIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/status_options/{status_option_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"status_option_id"+"}", url.PathEscape(parameterValueToString(r.statusOptionId, "statusOptionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "include", r.include, "form", "csv")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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

type ApiStatusOptionsStatusOptionIdPutRequest struct {
	ctx context.Context
	ApiService *StatusOptionsAPIService
	statusOptionId int64
	statusOptionsPut *StatusOptionsPut
}

func (r ApiStatusOptionsStatusOptionIdPutRequest) StatusOptionsPut(statusOptionsPut StatusOptionsPut) ApiStatusOptionsStatusOptionIdPutRequest {
	r.statusOptionsPut = &statusOptionsPut
	return r
}

func (r ApiStatusOptionsStatusOptionIdPutRequest) Execute() (*StatusOptions, *http.Response, error) {
	return r.ApiService.StatusOptionsStatusOptionIdPutExecute(r)
}

/*
StatusOptionsStatusOptionIdPut Method for StatusOptionsStatusOptionIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param statusOptionId Model id
 @return ApiStatusOptionsStatusOptionIdPutRequest
*/
func (a *StatusOptionsAPIService) StatusOptionsStatusOptionIdPut(ctx context.Context, statusOptionId int64) ApiStatusOptionsStatusOptionIdPutRequest {
	return ApiStatusOptionsStatusOptionIdPutRequest{
		ApiService: a,
		ctx: ctx,
		statusOptionId: statusOptionId,
	}
}

// Execute executes the request
//  @return StatusOptions
func (a *StatusOptionsAPIService) StatusOptionsStatusOptionIdPutExecute(r ApiStatusOptionsStatusOptionIdPutRequest) (*StatusOptions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StatusOptions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatusOptionsAPIService.StatusOptionsStatusOptionIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/status_options/{status_option_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"status_option_id"+"}", url.PathEscape(parameterValueToString(r.statusOptionId, "statusOptionId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.statusOptionsPut
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["Bearer"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
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
