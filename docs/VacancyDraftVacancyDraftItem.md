# VacancyDraftVacancyDraftItem

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
**ScheduledAt** | **NullableString** | Время запланированной публикации вакансии (в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**Areas** | [**[]VacancyAreaOutput**](VacancyAreaOutput.md) | Коды и названия регионов (фед. округа, субъекты федерации, города) | 
**AssignedManager** | Pointer to [**NullableVacancyDraftAssignedManager**](VacancyDraftAssignedManager.md) |  | [optional] 
**BillingType** | [**NullableVacancyBillingTypeOutput**](VacancyBillingTypeOutput.md) |  | 
**Name** | Pointer to **string** | Название вакансии | [optional] 
**PublicationType** | **string** | Тип публикации (справочник [vacancy_billing_type](#tag/Obshie-spravochniki/operation/get-dictionaries)) | 
**Url** | **string** | Url для запроса полной информации черновика | 
**VacancyType** | **NullableString** | Тип вакансии (справочник [vacancy_type](#tag/Obshie-spravochniki/operation/get-dictionaries)) | 

## Methods

### NewVacancyDraftVacancyDraftItem

`func NewVacancyDraftVacancyDraftItem(completedFieldsPercentage float32, draftId string, publicationReady bool, scheduledAt NullableString, areas []VacancyAreaOutput, billingType NullableVacancyBillingTypeOutput, publicationType string, url string, vacancyType NullableString, ) *VacancyDraftVacancyDraftItem`

NewVacancyDraftVacancyDraftItem instantiates a new VacancyDraftVacancyDraftItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftVacancyDraftItemWithDefaults

`func NewVacancyDraftVacancyDraftItemWithDefaults() *VacancyDraftVacancyDraftItem`

NewVacancyDraftVacancyDraftItemWithDefaults instantiates a new VacancyDraftVacancyDraftItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoPublication

`func (o *VacancyDraftVacancyDraftItem) GetAutoPublication() VacancyDraftAutoPublicationState`

GetAutoPublication returns the AutoPublication field if non-nil, zero value otherwise.

### GetAutoPublicationOk

`func (o *VacancyDraftVacancyDraftItem) GetAutoPublicationOk() (*VacancyDraftAutoPublicationState, bool)`

GetAutoPublicationOk returns a tuple with the AutoPublication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPublication

`func (o *VacancyDraftVacancyDraftItem) SetAutoPublication(v VacancyDraftAutoPublicationState)`

SetAutoPublication sets AutoPublication field to given value.

### HasAutoPublication

`func (o *VacancyDraftVacancyDraftItem) HasAutoPublication() bool`

HasAutoPublication returns a boolean if a field has been set.

### SetAutoPublicationNil

`func (o *VacancyDraftVacancyDraftItem) SetAutoPublicationNil(b bool)`

 SetAutoPublicationNil sets the value for AutoPublication to be an explicit nil

### UnsetAutoPublication
`func (o *VacancyDraftVacancyDraftItem) UnsetAutoPublication()`

UnsetAutoPublication ensures that no value is present for AutoPublication, not even an explicit nil
### GetCompletedFieldsPercentage

`func (o *VacancyDraftVacancyDraftItem) GetCompletedFieldsPercentage() float32`

GetCompletedFieldsPercentage returns the CompletedFieldsPercentage field if non-nil, zero value otherwise.

### GetCompletedFieldsPercentageOk

`func (o *VacancyDraftVacancyDraftItem) GetCompletedFieldsPercentageOk() (*float32, bool)`

GetCompletedFieldsPercentageOk returns a tuple with the CompletedFieldsPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedFieldsPercentage

`func (o *VacancyDraftVacancyDraftItem) SetCompletedFieldsPercentage(v float32)`

SetCompletedFieldsPercentage sets CompletedFieldsPercentage field to given value.


### GetDraftId

`func (o *VacancyDraftVacancyDraftItem) GetDraftId() string`

GetDraftId returns the DraftId field if non-nil, zero value otherwise.

### GetDraftIdOk

`func (o *VacancyDraftVacancyDraftItem) GetDraftIdOk() (*string, bool)`

GetDraftIdOk returns a tuple with the DraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftId

`func (o *VacancyDraftVacancyDraftItem) SetDraftId(v string)`

SetDraftId sets DraftId field to given value.


### GetInsufficientPublications

`func (o *VacancyDraftVacancyDraftItem) GetInsufficientPublications() []VacancyDraftPublications`

GetInsufficientPublications returns the InsufficientPublications field if non-nil, zero value otherwise.

### GetInsufficientPublicationsOk

`func (o *VacancyDraftVacancyDraftItem) GetInsufficientPublicationsOk() (*[]VacancyDraftPublications, bool)`

GetInsufficientPublicationsOk returns a tuple with the InsufficientPublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientPublications

`func (o *VacancyDraftVacancyDraftItem) SetInsufficientPublications(v []VacancyDraftPublications)`

SetInsufficientPublications sets InsufficientPublications field to given value.

### HasInsufficientPublications

`func (o *VacancyDraftVacancyDraftItem) HasInsufficientPublications() bool`

HasInsufficientPublications returns a boolean if a field has been set.

### SetInsufficientPublicationsNil

`func (o *VacancyDraftVacancyDraftItem) SetInsufficientPublicationsNil(b bool)`

 SetInsufficientPublicationsNil sets the value for InsufficientPublications to be an explicit nil

### UnsetInsufficientPublications
`func (o *VacancyDraftVacancyDraftItem) UnsetInsufficientPublications()`

UnsetInsufficientPublications ensures that no value is present for InsufficientPublications, not even an explicit nil
### GetInsufficientQuotas

`func (o *VacancyDraftVacancyDraftItem) GetInsufficientQuotas() []VacancyDraftPublications`

GetInsufficientQuotas returns the InsufficientQuotas field if non-nil, zero value otherwise.

### GetInsufficientQuotasOk

`func (o *VacancyDraftVacancyDraftItem) GetInsufficientQuotasOk() (*[]VacancyDraftPublications, bool)`

GetInsufficientQuotasOk returns a tuple with the InsufficientQuotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsufficientQuotas

`func (o *VacancyDraftVacancyDraftItem) SetInsufficientQuotas(v []VacancyDraftPublications)`

SetInsufficientQuotas sets InsufficientQuotas field to given value.

### HasInsufficientQuotas

`func (o *VacancyDraftVacancyDraftItem) HasInsufficientQuotas() bool`

HasInsufficientQuotas returns a boolean if a field has been set.

### SetInsufficientQuotasNil

`func (o *VacancyDraftVacancyDraftItem) SetInsufficientQuotasNil(b bool)`

 SetInsufficientQuotasNil sets the value for InsufficientQuotas to be an explicit nil

### UnsetInsufficientQuotas
`func (o *VacancyDraftVacancyDraftItem) UnsetInsufficientQuotas()`

UnsetInsufficientQuotas ensures that no value is present for InsufficientQuotas, not even an explicit nil
### GetLastChangeTime

`func (o *VacancyDraftVacancyDraftItem) GetLastChangeTime() string`

GetLastChangeTime returns the LastChangeTime field if non-nil, zero value otherwise.

### GetLastChangeTimeOk

`func (o *VacancyDraftVacancyDraftItem) GetLastChangeTimeOk() (*string, bool)`

GetLastChangeTimeOk returns a tuple with the LastChangeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChangeTime

`func (o *VacancyDraftVacancyDraftItem) SetLastChangeTime(v string)`

SetLastChangeTime sets LastChangeTime field to given value.

### HasLastChangeTime

`func (o *VacancyDraftVacancyDraftItem) HasLastChangeTime() bool`

HasLastChangeTime returns a boolean if a field has been set.

### SetLastChangeTimeNil

`func (o *VacancyDraftVacancyDraftItem) SetLastChangeTimeNil(b bool)`

 SetLastChangeTimeNil sets the value for LastChangeTime to be an explicit nil

### UnsetLastChangeTime
`func (o *VacancyDraftVacancyDraftItem) UnsetLastChangeTime()`

UnsetLastChangeTime ensures that no value is present for LastChangeTime, not even an explicit nil
### GetPublicationReady

`func (o *VacancyDraftVacancyDraftItem) GetPublicationReady() bool`

GetPublicationReady returns the PublicationReady field if non-nil, zero value otherwise.

### GetPublicationReadyOk

`func (o *VacancyDraftVacancyDraftItem) GetPublicationReadyOk() (*bool, bool)`

GetPublicationReadyOk returns a tuple with the PublicationReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationReady

`func (o *VacancyDraftVacancyDraftItem) SetPublicationReady(v bool)`

SetPublicationReady sets PublicationReady field to given value.


### GetRequiredPublications

`func (o *VacancyDraftVacancyDraftItem) GetRequiredPublications() []VacancyDraftPublications`

GetRequiredPublications returns the RequiredPublications field if non-nil, zero value otherwise.

### GetRequiredPublicationsOk

`func (o *VacancyDraftVacancyDraftItem) GetRequiredPublicationsOk() (*[]VacancyDraftPublications, bool)`

GetRequiredPublicationsOk returns a tuple with the RequiredPublications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredPublications

`func (o *VacancyDraftVacancyDraftItem) SetRequiredPublications(v []VacancyDraftPublications)`

SetRequiredPublications sets RequiredPublications field to given value.

### HasRequiredPublications

`func (o *VacancyDraftVacancyDraftItem) HasRequiredPublications() bool`

HasRequiredPublications returns a boolean if a field has been set.

### SetRequiredPublicationsNil

`func (o *VacancyDraftVacancyDraftItem) SetRequiredPublicationsNil(b bool)`

 SetRequiredPublicationsNil sets the value for RequiredPublications to be an explicit nil

### UnsetRequiredPublications
`func (o *VacancyDraftVacancyDraftItem) UnsetRequiredPublications()`

UnsetRequiredPublications ensures that no value is present for RequiredPublications, not even an explicit nil
### GetScheduledAt

`func (o *VacancyDraftVacancyDraftItem) GetScheduledAt() string`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *VacancyDraftVacancyDraftItem) GetScheduledAtOk() (*string, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *VacancyDraftVacancyDraftItem) SetScheduledAt(v string)`

SetScheduledAt sets ScheduledAt field to given value.


### SetScheduledAtNil

`func (o *VacancyDraftVacancyDraftItem) SetScheduledAtNil(b bool)`

 SetScheduledAtNil sets the value for ScheduledAt to be an explicit nil

### UnsetScheduledAt
`func (o *VacancyDraftVacancyDraftItem) UnsetScheduledAt()`

UnsetScheduledAt ensures that no value is present for ScheduledAt, not even an explicit nil
### GetAreas

`func (o *VacancyDraftVacancyDraftItem) GetAreas() []VacancyAreaOutput`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *VacancyDraftVacancyDraftItem) GetAreasOk() (*[]VacancyAreaOutput, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *VacancyDraftVacancyDraftItem) SetAreas(v []VacancyAreaOutput)`

SetAreas sets Areas field to given value.


### GetAssignedManager

`func (o *VacancyDraftVacancyDraftItem) GetAssignedManager() VacancyDraftAssignedManager`

GetAssignedManager returns the AssignedManager field if non-nil, zero value otherwise.

### GetAssignedManagerOk

`func (o *VacancyDraftVacancyDraftItem) GetAssignedManagerOk() (*VacancyDraftAssignedManager, bool)`

GetAssignedManagerOk returns a tuple with the AssignedManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedManager

`func (o *VacancyDraftVacancyDraftItem) SetAssignedManager(v VacancyDraftAssignedManager)`

SetAssignedManager sets AssignedManager field to given value.

### HasAssignedManager

`func (o *VacancyDraftVacancyDraftItem) HasAssignedManager() bool`

HasAssignedManager returns a boolean if a field has been set.

### SetAssignedManagerNil

`func (o *VacancyDraftVacancyDraftItem) SetAssignedManagerNil(b bool)`

 SetAssignedManagerNil sets the value for AssignedManager to be an explicit nil

### UnsetAssignedManager
`func (o *VacancyDraftVacancyDraftItem) UnsetAssignedManager()`

UnsetAssignedManager ensures that no value is present for AssignedManager, not even an explicit nil
### GetBillingType

`func (o *VacancyDraftVacancyDraftItem) GetBillingType() VacancyBillingTypeOutput`

GetBillingType returns the BillingType field if non-nil, zero value otherwise.

### GetBillingTypeOk

`func (o *VacancyDraftVacancyDraftItem) GetBillingTypeOk() (*VacancyBillingTypeOutput, bool)`

GetBillingTypeOk returns a tuple with the BillingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingType

`func (o *VacancyDraftVacancyDraftItem) SetBillingType(v VacancyBillingTypeOutput)`

SetBillingType sets BillingType field to given value.


### SetBillingTypeNil

`func (o *VacancyDraftVacancyDraftItem) SetBillingTypeNil(b bool)`

 SetBillingTypeNil sets the value for BillingType to be an explicit nil

### UnsetBillingType
`func (o *VacancyDraftVacancyDraftItem) UnsetBillingType()`

UnsetBillingType ensures that no value is present for BillingType, not even an explicit nil
### GetName

`func (o *VacancyDraftVacancyDraftItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyDraftVacancyDraftItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyDraftVacancyDraftItem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacancyDraftVacancyDraftItem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPublicationType

`func (o *VacancyDraftVacancyDraftItem) GetPublicationType() string`

GetPublicationType returns the PublicationType field if non-nil, zero value otherwise.

### GetPublicationTypeOk

`func (o *VacancyDraftVacancyDraftItem) GetPublicationTypeOk() (*string, bool)`

GetPublicationTypeOk returns a tuple with the PublicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicationType

`func (o *VacancyDraftVacancyDraftItem) SetPublicationType(v string)`

SetPublicationType sets PublicationType field to given value.


### GetUrl

`func (o *VacancyDraftVacancyDraftItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacancyDraftVacancyDraftItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacancyDraftVacancyDraftItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVacancyType

`func (o *VacancyDraftVacancyDraftItem) GetVacancyType() string`

GetVacancyType returns the VacancyType field if non-nil, zero value otherwise.

### GetVacancyTypeOk

`func (o *VacancyDraftVacancyDraftItem) GetVacancyTypeOk() (*string, bool)`

GetVacancyTypeOk returns a tuple with the VacancyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyType

`func (o *VacancyDraftVacancyDraftItem) SetVacancyType(v string)`

SetVacancyType sets VacancyType field to given value.


### SetVacancyTypeNil

`func (o *VacancyDraftVacancyDraftItem) SetVacancyTypeNil(b bool)`

 SetVacancyTypeNil sets the value for VacancyType to be an explicit nil

### UnsetVacancyType
`func (o *VacancyDraftVacancyDraftItem) UnsetVacancyType()`

UnsetVacancyType ensures that no value is present for VacancyType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


