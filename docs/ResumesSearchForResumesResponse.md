# ResumesSearchForResumesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**Items** | [**[]ResumesSearchForEmployerAndApplicant**](ResumesSearchForEmployerAndApplicant.md) | Массив найденных резюме | 

## Methods

### NewResumesSearchForResumesResponse

`func NewResumesSearchForResumesResponse(found float32, page float32, pages float32, perPage float32, items []ResumesSearchForEmployerAndApplicant, ) *ResumesSearchForResumesResponse`

NewResumesSearchForResumesResponse instantiates a new ResumesSearchForResumesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumesSearchForResumesResponseWithDefaults

`func NewResumesSearchForResumesResponseWithDefaults() *ResumesSearchForResumesResponse`

NewResumesSearchForResumesResponseWithDefaults instantiates a new ResumesSearchForResumesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *ResumesSearchForResumesResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *ResumesSearchForResumesResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *ResumesSearchForResumesResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *ResumesSearchForResumesResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ResumesSearchForResumesResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ResumesSearchForResumesResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *ResumesSearchForResumesResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *ResumesSearchForResumesResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *ResumesSearchForResumesResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *ResumesSearchForResumesResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *ResumesSearchForResumesResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *ResumesSearchForResumesResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetItems

`func (o *ResumesSearchForResumesResponse) GetItems() []ResumesSearchForEmployerAndApplicant`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ResumesSearchForResumesResponse) GetItemsOk() (*[]ResumesSearchForEmployerAndApplicant, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ResumesSearchForResumesResponse) SetItems(v []ResumesSearchForEmployerAndApplicant)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


