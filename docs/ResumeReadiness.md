# ResumeReadiness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModerationNote** | [**[]ResumeObjectsModerationNote**](ResumeObjectsModerationNote.md) | Замечания модератора. В некоторых случаях замечания могут сопровождаться [блокировкой резюме](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume). Полный список возможных замечаний доступен в поле &#x60;resume_moderation_note&#x60; [в справочнике полей](#tag/Obshie-spravochniki/operation/get-dictionaries)  | 
**Progress** | [**ResumeObjectsProgress**](ResumeObjectsProgress.md) |  | 
**PublishUrl** | **string** | URL для публикации или обновления резюме | 

## Methods

### NewResumeReadiness

`func NewResumeReadiness(moderationNote []ResumeObjectsModerationNote, progress ResumeObjectsProgress, publishUrl string, ) *ResumeReadiness`

NewResumeReadiness instantiates a new ResumeReadiness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeReadinessWithDefaults

`func NewResumeReadinessWithDefaults() *ResumeReadiness`

NewResumeReadinessWithDefaults instantiates a new ResumeReadiness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModerationNote

`func (o *ResumeReadiness) GetModerationNote() []ResumeObjectsModerationNote`

GetModerationNote returns the ModerationNote field if non-nil, zero value otherwise.

### GetModerationNoteOk

`func (o *ResumeReadiness) GetModerationNoteOk() (*[]ResumeObjectsModerationNote, bool)`

GetModerationNoteOk returns a tuple with the ModerationNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationNote

`func (o *ResumeReadiness) SetModerationNote(v []ResumeObjectsModerationNote)`

SetModerationNote sets ModerationNote field to given value.


### GetProgress

`func (o *ResumeReadiness) GetProgress() ResumeObjectsProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ResumeReadiness) GetProgressOk() (*ResumeObjectsProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ResumeReadiness) SetProgress(v ResumeObjectsProgress)`

SetProgress sets Progress field to given value.


### GetPublishUrl

`func (o *ResumeReadiness) GetPublishUrl() string`

GetPublishUrl returns the PublishUrl field if non-nil, zero value otherwise.

### GetPublishUrlOk

`func (o *ResumeReadiness) GetPublishUrlOk() (*string, bool)`

GetPublishUrlOk returns a tuple with the PublishUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishUrl

`func (o *ResumeReadiness) SetPublishUrl(v string)`

SetPublishUrl sets PublishUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


