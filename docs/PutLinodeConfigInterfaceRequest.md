# PutLinodeConfigInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpRanges** | Pointer to **[]string** | An array of IPv4 CIDR VPC Subnet ranges that are routed to this Interface.  - Array items are only allowed for &#x60;vpc&#x60; type Interfaces. - This must be empty for non-&#x60;vpc&#x60; type Interfaces.  For requests:  - Addresses in submitted ranges must not already be actively assigned. - Submitting values replaces any existing values. - Submitting an empty array removes any existing values. - Omitting this property results in no change to existing values. | [optional] 
**Ipv4** | Pointer to [**PostLinodeInstanceRequestAllOfInterfacesInnerIpv4**](PostLinodeInstanceRequestAllOfInterfacesInnerIpv4.md) |  | [optional] 
**Primary** | Pointer to **bool** | The primary Interface is configured as the default route to the Linode.  Each Configuration Profile can have up to one &#x60;\&quot;primary\&quot;: true&#x60; Interface at a time.  Must be &#x60;false&#x60; for &#x60;vlan&#x60; type Interfaces.  If no Interface is configured as the primary, the first non-&#x60;vlan&#x60; type Interface in the &#x60;interfaces&#x60; array is automatically treated as the primary Interface. | [optional] 

## Methods

### NewPutLinodeConfigInterfaceRequest

`func NewPutLinodeConfigInterfaceRequest() *PutLinodeConfigInterfaceRequest`

NewPutLinodeConfigInterfaceRequest instantiates a new PutLinodeConfigInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLinodeConfigInterfaceRequestWithDefaults

`func NewPutLinodeConfigInterfaceRequestWithDefaults() *PutLinodeConfigInterfaceRequest`

NewPutLinodeConfigInterfaceRequestWithDefaults instantiates a new PutLinodeConfigInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpRanges

`func (o *PutLinodeConfigInterfaceRequest) GetIpRanges() []string`

GetIpRanges returns the IpRanges field if non-nil, zero value otherwise.

### GetIpRangesOk

`func (o *PutLinodeConfigInterfaceRequest) GetIpRangesOk() (*[]string, bool)`

GetIpRangesOk returns a tuple with the IpRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRanges

`func (o *PutLinodeConfigInterfaceRequest) SetIpRanges(v []string)`

SetIpRanges sets IpRanges field to given value.

### HasIpRanges

`func (o *PutLinodeConfigInterfaceRequest) HasIpRanges() bool`

HasIpRanges returns a boolean if a field has been set.

### SetIpRangesNil

`func (o *PutLinodeConfigInterfaceRequest) SetIpRangesNil(b bool)`

 SetIpRangesNil sets the value for IpRanges to be an explicit nil

### UnsetIpRanges
`func (o *PutLinodeConfigInterfaceRequest) UnsetIpRanges()`

UnsetIpRanges ensures that no value is present for IpRanges, not even an explicit nil
### GetIpv4

`func (o *PutLinodeConfigInterfaceRequest) GetIpv4() PostLinodeInstanceRequestAllOfInterfacesInnerIpv4`

GetIpv4 returns the Ipv4 field if non-nil, zero value otherwise.

### GetIpv4Ok

`func (o *PutLinodeConfigInterfaceRequest) GetIpv4Ok() (*PostLinodeInstanceRequestAllOfInterfacesInnerIpv4, bool)`

GetIpv4Ok returns a tuple with the Ipv4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4

`func (o *PutLinodeConfigInterfaceRequest) SetIpv4(v PostLinodeInstanceRequestAllOfInterfacesInnerIpv4)`

SetIpv4 sets Ipv4 field to given value.

### HasIpv4

`func (o *PutLinodeConfigInterfaceRequest) HasIpv4() bool`

HasIpv4 returns a boolean if a field has been set.

### GetPrimary

`func (o *PutLinodeConfigInterfaceRequest) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PutLinodeConfigInterfaceRequest) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PutLinodeConfigInterfaceRequest) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *PutLinodeConfigInterfaceRequest) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


