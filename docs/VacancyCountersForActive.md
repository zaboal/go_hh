# VacancyCountersForActive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Calls** | Pointer to **float32** | Общее количество звонков | [optional] 
**Invitations** | **float32** | Количество приглашений на вакансию | 
**InvitationsAndResponses** | Pointer to **float32** | Количество откликнувшихся и приглашенных соискателей на вакансию | [optional] 
**NewMissedCalls** | Pointer to **float32** | Общее количество новых пропущенных звонков | [optional] 
**Responses** | **float32** | Количество откликов на вакансию | 
**ResumesInProgress** | **float32** | Количество резюме в работе на вакансию | 
**UnreadResponses** | **float32** | Количество непросмотренных откликов на вакансию | 
**Views** | **float32** | Количество просмотров вакансии | 

## Methods

### NewVacancyCountersForActive

`func NewVacancyCountersForActive(invitations float32, responses float32, resumesInProgress float32, unreadResponses float32, views float32, ) *VacancyCountersForActive`

NewVacancyCountersForActive instantiates a new VacancyCountersForActive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyCountersForActiveWithDefaults

`func NewVacancyCountersForActiveWithDefaults() *VacancyCountersForActive`

NewVacancyCountersForActiveWithDefaults instantiates a new VacancyCountersForActive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCalls

`func (o *VacancyCountersForActive) GetCalls() float32`

GetCalls returns the Calls field if non-nil, zero value otherwise.

### GetCallsOk

`func (o *VacancyCountersForActive) GetCallsOk() (*float32, bool)`

GetCallsOk returns a tuple with the Calls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCalls

`func (o *VacancyCountersForActive) SetCalls(v float32)`

SetCalls sets Calls field to given value.

### HasCalls

`func (o *VacancyCountersForActive) HasCalls() bool`

HasCalls returns a boolean if a field has been set.

### GetInvitations

`func (o *VacancyCountersForActive) GetInvitations() float32`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *VacancyCountersForActive) GetInvitationsOk() (*float32, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *VacancyCountersForActive) SetInvitations(v float32)`

SetInvitations sets Invitations field to given value.


### GetInvitationsAndResponses

`func (o *VacancyCountersForActive) GetInvitationsAndResponses() float32`

GetInvitationsAndResponses returns the InvitationsAndResponses field if non-nil, zero value otherwise.

### GetInvitationsAndResponsesOk

`func (o *VacancyCountersForActive) GetInvitationsAndResponsesOk() (*float32, bool)`

GetInvitationsAndResponsesOk returns a tuple with the InvitationsAndResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitationsAndResponses

`func (o *VacancyCountersForActive) SetInvitationsAndResponses(v float32)`

SetInvitationsAndResponses sets InvitationsAndResponses field to given value.

### HasInvitationsAndResponses

`func (o *VacancyCountersForActive) HasInvitationsAndResponses() bool`

HasInvitationsAndResponses returns a boolean if a field has been set.

### GetNewMissedCalls

`func (o *VacancyCountersForActive) GetNewMissedCalls() float32`

GetNewMissedCalls returns the NewMissedCalls field if non-nil, zero value otherwise.

### GetNewMissedCallsOk

`func (o *VacancyCountersForActive) GetNewMissedCallsOk() (*float32, bool)`

GetNewMissedCallsOk returns a tuple with the NewMissedCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMissedCalls

`func (o *VacancyCountersForActive) SetNewMissedCalls(v float32)`

SetNewMissedCalls sets NewMissedCalls field to given value.

### HasNewMissedCalls

`func (o *VacancyCountersForActive) HasNewMissedCalls() bool`

HasNewMissedCalls returns a boolean if a field has been set.

### GetResponses

`func (o *VacancyCountersForActive) GetResponses() float32`

GetResponses returns the Responses field if non-nil, zero value otherwise.

### GetResponsesOk

`func (o *VacancyCountersForActive) GetResponsesOk() (*float32, bool)`

GetResponsesOk returns a tuple with the Responses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponses

`func (o *VacancyCountersForActive) SetResponses(v float32)`

SetResponses sets Responses field to given value.


### GetResumesInProgress

`func (o *VacancyCountersForActive) GetResumesInProgress() float32`

GetResumesInProgress returns the ResumesInProgress field if non-nil, zero value otherwise.

### GetResumesInProgressOk

`func (o *VacancyCountersForActive) GetResumesInProgressOk() (*float32, bool)`

GetResumesInProgressOk returns a tuple with the ResumesInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumesInProgress

`func (o *VacancyCountersForActive) SetResumesInProgress(v float32)`

SetResumesInProgress sets ResumesInProgress field to given value.


### GetUnreadResponses

`func (o *VacancyCountersForActive) GetUnreadResponses() float32`

GetUnreadResponses returns the UnreadResponses field if non-nil, zero value otherwise.

### GetUnreadResponsesOk

`func (o *VacancyCountersForActive) GetUnreadResponsesOk() (*float32, bool)`

GetUnreadResponsesOk returns a tuple with the UnreadResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadResponses

`func (o *VacancyCountersForActive) SetUnreadResponses(v float32)`

SetUnreadResponses sets UnreadResponses field to given value.


### GetViews

`func (o *VacancyCountersForActive) GetViews() float32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *VacancyCountersForActive) GetViewsOk() (*float32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *VacancyCountersForActive) SetViews(v float32)`

SetViews sets Views field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


