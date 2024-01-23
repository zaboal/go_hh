# ResumeShouldSendSmsPhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | Префикс города номера телефона | 
**Country** | **string** | Префикс страны номера телефона | 
**Formatted** | **string** | Номер телефона отформатированный | 
**NeedVerification** | **bool** | Нужно ли верифицировать номер телефона | 
**Number** | **string** | Номер телефона без префиксов | 
**RestrictedCountry** | **bool** | Находится ли номер телефона в черном списке по отправке смс | 
**Verified** | **bool** | Верифицирован ли номер телефона пользователем | 

## Methods

### NewResumeShouldSendSmsPhone

`func NewResumeShouldSendSmsPhone(city string, country string, formatted string, needVerification bool, number string, restrictedCountry bool, verified bool, ) *ResumeShouldSendSmsPhone`

NewResumeShouldSendSmsPhone instantiates a new ResumeShouldSendSmsPhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeShouldSendSmsPhoneWithDefaults

`func NewResumeShouldSendSmsPhoneWithDefaults() *ResumeShouldSendSmsPhone`

NewResumeShouldSendSmsPhoneWithDefaults instantiates a new ResumeShouldSendSmsPhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *ResumeShouldSendSmsPhone) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ResumeShouldSendSmsPhone) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ResumeShouldSendSmsPhone) SetCity(v string)`

SetCity sets City field to given value.


### GetCountry

`func (o *ResumeShouldSendSmsPhone) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ResumeShouldSendSmsPhone) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ResumeShouldSendSmsPhone) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetFormatted

`func (o *ResumeShouldSendSmsPhone) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *ResumeShouldSendSmsPhone) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *ResumeShouldSendSmsPhone) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.


### GetNeedVerification

`func (o *ResumeShouldSendSmsPhone) GetNeedVerification() bool`

GetNeedVerification returns the NeedVerification field if non-nil, zero value otherwise.

### GetNeedVerificationOk

`func (o *ResumeShouldSendSmsPhone) GetNeedVerificationOk() (*bool, bool)`

GetNeedVerificationOk returns a tuple with the NeedVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedVerification

`func (o *ResumeShouldSendSmsPhone) SetNeedVerification(v bool)`

SetNeedVerification sets NeedVerification field to given value.


### GetNumber

`func (o *ResumeShouldSendSmsPhone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *ResumeShouldSendSmsPhone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *ResumeShouldSendSmsPhone) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetRestrictedCountry

`func (o *ResumeShouldSendSmsPhone) GetRestrictedCountry() bool`

GetRestrictedCountry returns the RestrictedCountry field if non-nil, zero value otherwise.

### GetRestrictedCountryOk

`func (o *ResumeShouldSendSmsPhone) GetRestrictedCountryOk() (*bool, bool)`

GetRestrictedCountryOk returns a tuple with the RestrictedCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCountry

`func (o *ResumeShouldSendSmsPhone) SetRestrictedCountry(v bool)`

SetRestrictedCountry sets RestrictedCountry field to given value.


### GetVerified

`func (o *ResumeShouldSendSmsPhone) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ResumeShouldSendSmsPhone) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ResumeShouldSendSmsPhone) SetVerified(v bool)`

SetVerified sets Verified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


