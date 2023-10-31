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

// checks if the V6CredentialGroupSearchResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V6CredentialGroupSearchResult{}

// V6CredentialGroupSearchResult 
type V6CredentialGroupSearchResult struct {
	Id *string `json:"Id,omitempty"`
	Name *string `json:"Name,omitempty"`
	FullPath *string `json:"FullPath,omitempty"`
}

// NewV6CredentialGroupSearchResult instantiates a new V6CredentialGroupSearchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV6CredentialGroupSearchResult() *V6CredentialGroupSearchResult {
	this := V6CredentialGroupSearchResult{}
	return &this
}

// NewV6CredentialGroupSearchResultWithDefaults instantiates a new V6CredentialGroupSearchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV6CredentialGroupSearchResultWithDefaults() *V6CredentialGroupSearchResult {
	this := V6CredentialGroupSearchResult{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V6CredentialGroupSearchResult) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupSearchResult) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V6CredentialGroupSearchResult) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V6CredentialGroupSearchResult) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V6CredentialGroupSearchResult) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupSearchResult) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V6CredentialGroupSearchResult) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V6CredentialGroupSearchResult) SetName(v string) {
	o.Name = &v
}

// GetFullPath returns the FullPath field value if set, zero value otherwise.
func (o *V6CredentialGroupSearchResult) GetFullPath() string {
	if o == nil || IsNil(o.FullPath) {
		var ret string
		return ret
	}
	return *o.FullPath
}

// GetFullPathOk returns a tuple with the FullPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V6CredentialGroupSearchResult) GetFullPathOk() (*string, bool) {
	if o == nil || IsNil(o.FullPath) {
		return nil, false
	}
	return o.FullPath, true
}

// HasFullPath returns a boolean if a field has been set.
func (o *V6CredentialGroupSearchResult) HasFullPath() bool {
	if o != nil && !IsNil(o.FullPath) {
		return true
	}

	return false
}

// SetFullPath gets a reference to the given string and assigns it to the FullPath field.
func (o *V6CredentialGroupSearchResult) SetFullPath(v string) {
	o.FullPath = &v
}

func (o V6CredentialGroupSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V6CredentialGroupSearchResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["Name"] = o.Name
	}
	if !IsNil(o.FullPath) {
		toSerialize["FullPath"] = o.FullPath
	}
	return toSerialize, nil
}

type NullableV6CredentialGroupSearchResult struct {
	value *V6CredentialGroupSearchResult
	isSet bool
}

func (v NullableV6CredentialGroupSearchResult) Get() *V6CredentialGroupSearchResult {
	return v.value
}

func (v *NullableV6CredentialGroupSearchResult) Set(val *V6CredentialGroupSearchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableV6CredentialGroupSearchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableV6CredentialGroupSearchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV6CredentialGroupSearchResult(val *V6CredentialGroupSearchResult) *NullableV6CredentialGroupSearchResult {
	return &NullableV6CredentialGroupSearchResult{value: val, isSet: true}
}

func (v NullableV6CredentialGroupSearchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV6CredentialGroupSearchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


