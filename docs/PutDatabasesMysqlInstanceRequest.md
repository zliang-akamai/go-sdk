# PutDatabasesMysqlInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowList** | Pointer to **[]string** | A list of IP addresses that can access the Managed Database. Each item can be a single IP address or a range in CIDR format.  By default, this is an empty array (&#x60;[]&#x60;), which blocks all connections (both public and private) to the Managed Database.  If &#x60;0.0.0.0/0&#x60; is a value in this list, then all IP addresses can access the Managed Database. | [optional] 
**Label** | Pointer to **string** | A unique, user-defined string referring to the Managed Database. | [optional] 
**Type** | Pointer to **string** | Request re-sizing of your cluster to a Linode Type with more disk space. For example, you could request a Linode Type that uses a higher plan.  - Needs to be a Linode Type with more disk space than your current Linode.  - Resizing to a larger Linode Type can accrue additional cost. Review the &#x60;price&#x60; output in the [List types](https://techdocs.akamai.com/linode-api/reference/get-linode-types) operation for more information.  - You can&#39;t update the &#x60;allow_list&#x60; and set a new &#x60;type&#x60; in the same request.  - Any active updates to your cluster need to complete before you can request a resize. The reverse is also true: An active resizing needs to complete before you can perform any other update. | [optional] 
**Updates** | Pointer to [**GetDatabasesInstances200ResponseAllOfDataInnerUpdates**](GetDatabasesInstances200ResponseAllOfDataInnerUpdates.md) |  | [optional] 

## Methods

### NewPutDatabasesMysqlInstanceRequest

`func NewPutDatabasesMysqlInstanceRequest() *PutDatabasesMysqlInstanceRequest`

NewPutDatabasesMysqlInstanceRequest instantiates a new PutDatabasesMysqlInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutDatabasesMysqlInstanceRequestWithDefaults

`func NewPutDatabasesMysqlInstanceRequestWithDefaults() *PutDatabasesMysqlInstanceRequest`

NewPutDatabasesMysqlInstanceRequestWithDefaults instantiates a new PutDatabasesMysqlInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowList

`func (o *PutDatabasesMysqlInstanceRequest) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *PutDatabasesMysqlInstanceRequest) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *PutDatabasesMysqlInstanceRequest) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *PutDatabasesMysqlInstanceRequest) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetLabel

`func (o *PutDatabasesMysqlInstanceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutDatabasesMysqlInstanceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutDatabasesMysqlInstanceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutDatabasesMysqlInstanceRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PutDatabasesMysqlInstanceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PutDatabasesMysqlInstanceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PutDatabasesMysqlInstanceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PutDatabasesMysqlInstanceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdates

`func (o *PutDatabasesMysqlInstanceRequest) GetUpdates() GetDatabasesInstances200ResponseAllOfDataInnerUpdates`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *PutDatabasesMysqlInstanceRequest) GetUpdatesOk() (*GetDatabasesInstances200ResponseAllOfDataInnerUpdates, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *PutDatabasesMysqlInstanceRequest) SetUpdates(v GetDatabasesInstances200ResponseAllOfDataInnerUpdates)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *PutDatabasesMysqlInstanceRequest) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


