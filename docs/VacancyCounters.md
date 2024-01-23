# VacancyCounters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Responses** | Pointer to **float32** | Количество откликов на вакансию с момента публикации | [optional] 
**TotalResponses** | Pointer to **float32** | Количество откликов на вакансию с момента создания | [optional] 

## Methods

### NewVacancyCounters

`func NewVacancyCounters() *VacancyCounters`

NewVacancyCounters instantiates a new VacancyCounters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyCountersWithDefaults

`func NewVacancyCountersWithDefaults() *VacancyCounters`

NewVacancyCountersWithDefaults instantiates a new VacancyCounters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResponses

`func (o *VacancyCounters) GetResponses() float32`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *VacancyCounters) GetResponsesOk() (*float32, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *VacancyCounters) SetResponses(v float32)`

SetResponses sets Responses field to given value.

### HasResponses

`func (o *VacancyCounters) HasResponses() bool`

HasResponses returns a boolean if a field has been set.

### GetTotalResponses

`func (o *VacancyCounters) GetTotalResponses() float32`

GetTotalResponses returns the TotalResponses field if non-nil, zero value otherwise.

### GetTotalResponsesOk

`func (o *VacancyCounters) GetTotalResponsesOk() (*float32, bool)`

GetTotalResponsesOk returns a tuple with the TotalResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResponses

`func (o *VacancyCounters) SetTotalResponses(v float32)`

SetTotalResponses sets TotalResponses field to given value.

### HasTotalResponses

`func (o *VacancyCounters) HasTotalResponses() bool`

HasTotalResponses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


