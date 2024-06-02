# VacancyDraftVacancyDraftBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoPublication** | Pointer to [**NullableVacancyDraftAutoPublicationState**](VacancyDraftAutoPublicationState.md) |  | [optional] 
**CompletedFieldsPercentage** | **float32** | Процент заполнения черновика | 
**DraftId** | **string** | Идентификатор черновика | 
**InsufficientPublications** | Pointer to [**[]VacancyDraftPublications**](VacancyDraftPublications.md) | Массив объектов с информацией о том, каких публикаций не хватает на счету для публикации вакансии из данного черновика | [optional] 
**InsufficientQuotas** | Pointer to [**[]VacancyDraftPublications**](VacancyDraftPublications.md) | Массив объектов с информацией о том, какие квоты превышены | [optional] 
**LastChangeTime** | Pointer to **NullableString** | Время изменения черновика (в формате [ISO 8601](https://ru.wikipedia.org/wiki/ISO_8601) с точностью до секунды &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60;) | [optional] 
**PublicationReady** | **bool** | Готовность черновика к публикации | 
**RequiredPublications** | Pointer to [**[]VacancyDraftPublications**](VacancyDraftPublications.md) | Массив объектов с информацией о необходимых публикациях на счету | [optional] 
**ScheduledAt** | **string** | Время запланированной публикации вакансии (в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 

## Methods

### NewVacancyDraftVacancyDraftBase

`func NewVacancyDraftVacancyDraftBase(completedFieldsPercentage float32, draftId string, publicationReady bool, scheduledAt string, ) *VacancyDraftVacancyDraftBase`

NewVacancyDraftVacancyDraftBase instantiates a new VacancyDraftVacancyDraftBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftVacancyDraftBaseWithDefaults

`func NewVacancyDraftVacancyDraftBaseWithDefaults() *VacancyDraftVacancyDraftBase`

NewVacancyDraftVacancyDraftBaseWithDefaults instantiates a new VacancyDraftVacancyDraftBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPublication

`func (o *VacancyDraftVacancyDraftBase) GetAutoPublication() VacancyDraftAutoPublicationState`

GetAutoPublication returns the AutoPublication field if non-nil, zero value otherwise.

### GetAutoPublicationOk

`func (o *VacancyDraftVacancyDraftBase) GetAutoPublicationOk() (*VacancyDraftAutoPublicationState, bool)`

GetAutoPublicationOk returns a tuple with the AutoPublication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPublication

`func (o *VacancyDraftVacancyDraftBase) SetAutoPublication(v VacancyDraftAutoPublicationState)`

SetAutoPublication sets AutoPublication field to given value.

### HasAutoPublication

`func (o *VacancyDraftVacancyDraftBase) HasAutoPublication() bool`

HasAutoPublication returns a boolean if a field has been set.

### SetAutoPublicationNil

`func (o *VacancyDraftVacancyDraftBase) SetAutoPublicationNil(b bool)`

 SetAutoPublicationNil sets the value for AutoPublication to be an explicit nil

### UnsetAutoPublication
`func (o *VacancyDraftVacancyDraftBase) UnsetAutoPublication()`

UnsetAutoPublication ensures that no value is present for AutoPublication, not even an explicit nil
### GetCompletedFieldsPercentage

`func (o *VacancyDraftVacancyDraftBase) GetCompletedFieldsPercentage() float32`

GetCompletedFieldsPercentage returns the CompletedFieldsPercentage field if non-nil, zero value otherwise.

### GetCompletedFieldsPercentageOk

`func (o *VacancyDraftVacancyDraftBase) GetCompletedFieldsPercentageOk() (*float32, bool)`

GetCompletedFieldsPercentageOk returns a tuple with the CompletedFieldsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedFieldsPercentage

`func (o *VacancyDraftVacancyDraftBase) SetCompletedFieldsPercentage(v float32)`

SetCompletedFieldsPercentage sets CompletedFieldsPercentage field to given value.


### GetDraftId

`func (o *VacancyDraftVacancyDraftBase) GetDraftId() string`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *VacancyDraftVacancyDraftBase) GetDraftIdOk() (*string, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *VacancyDraftVacancyDraftBase) SetDraftId(v string)`

SetDraftId sets DraftId field to given value.


### GetInsufficientPublications

`func (o *VacancyDraftVacancyDraftBase) GetInsufficientPublications() []VacancyDraftPublications`

GetInsufficientPublications returns the InsufficientPublications field if non-nil, zero value otherwise.

### GetInsufficientPublicationsOk

`func (o *VacancyDraftVacancyDraftBase) GetInsufficientPublicationsOk() (*[]VacancyDraftPublications, bool)`

GetInsufficientPublicationsOk returns a tuple with the InsufficientPublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientPublications

`func (o *VacancyDraftVacancyDraftBase) SetInsufficientPublications(v []VacancyDraftPublications)`

SetInsufficientPublications sets InsufficientPublications field to given value.

### HasInsufficientPublications

`func (o *VacancyDraftVacancyDraftBase) HasInsufficientPublications() bool`

HasInsufficientPublications returns a boolean if a field has been set.

### SetInsufficientPublicationsNil

`func (o *VacancyDraftVacancyDraftBase) SetInsufficientPublicationsNil(b bool)`

 SetInsufficientPublicationsNil sets the value for InsufficientPublications to be an explicit nil

### UnsetInsufficientPublications
`func (o *VacancyDraftVacancyDraftBase) UnsetInsufficientPublications()`

UnsetInsufficientPublications ensures that no value is present for InsufficientPublications, not even an explicit nil
### GetInsufficientQuotas

`func (o *VacancyDraftVacancyDraftBase) GetInsufficientQuotas() []VacancyDraftPublications`

GetInsufficientQuotas returns the InsufficientQuotas field if non-nil, zero value otherwise.

### GetInsufficientQuotasOk

`func (o *VacancyDraftVacancyDraftBase) GetInsufficientQuotasOk() (*[]VacancyDraftPublications, bool)`

GetInsufficientQuotasOk returns a tuple with the InsufficientQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientQuotas

`func (o *VacancyDraftVacancyDraftBase) SetInsufficientQuotas(v []VacancyDraftPublications)`

SetInsufficientQuotas sets InsufficientQuotas field to given value.

### HasInsufficientQuotas

`func (o *VacancyDraftVacancyDraftBase) HasInsufficientQuotas() bool`

HasInsufficientQuotas returns a boolean if a field has been set.

### SetInsufficientQuotasNil

`func (o *VacancyDraftVacancyDraftBase) SetInsufficientQuotasNil(b bool)`

 SetInsufficientQuotasNil sets the value for InsufficientQuotas to be an explicit nil

### UnsetInsufficientQuotas
`func (o *VacancyDraftVacancyDraftBase) UnsetInsufficientQuotas()`

UnsetInsufficientQuotas ensures that no value is present for InsufficientQuotas, not even an explicit nil
### GetLastChangeTime

`func (o *VacancyDraftVacancyDraftBase) GetLastChangeTime() string`

GetLastChangeTime returns the LastChangeTime field if non-nil, zero value otherwise.

### GetLastChangeTimeOk

`func (o *VacancyDraftVacancyDraftBase) GetLastChangeTimeOk() (*string, bool)`

GetLastChangeTimeOk returns a tuple with the LastChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeTime

`func (o *VacancyDraftVacancyDraftBase) SetLastChangeTime(v string)`

SetLastChangeTime sets LastChangeTime field to given value.

### HasLastChangeTime

`func (o *VacancyDraftVacancyDraftBase) HasLastChangeTime() bool`

HasLastChangeTime returns a boolean if a field has been set.

### SetLastChangeTimeNil

`func (o *VacancyDraftVacancyDraftBase) SetLastChangeTimeNil(b bool)`

 SetLastChangeTimeNil sets the value for LastChangeTime to be an explicit nil

### UnsetLastChangeTime
`func (o *VacancyDraftVacancyDraftBase) UnsetLastChangeTime()`

UnsetLastChangeTime ensures that no value is present for LastChangeTime, not even an explicit nil
### GetPublicationReady

`func (o *VacancyDraftVacancyDraftBase) GetPublicationReady() bool`

GetPublicationReady returns the PublicationReady field if non-nil, zero value otherwise.

### GetPublicationReadyOk

`func (o *VacancyDraftVacancyDraftBase) GetPublicationReadyOk() (*bool, bool)`

GetPublicationReadyOk returns a tuple with the PublicationReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationReady

`func (o *VacancyDraftVacancyDraftBase) SetPublicationReady(v bool)`

SetPublicationReady sets PublicationReady field to given value.


### GetRequiredPublications

`func (o *VacancyDraftVacancyDraftBase) GetRequiredPublications() []VacancyDraftPublications`

GetRequiredPublications returns the RequiredPublications field if non-nil, zero value otherwise.

### GetRequiredPublicationsOk

`func (o *VacancyDraftVacancyDraftBase) GetRequiredPublicationsOk() (*[]VacancyDraftPublications, bool)`

GetRequiredPublicationsOk returns a tuple with the RequiredPublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPublications

`func (o *VacancyDraftVacancyDraftBase) SetRequiredPublications(v []VacancyDraftPublications)`

SetRequiredPublications sets RequiredPublications field to given value.

### HasRequiredPublications

`func (o *VacancyDraftVacancyDraftBase) HasRequiredPublications() bool`

HasRequiredPublications returns a boolean if a field has been set.

### SetRequiredPublicationsNil

`func (o *VacancyDraftVacancyDraftBase) SetRequiredPublicationsNil(b bool)`

 SetRequiredPublicationsNil sets the value for RequiredPublications to be an explicit nil

### UnsetRequiredPublications
`func (o *VacancyDraftVacancyDraftBase) UnsetRequiredPublications()`

UnsetRequiredPublications ensures that no value is present for RequiredPublications, not even an explicit nil
### GetScheduledAt

`func (o *VacancyDraftVacancyDraftBase) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *VacancyDraftVacancyDraftBase) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *VacancyDraftVacancyDraftBase) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


