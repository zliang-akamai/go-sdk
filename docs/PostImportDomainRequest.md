# PostImportDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | **string** | The domain to import. | 
**RemoteNameserver** | **string** | The remote nameserver that allows zone transfers (AXFR). | 

## Methods

### NewPostImportDomainRequest

`func NewPostImportDomainRequest(domain string, remoteNameserver string, ) *PostImportDomainRequest`

NewPostImportDomainRequest instantiates a new PostImportDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostImportDomainRequestWithDefaults

`func NewPostImportDomainRequestWithDefaults() *PostImportDomainRequest`

NewPostImportDomainRequestWithDefaults instantiates a new PostImportDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *PostImportDomainRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *PostImportDomainRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *PostImportDomainRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.


### GetRemoteNameserver

`func (o *PostImportDomainRequest) GetRemoteNameserver() string`

GetRemoteNameserver returns the RemoteNameserver field if non-nil, zero value otherwise.

### GetRemoteNameserverOk

`func (o *PostImportDomainRequest) GetRemoteNameserverOk() (*string, bool)`

GetRemoteNameserverOk returns a tuple with the RemoteNameserver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteNameserver

`func (o *PostImportDomainRequest) SetRemoteNameserver(v string)`

SetRemoteNameserver sets RemoteNameserver field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


