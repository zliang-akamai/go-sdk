# GetSshKeys200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date this key was added. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique identifier of an SSH Key object. | [optional] [readonly] 
**Label** | Pointer to **string** | A label for the SSH Key. | [optional] 
**SshKey** | Pointer to **string** | The public SSH Key, which is used to authenticate to the root user of the Linodes you deploy.  Accepted formats:  - ssh-dss - ssh-rsa - ecdsa-sha2-nistp - ssh-ed25519 - sk-ecdsa-sha2-nistp256 (Akamai-specific) | [optional] 

## Methods

### NewGetSshKeys200ResponseDataInner

`func NewGetSshKeys200ResponseDataInner() *GetSshKeys200ResponseDataInner`

NewGetSshKeys200ResponseDataInner instantiates a new GetSshKeys200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSshKeys200ResponseDataInnerWithDefaults

`func NewGetSshKeys200ResponseDataInnerWithDefaults() *GetSshKeys200ResponseDataInner`

NewGetSshKeys200ResponseDataInnerWithDefaults instantiates a new GetSshKeys200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetSshKeys200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetSshKeys200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetSshKeys200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetSshKeys200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *GetSshKeys200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSshKeys200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSshKeys200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetSshKeys200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetSshKeys200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetSshKeys200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetSshKeys200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetSshKeys200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSshKey

`func (o *GetSshKeys200ResponseDataInner) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *GetSshKeys200ResponseDataInner) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *GetSshKeys200ResponseDataInner) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *GetSshKeys200ResponseDataInner) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


