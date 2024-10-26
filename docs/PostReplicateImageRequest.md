# PostReplicateImageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regions** | Pointer to **[]string** | The unique identifier for each compute &#x60;region&#x60;. See [Regions and images](https://techdocs.akamai.com/cloud-computing/docs/images#regions-and-images) for full details on support for &#x60;regions&#x60;. | [optional] 

## Methods

### NewPostReplicateImageRequest

`func NewPostReplicateImageRequest() *PostReplicateImageRequest`

NewPostReplicateImageRequest instantiates a new PostReplicateImageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostReplicateImageRequestWithDefaults

`func NewPostReplicateImageRequestWithDefaults() *PostReplicateImageRequest`

NewPostReplicateImageRequestWithDefaults instantiates a new PostReplicateImageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegions

`func (o *PostReplicateImageRequest) GetRegions() []string`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *PostReplicateImageRequest) GetRegionsOk() (*[]string, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *PostReplicateImageRequest) SetRegions(v []string)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *PostReplicateImageRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


