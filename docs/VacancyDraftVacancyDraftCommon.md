# VacancyDraftVacancyDraftCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptHandicapped** | **bool** | Указание, что вакансия доступна для соискателей с инвалидностью | 
**AcceptIncompleteResumes** | **bool** | Разрешен ли отклик на вакансию неполным резюме | 
**AcceptKids** | **bool** | Указание, что вакансия доступна для соискателей старше 14 лет [подробнее](https://github.com/hhru/api/blob/master/docs/employer_vacancies_accept_kids.md#accept-kids) | 
**AcceptTemporary** | Pointer to **NullableBool** | Указание, что вакансия доступна с временным трудоустройством | [optional] 
**AllowMessages** | **bool** | Возможность [переписки с кандидатами](https://inboxemp.hh.ru/) по данной вакансии | 
**BillingType** | [**NullableVacancyBillingTypeOutput**](VacancyBillingTypeOutput.md) |  | 
**Code** | Pointer to **NullableString** | Внутренний код вакансии | [optional] 
**Department** | Pointer to [**NullableVacancyDepartmentOutput**](VacancyDepartmentOutput.md) |  | [optional] 
**Description** | **string** | Описание в html, не менее 200 символов | 
**DriverLicenseTypes** | [**[]VacancyDriverLicenceTypeItem**](VacancyDriverLicenceTypeItem.md) | Список требуемых категорий водительских прав | 
**Employment** | Pointer to [**NullableVacancyEmploymentOutput**](VacancyEmploymentOutput.md) |  | [optional] 
**Experience** | [**NullableVacancyExperienceOutput**](VacancyExperienceOutput.md) |  | 
**HasTest** | **bool** | Информация о наличии прикрепленного тестового задании к вакансии | 
**KeySkills** | [**[]VacancyKeySkillItem**](VacancyKeySkillItem.md) | Список ключевых навыков, не более 30 | 
**Languages** | [**[]VacancyLanguageOutput**](VacancyLanguageOutput.md) | Список языков вакансии. Значения из справочника [/languages](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Manager** | [**VacancyManager**](VacancyManager.md) |  | 
**Name** | **string** | Название | 
**ProfessionalRoles** | [**[]VacancyProfessionalRoleItemOutput**](VacancyProfessionalRoleItemOutput.md) | Список профессиональных ролей | 
**ResponseLetterRequired** | **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | 
**ResponseNotifications** | **bool** | Уведомлять ли менеджера о новых откликах | 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | Pointer to [**NullableVacancySalary**](VacancySalary.md) |  | [optional] 
**Schedule** | [**VacancyScheduleOutput**](VacancyScheduleOutput.md) |  | 
**Test** | Pointer to [**NullableVacancyDraftTest**](VacancyDraftTest.md) |  | [optional] 
**Type** | [**VacancyTypeOutput**](VacancyTypeOutput.md) |  | 
**WorkingDays** | Pointer to [**[]VacancyWorkingDayItemOutput**](VacancyWorkingDayItemOutput.md) | Список рабочих дней | [optional] 
**WorkingTimeIntervals** | Pointer to [**[]VacancyWorkingTimeIntervalItemOutput**](VacancyWorkingTimeIntervalItemOutput.md) | Список с временными интервалами работы | [optional] 
**WorkingTimeModes** | Pointer to [**[]VacancyWorkingTimeModeItemOutput**](VacancyWorkingTimeModeItemOutput.md) | Список режимов времени работы | [optional] 

## Methods

### NewVacancyDraftVacancyDraftCommon

`func NewVacancyDraftVacancyDraftCommon(acceptHandicapped bool, acceptIncompleteResumes bool, acceptKids bool, allowMessages bool, billingType NullableVacancyBillingTypeOutput, description string, driverLicenseTypes []VacancyDriverLicenceTypeItem, experience NullableVacancyExperienceOutput, hasTest bool, keySkills []VacancyKeySkillItem, languages []VacancyLanguageOutput, manager VacancyManager, name string, professionalRoles []VacancyProfessionalRoleItemOutput, responseLetterRequired bool, responseNotifications bool, schedule VacancyScheduleOutput, type_ VacancyTypeOutput, ) *VacancyDraftVacancyDraftCommon`

NewVacancyDraftVacancyDraftCommon instantiates a new VacancyDraftVacancyDraftCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftVacancyDraftCommonWithDefaults

`func NewVacancyDraftVacancyDraftCommonWithDefaults() *VacancyDraftVacancyDraftCommon`

NewVacancyDraftVacancyDraftCommonWithDefaults instantiates a new VacancyDraftVacancyDraftCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacancyDraftVacancyDraftCommon) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacancyDraftVacancyDraftCommon) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacancyDraftVacancyDraftCommon) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.


### GetAcceptIncompleteResumes

`func (o *VacancyDraftVacancyDraftCommon) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacancyDraftVacancyDraftCommon) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacancyDraftVacancyDraftCommon) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.


### GetAcceptKids

`func (o *VacancyDraftVacancyDraftCommon) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacancyDraftVacancyDraftCommon) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacancyDraftVacancyDraftCommon) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.


### GetAcceptTemporary

`func (o *VacancyDraftVacancyDraftCommon) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacancyDraftVacancyDraftCommon) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacancyDraftVacancyDraftCommon) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacancyDraftVacancyDraftCommon) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacancyDraftVacancyDraftCommon) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacancyDraftVacancyDraftCommon) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAllowMessages

`func (o *VacancyDraftVacancyDraftCommon) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacancyDraftVacancyDraftCommon) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacancyDraftVacancyDraftCommon) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.


### GetBillingType

`func (o *VacancyDraftVacancyDraftCommon) GetBillingType() VacancyBillingTypeOutput`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacancyDraftVacancyDraftCommon) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacancyDraftVacancyDraftCommon) SetBillingType(v VacancyBillingTypeOutput)`

SetBillingType sets BillingType field to given value.


### SetBillingTypeNil

`func (o *VacancyDraftVacancyDraftCommon) SetBillingTypeNil(b bool)`

 SetBillingTypeNil sets the value for BillingType to be an explicit nil

### UnsetBillingType
`func (o *VacancyDraftVacancyDraftCommon) UnsetBillingType()`

UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
### GetCode

`func (o *VacancyDraftVacancyDraftCommon) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacancyDraftVacancyDraftCommon) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacancyDraftVacancyDraftCommon) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacancyDraftVacancyDraftCommon) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacancyDraftVacancyDraftCommon) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacancyDraftVacancyDraftCommon) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetDepartment

`func (o *VacancyDraftVacancyDraftCommon) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacancyDraftVacancyDraftCommon) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacancyDraftVacancyDraftCommon) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacancyDraftVacancyDraftCommon) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacancyDraftVacancyDraftCommon) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacancyDraftVacancyDraftCommon) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDescription

`func (o *VacancyDraftVacancyDraftCommon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacancyDraftVacancyDraftCommon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacancyDraftVacancyDraftCommon) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDriverLicenseTypes

`func (o *VacancyDraftVacancyDraftCommon) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacancyDraftVacancyDraftCommon) GetDriverLicenseTypesOk() (*[]VacancyDriverLicenceTypeItem, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacancyDraftVacancyDraftCommon) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEmployment

`func (o *VacancyDraftVacancyDraftCommon) GetEmployment() VacancyEmploymentOutput`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacancyDraftVacancyDraftCommon) GetEmploymentOk() (*VacancyEmploymentOutput, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacancyDraftVacancyDraftCommon) SetEmployment(v VacancyEmploymentOutput)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacancyDraftVacancyDraftCommon) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacancyDraftVacancyDraftCommon) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacancyDraftVacancyDraftCommon) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacancyDraftVacancyDraftCommon) GetExperience() VacancyExperienceOutput`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacancyDraftVacancyDraftCommon) GetExperienceOk() (*VacancyExperienceOutput, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacancyDraftVacancyDraftCommon) SetExperience(v VacancyExperienceOutput)`

SetExperience sets Experience field to given value.


### SetExperienceNil

`func (o *VacancyDraftVacancyDraftCommon) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacancyDraftVacancyDraftCommon) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetHasTest

`func (o *VacancyDraftVacancyDraftCommon) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacancyDraftVacancyDraftCommon) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacancyDraftVacancyDraftCommon) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetKeySkills

`func (o *VacancyDraftVacancyDraftCommon) GetKeySkills() []VacancyKeySkillItem`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacancyDraftVacancyDraftCommon) GetKeySkillsOk() (*[]VacancyKeySkillItem, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacancyDraftVacancyDraftCommon) SetKeySkills(v []VacancyKeySkillItem)`

SetKeySkills sets KeySkills field to given value.


### GetLanguages

`func (o *VacancyDraftVacancyDraftCommon) GetLanguages() []VacancyLanguageOutput`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacancyDraftVacancyDraftCommon) GetLanguagesOk() (*[]VacancyLanguageOutput, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacancyDraftVacancyDraftCommon) SetLanguages(v []VacancyLanguageOutput)`

SetLanguages sets Languages field to given value.


### GetManager

`func (o *VacancyDraftVacancyDraftCommon) GetManager() VacancyManager`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacancyDraftVacancyDraftCommon) GetManagerOk() (*VacancyManager, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacancyDraftVacancyDraftCommon) SetManager(v VacancyManager)`

SetManager sets Manager field to given value.


### GetName

`func (o *VacancyDraftVacancyDraftCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyDraftVacancyDraftCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyDraftVacancyDraftCommon) SetName(v string)`

SetName sets Name field to given value.


### GetProfessionalRoles

`func (o *VacancyDraftVacancyDraftCommon) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacancyDraftVacancyDraftCommon) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItemOutput, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacancyDraftVacancyDraftCommon) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput)`

SetProfessionalRoles sets ProfessionalRoles field to given value.


### GetResponseLetterRequired

`func (o *VacancyDraftVacancyDraftCommon) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacancyDraftVacancyDraftCommon) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacancyDraftVacancyDraftCommon) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseNotifications

`func (o *VacancyDraftVacancyDraftCommon) GetResponseNotifications() bool`

GetResponseNotifications returns the ResponseNotifications field if non-nil, zero value otherwise.

### GetResponseNotificationsOk

`func (o *VacancyDraftVacancyDraftCommon) GetResponseNotificationsOk() (*bool, bool)`

GetResponseNotificationsOk returns a tuple with the ResponseNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseNotifications

`func (o *VacancyDraftVacancyDraftCommon) SetResponseNotifications(v bool)`

SetResponseNotifications sets ResponseNotifications field to given value.


### GetResponseUrl

`func (o *VacancyDraftVacancyDraftCommon) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacancyDraftVacancyDraftCommon) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacancyDraftVacancyDraftCommon) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacancyDraftVacancyDraftCommon) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacancyDraftVacancyDraftCommon) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacancyDraftVacancyDraftCommon) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacancyDraftVacancyDraftCommon) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacancyDraftVacancyDraftCommon) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacancyDraftVacancyDraftCommon) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacancyDraftVacancyDraftCommon) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacancyDraftVacancyDraftCommon) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacancyDraftVacancyDraftCommon) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacancyDraftVacancyDraftCommon) GetSchedule() VacancyScheduleOutput`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacancyDraftVacancyDraftCommon) GetScheduleOk() (*VacancyScheduleOutput, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacancyDraftVacancyDraftCommon) SetSchedule(v VacancyScheduleOutput)`

SetSchedule sets Schedule field to given value.


### GetTest

`func (o *VacancyDraftVacancyDraftCommon) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacancyDraftVacancyDraftCommon) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacancyDraftVacancyDraftCommon) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacancyDraftVacancyDraftCommon) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacancyDraftVacancyDraftCommon) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacancyDraftVacancyDraftCommon) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetType

`func (o *VacancyDraftVacancyDraftCommon) GetType() VacancyTypeOutput`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacancyDraftVacancyDraftCommon) GetTypeOk() (*VacancyTypeOutput, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacancyDraftVacancyDraftCommon) SetType(v VacancyTypeOutput)`

SetType sets Type field to given value.


### GetWorkingDays

`func (o *VacancyDraftVacancyDraftCommon) GetWorkingDays() []VacancyWorkingDayItemOutput`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacancyDraftVacancyDraftCommon) GetWorkingDaysOk() (*[]VacancyWorkingDayItemOutput, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacancyDraftVacancyDraftCommon) SetWorkingDays(v []VacancyWorkingDayItemOutput)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacancyDraftVacancyDraftCommon) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacancyDraftVacancyDraftCommon) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacancyDraftVacancyDraftCommon) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacancyDraftVacancyDraftCommon) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItemOutput`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacancyDraftVacancyDraftCommon) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItemOutput, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacancyDraftVacancyDraftCommon) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItemOutput)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacancyDraftVacancyDraftCommon) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacancyDraftVacancyDraftCommon) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacancyDraftVacancyDraftCommon) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacancyDraftVacancyDraftCommon) GetWorkingTimeModes() []VacancyWorkingTimeModeItemOutput`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacancyDraftVacancyDraftCommon) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItemOutput, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacancyDraftVacancyDraftCommon) SetWorkingTimeModes(v []VacancyWorkingTimeModeItemOutput)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacancyDraftVacancyDraftCommon) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacancyDraftVacancyDraftCommon) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacancyDraftVacancyDraftCommon) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


