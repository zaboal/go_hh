# VacanciesMatchListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**VacanciesAddress**](VacanciesAddress.md) |  | [optional] 
**AlternateUrl** | **string** | Ссылка на представление вакансии на сайте | 
**ApplyAlternateUrl** | **string** | Ссылка на отклик на вакансию на сайте | 
**Archived** | **bool** | Находится ли данная вакансия в архиве | 
**Area** | [**IncludesArea**](IncludesArea.md) |  | 
**Department** | [**NullableVacancyDepartmentOutput**](VacancyDepartmentOutput.md) |  | 
**Employer** | [**VacanciesEmployerPublic**](VacanciesEmployerPublic.md) |  | 
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
**CanInvite** | **bool** | Можно ли пригласить соискателя на данную вакансию | 
**CreatedAt** | **string** | Дата и время публикации вакансии | 
**EmployerNegotiationsState** | [**IncludesIdName**](IncludesIdName.md) |  | 
**Manager** | [**VacancyManagerOutput**](VacancyManagerOutput.md) |  | 
**NegotiationsActions** | [**[]VacancyNegotiationActions**](VacancyNegotiationActions.md) | Действия для [создания отклика](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#add-invite). Если создать отклик невозможно (например, нет нужных услуг), то вернется пустой массив | 
**NegotiationsState** | [**IncludesIdName**](IncludesIdName.md) |  | 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций. Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60;  | [optional] 
**Templates** | Pointer to [**[]VacancyTemplates**](VacancyTemplates.md) | Шаблоны писем | [optional] 

## Methods

### NewVacanciesMatchListItem

`func NewVacanciesMatchListItem(alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, department NullableVacancyDepartmentOutput, employer VacanciesEmployerPublic, hasTest bool, id string, name string, premium bool, publishedAt string, relations []VacancyRelationItem, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string, canInvite bool, createdAt string, employerNegotiationsState IncludesIdName, manager VacancyManagerOutput, negotiationsActions []VacancyNegotiationActions, negotiationsState IncludesIdName, ) *VacanciesMatchListItem`

NewVacanciesMatchListItem instantiates a new VacanciesMatchListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesMatchListItemWithDefaults

`func NewVacanciesMatchListItemWithDefaults() *VacanciesMatchListItem`

NewVacanciesMatchListItemWithDefaults instantiates a new VacanciesMatchListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VacanciesMatchListItem) GetAddress() VacanciesAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesMatchListItem) GetAddressOk() (*VacanciesAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesMatchListItem) SetAddress(v VacanciesAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacanciesMatchListItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAlternateUrl

`func (o *VacanciesMatchListItem) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesMatchListItem) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesMatchListItem) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesMatchListItem) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesMatchListItem) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesMatchListItem) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetArchived

`func (o *VacanciesMatchListItem) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesMatchListItem) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesMatchListItem) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArea

`func (o *VacanciesMatchListItem) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesMatchListItem) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesMatchListItem) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetDepartment

`func (o *VacanciesMatchListItem) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesMatchListItem) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesMatchListItem) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.


### SetDepartmentNil

`func (o *VacanciesMatchListItem) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesMatchListItem) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetEmployer

`func (o *VacanciesMatchListItem) GetEmployer() VacanciesEmployerPublic`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesMatchListItem) GetEmployerOk() (*VacanciesEmployerPublic, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesMatchListItem) SetEmployer(v VacanciesEmployerPublic)`

SetEmployer sets Employer field to given value.


### GetHasTest

`func (o *VacanciesMatchListItem) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesMatchListItem) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesMatchListItem) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesMatchListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesMatchListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesMatchListItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VacanciesMatchListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesMatchListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesMatchListItem) SetName(v string)`

SetName sets Name field to given value.


### GetPremium

`func (o *VacanciesMatchListItem) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesMatchListItem) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesMatchListItem) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetPublishedAt

`func (o *VacanciesMatchListItem) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesMatchListItem) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesMatchListItem) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetRelations

`func (o *VacanciesMatchListItem) GetRelations() []VacancyRelationItem`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *VacanciesMatchListItem) GetRelationsOk() (*[]VacancyRelationItem, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *VacanciesMatchListItem) SetRelations(v []VacancyRelationItem)`

SetRelations sets Relations field to given value.


### GetResponseLetterRequired

`func (o *VacanciesMatchListItem) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesMatchListItem) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesMatchListItem) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesMatchListItem) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesMatchListItem) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesMatchListItem) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesMatchListItem) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesMatchListItem) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesMatchListItem) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesMatchListItem) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesMatchListItem) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesMatchListItem) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.


### SetSalaryNil

`func (o *VacanciesMatchListItem) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesMatchListItem) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetShowLogoInSearch

`func (o *VacanciesMatchListItem) GetShowLogoInSearch() bool`

GetShowLogoInSearch returns the ShowLogoInSearch field if non-nil, zero value otherwise.

### GetShowLogoInSearchOk

`func (o *VacanciesMatchListItem) GetShowLogoInSearchOk() (*bool, bool)`

GetShowLogoInSearchOk returns a tuple with the ShowLogoInSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowLogoInSearch

`func (o *VacanciesMatchListItem) SetShowLogoInSearch(v bool)`

SetShowLogoInSearch sets ShowLogoInSearch field to given value.

### HasShowLogoInSearch

`func (o *VacanciesMatchListItem) HasShowLogoInSearch() bool`

HasShowLogoInSearch returns a boolean if a field has been set.

### SetShowLogoInSearchNil

`func (o *VacanciesMatchListItem) SetShowLogoInSearchNil(b bool)`

 SetShowLogoInSearchNil sets the value for ShowLogoInSearch to be an explicit nil

### UnsetShowLogoInSearch
`func (o *VacanciesMatchListItem) UnsetShowLogoInSearch()`

UnsetShowLogoInSearch ensures that no value is present for ShowLogoInSearch, not even an explicit nil
### GetType

`func (o *VacanciesMatchListItem) GetType() VacancyTypeOutput`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesMatchListItem) GetTypeOk() (*VacancyTypeOutput, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesMatchListItem) SetType(v VacancyTypeOutput)`

SetType sets Type field to given value.


### GetUrl

`func (o *VacanciesMatchListItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesMatchListItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesMatchListItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetCanInvite

`func (o *VacanciesMatchListItem) GetCanInvite() bool`

GetCanInvite returns the CanInvite field if non-nil, zero value otherwise.

### GetCanInviteOk

`func (o *VacanciesMatchListItem) GetCanInviteOk() (*bool, bool)`

GetCanInviteOk returns a tuple with the CanInvite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanInvite

`func (o *VacanciesMatchListItem) SetCanInvite(v bool)`

SetCanInvite sets CanInvite field to given value.


### GetCreatedAt

`func (o *VacanciesMatchListItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesMatchListItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesMatchListItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmployerNegotiationsState

`func (o *VacanciesMatchListItem) GetEmployerNegotiationsState() IncludesIdName`

GetEmployerNegotiationsState returns the EmployerNegotiationsState field if non-nil, zero value otherwise.

### GetEmployerNegotiationsStateOk

`func (o *VacanciesMatchListItem) GetEmployerNegotiationsStateOk() (*IncludesIdName, bool)`

GetEmployerNegotiationsStateOk returns a tuple with the EmployerNegotiationsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerNegotiationsState

`func (o *VacanciesMatchListItem) SetEmployerNegotiationsState(v IncludesIdName)`

SetEmployerNegotiationsState sets EmployerNegotiationsState field to given value.


### GetManager

`func (o *VacanciesMatchListItem) GetManager() VacancyManagerOutput`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *VacanciesMatchListItem) GetManagerOk() (*VacancyManagerOutput, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *VacanciesMatchListItem) SetManager(v VacancyManagerOutput)`

SetManager sets Manager field to given value.


### GetNegotiationsActions

`func (o *VacanciesMatchListItem) GetNegotiationsActions() []VacancyNegotiationActions`

GetNegotiationsActions returns the NegotiationsActions field if non-nil, zero value otherwise.

### GetNegotiationsActionsOk

`func (o *VacanciesMatchListItem) GetNegotiationsActionsOk() (*[]VacancyNegotiationActions, bool)`

GetNegotiationsActionsOk returns a tuple with the NegotiationsActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsActions

`func (o *VacanciesMatchListItem) SetNegotiationsActions(v []VacancyNegotiationActions)`

SetNegotiationsActions sets NegotiationsActions field to given value.


### GetNegotiationsState

`func (o *VacanciesMatchListItem) GetNegotiationsState() IncludesIdName`

GetNegotiationsState returns the NegotiationsState field if non-nil, zero value otherwise.

### GetNegotiationsStateOk

`func (o *VacanciesMatchListItem) GetNegotiationsStateOk() (*IncludesIdName, bool)`

GetNegotiationsStateOk returns a tuple with the NegotiationsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsState

`func (o *VacanciesMatchListItem) SetNegotiationsState(v IncludesIdName)`

SetNegotiationsState sets NegotiationsState field to given value.


### GetSortPointDistance

`func (o *VacanciesMatchListItem) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesMatchListItem) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesMatchListItem) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesMatchListItem) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesMatchListItem) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesMatchListItem) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
### GetTemplates

`func (o *VacanciesMatchListItem) GetTemplates() []VacancyTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *VacanciesMatchListItem) GetTemplatesOk() (*[]VacancyTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *VacanciesMatchListItem) SetTemplates(v []VacancyTemplates)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *VacanciesMatchListItem) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


