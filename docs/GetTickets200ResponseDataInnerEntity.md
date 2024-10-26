# GetTickets200ResponseDataInnerEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID for this Ticket&#39;s entity. | [optional] [readonly] 
**Label** | Pointer to **string** | The current label of this entity. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of entity this is related to. | [optional] [readonly] 
**Url** | Pointer to **string** | The URL where you can access the object this event is for. If a relative URL, it is relative to the domain you retrieved the entity from. | [optional] [readonly] 

## Methods

### NewGetTickets200ResponseDataInnerEntity

`func NewGetTickets200ResponseDataInnerEntity() *GetTickets200ResponseDataInnerEntity`

NewGetTickets200ResponseDataInnerEntity instantiates a new GetTickets200ResponseDataInnerEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTickets200ResponseDataInnerEntityWithDefaults

`func NewGetTickets200ResponseDataInnerEntityWithDefaults() *GetTickets200ResponseDataInnerEntity`

NewGetTickets200ResponseDataInnerEntityWithDefaults instantiates a new GetTickets200ResponseDataInnerEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTickets200ResponseDataInnerEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTickets200ResponseDataInnerEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTickets200ResponseDataInnerEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTickets200ResponseDataInnerEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetTickets200ResponseDataInnerEntity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetTickets200ResponseDataInnerEntity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetTickets200ResponseDataInnerEntity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetTickets200ResponseDataInnerEntity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *GetTickets200ResponseDataInnerEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetTickets200ResponseDataInnerEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetTickets200ResponseDataInnerEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetTickets200ResponseDataInnerEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *GetTickets200ResponseDataInnerEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetTickets200ResponseDataInnerEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetTickets200ResponseDataInnerEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetTickets200ResponseDataInnerEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


