# ResumeEditResumeRequest

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
**Area** | Pointer to [**NullableId**](Id.md) |  | [optional] 
**Citizenship** | Pointer to [**[]IncludesId**](IncludesId.md) | Список гражданств соискателя. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas) | [optional] 
**Contact** | Pointer to [**[]IncludesContact**](IncludesContact.md) | Список контактов соискателя.  При заполнении контактов в резюме необходимо учитывать следующие условия:  * В резюме обязательно должен быть указан e-mail. Он может быть только один. * В резюме должен быть указан хотя бы один телефон, причём можно указывать только один телефон каждого типа. * Комментарий можно указывать только для телефонов, для e-mail комментарий не сохранится * Обязательно указать либо телефон полностью в поле &#x60;formatted&#x60;, либо все три части телефона по отдельности в трёх полях: &#x60;country&#x60;, &#x60;city&#x60; и &#x60;number&#x60;. Если указано и то, и то, используются данные из трёх полей. В поле &#x60;formatted&#x60; допустимо использовать пробелы, скобки и дефисы. В остальных полях допустимы только цифры  | [optional] 
**Education** | Pointer to [**NullableResumeObjectsEducation**](ResumeObjectsEducation.md) |  | [optional] 
**Experience** | Pointer to [**[]ResumeObjectsExperience**](ResumeObjectsExperience.md) | Опыт работы | [optional] 
**Gender** | Pointer to [**NullableId**](Id.md) |  | [optional] 
**Language** | Pointer to [**[]IncludesLanguageLevel**](IncludesLanguageLevel.md) | Список языков, которыми владеет соискатель. Элементы справочника [languages](#tag/Obshie-spravochniki/operation/get-languages) | [optional] 

## Methods

### NewResumeEditResumeRequest

`func NewResumeEditResumeRequest() *ResumeEditResumeRequest`

NewResumeEditResumeRequest instantiates a new ResumeEditResumeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeEditResumeRequestWithDefaults

`func NewResumeEditResumeRequestWithDefaults() *ResumeEditResumeRequest`

NewResumeEditResumeRequestWithDefaults instantiates a new ResumeEditResumeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ResumeEditResumeRequest) GetAccess() ResumeObjectsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ResumeEditResumeRequest) GetAccessOk() (*ResumeObjectsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ResumeEditResumeRequest) SetAccess(v ResumeObjectsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ResumeEditResumeRequest) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetBirthDate

`func (o *ResumeEditResumeRequest) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *ResumeEditResumeRequest) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *ResumeEditResumeRequest) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *ResumeEditResumeRequest) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *ResumeEditResumeRequest) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *ResumeEditResumeRequest) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetBusinessTripReadiness

`func (o *ResumeEditResumeRequest) GetBusinessTripReadiness() IncludesId`

GetBusinessTripReadiness returns the BusinessTripReadiness field if non-nil, zero value otherwise.

### GetBusinessTripReadinessOk

`func (o *ResumeEditResumeRequest) GetBusinessTripReadinessOk() (*IncludesId, bool)`

GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTripReadiness

`func (o *ResumeEditResumeRequest) SetBusinessTripReadiness(v IncludesId)`

SetBusinessTripReadiness sets BusinessTripReadiness field to given value.

### HasBusinessTripReadiness

`func (o *ResumeEditResumeRequest) HasBusinessTripReadiness() bool`

HasBusinessTripReadiness returns a boolean if a field has been set.

### GetCertificate

`func (o *ResumeEditResumeRequest) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeEditResumeRequest) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeEditResumeRequest) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ResumeEditResumeRequest) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ResumeEditResumeRequest) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ResumeEditResumeRequest) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetDriverLicenseTypes

`func (o *ResumeEditResumeRequest) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *ResumeEditResumeRequest) GetDriverLicenseTypesOk() (*[]ResumeObjectsDriverLicenseTypes, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *ResumeEditResumeRequest) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.

### HasDriverLicenseTypes

`func (o *ResumeEditResumeRequest) HasDriverLicenseTypes() bool`

HasDriverLicenseTypes returns a boolean if a field has been set.

### SetDriverLicenseTypesNil

`func (o *ResumeEditResumeRequest) SetDriverLicenseTypesNil(b bool)`

 SetDriverLicenseTypesNil sets the value for DriverLicenseTypes to be an explicit nil

### UnsetDriverLicenseTypes
`func (o *ResumeEditResumeRequest) UnsetDriverLicenseTypes()`

UnsetDriverLicenseTypes ensures that no value is present for DriverLicenseTypes, not even an explicit nil
### GetEmployments

`func (o *ResumeEditResumeRequest) GetEmployments() []IncludesIdName`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *ResumeEditResumeRequest) GetEmploymentsOk() (*[]IncludesIdName, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *ResumeEditResumeRequest) SetEmployments(v []IncludesIdName)`

SetEmployments sets Employments field to given value.

### HasEmployments

`func (o *ResumeEditResumeRequest) HasEmployments() bool`

HasEmployments returns a boolean if a field has been set.

### SetEmploymentsNil

`func (o *ResumeEditResumeRequest) SetEmploymentsNil(b bool)`

 SetEmploymentsNil sets the value for Employments to be an explicit nil

### UnsetEmployments
`func (o *ResumeEditResumeRequest) UnsetEmployments()`

UnsetEmployments ensures that no value is present for Employments, not even an explicit nil
### GetFirstName

`func (o *ResumeEditResumeRequest) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeEditResumeRequest) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeEditResumeRequest) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeEditResumeRequest) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeEditResumeRequest) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeEditResumeRequest) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetHasVehicle

`func (o *ResumeEditResumeRequest) GetHasVehicle() bool`

GetHasVehicle returns the HasVehicle field if non-nil, zero value otherwise.

### GetHasVehicleOk

`func (o *ResumeEditResumeRequest) GetHasVehicleOk() (*bool, bool)`

GetHasVehicleOk returns a tuple with the HasVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVehicle

`func (o *ResumeEditResumeRequest) SetHasVehicle(v bool)`

SetHasVehicle sets HasVehicle field to given value.

### HasHasVehicle

`func (o *ResumeEditResumeRequest) HasHasVehicle() bool`

HasHasVehicle returns a boolean if a field has been set.

### SetHasVehicleNil

`func (o *ResumeEditResumeRequest) SetHasVehicleNil(b bool)`

 SetHasVehicleNil sets the value for HasVehicle to be an explicit nil

### UnsetHasVehicle
`func (o *ResumeEditResumeRequest) UnsetHasVehicle()`

UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
### GetHiddenFields

`func (o *ResumeEditResumeRequest) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeEditResumeRequest) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeEditResumeRequest) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.

### HasHiddenFields

`func (o *ResumeEditResumeRequest) HasHiddenFields() bool`

HasHiddenFields returns a boolean if a field has been set.

### SetHiddenFieldsNil

`func (o *ResumeEditResumeRequest) SetHiddenFieldsNil(b bool)`

 SetHiddenFieldsNil sets the value for HiddenFields to be an explicit nil

### UnsetHiddenFields
`func (o *ResumeEditResumeRequest) UnsetHiddenFields()`

UnsetHiddenFields ensures that no value is present for HiddenFields, not even an explicit nil
### GetLastName

`func (o *ResumeEditResumeRequest) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeEditResumeRequest) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeEditResumeRequest) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeEditResumeRequest) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeEditResumeRequest) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeEditResumeRequest) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMetro

`func (o *ResumeEditResumeRequest) GetMetro() IncludesId`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *ResumeEditResumeRequest) GetMetroOk() (*IncludesId, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *ResumeEditResumeRequest) SetMetro(v IncludesId)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *ResumeEditResumeRequest) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### GetMiddleName

`func (o *ResumeEditResumeRequest) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeEditResumeRequest) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeEditResumeRequest) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeEditResumeRequest) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeEditResumeRequest) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeEditResumeRequest) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPhoto

`func (o *ResumeEditResumeRequest) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeEditResumeRequest) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeEditResumeRequest) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeEditResumeRequest) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeEditResumeRequest) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeEditResumeRequest) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPortfolio

`func (o *ResumeEditResumeRequest) GetPortfolio() []ResumeObjectsPortfolio`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *ResumeEditResumeRequest) GetPortfolioOk() (*[]ResumeObjectsPortfolio, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *ResumeEditResumeRequest) SetPortfolio(v []ResumeObjectsPortfolio)`

SetPortfolio sets Portfolio field to given value.

### HasPortfolio

`func (o *ResumeEditResumeRequest) HasPortfolio() bool`

HasPortfolio returns a boolean if a field has been set.

### SetPortfolioNil

`func (o *ResumeEditResumeRequest) SetPortfolioNil(b bool)`

 SetPortfolioNil sets the value for Portfolio to be an explicit nil

### UnsetPortfolio
`func (o *ResumeEditResumeRequest) UnsetPortfolio()`

UnsetPortfolio ensures that no value is present for Portfolio, not even an explicit nil
### GetProfessionalRoles

`func (o *ResumeEditResumeRequest) GetProfessionalRoles() []IncludesId`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *ResumeEditResumeRequest) GetProfessionalRolesOk() (*[]IncludesId, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *ResumeEditResumeRequest) SetProfessionalRoles(v []IncludesId)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *ResumeEditResumeRequest) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetRecommendation

`func (o *ResumeEditResumeRequest) GetRecommendation() []ResumeObjectsRecommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *ResumeEditResumeRequest) GetRecommendationOk() (*[]ResumeObjectsRecommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *ResumeEditResumeRequest) SetRecommendation(v []ResumeObjectsRecommendation)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *ResumeEditResumeRequest) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### SetRecommendationNil

`func (o *ResumeEditResumeRequest) SetRecommendationNil(b bool)`

 SetRecommendationNil sets the value for Recommendation to be an explicit nil

### UnsetRecommendation
`func (o *ResumeEditResumeRequest) UnsetRecommendation()`

UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
### GetRelocation

`func (o *ResumeEditResumeRequest) GetRelocation() ResumeObjectsRelocationPublic`

GetRelocation returns the Relocation field if non-nil, zero value otherwise.

### GetRelocationOk

`func (o *ResumeEditResumeRequest) GetRelocationOk() (*ResumeObjectsRelocationPublic, bool)`

GetRelocationOk returns a tuple with the Relocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocation

`func (o *ResumeEditResumeRequest) SetRelocation(v ResumeObjectsRelocationPublic)`

SetRelocation sets Relocation field to given value.

### HasRelocation

`func (o *ResumeEditResumeRequest) HasRelocation() bool`

HasRelocation returns a boolean if a field has been set.

### GetResumeLocale

`func (o *ResumeEditResumeRequest) GetResumeLocale() IncludesIdName`

GetResumeLocale returns the ResumeLocale field if non-nil, zero value otherwise.

### GetResumeLocaleOk

`func (o *ResumeEditResumeRequest) GetResumeLocaleOk() (*IncludesIdName, bool)`

GetResumeLocaleOk returns a tuple with the ResumeLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeLocale

`func (o *ResumeEditResumeRequest) SetResumeLocale(v IncludesIdName)`

SetResumeLocale sets ResumeLocale field to given value.

### HasResumeLocale

`func (o *ResumeEditResumeRequest) HasResumeLocale() bool`

HasResumeLocale returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeEditResumeRequest) GetSalary() ResumeObjectsSalaryAddEdit`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeEditResumeRequest) GetSalaryOk() (*ResumeObjectsSalaryAddEdit, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeEditResumeRequest) SetSalary(v ResumeObjectsSalaryAddEdit)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeEditResumeRequest) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### GetSchedules

`func (o *ResumeEditResumeRequest) GetSchedules() []IncludesIdName`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ResumeEditResumeRequest) GetSchedulesOk() (*[]IncludesIdName, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ResumeEditResumeRequest) SetSchedules(v []IncludesIdName)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ResumeEditResumeRequest) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### SetSchedulesNil

`func (o *ResumeEditResumeRequest) SetSchedulesNil(b bool)`

 SetSchedulesNil sets the value for Schedules to be an explicit nil

### UnsetSchedules
`func (o *ResumeEditResumeRequest) UnsetSchedules()`

UnsetSchedules ensures that no value is present for Schedules, not even an explicit nil
### GetSite

`func (o *ResumeEditResumeRequest) GetSite() []ResumeObjectsSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ResumeEditResumeRequest) GetSiteOk() (*[]ResumeObjectsSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ResumeEditResumeRequest) SetSite(v []ResumeObjectsSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *ResumeEditResumeRequest) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *ResumeEditResumeRequest) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *ResumeEditResumeRequest) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetSkillSet

`func (o *ResumeEditResumeRequest) GetSkillSet() []string`

GetSkillSet returns the SkillSet field if non-nil, zero value otherwise.

### GetSkillSetOk

`func (o *ResumeEditResumeRequest) GetSkillSetOk() (*[]string, bool)`

GetSkillSetOk returns a tuple with the SkillSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillSet

`func (o *ResumeEditResumeRequest) SetSkillSet(v []string)`

SetSkillSet sets SkillSet field to given value.

### HasSkillSet

`func (o *ResumeEditResumeRequest) HasSkillSet() bool`

HasSkillSet returns a boolean if a field has been set.

### SetSkillSetNil

`func (o *ResumeEditResumeRequest) SetSkillSetNil(b bool)`

 SetSkillSetNil sets the value for SkillSet to be an explicit nil

### UnsetSkillSet
`func (o *ResumeEditResumeRequest) UnsetSkillSet()`

UnsetSkillSet ensures that no value is present for SkillSet, not even an explicit nil
### GetSkills

`func (o *ResumeEditResumeRequest) GetSkills() string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ResumeEditResumeRequest) GetSkillsOk() (*string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ResumeEditResumeRequest) SetSkills(v string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ResumeEditResumeRequest) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkillsNil

`func (o *ResumeEditResumeRequest) SetSkillsNil(b bool)`

 SetSkillsNil sets the value for Skills to be an explicit nil

### UnsetSkills
`func (o *ResumeEditResumeRequest) UnsetSkills()`

UnsetSkills ensures that no value is present for Skills, not even an explicit nil
### GetTitle

`func (o *ResumeEditResumeRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeEditResumeRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeEditResumeRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResumeEditResumeRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ResumeEditResumeRequest) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeEditResumeRequest) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalExperience

`func (o *ResumeEditResumeRequest) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeEditResumeRequest) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeEditResumeRequest) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeEditResumeRequest) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeEditResumeRequest) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeEditResumeRequest) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetTravelTime

`func (o *ResumeEditResumeRequest) GetTravelTime() IncludesId`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *ResumeEditResumeRequest) GetTravelTimeOk() (*IncludesId, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *ResumeEditResumeRequest) SetTravelTime(v IncludesId)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *ResumeEditResumeRequest) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.

### GetWorkTicket

`func (o *ResumeEditResumeRequest) GetWorkTicket() []IncludesId`

GetWorkTicket returns the WorkTicket field if non-nil, zero value otherwise.

### GetWorkTicketOk

`func (o *ResumeEditResumeRequest) GetWorkTicketOk() (*[]IncludesId, bool)`

GetWorkTicketOk returns a tuple with the WorkTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTicket

`func (o *ResumeEditResumeRequest) SetWorkTicket(v []IncludesId)`

SetWorkTicket sets WorkTicket field to given value.

### HasWorkTicket

`func (o *ResumeEditResumeRequest) HasWorkTicket() bool`

HasWorkTicket returns a boolean if a field has been set.

### SetWorkTicketNil

`func (o *ResumeEditResumeRequest) SetWorkTicketNil(b bool)`

 SetWorkTicketNil sets the value for WorkTicket to be an explicit nil

### UnsetWorkTicket
`func (o *ResumeEditResumeRequest) UnsetWorkTicket()`

UnsetWorkTicket ensures that no value is present for WorkTicket, not even an explicit nil
### GetArea

`func (o *ResumeEditResumeRequest) GetArea() Id`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeEditResumeRequest) GetAreaOk() (*Id, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeEditResumeRequest) SetArea(v Id)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeEditResumeRequest) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeEditResumeRequest) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeEditResumeRequest) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCitizenship

`func (o *ResumeEditResumeRequest) GetCitizenship() []IncludesId`

GetCitizenship returns the Citizenship field if non-nil, zero value otherwise.

### GetCitizenshipOk

`func (o *ResumeEditResumeRequest) GetCitizenshipOk() (*[]IncludesId, bool)`

GetCitizenshipOk returns a tuple with the Citizenship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitizenship

`func (o *ResumeEditResumeRequest) SetCitizenship(v []IncludesId)`

SetCitizenship sets Citizenship field to given value.

### HasCitizenship

`func (o *ResumeEditResumeRequest) HasCitizenship() bool`

HasCitizenship returns a boolean if a field has been set.

### SetCitizenshipNil

`func (o *ResumeEditResumeRequest) SetCitizenshipNil(b bool)`

 SetCitizenshipNil sets the value for Citizenship to be an explicit nil

### UnsetCitizenship
`func (o *ResumeEditResumeRequest) UnsetCitizenship()`

UnsetCitizenship ensures that no value is present for Citizenship, not even an explicit nil
### GetContact

`func (o *ResumeEditResumeRequest) GetContact() []IncludesContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ResumeEditResumeRequest) GetContactOk() (*[]IncludesContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ResumeEditResumeRequest) SetContact(v []IncludesContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *ResumeEditResumeRequest) HasContact() bool`

HasContact returns a boolean if a field has been set.

### SetContactNil

`func (o *ResumeEditResumeRequest) SetContactNil(b bool)`

 SetContactNil sets the value for Contact to be an explicit nil

### UnsetContact
`func (o *ResumeEditResumeRequest) UnsetContact()`

UnsetContact ensures that no value is present for Contact, not even an explicit nil
### GetEducation

`func (o *ResumeEditResumeRequest) GetEducation() ResumeObjectsEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeEditResumeRequest) GetEducationOk() (*ResumeObjectsEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeEditResumeRequest) SetEducation(v ResumeObjectsEducation)`

SetEducation sets Education field to given value.

### HasEducation

`func (o *ResumeEditResumeRequest) HasEducation() bool`

HasEducation returns a boolean if a field has been set.

### SetEducationNil

`func (o *ResumeEditResumeRequest) SetEducationNil(b bool)`

 SetEducationNil sets the value for Education to be an explicit nil

### UnsetEducation
`func (o *ResumeEditResumeRequest) UnsetEducation()`

UnsetEducation ensures that no value is present for Education, not even an explicit nil
### GetExperience

`func (o *ResumeEditResumeRequest) GetExperience() []ResumeObjectsExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeEditResumeRequest) GetExperienceOk() (*[]ResumeObjectsExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeEditResumeRequest) SetExperience(v []ResumeObjectsExperience)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *ResumeEditResumeRequest) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *ResumeEditResumeRequest) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *ResumeEditResumeRequest) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetGender

`func (o *ResumeEditResumeRequest) GetGender() Id`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeEditResumeRequest) GetGenderOk() (*Id, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeEditResumeRequest) SetGender(v Id)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeEditResumeRequest) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ResumeEditResumeRequest) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ResumeEditResumeRequest) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetLanguage

`func (o *ResumeEditResumeRequest) GetLanguage() []IncludesLanguageLevel`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *ResumeEditResumeRequest) GetLanguageOk() (*[]IncludesLanguageLevel, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *ResumeEditResumeRequest) SetLanguage(v []IncludesLanguageLevel)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *ResumeEditResumeRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### SetLanguageNil

`func (o *ResumeEditResumeRequest) SetLanguageNil(b bool)`

 SetLanguageNil sets the value for Language to be an explicit nil

### UnsetLanguage
`func (o *ResumeEditResumeRequest) UnsetLanguage()`

UnsetLanguage ensures that no value is present for Language, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


