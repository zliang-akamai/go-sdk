# GetUserGrants200ResponseGlobal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountAccess** | Pointer to **NullableString** | The level of access this User has to Account-level actions, like billing information. A restricted User will never be able to manage users.  __Parent and child accounts__  In a [parent and child account](https://www.linode.com/docs/guides/parent-child-accounts/) environment, this grant can be added to a child account user, to give the user &#x60;read-write&#x60; access. This gives the child user unrestricted access to expected management operations, such as creating other child users. However, child users don&#39;t have write access to billing operations. The API issues a specific error message if a write operation is attempted by a child user. | [optional] 
**AddDatabases** | Pointer to **bool** | If true, this User may add Managed Databases. | [optional] 
**AddDomains** | Pointer to **bool** | If true, this User may add Domains. | [optional] 
**AddFirewalls** | Pointer to **bool** | If true, this User may add Firewalls. | [optional] 
**AddImages** | Pointer to **bool** | If true, this User may add Images. | [optional] 
**AddLinodes** | Pointer to **bool** | If true, this User may create Linodes. | [optional] 
**AddLoadbalancers** | Pointer to **bool** | If true, this User may add Cloud Load Balancers. | [optional] 
**AddLongview** | Pointer to **bool** | If true, this User may create Longview clients and view the current plan. | [optional] 
**AddNodebalancers** | Pointer to **bool** | If true, this User may add NodeBalancers. | [optional] 
**AddPlacementGroups** | Pointer to **bool** | If true, this User may add Placement Groups. | [optional] 
**AddStackscripts** | Pointer to **bool** | If true, this User may add StackScripts. | [optional] 
**AddVolumes** | Pointer to **bool** | If true, this User may add Volumes. | [optional] 
**AddVpcs** | Pointer to **bool** | If true, this User may add VPCs. | [optional] 
**CancelAccount** | Pointer to **bool** | If true, this User may cancel the entire Account. | [optional] 
**ChildAccountAccess** | Pointer to **NullableBool** | In a [parent and child account](https://www.linode.com/docs/guides/parent-child-accounts/) environment, this gives a parent account access to endpoints that can be used to manage child accounts. Unrestricted parent account users have access to this grant, while restricted parent users don&#39;t. An unrestricted parent user can set this to &#x60;true&#x60; to add this grant to a restricted parent user. Displayed as &#x60;null&#x60; for all non-parent accounts. | [optional] 
**LongviewSubscription** | Pointer to **bool** | If true, this User may manage the Account&#39;s Longview subscription. | [optional] 

## Methods

### NewGetUserGrants200ResponseGlobal

`func NewGetUserGrants200ResponseGlobal() *GetUserGrants200ResponseGlobal`

NewGetUserGrants200ResponseGlobal instantiates a new GetUserGrants200ResponseGlobal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserGrants200ResponseGlobalWithDefaults

`func NewGetUserGrants200ResponseGlobalWithDefaults() *GetUserGrants200ResponseGlobal`

NewGetUserGrants200ResponseGlobalWithDefaults instantiates a new GetUserGrants200ResponseGlobal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountAccess

`func (o *GetUserGrants200ResponseGlobal) GetAccountAccess() string`

GetAccountAccess returns the AccountAccess field if non-nil, zero value otherwise.

### GetAccountAccessOk

`func (o *GetUserGrants200ResponseGlobal) GetAccountAccessOk() (*string, bool)`

GetAccountAccessOk returns a tuple with the AccountAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountAccess

`func (o *GetUserGrants200ResponseGlobal) SetAccountAccess(v string)`

SetAccountAccess sets AccountAccess field to given value.

### HasAccountAccess

`func (o *GetUserGrants200ResponseGlobal) HasAccountAccess() bool`

HasAccountAccess returns a boolean if a field has been set.

### SetAccountAccessNil

`func (o *GetUserGrants200ResponseGlobal) SetAccountAccessNil(b bool)`

 SetAccountAccessNil sets the value for AccountAccess to be an explicit nil

### UnsetAccountAccess
`func (o *GetUserGrants200ResponseGlobal) UnsetAccountAccess()`

UnsetAccountAccess ensures that no value is present for AccountAccess, not even an explicit nil
### GetAddDatabases

`func (o *GetUserGrants200ResponseGlobal) GetAddDatabases() bool`

GetAddDatabases returns the AddDatabases field if non-nil, zero value otherwise.

### GetAddDatabasesOk

`func (o *GetUserGrants200ResponseGlobal) GetAddDatabasesOk() (*bool, bool)`

GetAddDatabasesOk returns a tuple with the AddDatabases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDatabases

`func (o *GetUserGrants200ResponseGlobal) SetAddDatabases(v bool)`

SetAddDatabases sets AddDatabases field to given value.

### HasAddDatabases

`func (o *GetUserGrants200ResponseGlobal) HasAddDatabases() bool`

HasAddDatabases returns a boolean if a field has been set.

### GetAddDomains

`func (o *GetUserGrants200ResponseGlobal) GetAddDomains() bool`

GetAddDomains returns the AddDomains field if non-nil, zero value otherwise.

### GetAddDomainsOk

`func (o *GetUserGrants200ResponseGlobal) GetAddDomainsOk() (*bool, bool)`

GetAddDomainsOk returns a tuple with the AddDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddDomains

`func (o *GetUserGrants200ResponseGlobal) SetAddDomains(v bool)`

SetAddDomains sets AddDomains field to given value.

### HasAddDomains

`func (o *GetUserGrants200ResponseGlobal) HasAddDomains() bool`

HasAddDomains returns a boolean if a field has been set.

### GetAddFirewalls

`func (o *GetUserGrants200ResponseGlobal) GetAddFirewalls() bool`

GetAddFirewalls returns the AddFirewalls field if non-nil, zero value otherwise.

### GetAddFirewallsOk

`func (o *GetUserGrants200ResponseGlobal) GetAddFirewallsOk() (*bool, bool)`

GetAddFirewallsOk returns a tuple with the AddFirewalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddFirewalls

`func (o *GetUserGrants200ResponseGlobal) SetAddFirewalls(v bool)`

SetAddFirewalls sets AddFirewalls field to given value.

### HasAddFirewalls

`func (o *GetUserGrants200ResponseGlobal) HasAddFirewalls() bool`

HasAddFirewalls returns a boolean if a field has been set.

### GetAddImages

`func (o *GetUserGrants200ResponseGlobal) GetAddImages() bool`

GetAddImages returns the AddImages field if non-nil, zero value otherwise.

### GetAddImagesOk

`func (o *GetUserGrants200ResponseGlobal) GetAddImagesOk() (*bool, bool)`

GetAddImagesOk returns a tuple with the AddImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddImages

`func (o *GetUserGrants200ResponseGlobal) SetAddImages(v bool)`

SetAddImages sets AddImages field to given value.

### HasAddImages

`func (o *GetUserGrants200ResponseGlobal) HasAddImages() bool`

HasAddImages returns a boolean if a field has been set.

### GetAddLinodes

`func (o *GetUserGrants200ResponseGlobal) GetAddLinodes() bool`

GetAddLinodes returns the AddLinodes field if non-nil, zero value otherwise.

### GetAddLinodesOk

`func (o *GetUserGrants200ResponseGlobal) GetAddLinodesOk() (*bool, bool)`

GetAddLinodesOk returns a tuple with the AddLinodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLinodes

`func (o *GetUserGrants200ResponseGlobal) SetAddLinodes(v bool)`

SetAddLinodes sets AddLinodes field to given value.

### HasAddLinodes

`func (o *GetUserGrants200ResponseGlobal) HasAddLinodes() bool`

HasAddLinodes returns a boolean if a field has been set.

### GetAddLoadbalancers

`func (o *GetUserGrants200ResponseGlobal) GetAddLoadbalancers() bool`

GetAddLoadbalancers returns the AddLoadbalancers field if non-nil, zero value otherwise.

### GetAddLoadbalancersOk

`func (o *GetUserGrants200ResponseGlobal) GetAddLoadbalancersOk() (*bool, bool)`

GetAddLoadbalancersOk returns a tuple with the AddLoadbalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLoadbalancers

`func (o *GetUserGrants200ResponseGlobal) SetAddLoadbalancers(v bool)`

SetAddLoadbalancers sets AddLoadbalancers field to given value.

### HasAddLoadbalancers

`func (o *GetUserGrants200ResponseGlobal) HasAddLoadbalancers() bool`

HasAddLoadbalancers returns a boolean if a field has been set.

### GetAddLongview

`func (o *GetUserGrants200ResponseGlobal) GetAddLongview() bool`

GetAddLongview returns the AddLongview field if non-nil, zero value otherwise.

### GetAddLongviewOk

`func (o *GetUserGrants200ResponseGlobal) GetAddLongviewOk() (*bool, bool)`

GetAddLongviewOk returns a tuple with the AddLongview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddLongview

`func (o *GetUserGrants200ResponseGlobal) SetAddLongview(v bool)`

SetAddLongview sets AddLongview field to given value.

### HasAddLongview

`func (o *GetUserGrants200ResponseGlobal) HasAddLongview() bool`

HasAddLongview returns a boolean if a field has been set.

### GetAddNodebalancers

`func (o *GetUserGrants200ResponseGlobal) GetAddNodebalancers() bool`

GetAddNodebalancers returns the AddNodebalancers field if non-nil, zero value otherwise.

### GetAddNodebalancersOk

`func (o *GetUserGrants200ResponseGlobal) GetAddNodebalancersOk() (*bool, bool)`

GetAddNodebalancersOk returns a tuple with the AddNodebalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddNodebalancers

`func (o *GetUserGrants200ResponseGlobal) SetAddNodebalancers(v bool)`

SetAddNodebalancers sets AddNodebalancers field to given value.

### HasAddNodebalancers

`func (o *GetUserGrants200ResponseGlobal) HasAddNodebalancers() bool`

HasAddNodebalancers returns a boolean if a field has been set.

### GetAddPlacementGroups

`func (o *GetUserGrants200ResponseGlobal) GetAddPlacementGroups() bool`

GetAddPlacementGroups returns the AddPlacementGroups field if non-nil, zero value otherwise.

### GetAddPlacementGroupsOk

`func (o *GetUserGrants200ResponseGlobal) GetAddPlacementGroupsOk() (*bool, bool)`

GetAddPlacementGroupsOk returns a tuple with the AddPlacementGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddPlacementGroups

`func (o *GetUserGrants200ResponseGlobal) SetAddPlacementGroups(v bool)`

SetAddPlacementGroups sets AddPlacementGroups field to given value.

### HasAddPlacementGroups

`func (o *GetUserGrants200ResponseGlobal) HasAddPlacementGroups() bool`

HasAddPlacementGroups returns a boolean if a field has been set.

### GetAddStackscripts

`func (o *GetUserGrants200ResponseGlobal) GetAddStackscripts() bool`

GetAddStackscripts returns the AddStackscripts field if non-nil, zero value otherwise.

### GetAddStackscriptsOk

`func (o *GetUserGrants200ResponseGlobal) GetAddStackscriptsOk() (*bool, bool)`

GetAddStackscriptsOk returns a tuple with the AddStackscripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddStackscripts

`func (o *GetUserGrants200ResponseGlobal) SetAddStackscripts(v bool)`

SetAddStackscripts sets AddStackscripts field to given value.

### HasAddStackscripts

`func (o *GetUserGrants200ResponseGlobal) HasAddStackscripts() bool`

HasAddStackscripts returns a boolean if a field has been set.

### GetAddVolumes

`func (o *GetUserGrants200ResponseGlobal) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *GetUserGrants200ResponseGlobal) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *GetUserGrants200ResponseGlobal) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *GetUserGrants200ResponseGlobal) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetAddVpcs

`func (o *GetUserGrants200ResponseGlobal) GetAddVpcs() bool`

GetAddVpcs returns the AddVpcs field if non-nil, zero value otherwise.

### GetAddVpcsOk

`func (o *GetUserGrants200ResponseGlobal) GetAddVpcsOk() (*bool, bool)`

GetAddVpcsOk returns a tuple with the AddVpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVpcs

`func (o *GetUserGrants200ResponseGlobal) SetAddVpcs(v bool)`

SetAddVpcs sets AddVpcs field to given value.

### HasAddVpcs

`func (o *GetUserGrants200ResponseGlobal) HasAddVpcs() bool`

HasAddVpcs returns a boolean if a field has been set.

### GetCancelAccount

`func (o *GetUserGrants200ResponseGlobal) GetCancelAccount() bool`

GetCancelAccount returns the CancelAccount field if non-nil, zero value otherwise.

### GetCancelAccountOk

`func (o *GetUserGrants200ResponseGlobal) GetCancelAccountOk() (*bool, bool)`

GetCancelAccountOk returns a tuple with the CancelAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelAccount

`func (o *GetUserGrants200ResponseGlobal) SetCancelAccount(v bool)`

SetCancelAccount sets CancelAccount field to given value.

### HasCancelAccount

`func (o *GetUserGrants200ResponseGlobal) HasCancelAccount() bool`

HasCancelAccount returns a boolean if a field has been set.

### GetChildAccountAccess

`func (o *GetUserGrants200ResponseGlobal) GetChildAccountAccess() bool`

GetChildAccountAccess returns the ChildAccountAccess field if non-nil, zero value otherwise.

### GetChildAccountAccessOk

`func (o *GetUserGrants200ResponseGlobal) GetChildAccountAccessOk() (*bool, bool)`

GetChildAccountAccessOk returns a tuple with the ChildAccountAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildAccountAccess

`func (o *GetUserGrants200ResponseGlobal) SetChildAccountAccess(v bool)`

SetChildAccountAccess sets ChildAccountAccess field to given value.

### HasChildAccountAccess

`func (o *GetUserGrants200ResponseGlobal) HasChildAccountAccess() bool`

HasChildAccountAccess returns a boolean if a field has been set.

### SetChildAccountAccessNil

`func (o *GetUserGrants200ResponseGlobal) SetChildAccountAccessNil(b bool)`

 SetChildAccountAccessNil sets the value for ChildAccountAccess to be an explicit nil

### UnsetChildAccountAccess
`func (o *GetUserGrants200ResponseGlobal) UnsetChildAccountAccess()`

UnsetChildAccountAccess ensures that no value is present for ChildAccountAccess, not even an explicit nil
### GetLongviewSubscription

`func (o *GetUserGrants200ResponseGlobal) GetLongviewSubscription() bool`

GetLongviewSubscription returns the LongviewSubscription field if non-nil, zero value otherwise.

### GetLongviewSubscriptionOk

`func (o *GetUserGrants200ResponseGlobal) GetLongviewSubscriptionOk() (*bool, bool)`

GetLongviewSubscriptionOk returns a tuple with the LongviewSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongviewSubscription

`func (o *GetUserGrants200ResponseGlobal) SetLongviewSubscription(v bool)`

SetLongviewSubscription sets LongviewSubscription field to given value.

### HasLongviewSubscription

`func (o *GetUserGrants200ResponseGlobal) HasLongviewSubscription() bool`

HasLongviewSubscription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


