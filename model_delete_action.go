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

// checks if the DeleteAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteAction{}

// DeleteAction 
type DeleteAction struct {
	Action *string `json:"Action,omitempty"`
	Comment *string `json:"Comment,omitempty"`
}

// NewDeleteAction instantiates a new DeleteAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteAction() *DeleteAction {
	this := DeleteAction{}
	return &this
}

// NewDeleteActionWithDefaults instantiates a new DeleteAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteActionWithDefaults() *DeleteAction {
	this := DeleteAction{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DeleteAction) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAction) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DeleteAction) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *DeleteAction) SetAction(v string) {
	o.Action = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *DeleteAction) GetComment() string {
	if o == nil || IsNil(o.Comment) {
		var ret string
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteAction) GetCommentOk() (*string, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *DeleteAction) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given string and assigns it to the Comment field.
func (o *DeleteAction) SetComment(v string) {
	o.Comment = &v
}

func (o DeleteAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["Action"] = o.Action
	}
	if !IsNil(o.Comment) {
		toSerialize["Comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableDeleteAction struct {
	value *DeleteAction
	isSet bool
}

func (v NullableDeleteAction) Get() *DeleteAction {
	return v.value
}

func (v *NullableDeleteAction) Set(val *DeleteAction) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteAction) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteAction(val *DeleteAction) *NullableDeleteAction {
	return &NullableDeleteAction{value: val, isSet: true}
}

func (v NullableDeleteAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


