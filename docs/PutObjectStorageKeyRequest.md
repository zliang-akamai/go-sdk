# PutObjectStorageKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The label for the Object Storage key. Used for display purposes. Omit this to leave the &#x60;label&#x60; unchanged. | [optional] 
**Regions** | Pointer to **[]string** | Replace the current list of &#x60;regions&#x60; set in a specific key. Include an existing region to maintain it or leave it out to remove it. If you include new &#x60;regions&#x60; in the key, they can&#39;t be used to manage content in buckets in that specific region. You can grant these keys this access using [bucket policies](https://www.linode.com/docs/products/storage/object-storage/guides/bucket-policies/). Omit this to leave the existing list unchanged.  &gt; 🚧 &gt; &gt; You can&#39;t remove a &#x60;region&#x60; from a limited key&#39;s original &#x60;bucket_access&#x60; list. If you include the &#x60;regions&#x60; array in this operation, you need to include all existing &#x60;region&#x60; entries from the &#x60;bucket_access&#x60; array. Otherwise, the operation fails with an error. Run [Get an Object Storage key](https://techdocs.akamai.com/linode-api/reference/get-object-storage-key) to review current &#x60;region&#x60; entries in a limited key. | [optional] 

## Methods

### NewPutObjectStorageKeyRequest

`func NewPutObjectStorageKeyRequest() *PutObjectStorageKeyRequest`

NewPutObjectStorageKeyRequest instantiates a new PutObjectStorageKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectStorageKeyRequestWithDefaults

`func NewPutObjectStorageKeyRequestWithDefaults() *PutObjectStorageKeyRequest`

NewPutObjectStorageKeyRequestWithDefaults instantiates a new PutObjectStorageKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PutObjectStorageKeyRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutObjectStorageKeyRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutObjectStorageKeyRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutObjectStorageKeyRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegions

`func (o *PutObjectStorageKeyRequest) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PutObjectStorageKeyRequest) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PutObjectStorageKeyRequest) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PutObjectStorageKeyRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


