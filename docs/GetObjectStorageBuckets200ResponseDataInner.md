# GetObjectStorageBuckets200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **string** | The legacy &#x60;clusterId&#x60; equivalent for the &#x60;regionId&#x60; where this bucket lives. The API maintains this for backward compatibility.  &gt; 📘 &gt; &gt; - This value and the &#x60;regionId&#x60; are interchangeable when used in requests. Best practice is to use the &#x60;regionId&#x60;. &gt; &gt; - This value is empty for newer regions that don&#39;t have a legacy &#x60;clusterId&#x60;. | [optional] 
**Created** | Pointer to **time.Time** | When this bucket was created. | [optional] 
**Hostname** | Pointer to **string** | The hostname where this bucket can be accessed. This hostname can be accessed through a browser if the bucket is made public. | [optional] 
**Label** | Pointer to **string** | The name of this bucket. | [optional] 
**Objects** | Pointer to **int32** | The number of objects stored in this bucket. | [optional] 
**Region** | Pointer to **string** | The &#x60;id&#x60; of the [region](https://techdocs.akamai.com/linode-api/reference/get-regions) where this Object Storage bucket lives. | [optional] 
**Size** | Pointer to **int32** | The size of the bucket in bytes. | [optional] 

## Methods

### NewGetObjectStorageBuckets200ResponseDataInner

`func NewGetObjectStorageBuckets200ResponseDataInner() *GetObjectStorageBuckets200ResponseDataInner`

NewGetObjectStorageBuckets200ResponseDataInner instantiates a new GetObjectStorageBuckets200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageBuckets200ResponseDataInnerWithDefaults

`func NewGetObjectStorageBuckets200ResponseDataInnerWithDefaults() *GetObjectStorageBuckets200ResponseDataInner`

NewGetObjectStorageBuckets200ResponseDataInnerWithDefaults instantiates a new GetObjectStorageBuckets200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *GetObjectStorageBuckets200ResponseDataInner) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *GetObjectStorageBuckets200ResponseDataInner) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetCreated

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetObjectStorageBuckets200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetObjectStorageBuckets200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetHostname

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *GetObjectStorageBuckets200ResponseDataInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *GetObjectStorageBuckets200ResponseDataInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetLabel

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetObjectStorageBuckets200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetObjectStorageBuckets200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetObjects

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetObjects() int32`

GetObjects returns the Objects field if non-nil, zero value otherwise.

### GetObjectsOk

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetObjectsOk() (*int32, bool)`

GetObjectsOk returns a tuple with the Objects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjects

`func (o *GetObjectStorageBuckets200ResponseDataInner) SetObjects(v int32)`

SetObjects sets Objects field to given value.

### HasObjects

`func (o *GetObjectStorageBuckets200ResponseDataInner) HasObjects() bool`

HasObjects returns a boolean if a field has been set.

### GetRegion

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetObjectStorageBuckets200ResponseDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetObjectStorageBuckets200ResponseDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSize

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetObjectStorageBuckets200ResponseDataInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetObjectStorageBuckets200ResponseDataInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetObjectStorageBuckets200ResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


