# ErrorObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** | The field in the request that caused this error. This may be a path, separated by periods in the case of nested fields. In some cases this may come back as &#x60;null&#x60; if the error is not specific to any single element of the request. | [optional] 
**Reason** | Pointer to **string** | What happened to cause this error. In most cases, this can be fixed immediately by changing the data you sent in the request, but in some cases you will be instructed to [Open a support ticket](https://techdocs.akamai.com/linode-api/reference/post-ticket) or perform some other action before you can complete the request successfully. | [optional] 

## Methods

### NewErrorObject

`func NewErrorObject() *ErrorObject`

NewErrorObject instantiates a new ErrorObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorObjectWithDefaults

`func NewErrorObjectWithDefaults() *ErrorObject`

NewErrorObjectWithDefaults instantiates a new ErrorObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ErrorObject) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ErrorObject) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ErrorObject) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ErrorObject) HasField() bool`

HasField returns a boolean if a field has been set.

### GetReason

`func (o *ErrorObject) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorObject) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorObject) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorObject) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


