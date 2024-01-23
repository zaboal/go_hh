# VacanciesVacancyForManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptHandicapped** | **bool** | Указание, что вакансия доступна для соискателей с инвалидностью | 
**AcceptIncompleteResumes** | **bool** | Разрешен ли отклик на вакансию неполным резюме | 
**AcceptKids** | **bool** | Указание, что вакансия доступна для соискателей старше 14 лет [подробнее](https://github.com/hhru/api/blob/master/docs/employer_vacancies_accept_kids.md#accept-kids) | 
**AcceptTemporary** | Pointer to **NullableBool** | Указание, что вакансия доступна с временным трудоустройством | [optional] 
**AllowMessages** | **bool** | Возможность [переписки с кандидатами](https://inboxemp.hh.ru/) по данной вакансии | 
**AlternateUrl** | **string** | Ссылка на представление вакансии на сайте | 
**ApplyAlternateUrl** | **string** | Ссылка на отклик на вакансию на сайте | 
**Archived** | **bool** | Находится ли данная вакансия в архиве | 
**Area** | [**IncludesArea**](IncludesArea.md) |  | 
**BillingType** | [**NullableVacancyBillingTypeOutput**](VacancyBillingTypeOutput.md) |  | 
**BrandedDescription** | Pointer to **string** | Строка с кодом HTML (возможно наличие &#x60;&lt;script/&gt;&#x60; и &#x60;&lt;style/&gt;&#x60;), которая является альтернативой стандартному описанию вакансии.  HTML адаптирован для мобильных устройств и корректно отображается без поддержки JavaScript. При этом:  * Содержимое растягивается на 100% ширины контейнера и умещается в 300px без прокрутки. * Содержимое рассчитано на то, что оно будет вставлено в обвязку, в которую входит название, требуемый опыт работы, регион, тип занятости и рабочего дня вакансии, а также ссылка на компанию, опубликовавшую вакансию. * Изображения, которые могут встретиться в таком описании, адаптированы под Retina-дисплеи. * Размер шрифта не меньше 12px, размер межстрочного интервала не меньше 16px.  Значение может быть &#x60;null&#x60;, если у вакансии отсутствует индивидуальное описание | [optional] 
**Code** | Pointer to **NullableString** | Внутренний код вакансии | [optional] 
**Contacts** | Pointer to [**NullableVacancyContactsOutput**](VacancyContactsOutput.md) |  | [optional] 
**CreatedAt** | **string** | Дата и время публикации вакансии | 
**Department** | Pointer to [**NullableVacancyDepartmentOutput**](VacancyDepartmentOutput.md) |  | [optional] 
**Description** | **string** | Описание в html, не менее 200 символов | 
**DriverLicenseTypes** | [**[]VacancyDriverLicenceTypeItem**](VacancyDriverLicenceTypeItem.md) | Список требуемых категорий водительских прав | 
**Employer** | Pointer to [**VacanciesVacancyEmployer**](VacanciesVacancyEmployer.md) |  | [optional] 
**Employment** | Pointer to [**NullableVacancyEmploymentOutput**](VacancyEmploymentOutput.md) |  | [optional] 
**Experience** | [**NullableVacancyExperienceOutput**](VacancyExperienceOutput.md) |  | 
**HasTest** | **bool** | Информация о наличии прикрепленного тестового задании к вакансии | 
**Id** | **string** | Идентификатор вакансии | 
**InitialCreatedAt** | **string** | Дата и время создания вакансии | 
**InsiderInterview** | Pointer to [**VacancyInsiderInterview**](VacancyInsiderInterview.md) |  | [optional] 
**KeySkills** | [**[]VacancyKeySkillItem**](VacancyKeySkillItem.md) | Список ключевых навыков, не более 30 | 
**Languages** | Pointer to [**[]VacancyLanguageOutput**](VacancyLanguageOutput.md) | Список языков вакансии. Значения из справочника [/languages](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**Name** | **string** | Название | 
**NegotiationsUrl** | Pointer to **NullableString** | Ссылка для получения списка откликов/приглашений | [optional] 
**Premium** | **bool** | Является ли данная вакансия премиум-вакансией | 
**ProfessionalRoles** | [**[]VacancyProfessionalRoleItemOutput**](VacancyProfessionalRoleItemOutput.md) | Список профессиональных ролей | 
**PublishedAt** | **string** | Дата и время публикации вакансии | 
**Relations** | Pointer to [**[]VacancyRelationItem**](VacancyRelationItem.md) | Возвращает связи соискателя с вакансией. Значения из поля &#x60;vacancy_relation&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**ResponseLetterRequired** | **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | Pointer to [**NullableVacancySalary**](VacancySalary.md) |  | [optional] 
**Schedule** | [**VacancyScheduleOutput**](VacancyScheduleOutput.md) |  | 
**SuitableResumesUrl** | Pointer to **NullableString** | Подходящие резюме на вакансию | [optional] 
**Test** | Pointer to [**NullableVacancyDraftTest**](VacancyDraftTest.md) |  | [optional] 
**Type** | [**IncludesIdName**](IncludesIdName.md) |  | 
**VacancyConstructorTemplate** | Pointer to [**VacancyVacancyConstructorTemplate**](VacancyVacancyConstructorTemplate.md) |  | [optional] 
**VideoVacancy** | Pointer to [**VacancyVideoVacancy**](VacancyVideoVacancy.md) |  | [optional] 
**WorkingDays** | Pointer to [**[]VacancyWorkingDayItemOutput**](VacancyWorkingDayItemOutput.md) | Список рабочих дней | [optional] 
**WorkingTimeIntervals** | Pointer to [**[]VacancyWorkingTimeIntervalItemOutput**](VacancyWorkingTimeIntervalItemOutput.md) | Список с временными интервалами работы | [optional] 
**WorkingTimeModes** | Pointer to [**[]VacancyWorkingTimeModeItemOutput**](VacancyWorkingTimeModeItemOutput.md) | Список режимов времени работы | [optional] 
**Address** | [**VacanciesAddress**](VacanciesAddress.md) |  | 
**ArchivedAt** | Pointer to **string** | Дата архивации вакансии | [optional] 
**BrandedTemplate** | [**NullableVacancyBrandedTemplate**](VacancyBrandedTemplate.md) |  | 
**CanUpgradeBillingType** | **bool** | Можно ли улучшить биллинговый тип вакансии | 
**Counters** | Pointer to [**VacancyCountersOutput**](VacancyCountersOutput.md) |  | [optional] 
**ExpiresAt** | **string** | Дата и время окончания публикации вакансии | 
**Hidden** | **bool** | Удалена ли вакансия (скрыта из архива) | 
**Manager** | [**VacancyManager**](VacancyManager.md) |  | 
**ResponseNotifications** | **bool** | Уведомлять ли менеджера о новых откликах | 
**PreviousId** | Pointer to **NullableString** | Идентификатор архивной вакансии, на основе которой была опубликована текущая вакансия. Если вакансия была создана самостоятельно - null | [optional] 

## Methods

### NewVacanciesVacancyForManager

`func NewVacanciesVacancyForManager(acceptHandicapped bool, acceptIncompleteResumes bool, acceptKids bool, allowMessages bool, alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, billingType NullableVacancyBillingTypeOutput, createdAt string, description string, driverLicenseTypes []VacancyDriverLicenceTypeItem, experience NullableVacancyExperienceOutput, hasTest bool, id string, initialCreatedAt string, keySkills []VacancyKeySkillItem, name string, premium bool, professionalRoles []VacancyProfessionalRoleItemOutput, publishedAt string, responseLetterRequired bool, schedule VacancyScheduleOutput, type_ IncludesIdName, address VacanciesAddress, brandedTemplate NullableVacancyBrandedTemplate, canUpgradeBillingType bool, expiresAt string, hidden bool, manager VacancyManager, responseNotifications bool, ) *VacanciesVacancyForManager`

NewVacanciesVacancyForManager instantiates a new VacanciesVacancyForManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyForManagerWithDefaults

`func NewVacanciesVacancyForManagerWithDefaults() *VacanciesVacancyForManager`

NewVacanciesVacancyForManagerWithDefaults instantiates a new VacanciesVacancyForManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacanciesVacancyForManager) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacanciesVacancyForManager) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacanciesVacancyForManager) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.


### GetAcceptIncompleteResumes

`func (o *VacanciesVacancyForManager) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacanciesVacancyForManager) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacanciesVacancyForManager) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.


### GetAcceptKids

`func (o *VacanciesVacancyForManager) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacanciesVacancyForManager) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacanciesVacancyForManager) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.


### GetAcceptTemporary

`func (o *VacanciesVacancyForManager) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacanciesVacancyForManager) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacanciesVacancyForManager) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacanciesVacancyForManager) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacanciesVacancyForManager) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacanciesVacancyForManager) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAllowMessages

`func (o *VacanciesVacancyForManager) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacanciesVacancyForManager) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacanciesVacancyForManager) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.


### GetAlternateUrl

`func (o *VacanciesVacancyForManager) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVacancyForManager) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVacancyForManager) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesVacancyForManager) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesVacancyForManager) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesVacancyForManager) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetArchived

`func (o *VacanciesVacancyForManager) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesVacancyForManager) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesVacancyForManager) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArea

`func (o *VacanciesVacancyForManager) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesVacancyForManager) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesVacancyForManager) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetBillingType

`func (o *VacanciesVacancyForManager) GetBillingType() VacancyBillingTypeOutput`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacanciesVacancyForManager) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacanciesVacancyForManager) SetBillingType(v VacancyBillingTypeOutput)`

SetBillingType sets BillingType field to given value.


### SetBillingTypeNil

`func (o *VacanciesVacancyForManager) SetBillingTypeNil(b bool)`

 SetBillingTypeNil sets the value for BillingType to be an explicit nil

### UnsetBillingType
`func (o *VacanciesVacancyForManager) UnsetBillingType()`

UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
### GetBrandedDescription

`func (o *VacanciesVacancyForManager) GetBrandedDescription() string`

GetBrandedDescription returns the BrandedDescription field if non-nil, zero value otherwise.

### GetBrandedDescriptionOk

`func (o *VacanciesVacancyForManager) GetBrandedDescriptionOk() (*string, bool)`

GetBrandedDescriptionOk returns a tuple with the BrandedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedDescription

`func (o *VacanciesVacancyForManager) SetBrandedDescription(v string)`

SetBrandedDescription sets BrandedDescription field to given value.

### HasBrandedDescription

`func (o *VacanciesVacancyForManager) HasBrandedDescription() bool`

HasBrandedDescription returns a boolean if a field has been set.

### GetCode

`func (o *VacanciesVacancyForManager) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacanciesVacancyForManager) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacanciesVacancyForManager) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacanciesVacancyForManager) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacanciesVacancyForManager) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacanciesVacancyForManager) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetContacts

`func (o *VacanciesVacancyForManager) GetContacts() VacancyContactsOutput`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacanciesVacancyForManager) GetContactsOk() (*VacancyContactsOutput, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacanciesVacancyForManager) SetContacts(v VacancyContactsOutput)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacanciesVacancyForManager) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacanciesVacancyForManager) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacanciesVacancyForManager) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCreatedAt

`func (o *VacanciesVacancyForManager) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesVacancyForManager) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesVacancyForManager) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDepartment

`func (o *VacanciesVacancyForManager) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesVacancyForManager) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesVacancyForManager) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacanciesVacancyForManager) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacanciesVacancyForManager) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesVacancyForManager) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDescription

`func (o *VacanciesVacancyForManager) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacanciesVacancyForManager) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacanciesVacancyForManager) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDriverLicenseTypes

`func (o *VacanciesVacancyForManager) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacanciesVacancyForManager) GetDriverLicenseTypesOk() (*[]VacancyDriverLicenceTypeItem, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacanciesVacancyForManager) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEmployer

`func (o *VacanciesVacancyForManager) GetEmployer() VacanciesVacancyEmployer`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesVacancyForManager) GetEmployerOk() (*VacanciesVacancyEmployer, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesVacancyForManager) SetEmployer(v VacanciesVacancyEmployer)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *VacanciesVacancyForManager) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetEmployment

`func (o *VacanciesVacancyForManager) GetEmployment() VacancyEmploymentOutput`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacanciesVacancyForManager) GetEmploymentOk() (*VacancyEmploymentOutput, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacanciesVacancyForManager) SetEmployment(v VacancyEmploymentOutput)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacanciesVacancyForManager) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacanciesVacancyForManager) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacanciesVacancyForManager) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacanciesVacancyForManager) GetExperience() VacancyExperienceOutput`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacanciesVacancyForManager) GetExperienceOk() (*VacancyExperienceOutput, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacanciesVacancyForManager) SetExperience(v VacancyExperienceOutput)`

SetExperience sets Experience field to given value.


### SetExperienceNil

`func (o *VacanciesVacancyForManager) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacanciesVacancyForManager) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetHasTest

`func (o *VacanciesVacancyForManager) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesVacancyForManager) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesVacancyForManager) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesVacancyForManager) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVacancyForManager) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVacancyForManager) SetId(v string)`

SetId sets Id field to given value.


### GetInitialCreatedAt

`func (o *VacanciesVacancyForManager) GetInitialCreatedAt() string`

GetInitialCreatedAt returns the InitialCreatedAt field if non-nil, zero value otherwise.

### GetInitialCreatedAtOk

`func (o *VacanciesVacancyForManager) GetInitialCreatedAtOk() (*string, bool)`

GetInitialCreatedAtOk returns a tuple with the InitialCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCreatedAt

`func (o *VacanciesVacancyForManager) SetInitialCreatedAt(v string)`

SetInitialCreatedAt sets InitialCreatedAt field to given value.


### GetInsiderInterview

`func (o *VacanciesVacancyForManager) GetInsiderInterview() VacancyInsiderInterview`

GetInsiderInterview returns the InsiderInterview field if non-nil, zero value otherwise.

### GetInsiderInterviewOk

`func (o *VacanciesVacancyForManager) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool)`

GetInsiderInterviewOk returns a tuple with the InsiderInterview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderInterview

`func (o *VacanciesVacancyForManager) SetInsiderInterview(v VacancyInsiderInterview)`

SetInsiderInterview sets InsiderInterview field to given value.

### HasInsiderInterview

`func (o *VacanciesVacancyForManager) HasInsiderInterview() bool`

HasInsiderInterview returns a boolean if a field has been set.

### GetKeySkills

`func (o *VacanciesVacancyForManager) GetKeySkills() []VacancyKeySkillItem`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacanciesVacancyForManager) GetKeySkillsOk() (*[]VacancyKeySkillItem, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacanciesVacancyForManager) SetKeySkills(v []VacancyKeySkillItem)`

SetKeySkills sets KeySkills field to given value.


### GetLanguages

`func (o *VacanciesVacancyForManager) GetLanguages() []VacancyLanguageOutput`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacanciesVacancyForManager) GetLanguagesOk() (*[]VacancyLanguageOutput, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacanciesVacancyForManager) SetLanguages(v []VacancyLanguageOutput)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VacanciesVacancyForManager) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetName

`func (o *VacanciesVacancyForManager) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesVacancyForManager) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesVacancyForManager) SetName(v string)`

SetName sets Name field to given value.


### GetNegotiationsUrl

`func (o *VacanciesVacancyForManager) GetNegotiationsUrl() string`

GetNegotiationsUrl returns the NegotiationsUrl field if non-nil, zero value otherwise.

### GetNegotiationsUrlOk

`func (o *VacanciesVacancyForManager) GetNegotiationsUrlOk() (*string, bool)`

GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsUrl

`func (o *VacanciesVacancyForManager) SetNegotiationsUrl(v string)`

SetNegotiationsUrl sets NegotiationsUrl field to given value.

### HasNegotiationsUrl

`func (o *VacanciesVacancyForManager) HasNegotiationsUrl() bool`

HasNegotiationsUrl returns a boolean if a field has been set.

### SetNegotiationsUrlNil

`func (o *VacanciesVacancyForManager) SetNegotiationsUrlNil(b bool)`

 SetNegotiationsUrlNil sets the value for NegotiationsUrl to be an explicit nil

### UnsetNegotiationsUrl
`func (o *VacanciesVacancyForManager) UnsetNegotiationsUrl()`

UnsetNegotiationsUrl ensures that no value is present for NegotiationsUrl, not even an explicit nil
### GetPremium

`func (o *VacanciesVacancyForManager) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesVacancyForManager) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesVacancyForManager) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetProfessionalRoles

`func (o *VacanciesVacancyForManager) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacanciesVacancyForManager) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItemOutput, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacanciesVacancyForManager) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput)`

SetProfessionalRoles sets ProfessionalRoles field to given value.


### GetPublishedAt

`func (o *VacanciesVacancyForManager) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesVacancyForManager) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesVacancyForManager) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetRelations

`func (o *VacanciesVacancyForManager) GetRelations() []VacancyRelationItem`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *VacanciesVacancyForManager) GetRelationsOk() (*[]VacancyRelationItem, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *VacanciesVacancyForManager) SetRelations(v []VacancyRelationItem)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *VacanciesVacancyForManager) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetResponseLetterRequired

`func (o *VacanciesVacancyForManager) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesVacancyForManager) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesVacancyForManager) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesVacancyForManager) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesVacancyForManager) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesVacancyForManager) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesVacancyForManager) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesVacancyForManager) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesVacancyForManager) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesVacancyForManager) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesVacancyForManager) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesVacancyForManager) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacanciesVacancyForManager) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacanciesVacancyForManager) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesVacancyForManager) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacanciesVacancyForManager) GetSchedule() VacancyScheduleOutput`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacanciesVacancyForManager) GetScheduleOk() (*VacancyScheduleOutput, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacanciesVacancyForManager) SetSchedule(v VacancyScheduleOutput)`

SetSchedule sets Schedule field to given value.


### GetSuitableResumesUrl

`func (o *VacanciesVacancyForManager) GetSuitableResumesUrl() string`

GetSuitableResumesUrl returns the SuitableResumesUrl field if non-nil, zero value otherwise.

### GetSuitableResumesUrlOk

`func (o *VacanciesVacancyForManager) GetSuitableResumesUrlOk() (*string, bool)`

GetSuitableResumesUrlOk returns a tuple with the SuitableResumesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitableResumesUrl

`func (o *VacanciesVacancyForManager) SetSuitableResumesUrl(v string)`

SetSuitableResumesUrl sets SuitableResumesUrl field to given value.

### HasSuitableResumesUrl

`func (o *VacanciesVacancyForManager) HasSuitableResumesUrl() bool`

HasSuitableResumesUrl returns a boolean if a field has been set.

### SetSuitableResumesUrlNil

`func (o *VacanciesVacancyForManager) SetSuitableResumesUrlNil(b bool)`

 SetSuitableResumesUrlNil sets the value for SuitableResumesUrl to be an explicit nil

### UnsetSuitableResumesUrl
`func (o *VacanciesVacancyForManager) UnsetSuitableResumesUrl()`

UnsetSuitableResumesUrl ensures that no value is present for SuitableResumesUrl, not even an explicit nil
### GetTest

`func (o *VacanciesVacancyForManager) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacanciesVacancyForManager) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacanciesVacancyForManager) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacanciesVacancyForManager) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacanciesVacancyForManager) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacanciesVacancyForManager) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetType

`func (o *VacanciesVacancyForManager) GetType() IncludesIdName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesVacancyForManager) GetTypeOk() (*IncludesIdName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesVacancyForManager) SetType(v IncludesIdName)`

SetType sets Type field to given value.


### GetVacancyConstructorTemplate

`func (o *VacanciesVacancyForManager) GetVacancyConstructorTemplate() VacancyVacancyConstructorTemplate`

GetVacancyConstructorTemplate returns the VacancyConstructorTemplate field if non-nil, zero value otherwise.

### GetVacancyConstructorTemplateOk

`func (o *VacanciesVacancyForManager) GetVacancyConstructorTemplateOk() (*VacancyVacancyConstructorTemplate, bool)`

GetVacancyConstructorTemplateOk returns a tuple with the VacancyConstructorTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyConstructorTemplate

`func (o *VacanciesVacancyForManager) SetVacancyConstructorTemplate(v VacancyVacancyConstructorTemplate)`

SetVacancyConstructorTemplate sets VacancyConstructorTemplate field to given value.

### HasVacancyConstructorTemplate

`func (o *VacanciesVacancyForManager) HasVacancyConstructorTemplate() bool`

HasVacancyConstructorTemplate returns a boolean if a field has been set.

### GetVideoVacancy

`func (o *VacanciesVacancyForManager) GetVideoVacancy() VacancyVideoVacancy`

GetVideoVacancy returns the VideoVacancy field if non-nil, zero value otherwise.

### GetVideoVacancyOk

`func (o *VacanciesVacancyForManager) GetVideoVacancyOk() (*VacancyVideoVacancy, bool)`

GetVideoVacancyOk returns a tuple with the VideoVacancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoVacancy

`func (o *VacanciesVacancyForManager) SetVideoVacancy(v VacancyVideoVacancy)`

SetVideoVacancy sets VideoVacancy field to given value.

### HasVideoVacancy

`func (o *VacanciesVacancyForManager) HasVideoVacancy() bool`

HasVideoVacancy returns a boolean if a field has been set.

### GetWorkingDays

`func (o *VacanciesVacancyForManager) GetWorkingDays() []VacancyWorkingDayItemOutput`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacanciesVacancyForManager) GetWorkingDaysOk() (*[]VacancyWorkingDayItemOutput, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacanciesVacancyForManager) SetWorkingDays(v []VacancyWorkingDayItemOutput)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacanciesVacancyForManager) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacanciesVacancyForManager) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacanciesVacancyForManager) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacanciesVacancyForManager) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItemOutput`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacanciesVacancyForManager) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItemOutput, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacanciesVacancyForManager) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItemOutput)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacanciesVacancyForManager) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacanciesVacancyForManager) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacanciesVacancyForManager) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacanciesVacancyForManager) GetWorkingTimeModes() []VacancyWorkingTimeModeItemOutput`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacanciesVacancyForManager) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItemOutput, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacanciesVacancyForManager) SetWorkingTimeModes(v []VacancyWorkingTimeModeItemOutput)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacanciesVacancyForManager) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacanciesVacancyForManager) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacanciesVacancyForManager) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil
### GetAddress

`func (o *VacanciesVacancyForManager) GetAddress() VacanciesAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesVacancyForManager) GetAddressOk() (*VacanciesAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesVacancyForManager) SetAddress(v VacanciesAddress)`

SetAddress sets Address field to given value.


### GetArchivedAt

`func (o *VacanciesVacancyForManager) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *VacanciesVacancyForManager) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *VacanciesVacancyForManager) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *VacanciesVacancyForManager) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetBrandedTemplate

`func (o *VacanciesVacancyForManager) GetBrandedTemplate() VacancyBrandedTemplate`

GetBrandedTemplate returns the BrandedTemplate field if non-nil, zero value otherwise.

### GetBrandedTemplateOk

`func (o *VacanciesVacancyForManager) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool)`

GetBrandedTemplateOk returns a tuple with the BrandedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedTemplate

`func (o *VacanciesVacancyForManager) SetBrandedTemplate(v VacancyBrandedTemplate)`

SetBrandedTemplate sets BrandedTemplate field to given value.


### SetBrandedTemplateNil

`func (o *VacanciesVacancyForManager) SetBrandedTemplateNil(b bool)`

 SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil

### UnsetBrandedTemplate
`func (o *VacanciesVacancyForManager) UnsetBrandedTemplate()`

UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
### GetCanUpgradeBillingType

`func (o *VacanciesVacancyForManager) GetCanUpgradeBillingType() bool`

GetCanUpgradeBillingType returns the CanUpgradeBillingType field if non-nil, zero value otherwise.

### GetCanUpgradeBillingTypeOk

`func (o *VacanciesVacancyForManager) GetCanUpgradeBillingTypeOk() (*bool, bool)`

GetCanUpgradeBillingTypeOk returns a tuple with the CanUpgradeBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgradeBillingType

`func (o *VacanciesVacancyForManager) SetCanUpgradeBillingType(v bool)`

SetCanUpgradeBillingType sets CanUpgradeBillingType field to given value.


### GetCounters

`func (o *VacanciesVacancyForManager) GetCounters() VacancyCountersOutput`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *VacanciesVacancyForManager) GetCountersOk() (*VacancyCountersOutput, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *VacanciesVacancyForManager) SetCounters(v VacancyCountersOutput)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *VacanciesVacancyForManager) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetExpiresAt

`func (o *VacanciesVacancyForManager) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VacanciesVacancyForManager) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VacanciesVacancyForManager) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetHidden

`func (o *VacanciesVacancyForManager) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *VacanciesVacancyForManager) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *VacanciesVacancyForManager) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetManager

`func (o *VacanciesVacancyForManager) GetManager() VacancyManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacanciesVacancyForManager) GetManagerOk() (*VacancyManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacanciesVacancyForManager) SetManager(v VacancyManager)`

SetManager sets Manager field to given value.


### GetResponseNotifications

`func (o *VacanciesVacancyForManager) GetResponseNotifications() bool`

GetResponseNotifications returns the ResponseNotifications field if non-nil, zero value otherwise.

### GetResponseNotificationsOk

`func (o *VacanciesVacancyForManager) GetResponseNotificationsOk() (*bool, bool)`

GetResponseNotificationsOk returns a tuple with the ResponseNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseNotifications

`func (o *VacanciesVacancyForManager) SetResponseNotifications(v bool)`

SetResponseNotifications sets ResponseNotifications field to given value.


### GetPreviousId

`func (o *VacanciesVacancyForManager) GetPreviousId() string`

GetPreviousId returns the PreviousId field if non-nil, zero value otherwise.

### GetPreviousIdOk

`func (o *VacanciesVacancyForManager) GetPreviousIdOk() (*string, bool)`

GetPreviousIdOk returns a tuple with the PreviousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousId

`func (o *VacanciesVacancyForManager) SetPreviousId(v string)`

SetPreviousId sets PreviousId field to given value.

### HasPreviousId

`func (o *VacanciesVacancyForManager) HasPreviousId() bool`

HasPreviousId returns a boolean if a field has been set.

### SetPreviousIdNil

`func (o *VacanciesVacancyForManager) SetPreviousIdNil(b bool)`

 SetPreviousIdNil sets the value for PreviousId to be an explicit nil

### UnsetPreviousId
`func (o *VacanciesVacancyForManager) UnsetPreviousId()`

UnsetPreviousId ensures that no value is present for PreviousId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


