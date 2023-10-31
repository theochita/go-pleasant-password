# V6SearchOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]V6CredentialSearchResult**](V6CredentialSearchResult.md) |  | [optional] 
**Groups** | Pointer to [**[]V6CredentialGroupSearchResult**](V6CredentialGroupSearchResult.md) |  | [optional] 

## Methods

### NewV6SearchOutput

`func NewV6SearchOutput() *V6SearchOutput`

NewV6SearchOutput instantiates a new V6SearchOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV6SearchOutputWithDefaults

`func NewV6SearchOutputWithDefaults() *V6SearchOutput`

NewV6SearchOutputWithDefaults instantiates a new V6SearchOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *V6SearchOutput) GetCredentials() []V6CredentialSearchResult`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *V6SearchOutput) GetCredentialsOk() (*[]V6CredentialSearchResult, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *V6SearchOutput) SetCredentials(v []V6CredentialSearchResult)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *V6SearchOutput) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetGroups

`func (o *V6SearchOutput) GetGroups() []V6CredentialGroupSearchResult`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *V6SearchOutput) GetGroupsOk() (*[]V6CredentialGroupSearchResult, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *V6SearchOutput) SetGroups(v []V6CredentialGroupSearchResult)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *V6SearchOutput) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


