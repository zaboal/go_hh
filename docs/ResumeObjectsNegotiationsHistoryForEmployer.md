# ResumeObjectsNegotiationsHistoryForEmployer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | URL, на который необходимо сделать GET-запрос, чтобы получить [подробную историю откликов/приглашений](#tag/Otklikipriglasheniya-rabotodatelya/operation/get-resume-negotiations-history) по данному резюме | 
**Vacancies** | Pointer to [**[]ResumesResumeNegotiationsHistoryVacancy**](ResumesResumeNegotiationsHistoryVacancy.md) | Массив вакансий | [optional] 

## Methods

### NewResumeObjectsNegotiationsHistoryForEmployer

`func NewResumeObjectsNegotiationsHistoryForEmployer(url string, ) *ResumeObjectsNegotiationsHistoryForEmployer`

NewResumeObjectsNegotiationsHistoryForEmployer instantiates a new ResumeObjectsNegotiationsHistoryForEmployer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsNegotiationsHistoryForEmployerWithDefaults

`func NewResumeObjectsNegotiationsHistoryForEmployerWithDefaults() *ResumeObjectsNegotiationsHistoryForEmployer`

NewResumeObjectsNegotiationsHistoryForEmployerWithDefaults instantiates a new ResumeObjectsNegotiationsHistoryForEmployer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ResumeObjectsNegotiationsHistoryForEmployer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumeObjectsNegotiationsHistoryForEmployer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumeObjectsNegotiationsHistoryForEmployer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVacancies

`func (o *ResumeObjectsNegotiationsHistoryForEmployer) GetVacancies() []ResumesResumeNegotiationsHistoryVacancy`

GetVacancies returns the Vacancies field if non-nil, zero value otherwise.

### GetVacanciesOk

`func (o *ResumeObjectsNegotiationsHistoryForEmployer) GetVacanciesOk() (*[]ResumesResumeNegotiationsHistoryVacancy, bool)`

GetVacanciesOk returns a tuple with the Vacancies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancies

`func (o *ResumeObjectsNegotiationsHistoryForEmployer) SetVacancies(v []ResumesResumeNegotiationsHistoryVacancy)`

SetVacancies sets Vacancies field to given value.

### HasVacancies

`func (o *ResumeObjectsNegotiationsHistoryForEmployer) HasVacancies() bool`

HasVacancies returns a boolean if a field has been set.

### SetVacanciesNil

`func (o *ResumeObjectsNegotiationsHistoryForEmployer) SetVacanciesNil(b bool)`

 SetVacanciesNil sets the value for Vacancies to be an explicit nil

### UnsetVacancies
`func (o *ResumeObjectsNegotiationsHistoryForEmployer) UnsetVacancies()`

UnsetVacancies ensures that no value is present for Vacancies, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


