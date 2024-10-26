# GetSecurityQuestions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityQuestions** | Pointer to [**[]GetSecurityQuestions200ResponseSecurityQuestionsInner**](GetSecurityQuestions200ResponseSecurityQuestionsInner.md) | Security questions and response objects. | [optional] 

## Methods

### NewGetSecurityQuestions200Response

`func NewGetSecurityQuestions200Response() *GetSecurityQuestions200Response`

NewGetSecurityQuestions200Response instantiates a new GetSecurityQuestions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSecurityQuestions200ResponseWithDefaults

`func NewGetSecurityQuestions200ResponseWithDefaults() *GetSecurityQuestions200Response`

NewGetSecurityQuestions200ResponseWithDefaults instantiates a new GetSecurityQuestions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityQuestions

`func (o *GetSecurityQuestions200Response) GetSecurityQuestions() []GetSecurityQuestions200ResponseSecurityQuestionsInner`

GetSecurityQuestions returns the SecurityQuestions field if non-nil, zero value otherwise.

### GetSecurityQuestionsOk

`func (o *GetSecurityQuestions200Response) GetSecurityQuestionsOk() (*[]GetSecurityQuestions200ResponseSecurityQuestionsInner, bool)`

GetSecurityQuestionsOk returns a tuple with the SecurityQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityQuestions

`func (o *GetSecurityQuestions200Response) SetSecurityQuestions(v []GetSecurityQuestions200ResponseSecurityQuestionsInner)`

SetSecurityQuestions sets SecurityQuestions field to given value.

### HasSecurityQuestions

`func (o *GetSecurityQuestions200Response) HasSecurityQuestions() bool`

HasSecurityQuestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


