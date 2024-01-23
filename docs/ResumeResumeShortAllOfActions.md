# ResumeResumeShortAllOfActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Download** | [**ResumeObjectsDownload**](ResumeObjectsDownload.md) | Ссылки для скачивания резюме в нескольких форматах ([подробнее](#tag/Prosmotr-rezyume/operation/get-resume)) (атрибут &#39;actions&#39;)
 | 
**DownloadWithContact** | Pointer to [**NullableResumeObjectsDownload**](ResumeObjectsDownload.md) |  | [optional] 
**GetWithContact** | Pointer to [**NullableIncludesUrl**](IncludesUrl.md) |  | [optional] 

## Methods

### NewResumeResumeShortAllOfActions

`func NewResumeResumeShortAllOfActions(download ResumeObjectsDownload, ) *ResumeResumeShortAllOfActions`

NewResumeResumeShortAllOfActions instantiates a new ResumeResumeShortAllOfActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeResumeShortAllOfActionsWithDefaults

`func NewResumeResumeShortAllOfActionsWithDefaults() *ResumeResumeShortAllOfActions`

NewResumeResumeShortAllOfActionsWithDefaults instantiates a new ResumeResumeShortAllOfActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownload

`func (o *ResumeResumeShortAllOfActions) GetDownload() ResumeObjectsDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeResumeShortAllOfActions) GetDownloadOk() (*ResumeObjectsDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeResumeShortAllOfActions) SetDownload(v ResumeObjectsDownload)`

SetDownload sets Download field to given value.


### GetDownloadWithContact

`func (o *ResumeResumeShortAllOfActions) GetDownloadWithContact() ResumeObjectsDownload`

GetDownloadWithContact returns the DownloadWithContact field if non-nil, zero value otherwise.

### GetDownloadWithContactOk

`func (o *ResumeResumeShortAllOfActions) GetDownloadWithContactOk() (*ResumeObjectsDownload, bool)`

GetDownloadWithContactOk returns a tuple with the DownloadWithContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadWithContact

`func (o *ResumeResumeShortAllOfActions) SetDownloadWithContact(v ResumeObjectsDownload)`

SetDownloadWithContact sets DownloadWithContact field to given value.

### HasDownloadWithContact

`func (o *ResumeResumeShortAllOfActions) HasDownloadWithContact() bool`

HasDownloadWithContact returns a boolean if a field has been set.

### SetDownloadWithContactNil

`func (o *ResumeResumeShortAllOfActions) SetDownloadWithContactNil(b bool)`

 SetDownloadWithContactNil sets the value for DownloadWithContact to be an explicit nil

### UnsetDownloadWithContact
`func (o *ResumeResumeShortAllOfActions) UnsetDownloadWithContact()`

UnsetDownloadWithContact ensures that no value is present for DownloadWithContact, not even an explicit nil
### GetGetWithContact

`func (o *ResumeResumeShortAllOfActions) GetGetWithContact() IncludesUrl`

GetGetWithContact returns the GetWithContact field if non-nil, zero value otherwise.

### GetGetWithContactOk

`func (o *ResumeResumeShortAllOfActions) GetGetWithContactOk() (*IncludesUrl, bool)`

GetGetWithContactOk returns a tuple with the GetWithContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetWithContact

`func (o *ResumeResumeShortAllOfActions) SetGetWithContact(v IncludesUrl)`

SetGetWithContact sets GetWithContact field to given value.

### HasGetWithContact

`func (o *ResumeResumeShortAllOfActions) HasGetWithContact() bool`

HasGetWithContact returns a boolean if a field has been set.

### SetGetWithContactNil

`func (o *ResumeResumeShortAllOfActions) SetGetWithContactNil(b bool)`

 SetGetWithContactNil sets the value for GetWithContact to be an explicit nil

### UnsetGetWithContact
`func (o *ResumeResumeShortAllOfActions) UnsetGetWithContact()`

UnsetGetWithContact ensures that no value is present for GetWithContact, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


