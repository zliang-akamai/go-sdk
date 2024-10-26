# GetClients200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The OAuth Client ID.  This is used to identify the client, and is a publicly known value (it is not a secret). | [optional] [readonly] 
**Label** | Pointer to **string** | The name of this application.  This will be presented to users when they are asked to grant it access to their Account. | [optional] 
**Public** | Pointer to **bool** | If this is a public or private OAuth Client.  Public clients have a slightly different authentication workflow than private clients.  See the [OAuth spec](https://oauth.net/2/) for more details. | [optional] [default to false]
**RedirectUri** | Pointer to **string** | The location a successful log in from [login.linode.com](https://login.linode.com) should be redirected to for this client.  The receiver of this redirect should be ready to accept an OAuth exchange code and finish the OAuth exchange. | [optional] 
**Secret** | Pointer to **string** | The OAuth Client secret, used in the OAuth exchange.  This is returned as &#x60;&lt;REDACTED&gt;&#x60; except when an OAuth Client is created or its secret is reset.  This is a secret, and should not be shared or disclosed publicly. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of this application.  &#x60;active&#x60; by default. | [optional] [readonly] 
**ThumbnailUrl** | Pointer to **NullableString** | The URL where this client&#39;s thumbnail may be viewed, or &#x60;null&#x60; if this client does not have a thumbnail set. | [optional] [readonly] 

## Methods

### NewGetClients200ResponseDataInner

`func NewGetClients200ResponseDataInner() *GetClients200ResponseDataInner`

NewGetClients200ResponseDataInner instantiates a new GetClients200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClients200ResponseDataInnerWithDefaults

`func NewGetClients200ResponseDataInnerWithDefaults() *GetClients200ResponseDataInner`

NewGetClients200ResponseDataInnerWithDefaults instantiates a new GetClients200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClients200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClients200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClients200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetClients200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetClients200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetClients200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetClients200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetClients200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPublic

`func (o *GetClients200ResponseDataInner) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *GetClients200ResponseDataInner) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *GetClients200ResponseDataInner) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *GetClients200ResponseDataInner) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRedirectUri

`func (o *GetClients200ResponseDataInner) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *GetClients200ResponseDataInner) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *GetClients200ResponseDataInner) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *GetClients200ResponseDataInner) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.

### GetSecret

`func (o *GetClients200ResponseDataInner) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *GetClients200ResponseDataInner) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *GetClients200ResponseDataInner) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *GetClients200ResponseDataInner) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatus

`func (o *GetClients200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClients200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClients200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClients200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *GetClients200ResponseDataInner) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *GetClients200ResponseDataInner) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *GetClients200ResponseDataInner) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *GetClients200ResponseDataInner) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *GetClients200ResponseDataInner) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *GetClients200ResponseDataInner) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


