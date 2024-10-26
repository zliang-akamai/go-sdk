# PostDatabasesMysqlInstancesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowList** | Pointer to **[]string** | A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.  By default, this is an empty array (&#x60;[]&#x60;), which blocks all connections (both public and private) to the Managed Database.  If &#x60;0.0.0.0/0&#x60; is a value in this list, then all IP addresses can access the Managed Database. | [optional] 
**ClusterSize** | Pointer to **int32** | The number of Linode Instance nodes deployed to the Managed Database.  Choosing 3 nodes creates a high availability cluster consisting of 1 primary node and 2 replica nodes. | [optional] [default to 1]
**Encrypted** | Pointer to **bool** | Whether the Managed Databases is encrypted. | [optional] [default to false]
**Engine** | **string** | The Managed Database engine in engine/version format. | 
**Label** | **string** | A unique, user-defined string referring to the Managed Database. | 
**Region** | **string** | The [Region](https://techdocs.akamai.com/linode-api/reference/get-regions) ID for the Managed Database. | 
**ReplicationType** | Pointer to **string** | The replication method used for the Managed Database.  Defaults to &#x60;none&#x60; for a single cluster and &#x60;semi_synch&#x60; for a high availability cluster.  Must be &#x60;none&#x60; for a single node cluster.  Must be &#x60;asynch&#x60; or &#x60;semi_synch&#x60; for a high availability cluster. | [optional] 
**SslConnection** | Pointer to **bool** | Whether to require SSL credentials to establish a connection to the Managed Database.  Run the [Get managed MySQL database credentials](https://techdocs.akamai.com/linode-api/reference/get-databases-mysql-instance-credentials) operation for access information. | [optional] [default to true]
**Type** | **string** | The Linode Instance type used by the Managed Database for its nodes. | 

## Methods

### NewPostDatabasesMysqlInstancesRequest

`func NewPostDatabasesMysqlInstancesRequest(engine string, label string, region string, type_ string, ) *PostDatabasesMysqlInstancesRequest`

NewPostDatabasesMysqlInstancesRequest instantiates a new PostDatabasesMysqlInstancesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostDatabasesMysqlInstancesRequestWithDefaults

`func NewPostDatabasesMysqlInstancesRequestWithDefaults() *PostDatabasesMysqlInstancesRequest`

NewPostDatabasesMysqlInstancesRequestWithDefaults instantiates a new PostDatabasesMysqlInstancesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowList

`func (o *PostDatabasesMysqlInstancesRequest) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *PostDatabasesMysqlInstancesRequest) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *PostDatabasesMysqlInstancesRequest) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *PostDatabasesMysqlInstancesRequest) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetClusterSize

`func (o *PostDatabasesMysqlInstancesRequest) GetClusterSize() int32`

GetClusterSize returns the ClusterSize field if non-nil, zero value otherwise.

### GetClusterSizeOk

`func (o *PostDatabasesMysqlInstancesRequest) GetClusterSizeOk() (*int32, bool)`

GetClusterSizeOk returns a tuple with the ClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSize

`func (o *PostDatabasesMysqlInstancesRequest) SetClusterSize(v int32)`

SetClusterSize sets ClusterSize field to given value.

### HasClusterSize

`func (o *PostDatabasesMysqlInstancesRequest) HasClusterSize() bool`

HasClusterSize returns a boolean if a field has been set.

### GetEncrypted

`func (o *PostDatabasesMysqlInstancesRequest) GetEncrypted() bool`

GetEncrypted returns the Encrypted field if non-nil, zero value otherwise.

### GetEncryptedOk

`func (o *PostDatabasesMysqlInstancesRequest) GetEncryptedOk() (*bool, bool)`

GetEncryptedOk returns a tuple with the Encrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncrypted

`func (o *PostDatabasesMysqlInstancesRequest) SetEncrypted(v bool)`

SetEncrypted sets Encrypted field to given value.

### HasEncrypted

`func (o *PostDatabasesMysqlInstancesRequest) HasEncrypted() bool`

HasEncrypted returns a boolean if a field has been set.

### GetEngine

`func (o *PostDatabasesMysqlInstancesRequest) GetEngine() string`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *PostDatabasesMysqlInstancesRequest) GetEngineOk() (*string, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *PostDatabasesMysqlInstancesRequest) SetEngine(v string)`

SetEngine sets Engine field to given value.


### GetLabel

`func (o *PostDatabasesMysqlInstancesRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostDatabasesMysqlInstancesRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostDatabasesMysqlInstancesRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRegion

`func (o *PostDatabasesMysqlInstancesRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostDatabasesMysqlInstancesRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostDatabasesMysqlInstancesRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetReplicationType

`func (o *PostDatabasesMysqlInstancesRequest) GetReplicationType() string`

GetReplicationType returns the ReplicationType field if non-nil, zero value otherwise.

### GetReplicationTypeOk

`func (o *PostDatabasesMysqlInstancesRequest) GetReplicationTypeOk() (*string, bool)`

GetReplicationTypeOk returns a tuple with the ReplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationType

`func (o *PostDatabasesMysqlInstancesRequest) SetReplicationType(v string)`

SetReplicationType sets ReplicationType field to given value.

### HasReplicationType

`func (o *PostDatabasesMysqlInstancesRequest) HasReplicationType() bool`

HasReplicationType returns a boolean if a field has been set.

### GetSslConnection

`func (o *PostDatabasesMysqlInstancesRequest) GetSslConnection() bool`

GetSslConnection returns the SslConnection field if non-nil, zero value otherwise.

### GetSslConnectionOk

`func (o *PostDatabasesMysqlInstancesRequest) GetSslConnectionOk() (*bool, bool)`

GetSslConnectionOk returns a tuple with the SslConnection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslConnection

`func (o *PostDatabasesMysqlInstancesRequest) SetSslConnection(v bool)`

SetSslConnection sets SslConnection field to given value.

### HasSslConnection

`func (o *PostDatabasesMysqlInstancesRequest) HasSslConnection() bool`

HasSslConnection returns a boolean if a field has been set.

### GetType

`func (o *PostDatabasesMysqlInstancesRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostDatabasesMysqlInstancesRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostDatabasesMysqlInstancesRequest) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


