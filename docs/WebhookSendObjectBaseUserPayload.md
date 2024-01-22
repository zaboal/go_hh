# WebhookSendObjectBaseUserPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerId** | **string** | Идентификатор работодателя | 
**NegotiationDate** | **map[string]interface{}** | Дата отклика в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**ResumeId** | **string** | Идентификатор резюме | 
**TopicId** | **string** | Идентификатор топика | 
**VacancyId** | **string** | Идентификатор вакансии | 
**ResponseDate** | **string** | Дата отклика или приглашения в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**ArchivationDate** | **string** | Дата архивации вакансии в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**CreationDate** | **string** | Дата создания вакансии в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm | 
**EmployerManagerId** | **string** | Идентификатор менеджера, совершившего перевод | 
**ProlongationDate** | **string** | Дата продления вакансии в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**ChangeDate** | **string** | Дата изменения вакансии в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**FromState** | **string** | С какого статуса перевели | 
**ToState** | **string** | На какой статус перевели | 
**TransferredAt** | **string** | Время перевода на новый этап | 

## Methods

### NewWebhookSendObjectBaseUserPayload

`func NewWebhookSendObjectBaseUserPayload(employerId string, negotiationDate map[string]interface{}, resumeId string, topicId string, vacancyId string, responseDate string, archivationDate string, creationDate string, employerManagerId string, prolongationDate string, changeDate string, fromState string, toState string, transferredAt string, ) *WebhookSendObjectBaseUserPayload`

NewWebhookSendObjectBaseUserPayload instantiates a new WebhookSendObjectBaseUserPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookSendObjectBaseUserPayloadWithDefaults

`func NewWebhookSendObjectBaseUserPayloadWithDefaults() *WebhookSendObjectBaseUserPayload`

NewWebhookSendObjectBaseUserPayloadWithDefaults instantiates a new WebhookSendObjectBaseUserPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerId

`func (o *WebhookSendObjectBaseUserPayload) GetEmployerId() string`

GetEmployerId returns the EmployerId field if non-nil, zero value otherwise.

### GetEmployerIdOk

`func (o *WebhookSendObjectBaseUserPayload) GetEmployerIdOk() (*string, bool)`

GetEmployerIdOk returns a tuple with the EmployerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerId

`func (o *WebhookSendObjectBaseUserPayload) SetEmployerId(v string)`

SetEmployerId sets EmployerId field to given value.


### GetNegotiationDate

`func (o *WebhookSendObjectBaseUserPayload) GetNegotiationDate() map[string]interface{}`

GetNegotiationDate returns the NegotiationDate field if non-nil, zero value otherwise.

### GetNegotiationDateOk

`func (o *WebhookSendObjectBaseUserPayload) GetNegotiationDateOk() (*map[string]interface{}, bool)`

GetNegotiationDateOk returns a tuple with the NegotiationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationDate

`func (o *WebhookSendObjectBaseUserPayload) SetNegotiationDate(v map[string]interface{})`

SetNegotiationDate sets NegotiationDate field to given value.


### GetResumeId

`func (o *WebhookSendObjectBaseUserPayload) GetResumeId() string`

GetResumeId returns the ResumeId field if non-nil, zero value otherwise.

### GetResumeIdOk

`func (o *WebhookSendObjectBaseUserPayload) GetResumeIdOk() (*string, bool)`

GetResumeIdOk returns a tuple with the ResumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeId

`func (o *WebhookSendObjectBaseUserPayload) SetResumeId(v string)`

SetResumeId sets ResumeId field to given value.


### GetTopicId

`func (o *WebhookSendObjectBaseUserPayload) GetTopicId() string`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *WebhookSendObjectBaseUserPayload) GetTopicIdOk() (*string, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *WebhookSendObjectBaseUserPayload) SetTopicId(v string)`

SetTopicId sets TopicId field to given value.


### GetVacancyId

`func (o *WebhookSendObjectBaseUserPayload) GetVacancyId() string`

GetVacancyId returns the VacancyId field if non-nil, zero value otherwise.

### GetVacancyIdOk

`func (o *WebhookSendObjectBaseUserPayload) GetVacancyIdOk() (*string, bool)`

GetVacancyIdOk returns a tuple with the VacancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyId

`func (o *WebhookSendObjectBaseUserPayload) SetVacancyId(v string)`

SetVacancyId sets VacancyId field to given value.


### GetResponseDate

`func (o *WebhookSendObjectBaseUserPayload) GetResponseDate() string`

GetResponseDate returns the ResponseDate field if non-nil, zero value otherwise.

### GetResponseDateOk

`func (o *WebhookSendObjectBaseUserPayload) GetResponseDateOk() (*string, bool)`

GetResponseDateOk returns a tuple with the ResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDate

`func (o *WebhookSendObjectBaseUserPayload) SetResponseDate(v string)`

SetResponseDate sets ResponseDate field to given value.


### GetArchivationDate

`func (o *WebhookSendObjectBaseUserPayload) GetArchivationDate() string`

GetArchivationDate returns the ArchivationDate field if non-nil, zero value otherwise.

### GetArchivationDateOk

`func (o *WebhookSendObjectBaseUserPayload) GetArchivationDateOk() (*string, bool)`

GetArchivationDateOk returns a tuple with the ArchivationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivationDate

`func (o *WebhookSendObjectBaseUserPayload) SetArchivationDate(v string)`

SetArchivationDate sets ArchivationDate field to given value.


### GetCreationDate

`func (o *WebhookSendObjectBaseUserPayload) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *WebhookSendObjectBaseUserPayload) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *WebhookSendObjectBaseUserPayload) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.


### GetEmployerManagerId

`func (o *WebhookSendObjectBaseUserPayload) GetEmployerManagerId() string`

GetEmployerManagerId returns the EmployerManagerId field if non-nil, zero value otherwise.

### GetEmployerManagerIdOk

`func (o *WebhookSendObjectBaseUserPayload) GetEmployerManagerIdOk() (*string, bool)`

GetEmployerManagerIdOk returns a tuple with the EmployerManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerManagerId

`func (o *WebhookSendObjectBaseUserPayload) SetEmployerManagerId(v string)`

SetEmployerManagerId sets EmployerManagerId field to given value.


### GetProlongationDate

`func (o *WebhookSendObjectBaseUserPayload) GetProlongationDate() string`

GetProlongationDate returns the ProlongationDate field if non-nil, zero value otherwise.

### GetProlongationDateOk

`func (o *WebhookSendObjectBaseUserPayload) GetProlongationDateOk() (*string, bool)`

GetProlongationDateOk returns a tuple with the ProlongationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongationDate

`func (o *WebhookSendObjectBaseUserPayload) SetProlongationDate(v string)`

SetProlongationDate sets ProlongationDate field to given value.


### GetChangeDate

`func (o *WebhookSendObjectBaseUserPayload) GetChangeDate() string`

GetChangeDate returns the ChangeDate field if non-nil, zero value otherwise.

### GetChangeDateOk

`func (o *WebhookSendObjectBaseUserPayload) GetChangeDateOk() (*string, bool)`

GetChangeDateOk returns a tuple with the ChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDate

`func (o *WebhookSendObjectBaseUserPayload) SetChangeDate(v string)`

SetChangeDate sets ChangeDate field to given value.


### GetFromState

`func (o *WebhookSendObjectBaseUserPayload) GetFromState() string`

GetFromState returns the FromState field if non-nil, zero value otherwise.

### GetFromStateOk

`func (o *WebhookSendObjectBaseUserPayload) GetFromStateOk() (*string, bool)`

GetFromStateOk returns a tuple with the FromState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromState

`func (o *WebhookSendObjectBaseUserPayload) SetFromState(v string)`

SetFromState sets FromState field to given value.


### GetToState

`func (o *WebhookSendObjectBaseUserPayload) GetToState() string`

GetToState returns the ToState field if non-nil, zero value otherwise.

### GetToStateOk

`func (o *WebhookSendObjectBaseUserPayload) GetToStateOk() (*string, bool)`

GetToStateOk returns a tuple with the ToState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToState

`func (o *WebhookSendObjectBaseUserPayload) SetToState(v string)`

SetToState sets ToState field to given value.


### GetTransferredAt

`func (o *WebhookSendObjectBaseUserPayload) GetTransferredAt() string`

GetTransferredAt returns the TransferredAt field if non-nil, zero value otherwise.

### GetTransferredAtOk

`func (o *WebhookSendObjectBaseUserPayload) GetTransferredAtOk() (*string, bool)`

GetTransferredAtOk returns a tuple with the TransferredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferredAt

`func (o *WebhookSendObjectBaseUserPayload) SetTransferredAt(v string)`

SetTransferredAt sets TransferredAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


