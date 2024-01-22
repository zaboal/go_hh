# ResumeResumeViewResponse

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
**Download** | [**ResumeResumeProfileAllOfDownload**](ResumeResumeProfileAllOfDownload.md) |  | 
**Education** | [**ResumeResumeProfileAllOfEducation**](ResumeResumeProfileAllOfEducation.md) |  | 
**Experience** | [**[]ResumeObjectsExperience**](ResumeObjectsExperience.md) | Опыт работы | 
**FirstName** | Pointer to **NullableString** | Имя | [optional] 
**Gender** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**HiddenFields** | [**[]IncludesIdName**](IncludesIdName.md) | Справочник [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле &#x60;resume_hidden_fields&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**LastName** | Pointer to **NullableString** | Фамилия | [optional] 
**Marked** | Pointer to **bool** | Выделено ли резюме в поиске | [optional] [default to false]
**MiddleName** | Pointer to **NullableString** | Отчество | [optional] 
**Platform** | Pointer to [**ResumeResumeProfileAllOfPlatform**](ResumeResumeProfileAllOfPlatform.md) |  | [optional] 
**Salary** | Pointer to [**NullableResumeObjectsSalaryProperties**](ResumeObjectsSalaryProperties.md) |  | [optional] 
**TotalExperience** | Pointer to [**NullableResumeObjectsTotalExperience**](ResumeObjectsTotalExperience.md) |  | [optional] 
**UpdatedAt** | **string** | Дата и время обновления резюме | 
**BirthDate** | Pointer to **NullableString** | День рождения (в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**BusinessTripReadiness** | [**ResumeResumeFullAllOfBusinessTripReadiness**](ResumeResumeFullAllOfBusinessTripReadiness.md) |  | 
**Citizenship** | [**[]IncludesIdNameUrl**](IncludesIdNameUrl.md) | Список гражданств соискателя. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas) | 
**Contact** | [**[]IncludesContact**](IncludesContact.md) | Список контактов соискателя | 
**Creds** | Pointer to [**NullableCredsCreds**](CredsCreds.md) |  | [optional] 
**DriverLicenseTypes** | [**[]ResumeObjectsDriverLicenseTypes**](ResumeObjectsDriverLicenseTypes.md) | Список категорий водительских прав соискателя | 
**Employment** | Pointer to [**ResumeResumeFullAllOfEmployment**](ResumeResumeFullAllOfEmployment.md) |  | [optional] 
**Employments** | [**[]IncludesIdName**](IncludesIdName.md) | Список подходящих соискателю типов занятостей. Элементы справочника [employment](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**HasVehicle** | Pointer to **NullableBool** | Наличие личного автомобиля у соискателя | [optional] 
**Language** | [**[]IncludesLanguageLevel**](IncludesLanguageLevel.md) | Список языков, которыми владеет соискатель. Элементы справочника [languages](#tag/Obshie-spravochniki/operation/get-languages) | 
**Metro** | Pointer to [**NullableResumeObjectsMetroStation**](ResumeObjectsMetroStation.md) |  | [optional] 
**PaidServices** | [**[]ResumeObjectsPaidServices**](ResumeObjectsPaidServices.md) | Платные услуги по резюме для автора | 
**ProfessionalRoles** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Массив объектов профролей | [optional] 
**Recommendation** | [**[]ResumeObjectsRecommendation**](ResumeObjectsRecommendation.md) | Список рекомендаций | 
**Relocation** | [**ResumeResumeFullAllOfRelocation**](ResumeResumeFullAllOfRelocation.md) |  | 
**ResumeLocale** | [**ResumeResumeFullAllOfResumeLocale**](ResumeResumeFullAllOfResumeLocale.md) |  | 
**Schedule** | [**ResumeResumeFullAllOfSchedule**](ResumeResumeFullAllOfSchedule.md) |  | 
**Schedules** | [**[]IncludesIdName**](IncludesIdName.md) | Список подходящих соискателю графиков работы. Элементы справочника [schedule](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Site** | [**[]ResumeObjectsSite**](ResumeObjectsSite.md) | Профили в соц. сетях и других сервисах | 
**SkillSet** | **[]string** | Ключевые навыки (список уникальных строк) | 
**Skills** | Pointer to **NullableString** | Дополнительная информация, описание навыков в свободной форме | [optional] 
**TravelTime** | [**ResumeResumeFullAllOfTravelTime**](ResumeResumeFullAllOfTravelTime.md) |  | 
**WorkTicket** | [**[]IncludesIdNameUrl**](IncludesIdNameUrl.md) | Список регионов, в которых соискатель имеет разрешение на работу. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas)  | 
**Actions** | [**ResumeObjectsActionsForOwner**](ResumeObjectsActionsForOwner.md) | Дополнительные действия | 
**Favorited** | **bool** | Добавлено ли резюме в избранные | 
**JobSearchStatus** | Pointer to [**IncludesIdName**](IncludesIdName.md) | Для получения данных нужно передать параметр &#x60;with_job_search_status&#x3D;true&#x60;. 
Возможные значения перечислены в поле &#x60;job_search_status&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)
 | [optional] 
**NegotiationsHistory** | [**ResumeObjectsNegotiationsHistoryForEmployer**](ResumeObjectsNegotiationsHistoryForEmployer.md) | Краткая история откликов/приглашений по резюме | 
**Owner** | [**ResumeObjectsOwner**](ResumeObjectsOwner.md) | Информация о владельце резюме | 
**Photo** | Pointer to [**NullableResumeObjectsPhoto**](ResumeObjectsPhoto.md) |  | [optional] 
**Portfolio** | [**[]ResumeObjectsPortfolio**](ResumeObjectsPortfolio.md) | Список изображений в портфолио пользователя | 
**Blocked** | **bool** | Заблокировано ли резюме ([подробнее](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume)) | 
**CanPublishOrUpdate** | Pointer to **NullableBool** | Можно ли опубликовать или обновить данное резюме | [optional] 
**Finished** | **bool** | Заполнено ли резюме | 
**Status** | [**IncludesIdName**](IncludesIdName.md) |  | 
**ModerationNote** | [**[]ResumeObjectsModerationNote**](ResumeObjectsModerationNote.md) | Замечания модератора. В некоторых случаях замечания могут сопровождаться [блокировкой резюме](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume). Полный список возможных замечаний доступен в поле &#x60;resume_moderation_note&#x60; [в справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**Progress** | [**ResumeObjectsProgress**](ResumeObjectsProgress.md) |  | 
**PublishUrl** | **string** | URL для публикации или обновления резюме | 
**Access** | [**ResumeObjectsAccess**](ResumeObjectsAccess.md) |  | 
**NewViews** | **float32** | Число новых просмотров. Данный счетчик сбрасывается при получении [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history)  | 
**NextPublishAt** | Pointer to **NullableString** | Дата и время следующей возможной публикации/обновления. Для неопубликованных резюме возвращается &#x60;null&#x60; | [optional] 
**TotalViews** | **float32** | Число просмотров резюме | 
**ViewsUrl** | **string** | URL, по которому необходимо сделать GET-запрос для получения [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history)  | 

## Methods

### NewResumeResumeViewResponse

`func NewResumeResumeViewResponse(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download ResumeResumeProfileAllOfDownload, education ResumeResumeProfileAllOfEducation, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, businessTripReadiness ResumeResumeFullAllOfBusinessTripReadiness, citizenship []IncludesIdNameUrl, contact []IncludesContact, driverLicenseTypes []ResumeObjectsDriverLicenseTypes, employments []IncludesIdName, language []IncludesLanguageLevel, paidServices []ResumeObjectsPaidServices, recommendation []ResumeObjectsRecommendation, relocation ResumeResumeFullAllOfRelocation, resumeLocale ResumeResumeFullAllOfResumeLocale, schedule ResumeResumeFullAllOfSchedule, schedules []IncludesIdName, site []ResumeObjectsSite, skillSet []string, travelTime ResumeResumeFullAllOfTravelTime, workTicket []IncludesIdNameUrl, actions ResumeObjectsActionsForOwner, favorited bool, negotiationsHistory ResumeObjectsNegotiationsHistoryForEmployer, owner ResumeObjectsOwner, portfolio []ResumeObjectsPortfolio, blocked bool, finished bool, status IncludesIdName, moderationNote []ResumeObjectsModerationNote, progress ResumeObjectsProgress, publishUrl string, access ResumeObjectsAccess, newViews float32, totalViews float32, viewsUrl string, ) *ResumeResumeViewResponse`

NewResumeResumeViewResponse instantiates a new ResumeResumeViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeViewResponseWithDefaults

`func NewResumeResumeViewResponseWithDefaults() *ResumeResumeViewResponse`

NewResumeResumeViewResponseWithDefaults instantiates a new ResumeResumeViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *ResumeResumeViewResponse) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumeResumeViewResponse) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumeResumeViewResponse) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *ResumeResumeViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeResumeViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeResumeViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ResumeResumeViewResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeResumeViewResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeResumeViewResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ResumeResumeViewResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeResumeViewResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAge

`func (o *ResumeResumeViewResponse) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumeResumeViewResponse) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumeResumeViewResponse) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumeResumeViewResponse) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumeResumeViewResponse) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumeResumeViewResponse) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetArea

`func (o *ResumeResumeViewResponse) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeResumeViewResponse) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeResumeViewResponse) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeResumeViewResponse) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeResumeViewResponse) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeResumeViewResponse) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumeResumeViewResponse) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumeResumeViewResponse) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumeResumeViewResponse) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumeResumeViewResponse) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumeResumeViewResponse) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumeResumeViewResponse) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumeResumeViewResponse) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeResumeViewResponse) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeResumeViewResponse) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumeResumeViewResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumeResumeViewResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumeResumeViewResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumeResumeViewResponse) GetDownload() ResumeResumeProfileAllOfDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeResumeViewResponse) GetDownloadOk() (*ResumeResumeProfileAllOfDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeResumeViewResponse) SetDownload(v ResumeResumeProfileAllOfDownload)`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumeResumeViewResponse) GetEducation() ResumeResumeProfileAllOfEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeResumeViewResponse) GetEducationOk() (*ResumeResumeProfileAllOfEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeResumeViewResponse) SetEducation(v ResumeResumeProfileAllOfEducation)`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumeResumeViewResponse) GetExperience() []ResumeObjectsExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeResumeViewResponse) GetExperienceOk() (*[]ResumeObjectsExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeResumeViewResponse) SetExperience(v []ResumeObjectsExperience)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumeResumeViewResponse) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeResumeViewResponse) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeResumeViewResponse) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeResumeViewResponse) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeResumeViewResponse) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeResumeViewResponse) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumeResumeViewResponse) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeResumeViewResponse) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeResumeViewResponse) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeResumeViewResponse) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ResumeResumeViewResponse) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ResumeResumeViewResponse) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *ResumeResumeViewResponse) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeResumeViewResponse) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeResumeViewResponse) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetLastName

`func (o *ResumeResumeViewResponse) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeResumeViewResponse) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeResumeViewResponse) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeResumeViewResponse) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeResumeViewResponse) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeResumeViewResponse) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumeResumeViewResponse) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumeResumeViewResponse) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumeResumeViewResponse) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *ResumeResumeViewResponse) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetMiddleName

`func (o *ResumeResumeViewResponse) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeResumeViewResponse) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeResumeViewResponse) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeResumeViewResponse) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeResumeViewResponse) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeResumeViewResponse) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPlatform

`func (o *ResumeResumeViewResponse) GetPlatform() ResumeResumeProfileAllOfPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumeResumeViewResponse) GetPlatformOk() (*ResumeResumeProfileAllOfPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumeResumeViewResponse) SetPlatform(v ResumeResumeProfileAllOfPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumeResumeViewResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeResumeViewResponse) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeResumeViewResponse) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeResumeViewResponse) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeResumeViewResponse) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumeResumeViewResponse) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumeResumeViewResponse) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTotalExperience

`func (o *ResumeResumeViewResponse) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeResumeViewResponse) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeResumeViewResponse) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeResumeViewResponse) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeResumeViewResponse) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeResumeViewResponse) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumeResumeViewResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumeResumeViewResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumeResumeViewResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBirthDate

`func (o *ResumeResumeViewResponse) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *ResumeResumeViewResponse) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *ResumeResumeViewResponse) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *ResumeResumeViewResponse) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *ResumeResumeViewResponse) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *ResumeResumeViewResponse) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetBusinessTripReadiness

`func (o *ResumeResumeViewResponse) GetBusinessTripReadiness() ResumeResumeFullAllOfBusinessTripReadiness`

GetBusinessTripReadiness returns the BusinessTripReadiness field if non-nil, zero value otherwise.

### GetBusinessTripReadinessOk

`func (o *ResumeResumeViewResponse) GetBusinessTripReadinessOk() (*ResumeResumeFullAllOfBusinessTripReadiness, bool)`

GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTripReadiness

`func (o *ResumeResumeViewResponse) SetBusinessTripReadiness(v ResumeResumeFullAllOfBusinessTripReadiness)`

SetBusinessTripReadiness sets BusinessTripReadiness field to given value.


### GetCitizenship

`func (o *ResumeResumeViewResponse) GetCitizenship() []IncludesIdNameUrl`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *ResumeResumeViewResponse) GetCitizenshipOk() (*[]IncludesIdNameUrl, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *ResumeResumeViewResponse) SetCitizenship(v []IncludesIdNameUrl)`

SetCitizenship sets Citizenship field to given value.


### GetContact

`func (o *ResumeResumeViewResponse) GetContact() []IncludesContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ResumeResumeViewResponse) GetContactOk() (*[]IncludesContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ResumeResumeViewResponse) SetContact(v []IncludesContact)`

SetContact sets Contact field to given value.


### GetCreds

`func (o *ResumeResumeViewResponse) GetCreds() CredsCreds`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *ResumeResumeViewResponse) GetCredsOk() (*CredsCreds, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *ResumeResumeViewResponse) SetCreds(v CredsCreds)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *ResumeResumeViewResponse) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### SetCredsNil

`func (o *ResumeResumeViewResponse) SetCredsNil(b bool)`

 SetCredsNil sets the value for Creds to be an explicit nil

### UnsetCreds
`func (o *ResumeResumeViewResponse) UnsetCreds()`

UnsetCreds ensures that no value is present for Creds, not even an explicit nil
### GetDriverLicenseTypes

`func (o *ResumeResumeViewResponse) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *ResumeResumeViewResponse) GetDriverLicenseTypesOk() (*[]ResumeObjectsDriverLicenseTypes, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *ResumeResumeViewResponse) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEmployment

`func (o *ResumeResumeViewResponse) GetEmployment() ResumeResumeFullAllOfEmployment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *ResumeResumeViewResponse) GetEmploymentOk() (*ResumeResumeFullAllOfEmployment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *ResumeResumeViewResponse) SetEmployment(v ResumeResumeFullAllOfEmployment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *ResumeResumeViewResponse) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### GetEmployments

`func (o *ResumeResumeViewResponse) GetEmployments() []IncludesIdName`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *ResumeResumeViewResponse) GetEmploymentsOk() (*[]IncludesIdName, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *ResumeResumeViewResponse) SetEmployments(v []IncludesIdName)`

SetEmployments sets Employments field to given value.


### GetHasVehicle

`func (o *ResumeResumeViewResponse) GetHasVehicle() bool`

GetHasVehicle returns the HasVehicle field if non-nil, zero value otherwise.

### GetHasVehicleOk

`func (o *ResumeResumeViewResponse) GetHasVehicleOk() (*bool, bool)`

GetHasVehicleOk returns a tuple with the HasVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVehicle

`func (o *ResumeResumeViewResponse) SetHasVehicle(v bool)`

SetHasVehicle sets HasVehicle field to given value.

### HasHasVehicle

`func (o *ResumeResumeViewResponse) HasHasVehicle() bool`

HasHasVehicle returns a boolean if a field has been set.

### SetHasVehicleNil

`func (o *ResumeResumeViewResponse) SetHasVehicleNil(b bool)`

 SetHasVehicleNil sets the value for HasVehicle to be an explicit nil

### UnsetHasVehicle
`func (o *ResumeResumeViewResponse) UnsetHasVehicle()`

UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
### GetLanguage

`func (o *ResumeResumeViewResponse) GetLanguage() []IncludesLanguageLevel`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ResumeResumeViewResponse) GetLanguageOk() (*[]IncludesLanguageLevel, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ResumeResumeViewResponse) SetLanguage(v []IncludesLanguageLevel)`

SetLanguage sets Language field to given value.


### GetMetro

`func (o *ResumeResumeViewResponse) GetMetro() ResumeObjectsMetroStation`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *ResumeResumeViewResponse) GetMetroOk() (*ResumeObjectsMetroStation, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *ResumeResumeViewResponse) SetMetro(v ResumeObjectsMetroStation)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *ResumeResumeViewResponse) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### SetMetroNil

`func (o *ResumeResumeViewResponse) SetMetroNil(b bool)`

 SetMetroNil sets the value for Metro to be an explicit nil

### UnsetMetro
`func (o *ResumeResumeViewResponse) UnsetMetro()`

UnsetMetro ensures that no value is present for Metro, not even an explicit nil
### GetPaidServices

`func (o *ResumeResumeViewResponse) GetPaidServices() []ResumeObjectsPaidServices`

GetPaidServices returns the PaidServices field if non-nil, zero value otherwise.

### GetPaidServicesOk

`func (o *ResumeResumeViewResponse) GetPaidServicesOk() (*[]ResumeObjectsPaidServices, bool)`

GetPaidServicesOk returns a tuple with the PaidServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServices

`func (o *ResumeResumeViewResponse) SetPaidServices(v []ResumeObjectsPaidServices)`

SetPaidServices sets PaidServices field to given value.


### GetProfessionalRoles

`func (o *ResumeResumeViewResponse) GetProfessionalRoles() []IncludesIdName`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *ResumeResumeViewResponse) GetProfessionalRolesOk() (*[]IncludesIdName, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *ResumeResumeViewResponse) SetProfessionalRoles(v []IncludesIdName)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *ResumeResumeViewResponse) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### SetProfessionalRolesNil

`func (o *ResumeResumeViewResponse) SetProfessionalRolesNil(b bool)`

 SetProfessionalRolesNil sets the value for ProfessionalRoles to be an explicit nil

### UnsetProfessionalRoles
`func (o *ResumeResumeViewResponse) UnsetProfessionalRoles()`

UnsetProfessionalRoles ensures that no value is present for ProfessionalRoles, not even an explicit nil
### GetRecommendation

`func (o *ResumeResumeViewResponse) GetRecommendation() []ResumeObjectsRecommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *ResumeResumeViewResponse) GetRecommendationOk() (*[]ResumeObjectsRecommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *ResumeResumeViewResponse) SetRecommendation(v []ResumeObjectsRecommendation)`

SetRecommendation sets Recommendation field to given value.


### GetRelocation

`func (o *ResumeResumeViewResponse) GetRelocation() ResumeResumeFullAllOfRelocation`

GetRelocation returns the Relocation field if non-nil, zero value otherwise.

### GetRelocationOk

`func (o *ResumeResumeViewResponse) GetRelocationOk() (*ResumeResumeFullAllOfRelocation, bool)`

GetRelocationOk returns a tuple with the Relocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocation

`func (o *ResumeResumeViewResponse) SetRelocation(v ResumeResumeFullAllOfRelocation)`

SetRelocation sets Relocation field to given value.


### GetResumeLocale

`func (o *ResumeResumeViewResponse) GetResumeLocale() ResumeResumeFullAllOfResumeLocale`

GetResumeLocale returns the ResumeLocale field if non-nil, zero value otherwise.

### GetResumeLocaleOk

`func (o *ResumeResumeViewResponse) GetResumeLocaleOk() (*ResumeResumeFullAllOfResumeLocale, bool)`

GetResumeLocaleOk returns a tuple with the ResumeLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeLocale

`func (o *ResumeResumeViewResponse) SetResumeLocale(v ResumeResumeFullAllOfResumeLocale)`

SetResumeLocale sets ResumeLocale field to given value.


### GetSchedule

`func (o *ResumeResumeViewResponse) GetSchedule() ResumeResumeFullAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ResumeResumeViewResponse) GetScheduleOk() (*ResumeResumeFullAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ResumeResumeViewResponse) SetSchedule(v ResumeResumeFullAllOfSchedule)`

SetSchedule sets Schedule field to given value.


### GetSchedules

`func (o *ResumeResumeViewResponse) GetSchedules() []IncludesIdName`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ResumeResumeViewResponse) GetSchedulesOk() (*[]IncludesIdName, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ResumeResumeViewResponse) SetSchedules(v []IncludesIdName)`

SetSchedules sets Schedules field to given value.


### GetSite

`func (o *ResumeResumeViewResponse) GetSite() []ResumeObjectsSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ResumeResumeViewResponse) GetSiteOk() (*[]ResumeObjectsSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ResumeResumeViewResponse) SetSite(v []ResumeObjectsSite)`

SetSite sets Site field to given value.


### GetSkillSet

`func (o *ResumeResumeViewResponse) GetSkillSet() []string`

GetSkillSet returns the SkillSet field if non-nil, zero value otherwise.

### GetSkillSetOk

`func (o *ResumeResumeViewResponse) GetSkillSetOk() (*[]string, bool)`

GetSkillSetOk returns a tuple with the SkillSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillSet

`func (o *ResumeResumeViewResponse) SetSkillSet(v []string)`

SetSkillSet sets SkillSet field to given value.


### GetSkills

`func (o *ResumeResumeViewResponse) GetSkills() string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ResumeResumeViewResponse) GetSkillsOk() (*string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ResumeResumeViewResponse) SetSkills(v string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ResumeResumeViewResponse) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkillsNil

`func (o *ResumeResumeViewResponse) SetSkillsNil(b bool)`

 SetSkillsNil sets the value for Skills to be an explicit nil

### UnsetSkills
`func (o *ResumeResumeViewResponse) UnsetSkills()`

UnsetSkills ensures that no value is present for Skills, not even an explicit nil
### GetTravelTime

`func (o *ResumeResumeViewResponse) GetTravelTime() ResumeResumeFullAllOfTravelTime`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *ResumeResumeViewResponse) GetTravelTimeOk() (*ResumeResumeFullAllOfTravelTime, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *ResumeResumeViewResponse) SetTravelTime(v ResumeResumeFullAllOfTravelTime)`

SetTravelTime sets TravelTime field to given value.


### GetWorkTicket

`func (o *ResumeResumeViewResponse) GetWorkTicket() []IncludesIdNameUrl`

GetWorkTicket returns the WorkTicket field if non-nil, zero value otherwise.

### GetWorkTicketOk

`func (o *ResumeResumeViewResponse) GetWorkTicketOk() (*[]IncludesIdNameUrl, bool)`

GetWorkTicketOk returns a tuple with the WorkTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTicket

`func (o *ResumeResumeViewResponse) SetWorkTicket(v []IncludesIdNameUrl)`

SetWorkTicket sets WorkTicket field to given value.


### GetActions

`func (o *ResumeResumeViewResponse) GetActions() ResumeObjectsActionsForOwner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumeResumeViewResponse) GetActionsOk() (*ResumeObjectsActionsForOwner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumeResumeViewResponse) SetActions(v ResumeObjectsActionsForOwner)`

SetActions sets Actions field to given value.


### GetFavorited

`func (o *ResumeResumeViewResponse) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *ResumeResumeViewResponse) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *ResumeResumeViewResponse) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetJobSearchStatus

`func (o *ResumeResumeViewResponse) GetJobSearchStatus() IncludesIdName`

GetJobSearchStatus returns the JobSearchStatus field if non-nil, zero value otherwise.

### GetJobSearchStatusOk

`func (o *ResumeResumeViewResponse) GetJobSearchStatusOk() (*IncludesIdName, bool)`

GetJobSearchStatusOk returns a tuple with the JobSearchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatus

`func (o *ResumeResumeViewResponse) SetJobSearchStatus(v IncludesIdName)`

SetJobSearchStatus sets JobSearchStatus field to given value.

### HasJobSearchStatus

`func (o *ResumeResumeViewResponse) HasJobSearchStatus() bool`

HasJobSearchStatus returns a boolean if a field has been set.

### GetNegotiationsHistory

`func (o *ResumeResumeViewResponse) GetNegotiationsHistory() ResumeObjectsNegotiationsHistoryForEmployer`

GetNegotiationsHistory returns the NegotiationsHistory field if non-nil, zero value otherwise.

### GetNegotiationsHistoryOk

`func (o *ResumeResumeViewResponse) GetNegotiationsHistoryOk() (*ResumeObjectsNegotiationsHistoryForEmployer, bool)`

GetNegotiationsHistoryOk returns a tuple with the NegotiationsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsHistory

`func (o *ResumeResumeViewResponse) SetNegotiationsHistory(v ResumeObjectsNegotiationsHistoryForEmployer)`

SetNegotiationsHistory sets NegotiationsHistory field to given value.


### GetOwner

`func (o *ResumeResumeViewResponse) GetOwner() ResumeObjectsOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResumeResumeViewResponse) GetOwnerOk() (*ResumeObjectsOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResumeResumeViewResponse) SetOwner(v ResumeObjectsOwner)`

SetOwner sets Owner field to given value.


### GetPhoto

`func (o *ResumeResumeViewResponse) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeResumeViewResponse) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeResumeViewResponse) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeResumeViewResponse) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeResumeViewResponse) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeResumeViewResponse) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPortfolio

`func (o *ResumeResumeViewResponse) GetPortfolio() []ResumeObjectsPortfolio`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *ResumeResumeViewResponse) GetPortfolioOk() (*[]ResumeObjectsPortfolio, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *ResumeResumeViewResponse) SetPortfolio(v []ResumeObjectsPortfolio)`

SetPortfolio sets Portfolio field to given value.


### GetBlocked

`func (o *ResumeResumeViewResponse) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *ResumeResumeViewResponse) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *ResumeResumeViewResponse) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetCanPublishOrUpdate

`func (o *ResumeResumeViewResponse) GetCanPublishOrUpdate() bool`

GetCanPublishOrUpdate returns the CanPublishOrUpdate field if non-nil, zero value otherwise.

### GetCanPublishOrUpdateOk

`func (o *ResumeResumeViewResponse) GetCanPublishOrUpdateOk() (*bool, bool)`

GetCanPublishOrUpdateOk returns a tuple with the CanPublishOrUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPublishOrUpdate

`func (o *ResumeResumeViewResponse) SetCanPublishOrUpdate(v bool)`

SetCanPublishOrUpdate sets CanPublishOrUpdate field to given value.

### HasCanPublishOrUpdate

`func (o *ResumeResumeViewResponse) HasCanPublishOrUpdate() bool`

HasCanPublishOrUpdate returns a boolean if a field has been set.

### SetCanPublishOrUpdateNil

`func (o *ResumeResumeViewResponse) SetCanPublishOrUpdateNil(b bool)`

 SetCanPublishOrUpdateNil sets the value for CanPublishOrUpdate to be an explicit nil

### UnsetCanPublishOrUpdate
`func (o *ResumeResumeViewResponse) UnsetCanPublishOrUpdate()`

UnsetCanPublishOrUpdate ensures that no value is present for CanPublishOrUpdate, not even an explicit nil
### GetFinished

`func (o *ResumeResumeViewResponse) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ResumeResumeViewResponse) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ResumeResumeViewResponse) SetFinished(v bool)`

SetFinished sets Finished field to given value.


### GetStatus

`func (o *ResumeResumeViewResponse) GetStatus() IncludesIdName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResumeResumeViewResponse) GetStatusOk() (*IncludesIdName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResumeResumeViewResponse) SetStatus(v IncludesIdName)`

SetStatus sets Status field to given value.


### GetModerationNote

`func (o *ResumeResumeViewResponse) GetModerationNote() []ResumeObjectsModerationNote`

GetModerationNote returns the ModerationNote field if non-nil, zero value otherwise.

### GetModerationNoteOk

`func (o *ResumeResumeViewResponse) GetModerationNoteOk() (*[]ResumeObjectsModerationNote, bool)`

GetModerationNoteOk returns a tuple with the ModerationNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationNote

`func (o *ResumeResumeViewResponse) SetModerationNote(v []ResumeObjectsModerationNote)`

SetModerationNote sets ModerationNote field to given value.


### GetProgress

`func (o *ResumeResumeViewResponse) GetProgress() ResumeObjectsProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ResumeResumeViewResponse) GetProgressOk() (*ResumeObjectsProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ResumeResumeViewResponse) SetProgress(v ResumeObjectsProgress)`

SetProgress sets Progress field to given value.


### GetPublishUrl

`func (o *ResumeResumeViewResponse) GetPublishUrl() string`

GetPublishUrl returns the PublishUrl field if non-nil, zero value otherwise.

### GetPublishUrlOk

`func (o *ResumeResumeViewResponse) GetPublishUrlOk() (*string, bool)`

GetPublishUrlOk returns a tuple with the PublishUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishUrl

`func (o *ResumeResumeViewResponse) SetPublishUrl(v string)`

SetPublishUrl sets PublishUrl field to given value.


### GetAccess

`func (o *ResumeResumeViewResponse) GetAccess() ResumeObjectsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ResumeResumeViewResponse) GetAccessOk() (*ResumeObjectsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ResumeResumeViewResponse) SetAccess(v ResumeObjectsAccess)`

SetAccess sets Access field to given value.


### GetNewViews

`func (o *ResumeResumeViewResponse) GetNewViews() float32`

GetNewViews returns the NewViews field if non-nil, zero value otherwise.

### GetNewViewsOk

`func (o *ResumeResumeViewResponse) GetNewViewsOk() (*float32, bool)`

GetNewViewsOk returns a tuple with the NewViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewViews

`func (o *ResumeResumeViewResponse) SetNewViews(v float32)`

SetNewViews sets NewViews field to given value.


### GetNextPublishAt

`func (o *ResumeResumeViewResponse) GetNextPublishAt() string`

GetNextPublishAt returns the NextPublishAt field if non-nil, zero value otherwise.

### GetNextPublishAtOk

`func (o *ResumeResumeViewResponse) GetNextPublishAtOk() (*string, bool)`

GetNextPublishAtOk returns a tuple with the NextPublishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPublishAt

`func (o *ResumeResumeViewResponse) SetNextPublishAt(v string)`

SetNextPublishAt sets NextPublishAt field to given value.

### HasNextPublishAt

`func (o *ResumeResumeViewResponse) HasNextPublishAt() bool`

HasNextPublishAt returns a boolean if a field has been set.

### SetNextPublishAtNil

`func (o *ResumeResumeViewResponse) SetNextPublishAtNil(b bool)`

 SetNextPublishAtNil sets the value for NextPublishAt to be an explicit nil

### UnsetNextPublishAt
`func (o *ResumeResumeViewResponse) UnsetNextPublishAt()`

UnsetNextPublishAt ensures that no value is present for NextPublishAt, not even an explicit nil
### GetTotalViews

`func (o *ResumeResumeViewResponse) GetTotalViews() float32`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *ResumeResumeViewResponse) GetTotalViewsOk() (*float32, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *ResumeResumeViewResponse) SetTotalViews(v float32)`

SetTotalViews sets TotalViews field to given value.


### GetViewsUrl

`func (o *ResumeResumeViewResponse) GetViewsUrl() string`

GetViewsUrl returns the ViewsUrl field if non-nil, zero value otherwise.

### GetViewsUrlOk

`func (o *ResumeResumeViewResponse) GetViewsUrlOk() (*string, bool)`

GetViewsUrlOk returns a tuple with the ViewsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsUrl

`func (o *ResumeResumeViewResponse) SetViewsUrl(v string)`

SetViewsUrl sets ViewsUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


