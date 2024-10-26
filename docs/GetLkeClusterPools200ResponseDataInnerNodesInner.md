# GetLkeClusterPools200ResponseDataInnerNodesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Node&#39;s ID. | [optional] 
**InstanceId** | Pointer to **string** | The Linode&#39;s ID. When no Linode is currently provisioned for this Node, this will be null. | [optional] 
**Status** | Pointer to **string** | The Node&#39;s status as it pertains to being a Kubernetes node. | [optional] 

## Methods

### NewGetLkeClusterPools200ResponseDataInnerNodesInner

`func NewGetLkeClusterPools200ResponseDataInnerNodesInner() *GetLkeClusterPools200ResponseDataInnerNodesInner`

NewGetLkeClusterPools200ResponseDataInnerNodesInner instantiates a new GetLkeClusterPools200ResponseDataInnerNodesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLkeClusterPools200ResponseDataInnerNodesInnerWithDefaults

`func NewGetLkeClusterPools200ResponseDataInnerNodesInnerWithDefaults() *GetLkeClusterPools200ResponseDataInnerNodesInner`

NewGetLkeClusterPools200ResponseDataInnerNodesInnerWithDefaults instantiates a new GetLkeClusterPools200ResponseDataInnerNodesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceId

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetStatus

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLkeClusterPools200ResponseDataInnerNodesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


