# NegotiationsObjectsTopicItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | Pointer to [**NegotiationsObjectsCounters**](NegotiationsObjectsCounters.md) |  | [optional] 
**CreatedAt** | **string** | Дата и время создания отклика | 
**DeclineAllowed** | **bool** | Можно ли [скрыть отклик](#tag/Perepiska-(otklikipriglasheniya)-dlya-soiskatelya/operation/hide-active-response) вместе с сообщением работодателю об отказе | 
**HasUpdates** | **bool** | Есть ли непросмотренные сообщения в отклике. Флаг сбрасывается при различных действиях с откликом, например, [просмотре списка сообщений](https://github.com/hhru/api/blob/master/docs/negotiations.md#get_messages) | 
**Hidden** | **bool** | Скрыт ли текущий отклик от соискателя | 
**Id** | **string** | Идентификатор отклика | 
**JobSearchStatus** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**MessagesUrl** | Pointer to **string** | URL, на который необходимо делать GET запрос для получения [списка сообщений в отклике/приглашении](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#get-messages). Если &#x60;can_edit&#x60; равно &#x60;false&#x60;, значение поля должно игнорироваться | [optional] 
**MessagingStatus** | **string** | Текущий статус переписки.  Возможные значения приведены в поле &#x60;messaging_status&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**PhoneCalls** | Pointer to [**NullableNegotiationsPhoneCalls**](NegotiationsPhoneCalls.md) |  | [optional] 
**ProfessionalRoles** | Pointer to [**IncludesIdName**](IncludesIdName.md) |  | [optional] 
**Source** | **string** | Источник отклика | 
**State** | [**IncludesIdName**](IncludesIdName.md) | Текущее состояние отклика.  Возможные значения приведены в поле &#x60;negotiations_state&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**UpdatedAt** | **string** | Дата и время последнего обновления отклика | 
**Url** | **string** | Ссылка на полную версию отклика | 
**Vacancy** | Pointer to [**NullableVacanciesNegotiationsVacancyShort**](VacanciesNegotiationsVacancyShort.md) |  | [optional] 
**ViewedByOpponent** | **bool** | Был ли отклик просмотрен работодателем | 

## Methods

### NewNegotiationsObjectsTopicItem

`func NewNegotiationsObjectsTopicItem(createdAt string, declineAllowed bool, hasUpdates bool, hidden bool, id string, messagingStatus string, source string, state IncludesIdName, updatedAt string, url string, viewedByOpponent bool, ) *NegotiationsObjectsTopicItem`

NewNegotiationsObjectsTopicItem instantiates a new NegotiationsObjectsTopicItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsObjectsTopicItemWithDefaults

`func NewNegotiationsObjectsTopicItemWithDefaults() *NegotiationsObjectsTopicItem`

NewNegotiationsObjectsTopicItemWithDefaults instantiates a new NegotiationsObjectsTopicItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *NegotiationsObjectsTopicItem) GetCounters() NegotiationsObjectsCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *NegotiationsObjectsTopicItem) GetCountersOk() (*NegotiationsObjectsCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *NegotiationsObjectsTopicItem) SetCounters(v NegotiationsObjectsCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *NegotiationsObjectsTopicItem) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NegotiationsObjectsTopicItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsObjectsTopicItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsObjectsTopicItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeclineAllowed

`func (o *NegotiationsObjectsTopicItem) GetDeclineAllowed() bool`

GetDeclineAllowed returns the DeclineAllowed field if non-nil, zero value otherwise.

### GetDeclineAllowedOk

`func (o *NegotiationsObjectsTopicItem) GetDeclineAllowedOk() (*bool, bool)`

GetDeclineAllowedOk returns a tuple with the DeclineAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineAllowed

`func (o *NegotiationsObjectsTopicItem) SetDeclineAllowed(v bool)`

SetDeclineAllowed sets DeclineAllowed field to given value.


### GetHasUpdates

`func (o *NegotiationsObjectsTopicItem) GetHasUpdates() bool`

GetHasUpdates returns the HasUpdates field if non-nil, zero value otherwise.

### GetHasUpdatesOk

`func (o *NegotiationsObjectsTopicItem) GetHasUpdatesOk() (*bool, bool)`

GetHasUpdatesOk returns a tuple with the HasUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdates

`func (o *NegotiationsObjectsTopicItem) SetHasUpdates(v bool)`

SetHasUpdates sets HasUpdates field to given value.


### GetHidden

`func (o *NegotiationsObjectsTopicItem) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *NegotiationsObjectsTopicItem) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *NegotiationsObjectsTopicItem) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetId

`func (o *NegotiationsObjectsTopicItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsObjectsTopicItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsObjectsTopicItem) SetId(v string)`

SetId sets Id field to given value.


### GetJobSearchStatus

`func (o *NegotiationsObjectsTopicItem) GetJobSearchStatus() IncludesIdName`

GetJobSearchStatus returns the JobSearchStatus field if non-nil, zero value otherwise.

### GetJobSearchStatusOk

`func (o *NegotiationsObjectsTopicItem) GetJobSearchStatusOk() (*IncludesIdName, bool)`

GetJobSearchStatusOk returns a tuple with the JobSearchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatus

`func (o *NegotiationsObjectsTopicItem) SetJobSearchStatus(v IncludesIdName)`

SetJobSearchStatus sets JobSearchStatus field to given value.

### HasJobSearchStatus

`func (o *NegotiationsObjectsTopicItem) HasJobSearchStatus() bool`

HasJobSearchStatus returns a boolean if a field has been set.

### SetJobSearchStatusNil

`func (o *NegotiationsObjectsTopicItem) SetJobSearchStatusNil(b bool)`

 SetJobSearchStatusNil sets the value for JobSearchStatus to be an explicit nil

### UnsetJobSearchStatus
`func (o *NegotiationsObjectsTopicItem) UnsetJobSearchStatus()`

UnsetJobSearchStatus ensures that no value is present for JobSearchStatus, not even an explicit nil
### GetMessagesUrl

`func (o *NegotiationsObjectsTopicItem) GetMessagesUrl() string`

GetMessagesUrl returns the MessagesUrl field if non-nil, zero value otherwise.

### GetMessagesUrlOk

`func (o *NegotiationsObjectsTopicItem) GetMessagesUrlOk() (*string, bool)`

GetMessagesUrlOk returns a tuple with the MessagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesUrl

`func (o *NegotiationsObjectsTopicItem) SetMessagesUrl(v string)`

SetMessagesUrl sets MessagesUrl field to given value.

### HasMessagesUrl

`func (o *NegotiationsObjectsTopicItem) HasMessagesUrl() bool`

HasMessagesUrl returns a boolean if a field has been set.

### GetMessagingStatus

`func (o *NegotiationsObjectsTopicItem) GetMessagingStatus() string`

GetMessagingStatus returns the MessagingStatus field if non-nil, zero value otherwise.

### GetMessagingStatusOk

`func (o *NegotiationsObjectsTopicItem) GetMessagingStatusOk() (*string, bool)`

GetMessagingStatusOk returns a tuple with the MessagingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingStatus

`func (o *NegotiationsObjectsTopicItem) SetMessagingStatus(v string)`

SetMessagingStatus sets MessagingStatus field to given value.


### GetPhoneCalls

`func (o *NegotiationsObjectsTopicItem) GetPhoneCalls() NegotiationsPhoneCalls`

GetPhoneCalls returns the PhoneCalls field if non-nil, zero value otherwise.

### GetPhoneCallsOk

`func (o *NegotiationsObjectsTopicItem) GetPhoneCallsOk() (*NegotiationsPhoneCalls, bool)`

GetPhoneCallsOk returns a tuple with the PhoneCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCalls

`func (o *NegotiationsObjectsTopicItem) SetPhoneCalls(v NegotiationsPhoneCalls)`

SetPhoneCalls sets PhoneCalls field to given value.

### HasPhoneCalls

`func (o *NegotiationsObjectsTopicItem) HasPhoneCalls() bool`

HasPhoneCalls returns a boolean if a field has been set.

### SetPhoneCallsNil

`func (o *NegotiationsObjectsTopicItem) SetPhoneCallsNil(b bool)`

 SetPhoneCallsNil sets the value for PhoneCalls to be an explicit nil

### UnsetPhoneCalls
`func (o *NegotiationsObjectsTopicItem) UnsetPhoneCalls()`

UnsetPhoneCalls ensures that no value is present for PhoneCalls, not even an explicit nil
### GetProfessionalRoles

`func (o *NegotiationsObjectsTopicItem) GetProfessionalRoles() IncludesIdName`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *NegotiationsObjectsTopicItem) GetProfessionalRolesOk() (*IncludesIdName, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *NegotiationsObjectsTopicItem) SetProfessionalRoles(v IncludesIdName)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *NegotiationsObjectsTopicItem) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetSource

`func (o *NegotiationsObjectsTopicItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NegotiationsObjectsTopicItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NegotiationsObjectsTopicItem) SetSource(v string)`

SetSource sets Source field to given value.


### GetState

`func (o *NegotiationsObjectsTopicItem) GetState() IncludesIdName`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NegotiationsObjectsTopicItem) GetStateOk() (*IncludesIdName, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NegotiationsObjectsTopicItem) SetState(v IncludesIdName)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *NegotiationsObjectsTopicItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NegotiationsObjectsTopicItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NegotiationsObjectsTopicItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *NegotiationsObjectsTopicItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NegotiationsObjectsTopicItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NegotiationsObjectsTopicItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVacancy

`func (o *NegotiationsObjectsTopicItem) GetVacancy() VacanciesNegotiationsVacancyShort`

GetVacancy returns the Vacancy field if non-nil, zero value otherwise.

### GetVacancyOk

`func (o *NegotiationsObjectsTopicItem) GetVacancyOk() (*VacanciesNegotiationsVacancyShort, bool)`

GetVacancyOk returns a tuple with the Vacancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancy

`func (o *NegotiationsObjectsTopicItem) SetVacancy(v VacanciesNegotiationsVacancyShort)`

SetVacancy sets Vacancy field to given value.

### HasVacancy

`func (o *NegotiationsObjectsTopicItem) HasVacancy() bool`

HasVacancy returns a boolean if a field has been set.

### SetVacancyNil

`func (o *NegotiationsObjectsTopicItem) SetVacancyNil(b bool)`

 SetVacancyNil sets the value for Vacancy to be an explicit nil

### UnsetVacancy
`func (o *NegotiationsObjectsTopicItem) UnsetVacancy()`

UnsetVacancy ensures that no value is present for Vacancy, not even an explicit nil
### GetViewedByOpponent

`func (o *NegotiationsObjectsTopicItem) GetViewedByOpponent() bool`

GetViewedByOpponent returns the ViewedByOpponent field if non-nil, zero value otherwise.

### GetViewedByOpponentOk

`func (o *NegotiationsObjectsTopicItem) GetViewedByOpponentOk() (*bool, bool)`

GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByOpponent

`func (o *NegotiationsObjectsTopicItem) SetViewedByOpponent(v bool)`

SetViewedByOpponent sets ViewedByOpponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


