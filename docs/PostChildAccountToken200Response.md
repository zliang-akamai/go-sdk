# PostChildAccountToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date and time this token was created. | [optional] [readonly] 
**Expiry** | Pointer to **time.Time** | When this token expires. This is default set to 15 minutes from the time of creation. Proxy user tokens can&#39;t be renewed. After this time, Akamai revokes the token and you need to generate a new one. | [optional] [readonly] 
**Id** | Pointer to **int32** | The proxy user token&#39;s unique ID, which can be used to revoke it. | [optional] [readonly] 
**Label** | Pointer to **string** | The name of the token. The API automatically sets this to &#x60;&lt;username&gt;_&lt;uid&gt;_&lt;time&gt;&#x60;. It&#39;s composed of the &#x60;username&#x60; for your parent account user, the unique &#x60;uid&#x60; Akamai assigned to identify your user, and the &#x60;time&#x60; the API generated the token. This is for display purposes only, but you can use it to help track how you&#39;re using each proxy user token. | [optional] 
**Scopes** | Pointer to **string** | The scopes this token was created with. Defaults to &#x60;*&#x60;. Proxy user tokens automatically inherit all the permissions of the proxy user. | [optional] [readonly] 
**Token** | Pointer to **string** | The proxy user token that can be used to access the API and CLI. After you [create](https://techdocs.akamai.com/linode-api/reference/post-child-account-token) a token, you can see the full token in the response. All other operations that contain this token only show the first 16 characters in their response. | [optional] [readonly] 

## Methods

### NewPostChildAccountToken200Response

`func NewPostChildAccountToken200Response() *PostChildAccountToken200Response`

NewPostChildAccountToken200Response instantiates a new PostChildAccountToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostChildAccountToken200ResponseWithDefaults

`func NewPostChildAccountToken200ResponseWithDefaults() *PostChildAccountToken200Response`

NewPostChildAccountToken200ResponseWithDefaults instantiates a new PostChildAccountToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PostChildAccountToken200Response) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PostChildAccountToken200Response) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PostChildAccountToken200Response) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PostChildAccountToken200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiry

`func (o *PostChildAccountToken200Response) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *PostChildAccountToken200Response) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *PostChildAccountToken200Response) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *PostChildAccountToken200Response) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetId

`func (o *PostChildAccountToken200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostChildAccountToken200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostChildAccountToken200Response) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostChildAccountToken200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PostChildAccountToken200Response) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostChildAccountToken200Response) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostChildAccountToken200Response) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostChildAccountToken200Response) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScopes

`func (o *PostChildAccountToken200Response) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *PostChildAccountToken200Response) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *PostChildAccountToken200Response) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *PostChildAccountToken200Response) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetToken

`func (o *PostChildAccountToken200Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PostChildAccountToken200Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PostChildAccountToken200Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PostChildAccountToken200Response) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


