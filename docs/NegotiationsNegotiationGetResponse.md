# NegotiationsNegotiationGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | Pointer to [**NegotiationsObjectsCounters**](NegotiationsObjectsCounters.md) |  | [optional] 
**CreatedAt** | **string** | Дата и время создания отклика | 
**DeclineAllowed** | **bool** | Можно ли [скрыть отклик](#tag/Perepiska-(otklikipriglasheniya)-dlya-soiskatelya/operation/hide-active-response) вместе с сообщением работодателю об отказе | 
**HasUpdates** | **bool** | Есть ли непросмотренные сообщения в отклике. Флаг сбрасывается при различных действиях с откликом, например, [просмотре списка сообщений](https://github.com/hhru/api/blob/master/docs/negotiations.md#get_messages) | 
**Hidden** | **bool** | Скрыт ли текущий отклик от соискателя | 
**Id** | **string** | Идентификатор отклика | 
**JobSearchStatus** | Pointer to [**IncludesIdName**](IncludesIdName.md) |  | [optional] 
**MessagesUrl** | Pointer to **string** | URL, на который необходимо делать GET запрос для получения [списка сообщений в отклике/приглашении](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#get-messages). Если &#x60;can_edit&#x60; равно &#x60;false&#x60;, значение поля должно игнорироваться | [optional] 
**MessagingStatus** | **string** | Текущий статус переписки.  Возможные значения приведены в поле &#x60;messaging_status&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**PhoneCalls** | Pointer to [**NegotiationsPhoneCalls**](NegotiationsPhoneCalls.md) |  | [optional] 
**ProfessionalRoles** | Pointer to [**IncludesIdName**](IncludesIdName.md) |  | [optional] 
**Source** | **string** | Источник отклика | 
**State** | [**IncludesIdName**](IncludesIdName.md) | Текущее состояние отклика.  Возможные значения приведены в поле &#x60;negotiations_state&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**UpdatedAt** | **string** | Дата и время последнего обновления отклика | 
**Url** | **string** | Ссылка на полную версию отклика | 
**Vacancy** | Pointer to [**VacanciesNegotiationsVacancyShort**](VacanciesNegotiationsVacancyShort.md) |  | [optional] 
**ViewedByOpponent** | **bool** | Был ли отклик просмотрен работодателем | 
**Resume** | Pointer to [**NullableNegotiationsObjectsTopicResume**](NegotiationsObjectsTopicResume.md) |  | [optional] 

## Methods

### NewNegotiationsNegotiationGetResponse

`func NewNegotiationsNegotiationGetResponse(createdAt string, declineAllowed bool, hasUpdates bool, hidden bool, id string, messagingStatus string, source string, state IncludesIdName, updatedAt string, url string, viewedByOpponent bool, ) *NegotiationsNegotiationGetResponse`

NewNegotiationsNegotiationGetResponse instantiates a new NegotiationsNegotiationGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsNegotiationGetResponseWithDefaults

`func NewNegotiationsNegotiationGetResponseWithDefaults() *NegotiationsNegotiationGetResponse`

NewNegotiationsNegotiationGetResponseWithDefaults instantiates a new NegotiationsNegotiationGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *NegotiationsNegotiationGetResponse) GetCounters() NegotiationsObjectsCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *NegotiationsNegotiationGetResponse) GetCountersOk() (*NegotiationsObjectsCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *NegotiationsNegotiationGetResponse) SetCounters(v NegotiationsObjectsCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *NegotiationsNegotiationGetResponse) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NegotiationsNegotiationGetResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsNegotiationGetResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsNegotiationGetResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetDeclineAllowed

`func (o *NegotiationsNegotiationGetResponse) GetDeclineAllowed() bool`

GetDeclineAllowed returns the DeclineAllowed field if non-nil, zero value otherwise.

### GetDeclineAllowedOk

`func (o *NegotiationsNegotiationGetResponse) GetDeclineAllowedOk() (*bool, bool)`

GetDeclineAllowedOk returns a tuple with the DeclineAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineAllowed

`func (o *NegotiationsNegotiationGetResponse) SetDeclineAllowed(v bool)`

SetDeclineAllowed sets DeclineAllowed field to given value.


### GetHasUpdates

`func (o *NegotiationsNegotiationGetResponse) GetHasUpdates() bool`

GetHasUpdates returns the HasUpdates field if non-nil, zero value otherwise.

### GetHasUpdatesOk

`func (o *NegotiationsNegotiationGetResponse) GetHasUpdatesOk() (*bool, bool)`

GetHasUpdatesOk returns a tuple with the HasUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdates

`func (o *NegotiationsNegotiationGetResponse) SetHasUpdates(v bool)`

SetHasUpdates sets HasUpdates field to given value.


### GetHidden

`func (o *NegotiationsNegotiationGetResponse) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *NegotiationsNegotiationGetResponse) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *NegotiationsNegotiationGetResponse) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetId

`func (o *NegotiationsNegotiationGetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsNegotiationGetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsNegotiationGetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetJobSearchStatus

`func (o *NegotiationsNegotiationGetResponse) GetJobSearchStatus() IncludesIdName`

GetJobSearchStatus returns the JobSearchStatus field if non-nil, zero value otherwise.

### GetJobSearchStatusOk

`func (o *NegotiationsNegotiationGetResponse) GetJobSearchStatusOk() (*IncludesIdName, bool)`

GetJobSearchStatusOk returns a tuple with the JobSearchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatus

`func (o *NegotiationsNegotiationGetResponse) SetJobSearchStatus(v IncludesIdName)`

SetJobSearchStatus sets JobSearchStatus field to given value.

### HasJobSearchStatus

`func (o *NegotiationsNegotiationGetResponse) HasJobSearchStatus() bool`

HasJobSearchStatus returns a boolean if a field has been set.

### GetMessagesUrl

`func (o *NegotiationsNegotiationGetResponse) GetMessagesUrl() string`

GetMessagesUrl returns the MessagesUrl field if non-nil, zero value otherwise.

### GetMessagesUrlOk

`func (o *NegotiationsNegotiationGetResponse) GetMessagesUrlOk() (*string, bool)`

GetMessagesUrlOk returns a tuple with the MessagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesUrl

`func (o *NegotiationsNegotiationGetResponse) SetMessagesUrl(v string)`

SetMessagesUrl sets MessagesUrl field to given value.

### HasMessagesUrl

`func (o *NegotiationsNegotiationGetResponse) HasMessagesUrl() bool`

HasMessagesUrl returns a boolean if a field has been set.

### GetMessagingStatus

`func (o *NegotiationsNegotiationGetResponse) GetMessagingStatus() string`

GetMessagingStatus returns the MessagingStatus field if non-nil, zero value otherwise.

### GetMessagingStatusOk

`func (o *NegotiationsNegotiationGetResponse) GetMessagingStatusOk() (*string, bool)`

GetMessagingStatusOk returns a tuple with the MessagingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingStatus

`func (o *NegotiationsNegotiationGetResponse) SetMessagingStatus(v string)`

SetMessagingStatus sets MessagingStatus field to given value.


### GetPhoneCalls

`func (o *NegotiationsNegotiationGetResponse) GetPhoneCalls() NegotiationsPhoneCalls`

GetPhoneCalls returns the PhoneCalls field if non-nil, zero value otherwise.

### GetPhoneCallsOk

`func (o *NegotiationsNegotiationGetResponse) GetPhoneCallsOk() (*NegotiationsPhoneCalls, bool)`

GetPhoneCallsOk returns a tuple with the PhoneCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCalls

`func (o *NegotiationsNegotiationGetResponse) SetPhoneCalls(v NegotiationsPhoneCalls)`

SetPhoneCalls sets PhoneCalls field to given value.

### HasPhoneCalls

`func (o *NegotiationsNegotiationGetResponse) HasPhoneCalls() bool`

HasPhoneCalls returns a boolean if a field has been set.

### GetProfessionalRoles

`func (o *NegotiationsNegotiationGetResponse) GetProfessionalRoles() IncludesIdName`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *NegotiationsNegotiationGetResponse) GetProfessionalRolesOk() (*IncludesIdName, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *NegotiationsNegotiationGetResponse) SetProfessionalRoles(v IncludesIdName)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *NegotiationsNegotiationGetResponse) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetSource

`func (o *NegotiationsNegotiationGetResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NegotiationsNegotiationGetResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NegotiationsNegotiationGetResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetState

`func (o *NegotiationsNegotiationGetResponse) GetState() IncludesIdName`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NegotiationsNegotiationGetResponse) GetStateOk() (*IncludesIdName, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NegotiationsNegotiationGetResponse) SetState(v IncludesIdName)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *NegotiationsNegotiationGetResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NegotiationsNegotiationGetResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NegotiationsNegotiationGetResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUrl

`func (o *NegotiationsNegotiationGetResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NegotiationsNegotiationGetResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NegotiationsNegotiationGetResponse) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVacancy

`func (o *NegotiationsNegotiationGetResponse) GetVacancy() VacanciesNegotiationsVacancyShort`

GetVacancy returns the Vacancy field if non-nil, zero value otherwise.

### GetVacancyOk

`func (o *NegotiationsNegotiationGetResponse) GetVacancyOk() (*VacanciesNegotiationsVacancyShort, bool)`

GetVacancyOk returns a tuple with the Vacancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancy

`func (o *NegotiationsNegotiationGetResponse) SetVacancy(v VacanciesNegotiationsVacancyShort)`

SetVacancy sets Vacancy field to given value.

### HasVacancy

`func (o *NegotiationsNegotiationGetResponse) HasVacancy() bool`

HasVacancy returns a boolean if a field has been set.

### GetViewedByOpponent

`func (o *NegotiationsNegotiationGetResponse) GetViewedByOpponent() bool`

GetViewedByOpponent returns the ViewedByOpponent field if non-nil, zero value otherwise.

### GetViewedByOpponentOk

`func (o *NegotiationsNegotiationGetResponse) GetViewedByOpponentOk() (*bool, bool)`

GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByOpponent

`func (o *NegotiationsNegotiationGetResponse) SetViewedByOpponent(v bool)`

SetViewedByOpponent sets ViewedByOpponent field to given value.


### GetResume

`func (o *NegotiationsNegotiationGetResponse) GetResume() NegotiationsObjectsTopicResume`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *NegotiationsNegotiationGetResponse) GetResumeOk() (*NegotiationsObjectsTopicResume, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *NegotiationsNegotiationGetResponse) SetResume(v NegotiationsObjectsTopicResume)`

SetResume sets Resume field to given value.

### HasResume

`func (o *NegotiationsNegotiationGetResponse) HasResume() bool`

HasResume returns a boolean if a field has been set.

### SetResumeNil

`func (o *NegotiationsNegotiationGetResponse) SetResumeNil(b bool)`

 SetResumeNil sets the value for Resume to be an explicit nil

### UnsetResume
`func (o *NegotiationsNegotiationGetResponse) UnsetResume()`

UnsetResume ensures that no value is present for Resume, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


