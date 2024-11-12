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


// OpsAuditSectionInstancesAPIService OpsAuditSectionInstancesAPI service
type OpsAuditSectionInstancesAPIService service

type ApiOpsAuditSectionInstancesGetRequest struct {
	ctx context.Context
	ApiService *OpsAuditSectionInstancesAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiOpsAuditSectionInstancesGetRequest) Include(include []string) ApiOpsAuditSectionInstancesGetRequest {
	r.include = &include
	return r
}

func (r ApiOpsAuditSectionInstancesGetRequest) Execute() (*OpsAuditSectionInstancesGet200Response, *http.Response, error) {
	return r.ApiService.OpsAuditSectionInstancesGetExecute(r)
}

/*
OpsAuditSectionInstancesGet Method for OpsAuditSectionInstancesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOpsAuditSectionInstancesGetRequest
*/
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesGet(ctx context.Context) ApiOpsAuditSectionInstancesGetRequest {
	return ApiOpsAuditSectionInstancesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OpsAuditSectionInstancesGet200Response
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesGetExecute(r ApiOpsAuditSectionInstancesGetRequest) (*OpsAuditSectionInstancesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpsAuditSectionInstancesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpsAuditSectionInstancesAPIService.OpsAuditSectionInstancesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ops_audit_section_instances"

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

type ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdDeleteRequest struct {
	ctx context.Context
	ApiService *OpsAuditSectionInstancesAPIService
	opsAuditSectionInstanceId int64
}

func (r ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdDeleteRequest) Execute() (*OpsAuditSectionInstances, *http.Response, error) {
	return r.ApiService.OpsAuditSectionInstancesOpsAuditSectionInstanceIdDeleteExecute(r)
}

/*
OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete Method for OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param opsAuditSectionInstanceId Model id
 @return ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdDeleteRequest
*/
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete(ctx context.Context, opsAuditSectionInstanceId int64) ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdDeleteRequest {
	return ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		opsAuditSectionInstanceId: opsAuditSectionInstanceId,
	}
}

// Execute executes the request
//  @return OpsAuditSectionInstances
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesOpsAuditSectionInstanceIdDeleteExecute(r ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdDeleteRequest) (*OpsAuditSectionInstances, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpsAuditSectionInstances
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpsAuditSectionInstancesAPIService.OpsAuditSectionInstancesOpsAuditSectionInstanceIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ops_audit_section_instances/{ops_audit_section_instance_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ops_audit_section_instance_id"+"}", url.PathEscape(parameterValueToString(r.opsAuditSectionInstanceId, "opsAuditSectionInstanceId")), -1)

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

type ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdGetRequest struct {
	ctx context.Context
	ApiService *OpsAuditSectionInstancesAPIService
	opsAuditSectionInstanceId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdGetRequest) Include(include []string) ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdGetRequest {
	r.include = &include
	return r
}

func (r ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdGetRequest) Execute() (*OpsAuditSectionInstances, *http.Response, error) {
	return r.ApiService.OpsAuditSectionInstancesOpsAuditSectionInstanceIdGetExecute(r)
}

/*
OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet Method for OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param opsAuditSectionInstanceId Model id
 @return ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdGetRequest
*/
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet(ctx context.Context, opsAuditSectionInstanceId int64) ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdGetRequest {
	return ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdGetRequest{
		ApiService: a,
		ctx: ctx,
		opsAuditSectionInstanceId: opsAuditSectionInstanceId,
	}
}

// Execute executes the request
//  @return OpsAuditSectionInstances
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesOpsAuditSectionInstanceIdGetExecute(r ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdGetRequest) (*OpsAuditSectionInstances, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpsAuditSectionInstances
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpsAuditSectionInstancesAPIService.OpsAuditSectionInstancesOpsAuditSectionInstanceIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ops_audit_section_instances/{ops_audit_section_instance_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ops_audit_section_instance_id"+"}", url.PathEscape(parameterValueToString(r.opsAuditSectionInstanceId, "opsAuditSectionInstanceId")), -1)

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

type ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdPutRequest struct {
	ctx context.Context
	ApiService *OpsAuditSectionInstancesAPIService
	opsAuditSectionInstanceId int64
	opsAuditSectionInstancesPut *OpsAuditSectionInstancesPut
}

func (r ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdPutRequest) OpsAuditSectionInstancesPut(opsAuditSectionInstancesPut OpsAuditSectionInstancesPut) ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdPutRequest {
	r.opsAuditSectionInstancesPut = &opsAuditSectionInstancesPut
	return r
}

func (r ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdPutRequest) Execute() (*OpsAuditSectionInstances, *http.Response, error) {
	return r.ApiService.OpsAuditSectionInstancesOpsAuditSectionInstanceIdPutExecute(r)
}

/*
OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut Method for OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param opsAuditSectionInstanceId Model id
 @return ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdPutRequest
*/
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut(ctx context.Context, opsAuditSectionInstanceId int64) ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdPutRequest {
	return ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdPutRequest{
		ApiService: a,
		ctx: ctx,
		opsAuditSectionInstanceId: opsAuditSectionInstanceId,
	}
}

// Execute executes the request
//  @return OpsAuditSectionInstances
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesOpsAuditSectionInstanceIdPutExecute(r ApiOpsAuditSectionInstancesOpsAuditSectionInstanceIdPutRequest) (*OpsAuditSectionInstances, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpsAuditSectionInstances
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpsAuditSectionInstancesAPIService.OpsAuditSectionInstancesOpsAuditSectionInstanceIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ops_audit_section_instances/{ops_audit_section_instance_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"ops_audit_section_instance_id"+"}", url.PathEscape(parameterValueToString(r.opsAuditSectionInstanceId, "opsAuditSectionInstanceId")), -1)

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
	localVarPostBody = r.opsAuditSectionInstancesPut
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

type ApiOpsAuditSectionInstancesPostRequest struct {
	ctx context.Context
	ApiService *OpsAuditSectionInstancesAPIService
	opsAuditSectionInstancesPostRequest *OpsAuditSectionInstancesPostRequest
}

func (r ApiOpsAuditSectionInstancesPostRequest) OpsAuditSectionInstancesPostRequest(opsAuditSectionInstancesPostRequest OpsAuditSectionInstancesPostRequest) ApiOpsAuditSectionInstancesPostRequest {
	r.opsAuditSectionInstancesPostRequest = &opsAuditSectionInstancesPostRequest
	return r
}

func (r ApiOpsAuditSectionInstancesPostRequest) Execute() (*OpsAuditSectionInstancesGet200Response, *http.Response, error) {
	return r.ApiService.OpsAuditSectionInstancesPostExecute(r)
}

/*
OpsAuditSectionInstancesPost Method for OpsAuditSectionInstancesPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiOpsAuditSectionInstancesPostRequest
*/
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesPost(ctx context.Context) ApiOpsAuditSectionInstancesPostRequest {
	return ApiOpsAuditSectionInstancesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OpsAuditSectionInstancesGet200Response
func (a *OpsAuditSectionInstancesAPIService) OpsAuditSectionInstancesPostExecute(r ApiOpsAuditSectionInstancesPostRequest) (*OpsAuditSectionInstancesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpsAuditSectionInstancesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpsAuditSectionInstancesAPIService.OpsAuditSectionInstancesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ops_audit_section_instances"

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
	localVarPostBody = r.opsAuditSectionInstancesPostRequest
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
