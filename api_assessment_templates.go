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


// AssessmentTemplatesAPIService AssessmentTemplatesAPI service
type AssessmentTemplatesAPIService service

type ApiAssessmentTemplatesAssessmentTemplateIdDeleteRequest struct {
	ctx context.Context
	ApiService *AssessmentTemplatesAPIService
	assessmentTemplateId int64
}

func (r ApiAssessmentTemplatesAssessmentTemplateIdDeleteRequest) Execute() (*AssessmentTemplates, *http.Response, error) {
	return r.ApiService.AssessmentTemplatesAssessmentTemplateIdDeleteExecute(r)
}

/*
AssessmentTemplatesAssessmentTemplateIdDelete Method for AssessmentTemplatesAssessmentTemplateIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assessmentTemplateId Model id
 @return ApiAssessmentTemplatesAssessmentTemplateIdDeleteRequest
*/
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesAssessmentTemplateIdDelete(ctx context.Context, assessmentTemplateId int64) ApiAssessmentTemplatesAssessmentTemplateIdDeleteRequest {
	return ApiAssessmentTemplatesAssessmentTemplateIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		assessmentTemplateId: assessmentTemplateId,
	}
}

// Execute executes the request
//  @return AssessmentTemplates
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesAssessmentTemplateIdDeleteExecute(r ApiAssessmentTemplatesAssessmentTemplateIdDeleteRequest) (*AssessmentTemplates, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentTemplates
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentTemplatesAPIService.AssessmentTemplatesAssessmentTemplateIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_templates/{assessment_template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"assessment_template_id"+"}", url.PathEscape(parameterValueToString(r.assessmentTemplateId, "assessmentTemplateId")), -1)

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

type ApiAssessmentTemplatesAssessmentTemplateIdGetRequest struct {
	ctx context.Context
	ApiService *AssessmentTemplatesAPIService
	assessmentTemplateId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAssessmentTemplatesAssessmentTemplateIdGetRequest) Include(include []string) ApiAssessmentTemplatesAssessmentTemplateIdGetRequest {
	r.include = &include
	return r
}

func (r ApiAssessmentTemplatesAssessmentTemplateIdGetRequest) Execute() (*AssessmentTemplates, *http.Response, error) {
	return r.ApiService.AssessmentTemplatesAssessmentTemplateIdGetExecute(r)
}

/*
AssessmentTemplatesAssessmentTemplateIdGet Method for AssessmentTemplatesAssessmentTemplateIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assessmentTemplateId Model id
 @return ApiAssessmentTemplatesAssessmentTemplateIdGetRequest
*/
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesAssessmentTemplateIdGet(ctx context.Context, assessmentTemplateId int64) ApiAssessmentTemplatesAssessmentTemplateIdGetRequest {
	return ApiAssessmentTemplatesAssessmentTemplateIdGetRequest{
		ApiService: a,
		ctx: ctx,
		assessmentTemplateId: assessmentTemplateId,
	}
}

// Execute executes the request
//  @return AssessmentTemplates
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesAssessmentTemplateIdGetExecute(r ApiAssessmentTemplatesAssessmentTemplateIdGetRequest) (*AssessmentTemplates, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentTemplates
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentTemplatesAPIService.AssessmentTemplatesAssessmentTemplateIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_templates/{assessment_template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"assessment_template_id"+"}", url.PathEscape(parameterValueToString(r.assessmentTemplateId, "assessmentTemplateId")), -1)

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

type ApiAssessmentTemplatesAssessmentTemplateIdPutRequest struct {
	ctx context.Context
	ApiService *AssessmentTemplatesAPIService
	assessmentTemplateId int64
	assessmentTemplatesPut *AssessmentTemplatesPut
}

func (r ApiAssessmentTemplatesAssessmentTemplateIdPutRequest) AssessmentTemplatesPut(assessmentTemplatesPut AssessmentTemplatesPut) ApiAssessmentTemplatesAssessmentTemplateIdPutRequest {
	r.assessmentTemplatesPut = &assessmentTemplatesPut
	return r
}

func (r ApiAssessmentTemplatesAssessmentTemplateIdPutRequest) Execute() (*AssessmentTemplates, *http.Response, error) {
	return r.ApiService.AssessmentTemplatesAssessmentTemplateIdPutExecute(r)
}

/*
AssessmentTemplatesAssessmentTemplateIdPut Method for AssessmentTemplatesAssessmentTemplateIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param assessmentTemplateId Model id
 @return ApiAssessmentTemplatesAssessmentTemplateIdPutRequest
*/
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesAssessmentTemplateIdPut(ctx context.Context, assessmentTemplateId int64) ApiAssessmentTemplatesAssessmentTemplateIdPutRequest {
	return ApiAssessmentTemplatesAssessmentTemplateIdPutRequest{
		ApiService: a,
		ctx: ctx,
		assessmentTemplateId: assessmentTemplateId,
	}
}

// Execute executes the request
//  @return AssessmentTemplates
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesAssessmentTemplateIdPutExecute(r ApiAssessmentTemplatesAssessmentTemplateIdPutRequest) (*AssessmentTemplates, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentTemplates
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentTemplatesAPIService.AssessmentTemplatesAssessmentTemplateIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_templates/{assessment_template_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"assessment_template_id"+"}", url.PathEscape(parameterValueToString(r.assessmentTemplateId, "assessmentTemplateId")), -1)

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
	localVarPostBody = r.assessmentTemplatesPut
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

type ApiAssessmentTemplatesGetRequest struct {
	ctx context.Context
	ApiService *AssessmentTemplatesAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAssessmentTemplatesGetRequest) Include(include []string) ApiAssessmentTemplatesGetRequest {
	r.include = &include
	return r
}

func (r ApiAssessmentTemplatesGetRequest) Execute() (*AssessmentTemplatesGet200Response, *http.Response, error) {
	return r.ApiService.AssessmentTemplatesGetExecute(r)
}

/*
AssessmentTemplatesGet Method for AssessmentTemplatesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAssessmentTemplatesGetRequest
*/
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesGet(ctx context.Context) ApiAssessmentTemplatesGetRequest {
	return ApiAssessmentTemplatesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AssessmentTemplatesGet200Response
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesGetExecute(r ApiAssessmentTemplatesGetRequest) (*AssessmentTemplatesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentTemplatesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentTemplatesAPIService.AssessmentTemplatesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_templates"

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

type ApiAssessmentTemplatesPostRequest struct {
	ctx context.Context
	ApiService *AssessmentTemplatesAPIService
	assessmentTemplatesPostRequest *AssessmentTemplatesPostRequest
}

func (r ApiAssessmentTemplatesPostRequest) AssessmentTemplatesPostRequest(assessmentTemplatesPostRequest AssessmentTemplatesPostRequest) ApiAssessmentTemplatesPostRequest {
	r.assessmentTemplatesPostRequest = &assessmentTemplatesPostRequest
	return r
}

func (r ApiAssessmentTemplatesPostRequest) Execute() (*AssessmentTemplatesGet200Response, *http.Response, error) {
	return r.ApiService.AssessmentTemplatesPostExecute(r)
}

/*
AssessmentTemplatesPost Method for AssessmentTemplatesPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAssessmentTemplatesPostRequest
*/
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesPost(ctx context.Context) ApiAssessmentTemplatesPostRequest {
	return ApiAssessmentTemplatesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AssessmentTemplatesGet200Response
func (a *AssessmentTemplatesAPIService) AssessmentTemplatesPostExecute(r ApiAssessmentTemplatesPostRequest) (*AssessmentTemplatesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AssessmentTemplatesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssessmentTemplatesAPIService.AssessmentTemplatesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/assessment_templates"

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
	localVarPostBody = r.assessmentTemplatesPostRequest
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