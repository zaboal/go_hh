# ManagerSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultCurrency** | [**ManagerSettingsCurrency**](ManagerSettingsCurrency.md) |  | 
**DefaultVacancyBrandedTemplate** | Pointer to [**NullableVacancyBrandedTemplate**](VacancyBrandedTemplate.md) |  | [optional] 
**UseSmsNotification** | **bool** | Предпочтение по использованию флага &#x60;send_sms&#x60; при [приглашении соискателя](#tag/Otklikipriglasheniya-rabotodatelya/operation/invite-applicant-to-vacancy)  | 

## Methods

### NewManagerSettings

`func NewManagerSettings(defaultCurrency ManagerSettingsCurrency, useSmsNotification bool, ) *ManagerSettings`

NewManagerSettings instantiates a new ManagerSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagerSettingsWithDefaults

`func NewManagerSettingsWithDefaults() *ManagerSettings`

NewManagerSettingsWithDefaults instantiates a new ManagerSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultCurrency

`func (o *ManagerSettings) GetDefaultCurrency() ManagerSettingsCurrency`

GetDefaultCurrency returns the DefaultCurrency field if non-nil, zero value otherwise.

### GetDefaultCurrencyOk

`func (o *ManagerSettings) GetDefaultCurrencyOk() (*ManagerSettingsCurrency, bool)`

GetDefaultCurrencyOk returns a tuple with the DefaultCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCurrency

`func (o *ManagerSettings) SetDefaultCurrency(v ManagerSettingsCurrency)`

SetDefaultCurrency sets DefaultCurrency field to given value.


### GetDefaultVacancyBrandedTemplate

`func (o *ManagerSettings) GetDefaultVacancyBrandedTemplate() VacancyBrandedTemplate`

GetDefaultVacancyBrandedTemplate returns the DefaultVacancyBrandedTemplate field if non-nil, zero value otherwise.

### GetDefaultVacancyBrandedTemplateOk

`func (o *ManagerSettings) GetDefaultVacancyBrandedTemplateOk() (*VacancyBrandedTemplate, bool)`

GetDefaultVacancyBrandedTemplateOk returns a tuple with the DefaultVacancyBrandedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVacancyBrandedTemplate

`func (o *ManagerSettings) SetDefaultVacancyBrandedTemplate(v VacancyBrandedTemplate)`

SetDefaultVacancyBrandedTemplate sets DefaultVacancyBrandedTemplate field to given value.

### HasDefaultVacancyBrandedTemplate

`func (o *ManagerSettings) HasDefaultVacancyBrandedTemplate() bool`

HasDefaultVacancyBrandedTemplate returns a boolean if a field has been set.

### SetDefaultVacancyBrandedTemplateNil

`func (o *ManagerSettings) SetDefaultVacancyBrandedTemplateNil(b bool)`

 SetDefaultVacancyBrandedTemplateNil sets the value for DefaultVacancyBrandedTemplate to be an explicit nil

### UnsetDefaultVacancyBrandedTemplate
`func (o *ManagerSettings) UnsetDefaultVacancyBrandedTemplate()`

UnsetDefaultVacancyBrandedTemplate ensures that no value is present for DefaultVacancyBrandedTemplate, not even an explicit nil
### GetUseSmsNotification

`func (o *ManagerSettings) GetUseSmsNotification() bool`

GetUseSmsNotification returns the UseSmsNotification field if non-nil, zero value otherwise.

### GetUseSmsNotificationOk

`func (o *ManagerSettings) GetUseSmsNotificationOk() (*bool, bool)`

GetUseSmsNotificationOk returns a tuple with the UseSmsNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSmsNotification

`func (o *ManagerSettings) SetUseSmsNotification(v bool)`

SetUseSmsNotification sets UseSmsNotification field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


