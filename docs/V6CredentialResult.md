# V6CredentialResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomUserFields** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomApplicationFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Attachments** | Pointer to [**[]V6AttachmentResult**](V6AttachmentResult.md) |  | [optional] 
**Tags** | Pointer to [**[]V6TagResult**](V6TagResult.md) |  | [optional] 
**HasModifyEntriesAccess** | Pointer to **bool** |  | [optional] 
**HasViewEntryContentsAccess** | Pointer to **bool** |  | [optional] 
**HasViewEntryPasswordAccess** | Pointer to **bool** |  | [optional] 
**HasViewTOTPAccess** | Pointer to **bool** |  | [optional] 
**HasModifyTOTPAccess** | Pointer to **bool** |  | [optional] 
**CommentPrompts** | Pointer to [**V6CommentPromptResult**](V6CommentPromptResult.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**GroupId** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Modified** | Pointer to **time.Time** |  | [optional] 
**Expires** | Pointer to **string** |  | [optional] 
**TOTPIssuer** | Pointer to **string** |  | [optional] 
**TOTPSecret** | Pointer to **string** |  | [optional] 
**TOTPDigits** | Pointer to **int32** |  | [optional] 
**TOTPPeriod** | Pointer to **int32** |  | [optional] 

## Methods

### NewV6CredentialResult

`func NewV6CredentialResult() *V6CredentialResult`

NewV6CredentialResult instantiates a new V6CredentialResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV6CredentialResultWithDefaults

`func NewV6CredentialResultWithDefaults() *V6CredentialResult`

NewV6CredentialResultWithDefaults instantiates a new V6CredentialResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomUserFields

`func (o *V6CredentialResult) GetCustomUserFields() map[string]interface{}`

GetCustomUserFields returns the CustomUserFields field if non-nil, zero value otherwise.

### GetCustomUserFieldsOk

`func (o *V6CredentialResult) GetCustomUserFieldsOk() (*map[string]interface{}, bool)`

GetCustomUserFieldsOk returns a tuple with the CustomUserFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomUserFields

`func (o *V6CredentialResult) SetCustomUserFields(v map[string]interface{})`

SetCustomUserFields sets CustomUserFields field to given value.

### HasCustomUserFields

`func (o *V6CredentialResult) HasCustomUserFields() bool`

HasCustomUserFields returns a boolean if a field has been set.

### GetCustomApplicationFields

`func (o *V6CredentialResult) GetCustomApplicationFields() map[string]interface{}`

GetCustomApplicationFields returns the CustomApplicationFields field if non-nil, zero value otherwise.

### GetCustomApplicationFieldsOk

`func (o *V6CredentialResult) GetCustomApplicationFieldsOk() (*map[string]interface{}, bool)`

GetCustomApplicationFieldsOk returns a tuple with the CustomApplicationFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomApplicationFields

`func (o *V6CredentialResult) SetCustomApplicationFields(v map[string]interface{})`

SetCustomApplicationFields sets CustomApplicationFields field to given value.

### HasCustomApplicationFields

`func (o *V6CredentialResult) HasCustomApplicationFields() bool`

HasCustomApplicationFields returns a boolean if a field has been set.

### GetAttachments

`func (o *V6CredentialResult) GetAttachments() []V6AttachmentResult`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *V6CredentialResult) GetAttachmentsOk() (*[]V6AttachmentResult, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *V6CredentialResult) SetAttachments(v []V6AttachmentResult)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *V6CredentialResult) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetTags

`func (o *V6CredentialResult) GetTags() []V6TagResult`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *V6CredentialResult) GetTagsOk() (*[]V6TagResult, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *V6CredentialResult) SetTags(v []V6TagResult)`

SetTags sets Tags field to given value.

### HasTags

`func (o *V6CredentialResult) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHasModifyEntriesAccess

`func (o *V6CredentialResult) GetHasModifyEntriesAccess() bool`

GetHasModifyEntriesAccess returns the HasModifyEntriesAccess field if non-nil, zero value otherwise.

### GetHasModifyEntriesAccessOk

`func (o *V6CredentialResult) GetHasModifyEntriesAccessOk() (*bool, bool)`

GetHasModifyEntriesAccessOk returns a tuple with the HasModifyEntriesAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasModifyEntriesAccess

`func (o *V6CredentialResult) SetHasModifyEntriesAccess(v bool)`

SetHasModifyEntriesAccess sets HasModifyEntriesAccess field to given value.

### HasHasModifyEntriesAccess

`func (o *V6CredentialResult) HasHasModifyEntriesAccess() bool`

HasHasModifyEntriesAccess returns a boolean if a field has been set.

### GetHasViewEntryContentsAccess

`func (o *V6CredentialResult) GetHasViewEntryContentsAccess() bool`

GetHasViewEntryContentsAccess returns the HasViewEntryContentsAccess field if non-nil, zero value otherwise.

### GetHasViewEntryContentsAccessOk

`func (o *V6CredentialResult) GetHasViewEntryContentsAccessOk() (*bool, bool)`

GetHasViewEntryContentsAccessOk returns a tuple with the HasViewEntryContentsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasViewEntryContentsAccess

`func (o *V6CredentialResult) SetHasViewEntryContentsAccess(v bool)`

SetHasViewEntryContentsAccess sets HasViewEntryContentsAccess field to given value.

### HasHasViewEntryContentsAccess

`func (o *V6CredentialResult) HasHasViewEntryContentsAccess() bool`

HasHasViewEntryContentsAccess returns a boolean if a field has been set.

### GetHasViewEntryPasswordAccess

`func (o *V6CredentialResult) GetHasViewEntryPasswordAccess() bool`

GetHasViewEntryPasswordAccess returns the HasViewEntryPasswordAccess field if non-nil, zero value otherwise.

### GetHasViewEntryPasswordAccessOk

`func (o *V6CredentialResult) GetHasViewEntryPasswordAccessOk() (*bool, bool)`

GetHasViewEntryPasswordAccessOk returns a tuple with the HasViewEntryPasswordAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasViewEntryPasswordAccess

`func (o *V6CredentialResult) SetHasViewEntryPasswordAccess(v bool)`

SetHasViewEntryPasswordAccess sets HasViewEntryPasswordAccess field to given value.

### HasHasViewEntryPasswordAccess

`func (o *V6CredentialResult) HasHasViewEntryPasswordAccess() bool`

HasHasViewEntryPasswordAccess returns a boolean if a field has been set.

### GetHasViewTOTPAccess

`func (o *V6CredentialResult) GetHasViewTOTPAccess() bool`

GetHasViewTOTPAccess returns the HasViewTOTPAccess field if non-nil, zero value otherwise.

### GetHasViewTOTPAccessOk

`func (o *V6CredentialResult) GetHasViewTOTPAccessOk() (*bool, bool)`

GetHasViewTOTPAccessOk returns a tuple with the HasViewTOTPAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasViewTOTPAccess

`func (o *V6CredentialResult) SetHasViewTOTPAccess(v bool)`

SetHasViewTOTPAccess sets HasViewTOTPAccess field to given value.

### HasHasViewTOTPAccess

`func (o *V6CredentialResult) HasHasViewTOTPAccess() bool`

HasHasViewTOTPAccess returns a boolean if a field has been set.

### GetHasModifyTOTPAccess

`func (o *V6CredentialResult) GetHasModifyTOTPAccess() bool`

GetHasModifyTOTPAccess returns the HasModifyTOTPAccess field if non-nil, zero value otherwise.

### GetHasModifyTOTPAccessOk

`func (o *V6CredentialResult) GetHasModifyTOTPAccessOk() (*bool, bool)`

GetHasModifyTOTPAccessOk returns a tuple with the HasModifyTOTPAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasModifyTOTPAccess

`func (o *V6CredentialResult) SetHasModifyTOTPAccess(v bool)`

SetHasModifyTOTPAccess sets HasModifyTOTPAccess field to given value.

### HasHasModifyTOTPAccess

`func (o *V6CredentialResult) HasHasModifyTOTPAccess() bool`

HasHasModifyTOTPAccess returns a boolean if a field has been set.

### GetCommentPrompts

`func (o *V6CredentialResult) GetCommentPrompts() V6CommentPromptResult`

GetCommentPrompts returns the CommentPrompts field if non-nil, zero value otherwise.

### GetCommentPromptsOk

`func (o *V6CredentialResult) GetCommentPromptsOk() (*V6CommentPromptResult, bool)`

GetCommentPromptsOk returns a tuple with the CommentPrompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentPrompts

`func (o *V6CredentialResult) SetCommentPrompts(v V6CommentPromptResult)`

SetCommentPrompts sets CommentPrompts field to given value.

### HasCommentPrompts

`func (o *V6CredentialResult) HasCommentPrompts() bool`

HasCommentPrompts returns a boolean if a field has been set.

### GetId

`func (o *V6CredentialResult) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V6CredentialResult) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V6CredentialResult) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V6CredentialResult) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V6CredentialResult) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V6CredentialResult) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V6CredentialResult) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V6CredentialResult) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *V6CredentialResult) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *V6CredentialResult) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *V6CredentialResult) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *V6CredentialResult) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *V6CredentialResult) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *V6CredentialResult) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *V6CredentialResult) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *V6CredentialResult) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUrl

`func (o *V6CredentialResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *V6CredentialResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *V6CredentialResult) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *V6CredentialResult) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNotes

`func (o *V6CredentialResult) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *V6CredentialResult) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *V6CredentialResult) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *V6CredentialResult) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetGroupId

`func (o *V6CredentialResult) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *V6CredentialResult) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *V6CredentialResult) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *V6CredentialResult) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetCreated

`func (o *V6CredentialResult) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *V6CredentialResult) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *V6CredentialResult) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *V6CredentialResult) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *V6CredentialResult) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *V6CredentialResult) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *V6CredentialResult) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *V6CredentialResult) HasModified() bool`

HasModified returns a boolean if a field has been set.

### GetExpires

`func (o *V6CredentialResult) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *V6CredentialResult) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *V6CredentialResult) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *V6CredentialResult) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetTOTPIssuer

`func (o *V6CredentialResult) GetTOTPIssuer() string`

GetTOTPIssuer returns the TOTPIssuer field if non-nil, zero value otherwise.

### GetTOTPIssuerOk

`func (o *V6CredentialResult) GetTOTPIssuerOk() (*string, bool)`

GetTOTPIssuerOk returns a tuple with the TOTPIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTOTPIssuer

`func (o *V6CredentialResult) SetTOTPIssuer(v string)`

SetTOTPIssuer sets TOTPIssuer field to given value.

### HasTOTPIssuer

`func (o *V6CredentialResult) HasTOTPIssuer() bool`

HasTOTPIssuer returns a boolean if a field has been set.

### GetTOTPSecret

`func (o *V6CredentialResult) GetTOTPSecret() string`

GetTOTPSecret returns the TOTPSecret field if non-nil, zero value otherwise.

### GetTOTPSecretOk

`func (o *V6CredentialResult) GetTOTPSecretOk() (*string, bool)`

GetTOTPSecretOk returns a tuple with the TOTPSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTOTPSecret

`func (o *V6CredentialResult) SetTOTPSecret(v string)`

SetTOTPSecret sets TOTPSecret field to given value.

### HasTOTPSecret

`func (o *V6CredentialResult) HasTOTPSecret() bool`

HasTOTPSecret returns a boolean if a field has been set.

### GetTOTPDigits

`func (o *V6CredentialResult) GetTOTPDigits() int32`

GetTOTPDigits returns the TOTPDigits field if non-nil, zero value otherwise.

### GetTOTPDigitsOk

`func (o *V6CredentialResult) GetTOTPDigitsOk() (*int32, bool)`

GetTOTPDigitsOk returns a tuple with the TOTPDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTOTPDigits

`func (o *V6CredentialResult) SetTOTPDigits(v int32)`

SetTOTPDigits sets TOTPDigits field to given value.

### HasTOTPDigits

`func (o *V6CredentialResult) HasTOTPDigits() bool`

HasTOTPDigits returns a boolean if a field has been set.

### GetTOTPPeriod

`func (o *V6CredentialResult) GetTOTPPeriod() int32`

GetTOTPPeriod returns the TOTPPeriod field if non-nil, zero value otherwise.

### GetTOTPPeriodOk

`func (o *V6CredentialResult) GetTOTPPeriodOk() (*int32, bool)`

GetTOTPPeriodOk returns a tuple with the TOTPPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTOTPPeriod

`func (o *V6CredentialResult) SetTOTPPeriod(v int32)`

SetTOTPPeriod sets TOTPPeriod field to given value.

### HasTOTPPeriod

`func (o *V6CredentialResult) HasTOTPPeriod() bool`

HasTOTPPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


