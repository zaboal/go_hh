# VacanciesVacancyManagerFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**VacanciesAddress**](VacanciesAddress.md) |  | 
**ArchivedAt** | Pointer to **string** | Дата архивации вакансии | [optional] 
**BrandedTemplate** | [**NullableVacancyBrandedTemplate**](VacancyBrandedTemplate.md) |  | 
**CanUpgradeBillingType** | **bool** | Можно ли улучшить биллинговый тип вакансии | 
**Counters** | Pointer to [**VacancyCountersOutput**](VacancyCountersOutput.md) |  | [optional] 
**ExpiresAt** | **string** | Дата и время окончания публикации вакансии | 
**Hidden** | **bool** | Удалена ли вакансия (скрыта из архива) | 
**Manager** | [**VacancyManager**](VacancyManager.md) |  | 
**ResponseNotifications** | **bool** | Уведомлять ли менеджера о новых откликах | 

## Methods

### NewVacanciesVacancyManagerFields

`func NewVacanciesVacancyManagerFields(address VacanciesAddress, brandedTemplate NullableVacancyBrandedTemplate, canUpgradeBillingType bool, expiresAt string, hidden bool, manager VacancyManager, responseNotifications bool, ) *VacanciesVacancyManagerFields`

NewVacanciesVacancyManagerFields instantiates a new VacanciesVacancyManagerFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyManagerFieldsWithDefaults

`func NewVacanciesVacancyManagerFieldsWithDefaults() *VacanciesVacancyManagerFields`

NewVacanciesVacancyManagerFieldsWithDefaults instantiates a new VacanciesVacancyManagerFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VacanciesVacancyManagerFields) GetAddress() VacanciesAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesVacancyManagerFields) GetAddressOk() (*VacanciesAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesVacancyManagerFields) SetAddress(v VacanciesAddress)`

SetAddress sets Address field to given value.


### GetArchivedAt

`func (o *VacanciesVacancyManagerFields) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *VacanciesVacancyManagerFields) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *VacanciesVacancyManagerFields) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *VacanciesVacancyManagerFields) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetBrandedTemplate

`func (o *VacanciesVacancyManagerFields) GetBrandedTemplate() VacancyBrandedTemplate`

GetBrandedTemplate returns the BrandedTemplate field if non-nil, zero value otherwise.

### GetBrandedTemplateOk

`func (o *VacanciesVacancyManagerFields) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool)`

GetBrandedTemplateOk returns a tuple with the BrandedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedTemplate

`func (o *VacanciesVacancyManagerFields) SetBrandedTemplate(v VacancyBrandedTemplate)`

SetBrandedTemplate sets BrandedTemplate field to given value.


### SetBrandedTemplateNil

`func (o *VacanciesVacancyManagerFields) SetBrandedTemplateNil(b bool)`

 SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil

### UnsetBrandedTemplate
`func (o *VacanciesVacancyManagerFields) UnsetBrandedTemplate()`

UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
### GetCanUpgradeBillingType

`func (o *VacanciesVacancyManagerFields) GetCanUpgradeBillingType() bool`

GetCanUpgradeBillingType returns the CanUpgradeBillingType field if non-nil, zero value otherwise.

### GetCanUpgradeBillingTypeOk

`func (o *VacanciesVacancyManagerFields) GetCanUpgradeBillingTypeOk() (*bool, bool)`

GetCanUpgradeBillingTypeOk returns a tuple with the CanUpgradeBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgradeBillingType

`func (o *VacanciesVacancyManagerFields) SetCanUpgradeBillingType(v bool)`

SetCanUpgradeBillingType sets CanUpgradeBillingType field to given value.


### GetCounters

`func (o *VacanciesVacancyManagerFields) GetCounters() VacancyCountersOutput`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *VacanciesVacancyManagerFields) GetCountersOk() (*VacancyCountersOutput, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *VacanciesVacancyManagerFields) SetCounters(v VacancyCountersOutput)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *VacanciesVacancyManagerFields) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetExpiresAt

`func (o *VacanciesVacancyManagerFields) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VacanciesVacancyManagerFields) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VacanciesVacancyManagerFields) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetHidden

`func (o *VacanciesVacancyManagerFields) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *VacanciesVacancyManagerFields) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *VacanciesVacancyManagerFields) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetManager

`func (o *VacanciesVacancyManagerFields) GetManager() VacancyManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacanciesVacancyManagerFields) GetManagerOk() (*VacancyManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacanciesVacancyManagerFields) SetManager(v VacancyManager)`

SetManager sets Manager field to given value.


### GetResponseNotifications

`func (o *VacanciesVacancyManagerFields) GetResponseNotifications() bool`

GetResponseNotifications returns the ResponseNotifications field if non-nil, zero value otherwise.

### GetResponseNotificationsOk

`func (o *VacanciesVacancyManagerFields) GetResponseNotificationsOk() (*bool, bool)`

GetResponseNotificationsOk returns a tuple with the ResponseNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseNotifications

`func (o *VacanciesVacancyManagerFields) SetResponseNotifications(v bool)`

SetResponseNotifications sets ResponseNotifications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


