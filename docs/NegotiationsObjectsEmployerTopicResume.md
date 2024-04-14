# NegotiationsObjectsEmployerTopicResume

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
**Experience** | [**[]ResumeObjectsExperience**](ResumeObjectsExperience.md) | Опыт работы | 
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
**Photo** | Pointer to [**NullableResumeObjectsPhoto**](ResumeObjectsPhoto.md) | Фотография пользователя | [optional] 
**Tags** | Pointer to [**[]IncludesId**](IncludesId.md) | Теги к резюме | [optional] 
**Viewed** | **bool** | Было ли резюме уже просмотрено работодателем | 
**Url** | **string** | Ссылка на получение элементов | 
**JobSearchStatus** | Pointer to [**IncludesIdNameLastChangeTime**](IncludesIdNameLastChangeTime.md) | Для получения данных нужно передать параметр &#x60;with_job_search_status&#x3D;true&#x60;.  Возможные значения перечислены в поле &#x60;job_search_statuses_employer&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | [optional] 

## Methods

### NewNegotiationsObjectsEmployerTopicResume

`func NewNegotiationsObjectsEmployerTopicResume(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download map[string]interface{}, education map[string]interface{}, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, actions ResumeObjectsActions, favorited bool, negotiationsHistory ResumeObjectsNegotiationsHistoryUrl, owner ResumeObjectsOwner, viewed bool, url string, ) *NegotiationsObjectsEmployerTopicResume`

NewNegotiationsObjectsEmployerTopicResume instantiates a new NegotiationsObjectsEmployerTopicResume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsObjectsEmployerTopicResumeWithDefaults

`func NewNegotiationsObjectsEmployerTopicResumeWithDefaults() *NegotiationsObjectsEmployerTopicResume`

NewNegotiationsObjectsEmployerTopicResumeWithDefaults instantiates a new NegotiationsObjectsEmployerTopicResume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *NegotiationsObjectsEmployerTopicResume) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *NegotiationsObjectsEmployerTopicResume) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *NegotiationsObjectsEmployerTopicResume) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsObjectsEmployerTopicResume) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *NegotiationsObjectsEmployerTopicResume) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NegotiationsObjectsEmployerTopicResume) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAge

`func (o *NegotiationsObjectsEmployerTopicResume) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *NegotiationsObjectsEmployerTopicResume) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *NegotiationsObjectsEmployerTopicResume) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetArea

`func (o *NegotiationsObjectsEmployerTopicResume) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *NegotiationsObjectsEmployerTopicResume) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *NegotiationsObjectsEmployerTopicResume) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCanViewFullInfo

`func (o *NegotiationsObjectsEmployerTopicResume) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *NegotiationsObjectsEmployerTopicResume) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *NegotiationsObjectsEmployerTopicResume) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *NegotiationsObjectsEmployerTopicResume) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *NegotiationsObjectsEmployerTopicResume) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *NegotiationsObjectsEmployerTopicResume) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsObjectsEmployerTopicResume) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *NegotiationsObjectsEmployerTopicResume) GetDownload() map[string]interface{}`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetDownloadOk() (*map[string]interface{}, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *NegotiationsObjectsEmployerTopicResume) SetDownload(v map[string]interface{})`

SetDownload sets Download field to given value.


### GetEducation

`func (o *NegotiationsObjectsEmployerTopicResume) GetEducation() map[string]interface{}`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetEducationOk() (*map[string]interface{}, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *NegotiationsObjectsEmployerTopicResume) SetEducation(v map[string]interface{})`

SetEducation sets Education field to given value.


### GetExperience

`func (o *NegotiationsObjectsEmployerTopicResume) GetExperience() []ResumeObjectsExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetExperienceOk() (*[]ResumeObjectsExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *NegotiationsObjectsEmployerTopicResume) SetExperience(v []ResumeObjectsExperience)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *NegotiationsObjectsEmployerTopicResume) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *NegotiationsObjectsEmployerTopicResume) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *NegotiationsObjectsEmployerTopicResume) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *NegotiationsObjectsEmployerTopicResume) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *NegotiationsObjectsEmployerTopicResume) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *NegotiationsObjectsEmployerTopicResume) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *NegotiationsObjectsEmployerTopicResume) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *NegotiationsObjectsEmployerTopicResume) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetLastName

`func (o *NegotiationsObjectsEmployerTopicResume) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *NegotiationsObjectsEmployerTopicResume) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *NegotiationsObjectsEmployerTopicResume) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *NegotiationsObjectsEmployerTopicResume) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *NegotiationsObjectsEmployerTopicResume) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *NegotiationsObjectsEmployerTopicResume) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetMiddleName

`func (o *NegotiationsObjectsEmployerTopicResume) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *NegotiationsObjectsEmployerTopicResume) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *NegotiationsObjectsEmployerTopicResume) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPlatform

`func (o *NegotiationsObjectsEmployerTopicResume) GetPlatform() map[string]interface{}`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetPlatformOk() (*map[string]interface{}, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *NegotiationsObjectsEmployerTopicResume) SetPlatform(v map[string]interface{})`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *NegotiationsObjectsEmployerTopicResume) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *NegotiationsObjectsEmployerTopicResume) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *NegotiationsObjectsEmployerTopicResume) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *NegotiationsObjectsEmployerTopicResume) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTotalExperience

`func (o *NegotiationsObjectsEmployerTopicResume) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *NegotiationsObjectsEmployerTopicResume) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *NegotiationsObjectsEmployerTopicResume) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *NegotiationsObjectsEmployerTopicResume) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NegotiationsObjectsEmployerTopicResume) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetActions

`func (o *NegotiationsObjectsEmployerTopicResume) GetActions() ResumeObjectsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetActionsOk() (*ResumeObjectsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *NegotiationsObjectsEmployerTopicResume) SetActions(v ResumeObjectsActions)`

SetActions sets Actions field to given value.


### GetFavorited

`func (o *NegotiationsObjectsEmployerTopicResume) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *NegotiationsObjectsEmployerTopicResume) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetNegotiationsHistory

`func (o *NegotiationsObjectsEmployerTopicResume) GetNegotiationsHistory() ResumeObjectsNegotiationsHistoryUrl`

GetNegotiationsHistory returns the NegotiationsHistory field if non-nil, zero value otherwise.

### GetNegotiationsHistoryOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetNegotiationsHistoryOk() (*ResumeObjectsNegotiationsHistoryUrl, bool)`

GetNegotiationsHistoryOk returns a tuple with the NegotiationsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsHistory

`func (o *NegotiationsObjectsEmployerTopicResume) SetNegotiationsHistory(v ResumeObjectsNegotiationsHistoryUrl)`

SetNegotiationsHistory sets NegotiationsHistory field to given value.


### GetOwner

`func (o *NegotiationsObjectsEmployerTopicResume) GetOwner() ResumeObjectsOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetOwnerOk() (*ResumeObjectsOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NegotiationsObjectsEmployerTopicResume) SetOwner(v ResumeObjectsOwner)`

SetOwner sets Owner field to given value.


### GetPhoto

`func (o *NegotiationsObjectsEmployerTopicResume) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *NegotiationsObjectsEmployerTopicResume) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *NegotiationsObjectsEmployerTopicResume) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *NegotiationsObjectsEmployerTopicResume) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *NegotiationsObjectsEmployerTopicResume) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetTags

`func (o *NegotiationsObjectsEmployerTopicResume) GetTags() []IncludesId`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetTagsOk() (*[]IncludesId, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *NegotiationsObjectsEmployerTopicResume) SetTags(v []IncludesId)`

SetTags sets Tags field to given value.

### HasTags

`func (o *NegotiationsObjectsEmployerTopicResume) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetViewed

`func (o *NegotiationsObjectsEmployerTopicResume) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *NegotiationsObjectsEmployerTopicResume) SetViewed(v bool)`

SetViewed sets Viewed field to given value.


### GetUrl

`func (o *NegotiationsObjectsEmployerTopicResume) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NegotiationsObjectsEmployerTopicResume) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetJobSearchStatus

`func (o *NegotiationsObjectsEmployerTopicResume) GetJobSearchStatus() IncludesIdNameLastChangeTime`

GetJobSearchStatus returns the JobSearchStatus field if non-nil, zero value otherwise.

### GetJobSearchStatusOk

`func (o *NegotiationsObjectsEmployerTopicResume) GetJobSearchStatusOk() (*IncludesIdNameLastChangeTime, bool)`

GetJobSearchStatusOk returns a tuple with the JobSearchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatus

`func (o *NegotiationsObjectsEmployerTopicResume) SetJobSearchStatus(v IncludesIdNameLastChangeTime)`

SetJobSearchStatus sets JobSearchStatus field to given value.

### HasJobSearchStatus

`func (o *NegotiationsObjectsEmployerTopicResume) HasJobSearchStatus() bool`

HasJobSearchStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


