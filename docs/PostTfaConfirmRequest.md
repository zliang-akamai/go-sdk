# PostTfaConfirmRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TfaCode** | Pointer to **string** | The Two Factor code you generated with your Two Factor secret. These codes are time-based, so be sure it is current. | [optional] 

## Methods

### NewPostTfaConfirmRequest

`func NewPostTfaConfirmRequest() *PostTfaConfirmRequest`

NewPostTfaConfirmRequest instantiates a new PostTfaConfirmRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTfaConfirmRequestWithDefaults

`func NewPostTfaConfirmRequestWithDefaults() *PostTfaConfirmRequest`

NewPostTfaConfirmRequestWithDefaults instantiates a new PostTfaConfirmRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTfaCode

`func (o *PostTfaConfirmRequest) GetTfaCode() string`

GetTfaCode returns the TfaCode field if non-nil, zero value otherwise.

### GetTfaCodeOk

`func (o *PostTfaConfirmRequest) GetTfaCodeOk() (*string, bool)`

GetTfaCodeOk returns a tuple with the TfaCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfaCode

`func (o *PostTfaConfirmRequest) SetTfaCode(v string)`

SetTfaCode sets TfaCode field to given value.

### HasTfaCode

`func (o *PostTfaConfirmRequest) HasTfaCode() bool`

HasTfaCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


