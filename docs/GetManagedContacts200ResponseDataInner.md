# GetManagedContacts200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | The address to email this Contact to alert them of issues. | [optional] 
**Group** | Pointer to **NullableString** | A grouping for this Contact. This is for display purposes only. | [optional] 
**Id** | Pointer to **int32** | This Contact&#39;s unique ID. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of this Contact. | [optional] 
**Phone** | Pointer to [**GetManagedContacts200ResponseDataInnerPhone**](GetManagedContacts200ResponseDataInnerPhone.md) |  | [optional] 
**Updated** | Pointer to **time.Time** | When this Contact was last updated. | [optional] [readonly] 

## Methods

### NewGetManagedContacts200ResponseDataInner

`func NewGetManagedContacts200ResponseDataInner() *GetManagedContacts200ResponseDataInner`

NewGetManagedContacts200ResponseDataInner instantiates a new GetManagedContacts200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedContacts200ResponseDataInnerWithDefaults

`func NewGetManagedContacts200ResponseDataInnerWithDefaults() *GetManagedContacts200ResponseDataInner`

NewGetManagedContacts200ResponseDataInnerWithDefaults instantiates a new GetManagedContacts200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GetManagedContacts200ResponseDataInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetManagedContacts200ResponseDataInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetManagedContacts200ResponseDataInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetManagedContacts200ResponseDataInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetGroup

`func (o *GetManagedContacts200ResponseDataInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetManagedContacts200ResponseDataInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetManagedContacts200ResponseDataInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetManagedContacts200ResponseDataInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### SetGroupNil

`func (o *GetManagedContacts200ResponseDataInner) SetGroupNil(b bool)`

 SetGroupNil sets the value for Group to be an explicit nil

### UnsetGroup
`func (o *GetManagedContacts200ResponseDataInner) UnsetGroup()`

UnsetGroup ensures that no value is present for Group, not even an explicit nil
### GetId

`func (o *GetManagedContacts200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetManagedContacts200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetManagedContacts200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetManagedContacts200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetManagedContacts200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetManagedContacts200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetManagedContacts200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetManagedContacts200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhone

`func (o *GetManagedContacts200ResponseDataInner) GetPhone() GetManagedContacts200ResponseDataInnerPhone`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetManagedContacts200ResponseDataInner) GetPhoneOk() (*GetManagedContacts200ResponseDataInnerPhone, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetManagedContacts200ResponseDataInner) SetPhone(v GetManagedContacts200ResponseDataInnerPhone)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetManagedContacts200ResponseDataInner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetUpdated

`func (o *GetManagedContacts200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetManagedContacts200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetManagedContacts200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetManagedContacts200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


