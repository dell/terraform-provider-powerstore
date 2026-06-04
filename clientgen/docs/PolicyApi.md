# \PolicyApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePolicyById**](PolicyApi.md#DeletePolicyById) | **Delete** /policy/{id} | Delete
[**GetAllPolicys**](PolicyApi.md#GetAllPolicys) | **Get** /policy | Collection Query
[**GetPolicyById**](PolicyApi.md#GetPolicyById) | **Get** /policy/{id} | Instance Query
[**PatchPolicyById**](PolicyApi.md#PatchPolicyById) | **Patch** /policy/{id} | Modify
[**PostAllPolicys**](PolicyApi.md#PostAllPolicys) | **Post** /policy | Create



## DeletePolicyById

> DeletePolicyById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the protectionor QoS policy to be deleted. name:{name} can be used instead of {id}.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyApi.DeletePolicyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.DeletePolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the protectionor QoS policy to be deleted. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyByIdRequest struct via the builder pattern


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


## GetAllPolicys

> []PolicyInstance GetAllPolicys(ctx).Execute()

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
    resp, r, err := apiClient.PolicyApi.GetAllPolicys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetAllPolicys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllPolicys`: []PolicyInstance
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetAllPolicys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPolicysRequest struct via the builder pattern


### Return type

[**[]PolicyInstance**](PolicyInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicyById

> PolicyInstance GetPolicyById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the policy. name:{name} can be used instead of {id}.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.GetPolicyById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.GetPolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicyById`: PolicyInstance
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.GetPolicyById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the policy. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PolicyInstance**](PolicyInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchPolicyById

> PatchPolicyById(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Unique identifier of the protectionor QoS policy to be modified. name:{name} can be used instead of {id}.
    body := *openapiclient.NewPolicyModify() // PolicyModify |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.PolicyApi.PatchPolicyById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.PatchPolicyById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the protectionor QoS policy to be modified. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchPolicyByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**PolicyModify**](PolicyModify.md) |  | 

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


## PostAllPolicys

> CreateResponse PostAllPolicys(ctx).Body(body).Execute()

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
    body := *openapiclient.NewPolicyCreate("Name_example") // PolicyCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyApi.PostAllPolicys(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyApi.PostAllPolicys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAllPolicys`: CreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyApi.PostAllPolicys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAllPolicysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PolicyCreate**](PolicyCreate.md) |  | 

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

