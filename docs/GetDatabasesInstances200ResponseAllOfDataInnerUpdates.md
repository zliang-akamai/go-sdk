# GetDatabasesInstances200ResponseAllOfDataInnerUpdates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **int32** | The day to perform maintenance. 1&#x3D;Monday, 2&#x3D;Tuesday, etc. | [optional] 
**Duration** | Pointer to **int32** | The maximum maintenance window time in hours. | [optional] 
**Frequency** | Pointer to **string** | Whether maintenance occurs on a weekly or monthly basis. | [optional] [default to "weekly"]
**HourOfDay** | Pointer to **int32** | The hour to begin maintenance based in UTC time. | [optional] 
**WeekOfMonth** | Pointer to **NullableInt32** | The week of the month to perform &#x60;monthly&#x60; frequency updates. Defaults to &#x60;null&#x60;.  - Required for &#x60;monthly&#x60; frequency updates.  - Must be &#x60;null&#x60; for &#x60;weekly&#x60; frequency updates. | [optional] 

## Methods

### NewGetDatabasesInstances200ResponseAllOfDataInnerUpdates

`func NewGetDatabasesInstances200ResponseAllOfDataInnerUpdates() *GetDatabasesInstances200ResponseAllOfDataInnerUpdates`

NewGetDatabasesInstances200ResponseAllOfDataInnerUpdates instantiates a new GetDatabasesInstances200ResponseAllOfDataInnerUpdates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasesInstances200ResponseAllOfDataInnerUpdatesWithDefaults

`func NewGetDatabasesInstances200ResponseAllOfDataInnerUpdatesWithDefaults() *GetDatabasesInstances200ResponseAllOfDataInnerUpdates`

NewGetDatabasesInstances200ResponseAllOfDataInnerUpdatesWithDefaults instantiates a new GetDatabasesInstances200ResponseAllOfDataInnerUpdates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetDuration

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFrequency

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetFrequency() string`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetFrequencyOk() (*string, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) SetFrequency(v string)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetHourOfDay

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetHourOfDay() int32`

GetHourOfDay returns the HourOfDay field if non-nil, zero value otherwise.

### GetHourOfDayOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetHourOfDayOk() (*int32, bool)`

GetHourOfDayOk returns a tuple with the HourOfDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHourOfDay

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) SetHourOfDay(v int32)`

SetHourOfDay sets HourOfDay field to given value.

### HasHourOfDay

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) HasHourOfDay() bool`

HasHourOfDay returns a boolean if a field has been set.

### GetWeekOfMonth

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetWeekOfMonth() int32`

GetWeekOfMonth returns the WeekOfMonth field if non-nil, zero value otherwise.

### GetWeekOfMonthOk

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) GetWeekOfMonthOk() (*int32, bool)`

GetWeekOfMonthOk returns a tuple with the WeekOfMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekOfMonth

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) SetWeekOfMonth(v int32)`

SetWeekOfMonth sets WeekOfMonth field to given value.

### HasWeekOfMonth

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) HasWeekOfMonth() bool`

HasWeekOfMonth returns a boolean if a field has been set.

### SetWeekOfMonthNil

`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) SetWeekOfMonthNil(b bool)`

 SetWeekOfMonthNil sets the value for WeekOfMonth to be an explicit nil

### UnsetWeekOfMonth
`func (o *GetDatabasesInstances200ResponseAllOfDataInnerUpdates) UnsetWeekOfMonth()`

UnsetWeekOfMonth ensures that no value is present for WeekOfMonth, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


