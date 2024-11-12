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


// GlobalAttributesAPIService GlobalAttributesAPI service
type GlobalAttributesAPIService service

type ApiGlobalAttributesGetRequest struct {
	ctx context.Context
	ApiService *GlobalAttributesAPIService
}

func (r ApiGlobalAttributesGetRequest) Execute() (*GlobalAttributesGet200Response, *http.Response, error) {
	return r.ApiService.GlobalAttributesGetExecute(r)
}

/*
GlobalAttributesGet Method for GlobalAttributesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGlobalAttributesGetRequest
*/
func (a *GlobalAttributesAPIService) GlobalAttributesGet(ctx context.Context) ApiGlobalAttributesGetRequest {
	return ApiGlobalAttributesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GlobalAttributesGet200Response
func (a *GlobalAttributesAPIService) GlobalAttributesGetExecute(r ApiGlobalAttributesGetRequest) (*GlobalAttributesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalAttributesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalAttributesAPIService.GlobalAttributesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/global_attributes"

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

type ApiGlobalAttributesGlobalAttributeIdDeleteRequest struct {
	ctx context.Context
	ApiService *GlobalAttributesAPIService
	globalAttributeId int64
}

func (r ApiGlobalAttributesGlobalAttributeIdDeleteRequest) Execute() (*GlobalAttributes, *http.Response, error) {
	return r.ApiService.GlobalAttributesGlobalAttributeIdDeleteExecute(r)
}

/*
GlobalAttributesGlobalAttributeIdDelete Method for GlobalAttributesGlobalAttributeIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param globalAttributeId Model id
 @return ApiGlobalAttributesGlobalAttributeIdDeleteRequest
*/
func (a *GlobalAttributesAPIService) GlobalAttributesGlobalAttributeIdDelete(ctx context.Context, globalAttributeId int64) ApiGlobalAttributesGlobalAttributeIdDeleteRequest {
	return ApiGlobalAttributesGlobalAttributeIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		globalAttributeId: globalAttributeId,
	}
}

// Execute executes the request
//  @return GlobalAttributes
func (a *GlobalAttributesAPIService) GlobalAttributesGlobalAttributeIdDeleteExecute(r ApiGlobalAttributesGlobalAttributeIdDeleteRequest) (*GlobalAttributes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalAttributes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalAttributesAPIService.GlobalAttributesGlobalAttributeIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/global_attributes/{global_attribute_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"global_attribute_id"+"}", url.PathEscape(parameterValueToString(r.globalAttributeId, "globalAttributeId")), -1)

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

type ApiGlobalAttributesGlobalAttributeIdGetRequest struct {
	ctx context.Context
	ApiService *GlobalAttributesAPIService
	globalAttributeId int64
}

func (r ApiGlobalAttributesGlobalAttributeIdGetRequest) Execute() (*GlobalAttributes, *http.Response, error) {
	return r.ApiService.GlobalAttributesGlobalAttributeIdGetExecute(r)
}

/*
GlobalAttributesGlobalAttributeIdGet Method for GlobalAttributesGlobalAttributeIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param globalAttributeId Model id
 @return ApiGlobalAttributesGlobalAttributeIdGetRequest
*/
func (a *GlobalAttributesAPIService) GlobalAttributesGlobalAttributeIdGet(ctx context.Context, globalAttributeId int64) ApiGlobalAttributesGlobalAttributeIdGetRequest {
	return ApiGlobalAttributesGlobalAttributeIdGetRequest{
		ApiService: a,
		ctx: ctx,
		globalAttributeId: globalAttributeId,
	}
}

// Execute executes the request
//  @return GlobalAttributes
func (a *GlobalAttributesAPIService) GlobalAttributesGlobalAttributeIdGetExecute(r ApiGlobalAttributesGlobalAttributeIdGetRequest) (*GlobalAttributes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalAttributes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalAttributesAPIService.GlobalAttributesGlobalAttributeIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/global_attributes/{global_attribute_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"global_attribute_id"+"}", url.PathEscape(parameterValueToString(r.globalAttributeId, "globalAttributeId")), -1)

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

type ApiGlobalAttributesGlobalAttributeIdPutRequest struct {
	ctx context.Context
	ApiService *GlobalAttributesAPIService
	globalAttributeId int64
	globalAttributesPut *GlobalAttributesPut
}

func (r ApiGlobalAttributesGlobalAttributeIdPutRequest) GlobalAttributesPut(globalAttributesPut GlobalAttributesPut) ApiGlobalAttributesGlobalAttributeIdPutRequest {
	r.globalAttributesPut = &globalAttributesPut
	return r
}

func (r ApiGlobalAttributesGlobalAttributeIdPutRequest) Execute() (*GlobalAttributes, *http.Response, error) {
	return r.ApiService.GlobalAttributesGlobalAttributeIdPutExecute(r)
}

/*
GlobalAttributesGlobalAttributeIdPut Method for GlobalAttributesGlobalAttributeIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param globalAttributeId Model id
 @return ApiGlobalAttributesGlobalAttributeIdPutRequest
*/
func (a *GlobalAttributesAPIService) GlobalAttributesGlobalAttributeIdPut(ctx context.Context, globalAttributeId int64) ApiGlobalAttributesGlobalAttributeIdPutRequest {
	return ApiGlobalAttributesGlobalAttributeIdPutRequest{
		ApiService: a,
		ctx: ctx,
		globalAttributeId: globalAttributeId,
	}
}

// Execute executes the request
//  @return GlobalAttributes
func (a *GlobalAttributesAPIService) GlobalAttributesGlobalAttributeIdPutExecute(r ApiGlobalAttributesGlobalAttributeIdPutRequest) (*GlobalAttributes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalAttributes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalAttributesAPIService.GlobalAttributesGlobalAttributeIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/global_attributes/{global_attribute_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"global_attribute_id"+"}", url.PathEscape(parameterValueToString(r.globalAttributeId, "globalAttributeId")), -1)

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
	localVarPostBody = r.globalAttributesPut
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

type ApiGlobalAttributesPostRequest struct {
	ctx context.Context
	ApiService *GlobalAttributesAPIService
	globalAttributesPostRequest *GlobalAttributesPostRequest
}

func (r ApiGlobalAttributesPostRequest) GlobalAttributesPostRequest(globalAttributesPostRequest GlobalAttributesPostRequest) ApiGlobalAttributesPostRequest {
	r.globalAttributesPostRequest = &globalAttributesPostRequest
	return r
}

func (r ApiGlobalAttributesPostRequest) Execute() (*GlobalAttributesGet200Response, *http.Response, error) {
	return r.ApiService.GlobalAttributesPostExecute(r)
}

/*
GlobalAttributesPost Method for GlobalAttributesPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGlobalAttributesPostRequest
*/
func (a *GlobalAttributesAPIService) GlobalAttributesPost(ctx context.Context) ApiGlobalAttributesPostRequest {
	return ApiGlobalAttributesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GlobalAttributesGet200Response
func (a *GlobalAttributesAPIService) GlobalAttributesPostExecute(r ApiGlobalAttributesPostRequest) (*GlobalAttributesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GlobalAttributesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GlobalAttributesAPIService.GlobalAttributesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/global_attributes"

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
	localVarPostBody = r.globalAttributesPostRequest
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
