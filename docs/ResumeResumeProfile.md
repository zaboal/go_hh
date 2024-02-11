# ResumeResumeProfile

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
**Education** | [**ResumeObjectsEducation**](ResumeObjectsEducation.md) | Образование соискателя. 

Особенности сохранения образования:

* Если передать и высшее и среднее образование и уровень образования &quot;средний&quot;, то сохранится только среднее образование.
* Если передать и высшее и среднее образование и уровень образования &quot;высшее&quot;, то сохранится только высшее образование
 | 
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

## Methods

### NewResumeResumeProfile

`func NewResumeResumeProfile(alternateUrl string, id string, title NullableString, certificate []ResumeObjectsCertificate, createdAt string, download ResumeObjectsDownload, education ResumeObjectsEducation, experience []ResumeObjectsExperience, hiddenFields []IncludesIdName, updatedAt string, ) *ResumeResumeProfile`

NewResumeResumeProfile instantiates a new ResumeResumeProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeProfileWithDefaults

`func NewResumeResumeProfileWithDefaults() *ResumeResumeProfile`

NewResumeResumeProfileWithDefaults instantiates a new ResumeResumeProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlternateUrl

`func (o *ResumeResumeProfile) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumeResumeProfile) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumeResumeProfile) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetId

`func (o *ResumeResumeProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeResumeProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeResumeProfile) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ResumeResumeProfile) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeResumeProfile) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeResumeProfile) SetTitle(v string)`

SetTitle sets Title field to given value.


### SetTitleNil

`func (o *ResumeResumeProfile) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeResumeProfile) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetAge

`func (o *ResumeResumeProfile) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumeResumeProfile) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumeResumeProfile) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumeResumeProfile) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumeResumeProfile) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumeResumeProfile) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetArea

`func (o *ResumeResumeProfile) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeResumeProfile) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeResumeProfile) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeResumeProfile) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeResumeProfile) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeResumeProfile) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumeResumeProfile) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumeResumeProfile) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumeResumeProfile) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumeResumeProfile) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumeResumeProfile) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumeResumeProfile) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumeResumeProfile) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeResumeProfile) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeResumeProfile) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumeResumeProfile) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumeResumeProfile) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumeResumeProfile) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumeResumeProfile) GetDownload() ResumeObjectsDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeResumeProfile) GetDownloadOk() (*ResumeObjectsDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeResumeProfile) SetDownload(v ResumeObjectsDownload)`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumeResumeProfile) GetEducation() ResumeObjectsEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeResumeProfile) GetEducationOk() (*ResumeObjectsEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeResumeProfile) SetEducation(v ResumeObjectsEducation)`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumeResumeProfile) GetExperience() []ResumeObjectsExperience`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeResumeProfile) GetExperienceOk() (*[]ResumeObjectsExperience, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeResumeProfile) SetExperience(v []ResumeObjectsExperience)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumeResumeProfile) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeResumeProfile) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeResumeProfile) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeResumeProfile) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeResumeProfile) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeResumeProfile) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumeResumeProfile) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeResumeProfile) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeResumeProfile) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeResumeProfile) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ResumeResumeProfile) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ResumeResumeProfile) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *ResumeResumeProfile) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeResumeProfile) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeResumeProfile) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetLastName

`func (o *ResumeResumeProfile) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeResumeProfile) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeResumeProfile) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeResumeProfile) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeResumeProfile) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeResumeProfile) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumeResumeProfile) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumeResumeProfile) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumeResumeProfile) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *ResumeResumeProfile) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetMiddleName

`func (o *ResumeResumeProfile) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeResumeProfile) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeResumeProfile) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeResumeProfile) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeResumeProfile) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeResumeProfile) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPlatform

`func (o *ResumeResumeProfile) GetPlatform() IncludesId`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumeResumeProfile) GetPlatformOk() (*IncludesId, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumeResumeProfile) SetPlatform(v IncludesId)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumeResumeProfile) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeResumeProfile) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeResumeProfile) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeResumeProfile) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeResumeProfile) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumeResumeProfile) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumeResumeProfile) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTotalExperience

`func (o *ResumeResumeProfile) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeResumeProfile) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeResumeProfile) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeResumeProfile) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeResumeProfile) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeResumeProfile) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumeResumeProfile) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumeResumeProfile) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumeResumeProfile) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


