# GetBackups200ResponseAutomaticInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Whether this Backup is available for restoration.  Backups undergoing maintenance are not available for restoration. | [optional] [readonly] 
**Configs** | Pointer to **[]string** | A list of the labels of the Configuration profiles that are part of the Backup. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The date the Backup was taken. | [optional] [readonly] 
**Disks** | Pointer to [**[]GetBackups200ResponseAutomaticInnerAllOfDisksInner**](GetBackups200ResponseAutomaticInnerAllOfDisksInner.md) | A list of the disks that are part of the Backup. | [optional] [readonly] 
**Finished** | Pointer to **time.Time** | The date the Backup completed. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID of this Backup. | [optional] [readonly] 
**Label** | Pointer to **NullableString** | A label for Backups that are of type &#x60;snapshot&#x60;. | [optional] 
**Status** | Pointer to **string** | The current state of a specific Backup. | [optional] [readonly] 
**Type** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **time.Time** | The date the Backup was most recently updated. | [optional] [readonly] 

## Methods

### NewGetBackups200ResponseAutomaticInner

`func NewGetBackups200ResponseAutomaticInner() *GetBackups200ResponseAutomaticInner`

NewGetBackups200ResponseAutomaticInner instantiates a new GetBackups200ResponseAutomaticInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBackups200ResponseAutomaticInnerWithDefaults

`func NewGetBackups200ResponseAutomaticInnerWithDefaults() *GetBackups200ResponseAutomaticInner`

NewGetBackups200ResponseAutomaticInnerWithDefaults instantiates a new GetBackups200ResponseAutomaticInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *GetBackups200ResponseAutomaticInner) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetBackups200ResponseAutomaticInner) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetBackups200ResponseAutomaticInner) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *GetBackups200ResponseAutomaticInner) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetConfigs

`func (o *GetBackups200ResponseAutomaticInner) GetConfigs() []string`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *GetBackups200ResponseAutomaticInner) GetConfigsOk() (*[]string, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *GetBackups200ResponseAutomaticInner) SetConfigs(v []string)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *GetBackups200ResponseAutomaticInner) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetCreated

`func (o *GetBackups200ResponseAutomaticInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetBackups200ResponseAutomaticInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetBackups200ResponseAutomaticInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetBackups200ResponseAutomaticInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDisks

`func (o *GetBackups200ResponseAutomaticInner) GetDisks() []GetBackups200ResponseAutomaticInnerAllOfDisksInner`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *GetBackups200ResponseAutomaticInner) GetDisksOk() (*[]GetBackups200ResponseAutomaticInnerAllOfDisksInner, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *GetBackups200ResponseAutomaticInner) SetDisks(v []GetBackups200ResponseAutomaticInnerAllOfDisksInner)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *GetBackups200ResponseAutomaticInner) HasDisks() bool`

HasDisks returns a boolean if a field has been set.

### GetFinished

`func (o *GetBackups200ResponseAutomaticInner) GetFinished() time.Time`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *GetBackups200ResponseAutomaticInner) GetFinishedOk() (*time.Time, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *GetBackups200ResponseAutomaticInner) SetFinished(v time.Time)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *GetBackups200ResponseAutomaticInner) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetId

`func (o *GetBackups200ResponseAutomaticInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBackups200ResponseAutomaticInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBackups200ResponseAutomaticInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetBackups200ResponseAutomaticInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetBackups200ResponseAutomaticInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetBackups200ResponseAutomaticInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetBackups200ResponseAutomaticInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetBackups200ResponseAutomaticInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *GetBackups200ResponseAutomaticInner) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *GetBackups200ResponseAutomaticInner) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStatus

`func (o *GetBackups200ResponseAutomaticInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBackups200ResponseAutomaticInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBackups200ResponseAutomaticInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetBackups200ResponseAutomaticInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GetBackups200ResponseAutomaticInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetBackups200ResponseAutomaticInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetBackups200ResponseAutomaticInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetBackups200ResponseAutomaticInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *GetBackups200ResponseAutomaticInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetBackups200ResponseAutomaticInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetBackups200ResponseAutomaticInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetBackups200ResponseAutomaticInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


