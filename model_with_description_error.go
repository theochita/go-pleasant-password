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

// checks if the WithDescriptionError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WithDescriptionError{}

// WithDescriptionError                 
type WithDescriptionError struct {
	Error *string `json:"error,omitempty"`
	ErrorDescription *string `json:"error_description,omitempty"`
}

// NewWithDescriptionError instantiates a new WithDescriptionError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWithDescriptionError() *WithDescriptionError {
	this := WithDescriptionError{}
	return &this
}

// NewWithDescriptionErrorWithDefaults instantiates a new WithDescriptionError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWithDescriptionErrorWithDefaults() *WithDescriptionError {
	this := WithDescriptionError{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *WithDescriptionError) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithDescriptionError) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *WithDescriptionError) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *WithDescriptionError) SetError(v string) {
	o.Error = &v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *WithDescriptionError) GetErrorDescription() string {
	if o == nil || IsNil(o.ErrorDescription) {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WithDescriptionError) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorDescription) {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *WithDescriptionError) HasErrorDescription() bool {
	if o != nil && !IsNil(o.ErrorDescription) {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *WithDescriptionError) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

func (o WithDescriptionError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WithDescriptionError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	if !IsNil(o.ErrorDescription) {
		toSerialize["error_description"] = o.ErrorDescription
	}
	return toSerialize, nil
}

type NullableWithDescriptionError struct {
	value *WithDescriptionError
	isSet bool
}

func (v NullableWithDescriptionError) Get() *WithDescriptionError {
	return v.value
}

func (v *NullableWithDescriptionError) Set(val *WithDescriptionError) {
	v.value = val
	v.isSet = true
}

func (v NullableWithDescriptionError) IsSet() bool {
	return v.isSet
}

func (v *NullableWithDescriptionError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWithDescriptionError(val *WithDescriptionError) *NullableWithDescriptionError {
	return &NullableWithDescriptionError{value: val, isSet: true}
}

func (v NullableWithDescriptionError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWithDescriptionError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


