# PostAddStackScriptRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | The date this StackScript was created. | [optional] [readonly] 
**DeploymentsActive** | Pointer to **int32** | Count of currently active, deployed Linodes created from this StackScript. | [optional] [readonly] 
**DeploymentsTotal** | Pointer to **int32** | The total number of times this StackScript has been deployed. | [optional] [readonly] 
**Description** | Pointer to **string** | A description for the StackScript. | [optional] 
**Id** | Pointer to **int32** | The unique ID of this StackScript. | [optional] [readonly] 
**Images** | **[]string** | An array of Image IDs. These are the Images that can be deployed with this StackScript.  &#x60;any/all&#x60; indicates that all available Images, including private Images, are accepted. | 
**IsPublic** | Pointer to **bool** | This determines whether other users can use your StackScript. __Once a StackScript is made public, it cannot be made private.__ | [optional] 
**Label** | **string** | The StackScript&#39;s label is for display purposes only. | 
**Mine** | Pointer to **bool** | Returns &#x60;true&#x60; if this StackScript is owned by the account of the user making the request, and the user making the request is unrestricted or has access to this StackScript. | [optional] [readonly] 
**RevNote** | Pointer to **string** | This field allows you to add notes for the set of revisions made to this StackScript. | [optional] 
**Script** | **string** | The script to execute when provisioning a new Linode with this StackScript. | 
**Updated** | Pointer to **time.Time** | The date this StackScript was last updated. | [optional] [readonly] 
**UserDefinedFields** | Pointer to [**[]GetStackScripts200ResponseDataInnerUserDefinedFieldsInner**](GetStackScripts200ResponseDataInnerUserDefinedFieldsInner.md) | This is a list of fields defined with a special syntax inside this StackScript that allow for supplying customized parameters during deployment. See [Declare User-Defined Fields (UDFs)](https://www.linode.com/docs/products/tools/stackscripts/guides/write-a-custom-script/#declare-user-defined-fields-udfs) for more information. | [optional] [readonly] 
**UserGravatarId** | Pointer to **string** | The Gravatar ID for the User who created the StackScript. | [optional] [readonly] 
**Username** | Pointer to **string** | The User who created the StackScript. | [optional] [readonly] 

## Methods

### NewPostAddStackScriptRequest

`func NewPostAddStackScriptRequest(images []string, label string, script string, ) *PostAddStackScriptRequest`

NewPostAddStackScriptRequest instantiates a new PostAddStackScriptRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAddStackScriptRequestWithDefaults

`func NewPostAddStackScriptRequestWithDefaults() *PostAddStackScriptRequest`

NewPostAddStackScriptRequestWithDefaults instantiates a new PostAddStackScriptRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *PostAddStackScriptRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PostAddStackScriptRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PostAddStackScriptRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PostAddStackScriptRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeploymentsActive

`func (o *PostAddStackScriptRequest) GetDeploymentsActive() int32`

GetDeploymentsActive returns the DeploymentsActive field if non-nil, zero value otherwise.

### GetDeploymentsActiveOk

`func (o *PostAddStackScriptRequest) GetDeploymentsActiveOk() (*int32, bool)`

GetDeploymentsActiveOk returns a tuple with the DeploymentsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsActive

`func (o *PostAddStackScriptRequest) SetDeploymentsActive(v int32)`

SetDeploymentsActive sets DeploymentsActive field to given value.

### HasDeploymentsActive

`func (o *PostAddStackScriptRequest) HasDeploymentsActive() bool`

HasDeploymentsActive returns a boolean if a field has been set.

### GetDeploymentsTotal

`func (o *PostAddStackScriptRequest) GetDeploymentsTotal() int32`

GetDeploymentsTotal returns the DeploymentsTotal field if non-nil, zero value otherwise.

### GetDeploymentsTotalOk

`func (o *PostAddStackScriptRequest) GetDeploymentsTotalOk() (*int32, bool)`

GetDeploymentsTotalOk returns a tuple with the DeploymentsTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentsTotal

`func (o *PostAddStackScriptRequest) SetDeploymentsTotal(v int32)`

SetDeploymentsTotal sets DeploymentsTotal field to given value.

### HasDeploymentsTotal

`func (o *PostAddStackScriptRequest) HasDeploymentsTotal() bool`

HasDeploymentsTotal returns a boolean if a field has been set.

### GetDescription

`func (o *PostAddStackScriptRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostAddStackScriptRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostAddStackScriptRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostAddStackScriptRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *PostAddStackScriptRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostAddStackScriptRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostAddStackScriptRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostAddStackScriptRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImages

`func (o *PostAddStackScriptRequest) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *PostAddStackScriptRequest) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *PostAddStackScriptRequest) SetImages(v []string)`

SetImages sets Images field to given value.


### GetIsPublic

`func (o *PostAddStackScriptRequest) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *PostAddStackScriptRequest) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *PostAddStackScriptRequest) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *PostAddStackScriptRequest) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetLabel

`func (o *PostAddStackScriptRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostAddStackScriptRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostAddStackScriptRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMine

`func (o *PostAddStackScriptRequest) GetMine() bool`

GetMine returns the Mine field if non-nil, zero value otherwise.

### GetMineOk

`func (o *PostAddStackScriptRequest) GetMineOk() (*bool, bool)`

GetMineOk returns a tuple with the Mine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMine

`func (o *PostAddStackScriptRequest) SetMine(v bool)`

SetMine sets Mine field to given value.

### HasMine

`func (o *PostAddStackScriptRequest) HasMine() bool`

HasMine returns a boolean if a field has been set.

### GetRevNote

`func (o *PostAddStackScriptRequest) GetRevNote() string`

GetRevNote returns the RevNote field if non-nil, zero value otherwise.

### GetRevNoteOk

`func (o *PostAddStackScriptRequest) GetRevNoteOk() (*string, bool)`

GetRevNoteOk returns a tuple with the RevNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevNote

`func (o *PostAddStackScriptRequest) SetRevNote(v string)`

SetRevNote sets RevNote field to given value.

### HasRevNote

`func (o *PostAddStackScriptRequest) HasRevNote() bool`

HasRevNote returns a boolean if a field has been set.

### GetScript

`func (o *PostAddStackScriptRequest) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *PostAddStackScriptRequest) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *PostAddStackScriptRequest) SetScript(v string)`

SetScript sets Script field to given value.


### GetUpdated

`func (o *PostAddStackScriptRequest) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PostAddStackScriptRequest) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PostAddStackScriptRequest) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PostAddStackScriptRequest) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetUserDefinedFields

`func (o *PostAddStackScriptRequest) GetUserDefinedFields() []GetStackScripts200ResponseDataInnerUserDefinedFieldsInner`

GetUserDefinedFields returns the UserDefinedFields field if non-nil, zero value otherwise.

### GetUserDefinedFieldsOk

`func (o *PostAddStackScriptRequest) GetUserDefinedFieldsOk() (*[]GetStackScripts200ResponseDataInnerUserDefinedFieldsInner, bool)`

GetUserDefinedFieldsOk returns a tuple with the UserDefinedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserDefinedFields

`func (o *PostAddStackScriptRequest) SetUserDefinedFields(v []GetStackScripts200ResponseDataInnerUserDefinedFieldsInner)`

SetUserDefinedFields sets UserDefinedFields field to given value.

### HasUserDefinedFields

`func (o *PostAddStackScriptRequest) HasUserDefinedFields() bool`

HasUserDefinedFields returns a boolean if a field has been set.

### GetUserGravatarId

`func (o *PostAddStackScriptRequest) GetUserGravatarId() string`

GetUserGravatarId returns the UserGravatarId field if non-nil, zero value otherwise.

### GetUserGravatarIdOk

`func (o *PostAddStackScriptRequest) GetUserGravatarIdOk() (*string, bool)`

GetUserGravatarIdOk returns a tuple with the UserGravatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGravatarId

`func (o *PostAddStackScriptRequest) SetUserGravatarId(v string)`

SetUserGravatarId sets UserGravatarId field to given value.

### HasUserGravatarId

`func (o *PostAddStackScriptRequest) HasUserGravatarId() bool`

HasUserGravatarId returns a boolean if a field has been set.

### GetUsername

`func (o *PostAddStackScriptRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PostAddStackScriptRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PostAddStackScriptRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PostAddStackScriptRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


