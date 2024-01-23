# ResumeStatusReadiness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | **bool** | Заблокировано ли резюме ([подробнее](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume)) | 
**CanPublishOrUpdate** | Pointer to **NullableBool** | Можно ли опубликовать или обновить данное резюме | [optional] 
**Finished** | **bool** | Заполнено ли резюме | 
**Status** | [**IncludesIdName**](IncludesIdName.md) |  | 
**ModerationNote** | [**[]ResumeObjectsModerationNote**](ResumeObjectsModerationNote.md) | Замечания модератора. В некоторых случаях замечания могут сопровождаться [блокировкой резюме](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume). Полный список возможных замечаний доступен в поле &#x60;resume_moderation_note&#x60; [в справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**Progress** | [**ResumeObjectsProgress**](ResumeObjectsProgress.md) |  | 
**PublishUrl** | **string** | URL для публикации или обновления резюме | 

## Methods

### NewResumeStatusReadiness

`func NewResumeStatusReadiness(blocked bool, finished bool, status IncludesIdName, moderationNote []ResumeObjectsModerationNote, progress ResumeObjectsProgress, publishUrl string, ) *ResumeStatusReadiness`

NewResumeStatusReadiness instantiates a new ResumeStatusReadiness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeStatusReadinessWithDefaults

`func NewResumeStatusReadinessWithDefaults() *ResumeStatusReadiness`

NewResumeStatusReadinessWithDefaults instantiates a new ResumeStatusReadiness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *ResumeStatusReadiness) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *ResumeStatusReadiness) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *ResumeStatusReadiness) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetCanPublishOrUpdate

`func (o *ResumeStatusReadiness) GetCanPublishOrUpdate() bool`

GetCanPublishOrUpdate returns the CanPublishOrUpdate field if non-nil, zero value otherwise.

### GetCanPublishOrUpdateOk

`func (o *ResumeStatusReadiness) GetCanPublishOrUpdateOk() (*bool, bool)`

GetCanPublishOrUpdateOk returns a tuple with the CanPublishOrUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPublishOrUpdate

`func (o *ResumeStatusReadiness) SetCanPublishOrUpdate(v bool)`

SetCanPublishOrUpdate sets CanPublishOrUpdate field to given value.

### HasCanPublishOrUpdate

`func (o *ResumeStatusReadiness) HasCanPublishOrUpdate() bool`

HasCanPublishOrUpdate returns a boolean if a field has been set.

### SetCanPublishOrUpdateNil

`func (o *ResumeStatusReadiness) SetCanPublishOrUpdateNil(b bool)`

 SetCanPublishOrUpdateNil sets the value for CanPublishOrUpdate to be an explicit nil

### UnsetCanPublishOrUpdate
`func (o *ResumeStatusReadiness) UnsetCanPublishOrUpdate()`

UnsetCanPublishOrUpdate ensures that no value is present for CanPublishOrUpdate, not even an explicit nil
### GetFinished

`func (o *ResumeStatusReadiness) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ResumeStatusReadiness) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ResumeStatusReadiness) SetFinished(v bool)`

SetFinished sets Finished field to given value.


### GetStatus

`func (o *ResumeStatusReadiness) GetStatus() IncludesIdName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResumeStatusReadiness) GetStatusOk() (*IncludesIdName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResumeStatusReadiness) SetStatus(v IncludesIdName)`

SetStatus sets Status field to given value.


### GetModerationNote

`func (o *ResumeStatusReadiness) GetModerationNote() []ResumeObjectsModerationNote`

GetModerationNote returns the ModerationNote field if non-nil, zero value otherwise.

### GetModerationNoteOk

`func (o *ResumeStatusReadiness) GetModerationNoteOk() (*[]ResumeObjectsModerationNote, bool)`

GetModerationNoteOk returns a tuple with the ModerationNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationNote

`func (o *ResumeStatusReadiness) SetModerationNote(v []ResumeObjectsModerationNote)`

SetModerationNote sets ModerationNote field to given value.


### GetProgress

`func (o *ResumeStatusReadiness) GetProgress() ResumeObjectsProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ResumeStatusReadiness) GetProgressOk() (*ResumeObjectsProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ResumeStatusReadiness) SetProgress(v ResumeObjectsProgress)`

SetProgress sets Progress field to given value.


### GetPublishUrl

`func (o *ResumeStatusReadiness) GetPublishUrl() string`

GetPublishUrl returns the PublishUrl field if non-nil, zero value otherwise.

### GetPublishUrlOk

`func (o *ResumeStatusReadiness) GetPublishUrlOk() (*string, bool)`

GetPublishUrlOk returns a tuple with the PublishUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishUrl

`func (o *ResumeStatusReadiness) SetPublishUrl(v string)`

SetPublishUrl sets PublishUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


