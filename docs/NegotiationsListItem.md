# NegotiationsListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | Pointer to [**NegotiationsObjectsCounters**](NegotiationsObjectsCounters.md) |  | [optional] 
**CreatedAt** | **string** | Дата и время создания отклика/приглашения | 
**HasUpdates** | **bool** | Есть ли в откликах/приглашениях по данной вакансии обновления, требующие внимания | 
**Id** | **string** | Идентификатор отклика/приглашения | 
**MessagesUrl** | Pointer to **string** | URL, на который необходимо делать GET-запрос для получения [списка сообщений в отклике/приглашении](#tag/Perepiska-(otklikipriglasheniya)-dlya-soiskatelya/operation/get-negotiation-messages). Если &#x60;can_edit&#x60; равно &#x60;false&#x60;, значение поля должно игнорироваться | [optional] 
**MessagingStatus** | **string** | Текущий статус переписки.  Возможные значения приведены в поле &#x60;messaging_status&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**ProfessionalRoles** | Pointer to [**[]VacancyProfessionalRoleItem**](VacancyProfessionalRoleItem.md) | Список профессиональных ролей | [optional] 
**Source** | Pointer to **string** | Источник отклика/приглашения | [optional] 
**State** | [**IncludesIdName**](IncludesIdName.md) | Текущее состояние отклика/приглашения.  Возможные значения приведены в поле &#x60;negotiations_state&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**UpdatedAt** | **string** | Дата и время последнего обновления отклика/приглашения | 
**ViewedByOpponent** | **bool** | Был ли отклик просмотрен работодателем | 
**DeclineAllowed** | **bool** | Можно ли [скрыть отклик](#tag/Perepiska-(otklikipriglasheniya)-dlya-soiskatelya/operation/hide-active-response) вместе с сообщением работодателю об отказе | 
**Hidden** | **bool** | Скрыт ли текущий отклик от соискателя | 
**JobSearchStatus** | Pointer to [**IncludesIdName**](IncludesIdName.md) |  | [optional] 
**PhoneCalls** | Pointer to [**NegotiationsPhoneCalls**](NegotiationsPhoneCalls.md) |  | [optional] 
**Vacancy** | Pointer to [**VacanciesNegotiationsVacancyShort**](VacanciesNegotiationsVacancyShort.md) |  | [optional] 
**Resume** | Pointer to [**NullableResumeResumeNanoWithUrl**](ResumeResumeNanoWithUrl.md) |  | [optional] 
**Url** | **string** | Ссылка на полную версию отклика | 

## Methods

### NewNegotiationsListItem

`func NewNegotiationsListItem(createdAt string, hasUpdates bool, id string, messagingStatus string, state IncludesIdName, updatedAt string, viewedByOpponent bool, declineAllowed bool, hidden bool, url string, ) *NegotiationsListItem`

NewNegotiationsListItem instantiates a new NegotiationsListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsListItemWithDefaults

`func NewNegotiationsListItemWithDefaults() *NegotiationsListItem`

NewNegotiationsListItemWithDefaults instantiates a new NegotiationsListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *NegotiationsListItem) GetCounters() NegotiationsObjectsCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *NegotiationsListItem) GetCountersOk() (*NegotiationsObjectsCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *NegotiationsListItem) SetCounters(v NegotiationsObjectsCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *NegotiationsListItem) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NegotiationsListItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsListItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsListItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetHasUpdates

`func (o *NegotiationsListItem) GetHasUpdates() bool`

GetHasUpdates returns the HasUpdates field if non-nil, zero value otherwise.

### GetHasUpdatesOk

`func (o *NegotiationsListItem) GetHasUpdatesOk() (*bool, bool)`

GetHasUpdatesOk returns a tuple with the HasUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdates

`func (o *NegotiationsListItem) SetHasUpdates(v bool)`

SetHasUpdates sets HasUpdates field to given value.


### GetId

`func (o *NegotiationsListItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsListItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsListItem) SetId(v string)`

SetId sets Id field to given value.


### GetMessagesUrl

`func (o *NegotiationsListItem) GetMessagesUrl() string`

GetMessagesUrl returns the MessagesUrl field if non-nil, zero value otherwise.

### GetMessagesUrlOk

`func (o *NegotiationsListItem) GetMessagesUrlOk() (*string, bool)`

GetMessagesUrlOk returns a tuple with the MessagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesUrl

`func (o *NegotiationsListItem) SetMessagesUrl(v string)`

SetMessagesUrl sets MessagesUrl field to given value.

### HasMessagesUrl

`func (o *NegotiationsListItem) HasMessagesUrl() bool`

HasMessagesUrl returns a boolean if a field has been set.

### GetMessagingStatus

`func (o *NegotiationsListItem) GetMessagingStatus() string`

GetMessagingStatus returns the MessagingStatus field if non-nil, zero value otherwise.

### GetMessagingStatusOk

`func (o *NegotiationsListItem) GetMessagingStatusOk() (*string, bool)`

GetMessagingStatusOk returns a tuple with the MessagingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingStatus

`func (o *NegotiationsListItem) SetMessagingStatus(v string)`

SetMessagingStatus sets MessagingStatus field to given value.


### GetProfessionalRoles

`func (o *NegotiationsListItem) GetProfessionalRoles() []VacancyProfessionalRoleItem`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *NegotiationsListItem) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItem, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *NegotiationsListItem) SetProfessionalRoles(v []VacancyProfessionalRoleItem)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *NegotiationsListItem) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetSource

`func (o *NegotiationsListItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NegotiationsListItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NegotiationsListItem) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NegotiationsListItem) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *NegotiationsListItem) GetState() IncludesIdName`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NegotiationsListItem) GetStateOk() (*IncludesIdName, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NegotiationsListItem) SetState(v IncludesIdName)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *NegotiationsListItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NegotiationsListItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NegotiationsListItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetViewedByOpponent

`func (o *NegotiationsListItem) GetViewedByOpponent() bool`

GetViewedByOpponent returns the ViewedByOpponent field if non-nil, zero value otherwise.

### GetViewedByOpponentOk

`func (o *NegotiationsListItem) GetViewedByOpponentOk() (*bool, bool)`

GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByOpponent

`func (o *NegotiationsListItem) SetViewedByOpponent(v bool)`

SetViewedByOpponent sets ViewedByOpponent field to given value.


### GetDeclineAllowed

`func (o *NegotiationsListItem) GetDeclineAllowed() bool`

GetDeclineAllowed returns the DeclineAllowed field if non-nil, zero value otherwise.

### GetDeclineAllowedOk

`func (o *NegotiationsListItem) GetDeclineAllowedOk() (*bool, bool)`

GetDeclineAllowedOk returns a tuple with the DeclineAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineAllowed

`func (o *NegotiationsListItem) SetDeclineAllowed(v bool)`

SetDeclineAllowed sets DeclineAllowed field to given value.


### GetHidden

`func (o *NegotiationsListItem) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *NegotiationsListItem) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *NegotiationsListItem) SetHidden(v bool)`

SetHidden sets Hidden field to given value.


### GetJobSearchStatus

`func (o *NegotiationsListItem) GetJobSearchStatus() IncludesIdName`

GetJobSearchStatus returns the JobSearchStatus field if non-nil, zero value otherwise.

### GetJobSearchStatusOk

`func (o *NegotiationsListItem) GetJobSearchStatusOk() (*IncludesIdName, bool)`

GetJobSearchStatusOk returns a tuple with the JobSearchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobSearchStatus

`func (o *NegotiationsListItem) SetJobSearchStatus(v IncludesIdName)`

SetJobSearchStatus sets JobSearchStatus field to given value.

### HasJobSearchStatus

`func (o *NegotiationsListItem) HasJobSearchStatus() bool`

HasJobSearchStatus returns a boolean if a field has been set.

### GetPhoneCalls

`func (o *NegotiationsListItem) GetPhoneCalls() NegotiationsPhoneCalls`

GetPhoneCalls returns the PhoneCalls field if non-nil, zero value otherwise.

### GetPhoneCallsOk

`func (o *NegotiationsListItem) GetPhoneCallsOk() (*NegotiationsPhoneCalls, bool)`

GetPhoneCallsOk returns a tuple with the PhoneCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneCalls

`func (o *NegotiationsListItem) SetPhoneCalls(v NegotiationsPhoneCalls)`

SetPhoneCalls sets PhoneCalls field to given value.

### HasPhoneCalls

`func (o *NegotiationsListItem) HasPhoneCalls() bool`

HasPhoneCalls returns a boolean if a field has been set.

### GetVacancy

`func (o *NegotiationsListItem) GetVacancy() VacanciesNegotiationsVacancyShort`

GetVacancy returns the Vacancy field if non-nil, zero value otherwise.

### GetVacancyOk

`func (o *NegotiationsListItem) GetVacancyOk() (*VacanciesNegotiationsVacancyShort, bool)`

GetVacancyOk returns a tuple with the Vacancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancy

`func (o *NegotiationsListItem) SetVacancy(v VacanciesNegotiationsVacancyShort)`

SetVacancy sets Vacancy field to given value.

### HasVacancy

`func (o *NegotiationsListItem) HasVacancy() bool`

HasVacancy returns a boolean if a field has been set.

### GetResume

`func (o *NegotiationsListItem) GetResume() ResumeResumeNanoWithUrl`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *NegotiationsListItem) GetResumeOk() (*ResumeResumeNanoWithUrl, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *NegotiationsListItem) SetResume(v ResumeResumeNanoWithUrl)`

SetResume sets Resume field to given value.

### HasResume

`func (o *NegotiationsListItem) HasResume() bool`

HasResume returns a boolean if a field has been set.

### SetResumeNil

`func (o *NegotiationsListItem) SetResumeNil(b bool)`

 SetResumeNil sets the value for Resume to be an explicit nil

### UnsetResume
`func (o *NegotiationsListItem) UnsetResume()`

UnsetResume ensures that no value is present for Resume, not even an explicit nil
### GetUrl

`func (o *NegotiationsListItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NegotiationsListItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NegotiationsListItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


