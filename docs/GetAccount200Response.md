# GetAccount200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivePromotions** | Pointer to [**[]GetAccount200ResponseActivePromotionsInner**](GetAccount200ResponseActivePromotionsInner.md) |  | [optional] [readonly] 
**ActiveSince** | Pointer to **time.Time** | The date and time the account was activated. | [optional] [readonly] 
**Address1** | Pointer to **string** | First line of this Account&#39;s billing address. | [optional] 
**Address2** | Pointer to **string** | Second line of this Account&#39;s billing address. | [optional] 
**Balance** | Pointer to **float32** | This Account&#39;s balance, in US dollars. | [optional] [readonly] 
**BalanceUninvoiced** | Pointer to **float32** | This Account&#39;s current estimated invoice in US dollars. This is not your final invoice balance. Transfer charges are not included in the estimate. | [optional] [readonly] 
**BillingSource** | Pointer to **string** | The source of service charges for this Account, as determined by its relationship with Akamai. Accounts that are associated with Akamai-specific customers return a value of &#x60;akamai&#x60;. All other Accounts return a value of &#x60;linode&#x60;. | [optional] [readonly] 
**Capabilities** | Pointer to **[]string** | A list of capabilities your account supports. | [optional] [readonly] 
**City** | Pointer to **string** | The city for this Account&#39;s billing address. | [optional] 
**Company** | Pointer to **string** | The company name associated with this Account.  Must not include any of the following characters: &#x60;&lt;&#x60; &#x60;&gt;&#x60; &#x60;(&#x60; &#x60;)&#x60; &#x60;\&quot;&#x60; &#x60;&#x3D;&#x60; | [optional] 
**Country** | Pointer to **string** | The two-letter ISO 3166 country code of this Account&#39;s billing address. | [optional] 
**CreditCard** | Pointer to [**GetAccount200ResponseCreditCard**](GetAccount200ResponseCreditCard.md) |  | [optional] 
**Email** | Pointer to **string** | The email address of the person associated with this Account. | [optional] 
**Euuid** | Pointer to **string** | An external unique identifier for this account. | [optional] [readonly] 
**FirstName** | Pointer to **string** | The first name of the person associated with this Account.  Must not include any of the following characters: &#x60;&lt;&#x60; &#x60;&gt;&#x60; &#x60;(&#x60; &#x60;)&#x60; &#x60;\&quot;&#x60; &#x60;&#x3D;&#x60; | [optional] 
**LastName** | Pointer to **string** | The last name of the person associated with this Account.  Must not include any of the following characters: &#x60;&lt;&#x60; &#x60;&gt;&#x60; &#x60;(&#x60; &#x60;)&#x60; &#x60;\&quot;&#x60; &#x60;&#x3D;&#x60; | [optional] 
**Phone** | Pointer to **string** | The phone number associated with this Account. | [optional] 
**State** | Pointer to **string** | If billing address is in the United States (US) or Canada (CA), only the two-letter ISO 3166 State or Province code are accepted. If entering a US military address, state abbreviations (AA, AE, AP) should be entered. If the address is outside the US or CA, this is the Province associated with the Account&#39;s billing address. | [optional] 
**TaxId** | Pointer to **string** | The tax identification number associated with this Account, for tax calculations in some countries. If you do not live in a country that collects tax, this should be an empty string (&#x60;\&quot;\&quot;&#x60;). | [optional] 
**Zip** | Pointer to **string** | The zip code of this Account&#39;s billing address. The following restrictions apply:  - Can only contain ASCII letters, numbers, and hyphens (&#x60;-&#x60;). - Must not contain more than 9 letter or number characters. | [optional] 

## Methods

### NewGetAccount200Response

`func NewGetAccount200Response() *GetAccount200Response`

NewGetAccount200Response instantiates a new GetAccount200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccount200ResponseWithDefaults

`func NewGetAccount200ResponseWithDefaults() *GetAccount200Response`

NewGetAccount200ResponseWithDefaults instantiates a new GetAccount200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivePromotions

`func (o *GetAccount200Response) GetActivePromotions() []GetAccount200ResponseActivePromotionsInner`

GetActivePromotions returns the ActivePromotions field if non-nil, zero value otherwise.

### GetActivePromotionsOk

`func (o *GetAccount200Response) GetActivePromotionsOk() (*[]GetAccount200ResponseActivePromotionsInner, bool)`

GetActivePromotionsOk returns a tuple with the ActivePromotions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePromotions

`func (o *GetAccount200Response) SetActivePromotions(v []GetAccount200ResponseActivePromotionsInner)`

SetActivePromotions sets ActivePromotions field to given value.

### HasActivePromotions

`func (o *GetAccount200Response) HasActivePromotions() bool`

HasActivePromotions returns a boolean if a field has been set.

### GetActiveSince

`func (o *GetAccount200Response) GetActiveSince() time.Time`

GetActiveSince returns the ActiveSince field if non-nil, zero value otherwise.

### GetActiveSinceOk

`func (o *GetAccount200Response) GetActiveSinceOk() (*time.Time, bool)`

GetActiveSinceOk returns a tuple with the ActiveSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSince

`func (o *GetAccount200Response) SetActiveSince(v time.Time)`

SetActiveSince sets ActiveSince field to given value.

### HasActiveSince

`func (o *GetAccount200Response) HasActiveSince() bool`

HasActiveSince returns a boolean if a field has been set.

### GetAddress1

`func (o *GetAccount200Response) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *GetAccount200Response) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *GetAccount200Response) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *GetAccount200Response) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *GetAccount200Response) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *GetAccount200Response) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *GetAccount200Response) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *GetAccount200Response) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetBalance

`func (o *GetAccount200Response) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetAccount200Response) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetAccount200Response) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *GetAccount200Response) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetBalanceUninvoiced

`func (o *GetAccount200Response) GetBalanceUninvoiced() float32`

GetBalanceUninvoiced returns the BalanceUninvoiced field if non-nil, zero value otherwise.

### GetBalanceUninvoicedOk

`func (o *GetAccount200Response) GetBalanceUninvoicedOk() (*float32, bool)`

GetBalanceUninvoicedOk returns a tuple with the BalanceUninvoiced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceUninvoiced

`func (o *GetAccount200Response) SetBalanceUninvoiced(v float32)`

SetBalanceUninvoiced sets BalanceUninvoiced field to given value.

### HasBalanceUninvoiced

`func (o *GetAccount200Response) HasBalanceUninvoiced() bool`

HasBalanceUninvoiced returns a boolean if a field has been set.

### GetBillingSource

`func (o *GetAccount200Response) GetBillingSource() string`

GetBillingSource returns the BillingSource field if non-nil, zero value otherwise.

### GetBillingSourceOk

`func (o *GetAccount200Response) GetBillingSourceOk() (*string, bool)`

GetBillingSourceOk returns a tuple with the BillingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSource

`func (o *GetAccount200Response) SetBillingSource(v string)`

SetBillingSource sets BillingSource field to given value.

### HasBillingSource

`func (o *GetAccount200Response) HasBillingSource() bool`

HasBillingSource returns a boolean if a field has been set.

### GetCapabilities

`func (o *GetAccount200Response) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetAccount200Response) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetAccount200Response) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *GetAccount200Response) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCity

`func (o *GetAccount200Response) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetAccount200Response) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetAccount200Response) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GetAccount200Response) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompany

`func (o *GetAccount200Response) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GetAccount200Response) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GetAccount200Response) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GetAccount200Response) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountry

`func (o *GetAccount200Response) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetAccount200Response) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetAccount200Response) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetAccount200Response) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreditCard

`func (o *GetAccount200Response) GetCreditCard() GetAccount200ResponseCreditCard`

GetCreditCard returns the CreditCard field if non-nil, zero value otherwise.

### GetCreditCardOk

`func (o *GetAccount200Response) GetCreditCardOk() (*GetAccount200ResponseCreditCard, bool)`

GetCreditCardOk returns a tuple with the CreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCard

`func (o *GetAccount200Response) SetCreditCard(v GetAccount200ResponseCreditCard)`

SetCreditCard sets CreditCard field to given value.

### HasCreditCard

`func (o *GetAccount200Response) HasCreditCard() bool`

HasCreditCard returns a boolean if a field has been set.

### GetEmail

`func (o *GetAccount200Response) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetAccount200Response) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetAccount200Response) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetAccount200Response) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEuuid

`func (o *GetAccount200Response) GetEuuid() string`

GetEuuid returns the Euuid field if non-nil, zero value otherwise.

### GetEuuidOk

`func (o *GetAccount200Response) GetEuuidOk() (*string, bool)`

GetEuuidOk returns a tuple with the Euuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEuuid

`func (o *GetAccount200Response) SetEuuid(v string)`

SetEuuid sets Euuid field to given value.

### HasEuuid

`func (o *GetAccount200Response) HasEuuid() bool`

HasEuuid returns a boolean if a field has been set.

### GetFirstName

`func (o *GetAccount200Response) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GetAccount200Response) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GetAccount200Response) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *GetAccount200Response) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *GetAccount200Response) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GetAccount200Response) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GetAccount200Response) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *GetAccount200Response) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *GetAccount200Response) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetAccount200Response) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetAccount200Response) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetAccount200Response) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *GetAccount200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetAccount200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetAccount200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetAccount200Response) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTaxId

`func (o *GetAccount200Response) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *GetAccount200Response) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *GetAccount200Response) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *GetAccount200Response) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetZip

`func (o *GetAccount200Response) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *GetAccount200Response) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *GetAccount200Response) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *GetAccount200Response) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


