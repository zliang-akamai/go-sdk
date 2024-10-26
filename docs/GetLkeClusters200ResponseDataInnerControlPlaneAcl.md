# GetLkeClusters200ResponseDataInnerControlPlaneAcl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to [**GetLkeClusters200ResponseDataInnerControlPlaneAclAddresses**](GetLkeClusters200ResponseDataInnerControlPlaneAclAddresses.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Defines Default Policy.  A value of true results in a default policy of DENY.  A value of false results in a default policy of ALLOW (i.e., access controls are disabled). Defaults to &#x60;true&#x60;. Creating a cluster with ACL (or upgrading a cluster to use ACL for LKE) is an __irreversible__ change: once upgraded, access controls can only be toggled through this property. | [optional] 
**RevisionId** | Pointer to **string** | Enables clients to track events related to ACL update requests and enforcements. Optional field. If omitted, defaults to a randomly generated string. | [optional] 

## Methods

### NewGetLkeClusters200ResponseDataInnerControlPlaneAcl

`func NewGetLkeClusters200ResponseDataInnerControlPlaneAcl() *GetLkeClusters200ResponseDataInnerControlPlaneAcl`

NewGetLkeClusters200ResponseDataInnerControlPlaneAcl instantiates a new GetLkeClusters200ResponseDataInnerControlPlaneAcl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLkeClusters200ResponseDataInnerControlPlaneAclWithDefaults

`func NewGetLkeClusters200ResponseDataInnerControlPlaneAclWithDefaults() *GetLkeClusters200ResponseDataInnerControlPlaneAcl`

NewGetLkeClusters200ResponseDataInnerControlPlaneAclWithDefaults instantiates a new GetLkeClusters200ResponseDataInnerControlPlaneAcl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) GetAddresses() GetLkeClusters200ResponseDataInnerControlPlaneAclAddresses`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) GetAddressesOk() (*GetLkeClusters200ResponseDataInnerControlPlaneAclAddresses, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) SetAddresses(v GetLkeClusters200ResponseDataInnerControlPlaneAclAddresses)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetEnabled

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetRevisionId

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) GetRevisionId() string`

GetRevisionId returns the RevisionId field if non-nil, zero value otherwise.

### GetRevisionIdOk

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) GetRevisionIdOk() (*string, bool)`

GetRevisionIdOk returns a tuple with the RevisionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisionId

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) SetRevisionId(v string)`

SetRevisionId sets RevisionId field to given value.

### HasRevisionId

`func (o *GetLkeClusters200ResponseDataInnerControlPlaneAcl) HasRevisionId() bool`

HasRevisionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


