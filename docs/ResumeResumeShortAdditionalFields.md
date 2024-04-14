# ResumeResumeShortAdditionalFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**ResumeObjectsActions**](ResumeObjectsActions.md) | Дополнительные действия | 
**Favorited** | **bool** | Добавлено ли резюме в избранные | 
**Marked** | Pointer to **bool** | Выделено ли резюме в поиске | [optional] 
**NegotiationsHistory** | [**ResumeObjectsNegotiationsHistoryUrl**](ResumeObjectsNegotiationsHistoryUrl.md) | Краткая история откликов/приглашений по резюме | 
**Owner** | [**ResumeObjectsOwner**](ResumeObjectsOwner.md) | Информация о владельце резюме | 
**Photo** | Pointer to [**NullableResumeObjectsPhoto**](ResumeObjectsPhoto.md) | Фотография пользователя | [optional] 
**Tags** | Pointer to [**[]IncludesId**](IncludesId.md) | Теги к резюме | [optional] 
**Viewed** | **bool** | Было ли резюме уже просмотрено работодателем | 

## Methods

### NewResumeResumeShortAdditionalFields

`func NewResumeResumeShortAdditionalFields(actions ResumeObjectsActions, favorited bool, negotiationsHistory ResumeObjectsNegotiationsHistoryUrl, owner ResumeObjectsOwner, viewed bool, ) *ResumeResumeShortAdditionalFields`

NewResumeResumeShortAdditionalFields instantiates a new ResumeResumeShortAdditionalFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeShortAdditionalFieldsWithDefaults

`func NewResumeResumeShortAdditionalFieldsWithDefaults() *ResumeResumeShortAdditionalFields`

NewResumeResumeShortAdditionalFieldsWithDefaults instantiates a new ResumeResumeShortAdditionalFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ResumeResumeShortAdditionalFields) GetActions() ResumeObjectsActions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ResumeResumeShortAdditionalFields) GetActionsOk() (*ResumeObjectsActions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ResumeResumeShortAdditionalFields) SetActions(v ResumeObjectsActions)`

SetActions sets Actions field to given value.


### GetFavorited

`func (o *ResumeResumeShortAdditionalFields) GetFavorited() bool`

GetFavorited returns the Favorited field if non-nil, zero value otherwise.

### GetFavoritedOk

`func (o *ResumeResumeShortAdditionalFields) GetFavoritedOk() (*bool, bool)`

GetFavoritedOk returns a tuple with the Favorited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorited

`func (o *ResumeResumeShortAdditionalFields) SetFavorited(v bool)`

SetFavorited sets Favorited field to given value.


### GetMarked

`func (o *ResumeResumeShortAdditionalFields) GetMarked() bool`

GetMarked returns the Marked field if non-nil, zero value otherwise.

### GetMarkedOk

`func (o *ResumeResumeShortAdditionalFields) GetMarkedOk() (*bool, bool)`

GetMarkedOk returns a tuple with the Marked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarked

`func (o *ResumeResumeShortAdditionalFields) SetMarked(v bool)`

SetMarked sets Marked field to given value.

### HasMarked

`func (o *ResumeResumeShortAdditionalFields) HasMarked() bool`

HasMarked returns a boolean if a field has been set.

### GetNegotiationsHistory

`func (o *ResumeResumeShortAdditionalFields) GetNegotiationsHistory() ResumeObjectsNegotiationsHistoryUrl`

GetNegotiationsHistory returns the NegotiationsHistory field if non-nil, zero value otherwise.

### GetNegotiationsHistoryOk

`func (o *ResumeResumeShortAdditionalFields) GetNegotiationsHistoryOk() (*ResumeObjectsNegotiationsHistoryUrl, bool)`

GetNegotiationsHistoryOk returns a tuple with the NegotiationsHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationsHistory

`func (o *ResumeResumeShortAdditionalFields) SetNegotiationsHistory(v ResumeObjectsNegotiationsHistoryUrl)`

SetNegotiationsHistory sets NegotiationsHistory field to given value.


### GetOwner

`func (o *ResumeResumeShortAdditionalFields) GetOwner() ResumeObjectsOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ResumeResumeShortAdditionalFields) GetOwnerOk() (*ResumeObjectsOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ResumeResumeShortAdditionalFields) SetOwner(v ResumeObjectsOwner)`

SetOwner sets Owner field to given value.


### GetPhoto

`func (o *ResumeResumeShortAdditionalFields) GetPhoto() ResumeObjectsPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *ResumeResumeShortAdditionalFields) GetPhotoOk() (*ResumeObjectsPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *ResumeResumeShortAdditionalFields) SetPhoto(v ResumeObjectsPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *ResumeResumeShortAdditionalFields) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *ResumeResumeShortAdditionalFields) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *ResumeResumeShortAdditionalFields) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetTags

`func (o *ResumeResumeShortAdditionalFields) GetTags() []IncludesId`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ResumeResumeShortAdditionalFields) GetTagsOk() (*[]IncludesId, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ResumeResumeShortAdditionalFields) SetTags(v []IncludesId)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ResumeResumeShortAdditionalFields) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetViewed

`func (o *ResumeResumeShortAdditionalFields) GetViewed() bool`

GetViewed returns the Viewed field if non-nil, zero value otherwise.

### GetViewedOk

`func (o *ResumeResumeShortAdditionalFields) GetViewedOk() (*bool, bool)`

GetViewedOk returns a tuple with the Viewed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewed

`func (o *ResumeResumeShortAdditionalFields) SetViewed(v bool)`

SetViewed sets Viewed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


