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


// AuditSurveysAPIService AuditSurveysAPI service
type AuditSurveysAPIService service

type ApiAuditSurveysAuditSurveyIdDeleteRequest struct {
	ctx context.Context
	ApiService *AuditSurveysAPIService
	auditSurveyId int64
}

func (r ApiAuditSurveysAuditSurveyIdDeleteRequest) Execute() (*AuditSurveys, *http.Response, error) {
	return r.ApiService.AuditSurveysAuditSurveyIdDeleteExecute(r)
}

/*
AuditSurveysAuditSurveyIdDelete Method for AuditSurveysAuditSurveyIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param auditSurveyId Model id
 @return ApiAuditSurveysAuditSurveyIdDeleteRequest
*/
func (a *AuditSurveysAPIService) AuditSurveysAuditSurveyIdDelete(ctx context.Context, auditSurveyId int64) ApiAuditSurveysAuditSurveyIdDeleteRequest {
	return ApiAuditSurveysAuditSurveyIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		auditSurveyId: auditSurveyId,
	}
}

// Execute executes the request
//  @return AuditSurveys
func (a *AuditSurveysAPIService) AuditSurveysAuditSurveyIdDeleteExecute(r ApiAuditSurveysAuditSurveyIdDeleteRequest) (*AuditSurveys, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditSurveys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditSurveysAPIService.AuditSurveysAuditSurveyIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/audit_surveys/{audit_survey_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"audit_survey_id"+"}", url.PathEscape(parameterValueToString(r.auditSurveyId, "auditSurveyId")), -1)

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

type ApiAuditSurveysAuditSurveyIdGetRequest struct {
	ctx context.Context
	ApiService *AuditSurveysAPIService
	auditSurveyId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAuditSurveysAuditSurveyIdGetRequest) Include(include []string) ApiAuditSurveysAuditSurveyIdGetRequest {
	r.include = &include
	return r
}

func (r ApiAuditSurveysAuditSurveyIdGetRequest) Execute() (*AuditSurveys, *http.Response, error) {
	return r.ApiService.AuditSurveysAuditSurveyIdGetExecute(r)
}

/*
AuditSurveysAuditSurveyIdGet Method for AuditSurveysAuditSurveyIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param auditSurveyId Model id
 @return ApiAuditSurveysAuditSurveyIdGetRequest
*/
func (a *AuditSurveysAPIService) AuditSurveysAuditSurveyIdGet(ctx context.Context, auditSurveyId int64) ApiAuditSurveysAuditSurveyIdGetRequest {
	return ApiAuditSurveysAuditSurveyIdGetRequest{
		ApiService: a,
		ctx: ctx,
		auditSurveyId: auditSurveyId,
	}
}

// Execute executes the request
//  @return AuditSurveys
func (a *AuditSurveysAPIService) AuditSurveysAuditSurveyIdGetExecute(r ApiAuditSurveysAuditSurveyIdGetRequest) (*AuditSurveys, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditSurveys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditSurveysAPIService.AuditSurveysAuditSurveyIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/audit_surveys/{audit_survey_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"audit_survey_id"+"}", url.PathEscape(parameterValueToString(r.auditSurveyId, "auditSurveyId")), -1)

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

type ApiAuditSurveysAuditSurveyIdPutRequest struct {
	ctx context.Context
	ApiService *AuditSurveysAPIService
	auditSurveyId int64
	auditSurveysPut *AuditSurveysPut
}

func (r ApiAuditSurveysAuditSurveyIdPutRequest) AuditSurveysPut(auditSurveysPut AuditSurveysPut) ApiAuditSurveysAuditSurveyIdPutRequest {
	r.auditSurveysPut = &auditSurveysPut
	return r
}

func (r ApiAuditSurveysAuditSurveyIdPutRequest) Execute() (*AuditSurveys, *http.Response, error) {
	return r.ApiService.AuditSurveysAuditSurveyIdPutExecute(r)
}

/*
AuditSurveysAuditSurveyIdPut Method for AuditSurveysAuditSurveyIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param auditSurveyId Model id
 @return ApiAuditSurveysAuditSurveyIdPutRequest
*/
func (a *AuditSurveysAPIService) AuditSurveysAuditSurveyIdPut(ctx context.Context, auditSurveyId int64) ApiAuditSurveysAuditSurveyIdPutRequest {
	return ApiAuditSurveysAuditSurveyIdPutRequest{
		ApiService: a,
		ctx: ctx,
		auditSurveyId: auditSurveyId,
	}
}

// Execute executes the request
//  @return AuditSurveys
func (a *AuditSurveysAPIService) AuditSurveysAuditSurveyIdPutExecute(r ApiAuditSurveysAuditSurveyIdPutRequest) (*AuditSurveys, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditSurveys
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditSurveysAPIService.AuditSurveysAuditSurveyIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/audit_surveys/{audit_survey_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"audit_survey_id"+"}", url.PathEscape(parameterValueToString(r.auditSurveyId, "auditSurveyId")), -1)

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
	localVarPostBody = r.auditSurveysPut
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

type ApiAuditSurveysGetRequest struct {
	ctx context.Context
	ApiService *AuditSurveysAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAuditSurveysGetRequest) Include(include []string) ApiAuditSurveysGetRequest {
	r.include = &include
	return r
}

func (r ApiAuditSurveysGetRequest) Execute() (*AuditSurveysGet200Response, *http.Response, error) {
	return r.ApiService.AuditSurveysGetExecute(r)
}

/*
AuditSurveysGet Method for AuditSurveysGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAuditSurveysGetRequest
*/
func (a *AuditSurveysAPIService) AuditSurveysGet(ctx context.Context) ApiAuditSurveysGetRequest {
	return ApiAuditSurveysGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AuditSurveysGet200Response
func (a *AuditSurveysAPIService) AuditSurveysGetExecute(r ApiAuditSurveysGetRequest) (*AuditSurveysGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuditSurveysGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuditSurveysAPIService.AuditSurveysGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/audit_surveys"

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
