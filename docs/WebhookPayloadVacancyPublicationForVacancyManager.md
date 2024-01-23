# WebhookPayloadVacancyPublicationForVacancyManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationDate** | **string** | Дата создания вакансии в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm | 
**EmployerManagerId** | **string** | Идентификатор модератора вакансии | 
**VacancyId** | **string** | Идентификатор вакансии | 

## Methods

### NewWebhookPayloadVacancyPublicationForVacancyManager

`func NewWebhookPayloadVacancyPublicationForVacancyManager(creationDate string, employerManagerId string, vacancyId string, ) *WebhookPayloadVacancyPublicationForVacancyManager`

NewWebhookPayloadVacancyPublicationForVacancyManager instantiates a new WebhookPayloadVacancyPublicationForVacancyManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadVacancyPublicationForVacancyManagerWithDefaults

`func NewWebhookPayloadVacancyPublicationForVacancyManagerWithDefaults() *WebhookPayloadVacancyPublicationForVacancyManager`

NewWebhookPayloadVacancyPublicationForVacancyManagerWithDefaults instantiates a new WebhookPayloadVacancyPublicationForVacancyManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationDate

`func (o *WebhookPayloadVacancyPublicationForVacancyManager) GetCreationDate() string`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *WebhookPayloadVacancyPublicationForVacancyManager) GetCreationDateOk() (*string, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *WebhookPayloadVacancyPublicationForVacancyManager) SetCreationDate(v string)`

SetCreationDate sets CreationDate field to given value.


### GetEmployerManagerId

`func (o *WebhookPayloadVacancyPublicationForVacancyManager) GetEmployerManagerId() string`

GetEmployerManagerId returns the EmployerManagerId field if non-nil, zero value otherwise.

### GetEmployerManagerIdOk

`func (o *WebhookPayloadVacancyPublicationForVacancyManager) GetEmployerManagerIdOk() (*string, bool)`

GetEmployerManagerIdOk returns a tuple with the EmployerManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerManagerId

`func (o *WebhookPayloadVacancyPublicationForVacancyManager) SetEmployerManagerId(v string)`

SetEmployerManagerId sets EmployerManagerId field to given value.


### GetVacancyId

`func (o *WebhookPayloadVacancyPublicationForVacancyManager) GetVacancyId() string`

GetVacancyId returns the VacancyId field if non-nil, zero value otherwise.

### GetVacancyIdOk

`func (o *WebhookPayloadVacancyPublicationForVacancyManager) GetVacancyIdOk() (*string, bool)`

GetVacancyIdOk returns a tuple with the VacancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyId

`func (o *WebhookPayloadVacancyPublicationForVacancyManager) SetVacancyId(v string)`

SetVacancyId sets VacancyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


