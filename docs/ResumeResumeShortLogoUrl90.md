# ResumeResumeShortLogoUrl90

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateUrl** | **string** | URL резюме на сайте | 
**Id** | **string** | Идентификатор резюме | 
**Title** | **NullableString** | Желаемая должность | 
**Url** | **string** | Ссылка на получение элементов | 
**Age** | Pointer to **NullableFloat32** | Возраст | [optional] 
**Area** | Pointer to [**NullableIncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**CanViewFullInfo** | Pointer to **NullableBool** | Доступен ли просмотр контактной информации в резюме текущему работодателю | [optional] 
**Certificate** | [**[]ResumeObjectsCertificate**](ResumeObjectsCertificate.md) | Список сертификатов соискателя | 
**CreatedAt** | **string** | Дата и время создания резюме | 
**Download** | **map[string]interface{}** | Ссылки для скачивания резюме в разных форматах | 
**Education** | **map[string]interface{}** | Образование соискателя.   Особенности сохранения образования:  * Если передать и высшее и среднее образование и уровень образования \&quot;средний\&quot;, то сохранится только среднее образование. * Если передать и высшее и среднее образование и уровень образования \&quot;высшее\&quot;, то сохранится только высшее образование  | 
**Experience** | [**[]ResumeObjectsExperienceLogoUrl90**](ResumeObjectsExperienceLogoUrl90.md) | Опыт работы | 
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
**Actions** | [**ResumeObjectsActions**](ResumeObjectsActions.md) | Дополнительные действия | 
**Favorited** | **bool** | Добавлено ли резюме в избранные | 
**NegotiationsHistory** | [**ResumeObjectsNegotiationsHistoryUrl**](ResumeObjectsNegotiationsHistoryUrl.md) | Краткая история откликов/приглашений по резюме | 
**Owner** | [**ResumeObjectsOwner**](ResumeObjectsOwner.md) | Информация о владельце резюме | 
**Photo** | Pointer to [**NullableResumeObjectsPhoto**](ResumeObjectsPhoto.md) |  | [optional] 
**Tags** | Pointer to [**[]IncludesId**](IncludesId.md) | Теги к резюме | [optional] 
**Viewed** | **bool** | Было ли резюме уже просмотрено работодателем | 
**JobSearchStatus** | Pointer to [**IncludesIdNameLastChangeTime**](IncludesIdNameLastChangeTime.md) | Для получения данных нужно передать параметр &#x60;with_job_search_status&#x3D;true&#x60;.  Возможные значения перечислены в поле &#x60;job_search_statuses_employer&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | [optional] 

## Methods

### NewResumeResumeShortLogoUrl90

`func NewResumeResumeShortLogoUrl90(alternateUrl string, id string, title NullableString, url string, certificate []ResumeObjectsCertificate, createdAt string, download map[string]interface{}, education map[string]interface{}, experience []ResumeObjectsExperienceLogoUrl90, hiddenFields []IncludesIdName, updatedAt string, actions ResumeObjectsActions, favorited bool, negotiationsHistory ResumeObjectsNegotiationsHistoryUrl, owner ResumeObjectsOwner, viewed bool, ) *ResumeResumeShortLogoUrl90`

NewResumeResumeShortLogoUrl90 instantiates a new ResumeResumeShortLogoUrl90 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeShortLogoUrl90WithDefaults

`func NewResumeResumeShortLogoUrl90WithDefaults() *ResumeResumeShortLogoUrl90`

NewResumeResumeShortLogoUrl90WithDefaults instantiates a new ResumeResumeShortLogoUrl90 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *ResumeResumeShortLogoUrl90) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumeResumeShortLogoUrl90) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumeResumeShortLogoUrl90) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *ResumeResumeShortLogoUrl90) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeResumeShortLogoUrl90) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeResumeShortLogoUrl90) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ResumeResumeShortLogoUrl90) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeResumeShortLogoUrl90) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeResumeShortLogoUrl90) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ResumeResumeShortLogoUrl90) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeResumeShortLogoUrl90) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *ResumeResumeShortLogoUrl90) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumeResumeShortLogoUrl90) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumeResumeShortLogoUrl90) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAge

`func (o *ResumeResumeShortLogoUrl90) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumeResumeShortLogoUrl90) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumeResumeShortLogoUrl90) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumeResumeShortLogoUrl90) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumeResumeShortLogoUrl90) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumeResumeShortLogoUrl90) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetArea

`func (o *ResumeResumeShortLogoUrl90) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeResumeShortLogoUrl90) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeResumeShortLogoUrl90) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeResumeShortLogoUrl90) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeResumeShortLogoUrl90) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeResumeShortLogoUrl90) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumeResumeShortLogoUrl90) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumeResumeShortLogoUrl90) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumeResumeShortLogoUrl90) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumeResumeShortLogoUrl90) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumeResumeShortLogoUrl90) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumeResumeShortLogoUrl90) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumeResumeShortLogoUrl90) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeResumeShortLogoUrl90) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeResumeShortLogoUrl90) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumeResumeShortLogoUrl90) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumeResumeShortLogoUrl90) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumeResumeShortLogoUrl90) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumeResumeShortLogoUrl90) GetDownload() map[string]interface{}`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeResumeShortLogoUrl90) GetDownloadOk() (*map[string]interface{}, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeResumeShortLogoUrl90) SetDownload(v map[string]interface{})`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumeResumeShortLogoUrl90) GetEducation() map[string]interface{}`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeResumeShortLogoUrl90) GetEducationOk() (*map[string]interface{}, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeResumeShortLogoUrl90) SetEducation(v map[string]interface{})`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumeResumeShortLogoUrl90) GetExperience() []ResumeObjectsExperienceLogoUrl90`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeResumeShortLogoUrl90) GetExperienceOk() (*[]ResumeObjectsExperienceLogoUrl90, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeResumeShortLogoUrl90) SetExperience(v []ResumeObjectsExperienceLogoUrl90)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumeResumeShortLogoUrl90) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeResumeShortLogoUrl90) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeResumeShortLogoUrl90) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeResumeShortLogoUrl90) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeResumeShortLogoUrl90) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeResumeShortLogoUrl90) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumeResumeShortLogoUrl90) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeResumeShortLogoUrl90) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeResumeShortLogoUrl90) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeResumeShortLogoUrl90) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ResumeResumeShortLogoUrl90) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ResumeResumeShortLogoUrl90) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *ResumeResumeShortLogoUrl90) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeResumeShortLogoUrl90) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeResumeShortLogoUrl90) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetLastName

`func (o *ResumeResumeShortLogoUrl90) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeResumeShortLogoUrl90) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeResumeShortLogoUrl90) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeResumeShortLogoUrl90) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeResumeShortLogoUrl90) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeResumeShortLogoUrl90) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumeResumeShortLogoUrl90) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumeResumeShortLogoUrl90) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumeResumeShortLogoUrl90) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *ResumeResumeShortLogoUrl90) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetMiddleName

`func (o *ResumeResumeShortLogoUrl90) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeResumeShortLogoUrl90) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeResumeShortLogoUrl90) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeResumeShortLogoUrl90) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeResumeShortLogoUrl90) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeResumeShortLogoUrl90) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPlatform

`func (o *ResumeResumeShortLogoUrl90) GetPlatform() map[string]interface{}`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumeResumeShortLogoUrl90) GetPlatformOk() (*map[string]interface{}, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumeResumeShortLogoUrl90) SetPlatform(v map[string]interface{})`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumeResumeShortLogoUrl90) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeResumeShortLogoUrl90) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeResumeShortLogoUrl90) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeResumeShortLogoUrl90) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeResumeShortLogoUrl90) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumeResumeShortLogoUrl90) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumeResumeShortLogoUrl90) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTotalExperience

`func (o *ResumeResumeShortLogoUrl90) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeResumeShortLogoUrl90) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeResumeShortLogoUrl90) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeResumeShortLogoUrl90) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeResumeShortLogoUrl90) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeResumeShortLogoUrl90) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumeResumeShortLogoUrl90) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumeResumeShortLogoUrl90) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumeResumeShortLogoUrl90) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetActions

`func (o *ResumeResumeShortLogoUrl90) GetActions() ResumeObjectsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumeResumeShortLogoUrl90) GetActionsOk() (*ResumeObjectsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumeResumeShortLogoUrl90) SetActions(v ResumeObjectsActions)`

SetActions sets Actions field to given value.


### GetFavorited

`func (o *ResumeResumeShortLogoUrl90) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *ResumeResumeShortLogoUrl90) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *ResumeResumeShortLogoUrl90) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetNegotiationsHistory

`func (o *ResumeResumeShortLogoUrl90) GetNegotiationsHistory() ResumeObjectsNegotiationsHistoryUrl`

GetNegotiationsHistory returns the NegotiationsHistory field if non-nil, zero value otherwise.

### GetNegotiationsHistoryOk

`func (o *ResumeResumeShortLogoUrl90) GetNegotiationsHistoryOk() (*ResumeObjectsNegotiationsHistoryUrl, bool)`

GetNegotiationsHistoryOk returns a tuple with the NegotiationsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsHistory

`func (o *ResumeResumeShortLogoUrl90) SetNegotiationsHistory(v ResumeObjectsNegotiationsHistoryUrl)`

SetNegotiationsHistory sets NegotiationsHistory field to given value.


### GetOwner

`func (o *ResumeResumeShortLogoUrl90) GetOwner() ResumeObjectsOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResumeResumeShortLogoUrl90) GetOwnerOk() (*ResumeObjectsOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResumeResumeShortLogoUrl90) SetOwner(v ResumeObjectsOwner)`

SetOwner sets Owner field to given value.


### GetPhoto

`func (o *ResumeResumeShortLogoUrl90) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeResumeShortLogoUrl90) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeResumeShortLogoUrl90) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeResumeShortLogoUrl90) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeResumeShortLogoUrl90) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeResumeShortLogoUrl90) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetTags

`func (o *ResumeResumeShortLogoUrl90) GetTags() []IncludesId`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResumeResumeShortLogoUrl90) GetTagsOk() (*[]IncludesId, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResumeResumeShortLogoUrl90) SetTags(v []IncludesId)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResumeResumeShortLogoUrl90) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetViewed

`func (o *ResumeResumeShortLogoUrl90) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *ResumeResumeShortLogoUrl90) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *ResumeResumeShortLogoUrl90) SetViewed(v bool)`

SetViewed sets Viewed field to given value.


### GetJobSearchStatus

`func (o *ResumeResumeShortLogoUrl90) GetJobSearchStatus() IncludesIdNameLastChangeTime`

GetJobSearchStatus returns the JobSearchStatus field if non-nil, zero value otherwise.

### GetJobSearchStatusOk

`func (o *ResumeResumeShortLogoUrl90) GetJobSearchStatusOk() (*IncludesIdNameLastChangeTime, bool)`

GetJobSearchStatusOk returns a tuple with the JobSearchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatus

`func (o *ResumeResumeShortLogoUrl90) SetJobSearchStatus(v IncludesIdNameLastChangeTime)`

SetJobSearchStatus sets JobSearchStatus field to given value.

### HasJobSearchStatus

`func (o *ResumeResumeShortLogoUrl90) HasJobSearchStatus() bool`

HasJobSearchStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


