# GetNodeBalancerStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**GetNodeBalancerStats200ResponseData**](GetNodeBalancerStats200ResponseData.md) |  | [optional] 
**Title** | Pointer to **string** | The title for the statistics generated in this response. | [optional] 

## Methods

### NewGetNodeBalancerStats200Response

`func NewGetNodeBalancerStats200Response() *GetNodeBalancerStats200Response`

NewGetNodeBalancerStats200Response instantiates a new GetNodeBalancerStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNodeBalancerStats200ResponseWithDefaults

`func NewGetNodeBalancerStats200ResponseWithDefaults() *GetNodeBalancerStats200Response`

NewGetNodeBalancerStats200ResponseWithDefaults instantiates a new GetNodeBalancerStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetNodeBalancerStats200Response) GetData() GetNodeBalancerStats200ResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetNodeBalancerStats200Response) GetDataOk() (*GetNodeBalancerStats200ResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetNodeBalancerStats200Response) SetData(v GetNodeBalancerStats200ResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *GetNodeBalancerStats200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetTitle

`func (o *GetNodeBalancerStats200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetNodeBalancerStats200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetNodeBalancerStats200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetNodeBalancerStats200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


