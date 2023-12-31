/*
pleasant password server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V6CommentPromptResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V6CommentPromptResult{}

// V6CommentPromptResult 
type V6CommentPromptResult struct {
	AskForCommentOnViewPassword *bool `json:"AskForCommentOnViewPassword,omitempty"`
	AskForCommentOnViewOffline *bool `json:"AskForCommentOnViewOffline,omitempty"`
	AskForCommentOnModifyEntries *bool `json:"AskForCommentOnModifyEntries,omitempty"`
	AskForCommentOnMoveEntries *bool `json:"AskForCommentOnMoveEntries,omitempty"`
	AskForCommentOnMoveFolders *bool `json:"AskForCommentOnMoveFolders,omitempty"`
	AskForCommentOnModifyFolders *bool `json:"AskForCommentOnModifyFolders,omitempty"`
}

// NewV6CommentPromptResult instantiates a new V6CommentPromptResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV6CommentPromptResult() *V6CommentPromptResult {
	this := V6CommentPromptResult{}
	return &this
}

// NewV6CommentPromptResultWithDefaults instantiates a new V6CommentPromptResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV6CommentPromptResultWithDefaults() *V6CommentPromptResult {
	this := V6CommentPromptResult{}
	return &this
}

// GetAskForCommentOnViewPassword returns the AskForCommentOnViewPassword field value if set, zero value otherwise.
func (o *V6CommentPromptResult) GetAskForCommentOnViewPassword() bool {
	if o == nil || IsNil(o.AskForCommentOnViewPassword) {
		var ret bool
		return ret
	}
	return *o.AskForCommentOnViewPassword
}

// GetAskForCommentOnViewPasswordOk returns a tuple with the AskForCommentOnViewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CommentPromptResult) GetAskForCommentOnViewPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.AskForCommentOnViewPassword) {
		return nil, false
	}
	return o.AskForCommentOnViewPassword, true
}

// HasAskForCommentOnViewPassword returns a boolean if a field has been set.
func (o *V6CommentPromptResult) HasAskForCommentOnViewPassword() bool {
	if o != nil && !IsNil(o.AskForCommentOnViewPassword) {
		return true
	}

	return false
}

// SetAskForCommentOnViewPassword gets a reference to the given bool and assigns it to the AskForCommentOnViewPassword field.
func (o *V6CommentPromptResult) SetAskForCommentOnViewPassword(v bool) {
	o.AskForCommentOnViewPassword = &v
}

// GetAskForCommentOnViewOffline returns the AskForCommentOnViewOffline field value if set, zero value otherwise.
func (o *V6CommentPromptResult) GetAskForCommentOnViewOffline() bool {
	if o == nil || IsNil(o.AskForCommentOnViewOffline) {
		var ret bool
		return ret
	}
	return *o.AskForCommentOnViewOffline
}

// GetAskForCommentOnViewOfflineOk returns a tuple with the AskForCommentOnViewOffline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CommentPromptResult) GetAskForCommentOnViewOfflineOk() (*bool, bool) {
	if o == nil || IsNil(o.AskForCommentOnViewOffline) {
		return nil, false
	}
	return o.AskForCommentOnViewOffline, true
}

// HasAskForCommentOnViewOffline returns a boolean if a field has been set.
func (o *V6CommentPromptResult) HasAskForCommentOnViewOffline() bool {
	if o != nil && !IsNil(o.AskForCommentOnViewOffline) {
		return true
	}

	return false
}

// SetAskForCommentOnViewOffline gets a reference to the given bool and assigns it to the AskForCommentOnViewOffline field.
func (o *V6CommentPromptResult) SetAskForCommentOnViewOffline(v bool) {
	o.AskForCommentOnViewOffline = &v
}

// GetAskForCommentOnModifyEntries returns the AskForCommentOnModifyEntries field value if set, zero value otherwise.
func (o *V6CommentPromptResult) GetAskForCommentOnModifyEntries() bool {
	if o == nil || IsNil(o.AskForCommentOnModifyEntries) {
		var ret bool
		return ret
	}
	return *o.AskForCommentOnModifyEntries
}

// GetAskForCommentOnModifyEntriesOk returns a tuple with the AskForCommentOnModifyEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CommentPromptResult) GetAskForCommentOnModifyEntriesOk() (*bool, bool) {
	if o == nil || IsNil(o.AskForCommentOnModifyEntries) {
		return nil, false
	}
	return o.AskForCommentOnModifyEntries, true
}

// HasAskForCommentOnModifyEntries returns a boolean if a field has been set.
func (o *V6CommentPromptResult) HasAskForCommentOnModifyEntries() bool {
	if o != nil && !IsNil(o.AskForCommentOnModifyEntries) {
		return true
	}

	return false
}

// SetAskForCommentOnModifyEntries gets a reference to the given bool and assigns it to the AskForCommentOnModifyEntries field.
func (o *V6CommentPromptResult) SetAskForCommentOnModifyEntries(v bool) {
	o.AskForCommentOnModifyEntries = &v
}

// GetAskForCommentOnMoveEntries returns the AskForCommentOnMoveEntries field value if set, zero value otherwise.
func (o *V6CommentPromptResult) GetAskForCommentOnMoveEntries() bool {
	if o == nil || IsNil(o.AskForCommentOnMoveEntries) {
		var ret bool
		return ret
	}
	return *o.AskForCommentOnMoveEntries
}

// GetAskForCommentOnMoveEntriesOk returns a tuple with the AskForCommentOnMoveEntries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CommentPromptResult) GetAskForCommentOnMoveEntriesOk() (*bool, bool) {
	if o == nil || IsNil(o.AskForCommentOnMoveEntries) {
		return nil, false
	}
	return o.AskForCommentOnMoveEntries, true
}

// HasAskForCommentOnMoveEntries returns a boolean if a field has been set.
func (o *V6CommentPromptResult) HasAskForCommentOnMoveEntries() bool {
	if o != nil && !IsNil(o.AskForCommentOnMoveEntries) {
		return true
	}

	return false
}

// SetAskForCommentOnMoveEntries gets a reference to the given bool and assigns it to the AskForCommentOnMoveEntries field.
func (o *V6CommentPromptResult) SetAskForCommentOnMoveEntries(v bool) {
	o.AskForCommentOnMoveEntries = &v
}

// GetAskForCommentOnMoveFolders returns the AskForCommentOnMoveFolders field value if set, zero value otherwise.
func (o *V6CommentPromptResult) GetAskForCommentOnMoveFolders() bool {
	if o == nil || IsNil(o.AskForCommentOnMoveFolders) {
		var ret bool
		return ret
	}
	return *o.AskForCommentOnMoveFolders
}

// GetAskForCommentOnMoveFoldersOk returns a tuple with the AskForCommentOnMoveFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CommentPromptResult) GetAskForCommentOnMoveFoldersOk() (*bool, bool) {
	if o == nil || IsNil(o.AskForCommentOnMoveFolders) {
		return nil, false
	}
	return o.AskForCommentOnMoveFolders, true
}

// HasAskForCommentOnMoveFolders returns a boolean if a field has been set.
func (o *V6CommentPromptResult) HasAskForCommentOnMoveFolders() bool {
	if o != nil && !IsNil(o.AskForCommentOnMoveFolders) {
		return true
	}

	return false
}

// SetAskForCommentOnMoveFolders gets a reference to the given bool and assigns it to the AskForCommentOnMoveFolders field.
func (o *V6CommentPromptResult) SetAskForCommentOnMoveFolders(v bool) {
	o.AskForCommentOnMoveFolders = &v
}

// GetAskForCommentOnModifyFolders returns the AskForCommentOnModifyFolders field value if set, zero value otherwise.
func (o *V6CommentPromptResult) GetAskForCommentOnModifyFolders() bool {
	if o == nil || IsNil(o.AskForCommentOnModifyFolders) {
		var ret bool
		return ret
	}
	return *o.AskForCommentOnModifyFolders
}

// GetAskForCommentOnModifyFoldersOk returns a tuple with the AskForCommentOnModifyFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CommentPromptResult) GetAskForCommentOnModifyFoldersOk() (*bool, bool) {
	if o == nil || IsNil(o.AskForCommentOnModifyFolders) {
		return nil, false
	}
	return o.AskForCommentOnModifyFolders, true
}

// HasAskForCommentOnModifyFolders returns a boolean if a field has been set.
func (o *V6CommentPromptResult) HasAskForCommentOnModifyFolders() bool {
	if o != nil && !IsNil(o.AskForCommentOnModifyFolders) {
		return true
	}

	return false
}

// SetAskForCommentOnModifyFolders gets a reference to the given bool and assigns it to the AskForCommentOnModifyFolders field.
func (o *V6CommentPromptResult) SetAskForCommentOnModifyFolders(v bool) {
	o.AskForCommentOnModifyFolders = &v
}

func (o V6CommentPromptResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V6CommentPromptResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AskForCommentOnViewPassword) {
		toSerialize["AskForCommentOnViewPassword"] = o.AskForCommentOnViewPassword
	}
	if !IsNil(o.AskForCommentOnViewOffline) {
		toSerialize["AskForCommentOnViewOffline"] = o.AskForCommentOnViewOffline
	}
	if !IsNil(o.AskForCommentOnModifyEntries) {
		toSerialize["AskForCommentOnModifyEntries"] = o.AskForCommentOnModifyEntries
	}
	if !IsNil(o.AskForCommentOnMoveEntries) {
		toSerialize["AskForCommentOnMoveEntries"] = o.AskForCommentOnMoveEntries
	}
	if !IsNil(o.AskForCommentOnMoveFolders) {
		toSerialize["AskForCommentOnMoveFolders"] = o.AskForCommentOnMoveFolders
	}
	if !IsNil(o.AskForCommentOnModifyFolders) {
		toSerialize["AskForCommentOnModifyFolders"] = o.AskForCommentOnModifyFolders
	}
	return toSerialize, nil
}

type NullableV6CommentPromptResult struct {
	value *V6CommentPromptResult
	isSet bool
}

func (v NullableV6CommentPromptResult) Get() *V6CommentPromptResult {
	return v.value
}

func (v *NullableV6CommentPromptResult) Set(val *V6CommentPromptResult) {
	v.value = val
	v.isSet = true
}

func (v NullableV6CommentPromptResult) IsSet() bool {
	return v.isSet
}

func (v *NullableV6CommentPromptResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV6CommentPromptResult(val *V6CommentPromptResult) *NullableV6CommentPromptResult {
	return &NullableV6CommentPromptResult{value: val, isSet: true}
}

func (v NullableV6CommentPromptResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV6CommentPromptResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


