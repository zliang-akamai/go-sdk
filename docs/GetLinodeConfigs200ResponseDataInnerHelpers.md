# GetLinodeConfigs200ResponseDataInnerHelpers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DevtmpfsAutomount** | Pointer to **bool** | Populates the /dev directory early during boot without udev.  Defaults to false. | [optional] 
**Distro** | Pointer to **bool** | Helps maintain correct inittab/upstart console device. | [optional] 
**ModulesDep** | Pointer to **bool** | Creates a modules dependency file for the Kernel you run. | [optional] 
**Network** | Pointer to **bool** | Automatically configures static networking. | [optional] 
**UpdatedbDisabled** | Pointer to **bool** | Disables updatedb cron job to avoid disk thrashing. | [optional] 

## Methods

### NewGetLinodeConfigs200ResponseDataInnerHelpers

`func NewGetLinodeConfigs200ResponseDataInnerHelpers() *GetLinodeConfigs200ResponseDataInnerHelpers`

NewGetLinodeConfigs200ResponseDataInnerHelpers instantiates a new GetLinodeConfigs200ResponseDataInnerHelpers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinodeConfigs200ResponseDataInnerHelpersWithDefaults

`func NewGetLinodeConfigs200ResponseDataInnerHelpersWithDefaults() *GetLinodeConfigs200ResponseDataInnerHelpers`

NewGetLinodeConfigs200ResponseDataInnerHelpersWithDefaults instantiates a new GetLinodeConfigs200ResponseDataInnerHelpers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevtmpfsAutomount

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetDevtmpfsAutomount() bool`

GetDevtmpfsAutomount returns the DevtmpfsAutomount field if non-nil, zero value otherwise.

### GetDevtmpfsAutomountOk

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetDevtmpfsAutomountOk() (*bool, bool)`

GetDevtmpfsAutomountOk returns a tuple with the DevtmpfsAutomount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevtmpfsAutomount

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetDevtmpfsAutomount(v bool)`

SetDevtmpfsAutomount sets DevtmpfsAutomount field to given value.

### HasDevtmpfsAutomount

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasDevtmpfsAutomount() bool`

HasDevtmpfsAutomount returns a boolean if a field has been set.

### GetDistro

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetDistro() bool`

GetDistro returns the Distro field if non-nil, zero value otherwise.

### GetDistroOk

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetDistroOk() (*bool, bool)`

GetDistroOk returns a tuple with the Distro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistro

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetDistro(v bool)`

SetDistro sets Distro field to given value.

### HasDistro

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasDistro() bool`

HasDistro returns a boolean if a field has been set.

### GetModulesDep

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetModulesDep() bool`

GetModulesDep returns the ModulesDep field if non-nil, zero value otherwise.

### GetModulesDepOk

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetModulesDepOk() (*bool, bool)`

GetModulesDepOk returns a tuple with the ModulesDep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulesDep

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetModulesDep(v bool)`

SetModulesDep sets ModulesDep field to given value.

### HasModulesDep

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasModulesDep() bool`

HasModulesDep returns a boolean if a field has been set.

### GetNetwork

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetNetwork() bool`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetNetworkOk() (*bool, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetNetwork(v bool)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetUpdatedbDisabled

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetUpdatedbDisabled() bool`

GetUpdatedbDisabled returns the UpdatedbDisabled field if non-nil, zero value otherwise.

### GetUpdatedbDisabledOk

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) GetUpdatedbDisabledOk() (*bool, bool)`

GetUpdatedbDisabledOk returns a tuple with the UpdatedbDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedbDisabled

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) SetUpdatedbDisabled(v bool)`

SetUpdatedbDisabled sets UpdatedbDisabled field to given value.

### HasUpdatedbDisabled

`func (o *GetLinodeConfigs200ResponseDataInnerHelpers) HasUpdatedbDisabled() bool`

HasUpdatedbDisabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


