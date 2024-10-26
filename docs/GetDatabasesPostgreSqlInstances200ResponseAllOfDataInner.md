# GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowList** | Pointer to **[]string** | A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.  By default, this is an empty array (&#x60;[]&#x60;), which blocks all connections (both public and private) to the Managed Database.  If &#x60;0.0.0.0/0&#x60; is a value in this list, then all IP addresses can access the Managed Database. | [optional] 
**ClusterSize** | Pointer to **int32** | The number of Linode Instance nodes deployed to the Managed Database.  Choosing 3 nodes creates a high availability cluster consisting of 1 primary node and 2 replica nodes. | [optional] [default to 1]
**Created** | Pointer to **time.Time** | When this Managed Database was created. | [optional] [readonly] 
**Encrypted** | Pointer to **bool** | Whether the Managed Databases is encrypted. | [optional] [default to false]
**Engine** | Pointer to **string** | The Managed Database engine type. | [optional] [readonly] 
**Hosts** | Pointer to [**GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts**](GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts.md) |  | [optional] 
**Id** | Pointer to **int32** | A unique ID that can be used to identify and reference the Managed Database. | [optional] [readonly] 
**Label** | Pointer to **string** | A unique, user-defined string referring to the Managed Database. | [optional] 
**Port** | Pointer to **int32** | The access port for this Managed Database. | [optional] 
**Region** | Pointer to **string** | The [Region](https://techdocs.akamai.com/linode-api/reference/get-regions) ID for the Managed Database. | [optional] 
**ReplicationCommitType** | Pointer to **string** | The synchronization level of the replicating server.  Must be &#x60;local&#x60; or &#x60;off&#x60; for the &#x60;asynch&#x60; replication type. | [optional] [default to "local"]
**ReplicationType** | Pointer to **string** | The replication method used for the Managed Database.  Defaults to &#x60;none&#x60; for a single cluster and &#x60;asynch&#x60; for a high availability cluster.  Must be &#x60;none&#x60; for a single node cluster.  Must be &#x60;asynch&#x60; for a high availability cluster. | [optional] 
**SslConnection** | Pointer to **bool** | Whether to require SSL credentials to establish a connection to the Managed Database.  Run the [Get managed PostgreSQL database credentials](https://techdocs.akamai.com/linode-api/reference/get-databases-postgre-sql-instance-credentials) operation for access information. | [optional] [default to true]
**Status** | Pointer to **string** | The operating status of the Managed Database. | [optional] [readonly] 
**TotalDiskSizeGb** | Pointer to **int32** | The total disk size of the database in GB. | [optional] 
**Type** | Pointer to **string** | The Linode Instance type used by the Managed Database for its nodes. | [optional] 
**Updated** | Pointer to **time.Time** | When this Managed Database was last updated. | [optional] [readonly] 
**Updates** | Pointer to [**GetDatabasesInstances200ResponseAllOfDataInnerUpdates**](GetDatabasesInstances200ResponseAllOfDataInnerUpdates.md) |  | [optional] 
**UsedDiskSizeGb** | Pointer to **int32** | The used space of the database in GB. | [optional] 
**Version** | Pointer to **string** | The Managed Database engine version. | [optional] [readonly] 

## Methods

### NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInner

`func NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInner() *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner`

NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInner instantiates a new GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerWithDefaults

`func NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerWithDefaults() *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner`

NewGetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerWithDefaults instantiates a new GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowList

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetClusterSize

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetClusterSize() int32`

GetClusterSize returns the ClusterSize field if non-nil, zero value otherwise.

### GetClusterSizeOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetClusterSizeOk() (*int32, bool)`

GetClusterSizeOk returns a tuple with the ClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSize

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetClusterSize(v int32)`

SetClusterSize sets ClusterSize field to given value.

### HasClusterSize

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasClusterSize() bool`

HasClusterSize returns a boolean if a field has been set.

### GetCreated

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEncrypted

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetEngine

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetEngine(v string)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetHosts

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetHosts() GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetHostsOk() (*GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetHosts(v GetDatabasesPostgreSqlInstances200ResponseAllOfDataInnerHosts)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetId

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPort

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRegion

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetReplicationCommitType

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetReplicationCommitType() string`

GetReplicationCommitType returns the ReplicationCommitType field if non-nil, zero value otherwise.

### GetReplicationCommitTypeOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetReplicationCommitTypeOk() (*string, bool)`

GetReplicationCommitTypeOk returns a tuple with the ReplicationCommitType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationCommitType

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetReplicationCommitType(v string)`

SetReplicationCommitType sets ReplicationCommitType field to given value.

### HasReplicationCommitType

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasReplicationCommitType() bool`

HasReplicationCommitType returns a boolean if a field has been set.

### GetReplicationType

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetReplicationType() string`

GetReplicationType returns the ReplicationType field if non-nil, zero value otherwise.

### GetReplicationTypeOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetReplicationTypeOk() (*string, bool)`

GetReplicationTypeOk returns a tuple with the ReplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationType

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetReplicationType(v string)`

SetReplicationType sets ReplicationType field to given value.

### HasReplicationType

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasReplicationType() bool`

HasReplicationType returns a boolean if a field has been set.

### GetSslConnection

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetSslConnection() bool`

GetSslConnection returns the SslConnection field if non-nil, zero value otherwise.

### GetSslConnectionOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetSslConnectionOk() (*bool, bool)`

GetSslConnectionOk returns a tuple with the SslConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnection

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetSslConnection(v bool)`

SetSslConnection sets SslConnection field to given value.

### HasSslConnection

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasSslConnection() bool`

HasSslConnection returns a boolean if a field has been set.

### GetStatus

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotalDiskSizeGb

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetTotalDiskSizeGb() int32`

GetTotalDiskSizeGb returns the TotalDiskSizeGb field if non-nil, zero value otherwise.

### GetTotalDiskSizeGbOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetTotalDiskSizeGbOk() (*int32, bool)`

GetTotalDiskSizeGbOk returns a tuple with the TotalDiskSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDiskSizeGb

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetTotalDiskSizeGb(v int32)`

SetTotalDiskSizeGb sets TotalDiskSizeGb field to given value.

### HasTotalDiskSizeGb

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasTotalDiskSizeGb() bool`

HasTotalDiskSizeGb returns a boolean if a field has been set.

### GetType

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpdates

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetUpdates() GetDatabasesInstances200ResponseAllOfDataInnerUpdates`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetUpdatesOk() (*GetDatabasesInstances200ResponseAllOfDataInnerUpdates, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetUpdates(v GetDatabasesInstances200ResponseAllOfDataInnerUpdates)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetUsedDiskSizeGb

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetUsedDiskSizeGb() int32`

GetUsedDiskSizeGb returns the UsedDiskSizeGb field if non-nil, zero value otherwise.

### GetUsedDiskSizeGbOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetUsedDiskSizeGbOk() (*int32, bool)`

GetUsedDiskSizeGbOk returns a tuple with the UsedDiskSizeGb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedDiskSizeGb

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetUsedDiskSizeGb(v int32)`

SetUsedDiskSizeGb sets UsedDiskSizeGb field to given value.

### HasUsedDiskSizeGb

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasUsedDiskSizeGb() bool`

HasUsedDiskSizeGb returns a boolean if a field has been set.

### GetVersion

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetDatabasesPostgreSqlInstances200ResponseAllOfDataInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


