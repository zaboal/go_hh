# VacanciesVacancyCommonFields

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
**Approved** | **bool** | Прошла ли вакансия модерацию | 
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
**Employer** | Pointer to [**NullableVacanciesVacancyEmployer**](VacanciesVacancyEmployer.md) |  | [optional] 
**Employment** | Pointer to [**NullableVacancyEmploymentOutput**](VacancyEmploymentOutput.md) |  | [optional] 
**Experience** | [**NullableVacancyExperienceOutput**](VacancyExperienceOutput.md) |  | 
**HasTest** | **bool** | Информация о наличии прикрепленного тестового задании к вакансии | 
**Id** | **string** | Идентификатор вакансии | 
**InitialCreatedAt** | **string** | Дата и время создания вакансии | 
**InsiderInterview** | Pointer to [**NullableVacancyInsiderInterview**](VacancyInsiderInterview.md) |  | [optional] 
**KeySkills** | [**[]VacancyKeySkillItem**](VacancyKeySkillItem.md) | Список ключевых навыков, не более 30 | 
**Languages** | Pointer to **[]string** |  | [optional] 
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
**VacancyConstructorTemplate** | Pointer to [**NullableVacancyVacancyConstructorTemplate**](VacancyVacancyConstructorTemplate.md) |  | [optional] 
**VideoVacancy** | Pointer to [**NullableVacancyVideoVacancy**](VacancyVideoVacancy.md) |  | [optional] 
**WorkingDays** | Pointer to [**[]VacancyWorkingDayItemOutput**](VacancyWorkingDayItemOutput.md) | Список рабочих дней | [optional] 
**WorkingTimeIntervals** | Pointer to [**[]VacancyWorkingTimeIntervalItemOutput**](VacancyWorkingTimeIntervalItemOutput.md) | Список с временными интервалами работы | [optional] 
**WorkingTimeModes** | Pointer to [**[]VacancyWorkingTimeModeItemOutput**](VacancyWorkingTimeModeItemOutput.md) | Список режимов времени работы | [optional] 

## Methods

### NewVacanciesVacancyCommonFields

`func NewVacanciesVacancyCommonFields(acceptHandicapped bool, acceptIncompleteResumes bool, acceptKids bool, allowMessages bool, alternateUrl string, applyAlternateUrl string, approved bool, archived bool, area IncludesArea, billingType NullableVacancyBillingTypeOutput, createdAt string, description string, driverLicenseTypes []VacancyDriverLicenceTypeItem, experience NullableVacancyExperienceOutput, hasTest bool, id string, initialCreatedAt string, keySkills []VacancyKeySkillItem, name string, premium bool, professionalRoles []VacancyProfessionalRoleItemOutput, publishedAt string, responseLetterRequired bool, schedule VacancyScheduleOutput, type_ IncludesIdName, ) *VacanciesVacancyCommonFields`

NewVacanciesVacancyCommonFields instantiates a new VacanciesVacancyCommonFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyCommonFieldsWithDefaults

`func NewVacanciesVacancyCommonFieldsWithDefaults() *VacanciesVacancyCommonFields`

NewVacanciesVacancyCommonFieldsWithDefaults instantiates a new VacanciesVacancyCommonFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacanciesVacancyCommonFields) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacanciesVacancyCommonFields) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacanciesVacancyCommonFields) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.


### GetAcceptIncompleteResumes

`func (o *VacanciesVacancyCommonFields) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacanciesVacancyCommonFields) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacanciesVacancyCommonFields) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.


### GetAcceptKids

`func (o *VacanciesVacancyCommonFields) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacanciesVacancyCommonFields) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacanciesVacancyCommonFields) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.


### GetAcceptTemporary

`func (o *VacanciesVacancyCommonFields) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacanciesVacancyCommonFields) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacanciesVacancyCommonFields) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacanciesVacancyCommonFields) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacanciesVacancyCommonFields) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacanciesVacancyCommonFields) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAllowMessages

`func (o *VacanciesVacancyCommonFields) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacanciesVacancyCommonFields) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacanciesVacancyCommonFields) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.


### GetAlternateUrl

`func (o *VacanciesVacancyCommonFields) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVacancyCommonFields) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVacancyCommonFields) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesVacancyCommonFields) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesVacancyCommonFields) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesVacancyCommonFields) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetApproved

`func (o *VacanciesVacancyCommonFields) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *VacanciesVacancyCommonFields) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *VacanciesVacancyCommonFields) SetApproved(v bool)`

SetApproved sets Approved field to given value.


### GetArchived

`func (o *VacanciesVacancyCommonFields) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesVacancyCommonFields) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesVacancyCommonFields) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArea

`func (o *VacanciesVacancyCommonFields) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesVacancyCommonFields) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesVacancyCommonFields) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetBillingType

`func (o *VacanciesVacancyCommonFields) GetBillingType() VacancyBillingTypeOutput`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacanciesVacancyCommonFields) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacanciesVacancyCommonFields) SetBillingType(v VacancyBillingTypeOutput)`

SetBillingType sets BillingType field to given value.


### SetBillingTypeNil

`func (o *VacanciesVacancyCommonFields) SetBillingTypeNil(b bool)`

 SetBillingTypeNil sets the value for BillingType to be an explicit nil

### UnsetBillingType
`func (o *VacanciesVacancyCommonFields) UnsetBillingType()`

UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
### GetBrandedDescription

`func (o *VacanciesVacancyCommonFields) GetBrandedDescription() string`

GetBrandedDescription returns the BrandedDescription field if non-nil, zero value otherwise.

### GetBrandedDescriptionOk

`func (o *VacanciesVacancyCommonFields) GetBrandedDescriptionOk() (*string, bool)`

GetBrandedDescriptionOk returns a tuple with the BrandedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedDescription

`func (o *VacanciesVacancyCommonFields) SetBrandedDescription(v string)`

SetBrandedDescription sets BrandedDescription field to given value.

### HasBrandedDescription

`func (o *VacanciesVacancyCommonFields) HasBrandedDescription() bool`

HasBrandedDescription returns a boolean if a field has been set.

### GetCode

`func (o *VacanciesVacancyCommonFields) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacanciesVacancyCommonFields) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacanciesVacancyCommonFields) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacanciesVacancyCommonFields) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacanciesVacancyCommonFields) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacanciesVacancyCommonFields) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetContacts

`func (o *VacanciesVacancyCommonFields) GetContacts() VacancyContactsOutput`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacanciesVacancyCommonFields) GetContactsOk() (*VacancyContactsOutput, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacanciesVacancyCommonFields) SetContacts(v VacancyContactsOutput)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacanciesVacancyCommonFields) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacanciesVacancyCommonFields) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacanciesVacancyCommonFields) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCreatedAt

`func (o *VacanciesVacancyCommonFields) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesVacancyCommonFields) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesVacancyCommonFields) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDepartment

`func (o *VacanciesVacancyCommonFields) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesVacancyCommonFields) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesVacancyCommonFields) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacanciesVacancyCommonFields) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacanciesVacancyCommonFields) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesVacancyCommonFields) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDescription

`func (o *VacanciesVacancyCommonFields) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacanciesVacancyCommonFields) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacanciesVacancyCommonFields) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDriverLicenseTypes

`func (o *VacanciesVacancyCommonFields) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacanciesVacancyCommonFields) GetDriverLicenseTypesOk() (*[]VacancyDriverLicenceTypeItem, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacanciesVacancyCommonFields) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEmployer

`func (o *VacanciesVacancyCommonFields) GetEmployer() VacanciesVacancyEmployer`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesVacancyCommonFields) GetEmployerOk() (*VacanciesVacancyEmployer, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesVacancyCommonFields) SetEmployer(v VacanciesVacancyEmployer)`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *VacanciesVacancyCommonFields) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### SetEmployerNil

`func (o *VacanciesVacancyCommonFields) SetEmployerNil(b bool)`

 SetEmployerNil sets the value for Employer to be an explicit nil

### UnsetEmployer
`func (o *VacanciesVacancyCommonFields) UnsetEmployer()`

UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
### GetEmployment

`func (o *VacanciesVacancyCommonFields) GetEmployment() VacancyEmploymentOutput`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacanciesVacancyCommonFields) GetEmploymentOk() (*VacancyEmploymentOutput, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacanciesVacancyCommonFields) SetEmployment(v VacancyEmploymentOutput)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacanciesVacancyCommonFields) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacanciesVacancyCommonFields) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacanciesVacancyCommonFields) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacanciesVacancyCommonFields) GetExperience() VacancyExperienceOutput`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacanciesVacancyCommonFields) GetExperienceOk() (*VacancyExperienceOutput, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacanciesVacancyCommonFields) SetExperience(v VacancyExperienceOutput)`

SetExperience sets Experience field to given value.


### SetExperienceNil

`func (o *VacanciesVacancyCommonFields) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacanciesVacancyCommonFields) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetHasTest

`func (o *VacanciesVacancyCommonFields) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesVacancyCommonFields) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesVacancyCommonFields) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesVacancyCommonFields) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVacancyCommonFields) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVacancyCommonFields) SetId(v string)`

SetId sets Id field to given value.


### GetInitialCreatedAt

`func (o *VacanciesVacancyCommonFields) GetInitialCreatedAt() string`

GetInitialCreatedAt returns the InitialCreatedAt field if non-nil, zero value otherwise.

### GetInitialCreatedAtOk

`func (o *VacanciesVacancyCommonFields) GetInitialCreatedAtOk() (*string, bool)`

GetInitialCreatedAtOk returns a tuple with the InitialCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialCreatedAt

`func (o *VacanciesVacancyCommonFields) SetInitialCreatedAt(v string)`

SetInitialCreatedAt sets InitialCreatedAt field to given value.


### GetInsiderInterview

`func (o *VacanciesVacancyCommonFields) GetInsiderInterview() VacancyInsiderInterview`

GetInsiderInterview returns the InsiderInterview field if non-nil, zero value otherwise.

### GetInsiderInterviewOk

`func (o *VacanciesVacancyCommonFields) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool)`

GetInsiderInterviewOk returns a tuple with the InsiderInterview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderInterview

`func (o *VacanciesVacancyCommonFields) SetInsiderInterview(v VacancyInsiderInterview)`

SetInsiderInterview sets InsiderInterview field to given value.

### HasInsiderInterview

`func (o *VacanciesVacancyCommonFields) HasInsiderInterview() bool`

HasInsiderInterview returns a boolean if a field has been set.

### SetInsiderInterviewNil

`func (o *VacanciesVacancyCommonFields) SetInsiderInterviewNil(b bool)`

 SetInsiderInterviewNil sets the value for InsiderInterview to be an explicit nil

### UnsetInsiderInterview
`func (o *VacanciesVacancyCommonFields) UnsetInsiderInterview()`

UnsetInsiderInterview ensures that no value is present for InsiderInterview, not even an explicit nil
### GetKeySkills

`func (o *VacanciesVacancyCommonFields) GetKeySkills() []VacancyKeySkillItem`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacanciesVacancyCommonFields) GetKeySkillsOk() (*[]VacancyKeySkillItem, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacanciesVacancyCommonFields) SetKeySkills(v []VacancyKeySkillItem)`

SetKeySkills sets KeySkills field to given value.


### GetLanguages

`func (o *VacanciesVacancyCommonFields) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacanciesVacancyCommonFields) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacanciesVacancyCommonFields) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VacanciesVacancyCommonFields) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *VacanciesVacancyCommonFields) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *VacanciesVacancyCommonFields) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetName

`func (o *VacanciesVacancyCommonFields) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesVacancyCommonFields) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesVacancyCommonFields) SetName(v string)`

SetName sets Name field to given value.


### GetNegotiationsUrl

`func (o *VacanciesVacancyCommonFields) GetNegotiationsUrl() string`

GetNegotiationsUrl returns the NegotiationsUrl field if non-nil, zero value otherwise.

### GetNegotiationsUrlOk

`func (o *VacanciesVacancyCommonFields) GetNegotiationsUrlOk() (*string, bool)`

GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsUrl

`func (o *VacanciesVacancyCommonFields) SetNegotiationsUrl(v string)`

SetNegotiationsUrl sets NegotiationsUrl field to given value.

### HasNegotiationsUrl

`func (o *VacanciesVacancyCommonFields) HasNegotiationsUrl() bool`

HasNegotiationsUrl returns a boolean if a field has been set.

### SetNegotiationsUrlNil

`func (o *VacanciesVacancyCommonFields) SetNegotiationsUrlNil(b bool)`

 SetNegotiationsUrlNil sets the value for NegotiationsUrl to be an explicit nil

### UnsetNegotiationsUrl
`func (o *VacanciesVacancyCommonFields) UnsetNegotiationsUrl()`

UnsetNegotiationsUrl ensures that no value is present for NegotiationsUrl, not even an explicit nil
### GetPremium

`func (o *VacanciesVacancyCommonFields) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesVacancyCommonFields) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesVacancyCommonFields) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetProfessionalRoles

`func (o *VacanciesVacancyCommonFields) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacanciesVacancyCommonFields) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItemOutput, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacanciesVacancyCommonFields) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput)`

SetProfessionalRoles sets ProfessionalRoles field to given value.


### GetPublishedAt

`func (o *VacanciesVacancyCommonFields) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesVacancyCommonFields) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesVacancyCommonFields) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetRelations

`func (o *VacanciesVacancyCommonFields) GetRelations() []VacancyRelationItem`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *VacanciesVacancyCommonFields) GetRelationsOk() (*[]VacancyRelationItem, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *VacanciesVacancyCommonFields) SetRelations(v []VacancyRelationItem)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *VacanciesVacancyCommonFields) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetResponseLetterRequired

`func (o *VacanciesVacancyCommonFields) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesVacancyCommonFields) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesVacancyCommonFields) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesVacancyCommonFields) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesVacancyCommonFields) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesVacancyCommonFields) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesVacancyCommonFields) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesVacancyCommonFields) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesVacancyCommonFields) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesVacancyCommonFields) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesVacancyCommonFields) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesVacancyCommonFields) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacanciesVacancyCommonFields) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacanciesVacancyCommonFields) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesVacancyCommonFields) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacanciesVacancyCommonFields) GetSchedule() VacancyScheduleOutput`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacanciesVacancyCommonFields) GetScheduleOk() (*VacancyScheduleOutput, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacanciesVacancyCommonFields) SetSchedule(v VacancyScheduleOutput)`

SetSchedule sets Schedule field to given value.


### GetSuitableResumesUrl

`func (o *VacanciesVacancyCommonFields) GetSuitableResumesUrl() string`

GetSuitableResumesUrl returns the SuitableResumesUrl field if non-nil, zero value otherwise.

### GetSuitableResumesUrlOk

`func (o *VacanciesVacancyCommonFields) GetSuitableResumesUrlOk() (*string, bool)`

GetSuitableResumesUrlOk returns a tuple with the SuitableResumesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitableResumesUrl

`func (o *VacanciesVacancyCommonFields) SetSuitableResumesUrl(v string)`

SetSuitableResumesUrl sets SuitableResumesUrl field to given value.

### HasSuitableResumesUrl

`func (o *VacanciesVacancyCommonFields) HasSuitableResumesUrl() bool`

HasSuitableResumesUrl returns a boolean if a field has been set.

### SetSuitableResumesUrlNil

`func (o *VacanciesVacancyCommonFields) SetSuitableResumesUrlNil(b bool)`

 SetSuitableResumesUrlNil sets the value for SuitableResumesUrl to be an explicit nil

### UnsetSuitableResumesUrl
`func (o *VacanciesVacancyCommonFields) UnsetSuitableResumesUrl()`

UnsetSuitableResumesUrl ensures that no value is present for SuitableResumesUrl, not even an explicit nil
### GetTest

`func (o *VacanciesVacancyCommonFields) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacanciesVacancyCommonFields) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacanciesVacancyCommonFields) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacanciesVacancyCommonFields) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacanciesVacancyCommonFields) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacanciesVacancyCommonFields) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetType

`func (o *VacanciesVacancyCommonFields) GetType() IncludesIdName`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesVacancyCommonFields) GetTypeOk() (*IncludesIdName, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesVacancyCommonFields) SetType(v IncludesIdName)`

SetType sets Type field to given value.


### GetVacancyConstructorTemplate

`func (o *VacanciesVacancyCommonFields) GetVacancyConstructorTemplate() VacancyVacancyConstructorTemplate`

GetVacancyConstructorTemplate returns the VacancyConstructorTemplate field if non-nil, zero value otherwise.

### GetVacancyConstructorTemplateOk

`func (o *VacanciesVacancyCommonFields) GetVacancyConstructorTemplateOk() (*VacancyVacancyConstructorTemplate, bool)`

GetVacancyConstructorTemplateOk returns a tuple with the VacancyConstructorTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyConstructorTemplate

`func (o *VacanciesVacancyCommonFields) SetVacancyConstructorTemplate(v VacancyVacancyConstructorTemplate)`

SetVacancyConstructorTemplate sets VacancyConstructorTemplate field to given value.

### HasVacancyConstructorTemplate

`func (o *VacanciesVacancyCommonFields) HasVacancyConstructorTemplate() bool`

HasVacancyConstructorTemplate returns a boolean if a field has been set.

### SetVacancyConstructorTemplateNil

`func (o *VacanciesVacancyCommonFields) SetVacancyConstructorTemplateNil(b bool)`

 SetVacancyConstructorTemplateNil sets the value for VacancyConstructorTemplate to be an explicit nil

### UnsetVacancyConstructorTemplate
`func (o *VacanciesVacancyCommonFields) UnsetVacancyConstructorTemplate()`

UnsetVacancyConstructorTemplate ensures that no value is present for VacancyConstructorTemplate, not even an explicit nil
### GetVideoVacancy

`func (o *VacanciesVacancyCommonFields) GetVideoVacancy() VacancyVideoVacancy`

GetVideoVacancy returns the VideoVacancy field if non-nil, zero value otherwise.

### GetVideoVacancyOk

`func (o *VacanciesVacancyCommonFields) GetVideoVacancyOk() (*VacancyVideoVacancy, bool)`

GetVideoVacancyOk returns a tuple with the VideoVacancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoVacancy

`func (o *VacanciesVacancyCommonFields) SetVideoVacancy(v VacancyVideoVacancy)`

SetVideoVacancy sets VideoVacancy field to given value.

### HasVideoVacancy

`func (o *VacanciesVacancyCommonFields) HasVideoVacancy() bool`

HasVideoVacancy returns a boolean if a field has been set.

### SetVideoVacancyNil

`func (o *VacanciesVacancyCommonFields) SetVideoVacancyNil(b bool)`

 SetVideoVacancyNil sets the value for VideoVacancy to be an explicit nil

### UnsetVideoVacancy
`func (o *VacanciesVacancyCommonFields) UnsetVideoVacancy()`

UnsetVideoVacancy ensures that no value is present for VideoVacancy, not even an explicit nil
### GetWorkingDays

`func (o *VacanciesVacancyCommonFields) GetWorkingDays() []VacancyWorkingDayItemOutput`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacanciesVacancyCommonFields) GetWorkingDaysOk() (*[]VacancyWorkingDayItemOutput, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacanciesVacancyCommonFields) SetWorkingDays(v []VacancyWorkingDayItemOutput)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacanciesVacancyCommonFields) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacanciesVacancyCommonFields) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacanciesVacancyCommonFields) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacanciesVacancyCommonFields) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItemOutput`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacanciesVacancyCommonFields) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItemOutput, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacanciesVacancyCommonFields) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItemOutput)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacanciesVacancyCommonFields) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacanciesVacancyCommonFields) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacanciesVacancyCommonFields) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacanciesVacancyCommonFields) GetWorkingTimeModes() []VacancyWorkingTimeModeItemOutput`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacanciesVacancyCommonFields) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItemOutput, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacanciesVacancyCommonFields) SetWorkingTimeModes(v []VacancyWorkingTimeModeItemOutput)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacanciesVacancyCommonFields) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacanciesVacancyCommonFields) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacanciesVacancyCommonFields) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


