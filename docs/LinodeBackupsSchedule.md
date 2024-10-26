# LinodeBackupsSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **NullableString** | The day of the week that your Linode&#39;s weekly backup is taken. If not set manually, a day will be chosen for you. Backups are taken every day, but backups taken on this day are preferred when selecting backups to retain for a longer period.  If not set manually, then when backups are initially enabled, this may come back as &#x60;Scheduling&#x60; until the &#x60;day&#x60; is automatically selected. | [optional] 
**Window** | Pointer to **NullableString** | When your backups will be taken, in UTC. A backups window is a two-hour span of time in which the backup may occur.  For example, &#x60;W10&#x60; indicates that your backups should be taken between 10:00 and 12:00. If you don&#39;t choose a backup window, the API automatically assigns one.  If not set manually, when backups are initially enabled this may come back as &#x60;Scheduling&#x60; until the &#x60;window&#x60; is automatically selected. | [optional] 

## Methods

### NewLinodeBackupsSchedule

`func NewLinodeBackupsSchedule() *LinodeBackupsSchedule`

NewLinodeBackupsSchedule instantiates a new LinodeBackupsSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinodeBackupsScheduleWithDefaults

`func NewLinodeBackupsScheduleWithDefaults() *LinodeBackupsSchedule`

NewLinodeBackupsScheduleWithDefaults instantiates a new LinodeBackupsSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *LinodeBackupsSchedule) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *LinodeBackupsSchedule) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *LinodeBackupsSchedule) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *LinodeBackupsSchedule) HasDay() bool`

HasDay returns a boolean if a field has been set.

### SetDayNil

`func (o *LinodeBackupsSchedule) SetDayNil(b bool)`

 SetDayNil sets the value for Day to be an explicit nil

### UnsetDay
`func (o *LinodeBackupsSchedule) UnsetDay()`

UnsetDay ensures that no value is present for Day, not even an explicit nil
### GetWindow

`func (o *LinodeBackupsSchedule) GetWindow() string`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *LinodeBackupsSchedule) GetWindowOk() (*string, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *LinodeBackupsSchedule) SetWindow(v string)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *LinodeBackupsSchedule) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### SetWindowNil

`func (o *LinodeBackupsSchedule) SetWindowNil(b bool)`

 SetWindowNil sets the value for Window to be an explicit nil

### UnsetWindow
`func (o *LinodeBackupsSchedule) UnsetWindow()`

UnsetWindow ensures that no value is present for Window, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


