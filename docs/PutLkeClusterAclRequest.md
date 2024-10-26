# PutLkeClusterAclRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to [**GetLkeClusters200ResponseDataInnerControlPlaneAcl**](GetLkeClusters200ResponseDataInnerControlPlaneAcl.md) |  | [optional] 

## Methods

### NewPutLkeClusterAclRequest

`func NewPutLkeClusterAclRequest() *PutLkeClusterAclRequest`

NewPutLkeClusterAclRequest instantiates a new PutLkeClusterAclRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLkeClusterAclRequestWithDefaults

`func NewPutLkeClusterAclRequestWithDefaults() *PutLkeClusterAclRequest`

NewPutLkeClusterAclRequestWithDefaults instantiates a new PutLkeClusterAclRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *PutLkeClusterAclRequest) GetAcl() GetLkeClusters200ResponseDataInnerControlPlaneAcl`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *PutLkeClusterAclRequest) GetAclOk() (*GetLkeClusters200ResponseDataInnerControlPlaneAcl, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *PutLkeClusterAclRequest) SetAcl(v GetLkeClusters200ResponseDataInnerControlPlaneAcl)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *PutLkeClusterAclRequest) HasAcl() bool`

HasAcl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


