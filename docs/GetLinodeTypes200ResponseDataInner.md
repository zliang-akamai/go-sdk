# GetLinodeTypes200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addons** | Pointer to [**GetLinodeTypes200ResponseDataInnerAddons**](GetLinodeTypes200ResponseDataInnerAddons.md) |  | [optional] 
**Class** | Pointer to **string** | The class of the Linode Type.  We currently offer six classes of compute instances:    - &#x60;nanode&#x60; - Nanode instances are good for low-duty workloads, where performance isn&#39;t critical. __Note__. As of June 16th, 2020, Nanodes became 1 GB Linodes in the Cloud Manager, however, the API, the CLI, and billing will continue to refer to these instances as Nanodes.   - &#x60;standard&#x60; - Standard Shared instances are good for medium-duty workloads and are a good mix of performance, resources, and price. __Note__. As of June 16th, 2020, Standard Linodes in the Cloud Manager became Shared Linodes, however, the API, the CLI, and billing will continue to refer to these instances as Standard Linodes.   - &#x60;dedicated&#x60; - Dedicated CPU instances are good for full-duty workloads where consistent performance is important.   - &#x60;premium&#x60; (limited Regions) - In addition to the features of Dedicated instances, Premium instances come equipped with the latest AMD EPYC&amp;trade; CPUs, ensuring your applications are running on the latest hardware with consistently high performance. Only available in [regions](https://techdocs.akamai.com/linode-api/reference/get-regions) with \&quot;Premium Plans\&quot; in their &#x60;capabilities&#x60;   - &#x60;gpu&#x60; (limited Regions) - Linodes with dedicated NVIDIA Quadro&amp;reg; RTX 6000 GPUs accelerate highly specialized applications such as machine learning, AI, and video transcoding. Only available in [regions](https://techdocs.akamai.com/linode-api/reference/get-regions) with \&quot;GPU Linodes\&quot; in their &#x60;capabilities&#x60;   - &#x60;highmem&#x60; - High Memory instances favor RAM over other resources, and can be good for memory hungry use cases like caching and in-memory databases. All High Memory plans contain dedicated CPU cores. | [optional] [readonly] 
**Disk** | Pointer to **int32** | The Disk size, in MB, of the Linode Type. | [optional] [readonly] 
**Gpus** | Pointer to **int32** | The number of GPUs this Linode Type offers. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID representing the Linode Type. | [optional] [readonly] 
**Label** | Pointer to **string** | The Linode Type&#39;s label is for display purposes only. | [optional] [readonly] 
**Memory** | Pointer to **int32** | Amount of RAM included in this Linode Type. | [optional] [readonly] 
**NetworkOut** | Pointer to **int32** | The Mbits outbound bandwidth allocation. | [optional] [readonly] 
**Price** | Pointer to [**GetLinodeTypes200ResponseDataInnerPrice**](GetLinodeTypes200ResponseDataInnerPrice.md) |  | [optional] 
**RegionPrices** | Pointer to [**[]GetLinodeTypes200ResponseDataInnerRegionPricesInner**](GetLinodeTypes200ResponseDataInnerRegionPricesInner.md) |  | [optional] 
**Successor** | Pointer to **NullableString** | The Linode Type that a [mutate](https://techdocs.akamai.com/linode-api/reference/post-mutate-linode-instance) will upgrade to for a Linode of this type.  If &#x60;null&#x60;, a Linode of this type may not mutate. | [optional] [readonly] 
**Transfer** | Pointer to **int32** | The monthly outbound transfer amount, in MB. | [optional] [readonly] 
**Vcpus** | Pointer to **int32** | The number of VCPU cores this Linode Type offers. | [optional] [readonly] 

## Methods

### NewGetLinodeTypes200ResponseDataInner

`func NewGetLinodeTypes200ResponseDataInner() *GetLinodeTypes200ResponseDataInner`

NewGetLinodeTypes200ResponseDataInner instantiates a new GetLinodeTypes200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeTypes200ResponseDataInnerWithDefaults

`func NewGetLinodeTypes200ResponseDataInnerWithDefaults() *GetLinodeTypes200ResponseDataInner`

NewGetLinodeTypes200ResponseDataInnerWithDefaults instantiates a new GetLinodeTypes200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddons

`func (o *GetLinodeTypes200ResponseDataInner) GetAddons() GetLinodeTypes200ResponseDataInnerAddons`

GetAddons returns the Addons field if non-nil, zero value otherwise.

### GetAddonsOk

`func (o *GetLinodeTypes200ResponseDataInner) GetAddonsOk() (*GetLinodeTypes200ResponseDataInnerAddons, bool)`

GetAddonsOk returns a tuple with the Addons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddons

`func (o *GetLinodeTypes200ResponseDataInner) SetAddons(v GetLinodeTypes200ResponseDataInnerAddons)`

SetAddons sets Addons field to given value.

### HasAddons

`func (o *GetLinodeTypes200ResponseDataInner) HasAddons() bool`

HasAddons returns a boolean if a field has been set.

### GetClass

`func (o *GetLinodeTypes200ResponseDataInner) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *GetLinodeTypes200ResponseDataInner) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *GetLinodeTypes200ResponseDataInner) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *GetLinodeTypes200ResponseDataInner) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetDisk

`func (o *GetLinodeTypes200ResponseDataInner) GetDisk() int32`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *GetLinodeTypes200ResponseDataInner) GetDiskOk() (*int32, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *GetLinodeTypes200ResponseDataInner) SetDisk(v int32)`

SetDisk sets Disk field to given value.

### HasDisk

`func (o *GetLinodeTypes200ResponseDataInner) HasDisk() bool`

HasDisk returns a boolean if a field has been set.

### GetGpus

`func (o *GetLinodeTypes200ResponseDataInner) GetGpus() int32`

GetGpus returns the Gpus field if non-nil, zero value otherwise.

### GetGpusOk

`func (o *GetLinodeTypes200ResponseDataInner) GetGpusOk() (*int32, bool)`

GetGpusOk returns a tuple with the Gpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGpus

`func (o *GetLinodeTypes200ResponseDataInner) SetGpus(v int32)`

SetGpus sets Gpus field to given value.

### HasGpus

`func (o *GetLinodeTypes200ResponseDataInner) HasGpus() bool`

HasGpus returns a boolean if a field has been set.

### GetId

`func (o *GetLinodeTypes200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLinodeTypes200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLinodeTypes200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetLinodeTypes200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetLinodeTypes200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetLinodeTypes200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetLinodeTypes200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetLinodeTypes200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMemory

`func (o *GetLinodeTypes200ResponseDataInner) GetMemory() int32`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *GetLinodeTypes200ResponseDataInner) GetMemoryOk() (*int32, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *GetLinodeTypes200ResponseDataInner) SetMemory(v int32)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *GetLinodeTypes200ResponseDataInner) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetNetworkOut

`func (o *GetLinodeTypes200ResponseDataInner) GetNetworkOut() int32`

GetNetworkOut returns the NetworkOut field if non-nil, zero value otherwise.

### GetNetworkOutOk

`func (o *GetLinodeTypes200ResponseDataInner) GetNetworkOutOk() (*int32, bool)`

GetNetworkOutOk returns a tuple with the NetworkOut field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOut

`func (o *GetLinodeTypes200ResponseDataInner) SetNetworkOut(v int32)`

SetNetworkOut sets NetworkOut field to given value.

### HasNetworkOut

`func (o *GetLinodeTypes200ResponseDataInner) HasNetworkOut() bool`

HasNetworkOut returns a boolean if a field has been set.

### GetPrice

`func (o *GetLinodeTypes200ResponseDataInner) GetPrice() GetLinodeTypes200ResponseDataInnerPrice`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetLinodeTypes200ResponseDataInner) GetPriceOk() (*GetLinodeTypes200ResponseDataInnerPrice, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetLinodeTypes200ResponseDataInner) SetPrice(v GetLinodeTypes200ResponseDataInnerPrice)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetLinodeTypes200ResponseDataInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetRegionPrices

`func (o *GetLinodeTypes200ResponseDataInner) GetRegionPrices() []GetLinodeTypes200ResponseDataInnerRegionPricesInner`

GetRegionPrices returns the RegionPrices field if non-nil, zero value otherwise.

### GetRegionPricesOk

`func (o *GetLinodeTypes200ResponseDataInner) GetRegionPricesOk() (*[]GetLinodeTypes200ResponseDataInnerRegionPricesInner, bool)`

GetRegionPricesOk returns a tuple with the RegionPrices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionPrices

`func (o *GetLinodeTypes200ResponseDataInner) SetRegionPrices(v []GetLinodeTypes200ResponseDataInnerRegionPricesInner)`

SetRegionPrices sets RegionPrices field to given value.

### HasRegionPrices

`func (o *GetLinodeTypes200ResponseDataInner) HasRegionPrices() bool`

HasRegionPrices returns a boolean if a field has been set.

### GetSuccessor

`func (o *GetLinodeTypes200ResponseDataInner) GetSuccessor() string`

GetSuccessor returns the Successor field if non-nil, zero value otherwise.

### GetSuccessorOk

`func (o *GetLinodeTypes200ResponseDataInner) GetSuccessorOk() (*string, bool)`

GetSuccessorOk returns a tuple with the Successor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessor

`func (o *GetLinodeTypes200ResponseDataInner) SetSuccessor(v string)`

SetSuccessor sets Successor field to given value.

### HasSuccessor

`func (o *GetLinodeTypes200ResponseDataInner) HasSuccessor() bool`

HasSuccessor returns a boolean if a field has been set.

### SetSuccessorNil

`func (o *GetLinodeTypes200ResponseDataInner) SetSuccessorNil(b bool)`

 SetSuccessorNil sets the value for Successor to be an explicit nil

### UnsetSuccessor
`func (o *GetLinodeTypes200ResponseDataInner) UnsetSuccessor()`

UnsetSuccessor ensures that no value is present for Successor, not even an explicit nil
### GetTransfer

`func (o *GetLinodeTypes200ResponseDataInner) GetTransfer() int32`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *GetLinodeTypes200ResponseDataInner) GetTransferOk() (*int32, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *GetLinodeTypes200ResponseDataInner) SetTransfer(v int32)`

SetTransfer sets Transfer field to given value.

### HasTransfer

`func (o *GetLinodeTypes200ResponseDataInner) HasTransfer() bool`

HasTransfer returns a boolean if a field has been set.

### GetVcpus

`func (o *GetLinodeTypes200ResponseDataInner) GetVcpus() int32`

GetVcpus returns the Vcpus field if non-nil, zero value otherwise.

### GetVcpusOk

`func (o *GetLinodeTypes200ResponseDataInner) GetVcpusOk() (*int32, bool)`

GetVcpusOk returns a tuple with the Vcpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcpus

`func (o *GetLinodeTypes200ResponseDataInner) SetVcpus(v int32)`

SetVcpus sets Vcpus field to given value.

### HasVcpus

`func (o *GetLinodeTypes200ResponseDataInner) HasVcpus() bool`

HasVcpus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


