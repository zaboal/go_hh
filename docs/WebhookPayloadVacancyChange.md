# WebhookPayloadVacancyChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeDate** | **string** | Дата изменения вакансии в формате [ISO 8601](http://en.wikipedia.org/wiki/ISO_8601) с точностью до секунды: &#x60;YYYY-MM-DDThh:mm:ss±hhmm&#x60; | 
**VacancyId** | **string** | Идентификатор вакансии | 

## Methods

### NewWebhookPayloadVacancyChange

`func NewWebhookPayloadVacancyChange(changeDate string, vacancyId string, ) *WebhookPayloadVacancyChange`

NewWebhookPayloadVacancyChange instantiates a new WebhookPayloadVacancyChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadVacancyChangeWithDefaults

`func NewWebhookPayloadVacancyChangeWithDefaults() *WebhookPayloadVacancyChange`

NewWebhookPayloadVacancyChangeWithDefaults instantiates a new WebhookPayloadVacancyChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeDate

`func (o *WebhookPayloadVacancyChange) GetChangeDate() string`

GetChangeDate returns the ChangeDate field if non-nil, zero value otherwise.

### GetChangeDateOk

`func (o *WebhookPayloadVacancyChange) GetChangeDateOk() (*string, bool)`

GetChangeDateOk returns a tuple with the ChangeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeDate

`func (o *WebhookPayloadVacancyChange) SetChangeDate(v string)`

SetChangeDate sets ChangeDate field to given value.


### GetVacancyId

`func (o *WebhookPayloadVacancyChange) GetVacancyId() string`

GetVacancyId returns the VacancyId field if non-nil, zero value otherwise.

### GetVacancyIdOk

`func (o *WebhookPayloadVacancyChange) GetVacancyIdOk() (*string, bool)`

GetVacancyIdOk returns a tuple with the VacancyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVacancyId

`func (o *WebhookPayloadVacancyChange) SetVacancyId(v string)`

SetVacancyId sets VacancyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


