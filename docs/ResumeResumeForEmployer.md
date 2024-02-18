# ResumeResumeForEmployer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternateUrl** | **string** | URL резюме на сайте | 
**Id** | **string** | Идентификатор резюме | 
**Title** | **NullableString** | Желаемая должность | 
**Age** | Pointer to **NullableFloat32** | Возраст | [optional] 
**Area** | Pointer to [**NullableIncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**CanViewFullInfo** | Pointer to **NullableBool** | Наличие права просмотра контактной информации в резюме | [optional] 
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
**PaidServices** | [**[]ResumeObjectsEmployerPaidServicesInner**](ResumeObjectsEmployerPaidServicesInner.md) | Платные услуги по резюме для работодателя  Работодателю может быть предложен список платных услуг по резюме.  Например, если полные данные по резюме недоступны, то будет выдано предложение покупки рекомендованной услуги, чтобы такой доступ получить  | 
**ProfessionalRoles** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Массив объектов профролей | [optional] 
**Recommendation** | [**[]ResumeObjectsRecommendation**](ResumeObjectsRecommendation.md) | Список рекомендаций | 
**Relocation** | [**ResumeObjectsRelocationPublic**](ResumeObjectsRelocationPublic.md) | Возможность переезда | 
**ResumeLocale** | [**IncludesIdName**](IncludesIdName.md) | Язык, на котором составлено резюме (локаль). Элемент справочника [локали резюме](#tag/Obshie-spravochniki/operation/get-locales) | 
**Schedule** | [**IncludesIdName**](IncludesIdName.md) |  | 
**Schedules** | [**[]IncludesIdName**](IncludesIdName.md) | Список подходящих соискателю графиков работы. Элементы справочника [schedule](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Site** | [**[]ResumeObjectsSite**](ResumeObjectsSite.md) | Профили в соц. сетях и других сервисах | 
**SkillSet** | **[]string** | Ключевые навыки (список уникальных строк) | 
**Skills** | Pointer to **NullableString** | Дополнительная информация, описание навыков в свободной форме | [optional] 
**TravelTime** | [**IncludesIdName**](IncludesIdName.md) | Желательное время в пути до работы. Элемент справочника [travel_time](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**WorkTicket** | [**[]IncludesIdNameUrl**](IncludesIdNameUrl.md) | Список регионов, в которых соискатель имеет разрешение на работу. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas)  | 
**Actions** | [**ResumeObjectsActions**](ResumeObjectsActions.md) | Дополнительные действия | 
**Favorited** | **bool** | Добавлено ли резюме в избранные | 
**JobSearchStatus** | Pointer to [**IncludesIdNameLastChangeTime**](IncludesIdNameLastChangeTime.md) | Для получения данных нужно передать параметр &#x60;with_job_search_status&#x3D;true&#x60;.  Возможные значения перечислены в поле &#x60;job_search_status&#x60; в [справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | [optional] 
**NegotiationsHistory** | [**ResumeObjectsNegotiationsHistoryForEmployer**](ResumeObjectsNegotiationsHistoryForEmployer.md) | Краткая история откликов/приглашений по резюме | 
**Owner** | [**ResumeObjectsOwner**](ResumeObjectsOwner.md) | Информация о владельце резюме | 
**Photo** | Pointer to [**NullableResumeObjectsPhotoNoId**](ResumeObjectsPhotoNoId.md) |  | [optional] 
**Portfolio** | [**[]ResumeObjectsPortfolioNoId**](ResumeObjectsPortfolioNoId.md) | Список изображений в портфолио пользователя | 

## Methods

### NewResumeResumeForEmployer

`func NewResumeResumeForEmployer(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download ResumeObjectsDownload, education ResumeObjectsEducation, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, businessTripReadiness IncludesIdName, citizenship []IncludesIdNameUrl, contact []IncludesContact, driverLicenseTypes []ResumeObjectsDriverLicenseTypes, employments []IncludesIdName, language []IncludesLanguageLevel, paidServices []ResumeObjectsEmployerPaidServicesInner, recommendation []ResumeObjectsRecommendation, relocation ResumeObjectsRelocationPublic, resumeLocale IncludesIdName, schedule IncludesIdName, schedules []IncludesIdName, site []ResumeObjectsSite, skillSet []string, travelTime IncludesIdName, workTicket []IncludesIdNameUrl, actions ResumeObjectsActions, favorited bool, negotiationsHistory ResumeObjectsNegotiationsHistoryForEmployer, owner ResumeObjectsOwner, portfolio []ResumeObjectsPortfolioNoId, ) *ResumeResumeForEmployer`

NewResumeResumeForEmployer instantiates a new ResumeResumeForEmployer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeForEmployerWithDefaults

`func NewResumeResumeForEmployerWithDefaults() *ResumeResumeForEmployer`

NewResumeResumeForEmployerWithDefaults instantiates a new ResumeResumeForEmployer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *ResumeResumeForEmployer) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumeResumeForEmployer) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumeResumeForEmployer) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *ResumeResumeForEmployer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeResumeForEmployer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeResumeForEmployer) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ResumeResumeForEmployer) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeResumeForEmployer) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeResumeForEmployer) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ResumeResumeForEmployer) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeResumeForEmployer) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAge

`func (o *ResumeResumeForEmployer) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumeResumeForEmployer) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumeResumeForEmployer) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumeResumeForEmployer) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumeResumeForEmployer) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumeResumeForEmployer) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetArea

`func (o *ResumeResumeForEmployer) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeResumeForEmployer) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeResumeForEmployer) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeResumeForEmployer) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeResumeForEmployer) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeResumeForEmployer) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumeResumeForEmployer) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumeResumeForEmployer) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumeResumeForEmployer) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumeResumeForEmployer) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumeResumeForEmployer) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumeResumeForEmployer) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumeResumeForEmployer) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeResumeForEmployer) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeResumeForEmployer) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumeResumeForEmployer) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumeResumeForEmployer) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumeResumeForEmployer) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumeResumeForEmployer) GetDownload() ResumeObjectsDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeResumeForEmployer) GetDownloadOk() (*ResumeObjectsDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeResumeForEmployer) SetDownload(v ResumeObjectsDownload)`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumeResumeForEmployer) GetEducation() ResumeObjectsEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeResumeForEmployer) GetEducationOk() (*ResumeObjectsEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeResumeForEmployer) SetEducation(v ResumeObjectsEducation)`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumeResumeForEmployer) GetExperience() []ResumeObjectsExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeResumeForEmployer) GetExperienceOk() (*[]ResumeObjectsExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeResumeForEmployer) SetExperience(v []ResumeObjectsExperience)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumeResumeForEmployer) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeResumeForEmployer) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeResumeForEmployer) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeResumeForEmployer) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeResumeForEmployer) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeResumeForEmployer) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumeResumeForEmployer) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeResumeForEmployer) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeResumeForEmployer) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeResumeForEmployer) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ResumeResumeForEmployer) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ResumeResumeForEmployer) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *ResumeResumeForEmployer) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeResumeForEmployer) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeResumeForEmployer) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetLastName

`func (o *ResumeResumeForEmployer) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeResumeForEmployer) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeResumeForEmployer) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeResumeForEmployer) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeResumeForEmployer) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeResumeForEmployer) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumeResumeForEmployer) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumeResumeForEmployer) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumeResumeForEmployer) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *ResumeResumeForEmployer) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetMiddleName

`func (o *ResumeResumeForEmployer) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeResumeForEmployer) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeResumeForEmployer) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeResumeForEmployer) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeResumeForEmployer) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeResumeForEmployer) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPlatform

`func (o *ResumeResumeForEmployer) GetPlatform() IncludesId`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumeResumeForEmployer) GetPlatformOk() (*IncludesId, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumeResumeForEmployer) SetPlatform(v IncludesId)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumeResumeForEmployer) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeResumeForEmployer) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeResumeForEmployer) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeResumeForEmployer) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeResumeForEmployer) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumeResumeForEmployer) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumeResumeForEmployer) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTotalExperience

`func (o *ResumeResumeForEmployer) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeResumeForEmployer) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeResumeForEmployer) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeResumeForEmployer) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeResumeForEmployer) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeResumeForEmployer) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumeResumeForEmployer) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumeResumeForEmployer) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumeResumeForEmployer) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBirthDate

`func (o *ResumeResumeForEmployer) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *ResumeResumeForEmployer) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *ResumeResumeForEmployer) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *ResumeResumeForEmployer) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *ResumeResumeForEmployer) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *ResumeResumeForEmployer) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetBusinessTripReadiness

`func (o *ResumeResumeForEmployer) GetBusinessTripReadiness() IncludesIdName`

GetBusinessTripReadiness returns the BusinessTripReadiness field if non-nil, zero value otherwise.

### GetBusinessTripReadinessOk

`func (o *ResumeResumeForEmployer) GetBusinessTripReadinessOk() (*IncludesIdName, bool)`

GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTripReadiness

`func (o *ResumeResumeForEmployer) SetBusinessTripReadiness(v IncludesIdName)`

SetBusinessTripReadiness sets BusinessTripReadiness field to given value.


### GetCitizenship

`func (o *ResumeResumeForEmployer) GetCitizenship() []IncludesIdNameUrl`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *ResumeResumeForEmployer) GetCitizenshipOk() (*[]IncludesIdNameUrl, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *ResumeResumeForEmployer) SetCitizenship(v []IncludesIdNameUrl)`

SetCitizenship sets Citizenship field to given value.


### GetContact

`func (o *ResumeResumeForEmployer) GetContact() []IncludesContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ResumeResumeForEmployer) GetContactOk() (*[]IncludesContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ResumeResumeForEmployer) SetContact(v []IncludesContact)`

SetContact sets Contact field to given value.


### GetCreds

`func (o *ResumeResumeForEmployer) GetCreds() CredsCreds`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *ResumeResumeForEmployer) GetCredsOk() (*CredsCreds, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *ResumeResumeForEmployer) SetCreds(v CredsCreds)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *ResumeResumeForEmployer) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### SetCredsNil

`func (o *ResumeResumeForEmployer) SetCredsNil(b bool)`

 SetCredsNil sets the value for Creds to be an explicit nil

### UnsetCreds
`func (o *ResumeResumeForEmployer) UnsetCreds()`

UnsetCreds ensures that no value is present for Creds, not even an explicit nil
### GetDriverLicenseTypes

`func (o *ResumeResumeForEmployer) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *ResumeResumeForEmployer) GetDriverLicenseTypesOk() (*[]ResumeObjectsDriverLicenseTypes, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *ResumeResumeForEmployer) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEmployment

`func (o *ResumeResumeForEmployer) GetEmployment() IncludesIdName`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *ResumeResumeForEmployer) GetEmploymentOk() (*IncludesIdName, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *ResumeResumeForEmployer) SetEmployment(v IncludesIdName)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *ResumeResumeForEmployer) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### GetEmployments

`func (o *ResumeResumeForEmployer) GetEmployments() []IncludesIdName`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *ResumeResumeForEmployer) GetEmploymentsOk() (*[]IncludesIdName, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *ResumeResumeForEmployer) SetEmployments(v []IncludesIdName)`

SetEmployments sets Employments field to given value.


### GetHasVehicle

`func (o *ResumeResumeForEmployer) GetHasVehicle() bool`

GetHasVehicle returns the HasVehicle field if non-nil, zero value otherwise.

### GetHasVehicleOk

`func (o *ResumeResumeForEmployer) GetHasVehicleOk() (*bool, bool)`

GetHasVehicleOk returns a tuple with the HasVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVehicle

`func (o *ResumeResumeForEmployer) SetHasVehicle(v bool)`

SetHasVehicle sets HasVehicle field to given value.

### HasHasVehicle

`func (o *ResumeResumeForEmployer) HasHasVehicle() bool`

HasHasVehicle returns a boolean if a field has been set.

### SetHasVehicleNil

`func (o *ResumeResumeForEmployer) SetHasVehicleNil(b bool)`

 SetHasVehicleNil sets the value for HasVehicle to be an explicit nil

### UnsetHasVehicle
`func (o *ResumeResumeForEmployer) UnsetHasVehicle()`

UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
### GetLanguage

`func (o *ResumeResumeForEmployer) GetLanguage() []IncludesLanguageLevel`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ResumeResumeForEmployer) GetLanguageOk() (*[]IncludesLanguageLevel, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ResumeResumeForEmployer) SetLanguage(v []IncludesLanguageLevel)`

SetLanguage sets Language field to given value.


### GetMetro

`func (o *ResumeResumeForEmployer) GetMetro() ResumeObjectsMetroStation`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *ResumeResumeForEmployer) GetMetroOk() (*ResumeObjectsMetroStation, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *ResumeResumeForEmployer) SetMetro(v ResumeObjectsMetroStation)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *ResumeResumeForEmployer) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### SetMetroNil

`func (o *ResumeResumeForEmployer) SetMetroNil(b bool)`

 SetMetroNil sets the value for Metro to be an explicit nil

### UnsetMetro
`func (o *ResumeResumeForEmployer) UnsetMetro()`

UnsetMetro ensures that no value is present for Metro, not even an explicit nil
### GetPaidServices

`func (o *ResumeResumeForEmployer) GetPaidServices() []ResumeObjectsEmployerPaidServicesInner`

GetPaidServices returns the PaidServices field if non-nil, zero value otherwise.

### GetPaidServicesOk

`func (o *ResumeResumeForEmployer) GetPaidServicesOk() (*[]ResumeObjectsEmployerPaidServicesInner, bool)`

GetPaidServicesOk returns a tuple with the PaidServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServices

`func (o *ResumeResumeForEmployer) SetPaidServices(v []ResumeObjectsEmployerPaidServicesInner)`

SetPaidServices sets PaidServices field to given value.


### GetProfessionalRoles

`func (o *ResumeResumeForEmployer) GetProfessionalRoles() []IncludesIdName`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *ResumeResumeForEmployer) GetProfessionalRolesOk() (*[]IncludesIdName, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *ResumeResumeForEmployer) SetProfessionalRoles(v []IncludesIdName)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *ResumeResumeForEmployer) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### SetProfessionalRolesNil

`func (o *ResumeResumeForEmployer) SetProfessionalRolesNil(b bool)`

 SetProfessionalRolesNil sets the value for ProfessionalRoles to be an explicit nil

### UnsetProfessionalRoles
`func (o *ResumeResumeForEmployer) UnsetProfessionalRoles()`

UnsetProfessionalRoles ensures that no value is present for ProfessionalRoles, not even an explicit nil
### GetRecommendation

`func (o *ResumeResumeForEmployer) GetRecommendation() []ResumeObjectsRecommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *ResumeResumeForEmployer) GetRecommendationOk() (*[]ResumeObjectsRecommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *ResumeResumeForEmployer) SetRecommendation(v []ResumeObjectsRecommendation)`

SetRecommendation sets Recommendation field to given value.


### GetRelocation

`func (o *ResumeResumeForEmployer) GetRelocation() ResumeObjectsRelocationPublic`

GetRelocation returns the Relocation field if non-nil, zero value otherwise.

### GetRelocationOk

`func (o *ResumeResumeForEmployer) GetRelocationOk() (*ResumeObjectsRelocationPublic, bool)`

GetRelocationOk returns a tuple with the Relocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocation

`func (o *ResumeResumeForEmployer) SetRelocation(v ResumeObjectsRelocationPublic)`

SetRelocation sets Relocation field to given value.


### GetResumeLocale

`func (o *ResumeResumeForEmployer) GetResumeLocale() IncludesIdName`

GetResumeLocale returns the ResumeLocale field if non-nil, zero value otherwise.

### GetResumeLocaleOk

`func (o *ResumeResumeForEmployer) GetResumeLocaleOk() (*IncludesIdName, bool)`

GetResumeLocaleOk returns a tuple with the ResumeLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeLocale

`func (o *ResumeResumeForEmployer) SetResumeLocale(v IncludesIdName)`

SetResumeLocale sets ResumeLocale field to given value.


### GetSchedule

`func (o *ResumeResumeForEmployer) GetSchedule() IncludesIdName`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ResumeResumeForEmployer) GetScheduleOk() (*IncludesIdName, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ResumeResumeForEmployer) SetSchedule(v IncludesIdName)`

SetSchedule sets Schedule field to given value.


### GetSchedules

`func (o *ResumeResumeForEmployer) GetSchedules() []IncludesIdName`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ResumeResumeForEmployer) GetSchedulesOk() (*[]IncludesIdName, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ResumeResumeForEmployer) SetSchedules(v []IncludesIdName)`

SetSchedules sets Schedules field to given value.


### GetSite

`func (o *ResumeResumeForEmployer) GetSite() []ResumeObjectsSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ResumeResumeForEmployer) GetSiteOk() (*[]ResumeObjectsSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ResumeResumeForEmployer) SetSite(v []ResumeObjectsSite)`

SetSite sets Site field to given value.


### GetSkillSet

`func (o *ResumeResumeForEmployer) GetSkillSet() []string`

GetSkillSet returns the SkillSet field if non-nil, zero value otherwise.

### GetSkillSetOk

`func (o *ResumeResumeForEmployer) GetSkillSetOk() (*[]string, bool)`

GetSkillSetOk returns a tuple with the SkillSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillSet

`func (o *ResumeResumeForEmployer) SetSkillSet(v []string)`

SetSkillSet sets SkillSet field to given value.


### GetSkills

`func (o *ResumeResumeForEmployer) GetSkills() string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ResumeResumeForEmployer) GetSkillsOk() (*string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ResumeResumeForEmployer) SetSkills(v string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ResumeResumeForEmployer) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkillsNil

`func (o *ResumeResumeForEmployer) SetSkillsNil(b bool)`

 SetSkillsNil sets the value for Skills to be an explicit nil

### UnsetSkills
`func (o *ResumeResumeForEmployer) UnsetSkills()`

UnsetSkills ensures that no value is present for Skills, not even an explicit nil
### GetTravelTime

`func (o *ResumeResumeForEmployer) GetTravelTime() IncludesIdName`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *ResumeResumeForEmployer) GetTravelTimeOk() (*IncludesIdName, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *ResumeResumeForEmployer) SetTravelTime(v IncludesIdName)`

SetTravelTime sets TravelTime field to given value.


### GetWorkTicket

`func (o *ResumeResumeForEmployer) GetWorkTicket() []IncludesIdNameUrl`

GetWorkTicket returns the WorkTicket field if non-nil, zero value otherwise.

### GetWorkTicketOk

`func (o *ResumeResumeForEmployer) GetWorkTicketOk() (*[]IncludesIdNameUrl, bool)`

GetWorkTicketOk returns a tuple with the WorkTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTicket

`func (o *ResumeResumeForEmployer) SetWorkTicket(v []IncludesIdNameUrl)`

SetWorkTicket sets WorkTicket field to given value.


### GetActions

`func (o *ResumeResumeForEmployer) GetActions() ResumeObjectsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumeResumeForEmployer) GetActionsOk() (*ResumeObjectsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumeResumeForEmployer) SetActions(v ResumeObjectsActions)`

SetActions sets Actions field to given value.


### GetFavorited

`func (o *ResumeResumeForEmployer) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *ResumeResumeForEmployer) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *ResumeResumeForEmployer) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetJobSearchStatus

`func (o *ResumeResumeForEmployer) GetJobSearchStatus() IncludesIdNameLastChangeTime`

GetJobSearchStatus returns the JobSearchStatus field if non-nil, zero value otherwise.

### GetJobSearchStatusOk

`func (o *ResumeResumeForEmployer) GetJobSearchStatusOk() (*IncludesIdNameLastChangeTime, bool)`

GetJobSearchStatusOk returns a tuple with the JobSearchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatus

`func (o *ResumeResumeForEmployer) SetJobSearchStatus(v IncludesIdNameLastChangeTime)`

SetJobSearchStatus sets JobSearchStatus field to given value.

### HasJobSearchStatus

`func (o *ResumeResumeForEmployer) HasJobSearchStatus() bool`

HasJobSearchStatus returns a boolean if a field has been set.

### GetNegotiationsHistory

`func (o *ResumeResumeForEmployer) GetNegotiationsHistory() ResumeObjectsNegotiationsHistoryForEmployer`

GetNegotiationsHistory returns the NegotiationsHistory field if non-nil, zero value otherwise.

### GetNegotiationsHistoryOk

`func (o *ResumeResumeForEmployer) GetNegotiationsHistoryOk() (*ResumeObjectsNegotiationsHistoryForEmployer, bool)`

GetNegotiationsHistoryOk returns a tuple with the NegotiationsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsHistory

`func (o *ResumeResumeForEmployer) SetNegotiationsHistory(v ResumeObjectsNegotiationsHistoryForEmployer)`

SetNegotiationsHistory sets NegotiationsHistory field to given value.


### GetOwner

`func (o *ResumeResumeForEmployer) GetOwner() ResumeObjectsOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResumeResumeForEmployer) GetOwnerOk() (*ResumeObjectsOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResumeResumeForEmployer) SetOwner(v ResumeObjectsOwner)`

SetOwner sets Owner field to given value.


### GetPhoto

`func (o *ResumeResumeForEmployer) GetPhoto() ResumeObjectsPhotoNoId`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeResumeForEmployer) GetPhotoOk() (*ResumeObjectsPhotoNoId, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeResumeForEmployer) SetPhoto(v ResumeObjectsPhotoNoId)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeResumeForEmployer) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeResumeForEmployer) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeResumeForEmployer) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPortfolio

`func (o *ResumeResumeForEmployer) GetPortfolio() []ResumeObjectsPortfolioNoId`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *ResumeResumeForEmployer) GetPortfolioOk() (*[]ResumeObjectsPortfolioNoId, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *ResumeResumeForEmployer) SetPortfolio(v []ResumeObjectsPortfolioNoId)`

SetPortfolio sets Portfolio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


