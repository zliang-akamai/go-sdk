# GetNotifications200ResponseDataInnerEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | The unique ID of the notification&#39;s entity, based on the entity type. Returns &#x60;null&#x60; for an &#x60;account&#x60; or &#x60;promotion&#x60; entity. | [optional] 
**Label** | Pointer to **NullableString** | The current label for this notification&#39;s entity.  Returns &#x60;null&#x60; for the following entity types:  - &#x60;entity_transfer&#x60; - &#x60;promotion&#x60; - &#x60;region&#x60; | [optional] 
**Type** | Pointer to **string** | The type of entity this is related to. | [optional] 
**Url** | Pointer to **NullableString** | The URL where you can access the notification&#39;s object. The URL is relative to the domain where you retrieved the notification. This value is &#x60;null&#x60; for the &#x60;promotion&#x60; entity type. | [optional] 

## Methods

### NewGetNotifications200ResponseDataInnerEntity

`func NewGetNotifications200ResponseDataInnerEntity() *GetNotifications200ResponseDataInnerEntity`

NewGetNotifications200ResponseDataInnerEntity instantiates a new GetNotifications200ResponseDataInnerEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotifications200ResponseDataInnerEntityWithDefaults

`func NewGetNotifications200ResponseDataInnerEntityWithDefaults() *GetNotifications200ResponseDataInnerEntity`

NewGetNotifications200ResponseDataInnerEntityWithDefaults instantiates a new GetNotifications200ResponseDataInnerEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNotifications200ResponseDataInnerEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNotifications200ResponseDataInnerEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNotifications200ResponseDataInnerEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetNotifications200ResponseDataInnerEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *GetNotifications200ResponseDataInnerEntity) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *GetNotifications200ResponseDataInnerEntity) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetLabel

`func (o *GetNotifications200ResponseDataInnerEntity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetNotifications200ResponseDataInnerEntity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetNotifications200ResponseDataInnerEntity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetNotifications200ResponseDataInnerEntity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *GetNotifications200ResponseDataInnerEntity) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *GetNotifications200ResponseDataInnerEntity) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetType

`func (o *GetNotifications200ResponseDataInnerEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNotifications200ResponseDataInnerEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNotifications200ResponseDataInnerEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNotifications200ResponseDataInnerEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *GetNotifications200ResponseDataInnerEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetNotifications200ResponseDataInnerEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetNotifications200ResponseDataInnerEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetNotifications200ResponseDataInnerEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GetNotifications200ResponseDataInnerEntity) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GetNotifications200ResponseDataInnerEntity) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


