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


// RcwRequestCategoriesAPIService RcwRequestCategoriesAPI service
type RcwRequestCategoriesAPIService service

type ApiRcwRequestCategoriesGetRequest struct {
	ctx context.Context
	ApiService *RcwRequestCategoriesAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiRcwRequestCategoriesGetRequest) Include(include []string) ApiRcwRequestCategoriesGetRequest {
	r.include = &include
	return r
}

func (r ApiRcwRequestCategoriesGetRequest) Execute() (*RcwRequestCategoriesGet200Response, *http.Response, error) {
	return r.ApiService.RcwRequestCategoriesGetExecute(r)
}

/*
RcwRequestCategoriesGet Method for RcwRequestCategoriesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRcwRequestCategoriesGetRequest
*/
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesGet(ctx context.Context) ApiRcwRequestCategoriesGetRequest {
	return ApiRcwRequestCategoriesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RcwRequestCategoriesGet200Response
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesGetExecute(r ApiRcwRequestCategoriesGetRequest) (*RcwRequestCategoriesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RcwRequestCategoriesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RcwRequestCategoriesAPIService.RcwRequestCategoriesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rcw_request_categories"

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

type ApiRcwRequestCategoriesPostRequest struct {
	ctx context.Context
	ApiService *RcwRequestCategoriesAPIService
	rcwRequestCategoriesPostRequest *RcwRequestCategoriesPostRequest
}

func (r ApiRcwRequestCategoriesPostRequest) RcwRequestCategoriesPostRequest(rcwRequestCategoriesPostRequest RcwRequestCategoriesPostRequest) ApiRcwRequestCategoriesPostRequest {
	r.rcwRequestCategoriesPostRequest = &rcwRequestCategoriesPostRequest
	return r
}

func (r ApiRcwRequestCategoriesPostRequest) Execute() (*RcwRequestCategoriesGet200Response, *http.Response, error) {
	return r.ApiService.RcwRequestCategoriesPostExecute(r)
}

/*
RcwRequestCategoriesPost Method for RcwRequestCategoriesPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRcwRequestCategoriesPostRequest
*/
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesPost(ctx context.Context) ApiRcwRequestCategoriesPostRequest {
	return ApiRcwRequestCategoriesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RcwRequestCategoriesGet200Response
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesPostExecute(r ApiRcwRequestCategoriesPostRequest) (*RcwRequestCategoriesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RcwRequestCategoriesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RcwRequestCategoriesAPIService.RcwRequestCategoriesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rcw_request_categories"

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
	localVarPostBody = r.rcwRequestCategoriesPostRequest
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

type ApiRcwRequestCategoriesRcwRequestCategoryIdDeleteRequest struct {
	ctx context.Context
	ApiService *RcwRequestCategoriesAPIService
	rcwRequestCategoryId int64
}

func (r ApiRcwRequestCategoriesRcwRequestCategoryIdDeleteRequest) Execute() (*RcwRequestCategories, *http.Response, error) {
	return r.ApiService.RcwRequestCategoriesRcwRequestCategoryIdDeleteExecute(r)
}

/*
RcwRequestCategoriesRcwRequestCategoryIdDelete Method for RcwRequestCategoriesRcwRequestCategoryIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rcwRequestCategoryId Model id
 @return ApiRcwRequestCategoriesRcwRequestCategoryIdDeleteRequest
*/
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesRcwRequestCategoryIdDelete(ctx context.Context, rcwRequestCategoryId int64) ApiRcwRequestCategoriesRcwRequestCategoryIdDeleteRequest {
	return ApiRcwRequestCategoriesRcwRequestCategoryIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		rcwRequestCategoryId: rcwRequestCategoryId,
	}
}

// Execute executes the request
//  @return RcwRequestCategories
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesRcwRequestCategoryIdDeleteExecute(r ApiRcwRequestCategoriesRcwRequestCategoryIdDeleteRequest) (*RcwRequestCategories, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RcwRequestCategories
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RcwRequestCategoriesAPIService.RcwRequestCategoriesRcwRequestCategoryIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rcw_request_categories/{rcw_request_category_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"rcw_request_category_id"+"}", url.PathEscape(parameterValueToString(r.rcwRequestCategoryId, "rcwRequestCategoryId")), -1)

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

type ApiRcwRequestCategoriesRcwRequestCategoryIdGetRequest struct {
	ctx context.Context
	ApiService *RcwRequestCategoriesAPIService
	rcwRequestCategoryId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiRcwRequestCategoriesRcwRequestCategoryIdGetRequest) Include(include []string) ApiRcwRequestCategoriesRcwRequestCategoryIdGetRequest {
	r.include = &include
	return r
}

func (r ApiRcwRequestCategoriesRcwRequestCategoryIdGetRequest) Execute() (*RcwRequestCategories, *http.Response, error) {
	return r.ApiService.RcwRequestCategoriesRcwRequestCategoryIdGetExecute(r)
}

/*
RcwRequestCategoriesRcwRequestCategoryIdGet Method for RcwRequestCategoriesRcwRequestCategoryIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rcwRequestCategoryId Model id
 @return ApiRcwRequestCategoriesRcwRequestCategoryIdGetRequest
*/
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesRcwRequestCategoryIdGet(ctx context.Context, rcwRequestCategoryId int64) ApiRcwRequestCategoriesRcwRequestCategoryIdGetRequest {
	return ApiRcwRequestCategoriesRcwRequestCategoryIdGetRequest{
		ApiService: a,
		ctx: ctx,
		rcwRequestCategoryId: rcwRequestCategoryId,
	}
}

// Execute executes the request
//  @return RcwRequestCategories
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesRcwRequestCategoryIdGetExecute(r ApiRcwRequestCategoriesRcwRequestCategoryIdGetRequest) (*RcwRequestCategories, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RcwRequestCategories
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RcwRequestCategoriesAPIService.RcwRequestCategoriesRcwRequestCategoryIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rcw_request_categories/{rcw_request_category_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"rcw_request_category_id"+"}", url.PathEscape(parameterValueToString(r.rcwRequestCategoryId, "rcwRequestCategoryId")), -1)

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

type ApiRcwRequestCategoriesRcwRequestCategoryIdPutRequest struct {
	ctx context.Context
	ApiService *RcwRequestCategoriesAPIService
	rcwRequestCategoryId int64
	rcwRequestCategoriesPut *RcwRequestCategoriesPut
}

func (r ApiRcwRequestCategoriesRcwRequestCategoryIdPutRequest) RcwRequestCategoriesPut(rcwRequestCategoriesPut RcwRequestCategoriesPut) ApiRcwRequestCategoriesRcwRequestCategoryIdPutRequest {
	r.rcwRequestCategoriesPut = &rcwRequestCategoriesPut
	return r
}

func (r ApiRcwRequestCategoriesRcwRequestCategoryIdPutRequest) Execute() (*RcwRequestCategories, *http.Response, error) {
	return r.ApiService.RcwRequestCategoriesRcwRequestCategoryIdPutExecute(r)
}

/*
RcwRequestCategoriesRcwRequestCategoryIdPut Method for RcwRequestCategoriesRcwRequestCategoryIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param rcwRequestCategoryId Model id
 @return ApiRcwRequestCategoriesRcwRequestCategoryIdPutRequest
*/
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesRcwRequestCategoryIdPut(ctx context.Context, rcwRequestCategoryId int64) ApiRcwRequestCategoriesRcwRequestCategoryIdPutRequest {
	return ApiRcwRequestCategoriesRcwRequestCategoryIdPutRequest{
		ApiService: a,
		ctx: ctx,
		rcwRequestCategoryId: rcwRequestCategoryId,
	}
}

// Execute executes the request
//  @return RcwRequestCategories
func (a *RcwRequestCategoriesAPIService) RcwRequestCategoriesRcwRequestCategoryIdPutExecute(r ApiRcwRequestCategoriesRcwRequestCategoryIdPutRequest) (*RcwRequestCategories, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RcwRequestCategories
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "RcwRequestCategoriesAPIService.RcwRequestCategoriesRcwRequestCategoryIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rcw_request_categories/{rcw_request_category_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"rcw_request_category_id"+"}", url.PathEscape(parameterValueToString(r.rcwRequestCategoryId, "rcwRequestCategoryId")), -1)

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
	localVarPostBody = r.rcwRequestCategoriesPut
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
