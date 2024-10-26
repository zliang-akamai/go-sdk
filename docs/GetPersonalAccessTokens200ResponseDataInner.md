# GetPersonalAccessTokens200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date and time this token was created. | [optional] [readonly] 
**Expiry** | Pointer to **time.Time** | When this token will expire.  Personal Access Tokens cannot be renewed, so after this time the token will be completely unusable and a new token will need to be generated.  Tokens may be created with &#x60;null&#x60; as their expiry and will never expire unless revoked. | [optional] [readonly] 
**Id** | Pointer to **int32** | This token&#39;s unique ID, which can be used to revoke it. | [optional] [readonly] 
**Label** | Pointer to **string** | This token&#39;s label.  This is for display purposes only, but can be used to more easily track what you&#39;re using each token for. | [optional] 
**Scopes** | Pointer to **string** | The scopes this token was created with. These define what parts of the Account the token can be used to access. Many command-line tools, such as the [Linode CLI](https://github.com/linode/linode-cli), require tokens with access to &#x60;*&#x60;. Tokens with more restrictive scopes are generally more secure. | [optional] [readonly] 
**Token** | Pointer to **string** | The token used to access the API.  When the token is created, the full token is returned here.  Otherwise, only the first 16 characters are returned. | [optional] [readonly] 

## Methods

### NewGetPersonalAccessTokens200ResponseDataInner

`func NewGetPersonalAccessTokens200ResponseDataInner() *GetPersonalAccessTokens200ResponseDataInner`

NewGetPersonalAccessTokens200ResponseDataInner instantiates a new GetPersonalAccessTokens200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPersonalAccessTokens200ResponseDataInnerWithDefaults

`func NewGetPersonalAccessTokens200ResponseDataInnerWithDefaults() *GetPersonalAccessTokens200ResponseDataInner`

NewGetPersonalAccessTokens200ResponseDataInnerWithDefaults instantiates a new GetPersonalAccessTokens200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetPersonalAccessTokens200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetPersonalAccessTokens200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiry

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GetPersonalAccessTokens200ResponseDataInner) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GetPersonalAccessTokens200ResponseDataInner) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetId

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPersonalAccessTokens200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetPersonalAccessTokens200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetPersonalAccessTokens200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetPersonalAccessTokens200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScopes

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetPersonalAccessTokens200ResponseDataInner) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GetPersonalAccessTokens200ResponseDataInner) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetToken

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetPersonalAccessTokens200ResponseDataInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetPersonalAccessTokens200ResponseDataInner) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetPersonalAccessTokens200ResponseDataInner) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


