# LinodeBackups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Whether Backups for this Linode are available for restoration.  Backups undergoing maintenance are not available for restoration. | [optional] [readonly] 
**Enabled** | Pointer to **bool** | If this Linode has the Backup service enabled. To enable backups, run [Enable backups](https://techdocs.akamai.com/linode-api/reference/post-enable-backups). | [optional] [readonly] 
**LastSuccessful** | Pointer to **time.Time** | The last successful backup time. Displayed as &#x60;null&#x60; if there was no previous backup. | [optional] [readonly] 
**Schedule** | Pointer to [**LinodeBackupsSchedule**](LinodeBackupsSchedule.md) |  | [optional] 

## Methods

### NewLinodeBackups

`func NewLinodeBackups() *LinodeBackups`

NewLinodeBackups instantiates a new LinodeBackups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeBackupsWithDefaults

`func NewLinodeBackupsWithDefaults() *LinodeBackups`

NewLinodeBackupsWithDefaults instantiates a new LinodeBackups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *LinodeBackups) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *LinodeBackups) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *LinodeBackups) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *LinodeBackups) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetEnabled

`func (o *LinodeBackups) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *LinodeBackups) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *LinodeBackups) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *LinodeBackups) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetLastSuccessful

`func (o *LinodeBackups) GetLastSuccessful() time.Time`

GetLastSuccessful returns the LastSuccessful field if non-nil, zero value otherwise.

### GetLastSuccessfulOk

`func (o *LinodeBackups) GetLastSuccessfulOk() (*time.Time, bool)`

GetLastSuccessfulOk returns a tuple with the LastSuccessful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSuccessful

`func (o *LinodeBackups) SetLastSuccessful(v time.Time)`

SetLastSuccessful sets LastSuccessful field to given value.

### HasLastSuccessful

`func (o *LinodeBackups) HasLastSuccessful() bool`

HasLastSuccessful returns a boolean if a field has been set.

### GetSchedule

`func (o *LinodeBackups) GetSchedule() LinodeBackupsSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *LinodeBackups) GetScheduleOk() (*LinodeBackupsSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *LinodeBackups) SetSchedule(v LinodeBackupsSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *LinodeBackups) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


