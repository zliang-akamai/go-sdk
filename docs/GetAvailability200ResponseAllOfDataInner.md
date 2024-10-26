# GetAvailability200ResponseAllOfDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **[]string** | A list of services _available_ to your account in the &#x60;region&#x60;. | [optional] [readonly] 
**Region** | Pointer to **string** | The Akamai cloud computing data center (region), represented by a slug value. You can view a full list of regions and their associated slugs with the [List regions](https://techdocs.akamai.com/linode-api/reference/get-regions) operation. | [optional] [readonly] 
**Unavailable** | Pointer to **[]string** | A list of services _unavailable_ to your account in the &#x60;region&#x60;. | [optional] [readonly] 

## Methods

### NewGetAvailability200ResponseAllOfDataInner

`func NewGetAvailability200ResponseAllOfDataInner() *GetAvailability200ResponseAllOfDataInner`

NewGetAvailability200ResponseAllOfDataInner instantiates a new GetAvailability200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAvailability200ResponseAllOfDataInnerWithDefaults

`func NewGetAvailability200ResponseAllOfDataInnerWithDefaults() *GetAvailability200ResponseAllOfDataInner`

NewGetAvailability200ResponseAllOfDataInnerWithDefaults instantiates a new GetAvailability200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *GetAvailability200ResponseAllOfDataInner) GetAvailable() []string`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetAvailability200ResponseAllOfDataInner) GetAvailableOk() (*[]string, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetAvailability200ResponseAllOfDataInner) SetAvailable(v []string)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *GetAvailability200ResponseAllOfDataInner) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetRegion

`func (o *GetAvailability200ResponseAllOfDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetAvailability200ResponseAllOfDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetAvailability200ResponseAllOfDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetAvailability200ResponseAllOfDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUnavailable

`func (o *GetAvailability200ResponseAllOfDataInner) GetUnavailable() []string`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *GetAvailability200ResponseAllOfDataInner) GetUnavailableOk() (*[]string, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *GetAvailability200ResponseAllOfDataInner) SetUnavailable(v []string)`

SetUnavailable sets Unavailable field to given value.

### HasUnavailable

`func (o *GetAvailability200ResponseAllOfDataInner) HasUnavailable() bool`

HasUnavailable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


