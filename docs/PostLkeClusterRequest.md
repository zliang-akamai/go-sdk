# PostLkeClusterRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlPlane** | Pointer to [**GetLkeClusters200ResponseDataInnerControlPlane**](GetLkeClusters200ResponseDataInnerControlPlane.md) |  | [optional] 
**K8sVersion** | **string** | The desired Kubernetes version for this Kubernetes cluster in the format of &amp;lt;major&amp;gt;.&amp;lt;minor&amp;gt;, and the latest supported patch version will be deployed. | 
**Label** | **string** | This Kubernetes cluster&#39;s unique label for display purposes only. Labels have the following constraints:    - UTF-8 characters will be returned by the API using escape sequences of their Unicode code points. For example, the Japanese character _か_ is 3 bytes in UTF-8 (&#x60;0xE382AB&#x60;). Its Unicode code point is 2 bytes (&#x60;0x30AB&#x60;). APIv4 supports this character and the API will return it as the escape sequence using six 1 byte characters which represent 2 bytes of Unicode code point (&#x60;\&quot;\\u30ab\&quot;&#x60;).    - 4 byte UTF-8 characters are not supported.    - If the label is entirely composed of UTF-8 characters, the API response will return the code points using up to 193 1 byte characters. | 
**NodePools** | [**[]PostLkeClusterRequestNodePoolsInner**](PostLkeClusterRequestNodePoolsInner.md) |  | 
**Region** | **string** | This Kubernetes cluster&#39;s location. | 
**Tags** | Pointer to **[]string** | An array of tags applied to the Kubernetes cluster. Tags are for organizational purposes only. | [optional] 

## Methods

### NewPostLkeClusterRequest

`func NewPostLkeClusterRequest(k8sVersion string, label string, nodePools []PostLkeClusterRequestNodePoolsInner, region string, ) *PostLkeClusterRequest`

NewPostLkeClusterRequest instantiates a new PostLkeClusterRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLkeClusterRequestWithDefaults

`func NewPostLkeClusterRequestWithDefaults() *PostLkeClusterRequest`

NewPostLkeClusterRequestWithDefaults instantiates a new PostLkeClusterRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlPlane

`func (o *PostLkeClusterRequest) GetControlPlane() GetLkeClusters200ResponseDataInnerControlPlane`

GetControlPlane returns the ControlPlane field if non-nil, zero value otherwise.

### GetControlPlaneOk

`func (o *PostLkeClusterRequest) GetControlPlaneOk() (*GetLkeClusters200ResponseDataInnerControlPlane, bool)`

GetControlPlaneOk returns a tuple with the ControlPlane field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlPlane

`func (o *PostLkeClusterRequest) SetControlPlane(v GetLkeClusters200ResponseDataInnerControlPlane)`

SetControlPlane sets ControlPlane field to given value.

### HasControlPlane

`func (o *PostLkeClusterRequest) HasControlPlane() bool`

HasControlPlane returns a boolean if a field has been set.

### GetK8sVersion

`func (o *PostLkeClusterRequest) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *PostLkeClusterRequest) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *PostLkeClusterRequest) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.


### GetLabel

`func (o *PostLkeClusterRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostLkeClusterRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostLkeClusterRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNodePools

`func (o *PostLkeClusterRequest) GetNodePools() []PostLkeClusterRequestNodePoolsInner`

GetNodePools returns the NodePools field if non-nil, zero value otherwise.

### GetNodePoolsOk

`func (o *PostLkeClusterRequest) GetNodePoolsOk() (*[]PostLkeClusterRequestNodePoolsInner, bool)`

GetNodePoolsOk returns a tuple with the NodePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodePools

`func (o *PostLkeClusterRequest) SetNodePools(v []PostLkeClusterRequestNodePoolsInner)`

SetNodePools sets NodePools field to given value.


### GetRegion

`func (o *PostLkeClusterRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostLkeClusterRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostLkeClusterRequest) SetRegion(v string)`

SetRegion sets Region field to given value.


### GetTags

`func (o *PostLkeClusterRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostLkeClusterRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostLkeClusterRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostLkeClusterRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


