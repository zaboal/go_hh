# ResumeObjectsActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Download** | [**ResumeObjectsDownload**](ResumeObjectsDownload.md) | Ссылки для скачивания резюме в нескольких форматах ([подробнее](#tag/Prosmotr-rezyume/operation/get-resume)) (атрибут &#39;actions&#39;)
 | 
**DownloadWithContact** | Pointer to [**NullableResumeObjectsDownload**](ResumeObjectsDownload.md) |  | [optional] 
**GetWithContact** | Pointer to [**NullableIncludesUrl**](IncludesUrl.md) |  | [optional] 

## Methods

### NewResumeObjectsActions

`func NewResumeObjectsActions(download ResumeObjectsDownload, ) *ResumeObjectsActions`

NewResumeObjectsActions instantiates a new ResumeObjectsActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeObjectsActionsWithDefaults

`func NewResumeObjectsActionsWithDefaults() *ResumeObjectsActions`

NewResumeObjectsActionsWithDefaults instantiates a new ResumeObjectsActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownload

`func (o *ResumeObjectsActions) GetDownload() ResumeObjectsDownload`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ResumeObjectsActions) GetDownloadOk() (*ResumeObjectsDownload, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ResumeObjectsActions) SetDownload(v ResumeObjectsDownload)`

SetDownload sets Download field to given value.


### GetDownloadWithContact

`func (o *ResumeObjectsActions) GetDownloadWithContact() ResumeObjectsDownload`

GetDownloadWithContact returns the DownloadWithContact field if non-nil, zero value otherwise.

### GetDownloadWithContactOk

`func (o *ResumeObjectsActions) GetDownloadWithContactOk() (*ResumeObjectsDownload, bool)`

GetDownloadWithContactOk returns a tuple with the DownloadWithContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadWithContact

`func (o *ResumeObjectsActions) SetDownloadWithContact(v ResumeObjectsDownload)`

SetDownloadWithContact sets DownloadWithContact field to given value.

### HasDownloadWithContact

`func (o *ResumeObjectsActions) HasDownloadWithContact() bool`

HasDownloadWithContact returns a boolean if a field has been set.

### SetDownloadWithContactNil

`func (o *ResumeObjectsActions) SetDownloadWithContactNil(b bool)`

 SetDownloadWithContactNil sets the value for DownloadWithContact to be an explicit nil

### UnsetDownloadWithContact
`func (o *ResumeObjectsActions) UnsetDownloadWithContact()`

UnsetDownloadWithContact ensures that no value is present for DownloadWithContact, not even an explicit nil
### GetGetWithContact

`func (o *ResumeObjectsActions) GetGetWithContact() IncludesUrl`

GetGetWithContact returns the GetWithContact field if non-nil, zero value otherwise.

### GetGetWithContactOk

`func (o *ResumeObjectsActions) GetGetWithContactOk() (*IncludesUrl, bool)`

GetGetWithContactOk returns a tuple with the GetWithContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGetWithContact

`func (o *ResumeObjectsActions) SetGetWithContact(v IncludesUrl)`

SetGetWithContact sets GetWithContact field to given value.

### HasGetWithContact

`func (o *ResumeObjectsActions) HasGetWithContact() bool`

HasGetWithContact returns a boolean if a field has been set.

### SetGetWithContactNil

`func (o *ResumeObjectsActions) SetGetWithContactNil(b bool)`

 SetGetWithContactNil sets the value for GetWithContact to be an explicit nil

### UnsetGetWithContact
`func (o *ResumeObjectsActions) UnsetGetWithContact()`

UnsetGetWithContact ensures that no value is present for GetWithContact, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


