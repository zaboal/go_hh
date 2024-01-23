# WebhookPayloadNewResponseOrInvitationVacancy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmployerId** | **string** | Идентификатор работодателя | 
**ResponseDate** | **string** | Дата отклика или приглашения в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**ResumeId** | **NullableString** | Идентификатор резюме | 
**TopicId** | **string** | Идентификатор топика | 
**VacancyId** | **NullableString** | Идентификатор вакансии | 

## Methods

### NewWebhookPayloadNewResponseOrInvitationVacancy

`func NewWebhookPayloadNewResponseOrInvitationVacancy(employerId string, responseDate string, resumeId NullableString, topicId string, vacancyId NullableString, ) *WebhookPayloadNewResponseOrInvitationVacancy`

NewWebhookPayloadNewResponseOrInvitationVacancy instantiates a new WebhookPayloadNewResponseOrInvitationVacancy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadNewResponseOrInvitationVacancyWithDefaults

`func NewWebhookPayloadNewResponseOrInvitationVacancyWithDefaults() *WebhookPayloadNewResponseOrInvitationVacancy`

NewWebhookPayloadNewResponseOrInvitationVacancyWithDefaults instantiates a new WebhookPayloadNewResponseOrInvitationVacancy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmployerId

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetEmployerId() string`

GetEmployerId returns the EmployerId field if non-nil, zero value otherwise.

### GetEmployerIdOk

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetEmployerIdOk() (*string, bool)`

GetEmployerIdOk returns a tuple with the EmployerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerId

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetEmployerId(v string)`

SetEmployerId sets EmployerId field to given value.


### GetResponseDate

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetResponseDate() string`

GetResponseDate returns the ResponseDate field if non-nil, zero value otherwise.

### GetResponseDateOk

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetResponseDateOk() (*string, bool)`

GetResponseDateOk returns a tuple with the ResponseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseDate

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetResponseDate(v string)`

SetResponseDate sets ResponseDate field to given value.


### GetResumeId

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetResumeId() string`

GetResumeId returns the ResumeId field if non-nil, zero value otherwise.

### GetResumeIdOk

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetResumeIdOk() (*string, bool)`

GetResumeIdOk returns a tuple with the ResumeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeId

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetResumeId(v string)`

SetResumeId sets ResumeId field to given value.


### SetResumeIdNil

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetResumeIdNil(b bool)`

 SetResumeIdNil sets the value for ResumeId to be an explicit nil

### UnsetResumeId
`func (o *WebhookPayloadNewResponseOrInvitationVacancy) UnsetResumeId()`

UnsetResumeId ensures that no value is present for ResumeId, not even an explicit nil
### GetTopicId

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetTopicId() string`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetTopicIdOk() (*string, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetTopicId(v string)`

SetTopicId sets TopicId field to given value.


### GetVacancyId

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetVacancyId() string`

GetVacancyId returns the VacancyId field if non-nil, zero value otherwise.

### GetVacancyIdOk

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) GetVacancyIdOk() (*string, bool)`

GetVacancyIdOk returns a tuple with the VacancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyId

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetVacancyId(v string)`

SetVacancyId sets VacancyId field to given value.


### SetVacancyIdNil

`func (o *WebhookPayloadNewResponseOrInvitationVacancy) SetVacancyIdNil(b bool)`

 SetVacancyIdNil sets the value for VacancyId to be an explicit nil

### UnsetVacancyId
`func (o *WebhookPayloadNewResponseOrInvitationVacancy) UnsetVacancyId()`

UnsetVacancyId ensures that no value is present for VacancyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


