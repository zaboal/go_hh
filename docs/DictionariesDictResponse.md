# DictionariesDictResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantCommentAccessType** | [**[]IncludesIdName**](IncludesIdName.md) | Тип доступа для комментария к соискателю | 
**ApplicantCommentsOrder** | [**[]IncludesIdName**](IncludesIdName.md) | Типы сортировки [списка комментариев к соискателю](#tag/Kommentarii-k-soiskatelyu/operation/get-applicant-comments-list) | 
**ApplicantNegotiationStatus** | [**[]IncludesIdName**](IncludesIdName.md) | Статусы откликов/приглашений | 
**BusinessTripReadiness** | [**[]IncludesIdName**](IncludesIdName.md) | Готовность к командировкам | 
**Currency** | [**[]DictionariesCurrencyItem**](DictionariesCurrencyItem.md) | Справочник валют | 
**DriverLicenseTypes** | [**[]IncludesId**](IncludesId.md) | Категории водительских прав | 
**EducationLevel** | [**[]IncludesIdName**](IncludesIdName.md) | Образование в резюме | 
**EmployerActiveVacanciesOrder** | [**[]IncludesIdName**](IncludesIdName.md) | Тип сортировки списка опубликованных вакансий | 
**EmployerArchivedVacanciesOrder** | [**[]IncludesIdName**](IncludesIdName.md) | Тип сортировки списка архивных вакансий | 
**EmployerHiddenVacanciesOrder** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Тип сортировки скрытых вакансий | [optional] 
**EmployerRelation** | [**[]IncludesIdName**](IncludesIdName.md) | Типы связи компании с пользователем | 
**EmployerType** | [**[]IncludesIdName**](IncludesIdName.md) | Тип работодателя | 
**Employment** | [**[]IncludesIdName**](IncludesIdName.md) | Тип занятости | 
**Experience** | [**[]IncludesIdName**](IncludesIdName.md) | Опыт работы | 
**Gender** | [**[]IncludesIdName**](IncludesIdName.md) | Пол | 
**JobSearchStatusesApplicant** | [**[]IncludesIdName**](IncludesIdName.md) | Статусы поиска соискателей для установки и отображения самому соискателю | 
**JobSearchStatusesEmployer** | [**[]IncludesIdName**](IncludesIdName.md) | Статусы поиска соискателей для отображения работодателям | 
**LanguageLevel** | [**[]IncludesIdName**](IncludesIdName.md) | Уровень владения языком | 
**MessagingStatus** | [**[]IncludesIdName**](IncludesIdName.md) | Статус возможности отправки сообщения в переписке | 
**NegotiationsOrder** | [**[]IncludesIdName**](IncludesIdName.md) | Типы порядка отображения откликов | 
**NegotiationsParticipantType** | [**[]IncludesIdName**](IncludesIdName.md) | Типы участников переписки | 
**NegotiationsState** | [**[]IncludesIdName**](IncludesIdName.md) | Типы состояний откликов | 
**PhoneCallStatus** | [**[]IncludesIdName**](IncludesIdName.md) | Статус звонка, зафиксированного в системе кол-трекинг | 
**PreferredContactType** | [**[]IncludesIdName**](IncludesIdName.md) | Предпочитаемый способ связи | 
**RelocationType** | [**[]IncludesIdName**](IncludesIdName.md) | Готовность к переезду | 
**ResumeAccessType** | [**[]IncludesIdName**](IncludesIdName.md) | Уровень доступа к резюме | 
**ResumeContactsSiteType** | [**[]IncludesIdName**](IncludesIdName.md) | Тип сайта в поле «контакты» | 
**ResumeHiddenFields** | [**[]IncludesIdName**](IncludesIdName.md) | Поля резюме, которые могут быть [скрыты](https://github.com/hhru/api/blob/master/docs/employer_resumes.md#hidden-fields) | 
**ResumeModerationNote** | [**[]IncludesIdName**](IncludesIdName.md) | Комментарий модератора | 
**ResumeSearchExperiencePeriod** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Условие поиска по опыту работы | [optional] 
**ResumeSearchFields** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Область поиска в резюме | [optional] 
**ResumeSearchLabel** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Метки поиска резюме | [optional] 
**ResumeSearchLogic** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Условие поиска резюме | [optional] 
**ResumeSearchOrder** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Тип сортировки резюме | [optional] 
**ResumeSearchRelocation** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Условие поиска по проживанию в регионе и готовности к переезду | [optional] 
**ResumeStatus** | [**[]IncludesIdName**](IncludesIdName.md) | Статус резюме | 
**Schedule** | [**[]IncludesIdNameUid**](IncludesIdNameUid.md) | График работы | 
**TravelTime** | [**[]IncludesIdName**](IncludesIdName.md) | Время в пути | 
**VacancyBillingType** | [**[]IncludesIdName**](IncludesIdName.md) | Варианты размещения вакансии с точки зрения биллинга | 
**VacancyCluster** | [**[]IncludesIdName**](IncludesIdName.md) | Тип кластеров | 
**VacancyLabel** | [**[]IncludesIdName**](IncludesIdName.md) | Метки вакансии | 
**VacancyNotProlongedReason** | [**[]IncludesIdName**](IncludesIdName.md) | Причины, из-за которых невозможно [продлить вакансию](#tag/Upravlenie-vakansiyami/operation/get-prolongation-vacancy-info) | 
**VacancyRelation** | [**[]IncludesIdName**](IncludesIdName.md) | Типы связи вакансии с пользователем | 
**VacancySearchFields** | [**[]IncludesIdName**](IncludesIdName.md) | Область поиска в вакансии | 
**VacancySearchOrder** | [**[]IncludesIdName**](IncludesIdName.md) | Тип сортировки вакансии | 
**VacancyType** | [**[]IncludesIdName**](IncludesIdName.md) | Тип вакансии | 
**WorkingDays** | [**[]IncludesIdName**](IncludesIdName.md) | Рабочие дни | 
**WorkingTimeIntervals** | [**[]IncludesIdName**](IncludesIdName.md) | Временные интервалы работы | 
**WorkingTimeModes** | [**[]IncludesIdName**](IncludesIdName.md) | Режимы времени работы | 

## Methods

### NewDictionariesDictResponse

`func NewDictionariesDictResponse(applicantCommentAccessType []IncludesIdName, applicantCommentsOrder []IncludesIdName, applicantNegotiationStatus []IncludesIdName, businessTripReadiness []IncludesIdName, currency []DictionariesCurrencyItem, driverLicenseTypes []IncludesId, educationLevel []IncludesIdName, employerActiveVacanciesOrder []IncludesIdName, employerArchivedVacanciesOrder []IncludesIdName, employerRelation []IncludesIdName, employerType []IncludesIdName, employment []IncludesIdName, experience []IncludesIdName, gender []IncludesIdName, jobSearchStatusesApplicant []IncludesIdName, jobSearchStatusesEmployer []IncludesIdName, languageLevel []IncludesIdName, messagingStatus []IncludesIdName, negotiationsOrder []IncludesIdName, negotiationsParticipantType []IncludesIdName, negotiationsState []IncludesIdName, phoneCallStatus []IncludesIdName, preferredContactType []IncludesIdName, relocationType []IncludesIdName, resumeAccessType []IncludesIdName, resumeContactsSiteType []IncludesIdName, resumeHiddenFields []IncludesIdName, resumeModerationNote []IncludesIdName, resumeStatus []IncludesIdName, schedule []IncludesIdNameUid, travelTime []IncludesIdName, vacancyBillingType []IncludesIdName, vacancyCluster []IncludesIdName, vacancyLabel []IncludesIdName, vacancyNotProlongedReason []IncludesIdName, vacancyRelation []IncludesIdName, vacancySearchFields []IncludesIdName, vacancySearchOrder []IncludesIdName, vacancyType []IncludesIdName, workingDays []IncludesIdName, workingTimeIntervals []IncludesIdName, workingTimeModes []IncludesIdName, ) *DictionariesDictResponse`

NewDictionariesDictResponse instantiates a new DictionariesDictResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDictionariesDictResponseWithDefaults

`func NewDictionariesDictResponseWithDefaults() *DictionariesDictResponse`

NewDictionariesDictResponseWithDefaults instantiates a new DictionariesDictResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantCommentAccessType

`func (o *DictionariesDictResponse) GetApplicantCommentAccessType() []IncludesIdName`

GetApplicantCommentAccessType returns the ApplicantCommentAccessType field if non-nil, zero value otherwise.

### GetApplicantCommentAccessTypeOk

`func (o *DictionariesDictResponse) GetApplicantCommentAccessTypeOk() (*[]IncludesIdName, bool)`

GetApplicantCommentAccessTypeOk returns a tuple with the ApplicantCommentAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantCommentAccessType

`func (o *DictionariesDictResponse) SetApplicantCommentAccessType(v []IncludesIdName)`

SetApplicantCommentAccessType sets ApplicantCommentAccessType field to given value.


### GetApplicantCommentsOrder

`func (o *DictionariesDictResponse) GetApplicantCommentsOrder() []IncludesIdName`

GetApplicantCommentsOrder returns the ApplicantCommentsOrder field if non-nil, zero value otherwise.

### GetApplicantCommentsOrderOk

`func (o *DictionariesDictResponse) GetApplicantCommentsOrderOk() (*[]IncludesIdName, bool)`

GetApplicantCommentsOrderOk returns a tuple with the ApplicantCommentsOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantCommentsOrder

`func (o *DictionariesDictResponse) SetApplicantCommentsOrder(v []IncludesIdName)`

SetApplicantCommentsOrder sets ApplicantCommentsOrder field to given value.


### GetApplicantNegotiationStatus

`func (o *DictionariesDictResponse) GetApplicantNegotiationStatus() []IncludesIdName`

GetApplicantNegotiationStatus returns the ApplicantNegotiationStatus field if non-nil, zero value otherwise.

### GetApplicantNegotiationStatusOk

`func (o *DictionariesDictResponse) GetApplicantNegotiationStatusOk() (*[]IncludesIdName, bool)`

GetApplicantNegotiationStatusOk returns a tuple with the ApplicantNegotiationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantNegotiationStatus

`func (o *DictionariesDictResponse) SetApplicantNegotiationStatus(v []IncludesIdName)`

SetApplicantNegotiationStatus sets ApplicantNegotiationStatus field to given value.


### GetBusinessTripReadiness

`func (o *DictionariesDictResponse) GetBusinessTripReadiness() []IncludesIdName`

GetBusinessTripReadiness returns the BusinessTripReadiness field if non-nil, zero value otherwise.

### GetBusinessTripReadinessOk

`func (o *DictionariesDictResponse) GetBusinessTripReadinessOk() (*[]IncludesIdName, bool)`

GetBusinessTripReadinessOk returns a tuple with the BusinessTripReadiness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessTripReadiness

`func (o *DictionariesDictResponse) SetBusinessTripReadiness(v []IncludesIdName)`

SetBusinessTripReadiness sets BusinessTripReadiness field to given value.


### GetCurrency

`func (o *DictionariesDictResponse) GetCurrency() []DictionariesCurrencyItem`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *DictionariesDictResponse) GetCurrencyOk() (*[]DictionariesCurrencyItem, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *DictionariesDictResponse) SetCurrency(v []DictionariesCurrencyItem)`

SetCurrency sets Currency field to given value.


### GetDriverLicenseTypes

`func (o *DictionariesDictResponse) GetDriverLicenseTypes() []IncludesId`

GetDriverLicenseTypes returns the DriverLicenseTypes field if non-nil, zero value otherwise.

### GetDriverLicenseTypesOk

`func (o *DictionariesDictResponse) GetDriverLicenseTypesOk() (*[]IncludesId, bool)`

GetDriverLicenseTypesOk returns a tuple with the DriverLicenseTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriverLicenseTypes

`func (o *DictionariesDictResponse) SetDriverLicenseTypes(v []IncludesId)`

SetDriverLicenseTypes sets DriverLicenseTypes field to given value.


### GetEducationLevel

`func (o *DictionariesDictResponse) GetEducationLevel() []IncludesIdName`

GetEducationLevel returns the EducationLevel field if non-nil, zero value otherwise.

### GetEducationLevelOk

`func (o *DictionariesDictResponse) GetEducationLevelOk() (*[]IncludesIdName, bool)`

GetEducationLevelOk returns a tuple with the EducationLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEducationLevel

`func (o *DictionariesDictResponse) SetEducationLevel(v []IncludesIdName)`

SetEducationLevel sets EducationLevel field to given value.


### GetEmployerActiveVacanciesOrder

`func (o *DictionariesDictResponse) GetEmployerActiveVacanciesOrder() []IncludesIdName`

GetEmployerActiveVacanciesOrder returns the EmployerActiveVacanciesOrder field if non-nil, zero value otherwise.

### GetEmployerActiveVacanciesOrderOk

`func (o *DictionariesDictResponse) GetEmployerActiveVacanciesOrderOk() (*[]IncludesIdName, bool)`

GetEmployerActiveVacanciesOrderOk returns a tuple with the EmployerActiveVacanciesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerActiveVacanciesOrder

`func (o *DictionariesDictResponse) SetEmployerActiveVacanciesOrder(v []IncludesIdName)`

SetEmployerActiveVacanciesOrder sets EmployerActiveVacanciesOrder field to given value.


### GetEmployerArchivedVacanciesOrder

`func (o *DictionariesDictResponse) GetEmployerArchivedVacanciesOrder() []IncludesIdName`

GetEmployerArchivedVacanciesOrder returns the EmployerArchivedVacanciesOrder field if non-nil, zero value otherwise.

### GetEmployerArchivedVacanciesOrderOk

`func (o *DictionariesDictResponse) GetEmployerArchivedVacanciesOrderOk() (*[]IncludesIdName, bool)`

GetEmployerArchivedVacanciesOrderOk returns a tuple with the EmployerArchivedVacanciesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerArchivedVacanciesOrder

`func (o *DictionariesDictResponse) SetEmployerArchivedVacanciesOrder(v []IncludesIdName)`

SetEmployerArchivedVacanciesOrder sets EmployerArchivedVacanciesOrder field to given value.


### GetEmployerHiddenVacanciesOrder

`func (o *DictionariesDictResponse) GetEmployerHiddenVacanciesOrder() []IncludesIdName`

GetEmployerHiddenVacanciesOrder returns the EmployerHiddenVacanciesOrder field if non-nil, zero value otherwise.

### GetEmployerHiddenVacanciesOrderOk

`func (o *DictionariesDictResponse) GetEmployerHiddenVacanciesOrderOk() (*[]IncludesIdName, bool)`

GetEmployerHiddenVacanciesOrderOk returns a tuple with the EmployerHiddenVacanciesOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerHiddenVacanciesOrder

`func (o *DictionariesDictResponse) SetEmployerHiddenVacanciesOrder(v []IncludesIdName)`

SetEmployerHiddenVacanciesOrder sets EmployerHiddenVacanciesOrder field to given value.

### HasEmployerHiddenVacanciesOrder

`func (o *DictionariesDictResponse) HasEmployerHiddenVacanciesOrder() bool`

HasEmployerHiddenVacanciesOrder returns a boolean if a field has been set.

### GetEmployerRelation

`func (o *DictionariesDictResponse) GetEmployerRelation() []IncludesIdName`

GetEmployerRelation returns the EmployerRelation field if non-nil, zero value otherwise.

### GetEmployerRelationOk

`func (o *DictionariesDictResponse) GetEmployerRelationOk() (*[]IncludesIdName, bool)`

GetEmployerRelationOk returns a tuple with the EmployerRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerRelation

`func (o *DictionariesDictResponse) SetEmployerRelation(v []IncludesIdName)`

SetEmployerRelation sets EmployerRelation field to given value.


### GetEmployerType

`func (o *DictionariesDictResponse) GetEmployerType() []IncludesIdName`

GetEmployerType returns the EmployerType field if non-nil, zero value otherwise.

### GetEmployerTypeOk

`func (o *DictionariesDictResponse) GetEmployerTypeOk() (*[]IncludesIdName, bool)`

GetEmployerTypeOk returns a tuple with the EmployerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerType

`func (o *DictionariesDictResponse) SetEmployerType(v []IncludesIdName)`

SetEmployerType sets EmployerType field to given value.


### GetEmployment

`func (o *DictionariesDictResponse) GetEmployment() []IncludesIdName`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *DictionariesDictResponse) GetEmploymentOk() (*[]IncludesIdName, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *DictionariesDictResponse) SetEmployment(v []IncludesIdName)`

SetEmployment sets Employment field to given value.


### GetExperience

`func (o *DictionariesDictResponse) GetExperience() []IncludesIdName`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *DictionariesDictResponse) GetExperienceOk() (*[]IncludesIdName, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *DictionariesDictResponse) SetExperience(v []IncludesIdName)`

SetExperience sets Experience field to given value.


### GetGender

`func (o *DictionariesDictResponse) GetGender() []IncludesIdName`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *DictionariesDictResponse) GetGenderOk() (*[]IncludesIdName, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *DictionariesDictResponse) SetGender(v []IncludesIdName)`

SetGender sets Gender field to given value.


### GetJobSearchStatusesApplicant

`func (o *DictionariesDictResponse) GetJobSearchStatusesApplicant() []IncludesIdName`

GetJobSearchStatusesApplicant returns the JobSearchStatusesApplicant field if non-nil, zero value otherwise.

### GetJobSearchStatusesApplicantOk

`func (o *DictionariesDictResponse) GetJobSearchStatusesApplicantOk() (*[]IncludesIdName, bool)`

GetJobSearchStatusesApplicantOk returns a tuple with the JobSearchStatusesApplicant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatusesApplicant

`func (o *DictionariesDictResponse) SetJobSearchStatusesApplicant(v []IncludesIdName)`

SetJobSearchStatusesApplicant sets JobSearchStatusesApplicant field to given value.


### GetJobSearchStatusesEmployer

`func (o *DictionariesDictResponse) GetJobSearchStatusesEmployer() []IncludesIdName`

GetJobSearchStatusesEmployer returns the JobSearchStatusesEmployer field if non-nil, zero value otherwise.

### GetJobSearchStatusesEmployerOk

`func (o *DictionariesDictResponse) GetJobSearchStatusesEmployerOk() (*[]IncludesIdName, bool)`

GetJobSearchStatusesEmployerOk returns a tuple with the JobSearchStatusesEmployer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatusesEmployer

`func (o *DictionariesDictResponse) SetJobSearchStatusesEmployer(v []IncludesIdName)`

SetJobSearchStatusesEmployer sets JobSearchStatusesEmployer field to given value.


### GetLanguageLevel

`func (o *DictionariesDictResponse) GetLanguageLevel() []IncludesIdName`

GetLanguageLevel returns the LanguageLevel field if non-nil, zero value otherwise.

### GetLanguageLevelOk

`func (o *DictionariesDictResponse) GetLanguageLevelOk() (*[]IncludesIdName, bool)`

GetLanguageLevelOk returns a tuple with the LanguageLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageLevel

`func (o *DictionariesDictResponse) SetLanguageLevel(v []IncludesIdName)`

SetLanguageLevel sets LanguageLevel field to given value.


### GetMessagingStatus

`func (o *DictionariesDictResponse) GetMessagingStatus() []IncludesIdName`

GetMessagingStatus returns the MessagingStatus field if non-nil, zero value otherwise.

### GetMessagingStatusOk

`func (o *DictionariesDictResponse) GetMessagingStatusOk() (*[]IncludesIdName, bool)`

GetMessagingStatusOk returns a tuple with the MessagingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingStatus

`func (o *DictionariesDictResponse) SetMessagingStatus(v []IncludesIdName)`

SetMessagingStatus sets MessagingStatus field to given value.


### GetNegotiationsOrder

`func (o *DictionariesDictResponse) GetNegotiationsOrder() []IncludesIdName`

GetNegotiationsOrder returns the NegotiationsOrder field if non-nil, zero value otherwise.

### GetNegotiationsOrderOk

`func (o *DictionariesDictResponse) GetNegotiationsOrderOk() (*[]IncludesIdName, bool)`

GetNegotiationsOrderOk returns a tuple with the NegotiationsOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsOrder

`func (o *DictionariesDictResponse) SetNegotiationsOrder(v []IncludesIdName)`

SetNegotiationsOrder sets NegotiationsOrder field to given value.


### GetNegotiationsParticipantType

`func (o *DictionariesDictResponse) GetNegotiationsParticipantType() []IncludesIdName`

GetNegotiationsParticipantType returns the NegotiationsParticipantType field if non-nil, zero value otherwise.

### GetNegotiationsParticipantTypeOk

`func (o *DictionariesDictResponse) GetNegotiationsParticipantTypeOk() (*[]IncludesIdName, bool)`

GetNegotiationsParticipantTypeOk returns a tuple with the NegotiationsParticipantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsParticipantType

`func (o *DictionariesDictResponse) SetNegotiationsParticipantType(v []IncludesIdName)`

SetNegotiationsParticipantType sets NegotiationsParticipantType field to given value.


### GetNegotiationsState

`func (o *DictionariesDictResponse) GetNegotiationsState() []IncludesIdName`

GetNegotiationsState returns the NegotiationsState field if non-nil, zero value otherwise.

### GetNegotiationsStateOk

`func (o *DictionariesDictResponse) GetNegotiationsStateOk() (*[]IncludesIdName, bool)`

GetNegotiationsStateOk returns a tuple with the NegotiationsState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsState

`func (o *DictionariesDictResponse) SetNegotiationsState(v []IncludesIdName)`

SetNegotiationsState sets NegotiationsState field to given value.


### GetPhoneCallStatus

`func (o *DictionariesDictResponse) GetPhoneCallStatus() []IncludesIdName`

GetPhoneCallStatus returns the PhoneCallStatus field if non-nil, zero value otherwise.

### GetPhoneCallStatusOk

`func (o *DictionariesDictResponse) GetPhoneCallStatusOk() (*[]IncludesIdName, bool)`

GetPhoneCallStatusOk returns a tuple with the PhoneCallStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCallStatus

`func (o *DictionariesDictResponse) SetPhoneCallStatus(v []IncludesIdName)`

SetPhoneCallStatus sets PhoneCallStatus field to given value.


### GetPreferredContactType

`func (o *DictionariesDictResponse) GetPreferredContactType() []IncludesIdName`

GetPreferredContactType returns the PreferredContactType field if non-nil, zero value otherwise.

### GetPreferredContactTypeOk

`func (o *DictionariesDictResponse) GetPreferredContactTypeOk() (*[]IncludesIdName, bool)`

GetPreferredContactTypeOk returns a tuple with the PreferredContactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredContactType

`func (o *DictionariesDictResponse) SetPreferredContactType(v []IncludesIdName)`

SetPreferredContactType sets PreferredContactType field to given value.


### GetRelocationType

`func (o *DictionariesDictResponse) GetRelocationType() []IncludesIdName`

GetRelocationType returns the RelocationType field if non-nil, zero value otherwise.

### GetRelocationTypeOk

`func (o *DictionariesDictResponse) GetRelocationTypeOk() (*[]IncludesIdName, bool)`

GetRelocationTypeOk returns a tuple with the RelocationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocationType

`func (o *DictionariesDictResponse) SetRelocationType(v []IncludesIdName)`

SetRelocationType sets RelocationType field to given value.


### GetResumeAccessType

`func (o *DictionariesDictResponse) GetResumeAccessType() []IncludesIdName`

GetResumeAccessType returns the ResumeAccessType field if non-nil, zero value otherwise.

### GetResumeAccessTypeOk

`func (o *DictionariesDictResponse) GetResumeAccessTypeOk() (*[]IncludesIdName, bool)`

GetResumeAccessTypeOk returns a tuple with the ResumeAccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAccessType

`func (o *DictionariesDictResponse) SetResumeAccessType(v []IncludesIdName)`

SetResumeAccessType sets ResumeAccessType field to given value.


### GetResumeContactsSiteType

`func (o *DictionariesDictResponse) GetResumeContactsSiteType() []IncludesIdName`

GetResumeContactsSiteType returns the ResumeContactsSiteType field if non-nil, zero value otherwise.

### GetResumeContactsSiteTypeOk

`func (o *DictionariesDictResponse) GetResumeContactsSiteTypeOk() (*[]IncludesIdName, bool)`

GetResumeContactsSiteTypeOk returns a tuple with the ResumeContactsSiteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeContactsSiteType

`func (o *DictionariesDictResponse) SetResumeContactsSiteType(v []IncludesIdName)`

SetResumeContactsSiteType sets ResumeContactsSiteType field to given value.


### GetResumeHiddenFields

`func (o *DictionariesDictResponse) GetResumeHiddenFields() []IncludesIdName`

GetResumeHiddenFields returns the ResumeHiddenFields field if non-nil, zero value otherwise.

### GetResumeHiddenFieldsOk

`func (o *DictionariesDictResponse) GetResumeHiddenFieldsOk() (*[]IncludesIdName, bool)`

GetResumeHiddenFieldsOk returns a tuple with the ResumeHiddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeHiddenFields

`func (o *DictionariesDictResponse) SetResumeHiddenFields(v []IncludesIdName)`

SetResumeHiddenFields sets ResumeHiddenFields field to given value.


### GetResumeModerationNote

`func (o *DictionariesDictResponse) GetResumeModerationNote() []IncludesIdName`

GetResumeModerationNote returns the ResumeModerationNote field if non-nil, zero value otherwise.

### GetResumeModerationNoteOk

`func (o *DictionariesDictResponse) GetResumeModerationNoteOk() (*[]IncludesIdName, bool)`

GetResumeModerationNoteOk returns a tuple with the ResumeModerationNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeModerationNote

`func (o *DictionariesDictResponse) SetResumeModerationNote(v []IncludesIdName)`

SetResumeModerationNote sets ResumeModerationNote field to given value.


### GetResumeSearchExperiencePeriod

`func (o *DictionariesDictResponse) GetResumeSearchExperiencePeriod() []IncludesIdName`

GetResumeSearchExperiencePeriod returns the ResumeSearchExperiencePeriod field if non-nil, zero value otherwise.

### GetResumeSearchExperiencePeriodOk

`func (o *DictionariesDictResponse) GetResumeSearchExperiencePeriodOk() (*[]IncludesIdName, bool)`

GetResumeSearchExperiencePeriodOk returns a tuple with the ResumeSearchExperiencePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeSearchExperiencePeriod

`func (o *DictionariesDictResponse) SetResumeSearchExperiencePeriod(v []IncludesIdName)`

SetResumeSearchExperiencePeriod sets ResumeSearchExperiencePeriod field to given value.

### HasResumeSearchExperiencePeriod

`func (o *DictionariesDictResponse) HasResumeSearchExperiencePeriod() bool`

HasResumeSearchExperiencePeriod returns a boolean if a field has been set.

### GetResumeSearchFields

`func (o *DictionariesDictResponse) GetResumeSearchFields() []IncludesIdName`

GetResumeSearchFields returns the ResumeSearchFields field if non-nil, zero value otherwise.

### GetResumeSearchFieldsOk

`func (o *DictionariesDictResponse) GetResumeSearchFieldsOk() (*[]IncludesIdName, bool)`

GetResumeSearchFieldsOk returns a tuple with the ResumeSearchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeSearchFields

`func (o *DictionariesDictResponse) SetResumeSearchFields(v []IncludesIdName)`

SetResumeSearchFields sets ResumeSearchFields field to given value.

### HasResumeSearchFields

`func (o *DictionariesDictResponse) HasResumeSearchFields() bool`

HasResumeSearchFields returns a boolean if a field has been set.

### GetResumeSearchLabel

`func (o *DictionariesDictResponse) GetResumeSearchLabel() []IncludesIdName`

GetResumeSearchLabel returns the ResumeSearchLabel field if non-nil, zero value otherwise.

### GetResumeSearchLabelOk

`func (o *DictionariesDictResponse) GetResumeSearchLabelOk() (*[]IncludesIdName, bool)`

GetResumeSearchLabelOk returns a tuple with the ResumeSearchLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeSearchLabel

`func (o *DictionariesDictResponse) SetResumeSearchLabel(v []IncludesIdName)`

SetResumeSearchLabel sets ResumeSearchLabel field to given value.

### HasResumeSearchLabel

`func (o *DictionariesDictResponse) HasResumeSearchLabel() bool`

HasResumeSearchLabel returns a boolean if a field has been set.

### GetResumeSearchLogic

`func (o *DictionariesDictResponse) GetResumeSearchLogic() []IncludesIdName`

GetResumeSearchLogic returns the ResumeSearchLogic field if non-nil, zero value otherwise.

### GetResumeSearchLogicOk

`func (o *DictionariesDictResponse) GetResumeSearchLogicOk() (*[]IncludesIdName, bool)`

GetResumeSearchLogicOk returns a tuple with the ResumeSearchLogic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeSearchLogic

`func (o *DictionariesDictResponse) SetResumeSearchLogic(v []IncludesIdName)`

SetResumeSearchLogic sets ResumeSearchLogic field to given value.

### HasResumeSearchLogic

`func (o *DictionariesDictResponse) HasResumeSearchLogic() bool`

HasResumeSearchLogic returns a boolean if a field has been set.

### GetResumeSearchOrder

`func (o *DictionariesDictResponse) GetResumeSearchOrder() []IncludesIdName`

GetResumeSearchOrder returns the ResumeSearchOrder field if non-nil, zero value otherwise.

### GetResumeSearchOrderOk

`func (o *DictionariesDictResponse) GetResumeSearchOrderOk() (*[]IncludesIdName, bool)`

GetResumeSearchOrderOk returns a tuple with the ResumeSearchOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeSearchOrder

`func (o *DictionariesDictResponse) SetResumeSearchOrder(v []IncludesIdName)`

SetResumeSearchOrder sets ResumeSearchOrder field to given value.

### HasResumeSearchOrder

`func (o *DictionariesDictResponse) HasResumeSearchOrder() bool`

HasResumeSearchOrder returns a boolean if a field has been set.

### GetResumeSearchRelocation

`func (o *DictionariesDictResponse) GetResumeSearchRelocation() []IncludesIdName`

GetResumeSearchRelocation returns the ResumeSearchRelocation field if non-nil, zero value otherwise.

### GetResumeSearchRelocationOk

`func (o *DictionariesDictResponse) GetResumeSearchRelocationOk() (*[]IncludesIdName, bool)`

GetResumeSearchRelocationOk returns a tuple with the ResumeSearchRelocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeSearchRelocation

`func (o *DictionariesDictResponse) SetResumeSearchRelocation(v []IncludesIdName)`

SetResumeSearchRelocation sets ResumeSearchRelocation field to given value.

### HasResumeSearchRelocation

`func (o *DictionariesDictResponse) HasResumeSearchRelocation() bool`

HasResumeSearchRelocation returns a boolean if a field has been set.

### GetResumeStatus

`func (o *DictionariesDictResponse) GetResumeStatus() []IncludesIdName`

GetResumeStatus returns the ResumeStatus field if non-nil, zero value otherwise.

### GetResumeStatusOk

`func (o *DictionariesDictResponse) GetResumeStatusOk() (*[]IncludesIdName, bool)`

GetResumeStatusOk returns a tuple with the ResumeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeStatus

`func (o *DictionariesDictResponse) SetResumeStatus(v []IncludesIdName)`

SetResumeStatus sets ResumeStatus field to given value.


### GetSchedule

`func (o *DictionariesDictResponse) GetSchedule() []IncludesIdNameUid`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *DictionariesDictResponse) GetScheduleOk() (*[]IncludesIdNameUid, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *DictionariesDictResponse) SetSchedule(v []IncludesIdNameUid)`

SetSchedule sets Schedule field to given value.


### GetTravelTime

`func (o *DictionariesDictResponse) GetTravelTime() []IncludesIdName`

GetTravelTime returns the TravelTime field if non-nil, zero value otherwise.

### GetTravelTimeOk

`func (o *DictionariesDictResponse) GetTravelTimeOk() (*[]IncludesIdName, bool)`

GetTravelTimeOk returns a tuple with the TravelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTravelTime

`func (o *DictionariesDictResponse) SetTravelTime(v []IncludesIdName)`

SetTravelTime sets TravelTime field to given value.


### GetVacancyBillingType

`func (o *DictionariesDictResponse) GetVacancyBillingType() []IncludesIdName`

GetVacancyBillingType returns the VacancyBillingType field if non-nil, zero value otherwise.

### GetVacancyBillingTypeOk

`func (o *DictionariesDictResponse) GetVacancyBillingTypeOk() (*[]IncludesIdName, bool)`

GetVacancyBillingTypeOk returns a tuple with the VacancyBillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyBillingType

`func (o *DictionariesDictResponse) SetVacancyBillingType(v []IncludesIdName)`

SetVacancyBillingType sets VacancyBillingType field to given value.


### GetVacancyCluster

`func (o *DictionariesDictResponse) GetVacancyCluster() []IncludesIdName`

GetVacancyCluster returns the VacancyCluster field if non-nil, zero value otherwise.

### GetVacancyClusterOk

`func (o *DictionariesDictResponse) GetVacancyClusterOk() (*[]IncludesIdName, bool)`

GetVacancyClusterOk returns a tuple with the VacancyCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyCluster

`func (o *DictionariesDictResponse) SetVacancyCluster(v []IncludesIdName)`

SetVacancyCluster sets VacancyCluster field to given value.


### GetVacancyLabel

`func (o *DictionariesDictResponse) GetVacancyLabel() []IncludesIdName`

GetVacancyLabel returns the VacancyLabel field if non-nil, zero value otherwise.

### GetVacancyLabelOk

`func (o *DictionariesDictResponse) GetVacancyLabelOk() (*[]IncludesIdName, bool)`

GetVacancyLabelOk returns a tuple with the VacancyLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyLabel

`func (o *DictionariesDictResponse) SetVacancyLabel(v []IncludesIdName)`

SetVacancyLabel sets VacancyLabel field to given value.


### GetVacancyNotProlongedReason

`func (o *DictionariesDictResponse) GetVacancyNotProlongedReason() []IncludesIdName`

GetVacancyNotProlongedReason returns the VacancyNotProlongedReason field if non-nil, zero value otherwise.

### GetVacancyNotProlongedReasonOk

`func (o *DictionariesDictResponse) GetVacancyNotProlongedReasonOk() (*[]IncludesIdName, bool)`

GetVacancyNotProlongedReasonOk returns a tuple with the VacancyNotProlongedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyNotProlongedReason

`func (o *DictionariesDictResponse) SetVacancyNotProlongedReason(v []IncludesIdName)`

SetVacancyNotProlongedReason sets VacancyNotProlongedReason field to given value.


### GetVacancyRelation

`func (o *DictionariesDictResponse) GetVacancyRelation() []IncludesIdName`

GetVacancyRelation returns the VacancyRelation field if non-nil, zero value otherwise.

### GetVacancyRelationOk

`func (o *DictionariesDictResponse) GetVacancyRelationOk() (*[]IncludesIdName, bool)`

GetVacancyRelationOk returns a tuple with the VacancyRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyRelation

`func (o *DictionariesDictResponse) SetVacancyRelation(v []IncludesIdName)`

SetVacancyRelation sets VacancyRelation field to given value.


### GetVacancySearchFields

`func (o *DictionariesDictResponse) GetVacancySearchFields() []IncludesIdName`

GetVacancySearchFields returns the VacancySearchFields field if non-nil, zero value otherwise.

### GetVacancySearchFieldsOk

`func (o *DictionariesDictResponse) GetVacancySearchFieldsOk() (*[]IncludesIdName, bool)`

GetVacancySearchFieldsOk returns a tuple with the VacancySearchFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancySearchFields

`func (o *DictionariesDictResponse) SetVacancySearchFields(v []IncludesIdName)`

SetVacancySearchFields sets VacancySearchFields field to given value.


### GetVacancySearchOrder

`func (o *DictionariesDictResponse) GetVacancySearchOrder() []IncludesIdName`

GetVacancySearchOrder returns the VacancySearchOrder field if non-nil, zero value otherwise.

### GetVacancySearchOrderOk

`func (o *DictionariesDictResponse) GetVacancySearchOrderOk() (*[]IncludesIdName, bool)`

GetVacancySearchOrderOk returns a tuple with the VacancySearchOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancySearchOrder

`func (o *DictionariesDictResponse) SetVacancySearchOrder(v []IncludesIdName)`

SetVacancySearchOrder sets VacancySearchOrder field to given value.


### GetVacancyType

`func (o *DictionariesDictResponse) GetVacancyType() []IncludesIdName`

GetVacancyType returns the VacancyType field if non-nil, zero value otherwise.

### GetVacancyTypeOk

`func (o *DictionariesDictResponse) GetVacancyTypeOk() (*[]IncludesIdName, bool)`

GetVacancyTypeOk returns a tuple with the VacancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyType

`func (o *DictionariesDictResponse) SetVacancyType(v []IncludesIdName)`

SetVacancyType sets VacancyType field to given value.


### GetWorkingDays

`func (o *DictionariesDictResponse) GetWorkingDays() []IncludesIdName`

GetWorkingDays returns the WorkingDays field if non-nil, zero value otherwise.

### GetWorkingDaysOk

`func (o *DictionariesDictResponse) GetWorkingDaysOk() (*[]IncludesIdName, bool)`

GetWorkingDaysOk returns a tuple with the WorkingDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDays

`func (o *DictionariesDictResponse) SetWorkingDays(v []IncludesIdName)`

SetWorkingDays sets WorkingDays field to given value.


### GetWorkingTimeIntervals

`func (o *DictionariesDictResponse) GetWorkingTimeIntervals() []IncludesIdName`

GetWorkingTimeIntervals returns the WorkingTimeIntervals field if non-nil, zero value otherwise.

### GetWorkingTimeIntervalsOk

`func (o *DictionariesDictResponse) GetWorkingTimeIntervalsOk() (*[]IncludesIdName, bool)`

GetWorkingTimeIntervalsOk returns a tuple with the WorkingTimeIntervals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeIntervals

`func (o *DictionariesDictResponse) SetWorkingTimeIntervals(v []IncludesIdName)`

SetWorkingTimeIntervals sets WorkingTimeIntervals field to given value.


### GetWorkingTimeModes

`func (o *DictionariesDictResponse) GetWorkingTimeModes() []IncludesIdName`

GetWorkingTimeModes returns the WorkingTimeModes field if non-nil, zero value otherwise.

### GetWorkingTimeModesOk

`func (o *DictionariesDictResponse) GetWorkingTimeModesOk() (*[]IncludesIdName, bool)`

GetWorkingTimeModesOk returns a tuple with the WorkingTimeModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingTimeModes

`func (o *DictionariesDictResponse) SetWorkingTimeModes(v []IncludesIdName)`

SetWorkingTimeModes sets WorkingTimeModes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


