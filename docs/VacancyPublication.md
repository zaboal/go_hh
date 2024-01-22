# VacancyPublication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreasUrl** | **string** | URL на список регионов, в которых можно опубликовать вакансию данного типа. Список возвращается в древовидной структуре и публикация вакансий возможна только в конечных (листовых) узлах дерева. Они помечены флагом &#x60;can_publish&#x3D;true&#x60; | 
**Count** | **float32** | Количество публикаций в регионе, доступных работодателю | 
**Name** | **string** | Название региона | 

## Methods

### NewVacancyPublication

`func NewVacancyPublication(areasUrl string, count float32, name string, ) *VacancyPublication`

NewVacancyPublication instantiates a new VacancyPublication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyPublicationWithDefaults

`func NewVacancyPublicationWithDefaults() *VacancyPublication`

NewVacancyPublicationWithDefaults instantiates a new VacancyPublication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreasUrl

`func (o *VacancyPublication) GetAreasUrl() string`

GetAreasUrl returns the AreasUrl field if non-nil, zero value otherwise.

### GetAreasUrlOk

`func (o *VacancyPublication) GetAreasUrlOk() (*string, bool)`

GetAreasUrlOk returns a tuple with the AreasUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreasUrl

`func (o *VacancyPublication) SetAreasUrl(v string)`

SetAreasUrl sets AreasUrl field to given value.


### GetCount

`func (o *VacancyPublication) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VacancyPublication) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VacancyPublication) SetCount(v float32)`

SetCount sets Count field to given value.


### GetName

`func (o *VacancyPublication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyPublication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyPublication) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


