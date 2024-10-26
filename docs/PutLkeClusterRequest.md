# PutLkeClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlPlane** | Pointer to [**PutLkeClusterRequestControlPlane**](PutLkeClusterRequestControlPlane.md) |  | [optional] 
**K8sVersion** | Pointer to **string** | The desired Kubernetes version for this Kubernetes cluster in the format of &amp;lt;major&amp;gt;.&amp;lt;minor&amp;gt;. New and recycled Nodes in this cluster will be installed with the latest available patch for the Cluster&#39;s Kubernetes version.  When upgrading the Kubernetes version, only the next latest minor version following the current version can be deployed. For example, a cluster with Kubernetes version 1.19 can be upgraded to version 1.20, but not directly to 1.21.  The Kubernetes version of a cluster can not be downgraded. | [optional] 
**Label** | Pointer to **string** | This Kubernetes cluster&#39;s unique label for display purposes only. Labels have the following constraints:    - UTF-8 characters will be returned by the API using escape sequences of their Unicode code points. For example, the Japanese character _か_ is 3 bytes in UTF-8 (&#x60;0xE382AB&#x60;). Its Unicode code point is 2 bytes (&#x60;0x30AB&#x60;). APIv4 supports this character and the API will return it as the escape sequence using six 1 byte characters which represent 2 bytes of Unicode code point (&#x60;\&quot;\\u30ab\&quot;&#x60;).    - 4 byte UTF-8 characters are not supported.    - If the label is entirely composed of UTF-8 characters, the API response will return the code points using up to 193 1 byte characters. | [optional] 
**Tags** | Pointer to **[]string** | An array of tags applied to the Kubernetes cluster. Tags are for organizational purposes only. To delete a tag, exclude it from your &#x60;tags&#x60; array. | [optional] 

## Methods

### NewPutLkeClusterRequest

`func NewPutLkeClusterRequest() *PutLkeClusterRequest`

NewPutLkeClusterRequest instantiates a new PutLkeClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLkeClusterRequestWithDefaults

`func NewPutLkeClusterRequestWithDefaults() *PutLkeClusterRequest`

NewPutLkeClusterRequestWithDefaults instantiates a new PutLkeClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlPlane

`func (o *PutLkeClusterRequest) GetControlPlane() PutLkeClusterRequestControlPlane`

GetControlPlane returns the ControlPlane field if non-nil, zero value otherwise.

### GetControlPlaneOk

`func (o *PutLkeClusterRequest) GetControlPlaneOk() (*PutLkeClusterRequestControlPlane, bool)`

GetControlPlaneOk returns a tuple with the ControlPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlane

`func (o *PutLkeClusterRequest) SetControlPlane(v PutLkeClusterRequestControlPlane)`

SetControlPlane sets ControlPlane field to given value.

### HasControlPlane

`func (o *PutLkeClusterRequest) HasControlPlane() bool`

HasControlPlane returns a boolean if a field has been set.

### GetK8sVersion

`func (o *PutLkeClusterRequest) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *PutLkeClusterRequest) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *PutLkeClusterRequest) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *PutLkeClusterRequest) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetLabel

`func (o *PutLkeClusterRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutLkeClusterRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutLkeClusterRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutLkeClusterRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetTags

`func (o *PutLkeClusterRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PutLkeClusterRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PutLkeClusterRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PutLkeClusterRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


