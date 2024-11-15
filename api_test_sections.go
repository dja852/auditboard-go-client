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


// TestSectionsAPIService TestSectionsAPI service
type TestSectionsAPIService service

type ApiTestSectionsGetRequest struct {
	ctx context.Context
	ApiService *TestSectionsAPIService
	include *[]string
}

// Possible sideloaded relationships
func (r ApiTestSectionsGetRequest) Include(include []string) ApiTestSectionsGetRequest {
	r.include = &include
	return r
}

func (r ApiTestSectionsGetRequest) Execute() (*TestSectionsGet200Response, *http.Response, error) {
	return r.ApiService.TestSectionsGetExecute(r)
}

/*
TestSectionsGet Method for TestSectionsGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTestSectionsGetRequest
*/
func (a *TestSectionsAPIService) TestSectionsGet(ctx context.Context) ApiTestSectionsGetRequest {
	return ApiTestSectionsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TestSectionsGet200Response
func (a *TestSectionsAPIService) TestSectionsGetExecute(r ApiTestSectionsGetRequest) (*TestSectionsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestSectionsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestSectionsAPIService.TestSectionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/test_sections"

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

type ApiTestSectionsPostRequest struct {
	ctx context.Context
	ApiService *TestSectionsAPIService
	testSectionsPostRequest *TestSectionsPostRequest
}

func (r ApiTestSectionsPostRequest) TestSectionsPostRequest(testSectionsPostRequest TestSectionsPostRequest) ApiTestSectionsPostRequest {
	r.testSectionsPostRequest = &testSectionsPostRequest
	return r
}

func (r ApiTestSectionsPostRequest) Execute() (*TestSectionsGet200Response, *http.Response, error) {
	return r.ApiService.TestSectionsPostExecute(r)
}

/*
TestSectionsPost Method for TestSectionsPost

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiTestSectionsPostRequest
*/
func (a *TestSectionsAPIService) TestSectionsPost(ctx context.Context) ApiTestSectionsPostRequest {
	return ApiTestSectionsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TestSectionsGet200Response
func (a *TestSectionsAPIService) TestSectionsPostExecute(r ApiTestSectionsPostRequest) (*TestSectionsGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestSectionsGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestSectionsAPIService.TestSectionsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/test_sections"

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
	localVarPostBody = r.testSectionsPostRequest
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

type ApiTestSectionsTestSectionIdDeleteRequest struct {
	ctx context.Context
	ApiService *TestSectionsAPIService
	testSectionId int64
}

func (r ApiTestSectionsTestSectionIdDeleteRequest) Execute() (*TestSections, *http.Response, error) {
	return r.ApiService.TestSectionsTestSectionIdDeleteExecute(r)
}

/*
TestSectionsTestSectionIdDelete Method for TestSectionsTestSectionIdDelete

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param testSectionId Model id
 @return ApiTestSectionsTestSectionIdDeleteRequest
*/
func (a *TestSectionsAPIService) TestSectionsTestSectionIdDelete(ctx context.Context, testSectionId int64) ApiTestSectionsTestSectionIdDeleteRequest {
	return ApiTestSectionsTestSectionIdDeleteRequest{
		ApiService: a,
		ctx: ctx,
		testSectionId: testSectionId,
	}
}

// Execute executes the request
//  @return TestSections
func (a *TestSectionsAPIService) TestSectionsTestSectionIdDeleteExecute(r ApiTestSectionsTestSectionIdDeleteRequest) (*TestSections, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestSections
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestSectionsAPIService.TestSectionsTestSectionIdDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/test_sections/{test_section_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"test_section_id"+"}", url.PathEscape(parameterValueToString(r.testSectionId, "testSectionId")), -1)

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

type ApiTestSectionsTestSectionIdGetRequest struct {
	ctx context.Context
	ApiService *TestSectionsAPIService
	testSectionId int64
	include *[]string
}

// Possible sideloaded relationships
func (r ApiTestSectionsTestSectionIdGetRequest) Include(include []string) ApiTestSectionsTestSectionIdGetRequest {
	r.include = &include
	return r
}

func (r ApiTestSectionsTestSectionIdGetRequest) Execute() (*TestSections, *http.Response, error) {
	return r.ApiService.TestSectionsTestSectionIdGetExecute(r)
}

/*
TestSectionsTestSectionIdGet Method for TestSectionsTestSectionIdGet

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param testSectionId Model id
 @return ApiTestSectionsTestSectionIdGetRequest
*/
func (a *TestSectionsAPIService) TestSectionsTestSectionIdGet(ctx context.Context, testSectionId int64) ApiTestSectionsTestSectionIdGetRequest {
	return ApiTestSectionsTestSectionIdGetRequest{
		ApiService: a,
		ctx: ctx,
		testSectionId: testSectionId,
	}
}

// Execute executes the request
//  @return TestSections
func (a *TestSectionsAPIService) TestSectionsTestSectionIdGetExecute(r ApiTestSectionsTestSectionIdGetRequest) (*TestSections, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestSections
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestSectionsAPIService.TestSectionsTestSectionIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/test_sections/{test_section_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"test_section_id"+"}", url.PathEscape(parameterValueToString(r.testSectionId, "testSectionId")), -1)

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

type ApiTestSectionsTestSectionIdPutRequest struct {
	ctx context.Context
	ApiService *TestSectionsAPIService
	testSectionId int64
	testSectionsPut *TestSectionsPut
}

func (r ApiTestSectionsTestSectionIdPutRequest) TestSectionsPut(testSectionsPut TestSectionsPut) ApiTestSectionsTestSectionIdPutRequest {
	r.testSectionsPut = &testSectionsPut
	return r
}

func (r ApiTestSectionsTestSectionIdPutRequest) Execute() (*TestSections, *http.Response, error) {
	return r.ApiService.TestSectionsTestSectionIdPutExecute(r)
}

/*
TestSectionsTestSectionIdPut Method for TestSectionsTestSectionIdPut

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param testSectionId Model id
 @return ApiTestSectionsTestSectionIdPutRequest
*/
func (a *TestSectionsAPIService) TestSectionsTestSectionIdPut(ctx context.Context, testSectionId int64) ApiTestSectionsTestSectionIdPutRequest {
	return ApiTestSectionsTestSectionIdPutRequest{
		ApiService: a,
		ctx: ctx,
		testSectionId: testSectionId,
	}
}

// Execute executes the request
//  @return TestSections
func (a *TestSectionsAPIService) TestSectionsTestSectionIdPutExecute(r ApiTestSectionsTestSectionIdPutRequest) (*TestSections, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TestSections
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TestSectionsAPIService.TestSectionsTestSectionIdPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/test_sections/{test_section_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"test_section_id"+"}", url.PathEscape(parameterValueToString(r.testSectionId, "testSectionId")), -1)

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
	localVarPostBody = r.testSectionsPut
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
