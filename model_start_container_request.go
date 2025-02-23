/*
ProxMox VE API

ProxMox VE API

API version: 8.3
Contact: baldur@email.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pxapiflat

import (
	"encoding/json"
)

// checks if the StartContainerRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartContainerRequest{}

// StartContainerRequest struct for StartContainerRequest
type StartContainerRequest struct {
	// If set, enables very verbose debug log-level on start.
	Debug *bool `json:"debug,omitempty"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *bool `json:"skiplock,omitempty"`
}

// NewStartContainerRequest instantiates a new StartContainerRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartContainerRequest() *StartContainerRequest {
	this := StartContainerRequest{}
	return &this
}

// NewStartContainerRequestWithDefaults instantiates a new StartContainerRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartContainerRequestWithDefaults() *StartContainerRequest {
	this := StartContainerRequest{}
	return &this
}

// GetDebug returns the Debug field value if set, zero value otherwise.
func (o *StartContainerRequest) GetDebug() bool {
	if o == nil || IsNil(o.Debug) {
		var ret bool
		return ret
	}
	return *o.Debug
}

// GetDebugOk returns a tuple with the Debug field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartContainerRequest) GetDebugOk() (*bool, bool) {
	if o == nil || IsNil(o.Debug) {
		return nil, false
	}
	return o.Debug, true
}

// HasDebug returns a boolean if a field has been set.
func (o *StartContainerRequest) HasDebug() bool {
	if o != nil && !IsNil(o.Debug) {
		return true
	}

	return false
}

// SetDebug gets a reference to the given bool and assigns it to the Debug field.
func (o *StartContainerRequest) SetDebug(v bool) {
	o.Debug = &v
}

// GetSkiplock returns the Skiplock field value if set, zero value otherwise.
func (o *StartContainerRequest) GetSkiplock() bool {
	if o == nil || IsNil(o.Skiplock) {
		var ret bool
		return ret
	}
	return *o.Skiplock
}

// GetSkiplockOk returns a tuple with the Skiplock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartContainerRequest) GetSkiplockOk() (*bool, bool) {
	if o == nil || IsNil(o.Skiplock) {
		return nil, false
	}
	return o.Skiplock, true
}

// HasSkiplock returns a boolean if a field has been set.
func (o *StartContainerRequest) HasSkiplock() bool {
	if o != nil && !IsNil(o.Skiplock) {
		return true
	}

	return false
}

// SetSkiplock gets a reference to the given bool and assigns it to the Skiplock field.
func (o *StartContainerRequest) SetSkiplock(v bool) {
	o.Skiplock = &v
}

func (o StartContainerRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartContainerRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Debug) {
		toSerialize["debug"] = o.Debug
	}
	if !IsNil(o.Skiplock) {
		toSerialize["skiplock"] = o.Skiplock
	}
	return toSerialize, nil
}

type NullableStartContainerRequest struct {
	value *StartContainerRequest
	isSet bool
}

func (v NullableStartContainerRequest) Get() *StartContainerRequest {
	return v.value
}

func (v *NullableStartContainerRequest) Set(val *StartContainerRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStartContainerRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStartContainerRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartContainerRequest(val *StartContainerRequest) *NullableStartContainerRequest {
	return &NullableStartContainerRequest{value: val, isSet: true}
}

func (v NullableStartContainerRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartContainerRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


