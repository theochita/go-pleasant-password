# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiV6RestFoldersIdDelete**](DefaultAPI.md#ApiV6RestFoldersIdDelete) | **Delete** /api/v6/rest/folders/{id} | 
[**GetV6CredentialsByID**](DefaultAPI.md#GetV6CredentialsByID) | **Get** /api/v6/rest/entries/{id} | get_entry
[**GetV6Folders**](DefaultAPI.md#GetV6Folders) | **Get** /api/v6/rest/folders | 
[**GetV6FoldersByID**](DefaultAPI.md#GetV6FoldersByID) | **Get** /api/v6/rest/folders/{id} | 
[**GetV6ServerInfo**](DefaultAPI.md#GetV6ServerInfo) | **Get** /api/v6/rest/getserverinfo | 
[**PostV6Folders**](DefaultAPI.md#PostV6Folders) | **Post** /api/v6/rest/folders | Create New Folder
[**PostV6Search**](DefaultAPI.md#PostV6Search) | **Post** /api/v6/rest/search | 



## ApiV6RestFoldersIdDelete

> ApiV6RestFoldersIdDelete(ctx, id).DeleteAction(deleteAction).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/theochita/go-pleasant-password"
)

func main() {
    id := "id_example" // string | 
    deleteAction := *openapiclient.NewDeleteAction() // DeleteAction | Delete or archive  folder

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.ApiV6RestFoldersIdDelete(context.Background(), id).DeleteAction(deleteAction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.ApiV6RestFoldersIdDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApiV6RestFoldersIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteAction** | [**DeleteAction**](DeleteAction.md) | Delete or archive  folder | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV6CredentialsByID

> V6CredentialResult GetV6CredentialsByID(ctx, id).Execute()

get_entry

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/theochita/go-pleasant-password"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetV6CredentialsByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetV6CredentialsByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetV6CredentialsByID`: V6CredentialResult
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetV6CredentialsByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV6CredentialsByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V6CredentialResult**](V6CredentialResult.md)

### Authorization

[Bearer_token](../README.md#Bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV6Folders

> V6CredentialGroupOutput GetV6Folders(ctx).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/theochita/go-pleasant-password"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetV6Folders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetV6Folders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetV6Folders`: V6CredentialGroupOutput
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetV6Folders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV6FoldersRequest struct via the builder pattern


### Return type

[**V6CredentialGroupOutput**](V6CredentialGroupOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV6FoldersByID

> V6CredentialGroupOutput GetV6FoldersByID(ctx, id).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/theochita/go-pleasant-password"
)

func main() {
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetV6FoldersByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetV6FoldersByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetV6FoldersByID`: V6CredentialGroupOutput
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetV6FoldersByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV6FoldersByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**V6CredentialGroupOutput**](V6CredentialGroupOutput.md)

### Authorization

[Bearer_token](../README.md#Bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV6ServerInfo

> V6ServerInfoOutput GetV6ServerInfo(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/theochita/go-pleasant-password"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.GetV6ServerInfo(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetV6ServerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetV6ServerInfo`: V6ServerInfoOutput
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetV6ServerInfo`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV6ServerInfoRequest struct via the builder pattern


### Return type

[**V6ServerInfoOutput**](V6ServerInfoOutput.md)

### Authorization

[Bearer_token](../README.md#Bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV6Folders

> PostV6Folders(ctx).V6CredentialGroupOutput(v6CredentialGroupOutput).Execute()

Create New Folder

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/theochita/go-pleasant-password"
)

func main() {
    v6CredentialGroupOutput := *openapiclient.NewV6CredentialGroupOutput() // V6CredentialGroupOutput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.PostV6Folders(context.Background()).V6CredentialGroupOutput(v6CredentialGroupOutput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostV6Folders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV6FoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v6CredentialGroupOutput** | [**V6CredentialGroupOutput**](V6CredentialGroupOutput.md) |  | 

### Return type

 (empty response body)

### Authorization

[Bearer_token](../README.md#Bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV6Search

> V6SearchOutput PostV6Search(ctx).V6SearchInput(v6SearchInput).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/theochita/go-pleasant-password"
)

func main() {
    v6SearchInput := *openapiclient.NewV6SearchInput() // V6SearchInput | Searches

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.PostV6Search(context.Background()).V6SearchInput(v6SearchInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostV6Search``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostV6Search`: V6SearchOutput
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostV6Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV6SearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v6SearchInput** | [**V6SearchInput**](V6SearchInput.md) | Searches | 

### Return type

[**V6SearchOutput**](V6SearchOutput.md)

### Authorization

[Bearer_token](../README.md#Bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

