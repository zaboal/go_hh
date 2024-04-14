# VacancyCommonFields

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
**DriverLicenseTypes** | Pointer to **[]string** |  | [optional] 
**Employment** | Pointer to [**NullableVacancyEmployment**](VacancyEmployment.md) |  | [optional] 
**Experience** | Pointer to [**NullableVacancyExperience**](VacancyExperience.md) |  | [optional] 
**KeySkills** | Pointer to **[]string** |  | [optional] 
**Languages** | Pointer to **[]string** |  | [optional] 
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

## Methods

### NewVacancyCommonFields

`func NewVacancyCommonFields() *VacancyCommonFields`

NewVacancyCommonFields instantiates a new VacancyCommonFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyCommonFieldsWithDefaults

`func NewVacancyCommonFieldsWithDefaults() *VacancyCommonFields`

NewVacancyCommonFieldsWithDefaults instantiates a new VacancyCommonFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacancyCommonFields) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacancyCommonFields) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacancyCommonFields) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.

### HasAcceptHandicapped

`func (o *VacancyCommonFields) HasAcceptHandicapped() bool`

HasAcceptHandicapped returns a boolean if a field has been set.

### GetAcceptIncompleteResumes

`func (o *VacancyCommonFields) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacancyCommonFields) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacancyCommonFields) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.

### HasAcceptIncompleteResumes

`func (o *VacancyCommonFields) HasAcceptIncompleteResumes() bool`

HasAcceptIncompleteResumes returns a boolean if a field has been set.

### GetAcceptKids

`func (o *VacancyCommonFields) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacancyCommonFields) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacancyCommonFields) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.

### HasAcceptKids

`func (o *VacancyCommonFields) HasAcceptKids() bool`

HasAcceptKids returns a boolean if a field has been set.

### GetAcceptTemporary

`func (o *VacancyCommonFields) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacancyCommonFields) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacancyCommonFields) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacancyCommonFields) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacancyCommonFields) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacancyCommonFields) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAddress

`func (o *VacancyCommonFields) GetAddress() VacancyAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacancyCommonFields) GetAddressOk() (*VacancyAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacancyCommonFields) SetAddress(v VacancyAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacancyCommonFields) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacancyCommonFields) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacancyCommonFields) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAllowMessages

`func (o *VacancyCommonFields) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacancyCommonFields) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacancyCommonFields) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.

### HasAllowMessages

`func (o *VacancyCommonFields) HasAllowMessages() bool`

HasAllowMessages returns a boolean if a field has been set.

### GetBrandedTemplate

`func (o *VacancyCommonFields) GetBrandedTemplate() VacancyBrandedTemplate`

GetBrandedTemplate returns the BrandedTemplate field if non-nil, zero value otherwise.

### GetBrandedTemplateOk

`func (o *VacancyCommonFields) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool)`

GetBrandedTemplateOk returns a tuple with the BrandedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedTemplate

`func (o *VacancyCommonFields) SetBrandedTemplate(v VacancyBrandedTemplate)`

SetBrandedTemplate sets BrandedTemplate field to given value.

### HasBrandedTemplate

`func (o *VacancyCommonFields) HasBrandedTemplate() bool`

HasBrandedTemplate returns a boolean if a field has been set.

### SetBrandedTemplateNil

`func (o *VacancyCommonFields) SetBrandedTemplateNil(b bool)`

 SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil

### UnsetBrandedTemplate
`func (o *VacancyCommonFields) UnsetBrandedTemplate()`

UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
### GetCode

`func (o *VacancyCommonFields) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacancyCommonFields) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacancyCommonFields) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacancyCommonFields) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacancyCommonFields) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacancyCommonFields) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetContacts

`func (o *VacancyCommonFields) GetContacts() VacancyContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacancyCommonFields) GetContactsOk() (*VacancyContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacancyCommonFields) SetContacts(v VacancyContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacancyCommonFields) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacancyCommonFields) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacancyCommonFields) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCustomEmployerName

`func (o *VacancyCommonFields) GetCustomEmployerName() string`

GetCustomEmployerName returns the CustomEmployerName field if non-nil, zero value otherwise.

### GetCustomEmployerNameOk

`func (o *VacancyCommonFields) GetCustomEmployerNameOk() (*string, bool)`

GetCustomEmployerNameOk returns a tuple with the CustomEmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmployerName

`func (o *VacancyCommonFields) SetCustomEmployerName(v string)`

SetCustomEmployerName sets CustomEmployerName field to given value.

### HasCustomEmployerName

`func (o *VacancyCommonFields) HasCustomEmployerName() bool`

HasCustomEmployerName returns a boolean if a field has been set.

### SetCustomEmployerNameNil

`func (o *VacancyCommonFields) SetCustomEmployerNameNil(b bool)`

 SetCustomEmployerNameNil sets the value for CustomEmployerName to be an explicit nil

### UnsetCustomEmployerName
`func (o *VacancyCommonFields) UnsetCustomEmployerName()`

UnsetCustomEmployerName ensures that no value is present for CustomEmployerName, not even an explicit nil
### GetDepartment

`func (o *VacancyCommonFields) GetDepartment() VacancyDepartment`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacancyCommonFields) GetDepartmentOk() (*VacancyDepartment, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacancyCommonFields) SetDepartment(v VacancyDepartment)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacancyCommonFields) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacancyCommonFields) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacancyCommonFields) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDriverLicenseTypes

`func (o *VacancyCommonFields) GetDriverLicenseTypes() []string`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacancyCommonFields) GetDriverLicenseTypesOk() (*[]string, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacancyCommonFields) SetDriverLicenseTypes(v []string)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.

### HasDriverLicenseTypes

`func (o *VacancyCommonFields) HasDriverLicenseTypes() bool`

HasDriverLicenseTypes returns a boolean if a field has been set.

### SetDriverLicenseTypesNil

`func (o *VacancyCommonFields) SetDriverLicenseTypesNil(b bool)`

 SetDriverLicenseTypesNil sets the value for DriverLicenseTypes to be an explicit nil

### UnsetDriverLicenseTypes
`func (o *VacancyCommonFields) UnsetDriverLicenseTypes()`

UnsetDriverLicenseTypes ensures that no value is present for DriverLicenseTypes, not even an explicit nil
### GetEmployment

`func (o *VacancyCommonFields) GetEmployment() VacancyEmployment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacancyCommonFields) GetEmploymentOk() (*VacancyEmployment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacancyCommonFields) SetEmployment(v VacancyEmployment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacancyCommonFields) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacancyCommonFields) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacancyCommonFields) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacancyCommonFields) GetExperience() VacancyExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacancyCommonFields) GetExperienceOk() (*VacancyExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacancyCommonFields) SetExperience(v VacancyExperience)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *VacancyCommonFields) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *VacancyCommonFields) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacancyCommonFields) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetKeySkills

`func (o *VacancyCommonFields) GetKeySkills() []string`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacancyCommonFields) GetKeySkillsOk() (*[]string, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacancyCommonFields) SetKeySkills(v []string)`

SetKeySkills sets KeySkills field to given value.

### HasKeySkills

`func (o *VacancyCommonFields) HasKeySkills() bool`

HasKeySkills returns a boolean if a field has been set.

### SetKeySkillsNil

`func (o *VacancyCommonFields) SetKeySkillsNil(b bool)`

 SetKeySkillsNil sets the value for KeySkills to be an explicit nil

### UnsetKeySkills
`func (o *VacancyCommonFields) UnsetKeySkills()`

UnsetKeySkills ensures that no value is present for KeySkills, not even an explicit nil
### GetLanguages

`func (o *VacancyCommonFields) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacancyCommonFields) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacancyCommonFields) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VacancyCommonFields) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *VacancyCommonFields) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *VacancyCommonFields) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetProfessionalRoles

`func (o *VacancyCommonFields) GetProfessionalRoles() []VacancyProfessionalRoleItem`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacancyCommonFields) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItem, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacancyCommonFields) SetProfessionalRoles(v []VacancyProfessionalRoleItem)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *VacancyCommonFields) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetResponseLetterRequired

`func (o *VacancyCommonFields) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacancyCommonFields) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacancyCommonFields) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.

### HasResponseLetterRequired

`func (o *VacancyCommonFields) HasResponseLetterRequired() bool`

HasResponseLetterRequired returns a boolean if a field has been set.

### GetResponseNotifications

`func (o *VacancyCommonFields) GetResponseNotifications() bool`

GetResponseNotifications returns the ResponseNotifications field if non-nil, zero value otherwise.

### GetResponseNotificationsOk

`func (o *VacancyCommonFields) GetResponseNotificationsOk() (*bool, bool)`

GetResponseNotificationsOk returns a tuple with the ResponseNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseNotifications

`func (o *VacancyCommonFields) SetResponseNotifications(v bool)`

SetResponseNotifications sets ResponseNotifications field to given value.

### HasResponseNotifications

`func (o *VacancyCommonFields) HasResponseNotifications() bool`

HasResponseNotifications returns a boolean if a field has been set.

### GetResponseUrl

`func (o *VacancyCommonFields) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacancyCommonFields) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacancyCommonFields) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacancyCommonFields) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacancyCommonFields) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacancyCommonFields) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacancyCommonFields) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacancyCommonFields) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacancyCommonFields) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacancyCommonFields) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacancyCommonFields) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacancyCommonFields) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacancyCommonFields) GetSchedule() VacancySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacancyCommonFields) GetScheduleOk() (*VacancySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacancyCommonFields) SetSchedule(v VacancySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *VacancyCommonFields) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *VacancyCommonFields) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *VacancyCommonFields) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetTest

`func (o *VacancyCommonFields) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacancyCommonFields) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacancyCommonFields) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacancyCommonFields) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacancyCommonFields) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacancyCommonFields) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetWorkingDays

`func (o *VacancyCommonFields) GetWorkingDays() []VacancyWorkingDayItem`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacancyCommonFields) GetWorkingDaysOk() (*[]VacancyWorkingDayItem, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacancyCommonFields) SetWorkingDays(v []VacancyWorkingDayItem)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacancyCommonFields) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacancyCommonFields) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacancyCommonFields) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacancyCommonFields) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItem`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacancyCommonFields) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItem, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacancyCommonFields) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItem)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacancyCommonFields) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacancyCommonFields) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacancyCommonFields) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacancyCommonFields) GetWorkingTimeModes() []VacancyWorkingTimeModeItem`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacancyCommonFields) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItem, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacancyCommonFields) SetWorkingTimeModes(v []VacancyWorkingTimeModeItem)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacancyCommonFields) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacancyCommonFields) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacancyCommonFields) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


