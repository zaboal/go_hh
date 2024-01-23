# VacanciesVacanciesFavoritedItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**NullableVacancyAddressRawOutput**](VacancyAddressRawOutput.md) |  | [optional] 
**AdvResponseUrl** | Pointer to **NullableString** | URL для регистрации нажатия кнопки отклика | [optional] 
**AlternateUrl** | **string** | Ссылка на представление вакансии на сайте | 
**ApplyAlternateUrl** | **string** | Ссылка на отклик на вакансию на сайте | 
**Archived** | **bool** | Находится ли данная вакансия в архиве | 
**Area** | [**IncludesArea**](IncludesArea.md) |  | 
**CreatedAt** | Pointer to **string** | Дата и время публикации вакансии | [optional] 
**Department** | [**NullableVacancyDepartmentOutput**](VacancyDepartmentOutput.md) |  | 
**Employer** | [**VacanciesEmployerPublic**](VacanciesEmployerPublic.md) |  | 
**HasTest** | **bool** | Информация о наличии прикрепленного тестового задании к вакансии | 
**Id** | **string** | Идентификатор вакансии | 
**InsiderInterview** | Pointer to [**NullableVacancyInsiderInterview**](VacancyInsiderInterview.md) |  | [optional] 
**Name** | **string** | Название | 
**NegotiationsUrl** | Pointer to **NullableString** | Ссылка для получения списка откликов/приглашений | [optional] 
**Premium** | **bool** | Является ли данная вакансия премиум-вакансией | 
**PublishedAt** | **string** | Дата и время публикации вакансии | 
**Relations** | [**[]VacancyRelationItem**](VacancyRelationItem.md) | Возвращает связи соискателя с вакансией. Значения из поля &#x60;vacancy_relation&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**ResponseLetterRequired** | **bool** | Обязательно ли заполнять сообщение при отклике на вакансию | 
**ResponseUrl** | Pointer to **NullableString** | URL отклика для прямых вакансий (&#x60;type.id&#x3D;direct&#x60;) | [optional] 
**Salary** | [**NullableVacancySalary**](VacancySalary.md) |  | 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций.  Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60; | [optional] 
**SuitableResumesUrl** | Pointer to **NullableString** | Подходящие резюме на вакансию | [optional] 
**Type** | [**VacancyTypeOutput**](VacancyTypeOutput.md) |  | 
**Url** | **string** | URL вакансии | 

## Methods

### NewVacanciesVacanciesFavoritedItem

`func NewVacanciesVacanciesFavoritedItem(alternateUrl string, applyAlternateUrl string, archived bool, area IncludesArea, department NullableVacancyDepartmentOutput, employer VacanciesEmployerPublic, hasTest bool, id string, name string, premium bool, publishedAt string, relations []VacancyRelationItem, responseLetterRequired bool, salary NullableVacancySalary, type_ VacancyTypeOutput, url string, ) *VacanciesVacanciesFavoritedItem`

NewVacanciesVacanciesFavoritedItem instantiates a new VacanciesVacanciesFavoritedItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacanciesFavoritedItemWithDefaults

`func NewVacanciesVacanciesFavoritedItemWithDefaults() *VacanciesVacanciesFavoritedItem`

NewVacanciesVacanciesFavoritedItemWithDefaults instantiates a new VacanciesVacanciesFavoritedItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VacanciesVacanciesFavoritedItem) GetAddress() VacancyAddressRawOutput`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VacanciesVacanciesFavoritedItem) GetAddressOk() (*VacancyAddressRawOutput, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VacanciesVacanciesFavoritedItem) SetAddress(v VacancyAddressRawOutput)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VacanciesVacanciesFavoritedItem) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *VacanciesVacanciesFavoritedItem) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *VacanciesVacanciesFavoritedItem) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAdvResponseUrl

`func (o *VacanciesVacanciesFavoritedItem) GetAdvResponseUrl() string`

GetAdvResponseUrl returns the AdvResponseUrl field if non-nil, zero value otherwise.

### GetAdvResponseUrlOk

`func (o *VacanciesVacanciesFavoritedItem) GetAdvResponseUrlOk() (*string, bool)`

GetAdvResponseUrlOk returns a tuple with the AdvResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvResponseUrl

`func (o *VacanciesVacanciesFavoritedItem) SetAdvResponseUrl(v string)`

SetAdvResponseUrl sets AdvResponseUrl field to given value.

### HasAdvResponseUrl

`func (o *VacanciesVacanciesFavoritedItem) HasAdvResponseUrl() bool`

HasAdvResponseUrl returns a boolean if a field has been set.

### SetAdvResponseUrlNil

`func (o *VacanciesVacanciesFavoritedItem) SetAdvResponseUrlNil(b bool)`

 SetAdvResponseUrlNil sets the value for AdvResponseUrl to be an explicit nil

### UnsetAdvResponseUrl
`func (o *VacanciesVacanciesFavoritedItem) UnsetAdvResponseUrl()`

UnsetAdvResponseUrl ensures that no value is present for AdvResponseUrl, not even an explicit nil
### GetAlternateUrl

`func (o *VacanciesVacanciesFavoritedItem) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVacanciesFavoritedItem) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVacanciesFavoritedItem) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetApplyAlternateUrl

`func (o *VacanciesVacanciesFavoritedItem) GetApplyAlternateUrl() string`

GetApplyAlternateUrl returns the ApplyAlternateUrl field if non-nil, zero value otherwise.

### GetApplyAlternateUrlOk

`func (o *VacanciesVacanciesFavoritedItem) GetApplyAlternateUrlOk() (*string, bool)`

GetApplyAlternateUrlOk returns a tuple with the ApplyAlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyAlternateUrl

`func (o *VacanciesVacanciesFavoritedItem) SetApplyAlternateUrl(v string)`

SetApplyAlternateUrl sets ApplyAlternateUrl field to given value.


### GetArchived

`func (o *VacanciesVacanciesFavoritedItem) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *VacanciesVacanciesFavoritedItem) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *VacanciesVacanciesFavoritedItem) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetArea

`func (o *VacanciesVacanciesFavoritedItem) GetArea() IncludesArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesVacanciesFavoritedItem) GetAreaOk() (*IncludesArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesVacanciesFavoritedItem) SetArea(v IncludesArea)`

SetArea sets Area field to given value.


### GetCreatedAt

`func (o *VacanciesVacanciesFavoritedItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesVacanciesFavoritedItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesVacanciesFavoritedItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VacanciesVacanciesFavoritedItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDepartment

`func (o *VacanciesVacanciesFavoritedItem) GetDepartment() VacancyDepartmentOutput`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *VacanciesVacanciesFavoritedItem) GetDepartmentOk() (*VacancyDepartmentOutput, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *VacanciesVacanciesFavoritedItem) SetDepartment(v VacancyDepartmentOutput)`

SetDepartment sets Department field to given value.


### SetDepartmentNil

`func (o *VacanciesVacanciesFavoritedItem) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *VacanciesVacanciesFavoritedItem) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetEmployer

`func (o *VacanciesVacanciesFavoritedItem) GetEmployer() VacanciesEmployerPublic`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *VacanciesVacanciesFavoritedItem) GetEmployerOk() (*VacanciesEmployerPublic, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *VacanciesVacanciesFavoritedItem) SetEmployer(v VacanciesEmployerPublic)`

SetEmployer sets Employer field to given value.


### GetHasTest

`func (o *VacanciesVacanciesFavoritedItem) GetHasTest() bool`

GetHasTest returns the HasTest field if non-nil, zero value otherwise.

### GetHasTestOk

`func (o *VacanciesVacanciesFavoritedItem) GetHasTestOk() (*bool, bool)`

GetHasTestOk returns a tuple with the HasTest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasTest

`func (o *VacanciesVacanciesFavoritedItem) SetHasTest(v bool)`

SetHasTest sets HasTest field to given value.


### GetId

`func (o *VacanciesVacanciesFavoritedItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVacanciesFavoritedItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVacanciesFavoritedItem) SetId(v string)`

SetId sets Id field to given value.


### GetInsiderInterview

`func (o *VacanciesVacanciesFavoritedItem) GetInsiderInterview() VacancyInsiderInterview`

GetInsiderInterview returns the InsiderInterview field if non-nil, zero value otherwise.

### GetInsiderInterviewOk

`func (o *VacanciesVacanciesFavoritedItem) GetInsiderInterviewOk() (*VacancyInsiderInterview, bool)`

GetInsiderInterviewOk returns a tuple with the InsiderInterview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsiderInterview

`func (o *VacanciesVacanciesFavoritedItem) SetInsiderInterview(v VacancyInsiderInterview)`

SetInsiderInterview sets InsiderInterview field to given value.

### HasInsiderInterview

`func (o *VacanciesVacanciesFavoritedItem) HasInsiderInterview() bool`

HasInsiderInterview returns a boolean if a field has been set.

### SetInsiderInterviewNil

`func (o *VacanciesVacanciesFavoritedItem) SetInsiderInterviewNil(b bool)`

 SetInsiderInterviewNil sets the value for InsiderInterview to be an explicit nil

### UnsetInsiderInterview
`func (o *VacanciesVacanciesFavoritedItem) UnsetInsiderInterview()`

UnsetInsiderInterview ensures that no value is present for InsiderInterview, not even an explicit nil
### GetName

`func (o *VacanciesVacanciesFavoritedItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesVacanciesFavoritedItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesVacanciesFavoritedItem) SetName(v string)`

SetName sets Name field to given value.


### GetNegotiationsUrl

`func (o *VacanciesVacanciesFavoritedItem) GetNegotiationsUrl() string`

GetNegotiationsUrl returns the NegotiationsUrl field if non-nil, zero value otherwise.

### GetNegotiationsUrlOk

`func (o *VacanciesVacanciesFavoritedItem) GetNegotiationsUrlOk() (*string, bool)`

GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsUrl

`func (o *VacanciesVacanciesFavoritedItem) SetNegotiationsUrl(v string)`

SetNegotiationsUrl sets NegotiationsUrl field to given value.

### HasNegotiationsUrl

`func (o *VacanciesVacanciesFavoritedItem) HasNegotiationsUrl() bool`

HasNegotiationsUrl returns a boolean if a field has been set.

### SetNegotiationsUrlNil

`func (o *VacanciesVacanciesFavoritedItem) SetNegotiationsUrlNil(b bool)`

 SetNegotiationsUrlNil sets the value for NegotiationsUrl to be an explicit nil

### UnsetNegotiationsUrl
`func (o *VacanciesVacanciesFavoritedItem) UnsetNegotiationsUrl()`

UnsetNegotiationsUrl ensures that no value is present for NegotiationsUrl, not even an explicit nil
### GetPremium

`func (o *VacanciesVacanciesFavoritedItem) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *VacanciesVacanciesFavoritedItem) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *VacanciesVacanciesFavoritedItem) SetPremium(v bool)`

SetPremium sets Premium field to given value.


### GetPublishedAt

`func (o *VacanciesVacanciesFavoritedItem) GetPublishedAt() string`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VacanciesVacanciesFavoritedItem) GetPublishedAtOk() (*string, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VacanciesVacanciesFavoritedItem) SetPublishedAt(v string)`

SetPublishedAt sets PublishedAt field to given value.


### GetRelations

`func (o *VacanciesVacanciesFavoritedItem) GetRelations() []VacancyRelationItem`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *VacanciesVacanciesFavoritedItem) GetRelationsOk() (*[]VacancyRelationItem, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *VacanciesVacanciesFavoritedItem) SetRelations(v []VacancyRelationItem)`

SetRelations sets Relations field to given value.


### GetResponseLetterRequired

`func (o *VacanciesVacanciesFavoritedItem) GetResponseLetterRequired() bool`

GetResponseLetterRequired returns the ResponseLetterRequired field if non-nil, zero value otherwise.

### GetResponseLetterRequiredOk

`func (o *VacanciesVacanciesFavoritedItem) GetResponseLetterRequiredOk() (*bool, bool)`

GetResponseLetterRequiredOk returns a tuple with the ResponseLetterRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseLetterRequired

`func (o *VacanciesVacanciesFavoritedItem) SetResponseLetterRequired(v bool)`

SetResponseLetterRequired sets ResponseLetterRequired field to given value.


### GetResponseUrl

`func (o *VacanciesVacanciesFavoritedItem) GetResponseUrl() string`

GetResponseUrl returns the ResponseUrl field if non-nil, zero value otherwise.

### GetResponseUrlOk

`func (o *VacanciesVacanciesFavoritedItem) GetResponseUrlOk() (*string, bool)`

GetResponseUrlOk returns a tuple with the ResponseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseUrl

`func (o *VacanciesVacanciesFavoritedItem) SetResponseUrl(v string)`

SetResponseUrl sets ResponseUrl field to given value.

### HasResponseUrl

`func (o *VacanciesVacanciesFavoritedItem) HasResponseUrl() bool`

HasResponseUrl returns a boolean if a field has been set.

### SetResponseUrlNil

`func (o *VacanciesVacanciesFavoritedItem) SetResponseUrlNil(b bool)`

 SetResponseUrlNil sets the value for ResponseUrl to be an explicit nil

### UnsetResponseUrl
`func (o *VacanciesVacanciesFavoritedItem) UnsetResponseUrl()`

UnsetResponseUrl ensures that no value is present for ResponseUrl, not even an explicit nil
### GetSalary

`func (o *VacanciesVacanciesFavoritedItem) GetSalary() VacancySalary`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesVacanciesFavoritedItem) GetSalaryOk() (*VacancySalary, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesVacanciesFavoritedItem) SetSalary(v VacancySalary)`

SetSalary sets Salary field to given value.


### SetSalaryNil

`func (o *VacanciesVacanciesFavoritedItem) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesVacanciesFavoritedItem) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSortPointDistance

`func (o *VacanciesVacanciesFavoritedItem) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesVacanciesFavoritedItem) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesVacanciesFavoritedItem) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesVacanciesFavoritedItem) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesVacanciesFavoritedItem) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesVacanciesFavoritedItem) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil
### GetSuitableResumesUrl

`func (o *VacanciesVacanciesFavoritedItem) GetSuitableResumesUrl() string`

GetSuitableResumesUrl returns the SuitableResumesUrl field if non-nil, zero value otherwise.

### GetSuitableResumesUrlOk

`func (o *VacanciesVacanciesFavoritedItem) GetSuitableResumesUrlOk() (*string, bool)`

GetSuitableResumesUrlOk returns a tuple with the SuitableResumesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitableResumesUrl

`func (o *VacanciesVacanciesFavoritedItem) SetSuitableResumesUrl(v string)`

SetSuitableResumesUrl sets SuitableResumesUrl field to given value.

### HasSuitableResumesUrl

`func (o *VacanciesVacanciesFavoritedItem) HasSuitableResumesUrl() bool`

HasSuitableResumesUrl returns a boolean if a field has been set.

### SetSuitableResumesUrlNil

`func (o *VacanciesVacanciesFavoritedItem) SetSuitableResumesUrlNil(b bool)`

 SetSuitableResumesUrlNil sets the value for SuitableResumesUrl to be an explicit nil

### UnsetSuitableResumesUrl
`func (o *VacanciesVacanciesFavoritedItem) UnsetSuitableResumesUrl()`

UnsetSuitableResumesUrl ensures that no value is present for SuitableResumesUrl, not even an explicit nil
### GetType

`func (o *VacanciesVacanciesFavoritedItem) GetType() VacancyTypeOutput`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesVacanciesFavoritedItem) GetTypeOk() (*VacancyTypeOutput, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesVacanciesFavoritedItem) SetType(v VacancyTypeOutput)`

SetType sets Type field to given value.


### GetUrl

`func (o *VacanciesVacanciesFavoritedItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesVacanciesFavoritedItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesVacanciesFavoritedItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


