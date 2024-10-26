# PostSecurityQuestionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityQuestions** | Pointer to [**[]PostSecurityQuestionsRequestSecurityQuestionsInner**](PostSecurityQuestionsRequestSecurityQuestionsInner.md) | Security questions and response objects. | [optional] 

## Methods

### NewPostSecurityQuestionsRequest

`func NewPostSecurityQuestionsRequest() *PostSecurityQuestionsRequest`

NewPostSecurityQuestionsRequest instantiates a new PostSecurityQuestionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSecurityQuestionsRequestWithDefaults

`func NewPostSecurityQuestionsRequestWithDefaults() *PostSecurityQuestionsRequest`

NewPostSecurityQuestionsRequestWithDefaults instantiates a new PostSecurityQuestionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityQuestions

`func (o *PostSecurityQuestionsRequest) GetSecurityQuestions() []PostSecurityQuestionsRequestSecurityQuestionsInner`

GetSecurityQuestions returns the SecurityQuestions field if non-nil, zero value otherwise.

### GetSecurityQuestionsOk

`func (o *PostSecurityQuestionsRequest) GetSecurityQuestionsOk() (*[]PostSecurityQuestionsRequestSecurityQuestionsInner, bool)`

GetSecurityQuestionsOk returns a tuple with the SecurityQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityQuestions

`func (o *PostSecurityQuestionsRequest) SetSecurityQuestions(v []PostSecurityQuestionsRequestSecurityQuestionsInner)`

SetSecurityQuestions sets SecurityQuestions field to given value.

### HasSecurityQuestions

`func (o *PostSecurityQuestionsRequest) HasSecurityQuestions() bool`

HasSecurityQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


