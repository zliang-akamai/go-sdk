# GetLkeClusterKubeconfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kubeconfig** | Pointer to **string** | The Base64-encoded Kubeconfig file for this Cluster. | [optional] [readonly] 

## Methods

### NewGetLkeClusterKubeconfig200Response

`func NewGetLkeClusterKubeconfig200Response() *GetLkeClusterKubeconfig200Response`

NewGetLkeClusterKubeconfig200Response instantiates a new GetLkeClusterKubeconfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLkeClusterKubeconfig200ResponseWithDefaults

`func NewGetLkeClusterKubeconfig200ResponseWithDefaults() *GetLkeClusterKubeconfig200Response`

NewGetLkeClusterKubeconfig200ResponseWithDefaults instantiates a new GetLkeClusterKubeconfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubeconfig

`func (o *GetLkeClusterKubeconfig200Response) GetKubeconfig() string`

GetKubeconfig returns the Kubeconfig field if non-nil, zero value otherwise.

### GetKubeconfigOk

`func (o *GetLkeClusterKubeconfig200Response) GetKubeconfigOk() (*string, bool)`

GetKubeconfigOk returns a tuple with the Kubeconfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeconfig

`func (o *GetLkeClusterKubeconfig200Response) SetKubeconfig(v string)`

SetKubeconfig sets Kubeconfig field to given value.

### HasKubeconfig

`func (o *GetLkeClusterKubeconfig200Response) HasKubeconfig() bool`

HasKubeconfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


