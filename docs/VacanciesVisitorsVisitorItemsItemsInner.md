# VacanciesVisitorsVisitorItemsItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateUrl** | **string** | URL резюме на сайте | 
**Id** | **string** | Идентификатор резюме | 
**Title** | **NullableString** | Желаемая должность | 
**Age** | Pointer to **NullableFloat32** | Возраст | [optional] 
**Area** | Pointer to [**NullableIncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**CanViewFullInfo** | Pointer to **NullableBool** | Доступен ли просмотр контактной информации в резюме текущему работодателю | [optional] 
**Certificate** | [**[]ResumeObjectsCertificate**](ResumeObjectsCertificate.md) | Список сертификатов соискателя | 
**CreatedAt** | **string** | Дата и время создания резюме | 
**Download** | **map[string]interface{}** | Ссылки для скачивания резюме в разных форматах | 
**Education** | **map[string]interface{}** | Образование соискателя.   Особенности сохранения образования:  * Если передать и высшее и среднее образование и уровень образования \&quot;средний\&quot;, то сохранится только среднее образование. * Если передать и высшее и среднее образование и уровень образования \&quot;высшее\&quot;, то сохранится только высшее образование  | 
**FirstName** | Pointer to **NullableString** | Имя | [optional] 
**Gender** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**HiddenFields** | [**[]IncludesIdName**](IncludesIdName.md) | Справочник [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле &#x60;resume_hidden_fields&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**LastName** | Pointer to **NullableString** | Фамилия | [optional] 
**Marked** | Pointer to **bool** | Выделено ли резюме в поиске | [optional] 
**MiddleName** | Pointer to **NullableString** | Отчество | [optional] 
**Platform** | Pointer to **map[string]interface{}** | Ресурс, на котором было размещено резюме | [optional] 
**Salary** | Pointer to [**NullableResumeObjectsSalaryProperties**](ResumeObjectsSalaryProperties.md) |  | [optional] 
**TotalExperience** | Pointer to [**NullableResumeObjectsTotalExperience**](ResumeObjectsTotalExperience.md) |  | [optional] 
**UpdatedAt** | **string** | Дата и время обновления резюме | 
**Experience** | [**[]ResumeObjectsExperienceShort**](ResumeObjectsExperienceShort.md) | Опыт работы. В объекте опыта отсутствует описание (поле description), а также должность (поле position) доступна только в последнем опыте | 
**Actions** | [**ResumeObjectsActions**](ResumeObjectsActions.md) | Дополнительные действия | 
**Favorited** | **bool** | Добавлено ли резюме в избранные | 
**NegotiationsHistory** | [**ResumeObjectsNegotiationsHistoryUrl**](ResumeObjectsNegotiationsHistoryUrl.md) | Краткая история откликов/приглашений по резюме | 
**Owner** | [**ResumeObjectsOwner**](ResumeObjectsOwner.md) | Информация о владельце резюме | 
**Photo** | Pointer to [**NullableResumeObjectsPhoto**](ResumeObjectsPhoto.md) |  | [optional] 
**Tags** | Pointer to [**[]IncludesId**](IncludesId.md) | Теги к резюме | [optional] 
**Viewed** | **bool** | Было ли резюме уже просмотрено работодателем | 
**Url** | **string** | Ссылка на получение элементов | 

## Methods

### NewVacanciesVisitorsVisitorItemsItemsInner

`func NewVacanciesVisitorsVisitorItemsItemsInner(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download map[string]interface{}, education map[string]interface{}, hiddenFields []IncludesIdName, updatedAt string, experience []ResumeObjectsExperienceShort, actions ResumeObjectsActions, favorited bool, negotiationsHistory ResumeObjectsNegotiationsHistoryUrl, owner ResumeObjectsOwner, viewed bool, url string, ) *VacanciesVisitorsVisitorItemsItemsInner`

NewVacanciesVisitorsVisitorItemsItemsInner instantiates a new VacanciesVisitorsVisitorItemsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVisitorsVisitorItemsItemsInnerWithDefaults

`func NewVacanciesVisitorsVisitorItemsItemsInnerWithDefaults() *VacanciesVisitorsVisitorItemsItemsInner`

NewVacanciesVisitorsVisitorItemsItemsInnerWithDefaults instantiates a new VacanciesVisitorsVisitorItemsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAge

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetArea

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCanViewFullInfo

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetDownload() map[string]interface{}`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetDownloadOk() (*map[string]interface{}, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetDownload(v map[string]interface{})`

SetDownload sets Download field to given value.


### GetEducation

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetEducation() map[string]interface{}`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetEducationOk() (*map[string]interface{}, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetEducation(v map[string]interface{})`

SetEducation sets Education field to given value.


### GetFirstName

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetLastName

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetMiddleName

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPlatform

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetPlatform() map[string]interface{}`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetPlatformOk() (*map[string]interface{}, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetPlatform(v map[string]interface{})`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTotalExperience

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetExperience

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetExperience() []ResumeObjectsExperienceShort`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetExperienceOk() (*[]ResumeObjectsExperienceShort, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetExperience(v []ResumeObjectsExperienceShort)`

SetExperience sets Experience field to given value.


### GetActions

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetActions() ResumeObjectsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetActionsOk() (*ResumeObjectsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetActions(v ResumeObjectsActions)`

SetActions sets Actions field to given value.


### GetFavorited

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetNegotiationsHistory

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetNegotiationsHistory() ResumeObjectsNegotiationsHistoryUrl`

GetNegotiationsHistory returns the NegotiationsHistory field if non-nil, zero value otherwise.

### GetNegotiationsHistoryOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetNegotiationsHistoryOk() (*ResumeObjectsNegotiationsHistoryUrl, bool)`

GetNegotiationsHistoryOk returns a tuple with the NegotiationsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsHistory

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetNegotiationsHistory(v ResumeObjectsNegotiationsHistoryUrl)`

SetNegotiationsHistory sets NegotiationsHistory field to given value.


### GetOwner

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetOwner() ResumeObjectsOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetOwnerOk() (*ResumeObjectsOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetOwner(v ResumeObjectsOwner)`

SetOwner sets Owner field to given value.


### GetPhoto

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *VacanciesVisitorsVisitorItemsItemsInner) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetTags

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTags() []IncludesId`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetTagsOk() (*[]IncludesId, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetTags(v []IncludesId)`

SetTags sets Tags field to given value.

### HasTags

`func (o *VacanciesVisitorsVisitorItemsItemsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetViewed

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetViewed(v bool)`

SetViewed sets Viewed field to given value.


### GetUrl

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesVisitorsVisitorItemsItemsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesVisitorsVisitorItemsItemsInner) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


