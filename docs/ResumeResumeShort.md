# ResumeResumeShort

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
**Download** | [**ResumeObjectsDownload**](ResumeObjectsDownload.md) | Ссылки для скачивания резюме в разных форматах | 
**Education** | [**ResumeObjectsEducation**](ResumeObjectsEducation.md) | Образование соискателя.   Особенности сохранения образования:  * Если передать и высшее и среднее образование и уровень образования \&quot;средний\&quot;, то сохранится только среднее образование. * Если передать и высшее и среднее образование и уровень образования \&quot;высшее\&quot;, то сохранится только высшее образование  | 
**Experience** | [**[]ResumeObjectsExperience**](ResumeObjectsExperience.md) | Опыт работы | 
**FirstName** | Pointer to **NullableString** | Имя | [optional] 
**Gender** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**HiddenFields** | [**[]IncludesIdName**](IncludesIdName.md) | Справочник [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле &#x60;resume_hidden_fields&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**LastName** | Pointer to **NullableString** | Фамилия | [optional] 
**Marked** | Pointer to **bool** | Выделено ли резюме в поиске | [optional] 
**MiddleName** | Pointer to **NullableString** | Отчество | [optional] 
**Platform** | Pointer to [**IncludesId**](IncludesId.md) | Ресурс, на котором было размещено резюме | [optional] 
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

## Methods

### NewResumeResumeShort

`func NewResumeResumeShort(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download ResumeObjectsDownload, education ResumeObjectsEducation, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, actions ResumeObjectsActions, favorited bool, negotiationsHistory ResumeObjectsNegotiationsHistoryUrl, owner ResumeObjectsOwner, viewed bool, ) *ResumeResumeShort`

NewResumeResumeShort instantiates a new ResumeResumeShort object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeShortWithDefaults

`func NewResumeResumeShortWithDefaults() *ResumeResumeShort`

NewResumeResumeShortWithDefaults instantiates a new ResumeResumeShort object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *ResumeResumeShort) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumeResumeShort) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumeResumeShort) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *ResumeResumeShort) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeResumeShort) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeResumeShort) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ResumeResumeShort) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeResumeShort) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeResumeShort) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ResumeResumeShort) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeResumeShort) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAge

`func (o *ResumeResumeShort) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumeResumeShort) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumeResumeShort) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumeResumeShort) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumeResumeShort) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumeResumeShort) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetArea

`func (o *ResumeResumeShort) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeResumeShort) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeResumeShort) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeResumeShort) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeResumeShort) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeResumeShort) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumeResumeShort) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumeResumeShort) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumeResumeShort) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumeResumeShort) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumeResumeShort) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumeResumeShort) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumeResumeShort) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeResumeShort) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeResumeShort) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumeResumeShort) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumeResumeShort) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumeResumeShort) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumeResumeShort) GetDownload() ResumeObjectsDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeResumeShort) GetDownloadOk() (*ResumeObjectsDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeResumeShort) SetDownload(v ResumeObjectsDownload)`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumeResumeShort) GetEducation() ResumeObjectsEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeResumeShort) GetEducationOk() (*ResumeObjectsEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeResumeShort) SetEducation(v ResumeObjectsEducation)`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumeResumeShort) GetExperience() []ResumeObjectsExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeResumeShort) GetExperienceOk() (*[]ResumeObjectsExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeResumeShort) SetExperience(v []ResumeObjectsExperience)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumeResumeShort) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeResumeShort) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeResumeShort) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeResumeShort) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeResumeShort) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeResumeShort) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumeResumeShort) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeResumeShort) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeResumeShort) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeResumeShort) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ResumeResumeShort) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ResumeResumeShort) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *ResumeResumeShort) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeResumeShort) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeResumeShort) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetLastName

`func (o *ResumeResumeShort) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeResumeShort) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeResumeShort) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeResumeShort) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeResumeShort) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeResumeShort) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumeResumeShort) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumeResumeShort) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumeResumeShort) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *ResumeResumeShort) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetMiddleName

`func (o *ResumeResumeShort) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeResumeShort) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeResumeShort) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeResumeShort) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeResumeShort) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeResumeShort) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPlatform

`func (o *ResumeResumeShort) GetPlatform() IncludesId`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumeResumeShort) GetPlatformOk() (*IncludesId, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumeResumeShort) SetPlatform(v IncludesId)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumeResumeShort) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeResumeShort) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeResumeShort) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeResumeShort) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeResumeShort) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumeResumeShort) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumeResumeShort) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTotalExperience

`func (o *ResumeResumeShort) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeResumeShort) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeResumeShort) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeResumeShort) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeResumeShort) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeResumeShort) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumeResumeShort) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumeResumeShort) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumeResumeShort) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetActions

`func (o *ResumeResumeShort) GetActions() ResumeObjectsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumeResumeShort) GetActionsOk() (*ResumeObjectsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumeResumeShort) SetActions(v ResumeObjectsActions)`

SetActions sets Actions field to given value.


### GetFavorited

`func (o *ResumeResumeShort) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *ResumeResumeShort) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *ResumeResumeShort) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetNegotiationsHistory

`func (o *ResumeResumeShort) GetNegotiationsHistory() ResumeObjectsNegotiationsHistoryUrl`

GetNegotiationsHistory returns the NegotiationsHistory field if non-nil, zero value otherwise.

### GetNegotiationsHistoryOk

`func (o *ResumeResumeShort) GetNegotiationsHistoryOk() (*ResumeObjectsNegotiationsHistoryUrl, bool)`

GetNegotiationsHistoryOk returns a tuple with the NegotiationsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsHistory

`func (o *ResumeResumeShort) SetNegotiationsHistory(v ResumeObjectsNegotiationsHistoryUrl)`

SetNegotiationsHistory sets NegotiationsHistory field to given value.


### GetOwner

`func (o *ResumeResumeShort) GetOwner() ResumeObjectsOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResumeResumeShort) GetOwnerOk() (*ResumeObjectsOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResumeResumeShort) SetOwner(v ResumeObjectsOwner)`

SetOwner sets Owner field to given value.


### GetPhoto

`func (o *ResumeResumeShort) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeResumeShort) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeResumeShort) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeResumeShort) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeResumeShort) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeResumeShort) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetTags

`func (o *ResumeResumeShort) GetTags() []IncludesId`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResumeResumeShort) GetTagsOk() (*[]IncludesId, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResumeResumeShort) SetTags(v []IncludesId)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResumeResumeShort) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetViewed

`func (o *ResumeResumeShort) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *ResumeResumeShort) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *ResumeResumeShort) SetViewed(v bool)`

SetViewed sets Viewed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


