# ResumeVMRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nocheck** | Pointer to **int32** |  | [optional] 
**Skiplock** | Pointer to **int32** | Ignore locks - only root is allowed to use this option. | [optional] 

## Methods

### NewResumeVMRequest

`func NewResumeVMRequest() *ResumeVMRequest`

NewResumeVMRequest instantiates a new ResumeVMRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeVMRequestWithDefaults

`func NewResumeVMRequestWithDefaults() *ResumeVMRequest`

NewResumeVMRequestWithDefaults instantiates a new ResumeVMRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNocheck

`func (o *ResumeVMRequest) GetNocheck() int32`

GetNocheck returns the Nocheck field if non-nil, zero value otherwise.

### GetNocheckOk

`func (o *ResumeVMRequest) GetNocheckOk() (*int32, bool)`

GetNocheckOk returns a tuple with the Nocheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNocheck

`func (o *ResumeVMRequest) SetNocheck(v int32)`

SetNocheck sets Nocheck field to given value.

### HasNocheck

`func (o *ResumeVMRequest) HasNocheck() bool`

HasNocheck returns a boolean if a field has been set.

### GetSkiplock

`func (o *ResumeVMRequest) GetSkiplock() int32`

GetSkiplock returns the Skiplock field if non-nil, zero value otherwise.

### GetSkiplockOk

`func (o *ResumeVMRequest) GetSkiplockOk() (*int32, bool)`

GetSkiplockOk returns a tuple with the Skiplock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkiplock

`func (o *ResumeVMRequest) SetSkiplock(v int32)`

SetSkiplock sets Skiplock field to given value.

### HasSkiplock

`func (o *ResumeVMRequest) HasSkiplock() bool`

HasSkiplock returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


