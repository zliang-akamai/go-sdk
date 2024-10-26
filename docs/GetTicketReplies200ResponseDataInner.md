# GetTicketReplies200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date and time this Ticket reply was created. | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The User who submitted this reply. | [optional] [readonly] 
**Description** | Pointer to **string** | The body of this Support Ticket reply. | [optional] [readonly] 
**FromLinode** | Pointer to **bool** | If set to true, this reply came from a Linode employee. | [optional] [readonly] 
**GravatarId** | Pointer to **string** | The Gravatar ID of the User who created this reply. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID of this Support Ticket reply. | [optional] [readonly] 

## Methods

### NewGetTicketReplies200ResponseDataInner

`func NewGetTicketReplies200ResponseDataInner() *GetTicketReplies200ResponseDataInner`

NewGetTicketReplies200ResponseDataInner instantiates a new GetTicketReplies200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTicketReplies200ResponseDataInnerWithDefaults

`func NewGetTicketReplies200ResponseDataInnerWithDefaults() *GetTicketReplies200ResponseDataInner`

NewGetTicketReplies200ResponseDataInnerWithDefaults instantiates a new GetTicketReplies200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetTicketReplies200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetTicketReplies200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetTicketReplies200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetTicketReplies200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetTicketReplies200ResponseDataInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetTicketReplies200ResponseDataInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetTicketReplies200ResponseDataInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetTicketReplies200ResponseDataInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDescription

`func (o *GetTicketReplies200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTicketReplies200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTicketReplies200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTicketReplies200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFromLinode

`func (o *GetTicketReplies200ResponseDataInner) GetFromLinode() bool`

GetFromLinode returns the FromLinode field if non-nil, zero value otherwise.

### GetFromLinodeOk

`func (o *GetTicketReplies200ResponseDataInner) GetFromLinodeOk() (*bool, bool)`

GetFromLinodeOk returns a tuple with the FromLinode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromLinode

`func (o *GetTicketReplies200ResponseDataInner) SetFromLinode(v bool)`

SetFromLinode sets FromLinode field to given value.

### HasFromLinode

`func (o *GetTicketReplies200ResponseDataInner) HasFromLinode() bool`

HasFromLinode returns a boolean if a field has been set.

### GetGravatarId

`func (o *GetTicketReplies200ResponseDataInner) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *GetTicketReplies200ResponseDataInner) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *GetTicketReplies200ResponseDataInner) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.

### HasGravatarId

`func (o *GetTicketReplies200ResponseDataInner) HasGravatarId() bool`

HasGravatarId returns a boolean if a field has been set.

### GetId

`func (o *GetTicketReplies200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTicketReplies200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTicketReplies200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTicketReplies200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


