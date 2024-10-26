# GetObjectStorageBucketAcl200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **string** | The Access Control Level of the bucket, as a canned ACL string. For more fine-grained control of ACLs, use the S3 API directly. | [optional] 
**AclXml** | Pointer to **string** | The full XML of the object&#39;s ACL policy. | [optional] 

## Methods

### NewGetObjectStorageBucketAcl200Response

`func NewGetObjectStorageBucketAcl200Response() *GetObjectStorageBucketAcl200Response`

NewGetObjectStorageBucketAcl200Response instantiates a new GetObjectStorageBucketAcl200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetObjectStorageBucketAcl200ResponseWithDefaults

`func NewGetObjectStorageBucketAcl200ResponseWithDefaults() *GetObjectStorageBucketAcl200Response`

NewGetObjectStorageBucketAcl200ResponseWithDefaults instantiates a new GetObjectStorageBucketAcl200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *GetObjectStorageBucketAcl200Response) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *GetObjectStorageBucketAcl200Response) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *GetObjectStorageBucketAcl200Response) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *GetObjectStorageBucketAcl200Response) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetAclXml

`func (o *GetObjectStorageBucketAcl200Response) GetAclXml() string`

GetAclXml returns the AclXml field if non-nil, zero value otherwise.

### GetAclXmlOk

`func (o *GetObjectStorageBucketAcl200Response) GetAclXmlOk() (*string, bool)`

GetAclXmlOk returns a tuple with the AclXml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclXml

`func (o *GetObjectStorageBucketAcl200Response) SetAclXml(v string)`

SetAclXml sets AclXml field to given value.

### HasAclXml

`func (o *GetObjectStorageBucketAcl200Response) HasAclXml() bool`

HasAclXml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


