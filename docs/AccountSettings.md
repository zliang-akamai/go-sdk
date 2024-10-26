# AccountSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupsEnabled** | Pointer to **bool** | Account-wide backups default.  If &#x60;true&#x60;, all Linodes created will automatically be enrolled in the Backups service.  If &#x60;false&#x60;, Linodes will not be enrolled by default, but may still be enrolled on creation or later. | [optional] 
**LongviewSubscription** | Pointer to **string** | The Longview Pro tier you are currently subscribed to. The value must be a [Longview subscription](https://techdocs.akamai.com/linode-api/reference/get-longview-subscriptions) ID or &#x60;null&#x60; for Longview Free. | [optional] [readonly] 
**Managed** | Pointer to **bool** | Our 24/7 incident response service. This robust, multi-homed monitoring system distributes monitoring checks to ensure that your servers remain online and available at all times. Linode Managed can monitor any service or software stack reachable over TCP or HTTP. Once you add a service to Linode Managed, we&#39;ll monitor it for connectivity, response, and total request time. | [optional] [readonly] 
**NetworkHelper** | Pointer to **bool** | Enables network helper across all users by default for new Linodes and Linode Configs. | [optional] 
**ObjectStorage** | Pointer to **string** | A string describing the status of this account&#39;s Object Storage service enrollment. | [optional] [readonly] [default to "disabled"]

## Methods

### NewAccountSettings

`func NewAccountSettings() *AccountSettings`

NewAccountSettings instantiates a new AccountSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountSettingsWithDefaults

`func NewAccountSettingsWithDefaults() *AccountSettings`

NewAccountSettingsWithDefaults instantiates a new AccountSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupsEnabled

`func (o *AccountSettings) GetBackupsEnabled() bool`

GetBackupsEnabled returns the BackupsEnabled field if non-nil, zero value otherwise.

### GetBackupsEnabledOk

`func (o *AccountSettings) GetBackupsEnabledOk() (*bool, bool)`

GetBackupsEnabledOk returns a tuple with the BackupsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupsEnabled

`func (o *AccountSettings) SetBackupsEnabled(v bool)`

SetBackupsEnabled sets BackupsEnabled field to given value.

### HasBackupsEnabled

`func (o *AccountSettings) HasBackupsEnabled() bool`

HasBackupsEnabled returns a boolean if a field has been set.

### GetLongviewSubscription

`func (o *AccountSettings) GetLongviewSubscription() string`

GetLongviewSubscription returns the LongviewSubscription field if non-nil, zero value otherwise.

### GetLongviewSubscriptionOk

`func (o *AccountSettings) GetLongviewSubscriptionOk() (*string, bool)`

GetLongviewSubscriptionOk returns a tuple with the LongviewSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongviewSubscription

`func (o *AccountSettings) SetLongviewSubscription(v string)`

SetLongviewSubscription sets LongviewSubscription field to given value.

### HasLongviewSubscription

`func (o *AccountSettings) HasLongviewSubscription() bool`

HasLongviewSubscription returns a boolean if a field has been set.

### GetManaged

`func (o *AccountSettings) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *AccountSettings) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *AccountSettings) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *AccountSettings) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetNetworkHelper

`func (o *AccountSettings) GetNetworkHelper() bool`

GetNetworkHelper returns the NetworkHelper field if non-nil, zero value otherwise.

### GetNetworkHelperOk

`func (o *AccountSettings) GetNetworkHelperOk() (*bool, bool)`

GetNetworkHelperOk returns a tuple with the NetworkHelper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkHelper

`func (o *AccountSettings) SetNetworkHelper(v bool)`

SetNetworkHelper sets NetworkHelper field to given value.

### HasNetworkHelper

`func (o *AccountSettings) HasNetworkHelper() bool`

HasNetworkHelper returns a boolean if a field has been set.

### GetObjectStorage

`func (o *AccountSettings) GetObjectStorage() string`

GetObjectStorage returns the ObjectStorage field if non-nil, zero value otherwise.

### GetObjectStorageOk

`func (o *AccountSettings) GetObjectStorageOk() (*string, bool)`

GetObjectStorageOk returns a tuple with the ObjectStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStorage

`func (o *AccountSettings) SetObjectStorage(v string)`

SetObjectStorage sets ObjectStorage field to given value.

### HasObjectStorage

`func (o *AccountSettings) HasObjectStorage() bool`

HasObjectStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


