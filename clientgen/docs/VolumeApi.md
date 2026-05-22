# \VolumeApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumeConfigureMetro**](VolumeApi.md#VolumeConfigureMetro) | **Post** /volume/{id}/configure_metro | Configure Metro
[**VolumeEndMetro**](VolumeApi.md#VolumeEndMetro) | **Post** /volume/{id}/end_metro | End Metro Configuration



## VolumeConfigureMetro

> VolumeConfigureMetroResponse VolumeConfigureMetro(ctx, id).Body(body).Execute()

Configure Metro



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
    id := "id_example" // string | Unique identifier of volume to configure. name:{name} can be used instead of {id}.
    body := *openapiclient.NewVolumeConfigureMetro("RemoteSystemId_example") // VolumeConfigureMetro | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.VolumeApi.VolumeConfigureMetro(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumeApi.VolumeConfigureMetro``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VolumeConfigureMetro`: VolumeConfigureMetroResponse
    fmt.Fprintf(os.Stdout, "Response from `VolumeApi.VolumeConfigureMetro`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of volume to configure. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeConfigureMetroRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeConfigureMetro**](VolumeConfigureMetro.md) |  | 

### Return type

[**VolumeConfigureMetroResponse**](VolumeConfigureMetroResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VolumeEndMetro

> VolumeEndMetro(ctx, id).Body(body).Execute()

End Metro Configuration



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
    id := "id_example" // string | Unique identifier of volume for which to end the metro configuration. name:{name} can be used instead of {id}.
    body := *openapiclient.NewVolumeEndMetro() // VolumeEndMetro |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.VolumeApi.VolumeEndMetro(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `VolumeApi.VolumeEndMetro``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of volume for which to end the metro configuration. name:{name} can be used instead of {id}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVolumeEndMetroRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**VolumeEndMetro**](VolumeEndMetro.md) |  | 

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

