# PostResetLinodePasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RootPass** | **string** | The root user&#39;s password on this Linode. Linode passwords must meet a password strength score requirement that is calculated internally by the API. If the strength requirement is not met, you will receive a Password does not meet strength requirement error. | 

## Methods

### NewPostResetLinodePasswordRequest

`func NewPostResetLinodePasswordRequest(rootPass string, ) *PostResetLinodePasswordRequest`

NewPostResetLinodePasswordRequest instantiates a new PostResetLinodePasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostResetLinodePasswordRequestWithDefaults

`func NewPostResetLinodePasswordRequestWithDefaults() *PostResetLinodePasswordRequest`

NewPostResetLinodePasswordRequestWithDefaults instantiates a new PostResetLinodePasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRootPass

`func (o *PostResetLinodePasswordRequest) GetRootPass() string`

GetRootPass returns the RootPass field if non-nil, zero value otherwise.

### GetRootPassOk

`func (o *PostResetLinodePasswordRequest) GetRootPassOk() (*string, bool)`

GetRootPassOk returns a tuple with the RootPass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPass

`func (o *PostResetLinodePasswordRequest) SetRootPass(v string)`

SetRootPass sets RootPass field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


