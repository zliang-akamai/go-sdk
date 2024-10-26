# PostLkeCluster200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this Kubernetes cluster was created. | [optional] [readonly] 
**K8sVersion** | Pointer to **string** | The desired Kubernetes version for this Kubernetes cluster in the format of &amp;lt;major&amp;gt;.&amp;lt;minor&amp;gt;, and the latest supported patch version will be deployed. | [optional] 
**Label** | Pointer to **string** | This Kubernetes cluster&#39;s unique label for display purposes only. Labels have the following constraints:    - UTF-8 characters will be returned by the API using escape sequences of their Unicode code points. For example, the Japanese character _か_ is 3 bytes in UTF-8 (&#x60;0xE382AB&#x60;). Its Unicode code point is 2 bytes (&#x60;0x30AB&#x60;). APIv4 supports this character and the API will return it as the escape sequence using six 1 byte characters which represent 2 bytes of Unicode code point (&#x60;\&quot;\\u30ab\&quot;&#x60;).    - 4 byte UTF-8 characters are not supported.    - If the label is entirely composed of UTF-8 characters, the API response will return the code points using up to 193 1 byte characters. | [optional] 
**Region** | Pointer to **string** | This Kubernetes cluster&#39;s location. | [optional] 
**Tags** | Pointer to **[]string** | An array of tags applied to the Kubernetes cluster. Tags are for organizational purposes only. To delete a tag, exclude it from your &#x60;tags&#x60; array. | [optional] 
**Updated** | Pointer to **time.Time** | When this Kubernetes cluster was updated. | [optional] [readonly] 

## Methods

### NewPostLkeCluster200Response

`func NewPostLkeCluster200Response() *PostLkeCluster200Response`

NewPostLkeCluster200Response instantiates a new PostLkeCluster200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLkeCluster200ResponseWithDefaults

`func NewPostLkeCluster200ResponseWithDefaults() *PostLkeCluster200Response`

NewPostLkeCluster200ResponseWithDefaults instantiates a new PostLkeCluster200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PostLkeCluster200Response) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PostLkeCluster200Response) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PostLkeCluster200Response) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PostLkeCluster200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetK8sVersion

`func (o *PostLkeCluster200Response) GetK8sVersion() string`

GetK8sVersion returns the K8sVersion field if non-nil, zero value otherwise.

### GetK8sVersionOk

`func (o *PostLkeCluster200Response) GetK8sVersionOk() (*string, bool)`

GetK8sVersionOk returns a tuple with the K8sVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sVersion

`func (o *PostLkeCluster200Response) SetK8sVersion(v string)`

SetK8sVersion sets K8sVersion field to given value.

### HasK8sVersion

`func (o *PostLkeCluster200Response) HasK8sVersion() bool`

HasK8sVersion returns a boolean if a field has been set.

### GetLabel

`func (o *PostLkeCluster200Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostLkeCluster200Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostLkeCluster200Response) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostLkeCluster200Response) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegion

`func (o *PostLkeCluster200Response) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostLkeCluster200Response) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostLkeCluster200Response) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostLkeCluster200Response) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetTags

`func (o *PostLkeCluster200Response) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostLkeCluster200Response) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostLkeCluster200Response) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostLkeCluster200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdated

`func (o *PostLkeCluster200Response) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PostLkeCluster200Response) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PostLkeCluster200Response) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PostLkeCluster200Response) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


