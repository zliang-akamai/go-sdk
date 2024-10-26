# GetMaintenance200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to [**GetMaintenance200ResponseDataInnerEntity**](GetMaintenance200ResponseDataInnerEntity.md) |  | [optional] 
**Reason** | Pointer to **string** | The reason maintenance is being performed. | [optional] 
**Status** | Pointer to **string** | The maintenance status.  Maintenance progresses in the following sequence: pending, started, then completed. | [optional] 
**Type** | Pointer to **string** | The type of maintenance. | [optional] 
**When** | Pointer to **time.Time** | When the maintenance will begin.  [Filterable](https://techdocs.akamai.com/linode-api/reference/filtering-and-sorting) with the following parameters:  - A single value in date-time string format (&#x60;%Y-%m-%dT%H:%M:%S&#x60;), which returns only matches to that value.  - A dictionary containing pairs of inequality operator string keys (&#x60;+or&#x60;, &#x60;+gt&#x60;, &#x60;+gte&#x60;, &#x60;+lt&#x60;, &#x60;+lte&#x60;, or &#x60;+neq&#x60;) and single date-time string format values (&#x60;%Y-%m-%dT%H:%M:%S&#x60;). &#x60;+or&#x60; accepts an array of values that may consist of single date-time strings or dictionaries of inequality operator pairs. | [optional] 

## Methods

### NewGetMaintenance200ResponseDataInner

`func NewGetMaintenance200ResponseDataInner() *GetMaintenance200ResponseDataInner`

NewGetMaintenance200ResponseDataInner instantiates a new GetMaintenance200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMaintenance200ResponseDataInnerWithDefaults

`func NewGetMaintenance200ResponseDataInnerWithDefaults() *GetMaintenance200ResponseDataInner`

NewGetMaintenance200ResponseDataInnerWithDefaults instantiates a new GetMaintenance200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *GetMaintenance200ResponseDataInner) GetEntity() GetMaintenance200ResponseDataInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GetMaintenance200ResponseDataInner) GetEntityOk() (*GetMaintenance200ResponseDataInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GetMaintenance200ResponseDataInner) SetEntity(v GetMaintenance200ResponseDataInnerEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GetMaintenance200ResponseDataInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetReason

`func (o *GetMaintenance200ResponseDataInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetMaintenance200ResponseDataInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetMaintenance200ResponseDataInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetMaintenance200ResponseDataInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *GetMaintenance200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMaintenance200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMaintenance200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetMaintenance200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *GetMaintenance200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetMaintenance200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetMaintenance200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetMaintenance200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWhen

`func (o *GetMaintenance200ResponseDataInner) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *GetMaintenance200ResponseDataInner) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *GetMaintenance200ResponseDataInner) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *GetMaintenance200ResponseDataInner) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


