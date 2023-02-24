# \CommonModelScopesApi

All URIs are relative to *https://api.merge.dev/api/ticketing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommonModelScopesCreate**](CommonModelScopesApi.md#CommonModelScopesCreate) | **Post** /common-model-scopes | 
[**CommonModelScopesRetrieve**](CommonModelScopesApi.md#CommonModelScopesRetrieve) | **Get** /common-model-scopes | 



## CommonModelScopesCreate

> CommonModelScopes CommonModelScopesCreate(ctx).CommonModelScopesUpdateSerializer(commonModelScopesUpdateSerializer).LinkedAccountId(linkedAccountId).Execute()





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
    commonModelScopesUpdateSerializer := *openapiclient.NewCommonModelScopesUpdateSerializer([]openapiclient.CommonModelScopesPostInnerDeserializerRequest{*openapiclient.NewCommonModelScopesPostInnerDeserializerRequest("hris.Employee", []openapiclient.EnabledActionsEnum{openapiclient.EnabledActionsEnum("READ")}, []string{"DisabledFields_example"})}) // CommonModelScopesUpdateSerializer | 
    linkedAccountId := "linkedAccountId_example" // string | ID of specific linked account to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommonModelScopesApi.CommonModelScopesCreate(context.Background()).CommonModelScopesUpdateSerializer(commonModelScopesUpdateSerializer).LinkedAccountId(linkedAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonModelScopesApi.CommonModelScopesCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommonModelScopesCreate`: CommonModelScopes
    fmt.Fprintf(os.Stdout, "Response from `CommonModelScopesApi.CommonModelScopesCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommonModelScopesCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonModelScopesUpdateSerializer** | [**CommonModelScopesUpdateSerializer**](CommonModelScopesUpdateSerializer.md) |  | 
 **linkedAccountId** | **string** | ID of specific linked account to fetch | 

### Return type

[**CommonModelScopes**](CommonModelScopes.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: application/json, application/x-www-form-urlencoded, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CommonModelScopesRetrieve

> CommonModelScopes CommonModelScopesRetrieve(ctx).LinkedAccountId(linkedAccountId).Execute()





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
    linkedAccountId := "linkedAccountId_example" // string | ID of specific linked account to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommonModelScopesApi.CommonModelScopesRetrieve(context.Background()).LinkedAccountId(linkedAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonModelScopesApi.CommonModelScopesRetrieve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommonModelScopesRetrieve`: CommonModelScopes
    fmt.Fprintf(os.Stdout, "Response from `CommonModelScopesApi.CommonModelScopesRetrieve`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommonModelScopesRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkedAccountId** | **string** | ID of specific linked account to fetch | 

### Return type

[**CommonModelScopes**](CommonModelScopes.md)

### Authorization

[tokenAuth](../README.md#tokenAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

