# VacanciesActiveListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**NullableVacanciesAddress**](VacanciesAddress.md) |  | [optional] 
**AlternateUrl** | **string** | Ссылка на представление вакансии на сайте | 
**ApplyAlternateUrl** | **string** | Ссылка на отклик на вакансию на сайте | 
**Archived** | **bool** | Находится ли данная вакансия в архиве | 
**Area** | [**IncludesArea**](IncludesArea.md) |  | 
**Department** | [**NullableVacancyDepartmentOutput**](VacancyDepartmentOutput.md) |  | 
**Employer** | [**NullableVacanciesEmployerPublic**](VacanciesEmployerPublic.md) |  | 
**HasTest** | **bool** | Информация о наличии прикрепленного тестового задании к вакансии | 
**Id** | **string** | Идентификатор вакансии | 
**Name** | **string** | Название вакансии | 
**Premium** | **bool** | Является ли данная вакансия премиум-вакансией | 
**PublishedAt** | **string** | Дата и время публикации вакансии | 
**Relations** | [**[]VacancyRelationItem**](VacancyRelationItem.md) | Возвращает связи соискателя с вакансией. Значения из поля &#x60;vacancy_relation&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**ResponseLetterRequired** | **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | [**NullableVacancySalary**](VacancySalary.md) |  | 
**ShowLogoInSearch** | Pointer to **NullableBool** | Отображать ли лого для вакансии в поисковой выдаче | [optional] 
**Type** | [**VacancyTypeOutput**](VacancyTypeOutput.md) |  | 
**Url** | **string** | URL вакансии | 
**BillingType** | Pointer to [**NullableVacancyBillingTypeOutput**](VacancyBillingTypeOutput.md) |  | [optional] 
**CanUpgradeBillingType** | **bool** | Можно ли улучшить биллинговый тип вакансии | 
**Counters** | [**VacancyCountersForActive**](VacancyCountersForActive.md) |  | 
**CreatedAt** | **string** | Дата и время публикации вакансии | 
**ExpiresAt** | **string** | Дата и время окончания публикации вакансии | 
**HasUpdates** | **bool** | Есть ли в откликах/приглашениях по данной вакансии обновления, требующие внимания | 
**Manager** | [**VacancyManagerOutput**](VacancyManagerOutput.md) |  | 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций. Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60;  | [optional] 

## Methods

### NewVacanciesActiveListItem

`func NewVacanciesActiveListItem(alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, department NullableVacancyDepartmentOutput, employer NullableVacanciesEmployerPublic, hasTest bool, id string, name string, premium bool, publishedAt string, relations []VacancyRelationItem, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string, canUpgradeBillingType bool, counters VacancyCountersForActive, createdAt string, expiresAt string, hasUpdates bool, manager VacancyManagerOutput, ) *VacanciesActiveListItem`

NewVacanciesActiveListItem instantiates a new VacanciesActiveListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesActiveListItemWithDefaults

`func NewVacanciesActiveListItemWithDefaults() *VacanciesActiveListItem`

NewVacanciesActiveListItemWithDefaults instantiates a new VacanciesActiveListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VacanciesActiveListItem) GetAddress() VacanciesAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesActiveListItem) GetAddressOk() (*VacanciesAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesActiveListItem) SetAddress(v VacanciesAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacanciesActiveListItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacanciesActiveListItem) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacanciesActiveListItem) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAlternateUrl

`func (o *VacanciesActiveListItem) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesActiveListItem) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesActiveListItem) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesActiveListItem) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesActiveListItem) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesActiveListItem) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetArchived

`func (o *VacanciesActiveListItem) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesActiveListItem) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesActiveListItem) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArea

`func (o *VacanciesActiveListItem) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesActiveListItem) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesActiveListItem) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetDepartment

`func (o *VacanciesActiveListItem) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesActiveListItem) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesActiveListItem) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.


### SetDepartmentNil

`func (o *VacanciesActiveListItem) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesActiveListItem) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetEmployer

`func (o *VacanciesActiveListItem) GetEmployer() VacanciesEmployerPublic`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesActiveListItem) GetEmployerOk() (*VacanciesEmployerPublic, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesActiveListItem) SetEmployer(v VacanciesEmployerPublic)`

SetEmployer sets Employer field to given value.


### SetEmployerNil

`func (o *VacanciesActiveListItem) SetEmployerNil(b bool)`

 SetEmployerNil sets the value for Employer to be an explicit nil

### UnsetEmployer
`func (o *VacanciesActiveListItem) UnsetEmployer()`

UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
### GetHasTest

`func (o *VacanciesActiveListItem) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesActiveListItem) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesActiveListItem) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesActiveListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesActiveListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesActiveListItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VacanciesActiveListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesActiveListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesActiveListItem) SetName(v string)`

SetName sets Name field to given value.


### GetPremium

`func (o *VacanciesActiveListItem) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesActiveListItem) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesActiveListItem) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetPublishedAt

`func (o *VacanciesActiveListItem) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesActiveListItem) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesActiveListItem) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetRelations

`func (o *VacanciesActiveListItem) GetRelations() []VacancyRelationItem`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *VacanciesActiveListItem) GetRelationsOk() (*[]VacancyRelationItem, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *VacanciesActiveListItem) SetRelations(v []VacancyRelationItem)`

SetRelations sets Relations field to given value.


### GetResponseLetterRequired

`func (o *VacanciesActiveListItem) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesActiveListItem) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesActiveListItem) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesActiveListItem) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesActiveListItem) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesActiveListItem) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesActiveListItem) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesActiveListItem) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesActiveListItem) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesActiveListItem) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesActiveListItem) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesActiveListItem) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.


### SetSalaryNil

`func (o *VacanciesActiveListItem) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesActiveListItem) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetShowLogoInSearch

`func (o *VacanciesActiveListItem) GetShowLogoInSearch() bool`

GetShowLogoInSearch returns the ShowLogoInSearch field if non-nil, zero value otherwise.

### GetShowLogoInSearchOk

`func (o *VacanciesActiveListItem) GetShowLogoInSearchOk() (*bool, bool)`

GetShowLogoInSearchOk returns a tuple with the ShowLogoInSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLogoInSearch

`func (o *VacanciesActiveListItem) SetShowLogoInSearch(v bool)`

SetShowLogoInSearch sets ShowLogoInSearch field to given value.

### HasShowLogoInSearch

`func (o *VacanciesActiveListItem) HasShowLogoInSearch() bool`

HasShowLogoInSearch returns a boolean if a field has been set.

### SetShowLogoInSearchNil

`func (o *VacanciesActiveListItem) SetShowLogoInSearchNil(b bool)`

 SetShowLogoInSearchNil sets the value for ShowLogoInSearch to be an explicit nil

### UnsetShowLogoInSearch
`func (o *VacanciesActiveListItem) UnsetShowLogoInSearch()`

UnsetShowLogoInSearch ensures that no value is present for ShowLogoInSearch, not even an explicit nil
### GetType

`func (o *VacanciesActiveListItem) GetType() VacancyTypeOutput`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesActiveListItem) GetTypeOk() (*VacancyTypeOutput, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesActiveListItem) SetType(v VacancyTypeOutput)`

SetType sets Type field to given value.


### GetUrl

`func (o *VacanciesActiveListItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesActiveListItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesActiveListItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBillingType

`func (o *VacanciesActiveListItem) GetBillingType() VacancyBillingTypeOutput`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacanciesActiveListItem) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacanciesActiveListItem) SetBillingType(v VacancyBillingTypeOutput)`

SetBillingType sets BillingType field to given value.

### HasBillingType

`func (o *VacanciesActiveListItem) HasBillingType() bool`

HasBillingType returns a boolean if a field has been set.

### SetBillingTypeNil

`func (o *VacanciesActiveListItem) SetBillingTypeNil(b bool)`

 SetBillingTypeNil sets the value for BillingType to be an explicit nil

### UnsetBillingType
`func (o *VacanciesActiveListItem) UnsetBillingType()`

UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
### GetCanUpgradeBillingType

`func (o *VacanciesActiveListItem) GetCanUpgradeBillingType() bool`

GetCanUpgradeBillingType returns the CanUpgradeBillingType field if non-nil, zero value otherwise.

### GetCanUpgradeBillingTypeOk

`func (o *VacanciesActiveListItem) GetCanUpgradeBillingTypeOk() (*bool, bool)`

GetCanUpgradeBillingTypeOk returns a tuple with the CanUpgradeBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpgradeBillingType

`func (o *VacanciesActiveListItem) SetCanUpgradeBillingType(v bool)`

SetCanUpgradeBillingType sets CanUpgradeBillingType field to given value.


### GetCounters

`func (o *VacanciesActiveListItem) GetCounters() VacancyCountersForActive`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *VacanciesActiveListItem) GetCountersOk() (*VacancyCountersForActive, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *VacanciesActiveListItem) SetCounters(v VacancyCountersForActive)`

SetCounters sets Counters field to given value.


### GetCreatedAt

`func (o *VacanciesActiveListItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesActiveListItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesActiveListItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiresAt

`func (o *VacanciesActiveListItem) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VacanciesActiveListItem) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VacanciesActiveListItem) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetHasUpdates

`func (o *VacanciesActiveListItem) GetHasUpdates() bool`

GetHasUpdates returns the HasUpdates field if non-nil, zero value otherwise.

### GetHasUpdatesOk

`func (o *VacanciesActiveListItem) GetHasUpdatesOk() (*bool, bool)`

GetHasUpdatesOk returns a tuple with the HasUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdates

`func (o *VacanciesActiveListItem) SetHasUpdates(v bool)`

SetHasUpdates sets HasUpdates field to given value.


### GetManager

`func (o *VacanciesActiveListItem) GetManager() VacancyManagerOutput`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacanciesActiveListItem) GetManagerOk() (*VacancyManagerOutput, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacanciesActiveListItem) SetManager(v VacancyManagerOutput)`

SetManager sets Manager field to given value.


### GetSortPointDistance

`func (o *VacanciesActiveListItem) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesActiveListItem) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesActiveListItem) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesActiveListItem) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesActiveListItem) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesActiveListItem) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


