# WebhookPayloadVacancyArchivation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivationDate** | **string** | Дата архивации вакансии в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**VacancyId** | **NullableString** | Идентификатор вакансии | 

## Methods

### NewWebhookPayloadVacancyArchivation

`func NewWebhookPayloadVacancyArchivation(archivationDate string, vacancyId NullableString, ) *WebhookPayloadVacancyArchivation`

NewWebhookPayloadVacancyArchivation instantiates a new WebhookPayloadVacancyArchivation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadVacancyArchivationWithDefaults

`func NewWebhookPayloadVacancyArchivationWithDefaults() *WebhookPayloadVacancyArchivation`

NewWebhookPayloadVacancyArchivationWithDefaults instantiates a new WebhookPayloadVacancyArchivation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivationDate

`func (o *WebhookPayloadVacancyArchivation) GetArchivationDate() string`

GetArchivationDate returns the ArchivationDate field if non-nil, zero value otherwise.

### GetArchivationDateOk

`func (o *WebhookPayloadVacancyArchivation) GetArchivationDateOk() (*string, bool)`

GetArchivationDateOk returns a tuple with the ArchivationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivationDate

`func (o *WebhookPayloadVacancyArchivation) SetArchivationDate(v string)`

SetArchivationDate sets ArchivationDate field to given value.


### GetVacancyId

`func (o *WebhookPayloadVacancyArchivation) GetVacancyId() string`

GetVacancyId returns the VacancyId field if non-nil, zero value otherwise.

### GetVacancyIdOk

`func (o *WebhookPayloadVacancyArchivation) GetVacancyIdOk() (*string, bool)`

GetVacancyIdOk returns a tuple with the VacancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyId

`func (o *WebhookPayloadVacancyArchivation) SetVacancyId(v string)`

SetVacancyId sets VacancyId field to given value.


### SetVacancyIdNil

`func (o *WebhookPayloadVacancyArchivation) SetVacancyIdNil(b bool)`

 SetVacancyIdNil sets the value for VacancyId to be an explicit nil

### UnsetVacancyId
`func (o *WebhookPayloadVacancyArchivation) UnsetVacancyId()`

UnsetVacancyId ensures that no value is present for VacancyId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


