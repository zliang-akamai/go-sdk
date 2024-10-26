# GetChildAccounts200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveSince** | Pointer to **time.Time** | The activation date and time for the child account. | [optional] [readonly] 
**Address1** | Pointer to **string** | First line of this child account&#39;s billing address. | [optional] 
**Address2** | Pointer to **string** | Second line of this child account&#39;s billing address, if applicable. | [optional] 
**Balance** | Pointer to **float32** | This child account&#39;s balance, in US dollars. | [optional] [readonly] 
**BalanceUninvoiced** | Pointer to **float32** | This child account&#39;s current estimated invoice in US dollars. This is not your final invoice balance. Transfer charges are not included in the estimate. | [optional] [readonly] 
**BillingSource** | Pointer to **string** | The source of service charges for this account, as determined by its relationship with Akamai. The API returns a value of &#x60;external&#x60; to describe a child account in a parent-child account environment. | [optional] [readonly] 
**Capabilities** | Pointer to **[]string** | A list of the capabilities the child account supports. | [optional] [readonly] 
**City** | Pointer to **string** | The city for this child account&#39;s billing address. | [optional] 
**Company** | Pointer to **string** | The company name for the owner of this child account. It can&#39;t include any of these characters: &#x60;&lt;&#x60; &#x60;&gt;&#x60; &#x60;(&#x60; &#x60;)&#x60; &#x60;\&quot;&#x60; &#x60;&#x3D;&#x60;. You can&#39;t change this value yourself. We use it to create the proxy users that a parent account uses to access a child account. Talk to your account team if you need to change this value. | [optional] 
**Country** | Pointer to **string** | The two-letter ISO 3166 country code for this child account&#39;s billing address. | [optional] 
**CreditCard** | Pointer to [**GetChildAccounts200ResponseDataInnerCreditCard**](GetChildAccounts200ResponseDataInnerCreditCard.md) |  | [optional] 
**Email** | Pointer to **string** | The email address of the owner of this child account. | [optional] 
**Euuid** | Pointer to **string** | An external, unique identifier that Akamai assigned to the child account. | [optional] [readonly] 
**FirstName** | Pointer to **string** | The first name of the owner of this child account. It can&#39;t include any of these characters: &#x60;&lt;&#x60; &#x60;&gt;&#x60; &#x60;(&#x60; &#x60;)&#x60; &#x60;\&quot;&#x60; &#x60;&#x3D;&#x60;. | [optional] 
**LastName** | Pointer to **string** | The last name of the owner of this child account. It can&#39;t include any of these characters: &#x60;&lt;&#x60; &#x60;&gt;&#x60; &#x60;(&#x60; &#x60;)&#x60; &#x60;\&quot;&#x60; &#x60;&#x3D;&#x60;. | [optional] 
**Phone** | Pointer to **string** | The phone number for the owner of this child account. | [optional] 
**State** | Pointer to **string** | The state or province for the billing address (&#x60;address_1&#x60; and &#x60;address_2, if applicable&#x60;). If in the United States (US) or Canada (CA), this is the two-letter ISO 3166 State or Province code.  __Note__. If this is a US military address, use state abbreviations (AA, AE, AP). | [optional] 
**TaxId** | Pointer to **string** | The tax identification number for this child account. Use this for tax calculations in some countries. If you live in a country that doesn&#39;t collect taxes, ensure this is an empty string (&#x60;\&quot;\&quot;&#x60;). | [optional] 
**Zip** | Pointer to **string** | The zip code of this Account&#39;s billing address. The following restrictions apply:  - Can only contain ASCII letters, numbers, and hyphens (&#x60;-&#x60;). - Can&#39;t contain more than 9 letter or number characters. | [optional] 

## Methods

### NewGetChildAccounts200ResponseDataInner

`func NewGetChildAccounts200ResponseDataInner() *GetChildAccounts200ResponseDataInner`

NewGetChildAccounts200ResponseDataInner instantiates a new GetChildAccounts200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetChildAccounts200ResponseDataInnerWithDefaults

`func NewGetChildAccounts200ResponseDataInnerWithDefaults() *GetChildAccounts200ResponseDataInner`

NewGetChildAccounts200ResponseDataInnerWithDefaults instantiates a new GetChildAccounts200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveSince

`func (o *GetChildAccounts200ResponseDataInner) GetActiveSince() time.Time`

GetActiveSince returns the ActiveSince field if non-nil, zero value otherwise.

### GetActiveSinceOk

`func (o *GetChildAccounts200ResponseDataInner) GetActiveSinceOk() (*time.Time, bool)`

GetActiveSinceOk returns a tuple with the ActiveSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveSince

`func (o *GetChildAccounts200ResponseDataInner) SetActiveSince(v time.Time)`

SetActiveSince sets ActiveSince field to given value.

### HasActiveSince

`func (o *GetChildAccounts200ResponseDataInner) HasActiveSince() bool`

HasActiveSince returns a boolean if a field has been set.

### GetAddress1

`func (o *GetChildAccounts200ResponseDataInner) GetAddress1() string`

GetAddress1 returns the Address1 field if non-nil, zero value otherwise.

### GetAddress1Ok

`func (o *GetChildAccounts200ResponseDataInner) GetAddress1Ok() (*string, bool)`

GetAddress1Ok returns a tuple with the Address1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress1

`func (o *GetChildAccounts200ResponseDataInner) SetAddress1(v string)`

SetAddress1 sets Address1 field to given value.

### HasAddress1

`func (o *GetChildAccounts200ResponseDataInner) HasAddress1() bool`

HasAddress1 returns a boolean if a field has been set.

### GetAddress2

`func (o *GetChildAccounts200ResponseDataInner) GetAddress2() string`

GetAddress2 returns the Address2 field if non-nil, zero value otherwise.

### GetAddress2Ok

`func (o *GetChildAccounts200ResponseDataInner) GetAddress2Ok() (*string, bool)`

GetAddress2Ok returns a tuple with the Address2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress2

`func (o *GetChildAccounts200ResponseDataInner) SetAddress2(v string)`

SetAddress2 sets Address2 field to given value.

### HasAddress2

`func (o *GetChildAccounts200ResponseDataInner) HasAddress2() bool`

HasAddress2 returns a boolean if a field has been set.

### GetBalance

`func (o *GetChildAccounts200ResponseDataInner) GetBalance() float32`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *GetChildAccounts200ResponseDataInner) GetBalanceOk() (*float32, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *GetChildAccounts200ResponseDataInner) SetBalance(v float32)`

SetBalance sets Balance field to given value.

### HasBalance

`func (o *GetChildAccounts200ResponseDataInner) HasBalance() bool`

HasBalance returns a boolean if a field has been set.

### GetBalanceUninvoiced

`func (o *GetChildAccounts200ResponseDataInner) GetBalanceUninvoiced() float32`

GetBalanceUninvoiced returns the BalanceUninvoiced field if non-nil, zero value otherwise.

### GetBalanceUninvoicedOk

`func (o *GetChildAccounts200ResponseDataInner) GetBalanceUninvoicedOk() (*float32, bool)`

GetBalanceUninvoicedOk returns a tuple with the BalanceUninvoiced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceUninvoiced

`func (o *GetChildAccounts200ResponseDataInner) SetBalanceUninvoiced(v float32)`

SetBalanceUninvoiced sets BalanceUninvoiced field to given value.

### HasBalanceUninvoiced

`func (o *GetChildAccounts200ResponseDataInner) HasBalanceUninvoiced() bool`

HasBalanceUninvoiced returns a boolean if a field has been set.

### GetBillingSource

`func (o *GetChildAccounts200ResponseDataInner) GetBillingSource() string`

GetBillingSource returns the BillingSource field if non-nil, zero value otherwise.

### GetBillingSourceOk

`func (o *GetChildAccounts200ResponseDataInner) GetBillingSourceOk() (*string, bool)`

GetBillingSourceOk returns a tuple with the BillingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSource

`func (o *GetChildAccounts200ResponseDataInner) SetBillingSource(v string)`

SetBillingSource sets BillingSource field to given value.

### HasBillingSource

`func (o *GetChildAccounts200ResponseDataInner) HasBillingSource() bool`

HasBillingSource returns a boolean if a field has been set.

### GetCapabilities

`func (o *GetChildAccounts200ResponseDataInner) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetChildAccounts200ResponseDataInner) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetChildAccounts200ResponseDataInner) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *GetChildAccounts200ResponseDataInner) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCity

`func (o *GetChildAccounts200ResponseDataInner) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *GetChildAccounts200ResponseDataInner) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *GetChildAccounts200ResponseDataInner) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *GetChildAccounts200ResponseDataInner) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompany

`func (o *GetChildAccounts200ResponseDataInner) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GetChildAccounts200ResponseDataInner) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GetChildAccounts200ResponseDataInner) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GetChildAccounts200ResponseDataInner) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetCountry

`func (o *GetChildAccounts200ResponseDataInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetChildAccounts200ResponseDataInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetChildAccounts200ResponseDataInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetChildAccounts200ResponseDataInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreditCard

`func (o *GetChildAccounts200ResponseDataInner) GetCreditCard() GetChildAccounts200ResponseDataInnerCreditCard`

GetCreditCard returns the CreditCard field if non-nil, zero value otherwise.

### GetCreditCardOk

`func (o *GetChildAccounts200ResponseDataInner) GetCreditCardOk() (*GetChildAccounts200ResponseDataInnerCreditCard, bool)`

GetCreditCardOk returns a tuple with the CreditCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditCard

`func (o *GetChildAccounts200ResponseDataInner) SetCreditCard(v GetChildAccounts200ResponseDataInnerCreditCard)`

SetCreditCard sets CreditCard field to given value.

### HasCreditCard

`func (o *GetChildAccounts200ResponseDataInner) HasCreditCard() bool`

HasCreditCard returns a boolean if a field has been set.

### GetEmail

`func (o *GetChildAccounts200ResponseDataInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetChildAccounts200ResponseDataInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetChildAccounts200ResponseDataInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetChildAccounts200ResponseDataInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEuuid

`func (o *GetChildAccounts200ResponseDataInner) GetEuuid() string`

GetEuuid returns the Euuid field if non-nil, zero value otherwise.

### GetEuuidOk

`func (o *GetChildAccounts200ResponseDataInner) GetEuuidOk() (*string, bool)`

GetEuuidOk returns a tuple with the Euuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEuuid

`func (o *GetChildAccounts200ResponseDataInner) SetEuuid(v string)`

SetEuuid sets Euuid field to given value.

### HasEuuid

`func (o *GetChildAccounts200ResponseDataInner) HasEuuid() bool`

HasEuuid returns a boolean if a field has been set.

### GetFirstName

`func (o *GetChildAccounts200ResponseDataInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *GetChildAccounts200ResponseDataInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *GetChildAccounts200ResponseDataInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *GetChildAccounts200ResponseDataInner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *GetChildAccounts200ResponseDataInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *GetChildAccounts200ResponseDataInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *GetChildAccounts200ResponseDataInner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *GetChildAccounts200ResponseDataInner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPhone

`func (o *GetChildAccounts200ResponseDataInner) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetChildAccounts200ResponseDataInner) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetChildAccounts200ResponseDataInner) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetChildAccounts200ResponseDataInner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetState

`func (o *GetChildAccounts200ResponseDataInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetChildAccounts200ResponseDataInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetChildAccounts200ResponseDataInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetChildAccounts200ResponseDataInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTaxId

`func (o *GetChildAccounts200ResponseDataInner) GetTaxId() string`

GetTaxId returns the TaxId field if non-nil, zero value otherwise.

### GetTaxIdOk

`func (o *GetChildAccounts200ResponseDataInner) GetTaxIdOk() (*string, bool)`

GetTaxIdOk returns a tuple with the TaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxId

`func (o *GetChildAccounts200ResponseDataInner) SetTaxId(v string)`

SetTaxId sets TaxId field to given value.

### HasTaxId

`func (o *GetChildAccounts200ResponseDataInner) HasTaxId() bool`

HasTaxId returns a boolean if a field has been set.

### GetZip

`func (o *GetChildAccounts200ResponseDataInner) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *GetChildAccounts200ResponseDataInner) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *GetChildAccounts200ResponseDataInner) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *GetChildAccounts200ResponseDataInner) HasZip() bool`

HasZip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


