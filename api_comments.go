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
	"strings"
	"time"
)

// Linger please
var (
	_ _context.Context
)

// CommentsApiService CommentsApi service
type CommentsApiService service

type ApiCommentsCreateRequest struct {
	ctx _context.Context
	ApiService *CommentsApiService
	xAccountToken *string
	commentEndpointRequest *CommentEndpointRequest
	isDebugMode *bool
	runAsync *bool
}

func (r ApiCommentsCreateRequest) XAccountToken(xAccountToken string) ApiCommentsCreateRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiCommentsCreateRequest) CommentEndpointRequest(commentEndpointRequest CommentEndpointRequest) ApiCommentsCreateRequest {
	r.commentEndpointRequest = &commentEndpointRequest
	return r
}
func (r ApiCommentsCreateRequest) IsDebugMode(isDebugMode bool) ApiCommentsCreateRequest {
	r.isDebugMode = &isDebugMode
	return r
}
func (r ApiCommentsCreateRequest) RunAsync(runAsync bool) ApiCommentsCreateRequest {
	r.runAsync = &runAsync
	return r
}

func (r ApiCommentsCreateRequest) Execute() (CommentResponse, *_nethttp.Response, error) {
	return r.ApiService.CommentsCreateExecute(r)
}

/*
 * CommentsCreate Method for CommentsCreate
 * Creates a `Comment` object with the given values.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCommentsCreateRequest
 */
func (a *CommentsApiService) CommentsCreate(ctx _context.Context) ApiCommentsCreateRequest {
	return ApiCommentsCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return CommentResponse
 */
func (a *CommentsApiService) CommentsCreateExecute(r ApiCommentsCreateRequest) (CommentResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  CommentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommentsApiService.CommentsCreate")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/comments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}
	if r.commentEndpointRequest == nil {
		return localVarReturnValue, nil, reportError("commentEndpointRequest is required and must be specified")
	}

	if r.isDebugMode != nil {
		localVarQueryParams.Add("is_debug_mode", parameterToString(*r.isDebugMode, ""))
	}
	if r.runAsync != nil {
		localVarQueryParams.Add("run_async", parameterToString(*r.runAsync, ""))
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
	localVarHeaderParams["X-Account-Token"] = parameterToString(*r.xAccountToken, "")
	// body params
	localVarPostBody = r.commentEndpointRequest
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

type ApiCommentsListRequest struct {
	ctx _context.Context
	ApiService *CommentsApiService
	xAccountToken *string
	createdAfter *time.Time
	createdBefore *time.Time
	cursor *string
	includeDeletedData *bool
	includeRemoteData *bool
	modifiedAfter *time.Time
	modifiedBefore *time.Time
	pageSize *int32
	remoteId *string
	ticketId *string
}

func (r ApiCommentsListRequest) XAccountToken(xAccountToken string) ApiCommentsListRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiCommentsListRequest) CreatedAfter(createdAfter time.Time) ApiCommentsListRequest {
	r.createdAfter = &createdAfter
	return r
}
func (r ApiCommentsListRequest) CreatedBefore(createdBefore time.Time) ApiCommentsListRequest {
	r.createdBefore = &createdBefore
	return r
}
func (r ApiCommentsListRequest) Cursor(cursor string) ApiCommentsListRequest {
	r.cursor = &cursor
	return r
}
func (r ApiCommentsListRequest) IncludeDeletedData(includeDeletedData bool) ApiCommentsListRequest {
	r.includeDeletedData = &includeDeletedData
	return r
}
func (r ApiCommentsListRequest) IncludeRemoteData(includeRemoteData bool) ApiCommentsListRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}
func (r ApiCommentsListRequest) ModifiedAfter(modifiedAfter time.Time) ApiCommentsListRequest {
	r.modifiedAfter = &modifiedAfter
	return r
}
func (r ApiCommentsListRequest) ModifiedBefore(modifiedBefore time.Time) ApiCommentsListRequest {
	r.modifiedBefore = &modifiedBefore
	return r
}
func (r ApiCommentsListRequest) PageSize(pageSize int32) ApiCommentsListRequest {
	r.pageSize = &pageSize
	return r
}
func (r ApiCommentsListRequest) RemoteId(remoteId string) ApiCommentsListRequest {
	r.remoteId = &remoteId
	return r
}
func (r ApiCommentsListRequest) TicketId(ticketId string) ApiCommentsListRequest {
	r.ticketId = &ticketId
	return r
}

func (r ApiCommentsListRequest) Execute() (PaginatedCommentList, *_nethttp.Response, error) {
	return r.ApiService.CommentsListExecute(r)
}

/*
 * CommentsList Method for CommentsList
 * Returns a list of `Comment` objects.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCommentsListRequest
 */
func (a *CommentsApiService) CommentsList(ctx _context.Context) ApiCommentsListRequest {
	return ApiCommentsListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return PaginatedCommentList
 */
func (a *CommentsApiService) CommentsListExecute(r ApiCommentsListRequest) (PaginatedCommentList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  PaginatedCommentList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommentsApiService.CommentsList")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/comments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.createdAfter != nil {
		localVarQueryParams.Add("created_after", parameterToString(*r.createdAfter, ""))
	}
	if r.createdBefore != nil {
		localVarQueryParams.Add("created_before", parameterToString(*r.createdBefore, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.includeDeletedData != nil {
		localVarQueryParams.Add("include_deleted_data", parameterToString(*r.includeDeletedData, ""))
	}
	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
	}
	if r.modifiedAfter != nil {
		localVarQueryParams.Add("modified_after", parameterToString(*r.modifiedAfter, ""))
	}
	if r.modifiedBefore != nil {
		localVarQueryParams.Add("modified_before", parameterToString(*r.modifiedBefore, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.remoteId != nil {
		localVarQueryParams.Add("remote_id", parameterToString(*r.remoteId, ""))
	}
	if r.ticketId != nil {
		localVarQueryParams.Add("ticket_id", parameterToString(*r.ticketId, ""))
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

type ApiCommentsMetaPostRetrieveRequest struct {
	ctx _context.Context
	ApiService *CommentsApiService
	xAccountToken *string
}

func (r ApiCommentsMetaPostRetrieveRequest) XAccountToken(xAccountToken string) ApiCommentsMetaPostRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}

func (r ApiCommentsMetaPostRetrieveRequest) Execute() (MetaResponse, *_nethttp.Response, error) {
	return r.ApiService.CommentsMetaPostRetrieveExecute(r)
}

/*
 * CommentsMetaPostRetrieve Method for CommentsMetaPostRetrieve
 * Returns metadata for `Comment` POSTs.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiCommentsMetaPostRetrieveRequest
 */
func (a *CommentsApiService) CommentsMetaPostRetrieve(ctx _context.Context) ApiCommentsMetaPostRetrieveRequest {
	return ApiCommentsMetaPostRetrieveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return MetaResponse
 */
func (a *CommentsApiService) CommentsMetaPostRetrieveExecute(r ApiCommentsMetaPostRetrieveRequest) (MetaResponse, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  MetaResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommentsApiService.CommentsMetaPostRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/comments/meta/post"

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

type ApiCommentsRetrieveRequest struct {
	ctx _context.Context
	ApiService *CommentsApiService
	xAccountToken *string
	id string
	includeRemoteData *bool
}

func (r ApiCommentsRetrieveRequest) XAccountToken(xAccountToken string) ApiCommentsRetrieveRequest {
	r.xAccountToken = &xAccountToken
	return r
}
func (r ApiCommentsRetrieveRequest) IncludeRemoteData(includeRemoteData bool) ApiCommentsRetrieveRequest {
	r.includeRemoteData = &includeRemoteData
	return r
}

func (r ApiCommentsRetrieveRequest) Execute() (Comment, *_nethttp.Response, error) {
	return r.ApiService.CommentsRetrieveExecute(r)
}

/*
 * CommentsRetrieve Method for CommentsRetrieve
 * Returns a `Comment` object with the given `id`.
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiCommentsRetrieveRequest
 */
func (a *CommentsApiService) CommentsRetrieve(ctx _context.Context, id string) ApiCommentsRetrieveRequest {
	return ApiCommentsRetrieveRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Comment
 */
func (a *CommentsApiService) CommentsRetrieveExecute(r ApiCommentsRetrieveRequest) (Comment, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Comment
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CommentsApiService.CommentsRetrieve")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/comments/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", _neturl.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.xAccountToken == nil {
		return localVarReturnValue, nil, reportError("xAccountToken is required and must be specified")
	}

	if r.includeRemoteData != nil {
		localVarQueryParams.Add("include_remote_data", parameterToString(*r.includeRemoteData, ""))
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
