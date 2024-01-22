# ResumeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | **bool** | Заблокировано ли резюме ([подробнее](#tag/Rezyume.-Prosmotr-informacii/Status-rezyume)) | 
**CanPublishOrUpdate** | Pointer to **NullableBool** | Можно ли опубликовать или обновить данное резюме | [optional] 
**Finished** | **bool** | Заполнено ли резюме | 
**Status** | [**IncludesIdName**](IncludesIdName.md) |  | 

## Methods

### NewResumeStatus

`func NewResumeStatus(blocked bool, finished bool, status IncludesIdName, ) *ResumeStatus`

NewResumeStatus instantiates a new ResumeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeStatusWithDefaults

`func NewResumeStatusWithDefaults() *ResumeStatus`

NewResumeStatusWithDefaults instantiates a new ResumeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *ResumeStatus) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *ResumeStatus) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *ResumeStatus) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetCanPublishOrUpdate

`func (o *ResumeStatus) GetCanPublishOrUpdate() bool`

GetCanPublishOrUpdate returns the CanPublishOrUpdate field if non-nil, zero value otherwise.

### GetCanPublishOrUpdateOk

`func (o *ResumeStatus) GetCanPublishOrUpdateOk() (*bool, bool)`

GetCanPublishOrUpdateOk returns a tuple with the CanPublishOrUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPublishOrUpdate

`func (o *ResumeStatus) SetCanPublishOrUpdate(v bool)`

SetCanPublishOrUpdate sets CanPublishOrUpdate field to given value.

### HasCanPublishOrUpdate

`func (o *ResumeStatus) HasCanPublishOrUpdate() bool`

HasCanPublishOrUpdate returns a boolean if a field has been set.

### SetCanPublishOrUpdateNil

`func (o *ResumeStatus) SetCanPublishOrUpdateNil(b bool)`

 SetCanPublishOrUpdateNil sets the value for CanPublishOrUpdate to be an explicit nil

### UnsetCanPublishOrUpdate
`func (o *ResumeStatus) UnsetCanPublishOrUpdate()`

UnsetCanPublishOrUpdate ensures that no value is present for CanPublishOrUpdate, not even an explicit nil
### GetFinished

`func (o *ResumeStatus) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *ResumeStatus) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *ResumeStatus) SetFinished(v bool)`

SetFinished sets Finished field to given value.


### GetStatus

`func (o *ResumeStatus) GetStatus() IncludesIdName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResumeStatus) GetStatusOk() (*IncludesIdName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResumeStatus) SetStatus(v IncludesIdName)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


