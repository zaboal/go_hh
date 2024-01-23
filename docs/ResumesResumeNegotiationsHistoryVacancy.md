# ResumesResumeNegotiationsHistoryVacancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | Признак того, что вакансия находится в архиве | 
**CanEdit** | **bool** | Признак того, что менеджер может редактировать данные вакансии, а также работать с информацией об откликах/приглашениях по этой вакансии | 
**Id** | **string** | Уникальный идентификатор вакансии | 
**Items** | [**[]ResumesResumeNegotiationsHistoryVacancyItem**](ResumesResumeNegotiationsHistoryVacancyItem.md) | Список последних изменений состояний откликов/приглашений по указанному резюме и данной вакансии | 
**MessagesUrl** | **string** | URL, на который необходимо делать GET запрос для получения [списка сообщений в отклике/приглашении](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#get-messages). Если &#x60;can_edit&#x60; равно &#x60;false&#x60;, значение поля должно игнорироваться | 
**Name** | **string** | Название вакансии | 
**NegotiationsUrl** | **string** | URL, на который необходимо делать GET-запрос для получения [данных об отклике/приглашении](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#get-negotiation). Если &#x60;can_edit&#x60; равно &#x60;false&#x60;, значение поля должно игнорироваться | 
**Url** | **string** | URL, на который необходимо делать GET-запрос для [получения данных о вакансии](#tag/Vakansii/operation/get-vacancy) | 

## Methods

### NewResumesResumeNegotiationsHistoryVacancy

`func NewResumesResumeNegotiationsHistoryVacancy(archived bool, canEdit bool, id string, items []ResumesResumeNegotiationsHistoryVacancyItem, messagesUrl string, name string, negotiationsUrl string, url string, ) *ResumesResumeNegotiationsHistoryVacancy`

NewResumesResumeNegotiationsHistoryVacancy instantiates a new ResumesResumeNegotiationsHistoryVacancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesResumeNegotiationsHistoryVacancyWithDefaults

`func NewResumesResumeNegotiationsHistoryVacancyWithDefaults() *ResumesResumeNegotiationsHistoryVacancy`

NewResumesResumeNegotiationsHistoryVacancyWithDefaults instantiates a new ResumesResumeNegotiationsHistoryVacancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchived

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *ResumesResumeNegotiationsHistoryVacancy) SetArchived(v bool)`

SetArchived sets Archived field to given value.


### GetCanEdit

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *ResumesResumeNegotiationsHistoryVacancy) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.


### GetId

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResumesResumeNegotiationsHistoryVacancy) SetId(v string)`

SetId sets Id field to given value.


### GetItems

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetItems() []ResumesResumeNegotiationsHistoryVacancyItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetItemsOk() (*[]ResumesResumeNegotiationsHistoryVacancyItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResumesResumeNegotiationsHistoryVacancy) SetItems(v []ResumesResumeNegotiationsHistoryVacancyItem)`

SetItems sets Items field to given value.


### GetMessagesUrl

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetMessagesUrl() string`

GetMessagesUrl returns the MessagesUrl field if non-nil, zero value otherwise.

### GetMessagesUrlOk

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetMessagesUrlOk() (*string, bool)`

GetMessagesUrlOk returns a tuple with the MessagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesUrl

`func (o *ResumesResumeNegotiationsHistoryVacancy) SetMessagesUrl(v string)`

SetMessagesUrl sets MessagesUrl field to given value.


### GetName

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResumesResumeNegotiationsHistoryVacancy) SetName(v string)`

SetName sets Name field to given value.


### GetNegotiationsUrl

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetNegotiationsUrl() string`

GetNegotiationsUrl returns the NegotiationsUrl field if non-nil, zero value otherwise.

### GetNegotiationsUrlOk

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetNegotiationsUrlOk() (*string, bool)`

GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsUrl

`func (o *ResumesResumeNegotiationsHistoryVacancy) SetNegotiationsUrl(v string)`

SetNegotiationsUrl sets NegotiationsUrl field to given value.


### GetUrl

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ResumesResumeNegotiationsHistoryVacancy) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ResumesResumeNegotiationsHistoryVacancy) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


