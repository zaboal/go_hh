# VacanciesVacancy

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
**Type** | [**IncludesIdName**](IncludesIdName.md) | Идентификатор типа вакансии из справочника [&#x60;vacancy_type&#x60;](https://api.hh.ru/openapi/redoc#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**VacancyConstructorTemplate** | Pointer to [**VacancyVacancyConstructorTemplate**](VacancyVacancyConstructorTemplate.md) |  | [optional] 
**VideoVacancy** | Pointer to [**VacancyVideoVacancy**](VacancyVideoVacancy.md) |  | [optional] 
**WorkingDays** | Pointer to [**[]VacancyWorkingDayItemOutput**](VacancyWorkingDayItemOutput.md) | Список рабочих дней | [optional] 
**WorkingTimeIntervals** | Pointer to [**[]VacancyWorkingTimeIntervalItemOutput**](VacancyWorkingTimeIntervalItemOutput.md) | Список с временными интервалами работы | [optional] 
**WorkingTimeModes** | Pointer to [**[]VacancyWorkingTimeModeItemOutput**](VacancyWorkingTimeModeItemOutput.md) | Список режимов времени работы | [optional] 
**Hidden** | **bool** | Удалена ли вакансия (скрыта из архива) | 
**Address** | [**VacanciesAddress**](VacanciesAddress.md) |  | 
**ArchivedAt** | Pointer to **string** | Дата архивации вакансии | [optional] 
**BrandedTemplate** | [**NullableVacancyBrandedTemplate**](VacancyBrandedTemplate.md) |  | 
**CanUpgradeBillingType** | **bool** | Можно ли улучшить биллинговый тип вакансии | 
**Counters** | Pointer to [**VacancyCountersOutput**](VacancyCountersOutput.md) |  | [optional] 
**ExpiresAt** | **string** | Дата и время окончания публикации вакансии | 
**Manager** | [**VacancyManager**](VacancyManager.md) |  | 
**ResponseNotifications** | **bool** | Уведомлять ли менеджера о новых откликах | 
**PreviousId** | Pointer to **NullableString** | Идентификатор архивной вакансии, на основе которой была опубликована текущая вакансия. Если вакансия была создана самостоятельно - null | [optional] 

## Methods

### NewVacanciesVacancy

`func NewVacanciesVacancy(acceptHandicapped bool, acceptIncompleteResumes bool, acceptKids bool, allowMessages bool, alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, billingType NullableVacancyBillingTypeOutput, createdAt string, description string, driverLicenseTypes []VacancyDriverLicenceTypeItem, experience NullableVacancyExperienceOutput, hasTest bool, id string, initialCreatedAt string, keySkills []VacancyKeySkillItem, name string, premium bool, professionalRoles []VacancyProfessionalRoleItemOutput, publishedAt string, responseLetterRequired bool, schedule VacancyScheduleOutput, type_ IncludesIdName, hidden bool, address VacanciesAddress, brandedTemplate NullableVacancyBrandedTemplate, canUpgradeBillingType bool, expiresAt string, manager VacancyManager, responseNotifications bool, ) *VacanciesVacancy`

NewVacanciesVacancy instantiates a new VacanciesVacancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyWithDefaults

`func NewVacanciesVacancyWithDefaults() *VacanciesVacancy`

NewVacanciesVacancyWithDefaults instantiates a new VacanciesVacancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacanciesVacancy) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacanciesVacancy) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacanciesVacancy) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.


### GetAcceptIncompleteResumes

`func (o *VacanciesVacancy) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacanciesVacancy) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacanciesVacancy) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.


### GetAcceptKids

`func (o *VacanciesVacancy) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacanciesVacancy) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacanciesVacancy) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.


### GetAcceptTemporary

`func (o *VacanciesVacancy) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacanciesVacancy) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacanciesVacancy) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacanciesVacancy) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacanciesVacancy) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacanciesVacancy) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAllowMessages

`func (o *VacanciesVacancy) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacanciesVacancy) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacanciesVacancy) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.


### GetAlternateUrl

`func (o *VacanciesVacancy) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVacancy) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVacancy) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesVacancy) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesVacancy) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesVacancy) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetArchived

`func (o *VacanciesVacancy) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesVacancy) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesVacancy) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArea

`func (o *VacanciesVacancy) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesVacancy) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesVacancy) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetBillingType

`func (o *VacanciesVacancy) GetBillingType() VacancyBillingTypeOutput`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacanciesVacancy) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacanciesVacancy) SetBillingType(v VacancyBillingTypeOutput)`

SetBillingType sets BillingType field to given value.


### SetBillingTypeNil

`func (o *VacanciesVacancy) SetBillingTypeNil(b bool)`

 SetBillingTypeNil sets the value for BillingType to be an explicit nil

### UnsetBillingType
`func (o *VacanciesVacancy) UnsetBillingType()`

UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
### GetBrandedDescription

`func (o *VacanciesVacancy) GetBrandedDescription() string`

GetBrandedDescription returns the BrandedDescription field if non-nil, zero value otherwise.

### GetBrandedDescriptionOk

`func (o *VacanciesVacancy) GetBrandedDescriptionOk() (*string, bool)`

GetBrandedDescriptionOk returns a tuple with the BrandedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedDescription

`func (o *VacanciesVacancy) SetBrandedDescription(v string)`

SetBrandedDescription sets BrandedDescription field to given value.

### HasBrandedDescription

`func (o *VacanciesVacancy) HasBrandedDescription() bool`

HasBrandedDescription returns a boolean if a field has been set.

### GetCode

`func (o *VacanciesVacancy) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacanciesVacancy) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacanciesVacancy) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacanciesVacancy) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacanciesVacancy) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacanciesVacancy) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetContacts

`func (o *VacanciesVacancy) GetContacts() VacancyContactsOutput`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacanciesVacancy) GetContactsOk() (*VacancyContactsOutput, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacanciesVacancy) SetContacts(v VacancyContactsOutput)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacanciesVacancy) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacanciesVacancy) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacanciesVacancy) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCreatedAt

`func (o *VacanciesVacancy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesVacancy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesVacancy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDepartment

`func (o *VacanciesVacancy) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesVacancy) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesVacancy) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacanciesVacancy) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacanciesVacancy) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesVacancy) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDescription

`func (o *VacanciesVacancy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacanciesVacancy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacanciesVacancy) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDriverLicenseTypes

`func (o *VacanciesVacancy) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacanciesVacancy) GetDriverLicenseTypesOk() (*[]VacancyDriverLicenceTypeItem, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacanciesVacancy) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEmployer

`func (o *VacanciesVacancy) GetEmployer() VacanciesVacancyEmployer`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesVacancy) GetEmployerOk() (*VacanciesVacancyEmployer, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesVacancy) SetEmployer(v VacanciesVacancyEmployer)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *VacanciesVacancy) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetEmployment

`func (o *VacanciesVacancy) GetEmployment() VacancyEmploymentOutput`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacanciesVacancy) GetEmploymentOk() (*VacancyEmploymentOutput, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacanciesVacancy) SetEmployment(v VacancyEmploymentOutput)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacanciesVacancy) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacanciesVacancy) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacanciesVacancy) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacanciesVacancy) GetExperience() VacancyExperienceOutput`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacanciesVacancy) GetExperienceOk() (*VacancyExperienceOutput, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacanciesVacancy) SetExperience(v VacancyExperienceOutput)`

SetExperience sets Experience field to given value.


### SetExperienceNil

`func (o *VacanciesVacancy) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacanciesVacancy) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetHasTest

`func (o *VacanciesVacancy) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesVacancy) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesVacancy) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesVacancy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVacancy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVacancy) SetId(v string)`

SetId sets Id field to given value.


### GetInitialCreatedAt

`func (o *VacanciesVacancy) GetInitialCreatedAt() string`

GetInitialCreatedAt returns the InitialCreatedAt field if non-nil, zero value otherwise.

### GetInitialCreatedAtOk

`func (o *VacanciesVacancy) GetInitialCreatedAtOk() (*string, bool)`

GetInitialCreatedAtOk returns a tuple with the InitialCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCreatedAt

`func (o *VacanciesVacancy) SetInitialCreatedAt(v string)`

SetInitialCreatedAt sets InitialCreatedAt field to given value.


### GetInsiderInterview

`func (o *VacanciesVacancy) GetInsiderInterview() VacancyInsiderInterview`

GetInsiderInterview returns the InsiderInterview field if non-nil, zero value otherwise.

### GetInsiderInterviewOk

`func (o *VacanciesVacancy) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool)`

GetInsiderInterviewOk returns a tuple with the InsiderInterview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderInterview

`func (o *VacanciesVacancy) SetInsiderInterview(v VacancyInsiderInterview)`

SetInsiderInterview sets InsiderInterview field to given value.

### HasInsiderInterview

`func (o *VacanciesVacancy) HasInsiderInterview() bool`

HasInsiderInterview returns a boolean if a field has been set.

### GetKeySkills

`func (o *VacanciesVacancy) GetKeySkills() []VacancyKeySkillItem`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacanciesVacancy) GetKeySkillsOk() (*[]VacancyKeySkillItem, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacanciesVacancy) SetKeySkills(v []VacancyKeySkillItem)`

SetKeySkills sets KeySkills field to given value.


### GetLanguages

`func (o *VacanciesVacancy) GetLanguages() []VacancyLanguageOutput`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacanciesVacancy) GetLanguagesOk() (*[]VacancyLanguageOutput, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacanciesVacancy) SetLanguages(v []VacancyLanguageOutput)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VacanciesVacancy) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetName

`func (o *VacanciesVacancy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesVacancy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesVacancy) SetName(v string)`

SetName sets Name field to given value.


### GetNegotiationsUrl

`func (o *VacanciesVacancy) GetNegotiationsUrl() string`

GetNegotiationsUrl returns the NegotiationsUrl field if non-nil, zero value otherwise.

### GetNegotiationsUrlOk

`func (o *VacanciesVacancy) GetNegotiationsUrlOk() (*string, bool)`

GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsUrl

`func (o *VacanciesVacancy) SetNegotiationsUrl(v string)`

SetNegotiationsUrl sets NegotiationsUrl field to given value.

### HasNegotiationsUrl

`func (o *VacanciesVacancy) HasNegotiationsUrl() bool`

HasNegotiationsUrl returns a boolean if a field has been set.

### SetNegotiationsUrlNil

`func (o *VacanciesVacancy) SetNegotiationsUrlNil(b bool)`

 SetNegotiationsUrlNil sets the value for NegotiationsUrl to be an explicit nil

### UnsetNegotiationsUrl
`func (o *VacanciesVacancy) UnsetNegotiationsUrl()`

UnsetNegotiationsUrl ensures that no value is present for NegotiationsUrl, not even an explicit nil
### GetPremium

`func (o *VacanciesVacancy) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesVacancy) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesVacancy) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetProfessionalRoles

`func (o *VacanciesVacancy) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacanciesVacancy) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItemOutput, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacanciesVacancy) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput)`

SetProfessionalRoles sets ProfessionalRoles field to given value.


### GetPublishedAt

`func (o *VacanciesVacancy) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesVacancy) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesVacancy) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetRelations

`func (o *VacanciesVacancy) GetRelations() []VacancyRelationItem`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *VacanciesVacancy) GetRelationsOk() (*[]VacancyRelationItem, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *VacanciesVacancy) SetRelations(v []VacancyRelationItem)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *VacanciesVacancy) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetResponseLetterRequired

`func (o *VacanciesVacancy) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesVacancy) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesVacancy) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesVacancy) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesVacancy) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesVacancy) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesVacancy) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesVacancy) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesVacancy) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesVacancy) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesVacancy) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesVacancy) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacanciesVacancy) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacanciesVacancy) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesVacancy) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacanciesVacancy) GetSchedule() VacancyScheduleOutput`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacanciesVacancy) GetScheduleOk() (*VacancyScheduleOutput, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacanciesVacancy) SetSchedule(v VacancyScheduleOutput)`

SetSchedule sets Schedule field to given value.


### GetSuitableResumesUrl

`func (o *VacanciesVacancy) GetSuitableResumesUrl() string`

GetSuitableResumesUrl returns the SuitableResumesUrl field if non-nil, zero value otherwise.

### GetSuitableResumesUrlOk

`func (o *VacanciesVacancy) GetSuitableResumesUrlOk() (*string, bool)`

GetSuitableResumesUrlOk returns a tuple with the SuitableResumesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitableResumesUrl

`func (o *VacanciesVacancy) SetSuitableResumesUrl(v string)`

SetSuitableResumesUrl sets SuitableResumesUrl field to given value.

### HasSuitableResumesUrl

`func (o *VacanciesVacancy) HasSuitableResumesUrl() bool`

HasSuitableResumesUrl returns a boolean if a field has been set.

### SetSuitableResumesUrlNil

`func (o *VacanciesVacancy) SetSuitableResumesUrlNil(b bool)`

 SetSuitableResumesUrlNil sets the value for SuitableResumesUrl to be an explicit nil

### UnsetSuitableResumesUrl
`func (o *VacanciesVacancy) UnsetSuitableResumesUrl()`

UnsetSuitableResumesUrl ensures that no value is present for SuitableResumesUrl, not even an explicit nil
### GetTest

`func (o *VacanciesVacancy) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacanciesVacancy) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacanciesVacancy) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacanciesVacancy) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacanciesVacancy) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacanciesVacancy) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetType

`func (o *VacanciesVacancy) GetType() IncludesIdName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesVacancy) GetTypeOk() (*IncludesIdName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesVacancy) SetType(v IncludesIdName)`

SetType sets Type field to given value.


### GetVacancyConstructorTemplate

`func (o *VacanciesVacancy) GetVacancyConstructorTemplate() VacancyVacancyConstructorTemplate`

GetVacancyConstructorTemplate returns the VacancyConstructorTemplate field if non-nil, zero value otherwise.

### GetVacancyConstructorTemplateOk

`func (o *VacanciesVacancy) GetVacancyConstructorTemplateOk() (*VacancyVacancyConstructorTemplate, bool)`

GetVacancyConstructorTemplateOk returns a tuple with the VacancyConstructorTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyConstructorTemplate

`func (o *VacanciesVacancy) SetVacancyConstructorTemplate(v VacancyVacancyConstructorTemplate)`

SetVacancyConstructorTemplate sets VacancyConstructorTemplate field to given value.

### HasVacancyConstructorTemplate

`func (o *VacanciesVacancy) HasVacancyConstructorTemplate() bool`

HasVacancyConstructorTemplate returns a boolean if a field has been set.

### GetVideoVacancy

`func (o *VacanciesVacancy) GetVideoVacancy() VacancyVideoVacancy`

GetVideoVacancy returns the VideoVacancy field if non-nil, zero value otherwise.

### GetVideoVacancyOk

`func (o *VacanciesVacancy) GetVideoVacancyOk() (*VacancyVideoVacancy, bool)`

GetVideoVacancyOk returns a tuple with the VideoVacancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoVacancy

`func (o *VacanciesVacancy) SetVideoVacancy(v VacancyVideoVacancy)`

SetVideoVacancy sets VideoVacancy field to given value.

### HasVideoVacancy

`func (o *VacanciesVacancy) HasVideoVacancy() bool`

HasVideoVacancy returns a boolean if a field has been set.

### GetWorkingDays

`func (o *VacanciesVacancy) GetWorkingDays() []VacancyWorkingDayItemOutput`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacanciesVacancy) GetWorkingDaysOk() (*[]VacancyWorkingDayItemOutput, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacanciesVacancy) SetWorkingDays(v []VacancyWorkingDayItemOutput)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacanciesVacancy) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacanciesVacancy) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacanciesVacancy) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacanciesVacancy) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItemOutput`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacanciesVacancy) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItemOutput, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacanciesVacancy) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItemOutput)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacanciesVacancy) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacanciesVacancy) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacanciesVacancy) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacanciesVacancy) GetWorkingTimeModes() []VacancyWorkingTimeModeItemOutput`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacanciesVacancy) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItemOutput, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacanciesVacancy) SetWorkingTimeModes(v []VacancyWorkingTimeModeItemOutput)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacanciesVacancy) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacanciesVacancy) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacanciesVacancy) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil
### GetHidden

`func (o *VacanciesVacancy) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *VacanciesVacancy) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *VacanciesVacancy) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetAddress

`func (o *VacanciesVacancy) GetAddress() VacanciesAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesVacancy) GetAddressOk() (*VacanciesAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesVacancy) SetAddress(v VacanciesAddress)`

SetAddress sets Address field to given value.


### GetArchivedAt

`func (o *VacanciesVacancy) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *VacanciesVacancy) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *VacanciesVacancy) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.

### HasArchivedAt

`func (o *VacanciesVacancy) HasArchivedAt() bool`

HasArchivedAt returns a boolean if a field has been set.

### GetBrandedTemplate

`func (o *VacanciesVacancy) GetBrandedTemplate() VacancyBrandedTemplate`

GetBrandedTemplate returns the BrandedTemplate field if non-nil, zero value otherwise.

### GetBrandedTemplateOk

`func (o *VacanciesVacancy) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool)`

GetBrandedTemplateOk returns a tuple with the BrandedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedTemplate

`func (o *VacanciesVacancy) SetBrandedTemplate(v VacancyBrandedTemplate)`

SetBrandedTemplate sets BrandedTemplate field to given value.


### SetBrandedTemplateNil

`func (o *VacanciesVacancy) SetBrandedTemplateNil(b bool)`

 SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil

### UnsetBrandedTemplate
`func (o *VacanciesVacancy) UnsetBrandedTemplate()`

UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
### GetCanUpgradeBillingType

`func (o *VacanciesVacancy) GetCanUpgradeBillingType() bool`

GetCanUpgradeBillingType returns the CanUpgradeBillingType field if non-nil, zero value otherwise.

### GetCanUpgradeBillingTypeOk

`func (o *VacanciesVacancy) GetCanUpgradeBillingTypeOk() (*bool, bool)`

GetCanUpgradeBillingTypeOk returns a tuple with the CanUpgradeBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgradeBillingType

`func (o *VacanciesVacancy) SetCanUpgradeBillingType(v bool)`

SetCanUpgradeBillingType sets CanUpgradeBillingType field to given value.


### GetCounters

`func (o *VacanciesVacancy) GetCounters() VacancyCountersOutput`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *VacanciesVacancy) GetCountersOk() (*VacancyCountersOutput, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *VacanciesVacancy) SetCounters(v VacancyCountersOutput)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *VacanciesVacancy) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetExpiresAt

`func (o *VacanciesVacancy) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VacanciesVacancy) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VacanciesVacancy) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetManager

`func (o *VacanciesVacancy) GetManager() VacancyManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacanciesVacancy) GetManagerOk() (*VacancyManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacanciesVacancy) SetManager(v VacancyManager)`

SetManager sets Manager field to given value.


### GetResponseNotifications

`func (o *VacanciesVacancy) GetResponseNotifications() bool`

GetResponseNotifications returns the ResponseNotifications field if non-nil, zero value otherwise.

### GetResponseNotificationsOk

`func (o *VacanciesVacancy) GetResponseNotificationsOk() (*bool, bool)`

GetResponseNotificationsOk returns a tuple with the ResponseNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseNotifications

`func (o *VacanciesVacancy) SetResponseNotifications(v bool)`

SetResponseNotifications sets ResponseNotifications field to given value.


### GetPreviousId

`func (o *VacanciesVacancy) GetPreviousId() string`

GetPreviousId returns the PreviousId field if non-nil, zero value otherwise.

### GetPreviousIdOk

`func (o *VacanciesVacancy) GetPreviousIdOk() (*string, bool)`

GetPreviousIdOk returns a tuple with the PreviousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousId

`func (o *VacanciesVacancy) SetPreviousId(v string)`

SetPreviousId sets PreviousId field to given value.

### HasPreviousId

`func (o *VacanciesVacancy) HasPreviousId() bool`

HasPreviousId returns a boolean if a field has been set.

### SetPreviousIdNil

`func (o *VacanciesVacancy) SetPreviousIdNil(b bool)`

 SetPreviousIdNil sets the value for PreviousId to be an explicit nil

### UnsetPreviousId
`func (o *VacanciesVacancy) UnsetPreviousId()`

UnsetPreviousId ensures that no value is present for PreviousId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


