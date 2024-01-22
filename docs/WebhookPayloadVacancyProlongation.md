# WebhookPayloadVacancyProlongation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProlongationDate** | **string** | Дата продления вакансии в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**VacancyId** | **string** | Идентификатор вакансии | 

## Methods

### NewWebhookPayloadVacancyProlongation

`func NewWebhookPayloadVacancyProlongation(prolongationDate string, vacancyId string, ) *WebhookPayloadVacancyProlongation`

NewWebhookPayloadVacancyProlongation instantiates a new WebhookPayloadVacancyProlongation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadVacancyProlongationWithDefaults

`func NewWebhookPayloadVacancyProlongationWithDefaults() *WebhookPayloadVacancyProlongation`

NewWebhookPayloadVacancyProlongationWithDefaults instantiates a new WebhookPayloadVacancyProlongation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProlongationDate

`func (o *WebhookPayloadVacancyProlongation) GetProlongationDate() string`

GetProlongationDate returns the ProlongationDate field if non-nil, zero value otherwise.

### GetProlongationDateOk

`func (o *WebhookPayloadVacancyProlongation) GetProlongationDateOk() (*string, bool)`

GetProlongationDateOk returns a tuple with the ProlongationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProlongationDate

`func (o *WebhookPayloadVacancyProlongation) SetProlongationDate(v string)`

SetProlongationDate sets ProlongationDate field to given value.


### GetVacancyId

`func (o *WebhookPayloadVacancyProlongation) GetVacancyId() string`

GetVacancyId returns the VacancyId field if non-nil, zero value otherwise.

### GetVacancyIdOk

`func (o *WebhookPayloadVacancyProlongation) GetVacancyIdOk() (*string, bool)`

GetVacancyIdOk returns a tuple with the VacancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyId

`func (o *WebhookPayloadVacancyProlongation) SetVacancyId(v string)`

SetVacancyId sets VacancyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


