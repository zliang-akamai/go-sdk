# GetLkeClusterAcl400ResponseErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | The field in the request that caused this error. This may be a path, separated by periods in the case of nested fields. In some cases this may come back as null if the error is not specific to any single element of the request. | [optional] 
**Reason** | Pointer to **string** | A string explaining that the cluster does not support Control Plane ACL. | [optional] 

## Methods

### NewGetLkeClusterAcl400ResponseErrorsInner

`func NewGetLkeClusterAcl400ResponseErrorsInner() *GetLkeClusterAcl400ResponseErrorsInner`

NewGetLkeClusterAcl400ResponseErrorsInner instantiates a new GetLkeClusterAcl400ResponseErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLkeClusterAcl400ResponseErrorsInnerWithDefaults

`func NewGetLkeClusterAcl400ResponseErrorsInnerWithDefaults() *GetLkeClusterAcl400ResponseErrorsInner`

NewGetLkeClusterAcl400ResponseErrorsInnerWithDefaults instantiates a new GetLkeClusterAcl400ResponseErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *GetLkeClusterAcl400ResponseErrorsInner) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *GetLkeClusterAcl400ResponseErrorsInner) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *GetLkeClusterAcl400ResponseErrorsInner) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *GetLkeClusterAcl400ResponseErrorsInner) HasField() bool`

HasField returns a boolean if a field has been set.

### GetReason

`func (o *GetLkeClusterAcl400ResponseErrorsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetLkeClusterAcl400ResponseErrorsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetLkeClusterAcl400ResponseErrorsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetLkeClusterAcl400ResponseErrorsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


