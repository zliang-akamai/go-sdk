# Domain1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AxfrIps** | Pointer to **[]string** | The list of IPs that may perform a zone transfer for this Domain. The total combined length of all data within this array cannot exceed 1000 characters.  __Note__. This is potentially dangerous, and should be set to an empty list unless you intend to use it. | [optional] 
**Description** | Pointer to **string** | A description for this Domain. This is for display purposes only. | [optional] 
**Domain** | Pointer to **string** | The domain this Domain represents. Domain labels cannot be longer than 63 characters and must conform to [RFC1035](https://tools.ietf.org/html/rfc1035). Domains must be unique on Linode&#39;s platform, including across different Linode accounts; there cannot be two Domains representing the same domain. | [optional] 
**ExpireSec** | Pointer to **int32** | The amount of time in seconds that may pass before this Domain is no longer authoritative.  - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200.  - Any other value is rounded up to the nearest valid value.  - A value of 0 is equivalent to the default value of 1209600. | [optional] [default to 0]
**Group** | Pointer to **string** | The group this Domain belongs to.  This is for display purposes only. | [optional] 
**Id** | Pointer to **int32** | This Domain&#39;s unique ID. | [optional] [readonly] 
**MasterIps** | Pointer to **[]string** | The IP addresses representing the master DNS for this Domain. At least one value is required for &#x60;type&#x60; slave Domains. The total combined length of all data within this array cannot exceed 1000 characters. | [optional] 
**RefreshSec** | Pointer to **int32** | The amount of time in seconds before this Domain should be refreshed.  - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200.  - Any other value is rounded up to the nearest valid value.  - A value of 0 is equivalent to the default value of 14400. | [optional] [default to 0]
**RetrySec** | Pointer to **int32** | The interval, in seconds, at which a failed refresh should be retried.  - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200.  - Any other value is rounded up to the nearest valid value.  - A value of 0 is equivalent to the default value of 14400. | [optional] [default to 0]
**SoaEmail** | Pointer to **string** | Start of Authority email address. This is required for &#x60;type&#x60; master Domains. | [optional] 
**Status** | Pointer to **string** | Used to control whether this Domain is currently being rendered. | [optional] [default to "active"]
**Tags** | Pointer to **[]string** | An array of tags applied to this object.  Tags are for organizational purposes only. | [optional] 
**TtlSec** | Pointer to **int32** | \&quot;Time to Live\&quot; - the amount of time in seconds that this Domain&#39;s records may be cached by resolvers or other domain servers.  - Valid values are 0, 30, 120, 300, 3600, 7200, 14400, 28800, 57600, 86400, 172800, 345600, 604800, 1209600, and 2419200. - Any other value is rounded up to the nearest valid value. - A value of 0 is equivalent to the default value of 86400. | [optional] [default to 0]
**Type** | Pointer to **string** | Whether this Domain represents the authoritative source of information for the domain it describes (&#x60;master&#x60;), or whether it is a read-only copy of a master (&#x60;slave&#x60;). | [optional] 

## Methods

### NewDomain1

`func NewDomain1() *Domain1`

NewDomain1 instantiates a new Domain1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDomain1WithDefaults

`func NewDomain1WithDefaults() *Domain1`

NewDomain1WithDefaults instantiates a new Domain1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAxfrIps

`func (o *Domain1) GetAxfrIps() []string`

GetAxfrIps returns the AxfrIps field if non-nil, zero value otherwise.

### GetAxfrIpsOk

`func (o *Domain1) GetAxfrIpsOk() (*[]string, bool)`

GetAxfrIpsOk returns a tuple with the AxfrIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxfrIps

`func (o *Domain1) SetAxfrIps(v []string)`

SetAxfrIps sets AxfrIps field to given value.

### HasAxfrIps

`func (o *Domain1) HasAxfrIps() bool`

HasAxfrIps returns a boolean if a field has been set.

### GetDescription

`func (o *Domain1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Domain1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Domain1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Domain1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *Domain1) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Domain1) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Domain1) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Domain1) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExpireSec

`func (o *Domain1) GetExpireSec() int32`

GetExpireSec returns the ExpireSec field if non-nil, zero value otherwise.

### GetExpireSecOk

`func (o *Domain1) GetExpireSecOk() (*int32, bool)`

GetExpireSecOk returns a tuple with the ExpireSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireSec

`func (o *Domain1) SetExpireSec(v int32)`

SetExpireSec sets ExpireSec field to given value.

### HasExpireSec

`func (o *Domain1) HasExpireSec() bool`

HasExpireSec returns a boolean if a field has been set.

### GetGroup

`func (o *Domain1) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Domain1) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Domain1) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Domain1) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *Domain1) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Domain1) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Domain1) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Domain1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMasterIps

`func (o *Domain1) GetMasterIps() []string`

GetMasterIps returns the MasterIps field if non-nil, zero value otherwise.

### GetMasterIpsOk

`func (o *Domain1) GetMasterIpsOk() (*[]string, bool)`

GetMasterIpsOk returns a tuple with the MasterIps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterIps

`func (o *Domain1) SetMasterIps(v []string)`

SetMasterIps sets MasterIps field to given value.

### HasMasterIps

`func (o *Domain1) HasMasterIps() bool`

HasMasterIps returns a boolean if a field has been set.

### GetRefreshSec

`func (o *Domain1) GetRefreshSec() int32`

GetRefreshSec returns the RefreshSec field if non-nil, zero value otherwise.

### GetRefreshSecOk

`func (o *Domain1) GetRefreshSecOk() (*int32, bool)`

GetRefreshSecOk returns a tuple with the RefreshSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshSec

`func (o *Domain1) SetRefreshSec(v int32)`

SetRefreshSec sets RefreshSec field to given value.

### HasRefreshSec

`func (o *Domain1) HasRefreshSec() bool`

HasRefreshSec returns a boolean if a field has been set.

### GetRetrySec

`func (o *Domain1) GetRetrySec() int32`

GetRetrySec returns the RetrySec field if non-nil, zero value otherwise.

### GetRetrySecOk

`func (o *Domain1) GetRetrySecOk() (*int32, bool)`

GetRetrySecOk returns a tuple with the RetrySec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrySec

`func (o *Domain1) SetRetrySec(v int32)`

SetRetrySec sets RetrySec field to given value.

### HasRetrySec

`func (o *Domain1) HasRetrySec() bool`

HasRetrySec returns a boolean if a field has been set.

### GetSoaEmail

`func (o *Domain1) GetSoaEmail() string`

GetSoaEmail returns the SoaEmail field if non-nil, zero value otherwise.

### GetSoaEmailOk

`func (o *Domain1) GetSoaEmailOk() (*string, bool)`

GetSoaEmailOk returns a tuple with the SoaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoaEmail

`func (o *Domain1) SetSoaEmail(v string)`

SetSoaEmail sets SoaEmail field to given value.

### HasSoaEmail

`func (o *Domain1) HasSoaEmail() bool`

HasSoaEmail returns a boolean if a field has been set.

### GetStatus

`func (o *Domain1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Domain1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Domain1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Domain1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Domain1) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Domain1) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Domain1) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Domain1) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTtlSec

`func (o *Domain1) GetTtlSec() int32`

GetTtlSec returns the TtlSec field if non-nil, zero value otherwise.

### GetTtlSecOk

`func (o *Domain1) GetTtlSecOk() (*int32, bool)`

GetTtlSecOk returns a tuple with the TtlSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlSec

`func (o *Domain1) SetTtlSec(v int32)`

SetTtlSec sets TtlSec field to given value.

### HasTtlSec

`func (o *Domain1) HasTtlSec() bool`

HasTtlSec returns a boolean if a field has been set.

### GetType

`func (o *Domain1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Domain1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Domain1) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Domain1) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


