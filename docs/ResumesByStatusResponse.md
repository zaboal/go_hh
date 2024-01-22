# ResumesByStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyApplied** | [**[]ResumesSuitableResumeItem**](ResumesSuitableResumeItem.md) | Список резюме, уже использованных для отклика на данную вакансию | 
**Counters** | Pointer to [**NullableResumesByStatusCounters**](ResumesByStatusCounters.md) |  | [optional] 
**NotPublished** | [**[]ResumesSuitableResumeItem**](ResumesSuitableResumeItem.md) | Список неопубликованных резюме (в [статусе](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume) &#x60;not_published&#x60; или &#x60;blocked&#x60;) | 
**Suitable** | [**[]ResumesSuitableResumeItem**](ResumesSuitableResumeItem.md) | Список резюме, которыми можно откликнуться на данную вакансию | 
**Unavailable** | [**[]ResumesSuitableResumeItem**](ResumesSuitableResumeItem.md) | Список резюме, которыми невозможно откликнуться на данную вакансию (например, из-за конфликтующих настроек видимости) | 

## Methods

### NewResumesByStatusResponse

`func NewResumesByStatusResponse(alreadyApplied []ResumesSuitableResumeItem, notPublished []ResumesSuitableResumeItem, suitable []ResumesSuitableResumeItem, unavailable []ResumesSuitableResumeItem, ) *ResumesByStatusResponse`

NewResumesByStatusResponse instantiates a new ResumesByStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesByStatusResponseWithDefaults

`func NewResumesByStatusResponseWithDefaults() *ResumesByStatusResponse`

NewResumesByStatusResponseWithDefaults instantiates a new ResumesByStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyApplied

`func (o *ResumesByStatusResponse) GetAlreadyApplied() []ResumesSuitableResumeItem`

GetAlreadyApplied returns the AlreadyApplied field if non-nil, zero value otherwise.

### GetAlreadyAppliedOk

`func (o *ResumesByStatusResponse) GetAlreadyAppliedOk() (*[]ResumesSuitableResumeItem, bool)`

GetAlreadyAppliedOk returns a tuple with the AlreadyApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyApplied

`func (o *ResumesByStatusResponse) SetAlreadyApplied(v []ResumesSuitableResumeItem)`

SetAlreadyApplied sets AlreadyApplied field to given value.


### GetCounters

`func (o *ResumesByStatusResponse) GetCounters() ResumesByStatusCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *ResumesByStatusResponse) GetCountersOk() (*ResumesByStatusCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *ResumesByStatusResponse) SetCounters(v ResumesByStatusCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *ResumesByStatusResponse) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### SetCountersNil

`func (o *ResumesByStatusResponse) SetCountersNil(b bool)`

 SetCountersNil sets the value for Counters to be an explicit nil

### UnsetCounters
`func (o *ResumesByStatusResponse) UnsetCounters()`

UnsetCounters ensures that no value is present for Counters, not even an explicit nil
### GetNotPublished

`func (o *ResumesByStatusResponse) GetNotPublished() []ResumesSuitableResumeItem`

GetNotPublished returns the NotPublished field if non-nil, zero value otherwise.

### GetNotPublishedOk

`func (o *ResumesByStatusResponse) GetNotPublishedOk() (*[]ResumesSuitableResumeItem, bool)`

GetNotPublishedOk returns a tuple with the NotPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotPublished

`func (o *ResumesByStatusResponse) SetNotPublished(v []ResumesSuitableResumeItem)`

SetNotPublished sets NotPublished field to given value.


### GetSuitable

`func (o *ResumesByStatusResponse) GetSuitable() []ResumesSuitableResumeItem`

GetSuitable returns the Suitable field if non-nil, zero value otherwise.

### GetSuitableOk

`func (o *ResumesByStatusResponse) GetSuitableOk() (*[]ResumesSuitableResumeItem, bool)`

GetSuitableOk returns a tuple with the Suitable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuitable

`func (o *ResumesByStatusResponse) SetSuitable(v []ResumesSuitableResumeItem)`

SetSuitable sets Suitable field to given value.


### GetUnavailable

`func (o *ResumesByStatusResponse) GetUnavailable() []ResumesSuitableResumeItem`

GetUnavailable returns the Unavailable field if non-nil, zero value otherwise.

### GetUnavailableOk

`func (o *ResumesByStatusResponse) GetUnavailableOk() (*[]ResumesSuitableResumeItem, bool)`

GetUnavailableOk returns a tuple with the Unavailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailable

`func (o *ResumesByStatusResponse) SetUnavailable(v []ResumesSuitableResumeItem)`

SetUnavailable sets Unavailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


