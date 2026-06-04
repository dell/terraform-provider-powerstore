# \FileIoLimitRuleApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFileIoLimitRuleById**](FileIoLimitRuleApi.md#DeleteFileIoLimitRuleById) | **Delete** /file_io_limit_rule/{id} | Delete
[**GetAllFileIoLimitRules**](FileIoLimitRuleApi.md#GetAllFileIoLimitRules) | **Get** /file_io_limit_rule | Collection Query
[**GetFileIoLimitRuleById**](FileIoLimitRuleApi.md#GetFileIoLimitRuleById) | **Get** /file_io_limit_rule/{id} | Instance Query
[**PatchFileIoLimitRuleById**](FileIoLimitRuleApi.md#PatchFileIoLimitRuleById) | **Patch** /file_io_limit_rule/{id} | Modify
[**PostAllFileIoLimitRules**](FileIoLimitRuleApi.md#PostAllFileIoLimitRules) | **Post** /file_io_limit_rule | Create



## DeleteFileIoLimitRuleById

> DeleteFileIoLimitRuleById(ctx, id).Execute()

Delete



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
    id := "id_example" // string | Unique identifier of the file_io_limit_rule instance. name:{name} can be used instead of {id}. Was added in version 4.1.0.0.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FileIoLimitRuleApi.DeleteFileIoLimitRuleById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileIoLimitRuleApi.DeleteFileIoLimitRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the file_io_limit_rule instance. name:{name} can be used instead of {id}. Was added in version 4.1.0.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileIoLimitRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFileIoLimitRules

> []FileIoLimitRuleInstance GetAllFileIoLimitRules(ctx).Execute()

Collection Query



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileIoLimitRuleApi.GetAllFileIoLimitRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileIoLimitRuleApi.GetAllFileIoLimitRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllFileIoLimitRules`: []FileIoLimitRuleInstance
    fmt.Fprintf(os.Stdout, "Response from `FileIoLimitRuleApi.GetAllFileIoLimitRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFileIoLimitRulesRequest struct via the builder pattern


### Return type

[**[]FileIoLimitRuleInstance**](FileIoLimitRuleInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileIoLimitRuleById

> FileIoLimitRuleInstance GetFileIoLimitRuleById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the file_io_limit_rule instance. name:{name} can be used instead of {id}. Was added in version 4.1.0.0.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileIoLimitRuleApi.GetFileIoLimitRuleById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileIoLimitRuleApi.GetFileIoLimitRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFileIoLimitRuleById`: FileIoLimitRuleInstance
    fmt.Fprintf(os.Stdout, "Response from `FileIoLimitRuleApi.GetFileIoLimitRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the file_io_limit_rule instance. name:{name} can be used instead of {id}. Was added in version 4.1.0.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileIoLimitRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FileIoLimitRuleInstance**](FileIoLimitRuleInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchFileIoLimitRuleById

> PatchFileIoLimitRuleById(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Unique identifier of the file_io_limit_rule instance. name:{name} can be used instead of {id}. Was added in version 4.1.0.0.
    body := *openapiclient.NewFileIoLimitRuleModify() // FileIoLimitRuleModify | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.FileIoLimitRuleApi.PatchFileIoLimitRuleById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileIoLimitRuleApi.PatchFileIoLimitRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the file_io_limit_rule instance. name:{name} can be used instead of {id}. Was added in version 4.1.0.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchFileIoLimitRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FileIoLimitRuleModify**](FileIoLimitRuleModify.md) |  | 

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


## PostAllFileIoLimitRules

> CreateResponse PostAllFileIoLimitRules(ctx).Body(body).Execute()

Create



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
    body := *openapiclient.NewFileIoLimitRuleCreate("Name_example", int32(123)) // FileIoLimitRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FileIoLimitRuleApi.PostAllFileIoLimitRules(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FileIoLimitRuleApi.PostAllFileIoLimitRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAllFileIoLimitRules`: CreateResponse
    fmt.Fprintf(os.Stdout, "Response from `FileIoLimitRuleApi.PostAllFileIoLimitRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAllFileIoLimitRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FileIoLimitRuleCreate**](FileIoLimitRuleCreate.md) |  | 

### Return type

[**CreateResponse**](CreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

