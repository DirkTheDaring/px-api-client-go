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

// checks if the GetNodeTaskLog200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNodeTaskLog200Response{}

// GetNodeTaskLog200Response struct for GetNodeTaskLog200Response
type GetNodeTaskLog200Response struct {
	Data []GetNodeTaskLog200ResponseDataInner `json:"data,omitempty"`
	Errors []string `json:"errors,omitempty"`
}

// NewGetNodeTaskLog200Response instantiates a new GetNodeTaskLog200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNodeTaskLog200Response() *GetNodeTaskLog200Response {
	this := GetNodeTaskLog200Response{}
	return &this
}

// NewGetNodeTaskLog200ResponseWithDefaults instantiates a new GetNodeTaskLog200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNodeTaskLog200ResponseWithDefaults() *GetNodeTaskLog200Response {
	this := GetNodeTaskLog200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetNodeTaskLog200Response) GetData() []GetNodeTaskLog200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetNodeTaskLog200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTaskLog200Response) GetDataOk() ([]GetNodeTaskLog200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetNodeTaskLog200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetNodeTaskLog200ResponseDataInner and assigns it to the Data field.
func (o *GetNodeTaskLog200Response) SetData(v []GetNodeTaskLog200ResponseDataInner) {
	o.Data = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetNodeTaskLog200Response) GetErrors() []string {
	if o == nil || IsNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTaskLog200Response) GetErrorsOk() ([]string, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetNodeTaskLog200Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *GetNodeTaskLog200Response) SetErrors(v []string) {
	o.Errors = v
}

func (o GetNodeTaskLog200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNodeTaskLog200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetNodeTaskLog200Response struct {
	value *GetNodeTaskLog200Response
	isSet bool
}

func (v NullableGetNodeTaskLog200Response) Get() *GetNodeTaskLog200Response {
	return v.value
}

func (v *NullableGetNodeTaskLog200Response) Set(val *GetNodeTaskLog200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNodeTaskLog200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNodeTaskLog200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNodeTaskLog200Response(val *GetNodeTaskLog200Response) *NullableGetNodeTaskLog200Response {
	return &NullableGetNodeTaskLog200Response{value: val, isSet: true}
}

func (v NullableGetNodeTaskLog200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNodeTaskLog200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

