# \SecurityConfigApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostSecurityConfigById**](SecurityConfigApi.md#PostSecurityConfigById) | **Post** /security_config/generate_temp_credentials | Generate a temporary credential to access the Powerstore system.



## PostSecurityConfigById

> SecurityConfigGenerateTempCredentialsResponse PostSecurityConfigById(ctx).Execute()

Generate a temporary credential to access the Powerstore system.



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
    resp, r, err := apiClient.SecurityConfigApi.PostSecurityConfigById(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SecurityConfigApi.PostSecurityConfigById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSecurityConfigById`: SecurityConfigGenerateTempCredentialsResponse
    fmt.Fprintf(os.Stdout, "Response from `SecurityConfigApi.PostSecurityConfigById`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostSecurityConfigByIdRequest struct via the builder pattern


### Return type

[**SecurityConfigGenerateTempCredentialsResponse**](SecurityConfigGenerateTempCredentialsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

