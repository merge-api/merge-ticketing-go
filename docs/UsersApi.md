# \UsersApi

All URIs are relative to *https://api.merge.dev/api/ticketing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersList**](UsersApi.md#UsersList) | **Get** /users | 
[**UsersRetrieve**](UsersApi.md#UsersRetrieve) | **Get** /users/{id} | 



## UsersList

> PaginatedUserList UsersList(ctx).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).EmailAddress(emailAddress).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    emailAddress := "emailAddress_example" // string | If provided, will only return users with emails equal to this value (case insensitive). (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, only objects synced by Merge after this date time will be returned. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, only objects synced by Merge before this date time will be returned. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersList(context.Background()).XAccountToken(xAccountToken).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).EmailAddress(emailAddress).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).RemoteId(remoteId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersList`: PaginatedUserList
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUsersListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **emailAddress** | **string** | If provided, will only return users with emails equal to this value (case insensitive). | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **modifiedAfter** | **time.Time** | If provided, only objects synced by Merge after this date time will be returned. | 
 **modifiedBefore** | **time.Time** | If provided, only objects synced by Merge before this date time will be returned. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 

### Return type

[**PaginatedUserList**](PaginatedUserList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UsersRetrieve

> User UsersRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xAccountToken := "xAccountToken_example" // string | Token identifying the end user.
    id := TODO // string | 
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersApi.UsersRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.UsersRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UsersRetrieve`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.UsersRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUsersRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 

### Return type

[**User**](User.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

