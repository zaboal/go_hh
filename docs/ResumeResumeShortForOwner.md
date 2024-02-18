# ResumeResumeShortForOwner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**ResumeObjectsActionsForOwner**](ResumeObjectsActionsForOwner.md) | Дополнительные действия | 
**Age** | Pointer to **NullableFloat32** | Возраст | [optional] 
**AlternateUrl** | **string** | URL резюме на сайте | 
**Area** | Pointer to [**NullableIncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**AutoHideTime** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**CanViewFullInfo** | Pointer to **NullableBool** | Доступен ли просмотр контактной информации в резюме текущему работодателю | [optional] 
**Certificate** | [**[]ResumeObjectsCertificate**](ResumeObjectsCertificate.md) | Список сертификатов соискателя | 
**CreatedAt** | **string** | Дата и время создания резюме | 
**Download** | [**ResumeObjectsDownload**](ResumeObjectsDownload.md) | Ссылки для скачивания резюме в разных форматах | 
**Education** | [**ResumeObjectsEducation**](ResumeObjectsEducation.md) | Образование соискателя.   Особенности сохранения образования:  * Если передать и высшее и среднее образование и уровень образования \&quot;средний\&quot;, то сохранится только среднее образование. * Если передать и высшее и среднее образование и уровень образования \&quot;высшее\&quot;, то сохранится только высшее образование  | 
**Experience** | [**[]ResumeObjectsExperienceForOwner**](ResumeObjectsExperienceForOwner.md) | Опыт работы | 
**FirstName** | Pointer to **NullableString** | Имя | [optional] 
**Gender** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**HiddenFields** | [**[]IncludesIdName**](IncludesIdName.md) | Документация [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле &#x60;resume_hidden_fields&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Id** | **string** | Идентификатор резюме | 
**LastName** | Pointer to **NullableString** | Фамилия | [optional] 
**Marked** | **bool** | Выделено ли резюме в поиске | 
**MiddleName** | Pointer to **NullableString** | Отчество | [optional] 
**Photo** | Pointer to [**NullableProfilePhoto**](ProfilePhoto.md) |  | [optional] 
**Platform** | Pointer to [**IncludesId**](IncludesId.md) | Ресурс, на котором было размещено резюме | [optional] 
**Salary** | Pointer to [**NullableResumeObjectsSalaryProperties**](ResumeObjectsSalaryProperties.md) |  | [optional] 
**Title** | Pointer to **NullableString** | Желаемая должность | [optional] 
**TotalExperience** | Pointer to [**NullableResumeObjectsTotalExperience**](ResumeObjectsTotalExperience.md) |  | [optional] 
**UpdatedAt** | **string** | Дата и время обновления резюме | 
**Url** | **string** | URL резюме на сайте | 

## Methods

### NewResumeResumeShortForOwner

`func NewResumeResumeShortForOwner(actions ResumeObjectsActionsForOwner, alternateUrl string, certificate []ResumeObjectsCertificate, createdAt string, download ResumeObjectsDownload, education ResumeObjectsEducation, experience []ResumeObjectsExperienceForOwner, hiddenFields []IncludesIdName, id string, marked bool, updatedAt string, url string, ) *ResumeResumeShortForOwner`

NewResumeResumeShortForOwner instantiates a new ResumeResumeShortForOwner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeShortForOwnerWithDefaults

`func NewResumeResumeShortForOwnerWithDefaults() *ResumeResumeShortForOwner`

NewResumeResumeShortForOwnerWithDefaults instantiates a new ResumeResumeShortForOwner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ResumeResumeShortForOwner) GetActions() ResumeObjectsActionsForOwner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumeResumeShortForOwner) GetActionsOk() (*ResumeObjectsActionsForOwner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumeResumeShortForOwner) SetActions(v ResumeObjectsActionsForOwner)`

SetActions sets Actions field to given value.


### GetAge

`func (o *ResumeResumeShortForOwner) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumeResumeShortForOwner) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumeResumeShortForOwner) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumeResumeShortForOwner) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumeResumeShortForOwner) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumeResumeShortForOwner) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetAlternateUrl

`func (o *ResumeResumeShortForOwner) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumeResumeShortForOwner) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumeResumeShortForOwner) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetArea

`func (o *ResumeResumeShortForOwner) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumeResumeShortForOwner) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumeResumeShortForOwner) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumeResumeShortForOwner) HasArea() bool`

HasArea returns a boolean if a field has been set.

### SetAreaNil

`func (o *ResumeResumeShortForOwner) SetAreaNil(b bool)`

 SetAreaNil sets the value for Area to be an explicit nil

### UnsetArea
`func (o *ResumeResumeShortForOwner) UnsetArea()`

UnsetArea ensures that no value is present for Area, not even an explicit nil
### GetAutoHideTime

`func (o *ResumeResumeShortForOwner) GetAutoHideTime() IncludesIdName`

GetAutoHideTime returns the AutoHideTime field if non-nil, zero value otherwise.

### GetAutoHideTimeOk

`func (o *ResumeResumeShortForOwner) GetAutoHideTimeOk() (*IncludesIdName, bool)`

GetAutoHideTimeOk returns a tuple with the AutoHideTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoHideTime

`func (o *ResumeResumeShortForOwner) SetAutoHideTime(v IncludesIdName)`

SetAutoHideTime sets AutoHideTime field to given value.

### HasAutoHideTime

`func (o *ResumeResumeShortForOwner) HasAutoHideTime() bool`

HasAutoHideTime returns a boolean if a field has been set.

### SetAutoHideTimeNil

`func (o *ResumeResumeShortForOwner) SetAutoHideTimeNil(b bool)`

 SetAutoHideTimeNil sets the value for AutoHideTime to be an explicit nil

### UnsetAutoHideTime
`func (o *ResumeResumeShortForOwner) UnsetAutoHideTime()`

UnsetAutoHideTime ensures that no value is present for AutoHideTime, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumeResumeShortForOwner) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumeResumeShortForOwner) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumeResumeShortForOwner) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumeResumeShortForOwner) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumeResumeShortForOwner) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumeResumeShortForOwner) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumeResumeShortForOwner) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumeResumeShortForOwner) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumeResumeShortForOwner) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumeResumeShortForOwner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumeResumeShortForOwner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumeResumeShortForOwner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumeResumeShortForOwner) GetDownload() ResumeObjectsDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeResumeShortForOwner) GetDownloadOk() (*ResumeObjectsDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeResumeShortForOwner) SetDownload(v ResumeObjectsDownload)`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumeResumeShortForOwner) GetEducation() ResumeObjectsEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumeResumeShortForOwner) GetEducationOk() (*ResumeObjectsEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumeResumeShortForOwner) SetEducation(v ResumeObjectsEducation)`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumeResumeShortForOwner) GetExperience() []ResumeObjectsExperienceForOwner`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumeResumeShortForOwner) GetExperienceOk() (*[]ResumeObjectsExperienceForOwner, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumeResumeShortForOwner) SetExperience(v []ResumeObjectsExperienceForOwner)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumeResumeShortForOwner) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumeResumeShortForOwner) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumeResumeShortForOwner) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumeResumeShortForOwner) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumeResumeShortForOwner) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumeResumeShortForOwner) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumeResumeShortForOwner) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumeResumeShortForOwner) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumeResumeShortForOwner) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumeResumeShortForOwner) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ResumeResumeShortForOwner) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ResumeResumeShortForOwner) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetHiddenFields

`func (o *ResumeResumeShortForOwner) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumeResumeShortForOwner) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumeResumeShortForOwner) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetId

`func (o *ResumeResumeShortForOwner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumeResumeShortForOwner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumeResumeShortForOwner) SetId(v string)`

SetId sets Id field to given value.


### GetLastName

`func (o *ResumeResumeShortForOwner) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumeResumeShortForOwner) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumeResumeShortForOwner) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumeResumeShortForOwner) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumeResumeShortForOwner) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumeResumeShortForOwner) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumeResumeShortForOwner) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumeResumeShortForOwner) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumeResumeShortForOwner) SetMarked(v bool)`

SetMarked sets Marked field to given value.


### GetMiddleName

`func (o *ResumeResumeShortForOwner) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumeResumeShortForOwner) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumeResumeShortForOwner) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumeResumeShortForOwner) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumeResumeShortForOwner) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumeResumeShortForOwner) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPhoto

`func (o *ResumeResumeShortForOwner) GetPhoto() ProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeResumeShortForOwner) GetPhotoOk() (*ProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeResumeShortForOwner) SetPhoto(v ProfilePhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeResumeShortForOwner) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeResumeShortForOwner) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeResumeShortForOwner) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetPlatform

`func (o *ResumeResumeShortForOwner) GetPlatform() IncludesId`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumeResumeShortForOwner) GetPlatformOk() (*IncludesId, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumeResumeShortForOwner) SetPlatform(v IncludesId)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumeResumeShortForOwner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumeResumeShortForOwner) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumeResumeShortForOwner) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumeResumeShortForOwner) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumeResumeShortForOwner) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumeResumeShortForOwner) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumeResumeShortForOwner) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTitle

`func (o *ResumeResumeShortForOwner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumeResumeShortForOwner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumeResumeShortForOwner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResumeResumeShortForOwner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ResumeResumeShortForOwner) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumeResumeShortForOwner) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalExperience

`func (o *ResumeResumeShortForOwner) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumeResumeShortForOwner) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumeResumeShortForOwner) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumeResumeShortForOwner) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumeResumeShortForOwner) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumeResumeShortForOwner) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumeResumeShortForOwner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumeResumeShortForOwner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumeResumeShortForOwner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *ResumeResumeShortForOwner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumeResumeShortForOwner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumeResumeShortForOwner) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


