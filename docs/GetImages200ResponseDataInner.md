# GetImages200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | A list of the possible capabilities of this image.  - &#x60;cloud-init&#x60;. The image supports the cloud-init multi-distribution method with our [Metadata service](https://www.linode.com/docs/products/compute/compute-instances/guides/metadata/#troubleshoot-metadata-and-cloud-init). This only applies to public images.  - &#x60;distributed-sites&#x60;. Whether the image can be used in distributed compute regions. Compared to a core compute region, distributed compute regions offer limited functionality, but they&#39;re globally distributed. Your image can be geographically closer to you, potentially letting you deploy it quicker. See [Regions and images](https://techdocs.akamai.com/cloud-computing/docs/images#regions-and-images) for complete details. | [optional] [readonly] 
**Created** | Pointer to **time.Time** | When this image was created. | [optional] [readonly] 
**CreatedBy** | Pointer to **string** | The name of the user who created this image, or &#x60;linode&#x60; for public images. | [optional] [readonly] 
**Deprecated** | Pointer to **bool** | Whether this image is deprecated. Only public images can be deprecated. | [optional] [readonly] 
**Description** | Pointer to **NullableString** | A detailed description of this image. | [optional] 
**Eol** | Pointer to **NullableTime** | The date of the public image&#39;s planned removal from service. This is &#x60;null&#x60; for private images. | [optional] [readonly] 
**Expiry** | Pointer to **NullableTime** | Only images created automatically from a deleted compute instance (type&#x3D;automatic) expire. This is &#x60;null&#x60; for private images. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique identifier for each image. | [optional] [readonly] 
**IsPublic** | Pointer to **bool** | Revealed as &#x60;true&#x60; if the image is a public distribution image. Private, account-specific images are listed as &#x60;false&#x60;. | [optional] [readonly] 
**Label** | Pointer to **string** | A short description of the image. | [optional] 
**Regions** | Pointer to [**[]GetImages200ResponseDataInnerRegionsInner**](GetImages200ResponseDataInnerRegionsInner.md) | Details on the regions where this image is stored. See [Regions and images](https://techdocs.akamai.com/cloud-computing/docs/images#regions-and-images) for full details on support for &#x60;regions&#x60;. | [optional] [readonly] 
**Size** | Pointer to **int32** | The minimum size in MB this image needs to deploy. | [optional] [readonly] 
**Status** | Pointer to **string** | The current status of the image. Possible values are &#x60;available&#x60;, &#x60;creating&#x60;, and &#x60;pending_upload&#x60;.  &gt; 📘 &gt; &gt; The &#x60;+order_by&#x60; and &#x60;+order&#x60; operators are not available when [filtering](https://techdocs.akamai.com/linode-api/reference/filtering-and-sorting) on this key. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | Tags used for organizational purposes. A tag can be from 3 to 100 characters long, and an image can have a maximum of 500 total tags. | [optional] 
**TotalSize** | Pointer to **int32** | The total size in bytes of all instances of this image, in all &#x60;regions&#x60;.  &gt; 📘 &gt; &gt; This object is empty for existing images. It&#39;s intended for use with future functionality. | [optional] [readonly] 
**Type** | Pointer to **string** | How the image was created. Create a &#x60;manual&#x60; image at any time. An &#x60;automatic&#x60; image is created automatically from a deleted compute instance. | [optional] [readonly] 
**Updated** | Pointer to **time.Time** | When this image was last updated. | [optional] [readonly] 
**Vendor** | Pointer to **NullableString** | The upstream distribution vendor. This is &#x60;null&#x60; for private images. | [optional] [readonly] 

## Methods

### NewGetImages200ResponseDataInner

`func NewGetImages200ResponseDataInner() *GetImages200ResponseDataInner`

NewGetImages200ResponseDataInner instantiates a new GetImages200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImages200ResponseDataInnerWithDefaults

`func NewGetImages200ResponseDataInnerWithDefaults() *GetImages200ResponseDataInner`

NewGetImages200ResponseDataInnerWithDefaults instantiates a new GetImages200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *GetImages200ResponseDataInner) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetImages200ResponseDataInner) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetImages200ResponseDataInner) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *GetImages200ResponseDataInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCreated

`func (o *GetImages200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetImages200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetImages200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetImages200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetImages200ResponseDataInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetImages200ResponseDataInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetImages200ResponseDataInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetImages200ResponseDataInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDeprecated

`func (o *GetImages200ResponseDataInner) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *GetImages200ResponseDataInner) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *GetImages200ResponseDataInner) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *GetImages200ResponseDataInner) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDescription

`func (o *GetImages200ResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetImages200ResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetImages200ResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetImages200ResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetImages200ResponseDataInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetImages200ResponseDataInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEol

`func (o *GetImages200ResponseDataInner) GetEol() time.Time`

GetEol returns the Eol field if non-nil, zero value otherwise.

### GetEolOk

`func (o *GetImages200ResponseDataInner) GetEolOk() (*time.Time, bool)`

GetEolOk returns a tuple with the Eol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEol

`func (o *GetImages200ResponseDataInner) SetEol(v time.Time)`

SetEol sets Eol field to given value.

### HasEol

`func (o *GetImages200ResponseDataInner) HasEol() bool`

HasEol returns a boolean if a field has been set.

### SetEolNil

`func (o *GetImages200ResponseDataInner) SetEolNil(b bool)`

 SetEolNil sets the value for Eol to be an explicit nil

### UnsetEol
`func (o *GetImages200ResponseDataInner) UnsetEol()`

UnsetEol ensures that no value is present for Eol, not even an explicit nil
### GetExpiry

`func (o *GetImages200ResponseDataInner) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GetImages200ResponseDataInner) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GetImages200ResponseDataInner) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GetImages200ResponseDataInner) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### SetExpiryNil

`func (o *GetImages200ResponseDataInner) SetExpiryNil(b bool)`

 SetExpiryNil sets the value for Expiry to be an explicit nil

### UnsetExpiry
`func (o *GetImages200ResponseDataInner) UnsetExpiry()`

UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
### GetId

`func (o *GetImages200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetImages200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetImages200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetImages200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPublic

`func (o *GetImages200ResponseDataInner) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *GetImages200ResponseDataInner) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *GetImages200ResponseDataInner) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *GetImages200ResponseDataInner) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetLabel

`func (o *GetImages200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetImages200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetImages200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetImages200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetRegions

`func (o *GetImages200ResponseDataInner) GetRegions() []GetImages200ResponseDataInnerRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *GetImages200ResponseDataInner) GetRegionsOk() (*[]GetImages200ResponseDataInnerRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *GetImages200ResponseDataInner) SetRegions(v []GetImages200ResponseDataInnerRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *GetImages200ResponseDataInner) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetSize

`func (o *GetImages200ResponseDataInner) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetImages200ResponseDataInner) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetImages200ResponseDataInner) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetImages200ResponseDataInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *GetImages200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetImages200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetImages200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetImages200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *GetImages200ResponseDataInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetImages200ResponseDataInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetImages200ResponseDataInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetImages200ResponseDataInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTotalSize

`func (o *GetImages200ResponseDataInner) GetTotalSize() int32`

GetTotalSize returns the TotalSize field if non-nil, zero value otherwise.

### GetTotalSizeOk

`func (o *GetImages200ResponseDataInner) GetTotalSizeOk() (*int32, bool)`

GetTotalSizeOk returns a tuple with the TotalSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSize

`func (o *GetImages200ResponseDataInner) SetTotalSize(v int32)`

SetTotalSize sets TotalSize field to given value.

### HasTotalSize

`func (o *GetImages200ResponseDataInner) HasTotalSize() bool`

HasTotalSize returns a boolean if a field has been set.

### GetType

`func (o *GetImages200ResponseDataInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetImages200ResponseDataInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetImages200ResponseDataInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetImages200ResponseDataInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *GetImages200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetImages200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetImages200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetImages200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetVendor

`func (o *GetImages200ResponseDataInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GetImages200ResponseDataInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GetImages200ResponseDataInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *GetImages200ResponseDataInner) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### SetVendorNil

`func (o *GetImages200ResponseDataInner) SetVendorNil(b bool)`

 SetVendorNil sets the value for Vendor to be an explicit nil

### UnsetVendor
`func (o *GetImages200ResponseDataInner) UnsetVendor()`

UnsetVendor ensures that no value is present for Vendor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


