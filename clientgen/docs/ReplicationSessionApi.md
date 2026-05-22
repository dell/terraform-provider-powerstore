# \ReplicationSessionApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllReplicationSessions**](ReplicationSessionApi.md#GetAllReplicationSessions) | **Get** /replication_session | Collection Query
[**GetReplicationSessionById**](ReplicationSessionApi.md#GetReplicationSessionById) | **Get** /replication_session/{id} | Instance Query
[**PatchReplicationSessionById**](ReplicationSessionApi.md#PatchReplicationSessionById) | **Patch** /replication_session/{id} | Modify
[**ReplicationSessionFailover**](ReplicationSessionApi.md#ReplicationSessionFailover) | **Post** /replication_session/{id}/failover | Failover
[**ReplicationSessionPause**](ReplicationSessionApi.md#ReplicationSessionPause) | **Post** /replication_session/{id}/pause | Pause
[**ReplicationSessionReprotect**](ReplicationSessionApi.md#ReplicationSessionReprotect) | **Post** /replication_session/{id}/reprotect | Reprotect
[**ReplicationSessionResume**](ReplicationSessionApi.md#ReplicationSessionResume) | **Post** /replication_session/{id}/resume | Resume
[**ReplicationSessionStartFailoverTest**](ReplicationSessionApi.md#ReplicationSessionStartFailoverTest) | **Post** /replication_session/{id}/start_failover_test | Start DR Failover Simulation Test
[**ReplicationSessionStopFailoverTest**](ReplicationSessionApi.md#ReplicationSessionStopFailoverTest) | **Post** /replication_session/{id}/stop_failover_test | Stop DR Failover Simulation Test
[**ReplicationSessionSync**](ReplicationSessionApi.md#ReplicationSessionSync) | **Post** /replication_session/{id}/sync | Synchronize



## GetAllReplicationSessions

> []ReplicationSessionInstance GetAllReplicationSessions(ctx).Execute()

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
    resp, r, err := apiClient.ReplicationSessionApi.GetAllReplicationSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.GetAllReplicationSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllReplicationSessions`: []ReplicationSessionInstance
    fmt.Fprintf(os.Stdout, "Response from `ReplicationSessionApi.GetAllReplicationSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllReplicationSessionsRequest struct via the builder pattern


### Return type

[**[]ReplicationSessionInstance**](ReplicationSessionInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReplicationSessionById

> ReplicationSessionInstance GetReplicationSessionById(ctx, id).Execute()

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
    id := "id_example" // string | Unique identifier of the replication session. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationSessionApi.GetReplicationSessionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.GetReplicationSessionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReplicationSessionById`: ReplicationSessionInstance
    fmt.Fprintf(os.Stdout, "Response from `ReplicationSessionApi.GetReplicationSessionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the replication session.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReplicationSessionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReplicationSessionInstance**](ReplicationSessionInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PatchReplicationSessionById

> PatchReplicationSessionById(ctx, id).Body(body).Execute()

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
    id := "id_example" // string | Unique identifier of the replication session. 
    body := *openapiclient.NewReplicationSessionModify() // ReplicationSessionModify | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReplicationSessionApi.PatchReplicationSessionById(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.PatchReplicationSessionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the replication session.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPatchReplicationSessionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReplicationSessionModify**](ReplicationSessionModify.md) |  | 

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


## ReplicationSessionFailover

> ReplicationSessionFailover(ctx, id).Body(body).Execute()

Failover



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
    id := "id_example" // string | Unique identifier of the replication session. 
    body := *openapiclient.NewReplicationSessionFailover() // ReplicationSessionFailover |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReplicationSessionApi.ReplicationSessionFailover(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.ReplicationSessionFailover``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the replication session.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplicationSessionFailoverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReplicationSessionFailover**](ReplicationSessionFailover.md) |  | 

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


## ReplicationSessionPause

> ReplicationSessionPause(ctx, id).Execute()

Pause



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
    id := "id_example" // string | Unique identifier of the replication session. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReplicationSessionApi.ReplicationSessionPause(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.ReplicationSessionPause``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the replication session.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplicationSessionPauseRequest struct via the builder pattern


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


## ReplicationSessionReprotect

> ReplicationSessionReprotect(ctx, id).Body(body).Execute()

Reprotect



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
    id := "id_example" // string | Unique identifier of the replication session. 
    body := *openapiclient.NewReplicationSessionReprotect() // ReplicationSessionReprotect | Parameters to reprotect a replication session. Was added in version 4.0.0.0. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReplicationSessionApi.ReplicationSessionReprotect(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.ReplicationSessionReprotect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the replication session.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplicationSessionReprotectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReplicationSessionReprotect**](ReplicationSessionReprotect.md) | Parameters to reprotect a replication session. Was added in version 4.0.0.0. | 

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


## ReplicationSessionResume

> ReplicationSessionResume(ctx, id).Execute()

Resume



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
    id := "id_example" // string | Unique identifier of the replication session. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReplicationSessionApi.ReplicationSessionResume(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.ReplicationSessionResume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the replication session.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplicationSessionResumeRequest struct via the builder pattern


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


## ReplicationSessionStartFailoverTest

> ReplicationSessionStartFailoverTest(ctx, id).Body(body).Execute()

Start DR Failover Simulation Test



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
    id := "id_example" // string | Unique identifier of the replication session.  Was added in version 2.0.0.0.
    body := *openapiclient.NewReplicationStartFailoverTest() // ReplicationStartFailoverTest | Parameters to start a DR failover simulation test on a replication session. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReplicationSessionApi.ReplicationSessionStartFailoverTest(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.ReplicationSessionStartFailoverTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the replication session.  Was added in version 2.0.0.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplicationSessionStartFailoverTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReplicationStartFailoverTest**](ReplicationStartFailoverTest.md) | Parameters to start a DR failover simulation test on a replication session. | 

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


## ReplicationSessionStopFailoverTest

> ReplicationSessionStopFailoverTestResponse ReplicationSessionStopFailoverTest(ctx, id).Body(body).Execute()

Stop DR Failover Simulation Test



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
    id := "id_example" // string | Unique identifier of the replication session.  Was added in version 2.0.0.0.
    body := *openapiclient.NewReplicationStopFailoverTest() // ReplicationStopFailoverTest | Parameters to stop a DR failover simulation test on a replication session. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReplicationSessionApi.ReplicationSessionStopFailoverTest(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.ReplicationSessionStopFailoverTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReplicationSessionStopFailoverTest`: ReplicationSessionStopFailoverTestResponse
    fmt.Fprintf(os.Stdout, "Response from `ReplicationSessionApi.ReplicationSessionStopFailoverTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the replication session.  Was added in version 2.0.0.0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplicationSessionStopFailoverTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ReplicationStopFailoverTest**](ReplicationStopFailoverTest.md) | Parameters to stop a DR failover simulation test on a replication session. | 

### Return type

[**ReplicationSessionStopFailoverTestResponse**](ReplicationSessionStopFailoverTestResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplicationSessionSync

> ReplicationSessionSync(ctx, id).Execute()

Synchronize



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
    id := "id_example" // string | Unique identifier of the replication session. 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.ReplicationSessionApi.ReplicationSessionSync(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReplicationSessionApi.ReplicationSessionSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Unique identifier of the replication session.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplicationSessionSyncRequest struct via the builder pattern


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

