# GetDatabasesInstances200ResponseAllOfDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowList** | Pointer to **[]string** | A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.  By default, this is an empty array (&#x60;[]&#x60;), which blocks all connections (both public and private) to the Managed Database.  If &#x60;0.0.0.0/0&#x60; is a value in this list, then all IP addresses can access the Managed Database. | [optional] 
**ClusterSize** | Pointer to **int32** | The number of Linode Instance nodes deployed to the Managed Database.  Choosing 3 nodes creates a high availability cluster consisting of 1 primary node and 2 replica nodes. | [optional] [default to 1]
**Created** | Pointer to **time.Time** | When this Managed Database was created. | [optional] [readonly] 
**Encrypted** | Pointer to **bool** | Whether the Managed Databases is encrypted. | [optional] [default to false]
**Engine** | Pointer to **string** | The Managed Database engine type. | [optional] [readonly] 
**Hosts** | Pointer to [**GetDatabasesInstances200ResponseAllOfDataInnerHosts**](GetDatabasesInstances200ResponseAllOfDataInnerHosts.md) |  | [optional] 
**Id** | Pointer to **int32** | A unique ID that can be used to identify and reference the Managed Database. | [optional] [readonly] 
**InstanceUri** | Pointer to **string** | Append this to &#x60;https://api.linode.com&#x60; to run commands for the Managed Database. | [optional] [readonly] 
**Label** | Pointer to **string** | A unique, user-defined string referring to the Managed Database. | [optional] 
**Region** | Pointer to **string** | The [Region](https://techdocs.akamai.com/linode-api/reference/get-regions) ID for the Managed Database. | [optional] 
**Status** | Pointer to **string** | The operating status of the Managed Database. | [optional] [readonly] 
**Type** | Pointer to **string** | The Linode Instance type used by the Managed Database for its nodes. | [optional] 
**Updated** | Pointer to **time.Time** | When this Managed Database was last updated. | [optional] [readonly] 
**Updates** | Pointer to [**GetDatabasesInstances200ResponseAllOfDataInnerUpdates**](GetDatabasesInstances200ResponseAllOfDataInnerUpdates.md) |  | [optional] 
**Version** | Pointer to **string** | The Managed Database engine version. | [optional] [readonly] 

## Methods

### NewGetDatabasesInstances200ResponseAllOfDataInner

`func NewGetDatabasesInstances200ResponseAllOfDataInner() *GetDatabasesInstances200ResponseAllOfDataInner`

NewGetDatabasesInstances200ResponseAllOfDataInner instantiates a new GetDatabasesInstances200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasesInstances200ResponseAllOfDataInnerWithDefaults

`func NewGetDatabasesInstances200ResponseAllOfDataInnerWithDefaults() *GetDatabasesInstances200ResponseAllOfDataInner`

NewGetDatabasesInstances200ResponseAllOfDataInnerWithDefaults instantiates a new GetDatabasesInstances200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowList

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetClusterSize

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetClusterSize() int32`

GetClusterSize returns the ClusterSize field if non-nil, zero value otherwise.

### GetClusterSizeOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetClusterSizeOk() (*int32, bool)`

GetClusterSizeOk returns a tuple with the ClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSize

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetClusterSize(v int32)`

SetClusterSize sets ClusterSize field to given value.

### HasClusterSize

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasClusterSize() bool`

HasClusterSize returns a boolean if a field has been set.

### GetCreated

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEncrypted

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetEngine

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetHosts

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetHosts() GetDatabasesInstances200ResponseAllOfDataInnerHosts`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetHostsOk() (*GetDatabasesInstances200ResponseAllOfDataInnerHosts, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetHosts(v GetDatabasesInstances200ResponseAllOfDataInnerHosts)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceUri

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetInstanceUri() string`

GetInstanceUri returns the InstanceUri field if non-nil, zero value otherwise.

### GetInstanceUriOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetInstanceUriOk() (*string, bool)`

GetInstanceUriOk returns a tuple with the InstanceUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUri

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetInstanceUri(v string)`

SetInstanceUri sets InstanceUri field to given value.

### HasInstanceUri

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasInstanceUri() bool`

HasInstanceUri returns a boolean if a field has been set.

### GetLabel

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegion

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetStatus

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpdates

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetUpdates() GetDatabasesInstances200ResponseAllOfDataInnerUpdates`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetUpdatesOk() (*GetDatabasesInstances200ResponseAllOfDataInnerUpdates, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetUpdates(v GetDatabasesInstances200ResponseAllOfDataInnerUpdates)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetVersion

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetDatabasesInstances200ResponseAllOfDataInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


