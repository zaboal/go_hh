# ResumesResumeViewHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **string** | Дата создания записи (дата просмотра резюме работодателем) | 
**Employer** | [**ResumesResumeViewHistoryItemEmployer**](ResumesResumeViewHistoryItemEmployer.md) |  | 
**Viewed** | **bool** | Отметка о просмотре записи | 

## Methods

### NewResumesResumeViewHistoryItem

`func NewResumesResumeViewHistoryItem(createdAt string, employer ResumesResumeViewHistoryItemEmployer, viewed bool, ) *ResumesResumeViewHistoryItem`

NewResumesResumeViewHistoryItem instantiates a new ResumesResumeViewHistoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesResumeViewHistoryItemWithDefaults

`func NewResumesResumeViewHistoryItemWithDefaults() *ResumesResumeViewHistoryItem`

NewResumesResumeViewHistoryItemWithDefaults instantiates a new ResumesResumeViewHistoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *ResumesResumeViewHistoryItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResumesResumeViewHistoryItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResumesResumeViewHistoryItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmployer

`func (o *ResumesResumeViewHistoryItem) GetEmployer() ResumesResumeViewHistoryItemEmployer`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *ResumesResumeViewHistoryItem) GetEmployerOk() (*ResumesResumeViewHistoryItemEmployer, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *ResumesResumeViewHistoryItem) SetEmployer(v ResumesResumeViewHistoryItemEmployer)`

SetEmployer sets Employer field to given value.


### GetViewed

`func (o *ResumesResumeViewHistoryItem) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *ResumesResumeViewHistoryItem) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *ResumesResumeViewHistoryItem) SetViewed(v bool)`

SetViewed sets Viewed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


