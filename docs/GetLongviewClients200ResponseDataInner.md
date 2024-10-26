# GetLongviewClients200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | Pointer to **string** | The API key for this Client, used when configuring the Longview Client application on your Linode.  Returns as &#x60;[REDACTED]&#x60; if you do not have read-write access to this client. | [optional] [readonly] 
**Apps** | Pointer to [**GetLongviewClients200ResponseDataInnerApps**](GetLongviewClients200ResponseDataInnerApps.md) |  | [optional] 
**Created** | Pointer to **time.Time** | When this Longview Client was created. | [optional] [readonly] 
**Id** | Pointer to **int32** | This Client&#39;s unique ID. | [optional] [readonly] 
**InstallCode** | Pointer to **string** | The install code for this Client, used when configuring the Longview Client application on your Linode.  Returns as &#x60;[REDACTED]&#x60; if you do not have read-write access to this client. | [optional] [readonly] 
**Label** | Pointer to **string** | This Client&#39;s unique label. This is for display purposes only. | [optional] 
**Updated** | Pointer to **time.Time** | When this Longview Client was last updated. | [optional] [readonly] 

## Methods

### NewGetLongviewClients200ResponseDataInner

`func NewGetLongviewClients200ResponseDataInner() *GetLongviewClients200ResponseDataInner`

NewGetLongviewClients200ResponseDataInner instantiates a new GetLongviewClients200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLongviewClients200ResponseDataInnerWithDefaults

`func NewGetLongviewClients200ResponseDataInnerWithDefaults() *GetLongviewClients200ResponseDataInner`

NewGetLongviewClients200ResponseDataInnerWithDefaults instantiates a new GetLongviewClients200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *GetLongviewClients200ResponseDataInner) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *GetLongviewClients200ResponseDataInner) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *GetLongviewClients200ResponseDataInner) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *GetLongviewClients200ResponseDataInner) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetApps

`func (o *GetLongviewClients200ResponseDataInner) GetApps() GetLongviewClients200ResponseDataInnerApps`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *GetLongviewClients200ResponseDataInner) GetAppsOk() (*GetLongviewClients200ResponseDataInnerApps, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *GetLongviewClients200ResponseDataInner) SetApps(v GetLongviewClients200ResponseDataInnerApps)`

SetApps sets Apps field to given value.

### HasApps

`func (o *GetLongviewClients200ResponseDataInner) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetCreated

`func (o *GetLongviewClients200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetLongviewClients200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetLongviewClients200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetLongviewClients200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *GetLongviewClients200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLongviewClients200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLongviewClients200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetLongviewClients200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstallCode

`func (o *GetLongviewClients200ResponseDataInner) GetInstallCode() string`

GetInstallCode returns the InstallCode field if non-nil, zero value otherwise.

### GetInstallCodeOk

`func (o *GetLongviewClients200ResponseDataInner) GetInstallCodeOk() (*string, bool)`

GetInstallCodeOk returns a tuple with the InstallCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallCode

`func (o *GetLongviewClients200ResponseDataInner) SetInstallCode(v string)`

SetInstallCode sets InstallCode field to given value.

### HasInstallCode

`func (o *GetLongviewClients200ResponseDataInner) HasInstallCode() bool`

HasInstallCode returns a boolean if a field has been set.

### GetLabel

`func (o *GetLongviewClients200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetLongviewClients200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetLongviewClients200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetLongviewClients200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetUpdated

`func (o *GetLongviewClients200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetLongviewClients200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetLongviewClients200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetLongviewClients200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


