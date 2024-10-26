# GetLinodeStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | Pointer to **[][]float32** | Percentage of CPU used. | [optional] 
**Io** | Pointer to [**GetLinodeStats200ResponseIo**](GetLinodeStats200ResponseIo.md) |  | [optional] 
**Netv4** | Pointer to [**GetLinodeStats200ResponseNetv4**](GetLinodeStats200ResponseNetv4.md) |  | [optional] 
**Netv6** | Pointer to [**GetLinodeStats200ResponseNetv6**](GetLinodeStats200ResponseNetv6.md) |  | [optional] 
**Title** | Pointer to **string** | The title for this data set. | [optional] 

## Methods

### NewGetLinodeStats200Response

`func NewGetLinodeStats200Response() *GetLinodeStats200Response`

NewGetLinodeStats200Response instantiates a new GetLinodeStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeStats200ResponseWithDefaults

`func NewGetLinodeStats200ResponseWithDefaults() *GetLinodeStats200Response`

NewGetLinodeStats200ResponseWithDefaults instantiates a new GetLinodeStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *GetLinodeStats200Response) GetCpu() [][]float32`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *GetLinodeStats200Response) GetCpuOk() (*[][]float32, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *GetLinodeStats200Response) SetCpu(v [][]float32)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *GetLinodeStats200Response) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetIo

`func (o *GetLinodeStats200Response) GetIo() GetLinodeStats200ResponseIo`

GetIo returns the Io field if non-nil, zero value otherwise.

### GetIoOk

`func (o *GetLinodeStats200Response) GetIoOk() (*GetLinodeStats200ResponseIo, bool)`

GetIoOk returns a tuple with the Io field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIo

`func (o *GetLinodeStats200Response) SetIo(v GetLinodeStats200ResponseIo)`

SetIo sets Io field to given value.

### HasIo

`func (o *GetLinodeStats200Response) HasIo() bool`

HasIo returns a boolean if a field has been set.

### GetNetv4

`func (o *GetLinodeStats200Response) GetNetv4() GetLinodeStats200ResponseNetv4`

GetNetv4 returns the Netv4 field if non-nil, zero value otherwise.

### GetNetv4Ok

`func (o *GetLinodeStats200Response) GetNetv4Ok() (*GetLinodeStats200ResponseNetv4, bool)`

GetNetv4Ok returns a tuple with the Netv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetv4

`func (o *GetLinodeStats200Response) SetNetv4(v GetLinodeStats200ResponseNetv4)`

SetNetv4 sets Netv4 field to given value.

### HasNetv4

`func (o *GetLinodeStats200Response) HasNetv4() bool`

HasNetv4 returns a boolean if a field has been set.

### GetNetv6

`func (o *GetLinodeStats200Response) GetNetv6() GetLinodeStats200ResponseNetv6`

GetNetv6 returns the Netv6 field if non-nil, zero value otherwise.

### GetNetv6Ok

`func (o *GetLinodeStats200Response) GetNetv6Ok() (*GetLinodeStats200ResponseNetv6, bool)`

GetNetv6Ok returns a tuple with the Netv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetv6

`func (o *GetLinodeStats200Response) SetNetv6(v GetLinodeStats200ResponseNetv6)`

SetNetv6 sets Netv6 field to given value.

### HasNetv6

`func (o *GetLinodeStats200Response) HasNetv6() bool`

HasNetv6 returns a boolean if a field has been set.

### GetTitle

`func (o *GetLinodeStats200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetLinodeStats200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetLinodeStats200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetLinodeStats200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


