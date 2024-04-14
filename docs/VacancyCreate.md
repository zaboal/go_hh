# VacancyCreate

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
**ProfessionalRoles** | [**[]VacancyProfessionalRoleItem**](VacancyProfessionalRoleItem.md) | Список профессиональных ролей | 
**ResponseLetterRequired** | Pointer to **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | [optional] 
**ResponseNotifications** | Pointer to **bool** | Уведомлять ли менеджера о новых откликах | [optional] 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | Pointer to [**NullableVacancySalary**](VacancySalary.md) |  | [optional] 
**Schedule** | Pointer to [**NullableVacancySchedule**](VacancySchedule.md) |  | [optional] 
**Test** | Pointer to [**NullableVacancyDraftTest**](VacancyDraftTest.md) |  | [optional] 
**WorkingDays** | Pointer to [**[]VacancyWorkingDayItem**](VacancyWorkingDayItem.md) | Список рабочих дней | [optional] 
**WorkingTimeIntervals** | Pointer to [**[]VacancyWorkingTimeIntervalItem**](VacancyWorkingTimeIntervalItem.md) | Список с временными интервалами работы | [optional] 
**WorkingTimeModes** | Pointer to [**[]VacancyWorkingTimeModeItem**](VacancyWorkingTimeModeItem.md) | Список режимов времени работы | [optional] 
**Area** | [**VacancyArea**](VacancyArea.md) |  | 
**BillingType** | [**VacancyBillingType**](VacancyBillingType.md) |  | 
**Description** | **string** | Описание в html, не менее 200 символов | 
**Manager** | Pointer to [**VacancyManager**](VacancyManager.md) |  | [optional] 
**Name** | **string** | Название | 
**PreviousId** | Pointer to **NullableString** | Если этот параметр передан, то у новой вакансии дополнительно будет создана связь с предыдущей вакансией (поле previous_id). Этот параметр не влияет на другие и не связан с ними, их всё равно необходимо передавать.  Должен быть равен только ID архивной вакансии. ID архивной вакансии можно получить, запросив [список архивных вакансий](#tag/Upravlenie-vakansiyami/operation/get-archived-vacancies) &lt;a name&#x3D;&#39;previous_id&#39;&gt;&lt;/a&gt;  | [optional] 
**Type** | [**VacancyType**](VacancyType.md) |  | 

## Methods

### NewVacancyCreate

`func NewVacancyCreate(professionalRoles []VacancyProfessionalRoleItem, area VacancyArea, billingType VacancyBillingType, description string, name string, type_ VacancyType, ) *VacancyCreate`

NewVacancyCreate instantiates a new VacancyCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyCreateWithDefaults

`func NewVacancyCreateWithDefaults() *VacancyCreate`

NewVacancyCreateWithDefaults instantiates a new VacancyCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacancyCreate) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacancyCreate) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacancyCreate) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.

### HasAcceptHandicapped

`func (o *VacancyCreate) HasAcceptHandicapped() bool`

HasAcceptHandicapped returns a boolean if a field has been set.

### GetAcceptIncompleteResumes

`func (o *VacancyCreate) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacancyCreate) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacancyCreate) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.

### HasAcceptIncompleteResumes

`func (o *VacancyCreate) HasAcceptIncompleteResumes() bool`

HasAcceptIncompleteResumes returns a boolean if a field has been set.

### GetAcceptKids

`func (o *VacancyCreate) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacancyCreate) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacancyCreate) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.

### HasAcceptKids

`func (o *VacancyCreate) HasAcceptKids() bool`

HasAcceptKids returns a boolean if a field has been set.

### GetAcceptTemporary

`func (o *VacancyCreate) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacancyCreate) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacancyCreate) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacancyCreate) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacancyCreate) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacancyCreate) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAddress

`func (o *VacancyCreate) GetAddress() VacancyAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacancyCreate) GetAddressOk() (*VacancyAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacancyCreate) SetAddress(v VacancyAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacancyCreate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacancyCreate) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacancyCreate) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAllowMessages

`func (o *VacancyCreate) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacancyCreate) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacancyCreate) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.

### HasAllowMessages

`func (o *VacancyCreate) HasAllowMessages() bool`

HasAllowMessages returns a boolean if a field has been set.

### GetBrandedTemplate

`func (o *VacancyCreate) GetBrandedTemplate() VacancyBrandedTemplate`

GetBrandedTemplate returns the BrandedTemplate field if non-nil, zero value otherwise.

### GetBrandedTemplateOk

`func (o *VacancyCreate) GetBrandedTemplateOk() (*VacancyBrandedTemplate, bool)`

GetBrandedTemplateOk returns a tuple with the BrandedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedTemplate

`func (o *VacancyCreate) SetBrandedTemplate(v VacancyBrandedTemplate)`

SetBrandedTemplate sets BrandedTemplate field to given value.

### HasBrandedTemplate

`func (o *VacancyCreate) HasBrandedTemplate() bool`

HasBrandedTemplate returns a boolean if a field has been set.

### SetBrandedTemplateNil

`func (o *VacancyCreate) SetBrandedTemplateNil(b bool)`

 SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil

### UnsetBrandedTemplate
`func (o *VacancyCreate) UnsetBrandedTemplate()`

UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
### GetCode

`func (o *VacancyCreate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacancyCreate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacancyCreate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacancyCreate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacancyCreate) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacancyCreate) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetContacts

`func (o *VacancyCreate) GetContacts() VacancyContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacancyCreate) GetContactsOk() (*VacancyContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacancyCreate) SetContacts(v VacancyContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacancyCreate) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacancyCreate) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacancyCreate) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCustomEmployerName

`func (o *VacancyCreate) GetCustomEmployerName() string`

GetCustomEmployerName returns the CustomEmployerName field if non-nil, zero value otherwise.

### GetCustomEmployerNameOk

`func (o *VacancyCreate) GetCustomEmployerNameOk() (*string, bool)`

GetCustomEmployerNameOk returns a tuple with the CustomEmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmployerName

`func (o *VacancyCreate) SetCustomEmployerName(v string)`

SetCustomEmployerName sets CustomEmployerName field to given value.

### HasCustomEmployerName

`func (o *VacancyCreate) HasCustomEmployerName() bool`

HasCustomEmployerName returns a boolean if a field has been set.

### SetCustomEmployerNameNil

`func (o *VacancyCreate) SetCustomEmployerNameNil(b bool)`

 SetCustomEmployerNameNil sets the value for CustomEmployerName to be an explicit nil

### UnsetCustomEmployerName
`func (o *VacancyCreate) UnsetCustomEmployerName()`

UnsetCustomEmployerName ensures that no value is present for CustomEmployerName, not even an explicit nil
### GetDepartment

`func (o *VacancyCreate) GetDepartment() VacancyDepartment`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacancyCreate) GetDepartmentOk() (*VacancyDepartment, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacancyCreate) SetDepartment(v VacancyDepartment)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacancyCreate) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacancyCreate) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacancyCreate) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDriverLicenseTypes

`func (o *VacancyCreate) GetDriverLicenseTypes() []string`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacancyCreate) GetDriverLicenseTypesOk() (*[]string, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacancyCreate) SetDriverLicenseTypes(v []string)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.

### HasDriverLicenseTypes

`func (o *VacancyCreate) HasDriverLicenseTypes() bool`

HasDriverLicenseTypes returns a boolean if a field has been set.

### SetDriverLicenseTypesNil

`func (o *VacancyCreate) SetDriverLicenseTypesNil(b bool)`

 SetDriverLicenseTypesNil sets the value for DriverLicenseTypes to be an explicit nil

### UnsetDriverLicenseTypes
`func (o *VacancyCreate) UnsetDriverLicenseTypes()`

UnsetDriverLicenseTypes ensures that no value is present for DriverLicenseTypes, not even an explicit nil
### GetEmployment

`func (o *VacancyCreate) GetEmployment() VacancyEmployment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacancyCreate) GetEmploymentOk() (*VacancyEmployment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacancyCreate) SetEmployment(v VacancyEmployment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacancyCreate) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacancyCreate) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacancyCreate) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacancyCreate) GetExperience() VacancyExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacancyCreate) GetExperienceOk() (*VacancyExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacancyCreate) SetExperience(v VacancyExperience)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *VacancyCreate) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *VacancyCreate) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacancyCreate) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetKeySkills

`func (o *VacancyCreate) GetKeySkills() []string`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacancyCreate) GetKeySkillsOk() (*[]string, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacancyCreate) SetKeySkills(v []string)`

SetKeySkills sets KeySkills field to given value.

### HasKeySkills

`func (o *VacancyCreate) HasKeySkills() bool`

HasKeySkills returns a boolean if a field has been set.

### SetKeySkillsNil

`func (o *VacancyCreate) SetKeySkillsNil(b bool)`

 SetKeySkillsNil sets the value for KeySkills to be an explicit nil

### UnsetKeySkills
`func (o *VacancyCreate) UnsetKeySkills()`

UnsetKeySkills ensures that no value is present for KeySkills, not even an explicit nil
### GetLanguages

`func (o *VacancyCreate) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacancyCreate) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacancyCreate) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VacancyCreate) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### SetLanguagesNil

`func (o *VacancyCreate) SetLanguagesNil(b bool)`

 SetLanguagesNil sets the value for Languages to be an explicit nil

### UnsetLanguages
`func (o *VacancyCreate) UnsetLanguages()`

UnsetLanguages ensures that no value is present for Languages, not even an explicit nil
### GetProfessionalRoles

`func (o *VacancyCreate) GetProfessionalRoles() []VacancyProfessionalRoleItem`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacancyCreate) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItem, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacancyCreate) SetProfessionalRoles(v []VacancyProfessionalRoleItem)`

SetProfessionalRoles sets ProfessionalRoles field to given value.


### GetResponseLetterRequired

`func (o *VacancyCreate) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacancyCreate) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacancyCreate) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.

### HasResponseLetterRequired

`func (o *VacancyCreate) HasResponseLetterRequired() bool`

HasResponseLetterRequired returns a boolean if a field has been set.

### GetResponseNotifications

`func (o *VacancyCreate) GetResponseNotifications() bool`

GetResponseNotifications returns the ResponseNotifications field if non-nil, zero value otherwise.

### GetResponseNotificationsOk

`func (o *VacancyCreate) GetResponseNotificationsOk() (*bool, bool)`

GetResponseNotificationsOk returns a tuple with the ResponseNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseNotifications

`func (o *VacancyCreate) SetResponseNotifications(v bool)`

SetResponseNotifications sets ResponseNotifications field to given value.

### HasResponseNotifications

`func (o *VacancyCreate) HasResponseNotifications() bool`

HasResponseNotifications returns a boolean if a field has been set.

### GetResponseUrl

`func (o *VacancyCreate) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacancyCreate) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacancyCreate) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacancyCreate) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacancyCreate) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacancyCreate) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacancyCreate) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacancyCreate) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacancyCreate) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacancyCreate) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacancyCreate) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacancyCreate) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacancyCreate) GetSchedule() VacancySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacancyCreate) GetScheduleOk() (*VacancySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacancyCreate) SetSchedule(v VacancySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *VacancyCreate) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *VacancyCreate) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *VacancyCreate) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetTest

`func (o *VacancyCreate) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacancyCreate) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacancyCreate) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacancyCreate) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacancyCreate) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacancyCreate) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetWorkingDays

`func (o *VacancyCreate) GetWorkingDays() []VacancyWorkingDayItem`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacancyCreate) GetWorkingDaysOk() (*[]VacancyWorkingDayItem, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacancyCreate) SetWorkingDays(v []VacancyWorkingDayItem)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacancyCreate) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacancyCreate) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacancyCreate) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacancyCreate) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItem`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacancyCreate) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItem, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacancyCreate) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItem)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacancyCreate) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacancyCreate) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacancyCreate) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacancyCreate) GetWorkingTimeModes() []VacancyWorkingTimeModeItem`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacancyCreate) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItem, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacancyCreate) SetWorkingTimeModes(v []VacancyWorkingTimeModeItem)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacancyCreate) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacancyCreate) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacancyCreate) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil
### GetArea

`func (o *VacancyCreate) GetArea() VacancyArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacancyCreate) GetAreaOk() (*VacancyArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacancyCreate) SetArea(v VacancyArea)`

SetArea sets Area field to given value.


### GetBillingType

`func (o *VacancyCreate) GetBillingType() VacancyBillingType`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacancyCreate) GetBillingTypeOk() (*VacancyBillingType, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacancyCreate) SetBillingType(v VacancyBillingType)`

SetBillingType sets BillingType field to given value.


### GetDescription

`func (o *VacancyCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacancyCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacancyCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetManager

`func (o *VacancyCreate) GetManager() VacancyManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacancyCreate) GetManagerOk() (*VacancyManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacancyCreate) SetManager(v VacancyManager)`

SetManager sets Manager field to given value.

### HasManager

`func (o *VacancyCreate) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetName

`func (o *VacancyCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPreviousId

`func (o *VacancyCreate) GetPreviousId() string`

GetPreviousId returns the PreviousId field if non-nil, zero value otherwise.

### GetPreviousIdOk

`func (o *VacancyCreate) GetPreviousIdOk() (*string, bool)`

GetPreviousIdOk returns a tuple with the PreviousId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousId

`func (o *VacancyCreate) SetPreviousId(v string)`

SetPreviousId sets PreviousId field to given value.

### HasPreviousId

`func (o *VacancyCreate) HasPreviousId() bool`

HasPreviousId returns a boolean if a field has been set.

### SetPreviousIdNil

`func (o *VacancyCreate) SetPreviousIdNil(b bool)`

 SetPreviousIdNil sets the value for PreviousId to be an explicit nil

### UnsetPreviousId
`func (o *VacancyCreate) UnsetPreviousId()`

UnsetPreviousId ensures that no value is present for PreviousId, not even an explicit nil
### GetType

`func (o *VacancyCreate) GetType() VacancyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacancyCreate) GetTypeOk() (*VacancyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacancyCreate) SetType(v VacancyType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


