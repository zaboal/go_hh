# NegotiationsCollectionNegotiationsItemsInner

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
**Actions** | [**[]VacancyNegotiationActions**](VacancyNegotiationActions.md) | Возможные [действия по отклику/приглашению](https://github.com/hhru/api/blob/master/docs/employer_negotiations.md#actions-info)  | 
**EmployerState** | [**EmployersEmployersState**](EmployersEmployersState.md) |  | 
**FunnelStage** | Pointer to [**NullableEmployersFunnelStage**](EmployersFunnelStage.md) |  | [optional] 
**Templates** | [**[]VacancyTemplates**](VacancyTemplates.md) | Шаблоны писем | 
**TestResult** | Pointer to [**NullableSkillVerificationsTestResultWithUrl**](SkillVerificationsTestResultWithUrl.md) |  | [optional] 
**Resume** | Pointer to [**NullableResumeResumeShortLogoUrl90**](ResumeResumeShortLogoUrl90.md) |  | [optional] 
**Url** | **string** | Ссылка на полную версию отклика | 

## Methods

### NewNegotiationsCollectionNegotiationsItemsInner

`func NewNegotiationsCollectionNegotiationsItemsInner(createdAt string, hasUpdates bool, id string, messagingStatus string, state IncludesIdName, updatedAt string, viewedByOpponent bool, actions []VacancyNegotiationActions, employerState EmployersEmployersState, templates []VacancyTemplates, url string, ) *NegotiationsCollectionNegotiationsItemsInner`

NewNegotiationsCollectionNegotiationsItemsInner instantiates a new NegotiationsCollectionNegotiationsItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsCollectionNegotiationsItemsInnerWithDefaults

`func NewNegotiationsCollectionNegotiationsItemsInnerWithDefaults() *NegotiationsCollectionNegotiationsItemsInner`

NewNegotiationsCollectionNegotiationsItemsInnerWithDefaults instantiates a new NegotiationsCollectionNegotiationsItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetCounters() NegotiationsObjectsCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetCountersOk() (*NegotiationsObjectsCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetCounters(v NegotiationsObjectsCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *NegotiationsCollectionNegotiationsItemsInner) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetHasUpdates

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetHasUpdates() bool`

GetHasUpdates returns the HasUpdates field if non-nil, zero value otherwise.

### GetHasUpdatesOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetHasUpdatesOk() (*bool, bool)`

GetHasUpdatesOk returns a tuple with the HasUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdates

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetHasUpdates(v bool)`

SetHasUpdates sets HasUpdates field to given value.


### GetId

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetId(v string)`

SetId sets Id field to given value.


### GetMessagesUrl

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetMessagesUrl() string`

GetMessagesUrl returns the MessagesUrl field if non-nil, zero value otherwise.

### GetMessagesUrlOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetMessagesUrlOk() (*string, bool)`

GetMessagesUrlOk returns a tuple with the MessagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesUrl

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetMessagesUrl(v string)`

SetMessagesUrl sets MessagesUrl field to given value.

### HasMessagesUrl

`func (o *NegotiationsCollectionNegotiationsItemsInner) HasMessagesUrl() bool`

HasMessagesUrl returns a boolean if a field has been set.

### GetMessagingStatus

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetMessagingStatus() string`

GetMessagingStatus returns the MessagingStatus field if non-nil, zero value otherwise.

### GetMessagingStatusOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetMessagingStatusOk() (*string, bool)`

GetMessagingStatusOk returns a tuple with the MessagingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingStatus

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetMessagingStatus(v string)`

SetMessagingStatus sets MessagingStatus field to given value.


### GetProfessionalRoles

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetProfessionalRoles() []VacancyProfessionalRoleItem`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItem, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetProfessionalRoles(v []VacancyProfessionalRoleItem)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *NegotiationsCollectionNegotiationsItemsInner) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetSource

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NegotiationsCollectionNegotiationsItemsInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetState() IncludesIdName`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetStateOk() (*IncludesIdName, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetState(v IncludesIdName)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetViewedByOpponent

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetViewedByOpponent() bool`

GetViewedByOpponent returns the ViewedByOpponent field if non-nil, zero value otherwise.

### GetViewedByOpponentOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetViewedByOpponentOk() (*bool, bool)`

GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByOpponent

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetViewedByOpponent(v bool)`

SetViewedByOpponent sets ViewedByOpponent field to given value.


### GetActions

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetActions() []VacancyNegotiationActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetActionsOk() (*[]VacancyNegotiationActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetActions(v []VacancyNegotiationActions)`

SetActions sets Actions field to given value.


### GetEmployerState

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetEmployerState() EmployersEmployersState`

GetEmployerState returns the EmployerState field if non-nil, zero value otherwise.

### GetEmployerStateOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetEmployerStateOk() (*EmployersEmployersState, bool)`

GetEmployerStateOk returns a tuple with the EmployerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerState

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetEmployerState(v EmployersEmployersState)`

SetEmployerState sets EmployerState field to given value.


### GetFunnelStage

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetFunnelStage() EmployersFunnelStage`

GetFunnelStage returns the FunnelStage field if non-nil, zero value otherwise.

### GetFunnelStageOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetFunnelStageOk() (*EmployersFunnelStage, bool)`

GetFunnelStageOk returns a tuple with the FunnelStage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunnelStage

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetFunnelStage(v EmployersFunnelStage)`

SetFunnelStage sets FunnelStage field to given value.

### HasFunnelStage

`func (o *NegotiationsCollectionNegotiationsItemsInner) HasFunnelStage() bool`

HasFunnelStage returns a boolean if a field has been set.

### SetFunnelStageNil

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetFunnelStageNil(b bool)`

 SetFunnelStageNil sets the value for FunnelStage to be an explicit nil

### UnsetFunnelStage
`func (o *NegotiationsCollectionNegotiationsItemsInner) UnsetFunnelStage()`

UnsetFunnelStage ensures that no value is present for FunnelStage, not even an explicit nil
### GetTemplates

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetTemplates() []VacancyTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetTemplatesOk() (*[]VacancyTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetTemplates(v []VacancyTemplates)`

SetTemplates sets Templates field to given value.


### GetTestResult

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetTestResult() SkillVerificationsTestResultWithUrl`

GetTestResult returns the TestResult field if non-nil, zero value otherwise.

### GetTestResultOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetTestResultOk() (*SkillVerificationsTestResultWithUrl, bool)`

GetTestResultOk returns a tuple with the TestResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestResult

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetTestResult(v SkillVerificationsTestResultWithUrl)`

SetTestResult sets TestResult field to given value.

### HasTestResult

`func (o *NegotiationsCollectionNegotiationsItemsInner) HasTestResult() bool`

HasTestResult returns a boolean if a field has been set.

### SetTestResultNil

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetTestResultNil(b bool)`

 SetTestResultNil sets the value for TestResult to be an explicit nil

### UnsetTestResult
`func (o *NegotiationsCollectionNegotiationsItemsInner) UnsetTestResult()`

UnsetTestResult ensures that no value is present for TestResult, not even an explicit nil
### GetResume

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetResume() ResumeResumeShortLogoUrl90`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetResumeOk() (*ResumeResumeShortLogoUrl90, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetResume(v ResumeResumeShortLogoUrl90)`

SetResume sets Resume field to given value.

### HasResume

`func (o *NegotiationsCollectionNegotiationsItemsInner) HasResume() bool`

HasResume returns a boolean if a field has been set.

### SetResumeNil

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetResumeNil(b bool)`

 SetResumeNil sets the value for Resume to be an explicit nil

### UnsetResume
`func (o *NegotiationsCollectionNegotiationsItemsInner) UnsetResume()`

UnsetResume ensures that no value is present for Resume, not even an explicit nil
### GetUrl

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NegotiationsCollectionNegotiationsItemsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NegotiationsCollectionNegotiationsItemsInner) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


