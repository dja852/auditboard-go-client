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


// FinancialAccountsAPIService FinancialAccountsAPI service
type FinancialAccountsAPIService service

type ApiFinancialAccountsFinancialAccountIdDeleteRequest struct {
	ctx context.Context
	ApiService *FinancialAccountsAPIService
	financialAccountId int64
}

func (r ApiFinancialAccountsFinancialAccountIdDeleteRequest) Execute() (*FinancialAccounts, *http.Response, error) {
	return r.ApiService.FinancialAccountsFinancialAccountIdDeleteExecute(r)
}

/*
FinancialAccountsFinancialAccountIdDelete Method for FinancialAccountsFinancialAccountIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param financialAccountId Model id
 @return ApiFinancialAccountsFinancialAccountIdDeleteRequest
*/
func (a *FinancialAccountsAPIService) FinancialAccountsFinancialAccountIdDelete(ctx context.Context, financialAccountId int64) ApiFinancialAccountsFinancialAccountIdDeleteRequest {
	return ApiFinancialAccountsFinancialAccountIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		financialAccountId: financialAccountId,
	}
}

// Execute executes the request
//  @return FinancialAccounts
func (a *FinancialAccountsAPIService) FinancialAccountsFinancialAccountIdDeleteExecute(r ApiFinancialAccountsFinancialAccountIdDeleteRequest) (*FinancialAccounts, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FinancialAccounts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinancialAccountsAPIService.FinancialAccountsFinancialAccountIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/financial_accounts/{financial_account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"financial_account_id"+"}", url.PathEscape(parameterValueToString(r.financialAccountId, "financialAccountId")), -1)

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

type ApiFinancialAccountsFinancialAccountIdGetRequest struct {
	ctx context.Context
	ApiService *FinancialAccountsAPIService
	financialAccountId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiFinancialAccountsFinancialAccountIdGetRequest) Include(include []string) ApiFinancialAccountsFinancialAccountIdGetRequest {
	r.include = &include
	return r
}

func (r ApiFinancialAccountsFinancialAccountIdGetRequest) Execute() (*FinancialAccounts, *http.Response, error) {
	return r.ApiService.FinancialAccountsFinancialAccountIdGetExecute(r)
}

/*
FinancialAccountsFinancialAccountIdGet Method for FinancialAccountsFinancialAccountIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param financialAccountId Model id
 @return ApiFinancialAccountsFinancialAccountIdGetRequest
*/
func (a *FinancialAccountsAPIService) FinancialAccountsFinancialAccountIdGet(ctx context.Context, financialAccountId int64) ApiFinancialAccountsFinancialAccountIdGetRequest {
	return ApiFinancialAccountsFinancialAccountIdGetRequest{
		ApiService: a,
		ctx: ctx,
		financialAccountId: financialAccountId,
	}
}

// Execute executes the request
//  @return FinancialAccounts
func (a *FinancialAccountsAPIService) FinancialAccountsFinancialAccountIdGetExecute(r ApiFinancialAccountsFinancialAccountIdGetRequest) (*FinancialAccounts, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FinancialAccounts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinancialAccountsAPIService.FinancialAccountsFinancialAccountIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/financial_accounts/{financial_account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"financial_account_id"+"}", url.PathEscape(parameterValueToString(r.financialAccountId, "financialAccountId")), -1)

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

type ApiFinancialAccountsFinancialAccountIdPutRequest struct {
	ctx context.Context
	ApiService *FinancialAccountsAPIService
	financialAccountId int64
	financialAccountsPut *FinancialAccountsPut
}

func (r ApiFinancialAccountsFinancialAccountIdPutRequest) FinancialAccountsPut(financialAccountsPut FinancialAccountsPut) ApiFinancialAccountsFinancialAccountIdPutRequest {
	r.financialAccountsPut = &financialAccountsPut
	return r
}

func (r ApiFinancialAccountsFinancialAccountIdPutRequest) Execute() (*FinancialAccounts, *http.Response, error) {
	return r.ApiService.FinancialAccountsFinancialAccountIdPutExecute(r)
}

/*
FinancialAccountsFinancialAccountIdPut Method for FinancialAccountsFinancialAccountIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param financialAccountId Model id
 @return ApiFinancialAccountsFinancialAccountIdPutRequest
*/
func (a *FinancialAccountsAPIService) FinancialAccountsFinancialAccountIdPut(ctx context.Context, financialAccountId int64) ApiFinancialAccountsFinancialAccountIdPutRequest {
	return ApiFinancialAccountsFinancialAccountIdPutRequest{
		ApiService: a,
		ctx: ctx,
		financialAccountId: financialAccountId,
	}
}

// Execute executes the request
//  @return FinancialAccounts
func (a *FinancialAccountsAPIService) FinancialAccountsFinancialAccountIdPutExecute(r ApiFinancialAccountsFinancialAccountIdPutRequest) (*FinancialAccounts, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FinancialAccounts
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinancialAccountsAPIService.FinancialAccountsFinancialAccountIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/financial_accounts/{financial_account_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"financial_account_id"+"}", url.PathEscape(parameterValueToString(r.financialAccountId, "financialAccountId")), -1)

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
	localVarPostBody = r.financialAccountsPut
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

type ApiFinancialAccountsGetRequest struct {
	ctx context.Context
	ApiService *FinancialAccountsAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiFinancialAccountsGetRequest) Include(include []string) ApiFinancialAccountsGetRequest {
	r.include = &include
	return r
}

func (r ApiFinancialAccountsGetRequest) Execute() (*FinancialAccountsGet200Response, *http.Response, error) {
	return r.ApiService.FinancialAccountsGetExecute(r)
}

/*
FinancialAccountsGet Method for FinancialAccountsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFinancialAccountsGetRequest
*/
func (a *FinancialAccountsAPIService) FinancialAccountsGet(ctx context.Context) ApiFinancialAccountsGetRequest {
	return ApiFinancialAccountsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FinancialAccountsGet200Response
func (a *FinancialAccountsAPIService) FinancialAccountsGetExecute(r ApiFinancialAccountsGetRequest) (*FinancialAccountsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FinancialAccountsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinancialAccountsAPIService.FinancialAccountsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/financial_accounts"

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

type ApiFinancialAccountsPostRequest struct {
	ctx context.Context
	ApiService *FinancialAccountsAPIService
	financialAccountsPostRequest *FinancialAccountsPostRequest
}

func (r ApiFinancialAccountsPostRequest) FinancialAccountsPostRequest(financialAccountsPostRequest FinancialAccountsPostRequest) ApiFinancialAccountsPostRequest {
	r.financialAccountsPostRequest = &financialAccountsPostRequest
	return r
}

func (r ApiFinancialAccountsPostRequest) Execute() (*FinancialAccountsGet200Response, *http.Response, error) {
	return r.ApiService.FinancialAccountsPostExecute(r)
}

/*
FinancialAccountsPost Method for FinancialAccountsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFinancialAccountsPostRequest
*/
func (a *FinancialAccountsAPIService) FinancialAccountsPost(ctx context.Context) ApiFinancialAccountsPostRequest {
	return ApiFinancialAccountsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FinancialAccountsGet200Response
func (a *FinancialAccountsAPIService) FinancialAccountsPostExecute(r ApiFinancialAccountsPostRequest) (*FinancialAccountsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FinancialAccountsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FinancialAccountsAPIService.FinancialAccountsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/financial_accounts"

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
	localVarPostBody = r.financialAccountsPostRequest
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
