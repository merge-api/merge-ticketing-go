# \TicketsApi

All URIs are relative to *https://api.merge.dev/api/ticketing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TicketsCollaboratorsList**](TicketsApi.md#TicketsCollaboratorsList) | **Get** /tickets/{parent_id}/collaborators | 
[**TicketsCreate**](TicketsApi.md#TicketsCreate) | **Post** /tickets | 
[**TicketsList**](TicketsApi.md#TicketsList) | **Get** /tickets | 
[**TicketsMetaPatchRetrieve**](TicketsApi.md#TicketsMetaPatchRetrieve) | **Get** /tickets/meta/patch/{id} | 
[**TicketsMetaPostRetrieve**](TicketsApi.md#TicketsMetaPostRetrieve) | **Get** /tickets/meta/post | 
[**TicketsPartialUpdate**](TicketsApi.md#TicketsPartialUpdate) | **Patch** /tickets/{id} | 
[**TicketsRetrieve**](TicketsApi.md#TicketsRetrieve) | **Get** /tickets/{id} | 



## TicketsCollaboratorsList

> PaginatedUserList TicketsCollaboratorsList(ctx, parentId).XAccountToken(xAccountToken).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).PageSize(pageSize).Execute()





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
    parentId := TODO // string | 
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TicketsApi.TicketsCollaboratorsList(context.Background(), parentId).XAccountToken(xAccountToken).Cursor(cursor).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).PageSize(pageSize).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketsApi.TicketsCollaboratorsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketsCollaboratorsList`: PaginatedUserList
    fmt.Fprintf(os.Stdout, "Response from `TicketsApi.TicketsCollaboratorsList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentId** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketsCollaboratorsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **cursor** | **string** | The pagination cursor value. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **pageSize** | **int32** | Number of results to return per page. | 

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


## TicketsCreate

> TicketResponse TicketsCreate(ctx).XAccountToken(xAccountToken).TicketEndpointRequest(ticketEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    ticketEndpointRequest := *openapiclient.NewTicketEndpointRequest(*openapiclient.NewTicketRequest()) // TicketEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TicketsApi.TicketsCreate(context.Background()).XAccountToken(xAccountToken).TicketEndpointRequest(ticketEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketsApi.TicketsCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketsCreate`: TicketResponse
    fmt.Fprintf(os.Stdout, "Response from `TicketsApi.TicketsCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketsCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **ticketEndpointRequest** | [**TicketEndpointRequest**](TicketEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**TicketResponse**](TicketResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketsList

> PaginatedTicketList TicketsList(ctx).XAccountToken(xAccountToken).AccountId(accountId).AssigneeIds(assigneeIds).CollectionIds(collectionIds).CompletedAfter(completedAfter).CompletedBefore(completedBefore).ContactId(contactId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).DueAfter(dueAfter).DueBefore(dueBefore).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).ParentTicketId(parentTicketId).Priority(priority).ProjectId(projectId).RemoteCreatedAfter(remoteCreatedAfter).RemoteCreatedBefore(remoteCreatedBefore).RemoteFields(remoteFields).RemoteId(remoteId).RemoteUpdatedAfter(remoteUpdatedAfter).RemoteUpdatedBefore(remoteUpdatedBefore).ShowEnumOrigins(showEnumOrigins).Status(status).Tags(tags).TicketType(ticketType).Execute()





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
    accountId := "accountId_example" // string | If provided, will only return tickets for this account. (optional)
    assigneeIds := "assigneeIds_example" // string | If provided, will only return tickets assigned to the assignee_ids; multiple assignee_ids can be separated by commas. (optional)
    collectionIds := "collectionIds_example" // string | If provided, will only return tickets assigned to the collection_ids; multiple collection_ids can be separated by commas. (optional)
    completedAfter := time.Now() // time.Time | If provided, will only return tickets completed after this datetime. (optional)
    completedBefore := time.Now() // time.Time | If provided, will only return tickets completed before this datetime. (optional)
    contactId := "contactId_example" // string | If provided, will only return tickets for this contact. (optional)
    createdAfter := time.Now() // time.Time | If provided, will only return objects created after this datetime. (optional)
    createdBefore := time.Now() // time.Time | If provided, will only return objects created before this datetime. (optional)
    cursor := "cD0yMDIxLTAxLTA2KzAzJTNBMjQlM0E1My40MzQzMjYlMkIwMCUzQTAw" // string | The pagination cursor value. (optional)
    dueAfter := time.Now() // time.Time | If provided, will only return tickets due after this datetime. (optional)
    dueBefore := time.Now() // time.Time | If provided, will only return tickets due before this datetime. (optional)
    includeDeletedData := true // bool | Whether to include data that was marked as deleted by third party webhooks. (optional)
    includeRemoteData := true // bool | Whether to include the original data Merge fetched from the third-party to produce these models. (optional)
    modifiedAfter := time.Now() // time.Time | If provided, will only return objects modified after this datetime. (optional)
    modifiedBefore := time.Now() // time.Time | If provided, will only return objects modified before this datetime. (optional)
    pageSize := int32(56) // int32 | Number of results to return per page. (optional)
    parentTicketId := "parentTicketId_example" // string | If provided, will only return sub tickets of the parent_ticket_id. (optional)
    priority := "priority_example" // string | If provided, will only return tickets of this priority. (optional)
    projectId := "projectId_example" // string | If provided, will only return tickets for this project. (optional)
    remoteCreatedAfter := time.Now() // time.Time | If provided, will only return tickets created in the third party platform after this datetime. (optional)
    remoteCreatedBefore := time.Now() // time.Time | If provided, will only return tickets created in the third party platform before this datetime. (optional)
    remoteFields := "priority,status,ticket_type" // string | Deprecated. Use show_enum_origins. (optional)
    remoteId := "remoteId_example" // string | The API provider's ID for the given object. (optional)
    remoteUpdatedAfter := time.Now() // time.Time | If provided, will only return tickets updated in the third party platform after this datetime. (optional)
    remoteUpdatedBefore := time.Now() // time.Time | If provided, will only return tickets updated in the third party platform before this datetime. (optional)
    showEnumOrigins := "priority,status,ticket_type" // string | Which fields should be returned in non-normalized form. (optional)
    status := "status_example" // string | If provided, will only return tickets of this status. (optional)
    tags := "tags_example" // string | If provided, will only return tickets matching the tags; multiple tags can be separated by commas. (optional)
    ticketType := "ticketType_example" // string | If provided, will only return tickets of this type. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TicketsApi.TicketsList(context.Background()).XAccountToken(xAccountToken).AccountId(accountId).AssigneeIds(assigneeIds).CollectionIds(collectionIds).CompletedAfter(completedAfter).CompletedBefore(completedBefore).ContactId(contactId).CreatedAfter(createdAfter).CreatedBefore(createdBefore).Cursor(cursor).DueAfter(dueAfter).DueBefore(dueBefore).IncludeDeletedData(includeDeletedData).IncludeRemoteData(includeRemoteData).ModifiedAfter(modifiedAfter).ModifiedBefore(modifiedBefore).PageSize(pageSize).ParentTicketId(parentTicketId).Priority(priority).ProjectId(projectId).RemoteCreatedAfter(remoteCreatedAfter).RemoteCreatedBefore(remoteCreatedBefore).RemoteFields(remoteFields).RemoteId(remoteId).RemoteUpdatedAfter(remoteUpdatedAfter).RemoteUpdatedBefore(remoteUpdatedBefore).ShowEnumOrigins(showEnumOrigins).Status(status).Tags(tags).TicketType(ticketType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketsApi.TicketsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketsList`: PaginatedTicketList
    fmt.Fprintf(os.Stdout, "Response from `TicketsApi.TicketsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 
 **accountId** | **string** | If provided, will only return tickets for this account. | 
 **assigneeIds** | **string** | If provided, will only return tickets assigned to the assignee_ids; multiple assignee_ids can be separated by commas. | 
 **collectionIds** | **string** | If provided, will only return tickets assigned to the collection_ids; multiple collection_ids can be separated by commas. | 
 **completedAfter** | **time.Time** | If provided, will only return tickets completed after this datetime. | 
 **completedBefore** | **time.Time** | If provided, will only return tickets completed before this datetime. | 
 **contactId** | **string** | If provided, will only return tickets for this contact. | 
 **createdAfter** | **time.Time** | If provided, will only return objects created after this datetime. | 
 **createdBefore** | **time.Time** | If provided, will only return objects created before this datetime. | 
 **cursor** | **string** | The pagination cursor value. | 
 **dueAfter** | **time.Time** | If provided, will only return tickets due after this datetime. | 
 **dueBefore** | **time.Time** | If provided, will only return tickets due before this datetime. | 
 **includeDeletedData** | **bool** | Whether to include data that was marked as deleted by third party webhooks. | 
 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **modifiedAfter** | **time.Time** | If provided, will only return objects modified after this datetime. | 
 **modifiedBefore** | **time.Time** | If provided, will only return objects modified before this datetime. | 
 **pageSize** | **int32** | Number of results to return per page. | 
 **parentTicketId** | **string** | If provided, will only return sub tickets of the parent_ticket_id. | 
 **priority** | **string** | If provided, will only return tickets of this priority. | 
 **projectId** | **string** | If provided, will only return tickets for this project. | 
 **remoteCreatedAfter** | **time.Time** | If provided, will only return tickets created in the third party platform after this datetime. | 
 **remoteCreatedBefore** | **time.Time** | If provided, will only return tickets created in the third party platform before this datetime. | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **remoteId** | **string** | The API provider&#39;s ID for the given object. | 
 **remoteUpdatedAfter** | **time.Time** | If provided, will only return tickets updated in the third party platform after this datetime. | 
 **remoteUpdatedBefore** | **time.Time** | If provided, will only return tickets updated in the third party platform before this datetime. | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 
 **status** | **string** | If provided, will only return tickets of this status. | 
 **tags** | **string** | If provided, will only return tickets matching the tags; multiple tags can be separated by commas. | 
 **ticketType** | **string** | If provided, will only return tickets of this type. | 

### Return type

[**PaginatedTicketList**](PaginatedTicketList.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketsMetaPatchRetrieve

> MetaResponse TicketsMetaPatchRetrieve(ctx, id).XAccountToken(xAccountToken).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TicketsApi.TicketsMetaPatchRetrieve(context.Background(), id).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketsApi.TicketsMetaPatchRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketsMetaPatchRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `TicketsApi.TicketsMetaPatchRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketsMetaPatchRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 


### Return type

[**MetaResponse**](MetaResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketsMetaPostRetrieve

> MetaResponse TicketsMetaPostRetrieve(ctx).XAccountToken(xAccountToken).Execute()





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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TicketsApi.TicketsMetaPostRetrieve(context.Background()).XAccountToken(xAccountToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketsApi.TicketsMetaPostRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketsMetaPostRetrieve`: MetaResponse
    fmt.Fprintf(os.Stdout, "Response from `TicketsApi.TicketsMetaPostRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTicketsMetaPostRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

### Return type

[**MetaResponse**](MetaResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketsPartialUpdate

> TicketResponse TicketsPartialUpdate(ctx, id).XAccountToken(xAccountToken).PatchedTicketEndpointRequest(patchedTicketEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()





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
    patchedTicketEndpointRequest := *openapiclient.NewPatchedTicketEndpointRequest(*openapiclient.NewPatchedTicketRequest()) // PatchedTicketEndpointRequest | 
    isDebugMode := true // bool | Whether to include debug fields (such as log file links) in the response. (optional)
    runAsync := true // bool | Whether or not third-party updates should be run asynchronously. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TicketsApi.TicketsPartialUpdate(context.Background(), id).XAccountToken(xAccountToken).PatchedTicketEndpointRequest(patchedTicketEndpointRequest).IsDebugMode(isDebugMode).RunAsync(runAsync).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketsApi.TicketsPartialUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketsPartialUpdate`: TicketResponse
    fmt.Fprintf(os.Stdout, "Response from `TicketsApi.TicketsPartialUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketsPartialUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **patchedTicketEndpointRequest** | [**PatchedTicketEndpointRequest**](PatchedTicketEndpointRequest.md) |  | 
 **isDebugMode** | **bool** | Whether to include debug fields (such as log file links) in the response. | 
 **runAsync** | **bool** | Whether or not third-party updates should be run asynchronously. | 

### Return type

[**TicketResponse**](TicketResponse.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TicketsRetrieve

> Ticket TicketsRetrieve(ctx, id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()





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
    remoteFields := "priority,status,ticket_type" // string | Deprecated. Use show_enum_origins. (optional)
    showEnumOrigins := "priority,status,ticket_type" // string | Which fields should be returned in non-normalized form. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TicketsApi.TicketsRetrieve(context.Background(), id).XAccountToken(xAccountToken).IncludeRemoteData(includeRemoteData).RemoteFields(remoteFields).ShowEnumOrigins(showEnumOrigins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TicketsApi.TicketsRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TicketsRetrieve`: Ticket
    fmt.Fprintf(os.Stdout, "Response from `TicketsApi.TicketsRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | [**string**](.md) |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiTicketsRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAccountToken** | **string** | Token identifying the end user. | 

 **includeRemoteData** | **bool** | Whether to include the original data Merge fetched from the third-party to produce these models. | 
 **remoteFields** | **string** | Deprecated. Use show_enum_origins. | 
 **showEnumOrigins** | **string** | Which fields should be returned in non-normalized form. | 

### Return type

[**Ticket**](Ticket.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

