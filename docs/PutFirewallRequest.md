# PutFirewallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** | The Firewall&#39;s label, for display purposes only.  Firewall labels have the following constraints:    - Must begin and end with an alphanumeric character.   - May only consist of alphanumeric characters, hyphens (&#x60;-&#x60;), underscores (&#x60;_&#x60;) or periods (&#x60;.&#x60;).   - Cannot have two hyphens (&#x60;--&#x60;), underscores (&#x60;__&#x60;) or periods (&#x60;..&#x60;) in a row.   - Must be between 3 and 32 characters.   - Must be unique. | [optional] 
**Status** | Pointer to **string** | The status to be applied to this Firewall.   - When a Firewall is first created its status is &#x60;enabled&#x60;.  - Run the [Delete a firewall](https://techdocs.akamai.com/linode-api/reference/delete-firewall) operation to delete a Firewall. | [optional] 
**Tags** | Pointer to **[]string** | An array of tags applied to this object. Tags are for organizational purposes only. | [optional] 

## Methods

### NewPutFirewallRequest

`func NewPutFirewallRequest() *PutFirewallRequest`

NewPutFirewallRequest instantiates a new PutFirewallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPutFirewallRequestWithDefaults

`func NewPutFirewallRequestWithDefaults() *PutFirewallRequest`

NewPutFirewallRequestWithDefaults instantiates a new PutFirewallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *PutFirewallRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PutFirewallRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PutFirewallRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PutFirewallRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetStatus

`func (o *PutFirewallRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PutFirewallRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PutFirewallRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PutFirewallRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PutFirewallRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PutFirewallRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PutFirewallRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PutFirewallRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


