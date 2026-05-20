# \RecycleBinApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRecycleBinById**](RecycleBinApi.md#DeleteRecycleBinById) | **Delete** /recycle_bin/{id} | Delete
[**GetAllRecycleBins**](RecycleBinApi.md#GetAllRecycleBins) | **Get** /recycle_bin | Collection Query
[**GetRecycleBinById**](RecycleBinApi.md#GetRecycleBinById) | **Get** /recycle_bin/{id} | Instance Query
[**PostRecycleBinById**](RecycleBinApi.md#PostRecycleBinById) | **Post** /recycle_bin/empty | Empty
[**RecycleBinRecover**](RecycleBinApi.md#RecycleBinRecover) | **Post** /recycle_bin/{id}/recover | Recover



## DeleteRecycleBinById

> DeleteRecycleBinById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the storage object to delete. name:{name} can be used instead of {id}.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RecycleBinApi.DeleteRecycleBinById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecycleBinApi.DeleteRecycleBinById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the storage object to delete. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRecycleBinByIdRequest struct via the builder pattern


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


## GetAllRecycleBins

> []RecycleBinInstance GetAllRecycleBins(ctx).Execute()

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
    resp, r, err := apiClient.RecycleBinApi.GetAllRecycleBins(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecycleBinApi.GetAllRecycleBins``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRecycleBins`: []RecycleBinInstance
    fmt.Fprintf(os.Stdout, "Response from `RecycleBinApi.GetAllRecycleBins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRecycleBinsRequest struct via the builder pattern


### Return type

[**[]RecycleBinInstance**](RecycleBinInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRecycleBinById

> RecycleBinInstance GetRecycleBinById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the instance to retrieve. name:{name} can be used instead of {id}.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RecycleBinApi.GetRecycleBinById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecycleBinApi.GetRecycleBinById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRecycleBinById`: RecycleBinInstance
    fmt.Fprintf(os.Stdout, "Response from `RecycleBinApi.GetRecycleBinById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the instance to retrieve. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRecycleBinByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RecycleBinInstance**](RecycleBinInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostRecycleBinById

> PostRecycleBinById(ctx).Execute()

Empty



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
    r, err := apiClient.RecycleBinApi.PostRecycleBinById(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecycleBinApi.PostRecycleBinById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostRecycleBinByIdRequest struct via the builder pattern


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


## RecycleBinRecover

> RecycleBinRecover(ctx, id).Body(body).Execute()

Recover



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
    id := "id_example" // string | Unique identifier of the instance to recover. name:{name} can be used instead of {id}.
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RecycleBinApi.RecycleBinRecover(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RecycleBinApi.RecycleBinRecover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the instance to recover. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecycleBinRecoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

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

