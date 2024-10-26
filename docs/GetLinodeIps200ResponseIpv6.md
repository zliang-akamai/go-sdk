# GetLinodeIps200ResponseIpv6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Global** | Pointer to [**[]GetLinodeIps200ResponseIpv6GlobalInner**](GetLinodeIps200ResponseIpv6GlobalInner.md) | A list of IPv6 range objects assigned to this Linode. | [optional] 
**LinkLocal** | Pointer to [**GetLinodeIps200ResponseIpv6LinkLocal**](GetLinodeIps200ResponseIpv6LinkLocal.md) |  | [optional] 
**Slaac** | Pointer to [**GetLinodeIps200ResponseIpv6Slaac**](GetLinodeIps200ResponseIpv6Slaac.md) |  | [optional] 

## Methods

### NewGetLinodeIps200ResponseIpv6

`func NewGetLinodeIps200ResponseIpv6() *GetLinodeIps200ResponseIpv6`

NewGetLinodeIps200ResponseIpv6 instantiates a new GetLinodeIps200ResponseIpv6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeIps200ResponseIpv6WithDefaults

`func NewGetLinodeIps200ResponseIpv6WithDefaults() *GetLinodeIps200ResponseIpv6`

NewGetLinodeIps200ResponseIpv6WithDefaults instantiates a new GetLinodeIps200ResponseIpv6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobal

`func (o *GetLinodeIps200ResponseIpv6) GetGlobal() []GetLinodeIps200ResponseIpv6GlobalInner`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *GetLinodeIps200ResponseIpv6) GetGlobalOk() (*[]GetLinodeIps200ResponseIpv6GlobalInner, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *GetLinodeIps200ResponseIpv6) SetGlobal(v []GetLinodeIps200ResponseIpv6GlobalInner)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *GetLinodeIps200ResponseIpv6) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetLinkLocal

`func (o *GetLinodeIps200ResponseIpv6) GetLinkLocal() GetLinodeIps200ResponseIpv6LinkLocal`

GetLinkLocal returns the LinkLocal field if non-nil, zero value otherwise.

### GetLinkLocalOk

`func (o *GetLinodeIps200ResponseIpv6) GetLinkLocalOk() (*GetLinodeIps200ResponseIpv6LinkLocal, bool)`

GetLinkLocalOk returns a tuple with the LinkLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkLocal

`func (o *GetLinodeIps200ResponseIpv6) SetLinkLocal(v GetLinodeIps200ResponseIpv6LinkLocal)`

SetLinkLocal sets LinkLocal field to given value.

### HasLinkLocal

`func (o *GetLinodeIps200ResponseIpv6) HasLinkLocal() bool`

HasLinkLocal returns a boolean if a field has been set.

### GetSlaac

`func (o *GetLinodeIps200ResponseIpv6) GetSlaac() GetLinodeIps200ResponseIpv6Slaac`

GetSlaac returns the Slaac field if non-nil, zero value otherwise.

### GetSlaacOk

`func (o *GetLinodeIps200ResponseIpv6) GetSlaacOk() (*GetLinodeIps200ResponseIpv6Slaac, bool)`

GetSlaacOk returns a tuple with the Slaac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlaac

`func (o *GetLinodeIps200ResponseIpv6) SetSlaac(v GetLinodeIps200ResponseIpv6Slaac)`

SetSlaac sets Slaac field to given value.

### HasSlaac

`func (o *GetLinodeIps200ResponseIpv6) HasSlaac() bool`

HasSlaac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


