# ResumeNullableFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**NullableResumeObjectsAccess**](ResumeObjectsAccess.md) |  | [optional] 
**BirthDate** | Pointer to **NullableString** | День рождения (в формате &#x60;ГГГГ-ММ-ДД&#x60;) | [optional] 
**BusinessTripReadiness** | Pointer to [**NullableIncludesId**](IncludesId.md) |  | [optional] 
**Certificate** | Pointer to [**[]ResumeObjectsCertificate**](ResumeObjectsCertificate.md) | Список сертификатов соискателя | [optional] 
**DriverLicenseTypes** | Pointer to [**[]ResumeObjectsDriverLicenseTypes**](ResumeObjectsDriverLicenseTypes.md) | Список категорий водительских прав соискателя | [optional] 
**Employments** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Список подходящих соискателю типов занятостей. Элементы справочника [employment](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**FirstName** | Pointer to **NullableString** | Имя | [optional] 
**HasVehicle** | Pointer to **NullableBool** | Наличие личного автомобиля у соискателя | [optional] 
**HiddenFields** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Документация [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле &#x60;resume_hidden_fields&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**LastName** | Pointer to **NullableString** | Фамилия | [optional] 
**Metro** | Pointer to [**NullableIncludesId**](IncludesId.md) |  | [optional] 
**MiddleName** | Pointer to **NullableString** | Отчество | [optional] 
**Photo** | Pointer to [**NullableResumeObjectsPhoto**](ResumeObjectsPhoto.md) |  | [optional] 
**Portfolio** | Pointer to [**[]ResumeObjectsPortfolio**](ResumeObjectsPortfolio.md) | Список изображений в портфолио пользователя | [optional] 
**ProfessionalRoles** | Pointer to [**[]IncludesId**](IncludesId.md) | Массив объектов профролей. Элемент справочника [professional_roles](#tag/Obshie-spravochniki/operation/get-professional-roles-dictionary) | [optional] 
**Recommendation** | Pointer to [**[]ResumeObjectsRecommendation**](ResumeObjectsRecommendation.md) | Список рекомендаций | [optional] 
**Relocation** | Pointer to [**NullableResumeObjectsRelocationPublic**](ResumeObjectsRelocationPublic.md) |  | [optional] 
**ResumeLocale** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**Salary** | Pointer to [**NullableResumeObjectsSalaryAddEdit**](ResumeObjectsSalaryAddEdit.md) |  | [optional] 
**Schedules** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Список подходящих соискателю графиков работы. Элементы справочника [schedule](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**Site** | Pointer to [**[]ResumeObjectsSite**](ResumeObjectsSite.md) | Профили в соц. сетях и других сервисах | [optional] 
**SkillSet** | Pointer to **[]string** | Ключевые навыки (список уникальных строк) | [optional] 
**Skills** | Pointer to **NullableString** | Дополнительная информация, описание навыков в свободной форме | [optional] 
**Title** | Pointer to **NullableString** | Желаемая должность | [optional] 
**TotalExperience** | Pointer to [**NullableResumeObjectsTotalExperience**](ResumeObjectsTotalExperience.md) |  | [optional] 
**TravelTime** | Pointer to [**NullableIncludesId**](IncludesId.md) |  | [optional] 
**WorkTicket** | Pointer to [**[]IncludesId**](IncludesId.md) | Список регионов, в который соискатель имеет разрешение на работу. Элементы [справочника регионов](#tag/Obshie-spravochniki/operation/get-areas)  | [optional] 

## Methods

### NewResumeNullableFields

`func NewResumeNullableFields() *ResumeNullableFields`

NewResumeNullableFields instantiates a new ResumeNullableFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeNullableFieldsWithDefaults

`func NewResumeNullableFieldsWithDefaults() *ResumeNullableFields`

NewResumeNullableFieldsWithDefaults instantiates a new ResumeNullableFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *ResumeNullableFields) GetAccess() ResumeObjectsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ResumeNullableFields) GetAccessOk() (*ResumeObjectsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ResumeNullableFields) SetAccess(v ResumeObjectsAccess)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ResumeNullableFields) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccessNil

`func (o *ResumeNullableFields) SetAccessNil(b bool)`

 SetAccessNil sets the value for Access to be an explicit nil

### UnsetAccess
`func (o *ResumeNullableFields) UnsetAccess()`

UnsetAccess ensures that no value is present for Access, not even an explicit nil
### GetBirthDate

`func (o *ResumeNullableFields) GetBirthDate() string`

GetBirthDate returns the BirthDate field if non-nil, zero value otherwise.

### GetBirthDateOk

`func (o *ResumeNullableFields) GetBirthDateOk() (*string, bool)`

GetBirthDateOk returns a tuple with the BirthDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBirthDate

`func (o *ResumeNullableFields) SetBirthDate(v string)`

SetBirthDate sets BirthDate field to given value.

### HasBirthDate

`func (o *ResumeNullableFields) HasBirthDate() bool`

HasBirthDate returns a boolean if a field has been set.

### SetBirthDateNil

`func (o *ResumeNullableFields) SetBirthDateNil(b bool)`

 SetBirthDateNil sets the value for BirthDate to be an explicit nil

### UnsetBirthDate
`func (o *ResumeNullableFields) UnsetBirthDate()`

UnsetBirthDate ensures that no value is present for BirthDate, not even an explicit nil
### GetBusinessTripReadiness

`func (o *ResumeNullableFields) GetBusinessTripReadiness() IncludesId`

GetBusinessTripReadiness returns the BusinessTripReadiness field if non-nil, zero value otherwise.

### GetBusinessTripReadinessOk

`func (o *ResumeNullableFields) GetBusinessTripReadinessOk() (*IncludesId, bool)`

GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTripReadiness

`func (o *ResumeNullableFields) SetBusinessTripReadiness(v IncludesId)`

SetBusinessTripReadiness sets BusinessTripReadiness field to given value.

### HasBusinessTripReadiness

`func (o *ResumeNullableFields) HasBusinessTripReadiness() bool`

HasBusinessTripReadiness returns a boolean if a field has been set.

### SetBusinessTripReadinessNil

`func (o *ResumeNullableFields) SetBusinessTripReadinessNil(b bool)`

 SetBusinessTripReadinessNil sets the value for BusinessTripReadiness to be an explicit nil

### UnsetBusinessTripReadiness
`func (o *ResumeNullableFields) UnsetBusinessTripReadiness()`

UnsetBusinessTripReadiness ensures that no value is present for BusinessTripReadiness, not even an explicit nil
### GetCertificate

`func (o *ResumeNullableFields) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeNullableFields) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeNullableFields) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ResumeNullableFields) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### SetCertificateNil

`func (o *ResumeNullableFields) SetCertificateNil(b bool)`

 SetCertificateNil sets the value for Certificate to be an explicit nil

### UnsetCertificate
`func (o *ResumeNullableFields) UnsetCertificate()`

UnsetCertificate ensures that no value is present for Certificate, not even an explicit nil
### GetDriverLicenseTypes

`func (o *ResumeNullableFields) GetDriverLicenseTypes() []ResumeObjectsDriverLicenseTypes`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *ResumeNullableFields) GetDriverLicenseTypesOk() (*[]ResumeObjectsDriverLicenseTypes, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *ResumeNullableFields) SetDriverLicenseTypes(v []ResumeObjectsDriverLicenseTypes)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.

### HasDriverLicenseTypes

`func (o *ResumeNullableFields) HasDriverLicenseTypes() bool`

HasDriverLicenseTypes returns a boolean if a field has been set.

### SetDriverLicenseTypesNil

`func (o *ResumeNullableFields) SetDriverLicenseTypesNil(b bool)`

 SetDriverLicenseTypesNil sets the value for DriverLicenseTypes to be an explicit nil

### UnsetDriverLicenseTypes
`func (o *ResumeNullableFields) UnsetDriverLicenseTypes()`

UnsetDriverLicenseTypes ensures that no value is present for DriverLicenseTypes, not even an explicit nil
### GetEmployments

`func (o *ResumeNullableFields) GetEmployments() []IncludesIdName`

GetEmployments returns the Employments field if non-nil, zero value otherwise.

### GetEmploymentsOk

`func (o *ResumeNullableFields) GetEmploymentsOk() (*[]IncludesIdName, bool)`

GetEmploymentsOk returns a tuple with the Employments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployments

`func (o *ResumeNullableFields) SetEmployments(v []IncludesIdName)`

SetEmployments sets Employments field to given value.

### HasEmployments

`func (o *ResumeNullableFields) HasEmployments() bool`

HasEmployments returns a boolean if a field has been set.

### SetEmploymentsNil

`func (o *ResumeNullableFields) SetEmploymentsNil(b bool)`

 SetEmploymentsNil sets the value for Employments to be an explicit nil

### UnsetEmployments
`func (o *ResumeNullableFields) UnsetEmployments()`

UnsetEmployments ensures that no value is present for Employments, not even an explicit nil
### GetFirstName

`func (o *ResumeNullableFields) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeNullableFields) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeNullableFields) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeNullableFields) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeNullableFields) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeNullableFields) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetHasVehicle

`func (o *ResumeNullableFields) GetHasVehicle() bool`

GetHasVehicle returns the HasVehicle field if non-nil, zero value otherwise.

### GetHasVehicleOk

`func (o *ResumeNullableFields) GetHasVehicleOk() (*bool, bool)`

GetHasVehicleOk returns a tuple with the HasVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVehicle

`func (o *ResumeNullableFields) SetHasVehicle(v bool)`

SetHasVehicle sets HasVehicle field to given value.

### HasHasVehicle

`func (o *ResumeNullableFields) HasHasVehicle() bool`

HasHasVehicle returns a boolean if a field has been set.

### SetHasVehicleNil

`func (o *ResumeNullableFields) SetHasVehicleNil(b bool)`

 SetHasVehicleNil sets the value for HasVehicle to be an explicit nil

### UnsetHasVehicle
`func (o *ResumeNullableFields) UnsetHasVehicle()`

UnsetHasVehicle ensures that no value is present for HasVehicle, not even an explicit nil
### GetHiddenFields

`func (o *ResumeNullableFields) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeNullableFields) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeNullableFields) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.

### HasHiddenFields

`func (o *ResumeNullableFields) HasHiddenFields() bool`

HasHiddenFields returns a boolean if a field has been set.

### SetHiddenFieldsNil

`func (o *ResumeNullableFields) SetHiddenFieldsNil(b bool)`

 SetHiddenFieldsNil sets the value for HiddenFields to be an explicit nil

### UnsetHiddenFields
`func (o *ResumeNullableFields) UnsetHiddenFields()`

UnsetHiddenFields ensures that no value is present for HiddenFields, not even an explicit nil
### GetLastName

`func (o *ResumeNullableFields) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeNullableFields) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeNullableFields) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeNullableFields) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeNullableFields) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeNullableFields) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMetro

`func (o *ResumeNullableFields) GetMetro() IncludesId`

GetMetro returns the Metro field if non-nil, zero value otherwise.

### GetMetroOk

`func (o *ResumeNullableFields) GetMetroOk() (*IncludesId, bool)`

GetMetroOk returns a tuple with the Metro field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetro

`func (o *ResumeNullableFields) SetMetro(v IncludesId)`

SetMetro sets Metro field to given value.

### HasMetro

`func (o *ResumeNullableFields) HasMetro() bool`

HasMetro returns a boolean if a field has been set.

### SetMetroNil

`func (o *ResumeNullableFields) SetMetroNil(b bool)`

 SetMetroNil sets the value for Metro to be an explicit nil

### UnsetMetro
`func (o *ResumeNullableFields) UnsetMetro()`

UnsetMetro ensures that no value is present for Metro, not even an explicit nil
### GetMiddleName

`func (o *ResumeNullableFields) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeNullableFields) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeNullableFields) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeNullableFields) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeNullableFields) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeNullableFields) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPhoto

`func (o *ResumeNullableFields) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeNullableFields) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeNullableFields) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeNullableFields) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeNullableFields) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeNullableFields) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPortfolio

`func (o *ResumeNullableFields) GetPortfolio() []ResumeObjectsPortfolio`

GetPortfolio returns the Portfolio field if non-nil, zero value otherwise.

### GetPortfolioOk

`func (o *ResumeNullableFields) GetPortfolioOk() (*[]ResumeObjectsPortfolio, bool)`

GetPortfolioOk returns a tuple with the Portfolio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortfolio

`func (o *ResumeNullableFields) SetPortfolio(v []ResumeObjectsPortfolio)`

SetPortfolio sets Portfolio field to given value.

### HasPortfolio

`func (o *ResumeNullableFields) HasPortfolio() bool`

HasPortfolio returns a boolean if a field has been set.

### SetPortfolioNil

`func (o *ResumeNullableFields) SetPortfolioNil(b bool)`

 SetPortfolioNil sets the value for Portfolio to be an explicit nil

### UnsetPortfolio
`func (o *ResumeNullableFields) UnsetPortfolio()`

UnsetPortfolio ensures that no value is present for Portfolio, not even an explicit nil
### GetProfessionalRoles

`func (o *ResumeNullableFields) GetProfessionalRoles() []IncludesId`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *ResumeNullableFields) GetProfessionalRolesOk() (*[]IncludesId, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *ResumeNullableFields) SetProfessionalRoles(v []IncludesId)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *ResumeNullableFields) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetRecommendation

`func (o *ResumeNullableFields) GetRecommendation() []ResumeObjectsRecommendation`

GetRecommendation returns the Recommendation field if non-nil, zero value otherwise.

### GetRecommendationOk

`func (o *ResumeNullableFields) GetRecommendationOk() (*[]ResumeObjectsRecommendation, bool)`

GetRecommendationOk returns a tuple with the Recommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendation

`func (o *ResumeNullableFields) SetRecommendation(v []ResumeObjectsRecommendation)`

SetRecommendation sets Recommendation field to given value.

### HasRecommendation

`func (o *ResumeNullableFields) HasRecommendation() bool`

HasRecommendation returns a boolean if a field has been set.

### SetRecommendationNil

`func (o *ResumeNullableFields) SetRecommendationNil(b bool)`

 SetRecommendationNil sets the value for Recommendation to be an explicit nil

### UnsetRecommendation
`func (o *ResumeNullableFields) UnsetRecommendation()`

UnsetRecommendation ensures that no value is present for Recommendation, not even an explicit nil
### GetRelocation

`func (o *ResumeNullableFields) GetRelocation() ResumeObjectsRelocationPublic`

GetRelocation returns the Relocation field if non-nil, zero value otherwise.

### GetRelocationOk

`func (o *ResumeNullableFields) GetRelocationOk() (*ResumeObjectsRelocationPublic, bool)`

GetRelocationOk returns a tuple with the Relocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocation

`func (o *ResumeNullableFields) SetRelocation(v ResumeObjectsRelocationPublic)`

SetRelocation sets Relocation field to given value.

### HasRelocation

`func (o *ResumeNullableFields) HasRelocation() bool`

HasRelocation returns a boolean if a field has been set.

### SetRelocationNil

`func (o *ResumeNullableFields) SetRelocationNil(b bool)`

 SetRelocationNil sets the value for Relocation to be an explicit nil

### UnsetRelocation
`func (o *ResumeNullableFields) UnsetRelocation()`

UnsetRelocation ensures that no value is present for Relocation, not even an explicit nil
### GetResumeLocale

`func (o *ResumeNullableFields) GetResumeLocale() IncludesIdName`

GetResumeLocale returns the ResumeLocale field if non-nil, zero value otherwise.

### GetResumeLocaleOk

`func (o *ResumeNullableFields) GetResumeLocaleOk() (*IncludesIdName, bool)`

GetResumeLocaleOk returns a tuple with the ResumeLocale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeLocale

`func (o *ResumeNullableFields) SetResumeLocale(v IncludesIdName)`

SetResumeLocale sets ResumeLocale field to given value.

### HasResumeLocale

`func (o *ResumeNullableFields) HasResumeLocale() bool`

HasResumeLocale returns a boolean if a field has been set.

### SetResumeLocaleNil

`func (o *ResumeNullableFields) SetResumeLocaleNil(b bool)`

 SetResumeLocaleNil sets the value for ResumeLocale to be an explicit nil

### UnsetResumeLocale
`func (o *ResumeNullableFields) UnsetResumeLocale()`

UnsetResumeLocale ensures that no value is present for ResumeLocale, not even an explicit nil
### GetSalary

`func (o *ResumeNullableFields) GetSalary() ResumeObjectsSalaryAddEdit`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeNullableFields) GetSalaryOk() (*ResumeObjectsSalaryAddEdit, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeNullableFields) SetSalary(v ResumeObjectsSalaryAddEdit)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeNullableFields) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumeNullableFields) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumeNullableFields) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetSchedules

`func (o *ResumeNullableFields) GetSchedules() []IncludesIdName`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ResumeNullableFields) GetSchedulesOk() (*[]IncludesIdName, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ResumeNullableFields) SetSchedules(v []IncludesIdName)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ResumeNullableFields) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### SetSchedulesNil

`func (o *ResumeNullableFields) SetSchedulesNil(b bool)`

 SetSchedulesNil sets the value for Schedules to be an explicit nil

### UnsetSchedules
`func (o *ResumeNullableFields) UnsetSchedules()`

UnsetSchedules ensures that no value is present for Schedules, not even an explicit nil
### GetSite

`func (o *ResumeNullableFields) GetSite() []ResumeObjectsSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ResumeNullableFields) GetSiteOk() (*[]ResumeObjectsSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ResumeNullableFields) SetSite(v []ResumeObjectsSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *ResumeNullableFields) HasSite() bool`

HasSite returns a boolean if a field has been set.

### SetSiteNil

`func (o *ResumeNullableFields) SetSiteNil(b bool)`

 SetSiteNil sets the value for Site to be an explicit nil

### UnsetSite
`func (o *ResumeNullableFields) UnsetSite()`

UnsetSite ensures that no value is present for Site, not even an explicit nil
### GetSkillSet

`func (o *ResumeNullableFields) GetSkillSet() []string`

GetSkillSet returns the SkillSet field if non-nil, zero value otherwise.

### GetSkillSetOk

`func (o *ResumeNullableFields) GetSkillSetOk() (*[]string, bool)`

GetSkillSetOk returns a tuple with the SkillSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillSet

`func (o *ResumeNullableFields) SetSkillSet(v []string)`

SetSkillSet sets SkillSet field to given value.

### HasSkillSet

`func (o *ResumeNullableFields) HasSkillSet() bool`

HasSkillSet returns a boolean if a field has been set.

### SetSkillSetNil

`func (o *ResumeNullableFields) SetSkillSetNil(b bool)`

 SetSkillSetNil sets the value for SkillSet to be an explicit nil

### UnsetSkillSet
`func (o *ResumeNullableFields) UnsetSkillSet()`

UnsetSkillSet ensures that no value is present for SkillSet, not even an explicit nil
### GetSkills

`func (o *ResumeNullableFields) GetSkills() string`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *ResumeNullableFields) GetSkillsOk() (*string, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *ResumeNullableFields) SetSkills(v string)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *ResumeNullableFields) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkillsNil

`func (o *ResumeNullableFields) SetSkillsNil(b bool)`

 SetSkillsNil sets the value for Skills to be an explicit nil

### UnsetSkills
`func (o *ResumeNullableFields) UnsetSkills()`

UnsetSkills ensures that no value is present for Skills, not even an explicit nil
### GetTitle

`func (o *ResumeNullableFields) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeNullableFields) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeNullableFields) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResumeNullableFields) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ResumeNullableFields) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeNullableFields) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalExperience

`func (o *ResumeNullableFields) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeNullableFields) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeNullableFields) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeNullableFields) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeNullableFields) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeNullableFields) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetTravelTime

`func (o *ResumeNullableFields) GetTravelTime() IncludesId`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *ResumeNullableFields) GetTravelTimeOk() (*IncludesId, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *ResumeNullableFields) SetTravelTime(v IncludesId)`

SetTravelTime sets TravelTime field to given value.

### HasTravelTime

`func (o *ResumeNullableFields) HasTravelTime() bool`

HasTravelTime returns a boolean if a field has been set.

### SetTravelTimeNil

`func (o *ResumeNullableFields) SetTravelTimeNil(b bool)`

 SetTravelTimeNil sets the value for TravelTime to be an explicit nil

### UnsetTravelTime
`func (o *ResumeNullableFields) UnsetTravelTime()`

UnsetTravelTime ensures that no value is present for TravelTime, not even an explicit nil
### GetWorkTicket

`func (o *ResumeNullableFields) GetWorkTicket() []IncludesId`

GetWorkTicket returns the WorkTicket field if non-nil, zero value otherwise.

### GetWorkTicketOk

`func (o *ResumeNullableFields) GetWorkTicketOk() (*[]IncludesId, bool)`

GetWorkTicketOk returns a tuple with the WorkTicket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTicket

`func (o *ResumeNullableFields) SetWorkTicket(v []IncludesId)`

SetWorkTicket sets WorkTicket field to given value.

### HasWorkTicket

`func (o *ResumeNullableFields) HasWorkTicket() bool`

HasWorkTicket returns a boolean if a field has been set.

### SetWorkTicketNil

`func (o *ResumeNullableFields) SetWorkTicketNil(b bool)`

 SetWorkTicketNil sets the value for WorkTicket to be an explicit nil

### UnsetWorkTicket
`func (o *ResumeNullableFields) UnsetWorkTicket()`

UnsetWorkTicket ensures that no value is present for WorkTicket, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


