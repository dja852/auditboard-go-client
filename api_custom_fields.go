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


// CustomFieldsAPIService CustomFieldsAPI service
type CustomFieldsAPIService service

type ApiCustomFieldsCustomFieldIdDeleteRequest struct {
	ctx context.Context
	ApiService *CustomFieldsAPIService
	customFieldId int64
}

func (r ApiCustomFieldsCustomFieldIdDeleteRequest) Execute() (*CustomFields, *http.Response, error) {
	return r.ApiService.CustomFieldsCustomFieldIdDeleteExecute(r)
}

/*
CustomFieldsCustomFieldIdDelete Method for CustomFieldsCustomFieldIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customFieldId Model id
 @return ApiCustomFieldsCustomFieldIdDeleteRequest
*/
func (a *CustomFieldsAPIService) CustomFieldsCustomFieldIdDelete(ctx context.Context, customFieldId int64) ApiCustomFieldsCustomFieldIdDeleteRequest {
	return ApiCustomFieldsCustomFieldIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		customFieldId: customFieldId,
	}
}

// Execute executes the request
//  @return CustomFields
func (a *CustomFieldsAPIService) CustomFieldsCustomFieldIdDeleteExecute(r ApiCustomFieldsCustomFieldIdDeleteRequest) (*CustomFields, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomFields
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomFieldsAPIService.CustomFieldsCustomFieldIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom_fields/{custom_field_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_field_id"+"}", url.PathEscape(parameterValueToString(r.customFieldId, "customFieldId")), -1)

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

type ApiCustomFieldsCustomFieldIdGetRequest struct {
	ctx context.Context
	ApiService *CustomFieldsAPIService
	customFieldId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiCustomFieldsCustomFieldIdGetRequest) Include(include []string) ApiCustomFieldsCustomFieldIdGetRequest {
	r.include = &include
	return r
}

func (r ApiCustomFieldsCustomFieldIdGetRequest) Execute() (*CustomFields, *http.Response, error) {
	return r.ApiService.CustomFieldsCustomFieldIdGetExecute(r)
}

/*
CustomFieldsCustomFieldIdGet Method for CustomFieldsCustomFieldIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customFieldId Model id
 @return ApiCustomFieldsCustomFieldIdGetRequest
*/
func (a *CustomFieldsAPIService) CustomFieldsCustomFieldIdGet(ctx context.Context, customFieldId int64) ApiCustomFieldsCustomFieldIdGetRequest {
	return ApiCustomFieldsCustomFieldIdGetRequest{
		ApiService: a,
		ctx: ctx,
		customFieldId: customFieldId,
	}
}

// Execute executes the request
//  @return CustomFields
func (a *CustomFieldsAPIService) CustomFieldsCustomFieldIdGetExecute(r ApiCustomFieldsCustomFieldIdGetRequest) (*CustomFields, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomFields
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomFieldsAPIService.CustomFieldsCustomFieldIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom_fields/{custom_field_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_field_id"+"}", url.PathEscape(parameterValueToString(r.customFieldId, "customFieldId")), -1)

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

type ApiCustomFieldsCustomFieldIdPutRequest struct {
	ctx context.Context
	ApiService *CustomFieldsAPIService
	customFieldId int64
	customFieldsPut *CustomFieldsPut
}

func (r ApiCustomFieldsCustomFieldIdPutRequest) CustomFieldsPut(customFieldsPut CustomFieldsPut) ApiCustomFieldsCustomFieldIdPutRequest {
	r.customFieldsPut = &customFieldsPut
	return r
}

func (r ApiCustomFieldsCustomFieldIdPutRequest) Execute() (*CustomFields, *http.Response, error) {
	return r.ApiService.CustomFieldsCustomFieldIdPutExecute(r)
}

/*
CustomFieldsCustomFieldIdPut Method for CustomFieldsCustomFieldIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param customFieldId Model id
 @return ApiCustomFieldsCustomFieldIdPutRequest
*/
func (a *CustomFieldsAPIService) CustomFieldsCustomFieldIdPut(ctx context.Context, customFieldId int64) ApiCustomFieldsCustomFieldIdPutRequest {
	return ApiCustomFieldsCustomFieldIdPutRequest{
		ApiService: a,
		ctx: ctx,
		customFieldId: customFieldId,
	}
}

// Execute executes the request
//  @return CustomFields
func (a *CustomFieldsAPIService) CustomFieldsCustomFieldIdPutExecute(r ApiCustomFieldsCustomFieldIdPutRequest) (*CustomFields, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomFields
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomFieldsAPIService.CustomFieldsCustomFieldIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom_fields/{custom_field_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"custom_field_id"+"}", url.PathEscape(parameterValueToString(r.customFieldId, "customFieldId")), -1)

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
	localVarPostBody = r.customFieldsPut
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

type ApiCustomFieldsGetRequest struct {
	ctx context.Context
	ApiService *CustomFieldsAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiCustomFieldsGetRequest) Include(include []string) ApiCustomFieldsGetRequest {
	r.include = &include
	return r
}

func (r ApiCustomFieldsGetRequest) Execute() (*CustomFieldsGet200Response, *http.Response, error) {
	return r.ApiService.CustomFieldsGetExecute(r)
}

/*
CustomFieldsGet Method for CustomFieldsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCustomFieldsGetRequest
*/
func (a *CustomFieldsAPIService) CustomFieldsGet(ctx context.Context) ApiCustomFieldsGetRequest {
	return ApiCustomFieldsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomFieldsGet200Response
func (a *CustomFieldsAPIService) CustomFieldsGetExecute(r ApiCustomFieldsGetRequest) (*CustomFieldsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomFieldsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomFieldsAPIService.CustomFieldsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom_fields"

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

type ApiCustomFieldsPostRequest struct {
	ctx context.Context
	ApiService *CustomFieldsAPIService
	customFieldsPostRequest *CustomFieldsPostRequest
}

func (r ApiCustomFieldsPostRequest) CustomFieldsPostRequest(customFieldsPostRequest CustomFieldsPostRequest) ApiCustomFieldsPostRequest {
	r.customFieldsPostRequest = &customFieldsPostRequest
	return r
}

func (r ApiCustomFieldsPostRequest) Execute() (*CustomFieldsGet200Response, *http.Response, error) {
	return r.ApiService.CustomFieldsPostExecute(r)
}

/*
CustomFieldsPost Method for CustomFieldsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCustomFieldsPostRequest
*/
func (a *CustomFieldsAPIService) CustomFieldsPost(ctx context.Context) ApiCustomFieldsPostRequest {
	return ApiCustomFieldsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CustomFieldsGet200Response
func (a *CustomFieldsAPIService) CustomFieldsPostExecute(r ApiCustomFieldsPostRequest) (*CustomFieldsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CustomFieldsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomFieldsAPIService.CustomFieldsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/custom_fields"

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
	localVarPostBody = r.customFieldsPostRequest
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