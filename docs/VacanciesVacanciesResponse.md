# VacanciesVacanciesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]VacanciesVacanciesItem**](VacanciesVacanciesItem.md) | Список вакансий | 
**Found** | **float32** | Найдено результатов | 
**Page** | **float32** | Номер страницы | 
**Pages** | **float32** | Всего страниц | 
**PerPage** | **float32** | Результатов на странице | 
**Clusters** | Pointer to [**[]VacanciesClusterItem**](VacanciesClusterItem.md) | Массив [кластеров поиска](#tag/Poisk-vakansij/Klastery-v-poiske-vakansij) | [optional] 
**Arguments** | Pointer to [**[]VacanciesArgumentItem**](VacanciesArgumentItem.md) | Массив параметров поиска, переданных в запросе.  Возвращается только если в запросе передан параметр &#x60;describe_arguments&#x3D;true&#x60;. В массиве выдаются только те параметры, которые влияют на поиск вакансий. Неизвестные параметры игнорируются. Элемент списка с одним значением &#x60;argument&#x60; может повторяться несколько раз, если параметр имеет несколько значений  | [optional] 
**AlternateUrl** | Pointer to **NullableString** | Ссылка на вакансию | [optional] 
**Fixes** | Pointer to [**NullableVacanciesFixes**](VacanciesFixes.md) |  | [optional] 
**Suggests** | Pointer to [**NullableVacanciesSuggests**](VacanciesSuggests.md) |  | [optional] 

## Methods

### NewVacanciesVacanciesResponse

`func NewVacanciesVacanciesResponse(items []VacanciesVacanciesItem, found float32, page float32, pages float32, perPage float32, ) *VacanciesVacanciesResponse`

NewVacanciesVacanciesResponse instantiates a new VacanciesVacanciesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacanciesResponseWithDefaults

`func NewVacanciesVacanciesResponseWithDefaults() *VacanciesVacanciesResponse`

NewVacanciesVacanciesResponseWithDefaults instantiates a new VacanciesVacanciesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *VacanciesVacanciesResponse) GetItems() []VacanciesVacanciesItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *VacanciesVacanciesResponse) GetItemsOk() (*[]VacanciesVacanciesItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *VacanciesVacanciesResponse) SetItems(v []VacanciesVacanciesItem)`

SetItems sets Items field to given value.


### GetFound

`func (o *VacanciesVacanciesResponse) GetFound() float32`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *VacanciesVacanciesResponse) GetFoundOk() (*float32, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *VacanciesVacanciesResponse) SetFound(v float32)`

SetFound sets Found field to given value.


### GetPage

`func (o *VacanciesVacanciesResponse) GetPage() float32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *VacanciesVacanciesResponse) GetPageOk() (*float32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *VacanciesVacanciesResponse) SetPage(v float32)`

SetPage sets Page field to given value.


### GetPages

`func (o *VacanciesVacanciesResponse) GetPages() float32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *VacanciesVacanciesResponse) GetPagesOk() (*float32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *VacanciesVacanciesResponse) SetPages(v float32)`

SetPages sets Pages field to given value.


### GetPerPage

`func (o *VacanciesVacanciesResponse) GetPerPage() float32`

GetPerPage returns the PerPage field if non-nil, zero value otherwise.

### GetPerPageOk

`func (o *VacanciesVacanciesResponse) GetPerPageOk() (*float32, bool)`

GetPerPageOk returns a tuple with the PerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPage

`func (o *VacanciesVacanciesResponse) SetPerPage(v float32)`

SetPerPage sets PerPage field to given value.


### GetClusters

`func (o *VacanciesVacanciesResponse) GetClusters() []VacanciesClusterItem`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *VacanciesVacanciesResponse) GetClustersOk() (*[]VacanciesClusterItem, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *VacanciesVacanciesResponse) SetClusters(v []VacanciesClusterItem)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *VacanciesVacanciesResponse) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### SetClustersNil

`func (o *VacanciesVacanciesResponse) SetClustersNil(b bool)`

 SetClustersNil sets the value for Clusters to be an explicit nil

### UnsetClusters
`func (o *VacanciesVacanciesResponse) UnsetClusters()`

UnsetClusters ensures that no value is present for Clusters, not even an explicit nil
### GetArguments

`func (o *VacanciesVacanciesResponse) GetArguments() []VacanciesArgumentItem`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *VacanciesVacanciesResponse) GetArgumentsOk() (*[]VacanciesArgumentItem, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *VacanciesVacanciesResponse) SetArguments(v []VacanciesArgumentItem)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *VacanciesVacanciesResponse) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### SetArgumentsNil

`func (o *VacanciesVacanciesResponse) SetArgumentsNil(b bool)`

 SetArgumentsNil sets the value for Arguments to be an explicit nil

### UnsetArguments
`func (o *VacanciesVacanciesResponse) UnsetArguments()`

UnsetArguments ensures that no value is present for Arguments, not even an explicit nil
### GetAlternateUrl

`func (o *VacanciesVacanciesResponse) GetAlternateUrl() string`

GetAlternateUrl returns the AlternateUrl field if non-nil, zero value otherwise.

### GetAlternateUrlOk

`func (o *VacanciesVacanciesResponse) GetAlternateUrlOk() (*string, bool)`

GetAlternateUrlOk returns a tuple with the AlternateUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternateUrl

`func (o *VacanciesVacanciesResponse) SetAlternateUrl(v string)`

SetAlternateUrl sets AlternateUrl field to given value.

### HasAlternateUrl

`func (o *VacanciesVacanciesResponse) HasAlternateUrl() bool`

HasAlternateUrl returns a boolean if a field has been set.

### SetAlternateUrlNil

`func (o *VacanciesVacanciesResponse) SetAlternateUrlNil(b bool)`

 SetAlternateUrlNil sets the value for AlternateUrl to be an explicit nil

### UnsetAlternateUrl
`func (o *VacanciesVacanciesResponse) UnsetAlternateUrl()`

UnsetAlternateUrl ensures that no value is present for AlternateUrl, not even an explicit nil
### GetFixes

`func (o *VacanciesVacanciesResponse) GetFixes() VacanciesFixes`

GetFixes returns the Fixes field if non-nil, zero value otherwise.

### GetFixesOk

`func (o *VacanciesVacanciesResponse) GetFixesOk() (*VacanciesFixes, bool)`

GetFixesOk returns a tuple with the Fixes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixes

`func (o *VacanciesVacanciesResponse) SetFixes(v VacanciesFixes)`

SetFixes sets Fixes field to given value.

### HasFixes

`func (o *VacanciesVacanciesResponse) HasFixes() bool`

HasFixes returns a boolean if a field has been set.

### SetFixesNil

`func (o *VacanciesVacanciesResponse) SetFixesNil(b bool)`

 SetFixesNil sets the value for Fixes to be an explicit nil

### UnsetFixes
`func (o *VacanciesVacanciesResponse) UnsetFixes()`

UnsetFixes ensures that no value is present for Fixes, not even an explicit nil
### GetSuggests

`func (o *VacanciesVacanciesResponse) GetSuggests() VacanciesSuggests`

GetSuggests returns the Suggests field if non-nil, zero value otherwise.

### GetSuggestsOk

`func (o *VacanciesVacanciesResponse) GetSuggestsOk() (*VacanciesSuggests, bool)`

GetSuggestsOk returns a tuple with the Suggests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggests

`func (o *VacanciesVacanciesResponse) SetSuggests(v VacanciesSuggests)`

SetSuggests sets Suggests field to given value.

### HasSuggests

`func (o *VacanciesVacanciesResponse) HasSuggests() bool`

HasSuggests returns a boolean if a field has been set.

### SetSuggestsNil

`func (o *VacanciesVacanciesResponse) SetSuggestsNil(b bool)`

 SetSuggestsNil sets the value for Suggests to be an explicit nil

### UnsetSuggests
`func (o *VacanciesVacanciesResponse) UnsetSuggests()`

UnsetSuggests ensures that no value is present for Suggests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


