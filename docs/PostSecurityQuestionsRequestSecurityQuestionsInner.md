# PostSecurityQuestionsRequestSecurityQuestionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuestionId** | Pointer to **int32** | The ID representing the security question. | [optional] 
**Response** | Pointer to **string** | The security question response. | [optional] 
**SecurityQuestion** | Pointer to **string** | The security question. | [optional] [readonly] 

## Methods

### NewPostSecurityQuestionsRequestSecurityQuestionsInner

`func NewPostSecurityQuestionsRequestSecurityQuestionsInner() *PostSecurityQuestionsRequestSecurityQuestionsInner`

NewPostSecurityQuestionsRequestSecurityQuestionsInner instantiates a new PostSecurityQuestionsRequestSecurityQuestionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSecurityQuestionsRequestSecurityQuestionsInnerWithDefaults

`func NewPostSecurityQuestionsRequestSecurityQuestionsInnerWithDefaults() *PostSecurityQuestionsRequestSecurityQuestionsInner`

NewPostSecurityQuestionsRequestSecurityQuestionsInnerWithDefaults instantiates a new PostSecurityQuestionsRequestSecurityQuestionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestionId

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) GetQuestionId() int32`

GetQuestionId returns the QuestionId field if non-nil, zero value otherwise.

### GetQuestionIdOk

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) GetQuestionIdOk() (*int32, bool)`

GetQuestionIdOk returns a tuple with the QuestionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestionId

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) SetQuestionId(v int32)`

SetQuestionId sets QuestionId field to given value.

### HasQuestionId

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) HasQuestionId() bool`

HasQuestionId returns a boolean if a field has been set.

### GetResponse

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetSecurityQuestion

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) GetSecurityQuestion() string`

GetSecurityQuestion returns the SecurityQuestion field if non-nil, zero value otherwise.

### GetSecurityQuestionOk

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) GetSecurityQuestionOk() (*string, bool)`

GetSecurityQuestionOk returns a tuple with the SecurityQuestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityQuestion

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) SetSecurityQuestion(v string)`

SetSecurityQuestion sets SecurityQuestion field to given value.

### HasSecurityQuestion

`func (o *PostSecurityQuestionsRequestSecurityQuestionsInner) HasSecurityQuestion() bool`

HasSecurityQuestion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


