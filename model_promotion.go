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
)

// checks if the Promotion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Promotion{}

// Promotion Promotions generally offer a set amount of credit that can be used toward your Linode services, and the promotion expires after a specified date. As well, a monthly cap on the promotional offer is set.  Simply put, a promotion offers a certain amount of credit  month, until either the expiration date is passed, or until the total promotional credit is used, whichever comes first.
type Promotion struct {
	// The amount available to spend per month.
	CreditMonthlyCap *string `json:"credit_monthly_cap,omitempty"`
	// The total amount of credit left for this promotion.
	CreditRemaining *string `json:"credit_remaining,omitempty"`
	// A detailed description of this promotion.
	Description *string `json:"description,omitempty"`
	// When this promotion's credits expire.
	ExpireDt *string `json:"expire_dt,omitempty"`
	// The location of an image for this promotion.
	ImageUrl *string `json:"image_url,omitempty"`
	// The service to which this promotion applies.
	ServiceType *string `json:"service_type,omitempty"`
	// Short details of this promotion.
	Summary *string `json:"summary,omitempty"`
	// The amount of credit left for this month for this promotion.
	ThisMonthCreditRemaining *string `json:"this_month_credit_remaining,omitempty"`
}

// NewPromotion instantiates a new Promotion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotion() *Promotion {
	this := Promotion{}
	return &this
}

// NewPromotionWithDefaults instantiates a new Promotion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotionWithDefaults() *Promotion {
	this := Promotion{}
	return &this
}

// GetCreditMonthlyCap returns the CreditMonthlyCap field value if set, zero value otherwise.
func (o *Promotion) GetCreditMonthlyCap() string {
	if o == nil || IsNil(o.CreditMonthlyCap) {
		var ret string
		return ret
	}
	return *o.CreditMonthlyCap
}

// GetCreditMonthlyCapOk returns a tuple with the CreditMonthlyCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion) GetCreditMonthlyCapOk() (*string, bool) {
	if o == nil || IsNil(o.CreditMonthlyCap) {
		return nil, false
	}
	return o.CreditMonthlyCap, true
}

// HasCreditMonthlyCap returns a boolean if a field has been set.
func (o *Promotion) HasCreditMonthlyCap() bool {
	if o != nil && !IsNil(o.CreditMonthlyCap) {
		return true
	}

	return false
}

// SetCreditMonthlyCap gets a reference to the given string and assigns it to the CreditMonthlyCap field.
func (o *Promotion) SetCreditMonthlyCap(v string) {
	o.CreditMonthlyCap = &v
}

// GetCreditRemaining returns the CreditRemaining field value if set, zero value otherwise.
func (o *Promotion) GetCreditRemaining() string {
	if o == nil || IsNil(o.CreditRemaining) {
		var ret string
		return ret
	}
	return *o.CreditRemaining
}

// GetCreditRemainingOk returns a tuple with the CreditRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion) GetCreditRemainingOk() (*string, bool) {
	if o == nil || IsNil(o.CreditRemaining) {
		return nil, false
	}
	return o.CreditRemaining, true
}

// HasCreditRemaining returns a boolean if a field has been set.
func (o *Promotion) HasCreditRemaining() bool {
	if o != nil && !IsNil(o.CreditRemaining) {
		return true
	}

	return false
}

// SetCreditRemaining gets a reference to the given string and assigns it to the CreditRemaining field.
func (o *Promotion) SetCreditRemaining(v string) {
	o.CreditRemaining = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Promotion) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Promotion) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Promotion) SetDescription(v string) {
	o.Description = &v
}

// GetExpireDt returns the ExpireDt field value if set, zero value otherwise.
func (o *Promotion) GetExpireDt() string {
	if o == nil || IsNil(o.ExpireDt) {
		var ret string
		return ret
	}
	return *o.ExpireDt
}

// GetExpireDtOk returns a tuple with the ExpireDt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion) GetExpireDtOk() (*string, bool) {
	if o == nil || IsNil(o.ExpireDt) {
		return nil, false
	}
	return o.ExpireDt, true
}

// HasExpireDt returns a boolean if a field has been set.
func (o *Promotion) HasExpireDt() bool {
	if o != nil && !IsNil(o.ExpireDt) {
		return true
	}

	return false
}

// SetExpireDt gets a reference to the given string and assigns it to the ExpireDt field.
func (o *Promotion) SetExpireDt(v string) {
	o.ExpireDt = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *Promotion) GetImageUrl() string {
	if o == nil || IsNil(o.ImageUrl) {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion) GetImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ImageUrl) {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *Promotion) HasImageUrl() bool {
	if o != nil && !IsNil(o.ImageUrl) {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *Promotion) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetServiceType returns the ServiceType field value if set, zero value otherwise.
func (o *Promotion) GetServiceType() string {
	if o == nil || IsNil(o.ServiceType) {
		var ret string
		return ret
	}
	return *o.ServiceType
}

// GetServiceTypeOk returns a tuple with the ServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion) GetServiceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceType) {
		return nil, false
	}
	return o.ServiceType, true
}

// HasServiceType returns a boolean if a field has been set.
func (o *Promotion) HasServiceType() bool {
	if o != nil && !IsNil(o.ServiceType) {
		return true
	}

	return false
}

// SetServiceType gets a reference to the given string and assigns it to the ServiceType field.
func (o *Promotion) SetServiceType(v string) {
	o.ServiceType = &v
}

// GetSummary returns the Summary field value if set, zero value otherwise.
func (o *Promotion) GetSummary() string {
	if o == nil || IsNil(o.Summary) {
		var ret string
		return ret
	}
	return *o.Summary
}

// GetSummaryOk returns a tuple with the Summary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion) GetSummaryOk() (*string, bool) {
	if o == nil || IsNil(o.Summary) {
		return nil, false
	}
	return o.Summary, true
}

// HasSummary returns a boolean if a field has been set.
func (o *Promotion) HasSummary() bool {
	if o != nil && !IsNil(o.Summary) {
		return true
	}

	return false
}

// SetSummary gets a reference to the given string and assigns it to the Summary field.
func (o *Promotion) SetSummary(v string) {
	o.Summary = &v
}

// GetThisMonthCreditRemaining returns the ThisMonthCreditRemaining field value if set, zero value otherwise.
func (o *Promotion) GetThisMonthCreditRemaining() string {
	if o == nil || IsNil(o.ThisMonthCreditRemaining) {
		var ret string
		return ret
	}
	return *o.ThisMonthCreditRemaining
}

// GetThisMonthCreditRemainingOk returns a tuple with the ThisMonthCreditRemaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Promotion) GetThisMonthCreditRemainingOk() (*string, bool) {
	if o == nil || IsNil(o.ThisMonthCreditRemaining) {
		return nil, false
	}
	return o.ThisMonthCreditRemaining, true
}

// HasThisMonthCreditRemaining returns a boolean if a field has been set.
func (o *Promotion) HasThisMonthCreditRemaining() bool {
	if o != nil && !IsNil(o.ThisMonthCreditRemaining) {
		return true
	}

	return false
}

// SetThisMonthCreditRemaining gets a reference to the given string and assigns it to the ThisMonthCreditRemaining field.
func (o *Promotion) SetThisMonthCreditRemaining(v string) {
	o.ThisMonthCreditRemaining = &v
}

func (o Promotion) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Promotion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreditMonthlyCap) {
		toSerialize["credit_monthly_cap"] = o.CreditMonthlyCap
	}
	if !IsNil(o.CreditRemaining) {
		toSerialize["credit_remaining"] = o.CreditRemaining
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ExpireDt) {
		toSerialize["expire_dt"] = o.ExpireDt
	}
	if !IsNil(o.ImageUrl) {
		toSerialize["image_url"] = o.ImageUrl
	}
	if !IsNil(o.ServiceType) {
		toSerialize["service_type"] = o.ServiceType
	}
	if !IsNil(o.Summary) {
		toSerialize["summary"] = o.Summary
	}
	if !IsNil(o.ThisMonthCreditRemaining) {
		toSerialize["this_month_credit_remaining"] = o.ThisMonthCreditRemaining
	}
	return toSerialize, nil
}

type NullablePromotion struct {
	value *Promotion
	isSet bool
}

func (v NullablePromotion) Get() *Promotion {
	return v.value
}

func (v *NullablePromotion) Set(val *Promotion) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotion) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotion(val *Promotion) *NullablePromotion {
	return &NullablePromotion{value: val, isSet: true}
}

func (v NullablePromotion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


