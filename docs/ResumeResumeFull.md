# ResumeResumeFull

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
**PaidServices** | [**[]ResumeObjectsPaidServices**](ResumeObjectsPaidServices.md) | Платные услуги по резюме | 
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

## Methods

### NewResumeResumeFull

`func NewResumeResumeFull(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download ResumeResumeProfileAllOfDownload, education ResumeResumeProfileAllOfEducation, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, businessTripReadiness ResumeResumeFullAllOfBusinessTripReadiness, citizenship []IncludesIdNameUrl, contact []IncludesContact, driverLicenseTypes []ResumeObjectsDriverLicenseTypes, employments []IncludesIdName, language []IncludesLanguageLevel, paidServices []ResumeObjectsPaidServices, recommendation []ResumeObjectsRecommendation, relocation ResumeResumeFullAllOfRelocation, resumeLocale ResumeResumeFullAllOfResumeLocale, schedule ResumeResumeFullAllOfSchedule, schedules []IncludesIdName, site []ResumeObjectsSite, skillSet []string, travelTime ResumeResumeFullAllOfTravelTime, workTicket []IncludesIdNameUrl, ) *ResumeResumeFull`

NewResumeResumeFull instantiates a new ResumeResumeFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeFullWithDefaults

`func NewResumeResumeFullWithDefaults() *ResumeResumeFull`

NewResumeResumeFullWithDefaults instantiates a new ResumeResumeFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *ResumeResumeFull) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumeResumeFull) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumeResumeFull) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *ResumeResumeFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeResumeFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeResumeFull) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ResumeResumeFull) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeResumeFull) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeResumeFull) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ResumeResumeFull) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeResumeFull) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAge

`func (o *ResumeResumeFull) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumeResumeFull) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumeResumeFull) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumeResumeFull) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumeResumeFull) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumeResumeFull) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetArea

`func (o *ResumeResumeFull) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeResumeFull) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeResumeFull) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeResumeFull) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeResumeFull) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeResumeFull) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumeResumeFull) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumeResumeFull) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumeResumeFull) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumeResumeFull) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumeResumeFull) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumeResumeFull) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumeResumeFull) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeResumeFull) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeResumeFull) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumeResumeFull) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumeResumeFull) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumeResumeFull) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumeResumeFull) GetDownload() ResumeResumeProfileAllOfDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeResumeFull) GetDownloadOk() (*ResumeResumeProfileAllOfDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeResumeFull) SetDownload(v ResumeResumeProfileAllOfDownload)`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumeResumeFull) GetEducation() ResumeResumeProfileAllOfEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeResumeFull) GetEducationOk() (*ResumeResumeProfileAllOfEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeResumeFull) SetEducation(v ResumeResumeProfileAllOfEducation)`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumeResumeFull) GetExperience() []ResumeObjectsExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeResumeFull) GetExperienceOk() (*[]ResumeObjectsExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeResumeFull) SetExperience(v []ResumeObjectsExperience)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumeResumeFull) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeResumeFull) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeResumeFull) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeResumeFull) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeResumeFull) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeResumeFull) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumeResumeFull) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeResumeFull) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeResumeFull) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeResumeFull) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ResumeResumeFull) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ResumeResumeFull) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *ResumeResumeFull) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeResumeFull) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeResumeFull) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetLastName

`func (o *ResumeResumeFull) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeResumeFull) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeResumeFull) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeResumeFull) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeResumeFull) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeResumeFull) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumeResumeFull) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumeResumeFull) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumeResumeFull) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *ResumeResumeFull) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetMiddleName

`func (o *ResumeResumeFull) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeResumeFull) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeResumeFull) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeResumeFull) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeResumeFull) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeResumeFull) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPlatform

`func (o *ResumeResumeFull) GetPlatform() ResumeResumeProfileAllOfPlatform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumeResumeFull) GetPlatformOk() (*ResumeResumeProfileAllOfPlatform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumeResumeFull) SetPlatform(v ResumeResumeProfileAllOfPlatform)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumeResumeFull) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeResumeFull) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeResumeFull) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeResumeFull) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeResumeFull) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumeResumeFull) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumeResumeFull) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTotalExperience

`func (o *ResumeResumeFull) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeResumeFull) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeResumeFull) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeResumeFull) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeResumeFull) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeResumeFull) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumeResumeFull) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumeResumeFull) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumeResumeFull) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBirthDate

`func (o *ResumeResumeFull) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *ResumeResumeFull) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *ResumeResumeFull) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *ResumeResumeFull) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *ResumeResumeFull) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *ResumeResumeFull) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetBusinessTripReadiness

`func (o *ResumeResumeFull) GetBusinessTripReadiness() ResumeResumeFullAllOfBusinessTripReadiness`

GetBusinessTripReadiness returns the BusinessTripReadiness field if non-nil, zero value otherwise.

### GetBusinessTripReadinessOk

`func (o *ResumeResumeFull) GetBusinessTripReadinessOk() (*ResumeResumeFullAllOfBusinessTripReadiness, bool)`

GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTripReadiness

`func (o *ResumeResumeFull) SetBusinessTripReadiness(v ResumeResumeFullAllOfBusinessTripReadiness)`

SetBusinessTripReadiness sets BusinessTripReadiness field to given value.


### GetCitizenship

`func (o *ResumeResumeFull) GetCitizenship() []IncludesIdNameUrl`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *ResumeResumeFull) GetCitizenshipOk() (*[]IncludesIdNameUrl, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *ResumeResumeFull) SetCitizenship(v []IncludesIdNameUrl)`

SetCitizenship sets Citizenship field to given value.


### GetContact

`func (o *ResumeResumeFull) GetContact() []IncludesContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ResumeResumeFull) GetContactOk() (*[]IncludesContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ResumeResumeFull) SetContact(v []IncludesContact)`

SetContact sets Contact field to given value.


### GetCreds

`func (o *ResumeResumeFull) GetCreds() CredsCreds`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *ResumeResumeFull) GetCredsOk() (*CredsCreds, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *ResumeResumeFull) SetCreds(v CredsCreds)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *ResumeResumeFull) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### SetCredsNil

`func (o *ResumeResumeFull) SetCredsNil(b bool)`

 SetCredsNil sets the value for Creds to be an explicit nil

### UnsetCreds
`func (o *ResumeResumeFull) UnsetCreds()`

UnsetCreds ensures that no value is present for Creds, not even an explicit nil
### GetDriverLicenseTypes

`func (o *ResumeResumeFull) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *ResumeResumeFull) GetDriverLicenseTypesOk() (*[]ResumeObjectsDriverLicenseTypes, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *ResumeResumeFull) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEmployment

`func (o *ResumeResumeFull) GetEmployment() ResumeResumeFullAllOfEmployment`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *ResumeResumeFull) GetEmploymentOk() (*ResumeResumeFullAllOfEmployment, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *ResumeResumeFull) SetEmployment(v ResumeResumeFullAllOfEmployment)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *ResumeResumeFull) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### GetEmployments

`func (o *ResumeResumeFull) GetEmployments() []IncludesIdName`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *ResumeResumeFull) GetEmploymentsOk() (*[]IncludesIdName, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *ResumeResumeFull) SetEmployments(v []IncludesIdName)`

SetEmployments sets Employments field to given value.


### GetHasVehicle

`func (o *ResumeResumeFull) GetHasVehicle() bool`

GetHasVehicle returns the HasVehicle field if non-nil, zero value otherwise.

### GetHasVehicleOk

`func (o *ResumeResumeFull) GetHasVehicleOk() (*bool, bool)`

GetHasVehicleOk returns a tuple with the HasVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVehicle

`func (o *ResumeResumeFull) SetHasVehicle(v bool)`

SetHasVehicle sets HasVehicle field to given value.

### HasHasVehicle

`func (o *ResumeResumeFull) HasHasVehicle() bool`

HasHasVehicle returns a boolean if a field has been set.

### SetHasVehicleNil

`func (o *ResumeResumeFull) SetHasVehicleNil(b bool)`

 SetHasVehicleNil sets the value for HasVehicle to be an explicit nil

### UnsetHasVehicle
`func (o *ResumeResumeFull) UnsetHasVehicle()`

UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
### GetLanguage

`func (o *ResumeResumeFull) GetLanguage() []IncludesLanguageLevel`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ResumeResumeFull) GetLanguageOk() (*[]IncludesLanguageLevel, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ResumeResumeFull) SetLanguage(v []IncludesLanguageLevel)`

SetLanguage sets Language field to given value.


### GetMetro

`func (o *ResumeResumeFull) GetMetro() ResumeObjectsMetroStation`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *ResumeResumeFull) GetMetroOk() (*ResumeObjectsMetroStation, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *ResumeResumeFull) SetMetro(v ResumeObjectsMetroStation)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *ResumeResumeFull) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### SetMetroNil

`func (o *ResumeResumeFull) SetMetroNil(b bool)`

 SetMetroNil sets the value for Metro to be an explicit nil

### UnsetMetro
`func (o *ResumeResumeFull) UnsetMetro()`

UnsetMetro ensures that no value is present for Metro, not even an explicit nil
### GetPaidServices

`func (o *ResumeResumeFull) GetPaidServices() []ResumeObjectsPaidServices`

GetPaidServices returns the PaidServices field if non-nil, zero value otherwise.

### GetPaidServicesOk

`func (o *ResumeResumeFull) GetPaidServicesOk() (*[]ResumeObjectsPaidServices, bool)`

GetPaidServicesOk returns a tuple with the PaidServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServices

`func (o *ResumeResumeFull) SetPaidServices(v []ResumeObjectsPaidServices)`

SetPaidServices sets PaidServices field to given value.


### GetProfessionalRoles

`func (o *ResumeResumeFull) GetProfessionalRoles() []IncludesIdName`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *ResumeResumeFull) GetProfessionalRolesOk() (*[]IncludesIdName, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *ResumeResumeFull) SetProfessionalRoles(v []IncludesIdName)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *ResumeResumeFull) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### SetProfessionalRolesNil

`func (o *ResumeResumeFull) SetProfessionalRolesNil(b bool)`

 SetProfessionalRolesNil sets the value for ProfessionalRoles to be an explicit nil

### UnsetProfessionalRoles
`func (o *ResumeResumeFull) UnsetProfessionalRoles()`

UnsetProfessionalRoles ensures that no value is present for ProfessionalRoles, not even an explicit nil
### GetRecommendation

`func (o *ResumeResumeFull) GetRecommendation() []ResumeObjectsRecommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *ResumeResumeFull) GetRecommendationOk() (*[]ResumeObjectsRecommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *ResumeResumeFull) SetRecommendation(v []ResumeObjectsRecommendation)`

SetRecommendation sets Recommendation field to given value.


### GetRelocation

`func (o *ResumeResumeFull) GetRelocation() ResumeResumeFullAllOfRelocation`

GetRelocation returns the Relocation field if non-nil, zero value otherwise.

### GetRelocationOk

`func (o *ResumeResumeFull) GetRelocationOk() (*ResumeResumeFullAllOfRelocation, bool)`

GetRelocationOk returns a tuple with the Relocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocation

`func (o *ResumeResumeFull) SetRelocation(v ResumeResumeFullAllOfRelocation)`

SetRelocation sets Relocation field to given value.


### GetResumeLocale

`func (o *ResumeResumeFull) GetResumeLocale() ResumeResumeFullAllOfResumeLocale`

GetResumeLocale returns the ResumeLocale field if non-nil, zero value otherwise.

### GetResumeLocaleOk

`func (o *ResumeResumeFull) GetResumeLocaleOk() (*ResumeResumeFullAllOfResumeLocale, bool)`

GetResumeLocaleOk returns a tuple with the ResumeLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeLocale

`func (o *ResumeResumeFull) SetResumeLocale(v ResumeResumeFullAllOfResumeLocale)`

SetResumeLocale sets ResumeLocale field to given value.


### GetSchedule

`func (o *ResumeResumeFull) GetSchedule() ResumeResumeFullAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ResumeResumeFull) GetScheduleOk() (*ResumeResumeFullAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ResumeResumeFull) SetSchedule(v ResumeResumeFullAllOfSchedule)`

SetSchedule sets Schedule field to given value.


### GetSchedules

`func (o *ResumeResumeFull) GetSchedules() []IncludesIdName`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ResumeResumeFull) GetSchedulesOk() (*[]IncludesIdName, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ResumeResumeFull) SetSchedules(v []IncludesIdName)`

SetSchedules sets Schedules field to given value.


### GetSite

`func (o *ResumeResumeFull) GetSite() []ResumeObjectsSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ResumeResumeFull) GetSiteOk() (*[]ResumeObjectsSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ResumeResumeFull) SetSite(v []ResumeObjectsSite)`

SetSite sets Site field to given value.


### GetSkillSet

`func (o *ResumeResumeFull) GetSkillSet() []string`

GetSkillSet returns the SkillSet field if non-nil, zero value otherwise.

### GetSkillSetOk

`func (o *ResumeResumeFull) GetSkillSetOk() (*[]string, bool)`

GetSkillSetOk returns a tuple with the SkillSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillSet

`func (o *ResumeResumeFull) SetSkillSet(v []string)`

SetSkillSet sets SkillSet field to given value.


### GetSkills

`func (o *ResumeResumeFull) GetSkills() string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ResumeResumeFull) GetSkillsOk() (*string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ResumeResumeFull) SetSkills(v string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ResumeResumeFull) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkillsNil

`func (o *ResumeResumeFull) SetSkillsNil(b bool)`

 SetSkillsNil sets the value for Skills to be an explicit nil

### UnsetSkills
`func (o *ResumeResumeFull) UnsetSkills()`

UnsetSkills ensures that no value is present for Skills, not even an explicit nil
### GetTravelTime

`func (o *ResumeResumeFull) GetTravelTime() ResumeResumeFullAllOfTravelTime`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *ResumeResumeFull) GetTravelTimeOk() (*ResumeResumeFullAllOfTravelTime, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *ResumeResumeFull) SetTravelTime(v ResumeResumeFullAllOfTravelTime)`

SetTravelTime sets TravelTime field to given value.


### GetWorkTicket

`func (o *ResumeResumeFull) GetWorkTicket() []IncludesIdNameUrl`

GetWorkTicket returns the WorkTicket field if non-nil, zero value otherwise.

### GetWorkTicketOk

`func (o *ResumeResumeFull) GetWorkTicketOk() (*[]IncludesIdNameUrl, bool)`

GetWorkTicketOk returns a tuple with the WorkTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTicket

`func (o *ResumeResumeFull) SetWorkTicket(v []IncludesIdNameUrl)`

SetWorkTicket sets WorkTicket field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


