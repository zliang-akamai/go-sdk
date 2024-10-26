# PutStorageBucketAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **string** | The Access Control Level of the bucket, as a canned ACL string. For more fine-grained control of ACLs, use the S3 API directly. | [optional] 
**CorsEnabled** | Pointer to **bool** | If true, the bucket will be created with CORS enabled for all origins. For more fine-grained controls of CORS, use the S3 API directly. | [optional] 

## Methods

### NewPutStorageBucketAccessRequest

`func NewPutStorageBucketAccessRequest() *PutStorageBucketAccessRequest`

NewPutStorageBucketAccessRequest instantiates a new PutStorageBucketAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutStorageBucketAccessRequestWithDefaults

`func NewPutStorageBucketAccessRequestWithDefaults() *PutStorageBucketAccessRequest`

NewPutStorageBucketAccessRequestWithDefaults instantiates a new PutStorageBucketAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *PutStorageBucketAccessRequest) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *PutStorageBucketAccessRequest) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *PutStorageBucketAccessRequest) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *PutStorageBucketAccessRequest) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetCorsEnabled

`func (o *PutStorageBucketAccessRequest) GetCorsEnabled() bool`

GetCorsEnabled returns the CorsEnabled field if non-nil, zero value otherwise.

### GetCorsEnabledOk

`func (o *PutStorageBucketAccessRequest) GetCorsEnabledOk() (*bool, bool)`

GetCorsEnabledOk returns a tuple with the CorsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsEnabled

`func (o *PutStorageBucketAccessRequest) SetCorsEnabled(v bool)`

SetCorsEnabled sets CorsEnabled field to given value.

### HasCorsEnabled

`func (o *PutStorageBucketAccessRequest) HasCorsEnabled() bool`

HasCorsEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


