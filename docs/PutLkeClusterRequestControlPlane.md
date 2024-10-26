# PutLkeClusterRequestControlPlane

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Acl** | Pointer to [**GetLkeClusters200ResponseDataInnerControlPlaneAcl**](GetLkeClusters200ResponseDataInnerControlPlaneAcl.md) |  | [optional] 
**HighAvailability** | Pointer to **bool** | Enables High Availability for the Control Plane Components of the cluster. Defaults to &#x60;false&#x60;. Enabling High Availability for LKE is an __irreversible__ change. | [optional] [default to false]

## Methods

### NewPutLkeClusterRequestControlPlane

`func NewPutLkeClusterRequestControlPlane() *PutLkeClusterRequestControlPlane`

NewPutLkeClusterRequestControlPlane instantiates a new PutLkeClusterRequestControlPlane object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLkeClusterRequestControlPlaneWithDefaults

`func NewPutLkeClusterRequestControlPlaneWithDefaults() *PutLkeClusterRequestControlPlane`

NewPutLkeClusterRequestControlPlaneWithDefaults instantiates a new PutLkeClusterRequestControlPlane object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcl

`func (o *PutLkeClusterRequestControlPlane) GetAcl() GetLkeClusters200ResponseDataInnerControlPlaneAcl`

GetAcl returns the Acl field if non-nil, zero value otherwise.

### GetAclOk

`func (o *PutLkeClusterRequestControlPlane) GetAclOk() (*GetLkeClusters200ResponseDataInnerControlPlaneAcl, bool)`

GetAclOk returns a tuple with the Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl

`func (o *PutLkeClusterRequestControlPlane) SetAcl(v GetLkeClusters200ResponseDataInnerControlPlaneAcl)`

SetAcl sets Acl field to given value.

### HasAcl

`func (o *PutLkeClusterRequestControlPlane) HasAcl() bool`

HasAcl returns a boolean if a field has been set.

### GetHighAvailability

`func (o *PutLkeClusterRequestControlPlane) GetHighAvailability() bool`

GetHighAvailability returns the HighAvailability field if non-nil, zero value otherwise.

### GetHighAvailabilityOk

`func (o *PutLkeClusterRequestControlPlane) GetHighAvailabilityOk() (*bool, bool)`

GetHighAvailabilityOk returns a tuple with the HighAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighAvailability

`func (o *PutLkeClusterRequestControlPlane) SetHighAvailability(v bool)`

SetHighAvailability sets HighAvailability field to given value.

### HasHighAvailability

`func (o *PutLkeClusterRequestControlPlane) HasHighAvailability() bool`

HasHighAvailability returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


