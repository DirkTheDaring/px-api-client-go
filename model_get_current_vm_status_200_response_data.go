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

// checks if the GetCurrentVMStatus200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCurrentVMStatus200ResponseData{}

// GetCurrentVMStatus200ResponseData 
type GetCurrentVMStatus200ResponseData struct {
	// QEMU Guest Agent is enabled in config.
	Agent *int32 `json:"agent,omitempty"`
	// Enable a specific clipboard. If not set, depending on the display type the SPICE one will be added.
	Clipboard *string `json:"clipboard,omitempty"`
	// Maximum usable CPUs.
	Cpus *float32 `json:"cpus,omitempty"`
	// The amount of bytes the guest read from it's block devices since the guest was started. (Note: This info is not available for all storage types.)
	Diskread *int64 `json:"diskread,omitempty"`
	// The amount of bytes the guest wrote from it's block devices since the guest was started. (Note: This info is not available for all storage types.)
	Diskwrite *int64 `json:"diskwrite,omitempty"`
	// HA manager service status.
	Ha map[string]interface{} `json:"ha,omitempty"`
	// The current config lock, if any.
	Lock *string `json:"lock,omitempty"`
	// Root disk size in bytes.
	Maxdisk *int64 `json:"maxdisk,omitempty"`
	// Maximum memory in bytes.
	Maxmem *int64 `json:"maxmem,omitempty"`
	// VM (host)name.
	Name *string `json:"name,omitempty"`
	// The amount of traffic in bytes that was sent to the guest over the network since it was started.
	Netin *int64 `json:"netin,omitempty"`
	// The amount of traffic in bytes that was sent from the guest over the network since it was started.
	Netout *int64 `json:"netout,omitempty"`
	// PID of the QEMU process, if the VM is running.
	Pid *int64 `json:"pid,omitempty"`
	// VM run state from the 'query-status' QMP monitor command.
	Qmpstatus *string `json:"qmpstatus,omitempty"`
	// The currently running machine type (if running).
	RunningMachine *string `json:"running-machine,omitempty"`
	// The QEMU version the VM is currently using (if running).
	RunningQemu *string `json:"running-qemu,omitempty"`
	// QEMU VGA configuration supports spice.
	Spice *int32 `json:"spice,omitempty"`
	// QEMU process status.
	Status *string `json:"status,omitempty"`
	// The current configured tags, if any
	Tags *string `json:"tags,omitempty"`
	// Determines if the guest is a template.
	Template *int32 `json:"template,omitempty"`
	// Uptime in seconds.
	Uptime *int64 `json:"uptime,omitempty"`
	// The (unique) ID of the VM.
	Vmid *int64 `json:"vmid,omitempty"`
}

// NewGetCurrentVMStatus200ResponseData instantiates a new GetCurrentVMStatus200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCurrentVMStatus200ResponseData() *GetCurrentVMStatus200ResponseData {
	this := GetCurrentVMStatus200ResponseData{}
	return &this
}

// NewGetCurrentVMStatus200ResponseDataWithDefaults instantiates a new GetCurrentVMStatus200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCurrentVMStatus200ResponseDataWithDefaults() *GetCurrentVMStatus200ResponseData {
	this := GetCurrentVMStatus200ResponseData{}
	return &this
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetAgent() int32 {
	if o == nil || IsNil(o.Agent) {
		var ret int32
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetAgentOk() (*int32, bool) {
	if o == nil || IsNil(o.Agent) {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasAgent() bool {
	if o != nil && !IsNil(o.Agent) {
		return true
	}

	return false
}

// SetAgent gets a reference to the given int32 and assigns it to the Agent field.
func (o *GetCurrentVMStatus200ResponseData) SetAgent(v int32) {
	o.Agent = &v
}

// GetClipboard returns the Clipboard field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetClipboard() string {
	if o == nil || IsNil(o.Clipboard) {
		var ret string
		return ret
	}
	return *o.Clipboard
}

// GetClipboardOk returns a tuple with the Clipboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetClipboardOk() (*string, bool) {
	if o == nil || IsNil(o.Clipboard) {
		return nil, false
	}
	return o.Clipboard, true
}

// HasClipboard returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasClipboard() bool {
	if o != nil && !IsNil(o.Clipboard) {
		return true
	}

	return false
}

// SetClipboard gets a reference to the given string and assigns it to the Clipboard field.
func (o *GetCurrentVMStatus200ResponseData) SetClipboard(v string) {
	o.Clipboard = &v
}

// GetCpus returns the Cpus field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetCpus() float32 {
	if o == nil || IsNil(o.Cpus) {
		var ret float32
		return ret
	}
	return *o.Cpus
}

// GetCpusOk returns a tuple with the Cpus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetCpusOk() (*float32, bool) {
	if o == nil || IsNil(o.Cpus) {
		return nil, false
	}
	return o.Cpus, true
}

// HasCpus returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasCpus() bool {
	if o != nil && !IsNil(o.Cpus) {
		return true
	}

	return false
}

// SetCpus gets a reference to the given float32 and assigns it to the Cpus field.
func (o *GetCurrentVMStatus200ResponseData) SetCpus(v float32) {
	o.Cpus = &v
}

// GetDiskread returns the Diskread field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetDiskread() int64 {
	if o == nil || IsNil(o.Diskread) {
		var ret int64
		return ret
	}
	return *o.Diskread
}

// GetDiskreadOk returns a tuple with the Diskread field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetDiskreadOk() (*int64, bool) {
	if o == nil || IsNil(o.Diskread) {
		return nil, false
	}
	return o.Diskread, true
}

// HasDiskread returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasDiskread() bool {
	if o != nil && !IsNil(o.Diskread) {
		return true
	}

	return false
}

// SetDiskread gets a reference to the given int64 and assigns it to the Diskread field.
func (o *GetCurrentVMStatus200ResponseData) SetDiskread(v int64) {
	o.Diskread = &v
}

// GetDiskwrite returns the Diskwrite field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetDiskwrite() int64 {
	if o == nil || IsNil(o.Diskwrite) {
		var ret int64
		return ret
	}
	return *o.Diskwrite
}

// GetDiskwriteOk returns a tuple with the Diskwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetDiskwriteOk() (*int64, bool) {
	if o == nil || IsNil(o.Diskwrite) {
		return nil, false
	}
	return o.Diskwrite, true
}

// HasDiskwrite returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasDiskwrite() bool {
	if o != nil && !IsNil(o.Diskwrite) {
		return true
	}

	return false
}

// SetDiskwrite gets a reference to the given int64 and assigns it to the Diskwrite field.
func (o *GetCurrentVMStatus200ResponseData) SetDiskwrite(v int64) {
	o.Diskwrite = &v
}

// GetHa returns the Ha field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetHa() map[string]interface{} {
	if o == nil || IsNil(o.Ha) {
		var ret map[string]interface{}
		return ret
	}
	return o.Ha
}

// GetHaOk returns a tuple with the Ha field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetHaOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Ha) {
		return map[string]interface{}{}, false
	}
	return o.Ha, true
}

// HasHa returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasHa() bool {
	if o != nil && !IsNil(o.Ha) {
		return true
	}

	return false
}

// SetHa gets a reference to the given map[string]interface{} and assigns it to the Ha field.
func (o *GetCurrentVMStatus200ResponseData) SetHa(v map[string]interface{}) {
	o.Ha = v
}

// GetLock returns the Lock field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetLock() string {
	if o == nil || IsNil(o.Lock) {
		var ret string
		return ret
	}
	return *o.Lock
}

// GetLockOk returns a tuple with the Lock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetLockOk() (*string, bool) {
	if o == nil || IsNil(o.Lock) {
		return nil, false
	}
	return o.Lock, true
}

// HasLock returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasLock() bool {
	if o != nil && !IsNil(o.Lock) {
		return true
	}

	return false
}

// SetLock gets a reference to the given string and assigns it to the Lock field.
func (o *GetCurrentVMStatus200ResponseData) SetLock(v string) {
	o.Lock = &v
}

// GetMaxdisk returns the Maxdisk field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetMaxdisk() int64 {
	if o == nil || IsNil(o.Maxdisk) {
		var ret int64
		return ret
	}
	return *o.Maxdisk
}

// GetMaxdiskOk returns a tuple with the Maxdisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetMaxdiskOk() (*int64, bool) {
	if o == nil || IsNil(o.Maxdisk) {
		return nil, false
	}
	return o.Maxdisk, true
}

// HasMaxdisk returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasMaxdisk() bool {
	if o != nil && !IsNil(o.Maxdisk) {
		return true
	}

	return false
}

// SetMaxdisk gets a reference to the given int64 and assigns it to the Maxdisk field.
func (o *GetCurrentVMStatus200ResponseData) SetMaxdisk(v int64) {
	o.Maxdisk = &v
}

// GetMaxmem returns the Maxmem field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetMaxmem() int64 {
	if o == nil || IsNil(o.Maxmem) {
		var ret int64
		return ret
	}
	return *o.Maxmem
}

// GetMaxmemOk returns a tuple with the Maxmem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetMaxmemOk() (*int64, bool) {
	if o == nil || IsNil(o.Maxmem) {
		return nil, false
	}
	return o.Maxmem, true
}

// HasMaxmem returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasMaxmem() bool {
	if o != nil && !IsNil(o.Maxmem) {
		return true
	}

	return false
}

// SetMaxmem gets a reference to the given int64 and assigns it to the Maxmem field.
func (o *GetCurrentVMStatus200ResponseData) SetMaxmem(v int64) {
	o.Maxmem = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetCurrentVMStatus200ResponseData) SetName(v string) {
	o.Name = &v
}

// GetNetin returns the Netin field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetNetin() int64 {
	if o == nil || IsNil(o.Netin) {
		var ret int64
		return ret
	}
	return *o.Netin
}

// GetNetinOk returns a tuple with the Netin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetNetinOk() (*int64, bool) {
	if o == nil || IsNil(o.Netin) {
		return nil, false
	}
	return o.Netin, true
}

// HasNetin returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasNetin() bool {
	if o != nil && !IsNil(o.Netin) {
		return true
	}

	return false
}

// SetNetin gets a reference to the given int64 and assigns it to the Netin field.
func (o *GetCurrentVMStatus200ResponseData) SetNetin(v int64) {
	o.Netin = &v
}

// GetNetout returns the Netout field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetNetout() int64 {
	if o == nil || IsNil(o.Netout) {
		var ret int64
		return ret
	}
	return *o.Netout
}

// GetNetoutOk returns a tuple with the Netout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetNetoutOk() (*int64, bool) {
	if o == nil || IsNil(o.Netout) {
		return nil, false
	}
	return o.Netout, true
}

// HasNetout returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasNetout() bool {
	if o != nil && !IsNil(o.Netout) {
		return true
	}

	return false
}

// SetNetout gets a reference to the given int64 and assigns it to the Netout field.
func (o *GetCurrentVMStatus200ResponseData) SetNetout(v int64) {
	o.Netout = &v
}

// GetPid returns the Pid field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetPid() int64 {
	if o == nil || IsNil(o.Pid) {
		var ret int64
		return ret
	}
	return *o.Pid
}

// GetPidOk returns a tuple with the Pid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetPidOk() (*int64, bool) {
	if o == nil || IsNil(o.Pid) {
		return nil, false
	}
	return o.Pid, true
}

// HasPid returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasPid() bool {
	if o != nil && !IsNil(o.Pid) {
		return true
	}

	return false
}

// SetPid gets a reference to the given int64 and assigns it to the Pid field.
func (o *GetCurrentVMStatus200ResponseData) SetPid(v int64) {
	o.Pid = &v
}

// GetQmpstatus returns the Qmpstatus field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetQmpstatus() string {
	if o == nil || IsNil(o.Qmpstatus) {
		var ret string
		return ret
	}
	return *o.Qmpstatus
}

// GetQmpstatusOk returns a tuple with the Qmpstatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetQmpstatusOk() (*string, bool) {
	if o == nil || IsNil(o.Qmpstatus) {
		return nil, false
	}
	return o.Qmpstatus, true
}

// HasQmpstatus returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasQmpstatus() bool {
	if o != nil && !IsNil(o.Qmpstatus) {
		return true
	}

	return false
}

// SetQmpstatus gets a reference to the given string and assigns it to the Qmpstatus field.
func (o *GetCurrentVMStatus200ResponseData) SetQmpstatus(v string) {
	o.Qmpstatus = &v
}

// GetRunningMachine returns the RunningMachine field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetRunningMachine() string {
	if o == nil || IsNil(o.RunningMachine) {
		var ret string
		return ret
	}
	return *o.RunningMachine
}

// GetRunningMachineOk returns a tuple with the RunningMachine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetRunningMachineOk() (*string, bool) {
	if o == nil || IsNil(o.RunningMachine) {
		return nil, false
	}
	return o.RunningMachine, true
}

// HasRunningMachine returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasRunningMachine() bool {
	if o != nil && !IsNil(o.RunningMachine) {
		return true
	}

	return false
}

// SetRunningMachine gets a reference to the given string and assigns it to the RunningMachine field.
func (o *GetCurrentVMStatus200ResponseData) SetRunningMachine(v string) {
	o.RunningMachine = &v
}

// GetRunningQemu returns the RunningQemu field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetRunningQemu() string {
	if o == nil || IsNil(o.RunningQemu) {
		var ret string
		return ret
	}
	return *o.RunningQemu
}

// GetRunningQemuOk returns a tuple with the RunningQemu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetRunningQemuOk() (*string, bool) {
	if o == nil || IsNil(o.RunningQemu) {
		return nil, false
	}
	return o.RunningQemu, true
}

// HasRunningQemu returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasRunningQemu() bool {
	if o != nil && !IsNil(o.RunningQemu) {
		return true
	}

	return false
}

// SetRunningQemu gets a reference to the given string and assigns it to the RunningQemu field.
func (o *GetCurrentVMStatus200ResponseData) SetRunningQemu(v string) {
	o.RunningQemu = &v
}

// GetSpice returns the Spice field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetSpice() int32 {
	if o == nil || IsNil(o.Spice) {
		var ret int32
		return ret
	}
	return *o.Spice
}

// GetSpiceOk returns a tuple with the Spice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetSpiceOk() (*int32, bool) {
	if o == nil || IsNil(o.Spice) {
		return nil, false
	}
	return o.Spice, true
}

// HasSpice returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasSpice() bool {
	if o != nil && !IsNil(o.Spice) {
		return true
	}

	return false
}

// SetSpice gets a reference to the given int32 and assigns it to the Spice field.
func (o *GetCurrentVMStatus200ResponseData) SetSpice(v int32) {
	o.Spice = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetCurrentVMStatus200ResponseData) SetStatus(v string) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetTags() string {
	if o == nil || IsNil(o.Tags) {
		var ret string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetTagsOk() (*string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given string and assigns it to the Tags field.
func (o *GetCurrentVMStatus200ResponseData) SetTags(v string) {
	o.Tags = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetTemplate() int32 {
	if o == nil || IsNil(o.Template) {
		var ret int32
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetTemplateOk() (*int32, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given int32 and assigns it to the Template field.
func (o *GetCurrentVMStatus200ResponseData) SetTemplate(v int32) {
	o.Template = &v
}

// GetUptime returns the Uptime field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetUptime() int64 {
	if o == nil || IsNil(o.Uptime) {
		var ret int64
		return ret
	}
	return *o.Uptime
}

// GetUptimeOk returns a tuple with the Uptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetUptimeOk() (*int64, bool) {
	if o == nil || IsNil(o.Uptime) {
		return nil, false
	}
	return o.Uptime, true
}

// HasUptime returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasUptime() bool {
	if o != nil && !IsNil(o.Uptime) {
		return true
	}

	return false
}

// SetUptime gets a reference to the given int64 and assigns it to the Uptime field.
func (o *GetCurrentVMStatus200ResponseData) SetUptime(v int64) {
	o.Uptime = &v
}

// GetVmid returns the Vmid field value if set, zero value otherwise.
func (o *GetCurrentVMStatus200ResponseData) GetVmid() int64 {
	if o == nil || IsNil(o.Vmid) {
		var ret int64
		return ret
	}
	return *o.Vmid
}

// GetVmidOk returns a tuple with the Vmid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCurrentVMStatus200ResponseData) GetVmidOk() (*int64, bool) {
	if o == nil || IsNil(o.Vmid) {
		return nil, false
	}
	return o.Vmid, true
}

// HasVmid returns a boolean if a field has been set.
func (o *GetCurrentVMStatus200ResponseData) HasVmid() bool {
	if o != nil && !IsNil(o.Vmid) {
		return true
	}

	return false
}

// SetVmid gets a reference to the given int64 and assigns it to the Vmid field.
func (o *GetCurrentVMStatus200ResponseData) SetVmid(v int64) {
	o.Vmid = &v
}

func (o GetCurrentVMStatus200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCurrentVMStatus200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Agent) {
		toSerialize["agent"] = o.Agent
	}
	if !IsNil(o.Clipboard) {
		toSerialize["clipboard"] = o.Clipboard
	}
	if !IsNil(o.Cpus) {
		toSerialize["cpus"] = o.Cpus
	}
	if !IsNil(o.Diskread) {
		toSerialize["diskread"] = o.Diskread
	}
	if !IsNil(o.Diskwrite) {
		toSerialize["diskwrite"] = o.Diskwrite
	}
	if !IsNil(o.Ha) {
		toSerialize["ha"] = o.Ha
	}
	if !IsNil(o.Lock) {
		toSerialize["lock"] = o.Lock
	}
	if !IsNil(o.Maxdisk) {
		toSerialize["maxdisk"] = o.Maxdisk
	}
	if !IsNil(o.Maxmem) {
		toSerialize["maxmem"] = o.Maxmem
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Netin) {
		toSerialize["netin"] = o.Netin
	}
	if !IsNil(o.Netout) {
		toSerialize["netout"] = o.Netout
	}
	if !IsNil(o.Pid) {
		toSerialize["pid"] = o.Pid
	}
	if !IsNil(o.Qmpstatus) {
		toSerialize["qmpstatus"] = o.Qmpstatus
	}
	if !IsNil(o.RunningMachine) {
		toSerialize["running-machine"] = o.RunningMachine
	}
	if !IsNil(o.RunningQemu) {
		toSerialize["running-qemu"] = o.RunningQemu
	}
	if !IsNil(o.Spice) {
		toSerialize["spice"] = o.Spice
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	if !IsNil(o.Uptime) {
		toSerialize["uptime"] = o.Uptime
	}
	if !IsNil(o.Vmid) {
		toSerialize["vmid"] = o.Vmid
	}
	return toSerialize, nil
}

type NullableGetCurrentVMStatus200ResponseData struct {
	value *GetCurrentVMStatus200ResponseData
	isSet bool
}

func (v NullableGetCurrentVMStatus200ResponseData) Get() *GetCurrentVMStatus200ResponseData {
	return v.value
}

func (v *NullableGetCurrentVMStatus200ResponseData) Set(val *GetCurrentVMStatus200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCurrentVMStatus200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCurrentVMStatus200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCurrentVMStatus200ResponseData(val *GetCurrentVMStatus200ResponseData) *NullableGetCurrentVMStatus200ResponseData {
	return &NullableGetCurrentVMStatus200ResponseData{value: val, isSet: true}
}

func (v NullableGetCurrentVMStatus200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCurrentVMStatus200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


