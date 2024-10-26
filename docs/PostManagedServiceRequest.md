# PostManagedServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | The URL at which this Service is monitored. URL parameters such as &#x60;?no-cache&#x3D;1&#x60; are preserved. URL fragments/anchors such as &#x60;#monitor&#x60; are __not__ preserved. | 
**Body** | Pointer to **NullableString** | What to expect to find in the response body for the Service to be considered up. | [optional] 
**ConsultationGroup** | Pointer to **string** | The group of ManagedContacts who should be notified or consulted with when an Issue is detected. | [optional] 
**Created** | Pointer to **time.Time** | When this Managed Service was created. | [optional] [readonly] 
**Credentials** | Pointer to **[]int32** | An array of ManagedCredential IDs that should be used when attempting to resolve issues with this Service. | [optional] 
**Id** | Pointer to **int32** | This Service&#39;s unique ID. | [optional] [readonly] 
**Label** | **string** | The label for this Service. This is for display purposes only. | 
**Notes** | Pointer to **NullableString** | Any information relevant to the Service that Linode special forces should know when attempting to resolve Issues. | [optional] 
**Region** | Pointer to **string** | The Region in which this Service is located. This is required if address is a private IP, and may not be set otherwise. | [optional] 
**ServiceType** | **string** | How this Service is monitored. | 
**Status** | Pointer to **string** | The current status of this Service. | [optional] [readonly] 
**Timeout** | **int32** | How long to wait, in seconds, for a response before considering the Service to be down. | 
**Updated** | Pointer to **time.Time** | When this Managed Service was last updated. | [optional] [readonly] 

## Methods

### NewPostManagedServiceRequest

`func NewPostManagedServiceRequest(address string, label string, serviceType string, timeout int32, ) *PostManagedServiceRequest`

NewPostManagedServiceRequest instantiates a new PostManagedServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostManagedServiceRequestWithDefaults

`func NewPostManagedServiceRequestWithDefaults() *PostManagedServiceRequest`

NewPostManagedServiceRequestWithDefaults instantiates a new PostManagedServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PostManagedServiceRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PostManagedServiceRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PostManagedServiceRequest) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetBody

`func (o *PostManagedServiceRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *PostManagedServiceRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *PostManagedServiceRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *PostManagedServiceRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *PostManagedServiceRequest) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *PostManagedServiceRequest) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetConsultationGroup

`func (o *PostManagedServiceRequest) GetConsultationGroup() string`

GetConsultationGroup returns the ConsultationGroup field if non-nil, zero value otherwise.

### GetConsultationGroupOk

`func (o *PostManagedServiceRequest) GetConsultationGroupOk() (*string, bool)`

GetConsultationGroupOk returns a tuple with the ConsultationGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsultationGroup

`func (o *PostManagedServiceRequest) SetConsultationGroup(v string)`

SetConsultationGroup sets ConsultationGroup field to given value.

### HasConsultationGroup

`func (o *PostManagedServiceRequest) HasConsultationGroup() bool`

HasConsultationGroup returns a boolean if a field has been set.

### GetCreated

`func (o *PostManagedServiceRequest) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PostManagedServiceRequest) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PostManagedServiceRequest) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PostManagedServiceRequest) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetCredentials

`func (o *PostManagedServiceRequest) GetCredentials() []int32`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *PostManagedServiceRequest) GetCredentialsOk() (*[]int32, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *PostManagedServiceRequest) SetCredentials(v []int32)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *PostManagedServiceRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetId

`func (o *PostManagedServiceRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostManagedServiceRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostManagedServiceRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PostManagedServiceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *PostManagedServiceRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PostManagedServiceRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PostManagedServiceRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNotes

`func (o *PostManagedServiceRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *PostManagedServiceRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *PostManagedServiceRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *PostManagedServiceRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *PostManagedServiceRequest) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *PostManagedServiceRequest) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetRegion

`func (o *PostManagedServiceRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *PostManagedServiceRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *PostManagedServiceRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *PostManagedServiceRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetServiceType

`func (o *PostManagedServiceRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PostManagedServiceRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PostManagedServiceRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetStatus

`func (o *PostManagedServiceRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostManagedServiceRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostManagedServiceRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostManagedServiceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTimeout

`func (o *PostManagedServiceRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *PostManagedServiceRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *PostManagedServiceRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.


### GetUpdated

`func (o *PostManagedServiceRequest) GetUpdated() time.Time`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PostManagedServiceRequest) GetUpdatedOk() (*time.Time, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PostManagedServiceRequest) SetUpdated(v time.Time)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PostManagedServiceRequest) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


