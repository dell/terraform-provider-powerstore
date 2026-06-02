# \ClusterApi

All URIs are relative to */api/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllClusters**](ClusterApi.md#GetAllClusters) | **Get** /cluster | Collection Query
[**PostAllClusters**](ClusterApi.md#PostAllClusters) | **Post** /cluster | Create



## GetAllClusters

> []ClusterInstance GetAllClusters(ctx).Execute()

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
    resp, r, err := apiClient.ClusterApi.GetAllClusters(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.GetAllClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllClusters`: []ClusterInstance
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.GetAllClusters`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllClustersRequest struct via the builder pattern


### Return type

[**[]ClusterInstance**](ClusterInstance.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostAllClusters

> CreateResponse PostAllClusters(ctx).Body(body).Execute()

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
    body := *openapiclient.NewClusterCreate(*openapiclient.NewClusterCreateCluster("Cluster-007"), []openapiclient.ClusterCreateAppliances{*openapiclient.NewClusterCreateAppliances("11.22.33.44")}, []string{"DnsServers_example"}, []string{"NtpServers_example"}, []openapiclient.ClusterCreateNetworks{*openapiclient.NewClusterCreateNetworks(openapiclient.NetworkTypeEnum("Management"), int32(64), []string{"Addresses_example"})}) // ClusterCreate | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ClusterApi.PostAllClusters(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClusterApi.PostAllClusters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostAllClusters`: CreateResponse
    fmt.Fprintf(os.Stdout, "Response from `ClusterApi.PostAllClusters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostAllClustersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ClusterCreate**](ClusterCreate.md) |  | 

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

