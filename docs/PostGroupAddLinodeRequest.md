# PostGroupAddLinodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Linodes** | Pointer to **[]int32** | The &#x60;linodeId&#x60; values for individual compute instances included in the placement group. | [optional] 

## Methods

### NewPostGroupAddLinodeRequest

`func NewPostGroupAddLinodeRequest() *PostGroupAddLinodeRequest`

NewPostGroupAddLinodeRequest instantiates a new PostGroupAddLinodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostGroupAddLinodeRequestWithDefaults

`func NewPostGroupAddLinodeRequestWithDefaults() *PostGroupAddLinodeRequest`

NewPostGroupAddLinodeRequestWithDefaults instantiates a new PostGroupAddLinodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinodes

`func (o *PostGroupAddLinodeRequest) GetLinodes() []int32`

GetLinodes returns the Linodes field if non-nil, zero value otherwise.

### GetLinodesOk

`func (o *PostGroupAddLinodeRequest) GetLinodesOk() (*[]int32, bool)`

GetLinodesOk returns a tuple with the Linodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinodes

`func (o *PostGroupAddLinodeRequest) SetLinodes(v []int32)`

SetLinodes sets Linodes field to given value.

### HasLinodes

`func (o *PostGroupAddLinodeRequest) HasLinodes() bool`

HasLinodes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


