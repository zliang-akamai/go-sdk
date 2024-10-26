# GetKernels200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Architecture** | Pointer to **string** | The architecture of this Kernel. | [optional] [readonly] 
**Built** | Pointer to **time.Time** | The date on which this Kernel was built. | [optional] [readonly] 
**Deprecated** | Pointer to **bool** | If this Kernel is marked as deprecated, this field has a value of true; otherwise, this field is false. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of this Kernel. | [optional] [readonly] 
**Kvm** | Pointer to **bool** | If this Kernel is suitable for KVM Linodes. | [optional] [readonly] 
**Label** | Pointer to **string** | The friendly name of this Kernel. | [optional] [readonly] 
**Pvops** | Pointer to **bool** | If this Kernel is suitable for paravirtualized operations. | [optional] [readonly] 
**Version** | Pointer to **string** | Linux Kernel version. | [optional] [readonly] 

## Methods

### NewGetKernels200ResponseDataInner

`func NewGetKernels200ResponseDataInner() *GetKernels200ResponseDataInner`

NewGetKernels200ResponseDataInner instantiates a new GetKernels200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetKernels200ResponseDataInnerWithDefaults

`func NewGetKernels200ResponseDataInnerWithDefaults() *GetKernels200ResponseDataInner`

NewGetKernels200ResponseDataInnerWithDefaults instantiates a new GetKernels200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchitecture

`func (o *GetKernels200ResponseDataInner) GetArchitecture() string`

GetArchitecture returns the Architecture field if non-nil, zero value otherwise.

### GetArchitectureOk

`func (o *GetKernels200ResponseDataInner) GetArchitectureOk() (*string, bool)`

GetArchitectureOk returns a tuple with the Architecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchitecture

`func (o *GetKernels200ResponseDataInner) SetArchitecture(v string)`

SetArchitecture sets Architecture field to given value.

### HasArchitecture

`func (o *GetKernels200ResponseDataInner) HasArchitecture() bool`

HasArchitecture returns a boolean if a field has been set.

### GetBuilt

`func (o *GetKernels200ResponseDataInner) GetBuilt() time.Time`

GetBuilt returns the Built field if non-nil, zero value otherwise.

### GetBuiltOk

`func (o *GetKernels200ResponseDataInner) GetBuiltOk() (*time.Time, bool)`

GetBuiltOk returns a tuple with the Built field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuilt

`func (o *GetKernels200ResponseDataInner) SetBuilt(v time.Time)`

SetBuilt sets Built field to given value.

### HasBuilt

`func (o *GetKernels200ResponseDataInner) HasBuilt() bool`

HasBuilt returns a boolean if a field has been set.

### GetDeprecated

`func (o *GetKernels200ResponseDataInner) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *GetKernels200ResponseDataInner) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *GetKernels200ResponseDataInner) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *GetKernels200ResponseDataInner) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetId

`func (o *GetKernels200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetKernels200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetKernels200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetKernels200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKvm

`func (o *GetKernels200ResponseDataInner) GetKvm() bool`

GetKvm returns the Kvm field if non-nil, zero value otherwise.

### GetKvmOk

`func (o *GetKernels200ResponseDataInner) GetKvmOk() (*bool, bool)`

GetKvmOk returns a tuple with the Kvm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvm

`func (o *GetKernels200ResponseDataInner) SetKvm(v bool)`

SetKvm sets Kvm field to given value.

### HasKvm

`func (o *GetKernels200ResponseDataInner) HasKvm() bool`

HasKvm returns a boolean if a field has been set.

### GetLabel

`func (o *GetKernels200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetKernels200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetKernels200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetKernels200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPvops

`func (o *GetKernels200ResponseDataInner) GetPvops() bool`

GetPvops returns the Pvops field if non-nil, zero value otherwise.

### GetPvopsOk

`func (o *GetKernels200ResponseDataInner) GetPvopsOk() (*bool, bool)`

GetPvopsOk returns a tuple with the Pvops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvops

`func (o *GetKernels200ResponseDataInner) SetPvops(v bool)`

SetPvops sets Pvops field to given value.

### HasPvops

`func (o *GetKernels200ResponseDataInner) HasPvops() bool`

HasPvops returns a boolean if a field has been set.

### GetVersion

`func (o *GetKernels200ResponseDataInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetKernels200ResponseDataInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetKernels200ResponseDataInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetKernels200ResponseDataInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


