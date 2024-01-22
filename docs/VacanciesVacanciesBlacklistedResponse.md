# VacanciesVacanciesBlacklistedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**Items** | Pointer to [**[]VacanciesVacanciesBlacklistedItem**](VacanciesVacanciesBlacklistedItem.md) | Массив скрытых вакансий | [optional] 
**LimitReached** | Pointer to **bool** | Превышено ли максимальное количество элементов в списке | [optional] 

## Methods

### NewVacanciesVacanciesBlacklistedResponse

`func NewVacanciesVacanciesBlacklistedResponse(found float32, page float32, pages float32, perPage float32, ) *VacanciesVacanciesBlacklistedResponse`

NewVacanciesVacanciesBlacklistedResponse instantiates a new VacanciesVacanciesBlacklistedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacanciesBlacklistedResponseWithDefaults

`func NewVacanciesVacanciesBlacklistedResponseWithDefaults() *VacanciesVacanciesBlacklistedResponse`

NewVacanciesVacanciesBlacklistedResponseWithDefaults instantiates a new VacanciesVacanciesBlacklistedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *VacanciesVacanciesBlacklistedResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *VacanciesVacanciesBlacklistedResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *VacanciesVacanciesBlacklistedResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *VacanciesVacanciesBlacklistedResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *VacanciesVacanciesBlacklistedResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *VacanciesVacanciesBlacklistedResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *VacanciesVacanciesBlacklistedResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *VacanciesVacanciesBlacklistedResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *VacanciesVacanciesBlacklistedResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *VacanciesVacanciesBlacklistedResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *VacanciesVacanciesBlacklistedResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *VacanciesVacanciesBlacklistedResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetItems

`func (o *VacanciesVacanciesBlacklistedResponse) GetItems() []VacanciesVacanciesBlacklistedItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VacanciesVacanciesBlacklistedResponse) GetItemsOk() (*[]VacanciesVacanciesBlacklistedItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VacanciesVacanciesBlacklistedResponse) SetItems(v []VacanciesVacanciesBlacklistedItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *VacanciesVacanciesBlacklistedResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetLimitReached

`func (o *VacanciesVacanciesBlacklistedResponse) GetLimitReached() bool`

GetLimitReached returns the LimitReached field if non-nil, zero value otherwise.

### GetLimitReachedOk

`func (o *VacanciesVacanciesBlacklistedResponse) GetLimitReachedOk() (*bool, bool)`

GetLimitReachedOk returns a tuple with the LimitReached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitReached

`func (o *VacanciesVacanciesBlacklistedResponse) SetLimitReached(v bool)`

SetLimitReached sets LimitReached field to given value.

### HasLimitReached

`func (o *VacanciesVacanciesBlacklistedResponse) HasLimitReached() bool`

HasLimitReached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


