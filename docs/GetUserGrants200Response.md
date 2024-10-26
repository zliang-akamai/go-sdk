# GetUserGrants200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each Database that is owned by this Account. | [optional] 
**Domain** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each Domain that is owned by this Account. | [optional] 
**Firewall** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each Firewall that is owned by this Account. | [optional] 
**Global** | Pointer to [**GetUserGrants200ResponseGlobal**](GetUserGrants200ResponseGlobal.md) |  | [optional] 
**Image** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each Image that is owned by this Account. | [optional] 
**Linode** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each Linode that is owned by this Account. | [optional] 
**Longview** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each Longview Client that is owned by this Account. | [optional] 
**Nodebalancer** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each NodeBalancer that is owned by this Account. | [optional] 
**PlacementGroup** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each Placement Group that is owned by this Account. | [optional] 
**Stackscript** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each StackScript that is owned by this Account. | [optional] 
**Volume** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each Block Storage Volume that is owned by this Account. | [optional] 
**Vpc** | Pointer to [**[]GetUserGrants200ResponseDatabaseInner**](GetUserGrants200ResponseDatabaseInner.md) | The grants this User has for each VPC that is owned by this Account. | [optional] 

## Methods

### NewGetUserGrants200Response

`func NewGetUserGrants200Response() *GetUserGrants200Response`

NewGetUserGrants200Response instantiates a new GetUserGrants200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserGrants200ResponseWithDefaults

`func NewGetUserGrants200ResponseWithDefaults() *GetUserGrants200Response`

NewGetUserGrants200ResponseWithDefaults instantiates a new GetUserGrants200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *GetUserGrants200Response) GetDatabase() []GetUserGrants200ResponseDatabaseInner`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *GetUserGrants200Response) GetDatabaseOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *GetUserGrants200Response) SetDatabase(v []GetUserGrants200ResponseDatabaseInner)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *GetUserGrants200Response) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetDomain

`func (o *GetUserGrants200Response) GetDomain() []GetUserGrants200ResponseDatabaseInner`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GetUserGrants200Response) GetDomainOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GetUserGrants200Response) SetDomain(v []GetUserGrants200ResponseDatabaseInner)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GetUserGrants200Response) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFirewall

`func (o *GetUserGrants200Response) GetFirewall() []GetUserGrants200ResponseDatabaseInner`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *GetUserGrants200Response) GetFirewallOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *GetUserGrants200Response) SetFirewall(v []GetUserGrants200ResponseDatabaseInner)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *GetUserGrants200Response) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetGlobal

`func (o *GetUserGrants200Response) GetGlobal() GetUserGrants200ResponseGlobal`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *GetUserGrants200Response) GetGlobalOk() (*GetUserGrants200ResponseGlobal, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *GetUserGrants200Response) SetGlobal(v GetUserGrants200ResponseGlobal)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *GetUserGrants200Response) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetImage

`func (o *GetUserGrants200Response) GetImage() []GetUserGrants200ResponseDatabaseInner`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GetUserGrants200Response) GetImageOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GetUserGrants200Response) SetImage(v []GetUserGrants200ResponseDatabaseInner)`

SetImage sets Image field to given value.

### HasImage

`func (o *GetUserGrants200Response) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLinode

`func (o *GetUserGrants200Response) GetLinode() []GetUserGrants200ResponseDatabaseInner`

GetLinode returns the Linode field if non-nil, zero value otherwise.

### GetLinodeOk

`func (o *GetUserGrants200Response) GetLinodeOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetLinodeOk returns a tuple with the Linode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinode

`func (o *GetUserGrants200Response) SetLinode(v []GetUserGrants200ResponseDatabaseInner)`

SetLinode sets Linode field to given value.

### HasLinode

`func (o *GetUserGrants200Response) HasLinode() bool`

HasLinode returns a boolean if a field has been set.

### GetLongview

`func (o *GetUserGrants200Response) GetLongview() []GetUserGrants200ResponseDatabaseInner`

GetLongview returns the Longview field if non-nil, zero value otherwise.

### GetLongviewOk

`func (o *GetUserGrants200Response) GetLongviewOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetLongviewOk returns a tuple with the Longview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongview

`func (o *GetUserGrants200Response) SetLongview(v []GetUserGrants200ResponseDatabaseInner)`

SetLongview sets Longview field to given value.

### HasLongview

`func (o *GetUserGrants200Response) HasLongview() bool`

HasLongview returns a boolean if a field has been set.

### GetNodebalancer

`func (o *GetUserGrants200Response) GetNodebalancer() []GetUserGrants200ResponseDatabaseInner`

GetNodebalancer returns the Nodebalancer field if non-nil, zero value otherwise.

### GetNodebalancerOk

`func (o *GetUserGrants200Response) GetNodebalancerOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetNodebalancerOk returns a tuple with the Nodebalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancer

`func (o *GetUserGrants200Response) SetNodebalancer(v []GetUserGrants200ResponseDatabaseInner)`

SetNodebalancer sets Nodebalancer field to given value.

### HasNodebalancer

`func (o *GetUserGrants200Response) HasNodebalancer() bool`

HasNodebalancer returns a boolean if a field has been set.

### GetPlacementGroup

`func (o *GetUserGrants200Response) GetPlacementGroup() []GetUserGrants200ResponseDatabaseInner`

GetPlacementGroup returns the PlacementGroup field if non-nil, zero value otherwise.

### GetPlacementGroupOk

`func (o *GetUserGrants200Response) GetPlacementGroupOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetPlacementGroupOk returns a tuple with the PlacementGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroup

`func (o *GetUserGrants200Response) SetPlacementGroup(v []GetUserGrants200ResponseDatabaseInner)`

SetPlacementGroup sets PlacementGroup field to given value.

### HasPlacementGroup

`func (o *GetUserGrants200Response) HasPlacementGroup() bool`

HasPlacementGroup returns a boolean if a field has been set.

### GetStackscript

`func (o *GetUserGrants200Response) GetStackscript() []GetUserGrants200ResponseDatabaseInner`

GetStackscript returns the Stackscript field if non-nil, zero value otherwise.

### GetStackscriptOk

`func (o *GetUserGrants200Response) GetStackscriptOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetStackscriptOk returns a tuple with the Stackscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscript

`func (o *GetUserGrants200Response) SetStackscript(v []GetUserGrants200ResponseDatabaseInner)`

SetStackscript sets Stackscript field to given value.

### HasStackscript

`func (o *GetUserGrants200Response) HasStackscript() bool`

HasStackscript returns a boolean if a field has been set.

### GetVolume

`func (o *GetUserGrants200Response) GetVolume() []GetUserGrants200ResponseDatabaseInner`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GetUserGrants200Response) GetVolumeOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GetUserGrants200Response) SetVolume(v []GetUserGrants200ResponseDatabaseInner)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GetUserGrants200Response) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetVpc

`func (o *GetUserGrants200Response) GetVpc() []GetUserGrants200ResponseDatabaseInner`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *GetUserGrants200Response) GetVpcOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *GetUserGrants200Response) SetVpc(v []GetUserGrants200ResponseDatabaseInner)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *GetUserGrants200Response) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


