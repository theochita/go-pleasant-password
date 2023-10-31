# \AuthenticationAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostOauthToken**](AuthenticationAPI.md#PostOauthToken) | **Post** /OAuth2/Token | 



## PostOauthToken

> Oauth2TokenOutput PostOauthToken(ctx).GrantType(grantType).Username(username).Password(password).XPleasantOTP(xPleasantOTP).XPleasantOTPProvider(xPleasantOTPProvider).Execute()



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
    grantType := "grantType_example" // string | 
    username := "username_example" // string | 
    password := "password_example" // string | 
    xPleasantOTP := "15686" // string | Only Required if user is Oauth (optional)
    xPleasantOTPProvider := "authenticator-app" // string | Only required if user is Oauth (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthenticationAPI.PostOauthToken(context.Background()).GrantType(grantType).Username(username).Password(password).XPleasantOTP(xPleasantOTP).XPleasantOTPProvider(xPleasantOTPProvider).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationAPI.PostOauthToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostOauthToken`: Oauth2TokenOutput
    fmt.Fprintf(os.Stdout, "Response from `AuthenticationAPI.PostOauthToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostOauthTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **username** | **string** |  | 
 **password** | **string** |  | 
 **xPleasantOTP** | **string** | Only Required if user is Oauth | 
 **xPleasantOTPProvider** | **string** | Only required if user is Oauth | 

### Return type

[**Oauth2TokenOutput**](Oauth2TokenOutput.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

