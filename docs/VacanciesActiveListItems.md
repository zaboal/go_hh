# VacanciesActiveListItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**Items** | [**[]VacanciesActiveListItem**](VacanciesActiveListItem.md) | Список опубликованных вакансий | 

## Methods

### NewVacanciesActiveListItems

`func NewVacanciesActiveListItems(found float32, page float32, pages float32, perPage float32, items []VacanciesActiveListItem, ) *VacanciesActiveListItems`

NewVacanciesActiveListItems instantiates a new VacanciesActiveListItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesActiveListItemsWithDefaults

`func NewVacanciesActiveListItemsWithDefaults() *VacanciesActiveListItems`

NewVacanciesActiveListItemsWithDefaults instantiates a new VacanciesActiveListItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *VacanciesActiveListItems) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *VacanciesActiveListItems) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *VacanciesActiveListItems) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *VacanciesActiveListItems) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *VacanciesActiveListItems) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *VacanciesActiveListItems) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *VacanciesActiveListItems) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *VacanciesActiveListItems) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *VacanciesActiveListItems) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *VacanciesActiveListItems) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *VacanciesActiveListItems) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *VacanciesActiveListItems) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetItems

`func (o *VacanciesActiveListItems) GetItems() []VacanciesActiveListItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VacanciesActiveListItems) GetItemsOk() (*[]VacanciesActiveListItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VacanciesActiveListItems) SetItems(v []VacanciesActiveListItem)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


