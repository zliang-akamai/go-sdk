# AddedGetAvailability200AllOfData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **[]string** | A list of services _available_ to your account in the &#x60;region&#x60;. | [optional] [readonly] 
**Region** | Pointer to **string** | The Akamai cloud computing data center (region), represented by a slug value. You can view a full list of regions and their associated slugs with the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation. | [optional] [readonly] 
**Unavailable** | Pointer to **[]string** | A list of services _unavailable_ to your account in the &#x60;region&#x60;. | [optional] [readonly] 

## Methods

### NewAddedGetAvailability200AllOfData

`func NewAddedGetAvailability200AllOfData() *AddedGetAvailability200AllOfData`

NewAddedGetAvailability200AllOfData instantiates a new AddedGetAvailability200AllOfData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddedGetAvailability200AllOfDataWithDefaults

`func NewAddedGetAvailability200AllOfDataWithDefaults() *AddedGetAvailability200AllOfData`

NewAddedGetAvailability200AllOfDataWithDefaults instantiates a new AddedGetAvailability200AllOfData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *AddedGetAvailability200AllOfData) GetAvailable() []string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *AddedGetAvailability200AllOfData) GetAvailableOk() (*[]string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *AddedGetAvailability200AllOfData) SetAvailable(v []string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *AddedGetAvailability200AllOfData) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetRegion

`func (o *AddedGetAvailability200AllOfData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddedGetAvailability200AllOfData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddedGetAvailability200AllOfData) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddedGetAvailability200AllOfData) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUnavailable

`func (o *AddedGetAvailability200AllOfData) GetUnavailable() []string`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *AddedGetAvailability200AllOfData) GetUnavailableOk() (*[]string, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *AddedGetAvailability200AllOfData) SetUnavailable(v []string)`

SetUnavailable sets Unavailable field to given value.

### HasUnavailable

`func (o *AddedGetAvailability200AllOfData) HasUnavailable() bool`

HasUnavailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


