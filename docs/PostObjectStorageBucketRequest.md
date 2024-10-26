# PostObjectStorageBucketRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to **string** | The Access Control Level of the bucket using a canned ACL string. For more fine-grained control of ACLs, use the S3 API directly. | [optional] [default to "private"]
**CorsEnabled** | Pointer to **bool** | If set to &#x60;false&#x60;, CORS is disabled for all origins in the bucket. For more fine-grained controls of CORS, use the S3 API directly. | [optional] [default to true]
**Label** | **string** | The name for this bucket.  * A bucket name can contain from 3 to 63 alphanumeric characters, dashes (&#x60;-&#x60;), or dots (&#x60;.&#x60;). * A bucket name can&#39;t end in a dash and you can&#39;t use two consecutive dashes. * A bucket name can&#39;t start or end in a dot, and you can&#39;t use two consecutive dots. As a best practice, only use dots if a certificate you&#39;re using with your bucket requires it. (For example, if you&#39;re using a custom TLS certificate.) * A bucket name needs to be unique in the &#x60;region&#x60; where you&#39;re creating the bucket. The API only reserves labels for the &#x60;region&#x60; where active buckets are created and stored. If you want to reserve this bucket&#39;s label in another &#x60;region&#x60;, create a new bucket with the same label in the new &#x60;region&#x60;. | 
**Region** | Pointer to **string** | The &#x60;id&#x60; assigned to the data center ([region](https://techdocs.akamai.com/linode-api/reference/get-regions)) where this Object Storage bucket should be created.  &gt; 📘 &gt; &gt; This supports legacy &#x60;clusterId&#x60; values that represented a specific region. For example, &#x60;us-east-1&#x60; is the legacy reference for the &#x60;us-east&#x60; region. | [optional] 

## Methods

### NewPostObjectStorageBucketRequest

`func NewPostObjectStorageBucketRequest(label string, ) *PostObjectStorageBucketRequest`

NewPostObjectStorageBucketRequest instantiates a new PostObjectStorageBucketRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostObjectStorageBucketRequestWithDefaults

`func NewPostObjectStorageBucketRequestWithDefaults() *PostObjectStorageBucketRequest`

NewPostObjectStorageBucketRequestWithDefaults instantiates a new PostObjectStorageBucketRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *PostObjectStorageBucketRequest) GetAcl() string`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *PostObjectStorageBucketRequest) GetAclOk() (*string, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *PostObjectStorageBucketRequest) SetAcl(v string)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *PostObjectStorageBucketRequest) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetCorsEnabled

`func (o *PostObjectStorageBucketRequest) GetCorsEnabled() bool`

GetCorsEnabled returns the CorsEnabled field if non-nil, zero value otherwise.

### GetCorsEnabledOk

`func (o *PostObjectStorageBucketRequest) GetCorsEnabledOk() (*bool, bool)`

GetCorsEnabledOk returns a tuple with the CorsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsEnabled

`func (o *PostObjectStorageBucketRequest) SetCorsEnabled(v bool)`

SetCorsEnabled sets CorsEnabled field to given value.

### HasCorsEnabled

`func (o *PostObjectStorageBucketRequest) HasCorsEnabled() bool`

HasCorsEnabled returns a boolean if a field has been set.

### GetLabel

`func (o *PostObjectStorageBucketRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostObjectStorageBucketRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostObjectStorageBucketRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetRegion

`func (o *PostObjectStorageBucketRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostObjectStorageBucketRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostObjectStorageBucketRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostObjectStorageBucketRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


