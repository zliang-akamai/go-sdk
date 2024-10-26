# PostAssignIpsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assignments** | [**[]PostAssignIpsRequestAssignmentsInner**](PostAssignIpsRequestAssignmentsInner.md) | The list of assignments to make. You must have read_write access to all IPs being assigned and all Linodes being assigned to in order for the assignments to succeed. | 
**Region** | **string** | The ID of the Region in which these assignments are to take place. All IPs and Linodes must exist in this Region. | 

## Methods

### NewPostAssignIpsRequest

`func NewPostAssignIpsRequest(assignments []PostAssignIpsRequestAssignmentsInner, region string, ) *PostAssignIpsRequest`

NewPostAssignIpsRequest instantiates a new PostAssignIpsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAssignIpsRequestWithDefaults

`func NewPostAssignIpsRequestWithDefaults() *PostAssignIpsRequest`

NewPostAssignIpsRequestWithDefaults instantiates a new PostAssignIpsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssignments

`func (o *PostAssignIpsRequest) GetAssignments() []PostAssignIpsRequestAssignmentsInner`

GetAssignments returns the Assignments field if non-nil, zero value otherwise.

### GetAssignmentsOk

`func (o *PostAssignIpsRequest) GetAssignmentsOk() (*[]PostAssignIpsRequestAssignmentsInner, bool)`

GetAssignmentsOk returns a tuple with the Assignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignments

`func (o *PostAssignIpsRequest) SetAssignments(v []PostAssignIpsRequestAssignmentsInner)`

SetAssignments sets Assignments field to given value.


### GetRegion

`func (o *PostAssignIpsRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostAssignIpsRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostAssignIpsRequest) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


