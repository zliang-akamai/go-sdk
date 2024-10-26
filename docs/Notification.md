# Notification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **NullableString** | A full description of this notification, in markdown format. Not all notifications include a &#x60;body&#x60;. | [optional] [readonly] 
**Entity** | Pointer to [**GetNotifications200ResponseDataInnerEntity**](GetNotifications200ResponseDataInnerEntity.md) |  | [optional] 
**Label** | Pointer to **string** | A short description of this notification. | [optional] [readonly] 
**Message** | Pointer to **string** | A human-readable description of the notification. | [optional] [readonly] 
**Severity** | Pointer to **string** | The severity of this notification. This field determines how prominently the notification is displayed and the color of the display text. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of notification. | [optional] [readonly] 
**Until** | Pointer to **time.Time** | If this notification has a duration, this is when the event or action will complete. For example, if there&#39;s scheduled maintenance for one of our systems, &#x60;until&#x60; represents the end of the maintenance window. | [optional] [readonly] 
**When** | Pointer to **time.Time** | If this notification is for an event in the future, this specifies when the action occurs. For example, if a compute instance needs to migrate in response to a security advisory, this field sets the approximate time the compute instance will be taken offline for migration. | [optional] [readonly] 

## Methods

### NewNotification

`func NewNotification() *Notification`

NewNotification instantiates a new Notification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationWithDefaults

`func NewNotificationWithDefaults() *Notification`

NewNotificationWithDefaults instantiates a new Notification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *Notification) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *Notification) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *Notification) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *Notification) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *Notification) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *Notification) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetEntity

`func (o *Notification) GetEntity() GetNotifications200ResponseDataInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *Notification) GetEntityOk() (*GetNotifications200ResponseDataInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *Notification) SetEntity(v GetNotifications200ResponseDataInnerEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *Notification) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetLabel

`func (o *Notification) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Notification) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Notification) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Notification) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMessage

`func (o *Notification) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Notification) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Notification) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Notification) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *Notification) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *Notification) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *Notification) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *Notification) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetType

`func (o *Notification) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Notification) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Notification) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Notification) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUntil

`func (o *Notification) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *Notification) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *Notification) SetUntil(v time.Time)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *Notification) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetWhen

`func (o *Notification) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *Notification) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *Notification) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *Notification) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


