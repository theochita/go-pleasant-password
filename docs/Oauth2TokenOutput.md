# Oauth2TokenOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | Pointer to **string** |  | [optional] 
**TokenType** | Pointer to **string** |  | [optional] 
**ExpiresIn** | Pointer to **int32** |  | [optional] 

## Methods

### NewOauth2TokenOutput

`func NewOauth2TokenOutput() *Oauth2TokenOutput`

NewOauth2TokenOutput instantiates a new Oauth2TokenOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2TokenOutputWithDefaults

`func NewOauth2TokenOutputWithDefaults() *Oauth2TokenOutput`

NewOauth2TokenOutputWithDefaults instantiates a new Oauth2TokenOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Oauth2TokenOutput) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Oauth2TokenOutput) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Oauth2TokenOutput) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *Oauth2TokenOutput) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### GetTokenType

`func (o *Oauth2TokenOutput) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Oauth2TokenOutput) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Oauth2TokenOutput) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.

### HasTokenType

`func (o *Oauth2TokenOutput) HasTokenType() bool`

HasTokenType returns a boolean if a field has been set.

### GetExpiresIn

`func (o *Oauth2TokenOutput) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *Oauth2TokenOutput) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *Oauth2TokenOutput) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *Oauth2TokenOutput) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


