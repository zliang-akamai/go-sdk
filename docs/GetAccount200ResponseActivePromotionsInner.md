# GetAccount200ResponseActivePromotionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditMonthlyCap** | Pointer to **string** | The amount available to spend per month. | [optional] 
**CreditRemaining** | Pointer to **string** | The total amount of credit left for this promotion. | [optional] 
**Description** | Pointer to **string** | A detailed description of this promotion. | [optional] 
**ExpireDt** | Pointer to **string** | When this promotion&#39;s credits expire. | [optional] 
**ImageUrl** | Pointer to **string** | The location of an image for this promotion. | [optional] 
**ServiceType** | Pointer to **string** | The service to which this promotion applies. | [optional] 
**Summary** | Pointer to **string** | Short details of this promotion. | [optional] 
**ThisMonthCreditRemaining** | Pointer to **string** | The amount of credit left for this month for this promotion. | [optional] 

## Methods

### NewGetAccount200ResponseActivePromotionsInner

`func NewGetAccount200ResponseActivePromotionsInner() *GetAccount200ResponseActivePromotionsInner`

NewGetAccount200ResponseActivePromotionsInner instantiates a new GetAccount200ResponseActivePromotionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccount200ResponseActivePromotionsInnerWithDefaults

`func NewGetAccount200ResponseActivePromotionsInnerWithDefaults() *GetAccount200ResponseActivePromotionsInner`

NewGetAccount200ResponseActivePromotionsInnerWithDefaults instantiates a new GetAccount200ResponseActivePromotionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditMonthlyCap

`func (o *GetAccount200ResponseActivePromotionsInner) GetCreditMonthlyCap() string`

GetCreditMonthlyCap returns the CreditMonthlyCap field if non-nil, zero value otherwise.

### GetCreditMonthlyCapOk

`func (o *GetAccount200ResponseActivePromotionsInner) GetCreditMonthlyCapOk() (*string, bool)`

GetCreditMonthlyCapOk returns a tuple with the CreditMonthlyCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditMonthlyCap

`func (o *GetAccount200ResponseActivePromotionsInner) SetCreditMonthlyCap(v string)`

SetCreditMonthlyCap sets CreditMonthlyCap field to given value.

### HasCreditMonthlyCap

`func (o *GetAccount200ResponseActivePromotionsInner) HasCreditMonthlyCap() bool`

HasCreditMonthlyCap returns a boolean if a field has been set.

### GetCreditRemaining

`func (o *GetAccount200ResponseActivePromotionsInner) GetCreditRemaining() string`

GetCreditRemaining returns the CreditRemaining field if non-nil, zero value otherwise.

### GetCreditRemainingOk

`func (o *GetAccount200ResponseActivePromotionsInner) GetCreditRemainingOk() (*string, bool)`

GetCreditRemainingOk returns a tuple with the CreditRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditRemaining

`func (o *GetAccount200ResponseActivePromotionsInner) SetCreditRemaining(v string)`

SetCreditRemaining sets CreditRemaining field to given value.

### HasCreditRemaining

`func (o *GetAccount200ResponseActivePromotionsInner) HasCreditRemaining() bool`

HasCreditRemaining returns a boolean if a field has been set.

### GetDescription

`func (o *GetAccount200ResponseActivePromotionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAccount200ResponseActivePromotionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAccount200ResponseActivePromotionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAccount200ResponseActivePromotionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpireDt

`func (o *GetAccount200ResponseActivePromotionsInner) GetExpireDt() string`

GetExpireDt returns the ExpireDt field if non-nil, zero value otherwise.

### GetExpireDtOk

`func (o *GetAccount200ResponseActivePromotionsInner) GetExpireDtOk() (*string, bool)`

GetExpireDtOk returns a tuple with the ExpireDt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDt

`func (o *GetAccount200ResponseActivePromotionsInner) SetExpireDt(v string)`

SetExpireDt sets ExpireDt field to given value.

### HasExpireDt

`func (o *GetAccount200ResponseActivePromotionsInner) HasExpireDt() bool`

HasExpireDt returns a boolean if a field has been set.

### GetImageUrl

`func (o *GetAccount200ResponseActivePromotionsInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GetAccount200ResponseActivePromotionsInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GetAccount200ResponseActivePromotionsInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GetAccount200ResponseActivePromotionsInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetServiceType

`func (o *GetAccount200ResponseActivePromotionsInner) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *GetAccount200ResponseActivePromotionsInner) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *GetAccount200ResponseActivePromotionsInner) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *GetAccount200ResponseActivePromotionsInner) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSummary

`func (o *GetAccount200ResponseActivePromotionsInner) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetAccount200ResponseActivePromotionsInner) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetAccount200ResponseActivePromotionsInner) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetAccount200ResponseActivePromotionsInner) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetThisMonthCreditRemaining

`func (o *GetAccount200ResponseActivePromotionsInner) GetThisMonthCreditRemaining() string`

GetThisMonthCreditRemaining returns the ThisMonthCreditRemaining field if non-nil, zero value otherwise.

### GetThisMonthCreditRemainingOk

`func (o *GetAccount200ResponseActivePromotionsInner) GetThisMonthCreditRemainingOk() (*string, bool)`

GetThisMonthCreditRemainingOk returns a tuple with the ThisMonthCreditRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThisMonthCreditRemaining

`func (o *GetAccount200ResponseActivePromotionsInner) SetThisMonthCreditRemaining(v string)`

SetThisMonthCreditRemaining sets ThisMonthCreditRemaining field to given value.

### HasThisMonthCreditRemaining

`func (o *GetAccount200ResponseActivePromotionsInner) HasThisMonthCreditRemaining() bool`

HasThisMonthCreditRemaining returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


