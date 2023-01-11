# \CommonModelScopesApi

All URIs are relative to *https://api.merge.dev/api/ticketing/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CommonModelScopesRetrieve**](CommonModelScopesApi.md#CommonModelScopesRetrieve) | **Get** /common-model-scopes | 
[**CommonModelScopesUpdate**](CommonModelScopesApi.md#CommonModelScopesUpdate) | **Put** /common-model-scopes | 



## CommonModelScopesRetrieve

> CommonModelScopes CommonModelScopesRetrieve(ctx).IntegrationSlug(integrationSlug).LinkedAccountId(linkedAccountId).Execute()





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
    integrationSlug := "integrationSlug_example" // string | Slug of the integration to fetch (optional)
    linkedAccountId := "linkedAccountId_example" // string | ID of specific linked account to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommonModelScopesApi.CommonModelScopesRetrieve(context.Background()).IntegrationSlug(integrationSlug).LinkedAccountId(linkedAccountId).Execute()
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
 **integrationSlug** | **string** | Slug of the integration to fetch | 
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


## CommonModelScopesUpdate

> CommonModelScopes CommonModelScopesUpdate(ctx).CommonModelScopesUpdateSerializer(commonModelScopesUpdateSerializer).IntegrationSlug(integrationSlug).LinkedAccountId(linkedAccountId).Execute()





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
    commonModelScopesUpdateSerializer := *openapiclient.NewCommonModelScopesUpdateSerializer([]openapiclient.CommonModelScopesPutInnerDeserializerRequest{*openapiclient.NewCommonModelScopesPutInnerDeserializerRequest("hris.Employee", []openapiclient.CommonModelScopesPutInnerDeserializerEnabledActionsEnum{openapiclient.CommonModelScopesPutInnerDeserializerEnabledActionsEnum("READ")}, []string{"DisabledFields_example"})}) // CommonModelScopesUpdateSerializer | 
    integrationSlug := "integrationSlug_example" // string | Slug of the integration to fetch (optional)
    linkedAccountId := "linkedAccountId_example" // string | ID of specific linked account to fetch (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CommonModelScopesApi.CommonModelScopesUpdate(context.Background()).CommonModelScopesUpdateSerializer(commonModelScopesUpdateSerializer).IntegrationSlug(integrationSlug).LinkedAccountId(linkedAccountId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CommonModelScopesApi.CommonModelScopesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CommonModelScopesUpdate`: CommonModelScopes
    fmt.Fprintf(os.Stdout, "Response from `CommonModelScopesApi.CommonModelScopesUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCommonModelScopesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **commonModelScopesUpdateSerializer** | [**CommonModelScopesUpdateSerializer**](CommonModelScopesUpdateSerializer.md) |  | 
 **integrationSlug** | **string** | Slug of the integration to fetch | 
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

