# GetEvents200ResponseDataInnerEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The unique ID for an Event&#39;s entity.  Some Event entities do not have IDs associated with them, so they will not be returned when filtering by ID. These Events include:    - &#x60;account&#x60;   - &#x60;profile&#x60;  Entities for some Events are assigned the ID of the Linode they correspond to. When filtering by ID for these Events, use the corresponding Linode&#39;s ID. These Events include:    - &#x60;disks&#x60;   - &#x60;backups&#x60;  Tag Events use a tag&#39;s name for the entity ID field. When filtering by ID for tag Events, supply the name of the tag. | [optional] 
**Label** | Pointer to **string** | The current label of this object. The label may reflect changes that occur with this Event. | [optional] 
**Type** | Pointer to **string** | The type of entity that is being referenced by the Event. | [optional] [readonly] 
**Url** | Pointer to **string** | The URL where you can access the object this Event is for. If a relative URL, it is relative to the domain you retrieved the Event from. | [optional] 

## Methods

### NewGetEvents200ResponseDataInnerEntity

`func NewGetEvents200ResponseDataInnerEntity() *GetEvents200ResponseDataInnerEntity`

NewGetEvents200ResponseDataInnerEntity instantiates a new GetEvents200ResponseDataInnerEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEvents200ResponseDataInnerEntityWithDefaults

`func NewGetEvents200ResponseDataInnerEntityWithDefaults() *GetEvents200ResponseDataInnerEntity`

NewGetEvents200ResponseDataInnerEntityWithDefaults instantiates a new GetEvents200ResponseDataInnerEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEvents200ResponseDataInnerEntity) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEvents200ResponseDataInnerEntity) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEvents200ResponseDataInnerEntity) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetEvents200ResponseDataInnerEntity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetEvents200ResponseDataInnerEntity) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetEvents200ResponseDataInnerEntity) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetEvents200ResponseDataInnerEntity) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetEvents200ResponseDataInnerEntity) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *GetEvents200ResponseDataInnerEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetEvents200ResponseDataInnerEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetEvents200ResponseDataInnerEntity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetEvents200ResponseDataInnerEntity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *GetEvents200ResponseDataInnerEntity) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetEvents200ResponseDataInnerEntity) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetEvents200ResponseDataInnerEntity) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetEvents200ResponseDataInnerEntity) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


