# VacancyDraftVacancyDraftCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptHandicapped** | Pointer to **bool** | Указание, что вакансия доступна для соискателей с инвалидностью | [optional] 
**AcceptIncompleteResumes** | Pointer to **bool** | Разрешен ли отклик на вакансию неполным резюме | [optional] 
**AcceptKids** | Pointer to **bool** | Указание, что вакансия доступна для соискателей старше 14 лет [подробнее](https://github.com/hhru/api/blob/master/docs/employer_vacancies_accept_kids.md#accept-kids) | [optional] 
**AcceptTemporary** | Pointer to **NullableBool** | Указание, что вакансия доступна с временным трудоустройством | [optional] 
**Address** | Pointer to [**NullableVacancyDraftAddress**](VacancyDraftAddress.md) |  | [optional] 
**AllowMessages** | Pointer to **bool** | Возможность [переписки с кандидатами](https://inboxemp.hh.ru/) по данной вакансии | [optional] 
**Area** | Pointer to [**NullableVacancyArea**](VacancyArea.md) |  | [optional] 
**Areas** | Pointer to [**[]VacancyArea**](VacancyArea.md) | Список регионов | [optional] 
**AssignedManagerId** | Pointer to **NullableString** | Идентификатор рабочего аккаунта менеджера, которому перейдет вакансия после публикации | [optional] 
**BillingType** | Pointer to [**NullableVacancyDraftBillingType**](VacancyDraftBillingType.md) |  | [optional] 
**BrandedTemplate** | Pointer to [**NullableVacancyDraftBrandedTemplate**](VacancyDraftBrandedTemplate.md) |  | [optional] 
**Code** | Pointer to **NullableString** | Внутренний код вакансии | [optional] 
**Contacts** | Pointer to [**NullableVacancyDraftContacts**](VacancyDraftContacts.md) |  | [optional] 
**CustomEmployerName** | Pointer to **NullableString** | Название компании для анонимных вакансий (&#x60;type.id&#x3D;anonymous&#x60;), например \&quot;крупный российский банк\&quot;. Поле становится обязательным при публикации вакансии **анонимного** типа | [optional] 
**Department** | Pointer to [**NullableVacancyDepartment**](VacancyDepartment.md) |  | [optional] 
**Description** | Pointer to **NullableString** | Описание в html, минимум 1 символ, для успешной публикации вакансии не менее 200 символов | [optional] 
**DriverLicenseTypes** | Pointer to [**[]VacancyDriverLicenceTypeItem**](VacancyDriverLicenceTypeItem.md) | Список требуемых категорий водительских прав | [optional] 
**Employment** | Pointer to [**NullableVacancyDraftEmployment**](VacancyDraftEmployment.md) |  | [optional] 
**Experience** | Pointer to [**NullableVacancyExperience**](VacancyExperience.md) |  | [optional] 
**KeySkills** | Pointer to [**[]VacancyDraftKeySkillItem**](VacancyDraftKeySkillItem.md) | Список ключевых навыков, не более 30 | [optional] 
**Languages** | Pointer to [**[]VacancyLanguage**](VacancyLanguage.md) | Список языков вакансии | [optional] 
**Name** | Pointer to **NullableString** | Название | [optional] 
**ProfessionalRoles** | Pointer to [**[]VacancyDraftProfessionalRoleItem**](VacancyDraftProfessionalRoleItem.md) | Список профессиональных ролей | [optional] 
**ResponseLetterRequired** | Pointer to **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | [optional] 
**ResponseNotifications** | Pointer to **bool** | Уведомлять ли менеджера о новых откликах | [optional] 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | Pointer to [**NullableVacancySalary**](VacancySalary.md) |  | [optional] 
**Schedule** | Pointer to [**NullableVacancySchedule**](VacancySchedule.md) |  | [optional] 
**Test** | Pointer to [**NullableVacancyDraftTest**](VacancyDraftTest.md) |  | [optional] 
**Type** | Pointer to [**NullableVacancyDraftType**](VacancyDraftType.md) |  | [optional] 
**WithZp** | Pointer to **bool** | Вашу вакансию увидят больше людей. Мы разместим ее дополнительно на сервисе Зарплата.ру | [optional] 
**WorkingDays** | Pointer to [**[]VacancyWorkingDayItem**](VacancyWorkingDayItem.md) | Список рабочих дней | [optional] 
**WorkingTimeIntervals** | Pointer to [**[]VacancyWorkingTimeIntervalItem**](VacancyWorkingTimeIntervalItem.md) | Список с временными интервалами работы | [optional] 
**WorkingTimeModes** | Pointer to [**[]VacancyWorkingTimeModeItem**](VacancyWorkingTimeModeItem.md) | Список режимов времени работы | [optional] 

## Methods

### NewVacancyDraftVacancyDraftCreate

`func NewVacancyDraftVacancyDraftCreate() *VacancyDraftVacancyDraftCreate`

NewVacancyDraftVacancyDraftCreate instantiates a new VacancyDraftVacancyDraftCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftVacancyDraftCreateWithDefaults

`func NewVacancyDraftVacancyDraftCreateWithDefaults() *VacancyDraftVacancyDraftCreate`

NewVacancyDraftVacancyDraftCreateWithDefaults instantiates a new VacancyDraftVacancyDraftCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptHandicapped

`func (o *VacancyDraftVacancyDraftCreate) GetAcceptHandicapped() bool`

GetAcceptHandicapped returns the AcceptHandicapped field if non-nil, zero value otherwise.

### GetAcceptHandicappedOk

`func (o *VacancyDraftVacancyDraftCreate) GetAcceptHandicappedOk() (*bool, bool)`

GetAcceptHandicappedOk returns a tuple with the AcceptHandicapped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptHandicapped

`func (o *VacancyDraftVacancyDraftCreate) SetAcceptHandicapped(v bool)`

SetAcceptHandicapped sets AcceptHandicapped field to given value.

### HasAcceptHandicapped

`func (o *VacancyDraftVacancyDraftCreate) HasAcceptHandicapped() bool`

HasAcceptHandicapped returns a boolean if a field has been set.

### GetAcceptIncompleteResumes

`func (o *VacancyDraftVacancyDraftCreate) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacancyDraftVacancyDraftCreate) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacancyDraftVacancyDraftCreate) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.

### HasAcceptIncompleteResumes

`func (o *VacancyDraftVacancyDraftCreate) HasAcceptIncompleteResumes() bool`

HasAcceptIncompleteResumes returns a boolean if a field has been set.

### GetAcceptKids

`func (o *VacancyDraftVacancyDraftCreate) GetAcceptKids() bool`

GetAcceptKids returns the AcceptKids field if non-nil, zero value otherwise.

### GetAcceptKidsOk

`func (o *VacancyDraftVacancyDraftCreate) GetAcceptKidsOk() (*bool, bool)`

GetAcceptKidsOk returns a tuple with the AcceptKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptKids

`func (o *VacancyDraftVacancyDraftCreate) SetAcceptKids(v bool)`

SetAcceptKids sets AcceptKids field to given value.

### HasAcceptKids

`func (o *VacancyDraftVacancyDraftCreate) HasAcceptKids() bool`

HasAcceptKids returns a boolean if a field has been set.

### GetAcceptTemporary

`func (o *VacancyDraftVacancyDraftCreate) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacancyDraftVacancyDraftCreate) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacancyDraftVacancyDraftCreate) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacancyDraftVacancyDraftCreate) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacancyDraftVacancyDraftCreate) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacancyDraftVacancyDraftCreate) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAddress

`func (o *VacancyDraftVacancyDraftCreate) GetAddress() VacancyDraftAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacancyDraftVacancyDraftCreate) GetAddressOk() (*VacancyDraftAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacancyDraftVacancyDraftCreate) SetAddress(v VacancyDraftAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacancyDraftVacancyDraftCreate) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacancyDraftVacancyDraftCreate) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacancyDraftVacancyDraftCreate) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAllowMessages

`func (o *VacancyDraftVacancyDraftCreate) GetAllowMessages() bool`

GetAllowMessages returns the AllowMessages field if non-nil, zero value otherwise.

### GetAllowMessagesOk

`func (o *VacancyDraftVacancyDraftCreate) GetAllowMessagesOk() (*bool, bool)`

GetAllowMessagesOk returns a tuple with the AllowMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMessages

`func (o *VacancyDraftVacancyDraftCreate) SetAllowMessages(v bool)`

SetAllowMessages sets AllowMessages field to given value.

### HasAllowMessages

`func (o *VacancyDraftVacancyDraftCreate) HasAllowMessages() bool`

HasAllowMessages returns a boolean if a field has been set.

### GetArea

`func (o *VacancyDraftVacancyDraftCreate) GetArea() VacancyArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacancyDraftVacancyDraftCreate) GetAreaOk() (*VacancyArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacancyDraftVacancyDraftCreate) SetArea(v VacancyArea)`

SetArea sets Area field to given value.

### HasArea

`func (o *VacancyDraftVacancyDraftCreate) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *VacancyDraftVacancyDraftCreate) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *VacancyDraftVacancyDraftCreate) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetAreas

`func (o *VacancyDraftVacancyDraftCreate) GetAreas() []VacancyArea`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *VacancyDraftVacancyDraftCreate) GetAreasOk() (*[]VacancyArea, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *VacancyDraftVacancyDraftCreate) SetAreas(v []VacancyArea)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *VacancyDraftVacancyDraftCreate) HasAreas() bool`

HasAreas returns a boolean if a field has been set.

### SetAreasNil

`func (o *VacancyDraftVacancyDraftCreate) SetAreasNil(b bool)`

 SetAreasNil sets the value for Areas to be an explicit nil

### UnsetAreas
`func (o *VacancyDraftVacancyDraftCreate) UnsetAreas()`

UnsetAreas ensures that no value is present for Areas, not even an explicit nil
### GetAssignedManagerId

`func (o *VacancyDraftVacancyDraftCreate) GetAssignedManagerId() string`

GetAssignedManagerId returns the AssignedManagerId field if non-nil, zero value otherwise.

### GetAssignedManagerIdOk

`func (o *VacancyDraftVacancyDraftCreate) GetAssignedManagerIdOk() (*string, bool)`

GetAssignedManagerIdOk returns a tuple with the AssignedManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedManagerId

`func (o *VacancyDraftVacancyDraftCreate) SetAssignedManagerId(v string)`

SetAssignedManagerId sets AssignedManagerId field to given value.

### HasAssignedManagerId

`func (o *VacancyDraftVacancyDraftCreate) HasAssignedManagerId() bool`

HasAssignedManagerId returns a boolean if a field has been set.

### SetAssignedManagerIdNil

`func (o *VacancyDraftVacancyDraftCreate) SetAssignedManagerIdNil(b bool)`

 SetAssignedManagerIdNil sets the value for AssignedManagerId to be an explicit nil

### UnsetAssignedManagerId
`func (o *VacancyDraftVacancyDraftCreate) UnsetAssignedManagerId()`

UnsetAssignedManagerId ensures that no value is present for AssignedManagerId, not even an explicit nil
### GetBillingType

`func (o *VacancyDraftVacancyDraftCreate) GetBillingType() VacancyDraftBillingType`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacancyDraftVacancyDraftCreate) GetBillingTypeOk() (*VacancyDraftBillingType, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacancyDraftVacancyDraftCreate) SetBillingType(v VacancyDraftBillingType)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *VacancyDraftVacancyDraftCreate) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### SetBillingTypeNil

`func (o *VacancyDraftVacancyDraftCreate) SetBillingTypeNil(b bool)`

 SetBillingTypeNil sets the value for BillingType to be an explicit nil

### UnsetBillingType
`func (o *VacancyDraftVacancyDraftCreate) UnsetBillingType()`

UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
### GetBrandedTemplate

`func (o *VacancyDraftVacancyDraftCreate) GetBrandedTemplate() VacancyDraftBrandedTemplate`

GetBrandedTemplate returns the BrandedTemplate field if non-nil, zero value otherwise.

### GetBrandedTemplateOk

`func (o *VacancyDraftVacancyDraftCreate) GetBrandedTemplateOk() (*VacancyDraftBrandedTemplate, bool)`

GetBrandedTemplateOk returns a tuple with the BrandedTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandedTemplate

`func (o *VacancyDraftVacancyDraftCreate) SetBrandedTemplate(v VacancyDraftBrandedTemplate)`

SetBrandedTemplate sets BrandedTemplate field to given value.

### HasBrandedTemplate

`func (o *VacancyDraftVacancyDraftCreate) HasBrandedTemplate() bool`

HasBrandedTemplate returns a boolean if a field has been set.

### SetBrandedTemplateNil

`func (o *VacancyDraftVacancyDraftCreate) SetBrandedTemplateNil(b bool)`

 SetBrandedTemplateNil sets the value for BrandedTemplate to be an explicit nil

### UnsetBrandedTemplate
`func (o *VacancyDraftVacancyDraftCreate) UnsetBrandedTemplate()`

UnsetBrandedTemplate ensures that no value is present for BrandedTemplate, not even an explicit nil
### GetCode

`func (o *VacancyDraftVacancyDraftCreate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *VacancyDraftVacancyDraftCreate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *VacancyDraftVacancyDraftCreate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *VacancyDraftVacancyDraftCreate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *VacancyDraftVacancyDraftCreate) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *VacancyDraftVacancyDraftCreate) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetContacts

`func (o *VacancyDraftVacancyDraftCreate) GetContacts() VacancyDraftContacts`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacancyDraftVacancyDraftCreate) GetContactsOk() (*VacancyDraftContacts, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacancyDraftVacancyDraftCreate) SetContacts(v VacancyDraftContacts)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacancyDraftVacancyDraftCreate) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacancyDraftVacancyDraftCreate) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacancyDraftVacancyDraftCreate) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCustomEmployerName

`func (o *VacancyDraftVacancyDraftCreate) GetCustomEmployerName() string`

GetCustomEmployerName returns the CustomEmployerName field if non-nil, zero value otherwise.

### GetCustomEmployerNameOk

`func (o *VacancyDraftVacancyDraftCreate) GetCustomEmployerNameOk() (*string, bool)`

GetCustomEmployerNameOk returns a tuple with the CustomEmployerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomEmployerName

`func (o *VacancyDraftVacancyDraftCreate) SetCustomEmployerName(v string)`

SetCustomEmployerName sets CustomEmployerName field to given value.

### HasCustomEmployerName

`func (o *VacancyDraftVacancyDraftCreate) HasCustomEmployerName() bool`

HasCustomEmployerName returns a boolean if a field has been set.

### SetCustomEmployerNameNil

`func (o *VacancyDraftVacancyDraftCreate) SetCustomEmployerNameNil(b bool)`

 SetCustomEmployerNameNil sets the value for CustomEmployerName to be an explicit nil

### UnsetCustomEmployerName
`func (o *VacancyDraftVacancyDraftCreate) UnsetCustomEmployerName()`

UnsetCustomEmployerName ensures that no value is present for CustomEmployerName, not even an explicit nil
### GetDepartment

`func (o *VacancyDraftVacancyDraftCreate) GetDepartment() VacancyDepartment`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacancyDraftVacancyDraftCreate) GetDepartmentOk() (*VacancyDepartment, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacancyDraftVacancyDraftCreate) SetDepartment(v VacancyDepartment)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *VacancyDraftVacancyDraftCreate) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *VacancyDraftVacancyDraftCreate) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacancyDraftVacancyDraftCreate) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetDescription

`func (o *VacancyDraftVacancyDraftCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VacancyDraftVacancyDraftCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VacancyDraftVacancyDraftCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VacancyDraftVacancyDraftCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *VacancyDraftVacancyDraftCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *VacancyDraftVacancyDraftCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDriverLicenseTypes

`func (o *VacancyDraftVacancyDraftCreate) GetDriverLicenseTypes() []VacancyDriverLicenceTypeItem`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *VacancyDraftVacancyDraftCreate) GetDriverLicenseTypesOk() (*[]VacancyDriverLicenceTypeItem, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *VacancyDraftVacancyDraftCreate) SetDriverLicenseTypes(v []VacancyDriverLicenceTypeItem)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.

### HasDriverLicenseTypes

`func (o *VacancyDraftVacancyDraftCreate) HasDriverLicenseTypes() bool`

HasDriverLicenseTypes returns a boolean if a field has been set.

### GetEmployment

`func (o *VacancyDraftVacancyDraftCreate) GetEmployment() VacancyDraftEmployment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacancyDraftVacancyDraftCreate) GetEmploymentOk() (*VacancyDraftEmployment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacancyDraftVacancyDraftCreate) SetEmployment(v VacancyDraftEmployment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacancyDraftVacancyDraftCreate) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacancyDraftVacancyDraftCreate) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacancyDraftVacancyDraftCreate) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacancyDraftVacancyDraftCreate) GetExperience() VacancyExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacancyDraftVacancyDraftCreate) GetExperienceOk() (*VacancyExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacancyDraftVacancyDraftCreate) SetExperience(v VacancyExperience)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *VacancyDraftVacancyDraftCreate) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *VacancyDraftVacancyDraftCreate) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacancyDraftVacancyDraftCreate) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetKeySkills

`func (o *VacancyDraftVacancyDraftCreate) GetKeySkills() []VacancyDraftKeySkillItem`

GetKeySkills returns the KeySkills field if non-nil, zero value otherwise.

### GetKeySkillsOk

`func (o *VacancyDraftVacancyDraftCreate) GetKeySkillsOk() (*[]VacancyDraftKeySkillItem, bool)`

GetKeySkillsOk returns a tuple with the KeySkills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeySkills

`func (o *VacancyDraftVacancyDraftCreate) SetKeySkills(v []VacancyDraftKeySkillItem)`

SetKeySkills sets KeySkills field to given value.

### HasKeySkills

`func (o *VacancyDraftVacancyDraftCreate) HasKeySkills() bool`

HasKeySkills returns a boolean if a field has been set.

### SetKeySkillsNil

`func (o *VacancyDraftVacancyDraftCreate) SetKeySkillsNil(b bool)`

 SetKeySkillsNil sets the value for KeySkills to be an explicit nil

### UnsetKeySkills
`func (o *VacancyDraftVacancyDraftCreate) UnsetKeySkills()`

UnsetKeySkills ensures that no value is present for KeySkills, not even an explicit nil
### GetLanguages

`func (o *VacancyDraftVacancyDraftCreate) GetLanguages() []VacancyLanguage`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *VacancyDraftVacancyDraftCreate) GetLanguagesOk() (*[]VacancyLanguage, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *VacancyDraftVacancyDraftCreate) SetLanguages(v []VacancyLanguage)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *VacancyDraftVacancyDraftCreate) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetName

`func (o *VacancyDraftVacancyDraftCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyDraftVacancyDraftCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyDraftVacancyDraftCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacancyDraftVacancyDraftCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *VacancyDraftVacancyDraftCreate) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *VacancyDraftVacancyDraftCreate) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetProfessionalRoles

`func (o *VacancyDraftVacancyDraftCreate) GetProfessionalRoles() []VacancyDraftProfessionalRoleItem`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacancyDraftVacancyDraftCreate) GetProfessionalRolesOk() (*[]VacancyDraftProfessionalRoleItem, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacancyDraftVacancyDraftCreate) SetProfessionalRoles(v []VacancyDraftProfessionalRoleItem)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *VacancyDraftVacancyDraftCreate) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### SetProfessionalRolesNil

`func (o *VacancyDraftVacancyDraftCreate) SetProfessionalRolesNil(b bool)`

 SetProfessionalRolesNil sets the value for ProfessionalRoles to be an explicit nil

### UnsetProfessionalRoles
`func (o *VacancyDraftVacancyDraftCreate) UnsetProfessionalRoles()`

UnsetProfessionalRoles ensures that no value is present for ProfessionalRoles, not even an explicit nil
### GetResponseLetterRequired

`func (o *VacancyDraftVacancyDraftCreate) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacancyDraftVacancyDraftCreate) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacancyDraftVacancyDraftCreate) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.

### HasResponseLetterRequired

`func (o *VacancyDraftVacancyDraftCreate) HasResponseLetterRequired() bool`

HasResponseLetterRequired returns a boolean if a field has been set.

### GetResponseNotifications

`func (o *VacancyDraftVacancyDraftCreate) GetResponseNotifications() bool`

GetResponseNotifications returns the ResponseNotifications field if non-nil, zero value otherwise.

### GetResponseNotificationsOk

`func (o *VacancyDraftVacancyDraftCreate) GetResponseNotificationsOk() (*bool, bool)`

GetResponseNotificationsOk returns a tuple with the ResponseNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseNotifications

`func (o *VacancyDraftVacancyDraftCreate) SetResponseNotifications(v bool)`

SetResponseNotifications sets ResponseNotifications field to given value.

### HasResponseNotifications

`func (o *VacancyDraftVacancyDraftCreate) HasResponseNotifications() bool`

HasResponseNotifications returns a boolean if a field has been set.

### GetResponseUrl

`func (o *VacancyDraftVacancyDraftCreate) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacancyDraftVacancyDraftCreate) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacancyDraftVacancyDraftCreate) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacancyDraftVacancyDraftCreate) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacancyDraftVacancyDraftCreate) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacancyDraftVacancyDraftCreate) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacancyDraftVacancyDraftCreate) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacancyDraftVacancyDraftCreate) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacancyDraftVacancyDraftCreate) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacancyDraftVacancyDraftCreate) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacancyDraftVacancyDraftCreate) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacancyDraftVacancyDraftCreate) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacancyDraftVacancyDraftCreate) GetSchedule() VacancySchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacancyDraftVacancyDraftCreate) GetScheduleOk() (*VacancySchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacancyDraftVacancyDraftCreate) SetSchedule(v VacancySchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *VacancyDraftVacancyDraftCreate) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *VacancyDraftVacancyDraftCreate) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *VacancyDraftVacancyDraftCreate) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetTest

`func (o *VacancyDraftVacancyDraftCreate) GetTest() VacancyDraftTest`

GetTest returns the Test field if non-nil, zero value otherwise.

### GetTestOk

`func (o *VacancyDraftVacancyDraftCreate) GetTestOk() (*VacancyDraftTest, bool)`

GetTestOk returns a tuple with the Test field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTest

`func (o *VacancyDraftVacancyDraftCreate) SetTest(v VacancyDraftTest)`

SetTest sets Test field to given value.

### HasTest

`func (o *VacancyDraftVacancyDraftCreate) HasTest() bool`

HasTest returns a boolean if a field has been set.

### SetTestNil

`func (o *VacancyDraftVacancyDraftCreate) SetTestNil(b bool)`

 SetTestNil sets the value for Test to be an explicit nil

### UnsetTest
`func (o *VacancyDraftVacancyDraftCreate) UnsetTest()`

UnsetTest ensures that no value is present for Test, not even an explicit nil
### GetType

`func (o *VacancyDraftVacancyDraftCreate) GetType() VacancyDraftType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacancyDraftVacancyDraftCreate) GetTypeOk() (*VacancyDraftType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacancyDraftVacancyDraftCreate) SetType(v VacancyDraftType)`

SetType sets Type field to given value.

### HasType

`func (o *VacancyDraftVacancyDraftCreate) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VacancyDraftVacancyDraftCreate) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VacancyDraftVacancyDraftCreate) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetWithZp

`func (o *VacancyDraftVacancyDraftCreate) GetWithZp() bool`

GetWithZp returns the WithZp field if non-nil, zero value otherwise.

### GetWithZpOk

`func (o *VacancyDraftVacancyDraftCreate) GetWithZpOk() (*bool, bool)`

GetWithZpOk returns a tuple with the WithZp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithZp

`func (o *VacancyDraftVacancyDraftCreate) SetWithZp(v bool)`

SetWithZp sets WithZp field to given value.

### HasWithZp

`func (o *VacancyDraftVacancyDraftCreate) HasWithZp() bool`

HasWithZp returns a boolean if a field has been set.

### GetWorkingDays

`func (o *VacancyDraftVacancyDraftCreate) GetWorkingDays() []VacancyWorkingDayItem`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacancyDraftVacancyDraftCreate) GetWorkingDaysOk() (*[]VacancyWorkingDayItem, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacancyDraftVacancyDraftCreate) SetWorkingDays(v []VacancyWorkingDayItem)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacancyDraftVacancyDraftCreate) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacancyDraftVacancyDraftCreate) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacancyDraftVacancyDraftCreate) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacancyDraftVacancyDraftCreate) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItem`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacancyDraftVacancyDraftCreate) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItem, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacancyDraftVacancyDraftCreate) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItem)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacancyDraftVacancyDraftCreate) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacancyDraftVacancyDraftCreate) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacancyDraftVacancyDraftCreate) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacancyDraftVacancyDraftCreate) GetWorkingTimeModes() []VacancyWorkingTimeModeItem`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacancyDraftVacancyDraftCreate) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItem, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacancyDraftVacancyDraftCreate) SetWorkingTimeModes(v []VacancyWorkingTimeModeItem)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacancyDraftVacancyDraftCreate) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacancyDraftVacancyDraftCreate) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacancyDraftVacancyDraftCreate) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


