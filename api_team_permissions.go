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


// TeamPermissionsAPIService TeamPermissionsAPI service
type TeamPermissionsAPIService service

type ApiTeamPermissionsGetRequest struct {
	ctx context.Context
	ApiService *TeamPermissionsAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiTeamPermissionsGetRequest) Include(include []string) ApiTeamPermissionsGetRequest {
	r.include = &include
	return r
}

func (r ApiTeamPermissionsGetRequest) Execute() (*TeamPermissionsGet200Response, *http.Response, error) {
	return r.ApiService.TeamPermissionsGetExecute(r)
}

/*
TeamPermissionsGet Method for TeamPermissionsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTeamPermissionsGetRequest
*/
func (a *TeamPermissionsAPIService) TeamPermissionsGet(ctx context.Context) ApiTeamPermissionsGetRequest {
	return ApiTeamPermissionsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TeamPermissionsGet200Response
func (a *TeamPermissionsAPIService) TeamPermissionsGetExecute(r ApiTeamPermissionsGetRequest) (*TeamPermissionsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamPermissionsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamPermissionsAPIService.TeamPermissionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/team_permissions"

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

type ApiTeamPermissionsPostRequest struct {
	ctx context.Context
	ApiService *TeamPermissionsAPIService
	teamPermissionsPostRequest *TeamPermissionsPostRequest
}

func (r ApiTeamPermissionsPostRequest) TeamPermissionsPostRequest(teamPermissionsPostRequest TeamPermissionsPostRequest) ApiTeamPermissionsPostRequest {
	r.teamPermissionsPostRequest = &teamPermissionsPostRequest
	return r
}

func (r ApiTeamPermissionsPostRequest) Execute() (*TeamPermissionsGet200Response, *http.Response, error) {
	return r.ApiService.TeamPermissionsPostExecute(r)
}

/*
TeamPermissionsPost Method for TeamPermissionsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTeamPermissionsPostRequest
*/
func (a *TeamPermissionsAPIService) TeamPermissionsPost(ctx context.Context) ApiTeamPermissionsPostRequest {
	return ApiTeamPermissionsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TeamPermissionsGet200Response
func (a *TeamPermissionsAPIService) TeamPermissionsPostExecute(r ApiTeamPermissionsPostRequest) (*TeamPermissionsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamPermissionsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamPermissionsAPIService.TeamPermissionsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/team_permissions"

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
	localVarPostBody = r.teamPermissionsPostRequest
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

type ApiTeamPermissionsTeamPermissionIdDeleteRequest struct {
	ctx context.Context
	ApiService *TeamPermissionsAPIService
	teamPermissionId int64
}

func (r ApiTeamPermissionsTeamPermissionIdDeleteRequest) Execute() (*TeamPermissions, *http.Response, error) {
	return r.ApiService.TeamPermissionsTeamPermissionIdDeleteExecute(r)
}

/*
TeamPermissionsTeamPermissionIdDelete Method for TeamPermissionsTeamPermissionIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamPermissionId Model id
 @return ApiTeamPermissionsTeamPermissionIdDeleteRequest
*/
func (a *TeamPermissionsAPIService) TeamPermissionsTeamPermissionIdDelete(ctx context.Context, teamPermissionId int64) ApiTeamPermissionsTeamPermissionIdDeleteRequest {
	return ApiTeamPermissionsTeamPermissionIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		teamPermissionId: teamPermissionId,
	}
}

// Execute executes the request
//  @return TeamPermissions
func (a *TeamPermissionsAPIService) TeamPermissionsTeamPermissionIdDeleteExecute(r ApiTeamPermissionsTeamPermissionIdDeleteRequest) (*TeamPermissions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamPermissions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamPermissionsAPIService.TeamPermissionsTeamPermissionIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/team_permissions/{team_permission_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_permission_id"+"}", url.PathEscape(parameterValueToString(r.teamPermissionId, "teamPermissionId")), -1)

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

type ApiTeamPermissionsTeamPermissionIdGetRequest struct {
	ctx context.Context
	ApiService *TeamPermissionsAPIService
	teamPermissionId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiTeamPermissionsTeamPermissionIdGetRequest) Include(include []string) ApiTeamPermissionsTeamPermissionIdGetRequest {
	r.include = &include
	return r
}

func (r ApiTeamPermissionsTeamPermissionIdGetRequest) Execute() (*TeamPermissions, *http.Response, error) {
	return r.ApiService.TeamPermissionsTeamPermissionIdGetExecute(r)
}

/*
TeamPermissionsTeamPermissionIdGet Method for TeamPermissionsTeamPermissionIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamPermissionId Model id
 @return ApiTeamPermissionsTeamPermissionIdGetRequest
*/
func (a *TeamPermissionsAPIService) TeamPermissionsTeamPermissionIdGet(ctx context.Context, teamPermissionId int64) ApiTeamPermissionsTeamPermissionIdGetRequest {
	return ApiTeamPermissionsTeamPermissionIdGetRequest{
		ApiService: a,
		ctx: ctx,
		teamPermissionId: teamPermissionId,
	}
}

// Execute executes the request
//  @return TeamPermissions
func (a *TeamPermissionsAPIService) TeamPermissionsTeamPermissionIdGetExecute(r ApiTeamPermissionsTeamPermissionIdGetRequest) (*TeamPermissions, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamPermissions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamPermissionsAPIService.TeamPermissionsTeamPermissionIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/team_permissions/{team_permission_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"team_permission_id"+"}", url.PathEscape(parameterValueToString(r.teamPermissionId, "teamPermissionId")), -1)

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
