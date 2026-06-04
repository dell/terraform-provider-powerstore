# \IoLimitRuleApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIoLimitRuleById**](IoLimitRuleApi.md#DeleteIoLimitRuleById) | **Delete** /io_limit_rule/{id} | Delete
[**GetAllIoLimitRules**](IoLimitRuleApi.md#GetAllIoLimitRules) | **Get** /io_limit_rule | Collection Query
[**GetIoLimitRuleById**](IoLimitRuleApi.md#GetIoLimitRuleById) | **Get** /io_limit_rule/{id} | Instance Query
[**PatchIoLimitRuleById**](IoLimitRuleApi.md#PatchIoLimitRuleById) | **Patch** /io_limit_rule/{id} | Modify
[**PostAllIoLimitRules**](IoLimitRuleApi.md#PostAllIoLimitRules) | **Post** /io_limit_rule | Create



## DeleteIoLimitRuleById

> DeleteIoLimitRuleById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the io_limit_rule. name:{name} can be used instead of {id}.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IoLimitRuleApi.DeleteIoLimitRuleById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IoLimitRuleApi.DeleteIoLimitRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the io_limit_rule. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIoLimitRuleByIdRequest struct via the builder pattern


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


## GetAllIoLimitRules

> []IoLimitRuleInstance GetAllIoLimitRules(ctx).Execute()

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
    resp, r, err := apiClient.IoLimitRuleApi.GetAllIoLimitRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IoLimitRuleApi.GetAllIoLimitRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllIoLimitRules`: []IoLimitRuleInstance
    fmt.Fprintf(os.Stdout, "Response from `IoLimitRuleApi.GetAllIoLimitRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIoLimitRulesRequest struct via the builder pattern


### Return type

[**[]IoLimitRuleInstance**](IoLimitRuleInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIoLimitRuleById

> IoLimitRuleInstance GetIoLimitRuleById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the io_limit_rule. name:{name} can be used instead of {id}.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IoLimitRuleApi.GetIoLimitRuleById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IoLimitRuleApi.GetIoLimitRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIoLimitRuleById`: IoLimitRuleInstance
    fmt.Fprintf(os.Stdout, "Response from `IoLimitRuleApi.GetIoLimitRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the io_limit_rule. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIoLimitRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IoLimitRuleInstance**](IoLimitRuleInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchIoLimitRuleById

> PatchIoLimitRuleById(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Unique identifier of the io_limit_rule. name:{name} can be used instead of {id}.
    body := *openapiclient.NewIoLimitRuleModify() // IoLimitRuleModify | Modify request arguments.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.IoLimitRuleApi.PatchIoLimitRuleById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IoLimitRuleApi.PatchIoLimitRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the io_limit_rule. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchIoLimitRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IoLimitRuleModify**](IoLimitRuleModify.md) | Modify request arguments. | 

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


## PostAllIoLimitRules

> CreateResponse PostAllIoLimitRules(ctx).Body(body).Execute()

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
    body := *openapiclient.NewIoLimitRuleCreate("Name_example", openapiclient.BandwidthLimitTypeEnum("Absolute")) // IoLimitRuleCreate | Create request arguments.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.IoLimitRuleApi.PostAllIoLimitRules(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IoLimitRuleApi.PostAllIoLimitRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAllIoLimitRules`: CreateResponse
    fmt.Fprintf(os.Stdout, "Response from `IoLimitRuleApi.PostAllIoLimitRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAllIoLimitRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IoLimitRuleCreate**](IoLimitRuleCreate.md) | Create request arguments. | 

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

