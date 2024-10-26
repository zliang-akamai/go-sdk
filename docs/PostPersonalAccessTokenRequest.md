# PostPersonalAccessTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiry** | Pointer to **time.Time** | When this token should be valid until.  If omitted, the new token will be valid until it is manually revoked. | [optional] 
**Label** | Pointer to **string** | This token&#39;s label.  This is for display purposes only, but can be used to more easily track what you&#39;re using each token for. | [optional] 
**Scopes** | Pointer to **string** | The access [scopes](https://techdocs.akamai.com/linode-api/reference/get-started#oauth-reference) to grant to the created token. These cannot be changed after creation, and may not exceed the scopes of the acting token.  If omitted or entered with a wildcard character (&#x60;*&#x60;), the new token will have the same scopes as the acting token.  Multiple scopes are separated by a space character (&#x60; &#x60;).  For example, &#x60;linodes:read_write account:read_only&#x60;. | [optional] 

## Methods

### NewPostPersonalAccessTokenRequest

`func NewPostPersonalAccessTokenRequest() *PostPersonalAccessTokenRequest`

NewPostPersonalAccessTokenRequest instantiates a new PostPersonalAccessTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPersonalAccessTokenRequestWithDefaults

`func NewPostPersonalAccessTokenRequestWithDefaults() *PostPersonalAccessTokenRequest`

NewPostPersonalAccessTokenRequestWithDefaults instantiates a new PostPersonalAccessTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiry

`func (o *PostPersonalAccessTokenRequest) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PostPersonalAccessTokenRequest) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PostPersonalAccessTokenRequest) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *PostPersonalAccessTokenRequest) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetLabel

`func (o *PostPersonalAccessTokenRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostPersonalAccessTokenRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostPersonalAccessTokenRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostPersonalAccessTokenRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScopes

`func (o *PostPersonalAccessTokenRequest) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostPersonalAccessTokenRequest) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostPersonalAccessTokenRequest) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PostPersonalAccessTokenRequest) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


