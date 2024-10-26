# GetNodeBalancerConfigs200ResponseDataInnerNodesStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Down** | Pointer to **int32** | The number of backends considered to be \&quot;DOWN\&quot; and unhealthy.  These are not in rotation, and not serving requests. | [optional] [readonly] 
**Up** | Pointer to **int32** | The number of backends considered to be \&quot;UP\&quot; and healthy, and that are serving requests. | [optional] [readonly] 

## Methods

### NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatus

`func NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatus() *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus`

NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatus instantiates a new GetNodeBalancerConfigs200ResponseDataInnerNodesStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatusWithDefaults

`func NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatusWithDefaults() *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus`

NewGetNodeBalancerConfigs200ResponseDataInnerNodesStatusWithDefaults instantiates a new GetNodeBalancerConfigs200ResponseDataInnerNodesStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDown

`func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) GetDown() int32`

GetDown returns the Down field if non-nil, zero value otherwise.

### GetDownOk

`func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) GetDownOk() (*int32, bool)`

GetDownOk returns a tuple with the Down field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDown

`func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) SetDown(v int32)`

SetDown sets Down field to given value.

### HasDown

`func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) HasDown() bool`

HasDown returns a boolean if a field has been set.

### GetUp

`func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) GetUp() int32`

GetUp returns the Up field if non-nil, zero value otherwise.

### GetUpOk

`func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) GetUpOk() (*int32, bool)`

GetUpOk returns a tuple with the Up field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUp

`func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) SetUp(v int32)`

SetUp sets Up field to given value.

### HasUp

`func (o *GetNodeBalancerConfigs200ResponseDataInnerNodesStatus) HasUp() bool`

HasUp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


