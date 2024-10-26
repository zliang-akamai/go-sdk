# GetNotifications200ResponseDataInner

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

### NewGetNotifications200ResponseDataInner

`func NewGetNotifications200ResponseDataInner() *GetNotifications200ResponseDataInner`

NewGetNotifications200ResponseDataInner instantiates a new GetNotifications200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotifications200ResponseDataInnerWithDefaults

`func NewGetNotifications200ResponseDataInnerWithDefaults() *GetNotifications200ResponseDataInner`

NewGetNotifications200ResponseDataInnerWithDefaults instantiates a new GetNotifications200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *GetNotifications200ResponseDataInner) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetNotifications200ResponseDataInner) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetNotifications200ResponseDataInner) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GetNotifications200ResponseDataInner) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *GetNotifications200ResponseDataInner) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *GetNotifications200ResponseDataInner) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetEntity

`func (o *GetNotifications200ResponseDataInner) GetEntity() GetNotifications200ResponseDataInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GetNotifications200ResponseDataInner) GetEntityOk() (*GetNotifications200ResponseDataInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GetNotifications200ResponseDataInner) SetEntity(v GetNotifications200ResponseDataInnerEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GetNotifications200ResponseDataInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetLabel

`func (o *GetNotifications200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetNotifications200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetNotifications200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetNotifications200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMessage

`func (o *GetNotifications200ResponseDataInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetNotifications200ResponseDataInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetNotifications200ResponseDataInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetNotifications200ResponseDataInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *GetNotifications200ResponseDataInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetNotifications200ResponseDataInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetNotifications200ResponseDataInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetNotifications200ResponseDataInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetType

`func (o *GetNotifications200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNotifications200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNotifications200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNotifications200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUntil

`func (o *GetNotifications200ResponseDataInner) GetUntil() time.Time`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *GetNotifications200ResponseDataInner) GetUntilOk() (*time.Time, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *GetNotifications200ResponseDataInner) SetUntil(v time.Time)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *GetNotifications200ResponseDataInner) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetWhen

`func (o *GetNotifications200ResponseDataInner) GetWhen() time.Time`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *GetNotifications200ResponseDataInner) GetWhenOk() (*time.Time, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *GetNotifications200ResponseDataInner) SetWhen(v time.Time)`

SetWhen sets When field to given value.

### HasWhen

`func (o *GetNotifications200ResponseDataInner) HasWhen() bool`

HasWhen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


