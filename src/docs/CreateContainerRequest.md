# CreateContainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arch** | Pointer to **string** | OS architecture type. | [optional] 
**Bwlimit** | Pointer to **float32** | Override I/O bandwidth limit (in KiB/s). | [optional] 
**Cmode** | Pointer to **string** | Console mode. By default, the console command tries to open a connection to one of the available tty devices. By setting cmode to &#39;console&#39; it tries to attach to /dev/console instead. If you set cmode to &#39;shell&#39;, it simply invokes a shell inside the container (no login). | [optional] 
**Console** | Pointer to **int32** | Attach a console device (/dev/console) to the container. | [optional] 
**Cores** | Pointer to **int64** | The number of cores assigned to the container. A container can use all available cores by default. | [optional] 
**Cpulimit** | Pointer to **float32** | Limit of CPU usage.  NOTE: If the computer has 2 CPUs, it has a total of &#39;2&#39; CPU time. Value &#39;0&#39; indicates no CPU limit. | [optional] 
**Cpuunits** | Pointer to **int64** | CPU weight for a container, will be clamped to [1, 10000] in cgroup v2. | [optional] 
**Debug** | Pointer to **int32** | Try to be more verbose. For now this only enables debug log-level on start. | [optional] 
**Description** | Pointer to **string** | Description for the Container. Shown in the web-interface CT&#39;s summary. This is saved as comment inside the configuration file. | [optional] 
**Features** | Pointer to **string** | Allow containers access to advanced features. | [optional] 
**Force** | Pointer to **int32** | Allow to overwrite existing container. | [optional] 
**Hookscript** | Pointer to **string** | Script that will be exectued during various steps in the containers lifetime. | [optional] 
**Hostname** | Pointer to **string** | Set a host name for the container. | [optional] 
**IgnoreUnpackErrors** | Pointer to **int32** | Ignore errors when extracting the template. | [optional] 
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
**Ostemplate** | **string** | The OS template or backup file. | 
**Ostype** | Pointer to **string** | OS type. This is used to setup configuration inside the container, and corresponds to lxc setup scripts in /usr/share/lxc/config/&lt;ostype&gt;.common.conf. Value &#39;unmanaged&#39; can be used to skip and OS specific setup. | [optional] 
**Password** | Pointer to **string** | Sets root password inside container. | [optional] 
**Pool** | Pointer to **string** | Add the VM to the specified pool. | [optional] 
**Protection** | Pointer to **int32** | Sets the protection flag of the container. This will prevent the CT or CT&#39;s disk remove/update operation. | [optional] 
**Restore** | Pointer to **int32** | Mark this as restore task. | [optional] 
**Rootfs** | Pointer to **string** | Use volume as container root. | [optional] 
**Searchdomain** | Pointer to **string** | Sets DNS search domains for a container. Create will automatically use the setting from the host if you neither set searchdomain nor nameserver. | [optional] 
**SshPublicKeys** | Pointer to **string** | Setup public SSH keys (one key per line, OpenSSH format). | [optional] 
**Start** | Pointer to **int32** | Start the CT after its creation finished successfully. | [optional] 
**Startup** | Pointer to **string** | Startup and shutdown behavior. Order is a non-negative number defining the general startup order. Shutdown in done with reverse ordering. Additionally you can set the &#39;up&#39; or &#39;down&#39; delay in seconds, which specifies a delay to wait before the next VM is started or stopped. | [optional] 
**Storage** | Pointer to **string** | Default Storage. | [optional] 
**Swap** | Pointer to **int64** | Amount of SWAP for the container in MB. | [optional] 
**Tags** | Pointer to **string** | Tags of the Container. This is only meta information. | [optional] 
**Template** | Pointer to **int32** | Enable/disable Template. | [optional] 
**Timezone** | Pointer to **string** | Time zone to use in the container. If option isn&#39;t set, then nothing will be done. Can be set to &#39;host&#39; to match the host time zone, or an arbitrary time zone option from /usr/share/zoneinfo/zone.tab | [optional] 
**Tty** | Pointer to **int64** | Specify the number of tty available to the container | [optional] 
**Unique** | Pointer to **int32** | Assign a unique random ethernet address. | [optional] 
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
**Vmid** | **int64** | The (unique) ID of the VM. | 

## Methods

### NewCreateContainerRequest

`func NewCreateContainerRequest(ostemplate string, vmid int64, ) *CreateContainerRequest`

NewCreateContainerRequest instantiates a new CreateContainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContainerRequestWithDefaults

`func NewCreateContainerRequestWithDefaults() *CreateContainerRequest`

NewCreateContainerRequestWithDefaults instantiates a new CreateContainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArch

`func (o *CreateContainerRequest) GetArch() string`

GetArch returns the Arch field if non-nil, zero value otherwise.

### GetArchOk

`func (o *CreateContainerRequest) GetArchOk() (*string, bool)`

GetArchOk returns a tuple with the Arch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArch

`func (o *CreateContainerRequest) SetArch(v string)`

SetArch sets Arch field to given value.

### HasArch

`func (o *CreateContainerRequest) HasArch() bool`

HasArch returns a boolean if a field has been set.

### GetBwlimit

`func (o *CreateContainerRequest) GetBwlimit() float32`

GetBwlimit returns the Bwlimit field if non-nil, zero value otherwise.

### GetBwlimitOk

`func (o *CreateContainerRequest) GetBwlimitOk() (*float32, bool)`

GetBwlimitOk returns a tuple with the Bwlimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBwlimit

`func (o *CreateContainerRequest) SetBwlimit(v float32)`

SetBwlimit sets Bwlimit field to given value.

### HasBwlimit

`func (o *CreateContainerRequest) HasBwlimit() bool`

HasBwlimit returns a boolean if a field has been set.

### GetCmode

`func (o *CreateContainerRequest) GetCmode() string`

GetCmode returns the Cmode field if non-nil, zero value otherwise.

### GetCmodeOk

`func (o *CreateContainerRequest) GetCmodeOk() (*string, bool)`

GetCmodeOk returns a tuple with the Cmode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmode

`func (o *CreateContainerRequest) SetCmode(v string)`

SetCmode sets Cmode field to given value.

### HasCmode

`func (o *CreateContainerRequest) HasCmode() bool`

HasCmode returns a boolean if a field has been set.

### GetConsole

`func (o *CreateContainerRequest) GetConsole() int32`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *CreateContainerRequest) GetConsoleOk() (*int32, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *CreateContainerRequest) SetConsole(v int32)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *CreateContainerRequest) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### GetCores

`func (o *CreateContainerRequest) GetCores() int64`

GetCores returns the Cores field if non-nil, zero value otherwise.

### GetCoresOk

`func (o *CreateContainerRequest) GetCoresOk() (*int64, bool)`

GetCoresOk returns a tuple with the Cores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCores

`func (o *CreateContainerRequest) SetCores(v int64)`

SetCores sets Cores field to given value.

### HasCores

`func (o *CreateContainerRequest) HasCores() bool`

HasCores returns a boolean if a field has been set.

### GetCpulimit

`func (o *CreateContainerRequest) GetCpulimit() float32`

GetCpulimit returns the Cpulimit field if non-nil, zero value otherwise.

### GetCpulimitOk

`func (o *CreateContainerRequest) GetCpulimitOk() (*float32, bool)`

GetCpulimitOk returns a tuple with the Cpulimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpulimit

`func (o *CreateContainerRequest) SetCpulimit(v float32)`

SetCpulimit sets Cpulimit field to given value.

### HasCpulimit

`func (o *CreateContainerRequest) HasCpulimit() bool`

HasCpulimit returns a boolean if a field has been set.

### GetCpuunits

`func (o *CreateContainerRequest) GetCpuunits() int64`

GetCpuunits returns the Cpuunits field if non-nil, zero value otherwise.

### GetCpuunitsOk

`func (o *CreateContainerRequest) GetCpuunitsOk() (*int64, bool)`

GetCpuunitsOk returns a tuple with the Cpuunits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuunits

`func (o *CreateContainerRequest) SetCpuunits(v int64)`

SetCpuunits sets Cpuunits field to given value.

### HasCpuunits

`func (o *CreateContainerRequest) HasCpuunits() bool`

HasCpuunits returns a boolean if a field has been set.

### GetDebug

`func (o *CreateContainerRequest) GetDebug() int32`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *CreateContainerRequest) GetDebugOk() (*int32, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *CreateContainerRequest) SetDebug(v int32)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *CreateContainerRequest) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetDescription

`func (o *CreateContainerRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateContainerRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateContainerRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateContainerRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *CreateContainerRequest) GetFeatures() string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *CreateContainerRequest) GetFeaturesOk() (*string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *CreateContainerRequest) SetFeatures(v string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *CreateContainerRequest) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetForce

`func (o *CreateContainerRequest) GetForce() int32`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *CreateContainerRequest) GetForceOk() (*int32, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *CreateContainerRequest) SetForce(v int32)`

SetForce sets Force field to given value.

### HasForce

`func (o *CreateContainerRequest) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetHookscript

`func (o *CreateContainerRequest) GetHookscript() string`

GetHookscript returns the Hookscript field if non-nil, zero value otherwise.

### GetHookscriptOk

`func (o *CreateContainerRequest) GetHookscriptOk() (*string, bool)`

GetHookscriptOk returns a tuple with the Hookscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHookscript

`func (o *CreateContainerRequest) SetHookscript(v string)`

SetHookscript sets Hookscript field to given value.

### HasHookscript

`func (o *CreateContainerRequest) HasHookscript() bool`

HasHookscript returns a boolean if a field has been set.

### GetHostname

`func (o *CreateContainerRequest) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateContainerRequest) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateContainerRequest) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateContainerRequest) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetIgnoreUnpackErrors

`func (o *CreateContainerRequest) GetIgnoreUnpackErrors() int32`

GetIgnoreUnpackErrors returns the IgnoreUnpackErrors field if non-nil, zero value otherwise.

### GetIgnoreUnpackErrorsOk

`func (o *CreateContainerRequest) GetIgnoreUnpackErrorsOk() (*int32, bool)`

GetIgnoreUnpackErrorsOk returns a tuple with the IgnoreUnpackErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreUnpackErrors

`func (o *CreateContainerRequest) SetIgnoreUnpackErrors(v int32)`

SetIgnoreUnpackErrors sets IgnoreUnpackErrors field to given value.

### HasIgnoreUnpackErrors

`func (o *CreateContainerRequest) HasIgnoreUnpackErrors() bool`

HasIgnoreUnpackErrors returns a boolean if a field has been set.

### GetLock

`func (o *CreateContainerRequest) GetLock() string`

GetLock returns the Lock field if non-nil, zero value otherwise.

### GetLockOk

`func (o *CreateContainerRequest) GetLockOk() (*string, bool)`

GetLockOk returns a tuple with the Lock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLock

`func (o *CreateContainerRequest) SetLock(v string)`

SetLock sets Lock field to given value.

### HasLock

`func (o *CreateContainerRequest) HasLock() bool`

HasLock returns a boolean if a field has been set.

### GetMemory

`func (o *CreateContainerRequest) GetMemory() int64`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *CreateContainerRequest) GetMemoryOk() (*int64, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *CreateContainerRequest) SetMemory(v int64)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *CreateContainerRequest) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetMp0

`func (o *CreateContainerRequest) GetMp0() string`

GetMp0 returns the Mp0 field if non-nil, zero value otherwise.

### GetMp0Ok

`func (o *CreateContainerRequest) GetMp0Ok() (*string, bool)`

GetMp0Ok returns a tuple with the Mp0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp0

`func (o *CreateContainerRequest) SetMp0(v string)`

SetMp0 sets Mp0 field to given value.

### HasMp0

`func (o *CreateContainerRequest) HasMp0() bool`

HasMp0 returns a boolean if a field has been set.

### GetMp1

`func (o *CreateContainerRequest) GetMp1() string`

GetMp1 returns the Mp1 field if non-nil, zero value otherwise.

### GetMp1Ok

`func (o *CreateContainerRequest) GetMp1Ok() (*string, bool)`

GetMp1Ok returns a tuple with the Mp1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp1

`func (o *CreateContainerRequest) SetMp1(v string)`

SetMp1 sets Mp1 field to given value.

### HasMp1

`func (o *CreateContainerRequest) HasMp1() bool`

HasMp1 returns a boolean if a field has been set.

### GetMp2

`func (o *CreateContainerRequest) GetMp2() string`

GetMp2 returns the Mp2 field if non-nil, zero value otherwise.

### GetMp2Ok

`func (o *CreateContainerRequest) GetMp2Ok() (*string, bool)`

GetMp2Ok returns a tuple with the Mp2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp2

`func (o *CreateContainerRequest) SetMp2(v string)`

SetMp2 sets Mp2 field to given value.

### HasMp2

`func (o *CreateContainerRequest) HasMp2() bool`

HasMp2 returns a boolean if a field has been set.

### GetMp3

`func (o *CreateContainerRequest) GetMp3() string`

GetMp3 returns the Mp3 field if non-nil, zero value otherwise.

### GetMp3Ok

`func (o *CreateContainerRequest) GetMp3Ok() (*string, bool)`

GetMp3Ok returns a tuple with the Mp3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp3

`func (o *CreateContainerRequest) SetMp3(v string)`

SetMp3 sets Mp3 field to given value.

### HasMp3

`func (o *CreateContainerRequest) HasMp3() bool`

HasMp3 returns a boolean if a field has been set.

### GetMp4

`func (o *CreateContainerRequest) GetMp4() string`

GetMp4 returns the Mp4 field if non-nil, zero value otherwise.

### GetMp4Ok

`func (o *CreateContainerRequest) GetMp4Ok() (*string, bool)`

GetMp4Ok returns a tuple with the Mp4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp4

`func (o *CreateContainerRequest) SetMp4(v string)`

SetMp4 sets Mp4 field to given value.

### HasMp4

`func (o *CreateContainerRequest) HasMp4() bool`

HasMp4 returns a boolean if a field has been set.

### GetMp5

`func (o *CreateContainerRequest) GetMp5() string`

GetMp5 returns the Mp5 field if non-nil, zero value otherwise.

### GetMp5Ok

`func (o *CreateContainerRequest) GetMp5Ok() (*string, bool)`

GetMp5Ok returns a tuple with the Mp5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp5

`func (o *CreateContainerRequest) SetMp5(v string)`

SetMp5 sets Mp5 field to given value.

### HasMp5

`func (o *CreateContainerRequest) HasMp5() bool`

HasMp5 returns a boolean if a field has been set.

### GetMp6

`func (o *CreateContainerRequest) GetMp6() string`

GetMp6 returns the Mp6 field if non-nil, zero value otherwise.

### GetMp6Ok

`func (o *CreateContainerRequest) GetMp6Ok() (*string, bool)`

GetMp6Ok returns a tuple with the Mp6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp6

`func (o *CreateContainerRequest) SetMp6(v string)`

SetMp6 sets Mp6 field to given value.

### HasMp6

`func (o *CreateContainerRequest) HasMp6() bool`

HasMp6 returns a boolean if a field has been set.

### GetMp7

`func (o *CreateContainerRequest) GetMp7() string`

GetMp7 returns the Mp7 field if non-nil, zero value otherwise.

### GetMp7Ok

`func (o *CreateContainerRequest) GetMp7Ok() (*string, bool)`

GetMp7Ok returns a tuple with the Mp7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp7

`func (o *CreateContainerRequest) SetMp7(v string)`

SetMp7 sets Mp7 field to given value.

### HasMp7

`func (o *CreateContainerRequest) HasMp7() bool`

HasMp7 returns a boolean if a field has been set.

### GetMp8

`func (o *CreateContainerRequest) GetMp8() string`

GetMp8 returns the Mp8 field if non-nil, zero value otherwise.

### GetMp8Ok

`func (o *CreateContainerRequest) GetMp8Ok() (*string, bool)`

GetMp8Ok returns a tuple with the Mp8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp8

`func (o *CreateContainerRequest) SetMp8(v string)`

SetMp8 sets Mp8 field to given value.

### HasMp8

`func (o *CreateContainerRequest) HasMp8() bool`

HasMp8 returns a boolean if a field has been set.

### GetMp9

`func (o *CreateContainerRequest) GetMp9() string`

GetMp9 returns the Mp9 field if non-nil, zero value otherwise.

### GetMp9Ok

`func (o *CreateContainerRequest) GetMp9Ok() (*string, bool)`

GetMp9Ok returns a tuple with the Mp9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp9

`func (o *CreateContainerRequest) SetMp9(v string)`

SetMp9 sets Mp9 field to given value.

### HasMp9

`func (o *CreateContainerRequest) HasMp9() bool`

HasMp9 returns a boolean if a field has been set.

### GetMp10

`func (o *CreateContainerRequest) GetMp10() string`

GetMp10 returns the Mp10 field if non-nil, zero value otherwise.

### GetMp10Ok

`func (o *CreateContainerRequest) GetMp10Ok() (*string, bool)`

GetMp10Ok returns a tuple with the Mp10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp10

`func (o *CreateContainerRequest) SetMp10(v string)`

SetMp10 sets Mp10 field to given value.

### HasMp10

`func (o *CreateContainerRequest) HasMp10() bool`

HasMp10 returns a boolean if a field has been set.

### GetMp11

`func (o *CreateContainerRequest) GetMp11() string`

GetMp11 returns the Mp11 field if non-nil, zero value otherwise.

### GetMp11Ok

`func (o *CreateContainerRequest) GetMp11Ok() (*string, bool)`

GetMp11Ok returns a tuple with the Mp11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp11

`func (o *CreateContainerRequest) SetMp11(v string)`

SetMp11 sets Mp11 field to given value.

### HasMp11

`func (o *CreateContainerRequest) HasMp11() bool`

HasMp11 returns a boolean if a field has been set.

### GetMp12

`func (o *CreateContainerRequest) GetMp12() string`

GetMp12 returns the Mp12 field if non-nil, zero value otherwise.

### GetMp12Ok

`func (o *CreateContainerRequest) GetMp12Ok() (*string, bool)`

GetMp12Ok returns a tuple with the Mp12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp12

`func (o *CreateContainerRequest) SetMp12(v string)`

SetMp12 sets Mp12 field to given value.

### HasMp12

`func (o *CreateContainerRequest) HasMp12() bool`

HasMp12 returns a boolean if a field has been set.

### GetMp13

`func (o *CreateContainerRequest) GetMp13() string`

GetMp13 returns the Mp13 field if non-nil, zero value otherwise.

### GetMp13Ok

`func (o *CreateContainerRequest) GetMp13Ok() (*string, bool)`

GetMp13Ok returns a tuple with the Mp13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp13

`func (o *CreateContainerRequest) SetMp13(v string)`

SetMp13 sets Mp13 field to given value.

### HasMp13

`func (o *CreateContainerRequest) HasMp13() bool`

HasMp13 returns a boolean if a field has been set.

### GetMp14

`func (o *CreateContainerRequest) GetMp14() string`

GetMp14 returns the Mp14 field if non-nil, zero value otherwise.

### GetMp14Ok

`func (o *CreateContainerRequest) GetMp14Ok() (*string, bool)`

GetMp14Ok returns a tuple with the Mp14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp14

`func (o *CreateContainerRequest) SetMp14(v string)`

SetMp14 sets Mp14 field to given value.

### HasMp14

`func (o *CreateContainerRequest) HasMp14() bool`

HasMp14 returns a boolean if a field has been set.

### GetMp15

`func (o *CreateContainerRequest) GetMp15() string`

GetMp15 returns the Mp15 field if non-nil, zero value otherwise.

### GetMp15Ok

`func (o *CreateContainerRequest) GetMp15Ok() (*string, bool)`

GetMp15Ok returns a tuple with the Mp15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp15

`func (o *CreateContainerRequest) SetMp15(v string)`

SetMp15 sets Mp15 field to given value.

### HasMp15

`func (o *CreateContainerRequest) HasMp15() bool`

HasMp15 returns a boolean if a field has been set.

### GetMp16

`func (o *CreateContainerRequest) GetMp16() string`

GetMp16 returns the Mp16 field if non-nil, zero value otherwise.

### GetMp16Ok

`func (o *CreateContainerRequest) GetMp16Ok() (*string, bool)`

GetMp16Ok returns a tuple with the Mp16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp16

`func (o *CreateContainerRequest) SetMp16(v string)`

SetMp16 sets Mp16 field to given value.

### HasMp16

`func (o *CreateContainerRequest) HasMp16() bool`

HasMp16 returns a boolean if a field has been set.

### GetMp17

`func (o *CreateContainerRequest) GetMp17() string`

GetMp17 returns the Mp17 field if non-nil, zero value otherwise.

### GetMp17Ok

`func (o *CreateContainerRequest) GetMp17Ok() (*string, bool)`

GetMp17Ok returns a tuple with the Mp17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp17

`func (o *CreateContainerRequest) SetMp17(v string)`

SetMp17 sets Mp17 field to given value.

### HasMp17

`func (o *CreateContainerRequest) HasMp17() bool`

HasMp17 returns a boolean if a field has been set.

### GetMp18

`func (o *CreateContainerRequest) GetMp18() string`

GetMp18 returns the Mp18 field if non-nil, zero value otherwise.

### GetMp18Ok

`func (o *CreateContainerRequest) GetMp18Ok() (*string, bool)`

GetMp18Ok returns a tuple with the Mp18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp18

`func (o *CreateContainerRequest) SetMp18(v string)`

SetMp18 sets Mp18 field to given value.

### HasMp18

`func (o *CreateContainerRequest) HasMp18() bool`

HasMp18 returns a boolean if a field has been set.

### GetMp19

`func (o *CreateContainerRequest) GetMp19() string`

GetMp19 returns the Mp19 field if non-nil, zero value otherwise.

### GetMp19Ok

`func (o *CreateContainerRequest) GetMp19Ok() (*string, bool)`

GetMp19Ok returns a tuple with the Mp19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp19

`func (o *CreateContainerRequest) SetMp19(v string)`

SetMp19 sets Mp19 field to given value.

### HasMp19

`func (o *CreateContainerRequest) HasMp19() bool`

HasMp19 returns a boolean if a field has been set.

### GetMp20

`func (o *CreateContainerRequest) GetMp20() string`

GetMp20 returns the Mp20 field if non-nil, zero value otherwise.

### GetMp20Ok

`func (o *CreateContainerRequest) GetMp20Ok() (*string, bool)`

GetMp20Ok returns a tuple with the Mp20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp20

`func (o *CreateContainerRequest) SetMp20(v string)`

SetMp20 sets Mp20 field to given value.

### HasMp20

`func (o *CreateContainerRequest) HasMp20() bool`

HasMp20 returns a boolean if a field has been set.

### GetMp21

`func (o *CreateContainerRequest) GetMp21() string`

GetMp21 returns the Mp21 field if non-nil, zero value otherwise.

### GetMp21Ok

`func (o *CreateContainerRequest) GetMp21Ok() (*string, bool)`

GetMp21Ok returns a tuple with the Mp21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp21

`func (o *CreateContainerRequest) SetMp21(v string)`

SetMp21 sets Mp21 field to given value.

### HasMp21

`func (o *CreateContainerRequest) HasMp21() bool`

HasMp21 returns a boolean if a field has been set.

### GetMp22

`func (o *CreateContainerRequest) GetMp22() string`

GetMp22 returns the Mp22 field if non-nil, zero value otherwise.

### GetMp22Ok

`func (o *CreateContainerRequest) GetMp22Ok() (*string, bool)`

GetMp22Ok returns a tuple with the Mp22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp22

`func (o *CreateContainerRequest) SetMp22(v string)`

SetMp22 sets Mp22 field to given value.

### HasMp22

`func (o *CreateContainerRequest) HasMp22() bool`

HasMp22 returns a boolean if a field has been set.

### GetMp23

`func (o *CreateContainerRequest) GetMp23() string`

GetMp23 returns the Mp23 field if non-nil, zero value otherwise.

### GetMp23Ok

`func (o *CreateContainerRequest) GetMp23Ok() (*string, bool)`

GetMp23Ok returns a tuple with the Mp23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp23

`func (o *CreateContainerRequest) SetMp23(v string)`

SetMp23 sets Mp23 field to given value.

### HasMp23

`func (o *CreateContainerRequest) HasMp23() bool`

HasMp23 returns a boolean if a field has been set.

### GetMp24

`func (o *CreateContainerRequest) GetMp24() string`

GetMp24 returns the Mp24 field if non-nil, zero value otherwise.

### GetMp24Ok

`func (o *CreateContainerRequest) GetMp24Ok() (*string, bool)`

GetMp24Ok returns a tuple with the Mp24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp24

`func (o *CreateContainerRequest) SetMp24(v string)`

SetMp24 sets Mp24 field to given value.

### HasMp24

`func (o *CreateContainerRequest) HasMp24() bool`

HasMp24 returns a boolean if a field has been set.

### GetMp25

`func (o *CreateContainerRequest) GetMp25() string`

GetMp25 returns the Mp25 field if non-nil, zero value otherwise.

### GetMp25Ok

`func (o *CreateContainerRequest) GetMp25Ok() (*string, bool)`

GetMp25Ok returns a tuple with the Mp25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp25

`func (o *CreateContainerRequest) SetMp25(v string)`

SetMp25 sets Mp25 field to given value.

### HasMp25

`func (o *CreateContainerRequest) HasMp25() bool`

HasMp25 returns a boolean if a field has been set.

### GetMp26

`func (o *CreateContainerRequest) GetMp26() string`

GetMp26 returns the Mp26 field if non-nil, zero value otherwise.

### GetMp26Ok

`func (o *CreateContainerRequest) GetMp26Ok() (*string, bool)`

GetMp26Ok returns a tuple with the Mp26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp26

`func (o *CreateContainerRequest) SetMp26(v string)`

SetMp26 sets Mp26 field to given value.

### HasMp26

`func (o *CreateContainerRequest) HasMp26() bool`

HasMp26 returns a boolean if a field has been set.

### GetMp27

`func (o *CreateContainerRequest) GetMp27() string`

GetMp27 returns the Mp27 field if non-nil, zero value otherwise.

### GetMp27Ok

`func (o *CreateContainerRequest) GetMp27Ok() (*string, bool)`

GetMp27Ok returns a tuple with the Mp27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp27

`func (o *CreateContainerRequest) SetMp27(v string)`

SetMp27 sets Mp27 field to given value.

### HasMp27

`func (o *CreateContainerRequest) HasMp27() bool`

HasMp27 returns a boolean if a field has been set.

### GetMp28

`func (o *CreateContainerRequest) GetMp28() string`

GetMp28 returns the Mp28 field if non-nil, zero value otherwise.

### GetMp28Ok

`func (o *CreateContainerRequest) GetMp28Ok() (*string, bool)`

GetMp28Ok returns a tuple with the Mp28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp28

`func (o *CreateContainerRequest) SetMp28(v string)`

SetMp28 sets Mp28 field to given value.

### HasMp28

`func (o *CreateContainerRequest) HasMp28() bool`

HasMp28 returns a boolean if a field has been set.

### GetMp29

`func (o *CreateContainerRequest) GetMp29() string`

GetMp29 returns the Mp29 field if non-nil, zero value otherwise.

### GetMp29Ok

`func (o *CreateContainerRequest) GetMp29Ok() (*string, bool)`

GetMp29Ok returns a tuple with the Mp29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp29

`func (o *CreateContainerRequest) SetMp29(v string)`

SetMp29 sets Mp29 field to given value.

### HasMp29

`func (o *CreateContainerRequest) HasMp29() bool`

HasMp29 returns a boolean if a field has been set.

### GetMp30

`func (o *CreateContainerRequest) GetMp30() string`

GetMp30 returns the Mp30 field if non-nil, zero value otherwise.

### GetMp30Ok

`func (o *CreateContainerRequest) GetMp30Ok() (*string, bool)`

GetMp30Ok returns a tuple with the Mp30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp30

`func (o *CreateContainerRequest) SetMp30(v string)`

SetMp30 sets Mp30 field to given value.

### HasMp30

`func (o *CreateContainerRequest) HasMp30() bool`

HasMp30 returns a boolean if a field has been set.

### GetMp31

`func (o *CreateContainerRequest) GetMp31() string`

GetMp31 returns the Mp31 field if non-nil, zero value otherwise.

### GetMp31Ok

`func (o *CreateContainerRequest) GetMp31Ok() (*string, bool)`

GetMp31Ok returns a tuple with the Mp31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp31

`func (o *CreateContainerRequest) SetMp31(v string)`

SetMp31 sets Mp31 field to given value.

### HasMp31

`func (o *CreateContainerRequest) HasMp31() bool`

HasMp31 returns a boolean if a field has been set.

### GetMp32

`func (o *CreateContainerRequest) GetMp32() string`

GetMp32 returns the Mp32 field if non-nil, zero value otherwise.

### GetMp32Ok

`func (o *CreateContainerRequest) GetMp32Ok() (*string, bool)`

GetMp32Ok returns a tuple with the Mp32 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp32

`func (o *CreateContainerRequest) SetMp32(v string)`

SetMp32 sets Mp32 field to given value.

### HasMp32

`func (o *CreateContainerRequest) HasMp32() bool`

HasMp32 returns a boolean if a field has been set.

### GetMp33

`func (o *CreateContainerRequest) GetMp33() string`

GetMp33 returns the Mp33 field if non-nil, zero value otherwise.

### GetMp33Ok

`func (o *CreateContainerRequest) GetMp33Ok() (*string, bool)`

GetMp33Ok returns a tuple with the Mp33 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp33

`func (o *CreateContainerRequest) SetMp33(v string)`

SetMp33 sets Mp33 field to given value.

### HasMp33

`func (o *CreateContainerRequest) HasMp33() bool`

HasMp33 returns a boolean if a field has been set.

### GetMp34

`func (o *CreateContainerRequest) GetMp34() string`

GetMp34 returns the Mp34 field if non-nil, zero value otherwise.

### GetMp34Ok

`func (o *CreateContainerRequest) GetMp34Ok() (*string, bool)`

GetMp34Ok returns a tuple with the Mp34 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp34

`func (o *CreateContainerRequest) SetMp34(v string)`

SetMp34 sets Mp34 field to given value.

### HasMp34

`func (o *CreateContainerRequest) HasMp34() bool`

HasMp34 returns a boolean if a field has been set.

### GetMp35

`func (o *CreateContainerRequest) GetMp35() string`

GetMp35 returns the Mp35 field if non-nil, zero value otherwise.

### GetMp35Ok

`func (o *CreateContainerRequest) GetMp35Ok() (*string, bool)`

GetMp35Ok returns a tuple with the Mp35 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp35

`func (o *CreateContainerRequest) SetMp35(v string)`

SetMp35 sets Mp35 field to given value.

### HasMp35

`func (o *CreateContainerRequest) HasMp35() bool`

HasMp35 returns a boolean if a field has been set.

### GetMp36

`func (o *CreateContainerRequest) GetMp36() string`

GetMp36 returns the Mp36 field if non-nil, zero value otherwise.

### GetMp36Ok

`func (o *CreateContainerRequest) GetMp36Ok() (*string, bool)`

GetMp36Ok returns a tuple with the Mp36 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp36

`func (o *CreateContainerRequest) SetMp36(v string)`

SetMp36 sets Mp36 field to given value.

### HasMp36

`func (o *CreateContainerRequest) HasMp36() bool`

HasMp36 returns a boolean if a field has been set.

### GetMp37

`func (o *CreateContainerRequest) GetMp37() string`

GetMp37 returns the Mp37 field if non-nil, zero value otherwise.

### GetMp37Ok

`func (o *CreateContainerRequest) GetMp37Ok() (*string, bool)`

GetMp37Ok returns a tuple with the Mp37 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp37

`func (o *CreateContainerRequest) SetMp37(v string)`

SetMp37 sets Mp37 field to given value.

### HasMp37

`func (o *CreateContainerRequest) HasMp37() bool`

HasMp37 returns a boolean if a field has been set.

### GetMp38

`func (o *CreateContainerRequest) GetMp38() string`

GetMp38 returns the Mp38 field if non-nil, zero value otherwise.

### GetMp38Ok

`func (o *CreateContainerRequest) GetMp38Ok() (*string, bool)`

GetMp38Ok returns a tuple with the Mp38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp38

`func (o *CreateContainerRequest) SetMp38(v string)`

SetMp38 sets Mp38 field to given value.

### HasMp38

`func (o *CreateContainerRequest) HasMp38() bool`

HasMp38 returns a boolean if a field has been set.

### GetMp39

`func (o *CreateContainerRequest) GetMp39() string`

GetMp39 returns the Mp39 field if non-nil, zero value otherwise.

### GetMp39Ok

`func (o *CreateContainerRequest) GetMp39Ok() (*string, bool)`

GetMp39Ok returns a tuple with the Mp39 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp39

`func (o *CreateContainerRequest) SetMp39(v string)`

SetMp39 sets Mp39 field to given value.

### HasMp39

`func (o *CreateContainerRequest) HasMp39() bool`

HasMp39 returns a boolean if a field has been set.

### GetMp40

`func (o *CreateContainerRequest) GetMp40() string`

GetMp40 returns the Mp40 field if non-nil, zero value otherwise.

### GetMp40Ok

`func (o *CreateContainerRequest) GetMp40Ok() (*string, bool)`

GetMp40Ok returns a tuple with the Mp40 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp40

`func (o *CreateContainerRequest) SetMp40(v string)`

SetMp40 sets Mp40 field to given value.

### HasMp40

`func (o *CreateContainerRequest) HasMp40() bool`

HasMp40 returns a boolean if a field has been set.

### GetMp41

`func (o *CreateContainerRequest) GetMp41() string`

GetMp41 returns the Mp41 field if non-nil, zero value otherwise.

### GetMp41Ok

`func (o *CreateContainerRequest) GetMp41Ok() (*string, bool)`

GetMp41Ok returns a tuple with the Mp41 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp41

`func (o *CreateContainerRequest) SetMp41(v string)`

SetMp41 sets Mp41 field to given value.

### HasMp41

`func (o *CreateContainerRequest) HasMp41() bool`

HasMp41 returns a boolean if a field has been set.

### GetMp42

`func (o *CreateContainerRequest) GetMp42() string`

GetMp42 returns the Mp42 field if non-nil, zero value otherwise.

### GetMp42Ok

`func (o *CreateContainerRequest) GetMp42Ok() (*string, bool)`

GetMp42Ok returns a tuple with the Mp42 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp42

`func (o *CreateContainerRequest) SetMp42(v string)`

SetMp42 sets Mp42 field to given value.

### HasMp42

`func (o *CreateContainerRequest) HasMp42() bool`

HasMp42 returns a boolean if a field has been set.

### GetMp43

`func (o *CreateContainerRequest) GetMp43() string`

GetMp43 returns the Mp43 field if non-nil, zero value otherwise.

### GetMp43Ok

`func (o *CreateContainerRequest) GetMp43Ok() (*string, bool)`

GetMp43Ok returns a tuple with the Mp43 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp43

`func (o *CreateContainerRequest) SetMp43(v string)`

SetMp43 sets Mp43 field to given value.

### HasMp43

`func (o *CreateContainerRequest) HasMp43() bool`

HasMp43 returns a boolean if a field has been set.

### GetMp44

`func (o *CreateContainerRequest) GetMp44() string`

GetMp44 returns the Mp44 field if non-nil, zero value otherwise.

### GetMp44Ok

`func (o *CreateContainerRequest) GetMp44Ok() (*string, bool)`

GetMp44Ok returns a tuple with the Mp44 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp44

`func (o *CreateContainerRequest) SetMp44(v string)`

SetMp44 sets Mp44 field to given value.

### HasMp44

`func (o *CreateContainerRequest) HasMp44() bool`

HasMp44 returns a boolean if a field has been set.

### GetMp45

`func (o *CreateContainerRequest) GetMp45() string`

GetMp45 returns the Mp45 field if non-nil, zero value otherwise.

### GetMp45Ok

`func (o *CreateContainerRequest) GetMp45Ok() (*string, bool)`

GetMp45Ok returns a tuple with the Mp45 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp45

`func (o *CreateContainerRequest) SetMp45(v string)`

SetMp45 sets Mp45 field to given value.

### HasMp45

`func (o *CreateContainerRequest) HasMp45() bool`

HasMp45 returns a boolean if a field has been set.

### GetMp46

`func (o *CreateContainerRequest) GetMp46() string`

GetMp46 returns the Mp46 field if non-nil, zero value otherwise.

### GetMp46Ok

`func (o *CreateContainerRequest) GetMp46Ok() (*string, bool)`

GetMp46Ok returns a tuple with the Mp46 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp46

`func (o *CreateContainerRequest) SetMp46(v string)`

SetMp46 sets Mp46 field to given value.

### HasMp46

`func (o *CreateContainerRequest) HasMp46() bool`

HasMp46 returns a boolean if a field has been set.

### GetMp47

`func (o *CreateContainerRequest) GetMp47() string`

GetMp47 returns the Mp47 field if non-nil, zero value otherwise.

### GetMp47Ok

`func (o *CreateContainerRequest) GetMp47Ok() (*string, bool)`

GetMp47Ok returns a tuple with the Mp47 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp47

`func (o *CreateContainerRequest) SetMp47(v string)`

SetMp47 sets Mp47 field to given value.

### HasMp47

`func (o *CreateContainerRequest) HasMp47() bool`

HasMp47 returns a boolean if a field has been set.

### GetMp48

`func (o *CreateContainerRequest) GetMp48() string`

GetMp48 returns the Mp48 field if non-nil, zero value otherwise.

### GetMp48Ok

`func (o *CreateContainerRequest) GetMp48Ok() (*string, bool)`

GetMp48Ok returns a tuple with the Mp48 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp48

`func (o *CreateContainerRequest) SetMp48(v string)`

SetMp48 sets Mp48 field to given value.

### HasMp48

`func (o *CreateContainerRequest) HasMp48() bool`

HasMp48 returns a boolean if a field has been set.

### GetMp49

`func (o *CreateContainerRequest) GetMp49() string`

GetMp49 returns the Mp49 field if non-nil, zero value otherwise.

### GetMp49Ok

`func (o *CreateContainerRequest) GetMp49Ok() (*string, bool)`

GetMp49Ok returns a tuple with the Mp49 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp49

`func (o *CreateContainerRequest) SetMp49(v string)`

SetMp49 sets Mp49 field to given value.

### HasMp49

`func (o *CreateContainerRequest) HasMp49() bool`

HasMp49 returns a boolean if a field has been set.

### GetMp50

`func (o *CreateContainerRequest) GetMp50() string`

GetMp50 returns the Mp50 field if non-nil, zero value otherwise.

### GetMp50Ok

`func (o *CreateContainerRequest) GetMp50Ok() (*string, bool)`

GetMp50Ok returns a tuple with the Mp50 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp50

`func (o *CreateContainerRequest) SetMp50(v string)`

SetMp50 sets Mp50 field to given value.

### HasMp50

`func (o *CreateContainerRequest) HasMp50() bool`

HasMp50 returns a boolean if a field has been set.

### GetMp51

`func (o *CreateContainerRequest) GetMp51() string`

GetMp51 returns the Mp51 field if non-nil, zero value otherwise.

### GetMp51Ok

`func (o *CreateContainerRequest) GetMp51Ok() (*string, bool)`

GetMp51Ok returns a tuple with the Mp51 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp51

`func (o *CreateContainerRequest) SetMp51(v string)`

SetMp51 sets Mp51 field to given value.

### HasMp51

`func (o *CreateContainerRequest) HasMp51() bool`

HasMp51 returns a boolean if a field has been set.

### GetMp52

`func (o *CreateContainerRequest) GetMp52() string`

GetMp52 returns the Mp52 field if non-nil, zero value otherwise.

### GetMp52Ok

`func (o *CreateContainerRequest) GetMp52Ok() (*string, bool)`

GetMp52Ok returns a tuple with the Mp52 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp52

`func (o *CreateContainerRequest) SetMp52(v string)`

SetMp52 sets Mp52 field to given value.

### HasMp52

`func (o *CreateContainerRequest) HasMp52() bool`

HasMp52 returns a boolean if a field has been set.

### GetMp53

`func (o *CreateContainerRequest) GetMp53() string`

GetMp53 returns the Mp53 field if non-nil, zero value otherwise.

### GetMp53Ok

`func (o *CreateContainerRequest) GetMp53Ok() (*string, bool)`

GetMp53Ok returns a tuple with the Mp53 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp53

`func (o *CreateContainerRequest) SetMp53(v string)`

SetMp53 sets Mp53 field to given value.

### HasMp53

`func (o *CreateContainerRequest) HasMp53() bool`

HasMp53 returns a boolean if a field has been set.

### GetMp54

`func (o *CreateContainerRequest) GetMp54() string`

GetMp54 returns the Mp54 field if non-nil, zero value otherwise.

### GetMp54Ok

`func (o *CreateContainerRequest) GetMp54Ok() (*string, bool)`

GetMp54Ok returns a tuple with the Mp54 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp54

`func (o *CreateContainerRequest) SetMp54(v string)`

SetMp54 sets Mp54 field to given value.

### HasMp54

`func (o *CreateContainerRequest) HasMp54() bool`

HasMp54 returns a boolean if a field has been set.

### GetMp55

`func (o *CreateContainerRequest) GetMp55() string`

GetMp55 returns the Mp55 field if non-nil, zero value otherwise.

### GetMp55Ok

`func (o *CreateContainerRequest) GetMp55Ok() (*string, bool)`

GetMp55Ok returns a tuple with the Mp55 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp55

`func (o *CreateContainerRequest) SetMp55(v string)`

SetMp55 sets Mp55 field to given value.

### HasMp55

`func (o *CreateContainerRequest) HasMp55() bool`

HasMp55 returns a boolean if a field has been set.

### GetMp56

`func (o *CreateContainerRequest) GetMp56() string`

GetMp56 returns the Mp56 field if non-nil, zero value otherwise.

### GetMp56Ok

`func (o *CreateContainerRequest) GetMp56Ok() (*string, bool)`

GetMp56Ok returns a tuple with the Mp56 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp56

`func (o *CreateContainerRequest) SetMp56(v string)`

SetMp56 sets Mp56 field to given value.

### HasMp56

`func (o *CreateContainerRequest) HasMp56() bool`

HasMp56 returns a boolean if a field has been set.

### GetMp57

`func (o *CreateContainerRequest) GetMp57() string`

GetMp57 returns the Mp57 field if non-nil, zero value otherwise.

### GetMp57Ok

`func (o *CreateContainerRequest) GetMp57Ok() (*string, bool)`

GetMp57Ok returns a tuple with the Mp57 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp57

`func (o *CreateContainerRequest) SetMp57(v string)`

SetMp57 sets Mp57 field to given value.

### HasMp57

`func (o *CreateContainerRequest) HasMp57() bool`

HasMp57 returns a boolean if a field has been set.

### GetMp58

`func (o *CreateContainerRequest) GetMp58() string`

GetMp58 returns the Mp58 field if non-nil, zero value otherwise.

### GetMp58Ok

`func (o *CreateContainerRequest) GetMp58Ok() (*string, bool)`

GetMp58Ok returns a tuple with the Mp58 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp58

`func (o *CreateContainerRequest) SetMp58(v string)`

SetMp58 sets Mp58 field to given value.

### HasMp58

`func (o *CreateContainerRequest) HasMp58() bool`

HasMp58 returns a boolean if a field has been set.

### GetMp59

`func (o *CreateContainerRequest) GetMp59() string`

GetMp59 returns the Mp59 field if non-nil, zero value otherwise.

### GetMp59Ok

`func (o *CreateContainerRequest) GetMp59Ok() (*string, bool)`

GetMp59Ok returns a tuple with the Mp59 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp59

`func (o *CreateContainerRequest) SetMp59(v string)`

SetMp59 sets Mp59 field to given value.

### HasMp59

`func (o *CreateContainerRequest) HasMp59() bool`

HasMp59 returns a boolean if a field has been set.

### GetMp60

`func (o *CreateContainerRequest) GetMp60() string`

GetMp60 returns the Mp60 field if non-nil, zero value otherwise.

### GetMp60Ok

`func (o *CreateContainerRequest) GetMp60Ok() (*string, bool)`

GetMp60Ok returns a tuple with the Mp60 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp60

`func (o *CreateContainerRequest) SetMp60(v string)`

SetMp60 sets Mp60 field to given value.

### HasMp60

`func (o *CreateContainerRequest) HasMp60() bool`

HasMp60 returns a boolean if a field has been set.

### GetMp61

`func (o *CreateContainerRequest) GetMp61() string`

GetMp61 returns the Mp61 field if non-nil, zero value otherwise.

### GetMp61Ok

`func (o *CreateContainerRequest) GetMp61Ok() (*string, bool)`

GetMp61Ok returns a tuple with the Mp61 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp61

`func (o *CreateContainerRequest) SetMp61(v string)`

SetMp61 sets Mp61 field to given value.

### HasMp61

`func (o *CreateContainerRequest) HasMp61() bool`

HasMp61 returns a boolean if a field has been set.

### GetMp62

`func (o *CreateContainerRequest) GetMp62() string`

GetMp62 returns the Mp62 field if non-nil, zero value otherwise.

### GetMp62Ok

`func (o *CreateContainerRequest) GetMp62Ok() (*string, bool)`

GetMp62Ok returns a tuple with the Mp62 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp62

`func (o *CreateContainerRequest) SetMp62(v string)`

SetMp62 sets Mp62 field to given value.

### HasMp62

`func (o *CreateContainerRequest) HasMp62() bool`

HasMp62 returns a boolean if a field has been set.

### GetMp63

`func (o *CreateContainerRequest) GetMp63() string`

GetMp63 returns the Mp63 field if non-nil, zero value otherwise.

### GetMp63Ok

`func (o *CreateContainerRequest) GetMp63Ok() (*string, bool)`

GetMp63Ok returns a tuple with the Mp63 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp63

`func (o *CreateContainerRequest) SetMp63(v string)`

SetMp63 sets Mp63 field to given value.

### HasMp63

`func (o *CreateContainerRequest) HasMp63() bool`

HasMp63 returns a boolean if a field has been set.

### GetMp64

`func (o *CreateContainerRequest) GetMp64() string`

GetMp64 returns the Mp64 field if non-nil, zero value otherwise.

### GetMp64Ok

`func (o *CreateContainerRequest) GetMp64Ok() (*string, bool)`

GetMp64Ok returns a tuple with the Mp64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp64

`func (o *CreateContainerRequest) SetMp64(v string)`

SetMp64 sets Mp64 field to given value.

### HasMp64

`func (o *CreateContainerRequest) HasMp64() bool`

HasMp64 returns a boolean if a field has been set.

### GetMp65

`func (o *CreateContainerRequest) GetMp65() string`

GetMp65 returns the Mp65 field if non-nil, zero value otherwise.

### GetMp65Ok

`func (o *CreateContainerRequest) GetMp65Ok() (*string, bool)`

GetMp65Ok returns a tuple with the Mp65 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp65

`func (o *CreateContainerRequest) SetMp65(v string)`

SetMp65 sets Mp65 field to given value.

### HasMp65

`func (o *CreateContainerRequest) HasMp65() bool`

HasMp65 returns a boolean if a field has been set.

### GetMp66

`func (o *CreateContainerRequest) GetMp66() string`

GetMp66 returns the Mp66 field if non-nil, zero value otherwise.

### GetMp66Ok

`func (o *CreateContainerRequest) GetMp66Ok() (*string, bool)`

GetMp66Ok returns a tuple with the Mp66 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp66

`func (o *CreateContainerRequest) SetMp66(v string)`

SetMp66 sets Mp66 field to given value.

### HasMp66

`func (o *CreateContainerRequest) HasMp66() bool`

HasMp66 returns a boolean if a field has been set.

### GetMp67

`func (o *CreateContainerRequest) GetMp67() string`

GetMp67 returns the Mp67 field if non-nil, zero value otherwise.

### GetMp67Ok

`func (o *CreateContainerRequest) GetMp67Ok() (*string, bool)`

GetMp67Ok returns a tuple with the Mp67 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp67

`func (o *CreateContainerRequest) SetMp67(v string)`

SetMp67 sets Mp67 field to given value.

### HasMp67

`func (o *CreateContainerRequest) HasMp67() bool`

HasMp67 returns a boolean if a field has been set.

### GetMp68

`func (o *CreateContainerRequest) GetMp68() string`

GetMp68 returns the Mp68 field if non-nil, zero value otherwise.

### GetMp68Ok

`func (o *CreateContainerRequest) GetMp68Ok() (*string, bool)`

GetMp68Ok returns a tuple with the Mp68 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp68

`func (o *CreateContainerRequest) SetMp68(v string)`

SetMp68 sets Mp68 field to given value.

### HasMp68

`func (o *CreateContainerRequest) HasMp68() bool`

HasMp68 returns a boolean if a field has been set.

### GetMp69

`func (o *CreateContainerRequest) GetMp69() string`

GetMp69 returns the Mp69 field if non-nil, zero value otherwise.

### GetMp69Ok

`func (o *CreateContainerRequest) GetMp69Ok() (*string, bool)`

GetMp69Ok returns a tuple with the Mp69 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp69

`func (o *CreateContainerRequest) SetMp69(v string)`

SetMp69 sets Mp69 field to given value.

### HasMp69

`func (o *CreateContainerRequest) HasMp69() bool`

HasMp69 returns a boolean if a field has been set.

### GetMp70

`func (o *CreateContainerRequest) GetMp70() string`

GetMp70 returns the Mp70 field if non-nil, zero value otherwise.

### GetMp70Ok

`func (o *CreateContainerRequest) GetMp70Ok() (*string, bool)`

GetMp70Ok returns a tuple with the Mp70 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp70

`func (o *CreateContainerRequest) SetMp70(v string)`

SetMp70 sets Mp70 field to given value.

### HasMp70

`func (o *CreateContainerRequest) HasMp70() bool`

HasMp70 returns a boolean if a field has been set.

### GetMp71

`func (o *CreateContainerRequest) GetMp71() string`

GetMp71 returns the Mp71 field if non-nil, zero value otherwise.

### GetMp71Ok

`func (o *CreateContainerRequest) GetMp71Ok() (*string, bool)`

GetMp71Ok returns a tuple with the Mp71 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp71

`func (o *CreateContainerRequest) SetMp71(v string)`

SetMp71 sets Mp71 field to given value.

### HasMp71

`func (o *CreateContainerRequest) HasMp71() bool`

HasMp71 returns a boolean if a field has been set.

### GetMp72

`func (o *CreateContainerRequest) GetMp72() string`

GetMp72 returns the Mp72 field if non-nil, zero value otherwise.

### GetMp72Ok

`func (o *CreateContainerRequest) GetMp72Ok() (*string, bool)`

GetMp72Ok returns a tuple with the Mp72 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp72

`func (o *CreateContainerRequest) SetMp72(v string)`

SetMp72 sets Mp72 field to given value.

### HasMp72

`func (o *CreateContainerRequest) HasMp72() bool`

HasMp72 returns a boolean if a field has been set.

### GetMp73

`func (o *CreateContainerRequest) GetMp73() string`

GetMp73 returns the Mp73 field if non-nil, zero value otherwise.

### GetMp73Ok

`func (o *CreateContainerRequest) GetMp73Ok() (*string, bool)`

GetMp73Ok returns a tuple with the Mp73 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp73

`func (o *CreateContainerRequest) SetMp73(v string)`

SetMp73 sets Mp73 field to given value.

### HasMp73

`func (o *CreateContainerRequest) HasMp73() bool`

HasMp73 returns a boolean if a field has been set.

### GetMp74

`func (o *CreateContainerRequest) GetMp74() string`

GetMp74 returns the Mp74 field if non-nil, zero value otherwise.

### GetMp74Ok

`func (o *CreateContainerRequest) GetMp74Ok() (*string, bool)`

GetMp74Ok returns a tuple with the Mp74 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp74

`func (o *CreateContainerRequest) SetMp74(v string)`

SetMp74 sets Mp74 field to given value.

### HasMp74

`func (o *CreateContainerRequest) HasMp74() bool`

HasMp74 returns a boolean if a field has been set.

### GetMp75

`func (o *CreateContainerRequest) GetMp75() string`

GetMp75 returns the Mp75 field if non-nil, zero value otherwise.

### GetMp75Ok

`func (o *CreateContainerRequest) GetMp75Ok() (*string, bool)`

GetMp75Ok returns a tuple with the Mp75 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp75

`func (o *CreateContainerRequest) SetMp75(v string)`

SetMp75 sets Mp75 field to given value.

### HasMp75

`func (o *CreateContainerRequest) HasMp75() bool`

HasMp75 returns a boolean if a field has been set.

### GetMp76

`func (o *CreateContainerRequest) GetMp76() string`

GetMp76 returns the Mp76 field if non-nil, zero value otherwise.

### GetMp76Ok

`func (o *CreateContainerRequest) GetMp76Ok() (*string, bool)`

GetMp76Ok returns a tuple with the Mp76 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp76

`func (o *CreateContainerRequest) SetMp76(v string)`

SetMp76 sets Mp76 field to given value.

### HasMp76

`func (o *CreateContainerRequest) HasMp76() bool`

HasMp76 returns a boolean if a field has been set.

### GetMp77

`func (o *CreateContainerRequest) GetMp77() string`

GetMp77 returns the Mp77 field if non-nil, zero value otherwise.

### GetMp77Ok

`func (o *CreateContainerRequest) GetMp77Ok() (*string, bool)`

GetMp77Ok returns a tuple with the Mp77 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp77

`func (o *CreateContainerRequest) SetMp77(v string)`

SetMp77 sets Mp77 field to given value.

### HasMp77

`func (o *CreateContainerRequest) HasMp77() bool`

HasMp77 returns a boolean if a field has been set.

### GetMp78

`func (o *CreateContainerRequest) GetMp78() string`

GetMp78 returns the Mp78 field if non-nil, zero value otherwise.

### GetMp78Ok

`func (o *CreateContainerRequest) GetMp78Ok() (*string, bool)`

GetMp78Ok returns a tuple with the Mp78 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp78

`func (o *CreateContainerRequest) SetMp78(v string)`

SetMp78 sets Mp78 field to given value.

### HasMp78

`func (o *CreateContainerRequest) HasMp78() bool`

HasMp78 returns a boolean if a field has been set.

### GetMp79

`func (o *CreateContainerRequest) GetMp79() string`

GetMp79 returns the Mp79 field if non-nil, zero value otherwise.

### GetMp79Ok

`func (o *CreateContainerRequest) GetMp79Ok() (*string, bool)`

GetMp79Ok returns a tuple with the Mp79 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp79

`func (o *CreateContainerRequest) SetMp79(v string)`

SetMp79 sets Mp79 field to given value.

### HasMp79

`func (o *CreateContainerRequest) HasMp79() bool`

HasMp79 returns a boolean if a field has been set.

### GetMp80

`func (o *CreateContainerRequest) GetMp80() string`

GetMp80 returns the Mp80 field if non-nil, zero value otherwise.

### GetMp80Ok

`func (o *CreateContainerRequest) GetMp80Ok() (*string, bool)`

GetMp80Ok returns a tuple with the Mp80 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp80

`func (o *CreateContainerRequest) SetMp80(v string)`

SetMp80 sets Mp80 field to given value.

### HasMp80

`func (o *CreateContainerRequest) HasMp80() bool`

HasMp80 returns a boolean if a field has been set.

### GetMp81

`func (o *CreateContainerRequest) GetMp81() string`

GetMp81 returns the Mp81 field if non-nil, zero value otherwise.

### GetMp81Ok

`func (o *CreateContainerRequest) GetMp81Ok() (*string, bool)`

GetMp81Ok returns a tuple with the Mp81 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp81

`func (o *CreateContainerRequest) SetMp81(v string)`

SetMp81 sets Mp81 field to given value.

### HasMp81

`func (o *CreateContainerRequest) HasMp81() bool`

HasMp81 returns a boolean if a field has been set.

### GetMp82

`func (o *CreateContainerRequest) GetMp82() string`

GetMp82 returns the Mp82 field if non-nil, zero value otherwise.

### GetMp82Ok

`func (o *CreateContainerRequest) GetMp82Ok() (*string, bool)`

GetMp82Ok returns a tuple with the Mp82 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp82

`func (o *CreateContainerRequest) SetMp82(v string)`

SetMp82 sets Mp82 field to given value.

### HasMp82

`func (o *CreateContainerRequest) HasMp82() bool`

HasMp82 returns a boolean if a field has been set.

### GetMp83

`func (o *CreateContainerRequest) GetMp83() string`

GetMp83 returns the Mp83 field if non-nil, zero value otherwise.

### GetMp83Ok

`func (o *CreateContainerRequest) GetMp83Ok() (*string, bool)`

GetMp83Ok returns a tuple with the Mp83 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp83

`func (o *CreateContainerRequest) SetMp83(v string)`

SetMp83 sets Mp83 field to given value.

### HasMp83

`func (o *CreateContainerRequest) HasMp83() bool`

HasMp83 returns a boolean if a field has been set.

### GetMp84

`func (o *CreateContainerRequest) GetMp84() string`

GetMp84 returns the Mp84 field if non-nil, zero value otherwise.

### GetMp84Ok

`func (o *CreateContainerRequest) GetMp84Ok() (*string, bool)`

GetMp84Ok returns a tuple with the Mp84 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp84

`func (o *CreateContainerRequest) SetMp84(v string)`

SetMp84 sets Mp84 field to given value.

### HasMp84

`func (o *CreateContainerRequest) HasMp84() bool`

HasMp84 returns a boolean if a field has been set.

### GetMp85

`func (o *CreateContainerRequest) GetMp85() string`

GetMp85 returns the Mp85 field if non-nil, zero value otherwise.

### GetMp85Ok

`func (o *CreateContainerRequest) GetMp85Ok() (*string, bool)`

GetMp85Ok returns a tuple with the Mp85 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp85

`func (o *CreateContainerRequest) SetMp85(v string)`

SetMp85 sets Mp85 field to given value.

### HasMp85

`func (o *CreateContainerRequest) HasMp85() bool`

HasMp85 returns a boolean if a field has been set.

### GetMp86

`func (o *CreateContainerRequest) GetMp86() string`

GetMp86 returns the Mp86 field if non-nil, zero value otherwise.

### GetMp86Ok

`func (o *CreateContainerRequest) GetMp86Ok() (*string, bool)`

GetMp86Ok returns a tuple with the Mp86 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp86

`func (o *CreateContainerRequest) SetMp86(v string)`

SetMp86 sets Mp86 field to given value.

### HasMp86

`func (o *CreateContainerRequest) HasMp86() bool`

HasMp86 returns a boolean if a field has been set.

### GetMp87

`func (o *CreateContainerRequest) GetMp87() string`

GetMp87 returns the Mp87 field if non-nil, zero value otherwise.

### GetMp87Ok

`func (o *CreateContainerRequest) GetMp87Ok() (*string, bool)`

GetMp87Ok returns a tuple with the Mp87 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp87

`func (o *CreateContainerRequest) SetMp87(v string)`

SetMp87 sets Mp87 field to given value.

### HasMp87

`func (o *CreateContainerRequest) HasMp87() bool`

HasMp87 returns a boolean if a field has been set.

### GetMp88

`func (o *CreateContainerRequest) GetMp88() string`

GetMp88 returns the Mp88 field if non-nil, zero value otherwise.

### GetMp88Ok

`func (o *CreateContainerRequest) GetMp88Ok() (*string, bool)`

GetMp88Ok returns a tuple with the Mp88 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp88

`func (o *CreateContainerRequest) SetMp88(v string)`

SetMp88 sets Mp88 field to given value.

### HasMp88

`func (o *CreateContainerRequest) HasMp88() bool`

HasMp88 returns a boolean if a field has been set.

### GetMp89

`func (o *CreateContainerRequest) GetMp89() string`

GetMp89 returns the Mp89 field if non-nil, zero value otherwise.

### GetMp89Ok

`func (o *CreateContainerRequest) GetMp89Ok() (*string, bool)`

GetMp89Ok returns a tuple with the Mp89 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp89

`func (o *CreateContainerRequest) SetMp89(v string)`

SetMp89 sets Mp89 field to given value.

### HasMp89

`func (o *CreateContainerRequest) HasMp89() bool`

HasMp89 returns a boolean if a field has been set.

### GetMp90

`func (o *CreateContainerRequest) GetMp90() string`

GetMp90 returns the Mp90 field if non-nil, zero value otherwise.

### GetMp90Ok

`func (o *CreateContainerRequest) GetMp90Ok() (*string, bool)`

GetMp90Ok returns a tuple with the Mp90 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp90

`func (o *CreateContainerRequest) SetMp90(v string)`

SetMp90 sets Mp90 field to given value.

### HasMp90

`func (o *CreateContainerRequest) HasMp90() bool`

HasMp90 returns a boolean if a field has been set.

### GetMp91

`func (o *CreateContainerRequest) GetMp91() string`

GetMp91 returns the Mp91 field if non-nil, zero value otherwise.

### GetMp91Ok

`func (o *CreateContainerRequest) GetMp91Ok() (*string, bool)`

GetMp91Ok returns a tuple with the Mp91 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp91

`func (o *CreateContainerRequest) SetMp91(v string)`

SetMp91 sets Mp91 field to given value.

### HasMp91

`func (o *CreateContainerRequest) HasMp91() bool`

HasMp91 returns a boolean if a field has been set.

### GetMp92

`func (o *CreateContainerRequest) GetMp92() string`

GetMp92 returns the Mp92 field if non-nil, zero value otherwise.

### GetMp92Ok

`func (o *CreateContainerRequest) GetMp92Ok() (*string, bool)`

GetMp92Ok returns a tuple with the Mp92 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp92

`func (o *CreateContainerRequest) SetMp92(v string)`

SetMp92 sets Mp92 field to given value.

### HasMp92

`func (o *CreateContainerRequest) HasMp92() bool`

HasMp92 returns a boolean if a field has been set.

### GetMp93

`func (o *CreateContainerRequest) GetMp93() string`

GetMp93 returns the Mp93 field if non-nil, zero value otherwise.

### GetMp93Ok

`func (o *CreateContainerRequest) GetMp93Ok() (*string, bool)`

GetMp93Ok returns a tuple with the Mp93 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp93

`func (o *CreateContainerRequest) SetMp93(v string)`

SetMp93 sets Mp93 field to given value.

### HasMp93

`func (o *CreateContainerRequest) HasMp93() bool`

HasMp93 returns a boolean if a field has been set.

### GetMp94

`func (o *CreateContainerRequest) GetMp94() string`

GetMp94 returns the Mp94 field if non-nil, zero value otherwise.

### GetMp94Ok

`func (o *CreateContainerRequest) GetMp94Ok() (*string, bool)`

GetMp94Ok returns a tuple with the Mp94 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp94

`func (o *CreateContainerRequest) SetMp94(v string)`

SetMp94 sets Mp94 field to given value.

### HasMp94

`func (o *CreateContainerRequest) HasMp94() bool`

HasMp94 returns a boolean if a field has been set.

### GetMp95

`func (o *CreateContainerRequest) GetMp95() string`

GetMp95 returns the Mp95 field if non-nil, zero value otherwise.

### GetMp95Ok

`func (o *CreateContainerRequest) GetMp95Ok() (*string, bool)`

GetMp95Ok returns a tuple with the Mp95 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp95

`func (o *CreateContainerRequest) SetMp95(v string)`

SetMp95 sets Mp95 field to given value.

### HasMp95

`func (o *CreateContainerRequest) HasMp95() bool`

HasMp95 returns a boolean if a field has been set.

### GetMp96

`func (o *CreateContainerRequest) GetMp96() string`

GetMp96 returns the Mp96 field if non-nil, zero value otherwise.

### GetMp96Ok

`func (o *CreateContainerRequest) GetMp96Ok() (*string, bool)`

GetMp96Ok returns a tuple with the Mp96 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp96

`func (o *CreateContainerRequest) SetMp96(v string)`

SetMp96 sets Mp96 field to given value.

### HasMp96

`func (o *CreateContainerRequest) HasMp96() bool`

HasMp96 returns a boolean if a field has been set.

### GetMp97

`func (o *CreateContainerRequest) GetMp97() string`

GetMp97 returns the Mp97 field if non-nil, zero value otherwise.

### GetMp97Ok

`func (o *CreateContainerRequest) GetMp97Ok() (*string, bool)`

GetMp97Ok returns a tuple with the Mp97 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp97

`func (o *CreateContainerRequest) SetMp97(v string)`

SetMp97 sets Mp97 field to given value.

### HasMp97

`func (o *CreateContainerRequest) HasMp97() bool`

HasMp97 returns a boolean if a field has been set.

### GetMp98

`func (o *CreateContainerRequest) GetMp98() string`

GetMp98 returns the Mp98 field if non-nil, zero value otherwise.

### GetMp98Ok

`func (o *CreateContainerRequest) GetMp98Ok() (*string, bool)`

GetMp98Ok returns a tuple with the Mp98 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp98

`func (o *CreateContainerRequest) SetMp98(v string)`

SetMp98 sets Mp98 field to given value.

### HasMp98

`func (o *CreateContainerRequest) HasMp98() bool`

HasMp98 returns a boolean if a field has been set.

### GetMp99

`func (o *CreateContainerRequest) GetMp99() string`

GetMp99 returns the Mp99 field if non-nil, zero value otherwise.

### GetMp99Ok

`func (o *CreateContainerRequest) GetMp99Ok() (*string, bool)`

GetMp99Ok returns a tuple with the Mp99 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp99

`func (o *CreateContainerRequest) SetMp99(v string)`

SetMp99 sets Mp99 field to given value.

### HasMp99

`func (o *CreateContainerRequest) HasMp99() bool`

HasMp99 returns a boolean if a field has been set.

### GetMp100

`func (o *CreateContainerRequest) GetMp100() string`

GetMp100 returns the Mp100 field if non-nil, zero value otherwise.

### GetMp100Ok

`func (o *CreateContainerRequest) GetMp100Ok() (*string, bool)`

GetMp100Ok returns a tuple with the Mp100 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp100

`func (o *CreateContainerRequest) SetMp100(v string)`

SetMp100 sets Mp100 field to given value.

### HasMp100

`func (o *CreateContainerRequest) HasMp100() bool`

HasMp100 returns a boolean if a field has been set.

### GetMp101

`func (o *CreateContainerRequest) GetMp101() string`

GetMp101 returns the Mp101 field if non-nil, zero value otherwise.

### GetMp101Ok

`func (o *CreateContainerRequest) GetMp101Ok() (*string, bool)`

GetMp101Ok returns a tuple with the Mp101 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp101

`func (o *CreateContainerRequest) SetMp101(v string)`

SetMp101 sets Mp101 field to given value.

### HasMp101

`func (o *CreateContainerRequest) HasMp101() bool`

HasMp101 returns a boolean if a field has been set.

### GetMp102

`func (o *CreateContainerRequest) GetMp102() string`

GetMp102 returns the Mp102 field if non-nil, zero value otherwise.

### GetMp102Ok

`func (o *CreateContainerRequest) GetMp102Ok() (*string, bool)`

GetMp102Ok returns a tuple with the Mp102 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp102

`func (o *CreateContainerRequest) SetMp102(v string)`

SetMp102 sets Mp102 field to given value.

### HasMp102

`func (o *CreateContainerRequest) HasMp102() bool`

HasMp102 returns a boolean if a field has been set.

### GetMp103

`func (o *CreateContainerRequest) GetMp103() string`

GetMp103 returns the Mp103 field if non-nil, zero value otherwise.

### GetMp103Ok

`func (o *CreateContainerRequest) GetMp103Ok() (*string, bool)`

GetMp103Ok returns a tuple with the Mp103 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp103

`func (o *CreateContainerRequest) SetMp103(v string)`

SetMp103 sets Mp103 field to given value.

### HasMp103

`func (o *CreateContainerRequest) HasMp103() bool`

HasMp103 returns a boolean if a field has been set.

### GetMp104

`func (o *CreateContainerRequest) GetMp104() string`

GetMp104 returns the Mp104 field if non-nil, zero value otherwise.

### GetMp104Ok

`func (o *CreateContainerRequest) GetMp104Ok() (*string, bool)`

GetMp104Ok returns a tuple with the Mp104 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp104

`func (o *CreateContainerRequest) SetMp104(v string)`

SetMp104 sets Mp104 field to given value.

### HasMp104

`func (o *CreateContainerRequest) HasMp104() bool`

HasMp104 returns a boolean if a field has been set.

### GetMp105

`func (o *CreateContainerRequest) GetMp105() string`

GetMp105 returns the Mp105 field if non-nil, zero value otherwise.

### GetMp105Ok

`func (o *CreateContainerRequest) GetMp105Ok() (*string, bool)`

GetMp105Ok returns a tuple with the Mp105 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp105

`func (o *CreateContainerRequest) SetMp105(v string)`

SetMp105 sets Mp105 field to given value.

### HasMp105

`func (o *CreateContainerRequest) HasMp105() bool`

HasMp105 returns a boolean if a field has been set.

### GetMp106

`func (o *CreateContainerRequest) GetMp106() string`

GetMp106 returns the Mp106 field if non-nil, zero value otherwise.

### GetMp106Ok

`func (o *CreateContainerRequest) GetMp106Ok() (*string, bool)`

GetMp106Ok returns a tuple with the Mp106 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp106

`func (o *CreateContainerRequest) SetMp106(v string)`

SetMp106 sets Mp106 field to given value.

### HasMp106

`func (o *CreateContainerRequest) HasMp106() bool`

HasMp106 returns a boolean if a field has been set.

### GetMp107

`func (o *CreateContainerRequest) GetMp107() string`

GetMp107 returns the Mp107 field if non-nil, zero value otherwise.

### GetMp107Ok

`func (o *CreateContainerRequest) GetMp107Ok() (*string, bool)`

GetMp107Ok returns a tuple with the Mp107 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp107

`func (o *CreateContainerRequest) SetMp107(v string)`

SetMp107 sets Mp107 field to given value.

### HasMp107

`func (o *CreateContainerRequest) HasMp107() bool`

HasMp107 returns a boolean if a field has been set.

### GetMp108

`func (o *CreateContainerRequest) GetMp108() string`

GetMp108 returns the Mp108 field if non-nil, zero value otherwise.

### GetMp108Ok

`func (o *CreateContainerRequest) GetMp108Ok() (*string, bool)`

GetMp108Ok returns a tuple with the Mp108 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp108

`func (o *CreateContainerRequest) SetMp108(v string)`

SetMp108 sets Mp108 field to given value.

### HasMp108

`func (o *CreateContainerRequest) HasMp108() bool`

HasMp108 returns a boolean if a field has been set.

### GetMp109

`func (o *CreateContainerRequest) GetMp109() string`

GetMp109 returns the Mp109 field if non-nil, zero value otherwise.

### GetMp109Ok

`func (o *CreateContainerRequest) GetMp109Ok() (*string, bool)`

GetMp109Ok returns a tuple with the Mp109 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp109

`func (o *CreateContainerRequest) SetMp109(v string)`

SetMp109 sets Mp109 field to given value.

### HasMp109

`func (o *CreateContainerRequest) HasMp109() bool`

HasMp109 returns a boolean if a field has been set.

### GetMp110

`func (o *CreateContainerRequest) GetMp110() string`

GetMp110 returns the Mp110 field if non-nil, zero value otherwise.

### GetMp110Ok

`func (o *CreateContainerRequest) GetMp110Ok() (*string, bool)`

GetMp110Ok returns a tuple with the Mp110 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp110

`func (o *CreateContainerRequest) SetMp110(v string)`

SetMp110 sets Mp110 field to given value.

### HasMp110

`func (o *CreateContainerRequest) HasMp110() bool`

HasMp110 returns a boolean if a field has been set.

### GetMp111

`func (o *CreateContainerRequest) GetMp111() string`

GetMp111 returns the Mp111 field if non-nil, zero value otherwise.

### GetMp111Ok

`func (o *CreateContainerRequest) GetMp111Ok() (*string, bool)`

GetMp111Ok returns a tuple with the Mp111 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp111

`func (o *CreateContainerRequest) SetMp111(v string)`

SetMp111 sets Mp111 field to given value.

### HasMp111

`func (o *CreateContainerRequest) HasMp111() bool`

HasMp111 returns a boolean if a field has been set.

### GetMp112

`func (o *CreateContainerRequest) GetMp112() string`

GetMp112 returns the Mp112 field if non-nil, zero value otherwise.

### GetMp112Ok

`func (o *CreateContainerRequest) GetMp112Ok() (*string, bool)`

GetMp112Ok returns a tuple with the Mp112 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp112

`func (o *CreateContainerRequest) SetMp112(v string)`

SetMp112 sets Mp112 field to given value.

### HasMp112

`func (o *CreateContainerRequest) HasMp112() bool`

HasMp112 returns a boolean if a field has been set.

### GetMp113

`func (o *CreateContainerRequest) GetMp113() string`

GetMp113 returns the Mp113 field if non-nil, zero value otherwise.

### GetMp113Ok

`func (o *CreateContainerRequest) GetMp113Ok() (*string, bool)`

GetMp113Ok returns a tuple with the Mp113 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp113

`func (o *CreateContainerRequest) SetMp113(v string)`

SetMp113 sets Mp113 field to given value.

### HasMp113

`func (o *CreateContainerRequest) HasMp113() bool`

HasMp113 returns a boolean if a field has been set.

### GetMp114

`func (o *CreateContainerRequest) GetMp114() string`

GetMp114 returns the Mp114 field if non-nil, zero value otherwise.

### GetMp114Ok

`func (o *CreateContainerRequest) GetMp114Ok() (*string, bool)`

GetMp114Ok returns a tuple with the Mp114 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp114

`func (o *CreateContainerRequest) SetMp114(v string)`

SetMp114 sets Mp114 field to given value.

### HasMp114

`func (o *CreateContainerRequest) HasMp114() bool`

HasMp114 returns a boolean if a field has been set.

### GetMp115

`func (o *CreateContainerRequest) GetMp115() string`

GetMp115 returns the Mp115 field if non-nil, zero value otherwise.

### GetMp115Ok

`func (o *CreateContainerRequest) GetMp115Ok() (*string, bool)`

GetMp115Ok returns a tuple with the Mp115 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp115

`func (o *CreateContainerRequest) SetMp115(v string)`

SetMp115 sets Mp115 field to given value.

### HasMp115

`func (o *CreateContainerRequest) HasMp115() bool`

HasMp115 returns a boolean if a field has been set.

### GetMp116

`func (o *CreateContainerRequest) GetMp116() string`

GetMp116 returns the Mp116 field if non-nil, zero value otherwise.

### GetMp116Ok

`func (o *CreateContainerRequest) GetMp116Ok() (*string, bool)`

GetMp116Ok returns a tuple with the Mp116 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp116

`func (o *CreateContainerRequest) SetMp116(v string)`

SetMp116 sets Mp116 field to given value.

### HasMp116

`func (o *CreateContainerRequest) HasMp116() bool`

HasMp116 returns a boolean if a field has been set.

### GetMp117

`func (o *CreateContainerRequest) GetMp117() string`

GetMp117 returns the Mp117 field if non-nil, zero value otherwise.

### GetMp117Ok

`func (o *CreateContainerRequest) GetMp117Ok() (*string, bool)`

GetMp117Ok returns a tuple with the Mp117 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp117

`func (o *CreateContainerRequest) SetMp117(v string)`

SetMp117 sets Mp117 field to given value.

### HasMp117

`func (o *CreateContainerRequest) HasMp117() bool`

HasMp117 returns a boolean if a field has been set.

### GetMp118

`func (o *CreateContainerRequest) GetMp118() string`

GetMp118 returns the Mp118 field if non-nil, zero value otherwise.

### GetMp118Ok

`func (o *CreateContainerRequest) GetMp118Ok() (*string, bool)`

GetMp118Ok returns a tuple with the Mp118 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp118

`func (o *CreateContainerRequest) SetMp118(v string)`

SetMp118 sets Mp118 field to given value.

### HasMp118

`func (o *CreateContainerRequest) HasMp118() bool`

HasMp118 returns a boolean if a field has been set.

### GetMp119

`func (o *CreateContainerRequest) GetMp119() string`

GetMp119 returns the Mp119 field if non-nil, zero value otherwise.

### GetMp119Ok

`func (o *CreateContainerRequest) GetMp119Ok() (*string, bool)`

GetMp119Ok returns a tuple with the Mp119 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp119

`func (o *CreateContainerRequest) SetMp119(v string)`

SetMp119 sets Mp119 field to given value.

### HasMp119

`func (o *CreateContainerRequest) HasMp119() bool`

HasMp119 returns a boolean if a field has been set.

### GetMp120

`func (o *CreateContainerRequest) GetMp120() string`

GetMp120 returns the Mp120 field if non-nil, zero value otherwise.

### GetMp120Ok

`func (o *CreateContainerRequest) GetMp120Ok() (*string, bool)`

GetMp120Ok returns a tuple with the Mp120 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp120

`func (o *CreateContainerRequest) SetMp120(v string)`

SetMp120 sets Mp120 field to given value.

### HasMp120

`func (o *CreateContainerRequest) HasMp120() bool`

HasMp120 returns a boolean if a field has been set.

### GetMp121

`func (o *CreateContainerRequest) GetMp121() string`

GetMp121 returns the Mp121 field if non-nil, zero value otherwise.

### GetMp121Ok

`func (o *CreateContainerRequest) GetMp121Ok() (*string, bool)`

GetMp121Ok returns a tuple with the Mp121 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp121

`func (o *CreateContainerRequest) SetMp121(v string)`

SetMp121 sets Mp121 field to given value.

### HasMp121

`func (o *CreateContainerRequest) HasMp121() bool`

HasMp121 returns a boolean if a field has been set.

### GetMp122

`func (o *CreateContainerRequest) GetMp122() string`

GetMp122 returns the Mp122 field if non-nil, zero value otherwise.

### GetMp122Ok

`func (o *CreateContainerRequest) GetMp122Ok() (*string, bool)`

GetMp122Ok returns a tuple with the Mp122 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp122

`func (o *CreateContainerRequest) SetMp122(v string)`

SetMp122 sets Mp122 field to given value.

### HasMp122

`func (o *CreateContainerRequest) HasMp122() bool`

HasMp122 returns a boolean if a field has been set.

### GetMp123

`func (o *CreateContainerRequest) GetMp123() string`

GetMp123 returns the Mp123 field if non-nil, zero value otherwise.

### GetMp123Ok

`func (o *CreateContainerRequest) GetMp123Ok() (*string, bool)`

GetMp123Ok returns a tuple with the Mp123 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp123

`func (o *CreateContainerRequest) SetMp123(v string)`

SetMp123 sets Mp123 field to given value.

### HasMp123

`func (o *CreateContainerRequest) HasMp123() bool`

HasMp123 returns a boolean if a field has been set.

### GetMp124

`func (o *CreateContainerRequest) GetMp124() string`

GetMp124 returns the Mp124 field if non-nil, zero value otherwise.

### GetMp124Ok

`func (o *CreateContainerRequest) GetMp124Ok() (*string, bool)`

GetMp124Ok returns a tuple with the Mp124 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp124

`func (o *CreateContainerRequest) SetMp124(v string)`

SetMp124 sets Mp124 field to given value.

### HasMp124

`func (o *CreateContainerRequest) HasMp124() bool`

HasMp124 returns a boolean if a field has been set.

### GetMp125

`func (o *CreateContainerRequest) GetMp125() string`

GetMp125 returns the Mp125 field if non-nil, zero value otherwise.

### GetMp125Ok

`func (o *CreateContainerRequest) GetMp125Ok() (*string, bool)`

GetMp125Ok returns a tuple with the Mp125 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp125

`func (o *CreateContainerRequest) SetMp125(v string)`

SetMp125 sets Mp125 field to given value.

### HasMp125

`func (o *CreateContainerRequest) HasMp125() bool`

HasMp125 returns a boolean if a field has been set.

### GetMp126

`func (o *CreateContainerRequest) GetMp126() string`

GetMp126 returns the Mp126 field if non-nil, zero value otherwise.

### GetMp126Ok

`func (o *CreateContainerRequest) GetMp126Ok() (*string, bool)`

GetMp126Ok returns a tuple with the Mp126 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp126

`func (o *CreateContainerRequest) SetMp126(v string)`

SetMp126 sets Mp126 field to given value.

### HasMp126

`func (o *CreateContainerRequest) HasMp126() bool`

HasMp126 returns a boolean if a field has been set.

### GetMp127

`func (o *CreateContainerRequest) GetMp127() string`

GetMp127 returns the Mp127 field if non-nil, zero value otherwise.

### GetMp127Ok

`func (o *CreateContainerRequest) GetMp127Ok() (*string, bool)`

GetMp127Ok returns a tuple with the Mp127 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp127

`func (o *CreateContainerRequest) SetMp127(v string)`

SetMp127 sets Mp127 field to given value.

### HasMp127

`func (o *CreateContainerRequest) HasMp127() bool`

HasMp127 returns a boolean if a field has been set.

### GetMp128

`func (o *CreateContainerRequest) GetMp128() string`

GetMp128 returns the Mp128 field if non-nil, zero value otherwise.

### GetMp128Ok

`func (o *CreateContainerRequest) GetMp128Ok() (*string, bool)`

GetMp128Ok returns a tuple with the Mp128 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp128

`func (o *CreateContainerRequest) SetMp128(v string)`

SetMp128 sets Mp128 field to given value.

### HasMp128

`func (o *CreateContainerRequest) HasMp128() bool`

HasMp128 returns a boolean if a field has been set.

### GetMp129

`func (o *CreateContainerRequest) GetMp129() string`

GetMp129 returns the Mp129 field if non-nil, zero value otherwise.

### GetMp129Ok

`func (o *CreateContainerRequest) GetMp129Ok() (*string, bool)`

GetMp129Ok returns a tuple with the Mp129 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp129

`func (o *CreateContainerRequest) SetMp129(v string)`

SetMp129 sets Mp129 field to given value.

### HasMp129

`func (o *CreateContainerRequest) HasMp129() bool`

HasMp129 returns a boolean if a field has been set.

### GetMp130

`func (o *CreateContainerRequest) GetMp130() string`

GetMp130 returns the Mp130 field if non-nil, zero value otherwise.

### GetMp130Ok

`func (o *CreateContainerRequest) GetMp130Ok() (*string, bool)`

GetMp130Ok returns a tuple with the Mp130 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp130

`func (o *CreateContainerRequest) SetMp130(v string)`

SetMp130 sets Mp130 field to given value.

### HasMp130

`func (o *CreateContainerRequest) HasMp130() bool`

HasMp130 returns a boolean if a field has been set.

### GetMp131

`func (o *CreateContainerRequest) GetMp131() string`

GetMp131 returns the Mp131 field if non-nil, zero value otherwise.

### GetMp131Ok

`func (o *CreateContainerRequest) GetMp131Ok() (*string, bool)`

GetMp131Ok returns a tuple with the Mp131 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp131

`func (o *CreateContainerRequest) SetMp131(v string)`

SetMp131 sets Mp131 field to given value.

### HasMp131

`func (o *CreateContainerRequest) HasMp131() bool`

HasMp131 returns a boolean if a field has been set.

### GetMp132

`func (o *CreateContainerRequest) GetMp132() string`

GetMp132 returns the Mp132 field if non-nil, zero value otherwise.

### GetMp132Ok

`func (o *CreateContainerRequest) GetMp132Ok() (*string, bool)`

GetMp132Ok returns a tuple with the Mp132 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp132

`func (o *CreateContainerRequest) SetMp132(v string)`

SetMp132 sets Mp132 field to given value.

### HasMp132

`func (o *CreateContainerRequest) HasMp132() bool`

HasMp132 returns a boolean if a field has been set.

### GetMp133

`func (o *CreateContainerRequest) GetMp133() string`

GetMp133 returns the Mp133 field if non-nil, zero value otherwise.

### GetMp133Ok

`func (o *CreateContainerRequest) GetMp133Ok() (*string, bool)`

GetMp133Ok returns a tuple with the Mp133 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp133

`func (o *CreateContainerRequest) SetMp133(v string)`

SetMp133 sets Mp133 field to given value.

### HasMp133

`func (o *CreateContainerRequest) HasMp133() bool`

HasMp133 returns a boolean if a field has been set.

### GetMp134

`func (o *CreateContainerRequest) GetMp134() string`

GetMp134 returns the Mp134 field if non-nil, zero value otherwise.

### GetMp134Ok

`func (o *CreateContainerRequest) GetMp134Ok() (*string, bool)`

GetMp134Ok returns a tuple with the Mp134 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp134

`func (o *CreateContainerRequest) SetMp134(v string)`

SetMp134 sets Mp134 field to given value.

### HasMp134

`func (o *CreateContainerRequest) HasMp134() bool`

HasMp134 returns a boolean if a field has been set.

### GetMp135

`func (o *CreateContainerRequest) GetMp135() string`

GetMp135 returns the Mp135 field if non-nil, zero value otherwise.

### GetMp135Ok

`func (o *CreateContainerRequest) GetMp135Ok() (*string, bool)`

GetMp135Ok returns a tuple with the Mp135 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp135

`func (o *CreateContainerRequest) SetMp135(v string)`

SetMp135 sets Mp135 field to given value.

### HasMp135

`func (o *CreateContainerRequest) HasMp135() bool`

HasMp135 returns a boolean if a field has been set.

### GetMp136

`func (o *CreateContainerRequest) GetMp136() string`

GetMp136 returns the Mp136 field if non-nil, zero value otherwise.

### GetMp136Ok

`func (o *CreateContainerRequest) GetMp136Ok() (*string, bool)`

GetMp136Ok returns a tuple with the Mp136 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp136

`func (o *CreateContainerRequest) SetMp136(v string)`

SetMp136 sets Mp136 field to given value.

### HasMp136

`func (o *CreateContainerRequest) HasMp136() bool`

HasMp136 returns a boolean if a field has been set.

### GetMp137

`func (o *CreateContainerRequest) GetMp137() string`

GetMp137 returns the Mp137 field if non-nil, zero value otherwise.

### GetMp137Ok

`func (o *CreateContainerRequest) GetMp137Ok() (*string, bool)`

GetMp137Ok returns a tuple with the Mp137 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp137

`func (o *CreateContainerRequest) SetMp137(v string)`

SetMp137 sets Mp137 field to given value.

### HasMp137

`func (o *CreateContainerRequest) HasMp137() bool`

HasMp137 returns a boolean if a field has been set.

### GetMp138

`func (o *CreateContainerRequest) GetMp138() string`

GetMp138 returns the Mp138 field if non-nil, zero value otherwise.

### GetMp138Ok

`func (o *CreateContainerRequest) GetMp138Ok() (*string, bool)`

GetMp138Ok returns a tuple with the Mp138 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp138

`func (o *CreateContainerRequest) SetMp138(v string)`

SetMp138 sets Mp138 field to given value.

### HasMp138

`func (o *CreateContainerRequest) HasMp138() bool`

HasMp138 returns a boolean if a field has been set.

### GetMp139

`func (o *CreateContainerRequest) GetMp139() string`

GetMp139 returns the Mp139 field if non-nil, zero value otherwise.

### GetMp139Ok

`func (o *CreateContainerRequest) GetMp139Ok() (*string, bool)`

GetMp139Ok returns a tuple with the Mp139 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp139

`func (o *CreateContainerRequest) SetMp139(v string)`

SetMp139 sets Mp139 field to given value.

### HasMp139

`func (o *CreateContainerRequest) HasMp139() bool`

HasMp139 returns a boolean if a field has been set.

### GetMp140

`func (o *CreateContainerRequest) GetMp140() string`

GetMp140 returns the Mp140 field if non-nil, zero value otherwise.

### GetMp140Ok

`func (o *CreateContainerRequest) GetMp140Ok() (*string, bool)`

GetMp140Ok returns a tuple with the Mp140 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp140

`func (o *CreateContainerRequest) SetMp140(v string)`

SetMp140 sets Mp140 field to given value.

### HasMp140

`func (o *CreateContainerRequest) HasMp140() bool`

HasMp140 returns a boolean if a field has been set.

### GetMp141

`func (o *CreateContainerRequest) GetMp141() string`

GetMp141 returns the Mp141 field if non-nil, zero value otherwise.

### GetMp141Ok

`func (o *CreateContainerRequest) GetMp141Ok() (*string, bool)`

GetMp141Ok returns a tuple with the Mp141 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp141

`func (o *CreateContainerRequest) SetMp141(v string)`

SetMp141 sets Mp141 field to given value.

### HasMp141

`func (o *CreateContainerRequest) HasMp141() bool`

HasMp141 returns a boolean if a field has been set.

### GetMp142

`func (o *CreateContainerRequest) GetMp142() string`

GetMp142 returns the Mp142 field if non-nil, zero value otherwise.

### GetMp142Ok

`func (o *CreateContainerRequest) GetMp142Ok() (*string, bool)`

GetMp142Ok returns a tuple with the Mp142 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp142

`func (o *CreateContainerRequest) SetMp142(v string)`

SetMp142 sets Mp142 field to given value.

### HasMp142

`func (o *CreateContainerRequest) HasMp142() bool`

HasMp142 returns a boolean if a field has been set.

### GetMp143

`func (o *CreateContainerRequest) GetMp143() string`

GetMp143 returns the Mp143 field if non-nil, zero value otherwise.

### GetMp143Ok

`func (o *CreateContainerRequest) GetMp143Ok() (*string, bool)`

GetMp143Ok returns a tuple with the Mp143 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp143

`func (o *CreateContainerRequest) SetMp143(v string)`

SetMp143 sets Mp143 field to given value.

### HasMp143

`func (o *CreateContainerRequest) HasMp143() bool`

HasMp143 returns a boolean if a field has been set.

### GetMp144

`func (o *CreateContainerRequest) GetMp144() string`

GetMp144 returns the Mp144 field if non-nil, zero value otherwise.

### GetMp144Ok

`func (o *CreateContainerRequest) GetMp144Ok() (*string, bool)`

GetMp144Ok returns a tuple with the Mp144 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp144

`func (o *CreateContainerRequest) SetMp144(v string)`

SetMp144 sets Mp144 field to given value.

### HasMp144

`func (o *CreateContainerRequest) HasMp144() bool`

HasMp144 returns a boolean if a field has been set.

### GetMp145

`func (o *CreateContainerRequest) GetMp145() string`

GetMp145 returns the Mp145 field if non-nil, zero value otherwise.

### GetMp145Ok

`func (o *CreateContainerRequest) GetMp145Ok() (*string, bool)`

GetMp145Ok returns a tuple with the Mp145 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp145

`func (o *CreateContainerRequest) SetMp145(v string)`

SetMp145 sets Mp145 field to given value.

### HasMp145

`func (o *CreateContainerRequest) HasMp145() bool`

HasMp145 returns a boolean if a field has been set.

### GetMp146

`func (o *CreateContainerRequest) GetMp146() string`

GetMp146 returns the Mp146 field if non-nil, zero value otherwise.

### GetMp146Ok

`func (o *CreateContainerRequest) GetMp146Ok() (*string, bool)`

GetMp146Ok returns a tuple with the Mp146 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp146

`func (o *CreateContainerRequest) SetMp146(v string)`

SetMp146 sets Mp146 field to given value.

### HasMp146

`func (o *CreateContainerRequest) HasMp146() bool`

HasMp146 returns a boolean if a field has been set.

### GetMp147

`func (o *CreateContainerRequest) GetMp147() string`

GetMp147 returns the Mp147 field if non-nil, zero value otherwise.

### GetMp147Ok

`func (o *CreateContainerRequest) GetMp147Ok() (*string, bool)`

GetMp147Ok returns a tuple with the Mp147 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp147

`func (o *CreateContainerRequest) SetMp147(v string)`

SetMp147 sets Mp147 field to given value.

### HasMp147

`func (o *CreateContainerRequest) HasMp147() bool`

HasMp147 returns a boolean if a field has been set.

### GetMp148

`func (o *CreateContainerRequest) GetMp148() string`

GetMp148 returns the Mp148 field if non-nil, zero value otherwise.

### GetMp148Ok

`func (o *CreateContainerRequest) GetMp148Ok() (*string, bool)`

GetMp148Ok returns a tuple with the Mp148 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp148

`func (o *CreateContainerRequest) SetMp148(v string)`

SetMp148 sets Mp148 field to given value.

### HasMp148

`func (o *CreateContainerRequest) HasMp148() bool`

HasMp148 returns a boolean if a field has been set.

### GetMp149

`func (o *CreateContainerRequest) GetMp149() string`

GetMp149 returns the Mp149 field if non-nil, zero value otherwise.

### GetMp149Ok

`func (o *CreateContainerRequest) GetMp149Ok() (*string, bool)`

GetMp149Ok returns a tuple with the Mp149 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp149

`func (o *CreateContainerRequest) SetMp149(v string)`

SetMp149 sets Mp149 field to given value.

### HasMp149

`func (o *CreateContainerRequest) HasMp149() bool`

HasMp149 returns a boolean if a field has been set.

### GetMp150

`func (o *CreateContainerRequest) GetMp150() string`

GetMp150 returns the Mp150 field if non-nil, zero value otherwise.

### GetMp150Ok

`func (o *CreateContainerRequest) GetMp150Ok() (*string, bool)`

GetMp150Ok returns a tuple with the Mp150 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp150

`func (o *CreateContainerRequest) SetMp150(v string)`

SetMp150 sets Mp150 field to given value.

### HasMp150

`func (o *CreateContainerRequest) HasMp150() bool`

HasMp150 returns a boolean if a field has been set.

### GetMp151

`func (o *CreateContainerRequest) GetMp151() string`

GetMp151 returns the Mp151 field if non-nil, zero value otherwise.

### GetMp151Ok

`func (o *CreateContainerRequest) GetMp151Ok() (*string, bool)`

GetMp151Ok returns a tuple with the Mp151 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp151

`func (o *CreateContainerRequest) SetMp151(v string)`

SetMp151 sets Mp151 field to given value.

### HasMp151

`func (o *CreateContainerRequest) HasMp151() bool`

HasMp151 returns a boolean if a field has been set.

### GetMp152

`func (o *CreateContainerRequest) GetMp152() string`

GetMp152 returns the Mp152 field if non-nil, zero value otherwise.

### GetMp152Ok

`func (o *CreateContainerRequest) GetMp152Ok() (*string, bool)`

GetMp152Ok returns a tuple with the Mp152 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp152

`func (o *CreateContainerRequest) SetMp152(v string)`

SetMp152 sets Mp152 field to given value.

### HasMp152

`func (o *CreateContainerRequest) HasMp152() bool`

HasMp152 returns a boolean if a field has been set.

### GetMp153

`func (o *CreateContainerRequest) GetMp153() string`

GetMp153 returns the Mp153 field if non-nil, zero value otherwise.

### GetMp153Ok

`func (o *CreateContainerRequest) GetMp153Ok() (*string, bool)`

GetMp153Ok returns a tuple with the Mp153 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp153

`func (o *CreateContainerRequest) SetMp153(v string)`

SetMp153 sets Mp153 field to given value.

### HasMp153

`func (o *CreateContainerRequest) HasMp153() bool`

HasMp153 returns a boolean if a field has been set.

### GetMp154

`func (o *CreateContainerRequest) GetMp154() string`

GetMp154 returns the Mp154 field if non-nil, zero value otherwise.

### GetMp154Ok

`func (o *CreateContainerRequest) GetMp154Ok() (*string, bool)`

GetMp154Ok returns a tuple with the Mp154 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp154

`func (o *CreateContainerRequest) SetMp154(v string)`

SetMp154 sets Mp154 field to given value.

### HasMp154

`func (o *CreateContainerRequest) HasMp154() bool`

HasMp154 returns a boolean if a field has been set.

### GetMp155

`func (o *CreateContainerRequest) GetMp155() string`

GetMp155 returns the Mp155 field if non-nil, zero value otherwise.

### GetMp155Ok

`func (o *CreateContainerRequest) GetMp155Ok() (*string, bool)`

GetMp155Ok returns a tuple with the Mp155 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp155

`func (o *CreateContainerRequest) SetMp155(v string)`

SetMp155 sets Mp155 field to given value.

### HasMp155

`func (o *CreateContainerRequest) HasMp155() bool`

HasMp155 returns a boolean if a field has been set.

### GetMp156

`func (o *CreateContainerRequest) GetMp156() string`

GetMp156 returns the Mp156 field if non-nil, zero value otherwise.

### GetMp156Ok

`func (o *CreateContainerRequest) GetMp156Ok() (*string, bool)`

GetMp156Ok returns a tuple with the Mp156 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp156

`func (o *CreateContainerRequest) SetMp156(v string)`

SetMp156 sets Mp156 field to given value.

### HasMp156

`func (o *CreateContainerRequest) HasMp156() bool`

HasMp156 returns a boolean if a field has been set.

### GetMp157

`func (o *CreateContainerRequest) GetMp157() string`

GetMp157 returns the Mp157 field if non-nil, zero value otherwise.

### GetMp157Ok

`func (o *CreateContainerRequest) GetMp157Ok() (*string, bool)`

GetMp157Ok returns a tuple with the Mp157 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp157

`func (o *CreateContainerRequest) SetMp157(v string)`

SetMp157 sets Mp157 field to given value.

### HasMp157

`func (o *CreateContainerRequest) HasMp157() bool`

HasMp157 returns a boolean if a field has been set.

### GetMp158

`func (o *CreateContainerRequest) GetMp158() string`

GetMp158 returns the Mp158 field if non-nil, zero value otherwise.

### GetMp158Ok

`func (o *CreateContainerRequest) GetMp158Ok() (*string, bool)`

GetMp158Ok returns a tuple with the Mp158 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp158

`func (o *CreateContainerRequest) SetMp158(v string)`

SetMp158 sets Mp158 field to given value.

### HasMp158

`func (o *CreateContainerRequest) HasMp158() bool`

HasMp158 returns a boolean if a field has been set.

### GetMp159

`func (o *CreateContainerRequest) GetMp159() string`

GetMp159 returns the Mp159 field if non-nil, zero value otherwise.

### GetMp159Ok

`func (o *CreateContainerRequest) GetMp159Ok() (*string, bool)`

GetMp159Ok returns a tuple with the Mp159 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp159

`func (o *CreateContainerRequest) SetMp159(v string)`

SetMp159 sets Mp159 field to given value.

### HasMp159

`func (o *CreateContainerRequest) HasMp159() bool`

HasMp159 returns a boolean if a field has been set.

### GetMp160

`func (o *CreateContainerRequest) GetMp160() string`

GetMp160 returns the Mp160 field if non-nil, zero value otherwise.

### GetMp160Ok

`func (o *CreateContainerRequest) GetMp160Ok() (*string, bool)`

GetMp160Ok returns a tuple with the Mp160 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp160

`func (o *CreateContainerRequest) SetMp160(v string)`

SetMp160 sets Mp160 field to given value.

### HasMp160

`func (o *CreateContainerRequest) HasMp160() bool`

HasMp160 returns a boolean if a field has been set.

### GetMp161

`func (o *CreateContainerRequest) GetMp161() string`

GetMp161 returns the Mp161 field if non-nil, zero value otherwise.

### GetMp161Ok

`func (o *CreateContainerRequest) GetMp161Ok() (*string, bool)`

GetMp161Ok returns a tuple with the Mp161 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp161

`func (o *CreateContainerRequest) SetMp161(v string)`

SetMp161 sets Mp161 field to given value.

### HasMp161

`func (o *CreateContainerRequest) HasMp161() bool`

HasMp161 returns a boolean if a field has been set.

### GetMp162

`func (o *CreateContainerRequest) GetMp162() string`

GetMp162 returns the Mp162 field if non-nil, zero value otherwise.

### GetMp162Ok

`func (o *CreateContainerRequest) GetMp162Ok() (*string, bool)`

GetMp162Ok returns a tuple with the Mp162 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp162

`func (o *CreateContainerRequest) SetMp162(v string)`

SetMp162 sets Mp162 field to given value.

### HasMp162

`func (o *CreateContainerRequest) HasMp162() bool`

HasMp162 returns a boolean if a field has been set.

### GetMp163

`func (o *CreateContainerRequest) GetMp163() string`

GetMp163 returns the Mp163 field if non-nil, zero value otherwise.

### GetMp163Ok

`func (o *CreateContainerRequest) GetMp163Ok() (*string, bool)`

GetMp163Ok returns a tuple with the Mp163 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp163

`func (o *CreateContainerRequest) SetMp163(v string)`

SetMp163 sets Mp163 field to given value.

### HasMp163

`func (o *CreateContainerRequest) HasMp163() bool`

HasMp163 returns a boolean if a field has been set.

### GetMp164

`func (o *CreateContainerRequest) GetMp164() string`

GetMp164 returns the Mp164 field if non-nil, zero value otherwise.

### GetMp164Ok

`func (o *CreateContainerRequest) GetMp164Ok() (*string, bool)`

GetMp164Ok returns a tuple with the Mp164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp164

`func (o *CreateContainerRequest) SetMp164(v string)`

SetMp164 sets Mp164 field to given value.

### HasMp164

`func (o *CreateContainerRequest) HasMp164() bool`

HasMp164 returns a boolean if a field has been set.

### GetMp165

`func (o *CreateContainerRequest) GetMp165() string`

GetMp165 returns the Mp165 field if non-nil, zero value otherwise.

### GetMp165Ok

`func (o *CreateContainerRequest) GetMp165Ok() (*string, bool)`

GetMp165Ok returns a tuple with the Mp165 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp165

`func (o *CreateContainerRequest) SetMp165(v string)`

SetMp165 sets Mp165 field to given value.

### HasMp165

`func (o *CreateContainerRequest) HasMp165() bool`

HasMp165 returns a boolean if a field has been set.

### GetMp166

`func (o *CreateContainerRequest) GetMp166() string`

GetMp166 returns the Mp166 field if non-nil, zero value otherwise.

### GetMp166Ok

`func (o *CreateContainerRequest) GetMp166Ok() (*string, bool)`

GetMp166Ok returns a tuple with the Mp166 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp166

`func (o *CreateContainerRequest) SetMp166(v string)`

SetMp166 sets Mp166 field to given value.

### HasMp166

`func (o *CreateContainerRequest) HasMp166() bool`

HasMp166 returns a boolean if a field has been set.

### GetMp167

`func (o *CreateContainerRequest) GetMp167() string`

GetMp167 returns the Mp167 field if non-nil, zero value otherwise.

### GetMp167Ok

`func (o *CreateContainerRequest) GetMp167Ok() (*string, bool)`

GetMp167Ok returns a tuple with the Mp167 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp167

`func (o *CreateContainerRequest) SetMp167(v string)`

SetMp167 sets Mp167 field to given value.

### HasMp167

`func (o *CreateContainerRequest) HasMp167() bool`

HasMp167 returns a boolean if a field has been set.

### GetMp168

`func (o *CreateContainerRequest) GetMp168() string`

GetMp168 returns the Mp168 field if non-nil, zero value otherwise.

### GetMp168Ok

`func (o *CreateContainerRequest) GetMp168Ok() (*string, bool)`

GetMp168Ok returns a tuple with the Mp168 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp168

`func (o *CreateContainerRequest) SetMp168(v string)`

SetMp168 sets Mp168 field to given value.

### HasMp168

`func (o *CreateContainerRequest) HasMp168() bool`

HasMp168 returns a boolean if a field has been set.

### GetMp169

`func (o *CreateContainerRequest) GetMp169() string`

GetMp169 returns the Mp169 field if non-nil, zero value otherwise.

### GetMp169Ok

`func (o *CreateContainerRequest) GetMp169Ok() (*string, bool)`

GetMp169Ok returns a tuple with the Mp169 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp169

`func (o *CreateContainerRequest) SetMp169(v string)`

SetMp169 sets Mp169 field to given value.

### HasMp169

`func (o *CreateContainerRequest) HasMp169() bool`

HasMp169 returns a boolean if a field has been set.

### GetMp170

`func (o *CreateContainerRequest) GetMp170() string`

GetMp170 returns the Mp170 field if non-nil, zero value otherwise.

### GetMp170Ok

`func (o *CreateContainerRequest) GetMp170Ok() (*string, bool)`

GetMp170Ok returns a tuple with the Mp170 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp170

`func (o *CreateContainerRequest) SetMp170(v string)`

SetMp170 sets Mp170 field to given value.

### HasMp170

`func (o *CreateContainerRequest) HasMp170() bool`

HasMp170 returns a boolean if a field has been set.

### GetMp171

`func (o *CreateContainerRequest) GetMp171() string`

GetMp171 returns the Mp171 field if non-nil, zero value otherwise.

### GetMp171Ok

`func (o *CreateContainerRequest) GetMp171Ok() (*string, bool)`

GetMp171Ok returns a tuple with the Mp171 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp171

`func (o *CreateContainerRequest) SetMp171(v string)`

SetMp171 sets Mp171 field to given value.

### HasMp171

`func (o *CreateContainerRequest) HasMp171() bool`

HasMp171 returns a boolean if a field has been set.

### GetMp172

`func (o *CreateContainerRequest) GetMp172() string`

GetMp172 returns the Mp172 field if non-nil, zero value otherwise.

### GetMp172Ok

`func (o *CreateContainerRequest) GetMp172Ok() (*string, bool)`

GetMp172Ok returns a tuple with the Mp172 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp172

`func (o *CreateContainerRequest) SetMp172(v string)`

SetMp172 sets Mp172 field to given value.

### HasMp172

`func (o *CreateContainerRequest) HasMp172() bool`

HasMp172 returns a boolean if a field has been set.

### GetMp173

`func (o *CreateContainerRequest) GetMp173() string`

GetMp173 returns the Mp173 field if non-nil, zero value otherwise.

### GetMp173Ok

`func (o *CreateContainerRequest) GetMp173Ok() (*string, bool)`

GetMp173Ok returns a tuple with the Mp173 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp173

`func (o *CreateContainerRequest) SetMp173(v string)`

SetMp173 sets Mp173 field to given value.

### HasMp173

`func (o *CreateContainerRequest) HasMp173() bool`

HasMp173 returns a boolean if a field has been set.

### GetMp174

`func (o *CreateContainerRequest) GetMp174() string`

GetMp174 returns the Mp174 field if non-nil, zero value otherwise.

### GetMp174Ok

`func (o *CreateContainerRequest) GetMp174Ok() (*string, bool)`

GetMp174Ok returns a tuple with the Mp174 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp174

`func (o *CreateContainerRequest) SetMp174(v string)`

SetMp174 sets Mp174 field to given value.

### HasMp174

`func (o *CreateContainerRequest) HasMp174() bool`

HasMp174 returns a boolean if a field has been set.

### GetMp175

`func (o *CreateContainerRequest) GetMp175() string`

GetMp175 returns the Mp175 field if non-nil, zero value otherwise.

### GetMp175Ok

`func (o *CreateContainerRequest) GetMp175Ok() (*string, bool)`

GetMp175Ok returns a tuple with the Mp175 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp175

`func (o *CreateContainerRequest) SetMp175(v string)`

SetMp175 sets Mp175 field to given value.

### HasMp175

`func (o *CreateContainerRequest) HasMp175() bool`

HasMp175 returns a boolean if a field has been set.

### GetMp176

`func (o *CreateContainerRequest) GetMp176() string`

GetMp176 returns the Mp176 field if non-nil, zero value otherwise.

### GetMp176Ok

`func (o *CreateContainerRequest) GetMp176Ok() (*string, bool)`

GetMp176Ok returns a tuple with the Mp176 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp176

`func (o *CreateContainerRequest) SetMp176(v string)`

SetMp176 sets Mp176 field to given value.

### HasMp176

`func (o *CreateContainerRequest) HasMp176() bool`

HasMp176 returns a boolean if a field has been set.

### GetMp177

`func (o *CreateContainerRequest) GetMp177() string`

GetMp177 returns the Mp177 field if non-nil, zero value otherwise.

### GetMp177Ok

`func (o *CreateContainerRequest) GetMp177Ok() (*string, bool)`

GetMp177Ok returns a tuple with the Mp177 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp177

`func (o *CreateContainerRequest) SetMp177(v string)`

SetMp177 sets Mp177 field to given value.

### HasMp177

`func (o *CreateContainerRequest) HasMp177() bool`

HasMp177 returns a boolean if a field has been set.

### GetMp178

`func (o *CreateContainerRequest) GetMp178() string`

GetMp178 returns the Mp178 field if non-nil, zero value otherwise.

### GetMp178Ok

`func (o *CreateContainerRequest) GetMp178Ok() (*string, bool)`

GetMp178Ok returns a tuple with the Mp178 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp178

`func (o *CreateContainerRequest) SetMp178(v string)`

SetMp178 sets Mp178 field to given value.

### HasMp178

`func (o *CreateContainerRequest) HasMp178() bool`

HasMp178 returns a boolean if a field has been set.

### GetMp179

`func (o *CreateContainerRequest) GetMp179() string`

GetMp179 returns the Mp179 field if non-nil, zero value otherwise.

### GetMp179Ok

`func (o *CreateContainerRequest) GetMp179Ok() (*string, bool)`

GetMp179Ok returns a tuple with the Mp179 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp179

`func (o *CreateContainerRequest) SetMp179(v string)`

SetMp179 sets Mp179 field to given value.

### HasMp179

`func (o *CreateContainerRequest) HasMp179() bool`

HasMp179 returns a boolean if a field has been set.

### GetMp180

`func (o *CreateContainerRequest) GetMp180() string`

GetMp180 returns the Mp180 field if non-nil, zero value otherwise.

### GetMp180Ok

`func (o *CreateContainerRequest) GetMp180Ok() (*string, bool)`

GetMp180Ok returns a tuple with the Mp180 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp180

`func (o *CreateContainerRequest) SetMp180(v string)`

SetMp180 sets Mp180 field to given value.

### HasMp180

`func (o *CreateContainerRequest) HasMp180() bool`

HasMp180 returns a boolean if a field has been set.

### GetMp181

`func (o *CreateContainerRequest) GetMp181() string`

GetMp181 returns the Mp181 field if non-nil, zero value otherwise.

### GetMp181Ok

`func (o *CreateContainerRequest) GetMp181Ok() (*string, bool)`

GetMp181Ok returns a tuple with the Mp181 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp181

`func (o *CreateContainerRequest) SetMp181(v string)`

SetMp181 sets Mp181 field to given value.

### HasMp181

`func (o *CreateContainerRequest) HasMp181() bool`

HasMp181 returns a boolean if a field has been set.

### GetMp182

`func (o *CreateContainerRequest) GetMp182() string`

GetMp182 returns the Mp182 field if non-nil, zero value otherwise.

### GetMp182Ok

`func (o *CreateContainerRequest) GetMp182Ok() (*string, bool)`

GetMp182Ok returns a tuple with the Mp182 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp182

`func (o *CreateContainerRequest) SetMp182(v string)`

SetMp182 sets Mp182 field to given value.

### HasMp182

`func (o *CreateContainerRequest) HasMp182() bool`

HasMp182 returns a boolean if a field has been set.

### GetMp183

`func (o *CreateContainerRequest) GetMp183() string`

GetMp183 returns the Mp183 field if non-nil, zero value otherwise.

### GetMp183Ok

`func (o *CreateContainerRequest) GetMp183Ok() (*string, bool)`

GetMp183Ok returns a tuple with the Mp183 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp183

`func (o *CreateContainerRequest) SetMp183(v string)`

SetMp183 sets Mp183 field to given value.

### HasMp183

`func (o *CreateContainerRequest) HasMp183() bool`

HasMp183 returns a boolean if a field has been set.

### GetMp184

`func (o *CreateContainerRequest) GetMp184() string`

GetMp184 returns the Mp184 field if non-nil, zero value otherwise.

### GetMp184Ok

`func (o *CreateContainerRequest) GetMp184Ok() (*string, bool)`

GetMp184Ok returns a tuple with the Mp184 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp184

`func (o *CreateContainerRequest) SetMp184(v string)`

SetMp184 sets Mp184 field to given value.

### HasMp184

`func (o *CreateContainerRequest) HasMp184() bool`

HasMp184 returns a boolean if a field has been set.

### GetMp185

`func (o *CreateContainerRequest) GetMp185() string`

GetMp185 returns the Mp185 field if non-nil, zero value otherwise.

### GetMp185Ok

`func (o *CreateContainerRequest) GetMp185Ok() (*string, bool)`

GetMp185Ok returns a tuple with the Mp185 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp185

`func (o *CreateContainerRequest) SetMp185(v string)`

SetMp185 sets Mp185 field to given value.

### HasMp185

`func (o *CreateContainerRequest) HasMp185() bool`

HasMp185 returns a boolean if a field has been set.

### GetMp186

`func (o *CreateContainerRequest) GetMp186() string`

GetMp186 returns the Mp186 field if non-nil, zero value otherwise.

### GetMp186Ok

`func (o *CreateContainerRequest) GetMp186Ok() (*string, bool)`

GetMp186Ok returns a tuple with the Mp186 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp186

`func (o *CreateContainerRequest) SetMp186(v string)`

SetMp186 sets Mp186 field to given value.

### HasMp186

`func (o *CreateContainerRequest) HasMp186() bool`

HasMp186 returns a boolean if a field has been set.

### GetMp187

`func (o *CreateContainerRequest) GetMp187() string`

GetMp187 returns the Mp187 field if non-nil, zero value otherwise.

### GetMp187Ok

`func (o *CreateContainerRequest) GetMp187Ok() (*string, bool)`

GetMp187Ok returns a tuple with the Mp187 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp187

`func (o *CreateContainerRequest) SetMp187(v string)`

SetMp187 sets Mp187 field to given value.

### HasMp187

`func (o *CreateContainerRequest) HasMp187() bool`

HasMp187 returns a boolean if a field has been set.

### GetMp188

`func (o *CreateContainerRequest) GetMp188() string`

GetMp188 returns the Mp188 field if non-nil, zero value otherwise.

### GetMp188Ok

`func (o *CreateContainerRequest) GetMp188Ok() (*string, bool)`

GetMp188Ok returns a tuple with the Mp188 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp188

`func (o *CreateContainerRequest) SetMp188(v string)`

SetMp188 sets Mp188 field to given value.

### HasMp188

`func (o *CreateContainerRequest) HasMp188() bool`

HasMp188 returns a boolean if a field has been set.

### GetMp189

`func (o *CreateContainerRequest) GetMp189() string`

GetMp189 returns the Mp189 field if non-nil, zero value otherwise.

### GetMp189Ok

`func (o *CreateContainerRequest) GetMp189Ok() (*string, bool)`

GetMp189Ok returns a tuple with the Mp189 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp189

`func (o *CreateContainerRequest) SetMp189(v string)`

SetMp189 sets Mp189 field to given value.

### HasMp189

`func (o *CreateContainerRequest) HasMp189() bool`

HasMp189 returns a boolean if a field has been set.

### GetMp190

`func (o *CreateContainerRequest) GetMp190() string`

GetMp190 returns the Mp190 field if non-nil, zero value otherwise.

### GetMp190Ok

`func (o *CreateContainerRequest) GetMp190Ok() (*string, bool)`

GetMp190Ok returns a tuple with the Mp190 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp190

`func (o *CreateContainerRequest) SetMp190(v string)`

SetMp190 sets Mp190 field to given value.

### HasMp190

`func (o *CreateContainerRequest) HasMp190() bool`

HasMp190 returns a boolean if a field has been set.

### GetMp191

`func (o *CreateContainerRequest) GetMp191() string`

GetMp191 returns the Mp191 field if non-nil, zero value otherwise.

### GetMp191Ok

`func (o *CreateContainerRequest) GetMp191Ok() (*string, bool)`

GetMp191Ok returns a tuple with the Mp191 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp191

`func (o *CreateContainerRequest) SetMp191(v string)`

SetMp191 sets Mp191 field to given value.

### HasMp191

`func (o *CreateContainerRequest) HasMp191() bool`

HasMp191 returns a boolean if a field has been set.

### GetMp192

`func (o *CreateContainerRequest) GetMp192() string`

GetMp192 returns the Mp192 field if non-nil, zero value otherwise.

### GetMp192Ok

`func (o *CreateContainerRequest) GetMp192Ok() (*string, bool)`

GetMp192Ok returns a tuple with the Mp192 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp192

`func (o *CreateContainerRequest) SetMp192(v string)`

SetMp192 sets Mp192 field to given value.

### HasMp192

`func (o *CreateContainerRequest) HasMp192() bool`

HasMp192 returns a boolean if a field has been set.

### GetMp193

`func (o *CreateContainerRequest) GetMp193() string`

GetMp193 returns the Mp193 field if non-nil, zero value otherwise.

### GetMp193Ok

`func (o *CreateContainerRequest) GetMp193Ok() (*string, bool)`

GetMp193Ok returns a tuple with the Mp193 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp193

`func (o *CreateContainerRequest) SetMp193(v string)`

SetMp193 sets Mp193 field to given value.

### HasMp193

`func (o *CreateContainerRequest) HasMp193() bool`

HasMp193 returns a boolean if a field has been set.

### GetMp194

`func (o *CreateContainerRequest) GetMp194() string`

GetMp194 returns the Mp194 field if non-nil, zero value otherwise.

### GetMp194Ok

`func (o *CreateContainerRequest) GetMp194Ok() (*string, bool)`

GetMp194Ok returns a tuple with the Mp194 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp194

`func (o *CreateContainerRequest) SetMp194(v string)`

SetMp194 sets Mp194 field to given value.

### HasMp194

`func (o *CreateContainerRequest) HasMp194() bool`

HasMp194 returns a boolean if a field has been set.

### GetMp195

`func (o *CreateContainerRequest) GetMp195() string`

GetMp195 returns the Mp195 field if non-nil, zero value otherwise.

### GetMp195Ok

`func (o *CreateContainerRequest) GetMp195Ok() (*string, bool)`

GetMp195Ok returns a tuple with the Mp195 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp195

`func (o *CreateContainerRequest) SetMp195(v string)`

SetMp195 sets Mp195 field to given value.

### HasMp195

`func (o *CreateContainerRequest) HasMp195() bool`

HasMp195 returns a boolean if a field has been set.

### GetMp196

`func (o *CreateContainerRequest) GetMp196() string`

GetMp196 returns the Mp196 field if non-nil, zero value otherwise.

### GetMp196Ok

`func (o *CreateContainerRequest) GetMp196Ok() (*string, bool)`

GetMp196Ok returns a tuple with the Mp196 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp196

`func (o *CreateContainerRequest) SetMp196(v string)`

SetMp196 sets Mp196 field to given value.

### HasMp196

`func (o *CreateContainerRequest) HasMp196() bool`

HasMp196 returns a boolean if a field has been set.

### GetMp197

`func (o *CreateContainerRequest) GetMp197() string`

GetMp197 returns the Mp197 field if non-nil, zero value otherwise.

### GetMp197Ok

`func (o *CreateContainerRequest) GetMp197Ok() (*string, bool)`

GetMp197Ok returns a tuple with the Mp197 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp197

`func (o *CreateContainerRequest) SetMp197(v string)`

SetMp197 sets Mp197 field to given value.

### HasMp197

`func (o *CreateContainerRequest) HasMp197() bool`

HasMp197 returns a boolean if a field has been set.

### GetMp198

`func (o *CreateContainerRequest) GetMp198() string`

GetMp198 returns the Mp198 field if non-nil, zero value otherwise.

### GetMp198Ok

`func (o *CreateContainerRequest) GetMp198Ok() (*string, bool)`

GetMp198Ok returns a tuple with the Mp198 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp198

`func (o *CreateContainerRequest) SetMp198(v string)`

SetMp198 sets Mp198 field to given value.

### HasMp198

`func (o *CreateContainerRequest) HasMp198() bool`

HasMp198 returns a boolean if a field has been set.

### GetMp199

`func (o *CreateContainerRequest) GetMp199() string`

GetMp199 returns the Mp199 field if non-nil, zero value otherwise.

### GetMp199Ok

`func (o *CreateContainerRequest) GetMp199Ok() (*string, bool)`

GetMp199Ok returns a tuple with the Mp199 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp199

`func (o *CreateContainerRequest) SetMp199(v string)`

SetMp199 sets Mp199 field to given value.

### HasMp199

`func (o *CreateContainerRequest) HasMp199() bool`

HasMp199 returns a boolean if a field has been set.

### GetMp200

`func (o *CreateContainerRequest) GetMp200() string`

GetMp200 returns the Mp200 field if non-nil, zero value otherwise.

### GetMp200Ok

`func (o *CreateContainerRequest) GetMp200Ok() (*string, bool)`

GetMp200Ok returns a tuple with the Mp200 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp200

`func (o *CreateContainerRequest) SetMp200(v string)`

SetMp200 sets Mp200 field to given value.

### HasMp200

`func (o *CreateContainerRequest) HasMp200() bool`

HasMp200 returns a boolean if a field has been set.

### GetMp201

`func (o *CreateContainerRequest) GetMp201() string`

GetMp201 returns the Mp201 field if non-nil, zero value otherwise.

### GetMp201Ok

`func (o *CreateContainerRequest) GetMp201Ok() (*string, bool)`

GetMp201Ok returns a tuple with the Mp201 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp201

`func (o *CreateContainerRequest) SetMp201(v string)`

SetMp201 sets Mp201 field to given value.

### HasMp201

`func (o *CreateContainerRequest) HasMp201() bool`

HasMp201 returns a boolean if a field has been set.

### GetMp202

`func (o *CreateContainerRequest) GetMp202() string`

GetMp202 returns the Mp202 field if non-nil, zero value otherwise.

### GetMp202Ok

`func (o *CreateContainerRequest) GetMp202Ok() (*string, bool)`

GetMp202Ok returns a tuple with the Mp202 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp202

`func (o *CreateContainerRequest) SetMp202(v string)`

SetMp202 sets Mp202 field to given value.

### HasMp202

`func (o *CreateContainerRequest) HasMp202() bool`

HasMp202 returns a boolean if a field has been set.

### GetMp203

`func (o *CreateContainerRequest) GetMp203() string`

GetMp203 returns the Mp203 field if non-nil, zero value otherwise.

### GetMp203Ok

`func (o *CreateContainerRequest) GetMp203Ok() (*string, bool)`

GetMp203Ok returns a tuple with the Mp203 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp203

`func (o *CreateContainerRequest) SetMp203(v string)`

SetMp203 sets Mp203 field to given value.

### HasMp203

`func (o *CreateContainerRequest) HasMp203() bool`

HasMp203 returns a boolean if a field has been set.

### GetMp204

`func (o *CreateContainerRequest) GetMp204() string`

GetMp204 returns the Mp204 field if non-nil, zero value otherwise.

### GetMp204Ok

`func (o *CreateContainerRequest) GetMp204Ok() (*string, bool)`

GetMp204Ok returns a tuple with the Mp204 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp204

`func (o *CreateContainerRequest) SetMp204(v string)`

SetMp204 sets Mp204 field to given value.

### HasMp204

`func (o *CreateContainerRequest) HasMp204() bool`

HasMp204 returns a boolean if a field has been set.

### GetMp205

`func (o *CreateContainerRequest) GetMp205() string`

GetMp205 returns the Mp205 field if non-nil, zero value otherwise.

### GetMp205Ok

`func (o *CreateContainerRequest) GetMp205Ok() (*string, bool)`

GetMp205Ok returns a tuple with the Mp205 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp205

`func (o *CreateContainerRequest) SetMp205(v string)`

SetMp205 sets Mp205 field to given value.

### HasMp205

`func (o *CreateContainerRequest) HasMp205() bool`

HasMp205 returns a boolean if a field has been set.

### GetMp206

`func (o *CreateContainerRequest) GetMp206() string`

GetMp206 returns the Mp206 field if non-nil, zero value otherwise.

### GetMp206Ok

`func (o *CreateContainerRequest) GetMp206Ok() (*string, bool)`

GetMp206Ok returns a tuple with the Mp206 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp206

`func (o *CreateContainerRequest) SetMp206(v string)`

SetMp206 sets Mp206 field to given value.

### HasMp206

`func (o *CreateContainerRequest) HasMp206() bool`

HasMp206 returns a boolean if a field has been set.

### GetMp207

`func (o *CreateContainerRequest) GetMp207() string`

GetMp207 returns the Mp207 field if non-nil, zero value otherwise.

### GetMp207Ok

`func (o *CreateContainerRequest) GetMp207Ok() (*string, bool)`

GetMp207Ok returns a tuple with the Mp207 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp207

`func (o *CreateContainerRequest) SetMp207(v string)`

SetMp207 sets Mp207 field to given value.

### HasMp207

`func (o *CreateContainerRequest) HasMp207() bool`

HasMp207 returns a boolean if a field has been set.

### GetMp208

`func (o *CreateContainerRequest) GetMp208() string`

GetMp208 returns the Mp208 field if non-nil, zero value otherwise.

### GetMp208Ok

`func (o *CreateContainerRequest) GetMp208Ok() (*string, bool)`

GetMp208Ok returns a tuple with the Mp208 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp208

`func (o *CreateContainerRequest) SetMp208(v string)`

SetMp208 sets Mp208 field to given value.

### HasMp208

`func (o *CreateContainerRequest) HasMp208() bool`

HasMp208 returns a boolean if a field has been set.

### GetMp209

`func (o *CreateContainerRequest) GetMp209() string`

GetMp209 returns the Mp209 field if non-nil, zero value otherwise.

### GetMp209Ok

`func (o *CreateContainerRequest) GetMp209Ok() (*string, bool)`

GetMp209Ok returns a tuple with the Mp209 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp209

`func (o *CreateContainerRequest) SetMp209(v string)`

SetMp209 sets Mp209 field to given value.

### HasMp209

`func (o *CreateContainerRequest) HasMp209() bool`

HasMp209 returns a boolean if a field has been set.

### GetMp210

`func (o *CreateContainerRequest) GetMp210() string`

GetMp210 returns the Mp210 field if non-nil, zero value otherwise.

### GetMp210Ok

`func (o *CreateContainerRequest) GetMp210Ok() (*string, bool)`

GetMp210Ok returns a tuple with the Mp210 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp210

`func (o *CreateContainerRequest) SetMp210(v string)`

SetMp210 sets Mp210 field to given value.

### HasMp210

`func (o *CreateContainerRequest) HasMp210() bool`

HasMp210 returns a boolean if a field has been set.

### GetMp211

`func (o *CreateContainerRequest) GetMp211() string`

GetMp211 returns the Mp211 field if non-nil, zero value otherwise.

### GetMp211Ok

`func (o *CreateContainerRequest) GetMp211Ok() (*string, bool)`

GetMp211Ok returns a tuple with the Mp211 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp211

`func (o *CreateContainerRequest) SetMp211(v string)`

SetMp211 sets Mp211 field to given value.

### HasMp211

`func (o *CreateContainerRequest) HasMp211() bool`

HasMp211 returns a boolean if a field has been set.

### GetMp212

`func (o *CreateContainerRequest) GetMp212() string`

GetMp212 returns the Mp212 field if non-nil, zero value otherwise.

### GetMp212Ok

`func (o *CreateContainerRequest) GetMp212Ok() (*string, bool)`

GetMp212Ok returns a tuple with the Mp212 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp212

`func (o *CreateContainerRequest) SetMp212(v string)`

SetMp212 sets Mp212 field to given value.

### HasMp212

`func (o *CreateContainerRequest) HasMp212() bool`

HasMp212 returns a boolean if a field has been set.

### GetMp213

`func (o *CreateContainerRequest) GetMp213() string`

GetMp213 returns the Mp213 field if non-nil, zero value otherwise.

### GetMp213Ok

`func (o *CreateContainerRequest) GetMp213Ok() (*string, bool)`

GetMp213Ok returns a tuple with the Mp213 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp213

`func (o *CreateContainerRequest) SetMp213(v string)`

SetMp213 sets Mp213 field to given value.

### HasMp213

`func (o *CreateContainerRequest) HasMp213() bool`

HasMp213 returns a boolean if a field has been set.

### GetMp214

`func (o *CreateContainerRequest) GetMp214() string`

GetMp214 returns the Mp214 field if non-nil, zero value otherwise.

### GetMp214Ok

`func (o *CreateContainerRequest) GetMp214Ok() (*string, bool)`

GetMp214Ok returns a tuple with the Mp214 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp214

`func (o *CreateContainerRequest) SetMp214(v string)`

SetMp214 sets Mp214 field to given value.

### HasMp214

`func (o *CreateContainerRequest) HasMp214() bool`

HasMp214 returns a boolean if a field has been set.

### GetMp215

`func (o *CreateContainerRequest) GetMp215() string`

GetMp215 returns the Mp215 field if non-nil, zero value otherwise.

### GetMp215Ok

`func (o *CreateContainerRequest) GetMp215Ok() (*string, bool)`

GetMp215Ok returns a tuple with the Mp215 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp215

`func (o *CreateContainerRequest) SetMp215(v string)`

SetMp215 sets Mp215 field to given value.

### HasMp215

`func (o *CreateContainerRequest) HasMp215() bool`

HasMp215 returns a boolean if a field has been set.

### GetMp216

`func (o *CreateContainerRequest) GetMp216() string`

GetMp216 returns the Mp216 field if non-nil, zero value otherwise.

### GetMp216Ok

`func (o *CreateContainerRequest) GetMp216Ok() (*string, bool)`

GetMp216Ok returns a tuple with the Mp216 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp216

`func (o *CreateContainerRequest) SetMp216(v string)`

SetMp216 sets Mp216 field to given value.

### HasMp216

`func (o *CreateContainerRequest) HasMp216() bool`

HasMp216 returns a boolean if a field has been set.

### GetMp217

`func (o *CreateContainerRequest) GetMp217() string`

GetMp217 returns the Mp217 field if non-nil, zero value otherwise.

### GetMp217Ok

`func (o *CreateContainerRequest) GetMp217Ok() (*string, bool)`

GetMp217Ok returns a tuple with the Mp217 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp217

`func (o *CreateContainerRequest) SetMp217(v string)`

SetMp217 sets Mp217 field to given value.

### HasMp217

`func (o *CreateContainerRequest) HasMp217() bool`

HasMp217 returns a boolean if a field has been set.

### GetMp218

`func (o *CreateContainerRequest) GetMp218() string`

GetMp218 returns the Mp218 field if non-nil, zero value otherwise.

### GetMp218Ok

`func (o *CreateContainerRequest) GetMp218Ok() (*string, bool)`

GetMp218Ok returns a tuple with the Mp218 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp218

`func (o *CreateContainerRequest) SetMp218(v string)`

SetMp218 sets Mp218 field to given value.

### HasMp218

`func (o *CreateContainerRequest) HasMp218() bool`

HasMp218 returns a boolean if a field has been set.

### GetMp219

`func (o *CreateContainerRequest) GetMp219() string`

GetMp219 returns the Mp219 field if non-nil, zero value otherwise.

### GetMp219Ok

`func (o *CreateContainerRequest) GetMp219Ok() (*string, bool)`

GetMp219Ok returns a tuple with the Mp219 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp219

`func (o *CreateContainerRequest) SetMp219(v string)`

SetMp219 sets Mp219 field to given value.

### HasMp219

`func (o *CreateContainerRequest) HasMp219() bool`

HasMp219 returns a boolean if a field has been set.

### GetMp220

`func (o *CreateContainerRequest) GetMp220() string`

GetMp220 returns the Mp220 field if non-nil, zero value otherwise.

### GetMp220Ok

`func (o *CreateContainerRequest) GetMp220Ok() (*string, bool)`

GetMp220Ok returns a tuple with the Mp220 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp220

`func (o *CreateContainerRequest) SetMp220(v string)`

SetMp220 sets Mp220 field to given value.

### HasMp220

`func (o *CreateContainerRequest) HasMp220() bool`

HasMp220 returns a boolean if a field has been set.

### GetMp221

`func (o *CreateContainerRequest) GetMp221() string`

GetMp221 returns the Mp221 field if non-nil, zero value otherwise.

### GetMp221Ok

`func (o *CreateContainerRequest) GetMp221Ok() (*string, bool)`

GetMp221Ok returns a tuple with the Mp221 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp221

`func (o *CreateContainerRequest) SetMp221(v string)`

SetMp221 sets Mp221 field to given value.

### HasMp221

`func (o *CreateContainerRequest) HasMp221() bool`

HasMp221 returns a boolean if a field has been set.

### GetMp222

`func (o *CreateContainerRequest) GetMp222() string`

GetMp222 returns the Mp222 field if non-nil, zero value otherwise.

### GetMp222Ok

`func (o *CreateContainerRequest) GetMp222Ok() (*string, bool)`

GetMp222Ok returns a tuple with the Mp222 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp222

`func (o *CreateContainerRequest) SetMp222(v string)`

SetMp222 sets Mp222 field to given value.

### HasMp222

`func (o *CreateContainerRequest) HasMp222() bool`

HasMp222 returns a boolean if a field has been set.

### GetMp223

`func (o *CreateContainerRequest) GetMp223() string`

GetMp223 returns the Mp223 field if non-nil, zero value otherwise.

### GetMp223Ok

`func (o *CreateContainerRequest) GetMp223Ok() (*string, bool)`

GetMp223Ok returns a tuple with the Mp223 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp223

`func (o *CreateContainerRequest) SetMp223(v string)`

SetMp223 sets Mp223 field to given value.

### HasMp223

`func (o *CreateContainerRequest) HasMp223() bool`

HasMp223 returns a boolean if a field has been set.

### GetMp224

`func (o *CreateContainerRequest) GetMp224() string`

GetMp224 returns the Mp224 field if non-nil, zero value otherwise.

### GetMp224Ok

`func (o *CreateContainerRequest) GetMp224Ok() (*string, bool)`

GetMp224Ok returns a tuple with the Mp224 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp224

`func (o *CreateContainerRequest) SetMp224(v string)`

SetMp224 sets Mp224 field to given value.

### HasMp224

`func (o *CreateContainerRequest) HasMp224() bool`

HasMp224 returns a boolean if a field has been set.

### GetMp225

`func (o *CreateContainerRequest) GetMp225() string`

GetMp225 returns the Mp225 field if non-nil, zero value otherwise.

### GetMp225Ok

`func (o *CreateContainerRequest) GetMp225Ok() (*string, bool)`

GetMp225Ok returns a tuple with the Mp225 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp225

`func (o *CreateContainerRequest) SetMp225(v string)`

SetMp225 sets Mp225 field to given value.

### HasMp225

`func (o *CreateContainerRequest) HasMp225() bool`

HasMp225 returns a boolean if a field has been set.

### GetMp226

`func (o *CreateContainerRequest) GetMp226() string`

GetMp226 returns the Mp226 field if non-nil, zero value otherwise.

### GetMp226Ok

`func (o *CreateContainerRequest) GetMp226Ok() (*string, bool)`

GetMp226Ok returns a tuple with the Mp226 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp226

`func (o *CreateContainerRequest) SetMp226(v string)`

SetMp226 sets Mp226 field to given value.

### HasMp226

`func (o *CreateContainerRequest) HasMp226() bool`

HasMp226 returns a boolean if a field has been set.

### GetMp227

`func (o *CreateContainerRequest) GetMp227() string`

GetMp227 returns the Mp227 field if non-nil, zero value otherwise.

### GetMp227Ok

`func (o *CreateContainerRequest) GetMp227Ok() (*string, bool)`

GetMp227Ok returns a tuple with the Mp227 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp227

`func (o *CreateContainerRequest) SetMp227(v string)`

SetMp227 sets Mp227 field to given value.

### HasMp227

`func (o *CreateContainerRequest) HasMp227() bool`

HasMp227 returns a boolean if a field has been set.

### GetMp228

`func (o *CreateContainerRequest) GetMp228() string`

GetMp228 returns the Mp228 field if non-nil, zero value otherwise.

### GetMp228Ok

`func (o *CreateContainerRequest) GetMp228Ok() (*string, bool)`

GetMp228Ok returns a tuple with the Mp228 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp228

`func (o *CreateContainerRequest) SetMp228(v string)`

SetMp228 sets Mp228 field to given value.

### HasMp228

`func (o *CreateContainerRequest) HasMp228() bool`

HasMp228 returns a boolean if a field has been set.

### GetMp229

`func (o *CreateContainerRequest) GetMp229() string`

GetMp229 returns the Mp229 field if non-nil, zero value otherwise.

### GetMp229Ok

`func (o *CreateContainerRequest) GetMp229Ok() (*string, bool)`

GetMp229Ok returns a tuple with the Mp229 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp229

`func (o *CreateContainerRequest) SetMp229(v string)`

SetMp229 sets Mp229 field to given value.

### HasMp229

`func (o *CreateContainerRequest) HasMp229() bool`

HasMp229 returns a boolean if a field has been set.

### GetMp230

`func (o *CreateContainerRequest) GetMp230() string`

GetMp230 returns the Mp230 field if non-nil, zero value otherwise.

### GetMp230Ok

`func (o *CreateContainerRequest) GetMp230Ok() (*string, bool)`

GetMp230Ok returns a tuple with the Mp230 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp230

`func (o *CreateContainerRequest) SetMp230(v string)`

SetMp230 sets Mp230 field to given value.

### HasMp230

`func (o *CreateContainerRequest) HasMp230() bool`

HasMp230 returns a boolean if a field has been set.

### GetMp231

`func (o *CreateContainerRequest) GetMp231() string`

GetMp231 returns the Mp231 field if non-nil, zero value otherwise.

### GetMp231Ok

`func (o *CreateContainerRequest) GetMp231Ok() (*string, bool)`

GetMp231Ok returns a tuple with the Mp231 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp231

`func (o *CreateContainerRequest) SetMp231(v string)`

SetMp231 sets Mp231 field to given value.

### HasMp231

`func (o *CreateContainerRequest) HasMp231() bool`

HasMp231 returns a boolean if a field has been set.

### GetMp232

`func (o *CreateContainerRequest) GetMp232() string`

GetMp232 returns the Mp232 field if non-nil, zero value otherwise.

### GetMp232Ok

`func (o *CreateContainerRequest) GetMp232Ok() (*string, bool)`

GetMp232Ok returns a tuple with the Mp232 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp232

`func (o *CreateContainerRequest) SetMp232(v string)`

SetMp232 sets Mp232 field to given value.

### HasMp232

`func (o *CreateContainerRequest) HasMp232() bool`

HasMp232 returns a boolean if a field has been set.

### GetMp233

`func (o *CreateContainerRequest) GetMp233() string`

GetMp233 returns the Mp233 field if non-nil, zero value otherwise.

### GetMp233Ok

`func (o *CreateContainerRequest) GetMp233Ok() (*string, bool)`

GetMp233Ok returns a tuple with the Mp233 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp233

`func (o *CreateContainerRequest) SetMp233(v string)`

SetMp233 sets Mp233 field to given value.

### HasMp233

`func (o *CreateContainerRequest) HasMp233() bool`

HasMp233 returns a boolean if a field has been set.

### GetMp234

`func (o *CreateContainerRequest) GetMp234() string`

GetMp234 returns the Mp234 field if non-nil, zero value otherwise.

### GetMp234Ok

`func (o *CreateContainerRequest) GetMp234Ok() (*string, bool)`

GetMp234Ok returns a tuple with the Mp234 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp234

`func (o *CreateContainerRequest) SetMp234(v string)`

SetMp234 sets Mp234 field to given value.

### HasMp234

`func (o *CreateContainerRequest) HasMp234() bool`

HasMp234 returns a boolean if a field has been set.

### GetMp235

`func (o *CreateContainerRequest) GetMp235() string`

GetMp235 returns the Mp235 field if non-nil, zero value otherwise.

### GetMp235Ok

`func (o *CreateContainerRequest) GetMp235Ok() (*string, bool)`

GetMp235Ok returns a tuple with the Mp235 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp235

`func (o *CreateContainerRequest) SetMp235(v string)`

SetMp235 sets Mp235 field to given value.

### HasMp235

`func (o *CreateContainerRequest) HasMp235() bool`

HasMp235 returns a boolean if a field has been set.

### GetMp236

`func (o *CreateContainerRequest) GetMp236() string`

GetMp236 returns the Mp236 field if non-nil, zero value otherwise.

### GetMp236Ok

`func (o *CreateContainerRequest) GetMp236Ok() (*string, bool)`

GetMp236Ok returns a tuple with the Mp236 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp236

`func (o *CreateContainerRequest) SetMp236(v string)`

SetMp236 sets Mp236 field to given value.

### HasMp236

`func (o *CreateContainerRequest) HasMp236() bool`

HasMp236 returns a boolean if a field has been set.

### GetMp237

`func (o *CreateContainerRequest) GetMp237() string`

GetMp237 returns the Mp237 field if non-nil, zero value otherwise.

### GetMp237Ok

`func (o *CreateContainerRequest) GetMp237Ok() (*string, bool)`

GetMp237Ok returns a tuple with the Mp237 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp237

`func (o *CreateContainerRequest) SetMp237(v string)`

SetMp237 sets Mp237 field to given value.

### HasMp237

`func (o *CreateContainerRequest) HasMp237() bool`

HasMp237 returns a boolean if a field has been set.

### GetMp238

`func (o *CreateContainerRequest) GetMp238() string`

GetMp238 returns the Mp238 field if non-nil, zero value otherwise.

### GetMp238Ok

`func (o *CreateContainerRequest) GetMp238Ok() (*string, bool)`

GetMp238Ok returns a tuple with the Mp238 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp238

`func (o *CreateContainerRequest) SetMp238(v string)`

SetMp238 sets Mp238 field to given value.

### HasMp238

`func (o *CreateContainerRequest) HasMp238() bool`

HasMp238 returns a boolean if a field has been set.

### GetMp239

`func (o *CreateContainerRequest) GetMp239() string`

GetMp239 returns the Mp239 field if non-nil, zero value otherwise.

### GetMp239Ok

`func (o *CreateContainerRequest) GetMp239Ok() (*string, bool)`

GetMp239Ok returns a tuple with the Mp239 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp239

`func (o *CreateContainerRequest) SetMp239(v string)`

SetMp239 sets Mp239 field to given value.

### HasMp239

`func (o *CreateContainerRequest) HasMp239() bool`

HasMp239 returns a boolean if a field has been set.

### GetMp240

`func (o *CreateContainerRequest) GetMp240() string`

GetMp240 returns the Mp240 field if non-nil, zero value otherwise.

### GetMp240Ok

`func (o *CreateContainerRequest) GetMp240Ok() (*string, bool)`

GetMp240Ok returns a tuple with the Mp240 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp240

`func (o *CreateContainerRequest) SetMp240(v string)`

SetMp240 sets Mp240 field to given value.

### HasMp240

`func (o *CreateContainerRequest) HasMp240() bool`

HasMp240 returns a boolean if a field has been set.

### GetMp241

`func (o *CreateContainerRequest) GetMp241() string`

GetMp241 returns the Mp241 field if non-nil, zero value otherwise.

### GetMp241Ok

`func (o *CreateContainerRequest) GetMp241Ok() (*string, bool)`

GetMp241Ok returns a tuple with the Mp241 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp241

`func (o *CreateContainerRequest) SetMp241(v string)`

SetMp241 sets Mp241 field to given value.

### HasMp241

`func (o *CreateContainerRequest) HasMp241() bool`

HasMp241 returns a boolean if a field has been set.

### GetMp242

`func (o *CreateContainerRequest) GetMp242() string`

GetMp242 returns the Mp242 field if non-nil, zero value otherwise.

### GetMp242Ok

`func (o *CreateContainerRequest) GetMp242Ok() (*string, bool)`

GetMp242Ok returns a tuple with the Mp242 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp242

`func (o *CreateContainerRequest) SetMp242(v string)`

SetMp242 sets Mp242 field to given value.

### HasMp242

`func (o *CreateContainerRequest) HasMp242() bool`

HasMp242 returns a boolean if a field has been set.

### GetMp243

`func (o *CreateContainerRequest) GetMp243() string`

GetMp243 returns the Mp243 field if non-nil, zero value otherwise.

### GetMp243Ok

`func (o *CreateContainerRequest) GetMp243Ok() (*string, bool)`

GetMp243Ok returns a tuple with the Mp243 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp243

`func (o *CreateContainerRequest) SetMp243(v string)`

SetMp243 sets Mp243 field to given value.

### HasMp243

`func (o *CreateContainerRequest) HasMp243() bool`

HasMp243 returns a boolean if a field has been set.

### GetMp244

`func (o *CreateContainerRequest) GetMp244() string`

GetMp244 returns the Mp244 field if non-nil, zero value otherwise.

### GetMp244Ok

`func (o *CreateContainerRequest) GetMp244Ok() (*string, bool)`

GetMp244Ok returns a tuple with the Mp244 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp244

`func (o *CreateContainerRequest) SetMp244(v string)`

SetMp244 sets Mp244 field to given value.

### HasMp244

`func (o *CreateContainerRequest) HasMp244() bool`

HasMp244 returns a boolean if a field has been set.

### GetMp245

`func (o *CreateContainerRequest) GetMp245() string`

GetMp245 returns the Mp245 field if non-nil, zero value otherwise.

### GetMp245Ok

`func (o *CreateContainerRequest) GetMp245Ok() (*string, bool)`

GetMp245Ok returns a tuple with the Mp245 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp245

`func (o *CreateContainerRequest) SetMp245(v string)`

SetMp245 sets Mp245 field to given value.

### HasMp245

`func (o *CreateContainerRequest) HasMp245() bool`

HasMp245 returns a boolean if a field has been set.

### GetMp246

`func (o *CreateContainerRequest) GetMp246() string`

GetMp246 returns the Mp246 field if non-nil, zero value otherwise.

### GetMp246Ok

`func (o *CreateContainerRequest) GetMp246Ok() (*string, bool)`

GetMp246Ok returns a tuple with the Mp246 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp246

`func (o *CreateContainerRequest) SetMp246(v string)`

SetMp246 sets Mp246 field to given value.

### HasMp246

`func (o *CreateContainerRequest) HasMp246() bool`

HasMp246 returns a boolean if a field has been set.

### GetMp247

`func (o *CreateContainerRequest) GetMp247() string`

GetMp247 returns the Mp247 field if non-nil, zero value otherwise.

### GetMp247Ok

`func (o *CreateContainerRequest) GetMp247Ok() (*string, bool)`

GetMp247Ok returns a tuple with the Mp247 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp247

`func (o *CreateContainerRequest) SetMp247(v string)`

SetMp247 sets Mp247 field to given value.

### HasMp247

`func (o *CreateContainerRequest) HasMp247() bool`

HasMp247 returns a boolean if a field has been set.

### GetMp248

`func (o *CreateContainerRequest) GetMp248() string`

GetMp248 returns the Mp248 field if non-nil, zero value otherwise.

### GetMp248Ok

`func (o *CreateContainerRequest) GetMp248Ok() (*string, bool)`

GetMp248Ok returns a tuple with the Mp248 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp248

`func (o *CreateContainerRequest) SetMp248(v string)`

SetMp248 sets Mp248 field to given value.

### HasMp248

`func (o *CreateContainerRequest) HasMp248() bool`

HasMp248 returns a boolean if a field has been set.

### GetMp249

`func (o *CreateContainerRequest) GetMp249() string`

GetMp249 returns the Mp249 field if non-nil, zero value otherwise.

### GetMp249Ok

`func (o *CreateContainerRequest) GetMp249Ok() (*string, bool)`

GetMp249Ok returns a tuple with the Mp249 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp249

`func (o *CreateContainerRequest) SetMp249(v string)`

SetMp249 sets Mp249 field to given value.

### HasMp249

`func (o *CreateContainerRequest) HasMp249() bool`

HasMp249 returns a boolean if a field has been set.

### GetMp250

`func (o *CreateContainerRequest) GetMp250() string`

GetMp250 returns the Mp250 field if non-nil, zero value otherwise.

### GetMp250Ok

`func (o *CreateContainerRequest) GetMp250Ok() (*string, bool)`

GetMp250Ok returns a tuple with the Mp250 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp250

`func (o *CreateContainerRequest) SetMp250(v string)`

SetMp250 sets Mp250 field to given value.

### HasMp250

`func (o *CreateContainerRequest) HasMp250() bool`

HasMp250 returns a boolean if a field has been set.

### GetMp251

`func (o *CreateContainerRequest) GetMp251() string`

GetMp251 returns the Mp251 field if non-nil, zero value otherwise.

### GetMp251Ok

`func (o *CreateContainerRequest) GetMp251Ok() (*string, bool)`

GetMp251Ok returns a tuple with the Mp251 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp251

`func (o *CreateContainerRequest) SetMp251(v string)`

SetMp251 sets Mp251 field to given value.

### HasMp251

`func (o *CreateContainerRequest) HasMp251() bool`

HasMp251 returns a boolean if a field has been set.

### GetMp252

`func (o *CreateContainerRequest) GetMp252() string`

GetMp252 returns the Mp252 field if non-nil, zero value otherwise.

### GetMp252Ok

`func (o *CreateContainerRequest) GetMp252Ok() (*string, bool)`

GetMp252Ok returns a tuple with the Mp252 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp252

`func (o *CreateContainerRequest) SetMp252(v string)`

SetMp252 sets Mp252 field to given value.

### HasMp252

`func (o *CreateContainerRequest) HasMp252() bool`

HasMp252 returns a boolean if a field has been set.

### GetMp253

`func (o *CreateContainerRequest) GetMp253() string`

GetMp253 returns the Mp253 field if non-nil, zero value otherwise.

### GetMp253Ok

`func (o *CreateContainerRequest) GetMp253Ok() (*string, bool)`

GetMp253Ok returns a tuple with the Mp253 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp253

`func (o *CreateContainerRequest) SetMp253(v string)`

SetMp253 sets Mp253 field to given value.

### HasMp253

`func (o *CreateContainerRequest) HasMp253() bool`

HasMp253 returns a boolean if a field has been set.

### GetMp254

`func (o *CreateContainerRequest) GetMp254() string`

GetMp254 returns the Mp254 field if non-nil, zero value otherwise.

### GetMp254Ok

`func (o *CreateContainerRequest) GetMp254Ok() (*string, bool)`

GetMp254Ok returns a tuple with the Mp254 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp254

`func (o *CreateContainerRequest) SetMp254(v string)`

SetMp254 sets Mp254 field to given value.

### HasMp254

`func (o *CreateContainerRequest) HasMp254() bool`

HasMp254 returns a boolean if a field has been set.

### GetMp255

`func (o *CreateContainerRequest) GetMp255() string`

GetMp255 returns the Mp255 field if non-nil, zero value otherwise.

### GetMp255Ok

`func (o *CreateContainerRequest) GetMp255Ok() (*string, bool)`

GetMp255Ok returns a tuple with the Mp255 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp255

`func (o *CreateContainerRequest) SetMp255(v string)`

SetMp255 sets Mp255 field to given value.

### HasMp255

`func (o *CreateContainerRequest) HasMp255() bool`

HasMp255 returns a boolean if a field has been set.

### GetNameserver

`func (o *CreateContainerRequest) GetNameserver() string`

GetNameserver returns the Nameserver field if non-nil, zero value otherwise.

### GetNameserverOk

`func (o *CreateContainerRequest) GetNameserverOk() (*string, bool)`

GetNameserverOk returns a tuple with the Nameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameserver

`func (o *CreateContainerRequest) SetNameserver(v string)`

SetNameserver sets Nameserver field to given value.

### HasNameserver

`func (o *CreateContainerRequest) HasNameserver() bool`

HasNameserver returns a boolean if a field has been set.

### GetNet0

`func (o *CreateContainerRequest) GetNet0() string`

GetNet0 returns the Net0 field if non-nil, zero value otherwise.

### GetNet0Ok

`func (o *CreateContainerRequest) GetNet0Ok() (*string, bool)`

GetNet0Ok returns a tuple with the Net0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet0

`func (o *CreateContainerRequest) SetNet0(v string)`

SetNet0 sets Net0 field to given value.

### HasNet0

`func (o *CreateContainerRequest) HasNet0() bool`

HasNet0 returns a boolean if a field has been set.

### GetNet1

`func (o *CreateContainerRequest) GetNet1() string`

GetNet1 returns the Net1 field if non-nil, zero value otherwise.

### GetNet1Ok

`func (o *CreateContainerRequest) GetNet1Ok() (*string, bool)`

GetNet1Ok returns a tuple with the Net1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet1

`func (o *CreateContainerRequest) SetNet1(v string)`

SetNet1 sets Net1 field to given value.

### HasNet1

`func (o *CreateContainerRequest) HasNet1() bool`

HasNet1 returns a boolean if a field has been set.

### GetNet2

`func (o *CreateContainerRequest) GetNet2() string`

GetNet2 returns the Net2 field if non-nil, zero value otherwise.

### GetNet2Ok

`func (o *CreateContainerRequest) GetNet2Ok() (*string, bool)`

GetNet2Ok returns a tuple with the Net2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet2

`func (o *CreateContainerRequest) SetNet2(v string)`

SetNet2 sets Net2 field to given value.

### HasNet2

`func (o *CreateContainerRequest) HasNet2() bool`

HasNet2 returns a boolean if a field has been set.

### GetNet3

`func (o *CreateContainerRequest) GetNet3() string`

GetNet3 returns the Net3 field if non-nil, zero value otherwise.

### GetNet3Ok

`func (o *CreateContainerRequest) GetNet3Ok() (*string, bool)`

GetNet3Ok returns a tuple with the Net3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet3

`func (o *CreateContainerRequest) SetNet3(v string)`

SetNet3 sets Net3 field to given value.

### HasNet3

`func (o *CreateContainerRequest) HasNet3() bool`

HasNet3 returns a boolean if a field has been set.

### GetNet4

`func (o *CreateContainerRequest) GetNet4() string`

GetNet4 returns the Net4 field if non-nil, zero value otherwise.

### GetNet4Ok

`func (o *CreateContainerRequest) GetNet4Ok() (*string, bool)`

GetNet4Ok returns a tuple with the Net4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet4

`func (o *CreateContainerRequest) SetNet4(v string)`

SetNet4 sets Net4 field to given value.

### HasNet4

`func (o *CreateContainerRequest) HasNet4() bool`

HasNet4 returns a boolean if a field has been set.

### GetNet5

`func (o *CreateContainerRequest) GetNet5() string`

GetNet5 returns the Net5 field if non-nil, zero value otherwise.

### GetNet5Ok

`func (o *CreateContainerRequest) GetNet5Ok() (*string, bool)`

GetNet5Ok returns a tuple with the Net5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet5

`func (o *CreateContainerRequest) SetNet5(v string)`

SetNet5 sets Net5 field to given value.

### HasNet5

`func (o *CreateContainerRequest) HasNet5() bool`

HasNet5 returns a boolean if a field has been set.

### GetNet6

`func (o *CreateContainerRequest) GetNet6() string`

GetNet6 returns the Net6 field if non-nil, zero value otherwise.

### GetNet6Ok

`func (o *CreateContainerRequest) GetNet6Ok() (*string, bool)`

GetNet6Ok returns a tuple with the Net6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet6

`func (o *CreateContainerRequest) SetNet6(v string)`

SetNet6 sets Net6 field to given value.

### HasNet6

`func (o *CreateContainerRequest) HasNet6() bool`

HasNet6 returns a boolean if a field has been set.

### GetNet7

`func (o *CreateContainerRequest) GetNet7() string`

GetNet7 returns the Net7 field if non-nil, zero value otherwise.

### GetNet7Ok

`func (o *CreateContainerRequest) GetNet7Ok() (*string, bool)`

GetNet7Ok returns a tuple with the Net7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet7

`func (o *CreateContainerRequest) SetNet7(v string)`

SetNet7 sets Net7 field to given value.

### HasNet7

`func (o *CreateContainerRequest) HasNet7() bool`

HasNet7 returns a boolean if a field has been set.

### GetNet8

`func (o *CreateContainerRequest) GetNet8() string`

GetNet8 returns the Net8 field if non-nil, zero value otherwise.

### GetNet8Ok

`func (o *CreateContainerRequest) GetNet8Ok() (*string, bool)`

GetNet8Ok returns a tuple with the Net8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet8

`func (o *CreateContainerRequest) SetNet8(v string)`

SetNet8 sets Net8 field to given value.

### HasNet8

`func (o *CreateContainerRequest) HasNet8() bool`

HasNet8 returns a boolean if a field has been set.

### GetNet9

`func (o *CreateContainerRequest) GetNet9() string`

GetNet9 returns the Net9 field if non-nil, zero value otherwise.

### GetNet9Ok

`func (o *CreateContainerRequest) GetNet9Ok() (*string, bool)`

GetNet9Ok returns a tuple with the Net9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet9

`func (o *CreateContainerRequest) SetNet9(v string)`

SetNet9 sets Net9 field to given value.

### HasNet9

`func (o *CreateContainerRequest) HasNet9() bool`

HasNet9 returns a boolean if a field has been set.

### GetNet10

`func (o *CreateContainerRequest) GetNet10() string`

GetNet10 returns the Net10 field if non-nil, zero value otherwise.

### GetNet10Ok

`func (o *CreateContainerRequest) GetNet10Ok() (*string, bool)`

GetNet10Ok returns a tuple with the Net10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet10

`func (o *CreateContainerRequest) SetNet10(v string)`

SetNet10 sets Net10 field to given value.

### HasNet10

`func (o *CreateContainerRequest) HasNet10() bool`

HasNet10 returns a boolean if a field has been set.

### GetNet11

`func (o *CreateContainerRequest) GetNet11() string`

GetNet11 returns the Net11 field if non-nil, zero value otherwise.

### GetNet11Ok

`func (o *CreateContainerRequest) GetNet11Ok() (*string, bool)`

GetNet11Ok returns a tuple with the Net11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet11

`func (o *CreateContainerRequest) SetNet11(v string)`

SetNet11 sets Net11 field to given value.

### HasNet11

`func (o *CreateContainerRequest) HasNet11() bool`

HasNet11 returns a boolean if a field has been set.

### GetNet12

`func (o *CreateContainerRequest) GetNet12() string`

GetNet12 returns the Net12 field if non-nil, zero value otherwise.

### GetNet12Ok

`func (o *CreateContainerRequest) GetNet12Ok() (*string, bool)`

GetNet12Ok returns a tuple with the Net12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet12

`func (o *CreateContainerRequest) SetNet12(v string)`

SetNet12 sets Net12 field to given value.

### HasNet12

`func (o *CreateContainerRequest) HasNet12() bool`

HasNet12 returns a boolean if a field has been set.

### GetNet13

`func (o *CreateContainerRequest) GetNet13() string`

GetNet13 returns the Net13 field if non-nil, zero value otherwise.

### GetNet13Ok

`func (o *CreateContainerRequest) GetNet13Ok() (*string, bool)`

GetNet13Ok returns a tuple with the Net13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet13

`func (o *CreateContainerRequest) SetNet13(v string)`

SetNet13 sets Net13 field to given value.

### HasNet13

`func (o *CreateContainerRequest) HasNet13() bool`

HasNet13 returns a boolean if a field has been set.

### GetNet14

`func (o *CreateContainerRequest) GetNet14() string`

GetNet14 returns the Net14 field if non-nil, zero value otherwise.

### GetNet14Ok

`func (o *CreateContainerRequest) GetNet14Ok() (*string, bool)`

GetNet14Ok returns a tuple with the Net14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet14

`func (o *CreateContainerRequest) SetNet14(v string)`

SetNet14 sets Net14 field to given value.

### HasNet14

`func (o *CreateContainerRequest) HasNet14() bool`

HasNet14 returns a boolean if a field has been set.

### GetNet15

`func (o *CreateContainerRequest) GetNet15() string`

GetNet15 returns the Net15 field if non-nil, zero value otherwise.

### GetNet15Ok

`func (o *CreateContainerRequest) GetNet15Ok() (*string, bool)`

GetNet15Ok returns a tuple with the Net15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet15

`func (o *CreateContainerRequest) SetNet15(v string)`

SetNet15 sets Net15 field to given value.

### HasNet15

`func (o *CreateContainerRequest) HasNet15() bool`

HasNet15 returns a boolean if a field has been set.

### GetNet16

`func (o *CreateContainerRequest) GetNet16() string`

GetNet16 returns the Net16 field if non-nil, zero value otherwise.

### GetNet16Ok

`func (o *CreateContainerRequest) GetNet16Ok() (*string, bool)`

GetNet16Ok returns a tuple with the Net16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet16

`func (o *CreateContainerRequest) SetNet16(v string)`

SetNet16 sets Net16 field to given value.

### HasNet16

`func (o *CreateContainerRequest) HasNet16() bool`

HasNet16 returns a boolean if a field has been set.

### GetNet17

`func (o *CreateContainerRequest) GetNet17() string`

GetNet17 returns the Net17 field if non-nil, zero value otherwise.

### GetNet17Ok

`func (o *CreateContainerRequest) GetNet17Ok() (*string, bool)`

GetNet17Ok returns a tuple with the Net17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet17

`func (o *CreateContainerRequest) SetNet17(v string)`

SetNet17 sets Net17 field to given value.

### HasNet17

`func (o *CreateContainerRequest) HasNet17() bool`

HasNet17 returns a boolean if a field has been set.

### GetNet18

`func (o *CreateContainerRequest) GetNet18() string`

GetNet18 returns the Net18 field if non-nil, zero value otherwise.

### GetNet18Ok

`func (o *CreateContainerRequest) GetNet18Ok() (*string, bool)`

GetNet18Ok returns a tuple with the Net18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet18

`func (o *CreateContainerRequest) SetNet18(v string)`

SetNet18 sets Net18 field to given value.

### HasNet18

`func (o *CreateContainerRequest) HasNet18() bool`

HasNet18 returns a boolean if a field has been set.

### GetNet19

`func (o *CreateContainerRequest) GetNet19() string`

GetNet19 returns the Net19 field if non-nil, zero value otherwise.

### GetNet19Ok

`func (o *CreateContainerRequest) GetNet19Ok() (*string, bool)`

GetNet19Ok returns a tuple with the Net19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet19

`func (o *CreateContainerRequest) SetNet19(v string)`

SetNet19 sets Net19 field to given value.

### HasNet19

`func (o *CreateContainerRequest) HasNet19() bool`

HasNet19 returns a boolean if a field has been set.

### GetNet20

`func (o *CreateContainerRequest) GetNet20() string`

GetNet20 returns the Net20 field if non-nil, zero value otherwise.

### GetNet20Ok

`func (o *CreateContainerRequest) GetNet20Ok() (*string, bool)`

GetNet20Ok returns a tuple with the Net20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet20

`func (o *CreateContainerRequest) SetNet20(v string)`

SetNet20 sets Net20 field to given value.

### HasNet20

`func (o *CreateContainerRequest) HasNet20() bool`

HasNet20 returns a boolean if a field has been set.

### GetNet21

`func (o *CreateContainerRequest) GetNet21() string`

GetNet21 returns the Net21 field if non-nil, zero value otherwise.

### GetNet21Ok

`func (o *CreateContainerRequest) GetNet21Ok() (*string, bool)`

GetNet21Ok returns a tuple with the Net21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet21

`func (o *CreateContainerRequest) SetNet21(v string)`

SetNet21 sets Net21 field to given value.

### HasNet21

`func (o *CreateContainerRequest) HasNet21() bool`

HasNet21 returns a boolean if a field has been set.

### GetNet22

`func (o *CreateContainerRequest) GetNet22() string`

GetNet22 returns the Net22 field if non-nil, zero value otherwise.

### GetNet22Ok

`func (o *CreateContainerRequest) GetNet22Ok() (*string, bool)`

GetNet22Ok returns a tuple with the Net22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet22

`func (o *CreateContainerRequest) SetNet22(v string)`

SetNet22 sets Net22 field to given value.

### HasNet22

`func (o *CreateContainerRequest) HasNet22() bool`

HasNet22 returns a boolean if a field has been set.

### GetNet23

`func (o *CreateContainerRequest) GetNet23() string`

GetNet23 returns the Net23 field if non-nil, zero value otherwise.

### GetNet23Ok

`func (o *CreateContainerRequest) GetNet23Ok() (*string, bool)`

GetNet23Ok returns a tuple with the Net23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet23

`func (o *CreateContainerRequest) SetNet23(v string)`

SetNet23 sets Net23 field to given value.

### HasNet23

`func (o *CreateContainerRequest) HasNet23() bool`

HasNet23 returns a boolean if a field has been set.

### GetNet24

`func (o *CreateContainerRequest) GetNet24() string`

GetNet24 returns the Net24 field if non-nil, zero value otherwise.

### GetNet24Ok

`func (o *CreateContainerRequest) GetNet24Ok() (*string, bool)`

GetNet24Ok returns a tuple with the Net24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet24

`func (o *CreateContainerRequest) SetNet24(v string)`

SetNet24 sets Net24 field to given value.

### HasNet24

`func (o *CreateContainerRequest) HasNet24() bool`

HasNet24 returns a boolean if a field has been set.

### GetNet25

`func (o *CreateContainerRequest) GetNet25() string`

GetNet25 returns the Net25 field if non-nil, zero value otherwise.

### GetNet25Ok

`func (o *CreateContainerRequest) GetNet25Ok() (*string, bool)`

GetNet25Ok returns a tuple with the Net25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet25

`func (o *CreateContainerRequest) SetNet25(v string)`

SetNet25 sets Net25 field to given value.

### HasNet25

`func (o *CreateContainerRequest) HasNet25() bool`

HasNet25 returns a boolean if a field has been set.

### GetNet26

`func (o *CreateContainerRequest) GetNet26() string`

GetNet26 returns the Net26 field if non-nil, zero value otherwise.

### GetNet26Ok

`func (o *CreateContainerRequest) GetNet26Ok() (*string, bool)`

GetNet26Ok returns a tuple with the Net26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet26

`func (o *CreateContainerRequest) SetNet26(v string)`

SetNet26 sets Net26 field to given value.

### HasNet26

`func (o *CreateContainerRequest) HasNet26() bool`

HasNet26 returns a boolean if a field has been set.

### GetNet27

`func (o *CreateContainerRequest) GetNet27() string`

GetNet27 returns the Net27 field if non-nil, zero value otherwise.

### GetNet27Ok

`func (o *CreateContainerRequest) GetNet27Ok() (*string, bool)`

GetNet27Ok returns a tuple with the Net27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet27

`func (o *CreateContainerRequest) SetNet27(v string)`

SetNet27 sets Net27 field to given value.

### HasNet27

`func (o *CreateContainerRequest) HasNet27() bool`

HasNet27 returns a boolean if a field has been set.

### GetNet28

`func (o *CreateContainerRequest) GetNet28() string`

GetNet28 returns the Net28 field if non-nil, zero value otherwise.

### GetNet28Ok

`func (o *CreateContainerRequest) GetNet28Ok() (*string, bool)`

GetNet28Ok returns a tuple with the Net28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet28

`func (o *CreateContainerRequest) SetNet28(v string)`

SetNet28 sets Net28 field to given value.

### HasNet28

`func (o *CreateContainerRequest) HasNet28() bool`

HasNet28 returns a boolean if a field has been set.

### GetNet29

`func (o *CreateContainerRequest) GetNet29() string`

GetNet29 returns the Net29 field if non-nil, zero value otherwise.

### GetNet29Ok

`func (o *CreateContainerRequest) GetNet29Ok() (*string, bool)`

GetNet29Ok returns a tuple with the Net29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet29

`func (o *CreateContainerRequest) SetNet29(v string)`

SetNet29 sets Net29 field to given value.

### HasNet29

`func (o *CreateContainerRequest) HasNet29() bool`

HasNet29 returns a boolean if a field has been set.

### GetNet30

`func (o *CreateContainerRequest) GetNet30() string`

GetNet30 returns the Net30 field if non-nil, zero value otherwise.

### GetNet30Ok

`func (o *CreateContainerRequest) GetNet30Ok() (*string, bool)`

GetNet30Ok returns a tuple with the Net30 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet30

`func (o *CreateContainerRequest) SetNet30(v string)`

SetNet30 sets Net30 field to given value.

### HasNet30

`func (o *CreateContainerRequest) HasNet30() bool`

HasNet30 returns a boolean if a field has been set.

### GetNet31

`func (o *CreateContainerRequest) GetNet31() string`

GetNet31 returns the Net31 field if non-nil, zero value otherwise.

### GetNet31Ok

`func (o *CreateContainerRequest) GetNet31Ok() (*string, bool)`

GetNet31Ok returns a tuple with the Net31 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNet31

`func (o *CreateContainerRequest) SetNet31(v string)`

SetNet31 sets Net31 field to given value.

### HasNet31

`func (o *CreateContainerRequest) HasNet31() bool`

HasNet31 returns a boolean if a field has been set.

### GetOnboot

`func (o *CreateContainerRequest) GetOnboot() int32`

GetOnboot returns the Onboot field if non-nil, zero value otherwise.

### GetOnbootOk

`func (o *CreateContainerRequest) GetOnbootOk() (*int32, bool)`

GetOnbootOk returns a tuple with the Onboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnboot

`func (o *CreateContainerRequest) SetOnboot(v int32)`

SetOnboot sets Onboot field to given value.

### HasOnboot

`func (o *CreateContainerRequest) HasOnboot() bool`

HasOnboot returns a boolean if a field has been set.

### GetOstemplate

`func (o *CreateContainerRequest) GetOstemplate() string`

GetOstemplate returns the Ostemplate field if non-nil, zero value otherwise.

### GetOstemplateOk

`func (o *CreateContainerRequest) GetOstemplateOk() (*string, bool)`

GetOstemplateOk returns a tuple with the Ostemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstemplate

`func (o *CreateContainerRequest) SetOstemplate(v string)`

SetOstemplate sets Ostemplate field to given value.


### GetOstype

`func (o *CreateContainerRequest) GetOstype() string`

GetOstype returns the Ostype field if non-nil, zero value otherwise.

### GetOstypeOk

`func (o *CreateContainerRequest) GetOstypeOk() (*string, bool)`

GetOstypeOk returns a tuple with the Ostype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOstype

`func (o *CreateContainerRequest) SetOstype(v string)`

SetOstype sets Ostype field to given value.

### HasOstype

`func (o *CreateContainerRequest) HasOstype() bool`

HasOstype returns a boolean if a field has been set.

### GetPassword

`func (o *CreateContainerRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateContainerRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateContainerRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateContainerRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPool

`func (o *CreateContainerRequest) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *CreateContainerRequest) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *CreateContainerRequest) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *CreateContainerRequest) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetProtection

`func (o *CreateContainerRequest) GetProtection() int32`

GetProtection returns the Protection field if non-nil, zero value otherwise.

### GetProtectionOk

`func (o *CreateContainerRequest) GetProtectionOk() (*int32, bool)`

GetProtectionOk returns a tuple with the Protection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtection

`func (o *CreateContainerRequest) SetProtection(v int32)`

SetProtection sets Protection field to given value.

### HasProtection

`func (o *CreateContainerRequest) HasProtection() bool`

HasProtection returns a boolean if a field has been set.

### GetRestore

`func (o *CreateContainerRequest) GetRestore() int32`

GetRestore returns the Restore field if non-nil, zero value otherwise.

### GetRestoreOk

`func (o *CreateContainerRequest) GetRestoreOk() (*int32, bool)`

GetRestoreOk returns a tuple with the Restore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestore

`func (o *CreateContainerRequest) SetRestore(v int32)`

SetRestore sets Restore field to given value.

### HasRestore

`func (o *CreateContainerRequest) HasRestore() bool`

HasRestore returns a boolean if a field has been set.

### GetRootfs

`func (o *CreateContainerRequest) GetRootfs() string`

GetRootfs returns the Rootfs field if non-nil, zero value otherwise.

### GetRootfsOk

`func (o *CreateContainerRequest) GetRootfsOk() (*string, bool)`

GetRootfsOk returns a tuple with the Rootfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootfs

`func (o *CreateContainerRequest) SetRootfs(v string)`

SetRootfs sets Rootfs field to given value.

### HasRootfs

`func (o *CreateContainerRequest) HasRootfs() bool`

HasRootfs returns a boolean if a field has been set.

### GetSearchdomain

`func (o *CreateContainerRequest) GetSearchdomain() string`

GetSearchdomain returns the Searchdomain field if non-nil, zero value otherwise.

### GetSearchdomainOk

`func (o *CreateContainerRequest) GetSearchdomainOk() (*string, bool)`

GetSearchdomainOk returns a tuple with the Searchdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchdomain

`func (o *CreateContainerRequest) SetSearchdomain(v string)`

SetSearchdomain sets Searchdomain field to given value.

### HasSearchdomain

`func (o *CreateContainerRequest) HasSearchdomain() bool`

HasSearchdomain returns a boolean if a field has been set.

### GetSshPublicKeys

`func (o *CreateContainerRequest) GetSshPublicKeys() string`

GetSshPublicKeys returns the SshPublicKeys field if non-nil, zero value otherwise.

### GetSshPublicKeysOk

`func (o *CreateContainerRequest) GetSshPublicKeysOk() (*string, bool)`

GetSshPublicKeysOk returns a tuple with the SshPublicKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPublicKeys

`func (o *CreateContainerRequest) SetSshPublicKeys(v string)`

SetSshPublicKeys sets SshPublicKeys field to given value.

### HasSshPublicKeys

`func (o *CreateContainerRequest) HasSshPublicKeys() bool`

HasSshPublicKeys returns a boolean if a field has been set.

### GetStart

`func (o *CreateContainerRequest) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *CreateContainerRequest) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *CreateContainerRequest) SetStart(v int32)`

SetStart sets Start field to given value.

### HasStart

`func (o *CreateContainerRequest) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetStartup

`func (o *CreateContainerRequest) GetStartup() string`

GetStartup returns the Startup field if non-nil, zero value otherwise.

### GetStartupOk

`func (o *CreateContainerRequest) GetStartupOk() (*string, bool)`

GetStartupOk returns a tuple with the Startup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartup

`func (o *CreateContainerRequest) SetStartup(v string)`

SetStartup sets Startup field to given value.

### HasStartup

`func (o *CreateContainerRequest) HasStartup() bool`

HasStartup returns a boolean if a field has been set.

### GetStorage

`func (o *CreateContainerRequest) GetStorage() string`

GetStorage returns the Storage field if non-nil, zero value otherwise.

### GetStorageOk

`func (o *CreateContainerRequest) GetStorageOk() (*string, bool)`

GetStorageOk returns a tuple with the Storage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorage

`func (o *CreateContainerRequest) SetStorage(v string)`

SetStorage sets Storage field to given value.

### HasStorage

`func (o *CreateContainerRequest) HasStorage() bool`

HasStorage returns a boolean if a field has been set.

### GetSwap

`func (o *CreateContainerRequest) GetSwap() int64`

GetSwap returns the Swap field if non-nil, zero value otherwise.

### GetSwapOk

`func (o *CreateContainerRequest) GetSwapOk() (*int64, bool)`

GetSwapOk returns a tuple with the Swap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwap

`func (o *CreateContainerRequest) SetSwap(v int64)`

SetSwap sets Swap field to given value.

### HasSwap

`func (o *CreateContainerRequest) HasSwap() bool`

HasSwap returns a boolean if a field has been set.

### GetTags

`func (o *CreateContainerRequest) GetTags() string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateContainerRequest) GetTagsOk() (*string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateContainerRequest) SetTags(v string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateContainerRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *CreateContainerRequest) GetTemplate() int32`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CreateContainerRequest) GetTemplateOk() (*int32, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CreateContainerRequest) SetTemplate(v int32)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CreateContainerRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetTimezone

`func (o *CreateContainerRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CreateContainerRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CreateContainerRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CreateContainerRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTty

`func (o *CreateContainerRequest) GetTty() int64`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *CreateContainerRequest) GetTtyOk() (*int64, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *CreateContainerRequest) SetTty(v int64)`

SetTty sets Tty field to given value.

### HasTty

`func (o *CreateContainerRequest) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetUnique

`func (o *CreateContainerRequest) GetUnique() int32`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *CreateContainerRequest) GetUniqueOk() (*int32, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *CreateContainerRequest) SetUnique(v int32)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *CreateContainerRequest) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetUnprivileged

`func (o *CreateContainerRequest) GetUnprivileged() int32`

GetUnprivileged returns the Unprivileged field if non-nil, zero value otherwise.

### GetUnprivilegedOk

`func (o *CreateContainerRequest) GetUnprivilegedOk() (*int32, bool)`

GetUnprivilegedOk returns a tuple with the Unprivileged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnprivileged

`func (o *CreateContainerRequest) SetUnprivileged(v int32)`

SetUnprivileged sets Unprivileged field to given value.

### HasUnprivileged

`func (o *CreateContainerRequest) HasUnprivileged() bool`

HasUnprivileged returns a boolean if a field has been set.

### GetUnused0

`func (o *CreateContainerRequest) GetUnused0() string`

GetUnused0 returns the Unused0 field if non-nil, zero value otherwise.

### GetUnused0Ok

`func (o *CreateContainerRequest) GetUnused0Ok() (*string, bool)`

GetUnused0Ok returns a tuple with the Unused0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused0

`func (o *CreateContainerRequest) SetUnused0(v string)`

SetUnused0 sets Unused0 field to given value.

### HasUnused0

`func (o *CreateContainerRequest) HasUnused0() bool`

HasUnused0 returns a boolean if a field has been set.

### GetUnused1

`func (o *CreateContainerRequest) GetUnused1() string`

GetUnused1 returns the Unused1 field if non-nil, zero value otherwise.

### GetUnused1Ok

`func (o *CreateContainerRequest) GetUnused1Ok() (*string, bool)`

GetUnused1Ok returns a tuple with the Unused1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused1

`func (o *CreateContainerRequest) SetUnused1(v string)`

SetUnused1 sets Unused1 field to given value.

### HasUnused1

`func (o *CreateContainerRequest) HasUnused1() bool`

HasUnused1 returns a boolean if a field has been set.

### GetUnused2

`func (o *CreateContainerRequest) GetUnused2() string`

GetUnused2 returns the Unused2 field if non-nil, zero value otherwise.

### GetUnused2Ok

`func (o *CreateContainerRequest) GetUnused2Ok() (*string, bool)`

GetUnused2Ok returns a tuple with the Unused2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused2

`func (o *CreateContainerRequest) SetUnused2(v string)`

SetUnused2 sets Unused2 field to given value.

### HasUnused2

`func (o *CreateContainerRequest) HasUnused2() bool`

HasUnused2 returns a boolean if a field has been set.

### GetUnused3

`func (o *CreateContainerRequest) GetUnused3() string`

GetUnused3 returns the Unused3 field if non-nil, zero value otherwise.

### GetUnused3Ok

`func (o *CreateContainerRequest) GetUnused3Ok() (*string, bool)`

GetUnused3Ok returns a tuple with the Unused3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused3

`func (o *CreateContainerRequest) SetUnused3(v string)`

SetUnused3 sets Unused3 field to given value.

### HasUnused3

`func (o *CreateContainerRequest) HasUnused3() bool`

HasUnused3 returns a boolean if a field has been set.

### GetUnused4

`func (o *CreateContainerRequest) GetUnused4() string`

GetUnused4 returns the Unused4 field if non-nil, zero value otherwise.

### GetUnused4Ok

`func (o *CreateContainerRequest) GetUnused4Ok() (*string, bool)`

GetUnused4Ok returns a tuple with the Unused4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused4

`func (o *CreateContainerRequest) SetUnused4(v string)`

SetUnused4 sets Unused4 field to given value.

### HasUnused4

`func (o *CreateContainerRequest) HasUnused4() bool`

HasUnused4 returns a boolean if a field has been set.

### GetUnused5

`func (o *CreateContainerRequest) GetUnused5() string`

GetUnused5 returns the Unused5 field if non-nil, zero value otherwise.

### GetUnused5Ok

`func (o *CreateContainerRequest) GetUnused5Ok() (*string, bool)`

GetUnused5Ok returns a tuple with the Unused5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused5

`func (o *CreateContainerRequest) SetUnused5(v string)`

SetUnused5 sets Unused5 field to given value.

### HasUnused5

`func (o *CreateContainerRequest) HasUnused5() bool`

HasUnused5 returns a boolean if a field has been set.

### GetUnused6

`func (o *CreateContainerRequest) GetUnused6() string`

GetUnused6 returns the Unused6 field if non-nil, zero value otherwise.

### GetUnused6Ok

`func (o *CreateContainerRequest) GetUnused6Ok() (*string, bool)`

GetUnused6Ok returns a tuple with the Unused6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused6

`func (o *CreateContainerRequest) SetUnused6(v string)`

SetUnused6 sets Unused6 field to given value.

### HasUnused6

`func (o *CreateContainerRequest) HasUnused6() bool`

HasUnused6 returns a boolean if a field has been set.

### GetUnused7

`func (o *CreateContainerRequest) GetUnused7() string`

GetUnused7 returns the Unused7 field if non-nil, zero value otherwise.

### GetUnused7Ok

`func (o *CreateContainerRequest) GetUnused7Ok() (*string, bool)`

GetUnused7Ok returns a tuple with the Unused7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused7

`func (o *CreateContainerRequest) SetUnused7(v string)`

SetUnused7 sets Unused7 field to given value.

### HasUnused7

`func (o *CreateContainerRequest) HasUnused7() bool`

HasUnused7 returns a boolean if a field has been set.

### GetUnused8

`func (o *CreateContainerRequest) GetUnused8() string`

GetUnused8 returns the Unused8 field if non-nil, zero value otherwise.

### GetUnused8Ok

`func (o *CreateContainerRequest) GetUnused8Ok() (*string, bool)`

GetUnused8Ok returns a tuple with the Unused8 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused8

`func (o *CreateContainerRequest) SetUnused8(v string)`

SetUnused8 sets Unused8 field to given value.

### HasUnused8

`func (o *CreateContainerRequest) HasUnused8() bool`

HasUnused8 returns a boolean if a field has been set.

### GetUnused9

`func (o *CreateContainerRequest) GetUnused9() string`

GetUnused9 returns the Unused9 field if non-nil, zero value otherwise.

### GetUnused9Ok

`func (o *CreateContainerRequest) GetUnused9Ok() (*string, bool)`

GetUnused9Ok returns a tuple with the Unused9 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused9

`func (o *CreateContainerRequest) SetUnused9(v string)`

SetUnused9 sets Unused9 field to given value.

### HasUnused9

`func (o *CreateContainerRequest) HasUnused9() bool`

HasUnused9 returns a boolean if a field has been set.

### GetUnused10

`func (o *CreateContainerRequest) GetUnused10() string`

GetUnused10 returns the Unused10 field if non-nil, zero value otherwise.

### GetUnused10Ok

`func (o *CreateContainerRequest) GetUnused10Ok() (*string, bool)`

GetUnused10Ok returns a tuple with the Unused10 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused10

`func (o *CreateContainerRequest) SetUnused10(v string)`

SetUnused10 sets Unused10 field to given value.

### HasUnused10

`func (o *CreateContainerRequest) HasUnused10() bool`

HasUnused10 returns a boolean if a field has been set.

### GetUnused11

`func (o *CreateContainerRequest) GetUnused11() string`

GetUnused11 returns the Unused11 field if non-nil, zero value otherwise.

### GetUnused11Ok

`func (o *CreateContainerRequest) GetUnused11Ok() (*string, bool)`

GetUnused11Ok returns a tuple with the Unused11 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused11

`func (o *CreateContainerRequest) SetUnused11(v string)`

SetUnused11 sets Unused11 field to given value.

### HasUnused11

`func (o *CreateContainerRequest) HasUnused11() bool`

HasUnused11 returns a boolean if a field has been set.

### GetUnused12

`func (o *CreateContainerRequest) GetUnused12() string`

GetUnused12 returns the Unused12 field if non-nil, zero value otherwise.

### GetUnused12Ok

`func (o *CreateContainerRequest) GetUnused12Ok() (*string, bool)`

GetUnused12Ok returns a tuple with the Unused12 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused12

`func (o *CreateContainerRequest) SetUnused12(v string)`

SetUnused12 sets Unused12 field to given value.

### HasUnused12

`func (o *CreateContainerRequest) HasUnused12() bool`

HasUnused12 returns a boolean if a field has been set.

### GetUnused13

`func (o *CreateContainerRequest) GetUnused13() string`

GetUnused13 returns the Unused13 field if non-nil, zero value otherwise.

### GetUnused13Ok

`func (o *CreateContainerRequest) GetUnused13Ok() (*string, bool)`

GetUnused13Ok returns a tuple with the Unused13 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused13

`func (o *CreateContainerRequest) SetUnused13(v string)`

SetUnused13 sets Unused13 field to given value.

### HasUnused13

`func (o *CreateContainerRequest) HasUnused13() bool`

HasUnused13 returns a boolean if a field has been set.

### GetUnused14

`func (o *CreateContainerRequest) GetUnused14() string`

GetUnused14 returns the Unused14 field if non-nil, zero value otherwise.

### GetUnused14Ok

`func (o *CreateContainerRequest) GetUnused14Ok() (*string, bool)`

GetUnused14Ok returns a tuple with the Unused14 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused14

`func (o *CreateContainerRequest) SetUnused14(v string)`

SetUnused14 sets Unused14 field to given value.

### HasUnused14

`func (o *CreateContainerRequest) HasUnused14() bool`

HasUnused14 returns a boolean if a field has been set.

### GetUnused15

`func (o *CreateContainerRequest) GetUnused15() string`

GetUnused15 returns the Unused15 field if non-nil, zero value otherwise.

### GetUnused15Ok

`func (o *CreateContainerRequest) GetUnused15Ok() (*string, bool)`

GetUnused15Ok returns a tuple with the Unused15 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused15

`func (o *CreateContainerRequest) SetUnused15(v string)`

SetUnused15 sets Unused15 field to given value.

### HasUnused15

`func (o *CreateContainerRequest) HasUnused15() bool`

HasUnused15 returns a boolean if a field has been set.

### GetUnused16

`func (o *CreateContainerRequest) GetUnused16() string`

GetUnused16 returns the Unused16 field if non-nil, zero value otherwise.

### GetUnused16Ok

`func (o *CreateContainerRequest) GetUnused16Ok() (*string, bool)`

GetUnused16Ok returns a tuple with the Unused16 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused16

`func (o *CreateContainerRequest) SetUnused16(v string)`

SetUnused16 sets Unused16 field to given value.

### HasUnused16

`func (o *CreateContainerRequest) HasUnused16() bool`

HasUnused16 returns a boolean if a field has been set.

### GetUnused17

`func (o *CreateContainerRequest) GetUnused17() string`

GetUnused17 returns the Unused17 field if non-nil, zero value otherwise.

### GetUnused17Ok

`func (o *CreateContainerRequest) GetUnused17Ok() (*string, bool)`

GetUnused17Ok returns a tuple with the Unused17 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused17

`func (o *CreateContainerRequest) SetUnused17(v string)`

SetUnused17 sets Unused17 field to given value.

### HasUnused17

`func (o *CreateContainerRequest) HasUnused17() bool`

HasUnused17 returns a boolean if a field has been set.

### GetUnused18

`func (o *CreateContainerRequest) GetUnused18() string`

GetUnused18 returns the Unused18 field if non-nil, zero value otherwise.

### GetUnused18Ok

`func (o *CreateContainerRequest) GetUnused18Ok() (*string, bool)`

GetUnused18Ok returns a tuple with the Unused18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused18

`func (o *CreateContainerRequest) SetUnused18(v string)`

SetUnused18 sets Unused18 field to given value.

### HasUnused18

`func (o *CreateContainerRequest) HasUnused18() bool`

HasUnused18 returns a boolean if a field has been set.

### GetUnused19

`func (o *CreateContainerRequest) GetUnused19() string`

GetUnused19 returns the Unused19 field if non-nil, zero value otherwise.

### GetUnused19Ok

`func (o *CreateContainerRequest) GetUnused19Ok() (*string, bool)`

GetUnused19Ok returns a tuple with the Unused19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused19

`func (o *CreateContainerRequest) SetUnused19(v string)`

SetUnused19 sets Unused19 field to given value.

### HasUnused19

`func (o *CreateContainerRequest) HasUnused19() bool`

HasUnused19 returns a boolean if a field has been set.

### GetUnused20

`func (o *CreateContainerRequest) GetUnused20() string`

GetUnused20 returns the Unused20 field if non-nil, zero value otherwise.

### GetUnused20Ok

`func (o *CreateContainerRequest) GetUnused20Ok() (*string, bool)`

GetUnused20Ok returns a tuple with the Unused20 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused20

`func (o *CreateContainerRequest) SetUnused20(v string)`

SetUnused20 sets Unused20 field to given value.

### HasUnused20

`func (o *CreateContainerRequest) HasUnused20() bool`

HasUnused20 returns a boolean if a field has been set.

### GetUnused21

`func (o *CreateContainerRequest) GetUnused21() string`

GetUnused21 returns the Unused21 field if non-nil, zero value otherwise.

### GetUnused21Ok

`func (o *CreateContainerRequest) GetUnused21Ok() (*string, bool)`

GetUnused21Ok returns a tuple with the Unused21 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused21

`func (o *CreateContainerRequest) SetUnused21(v string)`

SetUnused21 sets Unused21 field to given value.

### HasUnused21

`func (o *CreateContainerRequest) HasUnused21() bool`

HasUnused21 returns a boolean if a field has been set.

### GetUnused22

`func (o *CreateContainerRequest) GetUnused22() string`

GetUnused22 returns the Unused22 field if non-nil, zero value otherwise.

### GetUnused22Ok

`func (o *CreateContainerRequest) GetUnused22Ok() (*string, bool)`

GetUnused22Ok returns a tuple with the Unused22 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused22

`func (o *CreateContainerRequest) SetUnused22(v string)`

SetUnused22 sets Unused22 field to given value.

### HasUnused22

`func (o *CreateContainerRequest) HasUnused22() bool`

HasUnused22 returns a boolean if a field has been set.

### GetUnused23

`func (o *CreateContainerRequest) GetUnused23() string`

GetUnused23 returns the Unused23 field if non-nil, zero value otherwise.

### GetUnused23Ok

`func (o *CreateContainerRequest) GetUnused23Ok() (*string, bool)`

GetUnused23Ok returns a tuple with the Unused23 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused23

`func (o *CreateContainerRequest) SetUnused23(v string)`

SetUnused23 sets Unused23 field to given value.

### HasUnused23

`func (o *CreateContainerRequest) HasUnused23() bool`

HasUnused23 returns a boolean if a field has been set.

### GetUnused24

`func (o *CreateContainerRequest) GetUnused24() string`

GetUnused24 returns the Unused24 field if non-nil, zero value otherwise.

### GetUnused24Ok

`func (o *CreateContainerRequest) GetUnused24Ok() (*string, bool)`

GetUnused24Ok returns a tuple with the Unused24 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused24

`func (o *CreateContainerRequest) SetUnused24(v string)`

SetUnused24 sets Unused24 field to given value.

### HasUnused24

`func (o *CreateContainerRequest) HasUnused24() bool`

HasUnused24 returns a boolean if a field has been set.

### GetUnused25

`func (o *CreateContainerRequest) GetUnused25() string`

GetUnused25 returns the Unused25 field if non-nil, zero value otherwise.

### GetUnused25Ok

`func (o *CreateContainerRequest) GetUnused25Ok() (*string, bool)`

GetUnused25Ok returns a tuple with the Unused25 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused25

`func (o *CreateContainerRequest) SetUnused25(v string)`

SetUnused25 sets Unused25 field to given value.

### HasUnused25

`func (o *CreateContainerRequest) HasUnused25() bool`

HasUnused25 returns a boolean if a field has been set.

### GetUnused26

`func (o *CreateContainerRequest) GetUnused26() string`

GetUnused26 returns the Unused26 field if non-nil, zero value otherwise.

### GetUnused26Ok

`func (o *CreateContainerRequest) GetUnused26Ok() (*string, bool)`

GetUnused26Ok returns a tuple with the Unused26 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused26

`func (o *CreateContainerRequest) SetUnused26(v string)`

SetUnused26 sets Unused26 field to given value.

### HasUnused26

`func (o *CreateContainerRequest) HasUnused26() bool`

HasUnused26 returns a boolean if a field has been set.

### GetUnused27

`func (o *CreateContainerRequest) GetUnused27() string`

GetUnused27 returns the Unused27 field if non-nil, zero value otherwise.

### GetUnused27Ok

`func (o *CreateContainerRequest) GetUnused27Ok() (*string, bool)`

GetUnused27Ok returns a tuple with the Unused27 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused27

`func (o *CreateContainerRequest) SetUnused27(v string)`

SetUnused27 sets Unused27 field to given value.

### HasUnused27

`func (o *CreateContainerRequest) HasUnused27() bool`

HasUnused27 returns a boolean if a field has been set.

### GetUnused28

`func (o *CreateContainerRequest) GetUnused28() string`

GetUnused28 returns the Unused28 field if non-nil, zero value otherwise.

### GetUnused28Ok

`func (o *CreateContainerRequest) GetUnused28Ok() (*string, bool)`

GetUnused28Ok returns a tuple with the Unused28 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused28

`func (o *CreateContainerRequest) SetUnused28(v string)`

SetUnused28 sets Unused28 field to given value.

### HasUnused28

`func (o *CreateContainerRequest) HasUnused28() bool`

HasUnused28 returns a boolean if a field has been set.

### GetUnused29

`func (o *CreateContainerRequest) GetUnused29() string`

GetUnused29 returns the Unused29 field if non-nil, zero value otherwise.

### GetUnused29Ok

`func (o *CreateContainerRequest) GetUnused29Ok() (*string, bool)`

GetUnused29Ok returns a tuple with the Unused29 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnused29

`func (o *CreateContainerRequest) SetUnused29(v string)`

SetUnused29 sets Unused29 field to given value.

### HasUnused29

`func (o *CreateContainerRequest) HasUnused29() bool`

HasUnused29 returns a boolean if a field has been set.

### GetVmid

`func (o *CreateContainerRequest) GetVmid() int64`

GetVmid returns the Vmid field if non-nil, zero value otherwise.

### GetVmidOk

`func (o *CreateContainerRequest) GetVmidOk() (*int64, bool)`

GetVmidOk returns a tuple with the Vmid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmid

`func (o *CreateContainerRequest) SetVmid(v int64)`

SetVmid sets Vmid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


