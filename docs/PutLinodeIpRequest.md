# PutLinodeIpRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rdns** | **NullableString** | The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set. | 

## Methods

### NewPutLinodeIpRequest

`func NewPutLinodeIpRequest(rdns NullableString, ) *PutLinodeIpRequest`

NewPutLinodeIpRequest instantiates a new PutLinodeIpRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLinodeIpRequestWithDefaults

`func NewPutLinodeIpRequestWithDefaults() *PutLinodeIpRequest`

NewPutLinodeIpRequestWithDefaults instantiates a new PutLinodeIpRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRdns

`func (o *PutLinodeIpRequest) GetRdns() string`

GetRdns returns the Rdns field if non-nil, zero value otherwise.

### GetRdnsOk

`func (o *PutLinodeIpRequest) GetRdnsOk() (*string, bool)`

GetRdnsOk returns a tuple with the Rdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdns

`func (o *PutLinodeIpRequest) SetRdns(v string)`

SetRdns sets Rdns field to given value.


### SetRdnsNil

`func (o *PutLinodeIpRequest) SetRdnsNil(b bool)`

 SetRdnsNil sets the value for Rdns to be an explicit nil

### UnsetRdns
`func (o *PutLinodeIpRequest) UnsetRdns()`

UnsetRdns ensures that no value is present for Rdns, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


