# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV6CredentialsByID**](DefaultAPI.md#DeleteV6CredentialsByID) | **Delete** /api/v6/rest/entries/{id} | 
[**DeleteV6FoldersByID**](DefaultAPI.md#DeleteV6FoldersByID) | **Delete** /api/v6/rest/folders/{id} | 
[**GetV6CredentialPasswordByID**](DefaultAPI.md#GetV6CredentialPasswordByID) | **Get** /api/v6/rest/entries/{id}/password | get_entry
[**GetV6CredentialsByID**](DefaultAPI.md#GetV6CredentialsByID) | **Get** /api/v6/rest/entries/{id} | get_entry
[**GetV6Folders**](DefaultAPI.md#GetV6Folders) | **Get** /api/v6/rest/folders | 
[**GetV6FoldersByID**](DefaultAPI.md#GetV6FoldersByID) | **Get** /api/v6/rest/folders/{id} | 
[**GetV6FoldersRoot**](DefaultAPI.md#GetV6FoldersRoot) | **Get** /api/v6/rest/folders/root | 
[**GetV6ServerInfo**](DefaultAPI.md#GetV6ServerInfo) | **Get** /api/v6/rest/getserverinfo | 
[**PatchV6CredentialsByID**](DefaultAPI.md#PatchV6CredentialsByID) | **Patch** /api/v6/rest/entries/{id} | 
[**PatchV6FoldersByID**](DefaultAPI.md#PatchV6FoldersByID) | **Patch** /api/v6/rest/folders/{id} | 
[**PostV6Credentials**](DefaultAPI.md#PostV6Credentials) | **Post** /api/v6/rest/entries | 
[**PostV6Folders**](DefaultAPI.md#PostV6Folders) | **Post** /api/v6/rest/folders | Create New Folder
[**PostV6Search**](DefaultAPI.md#PostV6Search) | **Post** /api/v6/rest/search | 



## DeleteV6CredentialsByID

> DeleteV6CredentialsByID(ctx, id).Execute()



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
    r, err := apiClient.DefaultAPI.DeleteV6CredentialsByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteV6CredentialsByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteV6CredentialsByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer_token](../README.md#Bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteV6FoldersByID

> DeleteV6FoldersByID(ctx, id).Execute()



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
    r, err := apiClient.DefaultAPI.DeleteV6FoldersByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteV6FoldersByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteV6FoldersByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer_token](../README.md#Bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV6CredentialPasswordByID

> string GetV6CredentialPasswordByID(ctx, id).Execute()

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
    resp, r, err := apiClient.DefaultAPI.GetV6CredentialPasswordByID(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetV6CredentialPasswordByID``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetV6CredentialPasswordByID`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetV6CredentialPasswordByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV6CredentialPasswordByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[Bearer_token](../README.md#Bearer_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

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

[Bearer_token](../README.md#Bearer_token)

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


## GetV6FoldersRoot

> string GetV6FoldersRoot(ctx).Execute()





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
    resp, r, err := apiClient.DefaultAPI.GetV6FoldersRoot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetV6FoldersRoot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetV6FoldersRoot`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetV6FoldersRoot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV6FoldersRootRequest struct via the builder pattern


### Return type

**string**

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


## PatchV6CredentialsByID

> PatchV6CredentialsByID(ctx, id).V6CredentialInput(v6CredentialInput).Execute()



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
    v6CredentialInput := *openapiclient.NewV6CredentialInput() // V6CredentialInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.PatchV6CredentialsByID(context.Background(), id).V6CredentialInput(v6CredentialInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PatchV6CredentialsByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchV6CredentialsByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v6CredentialInput** | [**V6CredentialInput**](V6CredentialInput.md) |  | 

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


## PatchV6FoldersByID

> PatchV6FoldersByID(ctx, id).V6CredentialGroupInput(v6CredentialGroupInput).Execute()



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
    v6CredentialGroupInput := *openapiclient.NewV6CredentialGroupInput() // V6CredentialGroupInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.PatchV6FoldersByID(context.Background(), id).V6CredentialGroupInput(v6CredentialGroupInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PatchV6FoldersByID``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPatchV6FoldersByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **v6CredentialGroupInput** | [**V6CredentialGroupInput**](V6CredentialGroupInput.md) |  | 

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


## PostV6Credentials

> string PostV6Credentials(ctx).V6CredentialInput(v6CredentialInput).Execute()



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
    v6CredentialInput := *openapiclient.NewV6CredentialInput() // V6CredentialInput | Create Credential

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.PostV6Credentials(context.Background()).V6CredentialInput(v6CredentialInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostV6Credentials``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostV6Credentials`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostV6Credentials`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV6CredentialsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v6CredentialInput** | [**V6CredentialInput**](V6CredentialInput.md) | Create Credential | 

### Return type

**string**

### Authorization

[Bearer_token](../README.md#Bearer_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV6Folders

> string PostV6Folders(ctx).V6CredentialGroupInput(v6CredentialGroupInput).Execute()

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
    v6CredentialGroupInput := *openapiclient.NewV6CredentialGroupInput() // V6CredentialGroupInput | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultAPI.PostV6Folders(context.Background()).V6CredentialGroupInput(v6CredentialGroupInput).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.PostV6Folders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostV6Folders`: string
    fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.PostV6Folders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV6FoldersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **v6CredentialGroupInput** | [**V6CredentialGroupInput**](V6CredentialGroupInput.md) |  | 

### Return type

**string**

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

