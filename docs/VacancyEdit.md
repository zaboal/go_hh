# VacancyEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptHandicapped** | Pointer to **bool** | Указание, что вакансия доступна для соискателей с инвалидностью | [optional] 
**AcceptIncompleteResumes** | Pointer to **bool** | Разрешен ли отклик на вакансию неполным резюме | [optional] 
**AcceptKids** | Pointer to **bool** | Указание, что вакансия доступна для соискателей старше 14 лет [подробнее](https://github.com/hhru/api/blob/master/docs/employer_vacancies_accept_kids.md#accept-kids) | [optional] 
**AcceptTemporary** | Pointer to **NullableBool** | Указание, что вакансия доступна с временным трудоустройством | [optional] 
**Address** | Pointer to [**NullableVacancyAddress**](VacancyAddress.md) |  | [optional] 
**AllowMessages** | Pointer to **bool** | Возможность [переписки с кандидатами](https://inboxemp.hh.ru/) по данной вакансии | [optional] 
**BrandedTemplate** | Pointer to [**NullableVacancyBrandedTemplate**](VacancyBrandedTemplate.md) |  | [optional] 
**Code** | Pointer to **NullableString** | Внутренний код вакансии | [optional] 
**Contacts** | Pointer to [**NullableVacancyContacts**](VacancyContacts.md) |  | [optional] 
**CustomEmployerName** | Pointer to **NullableString** | Название компании для анонимных вакансий (&#x60;type.id&#x3D;anonymous&#x60;), например \&quot;крупный российский банк\&quot;. Поле становится обязательным при публикации вакансии **анонимного** типа | [optional] 
**Department** | Pointer to [**NullableVacancyDepartment**](VacancyDepartment.md) |  | [optional] 
**DriverLicenseTypes** | Pointer to [**[]VacancyDriverLicenceTypeItem**](VacancyDriverLicenceTypeItem.md) | Список требуемых категорий водительских прав | [optional] 
**Employment** | Pointer to [**NullableVacancyEmployment**](VacancyEmployment.md) |  | [optional] 
**Experience** | Pointer to [**NullableVacancyExperience**](VacancyExperience.md) |  | [optional] 
**KeySkills** | Pointer to [**[]VacancyKeySkillItem**](VacancyKeySkillItem.md) | Список ключевых навыков, не более 30 | [optional] 
**Languages** | Pointer to [**[]VacancyLanguage**](VacancyLanguage.md) | Список языков вакансии | [optional] 
**ProfessionalRoles** | Pointer to [**[]VacancyProfessionalRoleItem**](VacancyProfessionalRoleItem.md) | Список профессиональных ролей | [optional] 
**ResponseLetterRequired** | Pointer to **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | [optional] 
**ResponseNotifications** | Pointer to **bool** | Уведомлять ли менеджера о новых откликах | [optional] 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | Pointer to [**NullableVacancySalary**](VacancySalary.md) |  | [optional] 
**Schedule** | Pointer to [**NullableVacancySchedule**](VacancySchedule.md) |  | [optional] 
**Test** | Pointer to [**NullableVacancyDraftTest**](VacancyDraftTest.md) |  | [optional] 
**WorkingDays** | Pointer to [**[]VacancyWorkingDayItem**](VacancyWorkingDayItem.md) | Список рабочих дней | [optional] 
**WorkingTimeIntervals** | Pointer to [**[]VacancyWorkingTimeIntervalItem**](VacancyWorkingTimeIntervalItem.md) | Список с временными интервалами работы | [optional] 
**WorkingTimeModes** | Pointer to [**[]VacancyWorkingTimeModeItem**](VacancyWorkingTimeModeItem.md) | Список режимов времени работы | [optional] 
**Description** | Pointer to **string** | Описание в html, не менее 200 символов | [optional] 
**Name** | Pointer to **string** | Название | [optional] 
**BillingType** | [**VacancyBillingType**](VacancyBillingType.md) |  | 
**Manager** | [**NullableIncludesId**](IncludesId.md) | Идентификатор менеджера, которому надо передать автопоиск из ([списка менеджеров компании](#tag/Menedzhery-rabotodatelya/operation/get-employer-managers)) | 

## Methods

### NewVacancyEdit

`func NewVacancyEdit(billingType VacancyBillingType, manager NullableIncludesId, ) *VacancyEdit`

NewVacancyEdit instantiates a new VacancyEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyEditWithDefaults

`func NewVacancyEditWithDefaults() *VacancyEdit`

NewVacancyEditWithDefaults instantiates a new VacancyEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacancyEdit) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacancyEdit) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacancyEdit) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.

### HasAcceptHandicapped

`func (o *VacancyEdit) HasAcceptHandicapped() bool`

HasAcceptHandicapped returns a boolean if a field has been set.

### GetAcceptIncompleteResumes

`func (o *VacancyEdit) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacancyEdit) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacancyEdit) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.

### HasAcceptIncompleteResumes

`func (o *VacancyEdit) HasAcceptIncompleteResumes() bool`

HasAcceptIncompleteResumes returns a boolean if a field has been set.

### GetAcceptKids

`func (o *VacancyEdit) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacancyEdit) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacancyEdit) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.

### HasAcceptKids

`func (o *VacancyEdit) HasAcceptKids() bool`

HasAcceptKids returns a boolean if a field has been set.

### GetAcceptTemporary

`func (o *VacancyEdit) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacancyEdit) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacancyEdit) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacancyEdit) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacancyEdit) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacancyEdit) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAddress

`func (o *VacancyEdit) GetAddress() VacancyAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacancyEdit) GetAddressOk() (*VacancyAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacancyEdit) SetAddress(v VacancyAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacancyEdit) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacancyEdit) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacancyEdit) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAllowMessages

`func (o *VacancyEdit) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacancyEdit) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacancyEdit) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.

### HasAllowMessages

`func (o *VacancyEdit) HasAllowMessages() bool`

HasAllowMessages returns a boolean if a field has been set.

### GetBrandedTemplate

`func (o *VacancyEdit) GetBrandedTemplate() VacancyBrandedTemplate`

GetBrandedTemplate returns the BrandedTemplate field if non-nil, zero value otherwise.

### GetBrandedTemplateOk

`func (o *VacancyEdit) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool)`

GetBrandedTemplateOk returns a tuple with the BrandedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedTemplate

`func (o *VacancyEdit) SetBrandedTemplate(v VacancyBrandedTemplate)`

SetBrandedTemplate sets BrandedTemplate field to given value.

### HasBrandedTemplate

`func (o *VacancyEdit) HasBrandedTemplate() bool`

HasBrandedTemplate returns a boolean if a field has been set.

### SetBrandedTemplateNil

`func (o *VacancyEdit) SetBrandedTemplateNil(b bool)`

 SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil

### UnsetBrandedTemplate
`func (o *VacancyEdit) UnsetBrandedTemplate()`

UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
### GetCode

`func (o *VacancyEdit) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacancyEdit) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacancyEdit) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacancyEdit) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacancyEdit) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacancyEdit) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetContacts

`func (o *VacancyEdit) GetContacts() VacancyContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacancyEdit) GetContactsOk() (*VacancyContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacancyEdit) SetContacts(v VacancyContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacancyEdit) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacancyEdit) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacancyEdit) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCustomEmployerName

`func (o *VacancyEdit) GetCustomEmployerName() string`

GetCustomEmployerName returns the CustomEmployerName field if non-nil, zero value otherwise.

### GetCustomEmployerNameOk

`func (o *VacancyEdit) GetCustomEmployerNameOk() (*string, bool)`

GetCustomEmployerNameOk returns a tuple with the CustomEmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmployerName

`func (o *VacancyEdit) SetCustomEmployerName(v string)`

SetCustomEmployerName sets CustomEmployerName field to given value.

### HasCustomEmployerName

`func (o *VacancyEdit) HasCustomEmployerName() bool`

HasCustomEmployerName returns a boolean if a field has been set.

### SetCustomEmployerNameNil

`func (o *VacancyEdit) SetCustomEmployerNameNil(b bool)`

 SetCustomEmployerNameNil sets the value for CustomEmployerName to be an explicit nil

### UnsetCustomEmployerName
`func (o *VacancyEdit) UnsetCustomEmployerName()`

UnsetCustomEmployerName ensures that no value is present for CustomEmployerName, not even an explicit nil
### GetDepartment

`func (o *VacancyEdit) GetDepartment() VacancyDepartment`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacancyEdit) GetDepartmentOk() (*VacancyDepartment, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacancyEdit) SetDepartment(v VacancyDepartment)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacancyEdit) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacancyEdit) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacancyEdit) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDriverLicenseTypes

`func (o *VacancyEdit) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacancyEdit) GetDriverLicenseTypesOk() (*[]VacancyDriverLicenceTypeItem, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacancyEdit) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.

### HasDriverLicenseTypes

`func (o *VacancyEdit) HasDriverLicenseTypes() bool`

HasDriverLicenseTypes returns a boolean if a field has been set.

### GetEmployment

`func (o *VacancyEdit) GetEmployment() VacancyEmployment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacancyEdit) GetEmploymentOk() (*VacancyEmployment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacancyEdit) SetEmployment(v VacancyEmployment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacancyEdit) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacancyEdit) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacancyEdit) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacancyEdit) GetExperience() VacancyExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacancyEdit) GetExperienceOk() (*VacancyExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacancyEdit) SetExperience(v VacancyExperience)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *VacancyEdit) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *VacancyEdit) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacancyEdit) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetKeySkills

`func (o *VacancyEdit) GetKeySkills() []VacancyKeySkillItem`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacancyEdit) GetKeySkillsOk() (*[]VacancyKeySkillItem, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacancyEdit) SetKeySkills(v []VacancyKeySkillItem)`

SetKeySkills sets KeySkills field to given value.

### HasKeySkills

`func (o *VacancyEdit) HasKeySkills() bool`

HasKeySkills returns a boolean if a field has been set.

### GetLanguages

`func (o *VacancyEdit) GetLanguages() []VacancyLanguage`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacancyEdit) GetLanguagesOk() (*[]VacancyLanguage, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacancyEdit) SetLanguages(v []VacancyLanguage)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VacancyEdit) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetProfessionalRoles

`func (o *VacancyEdit) GetProfessionalRoles() []VacancyProfessionalRoleItem`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacancyEdit) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItem, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacancyEdit) SetProfessionalRoles(v []VacancyProfessionalRoleItem)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *VacancyEdit) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetResponseLetterRequired

`func (o *VacancyEdit) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacancyEdit) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacancyEdit) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.

### HasResponseLetterRequired

`func (o *VacancyEdit) HasResponseLetterRequired() bool`

HasResponseLetterRequired returns a boolean if a field has been set.

### GetResponseNotifications

`func (o *VacancyEdit) GetResponseNotifications() bool`

GetResponseNotifications returns the ResponseNotifications field if non-nil, zero value otherwise.

### GetResponseNotificationsOk

`func (o *VacancyEdit) GetResponseNotificationsOk() (*bool, bool)`

GetResponseNotificationsOk returns a tuple with the ResponseNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseNotifications

`func (o *VacancyEdit) SetResponseNotifications(v bool)`

SetResponseNotifications sets ResponseNotifications field to given value.

### HasResponseNotifications

`func (o *VacancyEdit) HasResponseNotifications() bool`

HasResponseNotifications returns a boolean if a field has been set.

### GetResponseUrl

`func (o *VacancyEdit) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacancyEdit) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacancyEdit) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacancyEdit) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacancyEdit) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacancyEdit) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacancyEdit) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacancyEdit) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacancyEdit) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacancyEdit) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacancyEdit) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacancyEdit) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacancyEdit) GetSchedule() VacancySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacancyEdit) GetScheduleOk() (*VacancySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacancyEdit) SetSchedule(v VacancySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *VacancyEdit) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *VacancyEdit) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *VacancyEdit) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetTest

`func (o *VacancyEdit) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacancyEdit) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacancyEdit) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacancyEdit) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacancyEdit) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacancyEdit) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetWorkingDays

`func (o *VacancyEdit) GetWorkingDays() []VacancyWorkingDayItem`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacancyEdit) GetWorkingDaysOk() (*[]VacancyWorkingDayItem, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacancyEdit) SetWorkingDays(v []VacancyWorkingDayItem)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacancyEdit) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacancyEdit) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacancyEdit) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacancyEdit) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItem`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacancyEdit) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItem, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacancyEdit) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItem)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacancyEdit) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacancyEdit) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacancyEdit) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacancyEdit) GetWorkingTimeModes() []VacancyWorkingTimeModeItem`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacancyEdit) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItem, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacancyEdit) SetWorkingTimeModes(v []VacancyWorkingTimeModeItem)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacancyEdit) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacancyEdit) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacancyEdit) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil
### GetDescription

`func (o *VacancyEdit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacancyEdit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacancyEdit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VacancyEdit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VacancyEdit) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyEdit) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyEdit) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacancyEdit) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBillingType

`func (o *VacancyEdit) GetBillingType() VacancyBillingType`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacancyEdit) GetBillingTypeOk() (*VacancyBillingType, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacancyEdit) SetBillingType(v VacancyBillingType)`

SetBillingType sets BillingType field to given value.


### GetManager

`func (o *VacancyEdit) GetManager() IncludesId`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacancyEdit) GetManagerOk() (*IncludesId, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacancyEdit) SetManager(v IncludesId)`

SetManager sets Manager field to given value.


### SetManagerNil

`func (o *VacancyEdit) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *VacancyEdit) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


