# PostLinodeInstanceRequestAllOfMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserData** | Pointer to **string** | Base64-encoded [cloud-config](https://www.linode.com/docs/products/compute/compute-instances/guides/metadata-cloud-config/) data.  Cannot be modified after provisioning. To update, use either the [Clone a Linode](https://techdocs.akamai.com/linode-api/reference/post-clone-linode-instance) or [Rebuild a Linode](https://techdocs.akamai.com/linode-api/reference/post-rebuild-linode-instance) operations.  Must not be included when cloning to an existing Linode.  Unencoded data must not exceed 65535 bytes, or about 16kb encoded. | [optional] 

## Methods

### NewPostLinodeInstanceRequestAllOfMetadata

`func NewPostLinodeInstanceRequestAllOfMetadata() *PostLinodeInstanceRequestAllOfMetadata`

NewPostLinodeInstanceRequestAllOfMetadata instantiates a new PostLinodeInstanceRequestAllOfMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostLinodeInstanceRequestAllOfMetadataWithDefaults

`func NewPostLinodeInstanceRequestAllOfMetadataWithDefaults() *PostLinodeInstanceRequestAllOfMetadata`

NewPostLinodeInstanceRequestAllOfMetadataWithDefaults instantiates a new PostLinodeInstanceRequestAllOfMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserData

`func (o *PostLinodeInstanceRequestAllOfMetadata) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *PostLinodeInstanceRequestAllOfMetadata) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *PostLinodeInstanceRequestAllOfMetadata) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *PostLinodeInstanceRequestAllOfMetadata) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


