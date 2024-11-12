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


// ControlsDataArchivesAPIService ControlsDataArchivesAPI service
type ControlsDataArchivesAPIService service

type ApiControlsDataArchivesControlsDataArchiveIdGetRequest struct {
	ctx context.Context
	ApiService *ControlsDataArchivesAPIService
	controlsDataArchiveId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiControlsDataArchivesControlsDataArchiveIdGetRequest) Include(include []string) ApiControlsDataArchivesControlsDataArchiveIdGetRequest {
	r.include = &include
	return r
}

func (r ApiControlsDataArchivesControlsDataArchiveIdGetRequest) Execute() (*ControlsDataArchives, *http.Response, error) {
	return r.ApiService.ControlsDataArchivesControlsDataArchiveIdGetExecute(r)
}

/*
ControlsDataArchivesControlsDataArchiveIdGet Method for ControlsDataArchivesControlsDataArchiveIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param controlsDataArchiveId Model id
 @return ApiControlsDataArchivesControlsDataArchiveIdGetRequest
*/
func (a *ControlsDataArchivesAPIService) ControlsDataArchivesControlsDataArchiveIdGet(ctx context.Context, controlsDataArchiveId int64) ApiControlsDataArchivesControlsDataArchiveIdGetRequest {
	return ApiControlsDataArchivesControlsDataArchiveIdGetRequest{
		ApiService: a,
		ctx: ctx,
		controlsDataArchiveId: controlsDataArchiveId,
	}
}

// Execute executes the request
//  @return ControlsDataArchives
func (a *ControlsDataArchivesAPIService) ControlsDataArchivesControlsDataArchiveIdGetExecute(r ApiControlsDataArchivesControlsDataArchiveIdGetRequest) (*ControlsDataArchives, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControlsDataArchives
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControlsDataArchivesAPIService.ControlsDataArchivesControlsDataArchiveIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controls_data_archives/{controls_data_archive_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"controls_data_archive_id"+"}", url.PathEscape(parameterValueToString(r.controlsDataArchiveId, "controlsDataArchiveId")), -1)

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

type ApiControlsDataArchivesGetRequest struct {
	ctx context.Context
	ApiService *ControlsDataArchivesAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiControlsDataArchivesGetRequest) Include(include []string) ApiControlsDataArchivesGetRequest {
	r.include = &include
	return r
}

func (r ApiControlsDataArchivesGetRequest) Execute() (*ControlsDataArchivesGet200Response, *http.Response, error) {
	return r.ApiService.ControlsDataArchivesGetExecute(r)
}

/*
ControlsDataArchivesGet Method for ControlsDataArchivesGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiControlsDataArchivesGetRequest
*/
func (a *ControlsDataArchivesAPIService) ControlsDataArchivesGet(ctx context.Context) ApiControlsDataArchivesGetRequest {
	return ApiControlsDataArchivesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ControlsDataArchivesGet200Response
func (a *ControlsDataArchivesAPIService) ControlsDataArchivesGetExecute(r ApiControlsDataArchivesGetRequest) (*ControlsDataArchivesGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ControlsDataArchivesGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ControlsDataArchivesAPIService.ControlsDataArchivesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/controls_data_archives"

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