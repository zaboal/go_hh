# VacanciesVacanciesFavoritedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**Items** | Pointer to [**[]VacanciesVacanciesFavoritedItem**](VacanciesVacanciesFavoritedItem.md) | Массив отобранных вакансий | [optional] 

## Methods

### NewVacanciesVacanciesFavoritedResponse

`func NewVacanciesVacanciesFavoritedResponse(found float32, page float32, pages float32, perPage float32, ) *VacanciesVacanciesFavoritedResponse`

NewVacanciesVacanciesFavoritedResponse instantiates a new VacanciesVacanciesFavoritedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacanciesFavoritedResponseWithDefaults

`func NewVacanciesVacanciesFavoritedResponseWithDefaults() *VacanciesVacanciesFavoritedResponse`

NewVacanciesVacanciesFavoritedResponseWithDefaults instantiates a new VacanciesVacanciesFavoritedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *VacanciesVacanciesFavoritedResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *VacanciesVacanciesFavoritedResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *VacanciesVacanciesFavoritedResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *VacanciesVacanciesFavoritedResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *VacanciesVacanciesFavoritedResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *VacanciesVacanciesFavoritedResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *VacanciesVacanciesFavoritedResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *VacanciesVacanciesFavoritedResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *VacanciesVacanciesFavoritedResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *VacanciesVacanciesFavoritedResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *VacanciesVacanciesFavoritedResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *VacanciesVacanciesFavoritedResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetItems

`func (o *VacanciesVacanciesFavoritedResponse) GetItems() []VacanciesVacanciesFavoritedItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VacanciesVacanciesFavoritedResponse) GetItemsOk() (*[]VacanciesVacanciesFavoritedItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VacanciesVacanciesFavoritedResponse) SetItems(v []VacanciesVacanciesFavoritedItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *VacanciesVacanciesFavoritedResponse) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


