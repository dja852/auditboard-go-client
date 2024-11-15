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


// KeyPerformanceIndicatorsAPIService KeyPerformanceIndicatorsAPI service
type KeyPerformanceIndicatorsAPIService service

type ApiKeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostRequest struct {
	ctx context.Context
	ApiService *KeyPerformanceIndicatorsAPIService
	keyPerformanceIndicatorId int64
	keyPerformanceIndicatorAddValuePost *KeyPerformanceIndicatorAddValuePost
}

func (r ApiKeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostRequest) KeyPerformanceIndicatorAddValuePost(keyPerformanceIndicatorAddValuePost KeyPerformanceIndicatorAddValuePost) ApiKeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostRequest {
	r.keyPerformanceIndicatorAddValuePost = &keyPerformanceIndicatorAddValuePost
	return r
}

func (r ApiKeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostRequest) Execute() (*KeyPerformanceIndicatorsAddValue, *http.Response, error) {
	return r.ApiService.KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostExecute(r)
}

/*
KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost Method for KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param keyPerformanceIndicatorId Key Performance Indicator Id
 @return ApiKeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostRequest
*/
func (a *KeyPerformanceIndicatorsAPIService) KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost(ctx context.Context, keyPerformanceIndicatorId int64) ApiKeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostRequest {
	return ApiKeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostRequest{
		ApiService: a,
		ctx: ctx,
		keyPerformanceIndicatorId: keyPerformanceIndicatorId,
	}
}

// Execute executes the request
//  @return KeyPerformanceIndicatorsAddValue
func (a *KeyPerformanceIndicatorsAPIService) KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostExecute(r ApiKeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePostRequest) (*KeyPerformanceIndicatorsAddValue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *KeyPerformanceIndicatorsAddValue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "KeyPerformanceIndicatorsAPIService.KeyPerformanceIndicatorsKeyPerformanceIndicatorIdAddValuePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/key_performance_indicators/{key_performance_indicator_id}/add_value"
	localVarPath = strings.Replace(localVarPath, "{"+"key_performance_indicator_id"+"}", url.PathEscape(parameterValueToString(r.keyPerformanceIndicatorId, "keyPerformanceIndicatorId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.keyPerformanceIndicatorAddValuePost == nil {
		return localVarReturnValue, nil, reportError("keyPerformanceIndicatorAddValuePost is required and must be specified")
	}

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
	localVarPostBody = r.keyPerformanceIndicatorAddValuePost
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
