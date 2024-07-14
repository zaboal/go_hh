# NegotiationsObjectsTopicItemCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicantQuestionState** | Pointer to **bool** | Является ли вопросом до отклика | [optional] 
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

## Methods

### NewNegotiationsObjectsTopicItemCommon

`func NewNegotiationsObjectsTopicItemCommon(createdAt string, hasUpdates bool, id string, messagingStatus string, state IncludesIdName, updatedAt string, viewedByOpponent bool, ) *NegotiationsObjectsTopicItemCommon`

NewNegotiationsObjectsTopicItemCommon instantiates a new NegotiationsObjectsTopicItemCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNegotiationsObjectsTopicItemCommonWithDefaults

`func NewNegotiationsObjectsTopicItemCommonWithDefaults() *NegotiationsObjectsTopicItemCommon`

NewNegotiationsObjectsTopicItemCommonWithDefaults instantiates a new NegotiationsObjectsTopicItemCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicantQuestionState

`func (o *NegotiationsObjectsTopicItemCommon) GetApplicantQuestionState() bool`

GetApplicantQuestionState returns the ApplicantQuestionState field if non-nil, zero value otherwise.

### GetApplicantQuestionStateOk

`func (o *NegotiationsObjectsTopicItemCommon) GetApplicantQuestionStateOk() (*bool, bool)`

GetApplicantQuestionStateOk returns a tuple with the ApplicantQuestionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicantQuestionState

`func (o *NegotiationsObjectsTopicItemCommon) SetApplicantQuestionState(v bool)`

SetApplicantQuestionState sets ApplicantQuestionState field to given value.

### HasApplicantQuestionState

`func (o *NegotiationsObjectsTopicItemCommon) HasApplicantQuestionState() bool`

HasApplicantQuestionState returns a boolean if a field has been set.

### GetCounters

`func (o *NegotiationsObjectsTopicItemCommon) GetCounters() NegotiationsObjectsCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *NegotiationsObjectsTopicItemCommon) GetCountersOk() (*NegotiationsObjectsCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *NegotiationsObjectsTopicItemCommon) SetCounters(v NegotiationsObjectsCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *NegotiationsObjectsTopicItemCommon) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NegotiationsObjectsTopicItemCommon) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NegotiationsObjectsTopicItemCommon) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NegotiationsObjectsTopicItemCommon) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetHasUpdates

`func (o *NegotiationsObjectsTopicItemCommon) GetHasUpdates() bool`

GetHasUpdates returns the HasUpdates field if non-nil, zero value otherwise.

### GetHasUpdatesOk

`func (o *NegotiationsObjectsTopicItemCommon) GetHasUpdatesOk() (*bool, bool)`

GetHasUpdatesOk returns a tuple with the HasUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasUpdates

`func (o *NegotiationsObjectsTopicItemCommon) SetHasUpdates(v bool)`

SetHasUpdates sets HasUpdates field to given value.


### GetId

`func (o *NegotiationsObjectsTopicItemCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NegotiationsObjectsTopicItemCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NegotiationsObjectsTopicItemCommon) SetId(v string)`

SetId sets Id field to given value.


### GetMessagesUrl

`func (o *NegotiationsObjectsTopicItemCommon) GetMessagesUrl() string`

GetMessagesUrl returns the MessagesUrl field if non-nil, zero value otherwise.

### GetMessagesUrlOk

`func (o *NegotiationsObjectsTopicItemCommon) GetMessagesUrlOk() (*string, bool)`

GetMessagesUrlOk returns a tuple with the MessagesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesUrl

`func (o *NegotiationsObjectsTopicItemCommon) SetMessagesUrl(v string)`

SetMessagesUrl sets MessagesUrl field to given value.

### HasMessagesUrl

`func (o *NegotiationsObjectsTopicItemCommon) HasMessagesUrl() bool`

HasMessagesUrl returns a boolean if a field has been set.

### GetMessagingStatus

`func (o *NegotiationsObjectsTopicItemCommon) GetMessagingStatus() string`

GetMessagingStatus returns the MessagingStatus field if non-nil, zero value otherwise.

### GetMessagingStatusOk

`func (o *NegotiationsObjectsTopicItemCommon) GetMessagingStatusOk() (*string, bool)`

GetMessagingStatusOk returns a tuple with the MessagingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingStatus

`func (o *NegotiationsObjectsTopicItemCommon) SetMessagingStatus(v string)`

SetMessagingStatus sets MessagingStatus field to given value.


### GetProfessionalRoles

`func (o *NegotiationsObjectsTopicItemCommon) GetProfessionalRoles() []VacancyProfessionalRoleItem`

GetProfessionalRoles returns the ProfessionalRoles field if non-nil, zero value otherwise.

### GetProfessionalRolesOk

`func (o *NegotiationsObjectsTopicItemCommon) GetProfessionalRolesOk() (*[]VacancyProfessionalRoleItem, bool)`

GetProfessionalRolesOk returns a tuple with the ProfessionalRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfessionalRoles

`func (o *NegotiationsObjectsTopicItemCommon) SetProfessionalRoles(v []VacancyProfessionalRoleItem)`

SetProfessionalRoles sets ProfessionalRoles field to given value.

### HasProfessionalRoles

`func (o *NegotiationsObjectsTopicItemCommon) HasProfessionalRoles() bool`

HasProfessionalRoles returns a boolean if a field has been set.

### GetSource

`func (o *NegotiationsObjectsTopicItemCommon) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *NegotiationsObjectsTopicItemCommon) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *NegotiationsObjectsTopicItemCommon) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *NegotiationsObjectsTopicItemCommon) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetState

`func (o *NegotiationsObjectsTopicItemCommon) GetState() IncludesIdName`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NegotiationsObjectsTopicItemCommon) GetStateOk() (*IncludesIdName, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NegotiationsObjectsTopicItemCommon) SetState(v IncludesIdName)`

SetState sets State field to given value.


### GetUpdatedAt

`func (o *NegotiationsObjectsTopicItemCommon) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NegotiationsObjectsTopicItemCommon) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NegotiationsObjectsTopicItemCommon) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetViewedByOpponent

`func (o *NegotiationsObjectsTopicItemCommon) GetViewedByOpponent() bool`

GetViewedByOpponent returns the ViewedByOpponent field if non-nil, zero value otherwise.

### GetViewedByOpponentOk

`func (o *NegotiationsObjectsTopicItemCommon) GetViewedByOpponentOk() (*bool, bool)`

GetViewedByOpponentOk returns a tuple with the ViewedByOpponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewedByOpponent

`func (o *NegotiationsObjectsTopicItemCommon) SetViewedByOpponent(v bool)`

SetViewedByOpponent sets ViewedByOpponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


