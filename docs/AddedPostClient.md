# AddedPostClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The OAuth Client ID.  This is used to identify the client, and is a publicly known value (it is not a secret). | [optional] [readonly] 
**Label** | **string** | The name of this application.  This will be presented to users when they are asked to grant it access to their Account. | 
**Public** | Pointer to **bool** | If this is a public or private OAuth Client.  Public clients have a slightly different authentication workflow than private clients.  See the [OAuth spec](https://oauth.net/2/) for more details. | [optional] [default to false]
**RedirectUri** | **string** | The location a successful log in from [login.linode.com](https://login.linode.com) should be redirected to for this client.  The receiver of this redirect should be ready to accept an OAuth exchange code and finish the OAuth exchange. | 
**Secret** | Pointer to **string** | The OAuth Client secret, used in the OAuth exchange.  This is returned as &#x60;&lt;REDACTED&gt;&#x60; except when an OAuth Client is created or its secret is reset.  This is a secret, and should not be shared or disclosed publicly. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of this application.  &#x60;active&#x60; by default. | [optional] [readonly] 
**ThumbnailUrl** | Pointer to **NullableString** | The URL where this client&#39;s thumbnail may be viewed, or &#x60;null&#x60; if this client does not have a thumbnail set. | [optional] [readonly] 

## Methods

### NewAddedPostClient

`func NewAddedPostClient(label string, redirectUri string, ) *AddedPostClient`

NewAddedPostClient instantiates a new AddedPostClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedPostClientWithDefaults

`func NewAddedPostClientWithDefaults() *AddedPostClient`

NewAddedPostClientWithDefaults instantiates a new AddedPostClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddedPostClient) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddedPostClient) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddedPostClient) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddedPostClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *AddedPostClient) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AddedPostClient) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AddedPostClient) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetPublic

`func (o *AddedPostClient) GetPublic() bool`

GetPublic returns the Public field if non-nil, zero value otherwise.

### GetPublicOk

`func (o *AddedPostClient) GetPublicOk() (*bool, bool)`

GetPublicOk returns a tuple with the Public field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublic

`func (o *AddedPostClient) SetPublic(v bool)`

SetPublic sets Public field to given value.

### HasPublic

`func (o *AddedPostClient) HasPublic() bool`

HasPublic returns a boolean if a field has been set.

### GetRedirectUri

`func (o *AddedPostClient) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *AddedPostClient) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *AddedPostClient) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.


### GetSecret

`func (o *AddedPostClient) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *AddedPostClient) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *AddedPostClient) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *AddedPostClient) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetStatus

`func (o *AddedPostClient) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddedPostClient) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddedPostClient) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddedPostClient) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *AddedPostClient) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *AddedPostClient) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *AddedPostClient) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *AddedPostClient) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *AddedPostClient) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *AddedPostClient) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


