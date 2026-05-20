# \RecycleBinConfigApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRecycleBinConfigById**](RecycleBinConfigApi.md#GetRecycleBinConfigById) | **Get** /recycle_bin_config/{id} | Instance Query
[**PatchRecycleBinConfigById**](RecycleBinConfigApi.md#PatchRecycleBinConfigById) | **Patch** /recycle_bin_config/{id} | Modify



## GetRecycleBinConfigById

> RecycleBinConfigInstance GetRecycleBinConfigById(ctx, id).Execute()

Instance Query



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    id := "id_example" // string | Unique identifier of the instance to retrieve (always = \"0\")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecycleBinConfigApi.GetRecycleBinConfigById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecycleBinConfigApi.GetRecycleBinConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecycleBinConfigById`: RecycleBinConfigInstance
    fmt.Fprintf(os.Stdout, "Response from `RecycleBinConfigApi.GetRecycleBinConfigById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the instance to retrieve (always &#x3D; \&quot;0\&quot;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecycleBinConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecycleBinConfigInstance**](RecycleBinConfigInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRecycleBinConfigById

> PatchRecycleBinConfigById(ctx, id).Body(body).Execute()

Modify



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID/clientgen"
)

func main() {
    id := "id_example" // string | Unique identifier of the instance to modify (always = \"0\")
    body := *openapiclient.NewRecycleBinConfigModify() // RecycleBinConfigModify | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RecycleBinConfigApi.PatchRecycleBinConfigById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecycleBinConfigApi.PatchRecycleBinConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the instance to modify (always &#x3D; \&quot;0\&quot;) | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRecycleBinConfigByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RecycleBinConfigModify**](RecycleBinConfigModify.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

