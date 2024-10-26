# GrantsResponse

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

### NewGrantsResponse

`func NewGrantsResponse() *GrantsResponse`

NewGrantsResponse instantiates a new GrantsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantsResponseWithDefaults

`func NewGrantsResponseWithDefaults() *GrantsResponse`

NewGrantsResponseWithDefaults instantiates a new GrantsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *GrantsResponse) GetDatabase() []GetUserGrants200ResponseDatabaseInner`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *GrantsResponse) GetDatabaseOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *GrantsResponse) SetDatabase(v []GetUserGrants200ResponseDatabaseInner)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *GrantsResponse) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetDomain

`func (o *GrantsResponse) GetDomain() []GetUserGrants200ResponseDatabaseInner`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GrantsResponse) GetDomainOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GrantsResponse) SetDomain(v []GetUserGrants200ResponseDatabaseInner)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GrantsResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetFirewall

`func (o *GrantsResponse) GetFirewall() []GetUserGrants200ResponseDatabaseInner`

GetFirewall returns the Firewall field if non-nil, zero value otherwise.

### GetFirewallOk

`func (o *GrantsResponse) GetFirewallOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetFirewallOk returns a tuple with the Firewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewall

`func (o *GrantsResponse) SetFirewall(v []GetUserGrants200ResponseDatabaseInner)`

SetFirewall sets Firewall field to given value.

### HasFirewall

`func (o *GrantsResponse) HasFirewall() bool`

HasFirewall returns a boolean if a field has been set.

### GetGlobal

`func (o *GrantsResponse) GetGlobal() GetUserGrants200ResponseGlobal`

GetGlobal returns the Global field if non-nil, zero value otherwise.

### GetGlobalOk

`func (o *GrantsResponse) GetGlobalOk() (*GetUserGrants200ResponseGlobal, bool)`

GetGlobalOk returns a tuple with the Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobal

`func (o *GrantsResponse) SetGlobal(v GetUserGrants200ResponseGlobal)`

SetGlobal sets Global field to given value.

### HasGlobal

`func (o *GrantsResponse) HasGlobal() bool`

HasGlobal returns a boolean if a field has been set.

### GetImage

`func (o *GrantsResponse) GetImage() []GetUserGrants200ResponseDatabaseInner`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *GrantsResponse) GetImageOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *GrantsResponse) SetImage(v []GetUserGrants200ResponseDatabaseInner)`

SetImage sets Image field to given value.

### HasImage

`func (o *GrantsResponse) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLinode

`func (o *GrantsResponse) GetLinode() []GetUserGrants200ResponseDatabaseInner`

GetLinode returns the Linode field if non-nil, zero value otherwise.

### GetLinodeOk

`func (o *GrantsResponse) GetLinodeOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetLinodeOk returns a tuple with the Linode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinode

`func (o *GrantsResponse) SetLinode(v []GetUserGrants200ResponseDatabaseInner)`

SetLinode sets Linode field to given value.

### HasLinode

`func (o *GrantsResponse) HasLinode() bool`

HasLinode returns a boolean if a field has been set.

### GetLongview

`func (o *GrantsResponse) GetLongview() []GetUserGrants200ResponseDatabaseInner`

GetLongview returns the Longview field if non-nil, zero value otherwise.

### GetLongviewOk

`func (o *GrantsResponse) GetLongviewOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetLongviewOk returns a tuple with the Longview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongview

`func (o *GrantsResponse) SetLongview(v []GetUserGrants200ResponseDatabaseInner)`

SetLongview sets Longview field to given value.

### HasLongview

`func (o *GrantsResponse) HasLongview() bool`

HasLongview returns a boolean if a field has been set.

### GetNodebalancer

`func (o *GrantsResponse) GetNodebalancer() []GetUserGrants200ResponseDatabaseInner`

GetNodebalancer returns the Nodebalancer field if non-nil, zero value otherwise.

### GetNodebalancerOk

`func (o *GrantsResponse) GetNodebalancerOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetNodebalancerOk returns a tuple with the Nodebalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodebalancer

`func (o *GrantsResponse) SetNodebalancer(v []GetUserGrants200ResponseDatabaseInner)`

SetNodebalancer sets Nodebalancer field to given value.

### HasNodebalancer

`func (o *GrantsResponse) HasNodebalancer() bool`

HasNodebalancer returns a boolean if a field has been set.

### GetPlacementGroup

`func (o *GrantsResponse) GetPlacementGroup() []GetUserGrants200ResponseDatabaseInner`

GetPlacementGroup returns the PlacementGroup field if non-nil, zero value otherwise.

### GetPlacementGroupOk

`func (o *GrantsResponse) GetPlacementGroupOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetPlacementGroupOk returns a tuple with the PlacementGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroup

`func (o *GrantsResponse) SetPlacementGroup(v []GetUserGrants200ResponseDatabaseInner)`

SetPlacementGroup sets PlacementGroup field to given value.

### HasPlacementGroup

`func (o *GrantsResponse) HasPlacementGroup() bool`

HasPlacementGroup returns a boolean if a field has been set.

### GetStackscript

`func (o *GrantsResponse) GetStackscript() []GetUserGrants200ResponseDatabaseInner`

GetStackscript returns the Stackscript field if non-nil, zero value otherwise.

### GetStackscriptOk

`func (o *GrantsResponse) GetStackscriptOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetStackscriptOk returns a tuple with the Stackscript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackscript

`func (o *GrantsResponse) SetStackscript(v []GetUserGrants200ResponseDatabaseInner)`

SetStackscript sets Stackscript field to given value.

### HasStackscript

`func (o *GrantsResponse) HasStackscript() bool`

HasStackscript returns a boolean if a field has been set.

### GetVolume

`func (o *GrantsResponse) GetVolume() []GetUserGrants200ResponseDatabaseInner`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *GrantsResponse) GetVolumeOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *GrantsResponse) SetVolume(v []GetUserGrants200ResponseDatabaseInner)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *GrantsResponse) HasVolume() bool`

HasVolume returns a boolean if a field has been set.

### GetVpc

`func (o *GrantsResponse) GetVpc() []GetUserGrants200ResponseDatabaseInner`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *GrantsResponse) GetVpcOk() (*[]GetUserGrants200ResponseDatabaseInner, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *GrantsResponse) SetVpc(v []GetUserGrants200ResponseDatabaseInner)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *GrantsResponse) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


