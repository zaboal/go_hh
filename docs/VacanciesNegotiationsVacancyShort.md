# VacanciesNegotiationsVacancyShort

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**NullableVacancyAddressRawOutput**](VacancyAddressRawOutput.md) |  | [optional] 
**AlternateUrl** | **string** | Ссылка на представление вакансии на сайте | 
**ApplyAlternateUrl** | **string** | Ссылка на отклик на вакансию на сайте | 
**Archived** | **bool** | Находится ли данная вакансия в архиве | 
**Area** | [**IncludesArea**](IncludesArea.md) |  | 
**CreatedAt** | Pointer to **string** | Дата и время публикации вакансии | [optional] 
**Department** | [**NullableVacancyDepartmentOutput**](VacancyDepartmentOutput.md) |  | 
**Employer** | [**NullableVacanciesEmployerPublic**](VacanciesEmployerPublic.md) |  | 
**HasTest** | **bool** | Информация о наличии прикрепленного тестового задании к вакансии | 
**Id** | **string** | Идентификатор вакансии | 
**InsiderInterview** | Pointer to [**NullableVacancyInsiderInterview**](VacancyInsiderInterview.md) |  | [optional] 
**Name** | **string** | Название вакансии | 
**Premium** | **bool** | Является ли данная вакансия премиум-вакансией | 
**ProfessionalRoles** | Pointer to [**[]VacancyProfessionalRoleItemOutput**](VacancyProfessionalRoleItemOutput.md) | Список профессиональных ролей | [optional] 
**PublishedAt** | **string** | Дата и время публикации вакансии | 
**ResponseLetterRequired** | **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | [**NullableVacancySalary**](VacancySalary.md) |  | 
**ShowLogoInSearch** | Pointer to **NullableBool** | Отображать ли лого для вакансии в поисковой выдаче | [optional] 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций. Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60;  | [optional] 
**Type** | [**VacancyTypeOutput**](VacancyTypeOutput.md) |  | 
**Url** | **string** | URL вакансии | 

## Methods

### NewVacanciesNegotiationsVacancyShort

`func NewVacanciesNegotiationsVacancyShort(alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, department NullableVacancyDepartmentOutput, employer NullableVacanciesEmployerPublic, hasTest bool, id string, name string, premium bool, publishedAt string, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string, ) *VacanciesNegotiationsVacancyShort`

NewVacanciesNegotiationsVacancyShort instantiates a new VacanciesNegotiationsVacancyShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesNegotiationsVacancyShortWithDefaults

`func NewVacanciesNegotiationsVacancyShortWithDefaults() *VacanciesNegotiationsVacancyShort`

NewVacanciesNegotiationsVacancyShortWithDefaults instantiates a new VacanciesNegotiationsVacancyShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VacanciesNegotiationsVacancyShort) GetAddress() VacancyAddressRawOutput`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesNegotiationsVacancyShort) GetAddressOk() (*VacancyAddressRawOutput, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesNegotiationsVacancyShort) SetAddress(v VacancyAddressRawOutput)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacanciesNegotiationsVacancyShort) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacanciesNegotiationsVacancyShort) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacanciesNegotiationsVacancyShort) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAlternateUrl

`func (o *VacanciesNegotiationsVacancyShort) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesNegotiationsVacancyShort) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesNegotiationsVacancyShort) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesNegotiationsVacancyShort) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesNegotiationsVacancyShort) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesNegotiationsVacancyShort) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetArchived

`func (o *VacanciesNegotiationsVacancyShort) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesNegotiationsVacancyShort) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesNegotiationsVacancyShort) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArea

`func (o *VacanciesNegotiationsVacancyShort) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesNegotiationsVacancyShort) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesNegotiationsVacancyShort) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetCreatedAt

`func (o *VacanciesNegotiationsVacancyShort) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesNegotiationsVacancyShort) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesNegotiationsVacancyShort) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VacanciesNegotiationsVacancyShort) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepartment

`func (o *VacanciesNegotiationsVacancyShort) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesNegotiationsVacancyShort) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesNegotiationsVacancyShort) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.


### SetDepartmentNil

`func (o *VacanciesNegotiationsVacancyShort) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesNegotiationsVacancyShort) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetEmployer

`func (o *VacanciesNegotiationsVacancyShort) GetEmployer() VacanciesEmployerPublic`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesNegotiationsVacancyShort) GetEmployerOk() (*VacanciesEmployerPublic, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesNegotiationsVacancyShort) SetEmployer(v VacanciesEmployerPublic)`

SetEmployer sets Employer field to given value.


### SetEmployerNil

`func (o *VacanciesNegotiationsVacancyShort) SetEmployerNil(b bool)`

 SetEmployerNil sets the value for Employer to be an explicit nil

### UnsetEmployer
`func (o *VacanciesNegotiationsVacancyShort) UnsetEmployer()`

UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
### GetHasTest

`func (o *VacanciesNegotiationsVacancyShort) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesNegotiationsVacancyShort) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesNegotiationsVacancyShort) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesNegotiationsVacancyShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesNegotiationsVacancyShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesNegotiationsVacancyShort) SetId(v string)`

SetId sets Id field to given value.


### GetInsiderInterview

`func (o *VacanciesNegotiationsVacancyShort) GetInsiderInterview() VacancyInsiderInterview`

GetInsiderInterview returns the InsiderInterview field if non-nil, zero value otherwise.

### GetInsiderInterviewOk

`func (o *VacanciesNegotiationsVacancyShort) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool)`

GetInsiderInterviewOk returns a tuple with the InsiderInterview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderInterview

`func (o *VacanciesNegotiationsVacancyShort) SetInsiderInterview(v VacancyInsiderInterview)`

SetInsiderInterview sets InsiderInterview field to given value.

### HasInsiderInterview

`func (o *VacanciesNegotiationsVacancyShort) HasInsiderInterview() bool`

HasInsiderInterview returns a boolean if a field has been set.

### SetInsiderInterviewNil

`func (o *VacanciesNegotiationsVacancyShort) SetInsiderInterviewNil(b bool)`

 SetInsiderInterviewNil sets the value for InsiderInterview to be an explicit nil

### UnsetInsiderInterview
`func (o *VacanciesNegotiationsVacancyShort) UnsetInsiderInterview()`

UnsetInsiderInterview ensures that no value is present for InsiderInterview, not even an explicit nil
### GetName

`func (o *VacanciesNegotiationsVacancyShort) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesNegotiationsVacancyShort) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesNegotiationsVacancyShort) SetName(v string)`

SetName sets Name field to given value.


### GetPremium

`func (o *VacanciesNegotiationsVacancyShort) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesNegotiationsVacancyShort) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesNegotiationsVacancyShort) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetProfessionalRoles

`func (o *VacanciesNegotiationsVacancyShort) GetProfessionalRoles() []VacancyProfessionalRoleItemOutput`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *VacanciesNegotiationsVacancyShort) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItemOutput, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *VacanciesNegotiationsVacancyShort) SetProfessionalRoles(v []VacancyProfessionalRoleItemOutput)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *VacanciesNegotiationsVacancyShort) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetPublishedAt

`func (o *VacanciesNegotiationsVacancyShort) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesNegotiationsVacancyShort) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesNegotiationsVacancyShort) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetResponseLetterRequired

`func (o *VacanciesNegotiationsVacancyShort) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesNegotiationsVacancyShort) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesNegotiationsVacancyShort) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesNegotiationsVacancyShort) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesNegotiationsVacancyShort) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesNegotiationsVacancyShort) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesNegotiationsVacancyShort) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesNegotiationsVacancyShort) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesNegotiationsVacancyShort) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesNegotiationsVacancyShort) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesNegotiationsVacancyShort) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesNegotiationsVacancyShort) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.


### SetSalaryNil

`func (o *VacanciesNegotiationsVacancyShort) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesNegotiationsVacancyShort) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetShowLogoInSearch

`func (o *VacanciesNegotiationsVacancyShort) GetShowLogoInSearch() bool`

GetShowLogoInSearch returns the ShowLogoInSearch field if non-nil, zero value otherwise.

### GetShowLogoInSearchOk

`func (o *VacanciesNegotiationsVacancyShort) GetShowLogoInSearchOk() (*bool, bool)`

GetShowLogoInSearchOk returns a tuple with the ShowLogoInSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLogoInSearch

`func (o *VacanciesNegotiationsVacancyShort) SetShowLogoInSearch(v bool)`

SetShowLogoInSearch sets ShowLogoInSearch field to given value.

### HasShowLogoInSearch

`func (o *VacanciesNegotiationsVacancyShort) HasShowLogoInSearch() bool`

HasShowLogoInSearch returns a boolean if a field has been set.

### SetShowLogoInSearchNil

`func (o *VacanciesNegotiationsVacancyShort) SetShowLogoInSearchNil(b bool)`

 SetShowLogoInSearchNil sets the value for ShowLogoInSearch to be an explicit nil

### UnsetShowLogoInSearch
`func (o *VacanciesNegotiationsVacancyShort) UnsetShowLogoInSearch()`

UnsetShowLogoInSearch ensures that no value is present for ShowLogoInSearch, not even an explicit nil
### GetSortPointDistance

`func (o *VacanciesNegotiationsVacancyShort) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesNegotiationsVacancyShort) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesNegotiationsVacancyShort) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesNegotiationsVacancyShort) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesNegotiationsVacancyShort) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesNegotiationsVacancyShort) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
### GetType

`func (o *VacanciesNegotiationsVacancyShort) GetType() VacancyTypeOutput`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesNegotiationsVacancyShort) GetTypeOk() (*VacancyTypeOutput, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesNegotiationsVacancyShort) SetType(v VacancyTypeOutput)`

SetType sets Type field to given value.


### GetUrl

`func (o *VacanciesNegotiationsVacancyShort) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesNegotiationsVacancyShort) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesNegotiationsVacancyShort) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


