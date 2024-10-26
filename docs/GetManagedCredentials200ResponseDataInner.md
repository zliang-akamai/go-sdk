# GetManagedCredentials200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | This Credential&#39;s unique ID. | [optional] [readonly] 
**Label** | Pointer to **string** | The unique label for this Credential. This is for display purposes only. | [optional] 
**LastDecrypted** | Pointer to **time.Time** | The date this Credential was last decrypted by a member of Linode special forces. | [optional] [readonly] 

## Methods

### NewGetManagedCredentials200ResponseDataInner

`func NewGetManagedCredentials200ResponseDataInner() *GetManagedCredentials200ResponseDataInner`

NewGetManagedCredentials200ResponseDataInner instantiates a new GetManagedCredentials200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedCredentials200ResponseDataInnerWithDefaults

`func NewGetManagedCredentials200ResponseDataInnerWithDefaults() *GetManagedCredentials200ResponseDataInner`

NewGetManagedCredentials200ResponseDataInnerWithDefaults instantiates a new GetManagedCredentials200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetManagedCredentials200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetManagedCredentials200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetManagedCredentials200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetManagedCredentials200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetManagedCredentials200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetManagedCredentials200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetManagedCredentials200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetManagedCredentials200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastDecrypted

`func (o *GetManagedCredentials200ResponseDataInner) GetLastDecrypted() time.Time`

GetLastDecrypted returns the LastDecrypted field if non-nil, zero value otherwise.

### GetLastDecryptedOk

`func (o *GetManagedCredentials200ResponseDataInner) GetLastDecryptedOk() (*time.Time, bool)`

GetLastDecryptedOk returns a tuple with the LastDecrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDecrypted

`func (o *GetManagedCredentials200ResponseDataInner) SetLastDecrypted(v time.Time)`

SetLastDecrypted sets LastDecrypted field to given value.

### HasLastDecrypted

`func (o *GetManagedCredentials200ResponseDataInner) HasLastDecrypted() bool`

HasLastDecrypted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


