# VacanciesVacanciesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptIncompleteResumes** | **bool** | Разрешен ли отклик на вакансию неполным резюме | 
**AcceptTemporary** | Pointer to **NullableBool** | Указание, что вакансия доступна с временным трудоустройством | [optional] 
**Address** | Pointer to [**NullableVacancyAddressRawOutput**](VacancyAddressRawOutput.md) |  | [optional] 
**AdvResponseUrl** | Pointer to **NullableString** | URL для регистрации нажатия кнопки отклика (устаревшее, сейчас для регистрации нажатия кнопки отклика нужно заполнять хедер adv-context  в запросе [Резюме, сгруппированные по возможности отклика на данную вакансию](#tag/Rezyume.-Prosmotr-informacii/operation/get-resumes-by-status))  | [optional] 
**AlternateUrl** | **string** | Ссылка на представление вакансии на сайте | 
**ApplyAlternateUrl** | **string** | Ссылка на отклик на вакансию на сайте | 
**Archived** | Pointer to **NullableBool** | Находится ли данная вакансия в архиве | [optional] 
**Area** | [**IncludesArea**](IncludesArea.md) |  | 
**Contacts** | Pointer to [**NullableVacancyContactsOutput**](VacancyContactsOutput.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | Дата и время публикации вакансии | [optional] 
**Department** | [**NullableVacancyDepartmentOutput**](VacancyDepartmentOutput.md) |  | 
**Employer** | [**VacanciesEmployerPublic**](VacanciesEmployerPublic.md) |  | 
**HasTest** | **bool** | Информация о наличии прикрепленного тестового задании к вакансии | 
**Id** | **string** | Идентификатор вакансии | 
**InsiderInterview** | Pointer to [**NullableVacancyInsiderInterview**](VacancyInsiderInterview.md) |  | [optional] 
**MetroStations** | Pointer to [**IncludesMetroStation**](IncludesMetroStation.md) |  | [optional] 
**Name** | **string** | Название | 
**Premium** | Pointer to **NullableBool** | Является ли данная вакансия премиум-вакансией | [optional] 
**ProfessionalRoles** | [**[]VacancyProfessionalRoleItemOutput**](VacancyProfessionalRoleItemOutput.md) | Список профессиональных ролей | 
**PublishedAt** | **string** | Дата и время публикации вакансии | 
**Relations** | [**[]VacancyRelationItem**](VacancyRelationItem.md) | Возвращает связи соискателя с вакансией. Значения из поля &#x60;vacancy_relation&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**ResponseLetterRequired** | **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | [**NullableVacancySalary**](VacancySalary.md) |  | 
**Schedule** | Pointer to [**VacancyScheduleOutput**](VacancyScheduleOutput.md) |  | [optional] 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций.  Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60; | [optional] 
**Type** | [**VacancyTypeOutput**](VacancyTypeOutput.md) |  | 
**Url** | **string** | URL вакансии | 
**WorkingDays** | Pointer to [**[]VacancyWorkingDayItemOutput**](VacancyWorkingDayItemOutput.md) | Список рабочих дней | [optional] 
**WorkingTimeIntervals** | Pointer to [**[]VacancyWorkingTimeIntervalItemOutput**](VacancyWorkingTimeIntervalItemOutput.md) | Список с временными интервалами работы | [optional] 
**WorkingTimeModes** | Pointer to [**[]VacancyWorkingTimeModeItemOutput**](VacancyWorkingTimeModeItemOutput.md) | Список режимов времени работы | [optional] 
**Counters** | Pointer to [**VacancyCounters**](VacancyCounters.md) |  | [optional] 
**Employment** | Pointer to [**NullableVacancyEmploymentOutput**](VacancyEmploymentOutput.md) |  | [optional] 
**Experience** | Pointer to [**NullableVacancyExperienceOutput**](VacancyExperienceOutput.md) |  | [optional] 
**Snippet** | [**VacancySnippet**](VacancySnippet.md) |  | 
**ShowLogoInSearch** | Pointer to **NullableBool** | Отображать ли лого для вакансии в поисковой выдаче | [optional] 

## Methods

### NewVacanciesVacanciesItem

`func NewVacanciesVacanciesItem(acceptIncompleteResumes bool, alternateUrl string, applyAlternateUrl string, area IncludesArea, department NullableVacancyDepartmentOutput, employer VacanciesEmployerPublic, hasTest bool, id string, name string, professionalRoles []VacancyProfessionalRoleItemOutput, publishedAt string, relations []VacancyRelationItem, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string, snippet VacancySnippet, ) *VacanciesVacanciesItem`

NewVacanciesVacanciesItem instantiates a new VacanciesVacanciesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacanciesItemWithDefaults

`func NewVacanciesVacanciesItemWithDefaults() *VacanciesVacanciesItem`

NewVacanciesVacanciesItemWithDefaults instantiates a new VacanciesVacanciesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptIncompleteResumes

`func (o *VacanciesVacanciesItem) GetAcceptIncompleteResumes() bool`

GetAcceptIncompleteResumes returns the AcceptIncompleteResumes field if non-nil, zero value otherwise.

### GetAcceptIncompleteResumesOk

`func (o *VacanciesVacanciesItem) GetAcceptIncompleteResumesOk() (*bool, bool)`

GetAcceptIncompleteResumesOk returns a tuple with the AcceptIncompleteResumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptIncompleteResumes

`func (o *VacanciesVacanciesItem) SetAcceptIncompleteResumes(v bool)`

SetAcceptIncompleteResumes sets AcceptIncompleteResumes field to given value.


### GetAcceptTemporary

`func (o *VacanciesVacanciesItem) GetAcceptTemporary() bool`

GetAcceptTemporary returns the AcceptTemporary field if non-nil, zero value otherwise.

### GetAcceptTemporaryOk

`func (o *VacanciesVacanciesItem) GetAcceptTemporaryOk() (*bool, bool)`

GetAcceptTemporaryOk returns a tuple with the AcceptTemporary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptTemporary

`func (o *VacanciesVacanciesItem) SetAcceptTemporary(v bool)`

SetAcceptTemporary sets AcceptTemporary field to given value.

### HasAcceptTemporary

`func (o *VacanciesVacanciesItem) HasAcceptTemporary() bool`

HasAcceptTemporary returns a boolean if a field has been set.

### SetAcceptTemporaryNil

`func (o *VacanciesVacanciesItem) SetAcceptTemporaryNil(b bool)`

 SetAcceptTemporaryNil sets the value for AcceptTemporary to be an explicit nil

### UnsetAcceptTemporary
`func (o *VacanciesVacanciesItem) UnsetAcceptTemporary()`

UnsetAcceptTemporary ensures that no value is present for AcceptTemporary, not even an explicit nil
### GetAddress

`func (o *VacanciesVacanciesItem) GetAddress() VacancyAddressRawOutput`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesVacanciesItem) GetAddressOk() (*VacancyAddressRawOutput, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesVacanciesItem) SetAddress(v VacancyAddressRawOutput)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacanciesVacanciesItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacanciesVacanciesItem) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacanciesVacanciesItem) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAdvResponseUrl

`func (o *VacanciesVacanciesItem) GetAdvResponseUrl() string`

GetAdvResponseUrl returns the AdvResponseUrl field if non-nil, zero value otherwise.

### GetAdvResponseUrlOk

`func (o *VacanciesVacanciesItem) GetAdvResponseUrlOk() (*string, bool)`

GetAdvResponseUrlOk returns a tuple with the AdvResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvResponseUrl

`func (o *VacanciesVacanciesItem) SetAdvResponseUrl(v string)`

SetAdvResponseUrl sets AdvResponseUrl field to given value.

### HasAdvResponseUrl

`func (o *VacanciesVacanciesItem) HasAdvResponseUrl() bool`

HasAdvResponseUrl returns a boolean if a field has been set.

### SetAdvResponseUrlNil

`func (o *VacanciesVacanciesItem) SetAdvResponseUrlNil(b bool)`

 SetAdvResponseUrlNil sets the value for AdvResponseUrl to be an explicit nil

### UnsetAdvResponseUrl
`func (o *VacanciesVacanciesItem) UnsetAdvResponseUrl()`

UnsetAdvResponseUrl ensures that no value is present for AdvResponseUrl, not even an explicit nil
### GetAlternateUrl

`func (o *VacanciesVacanciesItem) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVacanciesItem) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVacanciesItem) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesVacanciesItem) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesVacanciesItem) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesVacanciesItem) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetArchived

`func (o *VacanciesVacanciesItem) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesVacanciesItem) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesVacanciesItem) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *VacanciesVacanciesItem) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### SetArchivedNil

`func (o *VacanciesVacanciesItem) SetArchivedNil(b bool)`

 SetArchivedNil sets the value for Archived to be an explicit nil

### UnsetArchived
`func (o *VacanciesVacanciesItem) UnsetArchived()`

UnsetArchived ensures that no value is present for Archived, not even an explicit nil
### GetArea

`func (o *VacanciesVacanciesItem) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesVacanciesItem) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesVacanciesItem) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetContacts

`func (o *VacanciesVacanciesItem) GetContacts() VacancyContactsOutput`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *VacanciesVacanciesItem) GetContactsOk() (*VacancyContactsOutput, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *VacanciesVacanciesItem) SetContacts(v VacancyContactsOutput)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *VacanciesVacanciesItem) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### SetContactsNil

`func (o *VacanciesVacanciesItem) SetContactsNil(b bool)`

 SetContactsNil sets the value for Contacts to be an explicit nil

### UnsetContacts
`func (o *VacanciesVacanciesItem) UnsetContacts()`

UnsetContacts ensures that no value is present for Contacts, not even an explicit nil
### GetCreatedAt

`func (o *VacanciesVacanciesItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesVacanciesItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesVacanciesItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VacanciesVacanciesItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepartment

`func (o *VacanciesVacanciesItem) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesVacanciesItem) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesVacanciesItem) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.


### SetDepartmentNil

`func (o *VacanciesVacanciesItem) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesVacanciesItem) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetEmployer

`func (o *VacanciesVacanciesItem) GetEmployer() VacanciesEmployerPublic`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesVacanciesItem) GetEmployerOk() (*VacanciesEmployerPublic, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesVacanciesItem) SetEmployer(v VacanciesEmployerPublic)`

SetEmployer sets Employer field to given value.


### GetHasTest

`func (o *VacanciesVacanciesItem) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesVacanciesItem) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesVacanciesItem) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesVacanciesItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVacanciesItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVacanciesItem) SetId(v string)`

SetId sets Id field to given value.


### GetInsiderInterview

`func (o *VacanciesVacanciesItem) GetInsiderInterview() VacancyInsiderInterview`

GetInsiderInterview returns the InsiderInterview field if non-nil, zero value otherwise.

### GetInsiderInterviewOk

`func (o *VacanciesVacanciesItem) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool)`

GetInsiderInterviewOk returns a tuple with the InsiderInterview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderInterview

`func (o *VacanciesVacanciesItem) SetInsiderInterview(v VacancyInsiderInterview)`

SetInsiderInterview sets InsiderInterview field to given value.

### HasInsiderInterview

`func (o *VacanciesVacanciesItem) HasInsiderInterview() bool`

HasInsiderInterview returns a boolean if a field has been set.

### SetInsiderInterviewNil

`func (o *VacanciesVacanciesItem) SetInsiderInterviewNil(b bool)`

 SetInsiderInterviewNil sets the value for InsiderInterview to be an explicit nil

### UnsetInsiderInterview
`func (o *VacanciesVacanciesItem) UnsetInsiderInterview()`

UnsetInsiderInterview ensures that no value is present for InsiderInterview, not even an explicit nil
### GetMetroStations

`func (o *VacanciesVacanciesItem) GetMetroStations() IncludesMetroStation`

GetMetroStations returns the MetroStations field if non-nil, zero value otherwise.

### GetMetroStationsOk

`func (o *VacanciesVacanciesItem) GetMetroStationsOk() (*IncludesMetroStation, bool)`

GetMetroStationsOk returns a tuple with the MetroStations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroStations

`func (o *VacanciesVacanciesItem) SetMetroStations(v IncludesMetroStation)`

SetMetroStations sets MetroStations field to given value.

### HasMetroStations

`func (o *VacanciesVacanciesItem) HasMetroStations() bool`

HasMetroStations returns a boolean if a field has been set.

### GetName

`func (o *VacanciesVacanciesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesVacanciesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesVacanciesItem) SetName(v string)`

SetName sets Name field to given value.


### GetPremium

`func (o *VacanciesVacanciesItem) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesVacanciesItem) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesVacanciesItem) SetPremium(v bool)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *VacanciesVacanciesItem) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### SetPremiumNil

`func (o *VacanciesVacanciesItem) SetPremiumNil(b bool)`

 SetPremiumNil sets the value for Premium to be an explicit nil

### UnsetPremium
`func (o *VacanciesVacanciesItem) UnsetPremium()`

UnsetPremium ensures that no value is present for Premium, not even an explicit nil
### GetProfessionalRoles

`func (o *VacanciesVacanciesItem) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacanciesVacanciesItem) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItemOutput, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacanciesVacanciesItem) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput)`

SetProfessionalRoles sets ProfessionalRoles field to given value.


### GetPublishedAt

`func (o *VacanciesVacanciesItem) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesVacanciesItem) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesVacanciesItem) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetRelations

`func (o *VacanciesVacanciesItem) GetRelations() []VacancyRelationItem`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *VacanciesVacanciesItem) GetRelationsOk() (*[]VacancyRelationItem, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *VacanciesVacanciesItem) SetRelations(v []VacancyRelationItem)`

SetRelations sets Relations field to given value.


### GetResponseLetterRequired

`func (o *VacanciesVacanciesItem) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesVacanciesItem) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesVacanciesItem) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesVacanciesItem) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesVacanciesItem) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesVacanciesItem) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesVacanciesItem) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesVacanciesItem) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesVacanciesItem) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesVacanciesItem) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesVacanciesItem) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesVacanciesItem) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.


### SetSalaryNil

`func (o *VacanciesVacanciesItem) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesVacanciesItem) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedule

`func (o *VacanciesVacanciesItem) GetSchedule() VacancyScheduleOutput`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *VacanciesVacanciesItem) GetScheduleOk() (*VacancyScheduleOutput, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *VacanciesVacanciesItem) SetSchedule(v VacancyScheduleOutput)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *VacanciesVacanciesItem) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSortPointDistance

`func (o *VacanciesVacanciesItem) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesVacanciesItem) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesVacanciesItem) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesVacanciesItem) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesVacanciesItem) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesVacanciesItem) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
### GetType

`func (o *VacanciesVacanciesItem) GetType() VacancyTypeOutput`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesVacanciesItem) GetTypeOk() (*VacancyTypeOutput, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesVacanciesItem) SetType(v VacancyTypeOutput)`

SetType sets Type field to given value.


### GetUrl

`func (o *VacanciesVacanciesItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesVacanciesItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesVacanciesItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWorkingDays

`func (o *VacanciesVacanciesItem) GetWorkingDays() []VacancyWorkingDayItemOutput`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *VacanciesVacanciesItem) GetWorkingDaysOk() (*[]VacancyWorkingDayItemOutput, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *VacanciesVacanciesItem) SetWorkingDays(v []VacancyWorkingDayItemOutput)`

SetWorkingDays sets WorkingDays field to given value.

### HasWorkingDays

`func (o *VacanciesVacanciesItem) HasWorkingDays() bool`

HasWorkingDays returns a boolean if a field has been set.

### SetWorkingDaysNil

`func (o *VacanciesVacanciesItem) SetWorkingDaysNil(b bool)`

 SetWorkingDaysNil sets the value for WorkingDays to be an explicit nil

### UnsetWorkingDays
`func (o *VacanciesVacanciesItem) UnsetWorkingDays()`

UnsetWorkingDays ensures that no value is present for WorkingDays, not even an explicit nil
### GetWorkingTimeIntervals

`func (o *VacanciesVacanciesItem) GetWorkingTimeIntervals() []VacancyWorkingTimeIntervalItemOutput`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *VacanciesVacanciesItem) GetWorkingTimeIntervalsOk() (*[]VacancyWorkingTimeIntervalItemOutput, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *VacanciesVacanciesItem) SetWorkingTimeIntervals(v []VacancyWorkingTimeIntervalItemOutput)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.

### HasWorkingTimeIntervals

`func (o *VacanciesVacanciesItem) HasWorkingTimeIntervals() bool`

HasWorkingTimeIntervals returns a boolean if a field has been set.

### SetWorkingTimeIntervalsNil

`func (o *VacanciesVacanciesItem) SetWorkingTimeIntervalsNil(b bool)`

 SetWorkingTimeIntervalsNil sets the value for WorkingTimeIntervals to be an explicit nil

### UnsetWorkingTimeIntervals
`func (o *VacanciesVacanciesItem) UnsetWorkingTimeIntervals()`

UnsetWorkingTimeIntervals ensures that no value is present for WorkingTimeIntervals, not even an explicit nil
### GetWorkingTimeModes

`func (o *VacanciesVacanciesItem) GetWorkingTimeModes() []VacancyWorkingTimeModeItemOutput`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *VacanciesVacanciesItem) GetWorkingTimeModesOk() (*[]VacancyWorkingTimeModeItemOutput, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *VacanciesVacanciesItem) SetWorkingTimeModes(v []VacancyWorkingTimeModeItemOutput)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.

### HasWorkingTimeModes

`func (o *VacanciesVacanciesItem) HasWorkingTimeModes() bool`

HasWorkingTimeModes returns a boolean if a field has been set.

### SetWorkingTimeModesNil

`func (o *VacanciesVacanciesItem) SetWorkingTimeModesNil(b bool)`

 SetWorkingTimeModesNil sets the value for WorkingTimeModes to be an explicit nil

### UnsetWorkingTimeModes
`func (o *VacanciesVacanciesItem) UnsetWorkingTimeModes()`

UnsetWorkingTimeModes ensures that no value is present for WorkingTimeModes, not even an explicit nil
### GetCounters

`func (o *VacanciesVacanciesItem) GetCounters() VacancyCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *VacanciesVacanciesItem) GetCountersOk() (*VacancyCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *VacanciesVacanciesItem) SetCounters(v VacancyCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *VacanciesVacanciesItem) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetEmployment

`func (o *VacanciesVacanciesItem) GetEmployment() VacancyEmploymentOutput`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacanciesVacanciesItem) GetEmploymentOk() (*VacancyEmploymentOutput, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacanciesVacanciesItem) SetEmployment(v VacancyEmploymentOutput)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacanciesVacanciesItem) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacanciesVacanciesItem) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacanciesVacanciesItem) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacanciesVacanciesItem) GetExperience() VacancyExperienceOutput`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacanciesVacanciesItem) GetExperienceOk() (*VacancyExperienceOutput, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacanciesVacanciesItem) SetExperience(v VacancyExperienceOutput)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *VacanciesVacanciesItem) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *VacanciesVacanciesItem) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacanciesVacanciesItem) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetSnippet

`func (o *VacanciesVacanciesItem) GetSnippet() VacancySnippet`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *VacanciesVacanciesItem) GetSnippetOk() (*VacancySnippet, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *VacanciesVacanciesItem) SetSnippet(v VacancySnippet)`

SetSnippet sets Snippet field to given value.


### GetShowLogoInSearch

`func (o *VacanciesVacanciesItem) GetShowLogoInSearch() bool`

GetShowLogoInSearch returns the ShowLogoInSearch field if non-nil, zero value otherwise.

### GetShowLogoInSearchOk

`func (o *VacanciesVacanciesItem) GetShowLogoInSearchOk() (*bool, bool)`

GetShowLogoInSearchOk returns a tuple with the ShowLogoInSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLogoInSearch

`func (o *VacanciesVacanciesItem) SetShowLogoInSearch(v bool)`

SetShowLogoInSearch sets ShowLogoInSearch field to given value.

### HasShowLogoInSearch

`func (o *VacanciesVacanciesItem) HasShowLogoInSearch() bool`

HasShowLogoInSearch returns a boolean if a field has been set.

### SetShowLogoInSearchNil

`func (o *VacanciesVacanciesItem) SetShowLogoInSearchNil(b bool)`

 SetShowLogoInSearchNil sets the value for ShowLogoInSearch to be an explicit nil

### UnsetShowLogoInSearch
`func (o *VacanciesVacanciesItem) UnsetShowLogoInSearch()`

UnsetShowLogoInSearch ensures that no value is present for ShowLogoInSearch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


