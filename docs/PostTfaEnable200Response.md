# PostTfaEnable200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **time.Time** | When this Two Factor secret expires. | [optional] 
**Secret** | Pointer to **string** | Your Two Factor secret. This is used to generate time-based two factor codes required for logging in. Doing this will be required to confirm TFA an actually enable it. | [optional] 

## Methods

### NewPostTfaEnable200Response

`func NewPostTfaEnable200Response() *PostTfaEnable200Response`

NewPostTfaEnable200Response instantiates a new PostTfaEnable200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostTfaEnable200ResponseWithDefaults

`func NewPostTfaEnable200ResponseWithDefaults() *PostTfaEnable200Response`

NewPostTfaEnable200ResponseWithDefaults instantiates a new PostTfaEnable200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *PostTfaEnable200Response) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PostTfaEnable200Response) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PostTfaEnable200Response) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *PostTfaEnable200Response) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSecret

`func (o *PostTfaEnable200Response) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PostTfaEnable200Response) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PostTfaEnable200Response) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PostTfaEnable200Response) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


