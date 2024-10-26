# GetManagedIssues200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this Issue was created. Issues are created in response to issues detected with Managed Services, so this is also when the Issue was detected. | [optional] [readonly] 
**Entity** | Pointer to [**GetManagedIssues200ResponseDataInnerEntity**](GetManagedIssues200ResponseDataInnerEntity.md) |  | [optional] 
**Id** | Pointer to **int32** | This Issue&#39;s unique ID. | [optional] [readonly] 
**Services** | Pointer to **[]int32** | An array of Managed Service IDs that were affected by this Issue. | [optional] [readonly] 

## Methods

### NewGetManagedIssues200ResponseDataInner

`func NewGetManagedIssues200ResponseDataInner() *GetManagedIssues200ResponseDataInner`

NewGetManagedIssues200ResponseDataInner instantiates a new GetManagedIssues200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedIssues200ResponseDataInnerWithDefaults

`func NewGetManagedIssues200ResponseDataInnerWithDefaults() *GetManagedIssues200ResponseDataInner`

NewGetManagedIssues200ResponseDataInnerWithDefaults instantiates a new GetManagedIssues200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetManagedIssues200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetManagedIssues200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetManagedIssues200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetManagedIssues200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetEntity

`func (o *GetManagedIssues200ResponseDataInner) GetEntity() GetManagedIssues200ResponseDataInnerEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GetManagedIssues200ResponseDataInner) GetEntityOk() (*GetManagedIssues200ResponseDataInnerEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GetManagedIssues200ResponseDataInner) SetEntity(v GetManagedIssues200ResponseDataInnerEntity)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GetManagedIssues200ResponseDataInner) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetId

`func (o *GetManagedIssues200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetManagedIssues200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetManagedIssues200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetManagedIssues200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetServices

`func (o *GetManagedIssues200ResponseDataInner) GetServices() []int32`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GetManagedIssues200ResponseDataInner) GetServicesOk() (*[]int32, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GetManagedIssues200ResponseDataInner) SetServices(v []int32)`

SetServices sets Services field to given value.

### HasServices

`func (o *GetManagedIssues200ResponseDataInner) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


