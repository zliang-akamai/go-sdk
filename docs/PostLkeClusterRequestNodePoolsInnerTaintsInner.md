# PostLkeClusterRequestNodePoolsInnerTaintsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | **string** | The Kubernetes taint effect. For &#x60;NoSchedule&#x60;, &#x60;PreferNoSchedule&#x60; and &#x60;NoExecute&#x60; descriptions, see [Kubernetes Taints and Tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/). | 
**Key** | **string** | The Kubernetes taint key.  - A key can contain alphanumeric characters, dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;), or dots (&#x60;.&#x60;). Start and end it with an alphanumeric character.  - If the key begins with a DNS subdomain prefix followed by a single slash, for example &#x60;example.com/my-app&#x60;, the prefix part needs to adhere to [RFC 1123](https://datatracker.ietf.org/doc/html/rfc1123) DNS subdomain restrictions and be a maximum of 253 characters. | 
**Value** | **string** | The Kubernetes taint value.  - A key can contain alphanumeric characters, dashes (&#x60;-&#x60;), underscores (&#x60;_&#x60;), or dots (&#x60;.&#x60;). Start and end it with an alphanumeric character.  - Can be specified as an empty string value with &#x60;\&quot;\&quot;&#x60;. | 

## Methods

### NewPostLkeClusterRequestNodePoolsInnerTaintsInner

`func NewPostLkeClusterRequestNodePoolsInnerTaintsInner(effect string, key string, value string, ) *PostLkeClusterRequestNodePoolsInnerTaintsInner`

NewPostLkeClusterRequestNodePoolsInnerTaintsInner instantiates a new PostLkeClusterRequestNodePoolsInnerTaintsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLkeClusterRequestNodePoolsInnerTaintsInnerWithDefaults

`func NewPostLkeClusterRequestNodePoolsInnerTaintsInnerWithDefaults() *PostLkeClusterRequestNodePoolsInnerTaintsInner`

NewPostLkeClusterRequestNodePoolsInnerTaintsInnerWithDefaults instantiates a new PostLkeClusterRequestNodePoolsInnerTaintsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *PostLkeClusterRequestNodePoolsInnerTaintsInner) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *PostLkeClusterRequestNodePoolsInnerTaintsInner) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *PostLkeClusterRequestNodePoolsInnerTaintsInner) SetEffect(v string)`

SetEffect sets Effect field to given value.


### GetKey

`func (o *PostLkeClusterRequestNodePoolsInnerTaintsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PostLkeClusterRequestNodePoolsInnerTaintsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PostLkeClusterRequestNodePoolsInnerTaintsInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *PostLkeClusterRequestNodePoolsInnerTaintsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PostLkeClusterRequestNodePoolsInnerTaintsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PostLkeClusterRequestNodePoolsInnerTaintsInner) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


