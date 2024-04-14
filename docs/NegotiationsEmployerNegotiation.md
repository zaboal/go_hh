# NegotiationsEmployerNegotiation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | Pointer to [**NegotiationsObjectsCounters**](NegotiationsObjectsCounters.md) |  | [optional] 
**CreatedAt** | **string** | Дата и время создания отклика/приглашения | 
**HasUpdates** | **bool** | Есть ли в откликах/приглашениях по данной вакансии обновления, требующие внимания | 
**Id** | **string** | Идентификатор отклика/приглашения | 
**MessagesUrl** | Pointer to **string** | URL, на который необходимо делать GET-запрос для получения [списка сообщений в отклике/приглашении](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#get-messages). Если &#x60;can_edit&#x60; равно &#x60;false&#x60;, значение поля должно игнорироваться | [optional] 
**MessagingStatus** | **string** | Текущий статус переписки.  Возможные значения приведены в поле &#x60;messaging_status&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**ProfessionalRoles** | Pointer to [**[]VacancyProfessionalRoleItem**](VacancyProfessionalRoleItem.md) | Список профессиональных ролей | [optional] 
**Source** | Pointer to **string** | Источник отклика/приглашения | [optional] 
**State** | [**IncludesIdName**](IncludesIdName.md) | Текущее состояние отклика/приглашения.  Возможные значения приведены в поле &#x60;negotiations_state&#x60; [справочника полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**UpdatedAt** | **string** | Дата и время последнего обновления отклика/приглашения | 
**ViewedByOpponent** | **bool** | Был ли отклик просмотрен работодателем | 
**Actions** | Pointer to [**[]VacancyNegotiationActions**](VacancyNegotiationActions.md) | Возможные [действия по отклику/приглашению](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#actions-info)  | [optional] 
**EmployerState** | Pointer to [**EmployersEmployersState**](EmployersEmployersState.md) |  | [optional] 
**Templates** | Pointer to [**[]VacancyTemplates**](VacancyTemplates.md) | Шаблоны писем | [optional] 
**TestResult** | Pointer to [**SkillVerificationsTestResultWithUrl**](SkillVerificationsTestResultWithUrl.md) |  | [optional] 
**Resume** | Pointer to [**NullableNegotiationsObjectsEmployerTopicResume**](NegotiationsObjectsEmployerTopicResume.md) |  | [optional] 
**Vacancy** | Pointer to [**NullableVacanciesNegotiationsVacancyShort**](VacanciesNegotiationsVacancyShort.md) |  | [optional] 

## Methods

### NewNegotiationsEmployerNegotiation

`func NewNegotiationsEmployerNegotiation(createdAt string, hasUpdates bool, id string, messagingStatus string, state IncludesIdName, updatedAt string, viewedByOpponent bool, ) *NegotiationsEmployerNegotiation`

NewNegotiationsEmployerNegotiation instantiates a new NegotiationsEmployerNegotiation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsEmployerNegotiationWithDefaults

`func NewNegotiationsEmployerNegotiationWithDefaults() *NegotiationsEmployerNegotiation`

NewNegotiationsEmployerNegotiationWithDefaults instantiates a new NegotiationsEmployerNegotiation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *NegotiationsEmployerNegotiation) GetCounters() NegotiationsObjectsCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *NegotiationsEmployerNegotiation) GetCountersOk() (*NegotiationsObjectsCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *NegotiationsEmployerNegotiation) SetCounters(v NegotiationsObjectsCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *NegotiationsEmployerNegotiation) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NegotiationsEmployerNegotiation) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsEmployerNegotiation) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsEmployerNegotiation) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetHasUpdates

`func (o *NegotiationsEmployerNegotiation) GetHasUpdates() bool`

GetHasUpdates returns the HasUpdates field if non-nil, zero value otherwise.

### GetHasUpdatesOk

`func (o *NegotiationsEmployerNegotiation) GetHasUpdatesOk() (*bool, bool)`

GetHasUpdatesOk returns a tuple with the HasUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdates

`func (o *NegotiationsEmployerNegotiation) SetHasUpdates(v bool)`

SetHasUpdates sets HasUpdates field to given value.


### GetId

`func (o *NegotiationsEmployerNegotiation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsEmployerNegotiation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsEmployerNegotiation) SetId(v string)`

SetId sets Id field to given value.


### GetMessagesUrl

`func (o *NegotiationsEmployerNegotiation) GetMessagesUrl() string`

GetMessagesUrl returns the MessagesUrl field if non-nil, zero value otherwise.

### GetMessagesUrlOk

`func (o *NegotiationsEmployerNegotiation) GetMessagesUrlOk() (*string, bool)`

GetMessagesUrlOk returns a tuple with the MessagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesUrl

`func (o *NegotiationsEmployerNegotiation) SetMessagesUrl(v string)`

SetMessagesUrl sets MessagesUrl field to given value.

### HasMessagesUrl

`func (o *NegotiationsEmployerNegotiation) HasMessagesUrl() bool`

HasMessagesUrl returns a boolean if a field has been set.

### GetMessagingStatus

`func (o *NegotiationsEmployerNegotiation) GetMessagingStatus() string`

GetMessagingStatus returns the MessagingStatus field if non-nil, zero value otherwise.

### GetMessagingStatusOk

`func (o *NegotiationsEmployerNegotiation) GetMessagingStatusOk() (*string, bool)`

GetMessagingStatusOk returns a tuple with the MessagingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingStatus

`func (o *NegotiationsEmployerNegotiation) SetMessagingStatus(v string)`

SetMessagingStatus sets MessagingStatus field to given value.


### GetProfessionalRoles

`func (o *NegotiationsEmployerNegotiation) GetProfessionalRoles() []VacancyProfessionalRoleItem`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *NegotiationsEmployerNegotiation) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItem, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *NegotiationsEmployerNegotiation) SetProfessionalRoles(v []VacancyProfessionalRoleItem)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *NegotiationsEmployerNegotiation) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetSource

`func (o *NegotiationsEmployerNegotiation) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NegotiationsEmployerNegotiation) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NegotiationsEmployerNegotiation) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NegotiationsEmployerNegotiation) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *NegotiationsEmployerNegotiation) GetState() IncludesIdName`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NegotiationsEmployerNegotiation) GetStateOk() (*IncludesIdName, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NegotiationsEmployerNegotiation) SetState(v IncludesIdName)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *NegotiationsEmployerNegotiation) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NegotiationsEmployerNegotiation) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NegotiationsEmployerNegotiation) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetViewedByOpponent

`func (o *NegotiationsEmployerNegotiation) GetViewedByOpponent() bool`

GetViewedByOpponent returns the ViewedByOpponent field if non-nil, zero value otherwise.

### GetViewedByOpponentOk

`func (o *NegotiationsEmployerNegotiation) GetViewedByOpponentOk() (*bool, bool)`

GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByOpponent

`func (o *NegotiationsEmployerNegotiation) SetViewedByOpponent(v bool)`

SetViewedByOpponent sets ViewedByOpponent field to given value.


### GetActions

`func (o *NegotiationsEmployerNegotiation) GetActions() []VacancyNegotiationActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *NegotiationsEmployerNegotiation) GetActionsOk() (*[]VacancyNegotiationActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *NegotiationsEmployerNegotiation) SetActions(v []VacancyNegotiationActions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *NegotiationsEmployerNegotiation) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetEmployerState

`func (o *NegotiationsEmployerNegotiation) GetEmployerState() EmployersEmployersState`

GetEmployerState returns the EmployerState field if non-nil, zero value otherwise.

### GetEmployerStateOk

`func (o *NegotiationsEmployerNegotiation) GetEmployerStateOk() (*EmployersEmployersState, bool)`

GetEmployerStateOk returns a tuple with the EmployerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerState

`func (o *NegotiationsEmployerNegotiation) SetEmployerState(v EmployersEmployersState)`

SetEmployerState sets EmployerState field to given value.

### HasEmployerState

`func (o *NegotiationsEmployerNegotiation) HasEmployerState() bool`

HasEmployerState returns a boolean if a field has been set.

### GetTemplates

`func (o *NegotiationsEmployerNegotiation) GetTemplates() []VacancyTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *NegotiationsEmployerNegotiation) GetTemplatesOk() (*[]VacancyTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *NegotiationsEmployerNegotiation) SetTemplates(v []VacancyTemplates)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *NegotiationsEmployerNegotiation) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetTestResult

`func (o *NegotiationsEmployerNegotiation) GetTestResult() SkillVerificationsTestResultWithUrl`

GetTestResult returns the TestResult field if non-nil, zero value otherwise.

### GetTestResultOk

`func (o *NegotiationsEmployerNegotiation) GetTestResultOk() (*SkillVerificationsTestResultWithUrl, bool)`

GetTestResultOk returns a tuple with the TestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResult

`func (o *NegotiationsEmployerNegotiation) SetTestResult(v SkillVerificationsTestResultWithUrl)`

SetTestResult sets TestResult field to given value.

### HasTestResult

`func (o *NegotiationsEmployerNegotiation) HasTestResult() bool`

HasTestResult returns a boolean if a field has been set.

### GetResume

`func (o *NegotiationsEmployerNegotiation) GetResume() NegotiationsObjectsEmployerTopicResume`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *NegotiationsEmployerNegotiation) GetResumeOk() (*NegotiationsObjectsEmployerTopicResume, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *NegotiationsEmployerNegotiation) SetResume(v NegotiationsObjectsEmployerTopicResume)`

SetResume sets Resume field to given value.

### HasResume

`func (o *NegotiationsEmployerNegotiation) HasResume() bool`

HasResume returns a boolean if a field has been set.

### SetResumeNil

`func (o *NegotiationsEmployerNegotiation) SetResumeNil(b bool)`

 SetResumeNil sets the value for Resume to be an explicit nil

### UnsetResume
`func (o *NegotiationsEmployerNegotiation) UnsetResume()`

UnsetResume ensures that no value is present for Resume, not even an explicit nil
### GetVacancy

`func (o *NegotiationsEmployerNegotiation) GetVacancy() VacanciesNegotiationsVacancyShort`

GetVacancy returns the Vacancy field if non-nil, zero value otherwise.

### GetVacancyOk

`func (o *NegotiationsEmployerNegotiation) GetVacancyOk() (*VacanciesNegotiationsVacancyShort, bool)`

GetVacancyOk returns a tuple with the Vacancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancy

`func (o *NegotiationsEmployerNegotiation) SetVacancy(v VacanciesNegotiationsVacancyShort)`

SetVacancy sets Vacancy field to given value.

### HasVacancy

`func (o *NegotiationsEmployerNegotiation) HasVacancy() bool`

HasVacancy returns a boolean if a field has been set.

### SetVacancyNil

`func (o *NegotiationsEmployerNegotiation) SetVacancyNil(b bool)`

 SetVacancyNil sets the value for Vacancy to be an explicit nil

### UnsetVacancy
`func (o *NegotiationsEmployerNegotiation) UnsetVacancy()`

UnsetVacancy ensures that no value is present for Vacancy, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


