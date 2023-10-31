/*
pleasant password server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the V6CredentialGroupOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V6CredentialGroupOutput{}

// V6CredentialGroupOutput 
type V6CredentialGroupOutput struct {
	// 
	CustomUserFields map[string]interface{} `json:"CustomUserFields,omitempty"`
	// 
	CustomApplicationFields map[string]interface{} `json:"CustomApplicationFields,omitempty"`
	Children []V6CredentialGroupOutput `json:"Children,omitempty"`
	Credentials []V6CredentialResult `json:"Credentials,omitempty"`
	Tags []interface{} `json:"Tags,omitempty"`
	HasModifyEntriesAccess *bool `json:"HasModifyEntriesAccess,omitempty"`
	HasViewEntryContentsAccess *bool `json:"HasViewEntryContentsAccess,omitempty"`
	CommentPrompts *V6CommentPromptResult `json:"CommentPrompts,omitempty"`
	Id *string `json:"Id,omitempty"`
	Name *string `json:"Name,omitempty"`
	ParentId *string `json:"ParentId,omitempty"`
	Notes *string `json:"Notes,omitempty"`
	Created *time.Time `json:"Created,omitempty"`
	Modified *time.Time `json:"Modified,omitempty"`
	Expires *string `json:"Expires,omitempty"`
}

// NewV6CredentialGroupOutput instantiates a new V6CredentialGroupOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV6CredentialGroupOutput() *V6CredentialGroupOutput {
	this := V6CredentialGroupOutput{}
	return &this
}

// NewV6CredentialGroupOutputWithDefaults instantiates a new V6CredentialGroupOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV6CredentialGroupOutputWithDefaults() *V6CredentialGroupOutput {
	this := V6CredentialGroupOutput{}
	return &this
}

// GetCustomUserFields returns the CustomUserFields field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetCustomUserFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomUserFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomUserFields
}

// GetCustomUserFieldsOk returns a tuple with the CustomUserFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetCustomUserFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomUserFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomUserFields, true
}

// HasCustomUserFields returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasCustomUserFields() bool {
	if o != nil && !IsNil(o.CustomUserFields) {
		return true
	}

	return false
}

// SetCustomUserFields gets a reference to the given map[string]interface{} and assigns it to the CustomUserFields field.
func (o *V6CredentialGroupOutput) SetCustomUserFields(v map[string]interface{}) {
	o.CustomUserFields = v
}

// GetCustomApplicationFields returns the CustomApplicationFields field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetCustomApplicationFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomApplicationFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomApplicationFields
}

// GetCustomApplicationFieldsOk returns a tuple with the CustomApplicationFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetCustomApplicationFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomApplicationFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomApplicationFields, true
}

// HasCustomApplicationFields returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasCustomApplicationFields() bool {
	if o != nil && !IsNil(o.CustomApplicationFields) {
		return true
	}

	return false
}

// SetCustomApplicationFields gets a reference to the given map[string]interface{} and assigns it to the CustomApplicationFields field.
func (o *V6CredentialGroupOutput) SetCustomApplicationFields(v map[string]interface{}) {
	o.CustomApplicationFields = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetChildren() []V6CredentialGroupOutput {
	if o == nil || IsNil(o.Children) {
		var ret []V6CredentialGroupOutput
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetChildrenOk() ([]V6CredentialGroupOutput, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []V6CredentialGroupOutput and assigns it to the Children field.
func (o *V6CredentialGroupOutput) SetChildren(v []V6CredentialGroupOutput) {
	o.Children = v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetCredentials() []V6CredentialResult {
	if o == nil || IsNil(o.Credentials) {
		var ret []V6CredentialResult
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetCredentialsOk() ([]V6CredentialResult, bool) {
	if o == nil || IsNil(o.Credentials) {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasCredentials() bool {
	if o != nil && !IsNil(o.Credentials) {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []V6CredentialResult and assigns it to the Credentials field.
func (o *V6CredentialGroupOutput) SetCredentials(v []V6CredentialResult) {
	o.Credentials = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetTags() []interface{} {
	if o == nil || IsNil(o.Tags) {
		var ret []interface{}
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetTagsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []interface{} and assigns it to the Tags field.
func (o *V6CredentialGroupOutput) SetTags(v []interface{}) {
	o.Tags = v
}

// GetHasModifyEntriesAccess returns the HasModifyEntriesAccess field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetHasModifyEntriesAccess() bool {
	if o == nil || IsNil(o.HasModifyEntriesAccess) {
		var ret bool
		return ret
	}
	return *o.HasModifyEntriesAccess
}

// GetHasModifyEntriesAccessOk returns a tuple with the HasModifyEntriesAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetHasModifyEntriesAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.HasModifyEntriesAccess) {
		return nil, false
	}
	return o.HasModifyEntriesAccess, true
}

// HasHasModifyEntriesAccess returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasHasModifyEntriesAccess() bool {
	if o != nil && !IsNil(o.HasModifyEntriesAccess) {
		return true
	}

	return false
}

// SetHasModifyEntriesAccess gets a reference to the given bool and assigns it to the HasModifyEntriesAccess field.
func (o *V6CredentialGroupOutput) SetHasModifyEntriesAccess(v bool) {
	o.HasModifyEntriesAccess = &v
}

// GetHasViewEntryContentsAccess returns the HasViewEntryContentsAccess field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetHasViewEntryContentsAccess() bool {
	if o == nil || IsNil(o.HasViewEntryContentsAccess) {
		var ret bool
		return ret
	}
	return *o.HasViewEntryContentsAccess
}

// GetHasViewEntryContentsAccessOk returns a tuple with the HasViewEntryContentsAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetHasViewEntryContentsAccessOk() (*bool, bool) {
	if o == nil || IsNil(o.HasViewEntryContentsAccess) {
		return nil, false
	}
	return o.HasViewEntryContentsAccess, true
}

// HasHasViewEntryContentsAccess returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasHasViewEntryContentsAccess() bool {
	if o != nil && !IsNil(o.HasViewEntryContentsAccess) {
		return true
	}

	return false
}

// SetHasViewEntryContentsAccess gets a reference to the given bool and assigns it to the HasViewEntryContentsAccess field.
func (o *V6CredentialGroupOutput) SetHasViewEntryContentsAccess(v bool) {
	o.HasViewEntryContentsAccess = &v
}

// GetCommentPrompts returns the CommentPrompts field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetCommentPrompts() V6CommentPromptResult {
	if o == nil || IsNil(o.CommentPrompts) {
		var ret V6CommentPromptResult
		return ret
	}
	return *o.CommentPrompts
}

// GetCommentPromptsOk returns a tuple with the CommentPrompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetCommentPromptsOk() (*V6CommentPromptResult, bool) {
	if o == nil || IsNil(o.CommentPrompts) {
		return nil, false
	}
	return o.CommentPrompts, true
}

// HasCommentPrompts returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasCommentPrompts() bool {
	if o != nil && !IsNil(o.CommentPrompts) {
		return true
	}

	return false
}

// SetCommentPrompts gets a reference to the given V6CommentPromptResult and assigns it to the CommentPrompts field.
func (o *V6CredentialGroupOutput) SetCommentPrompts(v V6CommentPromptResult) {
	o.CommentPrompts = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V6CredentialGroupOutput) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V6CredentialGroupOutput) SetName(v string) {
	o.Name = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *V6CredentialGroupOutput) SetParentId(v string) {
	o.ParentId = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *V6CredentialGroupOutput) SetNotes(v string) {
	o.Notes = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *V6CredentialGroupOutput) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *V6CredentialGroupOutput) SetModified(v time.Time) {
	o.Modified = &v
}

// GetExpires returns the Expires field value if set, zero value otherwise.
func (o *V6CredentialGroupOutput) GetExpires() string {
	if o == nil || IsNil(o.Expires) {
		var ret string
		return ret
	}
	return *o.Expires
}

// GetExpiresOk returns a tuple with the Expires field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupOutput) GetExpiresOk() (*string, bool) {
	if o == nil || IsNil(o.Expires) {
		return nil, false
	}
	return o.Expires, true
}

// HasExpires returns a boolean if a field has been set.
func (o *V6CredentialGroupOutput) HasExpires() bool {
	if o != nil && !IsNil(o.Expires) {
		return true
	}

	return false
}

// SetExpires gets a reference to the given string and assigns it to the Expires field.
func (o *V6CredentialGroupOutput) SetExpires(v string) {
	o.Expires = &v
}

func (o V6CredentialGroupOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V6CredentialGroupOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CustomUserFields) {
		toSerialize["CustomUserFields"] = o.CustomUserFields
	}
	if !IsNil(o.CustomApplicationFields) {
		toSerialize["CustomApplicationFields"] = o.CustomApplicationFields
	}
	if !IsNil(o.Children) {
		toSerialize["Children"] = o.Children
	}
	if !IsNil(o.Credentials) {
		toSerialize["Credentials"] = o.Credentials
	}
	if !IsNil(o.Tags) {
		toSerialize["Tags"] = o.Tags
	}
	if !IsNil(o.HasModifyEntriesAccess) {
		toSerialize["HasModifyEntriesAccess"] = o.HasModifyEntriesAccess
	}
	if !IsNil(o.HasViewEntryContentsAccess) {
		toSerialize["HasViewEntryContentsAccess"] = o.HasViewEntryContentsAccess
	}
	if !IsNil(o.CommentPrompts) {
		toSerialize["CommentPrompts"] = o.CommentPrompts
	}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.ParentId) {
		toSerialize["ParentId"] = o.ParentId
	}
	if !IsNil(o.Notes) {
		toSerialize["Notes"] = o.Notes
	}
	if !IsNil(o.Created) {
		toSerialize["Created"] = o.Created
	}
	if !IsNil(o.Modified) {
		toSerialize["Modified"] = o.Modified
	}
	if !IsNil(o.Expires) {
		toSerialize["Expires"] = o.Expires
	}
	return toSerialize, nil
}

type NullableV6CredentialGroupOutput struct {
	value *V6CredentialGroupOutput
	isSet bool
}

func (v NullableV6CredentialGroupOutput) Get() *V6CredentialGroupOutput {
	return v.value
}

func (v *NullableV6CredentialGroupOutput) Set(val *V6CredentialGroupOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableV6CredentialGroupOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableV6CredentialGroupOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV6CredentialGroupOutput(val *V6CredentialGroupOutput) *NullableV6CredentialGroupOutput {
	return &NullableV6CredentialGroupOutput{value: val, isSet: true}
}

func (v NullableV6CredentialGroupOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV6CredentialGroupOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

