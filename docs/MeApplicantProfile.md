# MeApplicantProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | [**MeApplicantProfileCounters**](MeApplicantProfileCounters.md) |  | 
**Employer** | Pointer to **map[string]interface{}** |  | [optional] 
**IsInSearch** | **bool** | Имеет ли текущий пользователь статус \&quot;ищу работу\&quot; | 
**Manager** | Pointer to **map[string]interface{}** |  | [optional] 
**NegotiationsUrl** | **string** | URL, на который нужно сделать GET-запрос, чтобы получить список откликов/приглашений текущего пользователя  | 
**PersonalManager** | Pointer to **map[string]interface{}** |  | [optional] 
**ProfileVideos** | Pointer to [**ProfileVideosList**](ProfileVideosList.md) |  | [optional] 
**ResumesUrl** | **string** | URL, на который нужно сделать GET-запрос, чтобы получить список резюме текущего пользователя  | 
**UserStatuses** | Pointer to [**UserStatusesApplicant**](UserStatusesApplicant.md) |  | [optional] 

## Methods

### NewMeApplicantProfile

`func NewMeApplicantProfile(counters MeApplicantProfileCounters, isInSearch bool, negotiationsUrl string, resumesUrl string, ) *MeApplicantProfile`

NewMeApplicantProfile instantiates a new MeApplicantProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeApplicantProfileWithDefaults

`func NewMeApplicantProfileWithDefaults() *MeApplicantProfile`

NewMeApplicantProfileWithDefaults instantiates a new MeApplicantProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *MeApplicantProfile) GetCounters() MeApplicantProfileCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *MeApplicantProfile) GetCountersOk() (*MeApplicantProfileCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *MeApplicantProfile) SetCounters(v MeApplicantProfileCounters)`

SetCounters sets Counters field to given value.


### GetEmployer

`func (o *MeApplicantProfile) GetEmployer() map[string]interface{}`

GetEmployer returns the Employer field if non-nil, zero value otherwise.

### GetEmployerOk

`func (o *MeApplicantProfile) GetEmployerOk() (*map[string]interface{}, bool)`

GetEmployerOk returns a tuple with the Employer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployer

`func (o *MeApplicantProfile) SetEmployer(v map[string]interface{})`

SetEmployer sets Employer field to given value.

### HasEmployer

`func (o *MeApplicantProfile) HasEmployer() bool`

HasEmployer returns a boolean if a field has been set.

### SetEmployerNil

`func (o *MeApplicantProfile) SetEmployerNil(b bool)`

 SetEmployerNil sets the value for Employer to be an explicit nil

### UnsetEmployer
`func (o *MeApplicantProfile) UnsetEmployer()`

UnsetEmployer ensures that no value is present for Employer, not even an explicit nil
### GetIsInSearch

`func (o *MeApplicantProfile) GetIsInSearch() bool`

GetIsInSearch returns the IsInSearch field if non-nil, zero value otherwise.

### GetIsInSearchOk

`func (o *MeApplicantProfile) GetIsInSearchOk() (*bool, bool)`

GetIsInSearchOk returns a tuple with the IsInSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInSearch

`func (o *MeApplicantProfile) SetIsInSearch(v bool)`

SetIsInSearch sets IsInSearch field to given value.


### GetManager

`func (o *MeApplicantProfile) GetManager() map[string]interface{}`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *MeApplicantProfile) GetManagerOk() (*map[string]interface{}, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *MeApplicantProfile) SetManager(v map[string]interface{})`

SetManager sets Manager field to given value.

### HasManager

`func (o *MeApplicantProfile) HasManager() bool`

HasManager returns a boolean if a field has been set.

### SetManagerNil

`func (o *MeApplicantProfile) SetManagerNil(b bool)`

 SetManagerNil sets the value for Manager to be an explicit nil

### UnsetManager
`func (o *MeApplicantProfile) UnsetManager()`

UnsetManager ensures that no value is present for Manager, not even an explicit nil
### GetNegotiationsUrl

`func (o *MeApplicantProfile) GetNegotiationsUrl() string`

GetNegotiationsUrl returns the NegotiationsUrl field if non-nil, zero value otherwise.

### GetNegotiationsUrlOk

`func (o *MeApplicantProfile) GetNegotiationsUrlOk() (*string, bool)`

GetNegotiationsUrlOk returns a tuple with the NegotiationsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsUrl

`func (o *MeApplicantProfile) SetNegotiationsUrl(v string)`

SetNegotiationsUrl sets NegotiationsUrl field to given value.


### GetPersonalManager

`func (o *MeApplicantProfile) GetPersonalManager() map[string]interface{}`

GetPersonalManager returns the PersonalManager field if non-nil, zero value otherwise.

### GetPersonalManagerOk

`func (o *MeApplicantProfile) GetPersonalManagerOk() (*map[string]interface{}, bool)`

GetPersonalManagerOk returns a tuple with the PersonalManager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersonalManager

`func (o *MeApplicantProfile) SetPersonalManager(v map[string]interface{})`

SetPersonalManager sets PersonalManager field to given value.

### HasPersonalManager

`func (o *MeApplicantProfile) HasPersonalManager() bool`

HasPersonalManager returns a boolean if a field has been set.

### SetPersonalManagerNil

`func (o *MeApplicantProfile) SetPersonalManagerNil(b bool)`

 SetPersonalManagerNil sets the value for PersonalManager to be an explicit nil

### UnsetPersonalManager
`func (o *MeApplicantProfile) UnsetPersonalManager()`

UnsetPersonalManager ensures that no value is present for PersonalManager, not even an explicit nil
### GetProfileVideos

`func (o *MeApplicantProfile) GetProfileVideos() ProfileVideosList`

GetProfileVideos returns the ProfileVideos field if non-nil, zero value otherwise.

### GetProfileVideosOk

`func (o *MeApplicantProfile) GetProfileVideosOk() (*ProfileVideosList, bool)`

GetProfileVideosOk returns a tuple with the ProfileVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileVideos

`func (o *MeApplicantProfile) SetProfileVideos(v ProfileVideosList)`

SetProfileVideos sets ProfileVideos field to given value.

### HasProfileVideos

`func (o *MeApplicantProfile) HasProfileVideos() bool`

HasProfileVideos returns a boolean if a field has been set.

### GetResumesUrl

`func (o *MeApplicantProfile) GetResumesUrl() string`

GetResumesUrl returns the ResumesUrl field if non-nil, zero value otherwise.

### GetResumesUrlOk

`func (o *MeApplicantProfile) GetResumesUrlOk() (*string, bool)`

GetResumesUrlOk returns a tuple with the ResumesUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumesUrl

`func (o *MeApplicantProfile) SetResumesUrl(v string)`

SetResumesUrl sets ResumesUrl field to given value.


### GetUserStatuses

`func (o *MeApplicantProfile) GetUserStatuses() UserStatusesApplicant`

GetUserStatuses returns the UserStatuses field if non-nil, zero value otherwise.

### GetUserStatusesOk

`func (o *MeApplicantProfile) GetUserStatusesOk() (*UserStatusesApplicant, bool)`

GetUserStatusesOk returns a tuple with the UserStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserStatuses

`func (o *MeApplicantProfile) SetUserStatuses(v UserStatusesApplicant)`

SetUserStatuses sets UserStatuses field to given value.

### HasUserStatuses

`func (o *MeApplicantProfile) HasUserStatuses() bool`

HasUserStatuses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


