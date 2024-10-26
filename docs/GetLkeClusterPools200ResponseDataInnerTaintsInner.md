# GetLkeClusterPools200ResponseDataInnerTaintsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Effect** | Pointer to **string** | The Kubernetes taint effect. For &#x60;NoSchedule&#x60;, &#x60;PreferNoSchedule&#x60; and &#x60;NoExecute&#x60; descriptions, see [Kubernetes Taints and Tolerations](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/). | [optional] 
**Key** | Pointer to **string** | The Kubernetes taint key. | [optional] 
**Value** | Pointer to **string** | The Kubernetes taint value. | [optional] 

## Methods

### NewGetLkeClusterPools200ResponseDataInnerTaintsInner

`func NewGetLkeClusterPools200ResponseDataInnerTaintsInner() *GetLkeClusterPools200ResponseDataInnerTaintsInner`

NewGetLkeClusterPools200ResponseDataInnerTaintsInner instantiates a new GetLkeClusterPools200ResponseDataInnerTaintsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLkeClusterPools200ResponseDataInnerTaintsInnerWithDefaults

`func NewGetLkeClusterPools200ResponseDataInnerTaintsInnerWithDefaults() *GetLkeClusterPools200ResponseDataInnerTaintsInner`

NewGetLkeClusterPools200ResponseDataInnerTaintsInnerWithDefaults instantiates a new GetLkeClusterPools200ResponseDataInnerTaintsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffect

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) GetEffect() string`

GetEffect returns the Effect field if non-nil, zero value otherwise.

### GetEffectOk

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) GetEffectOk() (*string, bool)`

GetEffectOk returns a tuple with the Effect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffect

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) SetEffect(v string)`

SetEffect sets Effect field to given value.

### HasEffect

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) HasEffect() bool`

HasEffect returns a boolean if a field has been set.

### GetKey

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GetLkeClusterPools200ResponseDataInnerTaintsInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


