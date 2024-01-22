# MeApplicantProfileCounters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewResumeViews** | **int32** | Общее количество новых просмотров всех резюме текущего пользователя | 
**ResumesCount** | **int32** | Общее количество созданных резюме текущего пользователя | 
**UnreadNegotiations** | **int32** | Количество новых непрочитанных откликов (у которых &#x60;has_updates: true&#x60;) | 

## Methods

### NewMeApplicantProfileCounters

`func NewMeApplicantProfileCounters(newResumeViews int32, resumesCount int32, unreadNegotiations int32, ) *MeApplicantProfileCounters`

NewMeApplicantProfileCounters instantiates a new MeApplicantProfileCounters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeApplicantProfileCountersWithDefaults

`func NewMeApplicantProfileCountersWithDefaults() *MeApplicantProfileCounters`

NewMeApplicantProfileCountersWithDefaults instantiates a new MeApplicantProfileCounters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewResumeViews

`func (o *MeApplicantProfileCounters) GetNewResumeViews() int32`

GetNewResumeViews returns the NewResumeViews field if non-nil, zero value otherwise.

### GetNewResumeViewsOk

`func (o *MeApplicantProfileCounters) GetNewResumeViewsOk() (*int32, bool)`

GetNewResumeViewsOk returns a tuple with the NewResumeViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewResumeViews

`func (o *MeApplicantProfileCounters) SetNewResumeViews(v int32)`

SetNewResumeViews sets NewResumeViews field to given value.


### GetResumesCount

`func (o *MeApplicantProfileCounters) GetResumesCount() int32`

GetResumesCount returns the ResumesCount field if non-nil, zero value otherwise.

### GetResumesCountOk

`func (o *MeApplicantProfileCounters) GetResumesCountOk() (*int32, bool)`

GetResumesCountOk returns a tuple with the ResumesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumesCount

`func (o *MeApplicantProfileCounters) SetResumesCount(v int32)`

SetResumesCount sets ResumesCount field to given value.


### GetUnreadNegotiations

`func (o *MeApplicantProfileCounters) GetUnreadNegotiations() int32`

GetUnreadNegotiations returns the UnreadNegotiations field if non-nil, zero value otherwise.

### GetUnreadNegotiationsOk

`func (o *MeApplicantProfileCounters) GetUnreadNegotiationsOk() (*int32, bool)`

GetUnreadNegotiationsOk returns a tuple with the UnreadNegotiations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadNegotiations

`func (o *MeApplicantProfileCounters) SetUnreadNegotiations(v int32)`

SetUnreadNegotiations sets UnreadNegotiations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


