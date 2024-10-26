# AccountAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **[]string** | A list of services _available_ to your account in the &#x60;region&#x60;. | [optional] [readonly] 
**Region** | Pointer to **string** | The Akamai cloud computing data center (region), represented by a slug value. You can view a full list of regions and their associated slugs with the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation. | [optional] [readonly] 
**Unavailable** | Pointer to **[]string** | A list of services _unavailable_ to your account in the &#x60;region&#x60;. | [optional] [readonly] 

## Methods

### NewAccountAvailability

`func NewAccountAvailability() *AccountAvailability`

NewAccountAvailability instantiates a new AccountAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountAvailabilityWithDefaults

`func NewAccountAvailabilityWithDefaults() *AccountAvailability`

NewAccountAvailabilityWithDefaults instantiates a new AccountAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *AccountAvailability) GetAvailable() []string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AccountAvailability) GetAvailableOk() (*[]string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AccountAvailability) SetAvailable(v []string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AccountAvailability) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetRegion

`func (o *AccountAvailability) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AccountAvailability) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AccountAvailability) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AccountAvailability) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUnavailable

`func (o *AccountAvailability) GetUnavailable() []string`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *AccountAvailability) GetUnavailableOk() (*[]string, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *AccountAvailability) SetUnavailable(v []string)`

SetUnavailable sets Unavailable field to given value.

### HasUnavailable

`func (o *AccountAvailability) HasUnavailable() bool`

HasUnavailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


