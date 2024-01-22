# VacanciesVacanciesBlacklistedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**NullableVacancyAddress**](VacancyAddress.md) |  | [optional] 
**AdvResponseUrl** | Pointer to **string** | URL для регистрации нажатия кнопки отклика | [optional] 
**AlternateUrl** | **string** | Ссылка на представление вакансии на сайте | 
**ApplyAlternateUrl** | **string** | Ссылка на отклик на вакансию на сайте | 
**Archived** | **bool** | Находится ли данная вакансия в архиве | 
**Area** | [**IncludesArea**](IncludesArea.md) |  | 
**CreatedAt** | Pointer to **string** | Дата и время публикации вакансии | [optional] 
**Department** | [**NullableVacancyDepartment**](VacancyDepartment.md) |  | 
**Employer** | [**VacanciesEmployerPublic**](VacanciesEmployerPublic.md) |  | 
**HasTest** | **bool** | Информация о наличии прикрепленного тестового задании к вакансии | 
**Id** | **string** | Идентификатор вакансии | 
**InsiderInterview** | Pointer to [**NullableVacancyInsiderInterview**](VacancyInsiderInterview.md) |  | [optional] 
**Name** | **string** | Название | 
**Premium** | **bool** | Является ли данная вакансия премиум-вакансией | 
**PublishedAt** | **string** | Дата и время публикации вакансии | 
**Relations** | [**[]VacancyRelationItem**](VacancyRelationItem.md) | Возвращает связи соискателя с вакансией. Значения из поля &#x60;vacancy_relation&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**ResponseLetterRequired** | **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | [**NullableVacancySalary**](VacancySalary.md) |  | 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций.  Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60; | [optional] 
**Type** | [**VacancyTypeOutput**](VacancyTypeOutput.md) |  | 
**Url** | **string** | URL вакансии | 

## Methods

### NewVacanciesVacanciesBlacklistedItem

`func NewVacanciesVacanciesBlacklistedItem(alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, department NullableVacancyDepartment, employer VacanciesEmployerPublic, hasTest bool, id string, name string, premium bool, publishedAt string, relations []VacancyRelationItem, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string, ) *VacanciesVacanciesBlacklistedItem`

NewVacanciesVacanciesBlacklistedItem instantiates a new VacanciesVacanciesBlacklistedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacanciesBlacklistedItemWithDefaults

`func NewVacanciesVacanciesBlacklistedItemWithDefaults() *VacanciesVacanciesBlacklistedItem`

NewVacanciesVacanciesBlacklistedItemWithDefaults instantiates a new VacanciesVacanciesBlacklistedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VacanciesVacanciesBlacklistedItem) GetAddress() VacancyAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesVacanciesBlacklistedItem) GetAddressOk() (*VacancyAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesVacanciesBlacklistedItem) SetAddress(v VacancyAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacanciesVacanciesBlacklistedItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacanciesVacanciesBlacklistedItem) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacanciesVacanciesBlacklistedItem) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAdvResponseUrl

`func (o *VacanciesVacanciesBlacklistedItem) GetAdvResponseUrl() string`

GetAdvResponseUrl returns the AdvResponseUrl field if non-nil, zero value otherwise.

### GetAdvResponseUrlOk

`func (o *VacanciesVacanciesBlacklistedItem) GetAdvResponseUrlOk() (*string, bool)`

GetAdvResponseUrlOk returns a tuple with the AdvResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvResponseUrl

`func (o *VacanciesVacanciesBlacklistedItem) SetAdvResponseUrl(v string)`

SetAdvResponseUrl sets AdvResponseUrl field to given value.

### HasAdvResponseUrl

`func (o *VacanciesVacanciesBlacklistedItem) HasAdvResponseUrl() bool`

HasAdvResponseUrl returns a boolean if a field has been set.

### GetAlternateUrl

`func (o *VacanciesVacanciesBlacklistedItem) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVacanciesBlacklistedItem) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVacanciesBlacklistedItem) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesVacanciesBlacklistedItem) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesVacanciesBlacklistedItem) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesVacanciesBlacklistedItem) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetArchived

`func (o *VacanciesVacanciesBlacklistedItem) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesVacanciesBlacklistedItem) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesVacanciesBlacklistedItem) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArea

`func (o *VacanciesVacanciesBlacklistedItem) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesVacanciesBlacklistedItem) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesVacanciesBlacklistedItem) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetCreatedAt

`func (o *VacanciesVacanciesBlacklistedItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesVacanciesBlacklistedItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesVacanciesBlacklistedItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VacanciesVacanciesBlacklistedItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepartment

`func (o *VacanciesVacanciesBlacklistedItem) GetDepartment() VacancyDepartment`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesVacanciesBlacklistedItem) GetDepartmentOk() (*VacancyDepartment, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesVacanciesBlacklistedItem) SetDepartment(v VacancyDepartment)`

SetDepartment sets Department field to given value.


### SetDepartmentNil

`func (o *VacanciesVacanciesBlacklistedItem) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesVacanciesBlacklistedItem) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetEmployer

`func (o *VacanciesVacanciesBlacklistedItem) GetEmployer() VacanciesEmployerPublic`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesVacanciesBlacklistedItem) GetEmployerOk() (*VacanciesEmployerPublic, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesVacanciesBlacklistedItem) SetEmployer(v VacanciesEmployerPublic)`

SetEmployer sets Employer field to given value.


### GetHasTest

`func (o *VacanciesVacanciesBlacklistedItem) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesVacanciesBlacklistedItem) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesVacanciesBlacklistedItem) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesVacanciesBlacklistedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVacanciesBlacklistedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVacanciesBlacklistedItem) SetId(v string)`

SetId sets Id field to given value.


### GetInsiderInterview

`func (o *VacanciesVacanciesBlacklistedItem) GetInsiderInterview() VacancyInsiderInterview`

GetInsiderInterview returns the InsiderInterview field if non-nil, zero value otherwise.

### GetInsiderInterviewOk

`func (o *VacanciesVacanciesBlacklistedItem) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool)`

GetInsiderInterviewOk returns a tuple with the InsiderInterview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderInterview

`func (o *VacanciesVacanciesBlacklistedItem) SetInsiderInterview(v VacancyInsiderInterview)`

SetInsiderInterview sets InsiderInterview field to given value.

### HasInsiderInterview

`func (o *VacanciesVacanciesBlacklistedItem) HasInsiderInterview() bool`

HasInsiderInterview returns a boolean if a field has been set.

### SetInsiderInterviewNil

`func (o *VacanciesVacanciesBlacklistedItem) SetInsiderInterviewNil(b bool)`

 SetInsiderInterviewNil sets the value for InsiderInterview to be an explicit nil

### UnsetInsiderInterview
`func (o *VacanciesVacanciesBlacklistedItem) UnsetInsiderInterview()`

UnsetInsiderInterview ensures that no value is present for InsiderInterview, not even an explicit nil
### GetName

`func (o *VacanciesVacanciesBlacklistedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesVacanciesBlacklistedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesVacanciesBlacklistedItem) SetName(v string)`

SetName sets Name field to given value.


### GetPremium

`func (o *VacanciesVacanciesBlacklistedItem) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesVacanciesBlacklistedItem) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesVacanciesBlacklistedItem) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetPublishedAt

`func (o *VacanciesVacanciesBlacklistedItem) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesVacanciesBlacklistedItem) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesVacanciesBlacklistedItem) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetRelations

`func (o *VacanciesVacanciesBlacklistedItem) GetRelations() []VacancyRelationItem`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *VacanciesVacanciesBlacklistedItem) GetRelationsOk() (*[]VacancyRelationItem, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *VacanciesVacanciesBlacklistedItem) SetRelations(v []VacancyRelationItem)`

SetRelations sets Relations field to given value.


### GetResponseLetterRequired

`func (o *VacanciesVacanciesBlacklistedItem) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesVacanciesBlacklistedItem) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesVacanciesBlacklistedItem) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesVacanciesBlacklistedItem) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesVacanciesBlacklistedItem) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesVacanciesBlacklistedItem) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesVacanciesBlacklistedItem) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesVacanciesBlacklistedItem) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesVacanciesBlacklistedItem) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesVacanciesBlacklistedItem) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesVacanciesBlacklistedItem) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesVacanciesBlacklistedItem) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.


### SetSalaryNil

`func (o *VacanciesVacanciesBlacklistedItem) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesVacanciesBlacklistedItem) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSortPointDistance

`func (o *VacanciesVacanciesBlacklistedItem) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesVacanciesBlacklistedItem) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesVacanciesBlacklistedItem) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesVacanciesBlacklistedItem) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesVacanciesBlacklistedItem) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesVacanciesBlacklistedItem) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
### GetType

`func (o *VacanciesVacanciesBlacklistedItem) GetType() VacancyTypeOutput`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesVacanciesBlacklistedItem) GetTypeOk() (*VacancyTypeOutput, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesVacanciesBlacklistedItem) SetType(v VacancyTypeOutput)`

SetType sets Type field to given value.


### GetUrl

`func (o *VacanciesVacanciesBlacklistedItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesVacanciesBlacklistedItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesVacanciesBlacklistedItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


