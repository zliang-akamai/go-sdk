# GetRegions200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | A list of capabilities of this region. | [optional] [readonly] 
**Country** | Pointer to **string** | The country where this Region resides. | [optional] [readonly] 
**Id** | Pointer to **string** | The unique ID of this Region. | [optional] [readonly] 
**Label** | Pointer to **string** | Detailed location information for this Region, including city, state or region, and country. | [optional] [readonly] 
**PlacementGroupLimits** | Pointer to [**GetRegions200ResponseDataInnerPlacementGroupLimits**](GetRegions200ResponseDataInnerPlacementGroupLimits.md) |  | [optional] 
**Resolvers** | Pointer to [**GetRegions200ResponseDataInnerResolvers**](GetRegions200ResponseDataInnerResolvers.md) |  | [optional] 
**SiteType** | Pointer to **string** | This region&#39;s site type. A &#x60;core&#x60; region indicates a traditional cloud computing [region](https://www.linode.com/docs/products/platform/get-started/guides/choose-a-data-center/#product-availability) that offers all compute services. A &#x60;distributed&#x60; region indicates sites that are globally dispersed to be closer to end users and workloads. These regions offer limited services. | [optional] [readonly] 
**Status** | Pointer to **string** | This region&#39;s current operational status. | [optional] [readonly] 

## Methods

### NewGetRegions200ResponseDataInner

`func NewGetRegions200ResponseDataInner() *GetRegions200ResponseDataInner`

NewGetRegions200ResponseDataInner instantiates a new GetRegions200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRegions200ResponseDataInnerWithDefaults

`func NewGetRegions200ResponseDataInnerWithDefaults() *GetRegions200ResponseDataInner`

NewGetRegions200ResponseDataInnerWithDefaults instantiates a new GetRegions200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *GetRegions200ResponseDataInner) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetRegions200ResponseDataInner) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetRegions200ResponseDataInner) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *GetRegions200ResponseDataInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCountry

`func (o *GetRegions200ResponseDataInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetRegions200ResponseDataInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetRegions200ResponseDataInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetRegions200ResponseDataInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetId

`func (o *GetRegions200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRegions200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRegions200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetRegions200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetRegions200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetRegions200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetRegions200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetRegions200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPlacementGroupLimits

`func (o *GetRegions200ResponseDataInner) GetPlacementGroupLimits() GetRegions200ResponseDataInnerPlacementGroupLimits`

GetPlacementGroupLimits returns the PlacementGroupLimits field if non-nil, zero value otherwise.

### GetPlacementGroupLimitsOk

`func (o *GetRegions200ResponseDataInner) GetPlacementGroupLimitsOk() (*GetRegions200ResponseDataInnerPlacementGroupLimits, bool)`

GetPlacementGroupLimitsOk returns a tuple with the PlacementGroupLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacementGroupLimits

`func (o *GetRegions200ResponseDataInner) SetPlacementGroupLimits(v GetRegions200ResponseDataInnerPlacementGroupLimits)`

SetPlacementGroupLimits sets PlacementGroupLimits field to given value.

### HasPlacementGroupLimits

`func (o *GetRegions200ResponseDataInner) HasPlacementGroupLimits() bool`

HasPlacementGroupLimits returns a boolean if a field has been set.

### GetResolvers

`func (o *GetRegions200ResponseDataInner) GetResolvers() GetRegions200ResponseDataInnerResolvers`

GetResolvers returns the Resolvers field if non-nil, zero value otherwise.

### GetResolversOk

`func (o *GetRegions200ResponseDataInner) GetResolversOk() (*GetRegions200ResponseDataInnerResolvers, bool)`

GetResolversOk returns a tuple with the Resolvers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvers

`func (o *GetRegions200ResponseDataInner) SetResolvers(v GetRegions200ResponseDataInnerResolvers)`

SetResolvers sets Resolvers field to given value.

### HasResolvers

`func (o *GetRegions200ResponseDataInner) HasResolvers() bool`

HasResolvers returns a boolean if a field has been set.

### GetSiteType

`func (o *GetRegions200ResponseDataInner) GetSiteType() string`

GetSiteType returns the SiteType field if non-nil, zero value otherwise.

### GetSiteTypeOk

`func (o *GetRegions200ResponseDataInner) GetSiteTypeOk() (*string, bool)`

GetSiteTypeOk returns a tuple with the SiteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteType

`func (o *GetRegions200ResponseDataInner) SetSiteType(v string)`

SetSiteType sets SiteType field to given value.

### HasSiteType

`func (o *GetRegions200ResponseDataInner) HasSiteType() bool`

HasSiteType returns a boolean if a field has been set.

### GetStatus

`func (o *GetRegions200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetRegions200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetRegions200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetRegions200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


