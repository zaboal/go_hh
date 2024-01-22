# ResumesByStatusCounters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyApplied** | **float32** | Количество резюме, уже использованных для отклика на данную вакансию | 
**NotPublished** | **float32** | Количество неопубликованных резюме (в [статусе](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume) &#x60;not_published&#x60; или &#x60;blocked&#x60;) | 
**Suitable** | **float32** | Количество резюме, которыми можно откликнуться на данную вакансию | 
**Unavailable** | **float32** | Количество резюме, которыми невозможно откликнуться на данную вакансию (например, из-за конфликтующих настроек видимости) | 

## Methods

### NewResumesByStatusCounters

`func NewResumesByStatusCounters(alreadyApplied float32, notPublished float32, suitable float32, unavailable float32, ) *ResumesByStatusCounters`

NewResumesByStatusCounters instantiates a new ResumesByStatusCounters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesByStatusCountersWithDefaults

`func NewResumesByStatusCountersWithDefaults() *ResumesByStatusCounters`

NewResumesByStatusCountersWithDefaults instantiates a new ResumesByStatusCounters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyApplied

`func (o *ResumesByStatusCounters) GetAlreadyApplied() float32`

GetAlreadyApplied returns the AlreadyApplied field if non-nil, zero value otherwise.

### GetAlreadyAppliedOk

`func (o *ResumesByStatusCounters) GetAlreadyAppliedOk() (*float32, bool)`

GetAlreadyAppliedOk returns a tuple with the AlreadyApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyApplied

`func (o *ResumesByStatusCounters) SetAlreadyApplied(v float32)`

SetAlreadyApplied sets AlreadyApplied field to given value.


### GetNotPublished

`func (o *ResumesByStatusCounters) GetNotPublished() float32`

GetNotPublished returns the NotPublished field if non-nil, zero value otherwise.

### GetNotPublishedOk

`func (o *ResumesByStatusCounters) GetNotPublishedOk() (*float32, bool)`

GetNotPublishedOk returns a tuple with the NotPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotPublished

`func (o *ResumesByStatusCounters) SetNotPublished(v float32)`

SetNotPublished sets NotPublished field to given value.


### GetSuitable

`func (o *ResumesByStatusCounters) GetSuitable() float32`

GetSuitable returns the Suitable field if non-nil, zero value otherwise.

### GetSuitableOk

`func (o *ResumesByStatusCounters) GetSuitableOk() (*float32, bool)`

GetSuitableOk returns a tuple with the Suitable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitable

`func (o *ResumesByStatusCounters) SetSuitable(v float32)`

SetSuitable sets Suitable field to given value.


### GetUnavailable

`func (o *ResumesByStatusCounters) GetUnavailable() float32`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *ResumesByStatusCounters) GetUnavailableOk() (*float32, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *ResumesByStatusCounters) SetUnavailable(v float32)`

SetUnavailable sets Unavailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


