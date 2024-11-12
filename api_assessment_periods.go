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


// AssessmentPeriodsAPIService AssessmentPeriodsAPI service
type AssessmentPeriodsAPIService service

type ApiAssessmentPeriodsAssessmentPeriodIdDeleteRequest struct {
	ctx context.Context
	ApiService *AssessmentPeriodsAPIService
	assessmentPeriodId int64
}

func (r ApiAssessmentPeriodsAssessmentPeriodIdDeleteRequest) Execute() (*AssessmentPeriods, *http.Response, error) {
	return r.ApiService.AssessmentPeriodsAssessmentPeriodIdDeleteExecute(r)
}

/*
AssessmentPeriodsAssessmentPeriodIdDelete Method for AssessmentPeriodsAssessmentPeriodIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assessmentPeriodId Model id
 @return ApiAssessmentPeriodsAssessmentPeriodIdDeleteRequest
*/
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsAssessmentPeriodIdDelete(ctx context.Context, assessmentPeriodId int64) ApiAssessmentPeriodsAssessmentPeriodIdDeleteRequest {
	return ApiAssessmentPeriodsAssessmentPeriodIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		assessmentPeriodId: assessmentPeriodId,
	}
}

// Execute executes the request
//  @return AssessmentPeriods
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsAssessmentPeriodIdDeleteExecute(r ApiAssessmentPeriodsAssessmentPeriodIdDeleteRequest) (*AssessmentPeriods, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentPeriods
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentPeriodsAPIService.AssessmentPeriodsAssessmentPeriodIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_periods/{assessment_period_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"assessment_period_id"+"}", url.PathEscape(parameterValueToString(r.assessmentPeriodId, "assessmentPeriodId")), -1)

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

type ApiAssessmentPeriodsAssessmentPeriodIdGetRequest struct {
	ctx context.Context
	ApiService *AssessmentPeriodsAPIService
	assessmentPeriodId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAssessmentPeriodsAssessmentPeriodIdGetRequest) Include(include []string) ApiAssessmentPeriodsAssessmentPeriodIdGetRequest {
	r.include = &include
	return r
}

func (r ApiAssessmentPeriodsAssessmentPeriodIdGetRequest) Execute() (*AssessmentPeriods, *http.Response, error) {
	return r.ApiService.AssessmentPeriodsAssessmentPeriodIdGetExecute(r)
}

/*
AssessmentPeriodsAssessmentPeriodIdGet Method for AssessmentPeriodsAssessmentPeriodIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assessmentPeriodId Model id
 @return ApiAssessmentPeriodsAssessmentPeriodIdGetRequest
*/
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsAssessmentPeriodIdGet(ctx context.Context, assessmentPeriodId int64) ApiAssessmentPeriodsAssessmentPeriodIdGetRequest {
	return ApiAssessmentPeriodsAssessmentPeriodIdGetRequest{
		ApiService: a,
		ctx: ctx,
		assessmentPeriodId: assessmentPeriodId,
	}
}

// Execute executes the request
//  @return AssessmentPeriods
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsAssessmentPeriodIdGetExecute(r ApiAssessmentPeriodsAssessmentPeriodIdGetRequest) (*AssessmentPeriods, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentPeriods
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentPeriodsAPIService.AssessmentPeriodsAssessmentPeriodIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_periods/{assessment_period_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"assessment_period_id"+"}", url.PathEscape(parameterValueToString(r.assessmentPeriodId, "assessmentPeriodId")), -1)

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

type ApiAssessmentPeriodsAssessmentPeriodIdPutRequest struct {
	ctx context.Context
	ApiService *AssessmentPeriodsAPIService
	assessmentPeriodId int64
	assessmentPeriodsPut *AssessmentPeriodsPut
}

func (r ApiAssessmentPeriodsAssessmentPeriodIdPutRequest) AssessmentPeriodsPut(assessmentPeriodsPut AssessmentPeriodsPut) ApiAssessmentPeriodsAssessmentPeriodIdPutRequest {
	r.assessmentPeriodsPut = &assessmentPeriodsPut
	return r
}

func (r ApiAssessmentPeriodsAssessmentPeriodIdPutRequest) Execute() (*AssessmentPeriods, *http.Response, error) {
	return r.ApiService.AssessmentPeriodsAssessmentPeriodIdPutExecute(r)
}

/*
AssessmentPeriodsAssessmentPeriodIdPut Method for AssessmentPeriodsAssessmentPeriodIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assessmentPeriodId Model id
 @return ApiAssessmentPeriodsAssessmentPeriodIdPutRequest
*/
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsAssessmentPeriodIdPut(ctx context.Context, assessmentPeriodId int64) ApiAssessmentPeriodsAssessmentPeriodIdPutRequest {
	return ApiAssessmentPeriodsAssessmentPeriodIdPutRequest{
		ApiService: a,
		ctx: ctx,
		assessmentPeriodId: assessmentPeriodId,
	}
}

// Execute executes the request
//  @return AssessmentPeriods
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsAssessmentPeriodIdPutExecute(r ApiAssessmentPeriodsAssessmentPeriodIdPutRequest) (*AssessmentPeriods, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentPeriods
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentPeriodsAPIService.AssessmentPeriodsAssessmentPeriodIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_periods/{assessment_period_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"assessment_period_id"+"}", url.PathEscape(parameterValueToString(r.assessmentPeriodId, "assessmentPeriodId")), -1)

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
	localVarPostBody = r.assessmentPeriodsPut
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

type ApiAssessmentPeriodsGetRequest struct {
	ctx context.Context
	ApiService *AssessmentPeriodsAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAssessmentPeriodsGetRequest) Include(include []string) ApiAssessmentPeriodsGetRequest {
	r.include = &include
	return r
}

func (r ApiAssessmentPeriodsGetRequest) Execute() (*AssessmentPeriodsGet200Response, *http.Response, error) {
	return r.ApiService.AssessmentPeriodsGetExecute(r)
}

/*
AssessmentPeriodsGet Method for AssessmentPeriodsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAssessmentPeriodsGetRequest
*/
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsGet(ctx context.Context) ApiAssessmentPeriodsGetRequest {
	return ApiAssessmentPeriodsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AssessmentPeriodsGet200Response
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsGetExecute(r ApiAssessmentPeriodsGetRequest) (*AssessmentPeriodsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentPeriodsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentPeriodsAPIService.AssessmentPeriodsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_periods"

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

type ApiAssessmentPeriodsPostRequest struct {
	ctx context.Context
	ApiService *AssessmentPeriodsAPIService
	assessmentPeriodsPostRequest *AssessmentPeriodsPostRequest
}

func (r ApiAssessmentPeriodsPostRequest) AssessmentPeriodsPostRequest(assessmentPeriodsPostRequest AssessmentPeriodsPostRequest) ApiAssessmentPeriodsPostRequest {
	r.assessmentPeriodsPostRequest = &assessmentPeriodsPostRequest
	return r
}

func (r ApiAssessmentPeriodsPostRequest) Execute() (*AssessmentPeriodsGet200Response, *http.Response, error) {
	return r.ApiService.AssessmentPeriodsPostExecute(r)
}

/*
AssessmentPeriodsPost Method for AssessmentPeriodsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAssessmentPeriodsPostRequest
*/
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsPost(ctx context.Context) ApiAssessmentPeriodsPostRequest {
	return ApiAssessmentPeriodsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AssessmentPeriodsGet200Response
func (a *AssessmentPeriodsAPIService) AssessmentPeriodsPostExecute(r ApiAssessmentPeriodsPostRequest) (*AssessmentPeriodsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentPeriodsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentPeriodsAPIService.AssessmentPeriodsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_periods"

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
	localVarPostBody = r.assessmentPeriodsPostRequest
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
