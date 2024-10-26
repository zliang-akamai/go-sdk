# PutObjectStorageBucketAclRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | **string** | The Access Control Level of the bucket, as a canned ACL string. For more fine-grained control of ACLs, use the S3 API directly. | 
**Name** | **string** | The &#x60;name&#x60; of the object for which to update its Access Control List (ACL). Run the [List Object Storage bucket contents](https://techdocs.akamai.com/linode-api/reference/get-object-storage-bucket-content) operation to access all object names in a bucket. | 

## Methods

### NewPutObjectStorageBucketAclRequest

`func NewPutObjectStorageBucketAclRequest(acl string, name string, ) *PutObjectStorageBucketAclRequest`

NewPutObjectStorageBucketAclRequest instantiates a new PutObjectStorageBucketAclRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutObjectStorageBucketAclRequestWithDefaults

`func NewPutObjectStorageBucketAclRequestWithDefaults() *PutObjectStorageBucketAclRequest`

NewPutObjectStorageBucketAclRequestWithDefaults instantiates a new PutObjectStorageBucketAclRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *PutObjectStorageBucketAclRequest) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *PutObjectStorageBucketAclRequest) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *PutObjectStorageBucketAclRequest) SetAcl(v string)`

SetAcl sets Acl field to given value.


### GetName

`func (o *PutObjectStorageBucketAclRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PutObjectStorageBucketAclRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PutObjectStorageBucketAclRequest) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


