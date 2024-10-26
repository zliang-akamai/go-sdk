/*
Linode API

[Read the API documentation](https://techdocs.akamai.com/linode-api/reference/api).

API version: 4.189.2
Contact: support@linode.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the GetAccount200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAccount200Response{}

// GetAccount200Response Account object.
type GetAccount200Response struct {
	ActivePromotions []GetAccount200ResponseActivePromotionsInner `json:"active_promotions,omitempty"`
	// The date and time the account was activated.
	ActiveSince *time.Time `json:"active_since,omitempty"`
	// First line of this Account's billing address.
	Address1 *string `json:"address_1,omitempty"`
	// Second line of this Account's billing address.
	Address2 *string `json:"address_2,omitempty"`
	// This Account's balance, in US dollars.
	Balance *float32 `json:"balance,omitempty"`
	// This Account's current estimated invoice in US dollars. This is not your final invoice balance. Transfer charges are not included in the estimate.
	BalanceUninvoiced *float32 `json:"balance_uninvoiced,omitempty"`
	// The source of service charges for this Account, as determined by its relationship with Akamai. Accounts that are associated with Akamai-specific customers return a value of `akamai`. All other Accounts return a value of `linode`.
	BillingSource *string `json:"billing_source,omitempty"`
	// A list of capabilities your account supports.
	Capabilities []string `json:"capabilities,omitempty"`
	// The city for this Account's billing address.
	City *string `json:"city,omitempty"`
	// The company name associated with this Account.  Must not include any of the following characters: `<` `>` `(` `)` `\"` `=`
	Company *string `json:"company,omitempty"`
	// The two-letter ISO 3166 country code of this Account's billing address.
	Country *string `json:"country,omitempty"`
	CreditCard *GetAccount200ResponseCreditCard `json:"credit_card,omitempty"`
	// The email address of the person associated with this Account.
	Email *string `json:"email,omitempty"`
	// An external unique identifier for this account.
	Euuid *string `json:"euuid,omitempty"`
	// The first name of the person associated with this Account.  Must not include any of the following characters: `<` `>` `(` `)` `\"` `=`
	FirstName *string `json:"first_name,omitempty"`
	// The last name of the person associated with this Account.  Must not include any of the following characters: `<` `>` `(` `)` `\"` `=`
	LastName *string `json:"last_name,omitempty"`
	// The phone number associated with this Account.
	Phone *string `json:"phone,omitempty"`
	// If billing address is in the United States (US) or Canada (CA), only the two-letter ISO 3166 State or Province code are accepted. If entering a US military address, state abbreviations (AA, AE, AP) should be entered. If the address is outside the US or CA, this is the Province associated with the Account's billing address.
	State *string `json:"state,omitempty"`
	// The tax identification number associated with this Account, for tax calculations in some countries. If you do not live in a country that collects tax, this should be an empty string (`\"\"`).
	TaxId *string `json:"tax_id,omitempty"`
	// The zip code of this Account's billing address. The following restrictions apply:  - Can only contain ASCII letters, numbers, and hyphens (`-`). - Must not contain more than 9 letter or number characters.
	Zip *string `json:"zip,omitempty"`
}

// NewGetAccount200Response instantiates a new GetAccount200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccount200Response() *GetAccount200Response {
	this := GetAccount200Response{}
	return &this
}

// NewGetAccount200ResponseWithDefaults instantiates a new GetAccount200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccount200ResponseWithDefaults() *GetAccount200Response {
	this := GetAccount200Response{}
	return &this
}

// GetActivePromotions returns the ActivePromotions field value if set, zero value otherwise.
func (o *GetAccount200Response) GetActivePromotions() []GetAccount200ResponseActivePromotionsInner {
	if o == nil || IsNil(o.ActivePromotions) {
		var ret []GetAccount200ResponseActivePromotionsInner
		return ret
	}
	return o.ActivePromotions
}

// GetActivePromotionsOk returns a tuple with the ActivePromotions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetActivePromotionsOk() ([]GetAccount200ResponseActivePromotionsInner, bool) {
	if o == nil || IsNil(o.ActivePromotions) {
		return nil, false
	}
	return o.ActivePromotions, true
}

// HasActivePromotions returns a boolean if a field has been set.
func (o *GetAccount200Response) HasActivePromotions() bool {
	if o != nil && !IsNil(o.ActivePromotions) {
		return true
	}

	return false
}

// SetActivePromotions gets a reference to the given []GetAccount200ResponseActivePromotionsInner and assigns it to the ActivePromotions field.
func (o *GetAccount200Response) SetActivePromotions(v []GetAccount200ResponseActivePromotionsInner) {
	o.ActivePromotions = v
}

// GetActiveSince returns the ActiveSince field value if set, zero value otherwise.
func (o *GetAccount200Response) GetActiveSince() time.Time {
	if o == nil || IsNil(o.ActiveSince) {
		var ret time.Time
		return ret
	}
	return *o.ActiveSince
}

// GetActiveSinceOk returns a tuple with the ActiveSince field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetActiveSinceOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ActiveSince) {
		return nil, false
	}
	return o.ActiveSince, true
}

// HasActiveSince returns a boolean if a field has been set.
func (o *GetAccount200Response) HasActiveSince() bool {
	if o != nil && !IsNil(o.ActiveSince) {
		return true
	}

	return false
}

// SetActiveSince gets a reference to the given time.Time and assigns it to the ActiveSince field.
func (o *GetAccount200Response) SetActiveSince(v time.Time) {
	o.ActiveSince = &v
}

// GetAddress1 returns the Address1 field value if set, zero value otherwise.
func (o *GetAccount200Response) GetAddress1() string {
	if o == nil || IsNil(o.Address1) {
		var ret string
		return ret
	}
	return *o.Address1
}

// GetAddress1Ok returns a tuple with the Address1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetAddress1Ok() (*string, bool) {
	if o == nil || IsNil(o.Address1) {
		return nil, false
	}
	return o.Address1, true
}

// HasAddress1 returns a boolean if a field has been set.
func (o *GetAccount200Response) HasAddress1() bool {
	if o != nil && !IsNil(o.Address1) {
		return true
	}

	return false
}

// SetAddress1 gets a reference to the given string and assigns it to the Address1 field.
func (o *GetAccount200Response) SetAddress1(v string) {
	o.Address1 = &v
}

// GetAddress2 returns the Address2 field value if set, zero value otherwise.
func (o *GetAccount200Response) GetAddress2() string {
	if o == nil || IsNil(o.Address2) {
		var ret string
		return ret
	}
	return *o.Address2
}

// GetAddress2Ok returns a tuple with the Address2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetAddress2Ok() (*string, bool) {
	if o == nil || IsNil(o.Address2) {
		return nil, false
	}
	return o.Address2, true
}

// HasAddress2 returns a boolean if a field has been set.
func (o *GetAccount200Response) HasAddress2() bool {
	if o != nil && !IsNil(o.Address2) {
		return true
	}

	return false
}

// SetAddress2 gets a reference to the given string and assigns it to the Address2 field.
func (o *GetAccount200Response) SetAddress2(v string) {
	o.Address2 = &v
}

// GetBalance returns the Balance field value if set, zero value otherwise.
func (o *GetAccount200Response) GetBalance() float32 {
	if o == nil || IsNil(o.Balance) {
		var ret float32
		return ret
	}
	return *o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetBalanceOk() (*float32, bool) {
	if o == nil || IsNil(o.Balance) {
		return nil, false
	}
	return o.Balance, true
}

// HasBalance returns a boolean if a field has been set.
func (o *GetAccount200Response) HasBalance() bool {
	if o != nil && !IsNil(o.Balance) {
		return true
	}

	return false
}

// SetBalance gets a reference to the given float32 and assigns it to the Balance field.
func (o *GetAccount200Response) SetBalance(v float32) {
	o.Balance = &v
}

// GetBalanceUninvoiced returns the BalanceUninvoiced field value if set, zero value otherwise.
func (o *GetAccount200Response) GetBalanceUninvoiced() float32 {
	if o == nil || IsNil(o.BalanceUninvoiced) {
		var ret float32
		return ret
	}
	return *o.BalanceUninvoiced
}

// GetBalanceUninvoicedOk returns a tuple with the BalanceUninvoiced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetBalanceUninvoicedOk() (*float32, bool) {
	if o == nil || IsNil(o.BalanceUninvoiced) {
		return nil, false
	}
	return o.BalanceUninvoiced, true
}

// HasBalanceUninvoiced returns a boolean if a field has been set.
func (o *GetAccount200Response) HasBalanceUninvoiced() bool {
	if o != nil && !IsNil(o.BalanceUninvoiced) {
		return true
	}

	return false
}

// SetBalanceUninvoiced gets a reference to the given float32 and assigns it to the BalanceUninvoiced field.
func (o *GetAccount200Response) SetBalanceUninvoiced(v float32) {
	o.BalanceUninvoiced = &v
}

// GetBillingSource returns the BillingSource field value if set, zero value otherwise.
func (o *GetAccount200Response) GetBillingSource() string {
	if o == nil || IsNil(o.BillingSource) {
		var ret string
		return ret
	}
	return *o.BillingSource
}

// GetBillingSourceOk returns a tuple with the BillingSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetBillingSourceOk() (*string, bool) {
	if o == nil || IsNil(o.BillingSource) {
		return nil, false
	}
	return o.BillingSource, true
}

// HasBillingSource returns a boolean if a field has been set.
func (o *GetAccount200Response) HasBillingSource() bool {
	if o != nil && !IsNil(o.BillingSource) {
		return true
	}

	return false
}

// SetBillingSource gets a reference to the given string and assigns it to the BillingSource field.
func (o *GetAccount200Response) SetBillingSource(v string) {
	o.BillingSource = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *GetAccount200Response) GetCapabilities() []string {
	if o == nil || IsNil(o.Capabilities) {
		var ret []string
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetCapabilitiesOk() ([]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *GetAccount200Response) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []string and assigns it to the Capabilities field.
func (o *GetAccount200Response) SetCapabilities(v []string) {
	o.Capabilities = v
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *GetAccount200Response) GetCity() string {
	if o == nil || IsNil(o.City) {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetCityOk() (*string, bool) {
	if o == nil || IsNil(o.City) {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *GetAccount200Response) HasCity() bool {
	if o != nil && !IsNil(o.City) {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *GetAccount200Response) SetCity(v string) {
	o.City = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *GetAccount200Response) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *GetAccount200Response) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *GetAccount200Response) SetCompany(v string) {
	o.Company = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *GetAccount200Response) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *GetAccount200Response) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *GetAccount200Response) SetCountry(v string) {
	o.Country = &v
}

// GetCreditCard returns the CreditCard field value if set, zero value otherwise.
func (o *GetAccount200Response) GetCreditCard() GetAccount200ResponseCreditCard {
	if o == nil || IsNil(o.CreditCard) {
		var ret GetAccount200ResponseCreditCard
		return ret
	}
	return *o.CreditCard
}

// GetCreditCardOk returns a tuple with the CreditCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetCreditCardOk() (*GetAccount200ResponseCreditCard, bool) {
	if o == nil || IsNil(o.CreditCard) {
		return nil, false
	}
	return o.CreditCard, true
}

// HasCreditCard returns a boolean if a field has been set.
func (o *GetAccount200Response) HasCreditCard() bool {
	if o != nil && !IsNil(o.CreditCard) {
		return true
	}

	return false
}

// SetCreditCard gets a reference to the given GetAccount200ResponseCreditCard and assigns it to the CreditCard field.
func (o *GetAccount200Response) SetCreditCard(v GetAccount200ResponseCreditCard) {
	o.CreditCard = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *GetAccount200Response) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *GetAccount200Response) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *GetAccount200Response) SetEmail(v string) {
	o.Email = &v
}

// GetEuuid returns the Euuid field value if set, zero value otherwise.
func (o *GetAccount200Response) GetEuuid() string {
	if o == nil || IsNil(o.Euuid) {
		var ret string
		return ret
	}
	return *o.Euuid
}

// GetEuuidOk returns a tuple with the Euuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetEuuidOk() (*string, bool) {
	if o == nil || IsNil(o.Euuid) {
		return nil, false
	}
	return o.Euuid, true
}

// HasEuuid returns a boolean if a field has been set.
func (o *GetAccount200Response) HasEuuid() bool {
	if o != nil && !IsNil(o.Euuid) {
		return true
	}

	return false
}

// SetEuuid gets a reference to the given string and assigns it to the Euuid field.
func (o *GetAccount200Response) SetEuuid(v string) {
	o.Euuid = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *GetAccount200Response) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *GetAccount200Response) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *GetAccount200Response) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *GetAccount200Response) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *GetAccount200Response) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *GetAccount200Response) SetLastName(v string) {
	o.LastName = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *GetAccount200Response) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *GetAccount200Response) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *GetAccount200Response) SetPhone(v string) {
	o.Phone = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *GetAccount200Response) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *GetAccount200Response) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *GetAccount200Response) SetState(v string) {
	o.State = &v
}

// GetTaxId returns the TaxId field value if set, zero value otherwise.
func (o *GetAccount200Response) GetTaxId() string {
	if o == nil || IsNil(o.TaxId) {
		var ret string
		return ret
	}
	return *o.TaxId
}

// GetTaxIdOk returns a tuple with the TaxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetTaxIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaxId) {
		return nil, false
	}
	return o.TaxId, true
}

// HasTaxId returns a boolean if a field has been set.
func (o *GetAccount200Response) HasTaxId() bool {
	if o != nil && !IsNil(o.TaxId) {
		return true
	}

	return false
}

// SetTaxId gets a reference to the given string and assigns it to the TaxId field.
func (o *GetAccount200Response) SetTaxId(v string) {
	o.TaxId = &v
}

// GetZip returns the Zip field value if set, zero value otherwise.
func (o *GetAccount200Response) GetZip() string {
	if o == nil || IsNil(o.Zip) {
		var ret string
		return ret
	}
	return *o.Zip
}

// GetZipOk returns a tuple with the Zip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount200Response) GetZipOk() (*string, bool) {
	if o == nil || IsNil(o.Zip) {
		return nil, false
	}
	return o.Zip, true
}

// HasZip returns a boolean if a field has been set.
func (o *GetAccount200Response) HasZip() bool {
	if o != nil && !IsNil(o.Zip) {
		return true
	}

	return false
}

// SetZip gets a reference to the given string and assigns it to the Zip field.
func (o *GetAccount200Response) SetZip(v string) {
	o.Zip = &v
}

func (o GetAccount200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAccount200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivePromotions) {
		toSerialize["active_promotions"] = o.ActivePromotions
	}
	if !IsNil(o.ActiveSince) {
		toSerialize["active_since"] = o.ActiveSince
	}
	if !IsNil(o.Address1) {
		toSerialize["address_1"] = o.Address1
	}
	if !IsNil(o.Address2) {
		toSerialize["address_2"] = o.Address2
	}
	if !IsNil(o.Balance) {
		toSerialize["balance"] = o.Balance
	}
	if !IsNil(o.BalanceUninvoiced) {
		toSerialize["balance_uninvoiced"] = o.BalanceUninvoiced
	}
	if !IsNil(o.BillingSource) {
		toSerialize["billing_source"] = o.BillingSource
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.City) {
		toSerialize["city"] = o.City
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Country) {
		toSerialize["country"] = o.Country
	}
	if !IsNil(o.CreditCard) {
		toSerialize["credit_card"] = o.CreditCard
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Euuid) {
		toSerialize["euuid"] = o.Euuid
	}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.TaxId) {
		toSerialize["tax_id"] = o.TaxId
	}
	if !IsNil(o.Zip) {
		toSerialize["zip"] = o.Zip
	}
	return toSerialize, nil
}

type NullableGetAccount200Response struct {
	value *GetAccount200Response
	isSet bool
}

func (v NullableGetAccount200Response) Get() *GetAccount200Response {
	return v.value
}

func (v *NullableGetAccount200Response) Set(val *GetAccount200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccount200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccount200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccount200Response(val *GetAccount200Response) *NullableGetAccount200Response {
	return &NullableGetAccount200Response{value: val, isSet: true}
}

func (v NullableGetAccount200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccount200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


