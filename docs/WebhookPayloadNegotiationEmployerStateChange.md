# WebhookPayloadNegotiationEmployerStateChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerManagerId** | **string** | Идентификатор менеджера, совершившего перевод | 
**FromState** | **string** | С какого статуса перевели | 
**ResumeId** | **string** | Идентификатор резюме | 
**ToState** | **string** | На какой статус перевели | 
**TopicId** | **string** | Идентификатор топика | 
**TransferredAt** | **string** | Время перевода на новый этап | 
**VacancyId** | **string** | Идентификатор вакансии | 

## Methods

### NewWebhookPayloadNegotiationEmployerStateChange

`func NewWebhookPayloadNegotiationEmployerStateChange(employerManagerId string, fromState string, resumeId string, toState string, topicId string, transferredAt string, vacancyId string, ) *WebhookPayloadNegotiationEmployerStateChange`

NewWebhookPayloadNegotiationEmployerStateChange instantiates a new WebhookPayloadNegotiationEmployerStateChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadNegotiationEmployerStateChangeWithDefaults

`func NewWebhookPayloadNegotiationEmployerStateChangeWithDefaults() *WebhookPayloadNegotiationEmployerStateChange`

NewWebhookPayloadNegotiationEmployerStateChangeWithDefaults instantiates a new WebhookPayloadNegotiationEmployerStateChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerManagerId

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetEmployerManagerId() string`

GetEmployerManagerId returns the EmployerManagerId field if non-nil, zero value otherwise.

### GetEmployerManagerIdOk

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetEmployerManagerIdOk() (*string, bool)`

GetEmployerManagerIdOk returns a tuple with the EmployerManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerManagerId

`func (o *WebhookPayloadNegotiationEmployerStateChange) SetEmployerManagerId(v string)`

SetEmployerManagerId sets EmployerManagerId field to given value.


### GetFromState

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetFromState() string`

GetFromState returns the FromState field if non-nil, zero value otherwise.

### GetFromStateOk

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetFromStateOk() (*string, bool)`

GetFromStateOk returns a tuple with the FromState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromState

`func (o *WebhookPayloadNegotiationEmployerStateChange) SetFromState(v string)`

SetFromState sets FromState field to given value.


### GetResumeId

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetResumeId() string`

GetResumeId returns the ResumeId field if non-nil, zero value otherwise.

### GetResumeIdOk

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetResumeIdOk() (*string, bool)`

GetResumeIdOk returns a tuple with the ResumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeId

`func (o *WebhookPayloadNegotiationEmployerStateChange) SetResumeId(v string)`

SetResumeId sets ResumeId field to given value.


### GetToState

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetToState() string`

GetToState returns the ToState field if non-nil, zero value otherwise.

### GetToStateOk

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetToStateOk() (*string, bool)`

GetToStateOk returns a tuple with the ToState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToState

`func (o *WebhookPayloadNegotiationEmployerStateChange) SetToState(v string)`

SetToState sets ToState field to given value.


### GetTopicId

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetTopicId() string`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetTopicIdOk() (*string, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *WebhookPayloadNegotiationEmployerStateChange) SetTopicId(v string)`

SetTopicId sets TopicId field to given value.


### GetTransferredAt

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetTransferredAt() string`

GetTransferredAt returns the TransferredAt field if non-nil, zero value otherwise.

### GetTransferredAtOk

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetTransferredAtOk() (*string, bool)`

GetTransferredAtOk returns a tuple with the TransferredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferredAt

`func (o *WebhookPayloadNegotiationEmployerStateChange) SetTransferredAt(v string)`

SetTransferredAt sets TransferredAt field to given value.


### GetVacancyId

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetVacancyId() string`

GetVacancyId returns the VacancyId field if non-nil, zero value otherwise.

### GetVacancyIdOk

`func (o *WebhookPayloadNegotiationEmployerStateChange) GetVacancyIdOk() (*string, bool)`

GetVacancyIdOk returns a tuple with the VacancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyId

`func (o *WebhookPayloadNegotiationEmployerStateChange) SetVacancyId(v string)`

SetVacancyId sets VacancyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


