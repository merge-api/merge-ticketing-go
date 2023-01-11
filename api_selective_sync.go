/*
 * Merge Ticketing API
 *
 * The unified API for building rich integrations with multiple Ticketing platforms.
 *
 * API version: 1.0
 * Contact: hello@merge.dev
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package merge_ticketing_client

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// SelectiveSyncApiService SelectiveSyncApi service
type SelectiveSyncApiService service

type ApiSelectiveSyncConfigurationsListRequest struct {
	ctx _context.Context
	ApiService *SelectiveSyncApiService
	xAccountToken *string
}

func (r ApiSelectiveSyncConfigurationsListRequest) XAccountToken(xAccountToken string) ApiSelectiveSyncConfigurationsListRequest {
	r.xAccountToken = &xAccountToken
	return r
}

func (r ApiSelectiveSyncConfigurationsListRequest) Execute() ([]LinkedAccountSelectiveSyncConfiguration, *_nethttp.Response, error) {
	return r.ApiService.SelectiveSyncConfigurationsListExecute(r)
}

/*
 * SelectiveSyncConfigurationsList Method for SelectiveSyncConfigurationsList
 * Get a linked account's selective syncs.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSelectiveSyncConfigurationsListRequest
 */
func (a *SelectiveSyncApiService) SelectiveSyncConfigurationsList(ctx _context.Context) ApiSelectiveSyncConfigurationsListRequest {
	return ApiSelectiveSyncConfigurationsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []LinkedAccountSelectiveSyncConfiguration
 */
func (a *SelectiveSyncApiService) SelectiveSyncConfigurationsListExecute(r ApiSelectiveSyncConfigurationsListRequest) ([]LinkedAccountSelectiveSyncConfiguration, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []LinkedAccountSelectiveSyncConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelectiveSyncApiService.SelectiveSyncConfigurationsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/selective-sync/configurations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSelectiveSyncConfigurationsUpdateRequest struct {
	ctx _context.Context
	ApiService *SelectiveSyncApiService
	xAccountToken *string
	linkedAccountSelectiveSyncConfigurationListRequest *LinkedAccountSelectiveSyncConfigurationListRequest
}

func (r ApiSelectiveSyncConfigurationsUpdateRequest) XAccountToken(xAccountToken string) ApiSelectiveSyncConfigurationsUpdateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiSelectiveSyncConfigurationsUpdateRequest) LinkedAccountSelectiveSyncConfigurationListRequest(linkedAccountSelectiveSyncConfigurationListRequest LinkedAccountSelectiveSyncConfigurationListRequest) ApiSelectiveSyncConfigurationsUpdateRequest {
	r.linkedAccountSelectiveSyncConfigurationListRequest = &linkedAccountSelectiveSyncConfigurationListRequest
	return r
}

func (r ApiSelectiveSyncConfigurationsUpdateRequest) Execute() ([]LinkedAccountSelectiveSyncConfiguration, *_nethttp.Response, error) {
	return r.ApiService.SelectiveSyncConfigurationsUpdateExecute(r)
}

/*
 * SelectiveSyncConfigurationsUpdate Method for SelectiveSyncConfigurationsUpdate
 * Replace a linked account's selective syncs.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSelectiveSyncConfigurationsUpdateRequest
 */
func (a *SelectiveSyncApiService) SelectiveSyncConfigurationsUpdate(ctx _context.Context) ApiSelectiveSyncConfigurationsUpdateRequest {
	return ApiSelectiveSyncConfigurationsUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []LinkedAccountSelectiveSyncConfiguration
 */
func (a *SelectiveSyncApiService) SelectiveSyncConfigurationsUpdateExecute(r ApiSelectiveSyncConfigurationsUpdateRequest) ([]LinkedAccountSelectiveSyncConfiguration, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []LinkedAccountSelectiveSyncConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelectiveSyncApiService.SelectiveSyncConfigurationsUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/selective-sync/configurations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.linkedAccountSelectiveSyncConfigurationListRequest == nil {
		return localVarReturnValue, nil, reportError("linkedAccountSelectiveSyncConfigurationListRequest is required and must be specified")
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	// body params
	localVarPostBody = r.linkedAccountSelectiveSyncConfigurationListRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiSelectiveSyncMetaListRequest struct {
	ctx _context.Context
	ApiService *SelectiveSyncApiService
	xAccountToken *string
	commonModel *string
	cursor *string
	pageSize *int32
}

func (r ApiSelectiveSyncMetaListRequest) XAccountToken(xAccountToken string) ApiSelectiveSyncMetaListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiSelectiveSyncMetaListRequest) CommonModel(commonModel string) ApiSelectiveSyncMetaListRequest {
	r.commonModel = &commonModel
	return r
}
func (r ApiSelectiveSyncMetaListRequest) Cursor(cursor string) ApiSelectiveSyncMetaListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiSelectiveSyncMetaListRequest) PageSize(pageSize int32) ApiSelectiveSyncMetaListRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiSelectiveSyncMetaListRequest) Execute() (PaginatedConditionSchemaList, *_nethttp.Response, error) {
	return r.ApiService.SelectiveSyncMetaListExecute(r)
}

/*
 * SelectiveSyncMetaList Method for SelectiveSyncMetaList
 * Get metadata for the conditions available to a linked account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiSelectiveSyncMetaListRequest
 */
func (a *SelectiveSyncApiService) SelectiveSyncMetaList(ctx _context.Context) ApiSelectiveSyncMetaListRequest {
	return ApiSelectiveSyncMetaListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedConditionSchemaList
 */
func (a *SelectiveSyncApiService) SelectiveSyncMetaListExecute(r ApiSelectiveSyncMetaListRequest) (PaginatedConditionSchemaList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedConditionSchemaList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SelectiveSyncApiService.SelectiveSyncMetaList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/selective-sync/meta"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.commonModel != nil {
		localVarQueryParams.Add("common_model", parameterToString(*r.commonModel, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["tokenAuth"]; ok {
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
