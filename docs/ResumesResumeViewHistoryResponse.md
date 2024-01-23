# ResumesResumeViewHistoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**Items** | [**[]ResumesResumeViewHistoryItem**](ResumesResumeViewHistoryItem.md) | Список просмотров резюме | 
**Resume** | [**ResumesResumeViewHistoryResume**](ResumesResumeViewHistoryResume.md) |  | 

## Methods

### NewResumesResumeViewHistoryResponse

`func NewResumesResumeViewHistoryResponse(found float32, page float32, pages float32, perPage float32, items []ResumesResumeViewHistoryItem, resume ResumesResumeViewHistoryResume, ) *ResumesResumeViewHistoryResponse`

NewResumesResumeViewHistoryResponse instantiates a new ResumesResumeViewHistoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesResumeViewHistoryResponseWithDefaults

`func NewResumesResumeViewHistoryResponseWithDefaults() *ResumesResumeViewHistoryResponse`

NewResumesResumeViewHistoryResponseWithDefaults instantiates a new ResumesResumeViewHistoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *ResumesResumeViewHistoryResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *ResumesResumeViewHistoryResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *ResumesResumeViewHistoryResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *ResumesResumeViewHistoryResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ResumesResumeViewHistoryResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ResumesResumeViewHistoryResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *ResumesResumeViewHistoryResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ResumesResumeViewHistoryResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ResumesResumeViewHistoryResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *ResumesResumeViewHistoryResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ResumesResumeViewHistoryResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ResumesResumeViewHistoryResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetItems

`func (o *ResumesResumeViewHistoryResponse) GetItems() []ResumesResumeViewHistoryItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResumesResumeViewHistoryResponse) GetItemsOk() (*[]ResumesResumeViewHistoryItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResumesResumeViewHistoryResponse) SetItems(v []ResumesResumeViewHistoryItem)`

SetItems sets Items field to given value.


### GetResume

`func (o *ResumesResumeViewHistoryResponse) GetResume() ResumesResumeViewHistoryResume`

GetResume returns the Resume field if non-nil, zero value otherwise.

### GetResumeOk

`func (o *ResumesResumeViewHistoryResponse) GetResumeOk() (*ResumesResumeViewHistoryResume, bool)`

GetResumeOk returns a tuple with the Resume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResume

`func (o *ResumesResumeViewHistoryResponse) SetResume(v ResumesResumeViewHistoryResume)`

SetResume sets Resume field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


