# ServiceTransfer

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

### NewServiceTransfer

`func NewServiceTransfer() *ServiceTransfer`

NewServiceTransfer instantiates a new ServiceTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceTransferWithDefaults

`func NewServiceTransferWithDefaults() *ServiceTransfer`

NewServiceTransferWithDefaults instantiates a new ServiceTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ServiceTransfer) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ServiceTransfer) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ServiceTransfer) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ServiceTransfer) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEntities

`func (o *ServiceTransfer) GetEntities() GetServiceTransfers200ResponseDataInnerEntities`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *ServiceTransfer) GetEntitiesOk() (*GetServiceTransfers200ResponseDataInnerEntities, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *ServiceTransfer) SetEntities(v GetServiceTransfers200ResponseDataInnerEntities)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *ServiceTransfer) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetExpiry

`func (o *ServiceTransfer) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *ServiceTransfer) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *ServiceTransfer) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *ServiceTransfer) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetIsSender

`func (o *ServiceTransfer) GetIsSender() bool`

GetIsSender returns the IsSender field if non-nil, zero value otherwise.

### GetIsSenderOk

`func (o *ServiceTransfer) GetIsSenderOk() (*bool, bool)`

GetIsSenderOk returns a tuple with the IsSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSender

`func (o *ServiceTransfer) SetIsSender(v bool)`

SetIsSender sets IsSender field to given value.

### HasIsSender

`func (o *ServiceTransfer) HasIsSender() bool`

HasIsSender returns a boolean if a field has been set.

### GetStatus

`func (o *ServiceTransfer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ServiceTransfer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ServiceTransfer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ServiceTransfer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToken

`func (o *ServiceTransfer) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ServiceTransfer) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ServiceTransfer) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ServiceTransfer) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUpdated

`func (o *ServiceTransfer) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ServiceTransfer) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ServiceTransfer) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ServiceTransfer) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


