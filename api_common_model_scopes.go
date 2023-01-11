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

// CommonModelScopesApiService CommonModelScopesApi service
type CommonModelScopesApiService service

type ApiCommonModelScopesRetrieveRequest struct {
	ctx _context.Context
	ApiService *CommonModelScopesApiService
	integrationSlug *string
	linkedAccountId *string
}

func (r ApiCommonModelScopesRetrieveRequest) IntegrationSlug(integrationSlug string) ApiCommonModelScopesRetrieveRequest {
	r.integrationSlug = &integrationSlug
	return r
}
func (r ApiCommonModelScopesRetrieveRequest) LinkedAccountId(linkedAccountId string) ApiCommonModelScopesRetrieveRequest {
	r.linkedAccountId = &linkedAccountId
	return r
}

func (r ApiCommonModelScopesRetrieveRequest) Execute() (CommonModelScopes, *_nethttp.Response, error) {
	return r.ApiService.CommonModelScopesRetrieveExecute(r)
}

/*
 * CommonModelScopesRetrieve Method for CommonModelScopesRetrieve
 * Fetch the configuration of what data is saved by Merge when syncing. Common model scopes are set as a default across all linked accounts, but can be updated to have greater granularity per integration or account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCommonModelScopesRetrieveRequest
 */
func (a *CommonModelScopesApiService) CommonModelScopesRetrieve(ctx _context.Context) ApiCommonModelScopesRetrieveRequest {
	return ApiCommonModelScopesRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return CommonModelScopes
 */
func (a *CommonModelScopesApiService) CommonModelScopesRetrieveExecute(r ApiCommonModelScopesRetrieveRequest) (CommonModelScopes, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CommonModelScopes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommonModelScopesApiService.CommonModelScopesRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/common-model-scopes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.integrationSlug != nil {
		localVarQueryParams.Add("integration_slug", parameterToString(*r.integrationSlug, ""))
	}
	if r.linkedAccountId != nil {
		localVarQueryParams.Add("linked_account_id", parameterToString(*r.linkedAccountId, ""))
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

type ApiCommonModelScopesUpdateRequest struct {
	ctx _context.Context
	ApiService *CommonModelScopesApiService
	commonModelScopesUpdateSerializer *CommonModelScopesUpdateSerializer
	integrationSlug *string
	linkedAccountId *string
}

func (r ApiCommonModelScopesUpdateRequest) CommonModelScopesUpdateSerializer(commonModelScopesUpdateSerializer CommonModelScopesUpdateSerializer) ApiCommonModelScopesUpdateRequest {
	r.commonModelScopesUpdateSerializer = &commonModelScopesUpdateSerializer
	return r
}
func (r ApiCommonModelScopesUpdateRequest) IntegrationSlug(integrationSlug string) ApiCommonModelScopesUpdateRequest {
	r.integrationSlug = &integrationSlug
	return r
}
func (r ApiCommonModelScopesUpdateRequest) LinkedAccountId(linkedAccountId string) ApiCommonModelScopesUpdateRequest {
	r.linkedAccountId = &linkedAccountId
	return r
}

func (r ApiCommonModelScopesUpdateRequest) Execute() (CommonModelScopes, *_nethttp.Response, error) {
	return r.ApiService.CommonModelScopesUpdateExecute(r)
}

/*
 * CommonModelScopesUpdate Method for CommonModelScopesUpdate
 * Update the configuration of what data is saved by Merge when syncing. Common model scopes are set as a default across all linked accounts, but can be updated to have greater granularity per integration or account.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCommonModelScopesUpdateRequest
 */
func (a *CommonModelScopesApiService) CommonModelScopesUpdate(ctx _context.Context) ApiCommonModelScopesUpdateRequest {
	return ApiCommonModelScopesUpdateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return CommonModelScopes
 */
func (a *CommonModelScopesApiService) CommonModelScopesUpdateExecute(r ApiCommonModelScopesUpdateRequest) (CommonModelScopes, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CommonModelScopes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommonModelScopesApiService.CommonModelScopesUpdate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/common-model-scopes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.commonModelScopesUpdateSerializer == nil {
		return localVarReturnValue, nil, reportError("commonModelScopesUpdateSerializer is required and must be specified")
	}

	if r.integrationSlug != nil {
		localVarQueryParams.Add("integration_slug", parameterToString(*r.integrationSlug, ""))
	}
	if r.linkedAccountId != nil {
		localVarQueryParams.Add("linked_account_id", parameterToString(*r.linkedAccountId, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json", "application/x-www-form-urlencoded", "multipart/form-data"}

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
	localVarPostBody = r.commonModelScopesUpdateSerializer
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
