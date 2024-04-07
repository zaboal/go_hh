# ResumeResumeForApplicant

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
**Marked** | Pointer to **bool** | Выделено ли резюме в поиске | [optional] [default to false]
**MiddleName** | Pointer to **NullableString** | Отчество | [optional] 
**Platform** | Pointer to [**IncludesId**](IncludesId.md) | Ресурс, на котором было размещено резюме | [optional] 
**Salary** | Pointer to [**NullableResumeObjectsSalaryProperties**](ResumeObjectsSalaryProperties.md) |  | [optional] 
**TotalExperience** | Pointer to [**NullableResumeObjectsTotalExperience**](ResumeObjectsTotalExperience.md) |  | [optional] 
**UpdatedAt** | **string** | Дата и время обновления резюме | 
**BirthDate** | Pointer to **NullableString** | День рождения (в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**BusinessTripReadiness** | [**IncludesIdName**](IncludesIdName.md) | Готовность к командировкам. Элемент справочника [business_trip_readiness](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Citizenship** | [**[]IncludesIdNameUrl**](IncludesIdNameUrl.md) | Список гражданств соискателя. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas) | 
**Contact** | [**[]IncludesContact**](IncludesContact.md) | Список контактов соискателя | 
**Creds** | Pointer to [**NullableCredsCreds**](CredsCreds.md) |  | [optional] 
**DriverLicenseTypes** | [**[]ResumeObjectsDriverLicenseTypes**](ResumeObjectsDriverLicenseTypes.md) | Список категорий водительских прав соискателя | 
**Employment** | Pointer to [**IncludesIdName**](IncludesIdName.md) |  | [optional] 
**Employments** | [**[]IncludesIdName**](IncludesIdName.md) | Список подходящих соискателю типов занятостей. Элементы справочника [employment](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**HasVehicle** | Pointer to **NullableBool** | Наличие личного автомобиля у соискателя | [optional] 
**Language** | [**[]IncludesLanguageLevel**](IncludesLanguageLevel.md) | Список языков, которыми владеет соискатель. Элементы справочника [languages](#tag/Obshie-spravochniki/operation/get-languages) | 
**Metro** | Pointer to [**NullableResumeObjectsMetroStation**](ResumeObjectsMetroStation.md) |  | [optional] 
**PaidServices** | [**[]ResumeObjectsPaidServices**](ResumeObjectsPaidServices.md) | Платные услуги по резюме для автора | 
**ProfessionalRoles** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Массив объектов профролей | [optional] 
**Recommendation** | [**[]ResumeObjectsRecommendation**](ResumeObjectsRecommendation.md) | Список рекомендаций | 
**Relocation** | [**ResumeObjectsRelocationPublic**](ResumeObjectsRelocationPublic.md) | Возможность переезда | 
**ResumeLocale** | [**IncludesIdName**](IncludesIdName.md) | Язык, на котором составлено резюме (локаль). Элемент справочника [локали резюме](#tag/Obshie-spravochniki/operation/get-locales) | 
**Schedule** | [**IncludesIdName**](IncludesIdName.md) |  | 
**Schedules** | [**[]IncludesIdName**](IncludesIdName.md) | Список подходящих соискателю графиков работы. Элементы справочника [schedule](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Site** | [**[]ResumeObjectsSite**](ResumeObjectsSite.md) | Профили в соц. сетях и других сервисах | 
**SkillSet** | **[]string** | Ключевые навыки (список уникальных строк) | 
**Skills** | Pointer to **NullableString** | Дополнительная информация, описание навыков в свободной форме | [optional] 
**Tags** | Pointer to [**[]IncludesId**](IncludesId.md) | Теги к резюме | [optional] 
**TravelTime** | [**IncludesIdName**](IncludesIdName.md) | Желательное время в пути до работы. Элемент справочника [travel_time](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**WorkTicket** | [**[]IncludesIdNameUrl**](IncludesIdNameUrl.md) | Список регионов, в которых соискатель имеет разрешение на работу. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas)  | 
**Blocked** | **bool** | Заблокировано ли резюме ([подробнее](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume)) | 
**CanPublishOrUpdate** | Pointer to **NullableBool** | Можно ли опубликовать или обновить данное резюме | [optional] 
**Finished** | **bool** | Заполнено ли резюме | 
**Status** | [**IncludesIdName**](IncludesIdName.md) |  | 
**ModerationNote** | [**[]ResumeObjectsModerationNote**](ResumeObjectsModerationNote.md) | Замечания модератора. В некоторых случаях замечания могут сопровождаться [блокировкой резюме](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume). Полный список возможных замечаний доступен в поле &#x60;resume_moderation_note&#x60; [в справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**Progress** | [**ResumeObjectsProgress**](ResumeObjectsProgress.md) |  | 
**PublishUrl** | **string** | URL для публикации или обновления резюме | 
**Access** | [**ResumeObjectsAccess**](ResumeObjectsAccess.md) |  | 
**Actions** | [**ResumeObjectsActionsForOwner**](ResumeObjectsActionsForOwner.md) | Дополнительные действия | 
**NewViews** | **float32** | Число новых просмотров. Данный счетчик сбрасывается при получении [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history)  | 
**NextPublishAt** | Pointer to **NullableString** | Дата и время следующей возможной публикации/обновления. Для неопубликованных резюме возвращается &#x60;null&#x60; | [optional] 
**TotalViews** | **float32** | Число просмотров резюме | 
**ViewsUrl** | **string** | URL, по которому необходимо сделать GET-запрос для получения [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history)  | 
**Photo** | Pointer to [**NullableResumeObjectsPhoto**](ResumeObjectsPhoto.md) |  | [optional] 
**Portfolio** | [**[]ResumeObjectsPortfolio**](ResumeObjectsPortfolio.md) | Список изображений в портфолио пользователя | 

## Methods

### NewResumeResumeForApplicant

`func NewResumeResumeForApplicant(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download ResumeObjectsDownload, education ResumeObjectsEducation, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, businessTripReadiness IncludesIdName, citizenship []IncludesIdNameUrl, contact []IncludesContact, driverLicenseTypes []ResumeObjectsDriverLicenseTypes, employments []IncludesIdName, language []IncludesLanguageLevel, paidServices []ResumeObjectsPaidServices, recommendation []ResumeObjectsRecommendation, relocation ResumeObjectsRelocationPublic, resumeLocale IncludesIdName, schedule IncludesIdName, schedules []IncludesIdName, site []ResumeObjectsSite, skillSet []string, travelTime IncludesIdName, workTicket []IncludesIdNameUrl, blocked bool, finished bool, status IncludesIdName, moderationNote []ResumeObjectsModerationNote, progress ResumeObjectsProgress, publishUrl string, access ResumeObjectsAccess, actions ResumeObjectsActionsForOwner, newViews float32, totalViews float32, viewsUrl string, portfolio []ResumeObjectsPortfolio, ) *ResumeResumeForApplicant`

NewResumeResumeForApplicant instantiates a new ResumeResumeForApplicant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeForApplicantWithDefaults

`func NewResumeResumeForApplicantWithDefaults() *ResumeResumeForApplicant`

NewResumeResumeForApplicantWithDefaults instantiates a new ResumeResumeForApplicant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *ResumeResumeForApplicant) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumeResumeForApplicant) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumeResumeForApplicant) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *ResumeResumeForApplicant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeResumeForApplicant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeResumeForApplicant) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ResumeResumeForApplicant) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeResumeForApplicant) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeResumeForApplicant) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ResumeResumeForApplicant) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeResumeForApplicant) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAge

`func (o *ResumeResumeForApplicant) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumeResumeForApplicant) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumeResumeForApplicant) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumeResumeForApplicant) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumeResumeForApplicant) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumeResumeForApplicant) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetArea

`func (o *ResumeResumeForApplicant) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeResumeForApplicant) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeResumeForApplicant) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeResumeForApplicant) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeResumeForApplicant) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeResumeForApplicant) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumeResumeForApplicant) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumeResumeForApplicant) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumeResumeForApplicant) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumeResumeForApplicant) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumeResumeForApplicant) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumeResumeForApplicant) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumeResumeForApplicant) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeResumeForApplicant) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeResumeForApplicant) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumeResumeForApplicant) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumeResumeForApplicant) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumeResumeForApplicant) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumeResumeForApplicant) GetDownload() ResumeObjectsDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeResumeForApplicant) GetDownloadOk() (*ResumeObjectsDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeResumeForApplicant) SetDownload(v ResumeObjectsDownload)`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumeResumeForApplicant) GetEducation() ResumeObjectsEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeResumeForApplicant) GetEducationOk() (*ResumeObjectsEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeResumeForApplicant) SetEducation(v ResumeObjectsEducation)`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumeResumeForApplicant) GetExperience() []ResumeObjectsExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeResumeForApplicant) GetExperienceOk() (*[]ResumeObjectsExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeResumeForApplicant) SetExperience(v []ResumeObjectsExperience)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumeResumeForApplicant) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeResumeForApplicant) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeResumeForApplicant) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeResumeForApplicant) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeResumeForApplicant) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeResumeForApplicant) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumeResumeForApplicant) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeResumeForApplicant) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeResumeForApplicant) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeResumeForApplicant) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ResumeResumeForApplicant) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ResumeResumeForApplicant) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *ResumeResumeForApplicant) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeResumeForApplicant) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeResumeForApplicant) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetLastName

`func (o *ResumeResumeForApplicant) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeResumeForApplicant) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeResumeForApplicant) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeResumeForApplicant) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeResumeForApplicant) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeResumeForApplicant) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumeResumeForApplicant) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumeResumeForApplicant) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumeResumeForApplicant) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *ResumeResumeForApplicant) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetMiddleName

`func (o *ResumeResumeForApplicant) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeResumeForApplicant) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeResumeForApplicant) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeResumeForApplicant) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeResumeForApplicant) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeResumeForApplicant) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPlatform

`func (o *ResumeResumeForApplicant) GetPlatform() IncludesId`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumeResumeForApplicant) GetPlatformOk() (*IncludesId, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumeResumeForApplicant) SetPlatform(v IncludesId)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumeResumeForApplicant) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeResumeForApplicant) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeResumeForApplicant) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeResumeForApplicant) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeResumeForApplicant) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumeResumeForApplicant) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumeResumeForApplicant) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTotalExperience

`func (o *ResumeResumeForApplicant) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeResumeForApplicant) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeResumeForApplicant) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeResumeForApplicant) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeResumeForApplicant) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeResumeForApplicant) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumeResumeForApplicant) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumeResumeForApplicant) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumeResumeForApplicant) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBirthDate

`func (o *ResumeResumeForApplicant) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *ResumeResumeForApplicant) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *ResumeResumeForApplicant) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *ResumeResumeForApplicant) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *ResumeResumeForApplicant) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *ResumeResumeForApplicant) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetBusinessTripReadiness

`func (o *ResumeResumeForApplicant) GetBusinessTripReadiness() IncludesIdName`

GetBusinessTripReadiness returns the BusinessTripReadiness field if non-nil, zero value otherwise.

### GetBusinessTripReadinessOk

`func (o *ResumeResumeForApplicant) GetBusinessTripReadinessOk() (*IncludesIdName, bool)`

GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTripReadiness

`func (o *ResumeResumeForApplicant) SetBusinessTripReadiness(v IncludesIdName)`

SetBusinessTripReadiness sets BusinessTripReadiness field to given value.


### GetCitizenship

`func (o *ResumeResumeForApplicant) GetCitizenship() []IncludesIdNameUrl`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *ResumeResumeForApplicant) GetCitizenshipOk() (*[]IncludesIdNameUrl, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *ResumeResumeForApplicant) SetCitizenship(v []IncludesIdNameUrl)`

SetCitizenship sets Citizenship field to given value.


### GetContact

`func (o *ResumeResumeForApplicant) GetContact() []IncludesContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ResumeResumeForApplicant) GetContactOk() (*[]IncludesContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ResumeResumeForApplicant) SetContact(v []IncludesContact)`

SetContact sets Contact field to given value.


### GetCreds

`func (o *ResumeResumeForApplicant) GetCreds() CredsCreds`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *ResumeResumeForApplicant) GetCredsOk() (*CredsCreds, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *ResumeResumeForApplicant) SetCreds(v CredsCreds)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *ResumeResumeForApplicant) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### SetCredsNil

`func (o *ResumeResumeForApplicant) SetCredsNil(b bool)`

 SetCredsNil sets the value for Creds to be an explicit nil

### UnsetCreds
`func (o *ResumeResumeForApplicant) UnsetCreds()`

UnsetCreds ensures that no value is present for Creds, not even an explicit nil
### GetDriverLicenseTypes

`func (o *ResumeResumeForApplicant) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *ResumeResumeForApplicant) GetDriverLicenseTypesOk() (*[]ResumeObjectsDriverLicenseTypes, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *ResumeResumeForApplicant) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEmployment

`func (o *ResumeResumeForApplicant) GetEmployment() IncludesIdName`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *ResumeResumeForApplicant) GetEmploymentOk() (*IncludesIdName, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *ResumeResumeForApplicant) SetEmployment(v IncludesIdName)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *ResumeResumeForApplicant) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### GetEmployments

`func (o *ResumeResumeForApplicant) GetEmployments() []IncludesIdName`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *ResumeResumeForApplicant) GetEmploymentsOk() (*[]IncludesIdName, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *ResumeResumeForApplicant) SetEmployments(v []IncludesIdName)`

SetEmployments sets Employments field to given value.


### GetHasVehicle

`func (o *ResumeResumeForApplicant) GetHasVehicle() bool`

GetHasVehicle returns the HasVehicle field if non-nil, zero value otherwise.

### GetHasVehicleOk

`func (o *ResumeResumeForApplicant) GetHasVehicleOk() (*bool, bool)`

GetHasVehicleOk returns a tuple with the HasVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVehicle

`func (o *ResumeResumeForApplicant) SetHasVehicle(v bool)`

SetHasVehicle sets HasVehicle field to given value.

### HasHasVehicle

`func (o *ResumeResumeForApplicant) HasHasVehicle() bool`

HasHasVehicle returns a boolean if a field has been set.

### SetHasVehicleNil

`func (o *ResumeResumeForApplicant) SetHasVehicleNil(b bool)`

 SetHasVehicleNil sets the value for HasVehicle to be an explicit nil

### UnsetHasVehicle
`func (o *ResumeResumeForApplicant) UnsetHasVehicle()`

UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
### GetLanguage

`func (o *ResumeResumeForApplicant) GetLanguage() []IncludesLanguageLevel`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ResumeResumeForApplicant) GetLanguageOk() (*[]IncludesLanguageLevel, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ResumeResumeForApplicant) SetLanguage(v []IncludesLanguageLevel)`

SetLanguage sets Language field to given value.


### GetMetro

`func (o *ResumeResumeForApplicant) GetMetro() ResumeObjectsMetroStation`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *ResumeResumeForApplicant) GetMetroOk() (*ResumeObjectsMetroStation, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *ResumeResumeForApplicant) SetMetro(v ResumeObjectsMetroStation)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *ResumeResumeForApplicant) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### SetMetroNil

`func (o *ResumeResumeForApplicant) SetMetroNil(b bool)`

 SetMetroNil sets the value for Metro to be an explicit nil

### UnsetMetro
`func (o *ResumeResumeForApplicant) UnsetMetro()`

UnsetMetro ensures that no value is present for Metro, not even an explicit nil
### GetPaidServices

`func (o *ResumeResumeForApplicant) GetPaidServices() []ResumeObjectsPaidServices`

GetPaidServices returns the PaidServices field if non-nil, zero value otherwise.

### GetPaidServicesOk

`func (o *ResumeResumeForApplicant) GetPaidServicesOk() (*[]ResumeObjectsPaidServices, bool)`

GetPaidServicesOk returns a tuple with the PaidServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServices

`func (o *ResumeResumeForApplicant) SetPaidServices(v []ResumeObjectsPaidServices)`

SetPaidServices sets PaidServices field to given value.


### GetProfessionalRoles

`func (o *ResumeResumeForApplicant) GetProfessionalRoles() []IncludesIdName`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *ResumeResumeForApplicant) GetProfessionalRolesOk() (*[]IncludesIdName, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *ResumeResumeForApplicant) SetProfessionalRoles(v []IncludesIdName)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *ResumeResumeForApplicant) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### SetProfessionalRolesNil

`func (o *ResumeResumeForApplicant) SetProfessionalRolesNil(b bool)`

 SetProfessionalRolesNil sets the value for ProfessionalRoles to be an explicit nil

### UnsetProfessionalRoles
`func (o *ResumeResumeForApplicant) UnsetProfessionalRoles()`

UnsetProfessionalRoles ensures that no value is present for ProfessionalRoles, not even an explicit nil
### GetRecommendation

`func (o *ResumeResumeForApplicant) GetRecommendation() []ResumeObjectsRecommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *ResumeResumeForApplicant) GetRecommendationOk() (*[]ResumeObjectsRecommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *ResumeResumeForApplicant) SetRecommendation(v []ResumeObjectsRecommendation)`

SetRecommendation sets Recommendation field to given value.


### GetRelocation

`func (o *ResumeResumeForApplicant) GetRelocation() ResumeObjectsRelocationPublic`

GetRelocation returns the Relocation field if non-nil, zero value otherwise.

### GetRelocationOk

`func (o *ResumeResumeForApplicant) GetRelocationOk() (*ResumeObjectsRelocationPublic, bool)`

GetRelocationOk returns a tuple with the Relocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocation

`func (o *ResumeResumeForApplicant) SetRelocation(v ResumeObjectsRelocationPublic)`

SetRelocation sets Relocation field to given value.


### GetResumeLocale

`func (o *ResumeResumeForApplicant) GetResumeLocale() IncludesIdName`

GetResumeLocale returns the ResumeLocale field if non-nil, zero value otherwise.

### GetResumeLocaleOk

`func (o *ResumeResumeForApplicant) GetResumeLocaleOk() (*IncludesIdName, bool)`

GetResumeLocaleOk returns a tuple with the ResumeLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeLocale

`func (o *ResumeResumeForApplicant) SetResumeLocale(v IncludesIdName)`

SetResumeLocale sets ResumeLocale field to given value.


### GetSchedule

`func (o *ResumeResumeForApplicant) GetSchedule() IncludesIdName`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ResumeResumeForApplicant) GetScheduleOk() (*IncludesIdName, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ResumeResumeForApplicant) SetSchedule(v IncludesIdName)`

SetSchedule sets Schedule field to given value.


### GetSchedules

`func (o *ResumeResumeForApplicant) GetSchedules() []IncludesIdName`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ResumeResumeForApplicant) GetSchedulesOk() (*[]IncludesIdName, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ResumeResumeForApplicant) SetSchedules(v []IncludesIdName)`

SetSchedules sets Schedules field to given value.


### GetSite

`func (o *ResumeResumeForApplicant) GetSite() []ResumeObjectsSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ResumeResumeForApplicant) GetSiteOk() (*[]ResumeObjectsSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ResumeResumeForApplicant) SetSite(v []ResumeObjectsSite)`

SetSite sets Site field to given value.


### GetSkillSet

`func (o *ResumeResumeForApplicant) GetSkillSet() []string`

GetSkillSet returns the SkillSet field if non-nil, zero value otherwise.

### GetSkillSetOk

`func (o *ResumeResumeForApplicant) GetSkillSetOk() (*[]string, bool)`

GetSkillSetOk returns a tuple with the SkillSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillSet

`func (o *ResumeResumeForApplicant) SetSkillSet(v []string)`

SetSkillSet sets SkillSet field to given value.


### GetSkills

`func (o *ResumeResumeForApplicant) GetSkills() string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ResumeResumeForApplicant) GetSkillsOk() (*string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ResumeResumeForApplicant) SetSkills(v string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ResumeResumeForApplicant) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkillsNil

`func (o *ResumeResumeForApplicant) SetSkillsNil(b bool)`

 SetSkillsNil sets the value for Skills to be an explicit nil

### UnsetSkills
`func (o *ResumeResumeForApplicant) UnsetSkills()`

UnsetSkills ensures that no value is present for Skills, not even an explicit nil
### GetTags

`func (o *ResumeResumeForApplicant) GetTags() []IncludesId`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResumeResumeForApplicant) GetTagsOk() (*[]IncludesId, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResumeResumeForApplicant) SetTags(v []IncludesId)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResumeResumeForApplicant) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTravelTime

`func (o *ResumeResumeForApplicant) GetTravelTime() IncludesIdName`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *ResumeResumeForApplicant) GetTravelTimeOk() (*IncludesIdName, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *ResumeResumeForApplicant) SetTravelTime(v IncludesIdName)`

SetTravelTime sets TravelTime field to given value.


### GetWorkTicket

`func (o *ResumeResumeForApplicant) GetWorkTicket() []IncludesIdNameUrl`

GetWorkTicket returns the WorkTicket field if non-nil, zero value otherwise.

### GetWorkTicketOk

`func (o *ResumeResumeForApplicant) GetWorkTicketOk() (*[]IncludesIdNameUrl, bool)`

GetWorkTicketOk returns a tuple with the WorkTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTicket

`func (o *ResumeResumeForApplicant) SetWorkTicket(v []IncludesIdNameUrl)`

SetWorkTicket sets WorkTicket field to given value.


### GetBlocked

`func (o *ResumeResumeForApplicant) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *ResumeResumeForApplicant) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *ResumeResumeForApplicant) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetCanPublishOrUpdate

`func (o *ResumeResumeForApplicant) GetCanPublishOrUpdate() bool`

GetCanPublishOrUpdate returns the CanPublishOrUpdate field if non-nil, zero value otherwise.

### GetCanPublishOrUpdateOk

`func (o *ResumeResumeForApplicant) GetCanPublishOrUpdateOk() (*bool, bool)`

GetCanPublishOrUpdateOk returns a tuple with the CanPublishOrUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPublishOrUpdate

`func (o *ResumeResumeForApplicant) SetCanPublishOrUpdate(v bool)`

SetCanPublishOrUpdate sets CanPublishOrUpdate field to given value.

### HasCanPublishOrUpdate

`func (o *ResumeResumeForApplicant) HasCanPublishOrUpdate() bool`

HasCanPublishOrUpdate returns a boolean if a field has been set.

### SetCanPublishOrUpdateNil

`func (o *ResumeResumeForApplicant) SetCanPublishOrUpdateNil(b bool)`

 SetCanPublishOrUpdateNil sets the value for CanPublishOrUpdate to be an explicit nil

### UnsetCanPublishOrUpdate
`func (o *ResumeResumeForApplicant) UnsetCanPublishOrUpdate()`

UnsetCanPublishOrUpdate ensures that no value is present for CanPublishOrUpdate, not even an explicit nil
### GetFinished

`func (o *ResumeResumeForApplicant) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ResumeResumeForApplicant) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ResumeResumeForApplicant) SetFinished(v bool)`

SetFinished sets Finished field to given value.


### GetStatus

`func (o *ResumeResumeForApplicant) GetStatus() IncludesIdName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResumeResumeForApplicant) GetStatusOk() (*IncludesIdName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResumeResumeForApplicant) SetStatus(v IncludesIdName)`

SetStatus sets Status field to given value.


### GetModerationNote

`func (o *ResumeResumeForApplicant) GetModerationNote() []ResumeObjectsModerationNote`

GetModerationNote returns the ModerationNote field if non-nil, zero value otherwise.

### GetModerationNoteOk

`func (o *ResumeResumeForApplicant) GetModerationNoteOk() (*[]ResumeObjectsModerationNote, bool)`

GetModerationNoteOk returns a tuple with the ModerationNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationNote

`func (o *ResumeResumeForApplicant) SetModerationNote(v []ResumeObjectsModerationNote)`

SetModerationNote sets ModerationNote field to given value.


### GetProgress

`func (o *ResumeResumeForApplicant) GetProgress() ResumeObjectsProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ResumeResumeForApplicant) GetProgressOk() (*ResumeObjectsProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ResumeResumeForApplicant) SetProgress(v ResumeObjectsProgress)`

SetProgress sets Progress field to given value.


### GetPublishUrl

`func (o *ResumeResumeForApplicant) GetPublishUrl() string`

GetPublishUrl returns the PublishUrl field if non-nil, zero value otherwise.

### GetPublishUrlOk

`func (o *ResumeResumeForApplicant) GetPublishUrlOk() (*string, bool)`

GetPublishUrlOk returns a tuple with the PublishUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishUrl

`func (o *ResumeResumeForApplicant) SetPublishUrl(v string)`

SetPublishUrl sets PublishUrl field to given value.


### GetAccess

`func (o *ResumeResumeForApplicant) GetAccess() ResumeObjectsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ResumeResumeForApplicant) GetAccessOk() (*ResumeObjectsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ResumeResumeForApplicant) SetAccess(v ResumeObjectsAccess)`

SetAccess sets Access field to given value.


### GetActions

`func (o *ResumeResumeForApplicant) GetActions() ResumeObjectsActionsForOwner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumeResumeForApplicant) GetActionsOk() (*ResumeObjectsActionsForOwner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumeResumeForApplicant) SetActions(v ResumeObjectsActionsForOwner)`

SetActions sets Actions field to given value.


### GetNewViews

`func (o *ResumeResumeForApplicant) GetNewViews() float32`

GetNewViews returns the NewViews field if non-nil, zero value otherwise.

### GetNewViewsOk

`func (o *ResumeResumeForApplicant) GetNewViewsOk() (*float32, bool)`

GetNewViewsOk returns a tuple with the NewViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewViews

`func (o *ResumeResumeForApplicant) SetNewViews(v float32)`

SetNewViews sets NewViews field to given value.


### GetNextPublishAt

`func (o *ResumeResumeForApplicant) GetNextPublishAt() string`

GetNextPublishAt returns the NextPublishAt field if non-nil, zero value otherwise.

### GetNextPublishAtOk

`func (o *ResumeResumeForApplicant) GetNextPublishAtOk() (*string, bool)`

GetNextPublishAtOk returns a tuple with the NextPublishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPublishAt

`func (o *ResumeResumeForApplicant) SetNextPublishAt(v string)`

SetNextPublishAt sets NextPublishAt field to given value.

### HasNextPublishAt

`func (o *ResumeResumeForApplicant) HasNextPublishAt() bool`

HasNextPublishAt returns a boolean if a field has been set.

### SetNextPublishAtNil

`func (o *ResumeResumeForApplicant) SetNextPublishAtNil(b bool)`

 SetNextPublishAtNil sets the value for NextPublishAt to be an explicit nil

### UnsetNextPublishAt
`func (o *ResumeResumeForApplicant) UnsetNextPublishAt()`

UnsetNextPublishAt ensures that no value is present for NextPublishAt, not even an explicit nil
### GetTotalViews

`func (o *ResumeResumeForApplicant) GetTotalViews() float32`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *ResumeResumeForApplicant) GetTotalViewsOk() (*float32, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *ResumeResumeForApplicant) SetTotalViews(v float32)`

SetTotalViews sets TotalViews field to given value.


### GetViewsUrl

`func (o *ResumeResumeForApplicant) GetViewsUrl() string`

GetViewsUrl returns the ViewsUrl field if non-nil, zero value otherwise.

### GetViewsUrlOk

`func (o *ResumeResumeForApplicant) GetViewsUrlOk() (*string, bool)`

GetViewsUrlOk returns a tuple with the ViewsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsUrl

`func (o *ResumeResumeForApplicant) SetViewsUrl(v string)`

SetViewsUrl sets ViewsUrl field to given value.


### GetPhoto

`func (o *ResumeResumeForApplicant) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeResumeForApplicant) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeResumeForApplicant) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeResumeForApplicant) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeResumeForApplicant) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeResumeForApplicant) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPortfolio

`func (o *ResumeResumeForApplicant) GetPortfolio() []ResumeObjectsPortfolio`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *ResumeResumeForApplicant) GetPortfolioOk() (*[]ResumeObjectsPortfolio, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *ResumeResumeForApplicant) SetPortfolio(v []ResumeObjectsPortfolio)`

SetPortfolio sets Portfolio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


