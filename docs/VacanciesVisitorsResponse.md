# VacanciesVisitorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]VacanciesVisitorsVisitorItemsItemsInner**](VacanciesVisitorsVisitorItemsItemsInner.md) | Список сокращенных представлений резюме | 
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**HiddenOnPage** | **float32** | Количество удаленных или скрытых соискателями резюме на странице | 

## Methods

### NewVacanciesVisitorsResponse

`func NewVacanciesVisitorsResponse(items []VacanciesVisitorsVisitorItemsItemsInner, found float32, page float32, pages float32, perPage float32, hiddenOnPage float32, ) *VacanciesVisitorsResponse`

NewVacanciesVisitorsResponse instantiates a new VacanciesVisitorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVisitorsResponseWithDefaults

`func NewVacanciesVisitorsResponseWithDefaults() *VacanciesVisitorsResponse`

NewVacanciesVisitorsResponseWithDefaults instantiates a new VacanciesVisitorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *VacanciesVisitorsResponse) GetItems() []VacanciesVisitorsVisitorItemsItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VacanciesVisitorsResponse) GetItemsOk() (*[]VacanciesVisitorsVisitorItemsItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VacanciesVisitorsResponse) SetItems(v []VacanciesVisitorsVisitorItemsItemsInner)`

SetItems sets Items field to given value.


### GetFound

`func (o *VacanciesVisitorsResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *VacanciesVisitorsResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *VacanciesVisitorsResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *VacanciesVisitorsResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *VacanciesVisitorsResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *VacanciesVisitorsResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *VacanciesVisitorsResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *VacanciesVisitorsResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *VacanciesVisitorsResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *VacanciesVisitorsResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *VacanciesVisitorsResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *VacanciesVisitorsResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetHiddenOnPage

`func (o *VacanciesVisitorsResponse) GetHiddenOnPage() float32`

GetHiddenOnPage returns the HiddenOnPage field if non-nil, zero value otherwise.

### GetHiddenOnPageOk

`func (o *VacanciesVisitorsResponse) GetHiddenOnPageOk() (*float32, bool)`

GetHiddenOnPageOk returns a tuple with the HiddenOnPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenOnPage

`func (o *VacanciesVisitorsResponse) SetHiddenOnPage(v float32)`

SetHiddenOnPage sets HiddenOnPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


