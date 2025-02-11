# UpdateContainerConfigSyncRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | OS architecture type. | [optional] 
**Cmode** | Pointer to **string** | Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to &#39;console&#39; it tries to attach to /dev/console instead. If you set cmode to &#39;shell&#39;, it simply invokes a shell inside the container (no login). | [optional] 
**Console** | Pointer to **bool** | Attach a console device (/dev/console) to the container. | [optional] 
**Cores** | Pointer to **int64** | The number of cores assigned to the container. A container can use all available cores by default. | [optional] 
**Cpulimit** | Pointer to **float32** | Limit of CPU usage.  NOTE: If the computer has 2 CPUs, it has a total of &#39;2&#39; CPU time. Value &#39;0&#39; indicates no CPU limit. | [optional] 
**Cpuunits** | Pointer to **int64** | CPU weight for a container, will be clamped to [1, 10000] in cgroup v2. | [optional] 
**Debug** | Pointer to **bool** | Try to be more verbose. For now this only enables debug log-level on start. | [optional] 
**Delete** | Pointer to **string** | A list of settings you want to delete. | [optional] 
**Description** | Pointer to **string** | Description for the Container. Shown in the web-interface CT&#39;s summary. This is saved as comment inside the configuration file. | [optional] 
**Dev0** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev1** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev2** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev3** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev4** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev5** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev6** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev7** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev8** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev9** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev10** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev11** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev12** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev13** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev14** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev15** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev16** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev17** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev18** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev19** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev20** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev21** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev22** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev23** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev24** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev25** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev26** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev27** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev28** | Pointer to **string** | Device to pass through to the container | [optional] 
**Dev29** | Pointer to **string** | Device to pass through to the container | [optional] 
**Digest** | Pointer to **string** | Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications. | [optional] 
**Features** | Pointer to **string** | Allow containers access to advanced features. | [optional] 
**Hookscript** | Pointer to **string** | Script that will be executed during various steps in the containers lifetime. | [optional] 
**Hostname** | Pointer to **string** | Set a host name for the container. | [optional] 
**Lock** | Pointer to **string** | Lock/unlock the container. | [optional] 
**Memory** | Pointer to **int64** | Amount of RAM for the container in MB. | [optional] 
**Mp0** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp1** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp2** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp3** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp4** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp5** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp6** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp7** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp8** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp9** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp10** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp11** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp12** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp13** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp14** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp15** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp16** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp17** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp18** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp19** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp20** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp21** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp22** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp23** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp24** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp25** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp26** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp27** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp28** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp29** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Nameserver** | Pointer to **string** | Sets DNS server IP address for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver. | [optional] 
**Net0** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net1** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net2** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net3** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net4** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net5** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net6** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net7** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net8** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net9** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net10** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net11** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net12** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net13** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net14** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net15** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net16** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net17** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net18** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net19** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net20** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net21** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net22** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net23** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net24** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net25** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net26** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net27** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net28** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net29** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net30** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Net31** | Pointer to **string** | Specifies network interfaces for the container. | [optional] 
**Onboot** | Pointer to **bool** | Specifies whether a container will be started during system bootup. | [optional] 
**Ostype** | Pointer to **string** | OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/&lt;ostype&gt;.common.conf. Value &#39;unmanaged&#39; can be used to skip and OS specific setup. | [optional] 
**Protection** | Pointer to **bool** | Sets the protection flag of the container. This will prevent the CT or CT&#39;s disk remove/update operation. | [optional] 
**Revert** | Pointer to **string** | Revert a pending change. | [optional] 
**Rootfs** | Pointer to **string** | Use volume as container root. | [optional] 
**Searchdomain** | Pointer to **string** | Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver. | [optional] 
**Startup** | Pointer to **string** | Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the &#39;up&#39; or &#39;down&#39; delay in seconds, which specifies a delay to wait before the next VM is started or stopped. | [optional] 
**Swap** | Pointer to **int64** | Amount of SWAP for the container in MB. | [optional] 
**Tags** | Pointer to **string** | Tags of the Container. This is only meta information. | [optional] 
**Template** | Pointer to **bool** | Enable/disable Template. | [optional] 
**Timezone** | Pointer to **string** | Time zone to use in the container. If option isn&#39;t set, then nothing will be done. Can be set to &#39;host&#39; to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab | [optional] 
**Tty** | Pointer to **int64** | Specify the number of tty available to the container | [optional] 
**Unprivileged** | Pointer to **bool** | Makes the container run as unprivileged user. (Should not be modified manually.) | [optional] 
**Unused0** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused1** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused2** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused3** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused4** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused5** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused6** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused7** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused8** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused9** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused10** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused11** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused12** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused13** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused14** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused15** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused16** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused17** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused18** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused19** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused20** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused21** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused22** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused23** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused24** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused25** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused26** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused27** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused28** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 
**Unused29** | Pointer to **string** | Reference to unused volumes. This is used internally, and should not be modified manually. | [optional] 

## Methods

### NewUpdateContainerConfigSyncRequest

`func NewUpdateContainerConfigSyncRequest() *UpdateContainerConfigSyncRequest`

NewUpdateContainerConfigSyncRequest instantiates a new UpdateContainerConfigSyncRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContainerConfigSyncRequestWithDefaults

`func NewUpdateContainerConfigSyncRequestWithDefaults() *UpdateContainerConfigSyncRequest`

NewUpdateContainerConfigSyncRequestWithDefaults instantiates a new UpdateContainerConfigSyncRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *UpdateContainerConfigSyncRequest) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *UpdateContainerConfigSyncRequest) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *UpdateContainerConfigSyncRequest) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *UpdateContainerConfigSyncRequest) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCmode

`func (o *UpdateContainerConfigSyncRequest) GetCmode() string`

GetCmode returns the Cmode field if non-nil, zero value otherwise.

### GetCmodeOk

`func (o *UpdateContainerConfigSyncRequest) GetCmodeOk() (*string, bool)`

GetCmodeOk returns a tuple with the Cmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmode

`func (o *UpdateContainerConfigSyncRequest) SetCmode(v string)`

SetCmode sets Cmode field to given value.

### HasCmode

`func (o *UpdateContainerConfigSyncRequest) HasCmode() bool`

HasCmode returns a boolean if a field has been set.

### GetConsole

`func (o *UpdateContainerConfigSyncRequest) GetConsole() bool`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *UpdateContainerConfigSyncRequest) GetConsoleOk() (*bool, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *UpdateContainerConfigSyncRequest) SetConsole(v bool)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *UpdateContainerConfigSyncRequest) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### GetCores

`func (o *UpdateContainerConfigSyncRequest) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *UpdateContainerConfigSyncRequest) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *UpdateContainerConfigSyncRequest) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *UpdateContainerConfigSyncRequest) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpulimit

`func (o *UpdateContainerConfigSyncRequest) GetCpulimit() float32`

GetCpulimit returns the Cpulimit field if non-nil, zero value otherwise.

### GetCpulimitOk

`func (o *UpdateContainerConfigSyncRequest) GetCpulimitOk() (*float32, bool)`

GetCpulimitOk returns a tuple with the Cpulimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpulimit

`func (o *UpdateContainerConfigSyncRequest) SetCpulimit(v float32)`

SetCpulimit sets Cpulimit field to given value.

### HasCpulimit

`func (o *UpdateContainerConfigSyncRequest) HasCpulimit() bool`

HasCpulimit returns a boolean if a field has been set.

### GetCpuunits

`func (o *UpdateContainerConfigSyncRequest) GetCpuunits() int64`

GetCpuunits returns the Cpuunits field if non-nil, zero value otherwise.

### GetCpuunitsOk

`func (o *UpdateContainerConfigSyncRequest) GetCpuunitsOk() (*int64, bool)`

GetCpuunitsOk returns a tuple with the Cpuunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuunits

`func (o *UpdateContainerConfigSyncRequest) SetCpuunits(v int64)`

SetCpuunits sets Cpuunits field to given value.

### HasCpuunits

`func (o *UpdateContainerConfigSyncRequest) HasCpuunits() bool`

HasCpuunits returns a boolean if a field has been set.

### GetDebug

`func (o *UpdateContainerConfigSyncRequest) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *UpdateContainerConfigSyncRequest) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *UpdateContainerConfigSyncRequest) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *UpdateContainerConfigSyncRequest) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDelete

`func (o *UpdateContainerConfigSyncRequest) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpdateContainerConfigSyncRequest) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpdateContainerConfigSyncRequest) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpdateContainerConfigSyncRequest) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateContainerConfigSyncRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateContainerConfigSyncRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateContainerConfigSyncRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateContainerConfigSyncRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDev0

`func (o *UpdateContainerConfigSyncRequest) GetDev0() string`

GetDev0 returns the Dev0 field if non-nil, zero value otherwise.

### GetDev0Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev0Ok() (*string, bool)`

GetDev0Ok returns a tuple with the Dev0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev0

`func (o *UpdateContainerConfigSyncRequest) SetDev0(v string)`

SetDev0 sets Dev0 field to given value.

### HasDev0

`func (o *UpdateContainerConfigSyncRequest) HasDev0() bool`

HasDev0 returns a boolean if a field has been set.

### GetDev1

`func (o *UpdateContainerConfigSyncRequest) GetDev1() string`

GetDev1 returns the Dev1 field if non-nil, zero value otherwise.

### GetDev1Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev1Ok() (*string, bool)`

GetDev1Ok returns a tuple with the Dev1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev1

`func (o *UpdateContainerConfigSyncRequest) SetDev1(v string)`

SetDev1 sets Dev1 field to given value.

### HasDev1

`func (o *UpdateContainerConfigSyncRequest) HasDev1() bool`

HasDev1 returns a boolean if a field has been set.

### GetDev2

`func (o *UpdateContainerConfigSyncRequest) GetDev2() string`

GetDev2 returns the Dev2 field if non-nil, zero value otherwise.

### GetDev2Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev2Ok() (*string, bool)`

GetDev2Ok returns a tuple with the Dev2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev2

`func (o *UpdateContainerConfigSyncRequest) SetDev2(v string)`

SetDev2 sets Dev2 field to given value.

### HasDev2

`func (o *UpdateContainerConfigSyncRequest) HasDev2() bool`

HasDev2 returns a boolean if a field has been set.

### GetDev3

`func (o *UpdateContainerConfigSyncRequest) GetDev3() string`

GetDev3 returns the Dev3 field if non-nil, zero value otherwise.

### GetDev3Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev3Ok() (*string, bool)`

GetDev3Ok returns a tuple with the Dev3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev3

`func (o *UpdateContainerConfigSyncRequest) SetDev3(v string)`

SetDev3 sets Dev3 field to given value.

### HasDev3

`func (o *UpdateContainerConfigSyncRequest) HasDev3() bool`

HasDev3 returns a boolean if a field has been set.

### GetDev4

`func (o *UpdateContainerConfigSyncRequest) GetDev4() string`

GetDev4 returns the Dev4 field if non-nil, zero value otherwise.

### GetDev4Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev4Ok() (*string, bool)`

GetDev4Ok returns a tuple with the Dev4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev4

`func (o *UpdateContainerConfigSyncRequest) SetDev4(v string)`

SetDev4 sets Dev4 field to given value.

### HasDev4

`func (o *UpdateContainerConfigSyncRequest) HasDev4() bool`

HasDev4 returns a boolean if a field has been set.

### GetDev5

`func (o *UpdateContainerConfigSyncRequest) GetDev5() string`

GetDev5 returns the Dev5 field if non-nil, zero value otherwise.

### GetDev5Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev5Ok() (*string, bool)`

GetDev5Ok returns a tuple with the Dev5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev5

`func (o *UpdateContainerConfigSyncRequest) SetDev5(v string)`

SetDev5 sets Dev5 field to given value.

### HasDev5

`func (o *UpdateContainerConfigSyncRequest) HasDev5() bool`

HasDev5 returns a boolean if a field has been set.

### GetDev6

`func (o *UpdateContainerConfigSyncRequest) GetDev6() string`

GetDev6 returns the Dev6 field if non-nil, zero value otherwise.

### GetDev6Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev6Ok() (*string, bool)`

GetDev6Ok returns a tuple with the Dev6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev6

`func (o *UpdateContainerConfigSyncRequest) SetDev6(v string)`

SetDev6 sets Dev6 field to given value.

### HasDev6

`func (o *UpdateContainerConfigSyncRequest) HasDev6() bool`

HasDev6 returns a boolean if a field has been set.

### GetDev7

`func (o *UpdateContainerConfigSyncRequest) GetDev7() string`

GetDev7 returns the Dev7 field if non-nil, zero value otherwise.

### GetDev7Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev7Ok() (*string, bool)`

GetDev7Ok returns a tuple with the Dev7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev7

`func (o *UpdateContainerConfigSyncRequest) SetDev7(v string)`

SetDev7 sets Dev7 field to given value.

### HasDev7

`func (o *UpdateContainerConfigSyncRequest) HasDev7() bool`

HasDev7 returns a boolean if a field has been set.

### GetDev8

`func (o *UpdateContainerConfigSyncRequest) GetDev8() string`

GetDev8 returns the Dev8 field if non-nil, zero value otherwise.

### GetDev8Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev8Ok() (*string, bool)`

GetDev8Ok returns a tuple with the Dev8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev8

`func (o *UpdateContainerConfigSyncRequest) SetDev8(v string)`

SetDev8 sets Dev8 field to given value.

### HasDev8

`func (o *UpdateContainerConfigSyncRequest) HasDev8() bool`

HasDev8 returns a boolean if a field has been set.

### GetDev9

`func (o *UpdateContainerConfigSyncRequest) GetDev9() string`

GetDev9 returns the Dev9 field if non-nil, zero value otherwise.

### GetDev9Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev9Ok() (*string, bool)`

GetDev9Ok returns a tuple with the Dev9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev9

`func (o *UpdateContainerConfigSyncRequest) SetDev9(v string)`

SetDev9 sets Dev9 field to given value.

### HasDev9

`func (o *UpdateContainerConfigSyncRequest) HasDev9() bool`

HasDev9 returns a boolean if a field has been set.

### GetDev10

`func (o *UpdateContainerConfigSyncRequest) GetDev10() string`

GetDev10 returns the Dev10 field if non-nil, zero value otherwise.

### GetDev10Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev10Ok() (*string, bool)`

GetDev10Ok returns a tuple with the Dev10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev10

`func (o *UpdateContainerConfigSyncRequest) SetDev10(v string)`

SetDev10 sets Dev10 field to given value.

### HasDev10

`func (o *UpdateContainerConfigSyncRequest) HasDev10() bool`

HasDev10 returns a boolean if a field has been set.

### GetDev11

`func (o *UpdateContainerConfigSyncRequest) GetDev11() string`

GetDev11 returns the Dev11 field if non-nil, zero value otherwise.

### GetDev11Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev11Ok() (*string, bool)`

GetDev11Ok returns a tuple with the Dev11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev11

`func (o *UpdateContainerConfigSyncRequest) SetDev11(v string)`

SetDev11 sets Dev11 field to given value.

### HasDev11

`func (o *UpdateContainerConfigSyncRequest) HasDev11() bool`

HasDev11 returns a boolean if a field has been set.

### GetDev12

`func (o *UpdateContainerConfigSyncRequest) GetDev12() string`

GetDev12 returns the Dev12 field if non-nil, zero value otherwise.

### GetDev12Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev12Ok() (*string, bool)`

GetDev12Ok returns a tuple with the Dev12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev12

`func (o *UpdateContainerConfigSyncRequest) SetDev12(v string)`

SetDev12 sets Dev12 field to given value.

### HasDev12

`func (o *UpdateContainerConfigSyncRequest) HasDev12() bool`

HasDev12 returns a boolean if a field has been set.

### GetDev13

`func (o *UpdateContainerConfigSyncRequest) GetDev13() string`

GetDev13 returns the Dev13 field if non-nil, zero value otherwise.

### GetDev13Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev13Ok() (*string, bool)`

GetDev13Ok returns a tuple with the Dev13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev13

`func (o *UpdateContainerConfigSyncRequest) SetDev13(v string)`

SetDev13 sets Dev13 field to given value.

### HasDev13

`func (o *UpdateContainerConfigSyncRequest) HasDev13() bool`

HasDev13 returns a boolean if a field has been set.

### GetDev14

`func (o *UpdateContainerConfigSyncRequest) GetDev14() string`

GetDev14 returns the Dev14 field if non-nil, zero value otherwise.

### GetDev14Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev14Ok() (*string, bool)`

GetDev14Ok returns a tuple with the Dev14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev14

`func (o *UpdateContainerConfigSyncRequest) SetDev14(v string)`

SetDev14 sets Dev14 field to given value.

### HasDev14

`func (o *UpdateContainerConfigSyncRequest) HasDev14() bool`

HasDev14 returns a boolean if a field has been set.

### GetDev15

`func (o *UpdateContainerConfigSyncRequest) GetDev15() string`

GetDev15 returns the Dev15 field if non-nil, zero value otherwise.

### GetDev15Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev15Ok() (*string, bool)`

GetDev15Ok returns a tuple with the Dev15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev15

`func (o *UpdateContainerConfigSyncRequest) SetDev15(v string)`

SetDev15 sets Dev15 field to given value.

### HasDev15

`func (o *UpdateContainerConfigSyncRequest) HasDev15() bool`

HasDev15 returns a boolean if a field has been set.

### GetDev16

`func (o *UpdateContainerConfigSyncRequest) GetDev16() string`

GetDev16 returns the Dev16 field if non-nil, zero value otherwise.

### GetDev16Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev16Ok() (*string, bool)`

GetDev16Ok returns a tuple with the Dev16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev16

`func (o *UpdateContainerConfigSyncRequest) SetDev16(v string)`

SetDev16 sets Dev16 field to given value.

### HasDev16

`func (o *UpdateContainerConfigSyncRequest) HasDev16() bool`

HasDev16 returns a boolean if a field has been set.

### GetDev17

`func (o *UpdateContainerConfigSyncRequest) GetDev17() string`

GetDev17 returns the Dev17 field if non-nil, zero value otherwise.

### GetDev17Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev17Ok() (*string, bool)`

GetDev17Ok returns a tuple with the Dev17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev17

`func (o *UpdateContainerConfigSyncRequest) SetDev17(v string)`

SetDev17 sets Dev17 field to given value.

### HasDev17

`func (o *UpdateContainerConfigSyncRequest) HasDev17() bool`

HasDev17 returns a boolean if a field has been set.

### GetDev18

`func (o *UpdateContainerConfigSyncRequest) GetDev18() string`

GetDev18 returns the Dev18 field if non-nil, zero value otherwise.

### GetDev18Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev18Ok() (*string, bool)`

GetDev18Ok returns a tuple with the Dev18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev18

`func (o *UpdateContainerConfigSyncRequest) SetDev18(v string)`

SetDev18 sets Dev18 field to given value.

### HasDev18

`func (o *UpdateContainerConfigSyncRequest) HasDev18() bool`

HasDev18 returns a boolean if a field has been set.

### GetDev19

`func (o *UpdateContainerConfigSyncRequest) GetDev19() string`

GetDev19 returns the Dev19 field if non-nil, zero value otherwise.

### GetDev19Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev19Ok() (*string, bool)`

GetDev19Ok returns a tuple with the Dev19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev19

`func (o *UpdateContainerConfigSyncRequest) SetDev19(v string)`

SetDev19 sets Dev19 field to given value.

### HasDev19

`func (o *UpdateContainerConfigSyncRequest) HasDev19() bool`

HasDev19 returns a boolean if a field has been set.

### GetDev20

`func (o *UpdateContainerConfigSyncRequest) GetDev20() string`

GetDev20 returns the Dev20 field if non-nil, zero value otherwise.

### GetDev20Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev20Ok() (*string, bool)`

GetDev20Ok returns a tuple with the Dev20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev20

`func (o *UpdateContainerConfigSyncRequest) SetDev20(v string)`

SetDev20 sets Dev20 field to given value.

### HasDev20

`func (o *UpdateContainerConfigSyncRequest) HasDev20() bool`

HasDev20 returns a boolean if a field has been set.

### GetDev21

`func (o *UpdateContainerConfigSyncRequest) GetDev21() string`

GetDev21 returns the Dev21 field if non-nil, zero value otherwise.

### GetDev21Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev21Ok() (*string, bool)`

GetDev21Ok returns a tuple with the Dev21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev21

`func (o *UpdateContainerConfigSyncRequest) SetDev21(v string)`

SetDev21 sets Dev21 field to given value.

### HasDev21

`func (o *UpdateContainerConfigSyncRequest) HasDev21() bool`

HasDev21 returns a boolean if a field has been set.

### GetDev22

`func (o *UpdateContainerConfigSyncRequest) GetDev22() string`

GetDev22 returns the Dev22 field if non-nil, zero value otherwise.

### GetDev22Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev22Ok() (*string, bool)`

GetDev22Ok returns a tuple with the Dev22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev22

`func (o *UpdateContainerConfigSyncRequest) SetDev22(v string)`

SetDev22 sets Dev22 field to given value.

### HasDev22

`func (o *UpdateContainerConfigSyncRequest) HasDev22() bool`

HasDev22 returns a boolean if a field has been set.

### GetDev23

`func (o *UpdateContainerConfigSyncRequest) GetDev23() string`

GetDev23 returns the Dev23 field if non-nil, zero value otherwise.

### GetDev23Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev23Ok() (*string, bool)`

GetDev23Ok returns a tuple with the Dev23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev23

`func (o *UpdateContainerConfigSyncRequest) SetDev23(v string)`

SetDev23 sets Dev23 field to given value.

### HasDev23

`func (o *UpdateContainerConfigSyncRequest) HasDev23() bool`

HasDev23 returns a boolean if a field has been set.

### GetDev24

`func (o *UpdateContainerConfigSyncRequest) GetDev24() string`

GetDev24 returns the Dev24 field if non-nil, zero value otherwise.

### GetDev24Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev24Ok() (*string, bool)`

GetDev24Ok returns a tuple with the Dev24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev24

`func (o *UpdateContainerConfigSyncRequest) SetDev24(v string)`

SetDev24 sets Dev24 field to given value.

### HasDev24

`func (o *UpdateContainerConfigSyncRequest) HasDev24() bool`

HasDev24 returns a boolean if a field has been set.

### GetDev25

`func (o *UpdateContainerConfigSyncRequest) GetDev25() string`

GetDev25 returns the Dev25 field if non-nil, zero value otherwise.

### GetDev25Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev25Ok() (*string, bool)`

GetDev25Ok returns a tuple with the Dev25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev25

`func (o *UpdateContainerConfigSyncRequest) SetDev25(v string)`

SetDev25 sets Dev25 field to given value.

### HasDev25

`func (o *UpdateContainerConfigSyncRequest) HasDev25() bool`

HasDev25 returns a boolean if a field has been set.

### GetDev26

`func (o *UpdateContainerConfigSyncRequest) GetDev26() string`

GetDev26 returns the Dev26 field if non-nil, zero value otherwise.

### GetDev26Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev26Ok() (*string, bool)`

GetDev26Ok returns a tuple with the Dev26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev26

`func (o *UpdateContainerConfigSyncRequest) SetDev26(v string)`

SetDev26 sets Dev26 field to given value.

### HasDev26

`func (o *UpdateContainerConfigSyncRequest) HasDev26() bool`

HasDev26 returns a boolean if a field has been set.

### GetDev27

`func (o *UpdateContainerConfigSyncRequest) GetDev27() string`

GetDev27 returns the Dev27 field if non-nil, zero value otherwise.

### GetDev27Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev27Ok() (*string, bool)`

GetDev27Ok returns a tuple with the Dev27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev27

`func (o *UpdateContainerConfigSyncRequest) SetDev27(v string)`

SetDev27 sets Dev27 field to given value.

### HasDev27

`func (o *UpdateContainerConfigSyncRequest) HasDev27() bool`

HasDev27 returns a boolean if a field has been set.

### GetDev28

`func (o *UpdateContainerConfigSyncRequest) GetDev28() string`

GetDev28 returns the Dev28 field if non-nil, zero value otherwise.

### GetDev28Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev28Ok() (*string, bool)`

GetDev28Ok returns a tuple with the Dev28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev28

`func (o *UpdateContainerConfigSyncRequest) SetDev28(v string)`

SetDev28 sets Dev28 field to given value.

### HasDev28

`func (o *UpdateContainerConfigSyncRequest) HasDev28() bool`

HasDev28 returns a boolean if a field has been set.

### GetDev29

`func (o *UpdateContainerConfigSyncRequest) GetDev29() string`

GetDev29 returns the Dev29 field if non-nil, zero value otherwise.

### GetDev29Ok

`func (o *UpdateContainerConfigSyncRequest) GetDev29Ok() (*string, bool)`

GetDev29Ok returns a tuple with the Dev29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev29

`func (o *UpdateContainerConfigSyncRequest) SetDev29(v string)`

SetDev29 sets Dev29 field to given value.

### HasDev29

`func (o *UpdateContainerConfigSyncRequest) HasDev29() bool`

HasDev29 returns a boolean if a field has been set.

### GetDigest

`func (o *UpdateContainerConfigSyncRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *UpdateContainerConfigSyncRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *UpdateContainerConfigSyncRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *UpdateContainerConfigSyncRequest) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetFeatures

`func (o *UpdateContainerConfigSyncRequest) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *UpdateContainerConfigSyncRequest) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *UpdateContainerConfigSyncRequest) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *UpdateContainerConfigSyncRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHookscript

`func (o *UpdateContainerConfigSyncRequest) GetHookscript() string`

GetHookscript returns the Hookscript field if non-nil, zero value otherwise.

### GetHookscriptOk

`func (o *UpdateContainerConfigSyncRequest) GetHookscriptOk() (*string, bool)`

GetHookscriptOk returns a tuple with the Hookscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookscript

`func (o *UpdateContainerConfigSyncRequest) SetHookscript(v string)`

SetHookscript sets Hookscript field to given value.

### HasHookscript

`func (o *UpdateContainerConfigSyncRequest) HasHookscript() bool`

HasHookscript returns a boolean if a field has been set.

### GetHostname

`func (o *UpdateContainerConfigSyncRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateContainerConfigSyncRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateContainerConfigSyncRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *UpdateContainerConfigSyncRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLock

`func (o *UpdateContainerConfigSyncRequest) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *UpdateContainerConfigSyncRequest) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *UpdateContainerConfigSyncRequest) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *UpdateContainerConfigSyncRequest) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMemory

`func (o *UpdateContainerConfigSyncRequest) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *UpdateContainerConfigSyncRequest) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *UpdateContainerConfigSyncRequest) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *UpdateContainerConfigSyncRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMp0

`func (o *UpdateContainerConfigSyncRequest) GetMp0() string`

GetMp0 returns the Mp0 field if non-nil, zero value otherwise.

### GetMp0Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp0Ok() (*string, bool)`

GetMp0Ok returns a tuple with the Mp0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp0

`func (o *UpdateContainerConfigSyncRequest) SetMp0(v string)`

SetMp0 sets Mp0 field to given value.

### HasMp0

`func (o *UpdateContainerConfigSyncRequest) HasMp0() bool`

HasMp0 returns a boolean if a field has been set.

### GetMp1

`func (o *UpdateContainerConfigSyncRequest) GetMp1() string`

GetMp1 returns the Mp1 field if non-nil, zero value otherwise.

### GetMp1Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp1Ok() (*string, bool)`

GetMp1Ok returns a tuple with the Mp1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp1

`func (o *UpdateContainerConfigSyncRequest) SetMp1(v string)`

SetMp1 sets Mp1 field to given value.

### HasMp1

`func (o *UpdateContainerConfigSyncRequest) HasMp1() bool`

HasMp1 returns a boolean if a field has been set.

### GetMp2

`func (o *UpdateContainerConfigSyncRequest) GetMp2() string`

GetMp2 returns the Mp2 field if non-nil, zero value otherwise.

### GetMp2Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp2Ok() (*string, bool)`

GetMp2Ok returns a tuple with the Mp2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp2

`func (o *UpdateContainerConfigSyncRequest) SetMp2(v string)`

SetMp2 sets Mp2 field to given value.

### HasMp2

`func (o *UpdateContainerConfigSyncRequest) HasMp2() bool`

HasMp2 returns a boolean if a field has been set.

### GetMp3

`func (o *UpdateContainerConfigSyncRequest) GetMp3() string`

GetMp3 returns the Mp3 field if non-nil, zero value otherwise.

### GetMp3Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp3Ok() (*string, bool)`

GetMp3Ok returns a tuple with the Mp3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp3

`func (o *UpdateContainerConfigSyncRequest) SetMp3(v string)`

SetMp3 sets Mp3 field to given value.

### HasMp3

`func (o *UpdateContainerConfigSyncRequest) HasMp3() bool`

HasMp3 returns a boolean if a field has been set.

### GetMp4

`func (o *UpdateContainerConfigSyncRequest) GetMp4() string`

GetMp4 returns the Mp4 field if non-nil, zero value otherwise.

### GetMp4Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp4Ok() (*string, bool)`

GetMp4Ok returns a tuple with the Mp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4

`func (o *UpdateContainerConfigSyncRequest) SetMp4(v string)`

SetMp4 sets Mp4 field to given value.

### HasMp4

`func (o *UpdateContainerConfigSyncRequest) HasMp4() bool`

HasMp4 returns a boolean if a field has been set.

### GetMp5

`func (o *UpdateContainerConfigSyncRequest) GetMp5() string`

GetMp5 returns the Mp5 field if non-nil, zero value otherwise.

### GetMp5Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp5Ok() (*string, bool)`

GetMp5Ok returns a tuple with the Mp5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp5

`func (o *UpdateContainerConfigSyncRequest) SetMp5(v string)`

SetMp5 sets Mp5 field to given value.

### HasMp5

`func (o *UpdateContainerConfigSyncRequest) HasMp5() bool`

HasMp5 returns a boolean if a field has been set.

### GetMp6

`func (o *UpdateContainerConfigSyncRequest) GetMp6() string`

GetMp6 returns the Mp6 field if non-nil, zero value otherwise.

### GetMp6Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp6Ok() (*string, bool)`

GetMp6Ok returns a tuple with the Mp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp6

`func (o *UpdateContainerConfigSyncRequest) SetMp6(v string)`

SetMp6 sets Mp6 field to given value.

### HasMp6

`func (o *UpdateContainerConfigSyncRequest) HasMp6() bool`

HasMp6 returns a boolean if a field has been set.

### GetMp7

`func (o *UpdateContainerConfigSyncRequest) GetMp7() string`

GetMp7 returns the Mp7 field if non-nil, zero value otherwise.

### GetMp7Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp7Ok() (*string, bool)`

GetMp7Ok returns a tuple with the Mp7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp7

`func (o *UpdateContainerConfigSyncRequest) SetMp7(v string)`

SetMp7 sets Mp7 field to given value.

### HasMp7

`func (o *UpdateContainerConfigSyncRequest) HasMp7() bool`

HasMp7 returns a boolean if a field has been set.

### GetMp8

`func (o *UpdateContainerConfigSyncRequest) GetMp8() string`

GetMp8 returns the Mp8 field if non-nil, zero value otherwise.

### GetMp8Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp8Ok() (*string, bool)`

GetMp8Ok returns a tuple with the Mp8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp8

`func (o *UpdateContainerConfigSyncRequest) SetMp8(v string)`

SetMp8 sets Mp8 field to given value.

### HasMp8

`func (o *UpdateContainerConfigSyncRequest) HasMp8() bool`

HasMp8 returns a boolean if a field has been set.

### GetMp9

`func (o *UpdateContainerConfigSyncRequest) GetMp9() string`

GetMp9 returns the Mp9 field if non-nil, zero value otherwise.

### GetMp9Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp9Ok() (*string, bool)`

GetMp9Ok returns a tuple with the Mp9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp9

`func (o *UpdateContainerConfigSyncRequest) SetMp9(v string)`

SetMp9 sets Mp9 field to given value.

### HasMp9

`func (o *UpdateContainerConfigSyncRequest) HasMp9() bool`

HasMp9 returns a boolean if a field has been set.

### GetMp10

`func (o *UpdateContainerConfigSyncRequest) GetMp10() string`

GetMp10 returns the Mp10 field if non-nil, zero value otherwise.

### GetMp10Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp10Ok() (*string, bool)`

GetMp10Ok returns a tuple with the Mp10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp10

`func (o *UpdateContainerConfigSyncRequest) SetMp10(v string)`

SetMp10 sets Mp10 field to given value.

### HasMp10

`func (o *UpdateContainerConfigSyncRequest) HasMp10() bool`

HasMp10 returns a boolean if a field has been set.

### GetMp11

`func (o *UpdateContainerConfigSyncRequest) GetMp11() string`

GetMp11 returns the Mp11 field if non-nil, zero value otherwise.

### GetMp11Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp11Ok() (*string, bool)`

GetMp11Ok returns a tuple with the Mp11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp11

`func (o *UpdateContainerConfigSyncRequest) SetMp11(v string)`

SetMp11 sets Mp11 field to given value.

### HasMp11

`func (o *UpdateContainerConfigSyncRequest) HasMp11() bool`

HasMp11 returns a boolean if a field has been set.

### GetMp12

`func (o *UpdateContainerConfigSyncRequest) GetMp12() string`

GetMp12 returns the Mp12 field if non-nil, zero value otherwise.

### GetMp12Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp12Ok() (*string, bool)`

GetMp12Ok returns a tuple with the Mp12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp12

`func (o *UpdateContainerConfigSyncRequest) SetMp12(v string)`

SetMp12 sets Mp12 field to given value.

### HasMp12

`func (o *UpdateContainerConfigSyncRequest) HasMp12() bool`

HasMp12 returns a boolean if a field has been set.

### GetMp13

`func (o *UpdateContainerConfigSyncRequest) GetMp13() string`

GetMp13 returns the Mp13 field if non-nil, zero value otherwise.

### GetMp13Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp13Ok() (*string, bool)`

GetMp13Ok returns a tuple with the Mp13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp13

`func (o *UpdateContainerConfigSyncRequest) SetMp13(v string)`

SetMp13 sets Mp13 field to given value.

### HasMp13

`func (o *UpdateContainerConfigSyncRequest) HasMp13() bool`

HasMp13 returns a boolean if a field has been set.

### GetMp14

`func (o *UpdateContainerConfigSyncRequest) GetMp14() string`

GetMp14 returns the Mp14 field if non-nil, zero value otherwise.

### GetMp14Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp14Ok() (*string, bool)`

GetMp14Ok returns a tuple with the Mp14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp14

`func (o *UpdateContainerConfigSyncRequest) SetMp14(v string)`

SetMp14 sets Mp14 field to given value.

### HasMp14

`func (o *UpdateContainerConfigSyncRequest) HasMp14() bool`

HasMp14 returns a boolean if a field has been set.

### GetMp15

`func (o *UpdateContainerConfigSyncRequest) GetMp15() string`

GetMp15 returns the Mp15 field if non-nil, zero value otherwise.

### GetMp15Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp15Ok() (*string, bool)`

GetMp15Ok returns a tuple with the Mp15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp15

`func (o *UpdateContainerConfigSyncRequest) SetMp15(v string)`

SetMp15 sets Mp15 field to given value.

### HasMp15

`func (o *UpdateContainerConfigSyncRequest) HasMp15() bool`

HasMp15 returns a boolean if a field has been set.

### GetMp16

`func (o *UpdateContainerConfigSyncRequest) GetMp16() string`

GetMp16 returns the Mp16 field if non-nil, zero value otherwise.

### GetMp16Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp16Ok() (*string, bool)`

GetMp16Ok returns a tuple with the Mp16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp16

`func (o *UpdateContainerConfigSyncRequest) SetMp16(v string)`

SetMp16 sets Mp16 field to given value.

### HasMp16

`func (o *UpdateContainerConfigSyncRequest) HasMp16() bool`

HasMp16 returns a boolean if a field has been set.

### GetMp17

`func (o *UpdateContainerConfigSyncRequest) GetMp17() string`

GetMp17 returns the Mp17 field if non-nil, zero value otherwise.

### GetMp17Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp17Ok() (*string, bool)`

GetMp17Ok returns a tuple with the Mp17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp17

`func (o *UpdateContainerConfigSyncRequest) SetMp17(v string)`

SetMp17 sets Mp17 field to given value.

### HasMp17

`func (o *UpdateContainerConfigSyncRequest) HasMp17() bool`

HasMp17 returns a boolean if a field has been set.

### GetMp18

`func (o *UpdateContainerConfigSyncRequest) GetMp18() string`

GetMp18 returns the Mp18 field if non-nil, zero value otherwise.

### GetMp18Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp18Ok() (*string, bool)`

GetMp18Ok returns a tuple with the Mp18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp18

`func (o *UpdateContainerConfigSyncRequest) SetMp18(v string)`

SetMp18 sets Mp18 field to given value.

### HasMp18

`func (o *UpdateContainerConfigSyncRequest) HasMp18() bool`

HasMp18 returns a boolean if a field has been set.

### GetMp19

`func (o *UpdateContainerConfigSyncRequest) GetMp19() string`

GetMp19 returns the Mp19 field if non-nil, zero value otherwise.

### GetMp19Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp19Ok() (*string, bool)`

GetMp19Ok returns a tuple with the Mp19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp19

`func (o *UpdateContainerConfigSyncRequest) SetMp19(v string)`

SetMp19 sets Mp19 field to given value.

### HasMp19

`func (o *UpdateContainerConfigSyncRequest) HasMp19() bool`

HasMp19 returns a boolean if a field has been set.

### GetMp20

`func (o *UpdateContainerConfigSyncRequest) GetMp20() string`

GetMp20 returns the Mp20 field if non-nil, zero value otherwise.

### GetMp20Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp20Ok() (*string, bool)`

GetMp20Ok returns a tuple with the Mp20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp20

`func (o *UpdateContainerConfigSyncRequest) SetMp20(v string)`

SetMp20 sets Mp20 field to given value.

### HasMp20

`func (o *UpdateContainerConfigSyncRequest) HasMp20() bool`

HasMp20 returns a boolean if a field has been set.

### GetMp21

`func (o *UpdateContainerConfigSyncRequest) GetMp21() string`

GetMp21 returns the Mp21 field if non-nil, zero value otherwise.

### GetMp21Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp21Ok() (*string, bool)`

GetMp21Ok returns a tuple with the Mp21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp21

`func (o *UpdateContainerConfigSyncRequest) SetMp21(v string)`

SetMp21 sets Mp21 field to given value.

### HasMp21

`func (o *UpdateContainerConfigSyncRequest) HasMp21() bool`

HasMp21 returns a boolean if a field has been set.

### GetMp22

`func (o *UpdateContainerConfigSyncRequest) GetMp22() string`

GetMp22 returns the Mp22 field if non-nil, zero value otherwise.

### GetMp22Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp22Ok() (*string, bool)`

GetMp22Ok returns a tuple with the Mp22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp22

`func (o *UpdateContainerConfigSyncRequest) SetMp22(v string)`

SetMp22 sets Mp22 field to given value.

### HasMp22

`func (o *UpdateContainerConfigSyncRequest) HasMp22() bool`

HasMp22 returns a boolean if a field has been set.

### GetMp23

`func (o *UpdateContainerConfigSyncRequest) GetMp23() string`

GetMp23 returns the Mp23 field if non-nil, zero value otherwise.

### GetMp23Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp23Ok() (*string, bool)`

GetMp23Ok returns a tuple with the Mp23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp23

`func (o *UpdateContainerConfigSyncRequest) SetMp23(v string)`

SetMp23 sets Mp23 field to given value.

### HasMp23

`func (o *UpdateContainerConfigSyncRequest) HasMp23() bool`

HasMp23 returns a boolean if a field has been set.

### GetMp24

`func (o *UpdateContainerConfigSyncRequest) GetMp24() string`

GetMp24 returns the Mp24 field if non-nil, zero value otherwise.

### GetMp24Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp24Ok() (*string, bool)`

GetMp24Ok returns a tuple with the Mp24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp24

`func (o *UpdateContainerConfigSyncRequest) SetMp24(v string)`

SetMp24 sets Mp24 field to given value.

### HasMp24

`func (o *UpdateContainerConfigSyncRequest) HasMp24() bool`

HasMp24 returns a boolean if a field has been set.

### GetMp25

`func (o *UpdateContainerConfigSyncRequest) GetMp25() string`

GetMp25 returns the Mp25 field if non-nil, zero value otherwise.

### GetMp25Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp25Ok() (*string, bool)`

GetMp25Ok returns a tuple with the Mp25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp25

`func (o *UpdateContainerConfigSyncRequest) SetMp25(v string)`

SetMp25 sets Mp25 field to given value.

### HasMp25

`func (o *UpdateContainerConfigSyncRequest) HasMp25() bool`

HasMp25 returns a boolean if a field has been set.

### GetMp26

`func (o *UpdateContainerConfigSyncRequest) GetMp26() string`

GetMp26 returns the Mp26 field if non-nil, zero value otherwise.

### GetMp26Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp26Ok() (*string, bool)`

GetMp26Ok returns a tuple with the Mp26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp26

`func (o *UpdateContainerConfigSyncRequest) SetMp26(v string)`

SetMp26 sets Mp26 field to given value.

### HasMp26

`func (o *UpdateContainerConfigSyncRequest) HasMp26() bool`

HasMp26 returns a boolean if a field has been set.

### GetMp27

`func (o *UpdateContainerConfigSyncRequest) GetMp27() string`

GetMp27 returns the Mp27 field if non-nil, zero value otherwise.

### GetMp27Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp27Ok() (*string, bool)`

GetMp27Ok returns a tuple with the Mp27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp27

`func (o *UpdateContainerConfigSyncRequest) SetMp27(v string)`

SetMp27 sets Mp27 field to given value.

### HasMp27

`func (o *UpdateContainerConfigSyncRequest) HasMp27() bool`

HasMp27 returns a boolean if a field has been set.

### GetMp28

`func (o *UpdateContainerConfigSyncRequest) GetMp28() string`

GetMp28 returns the Mp28 field if non-nil, zero value otherwise.

### GetMp28Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp28Ok() (*string, bool)`

GetMp28Ok returns a tuple with the Mp28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp28

`func (o *UpdateContainerConfigSyncRequest) SetMp28(v string)`

SetMp28 sets Mp28 field to given value.

### HasMp28

`func (o *UpdateContainerConfigSyncRequest) HasMp28() bool`

HasMp28 returns a boolean if a field has been set.

### GetMp29

`func (o *UpdateContainerConfigSyncRequest) GetMp29() string`

GetMp29 returns the Mp29 field if non-nil, zero value otherwise.

### GetMp29Ok

`func (o *UpdateContainerConfigSyncRequest) GetMp29Ok() (*string, bool)`

GetMp29Ok returns a tuple with the Mp29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp29

`func (o *UpdateContainerConfigSyncRequest) SetMp29(v string)`

SetMp29 sets Mp29 field to given value.

### HasMp29

`func (o *UpdateContainerConfigSyncRequest) HasMp29() bool`

HasMp29 returns a boolean if a field has been set.

### GetNameserver

`func (o *UpdateContainerConfigSyncRequest) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *UpdateContainerConfigSyncRequest) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *UpdateContainerConfigSyncRequest) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *UpdateContainerConfigSyncRequest) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetNet0

`func (o *UpdateContainerConfigSyncRequest) GetNet0() string`

GetNet0 returns the Net0 field if non-nil, zero value otherwise.

### GetNet0Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet0Ok() (*string, bool)`

GetNet0Ok returns a tuple with the Net0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet0

`func (o *UpdateContainerConfigSyncRequest) SetNet0(v string)`

SetNet0 sets Net0 field to given value.

### HasNet0

`func (o *UpdateContainerConfigSyncRequest) HasNet0() bool`

HasNet0 returns a boolean if a field has been set.

### GetNet1

`func (o *UpdateContainerConfigSyncRequest) GetNet1() string`

GetNet1 returns the Net1 field if non-nil, zero value otherwise.

### GetNet1Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet1Ok() (*string, bool)`

GetNet1Ok returns a tuple with the Net1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet1

`func (o *UpdateContainerConfigSyncRequest) SetNet1(v string)`

SetNet1 sets Net1 field to given value.

### HasNet1

`func (o *UpdateContainerConfigSyncRequest) HasNet1() bool`

HasNet1 returns a boolean if a field has been set.

### GetNet2

`func (o *UpdateContainerConfigSyncRequest) GetNet2() string`

GetNet2 returns the Net2 field if non-nil, zero value otherwise.

### GetNet2Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet2Ok() (*string, bool)`

GetNet2Ok returns a tuple with the Net2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet2

`func (o *UpdateContainerConfigSyncRequest) SetNet2(v string)`

SetNet2 sets Net2 field to given value.

### HasNet2

`func (o *UpdateContainerConfigSyncRequest) HasNet2() bool`

HasNet2 returns a boolean if a field has been set.

### GetNet3

`func (o *UpdateContainerConfigSyncRequest) GetNet3() string`

GetNet3 returns the Net3 field if non-nil, zero value otherwise.

### GetNet3Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet3Ok() (*string, bool)`

GetNet3Ok returns a tuple with the Net3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet3

`func (o *UpdateContainerConfigSyncRequest) SetNet3(v string)`

SetNet3 sets Net3 field to given value.

### HasNet3

`func (o *UpdateContainerConfigSyncRequest) HasNet3() bool`

HasNet3 returns a boolean if a field has been set.

### GetNet4

`func (o *UpdateContainerConfigSyncRequest) GetNet4() string`

GetNet4 returns the Net4 field if non-nil, zero value otherwise.

### GetNet4Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet4Ok() (*string, bool)`

GetNet4Ok returns a tuple with the Net4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet4

`func (o *UpdateContainerConfigSyncRequest) SetNet4(v string)`

SetNet4 sets Net4 field to given value.

### HasNet4

`func (o *UpdateContainerConfigSyncRequest) HasNet4() bool`

HasNet4 returns a boolean if a field has been set.

### GetNet5

`func (o *UpdateContainerConfigSyncRequest) GetNet5() string`

GetNet5 returns the Net5 field if non-nil, zero value otherwise.

### GetNet5Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet5Ok() (*string, bool)`

GetNet5Ok returns a tuple with the Net5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet5

`func (o *UpdateContainerConfigSyncRequest) SetNet5(v string)`

SetNet5 sets Net5 field to given value.

### HasNet5

`func (o *UpdateContainerConfigSyncRequest) HasNet5() bool`

HasNet5 returns a boolean if a field has been set.

### GetNet6

`func (o *UpdateContainerConfigSyncRequest) GetNet6() string`

GetNet6 returns the Net6 field if non-nil, zero value otherwise.

### GetNet6Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet6Ok() (*string, bool)`

GetNet6Ok returns a tuple with the Net6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet6

`func (o *UpdateContainerConfigSyncRequest) SetNet6(v string)`

SetNet6 sets Net6 field to given value.

### HasNet6

`func (o *UpdateContainerConfigSyncRequest) HasNet6() bool`

HasNet6 returns a boolean if a field has been set.

### GetNet7

`func (o *UpdateContainerConfigSyncRequest) GetNet7() string`

GetNet7 returns the Net7 field if non-nil, zero value otherwise.

### GetNet7Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet7Ok() (*string, bool)`

GetNet7Ok returns a tuple with the Net7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet7

`func (o *UpdateContainerConfigSyncRequest) SetNet7(v string)`

SetNet7 sets Net7 field to given value.

### HasNet7

`func (o *UpdateContainerConfigSyncRequest) HasNet7() bool`

HasNet7 returns a boolean if a field has been set.

### GetNet8

`func (o *UpdateContainerConfigSyncRequest) GetNet8() string`

GetNet8 returns the Net8 field if non-nil, zero value otherwise.

### GetNet8Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet8Ok() (*string, bool)`

GetNet8Ok returns a tuple with the Net8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet8

`func (o *UpdateContainerConfigSyncRequest) SetNet8(v string)`

SetNet8 sets Net8 field to given value.

### HasNet8

`func (o *UpdateContainerConfigSyncRequest) HasNet8() bool`

HasNet8 returns a boolean if a field has been set.

### GetNet9

`func (o *UpdateContainerConfigSyncRequest) GetNet9() string`

GetNet9 returns the Net9 field if non-nil, zero value otherwise.

### GetNet9Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet9Ok() (*string, bool)`

GetNet9Ok returns a tuple with the Net9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet9

`func (o *UpdateContainerConfigSyncRequest) SetNet9(v string)`

SetNet9 sets Net9 field to given value.

### HasNet9

`func (o *UpdateContainerConfigSyncRequest) HasNet9() bool`

HasNet9 returns a boolean if a field has been set.

### GetNet10

`func (o *UpdateContainerConfigSyncRequest) GetNet10() string`

GetNet10 returns the Net10 field if non-nil, zero value otherwise.

### GetNet10Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet10Ok() (*string, bool)`

GetNet10Ok returns a tuple with the Net10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet10

`func (o *UpdateContainerConfigSyncRequest) SetNet10(v string)`

SetNet10 sets Net10 field to given value.

### HasNet10

`func (o *UpdateContainerConfigSyncRequest) HasNet10() bool`

HasNet10 returns a boolean if a field has been set.

### GetNet11

`func (o *UpdateContainerConfigSyncRequest) GetNet11() string`

GetNet11 returns the Net11 field if non-nil, zero value otherwise.

### GetNet11Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet11Ok() (*string, bool)`

GetNet11Ok returns a tuple with the Net11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet11

`func (o *UpdateContainerConfigSyncRequest) SetNet11(v string)`

SetNet11 sets Net11 field to given value.

### HasNet11

`func (o *UpdateContainerConfigSyncRequest) HasNet11() bool`

HasNet11 returns a boolean if a field has been set.

### GetNet12

`func (o *UpdateContainerConfigSyncRequest) GetNet12() string`

GetNet12 returns the Net12 field if non-nil, zero value otherwise.

### GetNet12Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet12Ok() (*string, bool)`

GetNet12Ok returns a tuple with the Net12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet12

`func (o *UpdateContainerConfigSyncRequest) SetNet12(v string)`

SetNet12 sets Net12 field to given value.

### HasNet12

`func (o *UpdateContainerConfigSyncRequest) HasNet12() bool`

HasNet12 returns a boolean if a field has been set.

### GetNet13

`func (o *UpdateContainerConfigSyncRequest) GetNet13() string`

GetNet13 returns the Net13 field if non-nil, zero value otherwise.

### GetNet13Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet13Ok() (*string, bool)`

GetNet13Ok returns a tuple with the Net13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet13

`func (o *UpdateContainerConfigSyncRequest) SetNet13(v string)`

SetNet13 sets Net13 field to given value.

### HasNet13

`func (o *UpdateContainerConfigSyncRequest) HasNet13() bool`

HasNet13 returns a boolean if a field has been set.

### GetNet14

`func (o *UpdateContainerConfigSyncRequest) GetNet14() string`

GetNet14 returns the Net14 field if non-nil, zero value otherwise.

### GetNet14Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet14Ok() (*string, bool)`

GetNet14Ok returns a tuple with the Net14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet14

`func (o *UpdateContainerConfigSyncRequest) SetNet14(v string)`

SetNet14 sets Net14 field to given value.

### HasNet14

`func (o *UpdateContainerConfigSyncRequest) HasNet14() bool`

HasNet14 returns a boolean if a field has been set.

### GetNet15

`func (o *UpdateContainerConfigSyncRequest) GetNet15() string`

GetNet15 returns the Net15 field if non-nil, zero value otherwise.

### GetNet15Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet15Ok() (*string, bool)`

GetNet15Ok returns a tuple with the Net15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet15

`func (o *UpdateContainerConfigSyncRequest) SetNet15(v string)`

SetNet15 sets Net15 field to given value.

### HasNet15

`func (o *UpdateContainerConfigSyncRequest) HasNet15() bool`

HasNet15 returns a boolean if a field has been set.

### GetNet16

`func (o *UpdateContainerConfigSyncRequest) GetNet16() string`

GetNet16 returns the Net16 field if non-nil, zero value otherwise.

### GetNet16Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet16Ok() (*string, bool)`

GetNet16Ok returns a tuple with the Net16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet16

`func (o *UpdateContainerConfigSyncRequest) SetNet16(v string)`

SetNet16 sets Net16 field to given value.

### HasNet16

`func (o *UpdateContainerConfigSyncRequest) HasNet16() bool`

HasNet16 returns a boolean if a field has been set.

### GetNet17

`func (o *UpdateContainerConfigSyncRequest) GetNet17() string`

GetNet17 returns the Net17 field if non-nil, zero value otherwise.

### GetNet17Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet17Ok() (*string, bool)`

GetNet17Ok returns a tuple with the Net17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet17

`func (o *UpdateContainerConfigSyncRequest) SetNet17(v string)`

SetNet17 sets Net17 field to given value.

### HasNet17

`func (o *UpdateContainerConfigSyncRequest) HasNet17() bool`

HasNet17 returns a boolean if a field has been set.

### GetNet18

`func (o *UpdateContainerConfigSyncRequest) GetNet18() string`

GetNet18 returns the Net18 field if non-nil, zero value otherwise.

### GetNet18Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet18Ok() (*string, bool)`

GetNet18Ok returns a tuple with the Net18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet18

`func (o *UpdateContainerConfigSyncRequest) SetNet18(v string)`

SetNet18 sets Net18 field to given value.

### HasNet18

`func (o *UpdateContainerConfigSyncRequest) HasNet18() bool`

HasNet18 returns a boolean if a field has been set.

### GetNet19

`func (o *UpdateContainerConfigSyncRequest) GetNet19() string`

GetNet19 returns the Net19 field if non-nil, zero value otherwise.

### GetNet19Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet19Ok() (*string, bool)`

GetNet19Ok returns a tuple with the Net19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet19

`func (o *UpdateContainerConfigSyncRequest) SetNet19(v string)`

SetNet19 sets Net19 field to given value.

### HasNet19

`func (o *UpdateContainerConfigSyncRequest) HasNet19() bool`

HasNet19 returns a boolean if a field has been set.

### GetNet20

`func (o *UpdateContainerConfigSyncRequest) GetNet20() string`

GetNet20 returns the Net20 field if non-nil, zero value otherwise.

### GetNet20Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet20Ok() (*string, bool)`

GetNet20Ok returns a tuple with the Net20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet20

`func (o *UpdateContainerConfigSyncRequest) SetNet20(v string)`

SetNet20 sets Net20 field to given value.

### HasNet20

`func (o *UpdateContainerConfigSyncRequest) HasNet20() bool`

HasNet20 returns a boolean if a field has been set.

### GetNet21

`func (o *UpdateContainerConfigSyncRequest) GetNet21() string`

GetNet21 returns the Net21 field if non-nil, zero value otherwise.

### GetNet21Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet21Ok() (*string, bool)`

GetNet21Ok returns a tuple with the Net21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet21

`func (o *UpdateContainerConfigSyncRequest) SetNet21(v string)`

SetNet21 sets Net21 field to given value.

### HasNet21

`func (o *UpdateContainerConfigSyncRequest) HasNet21() bool`

HasNet21 returns a boolean if a field has been set.

### GetNet22

`func (o *UpdateContainerConfigSyncRequest) GetNet22() string`

GetNet22 returns the Net22 field if non-nil, zero value otherwise.

### GetNet22Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet22Ok() (*string, bool)`

GetNet22Ok returns a tuple with the Net22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet22

`func (o *UpdateContainerConfigSyncRequest) SetNet22(v string)`

SetNet22 sets Net22 field to given value.

### HasNet22

`func (o *UpdateContainerConfigSyncRequest) HasNet22() bool`

HasNet22 returns a boolean if a field has been set.

### GetNet23

`func (o *UpdateContainerConfigSyncRequest) GetNet23() string`

GetNet23 returns the Net23 field if non-nil, zero value otherwise.

### GetNet23Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet23Ok() (*string, bool)`

GetNet23Ok returns a tuple with the Net23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet23

`func (o *UpdateContainerConfigSyncRequest) SetNet23(v string)`

SetNet23 sets Net23 field to given value.

### HasNet23

`func (o *UpdateContainerConfigSyncRequest) HasNet23() bool`

HasNet23 returns a boolean if a field has been set.

### GetNet24

`func (o *UpdateContainerConfigSyncRequest) GetNet24() string`

GetNet24 returns the Net24 field if non-nil, zero value otherwise.

### GetNet24Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet24Ok() (*string, bool)`

GetNet24Ok returns a tuple with the Net24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet24

`func (o *UpdateContainerConfigSyncRequest) SetNet24(v string)`

SetNet24 sets Net24 field to given value.

### HasNet24

`func (o *UpdateContainerConfigSyncRequest) HasNet24() bool`

HasNet24 returns a boolean if a field has been set.

### GetNet25

`func (o *UpdateContainerConfigSyncRequest) GetNet25() string`

GetNet25 returns the Net25 field if non-nil, zero value otherwise.

### GetNet25Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet25Ok() (*string, bool)`

GetNet25Ok returns a tuple with the Net25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet25

`func (o *UpdateContainerConfigSyncRequest) SetNet25(v string)`

SetNet25 sets Net25 field to given value.

### HasNet25

`func (o *UpdateContainerConfigSyncRequest) HasNet25() bool`

HasNet25 returns a boolean if a field has been set.

### GetNet26

`func (o *UpdateContainerConfigSyncRequest) GetNet26() string`

GetNet26 returns the Net26 field if non-nil, zero value otherwise.

### GetNet26Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet26Ok() (*string, bool)`

GetNet26Ok returns a tuple with the Net26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet26

`func (o *UpdateContainerConfigSyncRequest) SetNet26(v string)`

SetNet26 sets Net26 field to given value.

### HasNet26

`func (o *UpdateContainerConfigSyncRequest) HasNet26() bool`

HasNet26 returns a boolean if a field has been set.

### GetNet27

`func (o *UpdateContainerConfigSyncRequest) GetNet27() string`

GetNet27 returns the Net27 field if non-nil, zero value otherwise.

### GetNet27Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet27Ok() (*string, bool)`

GetNet27Ok returns a tuple with the Net27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet27

`func (o *UpdateContainerConfigSyncRequest) SetNet27(v string)`

SetNet27 sets Net27 field to given value.

### HasNet27

`func (o *UpdateContainerConfigSyncRequest) HasNet27() bool`

HasNet27 returns a boolean if a field has been set.

### GetNet28

`func (o *UpdateContainerConfigSyncRequest) GetNet28() string`

GetNet28 returns the Net28 field if non-nil, zero value otherwise.

### GetNet28Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet28Ok() (*string, bool)`

GetNet28Ok returns a tuple with the Net28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet28

`func (o *UpdateContainerConfigSyncRequest) SetNet28(v string)`

SetNet28 sets Net28 field to given value.

### HasNet28

`func (o *UpdateContainerConfigSyncRequest) HasNet28() bool`

HasNet28 returns a boolean if a field has been set.

### GetNet29

`func (o *UpdateContainerConfigSyncRequest) GetNet29() string`

GetNet29 returns the Net29 field if non-nil, zero value otherwise.

### GetNet29Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet29Ok() (*string, bool)`

GetNet29Ok returns a tuple with the Net29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet29

`func (o *UpdateContainerConfigSyncRequest) SetNet29(v string)`

SetNet29 sets Net29 field to given value.

### HasNet29

`func (o *UpdateContainerConfigSyncRequest) HasNet29() bool`

HasNet29 returns a boolean if a field has been set.

### GetNet30

`func (o *UpdateContainerConfigSyncRequest) GetNet30() string`

GetNet30 returns the Net30 field if non-nil, zero value otherwise.

### GetNet30Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet30Ok() (*string, bool)`

GetNet30Ok returns a tuple with the Net30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet30

`func (o *UpdateContainerConfigSyncRequest) SetNet30(v string)`

SetNet30 sets Net30 field to given value.

### HasNet30

`func (o *UpdateContainerConfigSyncRequest) HasNet30() bool`

HasNet30 returns a boolean if a field has been set.

### GetNet31

`func (o *UpdateContainerConfigSyncRequest) GetNet31() string`

GetNet31 returns the Net31 field if non-nil, zero value otherwise.

### GetNet31Ok

`func (o *UpdateContainerConfigSyncRequest) GetNet31Ok() (*string, bool)`

GetNet31Ok returns a tuple with the Net31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet31

`func (o *UpdateContainerConfigSyncRequest) SetNet31(v string)`

SetNet31 sets Net31 field to given value.

### HasNet31

`func (o *UpdateContainerConfigSyncRequest) HasNet31() bool`

HasNet31 returns a boolean if a field has been set.

### GetOnboot

`func (o *UpdateContainerConfigSyncRequest) GetOnboot() bool`

GetOnboot returns the Onboot field if non-nil, zero value otherwise.

### GetOnbootOk

`func (o *UpdateContainerConfigSyncRequest) GetOnbootOk() (*bool, bool)`

GetOnbootOk returns a tuple with the Onboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboot

`func (o *UpdateContainerConfigSyncRequest) SetOnboot(v bool)`

SetOnboot sets Onboot field to given value.

### HasOnboot

`func (o *UpdateContainerConfigSyncRequest) HasOnboot() bool`

HasOnboot returns a boolean if a field has been set.

### GetOstype

`func (o *UpdateContainerConfigSyncRequest) GetOstype() string`

GetOstype returns the Ostype field if non-nil, zero value otherwise.

### GetOstypeOk

`func (o *UpdateContainerConfigSyncRequest) GetOstypeOk() (*string, bool)`

GetOstypeOk returns a tuple with the Ostype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstype

`func (o *UpdateContainerConfigSyncRequest) SetOstype(v string)`

SetOstype sets Ostype field to given value.

### HasOstype

`func (o *UpdateContainerConfigSyncRequest) HasOstype() bool`

HasOstype returns a boolean if a field has been set.

### GetProtection

`func (o *UpdateContainerConfigSyncRequest) GetProtection() bool`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *UpdateContainerConfigSyncRequest) GetProtectionOk() (*bool, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *UpdateContainerConfigSyncRequest) SetProtection(v bool)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *UpdateContainerConfigSyncRequest) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetRevert

`func (o *UpdateContainerConfigSyncRequest) GetRevert() string`

GetRevert returns the Revert field if non-nil, zero value otherwise.

### GetRevertOk

`func (o *UpdateContainerConfigSyncRequest) GetRevertOk() (*string, bool)`

GetRevertOk returns a tuple with the Revert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevert

`func (o *UpdateContainerConfigSyncRequest) SetRevert(v string)`

SetRevert sets Revert field to given value.

### HasRevert

`func (o *UpdateContainerConfigSyncRequest) HasRevert() bool`

HasRevert returns a boolean if a field has been set.

### GetRootfs

`func (o *UpdateContainerConfigSyncRequest) GetRootfs() string`

GetRootfs returns the Rootfs field if non-nil, zero value otherwise.

### GetRootfsOk

`func (o *UpdateContainerConfigSyncRequest) GetRootfsOk() (*string, bool)`

GetRootfsOk returns a tuple with the Rootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfs

`func (o *UpdateContainerConfigSyncRequest) SetRootfs(v string)`

SetRootfs sets Rootfs field to given value.

### HasRootfs

`func (o *UpdateContainerConfigSyncRequest) HasRootfs() bool`

HasRootfs returns a boolean if a field has been set.

### GetSearchdomain

`func (o *UpdateContainerConfigSyncRequest) GetSearchdomain() string`

GetSearchdomain returns the Searchdomain field if non-nil, zero value otherwise.

### GetSearchdomainOk

`func (o *UpdateContainerConfigSyncRequest) GetSearchdomainOk() (*string, bool)`

GetSearchdomainOk returns a tuple with the Searchdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchdomain

`func (o *UpdateContainerConfigSyncRequest) SetSearchdomain(v string)`

SetSearchdomain sets Searchdomain field to given value.

### HasSearchdomain

`func (o *UpdateContainerConfigSyncRequest) HasSearchdomain() bool`

HasSearchdomain returns a boolean if a field has been set.

### GetStartup

`func (o *UpdateContainerConfigSyncRequest) GetStartup() string`

GetStartup returns the Startup field if non-nil, zero value otherwise.

### GetStartupOk

`func (o *UpdateContainerConfigSyncRequest) GetStartupOk() (*string, bool)`

GetStartupOk returns a tuple with the Startup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartup

`func (o *UpdateContainerConfigSyncRequest) SetStartup(v string)`

SetStartup sets Startup field to given value.

### HasStartup

`func (o *UpdateContainerConfigSyncRequest) HasStartup() bool`

HasStartup returns a boolean if a field has been set.

### GetSwap

`func (o *UpdateContainerConfigSyncRequest) GetSwap() int64`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *UpdateContainerConfigSyncRequest) GetSwapOk() (*int64, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *UpdateContainerConfigSyncRequest) SetSwap(v int64)`

SetSwap sets Swap field to given value.

### HasSwap

`func (o *UpdateContainerConfigSyncRequest) HasSwap() bool`

HasSwap returns a boolean if a field has been set.

### GetTags

`func (o *UpdateContainerConfigSyncRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateContainerConfigSyncRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateContainerConfigSyncRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateContainerConfigSyncRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *UpdateContainerConfigSyncRequest) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UpdateContainerConfigSyncRequest) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UpdateContainerConfigSyncRequest) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *UpdateContainerConfigSyncRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdateContainerConfigSyncRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateContainerConfigSyncRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateContainerConfigSyncRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdateContainerConfigSyncRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTty

`func (o *UpdateContainerConfigSyncRequest) GetTty() int64`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *UpdateContainerConfigSyncRequest) GetTtyOk() (*int64, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *UpdateContainerConfigSyncRequest) SetTty(v int64)`

SetTty sets Tty field to given value.

### HasTty

`func (o *UpdateContainerConfigSyncRequest) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetUnprivileged

`func (o *UpdateContainerConfigSyncRequest) GetUnprivileged() bool`

GetUnprivileged returns the Unprivileged field if non-nil, zero value otherwise.

### GetUnprivilegedOk

`func (o *UpdateContainerConfigSyncRequest) GetUnprivilegedOk() (*bool, bool)`

GetUnprivilegedOk returns a tuple with the Unprivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprivileged

`func (o *UpdateContainerConfigSyncRequest) SetUnprivileged(v bool)`

SetUnprivileged sets Unprivileged field to given value.

### HasUnprivileged

`func (o *UpdateContainerConfigSyncRequest) HasUnprivileged() bool`

HasUnprivileged returns a boolean if a field has been set.

### GetUnused0

`func (o *UpdateContainerConfigSyncRequest) GetUnused0() string`

GetUnused0 returns the Unused0 field if non-nil, zero value otherwise.

### GetUnused0Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused0Ok() (*string, bool)`

GetUnused0Ok returns a tuple with the Unused0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused0

`func (o *UpdateContainerConfigSyncRequest) SetUnused0(v string)`

SetUnused0 sets Unused0 field to given value.

### HasUnused0

`func (o *UpdateContainerConfigSyncRequest) HasUnused0() bool`

HasUnused0 returns a boolean if a field has been set.

### GetUnused1

`func (o *UpdateContainerConfigSyncRequest) GetUnused1() string`

GetUnused1 returns the Unused1 field if non-nil, zero value otherwise.

### GetUnused1Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused1Ok() (*string, bool)`

GetUnused1Ok returns a tuple with the Unused1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused1

`func (o *UpdateContainerConfigSyncRequest) SetUnused1(v string)`

SetUnused1 sets Unused1 field to given value.

### HasUnused1

`func (o *UpdateContainerConfigSyncRequest) HasUnused1() bool`

HasUnused1 returns a boolean if a field has been set.

### GetUnused2

`func (o *UpdateContainerConfigSyncRequest) GetUnused2() string`

GetUnused2 returns the Unused2 field if non-nil, zero value otherwise.

### GetUnused2Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused2Ok() (*string, bool)`

GetUnused2Ok returns a tuple with the Unused2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused2

`func (o *UpdateContainerConfigSyncRequest) SetUnused2(v string)`

SetUnused2 sets Unused2 field to given value.

### HasUnused2

`func (o *UpdateContainerConfigSyncRequest) HasUnused2() bool`

HasUnused2 returns a boolean if a field has been set.

### GetUnused3

`func (o *UpdateContainerConfigSyncRequest) GetUnused3() string`

GetUnused3 returns the Unused3 field if non-nil, zero value otherwise.

### GetUnused3Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused3Ok() (*string, bool)`

GetUnused3Ok returns a tuple with the Unused3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused3

`func (o *UpdateContainerConfigSyncRequest) SetUnused3(v string)`

SetUnused3 sets Unused3 field to given value.

### HasUnused3

`func (o *UpdateContainerConfigSyncRequest) HasUnused3() bool`

HasUnused3 returns a boolean if a field has been set.

### GetUnused4

`func (o *UpdateContainerConfigSyncRequest) GetUnused4() string`

GetUnused4 returns the Unused4 field if non-nil, zero value otherwise.

### GetUnused4Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused4Ok() (*string, bool)`

GetUnused4Ok returns a tuple with the Unused4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused4

`func (o *UpdateContainerConfigSyncRequest) SetUnused4(v string)`

SetUnused4 sets Unused4 field to given value.

### HasUnused4

`func (o *UpdateContainerConfigSyncRequest) HasUnused4() bool`

HasUnused4 returns a boolean if a field has been set.

### GetUnused5

`func (o *UpdateContainerConfigSyncRequest) GetUnused5() string`

GetUnused5 returns the Unused5 field if non-nil, zero value otherwise.

### GetUnused5Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused5Ok() (*string, bool)`

GetUnused5Ok returns a tuple with the Unused5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused5

`func (o *UpdateContainerConfigSyncRequest) SetUnused5(v string)`

SetUnused5 sets Unused5 field to given value.

### HasUnused5

`func (o *UpdateContainerConfigSyncRequest) HasUnused5() bool`

HasUnused5 returns a boolean if a field has been set.

### GetUnused6

`func (o *UpdateContainerConfigSyncRequest) GetUnused6() string`

GetUnused6 returns the Unused6 field if non-nil, zero value otherwise.

### GetUnused6Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused6Ok() (*string, bool)`

GetUnused6Ok returns a tuple with the Unused6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused6

`func (o *UpdateContainerConfigSyncRequest) SetUnused6(v string)`

SetUnused6 sets Unused6 field to given value.

### HasUnused6

`func (o *UpdateContainerConfigSyncRequest) HasUnused6() bool`

HasUnused6 returns a boolean if a field has been set.

### GetUnused7

`func (o *UpdateContainerConfigSyncRequest) GetUnused7() string`

GetUnused7 returns the Unused7 field if non-nil, zero value otherwise.

### GetUnused7Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused7Ok() (*string, bool)`

GetUnused7Ok returns a tuple with the Unused7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused7

`func (o *UpdateContainerConfigSyncRequest) SetUnused7(v string)`

SetUnused7 sets Unused7 field to given value.

### HasUnused7

`func (o *UpdateContainerConfigSyncRequest) HasUnused7() bool`

HasUnused7 returns a boolean if a field has been set.

### GetUnused8

`func (o *UpdateContainerConfigSyncRequest) GetUnused8() string`

GetUnused8 returns the Unused8 field if non-nil, zero value otherwise.

### GetUnused8Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused8Ok() (*string, bool)`

GetUnused8Ok returns a tuple with the Unused8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused8

`func (o *UpdateContainerConfigSyncRequest) SetUnused8(v string)`

SetUnused8 sets Unused8 field to given value.

### HasUnused8

`func (o *UpdateContainerConfigSyncRequest) HasUnused8() bool`

HasUnused8 returns a boolean if a field has been set.

### GetUnused9

`func (o *UpdateContainerConfigSyncRequest) GetUnused9() string`

GetUnused9 returns the Unused9 field if non-nil, zero value otherwise.

### GetUnused9Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused9Ok() (*string, bool)`

GetUnused9Ok returns a tuple with the Unused9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused9

`func (o *UpdateContainerConfigSyncRequest) SetUnused9(v string)`

SetUnused9 sets Unused9 field to given value.

### HasUnused9

`func (o *UpdateContainerConfigSyncRequest) HasUnused9() bool`

HasUnused9 returns a boolean if a field has been set.

### GetUnused10

`func (o *UpdateContainerConfigSyncRequest) GetUnused10() string`

GetUnused10 returns the Unused10 field if non-nil, zero value otherwise.

### GetUnused10Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused10Ok() (*string, bool)`

GetUnused10Ok returns a tuple with the Unused10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused10

`func (o *UpdateContainerConfigSyncRequest) SetUnused10(v string)`

SetUnused10 sets Unused10 field to given value.

### HasUnused10

`func (o *UpdateContainerConfigSyncRequest) HasUnused10() bool`

HasUnused10 returns a boolean if a field has been set.

### GetUnused11

`func (o *UpdateContainerConfigSyncRequest) GetUnused11() string`

GetUnused11 returns the Unused11 field if non-nil, zero value otherwise.

### GetUnused11Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused11Ok() (*string, bool)`

GetUnused11Ok returns a tuple with the Unused11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused11

`func (o *UpdateContainerConfigSyncRequest) SetUnused11(v string)`

SetUnused11 sets Unused11 field to given value.

### HasUnused11

`func (o *UpdateContainerConfigSyncRequest) HasUnused11() bool`

HasUnused11 returns a boolean if a field has been set.

### GetUnused12

`func (o *UpdateContainerConfigSyncRequest) GetUnused12() string`

GetUnused12 returns the Unused12 field if non-nil, zero value otherwise.

### GetUnused12Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused12Ok() (*string, bool)`

GetUnused12Ok returns a tuple with the Unused12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused12

`func (o *UpdateContainerConfigSyncRequest) SetUnused12(v string)`

SetUnused12 sets Unused12 field to given value.

### HasUnused12

`func (o *UpdateContainerConfigSyncRequest) HasUnused12() bool`

HasUnused12 returns a boolean if a field has been set.

### GetUnused13

`func (o *UpdateContainerConfigSyncRequest) GetUnused13() string`

GetUnused13 returns the Unused13 field if non-nil, zero value otherwise.

### GetUnused13Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused13Ok() (*string, bool)`

GetUnused13Ok returns a tuple with the Unused13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused13

`func (o *UpdateContainerConfigSyncRequest) SetUnused13(v string)`

SetUnused13 sets Unused13 field to given value.

### HasUnused13

`func (o *UpdateContainerConfigSyncRequest) HasUnused13() bool`

HasUnused13 returns a boolean if a field has been set.

### GetUnused14

`func (o *UpdateContainerConfigSyncRequest) GetUnused14() string`

GetUnused14 returns the Unused14 field if non-nil, zero value otherwise.

### GetUnused14Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused14Ok() (*string, bool)`

GetUnused14Ok returns a tuple with the Unused14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused14

`func (o *UpdateContainerConfigSyncRequest) SetUnused14(v string)`

SetUnused14 sets Unused14 field to given value.

### HasUnused14

`func (o *UpdateContainerConfigSyncRequest) HasUnused14() bool`

HasUnused14 returns a boolean if a field has been set.

### GetUnused15

`func (o *UpdateContainerConfigSyncRequest) GetUnused15() string`

GetUnused15 returns the Unused15 field if non-nil, zero value otherwise.

### GetUnused15Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused15Ok() (*string, bool)`

GetUnused15Ok returns a tuple with the Unused15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused15

`func (o *UpdateContainerConfigSyncRequest) SetUnused15(v string)`

SetUnused15 sets Unused15 field to given value.

### HasUnused15

`func (o *UpdateContainerConfigSyncRequest) HasUnused15() bool`

HasUnused15 returns a boolean if a field has been set.

### GetUnused16

`func (o *UpdateContainerConfigSyncRequest) GetUnused16() string`

GetUnused16 returns the Unused16 field if non-nil, zero value otherwise.

### GetUnused16Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused16Ok() (*string, bool)`

GetUnused16Ok returns a tuple with the Unused16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused16

`func (o *UpdateContainerConfigSyncRequest) SetUnused16(v string)`

SetUnused16 sets Unused16 field to given value.

### HasUnused16

`func (o *UpdateContainerConfigSyncRequest) HasUnused16() bool`

HasUnused16 returns a boolean if a field has been set.

### GetUnused17

`func (o *UpdateContainerConfigSyncRequest) GetUnused17() string`

GetUnused17 returns the Unused17 field if non-nil, zero value otherwise.

### GetUnused17Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused17Ok() (*string, bool)`

GetUnused17Ok returns a tuple with the Unused17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused17

`func (o *UpdateContainerConfigSyncRequest) SetUnused17(v string)`

SetUnused17 sets Unused17 field to given value.

### HasUnused17

`func (o *UpdateContainerConfigSyncRequest) HasUnused17() bool`

HasUnused17 returns a boolean if a field has been set.

### GetUnused18

`func (o *UpdateContainerConfigSyncRequest) GetUnused18() string`

GetUnused18 returns the Unused18 field if non-nil, zero value otherwise.

### GetUnused18Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused18Ok() (*string, bool)`

GetUnused18Ok returns a tuple with the Unused18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused18

`func (o *UpdateContainerConfigSyncRequest) SetUnused18(v string)`

SetUnused18 sets Unused18 field to given value.

### HasUnused18

`func (o *UpdateContainerConfigSyncRequest) HasUnused18() bool`

HasUnused18 returns a boolean if a field has been set.

### GetUnused19

`func (o *UpdateContainerConfigSyncRequest) GetUnused19() string`

GetUnused19 returns the Unused19 field if non-nil, zero value otherwise.

### GetUnused19Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused19Ok() (*string, bool)`

GetUnused19Ok returns a tuple with the Unused19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused19

`func (o *UpdateContainerConfigSyncRequest) SetUnused19(v string)`

SetUnused19 sets Unused19 field to given value.

### HasUnused19

`func (o *UpdateContainerConfigSyncRequest) HasUnused19() bool`

HasUnused19 returns a boolean if a field has been set.

### GetUnused20

`func (o *UpdateContainerConfigSyncRequest) GetUnused20() string`

GetUnused20 returns the Unused20 field if non-nil, zero value otherwise.

### GetUnused20Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused20Ok() (*string, bool)`

GetUnused20Ok returns a tuple with the Unused20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused20

`func (o *UpdateContainerConfigSyncRequest) SetUnused20(v string)`

SetUnused20 sets Unused20 field to given value.

### HasUnused20

`func (o *UpdateContainerConfigSyncRequest) HasUnused20() bool`

HasUnused20 returns a boolean if a field has been set.

### GetUnused21

`func (o *UpdateContainerConfigSyncRequest) GetUnused21() string`

GetUnused21 returns the Unused21 field if non-nil, zero value otherwise.

### GetUnused21Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused21Ok() (*string, bool)`

GetUnused21Ok returns a tuple with the Unused21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused21

`func (o *UpdateContainerConfigSyncRequest) SetUnused21(v string)`

SetUnused21 sets Unused21 field to given value.

### HasUnused21

`func (o *UpdateContainerConfigSyncRequest) HasUnused21() bool`

HasUnused21 returns a boolean if a field has been set.

### GetUnused22

`func (o *UpdateContainerConfigSyncRequest) GetUnused22() string`

GetUnused22 returns the Unused22 field if non-nil, zero value otherwise.

### GetUnused22Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused22Ok() (*string, bool)`

GetUnused22Ok returns a tuple with the Unused22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused22

`func (o *UpdateContainerConfigSyncRequest) SetUnused22(v string)`

SetUnused22 sets Unused22 field to given value.

### HasUnused22

`func (o *UpdateContainerConfigSyncRequest) HasUnused22() bool`

HasUnused22 returns a boolean if a field has been set.

### GetUnused23

`func (o *UpdateContainerConfigSyncRequest) GetUnused23() string`

GetUnused23 returns the Unused23 field if non-nil, zero value otherwise.

### GetUnused23Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused23Ok() (*string, bool)`

GetUnused23Ok returns a tuple with the Unused23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused23

`func (o *UpdateContainerConfigSyncRequest) SetUnused23(v string)`

SetUnused23 sets Unused23 field to given value.

### HasUnused23

`func (o *UpdateContainerConfigSyncRequest) HasUnused23() bool`

HasUnused23 returns a boolean if a field has been set.

### GetUnused24

`func (o *UpdateContainerConfigSyncRequest) GetUnused24() string`

GetUnused24 returns the Unused24 field if non-nil, zero value otherwise.

### GetUnused24Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused24Ok() (*string, bool)`

GetUnused24Ok returns a tuple with the Unused24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused24

`func (o *UpdateContainerConfigSyncRequest) SetUnused24(v string)`

SetUnused24 sets Unused24 field to given value.

### HasUnused24

`func (o *UpdateContainerConfigSyncRequest) HasUnused24() bool`

HasUnused24 returns a boolean if a field has been set.

### GetUnused25

`func (o *UpdateContainerConfigSyncRequest) GetUnused25() string`

GetUnused25 returns the Unused25 field if non-nil, zero value otherwise.

### GetUnused25Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused25Ok() (*string, bool)`

GetUnused25Ok returns a tuple with the Unused25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused25

`func (o *UpdateContainerConfigSyncRequest) SetUnused25(v string)`

SetUnused25 sets Unused25 field to given value.

### HasUnused25

`func (o *UpdateContainerConfigSyncRequest) HasUnused25() bool`

HasUnused25 returns a boolean if a field has been set.

### GetUnused26

`func (o *UpdateContainerConfigSyncRequest) GetUnused26() string`

GetUnused26 returns the Unused26 field if non-nil, zero value otherwise.

### GetUnused26Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused26Ok() (*string, bool)`

GetUnused26Ok returns a tuple with the Unused26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused26

`func (o *UpdateContainerConfigSyncRequest) SetUnused26(v string)`

SetUnused26 sets Unused26 field to given value.

### HasUnused26

`func (o *UpdateContainerConfigSyncRequest) HasUnused26() bool`

HasUnused26 returns a boolean if a field has been set.

### GetUnused27

`func (o *UpdateContainerConfigSyncRequest) GetUnused27() string`

GetUnused27 returns the Unused27 field if non-nil, zero value otherwise.

### GetUnused27Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused27Ok() (*string, bool)`

GetUnused27Ok returns a tuple with the Unused27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused27

`func (o *UpdateContainerConfigSyncRequest) SetUnused27(v string)`

SetUnused27 sets Unused27 field to given value.

### HasUnused27

`func (o *UpdateContainerConfigSyncRequest) HasUnused27() bool`

HasUnused27 returns a boolean if a field has been set.

### GetUnused28

`func (o *UpdateContainerConfigSyncRequest) GetUnused28() string`

GetUnused28 returns the Unused28 field if non-nil, zero value otherwise.

### GetUnused28Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused28Ok() (*string, bool)`

GetUnused28Ok returns a tuple with the Unused28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused28

`func (o *UpdateContainerConfigSyncRequest) SetUnused28(v string)`

SetUnused28 sets Unused28 field to given value.

### HasUnused28

`func (o *UpdateContainerConfigSyncRequest) HasUnused28() bool`

HasUnused28 returns a boolean if a field has been set.

### GetUnused29

`func (o *UpdateContainerConfigSyncRequest) GetUnused29() string`

GetUnused29 returns the Unused29 field if non-nil, zero value otherwise.

### GetUnused29Ok

`func (o *UpdateContainerConfigSyncRequest) GetUnused29Ok() (*string, bool)`

GetUnused29Ok returns a tuple with the Unused29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused29

`func (o *UpdateContainerConfigSyncRequest) SetUnused29(v string)`

SetUnused29 sets Unused29 field to given value.

### HasUnused29

`func (o *UpdateContainerConfigSyncRequest) HasUnused29() bool`

HasUnused29 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


