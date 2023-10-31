# Oauth2TokenInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantType** | **string** |  | 
**Username** | **string** |  | 
**Password** | **string** |  | 

## Methods

### NewOauth2TokenInput

`func NewOauth2TokenInput(grantType string, username string, password string, ) *Oauth2TokenInput`

NewOauth2TokenInput instantiates a new Oauth2TokenInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2TokenInputWithDefaults

`func NewOauth2TokenInputWithDefaults() *Oauth2TokenInput`

NewOauth2TokenInputWithDefaults instantiates a new Oauth2TokenInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantType

`func (o *Oauth2TokenInput) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *Oauth2TokenInput) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *Oauth2TokenInput) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.


### GetUsername

`func (o *Oauth2TokenInput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Oauth2TokenInput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Oauth2TokenInput) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *Oauth2TokenInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *Oauth2TokenInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *Oauth2TokenInput) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


