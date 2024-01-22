# VacanciesVacancyArchivedFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchivedAt** | **string** | Дата и время архивации вакансии | 
**Counters** | [**VacancyCountersForArchivedOrHidden**](VacancyCountersForArchivedOrHidden.md) |  | 
**CreatedAt** | **string** | Дата и время публикации вакансии | 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций.  Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60;  | [optional] 

## Methods

### NewVacanciesVacancyArchivedFields

`func NewVacanciesVacancyArchivedFields(archivedAt string, counters VacancyCountersForArchivedOrHidden, createdAt string, ) *VacanciesVacancyArchivedFields`

NewVacanciesVacancyArchivedFields instantiates a new VacanciesVacancyArchivedFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyArchivedFieldsWithDefaults

`func NewVacanciesVacancyArchivedFieldsWithDefaults() *VacanciesVacancyArchivedFields`

NewVacanciesVacancyArchivedFieldsWithDefaults instantiates a new VacanciesVacancyArchivedFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchivedAt

`func (o *VacanciesVacancyArchivedFields) GetArchivedAt() string`

GetArchivedAt returns the ArchivedAt field if non-nil, zero value otherwise.

### GetArchivedAtOk

`func (o *VacanciesVacancyArchivedFields) GetArchivedAtOk() (*string, bool)`

GetArchivedAtOk returns a tuple with the ArchivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchivedAt

`func (o *VacanciesVacancyArchivedFields) SetArchivedAt(v string)`

SetArchivedAt sets ArchivedAt field to given value.


### GetCounters

`func (o *VacanciesVacancyArchivedFields) GetCounters() VacancyCountersForArchivedOrHidden`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *VacanciesVacancyArchivedFields) GetCountersOk() (*VacancyCountersForArchivedOrHidden, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *VacanciesVacancyArchivedFields) SetCounters(v VacancyCountersForArchivedOrHidden)`

SetCounters sets Counters field to given value.


### GetCreatedAt

`func (o *VacanciesVacancyArchivedFields) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VacanciesVacancyArchivedFields) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VacanciesVacancyArchivedFields) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetSortPointDistance

`func (o *VacanciesVacancyArchivedFields) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesVacancyArchivedFields) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesVacancyArchivedFields) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesVacancyArchivedFields) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesVacancyArchivedFields) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesVacancyArchivedFields) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


