# UpdateVMConfigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acpi** | Pointer to **bool** | Enable/disable ACPI. | [optional] 
**Affinity** | Pointer to **string** | List of host cores used to execute guest processes, for example: 0,5,8-11 | [optional] 
**Agent** | Pointer to **string** | Enable/disable communication with the QEMU Guest Agent and its properties. | [optional] 
**AmdSev** | Pointer to **string** | Secure Encrypted Virtualization (SEV) features by AMD CPUs | [optional] 
**Arch** | Pointer to **string** | Virtual processor architecture. Defaults to the host. | [optional] 
**Args** | Pointer to **string** | Arbitrary arguments passed to kvm. | [optional] 
**Audio0** | Pointer to **string** | Configure a audio device, useful in combination with QXL/Spice. | [optional] 
**Autostart** | Pointer to **bool** | Automatic restart after crash (currently ignored). | [optional] 
**BackgroundDelay** | Pointer to **int64** | Time to wait for the task to finish. We return &#39;null&#39; if the task finish within that time. | [optional] 
**Balloon** | Pointer to **int64** | Amount of target RAM for the VM in MiB. Using zero disables the ballon driver. | [optional] 
**Bios** | Pointer to **string** | Select BIOS implementation. | [optional] 
**Boot** | Pointer to **string** | Specify guest boot order. Use the &#39;order&#x3D;&#39; sub-property as usage with no key or &#39;legacy&#x3D;&#39; is deprecated. | [optional] 
**Bootdisk** | Pointer to **string** | Enable booting from specified disk. Deprecated: Use &#39;boot: order&#x3D;foo;bar&#39; instead. | [optional] 
**Cdrom** | Pointer to **string** | This is an alias for option -ide2 | [optional] 
**Cicustom** | Pointer to **string** | cloud-init: Specify custom files to replace the automatically generated ones at start. | [optional] 
**Cipassword** | Pointer to **string** | cloud-init: Password to assign the user. Using this is generally not recommended. Use ssh keys instead. Also note that older cloud-init versions do not support hashed passwords. | [optional] 
**Citype** | Pointer to **string** | Specifies the cloud-init configuration format. The default depends on the configured operating system type (&#x60;ostype&#x60;. We use the &#x60;nocloud&#x60; format for Linux, and &#x60;configdrive2&#x60; for windows. | [optional] 
**Ciupgrade** | Pointer to **bool** | cloud-init: do an automatic package upgrade after the first boot. | [optional] 
**Ciuser** | Pointer to **string** | cloud-init: User name to change ssh keys and password for instead of the image&#39;s configured default user. | [optional] 
**Cores** | Pointer to **int64** | The number of cores per socket. | [optional] 
**Cpu** | Pointer to **string** | Emulated CPU type. | [optional] 
**Cpulimit** | Pointer to **float32** | Limit of CPU usage. | [optional] 
**Cpuunits** | Pointer to **int64** | CPU weight for a VM, will be clamped to [1, 10000] in cgroup v2. | [optional] 
**Delete** | Pointer to **string** | A list of settings you want to delete. | [optional] 
**Description** | Pointer to **string** | Description for the VM. Shown in the web-interface VM&#39;s summary. This is saved as comment inside the configuration file. | [optional] 
**Digest** | Pointer to **string** | Prevent changes if current configuration file has different SHA1 digest. This can be used to prevent concurrent modifications. | [optional] 
**Efidisk0** | Pointer to **string** | Configure a disk for storing EFI vars. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and that the default EFI vars are copied to the volume instead. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Force** | Pointer to **bool** | Force physical removal. Without this, we simple remove the disk from the config file and create an additional configuration entry called &#39;unused[n]&#39;, which contains the volume ID. Unlink of unused[n] always cause physical removal. | [optional] 
**Freeze** | Pointer to **bool** | Freeze CPU at startup (use &#39;c&#39; monitor command to start execution). | [optional] 
**Hookscript** | Pointer to **string** | Script that will be executed during various steps in the vms lifetime. | [optional] 
**Hostpci0** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci1** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci2** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci3** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci4** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci5** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci6** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci7** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci8** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci9** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci10** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci11** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci12** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci13** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci14** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci15** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci16** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci17** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci18** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci19** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci20** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci21** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci22** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci23** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci24** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci25** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci26** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci27** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci28** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hostpci29** | Pointer to **string** | Map host PCI devices into guest. | [optional] 
**Hotplug** | Pointer to **string** | Selectively enable hotplug features. This is a comma separated list of hotplug features: &#39;network&#39;, &#39;disk&#39;, &#39;cpu&#39;, &#39;memory&#39;, &#39;usb&#39; and &#39;cloudinit&#39;. Use &#39;0&#39; to disable hotplug completely. Using &#39;1&#39; as value is an alias for the default &#x60;network,disk,usb&#x60;. USB hotplugging is possible for guests with machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7. | [optional] 
**Hugepages** | Pointer to **string** | Enable/disable hugepages memory. | [optional] 
**Ide0** | Pointer to **string** | Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Ide1** | Pointer to **string** | Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Ide2** | Pointer to **string** | Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Ide3** | Pointer to **string** | Use volume as IDE hard disk or CD-ROM (n is 0 to 3). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**ImportWorkingStorage** | Pointer to **string** | A file-based storage with &#39;images&#39; content-type enabled, which is used as an intermediary extraction storage during import. Defaults to the source storage. | [optional] 
**Ipconfig0** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig1** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig2** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig3** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig4** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig5** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig6** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig7** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig8** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig9** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig10** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig11** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig12** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig13** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig14** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig15** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig16** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig17** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig18** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig19** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig20** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig21** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig22** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig23** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig24** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig25** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig26** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig27** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig28** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ipconfig29** | Pointer to **string** | cloud-init: Specify IP addresses and gateways for the corresponding interface.  IP addresses use CIDR notation, gateways are optional but need an IP of the same type specified.  The special string &#39;dhcp&#39; can be used for IP addresses to use DHCP, in which case no explicit gateway should be provided. For IPv6 the special string &#39;auto&#39; can be used to use stateless autoconfiguration. This requires cloud-init 19.4 or newer.  If cloud-init is enabled and neither an IPv4 nor an IPv6 address is specified, it defaults to using dhcp on IPv4.  | [optional] 
**Ivshmem** | Pointer to **string** | Inter-VM shared memory. Useful for direct communication between VMs, or to the host. | [optional] 
**Keephugepages** | Pointer to **bool** | Use together with hugepages. If enabled, hugepages will not not be deleted after VM shutdown and can be used for subsequent starts. | [optional] 
**Keyboard** | Pointer to **string** | Keyboard layout for VNC server. This option is generally not required and is often better handled from within the guest OS. | [optional] 
**Kvm** | Pointer to **bool** | Enable/disable KVM hardware virtualization. | [optional] 
**Localtime** | Pointer to **bool** | Set the real time clock (RTC) to local time. This is enabled by default if the &#x60;ostype&#x60; indicates a Microsoft Windows OS. | [optional] 
**Lock** | Pointer to **string** | Lock/unlock the VM. | [optional] 
**Machine** | Pointer to **string** | Specify the QEMU machine. | [optional] 
**Memory** | Pointer to **string** | Memory properties. | [optional] 
**MigrateDowntime** | Pointer to **float32** | Set maximum tolerated downtime (in seconds) for migrations. Should the migration not be able to converge in the very end, because too much newly dirtied RAM needs to be transferred, the limit will be increased automatically step-by-step until migration can converge. | [optional] 
**MigrateSpeed** | Pointer to **int64** | Set maximum speed (in MB/s) for migrations. Value 0 is no limit. | [optional] 
**Name** | Pointer to **string** | Set a name for the VM. Only used on the configuration web interface. | [optional] 
**Nameserver** | Pointer to **string** | cloud-init: Sets DNS server IP address for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set. | [optional] 
**Net0** | Pointer to **string** | Specify network devices. | [optional] 
**Net1** | Pointer to **string** | Specify network devices. | [optional] 
**Net2** | Pointer to **string** | Specify network devices. | [optional] 
**Net3** | Pointer to **string** | Specify network devices. | [optional] 
**Net4** | Pointer to **string** | Specify network devices. | [optional] 
**Net5** | Pointer to **string** | Specify network devices. | [optional] 
**Net6** | Pointer to **string** | Specify network devices. | [optional] 
**Net7** | Pointer to **string** | Specify network devices. | [optional] 
**Net8** | Pointer to **string** | Specify network devices. | [optional] 
**Net9** | Pointer to **string** | Specify network devices. | [optional] 
**Net10** | Pointer to **string** | Specify network devices. | [optional] 
**Net11** | Pointer to **string** | Specify network devices. | [optional] 
**Net12** | Pointer to **string** | Specify network devices. | [optional] 
**Net13** | Pointer to **string** | Specify network devices. | [optional] 
**Net14** | Pointer to **string** | Specify network devices. | [optional] 
**Net15** | Pointer to **string** | Specify network devices. | [optional] 
**Net16** | Pointer to **string** | Specify network devices. | [optional] 
**Net17** | Pointer to **string** | Specify network devices. | [optional] 
**Net18** | Pointer to **string** | Specify network devices. | [optional] 
**Net19** | Pointer to **string** | Specify network devices. | [optional] 
**Net20** | Pointer to **string** | Specify network devices. | [optional] 
**Net21** | Pointer to **string** | Specify network devices. | [optional] 
**Net22** | Pointer to **string** | Specify network devices. | [optional] 
**Net23** | Pointer to **string** | Specify network devices. | [optional] 
**Net24** | Pointer to **string** | Specify network devices. | [optional] 
**Net25** | Pointer to **string** | Specify network devices. | [optional] 
**Net26** | Pointer to **string** | Specify network devices. | [optional] 
**Net27** | Pointer to **string** | Specify network devices. | [optional] 
**Net28** | Pointer to **string** | Specify network devices. | [optional] 
**Net29** | Pointer to **string** | Specify network devices. | [optional] 
**Net30** | Pointer to **string** | Specify network devices. | [optional] 
**Net31** | Pointer to **string** | Specify network devices. | [optional] 
**Numa** | Pointer to **bool** | Enable/disable NUMA. | [optional] 
**Numa0** | Pointer to **string** | NUMA topology. | [optional] 
**Numa1** | Pointer to **string** | NUMA topology. | [optional] 
**Numa2** | Pointer to **string** | NUMA topology. | [optional] 
**Numa3** | Pointer to **string** | NUMA topology. | [optional] 
**Numa4** | Pointer to **string** | NUMA topology. | [optional] 
**Numa5** | Pointer to **string** | NUMA topology. | [optional] 
**Numa6** | Pointer to **string** | NUMA topology. | [optional] 
**Numa7** | Pointer to **string** | NUMA topology. | [optional] 
**Numa8** | Pointer to **string** | NUMA topology. | [optional] 
**Numa9** | Pointer to **string** | NUMA topology. | [optional] 
**Numa10** | Pointer to **string** | NUMA topology. | [optional] 
**Numa11** | Pointer to **string** | NUMA topology. | [optional] 
**Numa12** | Pointer to **string** | NUMA topology. | [optional] 
**Numa13** | Pointer to **string** | NUMA topology. | [optional] 
**Numa14** | Pointer to **string** | NUMA topology. | [optional] 
**Numa15** | Pointer to **string** | NUMA topology. | [optional] 
**Numa16** | Pointer to **string** | NUMA topology. | [optional] 
**Numa17** | Pointer to **string** | NUMA topology. | [optional] 
**Numa18** | Pointer to **string** | NUMA topology. | [optional] 
**Numa19** | Pointer to **string** | NUMA topology. | [optional] 
**Numa20** | Pointer to **string** | NUMA topology. | [optional] 
**Numa21** | Pointer to **string** | NUMA topology. | [optional] 
**Numa22** | Pointer to **string** | NUMA topology. | [optional] 
**Numa23** | Pointer to **string** | NUMA topology. | [optional] 
**Numa24** | Pointer to **string** | NUMA topology. | [optional] 
**Numa25** | Pointer to **string** | NUMA topology. | [optional] 
**Numa26** | Pointer to **string** | NUMA topology. | [optional] 
**Numa27** | Pointer to **string** | NUMA topology. | [optional] 
**Numa28** | Pointer to **string** | NUMA topology. | [optional] 
**Numa29** | Pointer to **string** | NUMA topology. | [optional] 
**Onboot** | Pointer to **bool** | Specifies whether a VM will be started during system bootup. | [optional] 
**Ostype** | Pointer to **string** | Specify guest operating system. | [optional] 
**Parallel0** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Parallel1** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Parallel2** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Parallel3** | Pointer to **string** | Map host parallel devices (n is 0 to 2). | [optional] 
**Protection** | Pointer to **bool** | Sets the protection flag of the VM. This will disable the remove VM and remove disk operations. | [optional] 
**Reboot** | Pointer to **bool** | Allow reboot. If set to &#39;0&#39; the VM exit on reboot. | [optional] 
**Revert** | Pointer to **string** | Revert a pending change. | [optional] 
**Rng0** | Pointer to **string** | Configure a VirtIO-based Random Number Generator. | [optional] 
**Sata0** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Sata1** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Sata2** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Sata3** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Sata4** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Sata5** | Pointer to **string** | Use volume as SATA hard disk or CD-ROM (n is 0 to 5). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi0** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi1** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi2** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi3** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi4** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi5** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi6** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi7** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi8** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi9** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi10** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi11** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi12** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi13** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi14** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi15** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi16** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi17** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi18** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi19** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi20** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi21** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi22** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi23** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi24** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi25** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi26** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi27** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi28** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsi29** | Pointer to **string** | Use volume as SCSI hard disk or CD-ROM (n is 0 to 30). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Scsihw** | Pointer to **string** | SCSI controller model | [optional] 
**Searchdomain** | Pointer to **string** | cloud-init: Sets DNS search domains for a container. Create will automatically use the setting from the host if neither searchdomain nor nameserver are set. | [optional] 
**Serial0** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Serial1** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Serial2** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Serial3** | Pointer to **string** | Create a serial device inside the VM (n is 0 to 3) | [optional] 
**Shares** | Pointer to **int64** | Amount of memory shares for auto-ballooning. The larger the number is, the more memory this VM gets. Number is relative to weights of all other running VMs. Using zero disables auto-ballooning. Auto-ballooning is done by pvestatd. | [optional] 
**Skiplock** | Pointer to **bool** | Ignore locks - only root is allowed to use this option. | [optional] 
**Smbios1** | Pointer to **string** | Specify SMBIOS type 1 fields. | [optional] 
**Smp** | Pointer to **int64** | The number of CPUs. Please use option -sockets instead. | [optional] 
**Sockets** | Pointer to **int64** | The number of CPU sockets. | [optional] 
**SpiceEnhancements** | Pointer to **string** | Configure additional enhancements for SPICE. | [optional] 
**Sshkeys** | Pointer to **string** | cloud-init: Setup public SSH keys (one key per line, OpenSSH format). | [optional] 
**Startdate** | Pointer to **string** | Set the initial date of the real time clock. Valid format for date are:&#39;now&#39; or &#39;2006-06-17T16:01:21&#39; or &#39;2006-06-17&#39;. | [optional] 
**Startup** | Pointer to **string** | Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the &#39;up&#39; or &#39;down&#39; delay in seconds, which specifies a delay to wait before the next VM is started or stopped. | [optional] 
**Tablet** | Pointer to **bool** | Enable/disable the USB tablet device. | [optional] 
**Tags** | Pointer to **string** | Tags of the VM. This is only meta information. | [optional] 
**Tdf** | Pointer to **bool** | Enable/disable time drift fix. | [optional] 
**Template** | Pointer to **bool** | Enable/disable Template. | [optional] 
**Tpmstate0** | Pointer to **string** | Configure a Disk for storing TPM state. The format is fixed to &#39;raw&#39;. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Note that SIZE_IN_GiB is ignored here and 4 MiB will be used instead. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
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
**Usb0** | Pointer to **string** | Configure an USB device (n is 0 to 4, for machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7, n can be up to 14). | [optional] 
**Usb1** | Pointer to **string** | Configure an USB device (n is 0 to 4, for machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7, n can be up to 14). | [optional] 
**Usb2** | Pointer to **string** | Configure an USB device (n is 0 to 4, for machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7, n can be up to 14). | [optional] 
**Usb3** | Pointer to **string** | Configure an USB device (n is 0 to 4, for machine version &gt;&#x3D; 7.1 and ostype l26 or windows &gt; 7, n can be up to 14). | [optional] 
**Vcpus** | Pointer to **int64** | Number of hotplugged vcpus. | [optional] 
**Vga** | Pointer to **string** | Configure the VGA hardware. | [optional] 
**Virtio0** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio1** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio2** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio3** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio4** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio5** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio6** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio7** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio8** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio9** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio10** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio11** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio12** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio13** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio14** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Virtio15** | Pointer to **string** | Use volume as VIRTIO hard disk (n is 0 to 15). Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. Use STORAGE_ID:0 and the &#39;import-from&#39; parameter to import from an existing volume. | [optional] 
**Vmgenid** | Pointer to **string** | Set VM Generation ID. Use &#39;1&#39; to autogenerate on create or update, pass &#39;0&#39; to disable explicitly. | [optional] 
**Vmstatestorage** | Pointer to **string** | Default storage for VM state volumes/files. | [optional] 
**Watchdog** | Pointer to **string** | Create a virtual hardware watchdog device. | [optional] 

## Methods

### NewUpdateVMConfigRequest

`func NewUpdateVMConfigRequest() *UpdateVMConfigRequest`

NewUpdateVMConfigRequest instantiates a new UpdateVMConfigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVMConfigRequestWithDefaults

`func NewUpdateVMConfigRequestWithDefaults() *UpdateVMConfigRequest`

NewUpdateVMConfigRequestWithDefaults instantiates a new UpdateVMConfigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcpi

`func (o *UpdateVMConfigRequest) GetAcpi() bool`

GetAcpi returns the Acpi field if non-nil, zero value otherwise.

### GetAcpiOk

`func (o *UpdateVMConfigRequest) GetAcpiOk() (*bool, bool)`

GetAcpiOk returns a tuple with the Acpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcpi

`func (o *UpdateVMConfigRequest) SetAcpi(v bool)`

SetAcpi sets Acpi field to given value.

### HasAcpi

`func (o *UpdateVMConfigRequest) HasAcpi() bool`

HasAcpi returns a boolean if a field has been set.

### GetAffinity

`func (o *UpdateVMConfigRequest) GetAffinity() string`

GetAffinity returns the Affinity field if non-nil, zero value otherwise.

### GetAffinityOk

`func (o *UpdateVMConfigRequest) GetAffinityOk() (*string, bool)`

GetAffinityOk returns a tuple with the Affinity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffinity

`func (o *UpdateVMConfigRequest) SetAffinity(v string)`

SetAffinity sets Affinity field to given value.

### HasAffinity

`func (o *UpdateVMConfigRequest) HasAffinity() bool`

HasAffinity returns a boolean if a field has been set.

### GetAgent

`func (o *UpdateVMConfigRequest) GetAgent() string`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *UpdateVMConfigRequest) GetAgentOk() (*string, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *UpdateVMConfigRequest) SetAgent(v string)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *UpdateVMConfigRequest) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAmdSev

`func (o *UpdateVMConfigRequest) GetAmdSev() string`

GetAmdSev returns the AmdSev field if non-nil, zero value otherwise.

### GetAmdSevOk

`func (o *UpdateVMConfigRequest) GetAmdSevOk() (*string, bool)`

GetAmdSevOk returns a tuple with the AmdSev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmdSev

`func (o *UpdateVMConfigRequest) SetAmdSev(v string)`

SetAmdSev sets AmdSev field to given value.

### HasAmdSev

`func (o *UpdateVMConfigRequest) HasAmdSev() bool`

HasAmdSev returns a boolean if a field has been set.

### GetArch

`func (o *UpdateVMConfigRequest) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *UpdateVMConfigRequest) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *UpdateVMConfigRequest) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *UpdateVMConfigRequest) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetArgs

`func (o *UpdateVMConfigRequest) GetArgs() string`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *UpdateVMConfigRequest) GetArgsOk() (*string, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *UpdateVMConfigRequest) SetArgs(v string)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *UpdateVMConfigRequest) HasArgs() bool`

HasArgs returns a boolean if a field has been set.

### GetAudio0

`func (o *UpdateVMConfigRequest) GetAudio0() string`

GetAudio0 returns the Audio0 field if non-nil, zero value otherwise.

### GetAudio0Ok

`func (o *UpdateVMConfigRequest) GetAudio0Ok() (*string, bool)`

GetAudio0Ok returns a tuple with the Audio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio0

`func (o *UpdateVMConfigRequest) SetAudio0(v string)`

SetAudio0 sets Audio0 field to given value.

### HasAudio0

`func (o *UpdateVMConfigRequest) HasAudio0() bool`

HasAudio0 returns a boolean if a field has been set.

### GetAutostart

`func (o *UpdateVMConfigRequest) GetAutostart() bool`

GetAutostart returns the Autostart field if non-nil, zero value otherwise.

### GetAutostartOk

`func (o *UpdateVMConfigRequest) GetAutostartOk() (*bool, bool)`

GetAutostartOk returns a tuple with the Autostart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutostart

`func (o *UpdateVMConfigRequest) SetAutostart(v bool)`

SetAutostart sets Autostart field to given value.

### HasAutostart

`func (o *UpdateVMConfigRequest) HasAutostart() bool`

HasAutostart returns a boolean if a field has been set.

### GetBackgroundDelay

`func (o *UpdateVMConfigRequest) GetBackgroundDelay() int64`

GetBackgroundDelay returns the BackgroundDelay field if non-nil, zero value otherwise.

### GetBackgroundDelayOk

`func (o *UpdateVMConfigRequest) GetBackgroundDelayOk() (*int64, bool)`

GetBackgroundDelayOk returns a tuple with the BackgroundDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundDelay

`func (o *UpdateVMConfigRequest) SetBackgroundDelay(v int64)`

SetBackgroundDelay sets BackgroundDelay field to given value.

### HasBackgroundDelay

`func (o *UpdateVMConfigRequest) HasBackgroundDelay() bool`

HasBackgroundDelay returns a boolean if a field has been set.

### GetBalloon

`func (o *UpdateVMConfigRequest) GetBalloon() int64`

GetBalloon returns the Balloon field if non-nil, zero value otherwise.

### GetBalloonOk

`func (o *UpdateVMConfigRequest) GetBalloonOk() (*int64, bool)`

GetBalloonOk returns a tuple with the Balloon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalloon

`func (o *UpdateVMConfigRequest) SetBalloon(v int64)`

SetBalloon sets Balloon field to given value.

### HasBalloon

`func (o *UpdateVMConfigRequest) HasBalloon() bool`

HasBalloon returns a boolean if a field has been set.

### GetBios

`func (o *UpdateVMConfigRequest) GetBios() string`

GetBios returns the Bios field if non-nil, zero value otherwise.

### GetBiosOk

`func (o *UpdateVMConfigRequest) GetBiosOk() (*string, bool)`

GetBiosOk returns a tuple with the Bios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBios

`func (o *UpdateVMConfigRequest) SetBios(v string)`

SetBios sets Bios field to given value.

### HasBios

`func (o *UpdateVMConfigRequest) HasBios() bool`

HasBios returns a boolean if a field has been set.

### GetBoot

`func (o *UpdateVMConfigRequest) GetBoot() string`

GetBoot returns the Boot field if non-nil, zero value otherwise.

### GetBootOk

`func (o *UpdateVMConfigRequest) GetBootOk() (*string, bool)`

GetBootOk returns a tuple with the Boot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoot

`func (o *UpdateVMConfigRequest) SetBoot(v string)`

SetBoot sets Boot field to given value.

### HasBoot

`func (o *UpdateVMConfigRequest) HasBoot() bool`

HasBoot returns a boolean if a field has been set.

### GetBootdisk

`func (o *UpdateVMConfigRequest) GetBootdisk() string`

GetBootdisk returns the Bootdisk field if non-nil, zero value otherwise.

### GetBootdiskOk

`func (o *UpdateVMConfigRequest) GetBootdiskOk() (*string, bool)`

GetBootdiskOk returns a tuple with the Bootdisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootdisk

`func (o *UpdateVMConfigRequest) SetBootdisk(v string)`

SetBootdisk sets Bootdisk field to given value.

### HasBootdisk

`func (o *UpdateVMConfigRequest) HasBootdisk() bool`

HasBootdisk returns a boolean if a field has been set.

### GetCdrom

`func (o *UpdateVMConfigRequest) GetCdrom() string`

GetCdrom returns the Cdrom field if non-nil, zero value otherwise.

### GetCdromOk

`func (o *UpdateVMConfigRequest) GetCdromOk() (*string, bool)`

GetCdromOk returns a tuple with the Cdrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrom

`func (o *UpdateVMConfigRequest) SetCdrom(v string)`

SetCdrom sets Cdrom field to given value.

### HasCdrom

`func (o *UpdateVMConfigRequest) HasCdrom() bool`

HasCdrom returns a boolean if a field has been set.

### GetCicustom

`func (o *UpdateVMConfigRequest) GetCicustom() string`

GetCicustom returns the Cicustom field if non-nil, zero value otherwise.

### GetCicustomOk

`func (o *UpdateVMConfigRequest) GetCicustomOk() (*string, bool)`

GetCicustomOk returns a tuple with the Cicustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCicustom

`func (o *UpdateVMConfigRequest) SetCicustom(v string)`

SetCicustom sets Cicustom field to given value.

### HasCicustom

`func (o *UpdateVMConfigRequest) HasCicustom() bool`

HasCicustom returns a boolean if a field has been set.

### GetCipassword

`func (o *UpdateVMConfigRequest) GetCipassword() string`

GetCipassword returns the Cipassword field if non-nil, zero value otherwise.

### GetCipasswordOk

`func (o *UpdateVMConfigRequest) GetCipasswordOk() (*string, bool)`

GetCipasswordOk returns a tuple with the Cipassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCipassword

`func (o *UpdateVMConfigRequest) SetCipassword(v string)`

SetCipassword sets Cipassword field to given value.

### HasCipassword

`func (o *UpdateVMConfigRequest) HasCipassword() bool`

HasCipassword returns a boolean if a field has been set.

### GetCitype

`func (o *UpdateVMConfigRequest) GetCitype() string`

GetCitype returns the Citype field if non-nil, zero value otherwise.

### GetCitypeOk

`func (o *UpdateVMConfigRequest) GetCitypeOk() (*string, bool)`

GetCitypeOk returns a tuple with the Citype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitype

`func (o *UpdateVMConfigRequest) SetCitype(v string)`

SetCitype sets Citype field to given value.

### HasCitype

`func (o *UpdateVMConfigRequest) HasCitype() bool`

HasCitype returns a boolean if a field has been set.

### GetCiupgrade

`func (o *UpdateVMConfigRequest) GetCiupgrade() bool`

GetCiupgrade returns the Ciupgrade field if non-nil, zero value otherwise.

### GetCiupgradeOk

`func (o *UpdateVMConfigRequest) GetCiupgradeOk() (*bool, bool)`

GetCiupgradeOk returns a tuple with the Ciupgrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiupgrade

`func (o *UpdateVMConfigRequest) SetCiupgrade(v bool)`

SetCiupgrade sets Ciupgrade field to given value.

### HasCiupgrade

`func (o *UpdateVMConfigRequest) HasCiupgrade() bool`

HasCiupgrade returns a boolean if a field has been set.

### GetCiuser

`func (o *UpdateVMConfigRequest) GetCiuser() string`

GetCiuser returns the Ciuser field if non-nil, zero value otherwise.

### GetCiuserOk

`func (o *UpdateVMConfigRequest) GetCiuserOk() (*string, bool)`

GetCiuserOk returns a tuple with the Ciuser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCiuser

`func (o *UpdateVMConfigRequest) SetCiuser(v string)`

SetCiuser sets Ciuser field to given value.

### HasCiuser

`func (o *UpdateVMConfigRequest) HasCiuser() bool`

HasCiuser returns a boolean if a field has been set.

### GetCores

`func (o *UpdateVMConfigRequest) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *UpdateVMConfigRequest) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *UpdateVMConfigRequest) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *UpdateVMConfigRequest) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpu

`func (o *UpdateVMConfigRequest) GetCpu() string`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *UpdateVMConfigRequest) GetCpuOk() (*string, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *UpdateVMConfigRequest) SetCpu(v string)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *UpdateVMConfigRequest) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetCpulimit

`func (o *UpdateVMConfigRequest) GetCpulimit() float32`

GetCpulimit returns the Cpulimit field if non-nil, zero value otherwise.

### GetCpulimitOk

`func (o *UpdateVMConfigRequest) GetCpulimitOk() (*float32, bool)`

GetCpulimitOk returns a tuple with the Cpulimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpulimit

`func (o *UpdateVMConfigRequest) SetCpulimit(v float32)`

SetCpulimit sets Cpulimit field to given value.

### HasCpulimit

`func (o *UpdateVMConfigRequest) HasCpulimit() bool`

HasCpulimit returns a boolean if a field has been set.

### GetCpuunits

`func (o *UpdateVMConfigRequest) GetCpuunits() int64`

GetCpuunits returns the Cpuunits field if non-nil, zero value otherwise.

### GetCpuunitsOk

`func (o *UpdateVMConfigRequest) GetCpuunitsOk() (*int64, bool)`

GetCpuunitsOk returns a tuple with the Cpuunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuunits

`func (o *UpdateVMConfigRequest) SetCpuunits(v int64)`

SetCpuunits sets Cpuunits field to given value.

### HasCpuunits

`func (o *UpdateVMConfigRequest) HasCpuunits() bool`

HasCpuunits returns a boolean if a field has been set.

### GetDelete

`func (o *UpdateVMConfigRequest) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *UpdateVMConfigRequest) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *UpdateVMConfigRequest) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *UpdateVMConfigRequest) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateVMConfigRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVMConfigRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVMConfigRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVMConfigRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDigest

`func (o *UpdateVMConfigRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *UpdateVMConfigRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *UpdateVMConfigRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *UpdateVMConfigRequest) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetEfidisk0

`func (o *UpdateVMConfigRequest) GetEfidisk0() string`

GetEfidisk0 returns the Efidisk0 field if non-nil, zero value otherwise.

### GetEfidisk0Ok

`func (o *UpdateVMConfigRequest) GetEfidisk0Ok() (*string, bool)`

GetEfidisk0Ok returns a tuple with the Efidisk0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEfidisk0

`func (o *UpdateVMConfigRequest) SetEfidisk0(v string)`

SetEfidisk0 sets Efidisk0 field to given value.

### HasEfidisk0

`func (o *UpdateVMConfigRequest) HasEfidisk0() bool`

HasEfidisk0 returns a boolean if a field has been set.

### GetForce

`func (o *UpdateVMConfigRequest) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *UpdateVMConfigRequest) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *UpdateVMConfigRequest) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *UpdateVMConfigRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetFreeze

`func (o *UpdateVMConfigRequest) GetFreeze() bool`

GetFreeze returns the Freeze field if non-nil, zero value otherwise.

### GetFreezeOk

`func (o *UpdateVMConfigRequest) GetFreezeOk() (*bool, bool)`

GetFreezeOk returns a tuple with the Freeze field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeze

`func (o *UpdateVMConfigRequest) SetFreeze(v bool)`

SetFreeze sets Freeze field to given value.

### HasFreeze

`func (o *UpdateVMConfigRequest) HasFreeze() bool`

HasFreeze returns a boolean if a field has been set.

### GetHookscript

`func (o *UpdateVMConfigRequest) GetHookscript() string`

GetHookscript returns the Hookscript field if non-nil, zero value otherwise.

### GetHookscriptOk

`func (o *UpdateVMConfigRequest) GetHookscriptOk() (*string, bool)`

GetHookscriptOk returns a tuple with the Hookscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookscript

`func (o *UpdateVMConfigRequest) SetHookscript(v string)`

SetHookscript sets Hookscript field to given value.

### HasHookscript

`func (o *UpdateVMConfigRequest) HasHookscript() bool`

HasHookscript returns a boolean if a field has been set.

### GetHostpci0

`func (o *UpdateVMConfigRequest) GetHostpci0() string`

GetHostpci0 returns the Hostpci0 field if non-nil, zero value otherwise.

### GetHostpci0Ok

`func (o *UpdateVMConfigRequest) GetHostpci0Ok() (*string, bool)`

GetHostpci0Ok returns a tuple with the Hostpci0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci0

`func (o *UpdateVMConfigRequest) SetHostpci0(v string)`

SetHostpci0 sets Hostpci0 field to given value.

### HasHostpci0

`func (o *UpdateVMConfigRequest) HasHostpci0() bool`

HasHostpci0 returns a boolean if a field has been set.

### GetHostpci1

`func (o *UpdateVMConfigRequest) GetHostpci1() string`

GetHostpci1 returns the Hostpci1 field if non-nil, zero value otherwise.

### GetHostpci1Ok

`func (o *UpdateVMConfigRequest) GetHostpci1Ok() (*string, bool)`

GetHostpci1Ok returns a tuple with the Hostpci1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci1

`func (o *UpdateVMConfigRequest) SetHostpci1(v string)`

SetHostpci1 sets Hostpci1 field to given value.

### HasHostpci1

`func (o *UpdateVMConfigRequest) HasHostpci1() bool`

HasHostpci1 returns a boolean if a field has been set.

### GetHostpci2

`func (o *UpdateVMConfigRequest) GetHostpci2() string`

GetHostpci2 returns the Hostpci2 field if non-nil, zero value otherwise.

### GetHostpci2Ok

`func (o *UpdateVMConfigRequest) GetHostpci2Ok() (*string, bool)`

GetHostpci2Ok returns a tuple with the Hostpci2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci2

`func (o *UpdateVMConfigRequest) SetHostpci2(v string)`

SetHostpci2 sets Hostpci2 field to given value.

### HasHostpci2

`func (o *UpdateVMConfigRequest) HasHostpci2() bool`

HasHostpci2 returns a boolean if a field has been set.

### GetHostpci3

`func (o *UpdateVMConfigRequest) GetHostpci3() string`

GetHostpci3 returns the Hostpci3 field if non-nil, zero value otherwise.

### GetHostpci3Ok

`func (o *UpdateVMConfigRequest) GetHostpci3Ok() (*string, bool)`

GetHostpci3Ok returns a tuple with the Hostpci3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci3

`func (o *UpdateVMConfigRequest) SetHostpci3(v string)`

SetHostpci3 sets Hostpci3 field to given value.

### HasHostpci3

`func (o *UpdateVMConfigRequest) HasHostpci3() bool`

HasHostpci3 returns a boolean if a field has been set.

### GetHostpci4

`func (o *UpdateVMConfigRequest) GetHostpci4() string`

GetHostpci4 returns the Hostpci4 field if non-nil, zero value otherwise.

### GetHostpci4Ok

`func (o *UpdateVMConfigRequest) GetHostpci4Ok() (*string, bool)`

GetHostpci4Ok returns a tuple with the Hostpci4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci4

`func (o *UpdateVMConfigRequest) SetHostpci4(v string)`

SetHostpci4 sets Hostpci4 field to given value.

### HasHostpci4

`func (o *UpdateVMConfigRequest) HasHostpci4() bool`

HasHostpci4 returns a boolean if a field has been set.

### GetHostpci5

`func (o *UpdateVMConfigRequest) GetHostpci5() string`

GetHostpci5 returns the Hostpci5 field if non-nil, zero value otherwise.

### GetHostpci5Ok

`func (o *UpdateVMConfigRequest) GetHostpci5Ok() (*string, bool)`

GetHostpci5Ok returns a tuple with the Hostpci5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci5

`func (o *UpdateVMConfigRequest) SetHostpci5(v string)`

SetHostpci5 sets Hostpci5 field to given value.

### HasHostpci5

`func (o *UpdateVMConfigRequest) HasHostpci5() bool`

HasHostpci5 returns a boolean if a field has been set.

### GetHostpci6

`func (o *UpdateVMConfigRequest) GetHostpci6() string`

GetHostpci6 returns the Hostpci6 field if non-nil, zero value otherwise.

### GetHostpci6Ok

`func (o *UpdateVMConfigRequest) GetHostpci6Ok() (*string, bool)`

GetHostpci6Ok returns a tuple with the Hostpci6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci6

`func (o *UpdateVMConfigRequest) SetHostpci6(v string)`

SetHostpci6 sets Hostpci6 field to given value.

### HasHostpci6

`func (o *UpdateVMConfigRequest) HasHostpci6() bool`

HasHostpci6 returns a boolean if a field has been set.

### GetHostpci7

`func (o *UpdateVMConfigRequest) GetHostpci7() string`

GetHostpci7 returns the Hostpci7 field if non-nil, zero value otherwise.

### GetHostpci7Ok

`func (o *UpdateVMConfigRequest) GetHostpci7Ok() (*string, bool)`

GetHostpci7Ok returns a tuple with the Hostpci7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci7

`func (o *UpdateVMConfigRequest) SetHostpci7(v string)`

SetHostpci7 sets Hostpci7 field to given value.

### HasHostpci7

`func (o *UpdateVMConfigRequest) HasHostpci7() bool`

HasHostpci7 returns a boolean if a field has been set.

### GetHostpci8

`func (o *UpdateVMConfigRequest) GetHostpci8() string`

GetHostpci8 returns the Hostpci8 field if non-nil, zero value otherwise.

### GetHostpci8Ok

`func (o *UpdateVMConfigRequest) GetHostpci8Ok() (*string, bool)`

GetHostpci8Ok returns a tuple with the Hostpci8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci8

`func (o *UpdateVMConfigRequest) SetHostpci8(v string)`

SetHostpci8 sets Hostpci8 field to given value.

### HasHostpci8

`func (o *UpdateVMConfigRequest) HasHostpci8() bool`

HasHostpci8 returns a boolean if a field has been set.

### GetHostpci9

`func (o *UpdateVMConfigRequest) GetHostpci9() string`

GetHostpci9 returns the Hostpci9 field if non-nil, zero value otherwise.

### GetHostpci9Ok

`func (o *UpdateVMConfigRequest) GetHostpci9Ok() (*string, bool)`

GetHostpci9Ok returns a tuple with the Hostpci9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci9

`func (o *UpdateVMConfigRequest) SetHostpci9(v string)`

SetHostpci9 sets Hostpci9 field to given value.

### HasHostpci9

`func (o *UpdateVMConfigRequest) HasHostpci9() bool`

HasHostpci9 returns a boolean if a field has been set.

### GetHostpci10

`func (o *UpdateVMConfigRequest) GetHostpci10() string`

GetHostpci10 returns the Hostpci10 field if non-nil, zero value otherwise.

### GetHostpci10Ok

`func (o *UpdateVMConfigRequest) GetHostpci10Ok() (*string, bool)`

GetHostpci10Ok returns a tuple with the Hostpci10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci10

`func (o *UpdateVMConfigRequest) SetHostpci10(v string)`

SetHostpci10 sets Hostpci10 field to given value.

### HasHostpci10

`func (o *UpdateVMConfigRequest) HasHostpci10() bool`

HasHostpci10 returns a boolean if a field has been set.

### GetHostpci11

`func (o *UpdateVMConfigRequest) GetHostpci11() string`

GetHostpci11 returns the Hostpci11 field if non-nil, zero value otherwise.

### GetHostpci11Ok

`func (o *UpdateVMConfigRequest) GetHostpci11Ok() (*string, bool)`

GetHostpci11Ok returns a tuple with the Hostpci11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci11

`func (o *UpdateVMConfigRequest) SetHostpci11(v string)`

SetHostpci11 sets Hostpci11 field to given value.

### HasHostpci11

`func (o *UpdateVMConfigRequest) HasHostpci11() bool`

HasHostpci11 returns a boolean if a field has been set.

### GetHostpci12

`func (o *UpdateVMConfigRequest) GetHostpci12() string`

GetHostpci12 returns the Hostpci12 field if non-nil, zero value otherwise.

### GetHostpci12Ok

`func (o *UpdateVMConfigRequest) GetHostpci12Ok() (*string, bool)`

GetHostpci12Ok returns a tuple with the Hostpci12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci12

`func (o *UpdateVMConfigRequest) SetHostpci12(v string)`

SetHostpci12 sets Hostpci12 field to given value.

### HasHostpci12

`func (o *UpdateVMConfigRequest) HasHostpci12() bool`

HasHostpci12 returns a boolean if a field has been set.

### GetHostpci13

`func (o *UpdateVMConfigRequest) GetHostpci13() string`

GetHostpci13 returns the Hostpci13 field if non-nil, zero value otherwise.

### GetHostpci13Ok

`func (o *UpdateVMConfigRequest) GetHostpci13Ok() (*string, bool)`

GetHostpci13Ok returns a tuple with the Hostpci13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci13

`func (o *UpdateVMConfigRequest) SetHostpci13(v string)`

SetHostpci13 sets Hostpci13 field to given value.

### HasHostpci13

`func (o *UpdateVMConfigRequest) HasHostpci13() bool`

HasHostpci13 returns a boolean if a field has been set.

### GetHostpci14

`func (o *UpdateVMConfigRequest) GetHostpci14() string`

GetHostpci14 returns the Hostpci14 field if non-nil, zero value otherwise.

### GetHostpci14Ok

`func (o *UpdateVMConfigRequest) GetHostpci14Ok() (*string, bool)`

GetHostpci14Ok returns a tuple with the Hostpci14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci14

`func (o *UpdateVMConfigRequest) SetHostpci14(v string)`

SetHostpci14 sets Hostpci14 field to given value.

### HasHostpci14

`func (o *UpdateVMConfigRequest) HasHostpci14() bool`

HasHostpci14 returns a boolean if a field has been set.

### GetHostpci15

`func (o *UpdateVMConfigRequest) GetHostpci15() string`

GetHostpci15 returns the Hostpci15 field if non-nil, zero value otherwise.

### GetHostpci15Ok

`func (o *UpdateVMConfigRequest) GetHostpci15Ok() (*string, bool)`

GetHostpci15Ok returns a tuple with the Hostpci15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci15

`func (o *UpdateVMConfigRequest) SetHostpci15(v string)`

SetHostpci15 sets Hostpci15 field to given value.

### HasHostpci15

`func (o *UpdateVMConfigRequest) HasHostpci15() bool`

HasHostpci15 returns a boolean if a field has been set.

### GetHostpci16

`func (o *UpdateVMConfigRequest) GetHostpci16() string`

GetHostpci16 returns the Hostpci16 field if non-nil, zero value otherwise.

### GetHostpci16Ok

`func (o *UpdateVMConfigRequest) GetHostpci16Ok() (*string, bool)`

GetHostpci16Ok returns a tuple with the Hostpci16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci16

`func (o *UpdateVMConfigRequest) SetHostpci16(v string)`

SetHostpci16 sets Hostpci16 field to given value.

### HasHostpci16

`func (o *UpdateVMConfigRequest) HasHostpci16() bool`

HasHostpci16 returns a boolean if a field has been set.

### GetHostpci17

`func (o *UpdateVMConfigRequest) GetHostpci17() string`

GetHostpci17 returns the Hostpci17 field if non-nil, zero value otherwise.

### GetHostpci17Ok

`func (o *UpdateVMConfigRequest) GetHostpci17Ok() (*string, bool)`

GetHostpci17Ok returns a tuple with the Hostpci17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci17

`func (o *UpdateVMConfigRequest) SetHostpci17(v string)`

SetHostpci17 sets Hostpci17 field to given value.

### HasHostpci17

`func (o *UpdateVMConfigRequest) HasHostpci17() bool`

HasHostpci17 returns a boolean if a field has been set.

### GetHostpci18

`func (o *UpdateVMConfigRequest) GetHostpci18() string`

GetHostpci18 returns the Hostpci18 field if non-nil, zero value otherwise.

### GetHostpci18Ok

`func (o *UpdateVMConfigRequest) GetHostpci18Ok() (*string, bool)`

GetHostpci18Ok returns a tuple with the Hostpci18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci18

`func (o *UpdateVMConfigRequest) SetHostpci18(v string)`

SetHostpci18 sets Hostpci18 field to given value.

### HasHostpci18

`func (o *UpdateVMConfigRequest) HasHostpci18() bool`

HasHostpci18 returns a boolean if a field has been set.

### GetHostpci19

`func (o *UpdateVMConfigRequest) GetHostpci19() string`

GetHostpci19 returns the Hostpci19 field if non-nil, zero value otherwise.

### GetHostpci19Ok

`func (o *UpdateVMConfigRequest) GetHostpci19Ok() (*string, bool)`

GetHostpci19Ok returns a tuple with the Hostpci19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci19

`func (o *UpdateVMConfigRequest) SetHostpci19(v string)`

SetHostpci19 sets Hostpci19 field to given value.

### HasHostpci19

`func (o *UpdateVMConfigRequest) HasHostpci19() bool`

HasHostpci19 returns a boolean if a field has been set.

### GetHostpci20

`func (o *UpdateVMConfigRequest) GetHostpci20() string`

GetHostpci20 returns the Hostpci20 field if non-nil, zero value otherwise.

### GetHostpci20Ok

`func (o *UpdateVMConfigRequest) GetHostpci20Ok() (*string, bool)`

GetHostpci20Ok returns a tuple with the Hostpci20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci20

`func (o *UpdateVMConfigRequest) SetHostpci20(v string)`

SetHostpci20 sets Hostpci20 field to given value.

### HasHostpci20

`func (o *UpdateVMConfigRequest) HasHostpci20() bool`

HasHostpci20 returns a boolean if a field has been set.

### GetHostpci21

`func (o *UpdateVMConfigRequest) GetHostpci21() string`

GetHostpci21 returns the Hostpci21 field if non-nil, zero value otherwise.

### GetHostpci21Ok

`func (o *UpdateVMConfigRequest) GetHostpci21Ok() (*string, bool)`

GetHostpci21Ok returns a tuple with the Hostpci21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci21

`func (o *UpdateVMConfigRequest) SetHostpci21(v string)`

SetHostpci21 sets Hostpci21 field to given value.

### HasHostpci21

`func (o *UpdateVMConfigRequest) HasHostpci21() bool`

HasHostpci21 returns a boolean if a field has been set.

### GetHostpci22

`func (o *UpdateVMConfigRequest) GetHostpci22() string`

GetHostpci22 returns the Hostpci22 field if non-nil, zero value otherwise.

### GetHostpci22Ok

`func (o *UpdateVMConfigRequest) GetHostpci22Ok() (*string, bool)`

GetHostpci22Ok returns a tuple with the Hostpci22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci22

`func (o *UpdateVMConfigRequest) SetHostpci22(v string)`

SetHostpci22 sets Hostpci22 field to given value.

### HasHostpci22

`func (o *UpdateVMConfigRequest) HasHostpci22() bool`

HasHostpci22 returns a boolean if a field has been set.

### GetHostpci23

`func (o *UpdateVMConfigRequest) GetHostpci23() string`

GetHostpci23 returns the Hostpci23 field if non-nil, zero value otherwise.

### GetHostpci23Ok

`func (o *UpdateVMConfigRequest) GetHostpci23Ok() (*string, bool)`

GetHostpci23Ok returns a tuple with the Hostpci23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci23

`func (o *UpdateVMConfigRequest) SetHostpci23(v string)`

SetHostpci23 sets Hostpci23 field to given value.

### HasHostpci23

`func (o *UpdateVMConfigRequest) HasHostpci23() bool`

HasHostpci23 returns a boolean if a field has been set.

### GetHostpci24

`func (o *UpdateVMConfigRequest) GetHostpci24() string`

GetHostpci24 returns the Hostpci24 field if non-nil, zero value otherwise.

### GetHostpci24Ok

`func (o *UpdateVMConfigRequest) GetHostpci24Ok() (*string, bool)`

GetHostpci24Ok returns a tuple with the Hostpci24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci24

`func (o *UpdateVMConfigRequest) SetHostpci24(v string)`

SetHostpci24 sets Hostpci24 field to given value.

### HasHostpci24

`func (o *UpdateVMConfigRequest) HasHostpci24() bool`

HasHostpci24 returns a boolean if a field has been set.

### GetHostpci25

`func (o *UpdateVMConfigRequest) GetHostpci25() string`

GetHostpci25 returns the Hostpci25 field if non-nil, zero value otherwise.

### GetHostpci25Ok

`func (o *UpdateVMConfigRequest) GetHostpci25Ok() (*string, bool)`

GetHostpci25Ok returns a tuple with the Hostpci25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci25

`func (o *UpdateVMConfigRequest) SetHostpci25(v string)`

SetHostpci25 sets Hostpci25 field to given value.

### HasHostpci25

`func (o *UpdateVMConfigRequest) HasHostpci25() bool`

HasHostpci25 returns a boolean if a field has been set.

### GetHostpci26

`func (o *UpdateVMConfigRequest) GetHostpci26() string`

GetHostpci26 returns the Hostpci26 field if non-nil, zero value otherwise.

### GetHostpci26Ok

`func (o *UpdateVMConfigRequest) GetHostpci26Ok() (*string, bool)`

GetHostpci26Ok returns a tuple with the Hostpci26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci26

`func (o *UpdateVMConfigRequest) SetHostpci26(v string)`

SetHostpci26 sets Hostpci26 field to given value.

### HasHostpci26

`func (o *UpdateVMConfigRequest) HasHostpci26() bool`

HasHostpci26 returns a boolean if a field has been set.

### GetHostpci27

`func (o *UpdateVMConfigRequest) GetHostpci27() string`

GetHostpci27 returns the Hostpci27 field if non-nil, zero value otherwise.

### GetHostpci27Ok

`func (o *UpdateVMConfigRequest) GetHostpci27Ok() (*string, bool)`

GetHostpci27Ok returns a tuple with the Hostpci27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci27

`func (o *UpdateVMConfigRequest) SetHostpci27(v string)`

SetHostpci27 sets Hostpci27 field to given value.

### HasHostpci27

`func (o *UpdateVMConfigRequest) HasHostpci27() bool`

HasHostpci27 returns a boolean if a field has been set.

### GetHostpci28

`func (o *UpdateVMConfigRequest) GetHostpci28() string`

GetHostpci28 returns the Hostpci28 field if non-nil, zero value otherwise.

### GetHostpci28Ok

`func (o *UpdateVMConfigRequest) GetHostpci28Ok() (*string, bool)`

GetHostpci28Ok returns a tuple with the Hostpci28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci28

`func (o *UpdateVMConfigRequest) SetHostpci28(v string)`

SetHostpci28 sets Hostpci28 field to given value.

### HasHostpci28

`func (o *UpdateVMConfigRequest) HasHostpci28() bool`

HasHostpci28 returns a boolean if a field has been set.

### GetHostpci29

`func (o *UpdateVMConfigRequest) GetHostpci29() string`

GetHostpci29 returns the Hostpci29 field if non-nil, zero value otherwise.

### GetHostpci29Ok

`func (o *UpdateVMConfigRequest) GetHostpci29Ok() (*string, bool)`

GetHostpci29Ok returns a tuple with the Hostpci29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostpci29

`func (o *UpdateVMConfigRequest) SetHostpci29(v string)`

SetHostpci29 sets Hostpci29 field to given value.

### HasHostpci29

`func (o *UpdateVMConfigRequest) HasHostpci29() bool`

HasHostpci29 returns a boolean if a field has been set.

### GetHotplug

`func (o *UpdateVMConfigRequest) GetHotplug() string`

GetHotplug returns the Hotplug field if non-nil, zero value otherwise.

### GetHotplugOk

`func (o *UpdateVMConfigRequest) GetHotplugOk() (*string, bool)`

GetHotplugOk returns a tuple with the Hotplug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotplug

`func (o *UpdateVMConfigRequest) SetHotplug(v string)`

SetHotplug sets Hotplug field to given value.

### HasHotplug

`func (o *UpdateVMConfigRequest) HasHotplug() bool`

HasHotplug returns a boolean if a field has been set.

### GetHugepages

`func (o *UpdateVMConfigRequest) GetHugepages() string`

GetHugepages returns the Hugepages field if non-nil, zero value otherwise.

### GetHugepagesOk

`func (o *UpdateVMConfigRequest) GetHugepagesOk() (*string, bool)`

GetHugepagesOk returns a tuple with the Hugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugepages

`func (o *UpdateVMConfigRequest) SetHugepages(v string)`

SetHugepages sets Hugepages field to given value.

### HasHugepages

`func (o *UpdateVMConfigRequest) HasHugepages() bool`

HasHugepages returns a boolean if a field has been set.

### GetIde0

`func (o *UpdateVMConfigRequest) GetIde0() string`

GetIde0 returns the Ide0 field if non-nil, zero value otherwise.

### GetIde0Ok

`func (o *UpdateVMConfigRequest) GetIde0Ok() (*string, bool)`

GetIde0Ok returns a tuple with the Ide0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde0

`func (o *UpdateVMConfigRequest) SetIde0(v string)`

SetIde0 sets Ide0 field to given value.

### HasIde0

`func (o *UpdateVMConfigRequest) HasIde0() bool`

HasIde0 returns a boolean if a field has been set.

### GetIde1

`func (o *UpdateVMConfigRequest) GetIde1() string`

GetIde1 returns the Ide1 field if non-nil, zero value otherwise.

### GetIde1Ok

`func (o *UpdateVMConfigRequest) GetIde1Ok() (*string, bool)`

GetIde1Ok returns a tuple with the Ide1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde1

`func (o *UpdateVMConfigRequest) SetIde1(v string)`

SetIde1 sets Ide1 field to given value.

### HasIde1

`func (o *UpdateVMConfigRequest) HasIde1() bool`

HasIde1 returns a boolean if a field has been set.

### GetIde2

`func (o *UpdateVMConfigRequest) GetIde2() string`

GetIde2 returns the Ide2 field if non-nil, zero value otherwise.

### GetIde2Ok

`func (o *UpdateVMConfigRequest) GetIde2Ok() (*string, bool)`

GetIde2Ok returns a tuple with the Ide2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde2

`func (o *UpdateVMConfigRequest) SetIde2(v string)`

SetIde2 sets Ide2 field to given value.

### HasIde2

`func (o *UpdateVMConfigRequest) HasIde2() bool`

HasIde2 returns a boolean if a field has been set.

### GetIde3

`func (o *UpdateVMConfigRequest) GetIde3() string`

GetIde3 returns the Ide3 field if non-nil, zero value otherwise.

### GetIde3Ok

`func (o *UpdateVMConfigRequest) GetIde3Ok() (*string, bool)`

GetIde3Ok returns a tuple with the Ide3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIde3

`func (o *UpdateVMConfigRequest) SetIde3(v string)`

SetIde3 sets Ide3 field to given value.

### HasIde3

`func (o *UpdateVMConfigRequest) HasIde3() bool`

HasIde3 returns a boolean if a field has been set.

### GetImportWorkingStorage

`func (o *UpdateVMConfigRequest) GetImportWorkingStorage() string`

GetImportWorkingStorage returns the ImportWorkingStorage field if non-nil, zero value otherwise.

### GetImportWorkingStorageOk

`func (o *UpdateVMConfigRequest) GetImportWorkingStorageOk() (*string, bool)`

GetImportWorkingStorageOk returns a tuple with the ImportWorkingStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportWorkingStorage

`func (o *UpdateVMConfigRequest) SetImportWorkingStorage(v string)`

SetImportWorkingStorage sets ImportWorkingStorage field to given value.

### HasImportWorkingStorage

`func (o *UpdateVMConfigRequest) HasImportWorkingStorage() bool`

HasImportWorkingStorage returns a boolean if a field has been set.

### GetIpconfig0

`func (o *UpdateVMConfigRequest) GetIpconfig0() string`

GetIpconfig0 returns the Ipconfig0 field if non-nil, zero value otherwise.

### GetIpconfig0Ok

`func (o *UpdateVMConfigRequest) GetIpconfig0Ok() (*string, bool)`

GetIpconfig0Ok returns a tuple with the Ipconfig0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig0

`func (o *UpdateVMConfigRequest) SetIpconfig0(v string)`

SetIpconfig0 sets Ipconfig0 field to given value.

### HasIpconfig0

`func (o *UpdateVMConfigRequest) HasIpconfig0() bool`

HasIpconfig0 returns a boolean if a field has been set.

### GetIpconfig1

`func (o *UpdateVMConfigRequest) GetIpconfig1() string`

GetIpconfig1 returns the Ipconfig1 field if non-nil, zero value otherwise.

### GetIpconfig1Ok

`func (o *UpdateVMConfigRequest) GetIpconfig1Ok() (*string, bool)`

GetIpconfig1Ok returns a tuple with the Ipconfig1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig1

`func (o *UpdateVMConfigRequest) SetIpconfig1(v string)`

SetIpconfig1 sets Ipconfig1 field to given value.

### HasIpconfig1

`func (o *UpdateVMConfigRequest) HasIpconfig1() bool`

HasIpconfig1 returns a boolean if a field has been set.

### GetIpconfig2

`func (o *UpdateVMConfigRequest) GetIpconfig2() string`

GetIpconfig2 returns the Ipconfig2 field if non-nil, zero value otherwise.

### GetIpconfig2Ok

`func (o *UpdateVMConfigRequest) GetIpconfig2Ok() (*string, bool)`

GetIpconfig2Ok returns a tuple with the Ipconfig2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig2

`func (o *UpdateVMConfigRequest) SetIpconfig2(v string)`

SetIpconfig2 sets Ipconfig2 field to given value.

### HasIpconfig2

`func (o *UpdateVMConfigRequest) HasIpconfig2() bool`

HasIpconfig2 returns a boolean if a field has been set.

### GetIpconfig3

`func (o *UpdateVMConfigRequest) GetIpconfig3() string`

GetIpconfig3 returns the Ipconfig3 field if non-nil, zero value otherwise.

### GetIpconfig3Ok

`func (o *UpdateVMConfigRequest) GetIpconfig3Ok() (*string, bool)`

GetIpconfig3Ok returns a tuple with the Ipconfig3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig3

`func (o *UpdateVMConfigRequest) SetIpconfig3(v string)`

SetIpconfig3 sets Ipconfig3 field to given value.

### HasIpconfig3

`func (o *UpdateVMConfigRequest) HasIpconfig3() bool`

HasIpconfig3 returns a boolean if a field has been set.

### GetIpconfig4

`func (o *UpdateVMConfigRequest) GetIpconfig4() string`

GetIpconfig4 returns the Ipconfig4 field if non-nil, zero value otherwise.

### GetIpconfig4Ok

`func (o *UpdateVMConfigRequest) GetIpconfig4Ok() (*string, bool)`

GetIpconfig4Ok returns a tuple with the Ipconfig4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig4

`func (o *UpdateVMConfigRequest) SetIpconfig4(v string)`

SetIpconfig4 sets Ipconfig4 field to given value.

### HasIpconfig4

`func (o *UpdateVMConfigRequest) HasIpconfig4() bool`

HasIpconfig4 returns a boolean if a field has been set.

### GetIpconfig5

`func (o *UpdateVMConfigRequest) GetIpconfig5() string`

GetIpconfig5 returns the Ipconfig5 field if non-nil, zero value otherwise.

### GetIpconfig5Ok

`func (o *UpdateVMConfigRequest) GetIpconfig5Ok() (*string, bool)`

GetIpconfig5Ok returns a tuple with the Ipconfig5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig5

`func (o *UpdateVMConfigRequest) SetIpconfig5(v string)`

SetIpconfig5 sets Ipconfig5 field to given value.

### HasIpconfig5

`func (o *UpdateVMConfigRequest) HasIpconfig5() bool`

HasIpconfig5 returns a boolean if a field has been set.

### GetIpconfig6

`func (o *UpdateVMConfigRequest) GetIpconfig6() string`

GetIpconfig6 returns the Ipconfig6 field if non-nil, zero value otherwise.

### GetIpconfig6Ok

`func (o *UpdateVMConfigRequest) GetIpconfig6Ok() (*string, bool)`

GetIpconfig6Ok returns a tuple with the Ipconfig6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig6

`func (o *UpdateVMConfigRequest) SetIpconfig6(v string)`

SetIpconfig6 sets Ipconfig6 field to given value.

### HasIpconfig6

`func (o *UpdateVMConfigRequest) HasIpconfig6() bool`

HasIpconfig6 returns a boolean if a field has been set.

### GetIpconfig7

`func (o *UpdateVMConfigRequest) GetIpconfig7() string`

GetIpconfig7 returns the Ipconfig7 field if non-nil, zero value otherwise.

### GetIpconfig7Ok

`func (o *UpdateVMConfigRequest) GetIpconfig7Ok() (*string, bool)`

GetIpconfig7Ok returns a tuple with the Ipconfig7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig7

`func (o *UpdateVMConfigRequest) SetIpconfig7(v string)`

SetIpconfig7 sets Ipconfig7 field to given value.

### HasIpconfig7

`func (o *UpdateVMConfigRequest) HasIpconfig7() bool`

HasIpconfig7 returns a boolean if a field has been set.

### GetIpconfig8

`func (o *UpdateVMConfigRequest) GetIpconfig8() string`

GetIpconfig8 returns the Ipconfig8 field if non-nil, zero value otherwise.

### GetIpconfig8Ok

`func (o *UpdateVMConfigRequest) GetIpconfig8Ok() (*string, bool)`

GetIpconfig8Ok returns a tuple with the Ipconfig8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig8

`func (o *UpdateVMConfigRequest) SetIpconfig8(v string)`

SetIpconfig8 sets Ipconfig8 field to given value.

### HasIpconfig8

`func (o *UpdateVMConfigRequest) HasIpconfig8() bool`

HasIpconfig8 returns a boolean if a field has been set.

### GetIpconfig9

`func (o *UpdateVMConfigRequest) GetIpconfig9() string`

GetIpconfig9 returns the Ipconfig9 field if non-nil, zero value otherwise.

### GetIpconfig9Ok

`func (o *UpdateVMConfigRequest) GetIpconfig9Ok() (*string, bool)`

GetIpconfig9Ok returns a tuple with the Ipconfig9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig9

`func (o *UpdateVMConfigRequest) SetIpconfig9(v string)`

SetIpconfig9 sets Ipconfig9 field to given value.

### HasIpconfig9

`func (o *UpdateVMConfigRequest) HasIpconfig9() bool`

HasIpconfig9 returns a boolean if a field has been set.

### GetIpconfig10

`func (o *UpdateVMConfigRequest) GetIpconfig10() string`

GetIpconfig10 returns the Ipconfig10 field if non-nil, zero value otherwise.

### GetIpconfig10Ok

`func (o *UpdateVMConfigRequest) GetIpconfig10Ok() (*string, bool)`

GetIpconfig10Ok returns a tuple with the Ipconfig10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig10

`func (o *UpdateVMConfigRequest) SetIpconfig10(v string)`

SetIpconfig10 sets Ipconfig10 field to given value.

### HasIpconfig10

`func (o *UpdateVMConfigRequest) HasIpconfig10() bool`

HasIpconfig10 returns a boolean if a field has been set.

### GetIpconfig11

`func (o *UpdateVMConfigRequest) GetIpconfig11() string`

GetIpconfig11 returns the Ipconfig11 field if non-nil, zero value otherwise.

### GetIpconfig11Ok

`func (o *UpdateVMConfigRequest) GetIpconfig11Ok() (*string, bool)`

GetIpconfig11Ok returns a tuple with the Ipconfig11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig11

`func (o *UpdateVMConfigRequest) SetIpconfig11(v string)`

SetIpconfig11 sets Ipconfig11 field to given value.

### HasIpconfig11

`func (o *UpdateVMConfigRequest) HasIpconfig11() bool`

HasIpconfig11 returns a boolean if a field has been set.

### GetIpconfig12

`func (o *UpdateVMConfigRequest) GetIpconfig12() string`

GetIpconfig12 returns the Ipconfig12 field if non-nil, zero value otherwise.

### GetIpconfig12Ok

`func (o *UpdateVMConfigRequest) GetIpconfig12Ok() (*string, bool)`

GetIpconfig12Ok returns a tuple with the Ipconfig12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig12

`func (o *UpdateVMConfigRequest) SetIpconfig12(v string)`

SetIpconfig12 sets Ipconfig12 field to given value.

### HasIpconfig12

`func (o *UpdateVMConfigRequest) HasIpconfig12() bool`

HasIpconfig12 returns a boolean if a field has been set.

### GetIpconfig13

`func (o *UpdateVMConfigRequest) GetIpconfig13() string`

GetIpconfig13 returns the Ipconfig13 field if non-nil, zero value otherwise.

### GetIpconfig13Ok

`func (o *UpdateVMConfigRequest) GetIpconfig13Ok() (*string, bool)`

GetIpconfig13Ok returns a tuple with the Ipconfig13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig13

`func (o *UpdateVMConfigRequest) SetIpconfig13(v string)`

SetIpconfig13 sets Ipconfig13 field to given value.

### HasIpconfig13

`func (o *UpdateVMConfigRequest) HasIpconfig13() bool`

HasIpconfig13 returns a boolean if a field has been set.

### GetIpconfig14

`func (o *UpdateVMConfigRequest) GetIpconfig14() string`

GetIpconfig14 returns the Ipconfig14 field if non-nil, zero value otherwise.

### GetIpconfig14Ok

`func (o *UpdateVMConfigRequest) GetIpconfig14Ok() (*string, bool)`

GetIpconfig14Ok returns a tuple with the Ipconfig14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig14

`func (o *UpdateVMConfigRequest) SetIpconfig14(v string)`

SetIpconfig14 sets Ipconfig14 field to given value.

### HasIpconfig14

`func (o *UpdateVMConfigRequest) HasIpconfig14() bool`

HasIpconfig14 returns a boolean if a field has been set.

### GetIpconfig15

`func (o *UpdateVMConfigRequest) GetIpconfig15() string`

GetIpconfig15 returns the Ipconfig15 field if non-nil, zero value otherwise.

### GetIpconfig15Ok

`func (o *UpdateVMConfigRequest) GetIpconfig15Ok() (*string, bool)`

GetIpconfig15Ok returns a tuple with the Ipconfig15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig15

`func (o *UpdateVMConfigRequest) SetIpconfig15(v string)`

SetIpconfig15 sets Ipconfig15 field to given value.

### HasIpconfig15

`func (o *UpdateVMConfigRequest) HasIpconfig15() bool`

HasIpconfig15 returns a boolean if a field has been set.

### GetIpconfig16

`func (o *UpdateVMConfigRequest) GetIpconfig16() string`

GetIpconfig16 returns the Ipconfig16 field if non-nil, zero value otherwise.

### GetIpconfig16Ok

`func (o *UpdateVMConfigRequest) GetIpconfig16Ok() (*string, bool)`

GetIpconfig16Ok returns a tuple with the Ipconfig16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig16

`func (o *UpdateVMConfigRequest) SetIpconfig16(v string)`

SetIpconfig16 sets Ipconfig16 field to given value.

### HasIpconfig16

`func (o *UpdateVMConfigRequest) HasIpconfig16() bool`

HasIpconfig16 returns a boolean if a field has been set.

### GetIpconfig17

`func (o *UpdateVMConfigRequest) GetIpconfig17() string`

GetIpconfig17 returns the Ipconfig17 field if non-nil, zero value otherwise.

### GetIpconfig17Ok

`func (o *UpdateVMConfigRequest) GetIpconfig17Ok() (*string, bool)`

GetIpconfig17Ok returns a tuple with the Ipconfig17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig17

`func (o *UpdateVMConfigRequest) SetIpconfig17(v string)`

SetIpconfig17 sets Ipconfig17 field to given value.

### HasIpconfig17

`func (o *UpdateVMConfigRequest) HasIpconfig17() bool`

HasIpconfig17 returns a boolean if a field has been set.

### GetIpconfig18

`func (o *UpdateVMConfigRequest) GetIpconfig18() string`

GetIpconfig18 returns the Ipconfig18 field if non-nil, zero value otherwise.

### GetIpconfig18Ok

`func (o *UpdateVMConfigRequest) GetIpconfig18Ok() (*string, bool)`

GetIpconfig18Ok returns a tuple with the Ipconfig18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig18

`func (o *UpdateVMConfigRequest) SetIpconfig18(v string)`

SetIpconfig18 sets Ipconfig18 field to given value.

### HasIpconfig18

`func (o *UpdateVMConfigRequest) HasIpconfig18() bool`

HasIpconfig18 returns a boolean if a field has been set.

### GetIpconfig19

`func (o *UpdateVMConfigRequest) GetIpconfig19() string`

GetIpconfig19 returns the Ipconfig19 field if non-nil, zero value otherwise.

### GetIpconfig19Ok

`func (o *UpdateVMConfigRequest) GetIpconfig19Ok() (*string, bool)`

GetIpconfig19Ok returns a tuple with the Ipconfig19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig19

`func (o *UpdateVMConfigRequest) SetIpconfig19(v string)`

SetIpconfig19 sets Ipconfig19 field to given value.

### HasIpconfig19

`func (o *UpdateVMConfigRequest) HasIpconfig19() bool`

HasIpconfig19 returns a boolean if a field has been set.

### GetIpconfig20

`func (o *UpdateVMConfigRequest) GetIpconfig20() string`

GetIpconfig20 returns the Ipconfig20 field if non-nil, zero value otherwise.

### GetIpconfig20Ok

`func (o *UpdateVMConfigRequest) GetIpconfig20Ok() (*string, bool)`

GetIpconfig20Ok returns a tuple with the Ipconfig20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig20

`func (o *UpdateVMConfigRequest) SetIpconfig20(v string)`

SetIpconfig20 sets Ipconfig20 field to given value.

### HasIpconfig20

`func (o *UpdateVMConfigRequest) HasIpconfig20() bool`

HasIpconfig20 returns a boolean if a field has been set.

### GetIpconfig21

`func (o *UpdateVMConfigRequest) GetIpconfig21() string`

GetIpconfig21 returns the Ipconfig21 field if non-nil, zero value otherwise.

### GetIpconfig21Ok

`func (o *UpdateVMConfigRequest) GetIpconfig21Ok() (*string, bool)`

GetIpconfig21Ok returns a tuple with the Ipconfig21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig21

`func (o *UpdateVMConfigRequest) SetIpconfig21(v string)`

SetIpconfig21 sets Ipconfig21 field to given value.

### HasIpconfig21

`func (o *UpdateVMConfigRequest) HasIpconfig21() bool`

HasIpconfig21 returns a boolean if a field has been set.

### GetIpconfig22

`func (o *UpdateVMConfigRequest) GetIpconfig22() string`

GetIpconfig22 returns the Ipconfig22 field if non-nil, zero value otherwise.

### GetIpconfig22Ok

`func (o *UpdateVMConfigRequest) GetIpconfig22Ok() (*string, bool)`

GetIpconfig22Ok returns a tuple with the Ipconfig22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig22

`func (o *UpdateVMConfigRequest) SetIpconfig22(v string)`

SetIpconfig22 sets Ipconfig22 field to given value.

### HasIpconfig22

`func (o *UpdateVMConfigRequest) HasIpconfig22() bool`

HasIpconfig22 returns a boolean if a field has been set.

### GetIpconfig23

`func (o *UpdateVMConfigRequest) GetIpconfig23() string`

GetIpconfig23 returns the Ipconfig23 field if non-nil, zero value otherwise.

### GetIpconfig23Ok

`func (o *UpdateVMConfigRequest) GetIpconfig23Ok() (*string, bool)`

GetIpconfig23Ok returns a tuple with the Ipconfig23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig23

`func (o *UpdateVMConfigRequest) SetIpconfig23(v string)`

SetIpconfig23 sets Ipconfig23 field to given value.

### HasIpconfig23

`func (o *UpdateVMConfigRequest) HasIpconfig23() bool`

HasIpconfig23 returns a boolean if a field has been set.

### GetIpconfig24

`func (o *UpdateVMConfigRequest) GetIpconfig24() string`

GetIpconfig24 returns the Ipconfig24 field if non-nil, zero value otherwise.

### GetIpconfig24Ok

`func (o *UpdateVMConfigRequest) GetIpconfig24Ok() (*string, bool)`

GetIpconfig24Ok returns a tuple with the Ipconfig24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig24

`func (o *UpdateVMConfigRequest) SetIpconfig24(v string)`

SetIpconfig24 sets Ipconfig24 field to given value.

### HasIpconfig24

`func (o *UpdateVMConfigRequest) HasIpconfig24() bool`

HasIpconfig24 returns a boolean if a field has been set.

### GetIpconfig25

`func (o *UpdateVMConfigRequest) GetIpconfig25() string`

GetIpconfig25 returns the Ipconfig25 field if non-nil, zero value otherwise.

### GetIpconfig25Ok

`func (o *UpdateVMConfigRequest) GetIpconfig25Ok() (*string, bool)`

GetIpconfig25Ok returns a tuple with the Ipconfig25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig25

`func (o *UpdateVMConfigRequest) SetIpconfig25(v string)`

SetIpconfig25 sets Ipconfig25 field to given value.

### HasIpconfig25

`func (o *UpdateVMConfigRequest) HasIpconfig25() bool`

HasIpconfig25 returns a boolean if a field has been set.

### GetIpconfig26

`func (o *UpdateVMConfigRequest) GetIpconfig26() string`

GetIpconfig26 returns the Ipconfig26 field if non-nil, zero value otherwise.

### GetIpconfig26Ok

`func (o *UpdateVMConfigRequest) GetIpconfig26Ok() (*string, bool)`

GetIpconfig26Ok returns a tuple with the Ipconfig26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig26

`func (o *UpdateVMConfigRequest) SetIpconfig26(v string)`

SetIpconfig26 sets Ipconfig26 field to given value.

### HasIpconfig26

`func (o *UpdateVMConfigRequest) HasIpconfig26() bool`

HasIpconfig26 returns a boolean if a field has been set.

### GetIpconfig27

`func (o *UpdateVMConfigRequest) GetIpconfig27() string`

GetIpconfig27 returns the Ipconfig27 field if non-nil, zero value otherwise.

### GetIpconfig27Ok

`func (o *UpdateVMConfigRequest) GetIpconfig27Ok() (*string, bool)`

GetIpconfig27Ok returns a tuple with the Ipconfig27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig27

`func (o *UpdateVMConfigRequest) SetIpconfig27(v string)`

SetIpconfig27 sets Ipconfig27 field to given value.

### HasIpconfig27

`func (o *UpdateVMConfigRequest) HasIpconfig27() bool`

HasIpconfig27 returns a boolean if a field has been set.

### GetIpconfig28

`func (o *UpdateVMConfigRequest) GetIpconfig28() string`

GetIpconfig28 returns the Ipconfig28 field if non-nil, zero value otherwise.

### GetIpconfig28Ok

`func (o *UpdateVMConfigRequest) GetIpconfig28Ok() (*string, bool)`

GetIpconfig28Ok returns a tuple with the Ipconfig28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig28

`func (o *UpdateVMConfigRequest) SetIpconfig28(v string)`

SetIpconfig28 sets Ipconfig28 field to given value.

### HasIpconfig28

`func (o *UpdateVMConfigRequest) HasIpconfig28() bool`

HasIpconfig28 returns a boolean if a field has been set.

### GetIpconfig29

`func (o *UpdateVMConfigRequest) GetIpconfig29() string`

GetIpconfig29 returns the Ipconfig29 field if non-nil, zero value otherwise.

### GetIpconfig29Ok

`func (o *UpdateVMConfigRequest) GetIpconfig29Ok() (*string, bool)`

GetIpconfig29Ok returns a tuple with the Ipconfig29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpconfig29

`func (o *UpdateVMConfigRequest) SetIpconfig29(v string)`

SetIpconfig29 sets Ipconfig29 field to given value.

### HasIpconfig29

`func (o *UpdateVMConfigRequest) HasIpconfig29() bool`

HasIpconfig29 returns a boolean if a field has been set.

### GetIvshmem

`func (o *UpdateVMConfigRequest) GetIvshmem() string`

GetIvshmem returns the Ivshmem field if non-nil, zero value otherwise.

### GetIvshmemOk

`func (o *UpdateVMConfigRequest) GetIvshmemOk() (*string, bool)`

GetIvshmemOk returns a tuple with the Ivshmem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvshmem

`func (o *UpdateVMConfigRequest) SetIvshmem(v string)`

SetIvshmem sets Ivshmem field to given value.

### HasIvshmem

`func (o *UpdateVMConfigRequest) HasIvshmem() bool`

HasIvshmem returns a boolean if a field has been set.

### GetKeephugepages

`func (o *UpdateVMConfigRequest) GetKeephugepages() bool`

GetKeephugepages returns the Keephugepages field if non-nil, zero value otherwise.

### GetKeephugepagesOk

`func (o *UpdateVMConfigRequest) GetKeephugepagesOk() (*bool, bool)`

GetKeephugepagesOk returns a tuple with the Keephugepages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeephugepages

`func (o *UpdateVMConfigRequest) SetKeephugepages(v bool)`

SetKeephugepages sets Keephugepages field to given value.

### HasKeephugepages

`func (o *UpdateVMConfigRequest) HasKeephugepages() bool`

HasKeephugepages returns a boolean if a field has been set.

### GetKeyboard

`func (o *UpdateVMConfigRequest) GetKeyboard() string`

GetKeyboard returns the Keyboard field if non-nil, zero value otherwise.

### GetKeyboardOk

`func (o *UpdateVMConfigRequest) GetKeyboardOk() (*string, bool)`

GetKeyboardOk returns a tuple with the Keyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboard

`func (o *UpdateVMConfigRequest) SetKeyboard(v string)`

SetKeyboard sets Keyboard field to given value.

### HasKeyboard

`func (o *UpdateVMConfigRequest) HasKeyboard() bool`

HasKeyboard returns a boolean if a field has been set.

### GetKvm

`func (o *UpdateVMConfigRequest) GetKvm() bool`

GetKvm returns the Kvm field if non-nil, zero value otherwise.

### GetKvmOk

`func (o *UpdateVMConfigRequest) GetKvmOk() (*bool, bool)`

GetKvmOk returns a tuple with the Kvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvm

`func (o *UpdateVMConfigRequest) SetKvm(v bool)`

SetKvm sets Kvm field to given value.

### HasKvm

`func (o *UpdateVMConfigRequest) HasKvm() bool`

HasKvm returns a boolean if a field has been set.

### GetLocaltime

`func (o *UpdateVMConfigRequest) GetLocaltime() bool`

GetLocaltime returns the Localtime field if non-nil, zero value otherwise.

### GetLocaltimeOk

`func (o *UpdateVMConfigRequest) GetLocaltimeOk() (*bool, bool)`

GetLocaltimeOk returns a tuple with the Localtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocaltime

`func (o *UpdateVMConfigRequest) SetLocaltime(v bool)`

SetLocaltime sets Localtime field to given value.

### HasLocaltime

`func (o *UpdateVMConfigRequest) HasLocaltime() bool`

HasLocaltime returns a boolean if a field has been set.

### GetLock

`func (o *UpdateVMConfigRequest) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *UpdateVMConfigRequest) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *UpdateVMConfigRequest) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *UpdateVMConfigRequest) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMachine

`func (o *UpdateVMConfigRequest) GetMachine() string`

GetMachine returns the Machine field if non-nil, zero value otherwise.

### GetMachineOk

`func (o *UpdateVMConfigRequest) GetMachineOk() (*string, bool)`

GetMachineOk returns a tuple with the Machine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachine

`func (o *UpdateVMConfigRequest) SetMachine(v string)`

SetMachine sets Machine field to given value.

### HasMachine

`func (o *UpdateVMConfigRequest) HasMachine() bool`

HasMachine returns a boolean if a field has been set.

### GetMemory

`func (o *UpdateVMConfigRequest) GetMemory() string`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *UpdateVMConfigRequest) GetMemoryOk() (*string, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *UpdateVMConfigRequest) SetMemory(v string)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *UpdateVMConfigRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMigrateDowntime

`func (o *UpdateVMConfigRequest) GetMigrateDowntime() float32`

GetMigrateDowntime returns the MigrateDowntime field if non-nil, zero value otherwise.

### GetMigrateDowntimeOk

`func (o *UpdateVMConfigRequest) GetMigrateDowntimeOk() (*float32, bool)`

GetMigrateDowntimeOk returns a tuple with the MigrateDowntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateDowntime

`func (o *UpdateVMConfigRequest) SetMigrateDowntime(v float32)`

SetMigrateDowntime sets MigrateDowntime field to given value.

### HasMigrateDowntime

`func (o *UpdateVMConfigRequest) HasMigrateDowntime() bool`

HasMigrateDowntime returns a boolean if a field has been set.

### GetMigrateSpeed

`func (o *UpdateVMConfigRequest) GetMigrateSpeed() int64`

GetMigrateSpeed returns the MigrateSpeed field if non-nil, zero value otherwise.

### GetMigrateSpeedOk

`func (o *UpdateVMConfigRequest) GetMigrateSpeedOk() (*int64, bool)`

GetMigrateSpeedOk returns a tuple with the MigrateSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigrateSpeed

`func (o *UpdateVMConfigRequest) SetMigrateSpeed(v int64)`

SetMigrateSpeed sets MigrateSpeed field to given value.

### HasMigrateSpeed

`func (o *UpdateVMConfigRequest) HasMigrateSpeed() bool`

HasMigrateSpeed returns a boolean if a field has been set.

### GetName

`func (o *UpdateVMConfigRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVMConfigRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVMConfigRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVMConfigRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNameserver

`func (o *UpdateVMConfigRequest) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *UpdateVMConfigRequest) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *UpdateVMConfigRequest) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *UpdateVMConfigRequest) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetNet0

`func (o *UpdateVMConfigRequest) GetNet0() string`

GetNet0 returns the Net0 field if non-nil, zero value otherwise.

### GetNet0Ok

`func (o *UpdateVMConfigRequest) GetNet0Ok() (*string, bool)`

GetNet0Ok returns a tuple with the Net0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet0

`func (o *UpdateVMConfigRequest) SetNet0(v string)`

SetNet0 sets Net0 field to given value.

### HasNet0

`func (o *UpdateVMConfigRequest) HasNet0() bool`

HasNet0 returns a boolean if a field has been set.

### GetNet1

`func (o *UpdateVMConfigRequest) GetNet1() string`

GetNet1 returns the Net1 field if non-nil, zero value otherwise.

### GetNet1Ok

`func (o *UpdateVMConfigRequest) GetNet1Ok() (*string, bool)`

GetNet1Ok returns a tuple with the Net1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet1

`func (o *UpdateVMConfigRequest) SetNet1(v string)`

SetNet1 sets Net1 field to given value.

### HasNet1

`func (o *UpdateVMConfigRequest) HasNet1() bool`

HasNet1 returns a boolean if a field has been set.

### GetNet2

`func (o *UpdateVMConfigRequest) GetNet2() string`

GetNet2 returns the Net2 field if non-nil, zero value otherwise.

### GetNet2Ok

`func (o *UpdateVMConfigRequest) GetNet2Ok() (*string, bool)`

GetNet2Ok returns a tuple with the Net2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet2

`func (o *UpdateVMConfigRequest) SetNet2(v string)`

SetNet2 sets Net2 field to given value.

### HasNet2

`func (o *UpdateVMConfigRequest) HasNet2() bool`

HasNet2 returns a boolean if a field has been set.

### GetNet3

`func (o *UpdateVMConfigRequest) GetNet3() string`

GetNet3 returns the Net3 field if non-nil, zero value otherwise.

### GetNet3Ok

`func (o *UpdateVMConfigRequest) GetNet3Ok() (*string, bool)`

GetNet3Ok returns a tuple with the Net3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet3

`func (o *UpdateVMConfigRequest) SetNet3(v string)`

SetNet3 sets Net3 field to given value.

### HasNet3

`func (o *UpdateVMConfigRequest) HasNet3() bool`

HasNet3 returns a boolean if a field has been set.

### GetNet4

`func (o *UpdateVMConfigRequest) GetNet4() string`

GetNet4 returns the Net4 field if non-nil, zero value otherwise.

### GetNet4Ok

`func (o *UpdateVMConfigRequest) GetNet4Ok() (*string, bool)`

GetNet4Ok returns a tuple with the Net4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet4

`func (o *UpdateVMConfigRequest) SetNet4(v string)`

SetNet4 sets Net4 field to given value.

### HasNet4

`func (o *UpdateVMConfigRequest) HasNet4() bool`

HasNet4 returns a boolean if a field has been set.

### GetNet5

`func (o *UpdateVMConfigRequest) GetNet5() string`

GetNet5 returns the Net5 field if non-nil, zero value otherwise.

### GetNet5Ok

`func (o *UpdateVMConfigRequest) GetNet5Ok() (*string, bool)`

GetNet5Ok returns a tuple with the Net5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet5

`func (o *UpdateVMConfigRequest) SetNet5(v string)`

SetNet5 sets Net5 field to given value.

### HasNet5

`func (o *UpdateVMConfigRequest) HasNet5() bool`

HasNet5 returns a boolean if a field has been set.

### GetNet6

`func (o *UpdateVMConfigRequest) GetNet6() string`

GetNet6 returns the Net6 field if non-nil, zero value otherwise.

### GetNet6Ok

`func (o *UpdateVMConfigRequest) GetNet6Ok() (*string, bool)`

GetNet6Ok returns a tuple with the Net6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet6

`func (o *UpdateVMConfigRequest) SetNet6(v string)`

SetNet6 sets Net6 field to given value.

### HasNet6

`func (o *UpdateVMConfigRequest) HasNet6() bool`

HasNet6 returns a boolean if a field has been set.

### GetNet7

`func (o *UpdateVMConfigRequest) GetNet7() string`

GetNet7 returns the Net7 field if non-nil, zero value otherwise.

### GetNet7Ok

`func (o *UpdateVMConfigRequest) GetNet7Ok() (*string, bool)`

GetNet7Ok returns a tuple with the Net7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet7

`func (o *UpdateVMConfigRequest) SetNet7(v string)`

SetNet7 sets Net7 field to given value.

### HasNet7

`func (o *UpdateVMConfigRequest) HasNet7() bool`

HasNet7 returns a boolean if a field has been set.

### GetNet8

`func (o *UpdateVMConfigRequest) GetNet8() string`

GetNet8 returns the Net8 field if non-nil, zero value otherwise.

### GetNet8Ok

`func (o *UpdateVMConfigRequest) GetNet8Ok() (*string, bool)`

GetNet8Ok returns a tuple with the Net8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet8

`func (o *UpdateVMConfigRequest) SetNet8(v string)`

SetNet8 sets Net8 field to given value.

### HasNet8

`func (o *UpdateVMConfigRequest) HasNet8() bool`

HasNet8 returns a boolean if a field has been set.

### GetNet9

`func (o *UpdateVMConfigRequest) GetNet9() string`

GetNet9 returns the Net9 field if non-nil, zero value otherwise.

### GetNet9Ok

`func (o *UpdateVMConfigRequest) GetNet9Ok() (*string, bool)`

GetNet9Ok returns a tuple with the Net9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet9

`func (o *UpdateVMConfigRequest) SetNet9(v string)`

SetNet9 sets Net9 field to given value.

### HasNet9

`func (o *UpdateVMConfigRequest) HasNet9() bool`

HasNet9 returns a boolean if a field has been set.

### GetNet10

`func (o *UpdateVMConfigRequest) GetNet10() string`

GetNet10 returns the Net10 field if non-nil, zero value otherwise.

### GetNet10Ok

`func (o *UpdateVMConfigRequest) GetNet10Ok() (*string, bool)`

GetNet10Ok returns a tuple with the Net10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet10

`func (o *UpdateVMConfigRequest) SetNet10(v string)`

SetNet10 sets Net10 field to given value.

### HasNet10

`func (o *UpdateVMConfigRequest) HasNet10() bool`

HasNet10 returns a boolean if a field has been set.

### GetNet11

`func (o *UpdateVMConfigRequest) GetNet11() string`

GetNet11 returns the Net11 field if non-nil, zero value otherwise.

### GetNet11Ok

`func (o *UpdateVMConfigRequest) GetNet11Ok() (*string, bool)`

GetNet11Ok returns a tuple with the Net11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet11

`func (o *UpdateVMConfigRequest) SetNet11(v string)`

SetNet11 sets Net11 field to given value.

### HasNet11

`func (o *UpdateVMConfigRequest) HasNet11() bool`

HasNet11 returns a boolean if a field has been set.

### GetNet12

`func (o *UpdateVMConfigRequest) GetNet12() string`

GetNet12 returns the Net12 field if non-nil, zero value otherwise.

### GetNet12Ok

`func (o *UpdateVMConfigRequest) GetNet12Ok() (*string, bool)`

GetNet12Ok returns a tuple with the Net12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet12

`func (o *UpdateVMConfigRequest) SetNet12(v string)`

SetNet12 sets Net12 field to given value.

### HasNet12

`func (o *UpdateVMConfigRequest) HasNet12() bool`

HasNet12 returns a boolean if a field has been set.

### GetNet13

`func (o *UpdateVMConfigRequest) GetNet13() string`

GetNet13 returns the Net13 field if non-nil, zero value otherwise.

### GetNet13Ok

`func (o *UpdateVMConfigRequest) GetNet13Ok() (*string, bool)`

GetNet13Ok returns a tuple with the Net13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet13

`func (o *UpdateVMConfigRequest) SetNet13(v string)`

SetNet13 sets Net13 field to given value.

### HasNet13

`func (o *UpdateVMConfigRequest) HasNet13() bool`

HasNet13 returns a boolean if a field has been set.

### GetNet14

`func (o *UpdateVMConfigRequest) GetNet14() string`

GetNet14 returns the Net14 field if non-nil, zero value otherwise.

### GetNet14Ok

`func (o *UpdateVMConfigRequest) GetNet14Ok() (*string, bool)`

GetNet14Ok returns a tuple with the Net14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet14

`func (o *UpdateVMConfigRequest) SetNet14(v string)`

SetNet14 sets Net14 field to given value.

### HasNet14

`func (o *UpdateVMConfigRequest) HasNet14() bool`

HasNet14 returns a boolean if a field has been set.

### GetNet15

`func (o *UpdateVMConfigRequest) GetNet15() string`

GetNet15 returns the Net15 field if non-nil, zero value otherwise.

### GetNet15Ok

`func (o *UpdateVMConfigRequest) GetNet15Ok() (*string, bool)`

GetNet15Ok returns a tuple with the Net15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet15

`func (o *UpdateVMConfigRequest) SetNet15(v string)`

SetNet15 sets Net15 field to given value.

### HasNet15

`func (o *UpdateVMConfigRequest) HasNet15() bool`

HasNet15 returns a boolean if a field has been set.

### GetNet16

`func (o *UpdateVMConfigRequest) GetNet16() string`

GetNet16 returns the Net16 field if non-nil, zero value otherwise.

### GetNet16Ok

`func (o *UpdateVMConfigRequest) GetNet16Ok() (*string, bool)`

GetNet16Ok returns a tuple with the Net16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet16

`func (o *UpdateVMConfigRequest) SetNet16(v string)`

SetNet16 sets Net16 field to given value.

### HasNet16

`func (o *UpdateVMConfigRequest) HasNet16() bool`

HasNet16 returns a boolean if a field has been set.

### GetNet17

`func (o *UpdateVMConfigRequest) GetNet17() string`

GetNet17 returns the Net17 field if non-nil, zero value otherwise.

### GetNet17Ok

`func (o *UpdateVMConfigRequest) GetNet17Ok() (*string, bool)`

GetNet17Ok returns a tuple with the Net17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet17

`func (o *UpdateVMConfigRequest) SetNet17(v string)`

SetNet17 sets Net17 field to given value.

### HasNet17

`func (o *UpdateVMConfigRequest) HasNet17() bool`

HasNet17 returns a boolean if a field has been set.

### GetNet18

`func (o *UpdateVMConfigRequest) GetNet18() string`

GetNet18 returns the Net18 field if non-nil, zero value otherwise.

### GetNet18Ok

`func (o *UpdateVMConfigRequest) GetNet18Ok() (*string, bool)`

GetNet18Ok returns a tuple with the Net18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet18

`func (o *UpdateVMConfigRequest) SetNet18(v string)`

SetNet18 sets Net18 field to given value.

### HasNet18

`func (o *UpdateVMConfigRequest) HasNet18() bool`

HasNet18 returns a boolean if a field has been set.

### GetNet19

`func (o *UpdateVMConfigRequest) GetNet19() string`

GetNet19 returns the Net19 field if non-nil, zero value otherwise.

### GetNet19Ok

`func (o *UpdateVMConfigRequest) GetNet19Ok() (*string, bool)`

GetNet19Ok returns a tuple with the Net19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet19

`func (o *UpdateVMConfigRequest) SetNet19(v string)`

SetNet19 sets Net19 field to given value.

### HasNet19

`func (o *UpdateVMConfigRequest) HasNet19() bool`

HasNet19 returns a boolean if a field has been set.

### GetNet20

`func (o *UpdateVMConfigRequest) GetNet20() string`

GetNet20 returns the Net20 field if non-nil, zero value otherwise.

### GetNet20Ok

`func (o *UpdateVMConfigRequest) GetNet20Ok() (*string, bool)`

GetNet20Ok returns a tuple with the Net20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet20

`func (o *UpdateVMConfigRequest) SetNet20(v string)`

SetNet20 sets Net20 field to given value.

### HasNet20

`func (o *UpdateVMConfigRequest) HasNet20() bool`

HasNet20 returns a boolean if a field has been set.

### GetNet21

`func (o *UpdateVMConfigRequest) GetNet21() string`

GetNet21 returns the Net21 field if non-nil, zero value otherwise.

### GetNet21Ok

`func (o *UpdateVMConfigRequest) GetNet21Ok() (*string, bool)`

GetNet21Ok returns a tuple with the Net21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet21

`func (o *UpdateVMConfigRequest) SetNet21(v string)`

SetNet21 sets Net21 field to given value.

### HasNet21

`func (o *UpdateVMConfigRequest) HasNet21() bool`

HasNet21 returns a boolean if a field has been set.

### GetNet22

`func (o *UpdateVMConfigRequest) GetNet22() string`

GetNet22 returns the Net22 field if non-nil, zero value otherwise.

### GetNet22Ok

`func (o *UpdateVMConfigRequest) GetNet22Ok() (*string, bool)`

GetNet22Ok returns a tuple with the Net22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet22

`func (o *UpdateVMConfigRequest) SetNet22(v string)`

SetNet22 sets Net22 field to given value.

### HasNet22

`func (o *UpdateVMConfigRequest) HasNet22() bool`

HasNet22 returns a boolean if a field has been set.

### GetNet23

`func (o *UpdateVMConfigRequest) GetNet23() string`

GetNet23 returns the Net23 field if non-nil, zero value otherwise.

### GetNet23Ok

`func (o *UpdateVMConfigRequest) GetNet23Ok() (*string, bool)`

GetNet23Ok returns a tuple with the Net23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet23

`func (o *UpdateVMConfigRequest) SetNet23(v string)`

SetNet23 sets Net23 field to given value.

### HasNet23

`func (o *UpdateVMConfigRequest) HasNet23() bool`

HasNet23 returns a boolean if a field has been set.

### GetNet24

`func (o *UpdateVMConfigRequest) GetNet24() string`

GetNet24 returns the Net24 field if non-nil, zero value otherwise.

### GetNet24Ok

`func (o *UpdateVMConfigRequest) GetNet24Ok() (*string, bool)`

GetNet24Ok returns a tuple with the Net24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet24

`func (o *UpdateVMConfigRequest) SetNet24(v string)`

SetNet24 sets Net24 field to given value.

### HasNet24

`func (o *UpdateVMConfigRequest) HasNet24() bool`

HasNet24 returns a boolean if a field has been set.

### GetNet25

`func (o *UpdateVMConfigRequest) GetNet25() string`

GetNet25 returns the Net25 field if non-nil, zero value otherwise.

### GetNet25Ok

`func (o *UpdateVMConfigRequest) GetNet25Ok() (*string, bool)`

GetNet25Ok returns a tuple with the Net25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet25

`func (o *UpdateVMConfigRequest) SetNet25(v string)`

SetNet25 sets Net25 field to given value.

### HasNet25

`func (o *UpdateVMConfigRequest) HasNet25() bool`

HasNet25 returns a boolean if a field has been set.

### GetNet26

`func (o *UpdateVMConfigRequest) GetNet26() string`

GetNet26 returns the Net26 field if non-nil, zero value otherwise.

### GetNet26Ok

`func (o *UpdateVMConfigRequest) GetNet26Ok() (*string, bool)`

GetNet26Ok returns a tuple with the Net26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet26

`func (o *UpdateVMConfigRequest) SetNet26(v string)`

SetNet26 sets Net26 field to given value.

### HasNet26

`func (o *UpdateVMConfigRequest) HasNet26() bool`

HasNet26 returns a boolean if a field has been set.

### GetNet27

`func (o *UpdateVMConfigRequest) GetNet27() string`

GetNet27 returns the Net27 field if non-nil, zero value otherwise.

### GetNet27Ok

`func (o *UpdateVMConfigRequest) GetNet27Ok() (*string, bool)`

GetNet27Ok returns a tuple with the Net27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet27

`func (o *UpdateVMConfigRequest) SetNet27(v string)`

SetNet27 sets Net27 field to given value.

### HasNet27

`func (o *UpdateVMConfigRequest) HasNet27() bool`

HasNet27 returns a boolean if a field has been set.

### GetNet28

`func (o *UpdateVMConfigRequest) GetNet28() string`

GetNet28 returns the Net28 field if non-nil, zero value otherwise.

### GetNet28Ok

`func (o *UpdateVMConfigRequest) GetNet28Ok() (*string, bool)`

GetNet28Ok returns a tuple with the Net28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet28

`func (o *UpdateVMConfigRequest) SetNet28(v string)`

SetNet28 sets Net28 field to given value.

### HasNet28

`func (o *UpdateVMConfigRequest) HasNet28() bool`

HasNet28 returns a boolean if a field has been set.

### GetNet29

`func (o *UpdateVMConfigRequest) GetNet29() string`

GetNet29 returns the Net29 field if non-nil, zero value otherwise.

### GetNet29Ok

`func (o *UpdateVMConfigRequest) GetNet29Ok() (*string, bool)`

GetNet29Ok returns a tuple with the Net29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet29

`func (o *UpdateVMConfigRequest) SetNet29(v string)`

SetNet29 sets Net29 field to given value.

### HasNet29

`func (o *UpdateVMConfigRequest) HasNet29() bool`

HasNet29 returns a boolean if a field has been set.

### GetNet30

`func (o *UpdateVMConfigRequest) GetNet30() string`

GetNet30 returns the Net30 field if non-nil, zero value otherwise.

### GetNet30Ok

`func (o *UpdateVMConfigRequest) GetNet30Ok() (*string, bool)`

GetNet30Ok returns a tuple with the Net30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet30

`func (o *UpdateVMConfigRequest) SetNet30(v string)`

SetNet30 sets Net30 field to given value.

### HasNet30

`func (o *UpdateVMConfigRequest) HasNet30() bool`

HasNet30 returns a boolean if a field has been set.

### GetNet31

`func (o *UpdateVMConfigRequest) GetNet31() string`

GetNet31 returns the Net31 field if non-nil, zero value otherwise.

### GetNet31Ok

`func (o *UpdateVMConfigRequest) GetNet31Ok() (*string, bool)`

GetNet31Ok returns a tuple with the Net31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet31

`func (o *UpdateVMConfigRequest) SetNet31(v string)`

SetNet31 sets Net31 field to given value.

### HasNet31

`func (o *UpdateVMConfigRequest) HasNet31() bool`

HasNet31 returns a boolean if a field has been set.

### GetNuma

`func (o *UpdateVMConfigRequest) GetNuma() bool`

GetNuma returns the Numa field if non-nil, zero value otherwise.

### GetNumaOk

`func (o *UpdateVMConfigRequest) GetNumaOk() (*bool, bool)`

GetNumaOk returns a tuple with the Numa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma

`func (o *UpdateVMConfigRequest) SetNuma(v bool)`

SetNuma sets Numa field to given value.

### HasNuma

`func (o *UpdateVMConfigRequest) HasNuma() bool`

HasNuma returns a boolean if a field has been set.

### GetNuma0

`func (o *UpdateVMConfigRequest) GetNuma0() string`

GetNuma0 returns the Numa0 field if non-nil, zero value otherwise.

### GetNuma0Ok

`func (o *UpdateVMConfigRequest) GetNuma0Ok() (*string, bool)`

GetNuma0Ok returns a tuple with the Numa0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma0

`func (o *UpdateVMConfigRequest) SetNuma0(v string)`

SetNuma0 sets Numa0 field to given value.

### HasNuma0

`func (o *UpdateVMConfigRequest) HasNuma0() bool`

HasNuma0 returns a boolean if a field has been set.

### GetNuma1

`func (o *UpdateVMConfigRequest) GetNuma1() string`

GetNuma1 returns the Numa1 field if non-nil, zero value otherwise.

### GetNuma1Ok

`func (o *UpdateVMConfigRequest) GetNuma1Ok() (*string, bool)`

GetNuma1Ok returns a tuple with the Numa1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma1

`func (o *UpdateVMConfigRequest) SetNuma1(v string)`

SetNuma1 sets Numa1 field to given value.

### HasNuma1

`func (o *UpdateVMConfigRequest) HasNuma1() bool`

HasNuma1 returns a boolean if a field has been set.

### GetNuma2

`func (o *UpdateVMConfigRequest) GetNuma2() string`

GetNuma2 returns the Numa2 field if non-nil, zero value otherwise.

### GetNuma2Ok

`func (o *UpdateVMConfigRequest) GetNuma2Ok() (*string, bool)`

GetNuma2Ok returns a tuple with the Numa2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma2

`func (o *UpdateVMConfigRequest) SetNuma2(v string)`

SetNuma2 sets Numa2 field to given value.

### HasNuma2

`func (o *UpdateVMConfigRequest) HasNuma2() bool`

HasNuma2 returns a boolean if a field has been set.

### GetNuma3

`func (o *UpdateVMConfigRequest) GetNuma3() string`

GetNuma3 returns the Numa3 field if non-nil, zero value otherwise.

### GetNuma3Ok

`func (o *UpdateVMConfigRequest) GetNuma3Ok() (*string, bool)`

GetNuma3Ok returns a tuple with the Numa3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma3

`func (o *UpdateVMConfigRequest) SetNuma3(v string)`

SetNuma3 sets Numa3 field to given value.

### HasNuma3

`func (o *UpdateVMConfigRequest) HasNuma3() bool`

HasNuma3 returns a boolean if a field has been set.

### GetNuma4

`func (o *UpdateVMConfigRequest) GetNuma4() string`

GetNuma4 returns the Numa4 field if non-nil, zero value otherwise.

### GetNuma4Ok

`func (o *UpdateVMConfigRequest) GetNuma4Ok() (*string, bool)`

GetNuma4Ok returns a tuple with the Numa4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma4

`func (o *UpdateVMConfigRequest) SetNuma4(v string)`

SetNuma4 sets Numa4 field to given value.

### HasNuma4

`func (o *UpdateVMConfigRequest) HasNuma4() bool`

HasNuma4 returns a boolean if a field has been set.

### GetNuma5

`func (o *UpdateVMConfigRequest) GetNuma5() string`

GetNuma5 returns the Numa5 field if non-nil, zero value otherwise.

### GetNuma5Ok

`func (o *UpdateVMConfigRequest) GetNuma5Ok() (*string, bool)`

GetNuma5Ok returns a tuple with the Numa5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma5

`func (o *UpdateVMConfigRequest) SetNuma5(v string)`

SetNuma5 sets Numa5 field to given value.

### HasNuma5

`func (o *UpdateVMConfigRequest) HasNuma5() bool`

HasNuma5 returns a boolean if a field has been set.

### GetNuma6

`func (o *UpdateVMConfigRequest) GetNuma6() string`

GetNuma6 returns the Numa6 field if non-nil, zero value otherwise.

### GetNuma6Ok

`func (o *UpdateVMConfigRequest) GetNuma6Ok() (*string, bool)`

GetNuma6Ok returns a tuple with the Numa6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma6

`func (o *UpdateVMConfigRequest) SetNuma6(v string)`

SetNuma6 sets Numa6 field to given value.

### HasNuma6

`func (o *UpdateVMConfigRequest) HasNuma6() bool`

HasNuma6 returns a boolean if a field has been set.

### GetNuma7

`func (o *UpdateVMConfigRequest) GetNuma7() string`

GetNuma7 returns the Numa7 field if non-nil, zero value otherwise.

### GetNuma7Ok

`func (o *UpdateVMConfigRequest) GetNuma7Ok() (*string, bool)`

GetNuma7Ok returns a tuple with the Numa7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma7

`func (o *UpdateVMConfigRequest) SetNuma7(v string)`

SetNuma7 sets Numa7 field to given value.

### HasNuma7

`func (o *UpdateVMConfigRequest) HasNuma7() bool`

HasNuma7 returns a boolean if a field has been set.

### GetNuma8

`func (o *UpdateVMConfigRequest) GetNuma8() string`

GetNuma8 returns the Numa8 field if non-nil, zero value otherwise.

### GetNuma8Ok

`func (o *UpdateVMConfigRequest) GetNuma8Ok() (*string, bool)`

GetNuma8Ok returns a tuple with the Numa8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma8

`func (o *UpdateVMConfigRequest) SetNuma8(v string)`

SetNuma8 sets Numa8 field to given value.

### HasNuma8

`func (o *UpdateVMConfigRequest) HasNuma8() bool`

HasNuma8 returns a boolean if a field has been set.

### GetNuma9

`func (o *UpdateVMConfigRequest) GetNuma9() string`

GetNuma9 returns the Numa9 field if non-nil, zero value otherwise.

### GetNuma9Ok

`func (o *UpdateVMConfigRequest) GetNuma9Ok() (*string, bool)`

GetNuma9Ok returns a tuple with the Numa9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma9

`func (o *UpdateVMConfigRequest) SetNuma9(v string)`

SetNuma9 sets Numa9 field to given value.

### HasNuma9

`func (o *UpdateVMConfigRequest) HasNuma9() bool`

HasNuma9 returns a boolean if a field has been set.

### GetNuma10

`func (o *UpdateVMConfigRequest) GetNuma10() string`

GetNuma10 returns the Numa10 field if non-nil, zero value otherwise.

### GetNuma10Ok

`func (o *UpdateVMConfigRequest) GetNuma10Ok() (*string, bool)`

GetNuma10Ok returns a tuple with the Numa10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma10

`func (o *UpdateVMConfigRequest) SetNuma10(v string)`

SetNuma10 sets Numa10 field to given value.

### HasNuma10

`func (o *UpdateVMConfigRequest) HasNuma10() bool`

HasNuma10 returns a boolean if a field has been set.

### GetNuma11

`func (o *UpdateVMConfigRequest) GetNuma11() string`

GetNuma11 returns the Numa11 field if non-nil, zero value otherwise.

### GetNuma11Ok

`func (o *UpdateVMConfigRequest) GetNuma11Ok() (*string, bool)`

GetNuma11Ok returns a tuple with the Numa11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma11

`func (o *UpdateVMConfigRequest) SetNuma11(v string)`

SetNuma11 sets Numa11 field to given value.

### HasNuma11

`func (o *UpdateVMConfigRequest) HasNuma11() bool`

HasNuma11 returns a boolean if a field has been set.

### GetNuma12

`func (o *UpdateVMConfigRequest) GetNuma12() string`

GetNuma12 returns the Numa12 field if non-nil, zero value otherwise.

### GetNuma12Ok

`func (o *UpdateVMConfigRequest) GetNuma12Ok() (*string, bool)`

GetNuma12Ok returns a tuple with the Numa12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma12

`func (o *UpdateVMConfigRequest) SetNuma12(v string)`

SetNuma12 sets Numa12 field to given value.

### HasNuma12

`func (o *UpdateVMConfigRequest) HasNuma12() bool`

HasNuma12 returns a boolean if a field has been set.

### GetNuma13

`func (o *UpdateVMConfigRequest) GetNuma13() string`

GetNuma13 returns the Numa13 field if non-nil, zero value otherwise.

### GetNuma13Ok

`func (o *UpdateVMConfigRequest) GetNuma13Ok() (*string, bool)`

GetNuma13Ok returns a tuple with the Numa13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma13

`func (o *UpdateVMConfigRequest) SetNuma13(v string)`

SetNuma13 sets Numa13 field to given value.

### HasNuma13

`func (o *UpdateVMConfigRequest) HasNuma13() bool`

HasNuma13 returns a boolean if a field has been set.

### GetNuma14

`func (o *UpdateVMConfigRequest) GetNuma14() string`

GetNuma14 returns the Numa14 field if non-nil, zero value otherwise.

### GetNuma14Ok

`func (o *UpdateVMConfigRequest) GetNuma14Ok() (*string, bool)`

GetNuma14Ok returns a tuple with the Numa14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma14

`func (o *UpdateVMConfigRequest) SetNuma14(v string)`

SetNuma14 sets Numa14 field to given value.

### HasNuma14

`func (o *UpdateVMConfigRequest) HasNuma14() bool`

HasNuma14 returns a boolean if a field has been set.

### GetNuma15

`func (o *UpdateVMConfigRequest) GetNuma15() string`

GetNuma15 returns the Numa15 field if non-nil, zero value otherwise.

### GetNuma15Ok

`func (o *UpdateVMConfigRequest) GetNuma15Ok() (*string, bool)`

GetNuma15Ok returns a tuple with the Numa15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma15

`func (o *UpdateVMConfigRequest) SetNuma15(v string)`

SetNuma15 sets Numa15 field to given value.

### HasNuma15

`func (o *UpdateVMConfigRequest) HasNuma15() bool`

HasNuma15 returns a boolean if a field has been set.

### GetNuma16

`func (o *UpdateVMConfigRequest) GetNuma16() string`

GetNuma16 returns the Numa16 field if non-nil, zero value otherwise.

### GetNuma16Ok

`func (o *UpdateVMConfigRequest) GetNuma16Ok() (*string, bool)`

GetNuma16Ok returns a tuple with the Numa16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma16

`func (o *UpdateVMConfigRequest) SetNuma16(v string)`

SetNuma16 sets Numa16 field to given value.

### HasNuma16

`func (o *UpdateVMConfigRequest) HasNuma16() bool`

HasNuma16 returns a boolean if a field has been set.

### GetNuma17

`func (o *UpdateVMConfigRequest) GetNuma17() string`

GetNuma17 returns the Numa17 field if non-nil, zero value otherwise.

### GetNuma17Ok

`func (o *UpdateVMConfigRequest) GetNuma17Ok() (*string, bool)`

GetNuma17Ok returns a tuple with the Numa17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma17

`func (o *UpdateVMConfigRequest) SetNuma17(v string)`

SetNuma17 sets Numa17 field to given value.

### HasNuma17

`func (o *UpdateVMConfigRequest) HasNuma17() bool`

HasNuma17 returns a boolean if a field has been set.

### GetNuma18

`func (o *UpdateVMConfigRequest) GetNuma18() string`

GetNuma18 returns the Numa18 field if non-nil, zero value otherwise.

### GetNuma18Ok

`func (o *UpdateVMConfigRequest) GetNuma18Ok() (*string, bool)`

GetNuma18Ok returns a tuple with the Numa18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma18

`func (o *UpdateVMConfigRequest) SetNuma18(v string)`

SetNuma18 sets Numa18 field to given value.

### HasNuma18

`func (o *UpdateVMConfigRequest) HasNuma18() bool`

HasNuma18 returns a boolean if a field has been set.

### GetNuma19

`func (o *UpdateVMConfigRequest) GetNuma19() string`

GetNuma19 returns the Numa19 field if non-nil, zero value otherwise.

### GetNuma19Ok

`func (o *UpdateVMConfigRequest) GetNuma19Ok() (*string, bool)`

GetNuma19Ok returns a tuple with the Numa19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma19

`func (o *UpdateVMConfigRequest) SetNuma19(v string)`

SetNuma19 sets Numa19 field to given value.

### HasNuma19

`func (o *UpdateVMConfigRequest) HasNuma19() bool`

HasNuma19 returns a boolean if a field has been set.

### GetNuma20

`func (o *UpdateVMConfigRequest) GetNuma20() string`

GetNuma20 returns the Numa20 field if non-nil, zero value otherwise.

### GetNuma20Ok

`func (o *UpdateVMConfigRequest) GetNuma20Ok() (*string, bool)`

GetNuma20Ok returns a tuple with the Numa20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma20

`func (o *UpdateVMConfigRequest) SetNuma20(v string)`

SetNuma20 sets Numa20 field to given value.

### HasNuma20

`func (o *UpdateVMConfigRequest) HasNuma20() bool`

HasNuma20 returns a boolean if a field has been set.

### GetNuma21

`func (o *UpdateVMConfigRequest) GetNuma21() string`

GetNuma21 returns the Numa21 field if non-nil, zero value otherwise.

### GetNuma21Ok

`func (o *UpdateVMConfigRequest) GetNuma21Ok() (*string, bool)`

GetNuma21Ok returns a tuple with the Numa21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma21

`func (o *UpdateVMConfigRequest) SetNuma21(v string)`

SetNuma21 sets Numa21 field to given value.

### HasNuma21

`func (o *UpdateVMConfigRequest) HasNuma21() bool`

HasNuma21 returns a boolean if a field has been set.

### GetNuma22

`func (o *UpdateVMConfigRequest) GetNuma22() string`

GetNuma22 returns the Numa22 field if non-nil, zero value otherwise.

### GetNuma22Ok

`func (o *UpdateVMConfigRequest) GetNuma22Ok() (*string, bool)`

GetNuma22Ok returns a tuple with the Numa22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma22

`func (o *UpdateVMConfigRequest) SetNuma22(v string)`

SetNuma22 sets Numa22 field to given value.

### HasNuma22

`func (o *UpdateVMConfigRequest) HasNuma22() bool`

HasNuma22 returns a boolean if a field has been set.

### GetNuma23

`func (o *UpdateVMConfigRequest) GetNuma23() string`

GetNuma23 returns the Numa23 field if non-nil, zero value otherwise.

### GetNuma23Ok

`func (o *UpdateVMConfigRequest) GetNuma23Ok() (*string, bool)`

GetNuma23Ok returns a tuple with the Numa23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma23

`func (o *UpdateVMConfigRequest) SetNuma23(v string)`

SetNuma23 sets Numa23 field to given value.

### HasNuma23

`func (o *UpdateVMConfigRequest) HasNuma23() bool`

HasNuma23 returns a boolean if a field has been set.

### GetNuma24

`func (o *UpdateVMConfigRequest) GetNuma24() string`

GetNuma24 returns the Numa24 field if non-nil, zero value otherwise.

### GetNuma24Ok

`func (o *UpdateVMConfigRequest) GetNuma24Ok() (*string, bool)`

GetNuma24Ok returns a tuple with the Numa24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma24

`func (o *UpdateVMConfigRequest) SetNuma24(v string)`

SetNuma24 sets Numa24 field to given value.

### HasNuma24

`func (o *UpdateVMConfigRequest) HasNuma24() bool`

HasNuma24 returns a boolean if a field has been set.

### GetNuma25

`func (o *UpdateVMConfigRequest) GetNuma25() string`

GetNuma25 returns the Numa25 field if non-nil, zero value otherwise.

### GetNuma25Ok

`func (o *UpdateVMConfigRequest) GetNuma25Ok() (*string, bool)`

GetNuma25Ok returns a tuple with the Numa25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma25

`func (o *UpdateVMConfigRequest) SetNuma25(v string)`

SetNuma25 sets Numa25 field to given value.

### HasNuma25

`func (o *UpdateVMConfigRequest) HasNuma25() bool`

HasNuma25 returns a boolean if a field has been set.

### GetNuma26

`func (o *UpdateVMConfigRequest) GetNuma26() string`

GetNuma26 returns the Numa26 field if non-nil, zero value otherwise.

### GetNuma26Ok

`func (o *UpdateVMConfigRequest) GetNuma26Ok() (*string, bool)`

GetNuma26Ok returns a tuple with the Numa26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma26

`func (o *UpdateVMConfigRequest) SetNuma26(v string)`

SetNuma26 sets Numa26 field to given value.

### HasNuma26

`func (o *UpdateVMConfigRequest) HasNuma26() bool`

HasNuma26 returns a boolean if a field has been set.

### GetNuma27

`func (o *UpdateVMConfigRequest) GetNuma27() string`

GetNuma27 returns the Numa27 field if non-nil, zero value otherwise.

### GetNuma27Ok

`func (o *UpdateVMConfigRequest) GetNuma27Ok() (*string, bool)`

GetNuma27Ok returns a tuple with the Numa27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma27

`func (o *UpdateVMConfigRequest) SetNuma27(v string)`

SetNuma27 sets Numa27 field to given value.

### HasNuma27

`func (o *UpdateVMConfigRequest) HasNuma27() bool`

HasNuma27 returns a boolean if a field has been set.

### GetNuma28

`func (o *UpdateVMConfigRequest) GetNuma28() string`

GetNuma28 returns the Numa28 field if non-nil, zero value otherwise.

### GetNuma28Ok

`func (o *UpdateVMConfigRequest) GetNuma28Ok() (*string, bool)`

GetNuma28Ok returns a tuple with the Numa28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma28

`func (o *UpdateVMConfigRequest) SetNuma28(v string)`

SetNuma28 sets Numa28 field to given value.

### HasNuma28

`func (o *UpdateVMConfigRequest) HasNuma28() bool`

HasNuma28 returns a boolean if a field has been set.

### GetNuma29

`func (o *UpdateVMConfigRequest) GetNuma29() string`

GetNuma29 returns the Numa29 field if non-nil, zero value otherwise.

### GetNuma29Ok

`func (o *UpdateVMConfigRequest) GetNuma29Ok() (*string, bool)`

GetNuma29Ok returns a tuple with the Numa29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNuma29

`func (o *UpdateVMConfigRequest) SetNuma29(v string)`

SetNuma29 sets Numa29 field to given value.

### HasNuma29

`func (o *UpdateVMConfigRequest) HasNuma29() bool`

HasNuma29 returns a boolean if a field has been set.

### GetOnboot

`func (o *UpdateVMConfigRequest) GetOnboot() bool`

GetOnboot returns the Onboot field if non-nil, zero value otherwise.

### GetOnbootOk

`func (o *UpdateVMConfigRequest) GetOnbootOk() (*bool, bool)`

GetOnbootOk returns a tuple with the Onboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboot

`func (o *UpdateVMConfigRequest) SetOnboot(v bool)`

SetOnboot sets Onboot field to given value.

### HasOnboot

`func (o *UpdateVMConfigRequest) HasOnboot() bool`

HasOnboot returns a boolean if a field has been set.

### GetOstype

`func (o *UpdateVMConfigRequest) GetOstype() string`

GetOstype returns the Ostype field if non-nil, zero value otherwise.

### GetOstypeOk

`func (o *UpdateVMConfigRequest) GetOstypeOk() (*string, bool)`

GetOstypeOk returns a tuple with the Ostype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstype

`func (o *UpdateVMConfigRequest) SetOstype(v string)`

SetOstype sets Ostype field to given value.

### HasOstype

`func (o *UpdateVMConfigRequest) HasOstype() bool`

HasOstype returns a boolean if a field has been set.

### GetParallel0

`func (o *UpdateVMConfigRequest) GetParallel0() string`

GetParallel0 returns the Parallel0 field if non-nil, zero value otherwise.

### GetParallel0Ok

`func (o *UpdateVMConfigRequest) GetParallel0Ok() (*string, bool)`

GetParallel0Ok returns a tuple with the Parallel0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel0

`func (o *UpdateVMConfigRequest) SetParallel0(v string)`

SetParallel0 sets Parallel0 field to given value.

### HasParallel0

`func (o *UpdateVMConfigRequest) HasParallel0() bool`

HasParallel0 returns a boolean if a field has been set.

### GetParallel1

`func (o *UpdateVMConfigRequest) GetParallel1() string`

GetParallel1 returns the Parallel1 field if non-nil, zero value otherwise.

### GetParallel1Ok

`func (o *UpdateVMConfigRequest) GetParallel1Ok() (*string, bool)`

GetParallel1Ok returns a tuple with the Parallel1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel1

`func (o *UpdateVMConfigRequest) SetParallel1(v string)`

SetParallel1 sets Parallel1 field to given value.

### HasParallel1

`func (o *UpdateVMConfigRequest) HasParallel1() bool`

HasParallel1 returns a boolean if a field has been set.

### GetParallel2

`func (o *UpdateVMConfigRequest) GetParallel2() string`

GetParallel2 returns the Parallel2 field if non-nil, zero value otherwise.

### GetParallel2Ok

`func (o *UpdateVMConfigRequest) GetParallel2Ok() (*string, bool)`

GetParallel2Ok returns a tuple with the Parallel2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel2

`func (o *UpdateVMConfigRequest) SetParallel2(v string)`

SetParallel2 sets Parallel2 field to given value.

### HasParallel2

`func (o *UpdateVMConfigRequest) HasParallel2() bool`

HasParallel2 returns a boolean if a field has been set.

### GetParallel3

`func (o *UpdateVMConfigRequest) GetParallel3() string`

GetParallel3 returns the Parallel3 field if non-nil, zero value otherwise.

### GetParallel3Ok

`func (o *UpdateVMConfigRequest) GetParallel3Ok() (*string, bool)`

GetParallel3Ok returns a tuple with the Parallel3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallel3

`func (o *UpdateVMConfigRequest) SetParallel3(v string)`

SetParallel3 sets Parallel3 field to given value.

### HasParallel3

`func (o *UpdateVMConfigRequest) HasParallel3() bool`

HasParallel3 returns a boolean if a field has been set.

### GetProtection

`func (o *UpdateVMConfigRequest) GetProtection() bool`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *UpdateVMConfigRequest) GetProtectionOk() (*bool, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *UpdateVMConfigRequest) SetProtection(v bool)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *UpdateVMConfigRequest) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetReboot

`func (o *UpdateVMConfigRequest) GetReboot() bool`

GetReboot returns the Reboot field if non-nil, zero value otherwise.

### GetRebootOk

`func (o *UpdateVMConfigRequest) GetRebootOk() (*bool, bool)`

GetRebootOk returns a tuple with the Reboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReboot

`func (o *UpdateVMConfigRequest) SetReboot(v bool)`

SetReboot sets Reboot field to given value.

### HasReboot

`func (o *UpdateVMConfigRequest) HasReboot() bool`

HasReboot returns a boolean if a field has been set.

### GetRevert

`func (o *UpdateVMConfigRequest) GetRevert() string`

GetRevert returns the Revert field if non-nil, zero value otherwise.

### GetRevertOk

`func (o *UpdateVMConfigRequest) GetRevertOk() (*string, bool)`

GetRevertOk returns a tuple with the Revert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevert

`func (o *UpdateVMConfigRequest) SetRevert(v string)`

SetRevert sets Revert field to given value.

### HasRevert

`func (o *UpdateVMConfigRequest) HasRevert() bool`

HasRevert returns a boolean if a field has been set.

### GetRng0

`func (o *UpdateVMConfigRequest) GetRng0() string`

GetRng0 returns the Rng0 field if non-nil, zero value otherwise.

### GetRng0Ok

`func (o *UpdateVMConfigRequest) GetRng0Ok() (*string, bool)`

GetRng0Ok returns a tuple with the Rng0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRng0

`func (o *UpdateVMConfigRequest) SetRng0(v string)`

SetRng0 sets Rng0 field to given value.

### HasRng0

`func (o *UpdateVMConfigRequest) HasRng0() bool`

HasRng0 returns a boolean if a field has been set.

### GetSata0

`func (o *UpdateVMConfigRequest) GetSata0() string`

GetSata0 returns the Sata0 field if non-nil, zero value otherwise.

### GetSata0Ok

`func (o *UpdateVMConfigRequest) GetSata0Ok() (*string, bool)`

GetSata0Ok returns a tuple with the Sata0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata0

`func (o *UpdateVMConfigRequest) SetSata0(v string)`

SetSata0 sets Sata0 field to given value.

### HasSata0

`func (o *UpdateVMConfigRequest) HasSata0() bool`

HasSata0 returns a boolean if a field has been set.

### GetSata1

`func (o *UpdateVMConfigRequest) GetSata1() string`

GetSata1 returns the Sata1 field if non-nil, zero value otherwise.

### GetSata1Ok

`func (o *UpdateVMConfigRequest) GetSata1Ok() (*string, bool)`

GetSata1Ok returns a tuple with the Sata1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata1

`func (o *UpdateVMConfigRequest) SetSata1(v string)`

SetSata1 sets Sata1 field to given value.

### HasSata1

`func (o *UpdateVMConfigRequest) HasSata1() bool`

HasSata1 returns a boolean if a field has been set.

### GetSata2

`func (o *UpdateVMConfigRequest) GetSata2() string`

GetSata2 returns the Sata2 field if non-nil, zero value otherwise.

### GetSata2Ok

`func (o *UpdateVMConfigRequest) GetSata2Ok() (*string, bool)`

GetSata2Ok returns a tuple with the Sata2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata2

`func (o *UpdateVMConfigRequest) SetSata2(v string)`

SetSata2 sets Sata2 field to given value.

### HasSata2

`func (o *UpdateVMConfigRequest) HasSata2() bool`

HasSata2 returns a boolean if a field has been set.

### GetSata3

`func (o *UpdateVMConfigRequest) GetSata3() string`

GetSata3 returns the Sata3 field if non-nil, zero value otherwise.

### GetSata3Ok

`func (o *UpdateVMConfigRequest) GetSata3Ok() (*string, bool)`

GetSata3Ok returns a tuple with the Sata3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata3

`func (o *UpdateVMConfigRequest) SetSata3(v string)`

SetSata3 sets Sata3 field to given value.

### HasSata3

`func (o *UpdateVMConfigRequest) HasSata3() bool`

HasSata3 returns a boolean if a field has been set.

### GetSata4

`func (o *UpdateVMConfigRequest) GetSata4() string`

GetSata4 returns the Sata4 field if non-nil, zero value otherwise.

### GetSata4Ok

`func (o *UpdateVMConfigRequest) GetSata4Ok() (*string, bool)`

GetSata4Ok returns a tuple with the Sata4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata4

`func (o *UpdateVMConfigRequest) SetSata4(v string)`

SetSata4 sets Sata4 field to given value.

### HasSata4

`func (o *UpdateVMConfigRequest) HasSata4() bool`

HasSata4 returns a boolean if a field has been set.

### GetSata5

`func (o *UpdateVMConfigRequest) GetSata5() string`

GetSata5 returns the Sata5 field if non-nil, zero value otherwise.

### GetSata5Ok

`func (o *UpdateVMConfigRequest) GetSata5Ok() (*string, bool)`

GetSata5Ok returns a tuple with the Sata5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSata5

`func (o *UpdateVMConfigRequest) SetSata5(v string)`

SetSata5 sets Sata5 field to given value.

### HasSata5

`func (o *UpdateVMConfigRequest) HasSata5() bool`

HasSata5 returns a boolean if a field has been set.

### GetScsi0

`func (o *UpdateVMConfigRequest) GetScsi0() string`

GetScsi0 returns the Scsi0 field if non-nil, zero value otherwise.

### GetScsi0Ok

`func (o *UpdateVMConfigRequest) GetScsi0Ok() (*string, bool)`

GetScsi0Ok returns a tuple with the Scsi0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi0

`func (o *UpdateVMConfigRequest) SetScsi0(v string)`

SetScsi0 sets Scsi0 field to given value.

### HasScsi0

`func (o *UpdateVMConfigRequest) HasScsi0() bool`

HasScsi0 returns a boolean if a field has been set.

### GetScsi1

`func (o *UpdateVMConfigRequest) GetScsi1() string`

GetScsi1 returns the Scsi1 field if non-nil, zero value otherwise.

### GetScsi1Ok

`func (o *UpdateVMConfigRequest) GetScsi1Ok() (*string, bool)`

GetScsi1Ok returns a tuple with the Scsi1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi1

`func (o *UpdateVMConfigRequest) SetScsi1(v string)`

SetScsi1 sets Scsi1 field to given value.

### HasScsi1

`func (o *UpdateVMConfigRequest) HasScsi1() bool`

HasScsi1 returns a boolean if a field has been set.

### GetScsi2

`func (o *UpdateVMConfigRequest) GetScsi2() string`

GetScsi2 returns the Scsi2 field if non-nil, zero value otherwise.

### GetScsi2Ok

`func (o *UpdateVMConfigRequest) GetScsi2Ok() (*string, bool)`

GetScsi2Ok returns a tuple with the Scsi2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi2

`func (o *UpdateVMConfigRequest) SetScsi2(v string)`

SetScsi2 sets Scsi2 field to given value.

### HasScsi2

`func (o *UpdateVMConfigRequest) HasScsi2() bool`

HasScsi2 returns a boolean if a field has been set.

### GetScsi3

`func (o *UpdateVMConfigRequest) GetScsi3() string`

GetScsi3 returns the Scsi3 field if non-nil, zero value otherwise.

### GetScsi3Ok

`func (o *UpdateVMConfigRequest) GetScsi3Ok() (*string, bool)`

GetScsi3Ok returns a tuple with the Scsi3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi3

`func (o *UpdateVMConfigRequest) SetScsi3(v string)`

SetScsi3 sets Scsi3 field to given value.

### HasScsi3

`func (o *UpdateVMConfigRequest) HasScsi3() bool`

HasScsi3 returns a boolean if a field has been set.

### GetScsi4

`func (o *UpdateVMConfigRequest) GetScsi4() string`

GetScsi4 returns the Scsi4 field if non-nil, zero value otherwise.

### GetScsi4Ok

`func (o *UpdateVMConfigRequest) GetScsi4Ok() (*string, bool)`

GetScsi4Ok returns a tuple with the Scsi4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi4

`func (o *UpdateVMConfigRequest) SetScsi4(v string)`

SetScsi4 sets Scsi4 field to given value.

### HasScsi4

`func (o *UpdateVMConfigRequest) HasScsi4() bool`

HasScsi4 returns a boolean if a field has been set.

### GetScsi5

`func (o *UpdateVMConfigRequest) GetScsi5() string`

GetScsi5 returns the Scsi5 field if non-nil, zero value otherwise.

### GetScsi5Ok

`func (o *UpdateVMConfigRequest) GetScsi5Ok() (*string, bool)`

GetScsi5Ok returns a tuple with the Scsi5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi5

`func (o *UpdateVMConfigRequest) SetScsi5(v string)`

SetScsi5 sets Scsi5 field to given value.

### HasScsi5

`func (o *UpdateVMConfigRequest) HasScsi5() bool`

HasScsi5 returns a boolean if a field has been set.

### GetScsi6

`func (o *UpdateVMConfigRequest) GetScsi6() string`

GetScsi6 returns the Scsi6 field if non-nil, zero value otherwise.

### GetScsi6Ok

`func (o *UpdateVMConfigRequest) GetScsi6Ok() (*string, bool)`

GetScsi6Ok returns a tuple with the Scsi6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi6

`func (o *UpdateVMConfigRequest) SetScsi6(v string)`

SetScsi6 sets Scsi6 field to given value.

### HasScsi6

`func (o *UpdateVMConfigRequest) HasScsi6() bool`

HasScsi6 returns a boolean if a field has been set.

### GetScsi7

`func (o *UpdateVMConfigRequest) GetScsi7() string`

GetScsi7 returns the Scsi7 field if non-nil, zero value otherwise.

### GetScsi7Ok

`func (o *UpdateVMConfigRequest) GetScsi7Ok() (*string, bool)`

GetScsi7Ok returns a tuple with the Scsi7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi7

`func (o *UpdateVMConfigRequest) SetScsi7(v string)`

SetScsi7 sets Scsi7 field to given value.

### HasScsi7

`func (o *UpdateVMConfigRequest) HasScsi7() bool`

HasScsi7 returns a boolean if a field has been set.

### GetScsi8

`func (o *UpdateVMConfigRequest) GetScsi8() string`

GetScsi8 returns the Scsi8 field if non-nil, zero value otherwise.

### GetScsi8Ok

`func (o *UpdateVMConfigRequest) GetScsi8Ok() (*string, bool)`

GetScsi8Ok returns a tuple with the Scsi8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi8

`func (o *UpdateVMConfigRequest) SetScsi8(v string)`

SetScsi8 sets Scsi8 field to given value.

### HasScsi8

`func (o *UpdateVMConfigRequest) HasScsi8() bool`

HasScsi8 returns a boolean if a field has been set.

### GetScsi9

`func (o *UpdateVMConfigRequest) GetScsi9() string`

GetScsi9 returns the Scsi9 field if non-nil, zero value otherwise.

### GetScsi9Ok

`func (o *UpdateVMConfigRequest) GetScsi9Ok() (*string, bool)`

GetScsi9Ok returns a tuple with the Scsi9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi9

`func (o *UpdateVMConfigRequest) SetScsi9(v string)`

SetScsi9 sets Scsi9 field to given value.

### HasScsi9

`func (o *UpdateVMConfigRequest) HasScsi9() bool`

HasScsi9 returns a boolean if a field has been set.

### GetScsi10

`func (o *UpdateVMConfigRequest) GetScsi10() string`

GetScsi10 returns the Scsi10 field if non-nil, zero value otherwise.

### GetScsi10Ok

`func (o *UpdateVMConfigRequest) GetScsi10Ok() (*string, bool)`

GetScsi10Ok returns a tuple with the Scsi10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi10

`func (o *UpdateVMConfigRequest) SetScsi10(v string)`

SetScsi10 sets Scsi10 field to given value.

### HasScsi10

`func (o *UpdateVMConfigRequest) HasScsi10() bool`

HasScsi10 returns a boolean if a field has been set.

### GetScsi11

`func (o *UpdateVMConfigRequest) GetScsi11() string`

GetScsi11 returns the Scsi11 field if non-nil, zero value otherwise.

### GetScsi11Ok

`func (o *UpdateVMConfigRequest) GetScsi11Ok() (*string, bool)`

GetScsi11Ok returns a tuple with the Scsi11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi11

`func (o *UpdateVMConfigRequest) SetScsi11(v string)`

SetScsi11 sets Scsi11 field to given value.

### HasScsi11

`func (o *UpdateVMConfigRequest) HasScsi11() bool`

HasScsi11 returns a boolean if a field has been set.

### GetScsi12

`func (o *UpdateVMConfigRequest) GetScsi12() string`

GetScsi12 returns the Scsi12 field if non-nil, zero value otherwise.

### GetScsi12Ok

`func (o *UpdateVMConfigRequest) GetScsi12Ok() (*string, bool)`

GetScsi12Ok returns a tuple with the Scsi12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi12

`func (o *UpdateVMConfigRequest) SetScsi12(v string)`

SetScsi12 sets Scsi12 field to given value.

### HasScsi12

`func (o *UpdateVMConfigRequest) HasScsi12() bool`

HasScsi12 returns a boolean if a field has been set.

### GetScsi13

`func (o *UpdateVMConfigRequest) GetScsi13() string`

GetScsi13 returns the Scsi13 field if non-nil, zero value otherwise.

### GetScsi13Ok

`func (o *UpdateVMConfigRequest) GetScsi13Ok() (*string, bool)`

GetScsi13Ok returns a tuple with the Scsi13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi13

`func (o *UpdateVMConfigRequest) SetScsi13(v string)`

SetScsi13 sets Scsi13 field to given value.

### HasScsi13

`func (o *UpdateVMConfigRequest) HasScsi13() bool`

HasScsi13 returns a boolean if a field has been set.

### GetScsi14

`func (o *UpdateVMConfigRequest) GetScsi14() string`

GetScsi14 returns the Scsi14 field if non-nil, zero value otherwise.

### GetScsi14Ok

`func (o *UpdateVMConfigRequest) GetScsi14Ok() (*string, bool)`

GetScsi14Ok returns a tuple with the Scsi14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi14

`func (o *UpdateVMConfigRequest) SetScsi14(v string)`

SetScsi14 sets Scsi14 field to given value.

### HasScsi14

`func (o *UpdateVMConfigRequest) HasScsi14() bool`

HasScsi14 returns a boolean if a field has been set.

### GetScsi15

`func (o *UpdateVMConfigRequest) GetScsi15() string`

GetScsi15 returns the Scsi15 field if non-nil, zero value otherwise.

### GetScsi15Ok

`func (o *UpdateVMConfigRequest) GetScsi15Ok() (*string, bool)`

GetScsi15Ok returns a tuple with the Scsi15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi15

`func (o *UpdateVMConfigRequest) SetScsi15(v string)`

SetScsi15 sets Scsi15 field to given value.

### HasScsi15

`func (o *UpdateVMConfigRequest) HasScsi15() bool`

HasScsi15 returns a boolean if a field has been set.

### GetScsi16

`func (o *UpdateVMConfigRequest) GetScsi16() string`

GetScsi16 returns the Scsi16 field if non-nil, zero value otherwise.

### GetScsi16Ok

`func (o *UpdateVMConfigRequest) GetScsi16Ok() (*string, bool)`

GetScsi16Ok returns a tuple with the Scsi16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi16

`func (o *UpdateVMConfigRequest) SetScsi16(v string)`

SetScsi16 sets Scsi16 field to given value.

### HasScsi16

`func (o *UpdateVMConfigRequest) HasScsi16() bool`

HasScsi16 returns a boolean if a field has been set.

### GetScsi17

`func (o *UpdateVMConfigRequest) GetScsi17() string`

GetScsi17 returns the Scsi17 field if non-nil, zero value otherwise.

### GetScsi17Ok

`func (o *UpdateVMConfigRequest) GetScsi17Ok() (*string, bool)`

GetScsi17Ok returns a tuple with the Scsi17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi17

`func (o *UpdateVMConfigRequest) SetScsi17(v string)`

SetScsi17 sets Scsi17 field to given value.

### HasScsi17

`func (o *UpdateVMConfigRequest) HasScsi17() bool`

HasScsi17 returns a boolean if a field has been set.

### GetScsi18

`func (o *UpdateVMConfigRequest) GetScsi18() string`

GetScsi18 returns the Scsi18 field if non-nil, zero value otherwise.

### GetScsi18Ok

`func (o *UpdateVMConfigRequest) GetScsi18Ok() (*string, bool)`

GetScsi18Ok returns a tuple with the Scsi18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi18

`func (o *UpdateVMConfigRequest) SetScsi18(v string)`

SetScsi18 sets Scsi18 field to given value.

### HasScsi18

`func (o *UpdateVMConfigRequest) HasScsi18() bool`

HasScsi18 returns a boolean if a field has been set.

### GetScsi19

`func (o *UpdateVMConfigRequest) GetScsi19() string`

GetScsi19 returns the Scsi19 field if non-nil, zero value otherwise.

### GetScsi19Ok

`func (o *UpdateVMConfigRequest) GetScsi19Ok() (*string, bool)`

GetScsi19Ok returns a tuple with the Scsi19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi19

`func (o *UpdateVMConfigRequest) SetScsi19(v string)`

SetScsi19 sets Scsi19 field to given value.

### HasScsi19

`func (o *UpdateVMConfigRequest) HasScsi19() bool`

HasScsi19 returns a boolean if a field has been set.

### GetScsi20

`func (o *UpdateVMConfigRequest) GetScsi20() string`

GetScsi20 returns the Scsi20 field if non-nil, zero value otherwise.

### GetScsi20Ok

`func (o *UpdateVMConfigRequest) GetScsi20Ok() (*string, bool)`

GetScsi20Ok returns a tuple with the Scsi20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi20

`func (o *UpdateVMConfigRequest) SetScsi20(v string)`

SetScsi20 sets Scsi20 field to given value.

### HasScsi20

`func (o *UpdateVMConfigRequest) HasScsi20() bool`

HasScsi20 returns a boolean if a field has been set.

### GetScsi21

`func (o *UpdateVMConfigRequest) GetScsi21() string`

GetScsi21 returns the Scsi21 field if non-nil, zero value otherwise.

### GetScsi21Ok

`func (o *UpdateVMConfigRequest) GetScsi21Ok() (*string, bool)`

GetScsi21Ok returns a tuple with the Scsi21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi21

`func (o *UpdateVMConfigRequest) SetScsi21(v string)`

SetScsi21 sets Scsi21 field to given value.

### HasScsi21

`func (o *UpdateVMConfigRequest) HasScsi21() bool`

HasScsi21 returns a boolean if a field has been set.

### GetScsi22

`func (o *UpdateVMConfigRequest) GetScsi22() string`

GetScsi22 returns the Scsi22 field if non-nil, zero value otherwise.

### GetScsi22Ok

`func (o *UpdateVMConfigRequest) GetScsi22Ok() (*string, bool)`

GetScsi22Ok returns a tuple with the Scsi22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi22

`func (o *UpdateVMConfigRequest) SetScsi22(v string)`

SetScsi22 sets Scsi22 field to given value.

### HasScsi22

`func (o *UpdateVMConfigRequest) HasScsi22() bool`

HasScsi22 returns a boolean if a field has been set.

### GetScsi23

`func (o *UpdateVMConfigRequest) GetScsi23() string`

GetScsi23 returns the Scsi23 field if non-nil, zero value otherwise.

### GetScsi23Ok

`func (o *UpdateVMConfigRequest) GetScsi23Ok() (*string, bool)`

GetScsi23Ok returns a tuple with the Scsi23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi23

`func (o *UpdateVMConfigRequest) SetScsi23(v string)`

SetScsi23 sets Scsi23 field to given value.

### HasScsi23

`func (o *UpdateVMConfigRequest) HasScsi23() bool`

HasScsi23 returns a boolean if a field has been set.

### GetScsi24

`func (o *UpdateVMConfigRequest) GetScsi24() string`

GetScsi24 returns the Scsi24 field if non-nil, zero value otherwise.

### GetScsi24Ok

`func (o *UpdateVMConfigRequest) GetScsi24Ok() (*string, bool)`

GetScsi24Ok returns a tuple with the Scsi24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi24

`func (o *UpdateVMConfigRequest) SetScsi24(v string)`

SetScsi24 sets Scsi24 field to given value.

### HasScsi24

`func (o *UpdateVMConfigRequest) HasScsi24() bool`

HasScsi24 returns a boolean if a field has been set.

### GetScsi25

`func (o *UpdateVMConfigRequest) GetScsi25() string`

GetScsi25 returns the Scsi25 field if non-nil, zero value otherwise.

### GetScsi25Ok

`func (o *UpdateVMConfigRequest) GetScsi25Ok() (*string, bool)`

GetScsi25Ok returns a tuple with the Scsi25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi25

`func (o *UpdateVMConfigRequest) SetScsi25(v string)`

SetScsi25 sets Scsi25 field to given value.

### HasScsi25

`func (o *UpdateVMConfigRequest) HasScsi25() bool`

HasScsi25 returns a boolean if a field has been set.

### GetScsi26

`func (o *UpdateVMConfigRequest) GetScsi26() string`

GetScsi26 returns the Scsi26 field if non-nil, zero value otherwise.

### GetScsi26Ok

`func (o *UpdateVMConfigRequest) GetScsi26Ok() (*string, bool)`

GetScsi26Ok returns a tuple with the Scsi26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi26

`func (o *UpdateVMConfigRequest) SetScsi26(v string)`

SetScsi26 sets Scsi26 field to given value.

### HasScsi26

`func (o *UpdateVMConfigRequest) HasScsi26() bool`

HasScsi26 returns a boolean if a field has been set.

### GetScsi27

`func (o *UpdateVMConfigRequest) GetScsi27() string`

GetScsi27 returns the Scsi27 field if non-nil, zero value otherwise.

### GetScsi27Ok

`func (o *UpdateVMConfigRequest) GetScsi27Ok() (*string, bool)`

GetScsi27Ok returns a tuple with the Scsi27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi27

`func (o *UpdateVMConfigRequest) SetScsi27(v string)`

SetScsi27 sets Scsi27 field to given value.

### HasScsi27

`func (o *UpdateVMConfigRequest) HasScsi27() bool`

HasScsi27 returns a boolean if a field has been set.

### GetScsi28

`func (o *UpdateVMConfigRequest) GetScsi28() string`

GetScsi28 returns the Scsi28 field if non-nil, zero value otherwise.

### GetScsi28Ok

`func (o *UpdateVMConfigRequest) GetScsi28Ok() (*string, bool)`

GetScsi28Ok returns a tuple with the Scsi28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi28

`func (o *UpdateVMConfigRequest) SetScsi28(v string)`

SetScsi28 sets Scsi28 field to given value.

### HasScsi28

`func (o *UpdateVMConfigRequest) HasScsi28() bool`

HasScsi28 returns a boolean if a field has been set.

### GetScsi29

`func (o *UpdateVMConfigRequest) GetScsi29() string`

GetScsi29 returns the Scsi29 field if non-nil, zero value otherwise.

### GetScsi29Ok

`func (o *UpdateVMConfigRequest) GetScsi29Ok() (*string, bool)`

GetScsi29Ok returns a tuple with the Scsi29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsi29

`func (o *UpdateVMConfigRequest) SetScsi29(v string)`

SetScsi29 sets Scsi29 field to given value.

### HasScsi29

`func (o *UpdateVMConfigRequest) HasScsi29() bool`

HasScsi29 returns a boolean if a field has been set.

### GetScsihw

`func (o *UpdateVMConfigRequest) GetScsihw() string`

GetScsihw returns the Scsihw field if non-nil, zero value otherwise.

### GetScsihwOk

`func (o *UpdateVMConfigRequest) GetScsihwOk() (*string, bool)`

GetScsihwOk returns a tuple with the Scsihw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScsihw

`func (o *UpdateVMConfigRequest) SetScsihw(v string)`

SetScsihw sets Scsihw field to given value.

### HasScsihw

`func (o *UpdateVMConfigRequest) HasScsihw() bool`

HasScsihw returns a boolean if a field has been set.

### GetSearchdomain

`func (o *UpdateVMConfigRequest) GetSearchdomain() string`

GetSearchdomain returns the Searchdomain field if non-nil, zero value otherwise.

### GetSearchdomainOk

`func (o *UpdateVMConfigRequest) GetSearchdomainOk() (*string, bool)`

GetSearchdomainOk returns a tuple with the Searchdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchdomain

`func (o *UpdateVMConfigRequest) SetSearchdomain(v string)`

SetSearchdomain sets Searchdomain field to given value.

### HasSearchdomain

`func (o *UpdateVMConfigRequest) HasSearchdomain() bool`

HasSearchdomain returns a boolean if a field has been set.

### GetSerial0

`func (o *UpdateVMConfigRequest) GetSerial0() string`

GetSerial0 returns the Serial0 field if non-nil, zero value otherwise.

### GetSerial0Ok

`func (o *UpdateVMConfigRequest) GetSerial0Ok() (*string, bool)`

GetSerial0Ok returns a tuple with the Serial0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial0

`func (o *UpdateVMConfigRequest) SetSerial0(v string)`

SetSerial0 sets Serial0 field to given value.

### HasSerial0

`func (o *UpdateVMConfigRequest) HasSerial0() bool`

HasSerial0 returns a boolean if a field has been set.

### GetSerial1

`func (o *UpdateVMConfigRequest) GetSerial1() string`

GetSerial1 returns the Serial1 field if non-nil, zero value otherwise.

### GetSerial1Ok

`func (o *UpdateVMConfigRequest) GetSerial1Ok() (*string, bool)`

GetSerial1Ok returns a tuple with the Serial1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial1

`func (o *UpdateVMConfigRequest) SetSerial1(v string)`

SetSerial1 sets Serial1 field to given value.

### HasSerial1

`func (o *UpdateVMConfigRequest) HasSerial1() bool`

HasSerial1 returns a boolean if a field has been set.

### GetSerial2

`func (o *UpdateVMConfigRequest) GetSerial2() string`

GetSerial2 returns the Serial2 field if non-nil, zero value otherwise.

### GetSerial2Ok

`func (o *UpdateVMConfigRequest) GetSerial2Ok() (*string, bool)`

GetSerial2Ok returns a tuple with the Serial2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial2

`func (o *UpdateVMConfigRequest) SetSerial2(v string)`

SetSerial2 sets Serial2 field to given value.

### HasSerial2

`func (o *UpdateVMConfigRequest) HasSerial2() bool`

HasSerial2 returns a boolean if a field has been set.

### GetSerial3

`func (o *UpdateVMConfigRequest) GetSerial3() string`

GetSerial3 returns the Serial3 field if non-nil, zero value otherwise.

### GetSerial3Ok

`func (o *UpdateVMConfigRequest) GetSerial3Ok() (*string, bool)`

GetSerial3Ok returns a tuple with the Serial3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial3

`func (o *UpdateVMConfigRequest) SetSerial3(v string)`

SetSerial3 sets Serial3 field to given value.

### HasSerial3

`func (o *UpdateVMConfigRequest) HasSerial3() bool`

HasSerial3 returns a boolean if a field has been set.

### GetShares

`func (o *UpdateVMConfigRequest) GetShares() int64`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *UpdateVMConfigRequest) GetSharesOk() (*int64, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *UpdateVMConfigRequest) SetShares(v int64)`

SetShares sets Shares field to given value.

### HasShares

`func (o *UpdateVMConfigRequest) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSkiplock

`func (o *UpdateVMConfigRequest) GetSkiplock() bool`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *UpdateVMConfigRequest) GetSkiplockOk() (*bool, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *UpdateVMConfigRequest) SetSkiplock(v bool)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *UpdateVMConfigRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.

### GetSmbios1

`func (o *UpdateVMConfigRequest) GetSmbios1() string`

GetSmbios1 returns the Smbios1 field if non-nil, zero value otherwise.

### GetSmbios1Ok

`func (o *UpdateVMConfigRequest) GetSmbios1Ok() (*string, bool)`

GetSmbios1Ok returns a tuple with the Smbios1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbios1

`func (o *UpdateVMConfigRequest) SetSmbios1(v string)`

SetSmbios1 sets Smbios1 field to given value.

### HasSmbios1

`func (o *UpdateVMConfigRequest) HasSmbios1() bool`

HasSmbios1 returns a boolean if a field has been set.

### GetSmp

`func (o *UpdateVMConfigRequest) GetSmp() int64`

GetSmp returns the Smp field if non-nil, zero value otherwise.

### GetSmpOk

`func (o *UpdateVMConfigRequest) GetSmpOk() (*int64, bool)`

GetSmpOk returns a tuple with the Smp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmp

`func (o *UpdateVMConfigRequest) SetSmp(v int64)`

SetSmp sets Smp field to given value.

### HasSmp

`func (o *UpdateVMConfigRequest) HasSmp() bool`

HasSmp returns a boolean if a field has been set.

### GetSockets

`func (o *UpdateVMConfigRequest) GetSockets() int64`

GetSockets returns the Sockets field if non-nil, zero value otherwise.

### GetSocketsOk

`func (o *UpdateVMConfigRequest) GetSocketsOk() (*int64, bool)`

GetSocketsOk returns a tuple with the Sockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSockets

`func (o *UpdateVMConfigRequest) SetSockets(v int64)`

SetSockets sets Sockets field to given value.

### HasSockets

`func (o *UpdateVMConfigRequest) HasSockets() bool`

HasSockets returns a boolean if a field has been set.

### GetSpiceEnhancements

`func (o *UpdateVMConfigRequest) GetSpiceEnhancements() string`

GetSpiceEnhancements returns the SpiceEnhancements field if non-nil, zero value otherwise.

### GetSpiceEnhancementsOk

`func (o *UpdateVMConfigRequest) GetSpiceEnhancementsOk() (*string, bool)`

GetSpiceEnhancementsOk returns a tuple with the SpiceEnhancements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpiceEnhancements

`func (o *UpdateVMConfigRequest) SetSpiceEnhancements(v string)`

SetSpiceEnhancements sets SpiceEnhancements field to given value.

### HasSpiceEnhancements

`func (o *UpdateVMConfigRequest) HasSpiceEnhancements() bool`

HasSpiceEnhancements returns a boolean if a field has been set.

### GetSshkeys

`func (o *UpdateVMConfigRequest) GetSshkeys() string`

GetSshkeys returns the Sshkeys field if non-nil, zero value otherwise.

### GetSshkeysOk

`func (o *UpdateVMConfigRequest) GetSshkeysOk() (*string, bool)`

GetSshkeysOk returns a tuple with the Sshkeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshkeys

`func (o *UpdateVMConfigRequest) SetSshkeys(v string)`

SetSshkeys sets Sshkeys field to given value.

### HasSshkeys

`func (o *UpdateVMConfigRequest) HasSshkeys() bool`

HasSshkeys returns a boolean if a field has been set.

### GetStartdate

`func (o *UpdateVMConfigRequest) GetStartdate() string`

GetStartdate returns the Startdate field if non-nil, zero value otherwise.

### GetStartdateOk

`func (o *UpdateVMConfigRequest) GetStartdateOk() (*string, bool)`

GetStartdateOk returns a tuple with the Startdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartdate

`func (o *UpdateVMConfigRequest) SetStartdate(v string)`

SetStartdate sets Startdate field to given value.

### HasStartdate

`func (o *UpdateVMConfigRequest) HasStartdate() bool`

HasStartdate returns a boolean if a field has been set.

### GetStartup

`func (o *UpdateVMConfigRequest) GetStartup() string`

GetStartup returns the Startup field if non-nil, zero value otherwise.

### GetStartupOk

`func (o *UpdateVMConfigRequest) GetStartupOk() (*string, bool)`

GetStartupOk returns a tuple with the Startup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartup

`func (o *UpdateVMConfigRequest) SetStartup(v string)`

SetStartup sets Startup field to given value.

### HasStartup

`func (o *UpdateVMConfigRequest) HasStartup() bool`

HasStartup returns a boolean if a field has been set.

### GetTablet

`func (o *UpdateVMConfigRequest) GetTablet() bool`

GetTablet returns the Tablet field if non-nil, zero value otherwise.

### GetTabletOk

`func (o *UpdateVMConfigRequest) GetTabletOk() (*bool, bool)`

GetTabletOk returns a tuple with the Tablet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablet

`func (o *UpdateVMConfigRequest) SetTablet(v bool)`

SetTablet sets Tablet field to given value.

### HasTablet

`func (o *UpdateVMConfigRequest) HasTablet() bool`

HasTablet returns a boolean if a field has been set.

### GetTags

`func (o *UpdateVMConfigRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateVMConfigRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateVMConfigRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateVMConfigRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTdf

`func (o *UpdateVMConfigRequest) GetTdf() bool`

GetTdf returns the Tdf field if non-nil, zero value otherwise.

### GetTdfOk

`func (o *UpdateVMConfigRequest) GetTdfOk() (*bool, bool)`

GetTdfOk returns a tuple with the Tdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdf

`func (o *UpdateVMConfigRequest) SetTdf(v bool)`

SetTdf sets Tdf field to given value.

### HasTdf

`func (o *UpdateVMConfigRequest) HasTdf() bool`

HasTdf returns a boolean if a field has been set.

### GetTemplate

`func (o *UpdateVMConfigRequest) GetTemplate() bool`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *UpdateVMConfigRequest) GetTemplateOk() (*bool, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *UpdateVMConfigRequest) SetTemplate(v bool)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *UpdateVMConfigRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTpmstate0

`func (o *UpdateVMConfigRequest) GetTpmstate0() string`

GetTpmstate0 returns the Tpmstate0 field if non-nil, zero value otherwise.

### GetTpmstate0Ok

`func (o *UpdateVMConfigRequest) GetTpmstate0Ok() (*string, bool)`

GetTpmstate0Ok returns a tuple with the Tpmstate0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTpmstate0

`func (o *UpdateVMConfigRequest) SetTpmstate0(v string)`

SetTpmstate0 sets Tpmstate0 field to given value.

### HasTpmstate0

`func (o *UpdateVMConfigRequest) HasTpmstate0() bool`

HasTpmstate0 returns a boolean if a field has been set.

### GetUnused0

`func (o *UpdateVMConfigRequest) GetUnused0() string`

GetUnused0 returns the Unused0 field if non-nil, zero value otherwise.

### GetUnused0Ok

`func (o *UpdateVMConfigRequest) GetUnused0Ok() (*string, bool)`

GetUnused0Ok returns a tuple with the Unused0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused0

`func (o *UpdateVMConfigRequest) SetUnused0(v string)`

SetUnused0 sets Unused0 field to given value.

### HasUnused0

`func (o *UpdateVMConfigRequest) HasUnused0() bool`

HasUnused0 returns a boolean if a field has been set.

### GetUnused1

`func (o *UpdateVMConfigRequest) GetUnused1() string`

GetUnused1 returns the Unused1 field if non-nil, zero value otherwise.

### GetUnused1Ok

`func (o *UpdateVMConfigRequest) GetUnused1Ok() (*string, bool)`

GetUnused1Ok returns a tuple with the Unused1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused1

`func (o *UpdateVMConfigRequest) SetUnused1(v string)`

SetUnused1 sets Unused1 field to given value.

### HasUnused1

`func (o *UpdateVMConfigRequest) HasUnused1() bool`

HasUnused1 returns a boolean if a field has been set.

### GetUnused2

`func (o *UpdateVMConfigRequest) GetUnused2() string`

GetUnused2 returns the Unused2 field if non-nil, zero value otherwise.

### GetUnused2Ok

`func (o *UpdateVMConfigRequest) GetUnused2Ok() (*string, bool)`

GetUnused2Ok returns a tuple with the Unused2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused2

`func (o *UpdateVMConfigRequest) SetUnused2(v string)`

SetUnused2 sets Unused2 field to given value.

### HasUnused2

`func (o *UpdateVMConfigRequest) HasUnused2() bool`

HasUnused2 returns a boolean if a field has been set.

### GetUnused3

`func (o *UpdateVMConfigRequest) GetUnused3() string`

GetUnused3 returns the Unused3 field if non-nil, zero value otherwise.

### GetUnused3Ok

`func (o *UpdateVMConfigRequest) GetUnused3Ok() (*string, bool)`

GetUnused3Ok returns a tuple with the Unused3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused3

`func (o *UpdateVMConfigRequest) SetUnused3(v string)`

SetUnused3 sets Unused3 field to given value.

### HasUnused3

`func (o *UpdateVMConfigRequest) HasUnused3() bool`

HasUnused3 returns a boolean if a field has been set.

### GetUnused4

`func (o *UpdateVMConfigRequest) GetUnused4() string`

GetUnused4 returns the Unused4 field if non-nil, zero value otherwise.

### GetUnused4Ok

`func (o *UpdateVMConfigRequest) GetUnused4Ok() (*string, bool)`

GetUnused4Ok returns a tuple with the Unused4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused4

`func (o *UpdateVMConfigRequest) SetUnused4(v string)`

SetUnused4 sets Unused4 field to given value.

### HasUnused4

`func (o *UpdateVMConfigRequest) HasUnused4() bool`

HasUnused4 returns a boolean if a field has been set.

### GetUnused5

`func (o *UpdateVMConfigRequest) GetUnused5() string`

GetUnused5 returns the Unused5 field if non-nil, zero value otherwise.

### GetUnused5Ok

`func (o *UpdateVMConfigRequest) GetUnused5Ok() (*string, bool)`

GetUnused5Ok returns a tuple with the Unused5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused5

`func (o *UpdateVMConfigRequest) SetUnused5(v string)`

SetUnused5 sets Unused5 field to given value.

### HasUnused5

`func (o *UpdateVMConfigRequest) HasUnused5() bool`

HasUnused5 returns a boolean if a field has been set.

### GetUnused6

`func (o *UpdateVMConfigRequest) GetUnused6() string`

GetUnused6 returns the Unused6 field if non-nil, zero value otherwise.

### GetUnused6Ok

`func (o *UpdateVMConfigRequest) GetUnused6Ok() (*string, bool)`

GetUnused6Ok returns a tuple with the Unused6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused6

`func (o *UpdateVMConfigRequest) SetUnused6(v string)`

SetUnused6 sets Unused6 field to given value.

### HasUnused6

`func (o *UpdateVMConfigRequest) HasUnused6() bool`

HasUnused6 returns a boolean if a field has been set.

### GetUnused7

`func (o *UpdateVMConfigRequest) GetUnused7() string`

GetUnused7 returns the Unused7 field if non-nil, zero value otherwise.

### GetUnused7Ok

`func (o *UpdateVMConfigRequest) GetUnused7Ok() (*string, bool)`

GetUnused7Ok returns a tuple with the Unused7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused7

`func (o *UpdateVMConfigRequest) SetUnused7(v string)`

SetUnused7 sets Unused7 field to given value.

### HasUnused7

`func (o *UpdateVMConfigRequest) HasUnused7() bool`

HasUnused7 returns a boolean if a field has been set.

### GetUnused8

`func (o *UpdateVMConfigRequest) GetUnused8() string`

GetUnused8 returns the Unused8 field if non-nil, zero value otherwise.

### GetUnused8Ok

`func (o *UpdateVMConfigRequest) GetUnused8Ok() (*string, bool)`

GetUnused8Ok returns a tuple with the Unused8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused8

`func (o *UpdateVMConfigRequest) SetUnused8(v string)`

SetUnused8 sets Unused8 field to given value.

### HasUnused8

`func (o *UpdateVMConfigRequest) HasUnused8() bool`

HasUnused8 returns a boolean if a field has been set.

### GetUnused9

`func (o *UpdateVMConfigRequest) GetUnused9() string`

GetUnused9 returns the Unused9 field if non-nil, zero value otherwise.

### GetUnused9Ok

`func (o *UpdateVMConfigRequest) GetUnused9Ok() (*string, bool)`

GetUnused9Ok returns a tuple with the Unused9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused9

`func (o *UpdateVMConfigRequest) SetUnused9(v string)`

SetUnused9 sets Unused9 field to given value.

### HasUnused9

`func (o *UpdateVMConfigRequest) HasUnused9() bool`

HasUnused9 returns a boolean if a field has been set.

### GetUnused10

`func (o *UpdateVMConfigRequest) GetUnused10() string`

GetUnused10 returns the Unused10 field if non-nil, zero value otherwise.

### GetUnused10Ok

`func (o *UpdateVMConfigRequest) GetUnused10Ok() (*string, bool)`

GetUnused10Ok returns a tuple with the Unused10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused10

`func (o *UpdateVMConfigRequest) SetUnused10(v string)`

SetUnused10 sets Unused10 field to given value.

### HasUnused10

`func (o *UpdateVMConfigRequest) HasUnused10() bool`

HasUnused10 returns a boolean if a field has been set.

### GetUnused11

`func (o *UpdateVMConfigRequest) GetUnused11() string`

GetUnused11 returns the Unused11 field if non-nil, zero value otherwise.

### GetUnused11Ok

`func (o *UpdateVMConfigRequest) GetUnused11Ok() (*string, bool)`

GetUnused11Ok returns a tuple with the Unused11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused11

`func (o *UpdateVMConfigRequest) SetUnused11(v string)`

SetUnused11 sets Unused11 field to given value.

### HasUnused11

`func (o *UpdateVMConfigRequest) HasUnused11() bool`

HasUnused11 returns a boolean if a field has been set.

### GetUnused12

`func (o *UpdateVMConfigRequest) GetUnused12() string`

GetUnused12 returns the Unused12 field if non-nil, zero value otherwise.

### GetUnused12Ok

`func (o *UpdateVMConfigRequest) GetUnused12Ok() (*string, bool)`

GetUnused12Ok returns a tuple with the Unused12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused12

`func (o *UpdateVMConfigRequest) SetUnused12(v string)`

SetUnused12 sets Unused12 field to given value.

### HasUnused12

`func (o *UpdateVMConfigRequest) HasUnused12() bool`

HasUnused12 returns a boolean if a field has been set.

### GetUnused13

`func (o *UpdateVMConfigRequest) GetUnused13() string`

GetUnused13 returns the Unused13 field if non-nil, zero value otherwise.

### GetUnused13Ok

`func (o *UpdateVMConfigRequest) GetUnused13Ok() (*string, bool)`

GetUnused13Ok returns a tuple with the Unused13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused13

`func (o *UpdateVMConfigRequest) SetUnused13(v string)`

SetUnused13 sets Unused13 field to given value.

### HasUnused13

`func (o *UpdateVMConfigRequest) HasUnused13() bool`

HasUnused13 returns a boolean if a field has been set.

### GetUnused14

`func (o *UpdateVMConfigRequest) GetUnused14() string`

GetUnused14 returns the Unused14 field if non-nil, zero value otherwise.

### GetUnused14Ok

`func (o *UpdateVMConfigRequest) GetUnused14Ok() (*string, bool)`

GetUnused14Ok returns a tuple with the Unused14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused14

`func (o *UpdateVMConfigRequest) SetUnused14(v string)`

SetUnused14 sets Unused14 field to given value.

### HasUnused14

`func (o *UpdateVMConfigRequest) HasUnused14() bool`

HasUnused14 returns a boolean if a field has been set.

### GetUnused15

`func (o *UpdateVMConfigRequest) GetUnused15() string`

GetUnused15 returns the Unused15 field if non-nil, zero value otherwise.

### GetUnused15Ok

`func (o *UpdateVMConfigRequest) GetUnused15Ok() (*string, bool)`

GetUnused15Ok returns a tuple with the Unused15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused15

`func (o *UpdateVMConfigRequest) SetUnused15(v string)`

SetUnused15 sets Unused15 field to given value.

### HasUnused15

`func (o *UpdateVMConfigRequest) HasUnused15() bool`

HasUnused15 returns a boolean if a field has been set.

### GetUnused16

`func (o *UpdateVMConfigRequest) GetUnused16() string`

GetUnused16 returns the Unused16 field if non-nil, zero value otherwise.

### GetUnused16Ok

`func (o *UpdateVMConfigRequest) GetUnused16Ok() (*string, bool)`

GetUnused16Ok returns a tuple with the Unused16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused16

`func (o *UpdateVMConfigRequest) SetUnused16(v string)`

SetUnused16 sets Unused16 field to given value.

### HasUnused16

`func (o *UpdateVMConfigRequest) HasUnused16() bool`

HasUnused16 returns a boolean if a field has been set.

### GetUnused17

`func (o *UpdateVMConfigRequest) GetUnused17() string`

GetUnused17 returns the Unused17 field if non-nil, zero value otherwise.

### GetUnused17Ok

`func (o *UpdateVMConfigRequest) GetUnused17Ok() (*string, bool)`

GetUnused17Ok returns a tuple with the Unused17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused17

`func (o *UpdateVMConfigRequest) SetUnused17(v string)`

SetUnused17 sets Unused17 field to given value.

### HasUnused17

`func (o *UpdateVMConfigRequest) HasUnused17() bool`

HasUnused17 returns a boolean if a field has been set.

### GetUnused18

`func (o *UpdateVMConfigRequest) GetUnused18() string`

GetUnused18 returns the Unused18 field if non-nil, zero value otherwise.

### GetUnused18Ok

`func (o *UpdateVMConfigRequest) GetUnused18Ok() (*string, bool)`

GetUnused18Ok returns a tuple with the Unused18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused18

`func (o *UpdateVMConfigRequest) SetUnused18(v string)`

SetUnused18 sets Unused18 field to given value.

### HasUnused18

`func (o *UpdateVMConfigRequest) HasUnused18() bool`

HasUnused18 returns a boolean if a field has been set.

### GetUnused19

`func (o *UpdateVMConfigRequest) GetUnused19() string`

GetUnused19 returns the Unused19 field if non-nil, zero value otherwise.

### GetUnused19Ok

`func (o *UpdateVMConfigRequest) GetUnused19Ok() (*string, bool)`

GetUnused19Ok returns a tuple with the Unused19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused19

`func (o *UpdateVMConfigRequest) SetUnused19(v string)`

SetUnused19 sets Unused19 field to given value.

### HasUnused19

`func (o *UpdateVMConfigRequest) HasUnused19() bool`

HasUnused19 returns a boolean if a field has been set.

### GetUnused20

`func (o *UpdateVMConfigRequest) GetUnused20() string`

GetUnused20 returns the Unused20 field if non-nil, zero value otherwise.

### GetUnused20Ok

`func (o *UpdateVMConfigRequest) GetUnused20Ok() (*string, bool)`

GetUnused20Ok returns a tuple with the Unused20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused20

`func (o *UpdateVMConfigRequest) SetUnused20(v string)`

SetUnused20 sets Unused20 field to given value.

### HasUnused20

`func (o *UpdateVMConfigRequest) HasUnused20() bool`

HasUnused20 returns a boolean if a field has been set.

### GetUnused21

`func (o *UpdateVMConfigRequest) GetUnused21() string`

GetUnused21 returns the Unused21 field if non-nil, zero value otherwise.

### GetUnused21Ok

`func (o *UpdateVMConfigRequest) GetUnused21Ok() (*string, bool)`

GetUnused21Ok returns a tuple with the Unused21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused21

`func (o *UpdateVMConfigRequest) SetUnused21(v string)`

SetUnused21 sets Unused21 field to given value.

### HasUnused21

`func (o *UpdateVMConfigRequest) HasUnused21() bool`

HasUnused21 returns a boolean if a field has been set.

### GetUnused22

`func (o *UpdateVMConfigRequest) GetUnused22() string`

GetUnused22 returns the Unused22 field if non-nil, zero value otherwise.

### GetUnused22Ok

`func (o *UpdateVMConfigRequest) GetUnused22Ok() (*string, bool)`

GetUnused22Ok returns a tuple with the Unused22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused22

`func (o *UpdateVMConfigRequest) SetUnused22(v string)`

SetUnused22 sets Unused22 field to given value.

### HasUnused22

`func (o *UpdateVMConfigRequest) HasUnused22() bool`

HasUnused22 returns a boolean if a field has been set.

### GetUnused23

`func (o *UpdateVMConfigRequest) GetUnused23() string`

GetUnused23 returns the Unused23 field if non-nil, zero value otherwise.

### GetUnused23Ok

`func (o *UpdateVMConfigRequest) GetUnused23Ok() (*string, bool)`

GetUnused23Ok returns a tuple with the Unused23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused23

`func (o *UpdateVMConfigRequest) SetUnused23(v string)`

SetUnused23 sets Unused23 field to given value.

### HasUnused23

`func (o *UpdateVMConfigRequest) HasUnused23() bool`

HasUnused23 returns a boolean if a field has been set.

### GetUnused24

`func (o *UpdateVMConfigRequest) GetUnused24() string`

GetUnused24 returns the Unused24 field if non-nil, zero value otherwise.

### GetUnused24Ok

`func (o *UpdateVMConfigRequest) GetUnused24Ok() (*string, bool)`

GetUnused24Ok returns a tuple with the Unused24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused24

`func (o *UpdateVMConfigRequest) SetUnused24(v string)`

SetUnused24 sets Unused24 field to given value.

### HasUnused24

`func (o *UpdateVMConfigRequest) HasUnused24() bool`

HasUnused24 returns a boolean if a field has been set.

### GetUnused25

`func (o *UpdateVMConfigRequest) GetUnused25() string`

GetUnused25 returns the Unused25 field if non-nil, zero value otherwise.

### GetUnused25Ok

`func (o *UpdateVMConfigRequest) GetUnused25Ok() (*string, bool)`

GetUnused25Ok returns a tuple with the Unused25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused25

`func (o *UpdateVMConfigRequest) SetUnused25(v string)`

SetUnused25 sets Unused25 field to given value.

### HasUnused25

`func (o *UpdateVMConfigRequest) HasUnused25() bool`

HasUnused25 returns a boolean if a field has been set.

### GetUnused26

`func (o *UpdateVMConfigRequest) GetUnused26() string`

GetUnused26 returns the Unused26 field if non-nil, zero value otherwise.

### GetUnused26Ok

`func (o *UpdateVMConfigRequest) GetUnused26Ok() (*string, bool)`

GetUnused26Ok returns a tuple with the Unused26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused26

`func (o *UpdateVMConfigRequest) SetUnused26(v string)`

SetUnused26 sets Unused26 field to given value.

### HasUnused26

`func (o *UpdateVMConfigRequest) HasUnused26() bool`

HasUnused26 returns a boolean if a field has been set.

### GetUnused27

`func (o *UpdateVMConfigRequest) GetUnused27() string`

GetUnused27 returns the Unused27 field if non-nil, zero value otherwise.

### GetUnused27Ok

`func (o *UpdateVMConfigRequest) GetUnused27Ok() (*string, bool)`

GetUnused27Ok returns a tuple with the Unused27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused27

`func (o *UpdateVMConfigRequest) SetUnused27(v string)`

SetUnused27 sets Unused27 field to given value.

### HasUnused27

`func (o *UpdateVMConfigRequest) HasUnused27() bool`

HasUnused27 returns a boolean if a field has been set.

### GetUnused28

`func (o *UpdateVMConfigRequest) GetUnused28() string`

GetUnused28 returns the Unused28 field if non-nil, zero value otherwise.

### GetUnused28Ok

`func (o *UpdateVMConfigRequest) GetUnused28Ok() (*string, bool)`

GetUnused28Ok returns a tuple with the Unused28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused28

`func (o *UpdateVMConfigRequest) SetUnused28(v string)`

SetUnused28 sets Unused28 field to given value.

### HasUnused28

`func (o *UpdateVMConfigRequest) HasUnused28() bool`

HasUnused28 returns a boolean if a field has been set.

### GetUnused29

`func (o *UpdateVMConfigRequest) GetUnused29() string`

GetUnused29 returns the Unused29 field if non-nil, zero value otherwise.

### GetUnused29Ok

`func (o *UpdateVMConfigRequest) GetUnused29Ok() (*string, bool)`

GetUnused29Ok returns a tuple with the Unused29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused29

`func (o *UpdateVMConfigRequest) SetUnused29(v string)`

SetUnused29 sets Unused29 field to given value.

### HasUnused29

`func (o *UpdateVMConfigRequest) HasUnused29() bool`

HasUnused29 returns a boolean if a field has been set.

### GetUsb0

`func (o *UpdateVMConfigRequest) GetUsb0() string`

GetUsb0 returns the Usb0 field if non-nil, zero value otherwise.

### GetUsb0Ok

`func (o *UpdateVMConfigRequest) GetUsb0Ok() (*string, bool)`

GetUsb0Ok returns a tuple with the Usb0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb0

`func (o *UpdateVMConfigRequest) SetUsb0(v string)`

SetUsb0 sets Usb0 field to given value.

### HasUsb0

`func (o *UpdateVMConfigRequest) HasUsb0() bool`

HasUsb0 returns a boolean if a field has been set.

### GetUsb1

`func (o *UpdateVMConfigRequest) GetUsb1() string`

GetUsb1 returns the Usb1 field if non-nil, zero value otherwise.

### GetUsb1Ok

`func (o *UpdateVMConfigRequest) GetUsb1Ok() (*string, bool)`

GetUsb1Ok returns a tuple with the Usb1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb1

`func (o *UpdateVMConfigRequest) SetUsb1(v string)`

SetUsb1 sets Usb1 field to given value.

### HasUsb1

`func (o *UpdateVMConfigRequest) HasUsb1() bool`

HasUsb1 returns a boolean if a field has been set.

### GetUsb2

`func (o *UpdateVMConfigRequest) GetUsb2() string`

GetUsb2 returns the Usb2 field if non-nil, zero value otherwise.

### GetUsb2Ok

`func (o *UpdateVMConfigRequest) GetUsb2Ok() (*string, bool)`

GetUsb2Ok returns a tuple with the Usb2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb2

`func (o *UpdateVMConfigRequest) SetUsb2(v string)`

SetUsb2 sets Usb2 field to given value.

### HasUsb2

`func (o *UpdateVMConfigRequest) HasUsb2() bool`

HasUsb2 returns a boolean if a field has been set.

### GetUsb3

`func (o *UpdateVMConfigRequest) GetUsb3() string`

GetUsb3 returns the Usb3 field if non-nil, zero value otherwise.

### GetUsb3Ok

`func (o *UpdateVMConfigRequest) GetUsb3Ok() (*string, bool)`

GetUsb3Ok returns a tuple with the Usb3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsb3

`func (o *UpdateVMConfigRequest) SetUsb3(v string)`

SetUsb3 sets Usb3 field to given value.

### HasUsb3

`func (o *UpdateVMConfigRequest) HasUsb3() bool`

HasUsb3 returns a boolean if a field has been set.

### GetVcpus

`func (o *UpdateVMConfigRequest) GetVcpus() int64`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *UpdateVMConfigRequest) GetVcpusOk() (*int64, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *UpdateVMConfigRequest) SetVcpus(v int64)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *UpdateVMConfigRequest) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.

### GetVga

`func (o *UpdateVMConfigRequest) GetVga() string`

GetVga returns the Vga field if non-nil, zero value otherwise.

### GetVgaOk

`func (o *UpdateVMConfigRequest) GetVgaOk() (*string, bool)`

GetVgaOk returns a tuple with the Vga field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVga

`func (o *UpdateVMConfigRequest) SetVga(v string)`

SetVga sets Vga field to given value.

### HasVga

`func (o *UpdateVMConfigRequest) HasVga() bool`

HasVga returns a boolean if a field has been set.

### GetVirtio0

`func (o *UpdateVMConfigRequest) GetVirtio0() string`

GetVirtio0 returns the Virtio0 field if non-nil, zero value otherwise.

### GetVirtio0Ok

`func (o *UpdateVMConfigRequest) GetVirtio0Ok() (*string, bool)`

GetVirtio0Ok returns a tuple with the Virtio0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio0

`func (o *UpdateVMConfigRequest) SetVirtio0(v string)`

SetVirtio0 sets Virtio0 field to given value.

### HasVirtio0

`func (o *UpdateVMConfigRequest) HasVirtio0() bool`

HasVirtio0 returns a boolean if a field has been set.

### GetVirtio1

`func (o *UpdateVMConfigRequest) GetVirtio1() string`

GetVirtio1 returns the Virtio1 field if non-nil, zero value otherwise.

### GetVirtio1Ok

`func (o *UpdateVMConfigRequest) GetVirtio1Ok() (*string, bool)`

GetVirtio1Ok returns a tuple with the Virtio1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio1

`func (o *UpdateVMConfigRequest) SetVirtio1(v string)`

SetVirtio1 sets Virtio1 field to given value.

### HasVirtio1

`func (o *UpdateVMConfigRequest) HasVirtio1() bool`

HasVirtio1 returns a boolean if a field has been set.

### GetVirtio2

`func (o *UpdateVMConfigRequest) GetVirtio2() string`

GetVirtio2 returns the Virtio2 field if non-nil, zero value otherwise.

### GetVirtio2Ok

`func (o *UpdateVMConfigRequest) GetVirtio2Ok() (*string, bool)`

GetVirtio2Ok returns a tuple with the Virtio2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio2

`func (o *UpdateVMConfigRequest) SetVirtio2(v string)`

SetVirtio2 sets Virtio2 field to given value.

### HasVirtio2

`func (o *UpdateVMConfigRequest) HasVirtio2() bool`

HasVirtio2 returns a boolean if a field has been set.

### GetVirtio3

`func (o *UpdateVMConfigRequest) GetVirtio3() string`

GetVirtio3 returns the Virtio3 field if non-nil, zero value otherwise.

### GetVirtio3Ok

`func (o *UpdateVMConfigRequest) GetVirtio3Ok() (*string, bool)`

GetVirtio3Ok returns a tuple with the Virtio3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio3

`func (o *UpdateVMConfigRequest) SetVirtio3(v string)`

SetVirtio3 sets Virtio3 field to given value.

### HasVirtio3

`func (o *UpdateVMConfigRequest) HasVirtio3() bool`

HasVirtio3 returns a boolean if a field has been set.

### GetVirtio4

`func (o *UpdateVMConfigRequest) GetVirtio4() string`

GetVirtio4 returns the Virtio4 field if non-nil, zero value otherwise.

### GetVirtio4Ok

`func (o *UpdateVMConfigRequest) GetVirtio4Ok() (*string, bool)`

GetVirtio4Ok returns a tuple with the Virtio4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio4

`func (o *UpdateVMConfigRequest) SetVirtio4(v string)`

SetVirtio4 sets Virtio4 field to given value.

### HasVirtio4

`func (o *UpdateVMConfigRequest) HasVirtio4() bool`

HasVirtio4 returns a boolean if a field has been set.

### GetVirtio5

`func (o *UpdateVMConfigRequest) GetVirtio5() string`

GetVirtio5 returns the Virtio5 field if non-nil, zero value otherwise.

### GetVirtio5Ok

`func (o *UpdateVMConfigRequest) GetVirtio5Ok() (*string, bool)`

GetVirtio5Ok returns a tuple with the Virtio5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio5

`func (o *UpdateVMConfigRequest) SetVirtio5(v string)`

SetVirtio5 sets Virtio5 field to given value.

### HasVirtio5

`func (o *UpdateVMConfigRequest) HasVirtio5() bool`

HasVirtio5 returns a boolean if a field has been set.

### GetVirtio6

`func (o *UpdateVMConfigRequest) GetVirtio6() string`

GetVirtio6 returns the Virtio6 field if non-nil, zero value otherwise.

### GetVirtio6Ok

`func (o *UpdateVMConfigRequest) GetVirtio6Ok() (*string, bool)`

GetVirtio6Ok returns a tuple with the Virtio6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio6

`func (o *UpdateVMConfigRequest) SetVirtio6(v string)`

SetVirtio6 sets Virtio6 field to given value.

### HasVirtio6

`func (o *UpdateVMConfigRequest) HasVirtio6() bool`

HasVirtio6 returns a boolean if a field has been set.

### GetVirtio7

`func (o *UpdateVMConfigRequest) GetVirtio7() string`

GetVirtio7 returns the Virtio7 field if non-nil, zero value otherwise.

### GetVirtio7Ok

`func (o *UpdateVMConfigRequest) GetVirtio7Ok() (*string, bool)`

GetVirtio7Ok returns a tuple with the Virtio7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio7

`func (o *UpdateVMConfigRequest) SetVirtio7(v string)`

SetVirtio7 sets Virtio7 field to given value.

### HasVirtio7

`func (o *UpdateVMConfigRequest) HasVirtio7() bool`

HasVirtio7 returns a boolean if a field has been set.

### GetVirtio8

`func (o *UpdateVMConfigRequest) GetVirtio8() string`

GetVirtio8 returns the Virtio8 field if non-nil, zero value otherwise.

### GetVirtio8Ok

`func (o *UpdateVMConfigRequest) GetVirtio8Ok() (*string, bool)`

GetVirtio8Ok returns a tuple with the Virtio8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio8

`func (o *UpdateVMConfigRequest) SetVirtio8(v string)`

SetVirtio8 sets Virtio8 field to given value.

### HasVirtio8

`func (o *UpdateVMConfigRequest) HasVirtio8() bool`

HasVirtio8 returns a boolean if a field has been set.

### GetVirtio9

`func (o *UpdateVMConfigRequest) GetVirtio9() string`

GetVirtio9 returns the Virtio9 field if non-nil, zero value otherwise.

### GetVirtio9Ok

`func (o *UpdateVMConfigRequest) GetVirtio9Ok() (*string, bool)`

GetVirtio9Ok returns a tuple with the Virtio9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio9

`func (o *UpdateVMConfigRequest) SetVirtio9(v string)`

SetVirtio9 sets Virtio9 field to given value.

### HasVirtio9

`func (o *UpdateVMConfigRequest) HasVirtio9() bool`

HasVirtio9 returns a boolean if a field has been set.

### GetVirtio10

`func (o *UpdateVMConfigRequest) GetVirtio10() string`

GetVirtio10 returns the Virtio10 field if non-nil, zero value otherwise.

### GetVirtio10Ok

`func (o *UpdateVMConfigRequest) GetVirtio10Ok() (*string, bool)`

GetVirtio10Ok returns a tuple with the Virtio10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio10

`func (o *UpdateVMConfigRequest) SetVirtio10(v string)`

SetVirtio10 sets Virtio10 field to given value.

### HasVirtio10

`func (o *UpdateVMConfigRequest) HasVirtio10() bool`

HasVirtio10 returns a boolean if a field has been set.

### GetVirtio11

`func (o *UpdateVMConfigRequest) GetVirtio11() string`

GetVirtio11 returns the Virtio11 field if non-nil, zero value otherwise.

### GetVirtio11Ok

`func (o *UpdateVMConfigRequest) GetVirtio11Ok() (*string, bool)`

GetVirtio11Ok returns a tuple with the Virtio11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio11

`func (o *UpdateVMConfigRequest) SetVirtio11(v string)`

SetVirtio11 sets Virtio11 field to given value.

### HasVirtio11

`func (o *UpdateVMConfigRequest) HasVirtio11() bool`

HasVirtio11 returns a boolean if a field has been set.

### GetVirtio12

`func (o *UpdateVMConfigRequest) GetVirtio12() string`

GetVirtio12 returns the Virtio12 field if non-nil, zero value otherwise.

### GetVirtio12Ok

`func (o *UpdateVMConfigRequest) GetVirtio12Ok() (*string, bool)`

GetVirtio12Ok returns a tuple with the Virtio12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio12

`func (o *UpdateVMConfigRequest) SetVirtio12(v string)`

SetVirtio12 sets Virtio12 field to given value.

### HasVirtio12

`func (o *UpdateVMConfigRequest) HasVirtio12() bool`

HasVirtio12 returns a boolean if a field has been set.

### GetVirtio13

`func (o *UpdateVMConfigRequest) GetVirtio13() string`

GetVirtio13 returns the Virtio13 field if non-nil, zero value otherwise.

### GetVirtio13Ok

`func (o *UpdateVMConfigRequest) GetVirtio13Ok() (*string, bool)`

GetVirtio13Ok returns a tuple with the Virtio13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio13

`func (o *UpdateVMConfigRequest) SetVirtio13(v string)`

SetVirtio13 sets Virtio13 field to given value.

### HasVirtio13

`func (o *UpdateVMConfigRequest) HasVirtio13() bool`

HasVirtio13 returns a boolean if a field has been set.

### GetVirtio14

`func (o *UpdateVMConfigRequest) GetVirtio14() string`

GetVirtio14 returns the Virtio14 field if non-nil, zero value otherwise.

### GetVirtio14Ok

`func (o *UpdateVMConfigRequest) GetVirtio14Ok() (*string, bool)`

GetVirtio14Ok returns a tuple with the Virtio14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio14

`func (o *UpdateVMConfigRequest) SetVirtio14(v string)`

SetVirtio14 sets Virtio14 field to given value.

### HasVirtio14

`func (o *UpdateVMConfigRequest) HasVirtio14() bool`

HasVirtio14 returns a boolean if a field has been set.

### GetVirtio15

`func (o *UpdateVMConfigRequest) GetVirtio15() string`

GetVirtio15 returns the Virtio15 field if non-nil, zero value otherwise.

### GetVirtio15Ok

`func (o *UpdateVMConfigRequest) GetVirtio15Ok() (*string, bool)`

GetVirtio15Ok returns a tuple with the Virtio15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtio15

`func (o *UpdateVMConfigRequest) SetVirtio15(v string)`

SetVirtio15 sets Virtio15 field to given value.

### HasVirtio15

`func (o *UpdateVMConfigRequest) HasVirtio15() bool`

HasVirtio15 returns a boolean if a field has been set.

### GetVmgenid

`func (o *UpdateVMConfigRequest) GetVmgenid() string`

GetVmgenid returns the Vmgenid field if non-nil, zero value otherwise.

### GetVmgenidOk

`func (o *UpdateVMConfigRequest) GetVmgenidOk() (*string, bool)`

GetVmgenidOk returns a tuple with the Vmgenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmgenid

`func (o *UpdateVMConfigRequest) SetVmgenid(v string)`

SetVmgenid sets Vmgenid field to given value.

### HasVmgenid

`func (o *UpdateVMConfigRequest) HasVmgenid() bool`

HasVmgenid returns a boolean if a field has been set.

### GetVmstatestorage

`func (o *UpdateVMConfigRequest) GetVmstatestorage() string`

GetVmstatestorage returns the Vmstatestorage field if non-nil, zero value otherwise.

### GetVmstatestorageOk

`func (o *UpdateVMConfigRequest) GetVmstatestorageOk() (*string, bool)`

GetVmstatestorageOk returns a tuple with the Vmstatestorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmstatestorage

`func (o *UpdateVMConfigRequest) SetVmstatestorage(v string)`

SetVmstatestorage sets Vmstatestorage field to given value.

### HasVmstatestorage

`func (o *UpdateVMConfigRequest) HasVmstatestorage() bool`

HasVmstatestorage returns a boolean if a field has been set.

### GetWatchdog

`func (o *UpdateVMConfigRequest) GetWatchdog() string`

GetWatchdog returns the Watchdog field if non-nil, zero value otherwise.

### GetWatchdogOk

`func (o *UpdateVMConfigRequest) GetWatchdogOk() (*string, bool)`

GetWatchdogOk returns a tuple with the Watchdog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWatchdog

`func (o *UpdateVMConfigRequest) SetWatchdog(v string)`

SetWatchdog sets Watchdog field to given value.

### HasWatchdog

`func (o *UpdateVMConfigRequest) HasWatchdog() bool`

HasWatchdog returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


