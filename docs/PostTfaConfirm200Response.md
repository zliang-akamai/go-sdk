# PostTfaConfirm200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scratch** | Pointer to **string** | A one-use code that can be used in place of your Two Factor code, in case you are unable to generate one.  Keep this in a safe place to avoid being locked out of your Account. | [optional] 

## Methods

### NewPostTfaConfirm200Response

`func NewPostTfaConfirm200Response() *PostTfaConfirm200Response`

NewPostTfaConfirm200Response instantiates a new PostTfaConfirm200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTfaConfirm200ResponseWithDefaults

`func NewPostTfaConfirm200ResponseWithDefaults() *PostTfaConfirm200Response`

NewPostTfaConfirm200ResponseWithDefaults instantiates a new PostTfaConfirm200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScratch

`func (o *PostTfaConfirm200Response) GetScratch() string`

GetScratch returns the Scratch field if non-nil, zero value otherwise.

### GetScratchOk

`func (o *PostTfaConfirm200Response) GetScratchOk() (*string, bool)`

GetScratchOk returns a tuple with the Scratch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScratch

`func (o *PostTfaConfirm200Response) SetScratch(v string)`

SetScratch sets Scratch field to given value.

### HasScratch

`func (o *PostTfaConfirm200Response) HasScratch() bool`

HasScratch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


