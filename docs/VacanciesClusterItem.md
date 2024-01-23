# VacanciesClusterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор кластера | 
**Items** | [**[]VacanciesItemsOfClusterItem**](VacanciesItemsOfClusterItem.md) | Массив поисковых запросов в данном кластере с указанием дополнительных параметров | 
**Name** | **string** | Название типа кластера | 

## Methods

### NewVacanciesClusterItem

`func NewVacanciesClusterItem(id string, items []VacanciesItemsOfClusterItem, name string, ) *VacanciesClusterItem`

NewVacanciesClusterItem instantiates a new VacanciesClusterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesClusterItemWithDefaults

`func NewVacanciesClusterItemWithDefaults() *VacanciesClusterItem`

NewVacanciesClusterItemWithDefaults instantiates a new VacanciesClusterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacanciesClusterItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesClusterItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesClusterItem) SetId(v string)`

SetId sets Id field to given value.


### GetItems

`func (o *VacanciesClusterItem) GetItems() []VacanciesItemsOfClusterItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VacanciesClusterItem) GetItemsOk() (*[]VacanciesItemsOfClusterItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VacanciesClusterItem) SetItems(v []VacanciesItemsOfClusterItem)`

SetItems sets Items field to given value.


### GetName

`func (o *VacanciesClusterItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesClusterItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesClusterItem) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


