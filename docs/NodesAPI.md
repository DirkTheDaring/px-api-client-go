# \NodesAPI

All URIs are relative to *https://127.0.0.1:8006/api2/json*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContainer**](NodesAPI.md#CreateContainer) | **Post** /nodes/{node}/lxc | createContainer
[**CreateContainerSnapshot**](NodesAPI.md#CreateContainerSnapshot) | **Post** /nodes/{node}/lxc/{vmid}/snapshot | createContainerSnapshot
[**CreateNodesSingleStorageSingleContent**](NodesAPI.md#CreateNodesSingleStorageSingleContent) | **Post** /nodes/{node}/storage/{storage}/content | createNodesSingleStorageSingleContent
[**CreateVM**](NodesAPI.md#CreateVM) | **Post** /nodes/{node}/qemu | createVM
[**CreateVMSnapshot**](NodesAPI.md#CreateVMSnapshot) | **Post** /nodes/{node}/qemu/{vmid}/snapshot | createVMSnapshot
[**DeleteContainer**](NodesAPI.md#DeleteContainer) | **Delete** /nodes/{node}/lxc/{vmid} | deleteContainer
[**DeleteContainerSnapshot**](NodesAPI.md#DeleteContainerSnapshot) | **Delete** /nodes/{node}/lxc/{vmid}/snapshot/{snapname} | deleteContainerSnapshot
[**DeleteVM**](NodesAPI.md#DeleteVM) | **Delete** /nodes/{node}/qemu/{vmid} | deleteVM
[**DeleteVMSnapshot**](NodesAPI.md#DeleteVMSnapshot) | **Delete** /nodes/{node}/qemu/{vmid}/snapshot/{snapname} | deleteVMSnapshot
[**GetContainer**](NodesAPI.md#GetContainer) | **Get** /nodes/{node}/lxc/{vmid} | getContainer
[**GetContainerConfig**](NodesAPI.md#GetContainerConfig) | **Get** /nodes/{node}/lxc/{vmid}/config | getContainerConfig
[**GetContainerConfigPending**](NodesAPI.md#GetContainerConfigPending) | **Get** /nodes/{node}/lxc/{vmid}/pending | getContainerConfigPending
[**GetContainerSnapshot**](NodesAPI.md#GetContainerSnapshot) | **Get** /nodes/{node}/lxc/{vmid}/snapshot/{snapname} | getContainerSnapshot
[**GetContainerSnapshotConfig**](NodesAPI.md#GetContainerSnapshotConfig) | **Get** /nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config | getContainerSnapshotConfig
[**GetContainerSnapshots**](NodesAPI.md#GetContainerSnapshots) | **Get** /nodes/{node}/lxc/{vmid}/snapshot | getContainerSnapshots
[**GetContainerStatus**](NodesAPI.md#GetContainerStatus) | **Get** /nodes/{node}/lxc/{vmid}/status | getContainerStatus
[**GetContainers**](NodesAPI.md#GetContainers) | **Get** /nodes/{node}/lxc | getContainers
[**GetCurrentContainerStatus**](NodesAPI.md#GetCurrentContainerStatus) | **Get** /nodes/{node}/lxc/{vmid}/status/current | getCurrentContainerStatus
[**GetCurrentVMStatus**](NodesAPI.md#GetCurrentVMStatus) | **Get** /nodes/{node}/qemu/{vmid}/status/current | getCurrentVMStatus
[**GetNodeTask**](NodesAPI.md#GetNodeTask) | **Get** /nodes/{node}/tasks/{upid} | getNodeTask
[**GetNodeTaskLog**](NodesAPI.md#GetNodeTaskLog) | **Get** /nodes/{node}/tasks/{upid}/log | getNodeTaskLog
[**GetNodeTaskStatus**](NodesAPI.md#GetNodeTaskStatus) | **Get** /nodes/{node}/tasks/{upid}/status | getNodeTaskStatus
[**GetNodeTasks**](NodesAPI.md#GetNodeTasks) | **Get** /nodes/{node}/tasks | getNodeTasks
[**GetStorageContent**](NodesAPI.md#GetStorageContent) | **Get** /nodes/{node}/storage/{storage}/content | getStorageContent
[**GetStorages**](NodesAPI.md#GetStorages) | **Get** /nodes/{node}/storage | getStorages
[**GetVM**](NodesAPI.md#GetVM) | **Get** /nodes/{node}/qemu/{vmid} | getVM
[**GetVMConfig**](NodesAPI.md#GetVMConfig) | **Get** /nodes/{node}/qemu/{vmid}/config | getVMConfig
[**GetVMConfigPending**](NodesAPI.md#GetVMConfigPending) | **Get** /nodes/{node}/qemu/{vmid}/pending | getVMConfigPending
[**GetVMSnapshot**](NodesAPI.md#GetVMSnapshot) | **Get** /nodes/{node}/qemu/{vmid}/snapshot/{snapname} | getVMSnapshot
[**GetVMSnapshotConfig**](NodesAPI.md#GetVMSnapshotConfig) | **Get** /nodes/{node}/qemu/{vmid}/snapshot/{snapname}/config | getVMSnapshotConfig
[**GetVMSnapshots**](NodesAPI.md#GetVMSnapshots) | **Get** /nodes/{node}/qemu/{vmid}/snapshot | getVMSnapshots
[**GetVMs**](NodesAPI.md#GetVMs) | **Get** /nodes/{node}/qemu | getVMs
[**RebootContainer**](NodesAPI.md#RebootContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/reboot | rebootContainer
[**RebootVM**](NodesAPI.md#RebootVM) | **Post** /nodes/{node}/qemu/{vmid}/status/reboot | rebootVM
[**ResizeContainerDisk**](NodesAPI.md#ResizeContainerDisk) | **Put** /nodes/{node}/lxc/{vmid}/resize | resizeContainerDisk
[**ResizeVMDisk**](NodesAPI.md#ResizeVMDisk) | **Put** /nodes/{node}/qemu/{vmid}/resize | resizeVMDisk
[**ResumeContainer**](NodesAPI.md#ResumeContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/resume | resumeContainer
[**ResumeVM**](NodesAPI.md#ResumeVM) | **Post** /nodes/{node}/qemu/{vmid}/status/resume | resumeVM
[**RollbackContainerSnapshot**](NodesAPI.md#RollbackContainerSnapshot) | **Post** /nodes/{node}/lxc/{vmid}/snapshot/{snapname}/rollback | rollbackContainerSnapshot
[**RollbackVMSnapshot**](NodesAPI.md#RollbackVMSnapshot) | **Post** /nodes/{node}/qemu/{vmid}/snapshot/{snapname}/rollback | rollbackVMSnapshot
[**ShutdownContainer**](NodesAPI.md#ShutdownContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/shutdown | shutdownContainer
[**ShutdownVM**](NodesAPI.md#ShutdownVM) | **Post** /nodes/{node}/qemu/{vmid}/status/shutdown | shutdownVM
[**StartContainer**](NodesAPI.md#StartContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/start | startContainer
[**StartVM**](NodesAPI.md#StartVM) | **Post** /nodes/{node}/qemu/{vmid}/status/start | startVM
[**StopContainer**](NodesAPI.md#StopContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/stop | stopContainer
[**StopNodeTask**](NodesAPI.md#StopNodeTask) | **Delete** /nodes/{node}/tasks/{upid} | stopNodeTask
[**StopVM**](NodesAPI.md#StopVM) | **Post** /nodes/{node}/qemu/{vmid}/status/stop | stopVM
[**SuspendContainer**](NodesAPI.md#SuspendContainer) | **Post** /nodes/{node}/lxc/{vmid}/status/suspend | suspendContainer
[**SuspendVM**](NodesAPI.md#SuspendVM) | **Post** /nodes/{node}/qemu/{vmid}/status/suspend | suspendVM
[**UpdateContainerConfigSync**](NodesAPI.md#UpdateContainerConfigSync) | **Put** /nodes/{node}/lxc/{vmid}/config | updateContainerConfigSync
[**UpdateContainerSnapshotConfig**](NodesAPI.md#UpdateContainerSnapshotConfig) | **Put** /nodes/{node}/lxc/{vmid}/snapshot/{snapname}/config | updateContainerSnapshotConfig
[**UpdateVMConfig**](NodesAPI.md#UpdateVMConfig) | **Post** /nodes/{node}/qemu/{vmid}/config | updateVMConfig
[**UpdateVMConfigSync**](NodesAPI.md#UpdateVMConfigSync) | **Put** /nodes/{node}/qemu/{vmid}/config | updateVMConfigSync
[**UpdateVMSnapshotConfig**](NodesAPI.md#UpdateVMSnapshotConfig) | **Put** /nodes/{node}/qemu/{vmid}/snapshot/{snapname}/config | updateVMSnapshotConfig
[**UploadFile**](NodesAPI.md#UploadFile) | **Post** /nodes/{node}/storage/{storage}/upload | uploadFile



## CreateContainer

> CreateVM200Response CreateContainer(ctx, node).CreateContainerRequest(createContainerRequest).Execute()

createContainer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    createContainerRequest := *openapiclient.NewCreateContainerRequest("Ostemplate_example", int64(123)) // CreateContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.CreateContainer(context.Background(), node).CreateContainerRequest(createContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.CreateContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.CreateContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createContainerRequest** | [**CreateContainerRequest**](CreateContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateContainerSnapshot

> TaskStartResponse CreateContainerSnapshot(ctx, node, vmid).CreateContainerSnapshotRequest(createContainerSnapshotRequest).Execute()

createContainerSnapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    createContainerSnapshotRequest := *openapiclient.NewCreateContainerSnapshotRequest("Snapname_example") // CreateContainerSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.CreateContainerSnapshot(context.Background(), node, vmid).CreateContainerSnapshotRequest(createContainerSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.CreateContainerSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateContainerSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.CreateContainerSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateContainerSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createContainerSnapshotRequest** | [**CreateContainerSnapshotRequest**](CreateContainerSnapshotRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateNodesSingleStorageSingleContent

> CreateNodesSingleStorageSingleContent200Response CreateNodesSingleStorageSingleContent(ctx, node, storage).CreateNodesSingleStorageSingleContentRequest(createNodesSingleStorageSingleContentRequest).Execute()

createNodesSingleStorageSingleContent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    storage := "storage_example" // string | The storage identifier.
    createNodesSingleStorageSingleContentRequest := *openapiclient.NewCreateNodesSingleStorageSingleContentRequest("TODO", "Size_example", int64(123)) // CreateNodesSingleStorageSingleContentRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.CreateNodesSingleStorageSingleContent(context.Background(), node, storage).CreateNodesSingleStorageSingleContentRequest(createNodesSingleStorageSingleContentRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.CreateNodesSingleStorageSingleContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNodesSingleStorageSingleContent`: CreateNodesSingleStorageSingleContent200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.CreateNodesSingleStorageSingleContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**storage** | **string** | The storage identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNodesSingleStorageSingleContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createNodesSingleStorageSingleContentRequest** | [**CreateNodesSingleStorageSingleContentRequest**](CreateNodesSingleStorageSingleContentRequest.md) |  | 

### Return type

[**CreateNodesSingleStorageSingleContent200Response**](CreateNodesSingleStorageSingleContent200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVM

> CreateVM200Response CreateVM(ctx, node).CreateVMRequest(createVMRequest).Execute()

createVM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    createVMRequest := *openapiclient.NewCreateVMRequest(int64(123)) // CreateVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.CreateVM(context.Background(), node).CreateVMRequest(createVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.CreateVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.CreateVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createVMRequest** | [**CreateVMRequest**](CreateVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVMSnapshot

> TaskStartResponse CreateVMSnapshot(ctx, node, vmid).CreateVMSnapshotRequest(createVMSnapshotRequest).Execute()

createVMSnapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    createVMSnapshotRequest := *openapiclient.NewCreateVMSnapshotRequest("Snapname_example") // CreateVMSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.CreateVMSnapshot(context.Background(), node, vmid).CreateVMSnapshotRequest(createVMSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.CreateVMSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVMSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.CreateVMSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVMSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createVMSnapshotRequest** | [**CreateVMSnapshotRequest**](CreateVMSnapshotRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainer

> CreateVM200Response DeleteContainer(ctx, node, vmid).DestroyUnreferencedDisks(destroyUnreferencedDisks).Force(force).Purge(purge).Execute()

deleteContainer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    destroyUnreferencedDisks := true // bool | If set, destroy additionally all disks with the VMID from all enabled storages which are not referenced in the config. (optional)
    force := true // bool | Force destroy, even if running. (optional)
    purge := true // bool | Remove container from all related configurations. For example, backup jobs, replication jobs or HA. Related ACLs and Firewall entries will *always* be removed. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.DeleteContainer(context.Background(), node, vmid).DestroyUnreferencedDisks(destroyUnreferencedDisks).Force(force).Purge(purge).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.DeleteContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.DeleteContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **destroyUnreferencedDisks** | **bool** | If set, destroy additionally all disks with the VMID from all enabled storages which are not referenced in the config. | 
 **force** | **bool** | Force destroy, even if running. | 
 **purge** | **bool** | Remove container from all related configurations. For example, backup jobs, replication jobs or HA. Related ACLs and Firewall entries will *always* be removed. | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteContainerSnapshot

> TaskStartResponse DeleteContainerSnapshot(ctx, node, vmid, snapname).Force(force).Execute()

deleteContainerSnapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.
    force := true // bool | For removal from config file, even if removing disk snapshots fails. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.DeleteContainerSnapshot(context.Background(), node, vmid, snapname).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.DeleteContainerSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteContainerSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.DeleteContainerSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteContainerSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | For removal from config file, even if removing disk snapshots fails. | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVM

> CreateVM200Response DeleteVM(ctx, node, vmid).DestroyUnreferencedDisks(destroyUnreferencedDisks).Purge(purge).Skiplock(skiplock).Execute()

deleteVM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    destroyUnreferencedDisks := true // bool | If set, destroy additionally all disks not referenced in the config but with a matching VMID from all enabled storages. (optional)
    purge := true // bool | Remove VMID from configurations, like backup & replication jobs and HA. (optional)
    skiplock := true // bool | Ignore locks - only root is allowed to use this option. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.DeleteVM(context.Background(), node, vmid).DestroyUnreferencedDisks(destroyUnreferencedDisks).Purge(purge).Skiplock(skiplock).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.DeleteVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.DeleteVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **destroyUnreferencedDisks** | **bool** | If set, destroy additionally all disks not referenced in the config but with a matching VMID from all enabled storages. | 
 **purge** | **bool** | Remove VMID from configurations, like backup &amp; replication jobs and HA. | 
 **skiplock** | **bool** | Ignore locks - only root is allowed to use this option. | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVMSnapshot

> TaskStartResponse DeleteVMSnapshot(ctx, node, vmid, snapname).Force(force).Execute()

deleteVMSnapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.
    force := true // bool | For removal from config file, even if removing disk snapshots fails. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.DeleteVMSnapshot(context.Background(), node, vmid, snapname).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.DeleteVMSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVMSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.DeleteVMSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVMSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **bool** | For removal from config file, even if removing disk snapshots fails. | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainer

> GetVM200Response GetContainer(ctx, node, vmid).Execute()

getContainer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetContainer(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainer`: GetVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVM200Response**](GetVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerConfig

> GetContainerConfig200Response GetContainerConfig(ctx, node, vmid).Current(current).Snapshot(snapshot).Execute()

getContainerConfig



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    current := true // bool | Get current values (instead of pending values). (optional)
    snapshot := "snapshot_example" // string | Fetch config values from given snapshot. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetContainerConfig(context.Background(), node, vmid).Current(current).Snapshot(snapshot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetContainerConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerConfig`: GetContainerConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetContainerConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **current** | **bool** | Get current values (instead of pending values). | 
 **snapshot** | **string** | Fetch config values from given snapshot. | 

### Return type

[**GetContainerConfig200Response**](GetContainerConfig200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerConfigPending

> GetContainerConfigPending200Response GetContainerConfigPending(ctx, node, vmid).Execute()

getContainerConfigPending



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetContainerConfigPending(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetContainerConfigPending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerConfigPending`: GetContainerConfigPending200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetContainerConfigPending`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerConfigPendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetContainerConfigPending200Response**](GetContainerConfigPending200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerSnapshot

> GetVMSnapshot200Response GetContainerSnapshot(ctx, node, vmid, snapname).Execute()

getContainerSnapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetContainerSnapshot(context.Background(), node, vmid, snapname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetContainerSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSnapshot`: GetVMSnapshot200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetContainerSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetVMSnapshot200Response**](GetVMSnapshot200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerSnapshotConfig

> GetVMSnapshotConfig200Response GetContainerSnapshotConfig(ctx, node, vmid, snapname).Execute()

getContainerSnapshotConfig



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetContainerSnapshotConfig(context.Background(), node, vmid, snapname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetContainerSnapshotConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSnapshotConfig`: GetVMSnapshotConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetContainerSnapshotConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSnapshotConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetVMSnapshotConfig200Response**](GetVMSnapshotConfig200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerSnapshots

> GetContainerSnapshots200Response GetContainerSnapshots(ctx, node, vmid).Execute()

getContainerSnapshots



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetContainerSnapshots(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetContainerSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerSnapshots`: GetContainerSnapshots200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetContainerSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetContainerSnapshots200Response**](GetContainerSnapshots200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainerStatus

> GetVM200Response GetContainerStatus(ctx, node, vmid).Execute()

getContainerStatus



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetContainerStatus(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetContainerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainerStatus`: GetVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetContainerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVM200Response**](GetVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContainers

> GetContainers200Response GetContainers(ctx, node).Execute()

getContainers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetContainers(context.Background(), node).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetContainers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContainers`: GetContainers200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetContainers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContainersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetContainers200Response**](GetContainers200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentContainerStatus

> GetCurrentContainerStatus200Response GetCurrentContainerStatus(ctx, node, vmid).Execute()

getCurrentContainerStatus



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetCurrentContainerStatus(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetCurrentContainerStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentContainerStatus`: GetCurrentContainerStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetCurrentContainerStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentContainerStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCurrentContainerStatus200Response**](GetCurrentContainerStatus200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrentVMStatus

> GetCurrentVMStatus200Response GetCurrentVMStatus(ctx, node, vmid).Execute()

getCurrentVMStatus



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetCurrentVMStatus(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetCurrentVMStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentVMStatus`: GetCurrentVMStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetCurrentVMStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentVMStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetCurrentVMStatus200Response**](GetCurrentVMStatus200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeTask

> GetVMSnapshot200Response GetNodeTask(ctx, node, upid).Execute()

getNodeTask



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    upid := "upid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetNodeTask(context.Background(), node, upid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetNodeTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeTask`: GetVMSnapshot200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetNodeTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**upid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVMSnapshot200Response**](GetVMSnapshot200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeTaskLog

> GetNodeTaskLog200Response GetNodeTaskLog(ctx, node, upid).Download(download).Limit(limit).Start(start).Execute()

getNodeTaskLog



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    upid := "upid_example" // string | The task's unique ID.
    download := true // bool | Whether the tasklog file should be downloaded. This parameter can't be used in conjunction with other parameters (optional)
    limit := int64(789) // int64 | The amount of lines to read from the tasklog. (optional)
    start := int64(789) // int64 | Start at this line when reading the tasklog (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetNodeTaskLog(context.Background(), node, upid).Download(download).Limit(limit).Start(start).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetNodeTaskLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeTaskLog`: GetNodeTaskLog200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetNodeTaskLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**upid** | **string** | The task&#39;s unique ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTaskLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **download** | **bool** | Whether the tasklog file should be downloaded. This parameter can&#39;t be used in conjunction with other parameters | 
 **limit** | **int64** | The amount of lines to read from the tasklog. | 
 **start** | **int64** | Start at this line when reading the tasklog | 

### Return type

[**GetNodeTaskLog200Response**](GetNodeTaskLog200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeTaskStatus

> GetNodeTaskStatus200Response GetNodeTaskStatus(ctx, node, upid).Execute()

getNodeTaskStatus



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    upid := "upid_example" // string | The task's unique ID.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetNodeTaskStatus(context.Background(), node, upid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetNodeTaskStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeTaskStatus`: GetNodeTaskStatus200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetNodeTaskStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**upid** | **string** | The task&#39;s unique ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTaskStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetNodeTaskStatus200Response**](GetNodeTaskStatus200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeTasks

> GetNodeTasks200Response GetNodeTasks(ctx, node).Errors(errors).Limit(limit).Since(since).Source(source).Start(start).Statusfilter(statusfilter).Typefilter(typefilter).Until(until).Userfilter(userfilter).Vmid(vmid).Execute()

getNodeTasks



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    errors := true // bool | Only list tasks with a status of ERROR. (optional)
    limit := int64(789) // int64 | Only list this amount of tasks. (optional)
    since := int64(789) // int64 | Only list tasks since this UNIX epoch. (optional)
    source := "source_example" // string | List archived, active or all tasks. (optional)
    start := int64(789) // int64 | List tasks beginning from this offset. (optional)
    statusfilter := "statusfilter_example" // string | List of Task States that should be returned. (optional)
    typefilter := "typefilter_example" // string | Only list tasks of this type (e.g., vzstart, vzdump). (optional)
    until := int64(789) // int64 | Only list tasks until this UNIX epoch. (optional)
    userfilter := "userfilter_example" // string | Only list tasks from this user. (optional)
    vmid := int64(789) // int64 | Only list tasks for this VM. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetNodeTasks(context.Background(), node).Errors(errors).Limit(limit).Since(since).Source(source).Start(start).Statusfilter(statusfilter).Typefilter(typefilter).Until(until).Userfilter(userfilter).Vmid(vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetNodeTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeTasks`: GetNodeTasks200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetNodeTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **errors** | **bool** | Only list tasks with a status of ERROR. | 
 **limit** | **int64** | Only list this amount of tasks. | 
 **since** | **int64** | Only list tasks since this UNIX epoch. | 
 **source** | **string** | List archived, active or all tasks. | 
 **start** | **int64** | List tasks beginning from this offset. | 
 **statusfilter** | **string** | List of Task States that should be returned. | 
 **typefilter** | **string** | Only list tasks of this type (e.g., vzstart, vzdump). | 
 **until** | **int64** | Only list tasks until this UNIX epoch. | 
 **userfilter** | **string** | Only list tasks from this user. | 
 **vmid** | **int64** | Only list tasks for this VM. | 

### Return type

[**GetNodeTasks200Response**](GetNodeTasks200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorageContent

> GetStorageContent200Response GetStorageContent(ctx, node, storage).Content(content).Vmid(vmid).Execute()

getStorageContent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    storage := "storage_example" // string | The storage identifier.
    content := "content_example" // string | Only list content of this type. (optional)
    vmid := int64(789) // int64 | Only list images for this VM (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetStorageContent(context.Background(), node, storage).Content(content).Vmid(vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetStorageContent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorageContent`: GetStorageContent200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetStorageContent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**storage** | **string** | The storage identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorageContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content** | **string** | Only list content of this type. | 
 **vmid** | **int64** | Only list images for this VM | 

### Return type

[**GetStorageContent200Response**](GetStorageContent200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStorages

> GetStorages200Response GetStorages(ctx, node).Content(content).Enabled(enabled).Format(format).Storage(storage).Target(target).Execute()

getStorages



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    content := "content_example" // string | Only list stores which support this content type. (optional)
    enabled := true // bool | Only list stores which are enabled (not disabled in config). (optional)
    format := true // bool | Include information about formats (optional)
    storage := "storage_example" // string | Only list status for  specified storage (optional)
    target := "target_example" // string | If target is different to 'node', we only lists shared storages which content is accessible on this 'node' and the specified 'target' node. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetStorages(context.Background(), node).Content(content).Enabled(enabled).Format(format).Storage(storage).Target(target).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetStorages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorages`: GetStorages200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetStorages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStoragesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **content** | **string** | Only list stores which support this content type. | 
 **enabled** | **bool** | Only list stores which are enabled (not disabled in config). | 
 **format** | **bool** | Include information about formats | 
 **storage** | **string** | Only list status for  specified storage | 
 **target** | **string** | If target is different to &#39;node&#39;, we only lists shared storages which content is accessible on this &#39;node&#39; and the specified &#39;target&#39; node. | 

### Return type

[**GetStorages200Response**](GetStorages200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVM

> GetVM200Response GetVM(ctx, node, vmid).Execute()

getVM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetVM(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVM`: GetVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVM200Response**](GetVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMConfig

> GetVMConfig200Response GetVMConfig(ctx, node, vmid).Current(current).Snapshot(snapshot).Execute()

getVMConfig



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    current := true // bool | Get current values (instead of pending values). (optional)
    snapshot := "snapshot_example" // string | Fetch config values from given snapshot. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetVMConfig(context.Background(), node, vmid).Current(current).Snapshot(snapshot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetVMConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMConfig`: GetVMConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetVMConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **current** | **bool** | Get current values (instead of pending values). | 
 **snapshot** | **string** | Fetch config values from given snapshot. | 

### Return type

[**GetVMConfig200Response**](GetVMConfig200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMConfigPending

> GetVMConfigPending200Response GetVMConfigPending(ctx, node, vmid).Execute()

getVMConfigPending



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetVMConfigPending(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetVMConfigPending``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMConfigPending`: GetVMConfigPending200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetVMConfigPending`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMConfigPendingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVMConfigPending200Response**](GetVMConfigPending200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMSnapshot

> GetVMSnapshot200Response GetVMSnapshot(ctx, node, vmid, snapname).Execute()

getVMSnapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetVMSnapshot(context.Background(), node, vmid, snapname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetVMSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMSnapshot`: GetVMSnapshot200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetVMSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetVMSnapshot200Response**](GetVMSnapshot200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMSnapshotConfig

> GetVMSnapshotConfig200Response GetVMSnapshotConfig(ctx, node, vmid, snapname).Execute()

getVMSnapshotConfig



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetVMSnapshotConfig(context.Background(), node, vmid, snapname).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetVMSnapshotConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMSnapshotConfig`: GetVMSnapshotConfig200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetVMSnapshotConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMSnapshotConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**GetVMSnapshotConfig200Response**](GetVMSnapshotConfig200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMSnapshots

> GetVMSnapshots200Response GetVMSnapshots(ctx, node, vmid).Execute()

getVMSnapshots



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetVMSnapshots(context.Background(), node, vmid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetVMSnapshots``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMSnapshots`: GetVMSnapshots200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetVMSnapshots`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMSnapshotsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetVMSnapshots200Response**](GetVMSnapshots200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVMs

> GetVMs200Response GetVMs(ctx, node).Full(full).Execute()

getVMs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    full := true // bool | Determine the full status of active VMs. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.GetVMs(context.Background(), node).Full(full).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.GetVMs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVMs`: GetVMs200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.GetVMs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVMsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **full** | **bool** | Determine the full status of active VMs. | 

### Return type

[**GetVMs200Response**](GetVMs200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootContainer

> CreateVM200Response RebootContainer(ctx, node, vmid).RebootContainerRequest(rebootContainerRequest).Execute()

rebootContainer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    rebootContainerRequest := *openapiclient.NewRebootContainerRequest() // RebootContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.RebootContainer(context.Background(), node, vmid).RebootContainerRequest(rebootContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.RebootContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.RebootContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rebootContainerRequest** | [**RebootContainerRequest**](RebootContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebootVM

> CreateVM200Response RebootVM(ctx, node, vmid).RebootVMRequest(rebootVMRequest).Execute()

rebootVM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    rebootVMRequest := *openapiclient.NewRebootVMRequest() // RebootVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.RebootVM(context.Background(), node, vmid).RebootVMRequest(rebootVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.RebootVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebootVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.RebootVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRebootVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **rebootVMRequest** | [**RebootVMRequest**](RebootVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeContainerDisk

> TaskStartResponse ResizeContainerDisk(ctx, node, vmid).ResizeContainerDiskRequest(resizeContainerDiskRequest).Execute()

resizeContainerDisk



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    resizeContainerDiskRequest := *openapiclient.NewResizeContainerDiskRequest("Disk_example", "Size_example") // ResizeContainerDiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.ResizeContainerDisk(context.Background(), node, vmid).ResizeContainerDiskRequest(resizeContainerDiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ResizeContainerDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeContainerDisk`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ResizeContainerDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeContainerDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resizeContainerDiskRequest** | [**ResizeContainerDiskRequest**](ResizeContainerDiskRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResizeVMDisk

> TaskStartResponse ResizeVMDisk(ctx, node, vmid).ResizeVMDiskRequest(resizeVMDiskRequest).Execute()

resizeVMDisk



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    resizeVMDiskRequest := *openapiclient.NewResizeVMDiskRequest("Disk_example", "Size_example") // ResizeVMDiskRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.ResizeVMDisk(context.Background(), node, vmid).ResizeVMDiskRequest(resizeVMDiskRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ResizeVMDisk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResizeVMDisk`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ResizeVMDisk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResizeVMDiskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resizeVMDiskRequest** | [**ResizeVMDiskRequest**](ResizeVMDiskRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeContainer

> CreateVM200Response ResumeContainer(ctx, node, vmid).Body(body).Execute()

resumeContainer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.ResumeContainer(context.Background(), node, vmid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ResumeContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ResumeContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeVM

> CreateVM200Response ResumeVM(ctx, node, vmid).ResumeVMRequest(resumeVMRequest).Execute()

resumeVM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    resumeVMRequest := *openapiclient.NewResumeVMRequest() // ResumeVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.ResumeVM(context.Background(), node, vmid).ResumeVMRequest(resumeVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ResumeVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResumeVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ResumeVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resumeVMRequest** | [**ResumeVMRequest**](ResumeVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackContainerSnapshot

> TaskStartResponse RollbackContainerSnapshot(ctx, node, vmid, snapname).RollbackContainerSnapshotRequest(rollbackContainerSnapshotRequest).Execute()

rollbackContainerSnapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.
    rollbackContainerSnapshotRequest := *openapiclient.NewRollbackContainerSnapshotRequest() // RollbackContainerSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.RollbackContainerSnapshot(context.Background(), node, vmid, snapname).RollbackContainerSnapshotRequest(rollbackContainerSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.RollbackContainerSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackContainerSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.RollbackContainerSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackContainerSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **rollbackContainerSnapshotRequest** | [**RollbackContainerSnapshotRequest**](RollbackContainerSnapshotRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RollbackVMSnapshot

> TaskStartResponse RollbackVMSnapshot(ctx, node, vmid, snapname).RollbackVMSnapshotRequest(rollbackVMSnapshotRequest).Execute()

rollbackVMSnapshot



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.
    rollbackVMSnapshotRequest := *openapiclient.NewRollbackVMSnapshotRequest() // RollbackVMSnapshotRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.RollbackVMSnapshot(context.Background(), node, vmid, snapname).RollbackVMSnapshotRequest(rollbackVMSnapshotRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.RollbackVMSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RollbackVMSnapshot`: TaskStartResponse
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.RollbackVMSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRollbackVMSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **rollbackVMSnapshotRequest** | [**RollbackVMSnapshotRequest**](RollbackVMSnapshotRequest.md) |  | 

### Return type

[**TaskStartResponse**](TaskStartResponse.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShutdownContainer

> CreateVM200Response ShutdownContainer(ctx, node, vmid).ShutdownContainerRequest(shutdownContainerRequest).Execute()

shutdownContainer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    shutdownContainerRequest := *openapiclient.NewShutdownContainerRequest() // ShutdownContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.ShutdownContainer(context.Background(), node, vmid).ShutdownContainerRequest(shutdownContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ShutdownContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShutdownContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ShutdownContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shutdownContainerRequest** | [**ShutdownContainerRequest**](ShutdownContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ShutdownVM

> CreateVM200Response ShutdownVM(ctx, node, vmid).ShutdownVMRequest(shutdownVMRequest).Execute()

shutdownVM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    shutdownVMRequest := *openapiclient.NewShutdownVMRequest() // ShutdownVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.ShutdownVM(context.Background(), node, vmid).ShutdownVMRequest(shutdownVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.ShutdownVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ShutdownVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.ShutdownVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiShutdownVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **shutdownVMRequest** | [**ShutdownVMRequest**](ShutdownVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartContainer

> CreateVM200Response StartContainer(ctx, node, vmid).StartContainerRequest(startContainerRequest).Execute()

startContainer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    startContainerRequest := *openapiclient.NewStartContainerRequest() // StartContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.StartContainer(context.Background(), node, vmid).StartContainerRequest(startContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.StartContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.StartContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startContainerRequest** | [**StartContainerRequest**](StartContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartVM

> CreateVM200Response StartVM(ctx, node, vmid).StartVMRequest(startVMRequest).Execute()

startVM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    startVMRequest := *openapiclient.NewStartVMRequest() // StartVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.StartVM(context.Background(), node, vmid).StartVMRequest(startVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.StartVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.StartVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startVMRequest** | [**StartVMRequest**](StartVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopContainer

> CreateVM200Response StopContainer(ctx, node, vmid).StopContainerRequest(stopContainerRequest).Execute()

stopContainer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    stopContainerRequest := *openapiclient.NewStopContainerRequest() // StopContainerRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.StopContainer(context.Background(), node, vmid).StopContainerRequest(stopContainerRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.StopContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.StopContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stopContainerRequest** | [**StopContainerRequest**](StopContainerRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopNodeTask

> CreateVM200Response StopNodeTask(ctx, node, upid).Execute()

stopNodeTask



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    upid := "upid_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.StopNodeTask(context.Background(), node, upid).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.StopNodeTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopNodeTask`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.StopNodeTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**upid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopNodeTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopVM

> CreateVM200Response StopVM(ctx, node, vmid).StopVMRequest(stopVMRequest).Execute()

stopVM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    stopVMRequest := *openapiclient.NewStopVMRequest() // StopVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.StopVM(context.Background(), node, vmid).StopVMRequest(stopVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.StopVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StopVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.StopVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stopVMRequest** | [**StopVMRequest**](StopVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendContainer

> CreateVM200Response SuspendContainer(ctx, node, vmid).Body(body).Execute()

suspendContainer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.SuspendContainer(context.Background(), node, vmid).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.SuspendContainer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendContainer`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.SuspendContainer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendContainerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuspendVM

> CreateVM200Response SuspendVM(ctx, node, vmid).SuspendVMRequest(suspendVMRequest).Execute()

suspendVM



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    suspendVMRequest := *openapiclient.NewSuspendVMRequest() // SuspendVMRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.SuspendVM(context.Background(), node, vmid).SuspendVMRequest(suspendVMRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.SuspendVM``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SuspendVM`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.SuspendVM`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSuspendVMRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **suspendVMRequest** | [**SuspendVMRequest**](SuspendVMRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerConfigSync

> CreateVM200Response UpdateContainerConfigSync(ctx, node, vmid).UpdateContainerConfigSyncRequest(updateContainerConfigSyncRequest).Execute()

updateContainerConfigSync



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    updateContainerConfigSyncRequest := *openapiclient.NewUpdateContainerConfigSyncRequest() // UpdateContainerConfigSyncRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.UpdateContainerConfigSync(context.Background(), node, vmid).UpdateContainerConfigSyncRequest(updateContainerConfigSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.UpdateContainerConfigSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContainerConfigSync`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.UpdateContainerConfigSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerConfigSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateContainerConfigSyncRequest** | [**UpdateContainerConfigSyncRequest**](UpdateContainerConfigSyncRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateContainerSnapshotConfig

> CreateVM200Response UpdateContainerSnapshotConfig(ctx, node, vmid, snapname).UpdateContainerSnapshotConfigRequest(updateContainerSnapshotConfigRequest).Execute()

updateContainerSnapshotConfig



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.
    updateContainerSnapshotConfigRequest := *openapiclient.NewUpdateContainerSnapshotConfigRequest() // UpdateContainerSnapshotConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.UpdateContainerSnapshotConfig(context.Background(), node, vmid, snapname).UpdateContainerSnapshotConfigRequest(updateContainerSnapshotConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.UpdateContainerSnapshotConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateContainerSnapshotConfig`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.UpdateContainerSnapshotConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateContainerSnapshotConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateContainerSnapshotConfigRequest** | [**UpdateContainerSnapshotConfigRequest**](UpdateContainerSnapshotConfigRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMConfig

> CreateVM200Response UpdateVMConfig(ctx, node, vmid).UpdateVMConfigRequest(updateVMConfigRequest).Execute()

updateVMConfig



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    updateVMConfigRequest := *openapiclient.NewUpdateVMConfigRequest() // UpdateVMConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.UpdateVMConfig(context.Background(), node, vmid).UpdateVMConfigRequest(updateVMConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.UpdateVMConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVMConfig`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.UpdateVMConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMConfigRequest** | [**UpdateVMConfigRequest**](UpdateVMConfigRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMConfigSync

> CreateVM200Response UpdateVMConfigSync(ctx, node, vmid).UpdateVMConfigSyncRequest(updateVMConfigSyncRequest).Execute()

updateVMConfigSync



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    updateVMConfigSyncRequest := *openapiclient.NewUpdateVMConfigSyncRequest() // UpdateVMConfigSyncRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.UpdateVMConfigSync(context.Background(), node, vmid).UpdateVMConfigSyncRequest(updateVMConfigSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.UpdateVMConfigSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVMConfigSync`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.UpdateVMConfigSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMConfigSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateVMConfigSyncRequest** | [**UpdateVMConfigSyncRequest**](UpdateVMConfigSyncRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVMSnapshotConfig

> CreateVM200Response UpdateVMSnapshotConfig(ctx, node, vmid, snapname).UpdateVMSnapshotConfigRequest(updateVMSnapshotConfigRequest).Execute()

updateVMSnapshotConfig



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    vmid := int64(789) // int64 | The (unique) ID of the VM.
    snapname := "snapname_example" // string | The name of the snapshot.
    updateVMSnapshotConfigRequest := *openapiclient.NewUpdateVMSnapshotConfigRequest() // UpdateVMSnapshotConfigRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.UpdateVMSnapshotConfig(context.Background(), node, vmid, snapname).UpdateVMSnapshotConfigRequest(updateVMSnapshotConfigRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.UpdateVMSnapshotConfig``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVMSnapshotConfig`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.UpdateVMSnapshotConfig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**vmid** | **int64** | The (unique) ID of the VM. | 
**snapname** | **string** | The name of the snapshot. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVMSnapshotConfigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **updateVMSnapshotConfigRequest** | [**UpdateVMSnapshotConfigRequest**](UpdateVMSnapshotConfigRequest.md) |  | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> CreateVM200Response UploadFile(ctx, node, storage).Content(content).Filename(filename).Checksum(checksum).ChecksumAlgorithm(checksumAlgorithm).Tmpfilename(tmpfilename).Execute()

uploadFile



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/DirkTheDaring/px-api-client-go"
)

func main() {
    node := "node_example" // string | The cluster node name.
    storage := "storage_example" // string | The storage identifier.
    content := "content_example" // string | Content type.
    filename := os.NewFile(1234, "some_file") // *os.File | The name of the file to create. Caution: This will be normalized!
    checksum := "checksum_example" // string | The expected checksum of the file. (optional)
    checksumAlgorithm := "checksumAlgorithm_example" // string | The algorithm to calculate the checksum of the file. (optional)
    tmpfilename := "tmpfilename_example" // string | The source file name. This parameter is usually set by the REST handler. You can only overwrite it when connecting to the trusted port on localhost. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NodesAPI.UploadFile(context.Background(), node, storage).Content(content).Filename(filename).Checksum(checksum).ChecksumAlgorithm(checksumAlgorithm).Tmpfilename(tmpfilename).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NodesAPI.UploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFile`: CreateVM200Response
    fmt.Fprintf(os.Stdout, "Response from `NodesAPI.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**node** | **string** | The cluster node name. | 
**storage** | **string** | The storage identifier. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **content** | **string** | Content type. | 
 **filename** | ***os.File** | The name of the file to create. Caution: This will be normalized! | 
 **checksum** | **string** | The expected checksum of the file. | 
 **checksumAlgorithm** | **string** | The algorithm to calculate the checksum of the file. | 
 **tmpfilename** | **string** | The source file name. This parameter is usually set by the REST handler. You can only overwrite it when connecting to the trusted port on localhost. | 

### Return type

[**CreateVM200Response**](CreateVM200Response.md)

### Authorization

[cookie](../README.md#cookie), [token](../README.md#token)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

