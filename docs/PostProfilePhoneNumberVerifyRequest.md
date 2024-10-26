# PostProfilePhoneNumberVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OtpCode** | **string** | The one-time code received via SMS message after running the [Send a phone number verification code](https://techdocs.akamai.com/linode-api/reference/post-profile-phone-number) operation. | 

## Methods

### NewPostProfilePhoneNumberVerifyRequest

`func NewPostProfilePhoneNumberVerifyRequest(otpCode string, ) *PostProfilePhoneNumberVerifyRequest`

NewPostProfilePhoneNumberVerifyRequest instantiates a new PostProfilePhoneNumberVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostProfilePhoneNumberVerifyRequestWithDefaults

`func NewPostProfilePhoneNumberVerifyRequestWithDefaults() *PostProfilePhoneNumberVerifyRequest`

NewPostProfilePhoneNumberVerifyRequestWithDefaults instantiates a new PostProfilePhoneNumberVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOtpCode

`func (o *PostProfilePhoneNumberVerifyRequest) GetOtpCode() string`

GetOtpCode returns the OtpCode field if non-nil, zero value otherwise.

### GetOtpCodeOk

`func (o *PostProfilePhoneNumberVerifyRequest) GetOtpCodeOk() (*string, bool)`

GetOtpCodeOk returns a tuple with the OtpCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpCode

`func (o *PostProfilePhoneNumberVerifyRequest) SetOtpCode(v string)`

SetOtpCode sets OtpCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


