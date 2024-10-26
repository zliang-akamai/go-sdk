# GetManagedSshKey200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SshKey** | Pointer to **string** | The unique SSH public key assigned to your Linode account&#39;s Managed service. | [optional] [readonly] 

## Methods

### NewGetManagedSshKey200Response

`func NewGetManagedSshKey200Response() *GetManagedSshKey200Response`

NewGetManagedSshKey200Response instantiates a new GetManagedSshKey200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedSshKey200ResponseWithDefaults

`func NewGetManagedSshKey200ResponseWithDefaults() *GetManagedSshKey200Response`

NewGetManagedSshKey200ResponseWithDefaults instantiates a new GetManagedSshKey200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSshKey

`func (o *GetManagedSshKey200Response) GetSshKey() string`

GetSshKey returns the SshKey field if non-nil, zero value otherwise.

### GetSshKeyOk

`func (o *GetManagedSshKey200Response) GetSshKeyOk() (*string, bool)`

GetSshKeyOk returns a tuple with the SshKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKey

`func (o *GetManagedSshKey200Response) SetSshKey(v string)`

SetSshKey sets SshKey field to given value.

### HasSshKey

`func (o *GetManagedSshKey200Response) HasSshKey() bool`

HasSshKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


