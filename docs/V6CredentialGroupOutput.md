# V6CredentialGroupOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomUserFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomApplicationFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Children** | Pointer to [**[]V6CredentialGroupOutput**](V6CredentialGroupOutput.md) |  | [optional] 
**Credentials** | Pointer to [**[]V6CredentialResult**](V6CredentialResult.md) |  | [optional] 
**Tags** | Pointer to [**[]V6TagResult**](V6TagResult.md) |  | [optional] 
**HasModifyEntriesAccess** | Pointer to **bool** |  | [optional] 
**HasViewEntryContentsAccess** | Pointer to **bool** |  | [optional] 
**CommentPrompts** | Pointer to [**V6CommentPromptResult**](V6CommentPromptResult.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Expires** | Pointer to **string** |  | [optional] 

## Methods

### NewV6CredentialGroupOutput

`func NewV6CredentialGroupOutput() *V6CredentialGroupOutput`

NewV6CredentialGroupOutput instantiates a new V6CredentialGroupOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV6CredentialGroupOutputWithDefaults

`func NewV6CredentialGroupOutputWithDefaults() *V6CredentialGroupOutput`

NewV6CredentialGroupOutputWithDefaults instantiates a new V6CredentialGroupOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomUserFields

`func (o *V6CredentialGroupOutput) GetCustomUserFields() map[string]interface{}`

GetCustomUserFields returns the CustomUserFields field if non-nil, zero value otherwise.

### GetCustomUserFieldsOk

`func (o *V6CredentialGroupOutput) GetCustomUserFieldsOk() (*map[string]interface{}, bool)`

GetCustomUserFieldsOk returns a tuple with the CustomUserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserFields

`func (o *V6CredentialGroupOutput) SetCustomUserFields(v map[string]interface{})`

SetCustomUserFields sets CustomUserFields field to given value.

### HasCustomUserFields

`func (o *V6CredentialGroupOutput) HasCustomUserFields() bool`

HasCustomUserFields returns a boolean if a field has been set.

### GetCustomApplicationFields

`func (o *V6CredentialGroupOutput) GetCustomApplicationFields() map[string]interface{}`

GetCustomApplicationFields returns the CustomApplicationFields field if non-nil, zero value otherwise.

### GetCustomApplicationFieldsOk

`func (o *V6CredentialGroupOutput) GetCustomApplicationFieldsOk() (*map[string]interface{}, bool)`

GetCustomApplicationFieldsOk returns a tuple with the CustomApplicationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomApplicationFields

`func (o *V6CredentialGroupOutput) SetCustomApplicationFields(v map[string]interface{})`

SetCustomApplicationFields sets CustomApplicationFields field to given value.

### HasCustomApplicationFields

`func (o *V6CredentialGroupOutput) HasCustomApplicationFields() bool`

HasCustomApplicationFields returns a boolean if a field has been set.

### GetChildren

`func (o *V6CredentialGroupOutput) GetChildren() []V6CredentialGroupOutput`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *V6CredentialGroupOutput) GetChildrenOk() (*[]V6CredentialGroupOutput, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *V6CredentialGroupOutput) SetChildren(v []V6CredentialGroupOutput)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *V6CredentialGroupOutput) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCredentials

`func (o *V6CredentialGroupOutput) GetCredentials() []V6CredentialResult`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *V6CredentialGroupOutput) GetCredentialsOk() (*[]V6CredentialResult, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *V6CredentialGroupOutput) SetCredentials(v []V6CredentialResult)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *V6CredentialGroupOutput) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetTags

`func (o *V6CredentialGroupOutput) GetTags() []V6TagResult`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V6CredentialGroupOutput) GetTagsOk() (*[]V6TagResult, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V6CredentialGroupOutput) SetTags(v []V6TagResult)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V6CredentialGroupOutput) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHasModifyEntriesAccess

`func (o *V6CredentialGroupOutput) GetHasModifyEntriesAccess() bool`

GetHasModifyEntriesAccess returns the HasModifyEntriesAccess field if non-nil, zero value otherwise.

### GetHasModifyEntriesAccessOk

`func (o *V6CredentialGroupOutput) GetHasModifyEntriesAccessOk() (*bool, bool)`

GetHasModifyEntriesAccessOk returns a tuple with the HasModifyEntriesAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasModifyEntriesAccess

`func (o *V6CredentialGroupOutput) SetHasModifyEntriesAccess(v bool)`

SetHasModifyEntriesAccess sets HasModifyEntriesAccess field to given value.

### HasHasModifyEntriesAccess

`func (o *V6CredentialGroupOutput) HasHasModifyEntriesAccess() bool`

HasHasModifyEntriesAccess returns a boolean if a field has been set.

### GetHasViewEntryContentsAccess

`func (o *V6CredentialGroupOutput) GetHasViewEntryContentsAccess() bool`

GetHasViewEntryContentsAccess returns the HasViewEntryContentsAccess field if non-nil, zero value otherwise.

### GetHasViewEntryContentsAccessOk

`func (o *V6CredentialGroupOutput) GetHasViewEntryContentsAccessOk() (*bool, bool)`

GetHasViewEntryContentsAccessOk returns a tuple with the HasViewEntryContentsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasViewEntryContentsAccess

`func (o *V6CredentialGroupOutput) SetHasViewEntryContentsAccess(v bool)`

SetHasViewEntryContentsAccess sets HasViewEntryContentsAccess field to given value.

### HasHasViewEntryContentsAccess

`func (o *V6CredentialGroupOutput) HasHasViewEntryContentsAccess() bool`

HasHasViewEntryContentsAccess returns a boolean if a field has been set.

### GetCommentPrompts

`func (o *V6CredentialGroupOutput) GetCommentPrompts() V6CommentPromptResult`

GetCommentPrompts returns the CommentPrompts field if non-nil, zero value otherwise.

### GetCommentPromptsOk

`func (o *V6CredentialGroupOutput) GetCommentPromptsOk() (*V6CommentPromptResult, bool)`

GetCommentPromptsOk returns a tuple with the CommentPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentPrompts

`func (o *V6CredentialGroupOutput) SetCommentPrompts(v V6CommentPromptResult)`

SetCommentPrompts sets CommentPrompts field to given value.

### HasCommentPrompts

`func (o *V6CredentialGroupOutput) HasCommentPrompts() bool`

HasCommentPrompts returns a boolean if a field has been set.

### GetId

`func (o *V6CredentialGroupOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V6CredentialGroupOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V6CredentialGroupOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V6CredentialGroupOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V6CredentialGroupOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V6CredentialGroupOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V6CredentialGroupOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V6CredentialGroupOutput) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentId

`func (o *V6CredentialGroupOutput) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *V6CredentialGroupOutput) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *V6CredentialGroupOutput) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *V6CredentialGroupOutput) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetNotes

`func (o *V6CredentialGroupOutput) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *V6CredentialGroupOutput) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *V6CredentialGroupOutput) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *V6CredentialGroupOutput) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCreated

`func (o *V6CredentialGroupOutput) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *V6CredentialGroupOutput) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *V6CredentialGroupOutput) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *V6CredentialGroupOutput) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *V6CredentialGroupOutput) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *V6CredentialGroupOutput) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *V6CredentialGroupOutput) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *V6CredentialGroupOutput) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetExpires

`func (o *V6CredentialGroupOutput) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *V6CredentialGroupOutput) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *V6CredentialGroupOutput) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *V6CredentialGroupOutput) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


