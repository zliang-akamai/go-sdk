# GetProfileApps200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this app was authorized. | [optional] [readonly] 
**Expiry** | Pointer to **NullableTime** | When the app&#39;s access to your account expires. If &#x60;null&#x60;, the app&#39;s access must be revoked manually. | [optional] [readonly] 
**Id** | Pointer to **int32** | This authorization&#39;s ID, used for revoking access. | [optional] [readonly] 
**Label** | Pointer to **string** | The name of the application you&#39;ve authorized. | [optional] [readonly] 
**Scopes** | Pointer to **string** | The OAuth scopes this app was authorized with.  This defines what parts of your Account the app is allowed to access. | [optional] [readonly] 
**ThumbnailUrl** | Pointer to **string** | The URL at which this app&#39;s thumbnail may be accessed. | [optional] [readonly] 
**Website** | Pointer to **string** | The website where you can get more information about this app. | [optional] [readonly] 

## Methods

### NewGetProfileApps200ResponseDataInner

`func NewGetProfileApps200ResponseDataInner() *GetProfileApps200ResponseDataInner`

NewGetProfileApps200ResponseDataInner instantiates a new GetProfileApps200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProfileApps200ResponseDataInnerWithDefaults

`func NewGetProfileApps200ResponseDataInnerWithDefaults() *GetProfileApps200ResponseDataInner`

NewGetProfileApps200ResponseDataInnerWithDefaults instantiates a new GetProfileApps200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetProfileApps200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetProfileApps200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetProfileApps200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetProfileApps200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiry

`func (o *GetProfileApps200ResponseDataInner) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GetProfileApps200ResponseDataInner) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GetProfileApps200ResponseDataInner) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GetProfileApps200ResponseDataInner) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### SetExpiryNil

`func (o *GetProfileApps200ResponseDataInner) SetExpiryNil(b bool)`

 SetExpiryNil sets the value for Expiry to be an explicit nil

### UnsetExpiry
`func (o *GetProfileApps200ResponseDataInner) UnsetExpiry()`

UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
### GetId

`func (o *GetProfileApps200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetProfileApps200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetProfileApps200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetProfileApps200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetProfileApps200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetProfileApps200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetProfileApps200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetProfileApps200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetScopes

`func (o *GetProfileApps200ResponseDataInner) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetProfileApps200ResponseDataInner) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetProfileApps200ResponseDataInner) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *GetProfileApps200ResponseDataInner) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *GetProfileApps200ResponseDataInner) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *GetProfileApps200ResponseDataInner) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *GetProfileApps200ResponseDataInner) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *GetProfileApps200ResponseDataInner) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetWebsite

`func (o *GetProfileApps200ResponseDataInner) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *GetProfileApps200ResponseDataInner) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *GetProfileApps200ResponseDataInner) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *GetProfileApps200ResponseDataInner) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


