# GetAccountAgreements200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EuModel** | Pointer to **bool** | The acknowledgement status for the [cross-border data transfer](https://www.akamai.com/legal/compliance/privacy-trust-center/cross-border-data-transfer-statement) agreement. | [optional] 
**MasterServiceAgreement** | Pointer to **bool** | The acknowledgement status for Akamai&#39;s [master service agreement](https://www.linode.com/legal-msa/). | [optional] 
**PrivacyPolicy** | Pointer to **bool** | The acknowledgement status for Akamai&#39;s [privacy statement](https://www.akamai.com/legal/privacy-statement). | [optional] 

## Methods

### NewGetAccountAgreements200Response

`func NewGetAccountAgreements200Response() *GetAccountAgreements200Response`

NewGetAccountAgreements200Response instantiates a new GetAccountAgreements200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountAgreements200ResponseWithDefaults

`func NewGetAccountAgreements200ResponseWithDefaults() *GetAccountAgreements200Response`

NewGetAccountAgreements200ResponseWithDefaults instantiates a new GetAccountAgreements200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEuModel

`func (o *GetAccountAgreements200Response) GetEuModel() bool`

GetEuModel returns the EuModel field if non-nil, zero value otherwise.

### GetEuModelOk

`func (o *GetAccountAgreements200Response) GetEuModelOk() (*bool, bool)`

GetEuModelOk returns a tuple with the EuModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEuModel

`func (o *GetAccountAgreements200Response) SetEuModel(v bool)`

SetEuModel sets EuModel field to given value.

### HasEuModel

`func (o *GetAccountAgreements200Response) HasEuModel() bool`

HasEuModel returns a boolean if a field has been set.

### GetMasterServiceAgreement

`func (o *GetAccountAgreements200Response) GetMasterServiceAgreement() bool`

GetMasterServiceAgreement returns the MasterServiceAgreement field if non-nil, zero value otherwise.

### GetMasterServiceAgreementOk

`func (o *GetAccountAgreements200Response) GetMasterServiceAgreementOk() (*bool, bool)`

GetMasterServiceAgreementOk returns a tuple with the MasterServiceAgreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterServiceAgreement

`func (o *GetAccountAgreements200Response) SetMasterServiceAgreement(v bool)`

SetMasterServiceAgreement sets MasterServiceAgreement field to given value.

### HasMasterServiceAgreement

`func (o *GetAccountAgreements200Response) HasMasterServiceAgreement() bool`

HasMasterServiceAgreement returns a boolean if a field has been set.

### GetPrivacyPolicy

`func (o *GetAccountAgreements200Response) GetPrivacyPolicy() bool`

GetPrivacyPolicy returns the PrivacyPolicy field if non-nil, zero value otherwise.

### GetPrivacyPolicyOk

`func (o *GetAccountAgreements200Response) GetPrivacyPolicyOk() (*bool, bool)`

GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicy

`func (o *GetAccountAgreements200Response) SetPrivacyPolicy(v bool)`

SetPrivacyPolicy sets PrivacyPolicy field to given value.

### HasPrivacyPolicy

`func (o *GetAccountAgreements200Response) HasPrivacyPolicy() bool`

HasPrivacyPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


