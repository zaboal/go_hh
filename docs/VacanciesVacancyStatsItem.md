# VacanciesVacancyStatsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** | Дата в формате &#x60;YYYY-MM-DD&#x60; | 
**Responses** | Pointer to **NullableFloat32** | Количество откликов на вакансию. &#x60;null&#x60; если дата в будущем или данных на эту дату нет | [optional] 
**Views** | Pointer to **NullableFloat32** | Количество просмотров вакансии. &#x60;null&#x60; если дата в будущем или данных на эту дату нет | [optional] 

## Methods

### NewVacanciesVacancyStatsItem

`func NewVacanciesVacancyStatsItem(date string, ) *VacanciesVacancyStatsItem`

NewVacanciesVacancyStatsItem instantiates a new VacanciesVacancyStatsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyStatsItemWithDefaults

`func NewVacanciesVacancyStatsItemWithDefaults() *VacanciesVacancyStatsItem`

NewVacanciesVacancyStatsItemWithDefaults instantiates a new VacanciesVacancyStatsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *VacanciesVacancyStatsItem) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *VacanciesVacancyStatsItem) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *VacanciesVacancyStatsItem) SetDate(v string)`

SetDate sets Date field to given value.


### GetResponses

`func (o *VacanciesVacancyStatsItem) GetResponses() float32`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *VacanciesVacancyStatsItem) GetResponsesOk() (*float32, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *VacanciesVacancyStatsItem) SetResponses(v float32)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *VacanciesVacancyStatsItem) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### SetResponsesNil

`func (o *VacanciesVacancyStatsItem) SetResponsesNil(b bool)`

 SetResponsesNil sets the value for Responses to be an explicit nil

### UnsetResponses
`func (o *VacanciesVacancyStatsItem) UnsetResponses()`

UnsetResponses ensures that no value is present for Responses, not even an explicit nil
### GetViews

`func (o *VacanciesVacancyStatsItem) GetViews() float32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *VacanciesVacancyStatsItem) GetViewsOk() (*float32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *VacanciesVacancyStatsItem) SetViews(v float32)`

SetViews sets Views field to given value.

### HasViews

`func (o *VacanciesVacancyStatsItem) HasViews() bool`

HasViews returns a boolean if a field has been set.

### SetViewsNil

`func (o *VacanciesVacancyStatsItem) SetViewsNil(b bool)`

 SetViewsNil sets the value for Views to be an explicit nil

### UnsetViews
`func (o *VacanciesVacancyStatsItem) UnsetViews()`

UnsetViews ensures that no value is present for Views, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


