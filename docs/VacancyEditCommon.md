# VacancyEditCommon

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

## Methods

### NewVacancyEditCommon

`func NewVacancyEditCommon() *VacancyEditCommon`

NewVacancyEditCommon instantiates a new VacancyEditCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyEditCommonWithDefaults

`func NewVacancyEditCommonWithDefaults() *VacancyEditCommon`

NewVacancyEditCommonWithDefaults instantiates a new VacancyEditCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacancyEditCommon) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacancyEditCommon) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacancyEditCommon) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.

### HasAcceptHandicapped

`func (o *VacancyEditCommon) HasAcceptHandicapped() bool`

HasAcceptHandicapped returns a boolean if a field has been set.

### GetAcceptIncompleteResumes

`func (o *VacancyEditCommon) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacancyEditCommon) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacancyEditCommon) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.

### HasAcceptIncompleteResumes

`func (o *VacancyEditCommon) HasAcceptIncompleteResumes() bool`

HasAcceptIncompleteResumes returns a boolean if a field has been set.

### GetAcceptKids

`func (o *VacancyEditCommon) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacancyEditCommon) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacancyEditCommon) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.

### HasAcceptKids

`func (o *VacancyEditCommon) HasAcceptKids() bool`

HasAcceptKids returns a boolean if a field has been set.

### GetAcceptTemporary

`func (o *VacancyEditCommon) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacancyEditCommon) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacancyEditCommon) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacancyEditCommon) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacancyEditCommon) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacancyEditCommon) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAddress

`func (o *VacancyEditCommon) GetAddress() VacancyAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacancyEditCommon) GetAddressOk() (*VacancyAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacancyEditCommon) SetAddress(v VacancyAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacancyEditCommon) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacancyEditCommon) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacancyEditCommon) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAllowMessages

`func (o *VacancyEditCommon) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacancyEditCommon) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacancyEditCommon) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.

### HasAllowMessages

`func (o *VacancyEditCommon) HasAllowMessages() bool`

HasAllowMessages returns a boolean if a field has been set.

### GetBrandedTemplate

`func (o *VacancyEditCommon) GetBrandedTemplate() VacancyBrandedTemplate`

GetBrandedTemplate returns the BrandedTemplate field if non-nil, zero value otherwise.

### GetBrandedTemplateOk

`func (o *VacancyEditCommon) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool)`

GetBrandedTemplateOk returns a tuple with the BrandedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedTemplate

`func (o *VacancyEditCommon) SetBrandedTemplate(v VacancyBrandedTemplate)`

SetBrandedTemplate sets BrandedTemplate field to given value.

### HasBrandedTemplate

`func (o *VacancyEditCommon) HasBrandedTemplate() bool`

HasBrandedTemplate returns a boolean if a field has been set.

### SetBrandedTemplateNil

`func (o *VacancyEditCommon) SetBrandedTemplateNil(b bool)`

 SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil

### UnsetBrandedTemplate
`func (o *VacancyEditCommon) UnsetBrandedTemplate()`

UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
### GetCode

`func (o *VacancyEditCommon) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacancyEditCommon) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacancyEditCommon) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacancyEditCommon) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacancyEditCommon) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacancyEditCommon) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetContacts

`func (o *VacancyEditCommon) GetContacts() VacancyContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacancyEditCommon) GetContactsOk() (*VacancyContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacancyEditCommon) SetContacts(v VacancyContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacancyEditCommon) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacancyEditCommon) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacancyEditCommon) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCustomEmployerName

`func (o *VacancyEditCommon) GetCustomEmployerName() string`

GetCustomEmployerName returns the CustomEmployerName field if non-nil, zero value otherwise.

### GetCustomEmployerNameOk

`func (o *VacancyEditCommon) GetCustomEmployerNameOk() (*string, bool)`

GetCustomEmployerNameOk returns a tuple with the CustomEmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmployerName

`func (o *VacancyEditCommon) SetCustomEmployerName(v string)`

SetCustomEmployerName sets CustomEmployerName field to given value.

### HasCustomEmployerName

`func (o *VacancyEditCommon) HasCustomEmployerName() bool`

HasCustomEmployerName returns a boolean if a field has been set.

### SetCustomEmployerNameNil

`func (o *VacancyEditCommon) SetCustomEmployerNameNil(b bool)`

 SetCustomEmployerNameNil sets the value for CustomEmployerName to be an explicit nil

### UnsetCustomEmployerName
`func (o *VacancyEditCommon) UnsetCustomEmployerName()`

UnsetCustomEmployerName ensures that no value is present for CustomEmployerName, not even an explicit nil
### GetDepartment

`func (o *VacancyEditCommon) GetDepartment() VacancyDepartment`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacancyEditCommon) GetDepartmentOk() (*VacancyDepartment, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacancyEditCommon) SetDepartment(v VacancyDepartment)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacancyEditCommon) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacancyEditCommon) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacancyEditCommon) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDriverLicenseTypes

`func (o *VacancyEditCommon) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacancyEditCommon) GetDriverLicenseTypesOk() (*[]VacancyDriverLicenceTypeItem, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacancyEditCommon) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.

### HasDriverLicenseTypes

`func (o *VacancyEditCommon) HasDriverLicenseTypes() bool`

HasDriverLicenseTypes returns a boolean if a field has been set.

### GetEmployment

`func (o *VacancyEditCommon) GetEmployment() VacancyEmployment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacancyEditCommon) GetEmploymentOk() (*VacancyEmployment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacancyEditCommon) SetEmployment(v VacancyEmployment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacancyEditCommon) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacancyEditCommon) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacancyEditCommon) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacancyEditCommon) GetExperience() VacancyExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacancyEditCommon) GetExperienceOk() (*VacancyExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacancyEditCommon) SetExperience(v VacancyExperience)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *VacancyEditCommon) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *VacancyEditCommon) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacancyEditCommon) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetKeySkills

`func (o *VacancyEditCommon) GetKeySkills() []VacancyKeySkillItem`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacancyEditCommon) GetKeySkillsOk() (*[]VacancyKeySkillItem, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacancyEditCommon) SetKeySkills(v []VacancyKeySkillItem)`

SetKeySkills sets KeySkills field to given value.

### HasKeySkills

`func (o *VacancyEditCommon) HasKeySkills() bool`

HasKeySkills returns a boolean if a field has been set.

### GetLanguages

`func (o *VacancyEditCommon) GetLanguages() []VacancyLanguage`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacancyEditCommon) GetLanguagesOk() (*[]VacancyLanguage, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacancyEditCommon) SetLanguages(v []VacancyLanguage)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VacancyEditCommon) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetProfessionalRoles

`func (o *VacancyEditCommon) GetProfessionalRoles() []VacancyProfessionalRoleItem`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacancyEditCommon) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItem, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacancyEditCommon) SetProfessionalRoles(v []VacancyProfessionalRoleItem)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *VacancyEditCommon) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetResponseLetterRequired

`func (o *VacancyEditCommon) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacancyEditCommon) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacancyEditCommon) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.

### HasResponseLetterRequired

`func (o *VacancyEditCommon) HasResponseLetterRequired() bool`

HasResponseLetterRequired returns a boolean if a field has been set.

### GetResponseNotifications

`func (o *VacancyEditCommon) GetResponseNotifications() bool`

GetResponseNotifications returns the ResponseNotifications field if non-nil, zero value otherwise.

### GetResponseNotificationsOk

`func (o *VacancyEditCommon) GetResponseNotificationsOk() (*bool, bool)`

GetResponseNotificationsOk returns a tuple with the ResponseNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseNotifications

`func (o *VacancyEditCommon) SetResponseNotifications(v bool)`

SetResponseNotifications sets ResponseNotifications field to given value.

### HasResponseNotifications

`func (o *VacancyEditCommon) HasResponseNotifications() bool`

HasResponseNotifications returns a boolean if a field has been set.

### GetResponseUrl

`func (o *VacancyEditCommon) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacancyEditCommon) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacancyEditCommon) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacancyEditCommon) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacancyEditCommon) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacancyEditCommon) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacancyEditCommon) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacancyEditCommon) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacancyEditCommon) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacancyEditCommon) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacancyEditCommon) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacancyEditCommon) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacancyEditCommon) GetSchedule() VacancySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacancyEditCommon) GetScheduleOk() (*VacancySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacancyEditCommon) SetSchedule(v VacancySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *VacancyEditCommon) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *VacancyEditCommon) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *VacancyEditCommon) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetTest

`func (o *VacancyEditCommon) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacancyEditCommon) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacancyEditCommon) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacancyEditCommon) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacancyEditCommon) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacancyEditCommon) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetWorkingDays

`func (o *VacancyEditCommon) GetWorkingDays() []VacancyWorkingDayItem`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacancyEditCommon) GetWorkingDaysOk() (*[]VacancyWorkingDayItem, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacancyEditCommon) SetWorkingDays(v []VacancyWorkingDayItem)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacancyEditCommon) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacancyEditCommon) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacancyEditCommon) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacancyEditCommon) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItem`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacancyEditCommon) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItem, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacancyEditCommon) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItem)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacancyEditCommon) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacancyEditCommon) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacancyEditCommon) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacancyEditCommon) GetWorkingTimeModes() []VacancyWorkingTimeModeItem`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacancyEditCommon) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItem, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacancyEditCommon) SetWorkingTimeModes(v []VacancyWorkingTimeModeItem)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacancyEditCommon) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacancyEditCommon) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacancyEditCommon) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil
### GetDescription

`func (o *VacancyEditCommon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacancyEditCommon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacancyEditCommon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VacancyEditCommon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *VacancyEditCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyEditCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyEditCommon) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacancyEditCommon) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


