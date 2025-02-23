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

// checks if the GetContainerSnapshots200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetContainerSnapshots200ResponseDataInner{}

// GetContainerSnapshots200ResponseDataInner 
type GetContainerSnapshots200ResponseDataInner struct {
	// Snapshot description.
	Description *string `json:"description,omitempty"`
	// Snapshot identifier. Value 'current' identifies the current VM.
	Name *string `json:"name,omitempty"`
	// Parent snapshot identifier.
	Parent *string `json:"parent,omitempty"`
	// Snapshot creation time
	Snaptime *int64 `json:"snaptime,omitempty"`
}

// NewGetContainerSnapshots200ResponseDataInner instantiates a new GetContainerSnapshots200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetContainerSnapshots200ResponseDataInner() *GetContainerSnapshots200ResponseDataInner {
	this := GetContainerSnapshots200ResponseDataInner{}
	return &this
}

// NewGetContainerSnapshots200ResponseDataInnerWithDefaults instantiates a new GetContainerSnapshots200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetContainerSnapshots200ResponseDataInnerWithDefaults() *GetContainerSnapshots200ResponseDataInner {
	this := GetContainerSnapshots200ResponseDataInner{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GetContainerSnapshots200ResponseDataInner) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerSnapshots200ResponseDataInner) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GetContainerSnapshots200ResponseDataInner) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GetContainerSnapshots200ResponseDataInner) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetContainerSnapshots200ResponseDataInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerSnapshots200ResponseDataInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetContainerSnapshots200ResponseDataInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetContainerSnapshots200ResponseDataInner) SetName(v string) {
	o.Name = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *GetContainerSnapshots200ResponseDataInner) GetParent() string {
	if o == nil || IsNil(o.Parent) {
		var ret string
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerSnapshots200ResponseDataInner) GetParentOk() (*string, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *GetContainerSnapshots200ResponseDataInner) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given string and assigns it to the Parent field.
func (o *GetContainerSnapshots200ResponseDataInner) SetParent(v string) {
	o.Parent = &v
}

// GetSnaptime returns the Snaptime field value if set, zero value otherwise.
func (o *GetContainerSnapshots200ResponseDataInner) GetSnaptime() int64 {
	if o == nil || IsNil(o.Snaptime) {
		var ret int64
		return ret
	}
	return *o.Snaptime
}

// GetSnaptimeOk returns a tuple with the Snaptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetContainerSnapshots200ResponseDataInner) GetSnaptimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Snaptime) {
		return nil, false
	}
	return o.Snaptime, true
}

// HasSnaptime returns a boolean if a field has been set.
func (o *GetContainerSnapshots200ResponseDataInner) HasSnaptime() bool {
	if o != nil && !IsNil(o.Snaptime) {
		return true
	}

	return false
}

// SetSnaptime gets a reference to the given int64 and assigns it to the Snaptime field.
func (o *GetContainerSnapshots200ResponseDataInner) SetSnaptime(v int64) {
	o.Snaptime = &v
}

func (o GetContainerSnapshots200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetContainerSnapshots200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	if !IsNil(o.Snaptime) {
		toSerialize["snaptime"] = o.Snaptime
	}
	return toSerialize, nil
}

type NullableGetContainerSnapshots200ResponseDataInner struct {
	value *GetContainerSnapshots200ResponseDataInner
	isSet bool
}

func (v NullableGetContainerSnapshots200ResponseDataInner) Get() *GetContainerSnapshots200ResponseDataInner {
	return v.value
}

func (v *NullableGetContainerSnapshots200ResponseDataInner) Set(val *GetContainerSnapshots200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetContainerSnapshots200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetContainerSnapshots200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetContainerSnapshots200ResponseDataInner(val *GetContainerSnapshots200ResponseDataInner) *NullableGetContainerSnapshots200ResponseDataInner {
	return &NullableGetContainerSnapshots200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetContainerSnapshots200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetContainerSnapshots200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


