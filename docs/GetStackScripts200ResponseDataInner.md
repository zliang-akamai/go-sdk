# GetStackScripts200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date this StackScript was created. | [optional] [readonly] 
**DeploymentsActive** | Pointer to **int32** | Count of currently active, deployed Linodes created from this StackScript. | [optional] [readonly] 
**DeploymentsTotal** | Pointer to **int32** | The total number of times this StackScript has been deployed. | [optional] [readonly] 
**Description** | Pointer to **string** | A description for the StackScript. | [optional] 
**Id** | Pointer to **int32** | The unique ID of this StackScript. | [optional] [readonly] 
**Images** | Pointer to **[]string** | An array of Image IDs. These are the Images that can be deployed with this StackScript.  &#x60;any/all&#x60; indicates that all available Images, including private Images, are accepted. | [optional] 
**IsPublic** | Pointer to **bool** | This determines whether other users can use your StackScript. __Once a StackScript is made public, it cannot be made private.__ | [optional] 
**Label** | Pointer to **string** | The StackScript&#39;s label is for display purposes only. | [optional] 
**Mine** | Pointer to **bool** | Returns &#x60;true&#x60; if this StackScript is owned by the account of the user making the request, and the user making the request is unrestricted or has access to this StackScript. | [optional] [readonly] 
**RevNote** | Pointer to **string** | This field allows you to add notes for the set of revisions made to this StackScript. | [optional] 
**Script** | Pointer to **string** | The script to execute when provisioning a new Linode with this StackScript. | [optional] 
**Updated** | Pointer to **time.Time** | The date this StackScript was last updated. | [optional] [readonly] 
**UserDefinedFields** | Pointer to [**[]GetStackScripts200ResponseDataInnerUserDefinedFieldsInner**](GetStackScripts200ResponseDataInnerUserDefinedFieldsInner.md) | This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment. See [Declare User-Defined Fields (UDFs)](https://www.linode.com/docs/products/tools/stackscripts/guides/write-a-custom-script/#declare-user-defined-fields-udfs) for more information. | [optional] [readonly] 
**UserGravatarId** | Pointer to **string** | The Gravatar ID for the User who created the StackScript. | [optional] [readonly] 
**Username** | Pointer to **string** | The User who created the StackScript. | [optional] [readonly] 

## Methods

### NewGetStackScripts200ResponseDataInner

`func NewGetStackScripts200ResponseDataInner() *GetStackScripts200ResponseDataInner`

NewGetStackScripts200ResponseDataInner instantiates a new GetStackScripts200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStackScripts200ResponseDataInnerWithDefaults

`func NewGetStackScripts200ResponseDataInnerWithDefaults() *GetStackScripts200ResponseDataInner`

NewGetStackScripts200ResponseDataInnerWithDefaults instantiates a new GetStackScripts200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetStackScripts200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetStackScripts200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetStackScripts200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetStackScripts200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeploymentsActive

`func (o *GetStackScripts200ResponseDataInner) GetDeploymentsActive() int32`

GetDeploymentsActive returns the DeploymentsActive field if non-nil, zero value otherwise.

### GetDeploymentsActiveOk

`func (o *GetStackScripts200ResponseDataInner) GetDeploymentsActiveOk() (*int32, bool)`

GetDeploymentsActiveOk returns a tuple with the DeploymentsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsActive

`func (o *GetStackScripts200ResponseDataInner) SetDeploymentsActive(v int32)`

SetDeploymentsActive sets DeploymentsActive field to given value.

### HasDeploymentsActive

`func (o *GetStackScripts200ResponseDataInner) HasDeploymentsActive() bool`

HasDeploymentsActive returns a boolean if a field has been set.

### GetDeploymentsTotal

`func (o *GetStackScripts200ResponseDataInner) GetDeploymentsTotal() int32`

GetDeploymentsTotal returns the DeploymentsTotal field if non-nil, zero value otherwise.

### GetDeploymentsTotalOk

`func (o *GetStackScripts200ResponseDataInner) GetDeploymentsTotalOk() (*int32, bool)`

GetDeploymentsTotalOk returns a tuple with the DeploymentsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsTotal

`func (o *GetStackScripts200ResponseDataInner) SetDeploymentsTotal(v int32)`

SetDeploymentsTotal sets DeploymentsTotal field to given value.

### HasDeploymentsTotal

`func (o *GetStackScripts200ResponseDataInner) HasDeploymentsTotal() bool`

HasDeploymentsTotal returns a boolean if a field has been set.

### GetDescription

`func (o *GetStackScripts200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetStackScripts200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetStackScripts200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetStackScripts200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *GetStackScripts200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetStackScripts200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetStackScripts200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetStackScripts200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImages

`func (o *GetStackScripts200ResponseDataInner) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetStackScripts200ResponseDataInner) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetStackScripts200ResponseDataInner) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *GetStackScripts200ResponseDataInner) HasImages() bool`

HasImages returns a boolean if a field has been set.

### GetIsPublic

`func (o *GetStackScripts200ResponseDataInner) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *GetStackScripts200ResponseDataInner) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *GetStackScripts200ResponseDataInner) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *GetStackScripts200ResponseDataInner) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetLabel

`func (o *GetStackScripts200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetStackScripts200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetStackScripts200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetStackScripts200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMine

`func (o *GetStackScripts200ResponseDataInner) GetMine() bool`

GetMine returns the Mine field if non-nil, zero value otherwise.

### GetMineOk

`func (o *GetStackScripts200ResponseDataInner) GetMineOk() (*bool, bool)`

GetMineOk returns a tuple with the Mine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMine

`func (o *GetStackScripts200ResponseDataInner) SetMine(v bool)`

SetMine sets Mine field to given value.

### HasMine

`func (o *GetStackScripts200ResponseDataInner) HasMine() bool`

HasMine returns a boolean if a field has been set.

### GetRevNote

`func (o *GetStackScripts200ResponseDataInner) GetRevNote() string`

GetRevNote returns the RevNote field if non-nil, zero value otherwise.

### GetRevNoteOk

`func (o *GetStackScripts200ResponseDataInner) GetRevNoteOk() (*string, bool)`

GetRevNoteOk returns a tuple with the RevNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevNote

`func (o *GetStackScripts200ResponseDataInner) SetRevNote(v string)`

SetRevNote sets RevNote field to given value.

### HasRevNote

`func (o *GetStackScripts200ResponseDataInner) HasRevNote() bool`

HasRevNote returns a boolean if a field has been set.

### GetScript

`func (o *GetStackScripts200ResponseDataInner) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *GetStackScripts200ResponseDataInner) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *GetStackScripts200ResponseDataInner) SetScript(v string)`

SetScript sets Script field to given value.

### HasScript

`func (o *GetStackScripts200ResponseDataInner) HasScript() bool`

HasScript returns a boolean if a field has been set.

### GetUpdated

`func (o *GetStackScripts200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetStackScripts200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetStackScripts200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetStackScripts200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *GetStackScripts200ResponseDataInner) GetUserDefinedFields() []GetStackScripts200ResponseDataInnerUserDefinedFieldsInner`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *GetStackScripts200ResponseDataInner) GetUserDefinedFieldsOk() (*[]GetStackScripts200ResponseDataInnerUserDefinedFieldsInner, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *GetStackScripts200ResponseDataInner) SetUserDefinedFields(v []GetStackScripts200ResponseDataInnerUserDefinedFieldsInner)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *GetStackScripts200ResponseDataInner) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetUserGravatarId

`func (o *GetStackScripts200ResponseDataInner) GetUserGravatarId() string`

GetUserGravatarId returns the UserGravatarId field if non-nil, zero value otherwise.

### GetUserGravatarIdOk

`func (o *GetStackScripts200ResponseDataInner) GetUserGravatarIdOk() (*string, bool)`

GetUserGravatarIdOk returns a tuple with the UserGravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGravatarId

`func (o *GetStackScripts200ResponseDataInner) SetUserGravatarId(v string)`

SetUserGravatarId sets UserGravatarId field to given value.

### HasUserGravatarId

`func (o *GetStackScripts200ResponseDataInner) HasUserGravatarId() bool`

HasUserGravatarId returns a boolean if a field has been set.

### GetUsername

`func (o *GetStackScripts200ResponseDataInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetStackScripts200ResponseDataInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetStackScripts200ResponseDataInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetStackScripts200ResponseDataInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


