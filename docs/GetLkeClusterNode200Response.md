# GetLkeClusterNode200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Node&#39;s ID. | [optional] [readonly] 
**InstanceId** | Pointer to **int32** | The Linode&#39;s ID. If no Linode is currently provisioned for this Node, this is &#x60;null&#x60;. | [optional] 
**Status** | Pointer to **string** | The creation status of this Node. This status is distinct from this Node&#39;s readiness as a Kubernetes Node Object as determined by the command &#x60;kubectl get nodes&#x60;.  &#x60;not_ready&#x60; indicates that the Linode is still being created.  &#x60;ready&#x60; indicates that the Linode has successfully been created and is running Kubernetes software. | [optional] 

## Methods

### NewGetLkeClusterNode200Response

`func NewGetLkeClusterNode200Response() *GetLkeClusterNode200Response`

NewGetLkeClusterNode200Response instantiates a new GetLkeClusterNode200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLkeClusterNode200ResponseWithDefaults

`func NewGetLkeClusterNode200ResponseWithDefaults() *GetLkeClusterNode200Response`

NewGetLkeClusterNode200ResponseWithDefaults instantiates a new GetLkeClusterNode200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLkeClusterNode200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLkeClusterNode200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLkeClusterNode200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetLkeClusterNode200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceId

`func (o *GetLkeClusterNode200Response) GetInstanceId() int32`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetLkeClusterNode200Response) GetInstanceIdOk() (*int32, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetLkeClusterNode200Response) SetInstanceId(v int32)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetLkeClusterNode200Response) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetStatus

`func (o *GetLkeClusterNode200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetLkeClusterNode200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetLkeClusterNode200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetLkeClusterNode200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


