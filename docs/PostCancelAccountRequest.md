# PostCancelAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to **string** | Any reason for cancelling the account, and any other comments you might have about your Linode service. | [optional] 

## Methods

### NewPostCancelAccountRequest

`func NewPostCancelAccountRequest() *PostCancelAccountRequest`

NewPostCancelAccountRequest instantiates a new PostCancelAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCancelAccountRequestWithDefaults

`func NewPostCancelAccountRequestWithDefaults() *PostCancelAccountRequest`

NewPostCancelAccountRequestWithDefaults instantiates a new PostCancelAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *PostCancelAccountRequest) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PostCancelAccountRequest) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PostCancelAccountRequest) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PostCancelAccountRequest) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


