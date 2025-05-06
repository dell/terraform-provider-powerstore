# \LoginSessionApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllLoginSessions**](LoginSessionApi.md#GetAllLoginSessions) | **Get** /login_session | Collection Query



## GetAllLoginSessions

> []LoginSessionInstance GetAllLoginSessions(ctx).Execute()

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
    resp, r, err := apiClient.LoginSessionApi.GetAllLoginSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LoginSessionApi.GetAllLoginSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllLoginSessions`: []LoginSessionInstance
    fmt.Fprintf(os.Stdout, "Response from `LoginSessionApi.GetAllLoginSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllLoginSessionsRequest struct via the builder pattern


### Return type

[**[]LoginSessionInstance**](LoginSessionInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

