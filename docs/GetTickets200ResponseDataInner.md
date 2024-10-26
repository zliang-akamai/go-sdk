# GetTickets200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Attachments** | Pointer to **[]string** | A list of filenames representing attached files associated with this Ticket. | [optional] [readonly] 
**Closable** | Pointer to **bool** | Whether the Support Ticket may be closed. | [optional] 
**Closed** | Pointer to **NullableTime** | The date and time this Ticket was closed. | [optional] [readonly] 
**Description** | Pointer to **string** | The full details of the issue or question. | [optional] [readonly] 
**Entity** | Pointer to [**NullableGetTickets200ResponseDataInnerEntity**](GetTickets200ResponseDataInnerEntity.md) |  | [optional] 
**GravatarId** | Pointer to **string** | The Gravatar ID of the User who opened this Ticket. | [optional] [readonly] 
**Id** | Pointer to **int32** | The ID of the Support Ticket. | [optional] [readonly] 
**Opened** | Pointer to **time.Time** | The date and time this Ticket was created. | [optional] [readonly] 
**OpenedBy** | Pointer to **string** | The User who opened this Ticket. | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of this Ticket. | [optional] [readonly] 
**Summary** | Pointer to **string** | The summary or title for this Ticket. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | The date and time this Ticket was last updated. | [optional] [readonly] 
**UpdatedBy** | Pointer to **NullableString** | The User who last updated this Ticket. | [optional] [readonly] 

## Methods

### NewGetTickets200ResponseDataInner

`func NewGetTickets200ResponseDataInner() *GetTickets200ResponseDataInner`

NewGetTickets200ResponseDataInner instantiates a new GetTickets200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTickets200ResponseDataInnerWithDefaults

`func NewGetTickets200ResponseDataInnerWithDefaults() *GetTickets200ResponseDataInner`

NewGetTickets200ResponseDataInnerWithDefaults instantiates a new GetTickets200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttachments

`func (o *GetTickets200ResponseDataInner) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetTickets200ResponseDataInner) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetTickets200ResponseDataInner) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetTickets200ResponseDataInner) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetClosable

`func (o *GetTickets200ResponseDataInner) GetClosable() bool`

GetClosable returns the Closable field if non-nil, zero value otherwise.

### GetClosableOk

`func (o *GetTickets200ResponseDataInner) GetClosableOk() (*bool, bool)`

GetClosableOk returns a tuple with the Closable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosable

`func (o *GetTickets200ResponseDataInner) SetClosable(v bool)`

SetClosable sets Closable field to given value.

### HasClosable

`func (o *GetTickets200ResponseDataInner) HasClosable() bool`

HasClosable returns a boolean if a field has been set.

### GetClosed

`func (o *GetTickets200ResponseDataInner) GetClosed() time.Time`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *GetTickets200ResponseDataInner) GetClosedOk() (*time.Time, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *GetTickets200ResponseDataInner) SetClosed(v time.Time)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *GetTickets200ResponseDataInner) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### SetClosedNil

`func (o *GetTickets200ResponseDataInner) SetClosedNil(b bool)`

 SetClosedNil sets the value for Closed to be an explicit nil

### UnsetClosed
`func (o *GetTickets200ResponseDataInner) UnsetClosed()`

UnsetClosed ensures that no value is present for Closed, not even an explicit nil
### GetDescription

`func (o *GetTickets200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetTickets200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetTickets200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetTickets200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEntity

`func (o *GetTickets200ResponseDataInner) GetEntity() GetTickets200ResponseDataInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GetTickets200ResponseDataInner) GetEntityOk() (*GetTickets200ResponseDataInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GetTickets200ResponseDataInner) SetEntity(v GetTickets200ResponseDataInnerEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GetTickets200ResponseDataInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### SetEntityNil

`func (o *GetTickets200ResponseDataInner) SetEntityNil(b bool)`

 SetEntityNil sets the value for Entity to be an explicit nil

### UnsetEntity
`func (o *GetTickets200ResponseDataInner) UnsetEntity()`

UnsetEntity ensures that no value is present for Entity, not even an explicit nil
### GetGravatarId

`func (o *GetTickets200ResponseDataInner) GetGravatarId() string`

GetGravatarId returns the GravatarId field if non-nil, zero value otherwise.

### GetGravatarIdOk

`func (o *GetTickets200ResponseDataInner) GetGravatarIdOk() (*string, bool)`

GetGravatarIdOk returns a tuple with the GravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGravatarId

`func (o *GetTickets200ResponseDataInner) SetGravatarId(v string)`

SetGravatarId sets GravatarId field to given value.

### HasGravatarId

`func (o *GetTickets200ResponseDataInner) HasGravatarId() bool`

HasGravatarId returns a boolean if a field has been set.

### GetId

`func (o *GetTickets200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTickets200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTickets200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetTickets200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOpened

`func (o *GetTickets200ResponseDataInner) GetOpened() time.Time`

GetOpened returns the Opened field if non-nil, zero value otherwise.

### GetOpenedOk

`func (o *GetTickets200ResponseDataInner) GetOpenedOk() (*time.Time, bool)`

GetOpenedOk returns a tuple with the Opened field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpened

`func (o *GetTickets200ResponseDataInner) SetOpened(v time.Time)`

SetOpened sets Opened field to given value.

### HasOpened

`func (o *GetTickets200ResponseDataInner) HasOpened() bool`

HasOpened returns a boolean if a field has been set.

### GetOpenedBy

`func (o *GetTickets200ResponseDataInner) GetOpenedBy() string`

GetOpenedBy returns the OpenedBy field if non-nil, zero value otherwise.

### GetOpenedByOk

`func (o *GetTickets200ResponseDataInner) GetOpenedByOk() (*string, bool)`

GetOpenedByOk returns a tuple with the OpenedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenedBy

`func (o *GetTickets200ResponseDataInner) SetOpenedBy(v string)`

SetOpenedBy sets OpenedBy field to given value.

### HasOpenedBy

`func (o *GetTickets200ResponseDataInner) HasOpenedBy() bool`

HasOpenedBy returns a boolean if a field has been set.

### GetStatus

`func (o *GetTickets200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetTickets200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetTickets200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetTickets200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSummary

`func (o *GetTickets200ResponseDataInner) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetTickets200ResponseDataInner) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetTickets200ResponseDataInner) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetTickets200ResponseDataInner) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetUpdated

`func (o *GetTickets200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetTickets200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetTickets200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetTickets200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetTickets200ResponseDataInner) GetUpdatedBy() string`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetTickets200ResponseDataInner) GetUpdatedByOk() (*string, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetTickets200ResponseDataInner) SetUpdatedBy(v string)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetTickets200ResponseDataInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### SetUpdatedByNil

`func (o *GetTickets200ResponseDataInner) SetUpdatedByNil(b bool)`

 SetUpdatedByNil sets the value for UpdatedBy to be an explicit nil

### UnsetUpdatedBy
`func (o *GetTickets200ResponseDataInner) UnsetUpdatedBy()`

UnsetUpdatedBy ensures that no value is present for UpdatedBy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


