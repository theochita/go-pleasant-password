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

// checks if the V6TagResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V6TagResult{}

// V6TagResult 
type V6TagResult struct {
	Name *string `json:"Name,omitempty"`
}

// NewV6TagResult instantiates a new V6TagResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV6TagResult() *V6TagResult {
	this := V6TagResult{}
	return &this
}

// NewV6TagResultWithDefaults instantiates a new V6TagResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV6TagResultWithDefaults() *V6TagResult {
	this := V6TagResult{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V6TagResult) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6TagResult) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V6TagResult) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V6TagResult) SetName(v string) {
	o.Name = &v
}

func (o V6TagResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V6TagResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV6TagResult struct {
	value *V6TagResult
	isSet bool
}

func (v NullableV6TagResult) Get() *V6TagResult {
	return v.value
}

func (v *NullableV6TagResult) Set(val *V6TagResult) {
	v.value = val
	v.isSet = true
}

func (v NullableV6TagResult) IsSet() bool {
	return v.isSet
}

func (v *NullableV6TagResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV6TagResult(val *V6TagResult) *NullableV6TagResult {
	return &NullableV6TagResult{value: val, isSet: true}
}

func (v NullableV6TagResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV6TagResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

