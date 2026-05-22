# \SnapshotRuleApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteSnapshotRuleById**](SnapshotRuleApi.md#DeleteSnapshotRuleById) | **Delete** /snapshot_rule/{id} | Delete
[**GetAllSnapshotRules**](SnapshotRuleApi.md#GetAllSnapshotRules) | **Get** /snapshot_rule | Collection Query
[**GetSnapshotRuleById**](SnapshotRuleApi.md#GetSnapshotRuleById) | **Get** /snapshot_rule/{id} | Instance Query
[**PatchSnapshotRuleById**](SnapshotRuleApi.md#PatchSnapshotRuleById) | **Patch** /snapshot_rule/{id} | Modify
[**PostAllSnapshotRules**](SnapshotRuleApi.md#PostAllSnapshotRules) | **Post** /snapshot_rule | Create



## DeleteSnapshotRuleById

> DeleteSnapshotRuleById(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Unique identifier of the snapshot rule. name:{name} can be used instead of {id}.
    body := *openapiclient.NewSnapshotRuleDelete() // SnapshotRuleDelete |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SnapshotRuleApi.DeleteSnapshotRuleById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotRuleApi.DeleteSnapshotRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the snapshot rule. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSnapshotRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SnapshotRuleDelete**](SnapshotRuleDelete.md) |  | 

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


## GetAllSnapshotRules

> []SnapshotRuleInstance GetAllSnapshotRules(ctx).Execute()

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
    resp, r, err := apiClient.SnapshotRuleApi.GetAllSnapshotRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotRuleApi.GetAllSnapshotRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllSnapshotRules`: []SnapshotRuleInstance
    fmt.Fprintf(os.Stdout, "Response from `SnapshotRuleApi.GetAllSnapshotRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSnapshotRulesRequest struct via the builder pattern


### Return type

[**[]SnapshotRuleInstance**](SnapshotRuleInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshotRuleById

> SnapshotRuleInstance GetSnapshotRuleById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the snapshot rule. name:{name} can be used instead of {id}.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotRuleApi.GetSnapshotRuleById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotRuleApi.GetSnapshotRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSnapshotRuleById`: SnapshotRuleInstance
    fmt.Fprintf(os.Stdout, "Response from `SnapshotRuleApi.GetSnapshotRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the snapshot rule. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SnapshotRuleInstance**](SnapshotRuleInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchSnapshotRuleById

> PatchSnapshotRuleById(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Unique identifier of the snapshot rule. name:{name} can be used instead of {id}.
    body := *openapiclient.NewSnapshotRuleModify() // SnapshotRuleModify |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SnapshotRuleApi.PatchSnapshotRuleById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotRuleApi.PatchSnapshotRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the snapshot rule. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchSnapshotRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SnapshotRuleModify**](SnapshotRuleModify.md) |  | 

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


## PostAllSnapshotRules

> CreateResponse PostAllSnapshotRules(ctx).Body(body).Execute()

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
    body := *openapiclient.NewSnapshotRuleCreate("Name_example", int32(123)) // SnapshotRuleCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SnapshotRuleApi.PostAllSnapshotRules(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SnapshotRuleApi.PostAllSnapshotRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAllSnapshotRules`: CreateResponse
    fmt.Fprintf(os.Stdout, "Response from `SnapshotRuleApi.PostAllSnapshotRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAllSnapshotRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SnapshotRuleCreate**](SnapshotRuleCreate.md) |  | 

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

