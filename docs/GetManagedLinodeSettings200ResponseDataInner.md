# GetManagedLinodeSettings200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **string** | The group of the Linode these Settings are for. This is for display purposes only. | [optional] [readonly] 
**Id** | Pointer to **int32** | The ID of the Linode these Settings are for. | [optional] [readonly] 
**Label** | Pointer to **string** | The label of the Linode these Settings are for. | [optional] [readonly] 
**Ssh** | Pointer to [**GetManagedLinodeSettings200ResponseDataInnerSsh**](GetManagedLinodeSettings200ResponseDataInnerSsh.md) |  | [optional] 

## Methods

### NewGetManagedLinodeSettings200ResponseDataInner

`func NewGetManagedLinodeSettings200ResponseDataInner() *GetManagedLinodeSettings200ResponseDataInner`

NewGetManagedLinodeSettings200ResponseDataInner instantiates a new GetManagedLinodeSettings200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedLinodeSettings200ResponseDataInnerWithDefaults

`func NewGetManagedLinodeSettings200ResponseDataInnerWithDefaults() *GetManagedLinodeSettings200ResponseDataInner`

NewGetManagedLinodeSettings200ResponseDataInnerWithDefaults instantiates a new GetManagedLinodeSettings200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *GetManagedLinodeSettings200ResponseDataInner) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetManagedLinodeSettings200ResponseDataInner) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetManagedLinodeSettings200ResponseDataInner) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetManagedLinodeSettings200ResponseDataInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *GetManagedLinodeSettings200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetManagedLinodeSettings200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetManagedLinodeSettings200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetManagedLinodeSettings200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetManagedLinodeSettings200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetManagedLinodeSettings200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetManagedLinodeSettings200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetManagedLinodeSettings200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSsh

`func (o *GetManagedLinodeSettings200ResponseDataInner) GetSsh() GetManagedLinodeSettings200ResponseDataInnerSsh`

GetSsh returns the Ssh field if non-nil, zero value otherwise.

### GetSshOk

`func (o *GetManagedLinodeSettings200ResponseDataInner) GetSshOk() (*GetManagedLinodeSettings200ResponseDataInnerSsh, bool)`

GetSshOk returns a tuple with the Ssh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsh

`func (o *GetManagedLinodeSettings200ResponseDataInner) SetSsh(v GetManagedLinodeSettings200ResponseDataInnerSsh)`

SetSsh sets Ssh field to given value.

### HasSsh

`func (o *GetManagedLinodeSettings200ResponseDataInner) HasSsh() bool`

HasSsh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


