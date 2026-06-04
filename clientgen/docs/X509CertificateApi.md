# \X509CertificateApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostX509CertificateById**](X509CertificateApi.md#PostX509CertificateById) | **Post** /x509_certificate/exchange | Exchange Certificates



## PostX509CertificateById

> PostX509CertificateById(ctx).Body(body).Execute()

Exchange Certificates



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
    body := *openapiclient.NewX509CertificateExchange(openapiclient.X509CertificateServiceEnum("CACPIV_HTTP"), "Address_example", int32(123), "Username_example", "Password_example") // X509CertificateExchange | Request body.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.X509CertificateApi.PostX509CertificateById(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `X509CertificateApi.PostX509CertificateById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostX509CertificateByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**X509CertificateExchange**](X509CertificateExchange.md) | Request body. | 

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

