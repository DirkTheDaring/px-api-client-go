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

// checks if the GetNodeTasks200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNodeTasks200ResponseDataInner{}

// GetNodeTasks200ResponseDataInner 
type GetNodeTasks200ResponseDataInner struct {
	// 
	Endtime *int64 `json:"endtime,omitempty"`
	// 
	Id *string `json:"id,omitempty"`
	// 
	Node *string `json:"node,omitempty"`
	// 
	Pid *int64 `json:"pid,omitempty"`
	// 
	Pstart *int64 `json:"pstart,omitempty"`
	// 
	Starttime *int64 `json:"starttime,omitempty"`
	// 
	Status *string `json:"status,omitempty"`
	// 
	Type *string `json:"type,omitempty"`
	// 
	Upid *string `json:"upid,omitempty"`
	// 
	User *string `json:"user,omitempty"`
}

// NewGetNodeTasks200ResponseDataInner instantiates a new GetNodeTasks200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNodeTasks200ResponseDataInner() *GetNodeTasks200ResponseDataInner {
	this := GetNodeTasks200ResponseDataInner{}
	return &this
}

// NewGetNodeTasks200ResponseDataInnerWithDefaults instantiates a new GetNodeTasks200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNodeTasks200ResponseDataInnerWithDefaults() *GetNodeTasks200ResponseDataInner {
	this := GetNodeTasks200ResponseDataInner{}
	return &this
}

// GetEndtime returns the Endtime field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetEndtime() int64 {
	if o == nil || IsNil(o.Endtime) {
		var ret int64
		return ret
	}
	return *o.Endtime
}

// GetEndtimeOk returns a tuple with the Endtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetEndtimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Endtime) {
		return nil, false
	}
	return o.Endtime, true
}

// HasEndtime returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasEndtime() bool {
	if o != nil && !IsNil(o.Endtime) {
		return true
	}

	return false
}

// SetEndtime gets a reference to the given int64 and assigns it to the Endtime field.
func (o *GetNodeTasks200ResponseDataInner) SetEndtime(v int64) {
	o.Endtime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetNodeTasks200ResponseDataInner) SetId(v string) {
	o.Id = &v
}

// GetNode returns the Node field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetNode() string {
	if o == nil || IsNil(o.Node) {
		var ret string
		return ret
	}
	return *o.Node
}

// GetNodeOk returns a tuple with the Node field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetNodeOk() (*string, bool) {
	if o == nil || IsNil(o.Node) {
		return nil, false
	}
	return o.Node, true
}

// HasNode returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasNode() bool {
	if o != nil && !IsNil(o.Node) {
		return true
	}

	return false
}

// SetNode gets a reference to the given string and assigns it to the Node field.
func (o *GetNodeTasks200ResponseDataInner) SetNode(v string) {
	o.Node = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetPid() int64 {
	if o == nil || IsNil(o.Pid) {
		var ret int64
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetPidOk() (*int64, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given int64 and assigns it to the Pid field.
func (o *GetNodeTasks200ResponseDataInner) SetPid(v int64) {
	o.Pid = &v
}

// GetPstart returns the Pstart field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetPstart() int64 {
	if o == nil || IsNil(o.Pstart) {
		var ret int64
		return ret
	}
	return *o.Pstart
}

// GetPstartOk returns a tuple with the Pstart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetPstartOk() (*int64, bool) {
	if o == nil || IsNil(o.Pstart) {
		return nil, false
	}
	return o.Pstart, true
}

// HasPstart returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasPstart() bool {
	if o != nil && !IsNil(o.Pstart) {
		return true
	}

	return false
}

// SetPstart gets a reference to the given int64 and assigns it to the Pstart field.
func (o *GetNodeTasks200ResponseDataInner) SetPstart(v int64) {
	o.Pstart = &v
}

// GetStarttime returns the Starttime field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetStarttime() int64 {
	if o == nil || IsNil(o.Starttime) {
		var ret int64
		return ret
	}
	return *o.Starttime
}

// GetStarttimeOk returns a tuple with the Starttime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetStarttimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Starttime) {
		return nil, false
	}
	return o.Starttime, true
}

// HasStarttime returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasStarttime() bool {
	if o != nil && !IsNil(o.Starttime) {
		return true
	}

	return false
}

// SetStarttime gets a reference to the given int64 and assigns it to the Starttime field.
func (o *GetNodeTasks200ResponseDataInner) SetStarttime(v int64) {
	o.Starttime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetNodeTasks200ResponseDataInner) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetNodeTasks200ResponseDataInner) SetType(v string) {
	o.Type = &v
}

// GetUpid returns the Upid field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetUpid() string {
	if o == nil || IsNil(o.Upid) {
		var ret string
		return ret
	}
	return *o.Upid
}

// GetUpidOk returns a tuple with the Upid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetUpidOk() (*string, bool) {
	if o == nil || IsNil(o.Upid) {
		return nil, false
	}
	return o.Upid, true
}

// HasUpid returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasUpid() bool {
	if o != nil && !IsNil(o.Upid) {
		return true
	}

	return false
}

// SetUpid gets a reference to the given string and assigns it to the Upid field.
func (o *GetNodeTasks200ResponseDataInner) SetUpid(v string) {
	o.Upid = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *GetNodeTasks200ResponseDataInner) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNodeTasks200ResponseDataInner) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *GetNodeTasks200ResponseDataInner) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *GetNodeTasks200ResponseDataInner) SetUser(v string) {
	o.User = &v
}

func (o GetNodeTasks200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNodeTasks200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Endtime) {
		toSerialize["endtime"] = o.Endtime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Node) {
		toSerialize["node"] = o.Node
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.Pstart) {
		toSerialize["pstart"] = o.Pstart
	}
	if !IsNil(o.Starttime) {
		toSerialize["starttime"] = o.Starttime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Upid) {
		toSerialize["upid"] = o.Upid
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableGetNodeTasks200ResponseDataInner struct {
	value *GetNodeTasks200ResponseDataInner
	isSet bool
}

func (v NullableGetNodeTasks200ResponseDataInner) Get() *GetNodeTasks200ResponseDataInner {
	return v.value
}

func (v *NullableGetNodeTasks200ResponseDataInner) Set(val *GetNodeTasks200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetNodeTasks200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetNodeTasks200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetNodeTasks200ResponseDataInner(val *GetNodeTasks200ResponseDataInner) *NullableGetNodeTasks200ResponseDataInner {
	return &NullableGetNodeTasks200ResponseDataInner{value: val, isSet: true}
}

func (v NullableGetNodeTasks200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetNodeTasks200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


