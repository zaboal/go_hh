# WebhookPayloadNewNegotiationVacancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerId** | **string** | Идентификатор работодателя | 
**NegotiationDate** | **interface{}** | Дата отклика в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**ResumeId** | **NullableString** | Идентификатор резюме | 
**TopicId** | **string** | Идентификатор топика | 
**VacancyId** | **NullableString** | Идентификатор вакансии | 

## Methods

### NewWebhookPayloadNewNegotiationVacancy

`func NewWebhookPayloadNewNegotiationVacancy(employerId string, negotiationDate interface{}, resumeId NullableString, topicId string, vacancyId NullableString, ) *WebhookPayloadNewNegotiationVacancy`

NewWebhookPayloadNewNegotiationVacancy instantiates a new WebhookPayloadNewNegotiationVacancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadNewNegotiationVacancyWithDefaults

`func NewWebhookPayloadNewNegotiationVacancyWithDefaults() *WebhookPayloadNewNegotiationVacancy`

NewWebhookPayloadNewNegotiationVacancyWithDefaults instantiates a new WebhookPayloadNewNegotiationVacancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerId

`func (o *WebhookPayloadNewNegotiationVacancy) GetEmployerId() string`

GetEmployerId returns the EmployerId field if non-nil, zero value otherwise.

### GetEmployerIdOk

`func (o *WebhookPayloadNewNegotiationVacancy) GetEmployerIdOk() (*string, bool)`

GetEmployerIdOk returns a tuple with the EmployerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerId

`func (o *WebhookPayloadNewNegotiationVacancy) SetEmployerId(v string)`

SetEmployerId sets EmployerId field to given value.


### GetNegotiationDate

`func (o *WebhookPayloadNewNegotiationVacancy) GetNegotiationDate() interface{}`

GetNegotiationDate returns the NegotiationDate field if non-nil, zero value otherwise.

### GetNegotiationDateOk

`func (o *WebhookPayloadNewNegotiationVacancy) GetNegotiationDateOk() (*interface{}, bool)`

GetNegotiationDateOk returns a tuple with the NegotiationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationDate

`func (o *WebhookPayloadNewNegotiationVacancy) SetNegotiationDate(v interface{})`

SetNegotiationDate sets NegotiationDate field to given value.


### SetNegotiationDateNil

`func (o *WebhookPayloadNewNegotiationVacancy) SetNegotiationDateNil(b bool)`

 SetNegotiationDateNil sets the value for NegotiationDate to be an explicit nil

### UnsetNegotiationDate
`func (o *WebhookPayloadNewNegotiationVacancy) UnsetNegotiationDate()`

UnsetNegotiationDate ensures that no value is present for NegotiationDate, not even an explicit nil
### GetResumeId

`func (o *WebhookPayloadNewNegotiationVacancy) GetResumeId() string`

GetResumeId returns the ResumeId field if non-nil, zero value otherwise.

### GetResumeIdOk

`func (o *WebhookPayloadNewNegotiationVacancy) GetResumeIdOk() (*string, bool)`

GetResumeIdOk returns a tuple with the ResumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeId

`func (o *WebhookPayloadNewNegotiationVacancy) SetResumeId(v string)`

SetResumeId sets ResumeId field to given value.


### SetResumeIdNil

`func (o *WebhookPayloadNewNegotiationVacancy) SetResumeIdNil(b bool)`

 SetResumeIdNil sets the value for ResumeId to be an explicit nil

### UnsetResumeId
`func (o *WebhookPayloadNewNegotiationVacancy) UnsetResumeId()`

UnsetResumeId ensures that no value is present for ResumeId, not even an explicit nil
### GetTopicId

`func (o *WebhookPayloadNewNegotiationVacancy) GetTopicId() string`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *WebhookPayloadNewNegotiationVacancy) GetTopicIdOk() (*string, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *WebhookPayloadNewNegotiationVacancy) SetTopicId(v string)`

SetTopicId sets TopicId field to given value.


### GetVacancyId

`func (o *WebhookPayloadNewNegotiationVacancy) GetVacancyId() string`

GetVacancyId returns the VacancyId field if non-nil, zero value otherwise.

### GetVacancyIdOk

`func (o *WebhookPayloadNewNegotiationVacancy) GetVacancyIdOk() (*string, bool)`

GetVacancyIdOk returns a tuple with the VacancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyId

`func (o *WebhookPayloadNewNegotiationVacancy) SetVacancyId(v string)`

SetVacancyId sets VacancyId field to given value.


### SetVacancyIdNil

`func (o *WebhookPayloadNewNegotiationVacancy) SetVacancyIdNil(b bool)`

 SetVacancyIdNil sets the value for VacancyId to be an explicit nil

### UnsetVacancyId
`func (o *WebhookPayloadNewNegotiationVacancy) UnsetVacancyId()`

UnsetVacancyId ensures that no value is present for VacancyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


