/*
ProxMox VE API

ProxMox VE API

API version: 8.0
Contact: baldur@email.de
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pxapiflat

import (
	"encoding/json"
)

// checks if the ResizeVMDiskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResizeVMDiskRequest{}

// ResizeVMDiskRequest struct for ResizeVMDiskRequest
type ResizeVMDiskRequest struct {
	// Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications.
	Digest *string `json:"digest,omitempty"`
	// The disk you want to resize.
	Disk string `json:"disk"`
	// The new size. With the `+` sign the value is added to the actual size of the volume and without it, the value is taken as an absolute one. Shrinking disk size is not supported.
	Size string `json:"size"`
	// Ignore locks - only root is allowed to use this option.
	Skiplock *int32 `json:"skiplock,omitempty"`
}

// NewResizeVMDiskRequest instantiates a new ResizeVMDiskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResizeVMDiskRequest(disk string, size string) *ResizeVMDiskRequest {
	this := ResizeVMDiskRequest{}
	this.Disk = disk
	this.Size = size
	return &this
}

// NewResizeVMDiskRequestWithDefaults instantiates a new ResizeVMDiskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResizeVMDiskRequestWithDefaults() *ResizeVMDiskRequest {
	this := ResizeVMDiskRequest{}
	return &this
}

// GetDigest returns the Digest field value if set, zero value otherwise.
func (o *ResizeVMDiskRequest) GetDigest() string {
	if o == nil || IsNil(o.Digest) {
		var ret string
		return ret
	}
	return *o.Digest
}

// GetDigestOk returns a tuple with the Digest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResizeVMDiskRequest) GetDigestOk() (*string, bool) {
	if o == nil || IsNil(o.Digest) {
		return nil, false
	}
	return o.Digest, true
}

// HasDigest returns a boolean if a field has been set.
func (o *ResizeVMDiskRequest) HasDigest() bool {
	if o != nil && !IsNil(o.Digest) {
		return true
	}

	return false
}

// SetDigest gets a reference to the given string and assigns it to the Digest field.
func (o *ResizeVMDiskRequest) SetDigest(v string) {
	o.Digest = &v
}

// GetDisk returns the Disk field value
func (o *ResizeVMDiskRequest) GetDisk() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Disk
}

// GetDiskOk returns a tuple with the Disk field value
// and a boolean to check if the value has been set.
func (o *ResizeVMDiskRequest) GetDiskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disk, true
}

// SetDisk sets field value
func (o *ResizeVMDiskRequest) SetDisk(v string) {
	o.Disk = v
}

// GetSize returns the Size field value
func (o *ResizeVMDiskRequest) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ResizeVMDiskRequest) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ResizeVMDiskRequest) SetSize(v string) {
	o.Size = v
}

// GetSkiplock returns the Skiplock field value if set, zero value otherwise.
func (o *ResizeVMDiskRequest) GetSkiplock() int32 {
	if o == nil || IsNil(o.Skiplock) {
		var ret int32
		return ret
	}
	return *o.Skiplock
}

// GetSkiplockOk returns a tuple with the Skiplock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResizeVMDiskRequest) GetSkiplockOk() (*int32, bool) {
	if o == nil || IsNil(o.Skiplock) {
		return nil, false
	}
	return o.Skiplock, true
}

// HasSkiplock returns a boolean if a field has been set.
func (o *ResizeVMDiskRequest) HasSkiplock() bool {
	if o != nil && !IsNil(o.Skiplock) {
		return true
	}

	return false
}

// SetSkiplock gets a reference to the given int32 and assigns it to the Skiplock field.
func (o *ResizeVMDiskRequest) SetSkiplock(v int32) {
	o.Skiplock = &v
}

func (o ResizeVMDiskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResizeVMDiskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Digest) {
		toSerialize["digest"] = o.Digest
	}
	toSerialize["disk"] = o.Disk
	toSerialize["size"] = o.Size
	if !IsNil(o.Skiplock) {
		toSerialize["skiplock"] = o.Skiplock
	}
	return toSerialize, nil
}

type NullableResizeVMDiskRequest struct {
	value *ResizeVMDiskRequest
	isSet bool
}

func (v NullableResizeVMDiskRequest) Get() *ResizeVMDiskRequest {
	return v.value
}

func (v *NullableResizeVMDiskRequest) Set(val *ResizeVMDiskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableResizeVMDiskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableResizeVMDiskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResizeVMDiskRequest(val *ResizeVMDiskRequest) *NullableResizeVMDiskRequest {
	return &NullableResizeVMDiskRequest{value: val, isSet: true}
}

func (v NullableResizeVMDiskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResizeVMDiskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

