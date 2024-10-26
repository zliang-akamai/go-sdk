# GetDatabasesMysqlInstanceSsl200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertificate** | Pointer to **string** | The base64-encoded SSL CA certificate for the Managed Database instance. | [optional] 

## Methods

### NewGetDatabasesMysqlInstanceSsl200Response

`func NewGetDatabasesMysqlInstanceSsl200Response() *GetDatabasesMysqlInstanceSsl200Response`

NewGetDatabasesMysqlInstanceSsl200Response instantiates a new GetDatabasesMysqlInstanceSsl200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDatabasesMysqlInstanceSsl200ResponseWithDefaults

`func NewGetDatabasesMysqlInstanceSsl200ResponseWithDefaults() *GetDatabasesMysqlInstanceSsl200Response`

NewGetDatabasesMysqlInstanceSsl200ResponseWithDefaults instantiates a new GetDatabasesMysqlInstanceSsl200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertificate

`func (o *GetDatabasesMysqlInstanceSsl200Response) GetCaCertificate() string`

GetCaCertificate returns the CaCertificate field if non-nil, zero value otherwise.

### GetCaCertificateOk

`func (o *GetDatabasesMysqlInstanceSsl200Response) GetCaCertificateOk() (*string, bool)`

GetCaCertificateOk returns a tuple with the CaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertificate

`func (o *GetDatabasesMysqlInstanceSsl200Response) SetCaCertificate(v string)`

SetCaCertificate sets CaCertificate field to given value.

### HasCaCertificate

`func (o *GetDatabasesMysqlInstanceSsl200Response) HasCaCertificate() bool`

HasCaCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


