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


// TimesheetEntriesAPIService TimesheetEntriesAPI service
type TimesheetEntriesAPIService service

type ApiTimesheetEntriesGetRequest struct {
	ctx context.Context
	ApiService *TimesheetEntriesAPIService
	start *string
	end *string
	skewForWeek *bool
	timesheetProjectId *int32
	timesheetableType *string
	timesheetableId *int32
	userId *int32
	limit *int32
}

// The start date for timesheets to find
func (r ApiTimesheetEntriesGetRequest) Start(start string) ApiTimesheetEntriesGetRequest {
	r.start = &start
	return r
}

// The end date for timesheets to find
func (r ApiTimesheetEntriesGetRequest) End(end string) ApiTimesheetEntriesGetRequest {
	r.end = &end
	return r
}

// \&quot;Skews\&quot; start and end date to the start and end of the week respectively
func (r ApiTimesheetEntriesGetRequest) SkewForWeek(skewForWeek bool) ApiTimesheetEntriesGetRequest {
	r.skewForWeek = &skewForWeek
	return r
}

func (r ApiTimesheetEntriesGetRequest) TimesheetProjectId(timesheetProjectId int32) ApiTimesheetEntriesGetRequest {
	r.timesheetProjectId = &timesheetProjectId
	return r
}

func (r ApiTimesheetEntriesGetRequest) TimesheetableType(timesheetableType string) ApiTimesheetEntriesGetRequest {
	r.timesheetableType = &timesheetableType
	return r
}

func (r ApiTimesheetEntriesGetRequest) TimesheetableId(timesheetableId int32) ApiTimesheetEntriesGetRequest {
	r.timesheetableId = &timesheetableId
	return r
}

func (r ApiTimesheetEntriesGetRequest) UserId(userId int32) ApiTimesheetEntriesGetRequest {
	r.userId = &userId
	return r
}

// Limits the returned result set
func (r ApiTimesheetEntriesGetRequest) Limit(limit int32) ApiTimesheetEntriesGetRequest {
	r.limit = &limit
	return r
}

func (r ApiTimesheetEntriesGetRequest) Execute() (*TimesheetEntriesGet200Response, *http.Response, error) {
	return r.ApiService.TimesheetEntriesGetExecute(r)
}

/*
TimesheetEntriesGet Method for TimesheetEntriesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTimesheetEntriesGetRequest
*/
func (a *TimesheetEntriesAPIService) TimesheetEntriesGet(ctx context.Context) ApiTimesheetEntriesGetRequest {
	return ApiTimesheetEntriesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TimesheetEntriesGet200Response
func (a *TimesheetEntriesAPIService) TimesheetEntriesGetExecute(r ApiTimesheetEntriesGetRequest) (*TimesheetEntriesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TimesheetEntriesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimesheetEntriesAPIService.TimesheetEntriesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/timesheet_entries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "form", "")
	}
	if r.skewForWeek != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "skewForWeek", r.skewForWeek, "form", "")
	}
	if r.timesheetProjectId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timesheet_project_id", r.timesheetProjectId, "form", "")
	}
	if r.timesheetableType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timesheetable_type", r.timesheetableType, "form", "")
	}
	if r.timesheetableId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "timesheetable_id", r.timesheetableId, "form", "")
	}
	if r.userId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "user_id", r.userId, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
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

type ApiTimesheetEntriesPostRequest struct {
	ctx context.Context
	ApiService *TimesheetEntriesAPIService
	timesheetEntriesPostRequest *TimesheetEntriesPostRequest
}

func (r ApiTimesheetEntriesPostRequest) TimesheetEntriesPostRequest(timesheetEntriesPostRequest TimesheetEntriesPostRequest) ApiTimesheetEntriesPostRequest {
	r.timesheetEntriesPostRequest = &timesheetEntriesPostRequest
	return r
}

func (r ApiTimesheetEntriesPostRequest) Execute() (*TimesheetEntriesGet200Response, *http.Response, error) {
	return r.ApiService.TimesheetEntriesPostExecute(r)
}

/*
TimesheetEntriesPost Method for TimesheetEntriesPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTimesheetEntriesPostRequest
*/
func (a *TimesheetEntriesAPIService) TimesheetEntriesPost(ctx context.Context) ApiTimesheetEntriesPostRequest {
	return ApiTimesheetEntriesPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TimesheetEntriesGet200Response
func (a *TimesheetEntriesAPIService) TimesheetEntriesPostExecute(r ApiTimesheetEntriesPostRequest) (*TimesheetEntriesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TimesheetEntriesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimesheetEntriesAPIService.TimesheetEntriesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/timesheet_entries"

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
	localVarPostBody = r.timesheetEntriesPostRequest
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

type ApiTimesheetEntriesTimesheetEntryIdDeleteRequest struct {
	ctx context.Context
	ApiService *TimesheetEntriesAPIService
	timesheetEntryId int64
}

func (r ApiTimesheetEntriesTimesheetEntryIdDeleteRequest) Execute() (*TimesheetEntries, *http.Response, error) {
	return r.ApiService.TimesheetEntriesTimesheetEntryIdDeleteExecute(r)
}

/*
TimesheetEntriesTimesheetEntryIdDelete Method for TimesheetEntriesTimesheetEntryIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param timesheetEntryId Model id
 @return ApiTimesheetEntriesTimesheetEntryIdDeleteRequest
*/
func (a *TimesheetEntriesAPIService) TimesheetEntriesTimesheetEntryIdDelete(ctx context.Context, timesheetEntryId int64) ApiTimesheetEntriesTimesheetEntryIdDeleteRequest {
	return ApiTimesheetEntriesTimesheetEntryIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		timesheetEntryId: timesheetEntryId,
	}
}

// Execute executes the request
//  @return TimesheetEntries
func (a *TimesheetEntriesAPIService) TimesheetEntriesTimesheetEntryIdDeleteExecute(r ApiTimesheetEntriesTimesheetEntryIdDeleteRequest) (*TimesheetEntries, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TimesheetEntries
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimesheetEntriesAPIService.TimesheetEntriesTimesheetEntryIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/timesheet_entries/{timesheet_entry_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"timesheet_entry_id"+"}", url.PathEscape(parameterValueToString(r.timesheetEntryId, "timesheetEntryId")), -1)

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

type ApiTimesheetEntriesTimesheetEntryIdGetRequest struct {
	ctx context.Context
	ApiService *TimesheetEntriesAPIService
	timesheetEntryId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiTimesheetEntriesTimesheetEntryIdGetRequest) Include(include []string) ApiTimesheetEntriesTimesheetEntryIdGetRequest {
	r.include = &include
	return r
}

func (r ApiTimesheetEntriesTimesheetEntryIdGetRequest) Execute() (*TimesheetEntries, *http.Response, error) {
	return r.ApiService.TimesheetEntriesTimesheetEntryIdGetExecute(r)
}

/*
TimesheetEntriesTimesheetEntryIdGet Method for TimesheetEntriesTimesheetEntryIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param timesheetEntryId Model id
 @return ApiTimesheetEntriesTimesheetEntryIdGetRequest
*/
func (a *TimesheetEntriesAPIService) TimesheetEntriesTimesheetEntryIdGet(ctx context.Context, timesheetEntryId int64) ApiTimesheetEntriesTimesheetEntryIdGetRequest {
	return ApiTimesheetEntriesTimesheetEntryIdGetRequest{
		ApiService: a,
		ctx: ctx,
		timesheetEntryId: timesheetEntryId,
	}
}

// Execute executes the request
//  @return TimesheetEntries
func (a *TimesheetEntriesAPIService) TimesheetEntriesTimesheetEntryIdGetExecute(r ApiTimesheetEntriesTimesheetEntryIdGetRequest) (*TimesheetEntries, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TimesheetEntries
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimesheetEntriesAPIService.TimesheetEntriesTimesheetEntryIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/timesheet_entries/{timesheet_entry_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"timesheet_entry_id"+"}", url.PathEscape(parameterValueToString(r.timesheetEntryId, "timesheetEntryId")), -1)

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

type ApiTimesheetEntriesTimesheetEntryIdPutRequest struct {
	ctx context.Context
	ApiService *TimesheetEntriesAPIService
	timesheetEntryId int64
	timesheetEntriesPut *TimesheetEntriesPut
}

func (r ApiTimesheetEntriesTimesheetEntryIdPutRequest) TimesheetEntriesPut(timesheetEntriesPut TimesheetEntriesPut) ApiTimesheetEntriesTimesheetEntryIdPutRequest {
	r.timesheetEntriesPut = &timesheetEntriesPut
	return r
}

func (r ApiTimesheetEntriesTimesheetEntryIdPutRequest) Execute() (*TimesheetEntries, *http.Response, error) {
	return r.ApiService.TimesheetEntriesTimesheetEntryIdPutExecute(r)
}

/*
TimesheetEntriesTimesheetEntryIdPut Method for TimesheetEntriesTimesheetEntryIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param timesheetEntryId Model id
 @return ApiTimesheetEntriesTimesheetEntryIdPutRequest
*/
func (a *TimesheetEntriesAPIService) TimesheetEntriesTimesheetEntryIdPut(ctx context.Context, timesheetEntryId int64) ApiTimesheetEntriesTimesheetEntryIdPutRequest {
	return ApiTimesheetEntriesTimesheetEntryIdPutRequest{
		ApiService: a,
		ctx: ctx,
		timesheetEntryId: timesheetEntryId,
	}
}

// Execute executes the request
//  @return TimesheetEntries
func (a *TimesheetEntriesAPIService) TimesheetEntriesTimesheetEntryIdPutExecute(r ApiTimesheetEntriesTimesheetEntryIdPutRequest) (*TimesheetEntries, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TimesheetEntries
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TimesheetEntriesAPIService.TimesheetEntriesTimesheetEntryIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/timesheet_entries/{timesheet_entry_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"timesheet_entry_id"+"}", url.PathEscape(parameterValueToString(r.timesheetEntryId, "timesheetEntryId")), -1)

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
	localVarPostBody = r.timesheetEntriesPut
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
