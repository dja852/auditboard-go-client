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


// AllowedUsersAPIService AllowedUsersAPI service
type AllowedUsersAPIService service

type ApiAllowedUsersAllowedUserIdDeleteRequest struct {
	ctx context.Context
	ApiService *AllowedUsersAPIService
	allowedUserId int64
}

func (r ApiAllowedUsersAllowedUserIdDeleteRequest) Execute() (*AllowedUsers, *http.Response, error) {
	return r.ApiService.AllowedUsersAllowedUserIdDeleteExecute(r)
}

/*
AllowedUsersAllowedUserIdDelete Method for AllowedUsersAllowedUserIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param allowedUserId Model id
 @return ApiAllowedUsersAllowedUserIdDeleteRequest
*/
func (a *AllowedUsersAPIService) AllowedUsersAllowedUserIdDelete(ctx context.Context, allowedUserId int64) ApiAllowedUsersAllowedUserIdDeleteRequest {
	return ApiAllowedUsersAllowedUserIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		allowedUserId: allowedUserId,
	}
}

// Execute executes the request
//  @return AllowedUsers
func (a *AllowedUsersAPIService) AllowedUsersAllowedUserIdDeleteExecute(r ApiAllowedUsersAllowedUserIdDeleteRequest) (*AllowedUsers, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AllowedUsers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllowedUsersAPIService.AllowedUsersAllowedUserIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/allowed_users/{allowed_user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"allowed_user_id"+"}", url.PathEscape(parameterValueToString(r.allowedUserId, "allowedUserId")), -1)

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

type ApiAllowedUsersAllowedUserIdGetRequest struct {
	ctx context.Context
	ApiService *AllowedUsersAPIService
	allowedUserId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAllowedUsersAllowedUserIdGetRequest) Include(include []string) ApiAllowedUsersAllowedUserIdGetRequest {
	r.include = &include
	return r
}

func (r ApiAllowedUsersAllowedUserIdGetRequest) Execute() (*AllowedUsers, *http.Response, error) {
	return r.ApiService.AllowedUsersAllowedUserIdGetExecute(r)
}

/*
AllowedUsersAllowedUserIdGet Method for AllowedUsersAllowedUserIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param allowedUserId Model id
 @return ApiAllowedUsersAllowedUserIdGetRequest
*/
func (a *AllowedUsersAPIService) AllowedUsersAllowedUserIdGet(ctx context.Context, allowedUserId int64) ApiAllowedUsersAllowedUserIdGetRequest {
	return ApiAllowedUsersAllowedUserIdGetRequest{
		ApiService: a,
		ctx: ctx,
		allowedUserId: allowedUserId,
	}
}

// Execute executes the request
//  @return AllowedUsers
func (a *AllowedUsersAPIService) AllowedUsersAllowedUserIdGetExecute(r ApiAllowedUsersAllowedUserIdGetRequest) (*AllowedUsers, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AllowedUsers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllowedUsersAPIService.AllowedUsersAllowedUserIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/allowed_users/{allowed_user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"allowed_user_id"+"}", url.PathEscape(parameterValueToString(r.allowedUserId, "allowedUserId")), -1)

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

type ApiAllowedUsersAllowedUserIdPutRequest struct {
	ctx context.Context
	ApiService *AllowedUsersAPIService
	allowedUserId int64
	allowedUsersPut *AllowedUsersPut
}

func (r ApiAllowedUsersAllowedUserIdPutRequest) AllowedUsersPut(allowedUsersPut AllowedUsersPut) ApiAllowedUsersAllowedUserIdPutRequest {
	r.allowedUsersPut = &allowedUsersPut
	return r
}

func (r ApiAllowedUsersAllowedUserIdPutRequest) Execute() (*AllowedUsers, *http.Response, error) {
	return r.ApiService.AllowedUsersAllowedUserIdPutExecute(r)
}

/*
AllowedUsersAllowedUserIdPut Method for AllowedUsersAllowedUserIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param allowedUserId Model id
 @return ApiAllowedUsersAllowedUserIdPutRequest
*/
func (a *AllowedUsersAPIService) AllowedUsersAllowedUserIdPut(ctx context.Context, allowedUserId int64) ApiAllowedUsersAllowedUserIdPutRequest {
	return ApiAllowedUsersAllowedUserIdPutRequest{
		ApiService: a,
		ctx: ctx,
		allowedUserId: allowedUserId,
	}
}

// Execute executes the request
//  @return AllowedUsers
func (a *AllowedUsersAPIService) AllowedUsersAllowedUserIdPutExecute(r ApiAllowedUsersAllowedUserIdPutRequest) (*AllowedUsers, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AllowedUsers
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllowedUsersAPIService.AllowedUsersAllowedUserIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/allowed_users/{allowed_user_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"allowed_user_id"+"}", url.PathEscape(parameterValueToString(r.allowedUserId, "allowedUserId")), -1)

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
	localVarPostBody = r.allowedUsersPut
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

type ApiAllowedUsersGetRequest struct {
	ctx context.Context
	ApiService *AllowedUsersAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiAllowedUsersGetRequest) Include(include []string) ApiAllowedUsersGetRequest {
	r.include = &include
	return r
}

func (r ApiAllowedUsersGetRequest) Execute() (*AllowedUsersGet200Response, *http.Response, error) {
	return r.ApiService.AllowedUsersGetExecute(r)
}

/*
AllowedUsersGet Method for AllowedUsersGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAllowedUsersGetRequest
*/
func (a *AllowedUsersAPIService) AllowedUsersGet(ctx context.Context) ApiAllowedUsersGetRequest {
	return ApiAllowedUsersGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AllowedUsersGet200Response
func (a *AllowedUsersAPIService) AllowedUsersGetExecute(r ApiAllowedUsersGetRequest) (*AllowedUsersGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AllowedUsersGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllowedUsersAPIService.AllowedUsersGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/allowed_users"

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

type ApiAllowedUsersPostRequest struct {
	ctx context.Context
	ApiService *AllowedUsersAPIService
	allowedUsersPostRequest *AllowedUsersPostRequest
}

func (r ApiAllowedUsersPostRequest) AllowedUsersPostRequest(allowedUsersPostRequest AllowedUsersPostRequest) ApiAllowedUsersPostRequest {
	r.allowedUsersPostRequest = &allowedUsersPostRequest
	return r
}

func (r ApiAllowedUsersPostRequest) Execute() (*AllowedUsersGet200Response, *http.Response, error) {
	return r.ApiService.AllowedUsersPostExecute(r)
}

/*
AllowedUsersPost Method for AllowedUsersPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAllowedUsersPostRequest
*/
func (a *AllowedUsersAPIService) AllowedUsersPost(ctx context.Context) ApiAllowedUsersPostRequest {
	return ApiAllowedUsersPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AllowedUsersGet200Response
func (a *AllowedUsersAPIService) AllowedUsersPostExecute(r ApiAllowedUsersPostRequest) (*AllowedUsersGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AllowedUsersGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AllowedUsersAPIService.AllowedUsersPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/allowed_users"

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
	localVarPostBody = r.allowedUsersPostRequest
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
