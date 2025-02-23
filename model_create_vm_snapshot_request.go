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

// checks if the CreateVMSnapshotRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVMSnapshotRequest{}

// CreateVMSnapshotRequest struct for CreateVMSnapshotRequest
type CreateVMSnapshotRequest struct {
	// A textual description or comment.
	Description *string `json:"description,omitempty"`
	// The name of the snapshot.
	Snapname string `json:"snapname"`
	// Save the vmstate
	Vmstate *bool `json:"vmstate,omitempty"`
}

// NewCreateVMSnapshotRequest instantiates a new CreateVMSnapshotRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVMSnapshotRequest(snapname string) *CreateVMSnapshotRequest {
	this := CreateVMSnapshotRequest{}
	this.Snapname = snapname
	return &this
}

// NewCreateVMSnapshotRequestWithDefaults instantiates a new CreateVMSnapshotRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVMSnapshotRequestWithDefaults() *CreateVMSnapshotRequest {
	this := CreateVMSnapshotRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateVMSnapshotRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMSnapshotRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateVMSnapshotRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateVMSnapshotRequest) SetDescription(v string) {
	o.Description = &v
}

// GetSnapname returns the Snapname field value
func (o *CreateVMSnapshotRequest) GetSnapname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Snapname
}

// GetSnapnameOk returns a tuple with the Snapname field value
// and a boolean to check if the value has been set.
func (o *CreateVMSnapshotRequest) GetSnapnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snapname, true
}

// SetSnapname sets field value
func (o *CreateVMSnapshotRequest) SetSnapname(v string) {
	o.Snapname = v
}

// GetVmstate returns the Vmstate field value if set, zero value otherwise.
func (o *CreateVMSnapshotRequest) GetVmstate() bool {
	if o == nil || IsNil(o.Vmstate) {
		var ret bool
		return ret
	}
	return *o.Vmstate
}

// GetVmstateOk returns a tuple with the Vmstate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVMSnapshotRequest) GetVmstateOk() (*bool, bool) {
	if o == nil || IsNil(o.Vmstate) {
		return nil, false
	}
	return o.Vmstate, true
}

// HasVmstate returns a boolean if a field has been set.
func (o *CreateVMSnapshotRequest) HasVmstate() bool {
	if o != nil && !IsNil(o.Vmstate) {
		return true
	}

	return false
}

// SetVmstate gets a reference to the given bool and assigns it to the Vmstate field.
func (o *CreateVMSnapshotRequest) SetVmstate(v bool) {
	o.Vmstate = &v
}

func (o CreateVMSnapshotRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVMSnapshotRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["snapname"] = o.Snapname
	if !IsNil(o.Vmstate) {
		toSerialize["vmstate"] = o.Vmstate
	}
	return toSerialize, nil
}

type NullableCreateVMSnapshotRequest struct {
	value *CreateVMSnapshotRequest
	isSet bool
}

func (v NullableCreateVMSnapshotRequest) Get() *CreateVMSnapshotRequest {
	return v.value
}

func (v *NullableCreateVMSnapshotRequest) Set(val *CreateVMSnapshotRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVMSnapshotRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVMSnapshotRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVMSnapshotRequest(val *CreateVMSnapshotRequest) *NullableCreateVMSnapshotRequest {
	return &NullableCreateVMSnapshotRequest{value: val, isSet: true}
}

func (v NullableCreateVMSnapshotRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVMSnapshotRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


