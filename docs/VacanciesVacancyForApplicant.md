# VacanciesVacancyForApplicant

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
**Hidden** | Pointer to **bool** | Удалена ли вакансия (скрыта из архива). Оставлено для обеспечения обратной совместимости | [optional] 
**Address** | Pointer to [**NullableVacancyAddressOutput**](VacancyAddressOutput.md) |  | [optional] 

## Methods

### NewVacanciesVacancyForApplicant

`func NewVacanciesVacancyForApplicant(acceptHandicapped bool, acceptIncompleteResumes bool, acceptKids bool, allowMessages bool, alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, billingType NullableVacancyBillingTypeOutput, createdAt string, description string, driverLicenseTypes []VacancyDriverLicenceTypeItem, experience NullableVacancyExperienceOutput, hasTest bool, id string, initialCreatedAt string, keySkills []VacancyKeySkillItem, name string, premium bool, professionalRoles []VacancyProfessionalRoleItemOutput, publishedAt string, responseLetterRequired bool, schedule VacancyScheduleOutput, type_ IncludesIdName, ) *VacanciesVacancyForApplicant`

NewVacanciesVacancyForApplicant instantiates a new VacanciesVacancyForApplicant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyForApplicantWithDefaults

`func NewVacanciesVacancyForApplicantWithDefaults() *VacanciesVacancyForApplicant`

NewVacanciesVacancyForApplicantWithDefaults instantiates a new VacanciesVacancyForApplicant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacanciesVacancyForApplicant) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacanciesVacancyForApplicant) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacanciesVacancyForApplicant) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.


### GetAcceptIncompleteResumes

`func (o *VacanciesVacancyForApplicant) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacanciesVacancyForApplicant) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacanciesVacancyForApplicant) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.


### GetAcceptKids

`func (o *VacanciesVacancyForApplicant) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacanciesVacancyForApplicant) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacanciesVacancyForApplicant) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.


### GetAcceptTemporary

`func (o *VacanciesVacancyForApplicant) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacanciesVacancyForApplicant) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacanciesVacancyForApplicant) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacanciesVacancyForApplicant) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacanciesVacancyForApplicant) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacanciesVacancyForApplicant) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAllowMessages

`func (o *VacanciesVacancyForApplicant) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacanciesVacancyForApplicant) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacanciesVacancyForApplicant) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.


### GetAlternateUrl

`func (o *VacanciesVacancyForApplicant) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVacancyForApplicant) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVacancyForApplicant) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesVacancyForApplicant) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesVacancyForApplicant) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesVacancyForApplicant) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetArchived

`func (o *VacanciesVacancyForApplicant) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesVacancyForApplicant) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesVacancyForApplicant) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArea

`func (o *VacanciesVacancyForApplicant) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesVacancyForApplicant) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesVacancyForApplicant) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetBillingType

`func (o *VacanciesVacancyForApplicant) GetBillingType() VacancyBillingTypeOutput`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacanciesVacancyForApplicant) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacanciesVacancyForApplicant) SetBillingType(v VacancyBillingTypeOutput)`

SetBillingType sets BillingType field to given value.


### SetBillingTypeNil

`func (o *VacanciesVacancyForApplicant) SetBillingTypeNil(b bool)`

 SetBillingTypeNil sets the value for BillingType to be an explicit nil

### UnsetBillingType
`func (o *VacanciesVacancyForApplicant) UnsetBillingType()`

UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
### GetBrandedDescription

`func (o *VacanciesVacancyForApplicant) GetBrandedDescription() string`

GetBrandedDescription returns the BrandedDescription field if non-nil, zero value otherwise.

### GetBrandedDescriptionOk

`func (o *VacanciesVacancyForApplicant) GetBrandedDescriptionOk() (*string, bool)`

GetBrandedDescriptionOk returns a tuple with the BrandedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedDescription

`func (o *VacanciesVacancyForApplicant) SetBrandedDescription(v string)`

SetBrandedDescription sets BrandedDescription field to given value.

### HasBrandedDescription

`func (o *VacanciesVacancyForApplicant) HasBrandedDescription() bool`

HasBrandedDescription returns a boolean if a field has been set.

### GetCode

`func (o *VacanciesVacancyForApplicant) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacanciesVacancyForApplicant) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacanciesVacancyForApplicant) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacanciesVacancyForApplicant) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacanciesVacancyForApplicant) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacanciesVacancyForApplicant) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetContacts

`func (o *VacanciesVacancyForApplicant) GetContacts() VacancyContactsOutput`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacanciesVacancyForApplicant) GetContactsOk() (*VacancyContactsOutput, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacanciesVacancyForApplicant) SetContacts(v VacancyContactsOutput)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacanciesVacancyForApplicant) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacanciesVacancyForApplicant) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacanciesVacancyForApplicant) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCreatedAt

`func (o *VacanciesVacancyForApplicant) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesVacancyForApplicant) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesVacancyForApplicant) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDepartment

`func (o *VacanciesVacancyForApplicant) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesVacancyForApplicant) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesVacancyForApplicant) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacanciesVacancyForApplicant) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacanciesVacancyForApplicant) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesVacancyForApplicant) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDescription

`func (o *VacanciesVacancyForApplicant) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacanciesVacancyForApplicant) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacanciesVacancyForApplicant) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDriverLicenseTypes

`func (o *VacanciesVacancyForApplicant) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacanciesVacancyForApplicant) GetDriverLicenseTypesOk() (*[]VacancyDriverLicenceTypeItem, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacanciesVacancyForApplicant) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEmployer

`func (o *VacanciesVacancyForApplicant) GetEmployer() VacanciesVacancyEmployer`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesVacancyForApplicant) GetEmployerOk() (*VacanciesVacancyEmployer, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesVacancyForApplicant) SetEmployer(v VacanciesVacancyEmployer)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *VacanciesVacancyForApplicant) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### GetEmployment

`func (o *VacanciesVacancyForApplicant) GetEmployment() VacancyEmploymentOutput`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacanciesVacancyForApplicant) GetEmploymentOk() (*VacancyEmploymentOutput, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacanciesVacancyForApplicant) SetEmployment(v VacancyEmploymentOutput)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacanciesVacancyForApplicant) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacanciesVacancyForApplicant) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacanciesVacancyForApplicant) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacanciesVacancyForApplicant) GetExperience() VacancyExperienceOutput`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacanciesVacancyForApplicant) GetExperienceOk() (*VacancyExperienceOutput, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacanciesVacancyForApplicant) SetExperience(v VacancyExperienceOutput)`

SetExperience sets Experience field to given value.


### SetExperienceNil

`func (o *VacanciesVacancyForApplicant) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacanciesVacancyForApplicant) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetHasTest

`func (o *VacanciesVacancyForApplicant) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesVacancyForApplicant) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesVacancyForApplicant) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesVacancyForApplicant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVacancyForApplicant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVacancyForApplicant) SetId(v string)`

SetId sets Id field to given value.


### GetInitialCreatedAt

`func (o *VacanciesVacancyForApplicant) GetInitialCreatedAt() string`

GetInitialCreatedAt returns the InitialCreatedAt field if non-nil, zero value otherwise.

### GetInitialCreatedAtOk

`func (o *VacanciesVacancyForApplicant) GetInitialCreatedAtOk() (*string, bool)`

GetInitialCreatedAtOk returns a tuple with the InitialCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCreatedAt

`func (o *VacanciesVacancyForApplicant) SetInitialCreatedAt(v string)`

SetInitialCreatedAt sets InitialCreatedAt field to given value.


### GetInsiderInterview

`func (o *VacanciesVacancyForApplicant) GetInsiderInterview() VacancyInsiderInterview`

GetInsiderInterview returns the InsiderInterview field if non-nil, zero value otherwise.

### GetInsiderInterviewOk

`func (o *VacanciesVacancyForApplicant) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool)`

GetInsiderInterviewOk returns a tuple with the InsiderInterview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderInterview

`func (o *VacanciesVacancyForApplicant) SetInsiderInterview(v VacancyInsiderInterview)`

SetInsiderInterview sets InsiderInterview field to given value.

### HasInsiderInterview

`func (o *VacanciesVacancyForApplicant) HasInsiderInterview() bool`

HasInsiderInterview returns a boolean if a field has been set.

### GetKeySkills

`func (o *VacanciesVacancyForApplicant) GetKeySkills() []VacancyKeySkillItem`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacanciesVacancyForApplicant) GetKeySkillsOk() (*[]VacancyKeySkillItem, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacanciesVacancyForApplicant) SetKeySkills(v []VacancyKeySkillItem)`

SetKeySkills sets KeySkills field to given value.


### GetLanguages

`func (o *VacanciesVacancyForApplicant) GetLanguages() []VacancyLanguageOutput`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacanciesVacancyForApplicant) GetLanguagesOk() (*[]VacancyLanguageOutput, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacanciesVacancyForApplicant) SetLanguages(v []VacancyLanguageOutput)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VacanciesVacancyForApplicant) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetName

`func (o *VacanciesVacancyForApplicant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesVacancyForApplicant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesVacancyForApplicant) SetName(v string)`

SetName sets Name field to given value.


### GetNegotiationsUrl

`func (o *VacanciesVacancyForApplicant) GetNegotiationsUrl() string`

GetNegotiationsUrl returns the NegotiationsUrl field if non-nil, zero value otherwise.

### GetNegotiationsUrlOk

`func (o *VacanciesVacancyForApplicant) GetNegotiationsUrlOk() (*string, bool)`

GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsUrl

`func (o *VacanciesVacancyForApplicant) SetNegotiationsUrl(v string)`

SetNegotiationsUrl sets NegotiationsUrl field to given value.

### HasNegotiationsUrl

`func (o *VacanciesVacancyForApplicant) HasNegotiationsUrl() bool`

HasNegotiationsUrl returns a boolean if a field has been set.

### SetNegotiationsUrlNil

`func (o *VacanciesVacancyForApplicant) SetNegotiationsUrlNil(b bool)`

 SetNegotiationsUrlNil sets the value for NegotiationsUrl to be an explicit nil

### UnsetNegotiationsUrl
`func (o *VacanciesVacancyForApplicant) UnsetNegotiationsUrl()`

UnsetNegotiationsUrl ensures that no value is present for NegotiationsUrl, not even an explicit nil
### GetPremium

`func (o *VacanciesVacancyForApplicant) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesVacancyForApplicant) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesVacancyForApplicant) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetProfessionalRoles

`func (o *VacanciesVacancyForApplicant) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacanciesVacancyForApplicant) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItemOutput, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacanciesVacancyForApplicant) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput)`

SetProfessionalRoles sets ProfessionalRoles field to given value.


### GetPublishedAt

`func (o *VacanciesVacancyForApplicant) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesVacancyForApplicant) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesVacancyForApplicant) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetRelations

`func (o *VacanciesVacancyForApplicant) GetRelations() []VacancyRelationItem`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *VacanciesVacancyForApplicant) GetRelationsOk() (*[]VacancyRelationItem, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *VacanciesVacancyForApplicant) SetRelations(v []VacancyRelationItem)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *VacanciesVacancyForApplicant) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetResponseLetterRequired

`func (o *VacanciesVacancyForApplicant) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesVacancyForApplicant) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesVacancyForApplicant) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesVacancyForApplicant) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesVacancyForApplicant) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesVacancyForApplicant) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesVacancyForApplicant) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesVacancyForApplicant) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesVacancyForApplicant) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesVacancyForApplicant) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesVacancyForApplicant) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesVacancyForApplicant) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacanciesVacancyForApplicant) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacanciesVacancyForApplicant) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesVacancyForApplicant) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacanciesVacancyForApplicant) GetSchedule() VacancyScheduleOutput`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacanciesVacancyForApplicant) GetScheduleOk() (*VacancyScheduleOutput, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacanciesVacancyForApplicant) SetSchedule(v VacancyScheduleOutput)`

SetSchedule sets Schedule field to given value.


### GetSuitableResumesUrl

`func (o *VacanciesVacancyForApplicant) GetSuitableResumesUrl() string`

GetSuitableResumesUrl returns the SuitableResumesUrl field if non-nil, zero value otherwise.

### GetSuitableResumesUrlOk

`func (o *VacanciesVacancyForApplicant) GetSuitableResumesUrlOk() (*string, bool)`

GetSuitableResumesUrlOk returns a tuple with the SuitableResumesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitableResumesUrl

`func (o *VacanciesVacancyForApplicant) SetSuitableResumesUrl(v string)`

SetSuitableResumesUrl sets SuitableResumesUrl field to given value.

### HasSuitableResumesUrl

`func (o *VacanciesVacancyForApplicant) HasSuitableResumesUrl() bool`

HasSuitableResumesUrl returns a boolean if a field has been set.

### SetSuitableResumesUrlNil

`func (o *VacanciesVacancyForApplicant) SetSuitableResumesUrlNil(b bool)`

 SetSuitableResumesUrlNil sets the value for SuitableResumesUrl to be an explicit nil

### UnsetSuitableResumesUrl
`func (o *VacanciesVacancyForApplicant) UnsetSuitableResumesUrl()`

UnsetSuitableResumesUrl ensures that no value is present for SuitableResumesUrl, not even an explicit nil
### GetTest

`func (o *VacanciesVacancyForApplicant) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacanciesVacancyForApplicant) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacanciesVacancyForApplicant) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacanciesVacancyForApplicant) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacanciesVacancyForApplicant) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacanciesVacancyForApplicant) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetType

`func (o *VacanciesVacancyForApplicant) GetType() IncludesIdName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesVacancyForApplicant) GetTypeOk() (*IncludesIdName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesVacancyForApplicant) SetType(v IncludesIdName)`

SetType sets Type field to given value.


### GetVacancyConstructorTemplate

`func (o *VacanciesVacancyForApplicant) GetVacancyConstructorTemplate() VacancyVacancyConstructorTemplate`

GetVacancyConstructorTemplate returns the VacancyConstructorTemplate field if non-nil, zero value otherwise.

### GetVacancyConstructorTemplateOk

`func (o *VacanciesVacancyForApplicant) GetVacancyConstructorTemplateOk() (*VacancyVacancyConstructorTemplate, bool)`

GetVacancyConstructorTemplateOk returns a tuple with the VacancyConstructorTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyConstructorTemplate

`func (o *VacanciesVacancyForApplicant) SetVacancyConstructorTemplate(v VacancyVacancyConstructorTemplate)`

SetVacancyConstructorTemplate sets VacancyConstructorTemplate field to given value.

### HasVacancyConstructorTemplate

`func (o *VacanciesVacancyForApplicant) HasVacancyConstructorTemplate() bool`

HasVacancyConstructorTemplate returns a boolean if a field has been set.

### GetVideoVacancy

`func (o *VacanciesVacancyForApplicant) GetVideoVacancy() VacancyVideoVacancy`

GetVideoVacancy returns the VideoVacancy field if non-nil, zero value otherwise.

### GetVideoVacancyOk

`func (o *VacanciesVacancyForApplicant) GetVideoVacancyOk() (*VacancyVideoVacancy, bool)`

GetVideoVacancyOk returns a tuple with the VideoVacancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoVacancy

`func (o *VacanciesVacancyForApplicant) SetVideoVacancy(v VacancyVideoVacancy)`

SetVideoVacancy sets VideoVacancy field to given value.

### HasVideoVacancy

`func (o *VacanciesVacancyForApplicant) HasVideoVacancy() bool`

HasVideoVacancy returns a boolean if a field has been set.

### GetWorkingDays

`func (o *VacanciesVacancyForApplicant) GetWorkingDays() []VacancyWorkingDayItemOutput`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacanciesVacancyForApplicant) GetWorkingDaysOk() (*[]VacancyWorkingDayItemOutput, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacanciesVacancyForApplicant) SetWorkingDays(v []VacancyWorkingDayItemOutput)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacanciesVacancyForApplicant) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacanciesVacancyForApplicant) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacanciesVacancyForApplicant) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacanciesVacancyForApplicant) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItemOutput`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacanciesVacancyForApplicant) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItemOutput, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacanciesVacancyForApplicant) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItemOutput)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacanciesVacancyForApplicant) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacanciesVacancyForApplicant) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacanciesVacancyForApplicant) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacanciesVacancyForApplicant) GetWorkingTimeModes() []VacancyWorkingTimeModeItemOutput`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacanciesVacancyForApplicant) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItemOutput, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacanciesVacancyForApplicant) SetWorkingTimeModes(v []VacancyWorkingTimeModeItemOutput)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacanciesVacancyForApplicant) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacanciesVacancyForApplicant) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacanciesVacancyForApplicant) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil
### GetHidden

`func (o *VacanciesVacancyForApplicant) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *VacanciesVacancyForApplicant) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *VacanciesVacancyForApplicant) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *VacanciesVacancyForApplicant) HasHidden() bool`

HasHidden returns a boolean if a field has been set.

### GetAddress

`func (o *VacanciesVacancyForApplicant) GetAddress() VacancyAddressOutput`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesVacancyForApplicant) GetAddressOk() (*VacancyAddressOutput, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesVacancyForApplicant) SetAddress(v VacancyAddressOutput)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacanciesVacancyForApplicant) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacanciesVacancyForApplicant) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacanciesVacancyForApplicant) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


