# PostProfilePhoneNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsoCode** | **string** | The two-letter ISO 3166 country code associated with the phone number. | 
**PhoneNumber** | **string** | A valid phone number. | 

## Methods

### NewPostProfilePhoneNumberRequest

`func NewPostProfilePhoneNumberRequest(isoCode string, phoneNumber string, ) *PostProfilePhoneNumberRequest`

NewPostProfilePhoneNumberRequest instantiates a new PostProfilePhoneNumberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostProfilePhoneNumberRequestWithDefaults

`func NewPostProfilePhoneNumberRequestWithDefaults() *PostProfilePhoneNumberRequest`

NewPostProfilePhoneNumberRequestWithDefaults instantiates a new PostProfilePhoneNumberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsoCode

`func (o *PostProfilePhoneNumberRequest) GetIsoCode() string`

GetIsoCode returns the IsoCode field if non-nil, zero value otherwise.

### GetIsoCodeOk

`func (o *PostProfilePhoneNumberRequest) GetIsoCodeOk() (*string, bool)`

GetIsoCodeOk returns a tuple with the IsoCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCode

`func (o *PostProfilePhoneNumberRequest) SetIsoCode(v string)`

SetIsoCode sets IsoCode field to given value.


### GetPhoneNumber

`func (o *PostProfilePhoneNumberRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PostProfilePhoneNumberRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PostProfilePhoneNumberRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


