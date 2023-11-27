# GetContainerConfig200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | OS architecture type. | [optional] 
**Cmode** | Pointer to **string** | Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to &#39;console&#39; it tries to attach to /dev/console instead. If you set cmode to &#39;shell&#39;, it simply invokes a shell inside the container (no login). | [optional] 
**Console** | Pointer to **int32** | Attach a console device (/dev/console) to the container. | [optional] 
**Cores** | Pointer to **int64** | The number of cores assigned to the container. A container can use all available cores by default. | [optional] 
**Cpulimit** | Pointer to **float32** | Limit of CPU usage.  NOTE: If the computer has 2 CPUs, it has a total of &#39;2&#39; CPU time. Value &#39;0&#39; indicates no CPU limit. | [optional] 
**Cpuunits** | Pointer to **int64** | CPU weight for a container, will be clamped to [1, 10000] in cgroup v2. | [optional] 
**Debug** | Pointer to **int32** | Try to be more verbose. For now this only enables debug log-level on start. | [optional] 
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
**Digest** | Pointer to **string** | SHA1 digest of configuration file. This can be used to prevent concurrent modifications. | [optional] 
**Features** | Pointer to **string** | Allow containers access to advanced features. | [optional] 
**Hookscript** | Pointer to **string** | Script that will be exectued during various steps in the containers lifetime. | [optional] 
**Hostname** | Pointer to **string** | Set a host name for the container. | [optional] 
**Lock** | Pointer to **string** | Lock/unlock the container. | [optional] 
**Lxc** | Pointer to **[][]string** | Array of lxc low-level configurations ([[key1, value1], [key2, value2] ...]). | [optional] 
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
**Mp30** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp31** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp32** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp33** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp34** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp35** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp36** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp37** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp38** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp39** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp40** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp41** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp42** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp43** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp44** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp45** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp46** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp47** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp48** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp49** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp50** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp51** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp52** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp53** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp54** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp55** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp56** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp57** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp58** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp59** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp60** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp61** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp62** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp63** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp64** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp65** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp66** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp67** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp68** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp69** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp70** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp71** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp72** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp73** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp74** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp75** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp76** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp77** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp78** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp79** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp80** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp81** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp82** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp83** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp84** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp85** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp86** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp87** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp88** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp89** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp90** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp91** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp92** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp93** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp94** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp95** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp96** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp97** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp98** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp99** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp100** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp101** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp102** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp103** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp104** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp105** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp106** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp107** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp108** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp109** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp110** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp111** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp112** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp113** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp114** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp115** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp116** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp117** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp118** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp119** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp120** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp121** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp122** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp123** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp124** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp125** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp126** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp127** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp128** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp129** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp130** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp131** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp132** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp133** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp134** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp135** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp136** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp137** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp138** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp139** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp140** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp141** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp142** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp143** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp144** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp145** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp146** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp147** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp148** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp149** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp150** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp151** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp152** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp153** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp154** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp155** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp156** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp157** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp158** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp159** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp160** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp161** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp162** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp163** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp164** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp165** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp166** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp167** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp168** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp169** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp170** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp171** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp172** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp173** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp174** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp175** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp176** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp177** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp178** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp179** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp180** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp181** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp182** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp183** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp184** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp185** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp186** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp187** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp188** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp189** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp190** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp191** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp192** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp193** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp194** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp195** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp196** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp197** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp198** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp199** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp200** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp201** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp202** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp203** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp204** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp205** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp206** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp207** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp208** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp209** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp210** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp211** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp212** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp213** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp214** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp215** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp216** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp217** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp218** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp219** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp220** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp221** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp222** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp223** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp224** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp225** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp226** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp227** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp228** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp229** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp230** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp231** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp232** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp233** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp234** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp235** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp236** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp237** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp238** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp239** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp240** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp241** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp242** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp243** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp244** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp245** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp246** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp247** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp248** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp249** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp250** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp251** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp252** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp253** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp254** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
**Mp255** | Pointer to **string** | Use volume as container mount point. Use the special syntax STORAGE_ID:SIZE_IN_GiB to allocate a new volume. | [optional] 
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
**Onboot** | Pointer to **int32** | Specifies whether a container will be started during system bootup. | [optional] 
**Ostype** | Pointer to **string** | OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/&lt;ostype&gt;.common.conf. Value &#39;unmanaged&#39; can be used to skip and OS specific setup. | [optional] 
**Protection** | Pointer to **int32** | Sets the protection flag of the container. This will prevent the CT or CT&#39;s disk remove/update operation. | [optional] 
**Rootfs** | Pointer to **string** | Use volume as container root. | [optional] 
**Searchdomain** | Pointer to **string** | Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver. | [optional] 
**Startup** | Pointer to **string** | Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the &#39;up&#39; or &#39;down&#39; delay in seconds, which specifies a delay to wait before the next VM is started or stopped. | [optional] 
**Swap** | Pointer to **int64** | Amount of SWAP for the container in MB. | [optional] 
**Tags** | Pointer to **string** | Tags of the Container. This is only meta information. | [optional] 
**Template** | Pointer to **int32** | Enable/disable Template. | [optional] 
**Timezone** | Pointer to **string** | Time zone to use in the container. If option isn&#39;t set, then nothing will be done. Can be set to &#39;host&#39; to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab | [optional] 
**Tty** | Pointer to **int64** | Specify the number of tty available to the container | [optional] 
**Unprivileged** | Pointer to **int32** | Makes the container run as unprivileged user. (Should not be modified manually.) | [optional] 
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

### NewGetContainerConfig200ResponseData

`func NewGetContainerConfig200ResponseData() *GetContainerConfig200ResponseData`

NewGetContainerConfig200ResponseData instantiates a new GetContainerConfig200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContainerConfig200ResponseDataWithDefaults

`func NewGetContainerConfig200ResponseDataWithDefaults() *GetContainerConfig200ResponseData`

NewGetContainerConfig200ResponseDataWithDefaults instantiates a new GetContainerConfig200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *GetContainerConfig200ResponseData) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *GetContainerConfig200ResponseData) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *GetContainerConfig200ResponseData) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *GetContainerConfig200ResponseData) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetCmode

`func (o *GetContainerConfig200ResponseData) GetCmode() string`

GetCmode returns the Cmode field if non-nil, zero value otherwise.

### GetCmodeOk

`func (o *GetContainerConfig200ResponseData) GetCmodeOk() (*string, bool)`

GetCmodeOk returns a tuple with the Cmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmode

`func (o *GetContainerConfig200ResponseData) SetCmode(v string)`

SetCmode sets Cmode field to given value.

### HasCmode

`func (o *GetContainerConfig200ResponseData) HasCmode() bool`

HasCmode returns a boolean if a field has been set.

### GetConsole

`func (o *GetContainerConfig200ResponseData) GetConsole() int32`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *GetContainerConfig200ResponseData) GetConsoleOk() (*int32, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *GetContainerConfig200ResponseData) SetConsole(v int32)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *GetContainerConfig200ResponseData) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### GetCores

`func (o *GetContainerConfig200ResponseData) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *GetContainerConfig200ResponseData) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *GetContainerConfig200ResponseData) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *GetContainerConfig200ResponseData) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpulimit

`func (o *GetContainerConfig200ResponseData) GetCpulimit() float32`

GetCpulimit returns the Cpulimit field if non-nil, zero value otherwise.

### GetCpulimitOk

`func (o *GetContainerConfig200ResponseData) GetCpulimitOk() (*float32, bool)`

GetCpulimitOk returns a tuple with the Cpulimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpulimit

`func (o *GetContainerConfig200ResponseData) SetCpulimit(v float32)`

SetCpulimit sets Cpulimit field to given value.

### HasCpulimit

`func (o *GetContainerConfig200ResponseData) HasCpulimit() bool`

HasCpulimit returns a boolean if a field has been set.

### GetCpuunits

`func (o *GetContainerConfig200ResponseData) GetCpuunits() int64`

GetCpuunits returns the Cpuunits field if non-nil, zero value otherwise.

### GetCpuunitsOk

`func (o *GetContainerConfig200ResponseData) GetCpuunitsOk() (*int64, bool)`

GetCpuunitsOk returns a tuple with the Cpuunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuunits

`func (o *GetContainerConfig200ResponseData) SetCpuunits(v int64)`

SetCpuunits sets Cpuunits field to given value.

### HasCpuunits

`func (o *GetContainerConfig200ResponseData) HasCpuunits() bool`

HasCpuunits returns a boolean if a field has been set.

### GetDebug

`func (o *GetContainerConfig200ResponseData) GetDebug() int32`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *GetContainerConfig200ResponseData) GetDebugOk() (*int32, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *GetContainerConfig200ResponseData) SetDebug(v int32)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *GetContainerConfig200ResponseData) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDescription

`func (o *GetContainerConfig200ResponseData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetContainerConfig200ResponseData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetContainerConfig200ResponseData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetContainerConfig200ResponseData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDev0

`func (o *GetContainerConfig200ResponseData) GetDev0() string`

GetDev0 returns the Dev0 field if non-nil, zero value otherwise.

### GetDev0Ok

`func (o *GetContainerConfig200ResponseData) GetDev0Ok() (*string, bool)`

GetDev0Ok returns a tuple with the Dev0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev0

`func (o *GetContainerConfig200ResponseData) SetDev0(v string)`

SetDev0 sets Dev0 field to given value.

### HasDev0

`func (o *GetContainerConfig200ResponseData) HasDev0() bool`

HasDev0 returns a boolean if a field has been set.

### GetDev1

`func (o *GetContainerConfig200ResponseData) GetDev1() string`

GetDev1 returns the Dev1 field if non-nil, zero value otherwise.

### GetDev1Ok

`func (o *GetContainerConfig200ResponseData) GetDev1Ok() (*string, bool)`

GetDev1Ok returns a tuple with the Dev1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev1

`func (o *GetContainerConfig200ResponseData) SetDev1(v string)`

SetDev1 sets Dev1 field to given value.

### HasDev1

`func (o *GetContainerConfig200ResponseData) HasDev1() bool`

HasDev1 returns a boolean if a field has been set.

### GetDev2

`func (o *GetContainerConfig200ResponseData) GetDev2() string`

GetDev2 returns the Dev2 field if non-nil, zero value otherwise.

### GetDev2Ok

`func (o *GetContainerConfig200ResponseData) GetDev2Ok() (*string, bool)`

GetDev2Ok returns a tuple with the Dev2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev2

`func (o *GetContainerConfig200ResponseData) SetDev2(v string)`

SetDev2 sets Dev2 field to given value.

### HasDev2

`func (o *GetContainerConfig200ResponseData) HasDev2() bool`

HasDev2 returns a boolean if a field has been set.

### GetDev3

`func (o *GetContainerConfig200ResponseData) GetDev3() string`

GetDev3 returns the Dev3 field if non-nil, zero value otherwise.

### GetDev3Ok

`func (o *GetContainerConfig200ResponseData) GetDev3Ok() (*string, bool)`

GetDev3Ok returns a tuple with the Dev3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev3

`func (o *GetContainerConfig200ResponseData) SetDev3(v string)`

SetDev3 sets Dev3 field to given value.

### HasDev3

`func (o *GetContainerConfig200ResponseData) HasDev3() bool`

HasDev3 returns a boolean if a field has been set.

### GetDev4

`func (o *GetContainerConfig200ResponseData) GetDev4() string`

GetDev4 returns the Dev4 field if non-nil, zero value otherwise.

### GetDev4Ok

`func (o *GetContainerConfig200ResponseData) GetDev4Ok() (*string, bool)`

GetDev4Ok returns a tuple with the Dev4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev4

`func (o *GetContainerConfig200ResponseData) SetDev4(v string)`

SetDev4 sets Dev4 field to given value.

### HasDev4

`func (o *GetContainerConfig200ResponseData) HasDev4() bool`

HasDev4 returns a boolean if a field has been set.

### GetDev5

`func (o *GetContainerConfig200ResponseData) GetDev5() string`

GetDev5 returns the Dev5 field if non-nil, zero value otherwise.

### GetDev5Ok

`func (o *GetContainerConfig200ResponseData) GetDev5Ok() (*string, bool)`

GetDev5Ok returns a tuple with the Dev5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev5

`func (o *GetContainerConfig200ResponseData) SetDev5(v string)`

SetDev5 sets Dev5 field to given value.

### HasDev5

`func (o *GetContainerConfig200ResponseData) HasDev5() bool`

HasDev5 returns a boolean if a field has been set.

### GetDev6

`func (o *GetContainerConfig200ResponseData) GetDev6() string`

GetDev6 returns the Dev6 field if non-nil, zero value otherwise.

### GetDev6Ok

`func (o *GetContainerConfig200ResponseData) GetDev6Ok() (*string, bool)`

GetDev6Ok returns a tuple with the Dev6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev6

`func (o *GetContainerConfig200ResponseData) SetDev6(v string)`

SetDev6 sets Dev6 field to given value.

### HasDev6

`func (o *GetContainerConfig200ResponseData) HasDev6() bool`

HasDev6 returns a boolean if a field has been set.

### GetDev7

`func (o *GetContainerConfig200ResponseData) GetDev7() string`

GetDev7 returns the Dev7 field if non-nil, zero value otherwise.

### GetDev7Ok

`func (o *GetContainerConfig200ResponseData) GetDev7Ok() (*string, bool)`

GetDev7Ok returns a tuple with the Dev7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev7

`func (o *GetContainerConfig200ResponseData) SetDev7(v string)`

SetDev7 sets Dev7 field to given value.

### HasDev7

`func (o *GetContainerConfig200ResponseData) HasDev7() bool`

HasDev7 returns a boolean if a field has been set.

### GetDev8

`func (o *GetContainerConfig200ResponseData) GetDev8() string`

GetDev8 returns the Dev8 field if non-nil, zero value otherwise.

### GetDev8Ok

`func (o *GetContainerConfig200ResponseData) GetDev8Ok() (*string, bool)`

GetDev8Ok returns a tuple with the Dev8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev8

`func (o *GetContainerConfig200ResponseData) SetDev8(v string)`

SetDev8 sets Dev8 field to given value.

### HasDev8

`func (o *GetContainerConfig200ResponseData) HasDev8() bool`

HasDev8 returns a boolean if a field has been set.

### GetDev9

`func (o *GetContainerConfig200ResponseData) GetDev9() string`

GetDev9 returns the Dev9 field if non-nil, zero value otherwise.

### GetDev9Ok

`func (o *GetContainerConfig200ResponseData) GetDev9Ok() (*string, bool)`

GetDev9Ok returns a tuple with the Dev9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev9

`func (o *GetContainerConfig200ResponseData) SetDev9(v string)`

SetDev9 sets Dev9 field to given value.

### HasDev9

`func (o *GetContainerConfig200ResponseData) HasDev9() bool`

HasDev9 returns a boolean if a field has been set.

### GetDev10

`func (o *GetContainerConfig200ResponseData) GetDev10() string`

GetDev10 returns the Dev10 field if non-nil, zero value otherwise.

### GetDev10Ok

`func (o *GetContainerConfig200ResponseData) GetDev10Ok() (*string, bool)`

GetDev10Ok returns a tuple with the Dev10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev10

`func (o *GetContainerConfig200ResponseData) SetDev10(v string)`

SetDev10 sets Dev10 field to given value.

### HasDev10

`func (o *GetContainerConfig200ResponseData) HasDev10() bool`

HasDev10 returns a boolean if a field has been set.

### GetDev11

`func (o *GetContainerConfig200ResponseData) GetDev11() string`

GetDev11 returns the Dev11 field if non-nil, zero value otherwise.

### GetDev11Ok

`func (o *GetContainerConfig200ResponseData) GetDev11Ok() (*string, bool)`

GetDev11Ok returns a tuple with the Dev11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev11

`func (o *GetContainerConfig200ResponseData) SetDev11(v string)`

SetDev11 sets Dev11 field to given value.

### HasDev11

`func (o *GetContainerConfig200ResponseData) HasDev11() bool`

HasDev11 returns a boolean if a field has been set.

### GetDev12

`func (o *GetContainerConfig200ResponseData) GetDev12() string`

GetDev12 returns the Dev12 field if non-nil, zero value otherwise.

### GetDev12Ok

`func (o *GetContainerConfig200ResponseData) GetDev12Ok() (*string, bool)`

GetDev12Ok returns a tuple with the Dev12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev12

`func (o *GetContainerConfig200ResponseData) SetDev12(v string)`

SetDev12 sets Dev12 field to given value.

### HasDev12

`func (o *GetContainerConfig200ResponseData) HasDev12() bool`

HasDev12 returns a boolean if a field has been set.

### GetDev13

`func (o *GetContainerConfig200ResponseData) GetDev13() string`

GetDev13 returns the Dev13 field if non-nil, zero value otherwise.

### GetDev13Ok

`func (o *GetContainerConfig200ResponseData) GetDev13Ok() (*string, bool)`

GetDev13Ok returns a tuple with the Dev13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev13

`func (o *GetContainerConfig200ResponseData) SetDev13(v string)`

SetDev13 sets Dev13 field to given value.

### HasDev13

`func (o *GetContainerConfig200ResponseData) HasDev13() bool`

HasDev13 returns a boolean if a field has been set.

### GetDev14

`func (o *GetContainerConfig200ResponseData) GetDev14() string`

GetDev14 returns the Dev14 field if non-nil, zero value otherwise.

### GetDev14Ok

`func (o *GetContainerConfig200ResponseData) GetDev14Ok() (*string, bool)`

GetDev14Ok returns a tuple with the Dev14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev14

`func (o *GetContainerConfig200ResponseData) SetDev14(v string)`

SetDev14 sets Dev14 field to given value.

### HasDev14

`func (o *GetContainerConfig200ResponseData) HasDev14() bool`

HasDev14 returns a boolean if a field has been set.

### GetDev15

`func (o *GetContainerConfig200ResponseData) GetDev15() string`

GetDev15 returns the Dev15 field if non-nil, zero value otherwise.

### GetDev15Ok

`func (o *GetContainerConfig200ResponseData) GetDev15Ok() (*string, bool)`

GetDev15Ok returns a tuple with the Dev15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev15

`func (o *GetContainerConfig200ResponseData) SetDev15(v string)`

SetDev15 sets Dev15 field to given value.

### HasDev15

`func (o *GetContainerConfig200ResponseData) HasDev15() bool`

HasDev15 returns a boolean if a field has been set.

### GetDev16

`func (o *GetContainerConfig200ResponseData) GetDev16() string`

GetDev16 returns the Dev16 field if non-nil, zero value otherwise.

### GetDev16Ok

`func (o *GetContainerConfig200ResponseData) GetDev16Ok() (*string, bool)`

GetDev16Ok returns a tuple with the Dev16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev16

`func (o *GetContainerConfig200ResponseData) SetDev16(v string)`

SetDev16 sets Dev16 field to given value.

### HasDev16

`func (o *GetContainerConfig200ResponseData) HasDev16() bool`

HasDev16 returns a boolean if a field has been set.

### GetDev17

`func (o *GetContainerConfig200ResponseData) GetDev17() string`

GetDev17 returns the Dev17 field if non-nil, zero value otherwise.

### GetDev17Ok

`func (o *GetContainerConfig200ResponseData) GetDev17Ok() (*string, bool)`

GetDev17Ok returns a tuple with the Dev17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev17

`func (o *GetContainerConfig200ResponseData) SetDev17(v string)`

SetDev17 sets Dev17 field to given value.

### HasDev17

`func (o *GetContainerConfig200ResponseData) HasDev17() bool`

HasDev17 returns a boolean if a field has been set.

### GetDev18

`func (o *GetContainerConfig200ResponseData) GetDev18() string`

GetDev18 returns the Dev18 field if non-nil, zero value otherwise.

### GetDev18Ok

`func (o *GetContainerConfig200ResponseData) GetDev18Ok() (*string, bool)`

GetDev18Ok returns a tuple with the Dev18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev18

`func (o *GetContainerConfig200ResponseData) SetDev18(v string)`

SetDev18 sets Dev18 field to given value.

### HasDev18

`func (o *GetContainerConfig200ResponseData) HasDev18() bool`

HasDev18 returns a boolean if a field has been set.

### GetDev19

`func (o *GetContainerConfig200ResponseData) GetDev19() string`

GetDev19 returns the Dev19 field if non-nil, zero value otherwise.

### GetDev19Ok

`func (o *GetContainerConfig200ResponseData) GetDev19Ok() (*string, bool)`

GetDev19Ok returns a tuple with the Dev19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev19

`func (o *GetContainerConfig200ResponseData) SetDev19(v string)`

SetDev19 sets Dev19 field to given value.

### HasDev19

`func (o *GetContainerConfig200ResponseData) HasDev19() bool`

HasDev19 returns a boolean if a field has been set.

### GetDev20

`func (o *GetContainerConfig200ResponseData) GetDev20() string`

GetDev20 returns the Dev20 field if non-nil, zero value otherwise.

### GetDev20Ok

`func (o *GetContainerConfig200ResponseData) GetDev20Ok() (*string, bool)`

GetDev20Ok returns a tuple with the Dev20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev20

`func (o *GetContainerConfig200ResponseData) SetDev20(v string)`

SetDev20 sets Dev20 field to given value.

### HasDev20

`func (o *GetContainerConfig200ResponseData) HasDev20() bool`

HasDev20 returns a boolean if a field has been set.

### GetDev21

`func (o *GetContainerConfig200ResponseData) GetDev21() string`

GetDev21 returns the Dev21 field if non-nil, zero value otherwise.

### GetDev21Ok

`func (o *GetContainerConfig200ResponseData) GetDev21Ok() (*string, bool)`

GetDev21Ok returns a tuple with the Dev21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev21

`func (o *GetContainerConfig200ResponseData) SetDev21(v string)`

SetDev21 sets Dev21 field to given value.

### HasDev21

`func (o *GetContainerConfig200ResponseData) HasDev21() bool`

HasDev21 returns a boolean if a field has been set.

### GetDev22

`func (o *GetContainerConfig200ResponseData) GetDev22() string`

GetDev22 returns the Dev22 field if non-nil, zero value otherwise.

### GetDev22Ok

`func (o *GetContainerConfig200ResponseData) GetDev22Ok() (*string, bool)`

GetDev22Ok returns a tuple with the Dev22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev22

`func (o *GetContainerConfig200ResponseData) SetDev22(v string)`

SetDev22 sets Dev22 field to given value.

### HasDev22

`func (o *GetContainerConfig200ResponseData) HasDev22() bool`

HasDev22 returns a boolean if a field has been set.

### GetDev23

`func (o *GetContainerConfig200ResponseData) GetDev23() string`

GetDev23 returns the Dev23 field if non-nil, zero value otherwise.

### GetDev23Ok

`func (o *GetContainerConfig200ResponseData) GetDev23Ok() (*string, bool)`

GetDev23Ok returns a tuple with the Dev23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev23

`func (o *GetContainerConfig200ResponseData) SetDev23(v string)`

SetDev23 sets Dev23 field to given value.

### HasDev23

`func (o *GetContainerConfig200ResponseData) HasDev23() bool`

HasDev23 returns a boolean if a field has been set.

### GetDev24

`func (o *GetContainerConfig200ResponseData) GetDev24() string`

GetDev24 returns the Dev24 field if non-nil, zero value otherwise.

### GetDev24Ok

`func (o *GetContainerConfig200ResponseData) GetDev24Ok() (*string, bool)`

GetDev24Ok returns a tuple with the Dev24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev24

`func (o *GetContainerConfig200ResponseData) SetDev24(v string)`

SetDev24 sets Dev24 field to given value.

### HasDev24

`func (o *GetContainerConfig200ResponseData) HasDev24() bool`

HasDev24 returns a boolean if a field has been set.

### GetDev25

`func (o *GetContainerConfig200ResponseData) GetDev25() string`

GetDev25 returns the Dev25 field if non-nil, zero value otherwise.

### GetDev25Ok

`func (o *GetContainerConfig200ResponseData) GetDev25Ok() (*string, bool)`

GetDev25Ok returns a tuple with the Dev25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev25

`func (o *GetContainerConfig200ResponseData) SetDev25(v string)`

SetDev25 sets Dev25 field to given value.

### HasDev25

`func (o *GetContainerConfig200ResponseData) HasDev25() bool`

HasDev25 returns a boolean if a field has been set.

### GetDev26

`func (o *GetContainerConfig200ResponseData) GetDev26() string`

GetDev26 returns the Dev26 field if non-nil, zero value otherwise.

### GetDev26Ok

`func (o *GetContainerConfig200ResponseData) GetDev26Ok() (*string, bool)`

GetDev26Ok returns a tuple with the Dev26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev26

`func (o *GetContainerConfig200ResponseData) SetDev26(v string)`

SetDev26 sets Dev26 field to given value.

### HasDev26

`func (o *GetContainerConfig200ResponseData) HasDev26() bool`

HasDev26 returns a boolean if a field has been set.

### GetDev27

`func (o *GetContainerConfig200ResponseData) GetDev27() string`

GetDev27 returns the Dev27 field if non-nil, zero value otherwise.

### GetDev27Ok

`func (o *GetContainerConfig200ResponseData) GetDev27Ok() (*string, bool)`

GetDev27Ok returns a tuple with the Dev27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev27

`func (o *GetContainerConfig200ResponseData) SetDev27(v string)`

SetDev27 sets Dev27 field to given value.

### HasDev27

`func (o *GetContainerConfig200ResponseData) HasDev27() bool`

HasDev27 returns a boolean if a field has been set.

### GetDev28

`func (o *GetContainerConfig200ResponseData) GetDev28() string`

GetDev28 returns the Dev28 field if non-nil, zero value otherwise.

### GetDev28Ok

`func (o *GetContainerConfig200ResponseData) GetDev28Ok() (*string, bool)`

GetDev28Ok returns a tuple with the Dev28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev28

`func (o *GetContainerConfig200ResponseData) SetDev28(v string)`

SetDev28 sets Dev28 field to given value.

### HasDev28

`func (o *GetContainerConfig200ResponseData) HasDev28() bool`

HasDev28 returns a boolean if a field has been set.

### GetDev29

`func (o *GetContainerConfig200ResponseData) GetDev29() string`

GetDev29 returns the Dev29 field if non-nil, zero value otherwise.

### GetDev29Ok

`func (o *GetContainerConfig200ResponseData) GetDev29Ok() (*string, bool)`

GetDev29Ok returns a tuple with the Dev29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDev29

`func (o *GetContainerConfig200ResponseData) SetDev29(v string)`

SetDev29 sets Dev29 field to given value.

### HasDev29

`func (o *GetContainerConfig200ResponseData) HasDev29() bool`

HasDev29 returns a boolean if a field has been set.

### GetDigest

`func (o *GetContainerConfig200ResponseData) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *GetContainerConfig200ResponseData) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *GetContainerConfig200ResponseData) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *GetContainerConfig200ResponseData) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetFeatures

`func (o *GetContainerConfig200ResponseData) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *GetContainerConfig200ResponseData) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *GetContainerConfig200ResponseData) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *GetContainerConfig200ResponseData) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetHookscript

`func (o *GetContainerConfig200ResponseData) GetHookscript() string`

GetHookscript returns the Hookscript field if non-nil, zero value otherwise.

### GetHookscriptOk

`func (o *GetContainerConfig200ResponseData) GetHookscriptOk() (*string, bool)`

GetHookscriptOk returns a tuple with the Hookscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookscript

`func (o *GetContainerConfig200ResponseData) SetHookscript(v string)`

SetHookscript sets Hookscript field to given value.

### HasHookscript

`func (o *GetContainerConfig200ResponseData) HasHookscript() bool`

HasHookscript returns a boolean if a field has been set.

### GetHostname

`func (o *GetContainerConfig200ResponseData) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetContainerConfig200ResponseData) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetContainerConfig200ResponseData) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetContainerConfig200ResponseData) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLock

`func (o *GetContainerConfig200ResponseData) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *GetContainerConfig200ResponseData) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *GetContainerConfig200ResponseData) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *GetContainerConfig200ResponseData) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetLxc

`func (o *GetContainerConfig200ResponseData) GetLxc() [][]string`

GetLxc returns the Lxc field if non-nil, zero value otherwise.

### GetLxcOk

`func (o *GetContainerConfig200ResponseData) GetLxcOk() (*[][]string, bool)`

GetLxcOk returns a tuple with the Lxc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLxc

`func (o *GetContainerConfig200ResponseData) SetLxc(v [][]string)`

SetLxc sets Lxc field to given value.

### HasLxc

`func (o *GetContainerConfig200ResponseData) HasLxc() bool`

HasLxc returns a boolean if a field has been set.

### GetMemory

`func (o *GetContainerConfig200ResponseData) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GetContainerConfig200ResponseData) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GetContainerConfig200ResponseData) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GetContainerConfig200ResponseData) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMp0

`func (o *GetContainerConfig200ResponseData) GetMp0() string`

GetMp0 returns the Mp0 field if non-nil, zero value otherwise.

### GetMp0Ok

`func (o *GetContainerConfig200ResponseData) GetMp0Ok() (*string, bool)`

GetMp0Ok returns a tuple with the Mp0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp0

`func (o *GetContainerConfig200ResponseData) SetMp0(v string)`

SetMp0 sets Mp0 field to given value.

### HasMp0

`func (o *GetContainerConfig200ResponseData) HasMp0() bool`

HasMp0 returns a boolean if a field has been set.

### GetMp1

`func (o *GetContainerConfig200ResponseData) GetMp1() string`

GetMp1 returns the Mp1 field if non-nil, zero value otherwise.

### GetMp1Ok

`func (o *GetContainerConfig200ResponseData) GetMp1Ok() (*string, bool)`

GetMp1Ok returns a tuple with the Mp1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp1

`func (o *GetContainerConfig200ResponseData) SetMp1(v string)`

SetMp1 sets Mp1 field to given value.

### HasMp1

`func (o *GetContainerConfig200ResponseData) HasMp1() bool`

HasMp1 returns a boolean if a field has been set.

### GetMp2

`func (o *GetContainerConfig200ResponseData) GetMp2() string`

GetMp2 returns the Mp2 field if non-nil, zero value otherwise.

### GetMp2Ok

`func (o *GetContainerConfig200ResponseData) GetMp2Ok() (*string, bool)`

GetMp2Ok returns a tuple with the Mp2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp2

`func (o *GetContainerConfig200ResponseData) SetMp2(v string)`

SetMp2 sets Mp2 field to given value.

### HasMp2

`func (o *GetContainerConfig200ResponseData) HasMp2() bool`

HasMp2 returns a boolean if a field has been set.

### GetMp3

`func (o *GetContainerConfig200ResponseData) GetMp3() string`

GetMp3 returns the Mp3 field if non-nil, zero value otherwise.

### GetMp3Ok

`func (o *GetContainerConfig200ResponseData) GetMp3Ok() (*string, bool)`

GetMp3Ok returns a tuple with the Mp3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp3

`func (o *GetContainerConfig200ResponseData) SetMp3(v string)`

SetMp3 sets Mp3 field to given value.

### HasMp3

`func (o *GetContainerConfig200ResponseData) HasMp3() bool`

HasMp3 returns a boolean if a field has been set.

### GetMp4

`func (o *GetContainerConfig200ResponseData) GetMp4() string`

GetMp4 returns the Mp4 field if non-nil, zero value otherwise.

### GetMp4Ok

`func (o *GetContainerConfig200ResponseData) GetMp4Ok() (*string, bool)`

GetMp4Ok returns a tuple with the Mp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4

`func (o *GetContainerConfig200ResponseData) SetMp4(v string)`

SetMp4 sets Mp4 field to given value.

### HasMp4

`func (o *GetContainerConfig200ResponseData) HasMp4() bool`

HasMp4 returns a boolean if a field has been set.

### GetMp5

`func (o *GetContainerConfig200ResponseData) GetMp5() string`

GetMp5 returns the Mp5 field if non-nil, zero value otherwise.

### GetMp5Ok

`func (o *GetContainerConfig200ResponseData) GetMp5Ok() (*string, bool)`

GetMp5Ok returns a tuple with the Mp5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp5

`func (o *GetContainerConfig200ResponseData) SetMp5(v string)`

SetMp5 sets Mp5 field to given value.

### HasMp5

`func (o *GetContainerConfig200ResponseData) HasMp5() bool`

HasMp5 returns a boolean if a field has been set.

### GetMp6

`func (o *GetContainerConfig200ResponseData) GetMp6() string`

GetMp6 returns the Mp6 field if non-nil, zero value otherwise.

### GetMp6Ok

`func (o *GetContainerConfig200ResponseData) GetMp6Ok() (*string, bool)`

GetMp6Ok returns a tuple with the Mp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp6

`func (o *GetContainerConfig200ResponseData) SetMp6(v string)`

SetMp6 sets Mp6 field to given value.

### HasMp6

`func (o *GetContainerConfig200ResponseData) HasMp6() bool`

HasMp6 returns a boolean if a field has been set.

### GetMp7

`func (o *GetContainerConfig200ResponseData) GetMp7() string`

GetMp7 returns the Mp7 field if non-nil, zero value otherwise.

### GetMp7Ok

`func (o *GetContainerConfig200ResponseData) GetMp7Ok() (*string, bool)`

GetMp7Ok returns a tuple with the Mp7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp7

`func (o *GetContainerConfig200ResponseData) SetMp7(v string)`

SetMp7 sets Mp7 field to given value.

### HasMp7

`func (o *GetContainerConfig200ResponseData) HasMp7() bool`

HasMp7 returns a boolean if a field has been set.

### GetMp8

`func (o *GetContainerConfig200ResponseData) GetMp8() string`

GetMp8 returns the Mp8 field if non-nil, zero value otherwise.

### GetMp8Ok

`func (o *GetContainerConfig200ResponseData) GetMp8Ok() (*string, bool)`

GetMp8Ok returns a tuple with the Mp8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp8

`func (o *GetContainerConfig200ResponseData) SetMp8(v string)`

SetMp8 sets Mp8 field to given value.

### HasMp8

`func (o *GetContainerConfig200ResponseData) HasMp8() bool`

HasMp8 returns a boolean if a field has been set.

### GetMp9

`func (o *GetContainerConfig200ResponseData) GetMp9() string`

GetMp9 returns the Mp9 field if non-nil, zero value otherwise.

### GetMp9Ok

`func (o *GetContainerConfig200ResponseData) GetMp9Ok() (*string, bool)`

GetMp9Ok returns a tuple with the Mp9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp9

`func (o *GetContainerConfig200ResponseData) SetMp9(v string)`

SetMp9 sets Mp9 field to given value.

### HasMp9

`func (o *GetContainerConfig200ResponseData) HasMp9() bool`

HasMp9 returns a boolean if a field has been set.

### GetMp10

`func (o *GetContainerConfig200ResponseData) GetMp10() string`

GetMp10 returns the Mp10 field if non-nil, zero value otherwise.

### GetMp10Ok

`func (o *GetContainerConfig200ResponseData) GetMp10Ok() (*string, bool)`

GetMp10Ok returns a tuple with the Mp10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp10

`func (o *GetContainerConfig200ResponseData) SetMp10(v string)`

SetMp10 sets Mp10 field to given value.

### HasMp10

`func (o *GetContainerConfig200ResponseData) HasMp10() bool`

HasMp10 returns a boolean if a field has been set.

### GetMp11

`func (o *GetContainerConfig200ResponseData) GetMp11() string`

GetMp11 returns the Mp11 field if non-nil, zero value otherwise.

### GetMp11Ok

`func (o *GetContainerConfig200ResponseData) GetMp11Ok() (*string, bool)`

GetMp11Ok returns a tuple with the Mp11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp11

`func (o *GetContainerConfig200ResponseData) SetMp11(v string)`

SetMp11 sets Mp11 field to given value.

### HasMp11

`func (o *GetContainerConfig200ResponseData) HasMp11() bool`

HasMp11 returns a boolean if a field has been set.

### GetMp12

`func (o *GetContainerConfig200ResponseData) GetMp12() string`

GetMp12 returns the Mp12 field if non-nil, zero value otherwise.

### GetMp12Ok

`func (o *GetContainerConfig200ResponseData) GetMp12Ok() (*string, bool)`

GetMp12Ok returns a tuple with the Mp12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp12

`func (o *GetContainerConfig200ResponseData) SetMp12(v string)`

SetMp12 sets Mp12 field to given value.

### HasMp12

`func (o *GetContainerConfig200ResponseData) HasMp12() bool`

HasMp12 returns a boolean if a field has been set.

### GetMp13

`func (o *GetContainerConfig200ResponseData) GetMp13() string`

GetMp13 returns the Mp13 field if non-nil, zero value otherwise.

### GetMp13Ok

`func (o *GetContainerConfig200ResponseData) GetMp13Ok() (*string, bool)`

GetMp13Ok returns a tuple with the Mp13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp13

`func (o *GetContainerConfig200ResponseData) SetMp13(v string)`

SetMp13 sets Mp13 field to given value.

### HasMp13

`func (o *GetContainerConfig200ResponseData) HasMp13() bool`

HasMp13 returns a boolean if a field has been set.

### GetMp14

`func (o *GetContainerConfig200ResponseData) GetMp14() string`

GetMp14 returns the Mp14 field if non-nil, zero value otherwise.

### GetMp14Ok

`func (o *GetContainerConfig200ResponseData) GetMp14Ok() (*string, bool)`

GetMp14Ok returns a tuple with the Mp14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp14

`func (o *GetContainerConfig200ResponseData) SetMp14(v string)`

SetMp14 sets Mp14 field to given value.

### HasMp14

`func (o *GetContainerConfig200ResponseData) HasMp14() bool`

HasMp14 returns a boolean if a field has been set.

### GetMp15

`func (o *GetContainerConfig200ResponseData) GetMp15() string`

GetMp15 returns the Mp15 field if non-nil, zero value otherwise.

### GetMp15Ok

`func (o *GetContainerConfig200ResponseData) GetMp15Ok() (*string, bool)`

GetMp15Ok returns a tuple with the Mp15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp15

`func (o *GetContainerConfig200ResponseData) SetMp15(v string)`

SetMp15 sets Mp15 field to given value.

### HasMp15

`func (o *GetContainerConfig200ResponseData) HasMp15() bool`

HasMp15 returns a boolean if a field has been set.

### GetMp16

`func (o *GetContainerConfig200ResponseData) GetMp16() string`

GetMp16 returns the Mp16 field if non-nil, zero value otherwise.

### GetMp16Ok

`func (o *GetContainerConfig200ResponseData) GetMp16Ok() (*string, bool)`

GetMp16Ok returns a tuple with the Mp16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp16

`func (o *GetContainerConfig200ResponseData) SetMp16(v string)`

SetMp16 sets Mp16 field to given value.

### HasMp16

`func (o *GetContainerConfig200ResponseData) HasMp16() bool`

HasMp16 returns a boolean if a field has been set.

### GetMp17

`func (o *GetContainerConfig200ResponseData) GetMp17() string`

GetMp17 returns the Mp17 field if non-nil, zero value otherwise.

### GetMp17Ok

`func (o *GetContainerConfig200ResponseData) GetMp17Ok() (*string, bool)`

GetMp17Ok returns a tuple with the Mp17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp17

`func (o *GetContainerConfig200ResponseData) SetMp17(v string)`

SetMp17 sets Mp17 field to given value.

### HasMp17

`func (o *GetContainerConfig200ResponseData) HasMp17() bool`

HasMp17 returns a boolean if a field has been set.

### GetMp18

`func (o *GetContainerConfig200ResponseData) GetMp18() string`

GetMp18 returns the Mp18 field if non-nil, zero value otherwise.

### GetMp18Ok

`func (o *GetContainerConfig200ResponseData) GetMp18Ok() (*string, bool)`

GetMp18Ok returns a tuple with the Mp18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp18

`func (o *GetContainerConfig200ResponseData) SetMp18(v string)`

SetMp18 sets Mp18 field to given value.

### HasMp18

`func (o *GetContainerConfig200ResponseData) HasMp18() bool`

HasMp18 returns a boolean if a field has been set.

### GetMp19

`func (o *GetContainerConfig200ResponseData) GetMp19() string`

GetMp19 returns the Mp19 field if non-nil, zero value otherwise.

### GetMp19Ok

`func (o *GetContainerConfig200ResponseData) GetMp19Ok() (*string, bool)`

GetMp19Ok returns a tuple with the Mp19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp19

`func (o *GetContainerConfig200ResponseData) SetMp19(v string)`

SetMp19 sets Mp19 field to given value.

### HasMp19

`func (o *GetContainerConfig200ResponseData) HasMp19() bool`

HasMp19 returns a boolean if a field has been set.

### GetMp20

`func (o *GetContainerConfig200ResponseData) GetMp20() string`

GetMp20 returns the Mp20 field if non-nil, zero value otherwise.

### GetMp20Ok

`func (o *GetContainerConfig200ResponseData) GetMp20Ok() (*string, bool)`

GetMp20Ok returns a tuple with the Mp20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp20

`func (o *GetContainerConfig200ResponseData) SetMp20(v string)`

SetMp20 sets Mp20 field to given value.

### HasMp20

`func (o *GetContainerConfig200ResponseData) HasMp20() bool`

HasMp20 returns a boolean if a field has been set.

### GetMp21

`func (o *GetContainerConfig200ResponseData) GetMp21() string`

GetMp21 returns the Mp21 field if non-nil, zero value otherwise.

### GetMp21Ok

`func (o *GetContainerConfig200ResponseData) GetMp21Ok() (*string, bool)`

GetMp21Ok returns a tuple with the Mp21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp21

`func (o *GetContainerConfig200ResponseData) SetMp21(v string)`

SetMp21 sets Mp21 field to given value.

### HasMp21

`func (o *GetContainerConfig200ResponseData) HasMp21() bool`

HasMp21 returns a boolean if a field has been set.

### GetMp22

`func (o *GetContainerConfig200ResponseData) GetMp22() string`

GetMp22 returns the Mp22 field if non-nil, zero value otherwise.

### GetMp22Ok

`func (o *GetContainerConfig200ResponseData) GetMp22Ok() (*string, bool)`

GetMp22Ok returns a tuple with the Mp22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp22

`func (o *GetContainerConfig200ResponseData) SetMp22(v string)`

SetMp22 sets Mp22 field to given value.

### HasMp22

`func (o *GetContainerConfig200ResponseData) HasMp22() bool`

HasMp22 returns a boolean if a field has been set.

### GetMp23

`func (o *GetContainerConfig200ResponseData) GetMp23() string`

GetMp23 returns the Mp23 field if non-nil, zero value otherwise.

### GetMp23Ok

`func (o *GetContainerConfig200ResponseData) GetMp23Ok() (*string, bool)`

GetMp23Ok returns a tuple with the Mp23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp23

`func (o *GetContainerConfig200ResponseData) SetMp23(v string)`

SetMp23 sets Mp23 field to given value.

### HasMp23

`func (o *GetContainerConfig200ResponseData) HasMp23() bool`

HasMp23 returns a boolean if a field has been set.

### GetMp24

`func (o *GetContainerConfig200ResponseData) GetMp24() string`

GetMp24 returns the Mp24 field if non-nil, zero value otherwise.

### GetMp24Ok

`func (o *GetContainerConfig200ResponseData) GetMp24Ok() (*string, bool)`

GetMp24Ok returns a tuple with the Mp24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp24

`func (o *GetContainerConfig200ResponseData) SetMp24(v string)`

SetMp24 sets Mp24 field to given value.

### HasMp24

`func (o *GetContainerConfig200ResponseData) HasMp24() bool`

HasMp24 returns a boolean if a field has been set.

### GetMp25

`func (o *GetContainerConfig200ResponseData) GetMp25() string`

GetMp25 returns the Mp25 field if non-nil, zero value otherwise.

### GetMp25Ok

`func (o *GetContainerConfig200ResponseData) GetMp25Ok() (*string, bool)`

GetMp25Ok returns a tuple with the Mp25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp25

`func (o *GetContainerConfig200ResponseData) SetMp25(v string)`

SetMp25 sets Mp25 field to given value.

### HasMp25

`func (o *GetContainerConfig200ResponseData) HasMp25() bool`

HasMp25 returns a boolean if a field has been set.

### GetMp26

`func (o *GetContainerConfig200ResponseData) GetMp26() string`

GetMp26 returns the Mp26 field if non-nil, zero value otherwise.

### GetMp26Ok

`func (o *GetContainerConfig200ResponseData) GetMp26Ok() (*string, bool)`

GetMp26Ok returns a tuple with the Mp26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp26

`func (o *GetContainerConfig200ResponseData) SetMp26(v string)`

SetMp26 sets Mp26 field to given value.

### HasMp26

`func (o *GetContainerConfig200ResponseData) HasMp26() bool`

HasMp26 returns a boolean if a field has been set.

### GetMp27

`func (o *GetContainerConfig200ResponseData) GetMp27() string`

GetMp27 returns the Mp27 field if non-nil, zero value otherwise.

### GetMp27Ok

`func (o *GetContainerConfig200ResponseData) GetMp27Ok() (*string, bool)`

GetMp27Ok returns a tuple with the Mp27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp27

`func (o *GetContainerConfig200ResponseData) SetMp27(v string)`

SetMp27 sets Mp27 field to given value.

### HasMp27

`func (o *GetContainerConfig200ResponseData) HasMp27() bool`

HasMp27 returns a boolean if a field has been set.

### GetMp28

`func (o *GetContainerConfig200ResponseData) GetMp28() string`

GetMp28 returns the Mp28 field if non-nil, zero value otherwise.

### GetMp28Ok

`func (o *GetContainerConfig200ResponseData) GetMp28Ok() (*string, bool)`

GetMp28Ok returns a tuple with the Mp28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp28

`func (o *GetContainerConfig200ResponseData) SetMp28(v string)`

SetMp28 sets Mp28 field to given value.

### HasMp28

`func (o *GetContainerConfig200ResponseData) HasMp28() bool`

HasMp28 returns a boolean if a field has been set.

### GetMp29

`func (o *GetContainerConfig200ResponseData) GetMp29() string`

GetMp29 returns the Mp29 field if non-nil, zero value otherwise.

### GetMp29Ok

`func (o *GetContainerConfig200ResponseData) GetMp29Ok() (*string, bool)`

GetMp29Ok returns a tuple with the Mp29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp29

`func (o *GetContainerConfig200ResponseData) SetMp29(v string)`

SetMp29 sets Mp29 field to given value.

### HasMp29

`func (o *GetContainerConfig200ResponseData) HasMp29() bool`

HasMp29 returns a boolean if a field has been set.

### GetMp30

`func (o *GetContainerConfig200ResponseData) GetMp30() string`

GetMp30 returns the Mp30 field if non-nil, zero value otherwise.

### GetMp30Ok

`func (o *GetContainerConfig200ResponseData) GetMp30Ok() (*string, bool)`

GetMp30Ok returns a tuple with the Mp30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp30

`func (o *GetContainerConfig200ResponseData) SetMp30(v string)`

SetMp30 sets Mp30 field to given value.

### HasMp30

`func (o *GetContainerConfig200ResponseData) HasMp30() bool`

HasMp30 returns a boolean if a field has been set.

### GetMp31

`func (o *GetContainerConfig200ResponseData) GetMp31() string`

GetMp31 returns the Mp31 field if non-nil, zero value otherwise.

### GetMp31Ok

`func (o *GetContainerConfig200ResponseData) GetMp31Ok() (*string, bool)`

GetMp31Ok returns a tuple with the Mp31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp31

`func (o *GetContainerConfig200ResponseData) SetMp31(v string)`

SetMp31 sets Mp31 field to given value.

### HasMp31

`func (o *GetContainerConfig200ResponseData) HasMp31() bool`

HasMp31 returns a boolean if a field has been set.

### GetMp32

`func (o *GetContainerConfig200ResponseData) GetMp32() string`

GetMp32 returns the Mp32 field if non-nil, zero value otherwise.

### GetMp32Ok

`func (o *GetContainerConfig200ResponseData) GetMp32Ok() (*string, bool)`

GetMp32Ok returns a tuple with the Mp32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp32

`func (o *GetContainerConfig200ResponseData) SetMp32(v string)`

SetMp32 sets Mp32 field to given value.

### HasMp32

`func (o *GetContainerConfig200ResponseData) HasMp32() bool`

HasMp32 returns a boolean if a field has been set.

### GetMp33

`func (o *GetContainerConfig200ResponseData) GetMp33() string`

GetMp33 returns the Mp33 field if non-nil, zero value otherwise.

### GetMp33Ok

`func (o *GetContainerConfig200ResponseData) GetMp33Ok() (*string, bool)`

GetMp33Ok returns a tuple with the Mp33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp33

`func (o *GetContainerConfig200ResponseData) SetMp33(v string)`

SetMp33 sets Mp33 field to given value.

### HasMp33

`func (o *GetContainerConfig200ResponseData) HasMp33() bool`

HasMp33 returns a boolean if a field has been set.

### GetMp34

`func (o *GetContainerConfig200ResponseData) GetMp34() string`

GetMp34 returns the Mp34 field if non-nil, zero value otherwise.

### GetMp34Ok

`func (o *GetContainerConfig200ResponseData) GetMp34Ok() (*string, bool)`

GetMp34Ok returns a tuple with the Mp34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp34

`func (o *GetContainerConfig200ResponseData) SetMp34(v string)`

SetMp34 sets Mp34 field to given value.

### HasMp34

`func (o *GetContainerConfig200ResponseData) HasMp34() bool`

HasMp34 returns a boolean if a field has been set.

### GetMp35

`func (o *GetContainerConfig200ResponseData) GetMp35() string`

GetMp35 returns the Mp35 field if non-nil, zero value otherwise.

### GetMp35Ok

`func (o *GetContainerConfig200ResponseData) GetMp35Ok() (*string, bool)`

GetMp35Ok returns a tuple with the Mp35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp35

`func (o *GetContainerConfig200ResponseData) SetMp35(v string)`

SetMp35 sets Mp35 field to given value.

### HasMp35

`func (o *GetContainerConfig200ResponseData) HasMp35() bool`

HasMp35 returns a boolean if a field has been set.

### GetMp36

`func (o *GetContainerConfig200ResponseData) GetMp36() string`

GetMp36 returns the Mp36 field if non-nil, zero value otherwise.

### GetMp36Ok

`func (o *GetContainerConfig200ResponseData) GetMp36Ok() (*string, bool)`

GetMp36Ok returns a tuple with the Mp36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp36

`func (o *GetContainerConfig200ResponseData) SetMp36(v string)`

SetMp36 sets Mp36 field to given value.

### HasMp36

`func (o *GetContainerConfig200ResponseData) HasMp36() bool`

HasMp36 returns a boolean if a field has been set.

### GetMp37

`func (o *GetContainerConfig200ResponseData) GetMp37() string`

GetMp37 returns the Mp37 field if non-nil, zero value otherwise.

### GetMp37Ok

`func (o *GetContainerConfig200ResponseData) GetMp37Ok() (*string, bool)`

GetMp37Ok returns a tuple with the Mp37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp37

`func (o *GetContainerConfig200ResponseData) SetMp37(v string)`

SetMp37 sets Mp37 field to given value.

### HasMp37

`func (o *GetContainerConfig200ResponseData) HasMp37() bool`

HasMp37 returns a boolean if a field has been set.

### GetMp38

`func (o *GetContainerConfig200ResponseData) GetMp38() string`

GetMp38 returns the Mp38 field if non-nil, zero value otherwise.

### GetMp38Ok

`func (o *GetContainerConfig200ResponseData) GetMp38Ok() (*string, bool)`

GetMp38Ok returns a tuple with the Mp38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp38

`func (o *GetContainerConfig200ResponseData) SetMp38(v string)`

SetMp38 sets Mp38 field to given value.

### HasMp38

`func (o *GetContainerConfig200ResponseData) HasMp38() bool`

HasMp38 returns a boolean if a field has been set.

### GetMp39

`func (o *GetContainerConfig200ResponseData) GetMp39() string`

GetMp39 returns the Mp39 field if non-nil, zero value otherwise.

### GetMp39Ok

`func (o *GetContainerConfig200ResponseData) GetMp39Ok() (*string, bool)`

GetMp39Ok returns a tuple with the Mp39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp39

`func (o *GetContainerConfig200ResponseData) SetMp39(v string)`

SetMp39 sets Mp39 field to given value.

### HasMp39

`func (o *GetContainerConfig200ResponseData) HasMp39() bool`

HasMp39 returns a boolean if a field has been set.

### GetMp40

`func (o *GetContainerConfig200ResponseData) GetMp40() string`

GetMp40 returns the Mp40 field if non-nil, zero value otherwise.

### GetMp40Ok

`func (o *GetContainerConfig200ResponseData) GetMp40Ok() (*string, bool)`

GetMp40Ok returns a tuple with the Mp40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp40

`func (o *GetContainerConfig200ResponseData) SetMp40(v string)`

SetMp40 sets Mp40 field to given value.

### HasMp40

`func (o *GetContainerConfig200ResponseData) HasMp40() bool`

HasMp40 returns a boolean if a field has been set.

### GetMp41

`func (o *GetContainerConfig200ResponseData) GetMp41() string`

GetMp41 returns the Mp41 field if non-nil, zero value otherwise.

### GetMp41Ok

`func (o *GetContainerConfig200ResponseData) GetMp41Ok() (*string, bool)`

GetMp41Ok returns a tuple with the Mp41 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp41

`func (o *GetContainerConfig200ResponseData) SetMp41(v string)`

SetMp41 sets Mp41 field to given value.

### HasMp41

`func (o *GetContainerConfig200ResponseData) HasMp41() bool`

HasMp41 returns a boolean if a field has been set.

### GetMp42

`func (o *GetContainerConfig200ResponseData) GetMp42() string`

GetMp42 returns the Mp42 field if non-nil, zero value otherwise.

### GetMp42Ok

`func (o *GetContainerConfig200ResponseData) GetMp42Ok() (*string, bool)`

GetMp42Ok returns a tuple with the Mp42 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp42

`func (o *GetContainerConfig200ResponseData) SetMp42(v string)`

SetMp42 sets Mp42 field to given value.

### HasMp42

`func (o *GetContainerConfig200ResponseData) HasMp42() bool`

HasMp42 returns a boolean if a field has been set.

### GetMp43

`func (o *GetContainerConfig200ResponseData) GetMp43() string`

GetMp43 returns the Mp43 field if non-nil, zero value otherwise.

### GetMp43Ok

`func (o *GetContainerConfig200ResponseData) GetMp43Ok() (*string, bool)`

GetMp43Ok returns a tuple with the Mp43 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp43

`func (o *GetContainerConfig200ResponseData) SetMp43(v string)`

SetMp43 sets Mp43 field to given value.

### HasMp43

`func (o *GetContainerConfig200ResponseData) HasMp43() bool`

HasMp43 returns a boolean if a field has been set.

### GetMp44

`func (o *GetContainerConfig200ResponseData) GetMp44() string`

GetMp44 returns the Mp44 field if non-nil, zero value otherwise.

### GetMp44Ok

`func (o *GetContainerConfig200ResponseData) GetMp44Ok() (*string, bool)`

GetMp44Ok returns a tuple with the Mp44 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp44

`func (o *GetContainerConfig200ResponseData) SetMp44(v string)`

SetMp44 sets Mp44 field to given value.

### HasMp44

`func (o *GetContainerConfig200ResponseData) HasMp44() bool`

HasMp44 returns a boolean if a field has been set.

### GetMp45

`func (o *GetContainerConfig200ResponseData) GetMp45() string`

GetMp45 returns the Mp45 field if non-nil, zero value otherwise.

### GetMp45Ok

`func (o *GetContainerConfig200ResponseData) GetMp45Ok() (*string, bool)`

GetMp45Ok returns a tuple with the Mp45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp45

`func (o *GetContainerConfig200ResponseData) SetMp45(v string)`

SetMp45 sets Mp45 field to given value.

### HasMp45

`func (o *GetContainerConfig200ResponseData) HasMp45() bool`

HasMp45 returns a boolean if a field has been set.

### GetMp46

`func (o *GetContainerConfig200ResponseData) GetMp46() string`

GetMp46 returns the Mp46 field if non-nil, zero value otherwise.

### GetMp46Ok

`func (o *GetContainerConfig200ResponseData) GetMp46Ok() (*string, bool)`

GetMp46Ok returns a tuple with the Mp46 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp46

`func (o *GetContainerConfig200ResponseData) SetMp46(v string)`

SetMp46 sets Mp46 field to given value.

### HasMp46

`func (o *GetContainerConfig200ResponseData) HasMp46() bool`

HasMp46 returns a boolean if a field has been set.

### GetMp47

`func (o *GetContainerConfig200ResponseData) GetMp47() string`

GetMp47 returns the Mp47 field if non-nil, zero value otherwise.

### GetMp47Ok

`func (o *GetContainerConfig200ResponseData) GetMp47Ok() (*string, bool)`

GetMp47Ok returns a tuple with the Mp47 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp47

`func (o *GetContainerConfig200ResponseData) SetMp47(v string)`

SetMp47 sets Mp47 field to given value.

### HasMp47

`func (o *GetContainerConfig200ResponseData) HasMp47() bool`

HasMp47 returns a boolean if a field has been set.

### GetMp48

`func (o *GetContainerConfig200ResponseData) GetMp48() string`

GetMp48 returns the Mp48 field if non-nil, zero value otherwise.

### GetMp48Ok

`func (o *GetContainerConfig200ResponseData) GetMp48Ok() (*string, bool)`

GetMp48Ok returns a tuple with the Mp48 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp48

`func (o *GetContainerConfig200ResponseData) SetMp48(v string)`

SetMp48 sets Mp48 field to given value.

### HasMp48

`func (o *GetContainerConfig200ResponseData) HasMp48() bool`

HasMp48 returns a boolean if a field has been set.

### GetMp49

`func (o *GetContainerConfig200ResponseData) GetMp49() string`

GetMp49 returns the Mp49 field if non-nil, zero value otherwise.

### GetMp49Ok

`func (o *GetContainerConfig200ResponseData) GetMp49Ok() (*string, bool)`

GetMp49Ok returns a tuple with the Mp49 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp49

`func (o *GetContainerConfig200ResponseData) SetMp49(v string)`

SetMp49 sets Mp49 field to given value.

### HasMp49

`func (o *GetContainerConfig200ResponseData) HasMp49() bool`

HasMp49 returns a boolean if a field has been set.

### GetMp50

`func (o *GetContainerConfig200ResponseData) GetMp50() string`

GetMp50 returns the Mp50 field if non-nil, zero value otherwise.

### GetMp50Ok

`func (o *GetContainerConfig200ResponseData) GetMp50Ok() (*string, bool)`

GetMp50Ok returns a tuple with the Mp50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp50

`func (o *GetContainerConfig200ResponseData) SetMp50(v string)`

SetMp50 sets Mp50 field to given value.

### HasMp50

`func (o *GetContainerConfig200ResponseData) HasMp50() bool`

HasMp50 returns a boolean if a field has been set.

### GetMp51

`func (o *GetContainerConfig200ResponseData) GetMp51() string`

GetMp51 returns the Mp51 field if non-nil, zero value otherwise.

### GetMp51Ok

`func (o *GetContainerConfig200ResponseData) GetMp51Ok() (*string, bool)`

GetMp51Ok returns a tuple with the Mp51 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp51

`func (o *GetContainerConfig200ResponseData) SetMp51(v string)`

SetMp51 sets Mp51 field to given value.

### HasMp51

`func (o *GetContainerConfig200ResponseData) HasMp51() bool`

HasMp51 returns a boolean if a field has been set.

### GetMp52

`func (o *GetContainerConfig200ResponseData) GetMp52() string`

GetMp52 returns the Mp52 field if non-nil, zero value otherwise.

### GetMp52Ok

`func (o *GetContainerConfig200ResponseData) GetMp52Ok() (*string, bool)`

GetMp52Ok returns a tuple with the Mp52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp52

`func (o *GetContainerConfig200ResponseData) SetMp52(v string)`

SetMp52 sets Mp52 field to given value.

### HasMp52

`func (o *GetContainerConfig200ResponseData) HasMp52() bool`

HasMp52 returns a boolean if a field has been set.

### GetMp53

`func (o *GetContainerConfig200ResponseData) GetMp53() string`

GetMp53 returns the Mp53 field if non-nil, zero value otherwise.

### GetMp53Ok

`func (o *GetContainerConfig200ResponseData) GetMp53Ok() (*string, bool)`

GetMp53Ok returns a tuple with the Mp53 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp53

`func (o *GetContainerConfig200ResponseData) SetMp53(v string)`

SetMp53 sets Mp53 field to given value.

### HasMp53

`func (o *GetContainerConfig200ResponseData) HasMp53() bool`

HasMp53 returns a boolean if a field has been set.

### GetMp54

`func (o *GetContainerConfig200ResponseData) GetMp54() string`

GetMp54 returns the Mp54 field if non-nil, zero value otherwise.

### GetMp54Ok

`func (o *GetContainerConfig200ResponseData) GetMp54Ok() (*string, bool)`

GetMp54Ok returns a tuple with the Mp54 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp54

`func (o *GetContainerConfig200ResponseData) SetMp54(v string)`

SetMp54 sets Mp54 field to given value.

### HasMp54

`func (o *GetContainerConfig200ResponseData) HasMp54() bool`

HasMp54 returns a boolean if a field has been set.

### GetMp55

`func (o *GetContainerConfig200ResponseData) GetMp55() string`

GetMp55 returns the Mp55 field if non-nil, zero value otherwise.

### GetMp55Ok

`func (o *GetContainerConfig200ResponseData) GetMp55Ok() (*string, bool)`

GetMp55Ok returns a tuple with the Mp55 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp55

`func (o *GetContainerConfig200ResponseData) SetMp55(v string)`

SetMp55 sets Mp55 field to given value.

### HasMp55

`func (o *GetContainerConfig200ResponseData) HasMp55() bool`

HasMp55 returns a boolean if a field has been set.

### GetMp56

`func (o *GetContainerConfig200ResponseData) GetMp56() string`

GetMp56 returns the Mp56 field if non-nil, zero value otherwise.

### GetMp56Ok

`func (o *GetContainerConfig200ResponseData) GetMp56Ok() (*string, bool)`

GetMp56Ok returns a tuple with the Mp56 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp56

`func (o *GetContainerConfig200ResponseData) SetMp56(v string)`

SetMp56 sets Mp56 field to given value.

### HasMp56

`func (o *GetContainerConfig200ResponseData) HasMp56() bool`

HasMp56 returns a boolean if a field has been set.

### GetMp57

`func (o *GetContainerConfig200ResponseData) GetMp57() string`

GetMp57 returns the Mp57 field if non-nil, zero value otherwise.

### GetMp57Ok

`func (o *GetContainerConfig200ResponseData) GetMp57Ok() (*string, bool)`

GetMp57Ok returns a tuple with the Mp57 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp57

`func (o *GetContainerConfig200ResponseData) SetMp57(v string)`

SetMp57 sets Mp57 field to given value.

### HasMp57

`func (o *GetContainerConfig200ResponseData) HasMp57() bool`

HasMp57 returns a boolean if a field has been set.

### GetMp58

`func (o *GetContainerConfig200ResponseData) GetMp58() string`

GetMp58 returns the Mp58 field if non-nil, zero value otherwise.

### GetMp58Ok

`func (o *GetContainerConfig200ResponseData) GetMp58Ok() (*string, bool)`

GetMp58Ok returns a tuple with the Mp58 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp58

`func (o *GetContainerConfig200ResponseData) SetMp58(v string)`

SetMp58 sets Mp58 field to given value.

### HasMp58

`func (o *GetContainerConfig200ResponseData) HasMp58() bool`

HasMp58 returns a boolean if a field has been set.

### GetMp59

`func (o *GetContainerConfig200ResponseData) GetMp59() string`

GetMp59 returns the Mp59 field if non-nil, zero value otherwise.

### GetMp59Ok

`func (o *GetContainerConfig200ResponseData) GetMp59Ok() (*string, bool)`

GetMp59Ok returns a tuple with the Mp59 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp59

`func (o *GetContainerConfig200ResponseData) SetMp59(v string)`

SetMp59 sets Mp59 field to given value.

### HasMp59

`func (o *GetContainerConfig200ResponseData) HasMp59() bool`

HasMp59 returns a boolean if a field has been set.

### GetMp60

`func (o *GetContainerConfig200ResponseData) GetMp60() string`

GetMp60 returns the Mp60 field if non-nil, zero value otherwise.

### GetMp60Ok

`func (o *GetContainerConfig200ResponseData) GetMp60Ok() (*string, bool)`

GetMp60Ok returns a tuple with the Mp60 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp60

`func (o *GetContainerConfig200ResponseData) SetMp60(v string)`

SetMp60 sets Mp60 field to given value.

### HasMp60

`func (o *GetContainerConfig200ResponseData) HasMp60() bool`

HasMp60 returns a boolean if a field has been set.

### GetMp61

`func (o *GetContainerConfig200ResponseData) GetMp61() string`

GetMp61 returns the Mp61 field if non-nil, zero value otherwise.

### GetMp61Ok

`func (o *GetContainerConfig200ResponseData) GetMp61Ok() (*string, bool)`

GetMp61Ok returns a tuple with the Mp61 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp61

`func (o *GetContainerConfig200ResponseData) SetMp61(v string)`

SetMp61 sets Mp61 field to given value.

### HasMp61

`func (o *GetContainerConfig200ResponseData) HasMp61() bool`

HasMp61 returns a boolean if a field has been set.

### GetMp62

`func (o *GetContainerConfig200ResponseData) GetMp62() string`

GetMp62 returns the Mp62 field if non-nil, zero value otherwise.

### GetMp62Ok

`func (o *GetContainerConfig200ResponseData) GetMp62Ok() (*string, bool)`

GetMp62Ok returns a tuple with the Mp62 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp62

`func (o *GetContainerConfig200ResponseData) SetMp62(v string)`

SetMp62 sets Mp62 field to given value.

### HasMp62

`func (o *GetContainerConfig200ResponseData) HasMp62() bool`

HasMp62 returns a boolean if a field has been set.

### GetMp63

`func (o *GetContainerConfig200ResponseData) GetMp63() string`

GetMp63 returns the Mp63 field if non-nil, zero value otherwise.

### GetMp63Ok

`func (o *GetContainerConfig200ResponseData) GetMp63Ok() (*string, bool)`

GetMp63Ok returns a tuple with the Mp63 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp63

`func (o *GetContainerConfig200ResponseData) SetMp63(v string)`

SetMp63 sets Mp63 field to given value.

### HasMp63

`func (o *GetContainerConfig200ResponseData) HasMp63() bool`

HasMp63 returns a boolean if a field has been set.

### GetMp64

`func (o *GetContainerConfig200ResponseData) GetMp64() string`

GetMp64 returns the Mp64 field if non-nil, zero value otherwise.

### GetMp64Ok

`func (o *GetContainerConfig200ResponseData) GetMp64Ok() (*string, bool)`

GetMp64Ok returns a tuple with the Mp64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp64

`func (o *GetContainerConfig200ResponseData) SetMp64(v string)`

SetMp64 sets Mp64 field to given value.

### HasMp64

`func (o *GetContainerConfig200ResponseData) HasMp64() bool`

HasMp64 returns a boolean if a field has been set.

### GetMp65

`func (o *GetContainerConfig200ResponseData) GetMp65() string`

GetMp65 returns the Mp65 field if non-nil, zero value otherwise.

### GetMp65Ok

`func (o *GetContainerConfig200ResponseData) GetMp65Ok() (*string, bool)`

GetMp65Ok returns a tuple with the Mp65 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp65

`func (o *GetContainerConfig200ResponseData) SetMp65(v string)`

SetMp65 sets Mp65 field to given value.

### HasMp65

`func (o *GetContainerConfig200ResponseData) HasMp65() bool`

HasMp65 returns a boolean if a field has been set.

### GetMp66

`func (o *GetContainerConfig200ResponseData) GetMp66() string`

GetMp66 returns the Mp66 field if non-nil, zero value otherwise.

### GetMp66Ok

`func (o *GetContainerConfig200ResponseData) GetMp66Ok() (*string, bool)`

GetMp66Ok returns a tuple with the Mp66 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp66

`func (o *GetContainerConfig200ResponseData) SetMp66(v string)`

SetMp66 sets Mp66 field to given value.

### HasMp66

`func (o *GetContainerConfig200ResponseData) HasMp66() bool`

HasMp66 returns a boolean if a field has been set.

### GetMp67

`func (o *GetContainerConfig200ResponseData) GetMp67() string`

GetMp67 returns the Mp67 field if non-nil, zero value otherwise.

### GetMp67Ok

`func (o *GetContainerConfig200ResponseData) GetMp67Ok() (*string, bool)`

GetMp67Ok returns a tuple with the Mp67 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp67

`func (o *GetContainerConfig200ResponseData) SetMp67(v string)`

SetMp67 sets Mp67 field to given value.

### HasMp67

`func (o *GetContainerConfig200ResponseData) HasMp67() bool`

HasMp67 returns a boolean if a field has been set.

### GetMp68

`func (o *GetContainerConfig200ResponseData) GetMp68() string`

GetMp68 returns the Mp68 field if non-nil, zero value otherwise.

### GetMp68Ok

`func (o *GetContainerConfig200ResponseData) GetMp68Ok() (*string, bool)`

GetMp68Ok returns a tuple with the Mp68 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp68

`func (o *GetContainerConfig200ResponseData) SetMp68(v string)`

SetMp68 sets Mp68 field to given value.

### HasMp68

`func (o *GetContainerConfig200ResponseData) HasMp68() bool`

HasMp68 returns a boolean if a field has been set.

### GetMp69

`func (o *GetContainerConfig200ResponseData) GetMp69() string`

GetMp69 returns the Mp69 field if non-nil, zero value otherwise.

### GetMp69Ok

`func (o *GetContainerConfig200ResponseData) GetMp69Ok() (*string, bool)`

GetMp69Ok returns a tuple with the Mp69 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp69

`func (o *GetContainerConfig200ResponseData) SetMp69(v string)`

SetMp69 sets Mp69 field to given value.

### HasMp69

`func (o *GetContainerConfig200ResponseData) HasMp69() bool`

HasMp69 returns a boolean if a field has been set.

### GetMp70

`func (o *GetContainerConfig200ResponseData) GetMp70() string`

GetMp70 returns the Mp70 field if non-nil, zero value otherwise.

### GetMp70Ok

`func (o *GetContainerConfig200ResponseData) GetMp70Ok() (*string, bool)`

GetMp70Ok returns a tuple with the Mp70 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp70

`func (o *GetContainerConfig200ResponseData) SetMp70(v string)`

SetMp70 sets Mp70 field to given value.

### HasMp70

`func (o *GetContainerConfig200ResponseData) HasMp70() bool`

HasMp70 returns a boolean if a field has been set.

### GetMp71

`func (o *GetContainerConfig200ResponseData) GetMp71() string`

GetMp71 returns the Mp71 field if non-nil, zero value otherwise.

### GetMp71Ok

`func (o *GetContainerConfig200ResponseData) GetMp71Ok() (*string, bool)`

GetMp71Ok returns a tuple with the Mp71 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp71

`func (o *GetContainerConfig200ResponseData) SetMp71(v string)`

SetMp71 sets Mp71 field to given value.

### HasMp71

`func (o *GetContainerConfig200ResponseData) HasMp71() bool`

HasMp71 returns a boolean if a field has been set.

### GetMp72

`func (o *GetContainerConfig200ResponseData) GetMp72() string`

GetMp72 returns the Mp72 field if non-nil, zero value otherwise.

### GetMp72Ok

`func (o *GetContainerConfig200ResponseData) GetMp72Ok() (*string, bool)`

GetMp72Ok returns a tuple with the Mp72 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp72

`func (o *GetContainerConfig200ResponseData) SetMp72(v string)`

SetMp72 sets Mp72 field to given value.

### HasMp72

`func (o *GetContainerConfig200ResponseData) HasMp72() bool`

HasMp72 returns a boolean if a field has been set.

### GetMp73

`func (o *GetContainerConfig200ResponseData) GetMp73() string`

GetMp73 returns the Mp73 field if non-nil, zero value otherwise.

### GetMp73Ok

`func (o *GetContainerConfig200ResponseData) GetMp73Ok() (*string, bool)`

GetMp73Ok returns a tuple with the Mp73 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp73

`func (o *GetContainerConfig200ResponseData) SetMp73(v string)`

SetMp73 sets Mp73 field to given value.

### HasMp73

`func (o *GetContainerConfig200ResponseData) HasMp73() bool`

HasMp73 returns a boolean if a field has been set.

### GetMp74

`func (o *GetContainerConfig200ResponseData) GetMp74() string`

GetMp74 returns the Mp74 field if non-nil, zero value otherwise.

### GetMp74Ok

`func (o *GetContainerConfig200ResponseData) GetMp74Ok() (*string, bool)`

GetMp74Ok returns a tuple with the Mp74 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp74

`func (o *GetContainerConfig200ResponseData) SetMp74(v string)`

SetMp74 sets Mp74 field to given value.

### HasMp74

`func (o *GetContainerConfig200ResponseData) HasMp74() bool`

HasMp74 returns a boolean if a field has been set.

### GetMp75

`func (o *GetContainerConfig200ResponseData) GetMp75() string`

GetMp75 returns the Mp75 field if non-nil, zero value otherwise.

### GetMp75Ok

`func (o *GetContainerConfig200ResponseData) GetMp75Ok() (*string, bool)`

GetMp75Ok returns a tuple with the Mp75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp75

`func (o *GetContainerConfig200ResponseData) SetMp75(v string)`

SetMp75 sets Mp75 field to given value.

### HasMp75

`func (o *GetContainerConfig200ResponseData) HasMp75() bool`

HasMp75 returns a boolean if a field has been set.

### GetMp76

`func (o *GetContainerConfig200ResponseData) GetMp76() string`

GetMp76 returns the Mp76 field if non-nil, zero value otherwise.

### GetMp76Ok

`func (o *GetContainerConfig200ResponseData) GetMp76Ok() (*string, bool)`

GetMp76Ok returns a tuple with the Mp76 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp76

`func (o *GetContainerConfig200ResponseData) SetMp76(v string)`

SetMp76 sets Mp76 field to given value.

### HasMp76

`func (o *GetContainerConfig200ResponseData) HasMp76() bool`

HasMp76 returns a boolean if a field has been set.

### GetMp77

`func (o *GetContainerConfig200ResponseData) GetMp77() string`

GetMp77 returns the Mp77 field if non-nil, zero value otherwise.

### GetMp77Ok

`func (o *GetContainerConfig200ResponseData) GetMp77Ok() (*string, bool)`

GetMp77Ok returns a tuple with the Mp77 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp77

`func (o *GetContainerConfig200ResponseData) SetMp77(v string)`

SetMp77 sets Mp77 field to given value.

### HasMp77

`func (o *GetContainerConfig200ResponseData) HasMp77() bool`

HasMp77 returns a boolean if a field has been set.

### GetMp78

`func (o *GetContainerConfig200ResponseData) GetMp78() string`

GetMp78 returns the Mp78 field if non-nil, zero value otherwise.

### GetMp78Ok

`func (o *GetContainerConfig200ResponseData) GetMp78Ok() (*string, bool)`

GetMp78Ok returns a tuple with the Mp78 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp78

`func (o *GetContainerConfig200ResponseData) SetMp78(v string)`

SetMp78 sets Mp78 field to given value.

### HasMp78

`func (o *GetContainerConfig200ResponseData) HasMp78() bool`

HasMp78 returns a boolean if a field has been set.

### GetMp79

`func (o *GetContainerConfig200ResponseData) GetMp79() string`

GetMp79 returns the Mp79 field if non-nil, zero value otherwise.

### GetMp79Ok

`func (o *GetContainerConfig200ResponseData) GetMp79Ok() (*string, bool)`

GetMp79Ok returns a tuple with the Mp79 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp79

`func (o *GetContainerConfig200ResponseData) SetMp79(v string)`

SetMp79 sets Mp79 field to given value.

### HasMp79

`func (o *GetContainerConfig200ResponseData) HasMp79() bool`

HasMp79 returns a boolean if a field has been set.

### GetMp80

`func (o *GetContainerConfig200ResponseData) GetMp80() string`

GetMp80 returns the Mp80 field if non-nil, zero value otherwise.

### GetMp80Ok

`func (o *GetContainerConfig200ResponseData) GetMp80Ok() (*string, bool)`

GetMp80Ok returns a tuple with the Mp80 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp80

`func (o *GetContainerConfig200ResponseData) SetMp80(v string)`

SetMp80 sets Mp80 field to given value.

### HasMp80

`func (o *GetContainerConfig200ResponseData) HasMp80() bool`

HasMp80 returns a boolean if a field has been set.

### GetMp81

`func (o *GetContainerConfig200ResponseData) GetMp81() string`

GetMp81 returns the Mp81 field if non-nil, zero value otherwise.

### GetMp81Ok

`func (o *GetContainerConfig200ResponseData) GetMp81Ok() (*string, bool)`

GetMp81Ok returns a tuple with the Mp81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp81

`func (o *GetContainerConfig200ResponseData) SetMp81(v string)`

SetMp81 sets Mp81 field to given value.

### HasMp81

`func (o *GetContainerConfig200ResponseData) HasMp81() bool`

HasMp81 returns a boolean if a field has been set.

### GetMp82

`func (o *GetContainerConfig200ResponseData) GetMp82() string`

GetMp82 returns the Mp82 field if non-nil, zero value otherwise.

### GetMp82Ok

`func (o *GetContainerConfig200ResponseData) GetMp82Ok() (*string, bool)`

GetMp82Ok returns a tuple with the Mp82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp82

`func (o *GetContainerConfig200ResponseData) SetMp82(v string)`

SetMp82 sets Mp82 field to given value.

### HasMp82

`func (o *GetContainerConfig200ResponseData) HasMp82() bool`

HasMp82 returns a boolean if a field has been set.

### GetMp83

`func (o *GetContainerConfig200ResponseData) GetMp83() string`

GetMp83 returns the Mp83 field if non-nil, zero value otherwise.

### GetMp83Ok

`func (o *GetContainerConfig200ResponseData) GetMp83Ok() (*string, bool)`

GetMp83Ok returns a tuple with the Mp83 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp83

`func (o *GetContainerConfig200ResponseData) SetMp83(v string)`

SetMp83 sets Mp83 field to given value.

### HasMp83

`func (o *GetContainerConfig200ResponseData) HasMp83() bool`

HasMp83 returns a boolean if a field has been set.

### GetMp84

`func (o *GetContainerConfig200ResponseData) GetMp84() string`

GetMp84 returns the Mp84 field if non-nil, zero value otherwise.

### GetMp84Ok

`func (o *GetContainerConfig200ResponseData) GetMp84Ok() (*string, bool)`

GetMp84Ok returns a tuple with the Mp84 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp84

`func (o *GetContainerConfig200ResponseData) SetMp84(v string)`

SetMp84 sets Mp84 field to given value.

### HasMp84

`func (o *GetContainerConfig200ResponseData) HasMp84() bool`

HasMp84 returns a boolean if a field has been set.

### GetMp85

`func (o *GetContainerConfig200ResponseData) GetMp85() string`

GetMp85 returns the Mp85 field if non-nil, zero value otherwise.

### GetMp85Ok

`func (o *GetContainerConfig200ResponseData) GetMp85Ok() (*string, bool)`

GetMp85Ok returns a tuple with the Mp85 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp85

`func (o *GetContainerConfig200ResponseData) SetMp85(v string)`

SetMp85 sets Mp85 field to given value.

### HasMp85

`func (o *GetContainerConfig200ResponseData) HasMp85() bool`

HasMp85 returns a boolean if a field has been set.

### GetMp86

`func (o *GetContainerConfig200ResponseData) GetMp86() string`

GetMp86 returns the Mp86 field if non-nil, zero value otherwise.

### GetMp86Ok

`func (o *GetContainerConfig200ResponseData) GetMp86Ok() (*string, bool)`

GetMp86Ok returns a tuple with the Mp86 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp86

`func (o *GetContainerConfig200ResponseData) SetMp86(v string)`

SetMp86 sets Mp86 field to given value.

### HasMp86

`func (o *GetContainerConfig200ResponseData) HasMp86() bool`

HasMp86 returns a boolean if a field has been set.

### GetMp87

`func (o *GetContainerConfig200ResponseData) GetMp87() string`

GetMp87 returns the Mp87 field if non-nil, zero value otherwise.

### GetMp87Ok

`func (o *GetContainerConfig200ResponseData) GetMp87Ok() (*string, bool)`

GetMp87Ok returns a tuple with the Mp87 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp87

`func (o *GetContainerConfig200ResponseData) SetMp87(v string)`

SetMp87 sets Mp87 field to given value.

### HasMp87

`func (o *GetContainerConfig200ResponseData) HasMp87() bool`

HasMp87 returns a boolean if a field has been set.

### GetMp88

`func (o *GetContainerConfig200ResponseData) GetMp88() string`

GetMp88 returns the Mp88 field if non-nil, zero value otherwise.

### GetMp88Ok

`func (o *GetContainerConfig200ResponseData) GetMp88Ok() (*string, bool)`

GetMp88Ok returns a tuple with the Mp88 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp88

`func (o *GetContainerConfig200ResponseData) SetMp88(v string)`

SetMp88 sets Mp88 field to given value.

### HasMp88

`func (o *GetContainerConfig200ResponseData) HasMp88() bool`

HasMp88 returns a boolean if a field has been set.

### GetMp89

`func (o *GetContainerConfig200ResponseData) GetMp89() string`

GetMp89 returns the Mp89 field if non-nil, zero value otherwise.

### GetMp89Ok

`func (o *GetContainerConfig200ResponseData) GetMp89Ok() (*string, bool)`

GetMp89Ok returns a tuple with the Mp89 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp89

`func (o *GetContainerConfig200ResponseData) SetMp89(v string)`

SetMp89 sets Mp89 field to given value.

### HasMp89

`func (o *GetContainerConfig200ResponseData) HasMp89() bool`

HasMp89 returns a boolean if a field has been set.

### GetMp90

`func (o *GetContainerConfig200ResponseData) GetMp90() string`

GetMp90 returns the Mp90 field if non-nil, zero value otherwise.

### GetMp90Ok

`func (o *GetContainerConfig200ResponseData) GetMp90Ok() (*string, bool)`

GetMp90Ok returns a tuple with the Mp90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp90

`func (o *GetContainerConfig200ResponseData) SetMp90(v string)`

SetMp90 sets Mp90 field to given value.

### HasMp90

`func (o *GetContainerConfig200ResponseData) HasMp90() bool`

HasMp90 returns a boolean if a field has been set.

### GetMp91

`func (o *GetContainerConfig200ResponseData) GetMp91() string`

GetMp91 returns the Mp91 field if non-nil, zero value otherwise.

### GetMp91Ok

`func (o *GetContainerConfig200ResponseData) GetMp91Ok() (*string, bool)`

GetMp91Ok returns a tuple with the Mp91 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp91

`func (o *GetContainerConfig200ResponseData) SetMp91(v string)`

SetMp91 sets Mp91 field to given value.

### HasMp91

`func (o *GetContainerConfig200ResponseData) HasMp91() bool`

HasMp91 returns a boolean if a field has been set.

### GetMp92

`func (o *GetContainerConfig200ResponseData) GetMp92() string`

GetMp92 returns the Mp92 field if non-nil, zero value otherwise.

### GetMp92Ok

`func (o *GetContainerConfig200ResponseData) GetMp92Ok() (*string, bool)`

GetMp92Ok returns a tuple with the Mp92 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp92

`func (o *GetContainerConfig200ResponseData) SetMp92(v string)`

SetMp92 sets Mp92 field to given value.

### HasMp92

`func (o *GetContainerConfig200ResponseData) HasMp92() bool`

HasMp92 returns a boolean if a field has been set.

### GetMp93

`func (o *GetContainerConfig200ResponseData) GetMp93() string`

GetMp93 returns the Mp93 field if non-nil, zero value otherwise.

### GetMp93Ok

`func (o *GetContainerConfig200ResponseData) GetMp93Ok() (*string, bool)`

GetMp93Ok returns a tuple with the Mp93 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp93

`func (o *GetContainerConfig200ResponseData) SetMp93(v string)`

SetMp93 sets Mp93 field to given value.

### HasMp93

`func (o *GetContainerConfig200ResponseData) HasMp93() bool`

HasMp93 returns a boolean if a field has been set.

### GetMp94

`func (o *GetContainerConfig200ResponseData) GetMp94() string`

GetMp94 returns the Mp94 field if non-nil, zero value otherwise.

### GetMp94Ok

`func (o *GetContainerConfig200ResponseData) GetMp94Ok() (*string, bool)`

GetMp94Ok returns a tuple with the Mp94 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp94

`func (o *GetContainerConfig200ResponseData) SetMp94(v string)`

SetMp94 sets Mp94 field to given value.

### HasMp94

`func (o *GetContainerConfig200ResponseData) HasMp94() bool`

HasMp94 returns a boolean if a field has been set.

### GetMp95

`func (o *GetContainerConfig200ResponseData) GetMp95() string`

GetMp95 returns the Mp95 field if non-nil, zero value otherwise.

### GetMp95Ok

`func (o *GetContainerConfig200ResponseData) GetMp95Ok() (*string, bool)`

GetMp95Ok returns a tuple with the Mp95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp95

`func (o *GetContainerConfig200ResponseData) SetMp95(v string)`

SetMp95 sets Mp95 field to given value.

### HasMp95

`func (o *GetContainerConfig200ResponseData) HasMp95() bool`

HasMp95 returns a boolean if a field has been set.

### GetMp96

`func (o *GetContainerConfig200ResponseData) GetMp96() string`

GetMp96 returns the Mp96 field if non-nil, zero value otherwise.

### GetMp96Ok

`func (o *GetContainerConfig200ResponseData) GetMp96Ok() (*string, bool)`

GetMp96Ok returns a tuple with the Mp96 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp96

`func (o *GetContainerConfig200ResponseData) SetMp96(v string)`

SetMp96 sets Mp96 field to given value.

### HasMp96

`func (o *GetContainerConfig200ResponseData) HasMp96() bool`

HasMp96 returns a boolean if a field has been set.

### GetMp97

`func (o *GetContainerConfig200ResponseData) GetMp97() string`

GetMp97 returns the Mp97 field if non-nil, zero value otherwise.

### GetMp97Ok

`func (o *GetContainerConfig200ResponseData) GetMp97Ok() (*string, bool)`

GetMp97Ok returns a tuple with the Mp97 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp97

`func (o *GetContainerConfig200ResponseData) SetMp97(v string)`

SetMp97 sets Mp97 field to given value.

### HasMp97

`func (o *GetContainerConfig200ResponseData) HasMp97() bool`

HasMp97 returns a boolean if a field has been set.

### GetMp98

`func (o *GetContainerConfig200ResponseData) GetMp98() string`

GetMp98 returns the Mp98 field if non-nil, zero value otherwise.

### GetMp98Ok

`func (o *GetContainerConfig200ResponseData) GetMp98Ok() (*string, bool)`

GetMp98Ok returns a tuple with the Mp98 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp98

`func (o *GetContainerConfig200ResponseData) SetMp98(v string)`

SetMp98 sets Mp98 field to given value.

### HasMp98

`func (o *GetContainerConfig200ResponseData) HasMp98() bool`

HasMp98 returns a boolean if a field has been set.

### GetMp99

`func (o *GetContainerConfig200ResponseData) GetMp99() string`

GetMp99 returns the Mp99 field if non-nil, zero value otherwise.

### GetMp99Ok

`func (o *GetContainerConfig200ResponseData) GetMp99Ok() (*string, bool)`

GetMp99Ok returns a tuple with the Mp99 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp99

`func (o *GetContainerConfig200ResponseData) SetMp99(v string)`

SetMp99 sets Mp99 field to given value.

### HasMp99

`func (o *GetContainerConfig200ResponseData) HasMp99() bool`

HasMp99 returns a boolean if a field has been set.

### GetMp100

`func (o *GetContainerConfig200ResponseData) GetMp100() string`

GetMp100 returns the Mp100 field if non-nil, zero value otherwise.

### GetMp100Ok

`func (o *GetContainerConfig200ResponseData) GetMp100Ok() (*string, bool)`

GetMp100Ok returns a tuple with the Mp100 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp100

`func (o *GetContainerConfig200ResponseData) SetMp100(v string)`

SetMp100 sets Mp100 field to given value.

### HasMp100

`func (o *GetContainerConfig200ResponseData) HasMp100() bool`

HasMp100 returns a boolean if a field has been set.

### GetMp101

`func (o *GetContainerConfig200ResponseData) GetMp101() string`

GetMp101 returns the Mp101 field if non-nil, zero value otherwise.

### GetMp101Ok

`func (o *GetContainerConfig200ResponseData) GetMp101Ok() (*string, bool)`

GetMp101Ok returns a tuple with the Mp101 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp101

`func (o *GetContainerConfig200ResponseData) SetMp101(v string)`

SetMp101 sets Mp101 field to given value.

### HasMp101

`func (o *GetContainerConfig200ResponseData) HasMp101() bool`

HasMp101 returns a boolean if a field has been set.

### GetMp102

`func (o *GetContainerConfig200ResponseData) GetMp102() string`

GetMp102 returns the Mp102 field if non-nil, zero value otherwise.

### GetMp102Ok

`func (o *GetContainerConfig200ResponseData) GetMp102Ok() (*string, bool)`

GetMp102Ok returns a tuple with the Mp102 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp102

`func (o *GetContainerConfig200ResponseData) SetMp102(v string)`

SetMp102 sets Mp102 field to given value.

### HasMp102

`func (o *GetContainerConfig200ResponseData) HasMp102() bool`

HasMp102 returns a boolean if a field has been set.

### GetMp103

`func (o *GetContainerConfig200ResponseData) GetMp103() string`

GetMp103 returns the Mp103 field if non-nil, zero value otherwise.

### GetMp103Ok

`func (o *GetContainerConfig200ResponseData) GetMp103Ok() (*string, bool)`

GetMp103Ok returns a tuple with the Mp103 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp103

`func (o *GetContainerConfig200ResponseData) SetMp103(v string)`

SetMp103 sets Mp103 field to given value.

### HasMp103

`func (o *GetContainerConfig200ResponseData) HasMp103() bool`

HasMp103 returns a boolean if a field has been set.

### GetMp104

`func (o *GetContainerConfig200ResponseData) GetMp104() string`

GetMp104 returns the Mp104 field if non-nil, zero value otherwise.

### GetMp104Ok

`func (o *GetContainerConfig200ResponseData) GetMp104Ok() (*string, bool)`

GetMp104Ok returns a tuple with the Mp104 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp104

`func (o *GetContainerConfig200ResponseData) SetMp104(v string)`

SetMp104 sets Mp104 field to given value.

### HasMp104

`func (o *GetContainerConfig200ResponseData) HasMp104() bool`

HasMp104 returns a boolean if a field has been set.

### GetMp105

`func (o *GetContainerConfig200ResponseData) GetMp105() string`

GetMp105 returns the Mp105 field if non-nil, zero value otherwise.

### GetMp105Ok

`func (o *GetContainerConfig200ResponseData) GetMp105Ok() (*string, bool)`

GetMp105Ok returns a tuple with the Mp105 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp105

`func (o *GetContainerConfig200ResponseData) SetMp105(v string)`

SetMp105 sets Mp105 field to given value.

### HasMp105

`func (o *GetContainerConfig200ResponseData) HasMp105() bool`

HasMp105 returns a boolean if a field has been set.

### GetMp106

`func (o *GetContainerConfig200ResponseData) GetMp106() string`

GetMp106 returns the Mp106 field if non-nil, zero value otherwise.

### GetMp106Ok

`func (o *GetContainerConfig200ResponseData) GetMp106Ok() (*string, bool)`

GetMp106Ok returns a tuple with the Mp106 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp106

`func (o *GetContainerConfig200ResponseData) SetMp106(v string)`

SetMp106 sets Mp106 field to given value.

### HasMp106

`func (o *GetContainerConfig200ResponseData) HasMp106() bool`

HasMp106 returns a boolean if a field has been set.

### GetMp107

`func (o *GetContainerConfig200ResponseData) GetMp107() string`

GetMp107 returns the Mp107 field if non-nil, zero value otherwise.

### GetMp107Ok

`func (o *GetContainerConfig200ResponseData) GetMp107Ok() (*string, bool)`

GetMp107Ok returns a tuple with the Mp107 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp107

`func (o *GetContainerConfig200ResponseData) SetMp107(v string)`

SetMp107 sets Mp107 field to given value.

### HasMp107

`func (o *GetContainerConfig200ResponseData) HasMp107() bool`

HasMp107 returns a boolean if a field has been set.

### GetMp108

`func (o *GetContainerConfig200ResponseData) GetMp108() string`

GetMp108 returns the Mp108 field if non-nil, zero value otherwise.

### GetMp108Ok

`func (o *GetContainerConfig200ResponseData) GetMp108Ok() (*string, bool)`

GetMp108Ok returns a tuple with the Mp108 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp108

`func (o *GetContainerConfig200ResponseData) SetMp108(v string)`

SetMp108 sets Mp108 field to given value.

### HasMp108

`func (o *GetContainerConfig200ResponseData) HasMp108() bool`

HasMp108 returns a boolean if a field has been set.

### GetMp109

`func (o *GetContainerConfig200ResponseData) GetMp109() string`

GetMp109 returns the Mp109 field if non-nil, zero value otherwise.

### GetMp109Ok

`func (o *GetContainerConfig200ResponseData) GetMp109Ok() (*string, bool)`

GetMp109Ok returns a tuple with the Mp109 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp109

`func (o *GetContainerConfig200ResponseData) SetMp109(v string)`

SetMp109 sets Mp109 field to given value.

### HasMp109

`func (o *GetContainerConfig200ResponseData) HasMp109() bool`

HasMp109 returns a boolean if a field has been set.

### GetMp110

`func (o *GetContainerConfig200ResponseData) GetMp110() string`

GetMp110 returns the Mp110 field if non-nil, zero value otherwise.

### GetMp110Ok

`func (o *GetContainerConfig200ResponseData) GetMp110Ok() (*string, bool)`

GetMp110Ok returns a tuple with the Mp110 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp110

`func (o *GetContainerConfig200ResponseData) SetMp110(v string)`

SetMp110 sets Mp110 field to given value.

### HasMp110

`func (o *GetContainerConfig200ResponseData) HasMp110() bool`

HasMp110 returns a boolean if a field has been set.

### GetMp111

`func (o *GetContainerConfig200ResponseData) GetMp111() string`

GetMp111 returns the Mp111 field if non-nil, zero value otherwise.

### GetMp111Ok

`func (o *GetContainerConfig200ResponseData) GetMp111Ok() (*string, bool)`

GetMp111Ok returns a tuple with the Mp111 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp111

`func (o *GetContainerConfig200ResponseData) SetMp111(v string)`

SetMp111 sets Mp111 field to given value.

### HasMp111

`func (o *GetContainerConfig200ResponseData) HasMp111() bool`

HasMp111 returns a boolean if a field has been set.

### GetMp112

`func (o *GetContainerConfig200ResponseData) GetMp112() string`

GetMp112 returns the Mp112 field if non-nil, zero value otherwise.

### GetMp112Ok

`func (o *GetContainerConfig200ResponseData) GetMp112Ok() (*string, bool)`

GetMp112Ok returns a tuple with the Mp112 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp112

`func (o *GetContainerConfig200ResponseData) SetMp112(v string)`

SetMp112 sets Mp112 field to given value.

### HasMp112

`func (o *GetContainerConfig200ResponseData) HasMp112() bool`

HasMp112 returns a boolean if a field has been set.

### GetMp113

`func (o *GetContainerConfig200ResponseData) GetMp113() string`

GetMp113 returns the Mp113 field if non-nil, zero value otherwise.

### GetMp113Ok

`func (o *GetContainerConfig200ResponseData) GetMp113Ok() (*string, bool)`

GetMp113Ok returns a tuple with the Mp113 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp113

`func (o *GetContainerConfig200ResponseData) SetMp113(v string)`

SetMp113 sets Mp113 field to given value.

### HasMp113

`func (o *GetContainerConfig200ResponseData) HasMp113() bool`

HasMp113 returns a boolean if a field has been set.

### GetMp114

`func (o *GetContainerConfig200ResponseData) GetMp114() string`

GetMp114 returns the Mp114 field if non-nil, zero value otherwise.

### GetMp114Ok

`func (o *GetContainerConfig200ResponseData) GetMp114Ok() (*string, bool)`

GetMp114Ok returns a tuple with the Mp114 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp114

`func (o *GetContainerConfig200ResponseData) SetMp114(v string)`

SetMp114 sets Mp114 field to given value.

### HasMp114

`func (o *GetContainerConfig200ResponseData) HasMp114() bool`

HasMp114 returns a boolean if a field has been set.

### GetMp115

`func (o *GetContainerConfig200ResponseData) GetMp115() string`

GetMp115 returns the Mp115 field if non-nil, zero value otherwise.

### GetMp115Ok

`func (o *GetContainerConfig200ResponseData) GetMp115Ok() (*string, bool)`

GetMp115Ok returns a tuple with the Mp115 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp115

`func (o *GetContainerConfig200ResponseData) SetMp115(v string)`

SetMp115 sets Mp115 field to given value.

### HasMp115

`func (o *GetContainerConfig200ResponseData) HasMp115() bool`

HasMp115 returns a boolean if a field has been set.

### GetMp116

`func (o *GetContainerConfig200ResponseData) GetMp116() string`

GetMp116 returns the Mp116 field if non-nil, zero value otherwise.

### GetMp116Ok

`func (o *GetContainerConfig200ResponseData) GetMp116Ok() (*string, bool)`

GetMp116Ok returns a tuple with the Mp116 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp116

`func (o *GetContainerConfig200ResponseData) SetMp116(v string)`

SetMp116 sets Mp116 field to given value.

### HasMp116

`func (o *GetContainerConfig200ResponseData) HasMp116() bool`

HasMp116 returns a boolean if a field has been set.

### GetMp117

`func (o *GetContainerConfig200ResponseData) GetMp117() string`

GetMp117 returns the Mp117 field if non-nil, zero value otherwise.

### GetMp117Ok

`func (o *GetContainerConfig200ResponseData) GetMp117Ok() (*string, bool)`

GetMp117Ok returns a tuple with the Mp117 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp117

`func (o *GetContainerConfig200ResponseData) SetMp117(v string)`

SetMp117 sets Mp117 field to given value.

### HasMp117

`func (o *GetContainerConfig200ResponseData) HasMp117() bool`

HasMp117 returns a boolean if a field has been set.

### GetMp118

`func (o *GetContainerConfig200ResponseData) GetMp118() string`

GetMp118 returns the Mp118 field if non-nil, zero value otherwise.

### GetMp118Ok

`func (o *GetContainerConfig200ResponseData) GetMp118Ok() (*string, bool)`

GetMp118Ok returns a tuple with the Mp118 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp118

`func (o *GetContainerConfig200ResponseData) SetMp118(v string)`

SetMp118 sets Mp118 field to given value.

### HasMp118

`func (o *GetContainerConfig200ResponseData) HasMp118() bool`

HasMp118 returns a boolean if a field has been set.

### GetMp119

`func (o *GetContainerConfig200ResponseData) GetMp119() string`

GetMp119 returns the Mp119 field if non-nil, zero value otherwise.

### GetMp119Ok

`func (o *GetContainerConfig200ResponseData) GetMp119Ok() (*string, bool)`

GetMp119Ok returns a tuple with the Mp119 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp119

`func (o *GetContainerConfig200ResponseData) SetMp119(v string)`

SetMp119 sets Mp119 field to given value.

### HasMp119

`func (o *GetContainerConfig200ResponseData) HasMp119() bool`

HasMp119 returns a boolean if a field has been set.

### GetMp120

`func (o *GetContainerConfig200ResponseData) GetMp120() string`

GetMp120 returns the Mp120 field if non-nil, zero value otherwise.

### GetMp120Ok

`func (o *GetContainerConfig200ResponseData) GetMp120Ok() (*string, bool)`

GetMp120Ok returns a tuple with the Mp120 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp120

`func (o *GetContainerConfig200ResponseData) SetMp120(v string)`

SetMp120 sets Mp120 field to given value.

### HasMp120

`func (o *GetContainerConfig200ResponseData) HasMp120() bool`

HasMp120 returns a boolean if a field has been set.

### GetMp121

`func (o *GetContainerConfig200ResponseData) GetMp121() string`

GetMp121 returns the Mp121 field if non-nil, zero value otherwise.

### GetMp121Ok

`func (o *GetContainerConfig200ResponseData) GetMp121Ok() (*string, bool)`

GetMp121Ok returns a tuple with the Mp121 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp121

`func (o *GetContainerConfig200ResponseData) SetMp121(v string)`

SetMp121 sets Mp121 field to given value.

### HasMp121

`func (o *GetContainerConfig200ResponseData) HasMp121() bool`

HasMp121 returns a boolean if a field has been set.

### GetMp122

`func (o *GetContainerConfig200ResponseData) GetMp122() string`

GetMp122 returns the Mp122 field if non-nil, zero value otherwise.

### GetMp122Ok

`func (o *GetContainerConfig200ResponseData) GetMp122Ok() (*string, bool)`

GetMp122Ok returns a tuple with the Mp122 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp122

`func (o *GetContainerConfig200ResponseData) SetMp122(v string)`

SetMp122 sets Mp122 field to given value.

### HasMp122

`func (o *GetContainerConfig200ResponseData) HasMp122() bool`

HasMp122 returns a boolean if a field has been set.

### GetMp123

`func (o *GetContainerConfig200ResponseData) GetMp123() string`

GetMp123 returns the Mp123 field if non-nil, zero value otherwise.

### GetMp123Ok

`func (o *GetContainerConfig200ResponseData) GetMp123Ok() (*string, bool)`

GetMp123Ok returns a tuple with the Mp123 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp123

`func (o *GetContainerConfig200ResponseData) SetMp123(v string)`

SetMp123 sets Mp123 field to given value.

### HasMp123

`func (o *GetContainerConfig200ResponseData) HasMp123() bool`

HasMp123 returns a boolean if a field has been set.

### GetMp124

`func (o *GetContainerConfig200ResponseData) GetMp124() string`

GetMp124 returns the Mp124 field if non-nil, zero value otherwise.

### GetMp124Ok

`func (o *GetContainerConfig200ResponseData) GetMp124Ok() (*string, bool)`

GetMp124Ok returns a tuple with the Mp124 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp124

`func (o *GetContainerConfig200ResponseData) SetMp124(v string)`

SetMp124 sets Mp124 field to given value.

### HasMp124

`func (o *GetContainerConfig200ResponseData) HasMp124() bool`

HasMp124 returns a boolean if a field has been set.

### GetMp125

`func (o *GetContainerConfig200ResponseData) GetMp125() string`

GetMp125 returns the Mp125 field if non-nil, zero value otherwise.

### GetMp125Ok

`func (o *GetContainerConfig200ResponseData) GetMp125Ok() (*string, bool)`

GetMp125Ok returns a tuple with the Mp125 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp125

`func (o *GetContainerConfig200ResponseData) SetMp125(v string)`

SetMp125 sets Mp125 field to given value.

### HasMp125

`func (o *GetContainerConfig200ResponseData) HasMp125() bool`

HasMp125 returns a boolean if a field has been set.

### GetMp126

`func (o *GetContainerConfig200ResponseData) GetMp126() string`

GetMp126 returns the Mp126 field if non-nil, zero value otherwise.

### GetMp126Ok

`func (o *GetContainerConfig200ResponseData) GetMp126Ok() (*string, bool)`

GetMp126Ok returns a tuple with the Mp126 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp126

`func (o *GetContainerConfig200ResponseData) SetMp126(v string)`

SetMp126 sets Mp126 field to given value.

### HasMp126

`func (o *GetContainerConfig200ResponseData) HasMp126() bool`

HasMp126 returns a boolean if a field has been set.

### GetMp127

`func (o *GetContainerConfig200ResponseData) GetMp127() string`

GetMp127 returns the Mp127 field if non-nil, zero value otherwise.

### GetMp127Ok

`func (o *GetContainerConfig200ResponseData) GetMp127Ok() (*string, bool)`

GetMp127Ok returns a tuple with the Mp127 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp127

`func (o *GetContainerConfig200ResponseData) SetMp127(v string)`

SetMp127 sets Mp127 field to given value.

### HasMp127

`func (o *GetContainerConfig200ResponseData) HasMp127() bool`

HasMp127 returns a boolean if a field has been set.

### GetMp128

`func (o *GetContainerConfig200ResponseData) GetMp128() string`

GetMp128 returns the Mp128 field if non-nil, zero value otherwise.

### GetMp128Ok

`func (o *GetContainerConfig200ResponseData) GetMp128Ok() (*string, bool)`

GetMp128Ok returns a tuple with the Mp128 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp128

`func (o *GetContainerConfig200ResponseData) SetMp128(v string)`

SetMp128 sets Mp128 field to given value.

### HasMp128

`func (o *GetContainerConfig200ResponseData) HasMp128() bool`

HasMp128 returns a boolean if a field has been set.

### GetMp129

`func (o *GetContainerConfig200ResponseData) GetMp129() string`

GetMp129 returns the Mp129 field if non-nil, zero value otherwise.

### GetMp129Ok

`func (o *GetContainerConfig200ResponseData) GetMp129Ok() (*string, bool)`

GetMp129Ok returns a tuple with the Mp129 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp129

`func (o *GetContainerConfig200ResponseData) SetMp129(v string)`

SetMp129 sets Mp129 field to given value.

### HasMp129

`func (o *GetContainerConfig200ResponseData) HasMp129() bool`

HasMp129 returns a boolean if a field has been set.

### GetMp130

`func (o *GetContainerConfig200ResponseData) GetMp130() string`

GetMp130 returns the Mp130 field if non-nil, zero value otherwise.

### GetMp130Ok

`func (o *GetContainerConfig200ResponseData) GetMp130Ok() (*string, bool)`

GetMp130Ok returns a tuple with the Mp130 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp130

`func (o *GetContainerConfig200ResponseData) SetMp130(v string)`

SetMp130 sets Mp130 field to given value.

### HasMp130

`func (o *GetContainerConfig200ResponseData) HasMp130() bool`

HasMp130 returns a boolean if a field has been set.

### GetMp131

`func (o *GetContainerConfig200ResponseData) GetMp131() string`

GetMp131 returns the Mp131 field if non-nil, zero value otherwise.

### GetMp131Ok

`func (o *GetContainerConfig200ResponseData) GetMp131Ok() (*string, bool)`

GetMp131Ok returns a tuple with the Mp131 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp131

`func (o *GetContainerConfig200ResponseData) SetMp131(v string)`

SetMp131 sets Mp131 field to given value.

### HasMp131

`func (o *GetContainerConfig200ResponseData) HasMp131() bool`

HasMp131 returns a boolean if a field has been set.

### GetMp132

`func (o *GetContainerConfig200ResponseData) GetMp132() string`

GetMp132 returns the Mp132 field if non-nil, zero value otherwise.

### GetMp132Ok

`func (o *GetContainerConfig200ResponseData) GetMp132Ok() (*string, bool)`

GetMp132Ok returns a tuple with the Mp132 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp132

`func (o *GetContainerConfig200ResponseData) SetMp132(v string)`

SetMp132 sets Mp132 field to given value.

### HasMp132

`func (o *GetContainerConfig200ResponseData) HasMp132() bool`

HasMp132 returns a boolean if a field has been set.

### GetMp133

`func (o *GetContainerConfig200ResponseData) GetMp133() string`

GetMp133 returns the Mp133 field if non-nil, zero value otherwise.

### GetMp133Ok

`func (o *GetContainerConfig200ResponseData) GetMp133Ok() (*string, bool)`

GetMp133Ok returns a tuple with the Mp133 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp133

`func (o *GetContainerConfig200ResponseData) SetMp133(v string)`

SetMp133 sets Mp133 field to given value.

### HasMp133

`func (o *GetContainerConfig200ResponseData) HasMp133() bool`

HasMp133 returns a boolean if a field has been set.

### GetMp134

`func (o *GetContainerConfig200ResponseData) GetMp134() string`

GetMp134 returns the Mp134 field if non-nil, zero value otherwise.

### GetMp134Ok

`func (o *GetContainerConfig200ResponseData) GetMp134Ok() (*string, bool)`

GetMp134Ok returns a tuple with the Mp134 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp134

`func (o *GetContainerConfig200ResponseData) SetMp134(v string)`

SetMp134 sets Mp134 field to given value.

### HasMp134

`func (o *GetContainerConfig200ResponseData) HasMp134() bool`

HasMp134 returns a boolean if a field has been set.

### GetMp135

`func (o *GetContainerConfig200ResponseData) GetMp135() string`

GetMp135 returns the Mp135 field if non-nil, zero value otherwise.

### GetMp135Ok

`func (o *GetContainerConfig200ResponseData) GetMp135Ok() (*string, bool)`

GetMp135Ok returns a tuple with the Mp135 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp135

`func (o *GetContainerConfig200ResponseData) SetMp135(v string)`

SetMp135 sets Mp135 field to given value.

### HasMp135

`func (o *GetContainerConfig200ResponseData) HasMp135() bool`

HasMp135 returns a boolean if a field has been set.

### GetMp136

`func (o *GetContainerConfig200ResponseData) GetMp136() string`

GetMp136 returns the Mp136 field if non-nil, zero value otherwise.

### GetMp136Ok

`func (o *GetContainerConfig200ResponseData) GetMp136Ok() (*string, bool)`

GetMp136Ok returns a tuple with the Mp136 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp136

`func (o *GetContainerConfig200ResponseData) SetMp136(v string)`

SetMp136 sets Mp136 field to given value.

### HasMp136

`func (o *GetContainerConfig200ResponseData) HasMp136() bool`

HasMp136 returns a boolean if a field has been set.

### GetMp137

`func (o *GetContainerConfig200ResponseData) GetMp137() string`

GetMp137 returns the Mp137 field if non-nil, zero value otherwise.

### GetMp137Ok

`func (o *GetContainerConfig200ResponseData) GetMp137Ok() (*string, bool)`

GetMp137Ok returns a tuple with the Mp137 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp137

`func (o *GetContainerConfig200ResponseData) SetMp137(v string)`

SetMp137 sets Mp137 field to given value.

### HasMp137

`func (o *GetContainerConfig200ResponseData) HasMp137() bool`

HasMp137 returns a boolean if a field has been set.

### GetMp138

`func (o *GetContainerConfig200ResponseData) GetMp138() string`

GetMp138 returns the Mp138 field if non-nil, zero value otherwise.

### GetMp138Ok

`func (o *GetContainerConfig200ResponseData) GetMp138Ok() (*string, bool)`

GetMp138Ok returns a tuple with the Mp138 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp138

`func (o *GetContainerConfig200ResponseData) SetMp138(v string)`

SetMp138 sets Mp138 field to given value.

### HasMp138

`func (o *GetContainerConfig200ResponseData) HasMp138() bool`

HasMp138 returns a boolean if a field has been set.

### GetMp139

`func (o *GetContainerConfig200ResponseData) GetMp139() string`

GetMp139 returns the Mp139 field if non-nil, zero value otherwise.

### GetMp139Ok

`func (o *GetContainerConfig200ResponseData) GetMp139Ok() (*string, bool)`

GetMp139Ok returns a tuple with the Mp139 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp139

`func (o *GetContainerConfig200ResponseData) SetMp139(v string)`

SetMp139 sets Mp139 field to given value.

### HasMp139

`func (o *GetContainerConfig200ResponseData) HasMp139() bool`

HasMp139 returns a boolean if a field has been set.

### GetMp140

`func (o *GetContainerConfig200ResponseData) GetMp140() string`

GetMp140 returns the Mp140 field if non-nil, zero value otherwise.

### GetMp140Ok

`func (o *GetContainerConfig200ResponseData) GetMp140Ok() (*string, bool)`

GetMp140Ok returns a tuple with the Mp140 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp140

`func (o *GetContainerConfig200ResponseData) SetMp140(v string)`

SetMp140 sets Mp140 field to given value.

### HasMp140

`func (o *GetContainerConfig200ResponseData) HasMp140() bool`

HasMp140 returns a boolean if a field has been set.

### GetMp141

`func (o *GetContainerConfig200ResponseData) GetMp141() string`

GetMp141 returns the Mp141 field if non-nil, zero value otherwise.

### GetMp141Ok

`func (o *GetContainerConfig200ResponseData) GetMp141Ok() (*string, bool)`

GetMp141Ok returns a tuple with the Mp141 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp141

`func (o *GetContainerConfig200ResponseData) SetMp141(v string)`

SetMp141 sets Mp141 field to given value.

### HasMp141

`func (o *GetContainerConfig200ResponseData) HasMp141() bool`

HasMp141 returns a boolean if a field has been set.

### GetMp142

`func (o *GetContainerConfig200ResponseData) GetMp142() string`

GetMp142 returns the Mp142 field if non-nil, zero value otherwise.

### GetMp142Ok

`func (o *GetContainerConfig200ResponseData) GetMp142Ok() (*string, bool)`

GetMp142Ok returns a tuple with the Mp142 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp142

`func (o *GetContainerConfig200ResponseData) SetMp142(v string)`

SetMp142 sets Mp142 field to given value.

### HasMp142

`func (o *GetContainerConfig200ResponseData) HasMp142() bool`

HasMp142 returns a boolean if a field has been set.

### GetMp143

`func (o *GetContainerConfig200ResponseData) GetMp143() string`

GetMp143 returns the Mp143 field if non-nil, zero value otherwise.

### GetMp143Ok

`func (o *GetContainerConfig200ResponseData) GetMp143Ok() (*string, bool)`

GetMp143Ok returns a tuple with the Mp143 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp143

`func (o *GetContainerConfig200ResponseData) SetMp143(v string)`

SetMp143 sets Mp143 field to given value.

### HasMp143

`func (o *GetContainerConfig200ResponseData) HasMp143() bool`

HasMp143 returns a boolean if a field has been set.

### GetMp144

`func (o *GetContainerConfig200ResponseData) GetMp144() string`

GetMp144 returns the Mp144 field if non-nil, zero value otherwise.

### GetMp144Ok

`func (o *GetContainerConfig200ResponseData) GetMp144Ok() (*string, bool)`

GetMp144Ok returns a tuple with the Mp144 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp144

`func (o *GetContainerConfig200ResponseData) SetMp144(v string)`

SetMp144 sets Mp144 field to given value.

### HasMp144

`func (o *GetContainerConfig200ResponseData) HasMp144() bool`

HasMp144 returns a boolean if a field has been set.

### GetMp145

`func (o *GetContainerConfig200ResponseData) GetMp145() string`

GetMp145 returns the Mp145 field if non-nil, zero value otherwise.

### GetMp145Ok

`func (o *GetContainerConfig200ResponseData) GetMp145Ok() (*string, bool)`

GetMp145Ok returns a tuple with the Mp145 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp145

`func (o *GetContainerConfig200ResponseData) SetMp145(v string)`

SetMp145 sets Mp145 field to given value.

### HasMp145

`func (o *GetContainerConfig200ResponseData) HasMp145() bool`

HasMp145 returns a boolean if a field has been set.

### GetMp146

`func (o *GetContainerConfig200ResponseData) GetMp146() string`

GetMp146 returns the Mp146 field if non-nil, zero value otherwise.

### GetMp146Ok

`func (o *GetContainerConfig200ResponseData) GetMp146Ok() (*string, bool)`

GetMp146Ok returns a tuple with the Mp146 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp146

`func (o *GetContainerConfig200ResponseData) SetMp146(v string)`

SetMp146 sets Mp146 field to given value.

### HasMp146

`func (o *GetContainerConfig200ResponseData) HasMp146() bool`

HasMp146 returns a boolean if a field has been set.

### GetMp147

`func (o *GetContainerConfig200ResponseData) GetMp147() string`

GetMp147 returns the Mp147 field if non-nil, zero value otherwise.

### GetMp147Ok

`func (o *GetContainerConfig200ResponseData) GetMp147Ok() (*string, bool)`

GetMp147Ok returns a tuple with the Mp147 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp147

`func (o *GetContainerConfig200ResponseData) SetMp147(v string)`

SetMp147 sets Mp147 field to given value.

### HasMp147

`func (o *GetContainerConfig200ResponseData) HasMp147() bool`

HasMp147 returns a boolean if a field has been set.

### GetMp148

`func (o *GetContainerConfig200ResponseData) GetMp148() string`

GetMp148 returns the Mp148 field if non-nil, zero value otherwise.

### GetMp148Ok

`func (o *GetContainerConfig200ResponseData) GetMp148Ok() (*string, bool)`

GetMp148Ok returns a tuple with the Mp148 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp148

`func (o *GetContainerConfig200ResponseData) SetMp148(v string)`

SetMp148 sets Mp148 field to given value.

### HasMp148

`func (o *GetContainerConfig200ResponseData) HasMp148() bool`

HasMp148 returns a boolean if a field has been set.

### GetMp149

`func (o *GetContainerConfig200ResponseData) GetMp149() string`

GetMp149 returns the Mp149 field if non-nil, zero value otherwise.

### GetMp149Ok

`func (o *GetContainerConfig200ResponseData) GetMp149Ok() (*string, bool)`

GetMp149Ok returns a tuple with the Mp149 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp149

`func (o *GetContainerConfig200ResponseData) SetMp149(v string)`

SetMp149 sets Mp149 field to given value.

### HasMp149

`func (o *GetContainerConfig200ResponseData) HasMp149() bool`

HasMp149 returns a boolean if a field has been set.

### GetMp150

`func (o *GetContainerConfig200ResponseData) GetMp150() string`

GetMp150 returns the Mp150 field if non-nil, zero value otherwise.

### GetMp150Ok

`func (o *GetContainerConfig200ResponseData) GetMp150Ok() (*string, bool)`

GetMp150Ok returns a tuple with the Mp150 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp150

`func (o *GetContainerConfig200ResponseData) SetMp150(v string)`

SetMp150 sets Mp150 field to given value.

### HasMp150

`func (o *GetContainerConfig200ResponseData) HasMp150() bool`

HasMp150 returns a boolean if a field has been set.

### GetMp151

`func (o *GetContainerConfig200ResponseData) GetMp151() string`

GetMp151 returns the Mp151 field if non-nil, zero value otherwise.

### GetMp151Ok

`func (o *GetContainerConfig200ResponseData) GetMp151Ok() (*string, bool)`

GetMp151Ok returns a tuple with the Mp151 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp151

`func (o *GetContainerConfig200ResponseData) SetMp151(v string)`

SetMp151 sets Mp151 field to given value.

### HasMp151

`func (o *GetContainerConfig200ResponseData) HasMp151() bool`

HasMp151 returns a boolean if a field has been set.

### GetMp152

`func (o *GetContainerConfig200ResponseData) GetMp152() string`

GetMp152 returns the Mp152 field if non-nil, zero value otherwise.

### GetMp152Ok

`func (o *GetContainerConfig200ResponseData) GetMp152Ok() (*string, bool)`

GetMp152Ok returns a tuple with the Mp152 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp152

`func (o *GetContainerConfig200ResponseData) SetMp152(v string)`

SetMp152 sets Mp152 field to given value.

### HasMp152

`func (o *GetContainerConfig200ResponseData) HasMp152() bool`

HasMp152 returns a boolean if a field has been set.

### GetMp153

`func (o *GetContainerConfig200ResponseData) GetMp153() string`

GetMp153 returns the Mp153 field if non-nil, zero value otherwise.

### GetMp153Ok

`func (o *GetContainerConfig200ResponseData) GetMp153Ok() (*string, bool)`

GetMp153Ok returns a tuple with the Mp153 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp153

`func (o *GetContainerConfig200ResponseData) SetMp153(v string)`

SetMp153 sets Mp153 field to given value.

### HasMp153

`func (o *GetContainerConfig200ResponseData) HasMp153() bool`

HasMp153 returns a boolean if a field has been set.

### GetMp154

`func (o *GetContainerConfig200ResponseData) GetMp154() string`

GetMp154 returns the Mp154 field if non-nil, zero value otherwise.

### GetMp154Ok

`func (o *GetContainerConfig200ResponseData) GetMp154Ok() (*string, bool)`

GetMp154Ok returns a tuple with the Mp154 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp154

`func (o *GetContainerConfig200ResponseData) SetMp154(v string)`

SetMp154 sets Mp154 field to given value.

### HasMp154

`func (o *GetContainerConfig200ResponseData) HasMp154() bool`

HasMp154 returns a boolean if a field has been set.

### GetMp155

`func (o *GetContainerConfig200ResponseData) GetMp155() string`

GetMp155 returns the Mp155 field if non-nil, zero value otherwise.

### GetMp155Ok

`func (o *GetContainerConfig200ResponseData) GetMp155Ok() (*string, bool)`

GetMp155Ok returns a tuple with the Mp155 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp155

`func (o *GetContainerConfig200ResponseData) SetMp155(v string)`

SetMp155 sets Mp155 field to given value.

### HasMp155

`func (o *GetContainerConfig200ResponseData) HasMp155() bool`

HasMp155 returns a boolean if a field has been set.

### GetMp156

`func (o *GetContainerConfig200ResponseData) GetMp156() string`

GetMp156 returns the Mp156 field if non-nil, zero value otherwise.

### GetMp156Ok

`func (o *GetContainerConfig200ResponseData) GetMp156Ok() (*string, bool)`

GetMp156Ok returns a tuple with the Mp156 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp156

`func (o *GetContainerConfig200ResponseData) SetMp156(v string)`

SetMp156 sets Mp156 field to given value.

### HasMp156

`func (o *GetContainerConfig200ResponseData) HasMp156() bool`

HasMp156 returns a boolean if a field has been set.

### GetMp157

`func (o *GetContainerConfig200ResponseData) GetMp157() string`

GetMp157 returns the Mp157 field if non-nil, zero value otherwise.

### GetMp157Ok

`func (o *GetContainerConfig200ResponseData) GetMp157Ok() (*string, bool)`

GetMp157Ok returns a tuple with the Mp157 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp157

`func (o *GetContainerConfig200ResponseData) SetMp157(v string)`

SetMp157 sets Mp157 field to given value.

### HasMp157

`func (o *GetContainerConfig200ResponseData) HasMp157() bool`

HasMp157 returns a boolean if a field has been set.

### GetMp158

`func (o *GetContainerConfig200ResponseData) GetMp158() string`

GetMp158 returns the Mp158 field if non-nil, zero value otherwise.

### GetMp158Ok

`func (o *GetContainerConfig200ResponseData) GetMp158Ok() (*string, bool)`

GetMp158Ok returns a tuple with the Mp158 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp158

`func (o *GetContainerConfig200ResponseData) SetMp158(v string)`

SetMp158 sets Mp158 field to given value.

### HasMp158

`func (o *GetContainerConfig200ResponseData) HasMp158() bool`

HasMp158 returns a boolean if a field has been set.

### GetMp159

`func (o *GetContainerConfig200ResponseData) GetMp159() string`

GetMp159 returns the Mp159 field if non-nil, zero value otherwise.

### GetMp159Ok

`func (o *GetContainerConfig200ResponseData) GetMp159Ok() (*string, bool)`

GetMp159Ok returns a tuple with the Mp159 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp159

`func (o *GetContainerConfig200ResponseData) SetMp159(v string)`

SetMp159 sets Mp159 field to given value.

### HasMp159

`func (o *GetContainerConfig200ResponseData) HasMp159() bool`

HasMp159 returns a boolean if a field has been set.

### GetMp160

`func (o *GetContainerConfig200ResponseData) GetMp160() string`

GetMp160 returns the Mp160 field if non-nil, zero value otherwise.

### GetMp160Ok

`func (o *GetContainerConfig200ResponseData) GetMp160Ok() (*string, bool)`

GetMp160Ok returns a tuple with the Mp160 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp160

`func (o *GetContainerConfig200ResponseData) SetMp160(v string)`

SetMp160 sets Mp160 field to given value.

### HasMp160

`func (o *GetContainerConfig200ResponseData) HasMp160() bool`

HasMp160 returns a boolean if a field has been set.

### GetMp161

`func (o *GetContainerConfig200ResponseData) GetMp161() string`

GetMp161 returns the Mp161 field if non-nil, zero value otherwise.

### GetMp161Ok

`func (o *GetContainerConfig200ResponseData) GetMp161Ok() (*string, bool)`

GetMp161Ok returns a tuple with the Mp161 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp161

`func (o *GetContainerConfig200ResponseData) SetMp161(v string)`

SetMp161 sets Mp161 field to given value.

### HasMp161

`func (o *GetContainerConfig200ResponseData) HasMp161() bool`

HasMp161 returns a boolean if a field has been set.

### GetMp162

`func (o *GetContainerConfig200ResponseData) GetMp162() string`

GetMp162 returns the Mp162 field if non-nil, zero value otherwise.

### GetMp162Ok

`func (o *GetContainerConfig200ResponseData) GetMp162Ok() (*string, bool)`

GetMp162Ok returns a tuple with the Mp162 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp162

`func (o *GetContainerConfig200ResponseData) SetMp162(v string)`

SetMp162 sets Mp162 field to given value.

### HasMp162

`func (o *GetContainerConfig200ResponseData) HasMp162() bool`

HasMp162 returns a boolean if a field has been set.

### GetMp163

`func (o *GetContainerConfig200ResponseData) GetMp163() string`

GetMp163 returns the Mp163 field if non-nil, zero value otherwise.

### GetMp163Ok

`func (o *GetContainerConfig200ResponseData) GetMp163Ok() (*string, bool)`

GetMp163Ok returns a tuple with the Mp163 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp163

`func (o *GetContainerConfig200ResponseData) SetMp163(v string)`

SetMp163 sets Mp163 field to given value.

### HasMp163

`func (o *GetContainerConfig200ResponseData) HasMp163() bool`

HasMp163 returns a boolean if a field has been set.

### GetMp164

`func (o *GetContainerConfig200ResponseData) GetMp164() string`

GetMp164 returns the Mp164 field if non-nil, zero value otherwise.

### GetMp164Ok

`func (o *GetContainerConfig200ResponseData) GetMp164Ok() (*string, bool)`

GetMp164Ok returns a tuple with the Mp164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp164

`func (o *GetContainerConfig200ResponseData) SetMp164(v string)`

SetMp164 sets Mp164 field to given value.

### HasMp164

`func (o *GetContainerConfig200ResponseData) HasMp164() bool`

HasMp164 returns a boolean if a field has been set.

### GetMp165

`func (o *GetContainerConfig200ResponseData) GetMp165() string`

GetMp165 returns the Mp165 field if non-nil, zero value otherwise.

### GetMp165Ok

`func (o *GetContainerConfig200ResponseData) GetMp165Ok() (*string, bool)`

GetMp165Ok returns a tuple with the Mp165 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp165

`func (o *GetContainerConfig200ResponseData) SetMp165(v string)`

SetMp165 sets Mp165 field to given value.

### HasMp165

`func (o *GetContainerConfig200ResponseData) HasMp165() bool`

HasMp165 returns a boolean if a field has been set.

### GetMp166

`func (o *GetContainerConfig200ResponseData) GetMp166() string`

GetMp166 returns the Mp166 field if non-nil, zero value otherwise.

### GetMp166Ok

`func (o *GetContainerConfig200ResponseData) GetMp166Ok() (*string, bool)`

GetMp166Ok returns a tuple with the Mp166 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp166

`func (o *GetContainerConfig200ResponseData) SetMp166(v string)`

SetMp166 sets Mp166 field to given value.

### HasMp166

`func (o *GetContainerConfig200ResponseData) HasMp166() bool`

HasMp166 returns a boolean if a field has been set.

### GetMp167

`func (o *GetContainerConfig200ResponseData) GetMp167() string`

GetMp167 returns the Mp167 field if non-nil, zero value otherwise.

### GetMp167Ok

`func (o *GetContainerConfig200ResponseData) GetMp167Ok() (*string, bool)`

GetMp167Ok returns a tuple with the Mp167 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp167

`func (o *GetContainerConfig200ResponseData) SetMp167(v string)`

SetMp167 sets Mp167 field to given value.

### HasMp167

`func (o *GetContainerConfig200ResponseData) HasMp167() bool`

HasMp167 returns a boolean if a field has been set.

### GetMp168

`func (o *GetContainerConfig200ResponseData) GetMp168() string`

GetMp168 returns the Mp168 field if non-nil, zero value otherwise.

### GetMp168Ok

`func (o *GetContainerConfig200ResponseData) GetMp168Ok() (*string, bool)`

GetMp168Ok returns a tuple with the Mp168 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp168

`func (o *GetContainerConfig200ResponseData) SetMp168(v string)`

SetMp168 sets Mp168 field to given value.

### HasMp168

`func (o *GetContainerConfig200ResponseData) HasMp168() bool`

HasMp168 returns a boolean if a field has been set.

### GetMp169

`func (o *GetContainerConfig200ResponseData) GetMp169() string`

GetMp169 returns the Mp169 field if non-nil, zero value otherwise.

### GetMp169Ok

`func (o *GetContainerConfig200ResponseData) GetMp169Ok() (*string, bool)`

GetMp169Ok returns a tuple with the Mp169 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp169

`func (o *GetContainerConfig200ResponseData) SetMp169(v string)`

SetMp169 sets Mp169 field to given value.

### HasMp169

`func (o *GetContainerConfig200ResponseData) HasMp169() bool`

HasMp169 returns a boolean if a field has been set.

### GetMp170

`func (o *GetContainerConfig200ResponseData) GetMp170() string`

GetMp170 returns the Mp170 field if non-nil, zero value otherwise.

### GetMp170Ok

`func (o *GetContainerConfig200ResponseData) GetMp170Ok() (*string, bool)`

GetMp170Ok returns a tuple with the Mp170 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp170

`func (o *GetContainerConfig200ResponseData) SetMp170(v string)`

SetMp170 sets Mp170 field to given value.

### HasMp170

`func (o *GetContainerConfig200ResponseData) HasMp170() bool`

HasMp170 returns a boolean if a field has been set.

### GetMp171

`func (o *GetContainerConfig200ResponseData) GetMp171() string`

GetMp171 returns the Mp171 field if non-nil, zero value otherwise.

### GetMp171Ok

`func (o *GetContainerConfig200ResponseData) GetMp171Ok() (*string, bool)`

GetMp171Ok returns a tuple with the Mp171 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp171

`func (o *GetContainerConfig200ResponseData) SetMp171(v string)`

SetMp171 sets Mp171 field to given value.

### HasMp171

`func (o *GetContainerConfig200ResponseData) HasMp171() bool`

HasMp171 returns a boolean if a field has been set.

### GetMp172

`func (o *GetContainerConfig200ResponseData) GetMp172() string`

GetMp172 returns the Mp172 field if non-nil, zero value otherwise.

### GetMp172Ok

`func (o *GetContainerConfig200ResponseData) GetMp172Ok() (*string, bool)`

GetMp172Ok returns a tuple with the Mp172 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp172

`func (o *GetContainerConfig200ResponseData) SetMp172(v string)`

SetMp172 sets Mp172 field to given value.

### HasMp172

`func (o *GetContainerConfig200ResponseData) HasMp172() bool`

HasMp172 returns a boolean if a field has been set.

### GetMp173

`func (o *GetContainerConfig200ResponseData) GetMp173() string`

GetMp173 returns the Mp173 field if non-nil, zero value otherwise.

### GetMp173Ok

`func (o *GetContainerConfig200ResponseData) GetMp173Ok() (*string, bool)`

GetMp173Ok returns a tuple with the Mp173 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp173

`func (o *GetContainerConfig200ResponseData) SetMp173(v string)`

SetMp173 sets Mp173 field to given value.

### HasMp173

`func (o *GetContainerConfig200ResponseData) HasMp173() bool`

HasMp173 returns a boolean if a field has been set.

### GetMp174

`func (o *GetContainerConfig200ResponseData) GetMp174() string`

GetMp174 returns the Mp174 field if non-nil, zero value otherwise.

### GetMp174Ok

`func (o *GetContainerConfig200ResponseData) GetMp174Ok() (*string, bool)`

GetMp174Ok returns a tuple with the Mp174 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp174

`func (o *GetContainerConfig200ResponseData) SetMp174(v string)`

SetMp174 sets Mp174 field to given value.

### HasMp174

`func (o *GetContainerConfig200ResponseData) HasMp174() bool`

HasMp174 returns a boolean if a field has been set.

### GetMp175

`func (o *GetContainerConfig200ResponseData) GetMp175() string`

GetMp175 returns the Mp175 field if non-nil, zero value otherwise.

### GetMp175Ok

`func (o *GetContainerConfig200ResponseData) GetMp175Ok() (*string, bool)`

GetMp175Ok returns a tuple with the Mp175 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp175

`func (o *GetContainerConfig200ResponseData) SetMp175(v string)`

SetMp175 sets Mp175 field to given value.

### HasMp175

`func (o *GetContainerConfig200ResponseData) HasMp175() bool`

HasMp175 returns a boolean if a field has been set.

### GetMp176

`func (o *GetContainerConfig200ResponseData) GetMp176() string`

GetMp176 returns the Mp176 field if non-nil, zero value otherwise.

### GetMp176Ok

`func (o *GetContainerConfig200ResponseData) GetMp176Ok() (*string, bool)`

GetMp176Ok returns a tuple with the Mp176 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp176

`func (o *GetContainerConfig200ResponseData) SetMp176(v string)`

SetMp176 sets Mp176 field to given value.

### HasMp176

`func (o *GetContainerConfig200ResponseData) HasMp176() bool`

HasMp176 returns a boolean if a field has been set.

### GetMp177

`func (o *GetContainerConfig200ResponseData) GetMp177() string`

GetMp177 returns the Mp177 field if non-nil, zero value otherwise.

### GetMp177Ok

`func (o *GetContainerConfig200ResponseData) GetMp177Ok() (*string, bool)`

GetMp177Ok returns a tuple with the Mp177 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp177

`func (o *GetContainerConfig200ResponseData) SetMp177(v string)`

SetMp177 sets Mp177 field to given value.

### HasMp177

`func (o *GetContainerConfig200ResponseData) HasMp177() bool`

HasMp177 returns a boolean if a field has been set.

### GetMp178

`func (o *GetContainerConfig200ResponseData) GetMp178() string`

GetMp178 returns the Mp178 field if non-nil, zero value otherwise.

### GetMp178Ok

`func (o *GetContainerConfig200ResponseData) GetMp178Ok() (*string, bool)`

GetMp178Ok returns a tuple with the Mp178 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp178

`func (o *GetContainerConfig200ResponseData) SetMp178(v string)`

SetMp178 sets Mp178 field to given value.

### HasMp178

`func (o *GetContainerConfig200ResponseData) HasMp178() bool`

HasMp178 returns a boolean if a field has been set.

### GetMp179

`func (o *GetContainerConfig200ResponseData) GetMp179() string`

GetMp179 returns the Mp179 field if non-nil, zero value otherwise.

### GetMp179Ok

`func (o *GetContainerConfig200ResponseData) GetMp179Ok() (*string, bool)`

GetMp179Ok returns a tuple with the Mp179 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp179

`func (o *GetContainerConfig200ResponseData) SetMp179(v string)`

SetMp179 sets Mp179 field to given value.

### HasMp179

`func (o *GetContainerConfig200ResponseData) HasMp179() bool`

HasMp179 returns a boolean if a field has been set.

### GetMp180

`func (o *GetContainerConfig200ResponseData) GetMp180() string`

GetMp180 returns the Mp180 field if non-nil, zero value otherwise.

### GetMp180Ok

`func (o *GetContainerConfig200ResponseData) GetMp180Ok() (*string, bool)`

GetMp180Ok returns a tuple with the Mp180 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp180

`func (o *GetContainerConfig200ResponseData) SetMp180(v string)`

SetMp180 sets Mp180 field to given value.

### HasMp180

`func (o *GetContainerConfig200ResponseData) HasMp180() bool`

HasMp180 returns a boolean if a field has been set.

### GetMp181

`func (o *GetContainerConfig200ResponseData) GetMp181() string`

GetMp181 returns the Mp181 field if non-nil, zero value otherwise.

### GetMp181Ok

`func (o *GetContainerConfig200ResponseData) GetMp181Ok() (*string, bool)`

GetMp181Ok returns a tuple with the Mp181 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp181

`func (o *GetContainerConfig200ResponseData) SetMp181(v string)`

SetMp181 sets Mp181 field to given value.

### HasMp181

`func (o *GetContainerConfig200ResponseData) HasMp181() bool`

HasMp181 returns a boolean if a field has been set.

### GetMp182

`func (o *GetContainerConfig200ResponseData) GetMp182() string`

GetMp182 returns the Mp182 field if non-nil, zero value otherwise.

### GetMp182Ok

`func (o *GetContainerConfig200ResponseData) GetMp182Ok() (*string, bool)`

GetMp182Ok returns a tuple with the Mp182 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp182

`func (o *GetContainerConfig200ResponseData) SetMp182(v string)`

SetMp182 sets Mp182 field to given value.

### HasMp182

`func (o *GetContainerConfig200ResponseData) HasMp182() bool`

HasMp182 returns a boolean if a field has been set.

### GetMp183

`func (o *GetContainerConfig200ResponseData) GetMp183() string`

GetMp183 returns the Mp183 field if non-nil, zero value otherwise.

### GetMp183Ok

`func (o *GetContainerConfig200ResponseData) GetMp183Ok() (*string, bool)`

GetMp183Ok returns a tuple with the Mp183 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp183

`func (o *GetContainerConfig200ResponseData) SetMp183(v string)`

SetMp183 sets Mp183 field to given value.

### HasMp183

`func (o *GetContainerConfig200ResponseData) HasMp183() bool`

HasMp183 returns a boolean if a field has been set.

### GetMp184

`func (o *GetContainerConfig200ResponseData) GetMp184() string`

GetMp184 returns the Mp184 field if non-nil, zero value otherwise.

### GetMp184Ok

`func (o *GetContainerConfig200ResponseData) GetMp184Ok() (*string, bool)`

GetMp184Ok returns a tuple with the Mp184 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp184

`func (o *GetContainerConfig200ResponseData) SetMp184(v string)`

SetMp184 sets Mp184 field to given value.

### HasMp184

`func (o *GetContainerConfig200ResponseData) HasMp184() bool`

HasMp184 returns a boolean if a field has been set.

### GetMp185

`func (o *GetContainerConfig200ResponseData) GetMp185() string`

GetMp185 returns the Mp185 field if non-nil, zero value otherwise.

### GetMp185Ok

`func (o *GetContainerConfig200ResponseData) GetMp185Ok() (*string, bool)`

GetMp185Ok returns a tuple with the Mp185 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp185

`func (o *GetContainerConfig200ResponseData) SetMp185(v string)`

SetMp185 sets Mp185 field to given value.

### HasMp185

`func (o *GetContainerConfig200ResponseData) HasMp185() bool`

HasMp185 returns a boolean if a field has been set.

### GetMp186

`func (o *GetContainerConfig200ResponseData) GetMp186() string`

GetMp186 returns the Mp186 field if non-nil, zero value otherwise.

### GetMp186Ok

`func (o *GetContainerConfig200ResponseData) GetMp186Ok() (*string, bool)`

GetMp186Ok returns a tuple with the Mp186 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp186

`func (o *GetContainerConfig200ResponseData) SetMp186(v string)`

SetMp186 sets Mp186 field to given value.

### HasMp186

`func (o *GetContainerConfig200ResponseData) HasMp186() bool`

HasMp186 returns a boolean if a field has been set.

### GetMp187

`func (o *GetContainerConfig200ResponseData) GetMp187() string`

GetMp187 returns the Mp187 field if non-nil, zero value otherwise.

### GetMp187Ok

`func (o *GetContainerConfig200ResponseData) GetMp187Ok() (*string, bool)`

GetMp187Ok returns a tuple with the Mp187 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp187

`func (o *GetContainerConfig200ResponseData) SetMp187(v string)`

SetMp187 sets Mp187 field to given value.

### HasMp187

`func (o *GetContainerConfig200ResponseData) HasMp187() bool`

HasMp187 returns a boolean if a field has been set.

### GetMp188

`func (o *GetContainerConfig200ResponseData) GetMp188() string`

GetMp188 returns the Mp188 field if non-nil, zero value otherwise.

### GetMp188Ok

`func (o *GetContainerConfig200ResponseData) GetMp188Ok() (*string, bool)`

GetMp188Ok returns a tuple with the Mp188 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp188

`func (o *GetContainerConfig200ResponseData) SetMp188(v string)`

SetMp188 sets Mp188 field to given value.

### HasMp188

`func (o *GetContainerConfig200ResponseData) HasMp188() bool`

HasMp188 returns a boolean if a field has been set.

### GetMp189

`func (o *GetContainerConfig200ResponseData) GetMp189() string`

GetMp189 returns the Mp189 field if non-nil, zero value otherwise.

### GetMp189Ok

`func (o *GetContainerConfig200ResponseData) GetMp189Ok() (*string, bool)`

GetMp189Ok returns a tuple with the Mp189 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp189

`func (o *GetContainerConfig200ResponseData) SetMp189(v string)`

SetMp189 sets Mp189 field to given value.

### HasMp189

`func (o *GetContainerConfig200ResponseData) HasMp189() bool`

HasMp189 returns a boolean if a field has been set.

### GetMp190

`func (o *GetContainerConfig200ResponseData) GetMp190() string`

GetMp190 returns the Mp190 field if non-nil, zero value otherwise.

### GetMp190Ok

`func (o *GetContainerConfig200ResponseData) GetMp190Ok() (*string, bool)`

GetMp190Ok returns a tuple with the Mp190 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp190

`func (o *GetContainerConfig200ResponseData) SetMp190(v string)`

SetMp190 sets Mp190 field to given value.

### HasMp190

`func (o *GetContainerConfig200ResponseData) HasMp190() bool`

HasMp190 returns a boolean if a field has been set.

### GetMp191

`func (o *GetContainerConfig200ResponseData) GetMp191() string`

GetMp191 returns the Mp191 field if non-nil, zero value otherwise.

### GetMp191Ok

`func (o *GetContainerConfig200ResponseData) GetMp191Ok() (*string, bool)`

GetMp191Ok returns a tuple with the Mp191 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp191

`func (o *GetContainerConfig200ResponseData) SetMp191(v string)`

SetMp191 sets Mp191 field to given value.

### HasMp191

`func (o *GetContainerConfig200ResponseData) HasMp191() bool`

HasMp191 returns a boolean if a field has been set.

### GetMp192

`func (o *GetContainerConfig200ResponseData) GetMp192() string`

GetMp192 returns the Mp192 field if non-nil, zero value otherwise.

### GetMp192Ok

`func (o *GetContainerConfig200ResponseData) GetMp192Ok() (*string, bool)`

GetMp192Ok returns a tuple with the Mp192 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp192

`func (o *GetContainerConfig200ResponseData) SetMp192(v string)`

SetMp192 sets Mp192 field to given value.

### HasMp192

`func (o *GetContainerConfig200ResponseData) HasMp192() bool`

HasMp192 returns a boolean if a field has been set.

### GetMp193

`func (o *GetContainerConfig200ResponseData) GetMp193() string`

GetMp193 returns the Mp193 field if non-nil, zero value otherwise.

### GetMp193Ok

`func (o *GetContainerConfig200ResponseData) GetMp193Ok() (*string, bool)`

GetMp193Ok returns a tuple with the Mp193 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp193

`func (o *GetContainerConfig200ResponseData) SetMp193(v string)`

SetMp193 sets Mp193 field to given value.

### HasMp193

`func (o *GetContainerConfig200ResponseData) HasMp193() bool`

HasMp193 returns a boolean if a field has been set.

### GetMp194

`func (o *GetContainerConfig200ResponseData) GetMp194() string`

GetMp194 returns the Mp194 field if non-nil, zero value otherwise.

### GetMp194Ok

`func (o *GetContainerConfig200ResponseData) GetMp194Ok() (*string, bool)`

GetMp194Ok returns a tuple with the Mp194 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp194

`func (o *GetContainerConfig200ResponseData) SetMp194(v string)`

SetMp194 sets Mp194 field to given value.

### HasMp194

`func (o *GetContainerConfig200ResponseData) HasMp194() bool`

HasMp194 returns a boolean if a field has been set.

### GetMp195

`func (o *GetContainerConfig200ResponseData) GetMp195() string`

GetMp195 returns the Mp195 field if non-nil, zero value otherwise.

### GetMp195Ok

`func (o *GetContainerConfig200ResponseData) GetMp195Ok() (*string, bool)`

GetMp195Ok returns a tuple with the Mp195 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp195

`func (o *GetContainerConfig200ResponseData) SetMp195(v string)`

SetMp195 sets Mp195 field to given value.

### HasMp195

`func (o *GetContainerConfig200ResponseData) HasMp195() bool`

HasMp195 returns a boolean if a field has been set.

### GetMp196

`func (o *GetContainerConfig200ResponseData) GetMp196() string`

GetMp196 returns the Mp196 field if non-nil, zero value otherwise.

### GetMp196Ok

`func (o *GetContainerConfig200ResponseData) GetMp196Ok() (*string, bool)`

GetMp196Ok returns a tuple with the Mp196 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp196

`func (o *GetContainerConfig200ResponseData) SetMp196(v string)`

SetMp196 sets Mp196 field to given value.

### HasMp196

`func (o *GetContainerConfig200ResponseData) HasMp196() bool`

HasMp196 returns a boolean if a field has been set.

### GetMp197

`func (o *GetContainerConfig200ResponseData) GetMp197() string`

GetMp197 returns the Mp197 field if non-nil, zero value otherwise.

### GetMp197Ok

`func (o *GetContainerConfig200ResponseData) GetMp197Ok() (*string, bool)`

GetMp197Ok returns a tuple with the Mp197 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp197

`func (o *GetContainerConfig200ResponseData) SetMp197(v string)`

SetMp197 sets Mp197 field to given value.

### HasMp197

`func (o *GetContainerConfig200ResponseData) HasMp197() bool`

HasMp197 returns a boolean if a field has been set.

### GetMp198

`func (o *GetContainerConfig200ResponseData) GetMp198() string`

GetMp198 returns the Mp198 field if non-nil, zero value otherwise.

### GetMp198Ok

`func (o *GetContainerConfig200ResponseData) GetMp198Ok() (*string, bool)`

GetMp198Ok returns a tuple with the Mp198 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp198

`func (o *GetContainerConfig200ResponseData) SetMp198(v string)`

SetMp198 sets Mp198 field to given value.

### HasMp198

`func (o *GetContainerConfig200ResponseData) HasMp198() bool`

HasMp198 returns a boolean if a field has been set.

### GetMp199

`func (o *GetContainerConfig200ResponseData) GetMp199() string`

GetMp199 returns the Mp199 field if non-nil, zero value otherwise.

### GetMp199Ok

`func (o *GetContainerConfig200ResponseData) GetMp199Ok() (*string, bool)`

GetMp199Ok returns a tuple with the Mp199 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp199

`func (o *GetContainerConfig200ResponseData) SetMp199(v string)`

SetMp199 sets Mp199 field to given value.

### HasMp199

`func (o *GetContainerConfig200ResponseData) HasMp199() bool`

HasMp199 returns a boolean if a field has been set.

### GetMp200

`func (o *GetContainerConfig200ResponseData) GetMp200() string`

GetMp200 returns the Mp200 field if non-nil, zero value otherwise.

### GetMp200Ok

`func (o *GetContainerConfig200ResponseData) GetMp200Ok() (*string, bool)`

GetMp200Ok returns a tuple with the Mp200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp200

`func (o *GetContainerConfig200ResponseData) SetMp200(v string)`

SetMp200 sets Mp200 field to given value.

### HasMp200

`func (o *GetContainerConfig200ResponseData) HasMp200() bool`

HasMp200 returns a boolean if a field has been set.

### GetMp201

`func (o *GetContainerConfig200ResponseData) GetMp201() string`

GetMp201 returns the Mp201 field if non-nil, zero value otherwise.

### GetMp201Ok

`func (o *GetContainerConfig200ResponseData) GetMp201Ok() (*string, bool)`

GetMp201Ok returns a tuple with the Mp201 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp201

`func (o *GetContainerConfig200ResponseData) SetMp201(v string)`

SetMp201 sets Mp201 field to given value.

### HasMp201

`func (o *GetContainerConfig200ResponseData) HasMp201() bool`

HasMp201 returns a boolean if a field has been set.

### GetMp202

`func (o *GetContainerConfig200ResponseData) GetMp202() string`

GetMp202 returns the Mp202 field if non-nil, zero value otherwise.

### GetMp202Ok

`func (o *GetContainerConfig200ResponseData) GetMp202Ok() (*string, bool)`

GetMp202Ok returns a tuple with the Mp202 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp202

`func (o *GetContainerConfig200ResponseData) SetMp202(v string)`

SetMp202 sets Mp202 field to given value.

### HasMp202

`func (o *GetContainerConfig200ResponseData) HasMp202() bool`

HasMp202 returns a boolean if a field has been set.

### GetMp203

`func (o *GetContainerConfig200ResponseData) GetMp203() string`

GetMp203 returns the Mp203 field if non-nil, zero value otherwise.

### GetMp203Ok

`func (o *GetContainerConfig200ResponseData) GetMp203Ok() (*string, bool)`

GetMp203Ok returns a tuple with the Mp203 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp203

`func (o *GetContainerConfig200ResponseData) SetMp203(v string)`

SetMp203 sets Mp203 field to given value.

### HasMp203

`func (o *GetContainerConfig200ResponseData) HasMp203() bool`

HasMp203 returns a boolean if a field has been set.

### GetMp204

`func (o *GetContainerConfig200ResponseData) GetMp204() string`

GetMp204 returns the Mp204 field if non-nil, zero value otherwise.

### GetMp204Ok

`func (o *GetContainerConfig200ResponseData) GetMp204Ok() (*string, bool)`

GetMp204Ok returns a tuple with the Mp204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp204

`func (o *GetContainerConfig200ResponseData) SetMp204(v string)`

SetMp204 sets Mp204 field to given value.

### HasMp204

`func (o *GetContainerConfig200ResponseData) HasMp204() bool`

HasMp204 returns a boolean if a field has been set.

### GetMp205

`func (o *GetContainerConfig200ResponseData) GetMp205() string`

GetMp205 returns the Mp205 field if non-nil, zero value otherwise.

### GetMp205Ok

`func (o *GetContainerConfig200ResponseData) GetMp205Ok() (*string, bool)`

GetMp205Ok returns a tuple with the Mp205 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp205

`func (o *GetContainerConfig200ResponseData) SetMp205(v string)`

SetMp205 sets Mp205 field to given value.

### HasMp205

`func (o *GetContainerConfig200ResponseData) HasMp205() bool`

HasMp205 returns a boolean if a field has been set.

### GetMp206

`func (o *GetContainerConfig200ResponseData) GetMp206() string`

GetMp206 returns the Mp206 field if non-nil, zero value otherwise.

### GetMp206Ok

`func (o *GetContainerConfig200ResponseData) GetMp206Ok() (*string, bool)`

GetMp206Ok returns a tuple with the Mp206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp206

`func (o *GetContainerConfig200ResponseData) SetMp206(v string)`

SetMp206 sets Mp206 field to given value.

### HasMp206

`func (o *GetContainerConfig200ResponseData) HasMp206() bool`

HasMp206 returns a boolean if a field has been set.

### GetMp207

`func (o *GetContainerConfig200ResponseData) GetMp207() string`

GetMp207 returns the Mp207 field if non-nil, zero value otherwise.

### GetMp207Ok

`func (o *GetContainerConfig200ResponseData) GetMp207Ok() (*string, bool)`

GetMp207Ok returns a tuple with the Mp207 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp207

`func (o *GetContainerConfig200ResponseData) SetMp207(v string)`

SetMp207 sets Mp207 field to given value.

### HasMp207

`func (o *GetContainerConfig200ResponseData) HasMp207() bool`

HasMp207 returns a boolean if a field has been set.

### GetMp208

`func (o *GetContainerConfig200ResponseData) GetMp208() string`

GetMp208 returns the Mp208 field if non-nil, zero value otherwise.

### GetMp208Ok

`func (o *GetContainerConfig200ResponseData) GetMp208Ok() (*string, bool)`

GetMp208Ok returns a tuple with the Mp208 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp208

`func (o *GetContainerConfig200ResponseData) SetMp208(v string)`

SetMp208 sets Mp208 field to given value.

### HasMp208

`func (o *GetContainerConfig200ResponseData) HasMp208() bool`

HasMp208 returns a boolean if a field has been set.

### GetMp209

`func (o *GetContainerConfig200ResponseData) GetMp209() string`

GetMp209 returns the Mp209 field if non-nil, zero value otherwise.

### GetMp209Ok

`func (o *GetContainerConfig200ResponseData) GetMp209Ok() (*string, bool)`

GetMp209Ok returns a tuple with the Mp209 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp209

`func (o *GetContainerConfig200ResponseData) SetMp209(v string)`

SetMp209 sets Mp209 field to given value.

### HasMp209

`func (o *GetContainerConfig200ResponseData) HasMp209() bool`

HasMp209 returns a boolean if a field has been set.

### GetMp210

`func (o *GetContainerConfig200ResponseData) GetMp210() string`

GetMp210 returns the Mp210 field if non-nil, zero value otherwise.

### GetMp210Ok

`func (o *GetContainerConfig200ResponseData) GetMp210Ok() (*string, bool)`

GetMp210Ok returns a tuple with the Mp210 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp210

`func (o *GetContainerConfig200ResponseData) SetMp210(v string)`

SetMp210 sets Mp210 field to given value.

### HasMp210

`func (o *GetContainerConfig200ResponseData) HasMp210() bool`

HasMp210 returns a boolean if a field has been set.

### GetMp211

`func (o *GetContainerConfig200ResponseData) GetMp211() string`

GetMp211 returns the Mp211 field if non-nil, zero value otherwise.

### GetMp211Ok

`func (o *GetContainerConfig200ResponseData) GetMp211Ok() (*string, bool)`

GetMp211Ok returns a tuple with the Mp211 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp211

`func (o *GetContainerConfig200ResponseData) SetMp211(v string)`

SetMp211 sets Mp211 field to given value.

### HasMp211

`func (o *GetContainerConfig200ResponseData) HasMp211() bool`

HasMp211 returns a boolean if a field has been set.

### GetMp212

`func (o *GetContainerConfig200ResponseData) GetMp212() string`

GetMp212 returns the Mp212 field if non-nil, zero value otherwise.

### GetMp212Ok

`func (o *GetContainerConfig200ResponseData) GetMp212Ok() (*string, bool)`

GetMp212Ok returns a tuple with the Mp212 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp212

`func (o *GetContainerConfig200ResponseData) SetMp212(v string)`

SetMp212 sets Mp212 field to given value.

### HasMp212

`func (o *GetContainerConfig200ResponseData) HasMp212() bool`

HasMp212 returns a boolean if a field has been set.

### GetMp213

`func (o *GetContainerConfig200ResponseData) GetMp213() string`

GetMp213 returns the Mp213 field if non-nil, zero value otherwise.

### GetMp213Ok

`func (o *GetContainerConfig200ResponseData) GetMp213Ok() (*string, bool)`

GetMp213Ok returns a tuple with the Mp213 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp213

`func (o *GetContainerConfig200ResponseData) SetMp213(v string)`

SetMp213 sets Mp213 field to given value.

### HasMp213

`func (o *GetContainerConfig200ResponseData) HasMp213() bool`

HasMp213 returns a boolean if a field has been set.

### GetMp214

`func (o *GetContainerConfig200ResponseData) GetMp214() string`

GetMp214 returns the Mp214 field if non-nil, zero value otherwise.

### GetMp214Ok

`func (o *GetContainerConfig200ResponseData) GetMp214Ok() (*string, bool)`

GetMp214Ok returns a tuple with the Mp214 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp214

`func (o *GetContainerConfig200ResponseData) SetMp214(v string)`

SetMp214 sets Mp214 field to given value.

### HasMp214

`func (o *GetContainerConfig200ResponseData) HasMp214() bool`

HasMp214 returns a boolean if a field has been set.

### GetMp215

`func (o *GetContainerConfig200ResponseData) GetMp215() string`

GetMp215 returns the Mp215 field if non-nil, zero value otherwise.

### GetMp215Ok

`func (o *GetContainerConfig200ResponseData) GetMp215Ok() (*string, bool)`

GetMp215Ok returns a tuple with the Mp215 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp215

`func (o *GetContainerConfig200ResponseData) SetMp215(v string)`

SetMp215 sets Mp215 field to given value.

### HasMp215

`func (o *GetContainerConfig200ResponseData) HasMp215() bool`

HasMp215 returns a boolean if a field has been set.

### GetMp216

`func (o *GetContainerConfig200ResponseData) GetMp216() string`

GetMp216 returns the Mp216 field if non-nil, zero value otherwise.

### GetMp216Ok

`func (o *GetContainerConfig200ResponseData) GetMp216Ok() (*string, bool)`

GetMp216Ok returns a tuple with the Mp216 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp216

`func (o *GetContainerConfig200ResponseData) SetMp216(v string)`

SetMp216 sets Mp216 field to given value.

### HasMp216

`func (o *GetContainerConfig200ResponseData) HasMp216() bool`

HasMp216 returns a boolean if a field has been set.

### GetMp217

`func (o *GetContainerConfig200ResponseData) GetMp217() string`

GetMp217 returns the Mp217 field if non-nil, zero value otherwise.

### GetMp217Ok

`func (o *GetContainerConfig200ResponseData) GetMp217Ok() (*string, bool)`

GetMp217Ok returns a tuple with the Mp217 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp217

`func (o *GetContainerConfig200ResponseData) SetMp217(v string)`

SetMp217 sets Mp217 field to given value.

### HasMp217

`func (o *GetContainerConfig200ResponseData) HasMp217() bool`

HasMp217 returns a boolean if a field has been set.

### GetMp218

`func (o *GetContainerConfig200ResponseData) GetMp218() string`

GetMp218 returns the Mp218 field if non-nil, zero value otherwise.

### GetMp218Ok

`func (o *GetContainerConfig200ResponseData) GetMp218Ok() (*string, bool)`

GetMp218Ok returns a tuple with the Mp218 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp218

`func (o *GetContainerConfig200ResponseData) SetMp218(v string)`

SetMp218 sets Mp218 field to given value.

### HasMp218

`func (o *GetContainerConfig200ResponseData) HasMp218() bool`

HasMp218 returns a boolean if a field has been set.

### GetMp219

`func (o *GetContainerConfig200ResponseData) GetMp219() string`

GetMp219 returns the Mp219 field if non-nil, zero value otherwise.

### GetMp219Ok

`func (o *GetContainerConfig200ResponseData) GetMp219Ok() (*string, bool)`

GetMp219Ok returns a tuple with the Mp219 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp219

`func (o *GetContainerConfig200ResponseData) SetMp219(v string)`

SetMp219 sets Mp219 field to given value.

### HasMp219

`func (o *GetContainerConfig200ResponseData) HasMp219() bool`

HasMp219 returns a boolean if a field has been set.

### GetMp220

`func (o *GetContainerConfig200ResponseData) GetMp220() string`

GetMp220 returns the Mp220 field if non-nil, zero value otherwise.

### GetMp220Ok

`func (o *GetContainerConfig200ResponseData) GetMp220Ok() (*string, bool)`

GetMp220Ok returns a tuple with the Mp220 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp220

`func (o *GetContainerConfig200ResponseData) SetMp220(v string)`

SetMp220 sets Mp220 field to given value.

### HasMp220

`func (o *GetContainerConfig200ResponseData) HasMp220() bool`

HasMp220 returns a boolean if a field has been set.

### GetMp221

`func (o *GetContainerConfig200ResponseData) GetMp221() string`

GetMp221 returns the Mp221 field if non-nil, zero value otherwise.

### GetMp221Ok

`func (o *GetContainerConfig200ResponseData) GetMp221Ok() (*string, bool)`

GetMp221Ok returns a tuple with the Mp221 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp221

`func (o *GetContainerConfig200ResponseData) SetMp221(v string)`

SetMp221 sets Mp221 field to given value.

### HasMp221

`func (o *GetContainerConfig200ResponseData) HasMp221() bool`

HasMp221 returns a boolean if a field has been set.

### GetMp222

`func (o *GetContainerConfig200ResponseData) GetMp222() string`

GetMp222 returns the Mp222 field if non-nil, zero value otherwise.

### GetMp222Ok

`func (o *GetContainerConfig200ResponseData) GetMp222Ok() (*string, bool)`

GetMp222Ok returns a tuple with the Mp222 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp222

`func (o *GetContainerConfig200ResponseData) SetMp222(v string)`

SetMp222 sets Mp222 field to given value.

### HasMp222

`func (o *GetContainerConfig200ResponseData) HasMp222() bool`

HasMp222 returns a boolean if a field has been set.

### GetMp223

`func (o *GetContainerConfig200ResponseData) GetMp223() string`

GetMp223 returns the Mp223 field if non-nil, zero value otherwise.

### GetMp223Ok

`func (o *GetContainerConfig200ResponseData) GetMp223Ok() (*string, bool)`

GetMp223Ok returns a tuple with the Mp223 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp223

`func (o *GetContainerConfig200ResponseData) SetMp223(v string)`

SetMp223 sets Mp223 field to given value.

### HasMp223

`func (o *GetContainerConfig200ResponseData) HasMp223() bool`

HasMp223 returns a boolean if a field has been set.

### GetMp224

`func (o *GetContainerConfig200ResponseData) GetMp224() string`

GetMp224 returns the Mp224 field if non-nil, zero value otherwise.

### GetMp224Ok

`func (o *GetContainerConfig200ResponseData) GetMp224Ok() (*string, bool)`

GetMp224Ok returns a tuple with the Mp224 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp224

`func (o *GetContainerConfig200ResponseData) SetMp224(v string)`

SetMp224 sets Mp224 field to given value.

### HasMp224

`func (o *GetContainerConfig200ResponseData) HasMp224() bool`

HasMp224 returns a boolean if a field has been set.

### GetMp225

`func (o *GetContainerConfig200ResponseData) GetMp225() string`

GetMp225 returns the Mp225 field if non-nil, zero value otherwise.

### GetMp225Ok

`func (o *GetContainerConfig200ResponseData) GetMp225Ok() (*string, bool)`

GetMp225Ok returns a tuple with the Mp225 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp225

`func (o *GetContainerConfig200ResponseData) SetMp225(v string)`

SetMp225 sets Mp225 field to given value.

### HasMp225

`func (o *GetContainerConfig200ResponseData) HasMp225() bool`

HasMp225 returns a boolean if a field has been set.

### GetMp226

`func (o *GetContainerConfig200ResponseData) GetMp226() string`

GetMp226 returns the Mp226 field if non-nil, zero value otherwise.

### GetMp226Ok

`func (o *GetContainerConfig200ResponseData) GetMp226Ok() (*string, bool)`

GetMp226Ok returns a tuple with the Mp226 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp226

`func (o *GetContainerConfig200ResponseData) SetMp226(v string)`

SetMp226 sets Mp226 field to given value.

### HasMp226

`func (o *GetContainerConfig200ResponseData) HasMp226() bool`

HasMp226 returns a boolean if a field has been set.

### GetMp227

`func (o *GetContainerConfig200ResponseData) GetMp227() string`

GetMp227 returns the Mp227 field if non-nil, zero value otherwise.

### GetMp227Ok

`func (o *GetContainerConfig200ResponseData) GetMp227Ok() (*string, bool)`

GetMp227Ok returns a tuple with the Mp227 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp227

`func (o *GetContainerConfig200ResponseData) SetMp227(v string)`

SetMp227 sets Mp227 field to given value.

### HasMp227

`func (o *GetContainerConfig200ResponseData) HasMp227() bool`

HasMp227 returns a boolean if a field has been set.

### GetMp228

`func (o *GetContainerConfig200ResponseData) GetMp228() string`

GetMp228 returns the Mp228 field if non-nil, zero value otherwise.

### GetMp228Ok

`func (o *GetContainerConfig200ResponseData) GetMp228Ok() (*string, bool)`

GetMp228Ok returns a tuple with the Mp228 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp228

`func (o *GetContainerConfig200ResponseData) SetMp228(v string)`

SetMp228 sets Mp228 field to given value.

### HasMp228

`func (o *GetContainerConfig200ResponseData) HasMp228() bool`

HasMp228 returns a boolean if a field has been set.

### GetMp229

`func (o *GetContainerConfig200ResponseData) GetMp229() string`

GetMp229 returns the Mp229 field if non-nil, zero value otherwise.

### GetMp229Ok

`func (o *GetContainerConfig200ResponseData) GetMp229Ok() (*string, bool)`

GetMp229Ok returns a tuple with the Mp229 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp229

`func (o *GetContainerConfig200ResponseData) SetMp229(v string)`

SetMp229 sets Mp229 field to given value.

### HasMp229

`func (o *GetContainerConfig200ResponseData) HasMp229() bool`

HasMp229 returns a boolean if a field has been set.

### GetMp230

`func (o *GetContainerConfig200ResponseData) GetMp230() string`

GetMp230 returns the Mp230 field if non-nil, zero value otherwise.

### GetMp230Ok

`func (o *GetContainerConfig200ResponseData) GetMp230Ok() (*string, bool)`

GetMp230Ok returns a tuple with the Mp230 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp230

`func (o *GetContainerConfig200ResponseData) SetMp230(v string)`

SetMp230 sets Mp230 field to given value.

### HasMp230

`func (o *GetContainerConfig200ResponseData) HasMp230() bool`

HasMp230 returns a boolean if a field has been set.

### GetMp231

`func (o *GetContainerConfig200ResponseData) GetMp231() string`

GetMp231 returns the Mp231 field if non-nil, zero value otherwise.

### GetMp231Ok

`func (o *GetContainerConfig200ResponseData) GetMp231Ok() (*string, bool)`

GetMp231Ok returns a tuple with the Mp231 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp231

`func (o *GetContainerConfig200ResponseData) SetMp231(v string)`

SetMp231 sets Mp231 field to given value.

### HasMp231

`func (o *GetContainerConfig200ResponseData) HasMp231() bool`

HasMp231 returns a boolean if a field has been set.

### GetMp232

`func (o *GetContainerConfig200ResponseData) GetMp232() string`

GetMp232 returns the Mp232 field if non-nil, zero value otherwise.

### GetMp232Ok

`func (o *GetContainerConfig200ResponseData) GetMp232Ok() (*string, bool)`

GetMp232Ok returns a tuple with the Mp232 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp232

`func (o *GetContainerConfig200ResponseData) SetMp232(v string)`

SetMp232 sets Mp232 field to given value.

### HasMp232

`func (o *GetContainerConfig200ResponseData) HasMp232() bool`

HasMp232 returns a boolean if a field has been set.

### GetMp233

`func (o *GetContainerConfig200ResponseData) GetMp233() string`

GetMp233 returns the Mp233 field if non-nil, zero value otherwise.

### GetMp233Ok

`func (o *GetContainerConfig200ResponseData) GetMp233Ok() (*string, bool)`

GetMp233Ok returns a tuple with the Mp233 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp233

`func (o *GetContainerConfig200ResponseData) SetMp233(v string)`

SetMp233 sets Mp233 field to given value.

### HasMp233

`func (o *GetContainerConfig200ResponseData) HasMp233() bool`

HasMp233 returns a boolean if a field has been set.

### GetMp234

`func (o *GetContainerConfig200ResponseData) GetMp234() string`

GetMp234 returns the Mp234 field if non-nil, zero value otherwise.

### GetMp234Ok

`func (o *GetContainerConfig200ResponseData) GetMp234Ok() (*string, bool)`

GetMp234Ok returns a tuple with the Mp234 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp234

`func (o *GetContainerConfig200ResponseData) SetMp234(v string)`

SetMp234 sets Mp234 field to given value.

### HasMp234

`func (o *GetContainerConfig200ResponseData) HasMp234() bool`

HasMp234 returns a boolean if a field has been set.

### GetMp235

`func (o *GetContainerConfig200ResponseData) GetMp235() string`

GetMp235 returns the Mp235 field if non-nil, zero value otherwise.

### GetMp235Ok

`func (o *GetContainerConfig200ResponseData) GetMp235Ok() (*string, bool)`

GetMp235Ok returns a tuple with the Mp235 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp235

`func (o *GetContainerConfig200ResponseData) SetMp235(v string)`

SetMp235 sets Mp235 field to given value.

### HasMp235

`func (o *GetContainerConfig200ResponseData) HasMp235() bool`

HasMp235 returns a boolean if a field has been set.

### GetMp236

`func (o *GetContainerConfig200ResponseData) GetMp236() string`

GetMp236 returns the Mp236 field if non-nil, zero value otherwise.

### GetMp236Ok

`func (o *GetContainerConfig200ResponseData) GetMp236Ok() (*string, bool)`

GetMp236Ok returns a tuple with the Mp236 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp236

`func (o *GetContainerConfig200ResponseData) SetMp236(v string)`

SetMp236 sets Mp236 field to given value.

### HasMp236

`func (o *GetContainerConfig200ResponseData) HasMp236() bool`

HasMp236 returns a boolean if a field has been set.

### GetMp237

`func (o *GetContainerConfig200ResponseData) GetMp237() string`

GetMp237 returns the Mp237 field if non-nil, zero value otherwise.

### GetMp237Ok

`func (o *GetContainerConfig200ResponseData) GetMp237Ok() (*string, bool)`

GetMp237Ok returns a tuple with the Mp237 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp237

`func (o *GetContainerConfig200ResponseData) SetMp237(v string)`

SetMp237 sets Mp237 field to given value.

### HasMp237

`func (o *GetContainerConfig200ResponseData) HasMp237() bool`

HasMp237 returns a boolean if a field has been set.

### GetMp238

`func (o *GetContainerConfig200ResponseData) GetMp238() string`

GetMp238 returns the Mp238 field if non-nil, zero value otherwise.

### GetMp238Ok

`func (o *GetContainerConfig200ResponseData) GetMp238Ok() (*string, bool)`

GetMp238Ok returns a tuple with the Mp238 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp238

`func (o *GetContainerConfig200ResponseData) SetMp238(v string)`

SetMp238 sets Mp238 field to given value.

### HasMp238

`func (o *GetContainerConfig200ResponseData) HasMp238() bool`

HasMp238 returns a boolean if a field has been set.

### GetMp239

`func (o *GetContainerConfig200ResponseData) GetMp239() string`

GetMp239 returns the Mp239 field if non-nil, zero value otherwise.

### GetMp239Ok

`func (o *GetContainerConfig200ResponseData) GetMp239Ok() (*string, bool)`

GetMp239Ok returns a tuple with the Mp239 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp239

`func (o *GetContainerConfig200ResponseData) SetMp239(v string)`

SetMp239 sets Mp239 field to given value.

### HasMp239

`func (o *GetContainerConfig200ResponseData) HasMp239() bool`

HasMp239 returns a boolean if a field has been set.

### GetMp240

`func (o *GetContainerConfig200ResponseData) GetMp240() string`

GetMp240 returns the Mp240 field if non-nil, zero value otherwise.

### GetMp240Ok

`func (o *GetContainerConfig200ResponseData) GetMp240Ok() (*string, bool)`

GetMp240Ok returns a tuple with the Mp240 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp240

`func (o *GetContainerConfig200ResponseData) SetMp240(v string)`

SetMp240 sets Mp240 field to given value.

### HasMp240

`func (o *GetContainerConfig200ResponseData) HasMp240() bool`

HasMp240 returns a boolean if a field has been set.

### GetMp241

`func (o *GetContainerConfig200ResponseData) GetMp241() string`

GetMp241 returns the Mp241 field if non-nil, zero value otherwise.

### GetMp241Ok

`func (o *GetContainerConfig200ResponseData) GetMp241Ok() (*string, bool)`

GetMp241Ok returns a tuple with the Mp241 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp241

`func (o *GetContainerConfig200ResponseData) SetMp241(v string)`

SetMp241 sets Mp241 field to given value.

### HasMp241

`func (o *GetContainerConfig200ResponseData) HasMp241() bool`

HasMp241 returns a boolean if a field has been set.

### GetMp242

`func (o *GetContainerConfig200ResponseData) GetMp242() string`

GetMp242 returns the Mp242 field if non-nil, zero value otherwise.

### GetMp242Ok

`func (o *GetContainerConfig200ResponseData) GetMp242Ok() (*string, bool)`

GetMp242Ok returns a tuple with the Mp242 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp242

`func (o *GetContainerConfig200ResponseData) SetMp242(v string)`

SetMp242 sets Mp242 field to given value.

### HasMp242

`func (o *GetContainerConfig200ResponseData) HasMp242() bool`

HasMp242 returns a boolean if a field has been set.

### GetMp243

`func (o *GetContainerConfig200ResponseData) GetMp243() string`

GetMp243 returns the Mp243 field if non-nil, zero value otherwise.

### GetMp243Ok

`func (o *GetContainerConfig200ResponseData) GetMp243Ok() (*string, bool)`

GetMp243Ok returns a tuple with the Mp243 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp243

`func (o *GetContainerConfig200ResponseData) SetMp243(v string)`

SetMp243 sets Mp243 field to given value.

### HasMp243

`func (o *GetContainerConfig200ResponseData) HasMp243() bool`

HasMp243 returns a boolean if a field has been set.

### GetMp244

`func (o *GetContainerConfig200ResponseData) GetMp244() string`

GetMp244 returns the Mp244 field if non-nil, zero value otherwise.

### GetMp244Ok

`func (o *GetContainerConfig200ResponseData) GetMp244Ok() (*string, bool)`

GetMp244Ok returns a tuple with the Mp244 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp244

`func (o *GetContainerConfig200ResponseData) SetMp244(v string)`

SetMp244 sets Mp244 field to given value.

### HasMp244

`func (o *GetContainerConfig200ResponseData) HasMp244() bool`

HasMp244 returns a boolean if a field has been set.

### GetMp245

`func (o *GetContainerConfig200ResponseData) GetMp245() string`

GetMp245 returns the Mp245 field if non-nil, zero value otherwise.

### GetMp245Ok

`func (o *GetContainerConfig200ResponseData) GetMp245Ok() (*string, bool)`

GetMp245Ok returns a tuple with the Mp245 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp245

`func (o *GetContainerConfig200ResponseData) SetMp245(v string)`

SetMp245 sets Mp245 field to given value.

### HasMp245

`func (o *GetContainerConfig200ResponseData) HasMp245() bool`

HasMp245 returns a boolean if a field has been set.

### GetMp246

`func (o *GetContainerConfig200ResponseData) GetMp246() string`

GetMp246 returns the Mp246 field if non-nil, zero value otherwise.

### GetMp246Ok

`func (o *GetContainerConfig200ResponseData) GetMp246Ok() (*string, bool)`

GetMp246Ok returns a tuple with the Mp246 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp246

`func (o *GetContainerConfig200ResponseData) SetMp246(v string)`

SetMp246 sets Mp246 field to given value.

### HasMp246

`func (o *GetContainerConfig200ResponseData) HasMp246() bool`

HasMp246 returns a boolean if a field has been set.

### GetMp247

`func (o *GetContainerConfig200ResponseData) GetMp247() string`

GetMp247 returns the Mp247 field if non-nil, zero value otherwise.

### GetMp247Ok

`func (o *GetContainerConfig200ResponseData) GetMp247Ok() (*string, bool)`

GetMp247Ok returns a tuple with the Mp247 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp247

`func (o *GetContainerConfig200ResponseData) SetMp247(v string)`

SetMp247 sets Mp247 field to given value.

### HasMp247

`func (o *GetContainerConfig200ResponseData) HasMp247() bool`

HasMp247 returns a boolean if a field has been set.

### GetMp248

`func (o *GetContainerConfig200ResponseData) GetMp248() string`

GetMp248 returns the Mp248 field if non-nil, zero value otherwise.

### GetMp248Ok

`func (o *GetContainerConfig200ResponseData) GetMp248Ok() (*string, bool)`

GetMp248Ok returns a tuple with the Mp248 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp248

`func (o *GetContainerConfig200ResponseData) SetMp248(v string)`

SetMp248 sets Mp248 field to given value.

### HasMp248

`func (o *GetContainerConfig200ResponseData) HasMp248() bool`

HasMp248 returns a boolean if a field has been set.

### GetMp249

`func (o *GetContainerConfig200ResponseData) GetMp249() string`

GetMp249 returns the Mp249 field if non-nil, zero value otherwise.

### GetMp249Ok

`func (o *GetContainerConfig200ResponseData) GetMp249Ok() (*string, bool)`

GetMp249Ok returns a tuple with the Mp249 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp249

`func (o *GetContainerConfig200ResponseData) SetMp249(v string)`

SetMp249 sets Mp249 field to given value.

### HasMp249

`func (o *GetContainerConfig200ResponseData) HasMp249() bool`

HasMp249 returns a boolean if a field has been set.

### GetMp250

`func (o *GetContainerConfig200ResponseData) GetMp250() string`

GetMp250 returns the Mp250 field if non-nil, zero value otherwise.

### GetMp250Ok

`func (o *GetContainerConfig200ResponseData) GetMp250Ok() (*string, bool)`

GetMp250Ok returns a tuple with the Mp250 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp250

`func (o *GetContainerConfig200ResponseData) SetMp250(v string)`

SetMp250 sets Mp250 field to given value.

### HasMp250

`func (o *GetContainerConfig200ResponseData) HasMp250() bool`

HasMp250 returns a boolean if a field has been set.

### GetMp251

`func (o *GetContainerConfig200ResponseData) GetMp251() string`

GetMp251 returns the Mp251 field if non-nil, zero value otherwise.

### GetMp251Ok

`func (o *GetContainerConfig200ResponseData) GetMp251Ok() (*string, bool)`

GetMp251Ok returns a tuple with the Mp251 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp251

`func (o *GetContainerConfig200ResponseData) SetMp251(v string)`

SetMp251 sets Mp251 field to given value.

### HasMp251

`func (o *GetContainerConfig200ResponseData) HasMp251() bool`

HasMp251 returns a boolean if a field has been set.

### GetMp252

`func (o *GetContainerConfig200ResponseData) GetMp252() string`

GetMp252 returns the Mp252 field if non-nil, zero value otherwise.

### GetMp252Ok

`func (o *GetContainerConfig200ResponseData) GetMp252Ok() (*string, bool)`

GetMp252Ok returns a tuple with the Mp252 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp252

`func (o *GetContainerConfig200ResponseData) SetMp252(v string)`

SetMp252 sets Mp252 field to given value.

### HasMp252

`func (o *GetContainerConfig200ResponseData) HasMp252() bool`

HasMp252 returns a boolean if a field has been set.

### GetMp253

`func (o *GetContainerConfig200ResponseData) GetMp253() string`

GetMp253 returns the Mp253 field if non-nil, zero value otherwise.

### GetMp253Ok

`func (o *GetContainerConfig200ResponseData) GetMp253Ok() (*string, bool)`

GetMp253Ok returns a tuple with the Mp253 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp253

`func (o *GetContainerConfig200ResponseData) SetMp253(v string)`

SetMp253 sets Mp253 field to given value.

### HasMp253

`func (o *GetContainerConfig200ResponseData) HasMp253() bool`

HasMp253 returns a boolean if a field has been set.

### GetMp254

`func (o *GetContainerConfig200ResponseData) GetMp254() string`

GetMp254 returns the Mp254 field if non-nil, zero value otherwise.

### GetMp254Ok

`func (o *GetContainerConfig200ResponseData) GetMp254Ok() (*string, bool)`

GetMp254Ok returns a tuple with the Mp254 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp254

`func (o *GetContainerConfig200ResponseData) SetMp254(v string)`

SetMp254 sets Mp254 field to given value.

### HasMp254

`func (o *GetContainerConfig200ResponseData) HasMp254() bool`

HasMp254 returns a boolean if a field has been set.

### GetMp255

`func (o *GetContainerConfig200ResponseData) GetMp255() string`

GetMp255 returns the Mp255 field if non-nil, zero value otherwise.

### GetMp255Ok

`func (o *GetContainerConfig200ResponseData) GetMp255Ok() (*string, bool)`

GetMp255Ok returns a tuple with the Mp255 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp255

`func (o *GetContainerConfig200ResponseData) SetMp255(v string)`

SetMp255 sets Mp255 field to given value.

### HasMp255

`func (o *GetContainerConfig200ResponseData) HasMp255() bool`

HasMp255 returns a boolean if a field has been set.

### GetNameserver

`func (o *GetContainerConfig200ResponseData) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *GetContainerConfig200ResponseData) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *GetContainerConfig200ResponseData) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *GetContainerConfig200ResponseData) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetNet0

`func (o *GetContainerConfig200ResponseData) GetNet0() string`

GetNet0 returns the Net0 field if non-nil, zero value otherwise.

### GetNet0Ok

`func (o *GetContainerConfig200ResponseData) GetNet0Ok() (*string, bool)`

GetNet0Ok returns a tuple with the Net0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet0

`func (o *GetContainerConfig200ResponseData) SetNet0(v string)`

SetNet0 sets Net0 field to given value.

### HasNet0

`func (o *GetContainerConfig200ResponseData) HasNet0() bool`

HasNet0 returns a boolean if a field has been set.

### GetNet1

`func (o *GetContainerConfig200ResponseData) GetNet1() string`

GetNet1 returns the Net1 field if non-nil, zero value otherwise.

### GetNet1Ok

`func (o *GetContainerConfig200ResponseData) GetNet1Ok() (*string, bool)`

GetNet1Ok returns a tuple with the Net1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet1

`func (o *GetContainerConfig200ResponseData) SetNet1(v string)`

SetNet1 sets Net1 field to given value.

### HasNet1

`func (o *GetContainerConfig200ResponseData) HasNet1() bool`

HasNet1 returns a boolean if a field has been set.

### GetNet2

`func (o *GetContainerConfig200ResponseData) GetNet2() string`

GetNet2 returns the Net2 field if non-nil, zero value otherwise.

### GetNet2Ok

`func (o *GetContainerConfig200ResponseData) GetNet2Ok() (*string, bool)`

GetNet2Ok returns a tuple with the Net2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet2

`func (o *GetContainerConfig200ResponseData) SetNet2(v string)`

SetNet2 sets Net2 field to given value.

### HasNet2

`func (o *GetContainerConfig200ResponseData) HasNet2() bool`

HasNet2 returns a boolean if a field has been set.

### GetNet3

`func (o *GetContainerConfig200ResponseData) GetNet3() string`

GetNet3 returns the Net3 field if non-nil, zero value otherwise.

### GetNet3Ok

`func (o *GetContainerConfig200ResponseData) GetNet3Ok() (*string, bool)`

GetNet3Ok returns a tuple with the Net3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet3

`func (o *GetContainerConfig200ResponseData) SetNet3(v string)`

SetNet3 sets Net3 field to given value.

### HasNet3

`func (o *GetContainerConfig200ResponseData) HasNet3() bool`

HasNet3 returns a boolean if a field has been set.

### GetNet4

`func (o *GetContainerConfig200ResponseData) GetNet4() string`

GetNet4 returns the Net4 field if non-nil, zero value otherwise.

### GetNet4Ok

`func (o *GetContainerConfig200ResponseData) GetNet4Ok() (*string, bool)`

GetNet4Ok returns a tuple with the Net4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet4

`func (o *GetContainerConfig200ResponseData) SetNet4(v string)`

SetNet4 sets Net4 field to given value.

### HasNet4

`func (o *GetContainerConfig200ResponseData) HasNet4() bool`

HasNet4 returns a boolean if a field has been set.

### GetNet5

`func (o *GetContainerConfig200ResponseData) GetNet5() string`

GetNet5 returns the Net5 field if non-nil, zero value otherwise.

### GetNet5Ok

`func (o *GetContainerConfig200ResponseData) GetNet5Ok() (*string, bool)`

GetNet5Ok returns a tuple with the Net5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet5

`func (o *GetContainerConfig200ResponseData) SetNet5(v string)`

SetNet5 sets Net5 field to given value.

### HasNet5

`func (o *GetContainerConfig200ResponseData) HasNet5() bool`

HasNet5 returns a boolean if a field has been set.

### GetNet6

`func (o *GetContainerConfig200ResponseData) GetNet6() string`

GetNet6 returns the Net6 field if non-nil, zero value otherwise.

### GetNet6Ok

`func (o *GetContainerConfig200ResponseData) GetNet6Ok() (*string, bool)`

GetNet6Ok returns a tuple with the Net6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet6

`func (o *GetContainerConfig200ResponseData) SetNet6(v string)`

SetNet6 sets Net6 field to given value.

### HasNet6

`func (o *GetContainerConfig200ResponseData) HasNet6() bool`

HasNet6 returns a boolean if a field has been set.

### GetNet7

`func (o *GetContainerConfig200ResponseData) GetNet7() string`

GetNet7 returns the Net7 field if non-nil, zero value otherwise.

### GetNet7Ok

`func (o *GetContainerConfig200ResponseData) GetNet7Ok() (*string, bool)`

GetNet7Ok returns a tuple with the Net7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet7

`func (o *GetContainerConfig200ResponseData) SetNet7(v string)`

SetNet7 sets Net7 field to given value.

### HasNet7

`func (o *GetContainerConfig200ResponseData) HasNet7() bool`

HasNet7 returns a boolean if a field has been set.

### GetNet8

`func (o *GetContainerConfig200ResponseData) GetNet8() string`

GetNet8 returns the Net8 field if non-nil, zero value otherwise.

### GetNet8Ok

`func (o *GetContainerConfig200ResponseData) GetNet8Ok() (*string, bool)`

GetNet8Ok returns a tuple with the Net8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet8

`func (o *GetContainerConfig200ResponseData) SetNet8(v string)`

SetNet8 sets Net8 field to given value.

### HasNet8

`func (o *GetContainerConfig200ResponseData) HasNet8() bool`

HasNet8 returns a boolean if a field has been set.

### GetNet9

`func (o *GetContainerConfig200ResponseData) GetNet9() string`

GetNet9 returns the Net9 field if non-nil, zero value otherwise.

### GetNet9Ok

`func (o *GetContainerConfig200ResponseData) GetNet9Ok() (*string, bool)`

GetNet9Ok returns a tuple with the Net9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet9

`func (o *GetContainerConfig200ResponseData) SetNet9(v string)`

SetNet9 sets Net9 field to given value.

### HasNet9

`func (o *GetContainerConfig200ResponseData) HasNet9() bool`

HasNet9 returns a boolean if a field has been set.

### GetNet10

`func (o *GetContainerConfig200ResponseData) GetNet10() string`

GetNet10 returns the Net10 field if non-nil, zero value otherwise.

### GetNet10Ok

`func (o *GetContainerConfig200ResponseData) GetNet10Ok() (*string, bool)`

GetNet10Ok returns a tuple with the Net10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet10

`func (o *GetContainerConfig200ResponseData) SetNet10(v string)`

SetNet10 sets Net10 field to given value.

### HasNet10

`func (o *GetContainerConfig200ResponseData) HasNet10() bool`

HasNet10 returns a boolean if a field has been set.

### GetNet11

`func (o *GetContainerConfig200ResponseData) GetNet11() string`

GetNet11 returns the Net11 field if non-nil, zero value otherwise.

### GetNet11Ok

`func (o *GetContainerConfig200ResponseData) GetNet11Ok() (*string, bool)`

GetNet11Ok returns a tuple with the Net11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet11

`func (o *GetContainerConfig200ResponseData) SetNet11(v string)`

SetNet11 sets Net11 field to given value.

### HasNet11

`func (o *GetContainerConfig200ResponseData) HasNet11() bool`

HasNet11 returns a boolean if a field has been set.

### GetNet12

`func (o *GetContainerConfig200ResponseData) GetNet12() string`

GetNet12 returns the Net12 field if non-nil, zero value otherwise.

### GetNet12Ok

`func (o *GetContainerConfig200ResponseData) GetNet12Ok() (*string, bool)`

GetNet12Ok returns a tuple with the Net12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet12

`func (o *GetContainerConfig200ResponseData) SetNet12(v string)`

SetNet12 sets Net12 field to given value.

### HasNet12

`func (o *GetContainerConfig200ResponseData) HasNet12() bool`

HasNet12 returns a boolean if a field has been set.

### GetNet13

`func (o *GetContainerConfig200ResponseData) GetNet13() string`

GetNet13 returns the Net13 field if non-nil, zero value otherwise.

### GetNet13Ok

`func (o *GetContainerConfig200ResponseData) GetNet13Ok() (*string, bool)`

GetNet13Ok returns a tuple with the Net13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet13

`func (o *GetContainerConfig200ResponseData) SetNet13(v string)`

SetNet13 sets Net13 field to given value.

### HasNet13

`func (o *GetContainerConfig200ResponseData) HasNet13() bool`

HasNet13 returns a boolean if a field has been set.

### GetNet14

`func (o *GetContainerConfig200ResponseData) GetNet14() string`

GetNet14 returns the Net14 field if non-nil, zero value otherwise.

### GetNet14Ok

`func (o *GetContainerConfig200ResponseData) GetNet14Ok() (*string, bool)`

GetNet14Ok returns a tuple with the Net14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet14

`func (o *GetContainerConfig200ResponseData) SetNet14(v string)`

SetNet14 sets Net14 field to given value.

### HasNet14

`func (o *GetContainerConfig200ResponseData) HasNet14() bool`

HasNet14 returns a boolean if a field has been set.

### GetNet15

`func (o *GetContainerConfig200ResponseData) GetNet15() string`

GetNet15 returns the Net15 field if non-nil, zero value otherwise.

### GetNet15Ok

`func (o *GetContainerConfig200ResponseData) GetNet15Ok() (*string, bool)`

GetNet15Ok returns a tuple with the Net15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet15

`func (o *GetContainerConfig200ResponseData) SetNet15(v string)`

SetNet15 sets Net15 field to given value.

### HasNet15

`func (o *GetContainerConfig200ResponseData) HasNet15() bool`

HasNet15 returns a boolean if a field has been set.

### GetNet16

`func (o *GetContainerConfig200ResponseData) GetNet16() string`

GetNet16 returns the Net16 field if non-nil, zero value otherwise.

### GetNet16Ok

`func (o *GetContainerConfig200ResponseData) GetNet16Ok() (*string, bool)`

GetNet16Ok returns a tuple with the Net16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet16

`func (o *GetContainerConfig200ResponseData) SetNet16(v string)`

SetNet16 sets Net16 field to given value.

### HasNet16

`func (o *GetContainerConfig200ResponseData) HasNet16() bool`

HasNet16 returns a boolean if a field has been set.

### GetNet17

`func (o *GetContainerConfig200ResponseData) GetNet17() string`

GetNet17 returns the Net17 field if non-nil, zero value otherwise.

### GetNet17Ok

`func (o *GetContainerConfig200ResponseData) GetNet17Ok() (*string, bool)`

GetNet17Ok returns a tuple with the Net17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet17

`func (o *GetContainerConfig200ResponseData) SetNet17(v string)`

SetNet17 sets Net17 field to given value.

### HasNet17

`func (o *GetContainerConfig200ResponseData) HasNet17() bool`

HasNet17 returns a boolean if a field has been set.

### GetNet18

`func (o *GetContainerConfig200ResponseData) GetNet18() string`

GetNet18 returns the Net18 field if non-nil, zero value otherwise.

### GetNet18Ok

`func (o *GetContainerConfig200ResponseData) GetNet18Ok() (*string, bool)`

GetNet18Ok returns a tuple with the Net18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet18

`func (o *GetContainerConfig200ResponseData) SetNet18(v string)`

SetNet18 sets Net18 field to given value.

### HasNet18

`func (o *GetContainerConfig200ResponseData) HasNet18() bool`

HasNet18 returns a boolean if a field has been set.

### GetNet19

`func (o *GetContainerConfig200ResponseData) GetNet19() string`

GetNet19 returns the Net19 field if non-nil, zero value otherwise.

### GetNet19Ok

`func (o *GetContainerConfig200ResponseData) GetNet19Ok() (*string, bool)`

GetNet19Ok returns a tuple with the Net19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet19

`func (o *GetContainerConfig200ResponseData) SetNet19(v string)`

SetNet19 sets Net19 field to given value.

### HasNet19

`func (o *GetContainerConfig200ResponseData) HasNet19() bool`

HasNet19 returns a boolean if a field has been set.

### GetNet20

`func (o *GetContainerConfig200ResponseData) GetNet20() string`

GetNet20 returns the Net20 field if non-nil, zero value otherwise.

### GetNet20Ok

`func (o *GetContainerConfig200ResponseData) GetNet20Ok() (*string, bool)`

GetNet20Ok returns a tuple with the Net20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet20

`func (o *GetContainerConfig200ResponseData) SetNet20(v string)`

SetNet20 sets Net20 field to given value.

### HasNet20

`func (o *GetContainerConfig200ResponseData) HasNet20() bool`

HasNet20 returns a boolean if a field has been set.

### GetNet21

`func (o *GetContainerConfig200ResponseData) GetNet21() string`

GetNet21 returns the Net21 field if non-nil, zero value otherwise.

### GetNet21Ok

`func (o *GetContainerConfig200ResponseData) GetNet21Ok() (*string, bool)`

GetNet21Ok returns a tuple with the Net21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet21

`func (o *GetContainerConfig200ResponseData) SetNet21(v string)`

SetNet21 sets Net21 field to given value.

### HasNet21

`func (o *GetContainerConfig200ResponseData) HasNet21() bool`

HasNet21 returns a boolean if a field has been set.

### GetNet22

`func (o *GetContainerConfig200ResponseData) GetNet22() string`

GetNet22 returns the Net22 field if non-nil, zero value otherwise.

### GetNet22Ok

`func (o *GetContainerConfig200ResponseData) GetNet22Ok() (*string, bool)`

GetNet22Ok returns a tuple with the Net22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet22

`func (o *GetContainerConfig200ResponseData) SetNet22(v string)`

SetNet22 sets Net22 field to given value.

### HasNet22

`func (o *GetContainerConfig200ResponseData) HasNet22() bool`

HasNet22 returns a boolean if a field has been set.

### GetNet23

`func (o *GetContainerConfig200ResponseData) GetNet23() string`

GetNet23 returns the Net23 field if non-nil, zero value otherwise.

### GetNet23Ok

`func (o *GetContainerConfig200ResponseData) GetNet23Ok() (*string, bool)`

GetNet23Ok returns a tuple with the Net23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet23

`func (o *GetContainerConfig200ResponseData) SetNet23(v string)`

SetNet23 sets Net23 field to given value.

### HasNet23

`func (o *GetContainerConfig200ResponseData) HasNet23() bool`

HasNet23 returns a boolean if a field has been set.

### GetNet24

`func (o *GetContainerConfig200ResponseData) GetNet24() string`

GetNet24 returns the Net24 field if non-nil, zero value otherwise.

### GetNet24Ok

`func (o *GetContainerConfig200ResponseData) GetNet24Ok() (*string, bool)`

GetNet24Ok returns a tuple with the Net24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet24

`func (o *GetContainerConfig200ResponseData) SetNet24(v string)`

SetNet24 sets Net24 field to given value.

### HasNet24

`func (o *GetContainerConfig200ResponseData) HasNet24() bool`

HasNet24 returns a boolean if a field has been set.

### GetNet25

`func (o *GetContainerConfig200ResponseData) GetNet25() string`

GetNet25 returns the Net25 field if non-nil, zero value otherwise.

### GetNet25Ok

`func (o *GetContainerConfig200ResponseData) GetNet25Ok() (*string, bool)`

GetNet25Ok returns a tuple with the Net25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet25

`func (o *GetContainerConfig200ResponseData) SetNet25(v string)`

SetNet25 sets Net25 field to given value.

### HasNet25

`func (o *GetContainerConfig200ResponseData) HasNet25() bool`

HasNet25 returns a boolean if a field has been set.

### GetNet26

`func (o *GetContainerConfig200ResponseData) GetNet26() string`

GetNet26 returns the Net26 field if non-nil, zero value otherwise.

### GetNet26Ok

`func (o *GetContainerConfig200ResponseData) GetNet26Ok() (*string, bool)`

GetNet26Ok returns a tuple with the Net26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet26

`func (o *GetContainerConfig200ResponseData) SetNet26(v string)`

SetNet26 sets Net26 field to given value.

### HasNet26

`func (o *GetContainerConfig200ResponseData) HasNet26() bool`

HasNet26 returns a boolean if a field has been set.

### GetNet27

`func (o *GetContainerConfig200ResponseData) GetNet27() string`

GetNet27 returns the Net27 field if non-nil, zero value otherwise.

### GetNet27Ok

`func (o *GetContainerConfig200ResponseData) GetNet27Ok() (*string, bool)`

GetNet27Ok returns a tuple with the Net27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet27

`func (o *GetContainerConfig200ResponseData) SetNet27(v string)`

SetNet27 sets Net27 field to given value.

### HasNet27

`func (o *GetContainerConfig200ResponseData) HasNet27() bool`

HasNet27 returns a boolean if a field has been set.

### GetNet28

`func (o *GetContainerConfig200ResponseData) GetNet28() string`

GetNet28 returns the Net28 field if non-nil, zero value otherwise.

### GetNet28Ok

`func (o *GetContainerConfig200ResponseData) GetNet28Ok() (*string, bool)`

GetNet28Ok returns a tuple with the Net28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet28

`func (o *GetContainerConfig200ResponseData) SetNet28(v string)`

SetNet28 sets Net28 field to given value.

### HasNet28

`func (o *GetContainerConfig200ResponseData) HasNet28() bool`

HasNet28 returns a boolean if a field has been set.

### GetNet29

`func (o *GetContainerConfig200ResponseData) GetNet29() string`

GetNet29 returns the Net29 field if non-nil, zero value otherwise.

### GetNet29Ok

`func (o *GetContainerConfig200ResponseData) GetNet29Ok() (*string, bool)`

GetNet29Ok returns a tuple with the Net29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet29

`func (o *GetContainerConfig200ResponseData) SetNet29(v string)`

SetNet29 sets Net29 field to given value.

### HasNet29

`func (o *GetContainerConfig200ResponseData) HasNet29() bool`

HasNet29 returns a boolean if a field has been set.

### GetNet30

`func (o *GetContainerConfig200ResponseData) GetNet30() string`

GetNet30 returns the Net30 field if non-nil, zero value otherwise.

### GetNet30Ok

`func (o *GetContainerConfig200ResponseData) GetNet30Ok() (*string, bool)`

GetNet30Ok returns a tuple with the Net30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet30

`func (o *GetContainerConfig200ResponseData) SetNet30(v string)`

SetNet30 sets Net30 field to given value.

### HasNet30

`func (o *GetContainerConfig200ResponseData) HasNet30() bool`

HasNet30 returns a boolean if a field has been set.

### GetNet31

`func (o *GetContainerConfig200ResponseData) GetNet31() string`

GetNet31 returns the Net31 field if non-nil, zero value otherwise.

### GetNet31Ok

`func (o *GetContainerConfig200ResponseData) GetNet31Ok() (*string, bool)`

GetNet31Ok returns a tuple with the Net31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet31

`func (o *GetContainerConfig200ResponseData) SetNet31(v string)`

SetNet31 sets Net31 field to given value.

### HasNet31

`func (o *GetContainerConfig200ResponseData) HasNet31() bool`

HasNet31 returns a boolean if a field has been set.

### GetOnboot

`func (o *GetContainerConfig200ResponseData) GetOnboot() int32`

GetOnboot returns the Onboot field if non-nil, zero value otherwise.

### GetOnbootOk

`func (o *GetContainerConfig200ResponseData) GetOnbootOk() (*int32, bool)`

GetOnbootOk returns a tuple with the Onboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboot

`func (o *GetContainerConfig200ResponseData) SetOnboot(v int32)`

SetOnboot sets Onboot field to given value.

### HasOnboot

`func (o *GetContainerConfig200ResponseData) HasOnboot() bool`

HasOnboot returns a boolean if a field has been set.

### GetOstype

`func (o *GetContainerConfig200ResponseData) GetOstype() string`

GetOstype returns the Ostype field if non-nil, zero value otherwise.

### GetOstypeOk

`func (o *GetContainerConfig200ResponseData) GetOstypeOk() (*string, bool)`

GetOstypeOk returns a tuple with the Ostype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstype

`func (o *GetContainerConfig200ResponseData) SetOstype(v string)`

SetOstype sets Ostype field to given value.

### HasOstype

`func (o *GetContainerConfig200ResponseData) HasOstype() bool`

HasOstype returns a boolean if a field has been set.

### GetProtection

`func (o *GetContainerConfig200ResponseData) GetProtection() int32`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *GetContainerConfig200ResponseData) GetProtectionOk() (*int32, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *GetContainerConfig200ResponseData) SetProtection(v int32)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *GetContainerConfig200ResponseData) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetRootfs

`func (o *GetContainerConfig200ResponseData) GetRootfs() string`

GetRootfs returns the Rootfs field if non-nil, zero value otherwise.

### GetRootfsOk

`func (o *GetContainerConfig200ResponseData) GetRootfsOk() (*string, bool)`

GetRootfsOk returns a tuple with the Rootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfs

`func (o *GetContainerConfig200ResponseData) SetRootfs(v string)`

SetRootfs sets Rootfs field to given value.

### HasRootfs

`func (o *GetContainerConfig200ResponseData) HasRootfs() bool`

HasRootfs returns a boolean if a field has been set.

### GetSearchdomain

`func (o *GetContainerConfig200ResponseData) GetSearchdomain() string`

GetSearchdomain returns the Searchdomain field if non-nil, zero value otherwise.

### GetSearchdomainOk

`func (o *GetContainerConfig200ResponseData) GetSearchdomainOk() (*string, bool)`

GetSearchdomainOk returns a tuple with the Searchdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchdomain

`func (o *GetContainerConfig200ResponseData) SetSearchdomain(v string)`

SetSearchdomain sets Searchdomain field to given value.

### HasSearchdomain

`func (o *GetContainerConfig200ResponseData) HasSearchdomain() bool`

HasSearchdomain returns a boolean if a field has been set.

### GetStartup

`func (o *GetContainerConfig200ResponseData) GetStartup() string`

GetStartup returns the Startup field if non-nil, zero value otherwise.

### GetStartupOk

`func (o *GetContainerConfig200ResponseData) GetStartupOk() (*string, bool)`

GetStartupOk returns a tuple with the Startup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartup

`func (o *GetContainerConfig200ResponseData) SetStartup(v string)`

SetStartup sets Startup field to given value.

### HasStartup

`func (o *GetContainerConfig200ResponseData) HasStartup() bool`

HasStartup returns a boolean if a field has been set.

### GetSwap

`func (o *GetContainerConfig200ResponseData) GetSwap() int64`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *GetContainerConfig200ResponseData) GetSwapOk() (*int64, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *GetContainerConfig200ResponseData) SetSwap(v int64)`

SetSwap sets Swap field to given value.

### HasSwap

`func (o *GetContainerConfig200ResponseData) HasSwap() bool`

HasSwap returns a boolean if a field has been set.

### GetTags

`func (o *GetContainerConfig200ResponseData) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetContainerConfig200ResponseData) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetContainerConfig200ResponseData) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetContainerConfig200ResponseData) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *GetContainerConfig200ResponseData) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetContainerConfig200ResponseData) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetContainerConfig200ResponseData) SetTemplate(v int32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetContainerConfig200ResponseData) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTimezone

`func (o *GetContainerConfig200ResponseData) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetContainerConfig200ResponseData) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetContainerConfig200ResponseData) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetContainerConfig200ResponseData) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTty

`func (o *GetContainerConfig200ResponseData) GetTty() int64`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *GetContainerConfig200ResponseData) GetTtyOk() (*int64, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *GetContainerConfig200ResponseData) SetTty(v int64)`

SetTty sets Tty field to given value.

### HasTty

`func (o *GetContainerConfig200ResponseData) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetUnprivileged

`func (o *GetContainerConfig200ResponseData) GetUnprivileged() int32`

GetUnprivileged returns the Unprivileged field if non-nil, zero value otherwise.

### GetUnprivilegedOk

`func (o *GetContainerConfig200ResponseData) GetUnprivilegedOk() (*int32, bool)`

GetUnprivilegedOk returns a tuple with the Unprivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprivileged

`func (o *GetContainerConfig200ResponseData) SetUnprivileged(v int32)`

SetUnprivileged sets Unprivileged field to given value.

### HasUnprivileged

`func (o *GetContainerConfig200ResponseData) HasUnprivileged() bool`

HasUnprivileged returns a boolean if a field has been set.

### GetUnused0

`func (o *GetContainerConfig200ResponseData) GetUnused0() string`

GetUnused0 returns the Unused0 field if non-nil, zero value otherwise.

### GetUnused0Ok

`func (o *GetContainerConfig200ResponseData) GetUnused0Ok() (*string, bool)`

GetUnused0Ok returns a tuple with the Unused0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused0

`func (o *GetContainerConfig200ResponseData) SetUnused0(v string)`

SetUnused0 sets Unused0 field to given value.

### HasUnused0

`func (o *GetContainerConfig200ResponseData) HasUnused0() bool`

HasUnused0 returns a boolean if a field has been set.

### GetUnused1

`func (o *GetContainerConfig200ResponseData) GetUnused1() string`

GetUnused1 returns the Unused1 field if non-nil, zero value otherwise.

### GetUnused1Ok

`func (o *GetContainerConfig200ResponseData) GetUnused1Ok() (*string, bool)`

GetUnused1Ok returns a tuple with the Unused1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused1

`func (o *GetContainerConfig200ResponseData) SetUnused1(v string)`

SetUnused1 sets Unused1 field to given value.

### HasUnused1

`func (o *GetContainerConfig200ResponseData) HasUnused1() bool`

HasUnused1 returns a boolean if a field has been set.

### GetUnused2

`func (o *GetContainerConfig200ResponseData) GetUnused2() string`

GetUnused2 returns the Unused2 field if non-nil, zero value otherwise.

### GetUnused2Ok

`func (o *GetContainerConfig200ResponseData) GetUnused2Ok() (*string, bool)`

GetUnused2Ok returns a tuple with the Unused2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused2

`func (o *GetContainerConfig200ResponseData) SetUnused2(v string)`

SetUnused2 sets Unused2 field to given value.

### HasUnused2

`func (o *GetContainerConfig200ResponseData) HasUnused2() bool`

HasUnused2 returns a boolean if a field has been set.

### GetUnused3

`func (o *GetContainerConfig200ResponseData) GetUnused3() string`

GetUnused3 returns the Unused3 field if non-nil, zero value otherwise.

### GetUnused3Ok

`func (o *GetContainerConfig200ResponseData) GetUnused3Ok() (*string, bool)`

GetUnused3Ok returns a tuple with the Unused3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused3

`func (o *GetContainerConfig200ResponseData) SetUnused3(v string)`

SetUnused3 sets Unused3 field to given value.

### HasUnused3

`func (o *GetContainerConfig200ResponseData) HasUnused3() bool`

HasUnused3 returns a boolean if a field has been set.

### GetUnused4

`func (o *GetContainerConfig200ResponseData) GetUnused4() string`

GetUnused4 returns the Unused4 field if non-nil, zero value otherwise.

### GetUnused4Ok

`func (o *GetContainerConfig200ResponseData) GetUnused4Ok() (*string, bool)`

GetUnused4Ok returns a tuple with the Unused4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused4

`func (o *GetContainerConfig200ResponseData) SetUnused4(v string)`

SetUnused4 sets Unused4 field to given value.

### HasUnused4

`func (o *GetContainerConfig200ResponseData) HasUnused4() bool`

HasUnused4 returns a boolean if a field has been set.

### GetUnused5

`func (o *GetContainerConfig200ResponseData) GetUnused5() string`

GetUnused5 returns the Unused5 field if non-nil, zero value otherwise.

### GetUnused5Ok

`func (o *GetContainerConfig200ResponseData) GetUnused5Ok() (*string, bool)`

GetUnused5Ok returns a tuple with the Unused5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused5

`func (o *GetContainerConfig200ResponseData) SetUnused5(v string)`

SetUnused5 sets Unused5 field to given value.

### HasUnused5

`func (o *GetContainerConfig200ResponseData) HasUnused5() bool`

HasUnused5 returns a boolean if a field has been set.

### GetUnused6

`func (o *GetContainerConfig200ResponseData) GetUnused6() string`

GetUnused6 returns the Unused6 field if non-nil, zero value otherwise.

### GetUnused6Ok

`func (o *GetContainerConfig200ResponseData) GetUnused6Ok() (*string, bool)`

GetUnused6Ok returns a tuple with the Unused6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused6

`func (o *GetContainerConfig200ResponseData) SetUnused6(v string)`

SetUnused6 sets Unused6 field to given value.

### HasUnused6

`func (o *GetContainerConfig200ResponseData) HasUnused6() bool`

HasUnused6 returns a boolean if a field has been set.

### GetUnused7

`func (o *GetContainerConfig200ResponseData) GetUnused7() string`

GetUnused7 returns the Unused7 field if non-nil, zero value otherwise.

### GetUnused7Ok

`func (o *GetContainerConfig200ResponseData) GetUnused7Ok() (*string, bool)`

GetUnused7Ok returns a tuple with the Unused7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused7

`func (o *GetContainerConfig200ResponseData) SetUnused7(v string)`

SetUnused7 sets Unused7 field to given value.

### HasUnused7

`func (o *GetContainerConfig200ResponseData) HasUnused7() bool`

HasUnused7 returns a boolean if a field has been set.

### GetUnused8

`func (o *GetContainerConfig200ResponseData) GetUnused8() string`

GetUnused8 returns the Unused8 field if non-nil, zero value otherwise.

### GetUnused8Ok

`func (o *GetContainerConfig200ResponseData) GetUnused8Ok() (*string, bool)`

GetUnused8Ok returns a tuple with the Unused8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused8

`func (o *GetContainerConfig200ResponseData) SetUnused8(v string)`

SetUnused8 sets Unused8 field to given value.

### HasUnused8

`func (o *GetContainerConfig200ResponseData) HasUnused8() bool`

HasUnused8 returns a boolean if a field has been set.

### GetUnused9

`func (o *GetContainerConfig200ResponseData) GetUnused9() string`

GetUnused9 returns the Unused9 field if non-nil, zero value otherwise.

### GetUnused9Ok

`func (o *GetContainerConfig200ResponseData) GetUnused9Ok() (*string, bool)`

GetUnused9Ok returns a tuple with the Unused9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused9

`func (o *GetContainerConfig200ResponseData) SetUnused9(v string)`

SetUnused9 sets Unused9 field to given value.

### HasUnused9

`func (o *GetContainerConfig200ResponseData) HasUnused9() bool`

HasUnused9 returns a boolean if a field has been set.

### GetUnused10

`func (o *GetContainerConfig200ResponseData) GetUnused10() string`

GetUnused10 returns the Unused10 field if non-nil, zero value otherwise.

### GetUnused10Ok

`func (o *GetContainerConfig200ResponseData) GetUnused10Ok() (*string, bool)`

GetUnused10Ok returns a tuple with the Unused10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused10

`func (o *GetContainerConfig200ResponseData) SetUnused10(v string)`

SetUnused10 sets Unused10 field to given value.

### HasUnused10

`func (o *GetContainerConfig200ResponseData) HasUnused10() bool`

HasUnused10 returns a boolean if a field has been set.

### GetUnused11

`func (o *GetContainerConfig200ResponseData) GetUnused11() string`

GetUnused11 returns the Unused11 field if non-nil, zero value otherwise.

### GetUnused11Ok

`func (o *GetContainerConfig200ResponseData) GetUnused11Ok() (*string, bool)`

GetUnused11Ok returns a tuple with the Unused11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused11

`func (o *GetContainerConfig200ResponseData) SetUnused11(v string)`

SetUnused11 sets Unused11 field to given value.

### HasUnused11

`func (o *GetContainerConfig200ResponseData) HasUnused11() bool`

HasUnused11 returns a boolean if a field has been set.

### GetUnused12

`func (o *GetContainerConfig200ResponseData) GetUnused12() string`

GetUnused12 returns the Unused12 field if non-nil, zero value otherwise.

### GetUnused12Ok

`func (o *GetContainerConfig200ResponseData) GetUnused12Ok() (*string, bool)`

GetUnused12Ok returns a tuple with the Unused12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused12

`func (o *GetContainerConfig200ResponseData) SetUnused12(v string)`

SetUnused12 sets Unused12 field to given value.

### HasUnused12

`func (o *GetContainerConfig200ResponseData) HasUnused12() bool`

HasUnused12 returns a boolean if a field has been set.

### GetUnused13

`func (o *GetContainerConfig200ResponseData) GetUnused13() string`

GetUnused13 returns the Unused13 field if non-nil, zero value otherwise.

### GetUnused13Ok

`func (o *GetContainerConfig200ResponseData) GetUnused13Ok() (*string, bool)`

GetUnused13Ok returns a tuple with the Unused13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused13

`func (o *GetContainerConfig200ResponseData) SetUnused13(v string)`

SetUnused13 sets Unused13 field to given value.

### HasUnused13

`func (o *GetContainerConfig200ResponseData) HasUnused13() bool`

HasUnused13 returns a boolean if a field has been set.

### GetUnused14

`func (o *GetContainerConfig200ResponseData) GetUnused14() string`

GetUnused14 returns the Unused14 field if non-nil, zero value otherwise.

### GetUnused14Ok

`func (o *GetContainerConfig200ResponseData) GetUnused14Ok() (*string, bool)`

GetUnused14Ok returns a tuple with the Unused14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused14

`func (o *GetContainerConfig200ResponseData) SetUnused14(v string)`

SetUnused14 sets Unused14 field to given value.

### HasUnused14

`func (o *GetContainerConfig200ResponseData) HasUnused14() bool`

HasUnused14 returns a boolean if a field has been set.

### GetUnused15

`func (o *GetContainerConfig200ResponseData) GetUnused15() string`

GetUnused15 returns the Unused15 field if non-nil, zero value otherwise.

### GetUnused15Ok

`func (o *GetContainerConfig200ResponseData) GetUnused15Ok() (*string, bool)`

GetUnused15Ok returns a tuple with the Unused15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused15

`func (o *GetContainerConfig200ResponseData) SetUnused15(v string)`

SetUnused15 sets Unused15 field to given value.

### HasUnused15

`func (o *GetContainerConfig200ResponseData) HasUnused15() bool`

HasUnused15 returns a boolean if a field has been set.

### GetUnused16

`func (o *GetContainerConfig200ResponseData) GetUnused16() string`

GetUnused16 returns the Unused16 field if non-nil, zero value otherwise.

### GetUnused16Ok

`func (o *GetContainerConfig200ResponseData) GetUnused16Ok() (*string, bool)`

GetUnused16Ok returns a tuple with the Unused16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused16

`func (o *GetContainerConfig200ResponseData) SetUnused16(v string)`

SetUnused16 sets Unused16 field to given value.

### HasUnused16

`func (o *GetContainerConfig200ResponseData) HasUnused16() bool`

HasUnused16 returns a boolean if a field has been set.

### GetUnused17

`func (o *GetContainerConfig200ResponseData) GetUnused17() string`

GetUnused17 returns the Unused17 field if non-nil, zero value otherwise.

### GetUnused17Ok

`func (o *GetContainerConfig200ResponseData) GetUnused17Ok() (*string, bool)`

GetUnused17Ok returns a tuple with the Unused17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused17

`func (o *GetContainerConfig200ResponseData) SetUnused17(v string)`

SetUnused17 sets Unused17 field to given value.

### HasUnused17

`func (o *GetContainerConfig200ResponseData) HasUnused17() bool`

HasUnused17 returns a boolean if a field has been set.

### GetUnused18

`func (o *GetContainerConfig200ResponseData) GetUnused18() string`

GetUnused18 returns the Unused18 field if non-nil, zero value otherwise.

### GetUnused18Ok

`func (o *GetContainerConfig200ResponseData) GetUnused18Ok() (*string, bool)`

GetUnused18Ok returns a tuple with the Unused18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused18

`func (o *GetContainerConfig200ResponseData) SetUnused18(v string)`

SetUnused18 sets Unused18 field to given value.

### HasUnused18

`func (o *GetContainerConfig200ResponseData) HasUnused18() bool`

HasUnused18 returns a boolean if a field has been set.

### GetUnused19

`func (o *GetContainerConfig200ResponseData) GetUnused19() string`

GetUnused19 returns the Unused19 field if non-nil, zero value otherwise.

### GetUnused19Ok

`func (o *GetContainerConfig200ResponseData) GetUnused19Ok() (*string, bool)`

GetUnused19Ok returns a tuple with the Unused19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused19

`func (o *GetContainerConfig200ResponseData) SetUnused19(v string)`

SetUnused19 sets Unused19 field to given value.

### HasUnused19

`func (o *GetContainerConfig200ResponseData) HasUnused19() bool`

HasUnused19 returns a boolean if a field has been set.

### GetUnused20

`func (o *GetContainerConfig200ResponseData) GetUnused20() string`

GetUnused20 returns the Unused20 field if non-nil, zero value otherwise.

### GetUnused20Ok

`func (o *GetContainerConfig200ResponseData) GetUnused20Ok() (*string, bool)`

GetUnused20Ok returns a tuple with the Unused20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused20

`func (o *GetContainerConfig200ResponseData) SetUnused20(v string)`

SetUnused20 sets Unused20 field to given value.

### HasUnused20

`func (o *GetContainerConfig200ResponseData) HasUnused20() bool`

HasUnused20 returns a boolean if a field has been set.

### GetUnused21

`func (o *GetContainerConfig200ResponseData) GetUnused21() string`

GetUnused21 returns the Unused21 field if non-nil, zero value otherwise.

### GetUnused21Ok

`func (o *GetContainerConfig200ResponseData) GetUnused21Ok() (*string, bool)`

GetUnused21Ok returns a tuple with the Unused21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused21

`func (o *GetContainerConfig200ResponseData) SetUnused21(v string)`

SetUnused21 sets Unused21 field to given value.

### HasUnused21

`func (o *GetContainerConfig200ResponseData) HasUnused21() bool`

HasUnused21 returns a boolean if a field has been set.

### GetUnused22

`func (o *GetContainerConfig200ResponseData) GetUnused22() string`

GetUnused22 returns the Unused22 field if non-nil, zero value otherwise.

### GetUnused22Ok

`func (o *GetContainerConfig200ResponseData) GetUnused22Ok() (*string, bool)`

GetUnused22Ok returns a tuple with the Unused22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused22

`func (o *GetContainerConfig200ResponseData) SetUnused22(v string)`

SetUnused22 sets Unused22 field to given value.

### HasUnused22

`func (o *GetContainerConfig200ResponseData) HasUnused22() bool`

HasUnused22 returns a boolean if a field has been set.

### GetUnused23

`func (o *GetContainerConfig200ResponseData) GetUnused23() string`

GetUnused23 returns the Unused23 field if non-nil, zero value otherwise.

### GetUnused23Ok

`func (o *GetContainerConfig200ResponseData) GetUnused23Ok() (*string, bool)`

GetUnused23Ok returns a tuple with the Unused23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused23

`func (o *GetContainerConfig200ResponseData) SetUnused23(v string)`

SetUnused23 sets Unused23 field to given value.

### HasUnused23

`func (o *GetContainerConfig200ResponseData) HasUnused23() bool`

HasUnused23 returns a boolean if a field has been set.

### GetUnused24

`func (o *GetContainerConfig200ResponseData) GetUnused24() string`

GetUnused24 returns the Unused24 field if non-nil, zero value otherwise.

### GetUnused24Ok

`func (o *GetContainerConfig200ResponseData) GetUnused24Ok() (*string, bool)`

GetUnused24Ok returns a tuple with the Unused24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused24

`func (o *GetContainerConfig200ResponseData) SetUnused24(v string)`

SetUnused24 sets Unused24 field to given value.

### HasUnused24

`func (o *GetContainerConfig200ResponseData) HasUnused24() bool`

HasUnused24 returns a boolean if a field has been set.

### GetUnused25

`func (o *GetContainerConfig200ResponseData) GetUnused25() string`

GetUnused25 returns the Unused25 field if non-nil, zero value otherwise.

### GetUnused25Ok

`func (o *GetContainerConfig200ResponseData) GetUnused25Ok() (*string, bool)`

GetUnused25Ok returns a tuple with the Unused25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused25

`func (o *GetContainerConfig200ResponseData) SetUnused25(v string)`

SetUnused25 sets Unused25 field to given value.

### HasUnused25

`func (o *GetContainerConfig200ResponseData) HasUnused25() bool`

HasUnused25 returns a boolean if a field has been set.

### GetUnused26

`func (o *GetContainerConfig200ResponseData) GetUnused26() string`

GetUnused26 returns the Unused26 field if non-nil, zero value otherwise.

### GetUnused26Ok

`func (o *GetContainerConfig200ResponseData) GetUnused26Ok() (*string, bool)`

GetUnused26Ok returns a tuple with the Unused26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused26

`func (o *GetContainerConfig200ResponseData) SetUnused26(v string)`

SetUnused26 sets Unused26 field to given value.

### HasUnused26

`func (o *GetContainerConfig200ResponseData) HasUnused26() bool`

HasUnused26 returns a boolean if a field has been set.

### GetUnused27

`func (o *GetContainerConfig200ResponseData) GetUnused27() string`

GetUnused27 returns the Unused27 field if non-nil, zero value otherwise.

### GetUnused27Ok

`func (o *GetContainerConfig200ResponseData) GetUnused27Ok() (*string, bool)`

GetUnused27Ok returns a tuple with the Unused27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused27

`func (o *GetContainerConfig200ResponseData) SetUnused27(v string)`

SetUnused27 sets Unused27 field to given value.

### HasUnused27

`func (o *GetContainerConfig200ResponseData) HasUnused27() bool`

HasUnused27 returns a boolean if a field has been set.

### GetUnused28

`func (o *GetContainerConfig200ResponseData) GetUnused28() string`

GetUnused28 returns the Unused28 field if non-nil, zero value otherwise.

### GetUnused28Ok

`func (o *GetContainerConfig200ResponseData) GetUnused28Ok() (*string, bool)`

GetUnused28Ok returns a tuple with the Unused28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused28

`func (o *GetContainerConfig200ResponseData) SetUnused28(v string)`

SetUnused28 sets Unused28 field to given value.

### HasUnused28

`func (o *GetContainerConfig200ResponseData) HasUnused28() bool`

HasUnused28 returns a boolean if a field has been set.

### GetUnused29

`func (o *GetContainerConfig200ResponseData) GetUnused29() string`

GetUnused29 returns the Unused29 field if non-nil, zero value otherwise.

### GetUnused29Ok

`func (o *GetContainerConfig200ResponseData) GetUnused29Ok() (*string, bool)`

GetUnused29Ok returns a tuple with the Unused29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused29

`func (o *GetContainerConfig200ResponseData) SetUnused29(v string)`

SetUnused29 sets Unused29 field to given value.

### HasUnused29

`func (o *GetContainerConfig200ResponseData) HasUnused29() bool`

HasUnused29 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


