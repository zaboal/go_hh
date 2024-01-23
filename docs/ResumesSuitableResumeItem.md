# ResumesSuitableResumeItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**ResumeObjectsActionsForOwner**](ResumeObjectsActionsForOwner.md) | Дополнительные действия | 
**Age** | Pointer to **NullableFloat32** | Возраст | [optional] 
**AlternateUrl** | **string** | URL резюме на сайте | 
**Area** | Pointer to [**IncludesIdNameUrl**](IncludesIdNameUrl.md) |  | [optional] 
**AutoHideTime** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**CanViewFullInfo** | Pointer to **NullableBool** | Доступен ли просмотр контактной информации в резюме текущему работодателю | [optional] 
**Certificate** | [**[]ResumeObjectsCertificate**](ResumeObjectsCertificate.md) | Список сертификатов соискателя | 
**CreatedAt** | **string** | Дата и время создания резюме | 
**Download** | [**ResumeObjectsDownload**](ResumeObjectsDownload.md) | Ссылки для скачивания резюме в разных форматах | 
**Education** | [**ResumeObjectsEducation**](ResumeObjectsEducation.md) | Образование соискателя. 

Особенности сохранения образования:

* Если передать и высшее и среднее образование и уровень образования &quot;средний&quot;, то сохранится только среднее образование.
* Если передать и высшее и среднее образование и уровень образования &quot;высшее&quot;, то сохранится только высшее образование
 | 
**Experience** | [**[]ResumeObjectsExperienceForOwner**](ResumeObjectsExperienceForOwner.md) | Опыт работы | 
**FirstName** | Pointer to **NullableString** | Имя | [optional] 
**Gender** | Pointer to [**IncludesIdName**](IncludesIdName.md) |  | [optional] 
**HiddenFields** | [**[]IncludesIdName**](IncludesIdName.md) | Документация [Список скрытых полей](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields). Возможные значения элементов приведены в поле &#x60;resume_hidden_fields&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Id** | **string** | Идентификатор резюме | 
**LastName** | Pointer to **NullableString** | Фамилия | [optional] 
**Marked** | **bool** | Выделено ли резюме в поиске | 
**MiddleName** | Pointer to **NullableString** | Отчество | [optional] 
**Photo** | Pointer to [**ProfilePhoto**](ProfilePhoto.md) |  | [optional] 
**Platform** | Pointer to [**IncludesId**](IncludesId.md) | Ресурс, на котором было размещено резюме | [optional] 
**Salary** | Pointer to [**NullableResumeObjectsSalaryProperties**](ResumeObjectsSalaryProperties.md) |  | [optional] 
**Title** | Pointer to **NullableString** | Желаемая должность | [optional] 
**TotalExperience** | Pointer to [**NullableResumeObjectsTotalExperience**](ResumeObjectsTotalExperience.md) |  | [optional] 
**UpdatedAt** | **string** | Дата и время обновления резюме | 
**Url** | **string** | URL резюме на сайте | 
**Access** | [**ResumeObjectsAccess**](ResumeObjectsAccess.md) |  | 
**Finished** | **bool** | Заполнено ли резюме | 
**RequiresCompletion** | **bool** | Принимает значение &#x60;true&#x60;, если резюме является неполным. Применимо только для вакансий, у которых не установлен флаг «принимать неполные резюме».   При получении &#x60;true&#x60; в данном поле, соискатель должен заполнить обязательные поля (доступны в [выдаче полного резюме](#tag/Prosmotr-rezyume/operation/get-resume)) перед откликом на данную вакансию  | 
**Status** | [**IncludesIdName**](IncludesIdName.md) |  | 

## Methods

### NewResumesSuitableResumeItem

`func NewResumesSuitableResumeItem(actions ResumeObjectsActionsForOwner, alternateUrl string, certificate []ResumeObjectsCertificate, createdAt string, download ResumeObjectsDownload, education ResumeObjectsEducation, experience []ResumeObjectsExperienceForOwner, hiddenFields []IncludesIdName, id string, marked bool, updatedAt string, url string, access ResumeObjectsAccess, finished bool, requiresCompletion bool, status IncludesIdName, ) *ResumesSuitableResumeItem`

NewResumesSuitableResumeItem instantiates a new ResumesSuitableResumeItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesSuitableResumeItemWithDefaults

`func NewResumesSuitableResumeItemWithDefaults() *ResumesSuitableResumeItem`

NewResumesSuitableResumeItemWithDefaults instantiates a new ResumesSuitableResumeItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ResumesSuitableResumeItem) GetActions() ResumeObjectsActionsForOwner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumesSuitableResumeItem) GetActionsOk() (*ResumeObjectsActionsForOwner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumesSuitableResumeItem) SetActions(v ResumeObjectsActionsForOwner)`

SetActions sets Actions field to given value.


### GetAge

`func (o *ResumesSuitableResumeItem) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumesSuitableResumeItem) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumesSuitableResumeItem) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumesSuitableResumeItem) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumesSuitableResumeItem) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumesSuitableResumeItem) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetAlternateUrl

`func (o *ResumesSuitableResumeItem) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumesSuitableResumeItem) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumesSuitableResumeItem) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetArea

`func (o *ResumesSuitableResumeItem) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumesSuitableResumeItem) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumesSuitableResumeItem) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumesSuitableResumeItem) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetAutoHideTime

`func (o *ResumesSuitableResumeItem) GetAutoHideTime() IncludesIdName`

GetAutoHideTime returns the AutoHideTime field if non-nil, zero value otherwise.

### GetAutoHideTimeOk

`func (o *ResumesSuitableResumeItem) GetAutoHideTimeOk() (*IncludesIdName, bool)`

GetAutoHideTimeOk returns a tuple with the AutoHideTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoHideTime

`func (o *ResumesSuitableResumeItem) SetAutoHideTime(v IncludesIdName)`

SetAutoHideTime sets AutoHideTime field to given value.

### HasAutoHideTime

`func (o *ResumesSuitableResumeItem) HasAutoHideTime() bool`

HasAutoHideTime returns a boolean if a field has been set.

### SetAutoHideTimeNil

`func (o *ResumesSuitableResumeItem) SetAutoHideTimeNil(b bool)`

 SetAutoHideTimeNil sets the value for AutoHideTime to be an explicit nil

### UnsetAutoHideTime
`func (o *ResumesSuitableResumeItem) UnsetAutoHideTime()`

UnsetAutoHideTime ensures that no value is present for AutoHideTime, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumesSuitableResumeItem) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumesSuitableResumeItem) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumesSuitableResumeItem) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumesSuitableResumeItem) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumesSuitableResumeItem) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumesSuitableResumeItem) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumesSuitableResumeItem) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumesSuitableResumeItem) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumesSuitableResumeItem) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumesSuitableResumeItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumesSuitableResumeItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumesSuitableResumeItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumesSuitableResumeItem) GetDownload() ResumeObjectsDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumesSuitableResumeItem) GetDownloadOk() (*ResumeObjectsDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumesSuitableResumeItem) SetDownload(v ResumeObjectsDownload)`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumesSuitableResumeItem) GetEducation() ResumeObjectsEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumesSuitableResumeItem) GetEducationOk() (*ResumeObjectsEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumesSuitableResumeItem) SetEducation(v ResumeObjectsEducation)`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumesSuitableResumeItem) GetExperience() []ResumeObjectsExperienceForOwner`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumesSuitableResumeItem) GetExperienceOk() (*[]ResumeObjectsExperienceForOwner, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumesSuitableResumeItem) SetExperience(v []ResumeObjectsExperienceForOwner)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumesSuitableResumeItem) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumesSuitableResumeItem) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumesSuitableResumeItem) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumesSuitableResumeItem) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumesSuitableResumeItem) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumesSuitableResumeItem) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumesSuitableResumeItem) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumesSuitableResumeItem) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumesSuitableResumeItem) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumesSuitableResumeItem) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetHiddenFields

`func (o *ResumesSuitableResumeItem) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumesSuitableResumeItem) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumesSuitableResumeItem) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetId

`func (o *ResumesSuitableResumeItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumesSuitableResumeItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumesSuitableResumeItem) SetId(v string)`

SetId sets Id field to given value.


### GetLastName

`func (o *ResumesSuitableResumeItem) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumesSuitableResumeItem) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumesSuitableResumeItem) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumesSuitableResumeItem) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumesSuitableResumeItem) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumesSuitableResumeItem) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumesSuitableResumeItem) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumesSuitableResumeItem) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumesSuitableResumeItem) SetMarked(v bool)`

SetMarked sets Marked field to given value.


### GetMiddleName

`func (o *ResumesSuitableResumeItem) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumesSuitableResumeItem) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumesSuitableResumeItem) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumesSuitableResumeItem) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumesSuitableResumeItem) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumesSuitableResumeItem) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPhoto

`func (o *ResumesSuitableResumeItem) GetPhoto() ProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumesSuitableResumeItem) GetPhotoOk() (*ProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumesSuitableResumeItem) SetPhoto(v ProfilePhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumesSuitableResumeItem) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetPlatform

`func (o *ResumesSuitableResumeItem) GetPlatform() IncludesId`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumesSuitableResumeItem) GetPlatformOk() (*IncludesId, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumesSuitableResumeItem) SetPlatform(v IncludesId)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumesSuitableResumeItem) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumesSuitableResumeItem) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumesSuitableResumeItem) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumesSuitableResumeItem) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumesSuitableResumeItem) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumesSuitableResumeItem) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumesSuitableResumeItem) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTitle

`func (o *ResumesSuitableResumeItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumesSuitableResumeItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumesSuitableResumeItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResumesSuitableResumeItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ResumesSuitableResumeItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumesSuitableResumeItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalExperience

`func (o *ResumesSuitableResumeItem) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumesSuitableResumeItem) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumesSuitableResumeItem) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumesSuitableResumeItem) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumesSuitableResumeItem) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumesSuitableResumeItem) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumesSuitableResumeItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumesSuitableResumeItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumesSuitableResumeItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *ResumesSuitableResumeItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumesSuitableResumeItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumesSuitableResumeItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetAccess

`func (o *ResumesSuitableResumeItem) GetAccess() ResumeObjectsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ResumesSuitableResumeItem) GetAccessOk() (*ResumeObjectsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ResumesSuitableResumeItem) SetAccess(v ResumeObjectsAccess)`

SetAccess sets Access field to given value.


### GetFinished

`func (o *ResumesSuitableResumeItem) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ResumesSuitableResumeItem) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ResumesSuitableResumeItem) SetFinished(v bool)`

SetFinished sets Finished field to given value.


### GetRequiresCompletion

`func (o *ResumesSuitableResumeItem) GetRequiresCompletion() bool`

GetRequiresCompletion returns the RequiresCompletion field if non-nil, zero value otherwise.

### GetRequiresCompletionOk

`func (o *ResumesSuitableResumeItem) GetRequiresCompletionOk() (*bool, bool)`

GetRequiresCompletionOk returns a tuple with the RequiresCompletion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresCompletion

`func (o *ResumesSuitableResumeItem) SetRequiresCompletion(v bool)`

SetRequiresCompletion sets RequiresCompletion field to given value.


### GetStatus

`func (o *ResumesSuitableResumeItem) GetStatus() IncludesIdName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResumesSuitableResumeItem) GetStatusOk() (*IncludesIdName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResumesSuitableResumeItem) SetStatus(v IncludesIdName)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


