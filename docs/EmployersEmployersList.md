# EmployersEmployersList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **float32** | Найдено результатов | 
**Items** | [**[]EmployersEmployerItem**](EmployersEmployerItem.md) | Найденные работодатели | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 

## Methods

### NewEmployersEmployersList

`func NewEmployersEmployersList(found float32, items []EmployersEmployerItem, page float32, pages float32, perPage float32, ) *EmployersEmployersList`

NewEmployersEmployersList instantiates a new EmployersEmployersList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployersEmployersListWithDefaults

`func NewEmployersEmployersListWithDefaults() *EmployersEmployersList`

NewEmployersEmployersListWithDefaults instantiates a new EmployersEmployersList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *EmployersEmployersList) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *EmployersEmployersList) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *EmployersEmployersList) SetFound(v float32)`

SetFound sets Found field to given value.


### GetItems

`func (o *EmployersEmployersList) GetItems() []EmployersEmployerItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EmployersEmployersList) GetItemsOk() (*[]EmployersEmployerItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EmployersEmployersList) SetItems(v []EmployersEmployerItem)`

SetItems sets Items field to given value.


### GetPage

`func (o *EmployersEmployersList) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EmployersEmployersList) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EmployersEmployersList) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *EmployersEmployersList) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *EmployersEmployersList) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *EmployersEmployersList) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *EmployersEmployersList) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *EmployersEmployersList) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *EmployersEmployersList) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


