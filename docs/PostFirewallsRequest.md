# PostFirewallsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this Firewall was created. | [optional] [readonly] 
**Id** | Pointer to **int32** | The Firewall&#39;s unique ID. | [optional] [readonly] 
**Label** | Pointer to **string** | The Firewall&#39;s label, for display purposes only.  Firewall labels have the following constraints:    - Must begin and end with an alphanumeric character.   - May only consist of alphanumeric characters, hyphens (&#x60;-&#x60;), underscores (&#x60;_&#x60;) or periods (&#x60;.&#x60;).   - Cannot have two hyphens (&#x60;--&#x60;), underscores (&#x60;__&#x60;) or periods (&#x60;..&#x60;) in a row.   - Must be between 3 and 32 characters.   - Must be unique. | [optional] 
**Rules** | [**PostFirewallsRequestAllOfRules**](PostFirewallsRequestAllOfRules.md) |  | 
**Status** | Pointer to **string** | The status of this Firewall.    - When a Firewall is first created its status is &#x60;enabled&#x60;.   - Run the [Update a firewall](https://techdocs.akamai.com/linode-api/reference/put-firewall) operation to set a Firewall&#39;s status to &#x60;enabled&#x60; or &#x60;disabled&#x60;.   - Run the [Delete a firewall](https://techdocs.akamai.com/linode-api/reference/delete-firewall) operation to delete a Firewall. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | An array of tags applied to this object. Tags are for organizational purposes only. | [optional] 
**Updated** | Pointer to **time.Time** | When this Firewall was last updated. | [optional] [readonly] 
**Devices** | Pointer to [**PostFirewallsRequestAllOfDevices**](PostFirewallsRequestAllOfDevices.md) |  | [optional] 

## Methods

### NewPostFirewallsRequest

`func NewPostFirewallsRequest(rules PostFirewallsRequestAllOfRules, ) *PostFirewallsRequest`

NewPostFirewallsRequest instantiates a new PostFirewallsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostFirewallsRequestWithDefaults

`func NewPostFirewallsRequestWithDefaults() *PostFirewallsRequest`

NewPostFirewallsRequestWithDefaults instantiates a new PostFirewallsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PostFirewallsRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PostFirewallsRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PostFirewallsRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PostFirewallsRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetId

`func (o *PostFirewallsRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostFirewallsRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostFirewallsRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostFirewallsRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PostFirewallsRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostFirewallsRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostFirewallsRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PostFirewallsRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRules

`func (o *PostFirewallsRequest) GetRules() PostFirewallsRequestAllOfRules`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *PostFirewallsRequest) GetRulesOk() (*PostFirewallsRequestAllOfRules, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *PostFirewallsRequest) SetRules(v PostFirewallsRequestAllOfRules)`

SetRules sets Rules field to given value.


### GetStatus

`func (o *PostFirewallsRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostFirewallsRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostFirewallsRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostFirewallsRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PostFirewallsRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PostFirewallsRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PostFirewallsRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PostFirewallsRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdated

`func (o *PostFirewallsRequest) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PostFirewallsRequest) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PostFirewallsRequest) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PostFirewallsRequest) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetDevices

`func (o *PostFirewallsRequest) GetDevices() PostFirewallsRequestAllOfDevices`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *PostFirewallsRequest) GetDevicesOk() (*PostFirewallsRequestAllOfDevices, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *PostFirewallsRequest) SetDevices(v PostFirewallsRequestAllOfDevices)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *PostFirewallsRequest) HasDevices() bool`

HasDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


