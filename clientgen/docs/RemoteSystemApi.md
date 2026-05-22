# \RemoteSystemApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRemoteSystemById**](RemoteSystemApi.md#DeleteRemoteSystemById) | **Delete** /remote_system/{id} | Delete
[**GetAllRemoteSystems**](RemoteSystemApi.md#GetAllRemoteSystems) | **Get** /remote_system | Collection Query
[**GetRemoteSystemById**](RemoteSystemApi.md#GetRemoteSystemById) | **Get** /remote_system/{id} | Instance Query
[**PatchRemoteSystemById**](RemoteSystemApi.md#PatchRemoteSystemById) | **Patch** /remote_system/{id} | Modify
[**PostAllRemoteSystems**](RemoteSystemApi.md#PostAllRemoteSystems) | **Post** /remote_system | Create



## DeleteRemoteSystemById

> DeleteRemoteSystemById(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Unique identifier of the remote system.  name:{name} can be used instead of {id}.
    body := map[string]interface{}{ ... } // map[string]interface{} | Parameters to delete a remote system.  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RemoteSystemApi.DeleteRemoteSystemById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteSystemApi.DeleteRemoteSystemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the remote system.  name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteSystemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** | Parameters to delete a remote system.  | 

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


## GetAllRemoteSystems

> []RemoteSystemInstance GetAllRemoteSystems(ctx).Execute()

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
    resp, r, err := apiClient.RemoteSystemApi.GetAllRemoteSystems(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteSystemApi.GetAllRemoteSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllRemoteSystems`: []RemoteSystemInstance
    fmt.Fprintf(os.Stdout, "Response from `RemoteSystemApi.GetAllRemoteSystems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllRemoteSystemsRequest struct via the builder pattern


### Return type

[**[]RemoteSystemInstance**](RemoteSystemInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteSystemById

> RemoteSystemInstance GetRemoteSystemById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the remote system.  name:{name} can be used instead of {id}.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteSystemApi.GetRemoteSystemById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteSystemApi.GetRemoteSystemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRemoteSystemById`: RemoteSystemInstance
    fmt.Fprintf(os.Stdout, "Response from `RemoteSystemApi.GetRemoteSystemById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the remote system.  name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteSystemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RemoteSystemInstance**](RemoteSystemInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchRemoteSystemById

> PatchRemoteSystemById(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Unique identifier of the remote system.  name:{name} can be used instead of {id}.
    body := *openapiclient.NewRemoteSystemModify() // RemoteSystemModify | Parameters to modify the remote system. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RemoteSystemApi.PatchRemoteSystemById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteSystemApi.PatchRemoteSystemById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the remote system.  name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchRemoteSystemByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RemoteSystemModify**](RemoteSystemModify.md) | Parameters to modify the remote system.  | 

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


## PostAllRemoteSystems

> CreateResponse PostAllRemoteSystems(ctx).Body(body).Execute()

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
    body := *openapiclient.NewRemoteSystemCreate() // RemoteSystemCreate | Parameters to create a remote system. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RemoteSystemApi.PostAllRemoteSystems(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RemoteSystemApi.PostAllRemoteSystems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAllRemoteSystems`: CreateResponse
    fmt.Fprintf(os.Stdout, "Response from `RemoteSystemApi.PostAllRemoteSystems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAllRemoteSystemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RemoteSystemCreate**](RemoteSystemCreate.md) | Parameters to create a remote system.  | 

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

