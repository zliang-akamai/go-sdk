# GetServiceTransfers200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this transfer was created. | [optional] 
**Entities** | Pointer to [**GetServiceTransfers200ResponseDataInnerEntities**](GetServiceTransfers200ResponseDataInnerEntities.md) |  | [optional] 
**Expiry** | Pointer to **time.Time** | When this transfer expires. Transfers will automatically expire 24 hours after creation. | [optional] 
**IsSender** | Pointer to **bool** | If the requesting account created this transfer. | [optional] 
**Status** | Pointer to **string** | The status of the transfer request.  &#x60;accepted&#x60;: The transfer has been accepted by another user and is currently in progress. Transfers can take up to 3 hours to complete.  &#x60;canceled&#x60;: The transfer has been canceled by the sender.  &#x60;completed&#x60;: The transfer has completed successfully.  &#x60;failed&#x60;: The transfer has failed after initiation.  &#x60;pending&#x60;: The transfer is ready to be accepted.  &#x60;stale&#x60;: The transfer has exceeded its expiration date. It can no longer be accepted or canceled. | [optional] 
**Token** | Pointer to **string** | The token used to identify and accept or cancel this transfer. | [optional] 
**Updated** | Pointer to **time.Time** | When this transfer was last updated. | [optional] 

## Methods

### NewGetServiceTransfers200ResponseDataInner

`func NewGetServiceTransfers200ResponseDataInner() *GetServiceTransfers200ResponseDataInner`

NewGetServiceTransfers200ResponseDataInner instantiates a new GetServiceTransfers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetServiceTransfers200ResponseDataInnerWithDefaults

`func NewGetServiceTransfers200ResponseDataInnerWithDefaults() *GetServiceTransfers200ResponseDataInner`

NewGetServiceTransfers200ResponseDataInnerWithDefaults instantiates a new GetServiceTransfers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetServiceTransfers200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetServiceTransfers200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetServiceTransfers200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetServiceTransfers200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEntities

`func (o *GetServiceTransfers200ResponseDataInner) GetEntities() GetServiceTransfers200ResponseDataInnerEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *GetServiceTransfers200ResponseDataInner) GetEntitiesOk() (*GetServiceTransfers200ResponseDataInnerEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *GetServiceTransfers200ResponseDataInner) SetEntities(v GetServiceTransfers200ResponseDataInnerEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *GetServiceTransfers200ResponseDataInner) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetExpiry

`func (o *GetServiceTransfers200ResponseDataInner) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GetServiceTransfers200ResponseDataInner) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GetServiceTransfers200ResponseDataInner) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GetServiceTransfers200ResponseDataInner) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetIsSender

`func (o *GetServiceTransfers200ResponseDataInner) GetIsSender() bool`

GetIsSender returns the IsSender field if non-nil, zero value otherwise.

### GetIsSenderOk

`func (o *GetServiceTransfers200ResponseDataInner) GetIsSenderOk() (*bool, bool)`

GetIsSenderOk returns a tuple with the IsSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSender

`func (o *GetServiceTransfers200ResponseDataInner) SetIsSender(v bool)`

SetIsSender sets IsSender field to given value.

### HasIsSender

`func (o *GetServiceTransfers200ResponseDataInner) HasIsSender() bool`

HasIsSender returns a boolean if a field has been set.

### GetStatus

`func (o *GetServiceTransfers200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetServiceTransfers200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetServiceTransfers200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetServiceTransfers200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *GetServiceTransfers200ResponseDataInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetServiceTransfers200ResponseDataInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetServiceTransfers200ResponseDataInner) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetServiceTransfers200ResponseDataInner) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdated

`func (o *GetServiceTransfers200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetServiceTransfers200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetServiceTransfers200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetServiceTransfers200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


