# ResumesMineItem

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
**Blocked** | **bool** | Заблокировано ли резюме ([подробнее](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume)) | 
**CanPublishOrUpdate** | Pointer to **NullableBool** | Можно ли опубликовать или обновить данное резюме | [optional] 
**Finished** | **bool** | Заполнено ли резюме | 
**Status** | [**IncludesIdName**](IncludesIdName.md) |  | 
**Access** | [**ResumeObjectsAccess**](ResumeObjectsAccess.md) |  | 
**NewViews** | **float32** | Число новых просмотров. Данный счетчик сбрасывается при получении [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history)  | 
**NextPublishAt** | Pointer to **NullableString** | Дата и время следующей возможной публикации/обновления. Для неопубликованных резюме возвращается &#x60;null&#x60; | [optional] 
**PaidServices** | [**[]ResumeObjectsPaidServices**](ResumeObjectsPaidServices.md) | Платные услуги по резюме для автора | 
**TotalViews** | **float32** | Число просмотров резюме | 
**ViewsUrl** | **string** | URL, по которому необходимо сделать GET-запрос для получения [детальной истории просмотров](#tag/Rezyume.-Prosmotr-informacii/operation/get-resume-view-history)  | 
**Contact** | [**[]IncludesContact**](IncludesContact.md) | Список контактов соискателя | 
**Created** | **string** | Дата и время создания резюме | 
**SimilarVacancies** | [**ResumeObjectsSimilarVacancies**](ResumeObjectsSimilarVacancies.md) |  | 
**Updated** | **string** | Дата и время обновления резюме | 
**Visible** | **bool** | Видно ли резюме в поиске | 

## Methods

### NewResumesMineItem

`func NewResumesMineItem(actions ResumeObjectsActionsForOwner, alternateUrl string, certificate []ResumeObjectsCertificate, createdAt string, download ResumeObjectsDownload, education ResumeObjectsEducation, experience []ResumeObjectsExperienceForOwner, hiddenFields []IncludesIdName, id string, marked bool, updatedAt string, url string, blocked bool, finished bool, status IncludesIdName, access ResumeObjectsAccess, newViews float32, paidServices []ResumeObjectsPaidServices, totalViews float32, viewsUrl string, contact []IncludesContact, created string, similarVacancies ResumeObjectsSimilarVacancies, updated string, visible bool, ) *ResumesMineItem`

NewResumesMineItem instantiates a new ResumesMineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesMineItemWithDefaults

`func NewResumesMineItemWithDefaults() *ResumesMineItem`

NewResumesMineItemWithDefaults instantiates a new ResumesMineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ResumesMineItem) GetActions() ResumeObjectsActionsForOwner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumesMineItem) GetActionsOk() (*ResumeObjectsActionsForOwner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumesMineItem) SetActions(v ResumeObjectsActionsForOwner)`

SetActions sets Actions field to given value.


### GetAge

`func (o *ResumesMineItem) GetAge() float32`

GetAge returns the Age field if non-nil, zero value otherwise.

### GetAgeOk

`func (o *ResumesMineItem) GetAgeOk() (*float32, bool)`

GetAgeOk returns a tuple with the Age field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAge

`func (o *ResumesMineItem) SetAge(v float32)`

SetAge sets Age field to given value.

### HasAge

`func (o *ResumesMineItem) HasAge() bool`

HasAge returns a boolean if a field has been set.

### SetAgeNil

`func (o *ResumesMineItem) SetAgeNil(b bool)`

 SetAgeNil sets the value for Age to be an explicit nil

### UnsetAge
`func (o *ResumesMineItem) UnsetAge()`

UnsetAge ensures that no value is present for Age, not even an explicit nil
### GetAlternateUrl

`func (o *ResumesMineItem) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *ResumesMineItem) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *ResumesMineItem) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.


### GetArea

`func (o *ResumesMineItem) GetArea() IncludesIdNameUrl`

GetArea returns the Area field if non-nil, zero value otherwise.

### GetAreaOk

`func (o *ResumesMineItem) GetAreaOk() (*IncludesIdNameUrl, bool)`

GetAreaOk returns a tuple with the Area field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArea

`func (o *ResumesMineItem) SetArea(v IncludesIdNameUrl)`

SetArea sets Area field to given value.

### HasArea

`func (o *ResumesMineItem) HasArea() bool`

HasArea returns a boolean if a field has been set.

### GetAutoHideTime

`func (o *ResumesMineItem) GetAutoHideTime() IncludesIdName`

GetAutoHideTime returns the AutoHideTime field if non-nil, zero value otherwise.

### GetAutoHideTimeOk

`func (o *ResumesMineItem) GetAutoHideTimeOk() (*IncludesIdName, bool)`

GetAutoHideTimeOk returns a tuple with the AutoHideTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoHideTime

`func (o *ResumesMineItem) SetAutoHideTime(v IncludesIdName)`

SetAutoHideTime sets AutoHideTime field to given value.

### HasAutoHideTime

`func (o *ResumesMineItem) HasAutoHideTime() bool`

HasAutoHideTime returns a boolean if a field has been set.

### SetAutoHideTimeNil

`func (o *ResumesMineItem) SetAutoHideTimeNil(b bool)`

 SetAutoHideTimeNil sets the value for AutoHideTime to be an explicit nil

### UnsetAutoHideTime
`func (o *ResumesMineItem) UnsetAutoHideTime()`

UnsetAutoHideTime ensures that no value is present for AutoHideTime, not even an explicit nil
### GetCanViewFullInfo

`func (o *ResumesMineItem) GetCanViewFullInfo() bool`

GetCanViewFullInfo returns the CanViewFullInfo field if non-nil, zero value otherwise.

### GetCanViewFullInfoOk

`func (o *ResumesMineItem) GetCanViewFullInfoOk() (*bool, bool)`

GetCanViewFullInfoOk returns a tuple with the CanViewFullInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanViewFullInfo

`func (o *ResumesMineItem) SetCanViewFullInfo(v bool)`

SetCanViewFullInfo sets CanViewFullInfo field to given value.

### HasCanViewFullInfo

`func (o *ResumesMineItem) HasCanViewFullInfo() bool`

HasCanViewFullInfo returns a boolean if a field has been set.

### SetCanViewFullInfoNil

`func (o *ResumesMineItem) SetCanViewFullInfoNil(b bool)`

 SetCanViewFullInfoNil sets the value for CanViewFullInfo to be an explicit nil

### UnsetCanViewFullInfo
`func (o *ResumesMineItem) UnsetCanViewFullInfo()`

UnsetCanViewFullInfo ensures that no value is present for CanViewFullInfo, not even an explicit nil
### GetCertificate

`func (o *ResumesMineItem) GetCertificate() []ResumeObjectsCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ResumesMineItem) GetCertificateOk() (*[]ResumeObjectsCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ResumesMineItem) SetCertificate(v []ResumeObjectsCertificate)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *ResumesMineItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumesMineItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumesMineItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDownload

`func (o *ResumesMineItem) GetDownload() ResumeObjectsDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumesMineItem) GetDownloadOk() (*ResumeObjectsDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumesMineItem) SetDownload(v ResumeObjectsDownload)`

SetDownload sets Download field to given value.


### GetEducation

`func (o *ResumesMineItem) GetEducation() ResumeObjectsEducation`

GetEducation returns the Education field if non-nil, zero value otherwise.

### GetEducationOk

`func (o *ResumesMineItem) GetEducationOk() (*ResumeObjectsEducation, bool)`

GetEducationOk returns a tuple with the Education field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducation

`func (o *ResumesMineItem) SetEducation(v ResumeObjectsEducation)`

SetEducation sets Education field to given value.


### GetExperience

`func (o *ResumesMineItem) GetExperience() []ResumeObjectsExperienceForOwner`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *ResumesMineItem) GetExperienceOk() (*[]ResumeObjectsExperienceForOwner, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *ResumesMineItem) SetExperience(v []ResumeObjectsExperienceForOwner)`

SetExperience sets Experience field to given value.


### GetFirstName

`func (o *ResumesMineItem) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ResumesMineItem) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ResumesMineItem) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ResumesMineItem) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *ResumesMineItem) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *ResumesMineItem) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetGender

`func (o *ResumesMineItem) GetGender() IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ResumesMineItem) GetGenderOk() (*IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ResumesMineItem) SetGender(v IncludesIdName)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ResumesMineItem) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetHiddenFields

`func (o *ResumesMineItem) GetHiddenFields() []IncludesIdName`

GetHiddenFields returns the HiddenFields field if non-nil, zero value otherwise.

### GetHiddenFieldsOk

`func (o *ResumesMineItem) GetHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetHiddenFieldsOk returns a tuple with the HiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFields

`func (o *ResumesMineItem) SetHiddenFields(v []IncludesIdName)`

SetHiddenFields sets HiddenFields field to given value.


### GetId

`func (o *ResumesMineItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumesMineItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumesMineItem) SetId(v string)`

SetId sets Id field to given value.


### GetLastName

`func (o *ResumesMineItem) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ResumesMineItem) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ResumesMineItem) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ResumesMineItem) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *ResumesMineItem) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *ResumesMineItem) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetMarked

`func (o *ResumesMineItem) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumesMineItem) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumesMineItem) SetMarked(v bool)`

SetMarked sets Marked field to given value.


### GetMiddleName

`func (o *ResumesMineItem) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *ResumesMineItem) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *ResumesMineItem) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.

### HasMiddleName

`func (o *ResumesMineItem) HasMiddleName() bool`

HasMiddleName returns a boolean if a field has been set.

### SetMiddleNameNil

`func (o *ResumesMineItem) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *ResumesMineItem) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil
### GetPhoto

`func (o *ResumesMineItem) GetPhoto() ProfilePhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumesMineItem) GetPhotoOk() (*ProfilePhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumesMineItem) SetPhoto(v ProfilePhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumesMineItem) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### GetPlatform

`func (o *ResumesMineItem) GetPlatform() IncludesId`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ResumesMineItem) GetPlatformOk() (*IncludesId, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ResumesMineItem) SetPlatform(v IncludesId)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ResumesMineItem) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSalary

`func (o *ResumesMineItem) GetSalary() ResumeObjectsSalaryProperties`

GetSalary returns the Salary field if non-nil, zero value otherwise.

### GetSalaryOk

`func (o *ResumesMineItem) GetSalaryOk() (*ResumeObjectsSalaryProperties, bool)`

GetSalaryOk returns a tuple with the Salary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalary

`func (o *ResumesMineItem) SetSalary(v ResumeObjectsSalaryProperties)`

SetSalary sets Salary field to given value.

### HasSalary

`func (o *ResumesMineItem) HasSalary() bool`

HasSalary returns a boolean if a field has been set.

### SetSalaryNil

`func (o *ResumesMineItem) SetSalaryNil(b bool)`

 SetSalaryNil sets the value for Salary to be an explicit nil

### UnsetSalary
`func (o *ResumesMineItem) UnsetSalary()`

UnsetSalary ensures that no value is present for Salary, not even an explicit nil
### GetTitle

`func (o *ResumesMineItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResumesMineItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResumesMineItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResumesMineItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ResumesMineItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResumesMineItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTotalExperience

`func (o *ResumesMineItem) GetTotalExperience() ResumeObjectsTotalExperience`

GetTotalExperience returns the TotalExperience field if non-nil, zero value otherwise.

### GetTotalExperienceOk

`func (o *ResumesMineItem) GetTotalExperienceOk() (*ResumeObjectsTotalExperience, bool)`

GetTotalExperienceOk returns a tuple with the TotalExperience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExperience

`func (o *ResumesMineItem) SetTotalExperience(v ResumeObjectsTotalExperience)`

SetTotalExperience sets TotalExperience field to given value.

### HasTotalExperience

`func (o *ResumesMineItem) HasTotalExperience() bool`

HasTotalExperience returns a boolean if a field has been set.

### SetTotalExperienceNil

`func (o *ResumesMineItem) SetTotalExperienceNil(b bool)`

 SetTotalExperienceNil sets the value for TotalExperience to be an explicit nil

### UnsetTotalExperience
`func (o *ResumesMineItem) UnsetTotalExperience()`

UnsetTotalExperience ensures that no value is present for TotalExperience, not even an explicit nil
### GetUpdatedAt

`func (o *ResumesMineItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResumesMineItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResumesMineItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *ResumesMineItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumesMineItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumesMineItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetBlocked

`func (o *ResumesMineItem) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *ResumesMineItem) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *ResumesMineItem) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetCanPublishOrUpdate

`func (o *ResumesMineItem) GetCanPublishOrUpdate() bool`

GetCanPublishOrUpdate returns the CanPublishOrUpdate field if non-nil, zero value otherwise.

### GetCanPublishOrUpdateOk

`func (o *ResumesMineItem) GetCanPublishOrUpdateOk() (*bool, bool)`

GetCanPublishOrUpdateOk returns a tuple with the CanPublishOrUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPublishOrUpdate

`func (o *ResumesMineItem) SetCanPublishOrUpdate(v bool)`

SetCanPublishOrUpdate sets CanPublishOrUpdate field to given value.

### HasCanPublishOrUpdate

`func (o *ResumesMineItem) HasCanPublishOrUpdate() bool`

HasCanPublishOrUpdate returns a boolean if a field has been set.

### SetCanPublishOrUpdateNil

`func (o *ResumesMineItem) SetCanPublishOrUpdateNil(b bool)`

 SetCanPublishOrUpdateNil sets the value for CanPublishOrUpdate to be an explicit nil

### UnsetCanPublishOrUpdate
`func (o *ResumesMineItem) UnsetCanPublishOrUpdate()`

UnsetCanPublishOrUpdate ensures that no value is present for CanPublishOrUpdate, not even an explicit nil
### GetFinished

`func (o *ResumesMineItem) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ResumesMineItem) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ResumesMineItem) SetFinished(v bool)`

SetFinished sets Finished field to given value.


### GetStatus

`func (o *ResumesMineItem) GetStatus() IncludesIdName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResumesMineItem) GetStatusOk() (*IncludesIdName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResumesMineItem) SetStatus(v IncludesIdName)`

SetStatus sets Status field to given value.


### GetAccess

`func (o *ResumesMineItem) GetAccess() ResumeObjectsAccess`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ResumesMineItem) GetAccessOk() (*ResumeObjectsAccess, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ResumesMineItem) SetAccess(v ResumeObjectsAccess)`

SetAccess sets Access field to given value.


### GetNewViews

`func (o *ResumesMineItem) GetNewViews() float32`

GetNewViews returns the NewViews field if non-nil, zero value otherwise.

### GetNewViewsOk

`func (o *ResumesMineItem) GetNewViewsOk() (*float32, bool)`

GetNewViewsOk returns a tuple with the NewViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewViews

`func (o *ResumesMineItem) SetNewViews(v float32)`

SetNewViews sets NewViews field to given value.


### GetNextPublishAt

`func (o *ResumesMineItem) GetNextPublishAt() string`

GetNextPublishAt returns the NextPublishAt field if non-nil, zero value otherwise.

### GetNextPublishAtOk

`func (o *ResumesMineItem) GetNextPublishAtOk() (*string, bool)`

GetNextPublishAtOk returns a tuple with the NextPublishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPublishAt

`func (o *ResumesMineItem) SetNextPublishAt(v string)`

SetNextPublishAt sets NextPublishAt field to given value.

### HasNextPublishAt

`func (o *ResumesMineItem) HasNextPublishAt() bool`

HasNextPublishAt returns a boolean if a field has been set.

### SetNextPublishAtNil

`func (o *ResumesMineItem) SetNextPublishAtNil(b bool)`

 SetNextPublishAtNil sets the value for NextPublishAt to be an explicit nil

### UnsetNextPublishAt
`func (o *ResumesMineItem) UnsetNextPublishAt()`

UnsetNextPublishAt ensures that no value is present for NextPublishAt, not even an explicit nil
### GetPaidServices

`func (o *ResumesMineItem) GetPaidServices() []ResumeObjectsPaidServices`

GetPaidServices returns the PaidServices field if non-nil, zero value otherwise.

### GetPaidServicesOk

`func (o *ResumesMineItem) GetPaidServicesOk() (*[]ResumeObjectsPaidServices, bool)`

GetPaidServicesOk returns a tuple with the PaidServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidServices

`func (o *ResumesMineItem) SetPaidServices(v []ResumeObjectsPaidServices)`

SetPaidServices sets PaidServices field to given value.


### GetTotalViews

`func (o *ResumesMineItem) GetTotalViews() float32`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *ResumesMineItem) GetTotalViewsOk() (*float32, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *ResumesMineItem) SetTotalViews(v float32)`

SetTotalViews sets TotalViews field to given value.


### GetViewsUrl

`func (o *ResumesMineItem) GetViewsUrl() string`

GetViewsUrl returns the ViewsUrl field if non-nil, zero value otherwise.

### GetViewsUrlOk

`func (o *ResumesMineItem) GetViewsUrlOk() (*string, bool)`

GetViewsUrlOk returns a tuple with the ViewsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewsUrl

`func (o *ResumesMineItem) SetViewsUrl(v string)`

SetViewsUrl sets ViewsUrl field to given value.


### GetContact

`func (o *ResumesMineItem) GetContact() []IncludesContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *ResumesMineItem) GetContactOk() (*[]IncludesContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *ResumesMineItem) SetContact(v []IncludesContact)`

SetContact sets Contact field to given value.


### GetCreated

`func (o *ResumesMineItem) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResumesMineItem) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResumesMineItem) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetSimilarVacancies

`func (o *ResumesMineItem) GetSimilarVacancies() ResumeObjectsSimilarVacancies`

GetSimilarVacancies returns the SimilarVacancies field if non-nil, zero value otherwise.

### GetSimilarVacanciesOk

`func (o *ResumesMineItem) GetSimilarVacanciesOk() (*ResumeObjectsSimilarVacancies, bool)`

GetSimilarVacanciesOk returns a tuple with the SimilarVacancies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimilarVacancies

`func (o *ResumesMineItem) SetSimilarVacancies(v ResumeObjectsSimilarVacancies)`

SetSimilarVacancies sets SimilarVacancies field to given value.


### GetUpdated

`func (o *ResumesMineItem) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ResumesMineItem) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ResumesMineItem) SetUpdated(v string)`

SetUpdated sets Updated field to given value.


### GetVisible

`func (o *ResumesMineItem) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ResumesMineItem) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ResumesMineItem) SetVisible(v bool)`

SetVisible sets Visible field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


