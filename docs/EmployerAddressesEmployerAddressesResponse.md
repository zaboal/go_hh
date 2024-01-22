# EmployerAddressesEmployerAddressesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]EmployerAddressesEmployerAddressItem**](EmployerAddressesEmployerAddressItem.md) | Список адресов работодателя | 
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 

## Methods

### NewEmployerAddressesEmployerAddressesResponse

`func NewEmployerAddressesEmployerAddressesResponse(items []EmployerAddressesEmployerAddressItem, found float32, page float32, pages float32, perPage float32, ) *EmployerAddressesEmployerAddressesResponse`

NewEmployerAddressesEmployerAddressesResponse instantiates a new EmployerAddressesEmployerAddressesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerAddressesEmployerAddressesResponseWithDefaults

`func NewEmployerAddressesEmployerAddressesResponseWithDefaults() *EmployerAddressesEmployerAddressesResponse`

NewEmployerAddressesEmployerAddressesResponseWithDefaults instantiates a new EmployerAddressesEmployerAddressesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *EmployerAddressesEmployerAddressesResponse) GetItems() []EmployerAddressesEmployerAddressItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *EmployerAddressesEmployerAddressesResponse) GetItemsOk() (*[]EmployerAddressesEmployerAddressItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *EmployerAddressesEmployerAddressesResponse) SetItems(v []EmployerAddressesEmployerAddressItem)`

SetItems sets Items field to given value.


### GetFound

`func (o *EmployerAddressesEmployerAddressesResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *EmployerAddressesEmployerAddressesResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *EmployerAddressesEmployerAddressesResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *EmployerAddressesEmployerAddressesResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *EmployerAddressesEmployerAddressesResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *EmployerAddressesEmployerAddressesResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *EmployerAddressesEmployerAddressesResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *EmployerAddressesEmployerAddressesResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *EmployerAddressesEmployerAddressesResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *EmployerAddressesEmployerAddressesResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *EmployerAddressesEmployerAddressesResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *EmployerAddressesEmployerAddressesResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


