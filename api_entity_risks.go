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


// EntityRisksAPIService EntityRisksAPI service
type EntityRisksAPIService service

type ApiEntityRisksEntityRiskIdGetRequest struct {
	ctx context.Context
	ApiService *EntityRisksAPIService
	entityRiskId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiEntityRisksEntityRiskIdGetRequest) Include(include []string) ApiEntityRisksEntityRiskIdGetRequest {
	r.include = &include
	return r
}

func (r ApiEntityRisksEntityRiskIdGetRequest) Execute() (*EntityRisks, *http.Response, error) {
	return r.ApiService.EntityRisksEntityRiskIdGetExecute(r)
}

/*
EntityRisksEntityRiskIdGet Method for EntityRisksEntityRiskIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entityRiskId Model id
 @return ApiEntityRisksEntityRiskIdGetRequest
*/
func (a *EntityRisksAPIService) EntityRisksEntityRiskIdGet(ctx context.Context, entityRiskId int64) ApiEntityRisksEntityRiskIdGetRequest {
	return ApiEntityRisksEntityRiskIdGetRequest{
		ApiService: a,
		ctx: ctx,
		entityRiskId: entityRiskId,
	}
}

// Execute executes the request
//  @return EntityRisks
func (a *EntityRisksAPIService) EntityRisksEntityRiskIdGetExecute(r ApiEntityRisksEntityRiskIdGetRequest) (*EntityRisks, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityRisks
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntityRisksAPIService.EntityRisksEntityRiskIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity_risks/{entity_risk_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity_risk_id"+"}", url.PathEscape(parameterValueToString(r.entityRiskId, "entityRiskId")), -1)

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

type ApiEntityRisksEntityRiskIdPutRequest struct {
	ctx context.Context
	ApiService *EntityRisksAPIService
	entityRiskId int64
	entityRisksPut *EntityRisksPut
}

func (r ApiEntityRisksEntityRiskIdPutRequest) EntityRisksPut(entityRisksPut EntityRisksPut) ApiEntityRisksEntityRiskIdPutRequest {
	r.entityRisksPut = &entityRisksPut
	return r
}

func (r ApiEntityRisksEntityRiskIdPutRequest) Execute() (*EntityRisks, *http.Response, error) {
	return r.ApiService.EntityRisksEntityRiskIdPutExecute(r)
}

/*
EntityRisksEntityRiskIdPut Method for EntityRisksEntityRiskIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param entityRiskId Model id
 @return ApiEntityRisksEntityRiskIdPutRequest
*/
func (a *EntityRisksAPIService) EntityRisksEntityRiskIdPut(ctx context.Context, entityRiskId int64) ApiEntityRisksEntityRiskIdPutRequest {
	return ApiEntityRisksEntityRiskIdPutRequest{
		ApiService: a,
		ctx: ctx,
		entityRiskId: entityRiskId,
	}
}

// Execute executes the request
//  @return EntityRisks
func (a *EntityRisksAPIService) EntityRisksEntityRiskIdPutExecute(r ApiEntityRisksEntityRiskIdPutRequest) (*EntityRisks, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityRisks
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntityRisksAPIService.EntityRisksEntityRiskIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity_risks/{entity_risk_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"entity_risk_id"+"}", url.PathEscape(parameterValueToString(r.entityRiskId, "entityRiskId")), -1)

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
	localVarPostBody = r.entityRisksPut
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

type ApiEntityRisksGetRequest struct {
	ctx context.Context
	ApiService *EntityRisksAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiEntityRisksGetRequest) Include(include []string) ApiEntityRisksGetRequest {
	r.include = &include
	return r
}

func (r ApiEntityRisksGetRequest) Execute() (*EntityRisksGet200Response, *http.Response, error) {
	return r.ApiService.EntityRisksGetExecute(r)
}

/*
EntityRisksGet Method for EntityRisksGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiEntityRisksGetRequest
*/
func (a *EntityRisksAPIService) EntityRisksGet(ctx context.Context) ApiEntityRisksGetRequest {
	return ApiEntityRisksGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EntityRisksGet200Response
func (a *EntityRisksAPIService) EntityRisksGetExecute(r ApiEntityRisksGetRequest) (*EntityRisksGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EntityRisksGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntityRisksAPIService.EntityRisksGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/entity_risks"

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
