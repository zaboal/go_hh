# VacanciesItemsOfClusterItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **float32** | Количество вакансий в данном элементе кластера | 
**MetroLine** | Pointer to [**NullableIncludesClusterMetroLine**](IncludesClusterMetroLine.md) |  | [optional] 
**MetroStation** | Pointer to [**NullableIncludesClusterMetroStation**](IncludesClusterMetroStation.md) |  | [optional] 
**Name** | **string** | Название элемента кластера | 
**Type** | Pointer to **NullableString** | Тип значения, связанного с группой | [optional] 
**Url** | **string** | Ссылка на поисковую выдачу по данному элементу кластера | 

## Methods

### NewVacanciesItemsOfClusterItem

`func NewVacanciesItemsOfClusterItem(count float32, name string, url string, ) *VacanciesItemsOfClusterItem`

NewVacanciesItemsOfClusterItem instantiates a new VacanciesItemsOfClusterItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesItemsOfClusterItemWithDefaults

`func NewVacanciesItemsOfClusterItemWithDefaults() *VacanciesItemsOfClusterItem`

NewVacanciesItemsOfClusterItemWithDefaults instantiates a new VacanciesItemsOfClusterItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VacanciesItemsOfClusterItem) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VacanciesItemsOfClusterItem) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VacanciesItemsOfClusterItem) SetCount(v float32)`

SetCount sets Count field to given value.


### GetMetroLine

`func (o *VacanciesItemsOfClusterItem) GetMetroLine() IncludesClusterMetroLine`

GetMetroLine returns the MetroLine field if non-nil, zero value otherwise.

### GetMetroLineOk

`func (o *VacanciesItemsOfClusterItem) GetMetroLineOk() (*IncludesClusterMetroLine, bool)`

GetMetroLineOk returns a tuple with the MetroLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroLine

`func (o *VacanciesItemsOfClusterItem) SetMetroLine(v IncludesClusterMetroLine)`

SetMetroLine sets MetroLine field to given value.

### HasMetroLine

`func (o *VacanciesItemsOfClusterItem) HasMetroLine() bool`

HasMetroLine returns a boolean if a field has been set.

### SetMetroLineNil

`func (o *VacanciesItemsOfClusterItem) SetMetroLineNil(b bool)`

 SetMetroLineNil sets the value for MetroLine to be an explicit nil

### UnsetMetroLine
`func (o *VacanciesItemsOfClusterItem) UnsetMetroLine()`

UnsetMetroLine ensures that no value is present for MetroLine, not even an explicit nil
### GetMetroStation

`func (o *VacanciesItemsOfClusterItem) GetMetroStation() IncludesClusterMetroStation`

GetMetroStation returns the MetroStation field if non-nil, zero value otherwise.

### GetMetroStationOk

`func (o *VacanciesItemsOfClusterItem) GetMetroStationOk() (*IncludesClusterMetroStation, bool)`

GetMetroStationOk returns a tuple with the MetroStation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetroStation

`func (o *VacanciesItemsOfClusterItem) SetMetroStation(v IncludesClusterMetroStation)`

SetMetroStation sets MetroStation field to given value.

### HasMetroStation

`func (o *VacanciesItemsOfClusterItem) HasMetroStation() bool`

HasMetroStation returns a boolean if a field has been set.

### SetMetroStationNil

`func (o *VacanciesItemsOfClusterItem) SetMetroStationNil(b bool)`

 SetMetroStationNil sets the value for MetroStation to be an explicit nil

### UnsetMetroStation
`func (o *VacanciesItemsOfClusterItem) UnsetMetroStation()`

UnsetMetroStation ensures that no value is present for MetroStation, not even an explicit nil
### GetName

`func (o *VacanciesItemsOfClusterItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacanciesItemsOfClusterItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacanciesItemsOfClusterItem) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *VacanciesItemsOfClusterItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VacanciesItemsOfClusterItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VacanciesItemsOfClusterItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VacanciesItemsOfClusterItem) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *VacanciesItemsOfClusterItem) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *VacanciesItemsOfClusterItem) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetUrl

`func (o *VacanciesItemsOfClusterItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacanciesItemsOfClusterItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacanciesItemsOfClusterItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


