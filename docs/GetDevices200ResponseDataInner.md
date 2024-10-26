# GetDevices200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** | When this Remember Me session was started.  This corresponds to the time of login with the \&quot;Remember Me\&quot; box checked. | [optional] [readonly] 
**Expiry** | Pointer to **time.Time** | When this TrustedDevice session expires.  Sessions typically last 30 days. | [optional] [readonly] 
**Id** | Pointer to **int32** | The unique ID for this TrustedDevice. | [optional] [readonly] 
**LastAuthenticated** | Pointer to **time.Time** | The last time this TrustedDevice was successfully used to authenticate to [login.linode.com](https://login.linode.com). | [optional] [readonly] 
**LastRemoteAddr** | Pointer to **string** | The last IP Address to successfully authenticate with this TrustedDevice. | [optional] [readonly] 
**UserAgent** | Pointer to **string** | The User Agent of the browser that created this TrustedDevice session. | [optional] [readonly] 

## Methods

### NewGetDevices200ResponseDataInner

`func NewGetDevices200ResponseDataInner() *GetDevices200ResponseDataInner`

NewGetDevices200ResponseDataInner instantiates a new GetDevices200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDevices200ResponseDataInnerWithDefaults

`func NewGetDevices200ResponseDataInnerWithDefaults() *GetDevices200ResponseDataInner`

NewGetDevices200ResponseDataInnerWithDefaults instantiates a new GetDevices200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *GetDevices200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *GetDevices200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *GetDevices200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *GetDevices200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExpiry

`func (o *GetDevices200ResponseDataInner) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *GetDevices200ResponseDataInner) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *GetDevices200ResponseDataInner) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *GetDevices200ResponseDataInner) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetId

`func (o *GetDevices200ResponseDataInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDevices200ResponseDataInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDevices200ResponseDataInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetDevices200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastAuthenticated

`func (o *GetDevices200ResponseDataInner) GetLastAuthenticated() time.Time`

GetLastAuthenticated returns the LastAuthenticated field if non-nil, zero value otherwise.

### GetLastAuthenticatedOk

`func (o *GetDevices200ResponseDataInner) GetLastAuthenticatedOk() (*time.Time, bool)`

GetLastAuthenticatedOk returns a tuple with the LastAuthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthenticated

`func (o *GetDevices200ResponseDataInner) SetLastAuthenticated(v time.Time)`

SetLastAuthenticated sets LastAuthenticated field to given value.

### HasLastAuthenticated

`func (o *GetDevices200ResponseDataInner) HasLastAuthenticated() bool`

HasLastAuthenticated returns a boolean if a field has been set.

### GetLastRemoteAddr

`func (o *GetDevices200ResponseDataInner) GetLastRemoteAddr() string`

GetLastRemoteAddr returns the LastRemoteAddr field if non-nil, zero value otherwise.

### GetLastRemoteAddrOk

`func (o *GetDevices200ResponseDataInner) GetLastRemoteAddrOk() (*string, bool)`

GetLastRemoteAddrOk returns a tuple with the LastRemoteAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRemoteAddr

`func (o *GetDevices200ResponseDataInner) SetLastRemoteAddr(v string)`

SetLastRemoteAddr sets LastRemoteAddr field to given value.

### HasLastRemoteAddr

`func (o *GetDevices200ResponseDataInner) HasLastRemoteAddr() bool`

HasLastRemoteAddr returns a boolean if a field has been set.

### GetUserAgent

`func (o *GetDevices200ResponseDataInner) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *GetDevices200ResponseDataInner) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *GetDevices200ResponseDataInner) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *GetDevices200ResponseDataInner) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


