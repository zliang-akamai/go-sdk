# PostFirewallDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The entity&#39;s ID. | 
**Label** | Pointer to **string** | The entity&#39;s label. | [optional] [readonly] 
**Type** | **string** | The entity&#39;s type. | 
**Url** | Pointer to **string** | The API URL path you can use to access this entity. | [optional] [readonly] 

## Methods

### NewPostFirewallDeviceRequest

`func NewPostFirewallDeviceRequest(id int32, type_ string, ) *PostFirewallDeviceRequest`

NewPostFirewallDeviceRequest instantiates a new PostFirewallDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFirewallDeviceRequestWithDefaults

`func NewPostFirewallDeviceRequestWithDefaults() *PostFirewallDeviceRequest`

NewPostFirewallDeviceRequestWithDefaults instantiates a new PostFirewallDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostFirewallDeviceRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostFirewallDeviceRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostFirewallDeviceRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetLabel

`func (o *PostFirewallDeviceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostFirewallDeviceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostFirewallDeviceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostFirewallDeviceRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetType

`func (o *PostFirewallDeviceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostFirewallDeviceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostFirewallDeviceRequest) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *PostFirewallDeviceRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PostFirewallDeviceRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PostFirewallDeviceRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PostFirewallDeviceRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


