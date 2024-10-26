# PostLkeClusterRegenerateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kubeconfig** | Pointer to **bool** | Whether to delete and regenerate the Kubeconfig file for this Cluster. | [optional] [default to false]
**Servicetoken** | Pointer to **bool** | Whether to delete and regenerate the service access token for this Cluster. | [optional] [default to false]

## Methods

### NewPostLkeClusterRegenerateRequest

`func NewPostLkeClusterRegenerateRequest() *PostLkeClusterRegenerateRequest`

NewPostLkeClusterRegenerateRequest instantiates a new PostLkeClusterRegenerateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLkeClusterRegenerateRequestWithDefaults

`func NewPostLkeClusterRegenerateRequestWithDefaults() *PostLkeClusterRegenerateRequest`

NewPostLkeClusterRegenerateRequestWithDefaults instantiates a new PostLkeClusterRegenerateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeconfig

`func (o *PostLkeClusterRegenerateRequest) GetKubeconfig() bool`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *PostLkeClusterRegenerateRequest) GetKubeconfigOk() (*bool, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *PostLkeClusterRegenerateRequest) SetKubeconfig(v bool)`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *PostLkeClusterRegenerateRequest) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.

### GetServicetoken

`func (o *PostLkeClusterRegenerateRequest) GetServicetoken() bool`

GetServicetoken returns the Servicetoken field if non-nil, zero value otherwise.

### GetServicetokenOk

`func (o *PostLkeClusterRegenerateRequest) GetServicetokenOk() (*bool, bool)`

GetServicetokenOk returns a tuple with the Servicetoken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicetoken

`func (o *PostLkeClusterRegenerateRequest) SetServicetoken(v bool)`

SetServicetoken sets Servicetoken field to given value.

### HasServicetoken

`func (o *PostLkeClusterRegenerateRequest) HasServicetoken() bool`

HasServicetoken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


