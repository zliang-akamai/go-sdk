# GetManagedServices200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | The URL at which this Service is monitored. URL parameters such as &#x60;?no-cache&#x3D;1&#x60; are preserved. URL fragments/anchors such as &#x60;#monitor&#x60; are __not__ preserved. | [optional] 
**Body** | Pointer to **NullableString** | What to expect to find in the response body for the Service to be considered up. | [optional] 
**ConsultationGroup** | Pointer to **string** | The group of ManagedContacts who should be notified or consulted with when an Issue is detected. | [optional] 
**Created** | Pointer to **time.Time** | When this Managed Service was created. | [optional] [readonly] 
**Credentials** | Pointer to **[]int32** | An array of ManagedCredential IDs that should be used when attempting to resolve issues with this Service. | [optional] 
**Id** | Pointer to **int32** | This Service&#39;s unique ID. | [optional] [readonly] 
**Label** | Pointer to **string** | The label for this Service. This is for display purposes only. | [optional] 
**Notes** | Pointer to **NullableString** | Any information relevant to the Service that Linode special forces should know when attempting to resolve Issues. | [optional] 
**Region** | Pointer to **string** | The Region in which this Service is located. This is required if address is a private IP, and may not be set otherwise. | [optional] 
**ServiceType** | Pointer to **string** | How this Service is monitored. | [optional] 
**Status** | Pointer to **string** | The current status of this Service. | [optional] [readonly] 
**Timeout** | Pointer to **int32** | How long to wait, in seconds, for a response before considering the Service to be down. | [optional] 
**Updated** | Pointer to **time.Time** | When this Managed Service was last updated. | [optional] [readonly] 

## Methods

### NewGetManagedServices200ResponseDataInner

`func NewGetManagedServices200ResponseDataInner() *GetManagedServices200ResponseDataInner`

NewGetManagedServices200ResponseDataInner instantiates a new GetManagedServices200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetManagedServices200ResponseDataInnerWithDefaults

`func NewGetManagedServices200ResponseDataInnerWithDefaults() *GetManagedServices200ResponseDataInner`

NewGetManagedServices200ResponseDataInnerWithDefaults instantiates a new GetManagedServices200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GetManagedServices200ResponseDataInner) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetManagedServices200ResponseDataInner) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetManagedServices200ResponseDataInner) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetManagedServices200ResponseDataInner) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBody

`func (o *GetManagedServices200ResponseDataInner) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetManagedServices200ResponseDataInner) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetManagedServices200ResponseDataInner) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GetManagedServices200ResponseDataInner) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *GetManagedServices200ResponseDataInner) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *GetManagedServices200ResponseDataInner) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetConsultationGroup

`func (o *GetManagedServices200ResponseDataInner) GetConsultationGroup() string`

GetConsultationGroup returns the ConsultationGroup field if non-nil, zero value otherwise.

### GetConsultationGroupOk

`func (o *GetManagedServices200ResponseDataInner) GetConsultationGroupOk() (*string, bool)`

GetConsultationGroupOk returns a tuple with the ConsultationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsultationGroup

`func (o *GetManagedServices200ResponseDataInner) SetConsultationGroup(v string)`

SetConsultationGroup sets ConsultationGroup field to given value.

### HasConsultationGroup

`func (o *GetManagedServices200ResponseDataInner) HasConsultationGroup() bool`

HasConsultationGroup returns a boolean if a field has been set.

### GetCreated

`func (o *GetManagedServices200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetManagedServices200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetManagedServices200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetManagedServices200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentials

`func (o *GetManagedServices200ResponseDataInner) GetCredentials() []int32`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GetManagedServices200ResponseDataInner) GetCredentialsOk() (*[]int32, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GetManagedServices200ResponseDataInner) SetCredentials(v []int32)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GetManagedServices200ResponseDataInner) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetId

`func (o *GetManagedServices200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetManagedServices200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetManagedServices200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetManagedServices200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GetManagedServices200ResponseDataInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetManagedServices200ResponseDataInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetManagedServices200ResponseDataInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetManagedServices200ResponseDataInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetNotes

`func (o *GetManagedServices200ResponseDataInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetManagedServices200ResponseDataInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetManagedServices200ResponseDataInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetManagedServices200ResponseDataInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *GetManagedServices200ResponseDataInner) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *GetManagedServices200ResponseDataInner) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetRegion

`func (o *GetManagedServices200ResponseDataInner) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *GetManagedServices200ResponseDataInner) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *GetManagedServices200ResponseDataInner) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *GetManagedServices200ResponseDataInner) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServiceType

`func (o *GetManagedServices200ResponseDataInner) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *GetManagedServices200ResponseDataInner) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *GetManagedServices200ResponseDataInner) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *GetManagedServices200ResponseDataInner) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetStatus

`func (o *GetManagedServices200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetManagedServices200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetManagedServices200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetManagedServices200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeout

`func (o *GetManagedServices200ResponseDataInner) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *GetManagedServices200ResponseDataInner) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *GetManagedServices200ResponseDataInner) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *GetManagedServices200ResponseDataInner) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUpdated

`func (o *GetManagedServices200ResponseDataInner) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *GetManagedServices200ResponseDataInner) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *GetManagedServices200ResponseDataInner) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *GetManagedServices200ResponseDataInner) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


