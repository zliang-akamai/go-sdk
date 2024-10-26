# WarningObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **string** | Specific information related to the warning. | [optional] 
**Title** | Pointer to **string** | The general warning message. | [optional] 

## Methods

### NewWarningObject

`func NewWarningObject() *WarningObject`

NewWarningObject instantiates a new WarningObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningObjectWithDefaults

`func NewWarningObjectWithDefaults() *WarningObject`

NewWarningObjectWithDefaults instantiates a new WarningObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *WarningObject) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *WarningObject) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *WarningObject) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *WarningObject) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetTitle

`func (o *WarningObject) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WarningObject) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WarningObject) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WarningObject) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


