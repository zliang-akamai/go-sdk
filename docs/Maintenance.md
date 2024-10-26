# Maintenance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**GetMaintenance200ResponseDataInnerEntity**](GetMaintenance200ResponseDataInnerEntity.md) |  | [optional] 
**Reason** | Pointer to **string** | The reason maintenance is being performed. | [optional] 
**Status** | Pointer to **string** | The maintenance status.  Maintenance progresses in the following sequence: pending, started, then completed. | [optional] 
**Type** | Pointer to **string** | The type of maintenance. | [optional] 
**When** | Pointer to **time.Time** | When the maintenance will begin.  [Filterable](https://techdocs.akamai.com/linode-api/reference/filtering-and-sorting) with the following parameters:  - A single value in date-time string format (&#x60;%Y-%m-%dT%H:%M:%S&#x60;), which returns only matches to that value.  - A dictionary containing pairs of inequality operator string keys (&#x60;+or&#x60;, &#x60;+gt&#x60;, &#x60;+gte&#x60;, &#x60;+lt&#x60;, &#x60;+lte&#x60;, or &#x60;+neq&#x60;) and single date-time string format values (&#x60;%Y-%m-%dT%H:%M:%S&#x60;). &#x60;+or&#x60; accepts an array of values that may consist of single date-time strings or dictionaries of inequality operator pairs. | [optional] 

## Methods

### NewMaintenance

`func NewMaintenance() *Maintenance`

NewMaintenance instantiates a new Maintenance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceWithDefaults

`func NewMaintenanceWithDefaults() *Maintenance`

NewMaintenanceWithDefaults instantiates a new Maintenance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *Maintenance) GetEntity() GetMaintenance200ResponseDataInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Maintenance) GetEntityOk() (*GetMaintenance200ResponseDataInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Maintenance) SetEntity(v GetMaintenance200ResponseDataInnerEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *Maintenance) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetReason

`func (o *Maintenance) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Maintenance) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Maintenance) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Maintenance) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *Maintenance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Maintenance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Maintenance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Maintenance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Maintenance) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Maintenance) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Maintenance) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Maintenance) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWhen

`func (o *Maintenance) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *Maintenance) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *Maintenance) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *Maintenance) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


