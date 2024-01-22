# ResumesResumeNegotiationsHistoryVacancyItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | Дата изменения состояния отклика/приглашения | 
**EmployerState** | [**ResumesNegotiationNanoEmployerState**](ResumesNegotiationNanoEmployerState.md) |  | 
**WithMessage** | **bool** | Признак того, что при изменении состояния отклика/приглашения было отправлено сообщение соискателю | 

## Methods

### NewResumesResumeNegotiationsHistoryVacancyItem

`func NewResumesResumeNegotiationsHistoryVacancyItem(createdAt string, employerState ResumesNegotiationNanoEmployerState, withMessage bool, ) *ResumesResumeNegotiationsHistoryVacancyItem`

NewResumesResumeNegotiationsHistoryVacancyItem instantiates a new ResumesResumeNegotiationsHistoryVacancyItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesResumeNegotiationsHistoryVacancyItemWithDefaults

`func NewResumesResumeNegotiationsHistoryVacancyItemWithDefaults() *ResumesResumeNegotiationsHistoryVacancyItem`

NewResumesResumeNegotiationsHistoryVacancyItemWithDefaults instantiates a new ResumesResumeNegotiationsHistoryVacancyItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ResumesResumeNegotiationsHistoryVacancyItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumesResumeNegotiationsHistoryVacancyItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumesResumeNegotiationsHistoryVacancyItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmployerState

`func (o *ResumesResumeNegotiationsHistoryVacancyItem) GetEmployerState() ResumesNegotiationNanoEmployerState`

GetEmployerState returns the EmployerState field if non-nil, zero value otherwise.

### GetEmployerStateOk

`func (o *ResumesResumeNegotiationsHistoryVacancyItem) GetEmployerStateOk() (*ResumesNegotiationNanoEmployerState, bool)`

GetEmployerStateOk returns a tuple with the EmployerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployerState

`func (o *ResumesResumeNegotiationsHistoryVacancyItem) SetEmployerState(v ResumesNegotiationNanoEmployerState)`

SetEmployerState sets EmployerState field to given value.


### GetWithMessage

`func (o *ResumesResumeNegotiationsHistoryVacancyItem) GetWithMessage() bool`

GetWithMessage returns the WithMessage field if non-nil, zero value otherwise.

### GetWithMessageOk

`func (o *ResumesResumeNegotiationsHistoryVacancyItem) GetWithMessageOk() (*bool, bool)`

GetWithMessageOk returns a tuple with the WithMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithMessage

`func (o *ResumesResumeNegotiationsHistoryVacancyItem) SetWithMessage(v bool)`

SetWithMessage sets WithMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


