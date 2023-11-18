# V6CredentialInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomUserFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to [**[]V6TagResult**](V6TagResult.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to **string** |  | [optional] 

## Methods

### NewV6CredentialInput

`func NewV6CredentialInput() *V6CredentialInput`

NewV6CredentialInput instantiates a new V6CredentialInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV6CredentialInputWithDefaults

`func NewV6CredentialInputWithDefaults() *V6CredentialInput`

NewV6CredentialInputWithDefaults instantiates a new V6CredentialInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomUserFields

`func (o *V6CredentialInput) GetCustomUserFields() map[string]interface{}`

GetCustomUserFields returns the CustomUserFields field if non-nil, zero value otherwise.

### GetCustomUserFieldsOk

`func (o *V6CredentialInput) GetCustomUserFieldsOk() (*map[string]interface{}, bool)`

GetCustomUserFieldsOk returns a tuple with the CustomUserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserFields

`func (o *V6CredentialInput) SetCustomUserFields(v map[string]interface{})`

SetCustomUserFields sets CustomUserFields field to given value.

### HasCustomUserFields

`func (o *V6CredentialInput) HasCustomUserFields() bool`

HasCustomUserFields returns a boolean if a field has been set.

### GetTags

`func (o *V6CredentialInput) GetTags() []V6TagResult`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V6CredentialInput) GetTagsOk() (*[]V6TagResult, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V6CredentialInput) SetTags(v []V6TagResult)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V6CredentialInput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetName

`func (o *V6CredentialInput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V6CredentialInput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V6CredentialInput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V6CredentialInput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *V6CredentialInput) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V6CredentialInput) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V6CredentialInput) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V6CredentialInput) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *V6CredentialInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V6CredentialInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V6CredentialInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V6CredentialInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrl

`func (o *V6CredentialInput) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V6CredentialInput) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V6CredentialInput) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V6CredentialInput) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNotes

`func (o *V6CredentialInput) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *V6CredentialInput) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *V6CredentialInput) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *V6CredentialInput) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetGroupId

`func (o *V6CredentialInput) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V6CredentialInput) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V6CredentialInput) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V6CredentialInput) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetExpires

`func (o *V6CredentialInput) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *V6CredentialInput) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *V6CredentialInput) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *V6CredentialInput) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


