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


// AuditableEntitiesAPIService AuditableEntitiesAPI service
type AuditableEntitiesAPIService service

type ApiAuditableEntitiesAuditableEntityIdDeleteRequest struct {
	ctx context.Context
	ApiService *AuditableEntitiesAPIService
	auditableEntityId int64
}

func (r ApiAuditableEntitiesAuditableEntityIdDeleteRequest) Execute() (*AuditableEntities, *http.Response, error) {
	return r.ApiService.AuditableEntitiesAuditableEntityIdDeleteExecute(r)
}

/*
AuditableEntitiesAuditableEntityIdDelete Method for AuditableEntitiesAuditableEntityIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param auditableEntityId Model id
 @return ApiAuditableEntitiesAuditableEntityIdDeleteRequest
*/
func (a *AuditableEntitiesAPIService) AuditableEntitiesAuditableEntityIdDelete(ctx context.Context, auditableEntityId int64) ApiAuditableEntitiesAuditableEntityIdDeleteRequest {
	return ApiAuditableEntitiesAuditableEntityIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		auditableEntityId: auditableEntityId,
	}
}

// Execute executes the request
//  @return AuditableEntities
func (a *AuditableEntitiesAPIService) AuditableEntitiesAuditableEntityIdDeleteExecute(r ApiAuditableEntitiesAuditableEntityIdDeleteRequest) (*AuditableEntities, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditableEntities
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditableEntitiesAPIService.AuditableEntitiesAuditableEntityIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditable_entities/{auditable_entity_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"auditable_entity_id"+"}", url.PathEscape(parameterValueToString(r.auditableEntityId, "auditableEntityId")), -1)

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

type ApiAuditableEntitiesAuditableEntityIdGetRequest struct {
	ctx context.Context
	ApiService *AuditableEntitiesAPIService
	auditableEntityId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAuditableEntitiesAuditableEntityIdGetRequest) Include(include []string) ApiAuditableEntitiesAuditableEntityIdGetRequest {
	r.include = &include
	return r
}

func (r ApiAuditableEntitiesAuditableEntityIdGetRequest) Execute() (*AuditableEntities, *http.Response, error) {
	return r.ApiService.AuditableEntitiesAuditableEntityIdGetExecute(r)
}

/*
AuditableEntitiesAuditableEntityIdGet Method for AuditableEntitiesAuditableEntityIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param auditableEntityId Model id
 @return ApiAuditableEntitiesAuditableEntityIdGetRequest
*/
func (a *AuditableEntitiesAPIService) AuditableEntitiesAuditableEntityIdGet(ctx context.Context, auditableEntityId int64) ApiAuditableEntitiesAuditableEntityIdGetRequest {
	return ApiAuditableEntitiesAuditableEntityIdGetRequest{
		ApiService: a,
		ctx: ctx,
		auditableEntityId: auditableEntityId,
	}
}

// Execute executes the request
//  @return AuditableEntities
func (a *AuditableEntitiesAPIService) AuditableEntitiesAuditableEntityIdGetExecute(r ApiAuditableEntitiesAuditableEntityIdGetRequest) (*AuditableEntities, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditableEntities
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditableEntitiesAPIService.AuditableEntitiesAuditableEntityIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditable_entities/{auditable_entity_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"auditable_entity_id"+"}", url.PathEscape(parameterValueToString(r.auditableEntityId, "auditableEntityId")), -1)

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

type ApiAuditableEntitiesAuditableEntityIdPutRequest struct {
	ctx context.Context
	ApiService *AuditableEntitiesAPIService
	auditableEntityId int64
	auditableEntitiesPut *AuditableEntitiesPut
}

func (r ApiAuditableEntitiesAuditableEntityIdPutRequest) AuditableEntitiesPut(auditableEntitiesPut AuditableEntitiesPut) ApiAuditableEntitiesAuditableEntityIdPutRequest {
	r.auditableEntitiesPut = &auditableEntitiesPut
	return r
}

func (r ApiAuditableEntitiesAuditableEntityIdPutRequest) Execute() (*AuditableEntities, *http.Response, error) {
	return r.ApiService.AuditableEntitiesAuditableEntityIdPutExecute(r)
}

/*
AuditableEntitiesAuditableEntityIdPut Method for AuditableEntitiesAuditableEntityIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param auditableEntityId Model id
 @return ApiAuditableEntitiesAuditableEntityIdPutRequest
*/
func (a *AuditableEntitiesAPIService) AuditableEntitiesAuditableEntityIdPut(ctx context.Context, auditableEntityId int64) ApiAuditableEntitiesAuditableEntityIdPutRequest {
	return ApiAuditableEntitiesAuditableEntityIdPutRequest{
		ApiService: a,
		ctx: ctx,
		auditableEntityId: auditableEntityId,
	}
}

// Execute executes the request
//  @return AuditableEntities
func (a *AuditableEntitiesAPIService) AuditableEntitiesAuditableEntityIdPutExecute(r ApiAuditableEntitiesAuditableEntityIdPutRequest) (*AuditableEntities, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditableEntities
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditableEntitiesAPIService.AuditableEntitiesAuditableEntityIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditable_entities/{auditable_entity_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"auditable_entity_id"+"}", url.PathEscape(parameterValueToString(r.auditableEntityId, "auditableEntityId")), -1)

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
	localVarPostBody = r.auditableEntitiesPut
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

type ApiAuditableEntitiesGetRequest struct {
	ctx context.Context
	ApiService *AuditableEntitiesAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAuditableEntitiesGetRequest) Include(include []string) ApiAuditableEntitiesGetRequest {
	r.include = &include
	return r
}

func (r ApiAuditableEntitiesGetRequest) Execute() (*AuditableEntitiesGet200Response, *http.Response, error) {
	return r.ApiService.AuditableEntitiesGetExecute(r)
}

/*
AuditableEntitiesGet Method for AuditableEntitiesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuditableEntitiesGetRequest
*/
func (a *AuditableEntitiesAPIService) AuditableEntitiesGet(ctx context.Context) ApiAuditableEntitiesGetRequest {
	return ApiAuditableEntitiesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuditableEntitiesGet200Response
func (a *AuditableEntitiesAPIService) AuditableEntitiesGetExecute(r ApiAuditableEntitiesGetRequest) (*AuditableEntitiesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditableEntitiesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditableEntitiesAPIService.AuditableEntitiesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditable_entities"

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

type ApiAuditableEntitiesPostRequest struct {
	ctx context.Context
	ApiService *AuditableEntitiesAPIService
	auditableEntitiesPostRequest *AuditableEntitiesPostRequest
}

func (r ApiAuditableEntitiesPostRequest) AuditableEntitiesPostRequest(auditableEntitiesPostRequest AuditableEntitiesPostRequest) ApiAuditableEntitiesPostRequest {
	r.auditableEntitiesPostRequest = &auditableEntitiesPostRequest
	return r
}

func (r ApiAuditableEntitiesPostRequest) Execute() (*AuditableEntitiesGet200Response, *http.Response, error) {
	return r.ApiService.AuditableEntitiesPostExecute(r)
}

/*
AuditableEntitiesPost Method for AuditableEntitiesPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuditableEntitiesPostRequest
*/
func (a *AuditableEntitiesAPIService) AuditableEntitiesPost(ctx context.Context) ApiAuditableEntitiesPostRequest {
	return ApiAuditableEntitiesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuditableEntitiesGet200Response
func (a *AuditableEntitiesAPIService) AuditableEntitiesPostExecute(r ApiAuditableEntitiesPostRequest) (*AuditableEntitiesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditableEntitiesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditableEntitiesAPIService.AuditableEntitiesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/auditable_entities"

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
	localVarPostBody = r.auditableEntitiesPostRequest
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