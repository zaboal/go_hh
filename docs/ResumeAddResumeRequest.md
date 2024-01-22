# ResumeAddResumeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**ResumeObjectsAccess**](ResumeObjectsAccess.md) |  | [optional] 
**BirthDate** | Pointer to **NullableString** | День рождения (в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**BusinessTripReadiness** | Pointer to [**IncludesId**](IncludesId.md) |  | [optional] 
**Certificate** | Pointer to [**[]ResumeObjectsCertificate**](ResumeObjectsCertificate.md) | Список сертификатов соискателя | [optional] 
**DriverLicenseTypes** | Pointer to [**[]ResumeObjectsDriverLicenseTypes**](ResumeObjectsDriverLicenseTypes.md) | Список категорий водительских прав соискателя | [optional] 
**Employments** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Список подходящих соискателю типов занятостей. Элементы справочника [employment](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**FirstName** | Pointer to **NullableString** | Имя | [optional] 
**HasVehicle** | Pointer to **NullableBool** | Наличие личного автомобиля у соискателя | [optional] 
**HiddenFields** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Документация [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле &#x60;resume_hidden_fields&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**LastName** | Pointer to **NullableString** | Фамилия | [optional] 
**Metro** | Pointer to [**IncludesId**](IncludesId.md) |  | [optional] 
**MiddleName** | Pointer to **NullableString** | Отчество | [optional] 
**Photo** | Pointer to [**NullableResumeObjectsPhoto**](ResumeObjectsPhoto.md) |  | [optional] 
**Portfolio** | Pointer to [**[]ResumeObjectsPortfolio**](ResumeObjectsPortfolio.md) | Список изображений в портфолио пользователя | [optional] 
**ProfessionalRoles** | Pointer to [**[]IncludesId**](IncludesId.md) | Массив объектов профролей. Элемент справочника [professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary) | [optional] 
**Recommendation** | Pointer to [**[]ResumeObjectsRecommendation**](ResumeObjectsRecommendation.md) | Список рекомендаций | [optional] 
**Relocation** | Pointer to [**ResumeObjectsRelocationPublic**](ResumeObjectsRelocationPublic.md) |  | [optional] 
**ResumeLocale** | Pointer to [**IncludesIdName**](IncludesIdName.md) |  | [optional] 
**Salary** | Pointer to [**ResumeObjectsSalaryAddEdit**](ResumeObjectsSalaryAddEdit.md) |  | [optional] 
**Schedules** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Список подходящих соискателю графиков работы. Элементы справочника [schedule](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**Site** | Pointer to [**[]ResumeObjectsSite**](ResumeObjectsSite.md) | Профили в соц. сетях и других сервисах | [optional] 
**SkillSet** | Pointer to **[]string** | Ключевые навыки (список уникальных строк) | [optional] 
**Skills** | Pointer to **NullableString** | Дополнительная информация, описание навыков в свободной форме | [optional] 
**Title** | Pointer to **NullableString** | Желаемая должность | [optional] 
**TotalExperience** | Pointer to [**NullableResumeObjectsTotalExperience**](ResumeObjectsTotalExperience.md) |  | [optional] 
**TravelTime** | Pointer to [**IncludesId**](IncludesId.md) |  | [optional] 
**WorkTicket** | Pointer to [**[]IncludesId**](IncludesId.md) | Список регионов, в который соискатель имеет разрешение на работу. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas)  | [optional] 
**Area** | Pointer to [**ResumeAddResumeRequestAllOfArea**](ResumeAddResumeRequestAllOfArea.md) |  | [optional] 
**Citizenship** | Pointer to [**[]IncludesId**](IncludesId.md) | Список гражданств соискателя. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas) | [optional] 
**Contact** | Pointer to [**[]ResumeObjectsContact**](ResumeObjectsContact.md) | Список контактов соискателя.  При заполнении контактов в резюме необходимо учитывать следующие условия:  * В резюме обязательно должен быть указан e-mail. Он может быть только один. * В резюме должен быть указан хотя бы один телефон, причём можно указывать только один телефон каждого типа. * Комментарий можно указывать только для телефонов, для e-mail комментарий не сохранится. * Обязательно указать либо телефон полностью в поле &#x60;formatted&#x60;, либо все три части телефона по отдельности в трёх полях: &#x60;country&#x60;, &#x60;city&#x60; и &#x60;number&#x60;. Если указано и то, и то, используются данные из трёх полей. В поле &#x60;formatted&#x60; допустимо использовать пробелы, скобки и дефисы. В остальных полях допустимы только цифры  | [optional] 
**Education** | Pointer to [**ResumeAddResumeRequestAllOfEducation**](ResumeAddResumeRequestAllOfEducation.md) |  | [optional] 
**Experience** | Pointer to [**[]ResumeObjectsExperienceCreateEditResume**](ResumeObjectsExperienceCreateEditResume.md) | Опыт работы | [optional] 
**Gender** | Pointer to [**ResumeAddResumeRequestAllOfGender**](ResumeAddResumeRequestAllOfGender.md) |  | [optional] 
**Language** | Pointer to [**[]ResumeObjectsLanguage**](ResumeObjectsLanguage.md) | Список языков, которыми владеет соискатель. Элементы справочника [languages](#tag/Obshie-spravochniki/operation/get-languages) | [optional] 

## Methods

### NewResumeAddResumeRequest

`func NewResumeAddResumeRequest() *ResumeAddResumeRequest`

NewResumeAddResumeRequest instantiates a new ResumeAddResumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeAddResumeRequestWithDefaults

`func NewResumeAddResumeRequestWithDefaults() *ResumeAddResumeRequest`

NewResumeAddResumeRequestWithDefaults instantiates a new ResumeAddResumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ResumeAddResumeRequest) GetAccess() ResumeObjectsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ResumeAddResumeRequest) GetAccessOk() (*ResumeObjectsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ResumeAddResumeRequest) SetAccess(v ResumeObjectsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ResumeAddResumeRequest) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetBirthDate

`func (o *ResumeAddResumeRequest) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *ResumeAddResumeRequest) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *ResumeAddResumeRequest) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *ResumeAddResumeRequest) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *ResumeAddResumeRequest) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *ResumeAddResumeRequest) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetBusinessTripReadiness

`func (o *ResumeAddResumeRequest) GetBusinessTripReadiness() IncludesId`

GetBusinessTripReadiness returns the BusinessTripReadiness field if non-nil, zero value otherwise.

### GetBusinessTripReadinessOk

`func (o *ResumeAddResumeRequest) GetBusinessTripReadinessOk() (*IncludesId, bool)`

GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTripReadiness

`func (o *ResumeAddResumeRequest) SetBusinessTripReadiness(v IncludesId)`

SetBusinessTripReadiness sets BusinessTripReadiness field to given value.

### HasBusinessTripReadiness

`func (o *ResumeAddResumeRequest) HasBusinessTripReadiness() bool`

HasBusinessTripReadiness returns a boolean if a field has been set.

### GetCertificate

`func (o *ResumeAddResumeRequest) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeAddResumeRequest) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeAddResumeRequest) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ResumeAddResumeRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ResumeAddResumeRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ResumeAddResumeRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetDriverLicenseTypes

`func (o *ResumeAddResumeRequest) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *ResumeAddResumeRequest) GetDriverLicenseTypesOk() (*[]ResumeObjectsDriverLicenseTypes, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *ResumeAddResumeRequest) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.

### HasDriverLicenseTypes

`func (o *ResumeAddResumeRequest) HasDriverLicenseTypes() bool`

HasDriverLicenseTypes returns a boolean if a field has been set.

### SetDriverLicenseTypesNil

`func (o *ResumeAddResumeRequest) SetDriverLicenseTypesNil(b bool)`

 SetDriverLicenseTypesNil sets the value for DriverLicenseTypes to be an explicit nil

### UnsetDriverLicenseTypes
`func (o *ResumeAddResumeRequest) UnsetDriverLicenseTypes()`

UnsetDriverLicenseTypes ensures that no value is present for DriverLicenseTypes, not even an explicit nil
### GetEmployments

`func (o *ResumeAddResumeRequest) GetEmployments() []IncludesIdName`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *ResumeAddResumeRequest) GetEmploymentsOk() (*[]IncludesIdName, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *ResumeAddResumeRequest) SetEmployments(v []IncludesIdName)`

SetEmployments sets Employments field to given value.

### HasEmployments

`func (o *ResumeAddResumeRequest) HasEmployments() bool`

HasEmployments returns a boolean if a field has been set.

### SetEmploymentsNil

`func (o *ResumeAddResumeRequest) SetEmploymentsNil(b bool)`

 SetEmploymentsNil sets the value for Employments to be an explicit nil

### UnsetEmployments
`func (o *ResumeAddResumeRequest) UnsetEmployments()`

UnsetEmployments ensures that no value is present for Employments, not even an explicit nil
### GetFirstName

`func (o *ResumeAddResumeRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeAddResumeRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeAddResumeRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeAddResumeRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeAddResumeRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeAddResumeRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetHasVehicle

`func (o *ResumeAddResumeRequest) GetHasVehicle() bool`

GetHasVehicle returns the HasVehicle field if non-nil, zero value otherwise.

### GetHasVehicleOk

`func (o *ResumeAddResumeRequest) GetHasVehicleOk() (*bool, bool)`

GetHasVehicleOk returns a tuple with the HasVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVehicle

`func (o *ResumeAddResumeRequest) SetHasVehicle(v bool)`

SetHasVehicle sets HasVehicle field to given value.

### HasHasVehicle

`func (o *ResumeAddResumeRequest) HasHasVehicle() bool`

HasHasVehicle returns a boolean if a field has been set.

### SetHasVehicleNil

`func (o *ResumeAddResumeRequest) SetHasVehicleNil(b bool)`

 SetHasVehicleNil sets the value for HasVehicle to be an explicit nil

### UnsetHasVehicle
`func (o *ResumeAddResumeRequest) UnsetHasVehicle()`

UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
### GetHiddenFields

`func (o *ResumeAddResumeRequest) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeAddResumeRequest) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeAddResumeRequest) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.

### HasHiddenFields

`func (o *ResumeAddResumeRequest) HasHiddenFields() bool`

HasHiddenFields returns a boolean if a field has been set.

### SetHiddenFieldsNil

`func (o *ResumeAddResumeRequest) SetHiddenFieldsNil(b bool)`

 SetHiddenFieldsNil sets the value for HiddenFields to be an explicit nil

### UnsetHiddenFields
`func (o *ResumeAddResumeRequest) UnsetHiddenFields()`

UnsetHiddenFields ensures that no value is present for HiddenFields, not even an explicit nil
### GetLastName

`func (o *ResumeAddResumeRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeAddResumeRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeAddResumeRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeAddResumeRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeAddResumeRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeAddResumeRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMetro

`func (o *ResumeAddResumeRequest) GetMetro() IncludesId`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *ResumeAddResumeRequest) GetMetroOk() (*IncludesId, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *ResumeAddResumeRequest) SetMetro(v IncludesId)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *ResumeAddResumeRequest) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetMiddleName

`func (o *ResumeAddResumeRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeAddResumeRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeAddResumeRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeAddResumeRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeAddResumeRequest) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeAddResumeRequest) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPhoto

`func (o *ResumeAddResumeRequest) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeAddResumeRequest) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeAddResumeRequest) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeAddResumeRequest) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeAddResumeRequest) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeAddResumeRequest) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPortfolio

`func (o *ResumeAddResumeRequest) GetPortfolio() []ResumeObjectsPortfolio`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *ResumeAddResumeRequest) GetPortfolioOk() (*[]ResumeObjectsPortfolio, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *ResumeAddResumeRequest) SetPortfolio(v []ResumeObjectsPortfolio)`

SetPortfolio sets Portfolio field to given value.

### HasPortfolio

`func (o *ResumeAddResumeRequest) HasPortfolio() bool`

HasPortfolio returns a boolean if a field has been set.

### SetPortfolioNil

`func (o *ResumeAddResumeRequest) SetPortfolioNil(b bool)`

 SetPortfolioNil sets the value for Portfolio to be an explicit nil

### UnsetPortfolio
`func (o *ResumeAddResumeRequest) UnsetPortfolio()`

UnsetPortfolio ensures that no value is present for Portfolio, not even an explicit nil
### GetProfessionalRoles

`func (o *ResumeAddResumeRequest) GetProfessionalRoles() []IncludesId`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *ResumeAddResumeRequest) GetProfessionalRolesOk() (*[]IncludesId, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *ResumeAddResumeRequest) SetProfessionalRoles(v []IncludesId)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *ResumeAddResumeRequest) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetRecommendation

`func (o *ResumeAddResumeRequest) GetRecommendation() []ResumeObjectsRecommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *ResumeAddResumeRequest) GetRecommendationOk() (*[]ResumeObjectsRecommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *ResumeAddResumeRequest) SetRecommendation(v []ResumeObjectsRecommendation)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *ResumeAddResumeRequest) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### SetRecommendationNil

`func (o *ResumeAddResumeRequest) SetRecommendationNil(b bool)`

 SetRecommendationNil sets the value for Recommendation to be an explicit nil

### UnsetRecommendation
`func (o *ResumeAddResumeRequest) UnsetRecommendation()`

UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
### GetRelocation

`func (o *ResumeAddResumeRequest) GetRelocation() ResumeObjectsRelocationPublic`

GetRelocation returns the Relocation field if non-nil, zero value otherwise.

### GetRelocationOk

`func (o *ResumeAddResumeRequest) GetRelocationOk() (*ResumeObjectsRelocationPublic, bool)`

GetRelocationOk returns a tuple with the Relocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocation

`func (o *ResumeAddResumeRequest) SetRelocation(v ResumeObjectsRelocationPublic)`

SetRelocation sets Relocation field to given value.

### HasRelocation

`func (o *ResumeAddResumeRequest) HasRelocation() bool`

HasRelocation returns a boolean if a field has been set.

### GetResumeLocale

`func (o *ResumeAddResumeRequest) GetResumeLocale() IncludesIdName`

GetResumeLocale returns the ResumeLocale field if non-nil, zero value otherwise.

### GetResumeLocaleOk

`func (o *ResumeAddResumeRequest) GetResumeLocaleOk() (*IncludesIdName, bool)`

GetResumeLocaleOk returns a tuple with the ResumeLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeLocale

`func (o *ResumeAddResumeRequest) SetResumeLocale(v IncludesIdName)`

SetResumeLocale sets ResumeLocale field to given value.

### HasResumeLocale

`func (o *ResumeAddResumeRequest) HasResumeLocale() bool`

HasResumeLocale returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeAddResumeRequest) GetSalary() ResumeObjectsSalaryAddEdit`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeAddResumeRequest) GetSalaryOk() (*ResumeObjectsSalaryAddEdit, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeAddResumeRequest) SetSalary(v ResumeObjectsSalaryAddEdit)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeAddResumeRequest) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### GetSchedules

`func (o *ResumeAddResumeRequest) GetSchedules() []IncludesIdName`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ResumeAddResumeRequest) GetSchedulesOk() (*[]IncludesIdName, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ResumeAddResumeRequest) SetSchedules(v []IncludesIdName)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ResumeAddResumeRequest) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### SetSchedulesNil

`func (o *ResumeAddResumeRequest) SetSchedulesNil(b bool)`

 SetSchedulesNil sets the value for Schedules to be an explicit nil

### UnsetSchedules
`func (o *ResumeAddResumeRequest) UnsetSchedules()`

UnsetSchedules ensures that no value is present for Schedules, not even an explicit nil
### GetSite

`func (o *ResumeAddResumeRequest) GetSite() []ResumeObjectsSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ResumeAddResumeRequest) GetSiteOk() (*[]ResumeObjectsSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ResumeAddResumeRequest) SetSite(v []ResumeObjectsSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *ResumeAddResumeRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *ResumeAddResumeRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *ResumeAddResumeRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetSkillSet

`func (o *ResumeAddResumeRequest) GetSkillSet() []string`

GetSkillSet returns the SkillSet field if non-nil, zero value otherwise.

### GetSkillSetOk

`func (o *ResumeAddResumeRequest) GetSkillSetOk() (*[]string, bool)`

GetSkillSetOk returns a tuple with the SkillSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillSet

`func (o *ResumeAddResumeRequest) SetSkillSet(v []string)`

SetSkillSet sets SkillSet field to given value.

### HasSkillSet

`func (o *ResumeAddResumeRequest) HasSkillSet() bool`

HasSkillSet returns a boolean if a field has been set.

### SetSkillSetNil

`func (o *ResumeAddResumeRequest) SetSkillSetNil(b bool)`

 SetSkillSetNil sets the value for SkillSet to be an explicit nil

### UnsetSkillSet
`func (o *ResumeAddResumeRequest) UnsetSkillSet()`

UnsetSkillSet ensures that no value is present for SkillSet, not even an explicit nil
### GetSkills

`func (o *ResumeAddResumeRequest) GetSkills() string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ResumeAddResumeRequest) GetSkillsOk() (*string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ResumeAddResumeRequest) SetSkills(v string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ResumeAddResumeRequest) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkillsNil

`func (o *ResumeAddResumeRequest) SetSkillsNil(b bool)`

 SetSkillsNil sets the value for Skills to be an explicit nil

### UnsetSkills
`func (o *ResumeAddResumeRequest) UnsetSkills()`

UnsetSkills ensures that no value is present for Skills, not even an explicit nil
### GetTitle

`func (o *ResumeAddResumeRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeAddResumeRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeAddResumeRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResumeAddResumeRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ResumeAddResumeRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeAddResumeRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalExperience

`func (o *ResumeAddResumeRequest) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeAddResumeRequest) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeAddResumeRequest) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeAddResumeRequest) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeAddResumeRequest) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeAddResumeRequest) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetTravelTime

`func (o *ResumeAddResumeRequest) GetTravelTime() IncludesId`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *ResumeAddResumeRequest) GetTravelTimeOk() (*IncludesId, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *ResumeAddResumeRequest) SetTravelTime(v IncludesId)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *ResumeAddResumeRequest) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.

### GetWorkTicket

`func (o *ResumeAddResumeRequest) GetWorkTicket() []IncludesId`

GetWorkTicket returns the WorkTicket field if non-nil, zero value otherwise.

### GetWorkTicketOk

`func (o *ResumeAddResumeRequest) GetWorkTicketOk() (*[]IncludesId, bool)`

GetWorkTicketOk returns a tuple with the WorkTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTicket

`func (o *ResumeAddResumeRequest) SetWorkTicket(v []IncludesId)`

SetWorkTicket sets WorkTicket field to given value.

### HasWorkTicket

`func (o *ResumeAddResumeRequest) HasWorkTicket() bool`

HasWorkTicket returns a boolean if a field has been set.

### SetWorkTicketNil

`func (o *ResumeAddResumeRequest) SetWorkTicketNil(b bool)`

 SetWorkTicketNil sets the value for WorkTicket to be an explicit nil

### UnsetWorkTicket
`func (o *ResumeAddResumeRequest) UnsetWorkTicket()`

UnsetWorkTicket ensures that no value is present for WorkTicket, not even an explicit nil
### GetArea

`func (o *ResumeAddResumeRequest) GetArea() ResumeAddResumeRequestAllOfArea`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeAddResumeRequest) GetAreaOk() (*ResumeAddResumeRequestAllOfArea, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeAddResumeRequest) SetArea(v ResumeAddResumeRequestAllOfArea)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeAddResumeRequest) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetCitizenship

`func (o *ResumeAddResumeRequest) GetCitizenship() []IncludesId`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *ResumeAddResumeRequest) GetCitizenshipOk() (*[]IncludesId, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *ResumeAddResumeRequest) SetCitizenship(v []IncludesId)`

SetCitizenship sets Citizenship field to given value.

### HasCitizenship

`func (o *ResumeAddResumeRequest) HasCitizenship() bool`

HasCitizenship returns a boolean if a field has been set.

### GetContact

`func (o *ResumeAddResumeRequest) GetContact() []ResumeObjectsContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ResumeAddResumeRequest) GetContactOk() (*[]ResumeObjectsContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ResumeAddResumeRequest) SetContact(v []ResumeObjectsContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ResumeAddResumeRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetEducation

`func (o *ResumeAddResumeRequest) GetEducation() ResumeAddResumeRequestAllOfEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeAddResumeRequest) GetEducationOk() (*ResumeAddResumeRequestAllOfEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeAddResumeRequest) SetEducation(v ResumeAddResumeRequestAllOfEducation)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *ResumeAddResumeRequest) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### GetExperience

`func (o *ResumeAddResumeRequest) GetExperience() []ResumeObjectsExperienceCreateEditResume`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeAddResumeRequest) GetExperienceOk() (*[]ResumeObjectsExperienceCreateEditResume, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeAddResumeRequest) SetExperience(v []ResumeObjectsExperienceCreateEditResume)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *ResumeAddResumeRequest) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### GetGender

`func (o *ResumeAddResumeRequest) GetGender() ResumeAddResumeRequestAllOfGender`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeAddResumeRequest) GetGenderOk() (*ResumeAddResumeRequestAllOfGender, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeAddResumeRequest) SetGender(v ResumeAddResumeRequestAllOfGender)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeAddResumeRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetLanguage

`func (o *ResumeAddResumeRequest) GetLanguage() []ResumeObjectsLanguage`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ResumeAddResumeRequest) GetLanguageOk() (*[]ResumeObjectsLanguage, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ResumeAddResumeRequest) SetLanguage(v []ResumeObjectsLanguage)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ResumeAddResumeRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


