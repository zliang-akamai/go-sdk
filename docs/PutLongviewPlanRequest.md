# PutLongviewPlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LongviewSubscription** | Pointer to **NullableString** | The subscription ID for a particular Longview plan. A value of &#x60;null&#x60; corresponds to Longview Free. You can send a request to the [List Longview subscriptions](https://techdocs.akamai.com/linode-api/reference/get-longview-subscriptions) operation to receive the details of each plan. | [optional] 

## Methods

### NewPutLongviewPlanRequest

`func NewPutLongviewPlanRequest() *PutLongviewPlanRequest`

NewPutLongviewPlanRequest instantiates a new PutLongviewPlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutLongviewPlanRequestWithDefaults

`func NewPutLongviewPlanRequestWithDefaults() *PutLongviewPlanRequest`

NewPutLongviewPlanRequestWithDefaults instantiates a new PutLongviewPlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLongviewSubscription

`func (o *PutLongviewPlanRequest) GetLongviewSubscription() string`

GetLongviewSubscription returns the LongviewSubscription field if non-nil, zero value otherwise.

### GetLongviewSubscriptionOk

`func (o *PutLongviewPlanRequest) GetLongviewSubscriptionOk() (*string, bool)`

GetLongviewSubscriptionOk returns a tuple with the LongviewSubscription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongviewSubscription

`func (o *PutLongviewPlanRequest) SetLongviewSubscription(v string)`

SetLongviewSubscription sets LongviewSubscription field to given value.

### HasLongviewSubscription

`func (o *PutLongviewPlanRequest) HasLongviewSubscription() bool`

HasLongviewSubscription returns a boolean if a field has been set.

### SetLongviewSubscriptionNil

`func (o *PutLongviewPlanRequest) SetLongviewSubscriptionNil(b bool)`

 SetLongviewSubscriptionNil sets the value for LongviewSubscription to be an explicit nil

### UnsetLongviewSubscription
`func (o *PutLongviewPlanRequest) UnsetLongviewSubscription()`

UnsetLongviewSubscription ensures that no value is present for LongviewSubscription, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


