# GetBackups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Automatic** | Pointer to [**[]GetBackups200ResponseAutomaticInner**](GetBackups200ResponseAutomaticInner.md) |  | [optional] 
**Snapshot** | Pointer to [**GetBackups200ResponseSnapshot**](GetBackups200ResponseSnapshot.md) |  | [optional] 

## Methods

### NewGetBackups200Response

`func NewGetBackups200Response() *GetBackups200Response`

NewGetBackups200Response instantiates a new GetBackups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBackups200ResponseWithDefaults

`func NewGetBackups200ResponseWithDefaults() *GetBackups200Response`

NewGetBackups200ResponseWithDefaults instantiates a new GetBackups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutomatic

`func (o *GetBackups200Response) GetAutomatic() []GetBackups200ResponseAutomaticInner`

GetAutomatic returns the Automatic field if non-nil, zero value otherwise.

### GetAutomaticOk

`func (o *GetBackups200Response) GetAutomaticOk() (*[]GetBackups200ResponseAutomaticInner, bool)`

GetAutomaticOk returns a tuple with the Automatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatic

`func (o *GetBackups200Response) SetAutomatic(v []GetBackups200ResponseAutomaticInner)`

SetAutomatic sets Automatic field to given value.

### HasAutomatic

`func (o *GetBackups200Response) HasAutomatic() bool`

HasAutomatic returns a boolean if a field has been set.

### GetSnapshot

`func (o *GetBackups200Response) GetSnapshot() GetBackups200ResponseSnapshot`

GetSnapshot returns the Snapshot field if non-nil, zero value otherwise.

### GetSnapshotOk

`func (o *GetBackups200Response) GetSnapshotOk() (*GetBackups200ResponseSnapshot, bool)`

GetSnapshotOk returns a tuple with the Snapshot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshot

`func (o *GetBackups200Response) SetSnapshot(v GetBackups200ResponseSnapshot)`

SetSnapshot sets Snapshot field to given value.

### HasSnapshot

`func (o *GetBackups200Response) HasSnapshot() bool`

HasSnapshot returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


