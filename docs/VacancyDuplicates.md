# VacancyDuplicates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **float32** | Общее количество дубликатов вакансии | 
**HasDuplicates** | **bool** | Существуют ли дубликаты вакансии | 
**Items** | **[]float32** | Список идентификаторов дубликатов вакансии | 

## Methods

### NewVacancyDuplicates

`func NewVacancyDuplicates(found float32, hasDuplicates bool, items []float32, ) *VacancyDuplicates`

NewVacancyDuplicates instantiates a new VacancyDuplicates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDuplicatesWithDefaults

`func NewVacancyDuplicatesWithDefaults() *VacancyDuplicates`

NewVacancyDuplicatesWithDefaults instantiates a new VacancyDuplicates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *VacancyDuplicates) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *VacancyDuplicates) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *VacancyDuplicates) SetFound(v float32)`

SetFound sets Found field to given value.


### GetHasDuplicates

`func (o *VacancyDuplicates) GetHasDuplicates() bool`

GetHasDuplicates returns the HasDuplicates field if non-nil, zero value otherwise.

### GetHasDuplicatesOk

`func (o *VacancyDuplicates) GetHasDuplicatesOk() (*bool, bool)`

GetHasDuplicatesOk returns a tuple with the HasDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDuplicates

`func (o *VacancyDuplicates) SetHasDuplicates(v bool)`

SetHasDuplicates sets HasDuplicates field to given value.


### GetItems

`func (o *VacancyDuplicates) GetItems() []float32`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VacancyDuplicates) GetItemsOk() (*[]float32, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VacancyDuplicates) SetItems(v []float32)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


