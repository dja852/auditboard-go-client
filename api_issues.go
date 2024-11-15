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


// IssuesAPIService IssuesAPI service
type IssuesAPIService service

type ApiIssuesGetRequest struct {
	ctx context.Context
	ApiService *IssuesAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiIssuesGetRequest) Include(include []string) ApiIssuesGetRequest {
	r.include = &include
	return r
}

func (r ApiIssuesGetRequest) Execute() (*IssuesGet200Response, *http.Response, error) {
	return r.ApiService.IssuesGetExecute(r)
}

/*
IssuesGet Method for IssuesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssuesGetRequest
*/
func (a *IssuesAPIService) IssuesGet(ctx context.Context) ApiIssuesGetRequest {
	return ApiIssuesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IssuesGet200Response
func (a *IssuesAPIService) IssuesGetExecute(r ApiIssuesGetRequest) (*IssuesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IssuesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesAPIService.IssuesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues"

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

type ApiIssuesIssueIdClosePostRequest struct {
	ctx context.Context
	ApiService *IssuesAPIService
	issueId int64
}

func (r ApiIssuesIssueIdClosePostRequest) Execute() (*ReportsGenerateTimesheetsPost200Response, *http.Response, error) {
	return r.ApiService.IssuesIssueIdClosePostExecute(r)
}

/*
IssuesIssueIdClosePost Method for IssuesIssueIdClosePost

Mark an Issue as closed

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId Issue id
 @return ApiIssuesIssueIdClosePostRequest
*/
func (a *IssuesAPIService) IssuesIssueIdClosePost(ctx context.Context, issueId int64) ApiIssuesIssueIdClosePostRequest {
	return ApiIssuesIssueIdClosePostRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
	}
}

// Execute executes the request
//  @return ReportsGenerateTimesheetsPost200Response
func (a *IssuesAPIService) IssuesIssueIdClosePostExecute(r ApiIssuesIssueIdClosePostRequest) (*ReportsGenerateTimesheetsPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReportsGenerateTimesheetsPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesAPIService.IssuesIssueIdClosePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues/{issue_id}/close"
	localVarPath = strings.Replace(localVarPath, "{"+"issue_id"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)

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

type ApiIssuesIssueIdDeleteRequest struct {
	ctx context.Context
	ApiService *IssuesAPIService
	issueId int64
}

func (r ApiIssuesIssueIdDeleteRequest) Execute() (*Issues, *http.Response, error) {
	return r.ApiService.IssuesIssueIdDeleteExecute(r)
}

/*
IssuesIssueIdDelete Method for IssuesIssueIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId Model id
 @return ApiIssuesIssueIdDeleteRequest
*/
func (a *IssuesAPIService) IssuesIssueIdDelete(ctx context.Context, issueId int64) ApiIssuesIssueIdDeleteRequest {
	return ApiIssuesIssueIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
	}
}

// Execute executes the request
//  @return Issues
func (a *IssuesAPIService) IssuesIssueIdDeleteExecute(r ApiIssuesIssueIdDeleteRequest) (*Issues, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Issues
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesAPIService.IssuesIssueIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues/{issue_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"issue_id"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)

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

type ApiIssuesIssueIdGetRequest struct {
	ctx context.Context
	ApiService *IssuesAPIService
	issueId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiIssuesIssueIdGetRequest) Include(include []string) ApiIssuesIssueIdGetRequest {
	r.include = &include
	return r
}

func (r ApiIssuesIssueIdGetRequest) Execute() (*Issues, *http.Response, error) {
	return r.ApiService.IssuesIssueIdGetExecute(r)
}

/*
IssuesIssueIdGet Method for IssuesIssueIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId Model id
 @return ApiIssuesIssueIdGetRequest
*/
func (a *IssuesAPIService) IssuesIssueIdGet(ctx context.Context, issueId int64) ApiIssuesIssueIdGetRequest {
	return ApiIssuesIssueIdGetRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
	}
}

// Execute executes the request
//  @return Issues
func (a *IssuesAPIService) IssuesIssueIdGetExecute(r ApiIssuesIssueIdGetRequest) (*Issues, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Issues
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesAPIService.IssuesIssueIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues/{issue_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"issue_id"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)

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

type ApiIssuesIssueIdMovePostRequest struct {
	ctx context.Context
	ApiService *IssuesAPIService
	issueId int64
	issuesIssueIdMovePostRequest *IssuesIssueIdMovePostRequest
}

func (r ApiIssuesIssueIdMovePostRequest) IssuesIssueIdMovePostRequest(issuesIssueIdMovePostRequest IssuesIssueIdMovePostRequest) ApiIssuesIssueIdMovePostRequest {
	r.issuesIssueIdMovePostRequest = &issuesIssueIdMovePostRequest
	return r
}

func (r ApiIssuesIssueIdMovePostRequest) Execute() (*ReportsGenerateTimesheetsPost200Response, *http.Response, error) {
	return r.ApiService.IssuesIssueIdMovePostExecute(r)
}

/*
IssuesIssueIdMovePost Method for IssuesIssueIdMovePost

Moves an Issue from one category/sourceable type to another, e.g. from Control to OpsAudit.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId Issue id
 @return ApiIssuesIssueIdMovePostRequest
*/
func (a *IssuesAPIService) IssuesIssueIdMovePost(ctx context.Context, issueId int64) ApiIssuesIssueIdMovePostRequest {
	return ApiIssuesIssueIdMovePostRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
	}
}

// Execute executes the request
//  @return ReportsGenerateTimesheetsPost200Response
func (a *IssuesAPIService) IssuesIssueIdMovePostExecute(r ApiIssuesIssueIdMovePostRequest) (*ReportsGenerateTimesheetsPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReportsGenerateTimesheetsPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesAPIService.IssuesIssueIdMovePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues/{issue_id}/move"
	localVarPath = strings.Replace(localVarPath, "{"+"issue_id"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)

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
	localVarPostBody = r.issuesIssueIdMovePostRequest
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

type ApiIssuesIssueIdPutRequest struct {
	ctx context.Context
	ApiService *IssuesAPIService
	issueId int64
	issuesPut *IssuesPut
}

func (r ApiIssuesIssueIdPutRequest) IssuesPut(issuesPut IssuesPut) ApiIssuesIssueIdPutRequest {
	r.issuesPut = &issuesPut
	return r
}

func (r ApiIssuesIssueIdPutRequest) Execute() (*Issues, *http.Response, error) {
	return r.ApiService.IssuesIssueIdPutExecute(r)
}

/*
IssuesIssueIdPut Method for IssuesIssueIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId Model id
 @return ApiIssuesIssueIdPutRequest
*/
func (a *IssuesAPIService) IssuesIssueIdPut(ctx context.Context, issueId int64) ApiIssuesIssueIdPutRequest {
	return ApiIssuesIssueIdPutRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
	}
}

// Execute executes the request
//  @return Issues
func (a *IssuesAPIService) IssuesIssueIdPutExecute(r ApiIssuesIssueIdPutRequest) (*Issues, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Issues
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesAPIService.IssuesIssueIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues/{issue_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"issue_id"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)

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
	localVarPostBody = r.issuesPut
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

type ApiIssuesIssueIdRemediatePostRequest struct {
	ctx context.Context
	ApiService *IssuesAPIService
	issueId int64
}

func (r ApiIssuesIssueIdRemediatePostRequest) Execute() (*ReportsGenerateTimesheetsPost200Response, *http.Response, error) {
	return r.ApiService.IssuesIssueIdRemediatePostExecute(r)
}

/*
IssuesIssueIdRemediatePost Method for IssuesIssueIdRemediatePost

Remediate an Issue

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param issueId Issue id
 @return ApiIssuesIssueIdRemediatePostRequest
*/
func (a *IssuesAPIService) IssuesIssueIdRemediatePost(ctx context.Context, issueId int64) ApiIssuesIssueIdRemediatePostRequest {
	return ApiIssuesIssueIdRemediatePostRequest{
		ApiService: a,
		ctx: ctx,
		issueId: issueId,
	}
}

// Execute executes the request
//  @return ReportsGenerateTimesheetsPost200Response
func (a *IssuesAPIService) IssuesIssueIdRemediatePostExecute(r ApiIssuesIssueIdRemediatePostRequest) (*ReportsGenerateTimesheetsPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReportsGenerateTimesheetsPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesAPIService.IssuesIssueIdRemediatePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues/{issue_id}/remediate"
	localVarPath = strings.Replace(localVarPath, "{"+"issue_id"+"}", url.PathEscape(parameterValueToString(r.issueId, "issueId")), -1)

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

type ApiIssuesPostRequest struct {
	ctx context.Context
	ApiService *IssuesAPIService
	issuesPostRequest *IssuesPostRequest
}

func (r ApiIssuesPostRequest) IssuesPostRequest(issuesPostRequest IssuesPostRequest) ApiIssuesPostRequest {
	r.issuesPostRequest = &issuesPostRequest
	return r
}

func (r ApiIssuesPostRequest) Execute() (*IssuesGet200Response, *http.Response, error) {
	return r.ApiService.IssuesPostExecute(r)
}

/*
IssuesPost Method for IssuesPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiIssuesPostRequest
*/
func (a *IssuesAPIService) IssuesPost(ctx context.Context) ApiIssuesPostRequest {
	return ApiIssuesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return IssuesGet200Response
func (a *IssuesAPIService) IssuesPostExecute(r ApiIssuesPostRequest) (*IssuesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *IssuesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IssuesAPIService.IssuesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/issues"

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
	localVarPostBody = r.issuesPostRequest
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
