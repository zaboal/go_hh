# VacanciesAdditionalVacancyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counters** | Pointer to [**VacancyCounters**](VacancyCounters.md) |  | [optional] 
**Employment** | Pointer to [**NullableVacancyEmploymentOutput**](VacancyEmploymentOutput.md) |  | [optional] 
**Experience** | Pointer to [**NullableVacancyExperienceOutput**](VacancyExperienceOutput.md) |  | [optional] 
**Snippet** | [**VacancySnippet**](VacancySnippet.md) |  | 
**SortPointDistance** | Pointer to **NullableFloat32** | Расстояние в метрах между центром сортировки (заданной параметрами &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;) и указанным в вакансии адресом. В случае, если в адресе указаны только станции метро, выдается расстояние между центром сортировки и средней геометрической точкой указанных станций.  Значение &#x60;sort_point_distance&#x60; выдается только в случае, если заданы параметры &#x60;sort_point_lat&#x60;, &#x60;sort_point_lng&#x60;, &#x60;order_by&#x3D;distance&#x60; | [optional] 

## Methods

### NewVacanciesAdditionalVacancyFields

`func NewVacanciesAdditionalVacancyFields(snippet VacancySnippet, ) *VacanciesAdditionalVacancyFields`

NewVacanciesAdditionalVacancyFields instantiates a new VacanciesAdditionalVacancyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesAdditionalVacancyFieldsWithDefaults

`func NewVacanciesAdditionalVacancyFieldsWithDefaults() *VacanciesAdditionalVacancyFields`

NewVacanciesAdditionalVacancyFieldsWithDefaults instantiates a new VacanciesAdditionalVacancyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounters

`func (o *VacanciesAdditionalVacancyFields) GetCounters() VacancyCounters`

GetCounters returns the Counters field if non-nil, zero value otherwise.

### GetCountersOk

`func (o *VacanciesAdditionalVacancyFields) GetCountersOk() (*VacancyCounters, bool)`

GetCountersOk returns a tuple with the Counters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounters

`func (o *VacanciesAdditionalVacancyFields) SetCounters(v VacancyCounters)`

SetCounters sets Counters field to given value.

### HasCounters

`func (o *VacanciesAdditionalVacancyFields) HasCounters() bool`

HasCounters returns a boolean if a field has been set.

### GetEmployment

`func (o *VacanciesAdditionalVacancyFields) GetEmployment() VacancyEmploymentOutput`

GetEmployment returns the Employment field if non-nil, zero value otherwise.

### GetEmploymentOk

`func (o *VacanciesAdditionalVacancyFields) GetEmploymentOk() (*VacancyEmploymentOutput, bool)`

GetEmploymentOk returns a tuple with the Employment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployment

`func (o *VacanciesAdditionalVacancyFields) SetEmployment(v VacancyEmploymentOutput)`

SetEmployment sets Employment field to given value.

### HasEmployment

`func (o *VacanciesAdditionalVacancyFields) HasEmployment() bool`

HasEmployment returns a boolean if a field has been set.

### SetEmploymentNil

`func (o *VacanciesAdditionalVacancyFields) SetEmploymentNil(b bool)`

 SetEmploymentNil sets the value for Employment to be an explicit nil

### UnsetEmployment
`func (o *VacanciesAdditionalVacancyFields) UnsetEmployment()`

UnsetEmployment ensures that no value is present for Employment, not even an explicit nil
### GetExperience

`func (o *VacanciesAdditionalVacancyFields) GetExperience() VacancyExperienceOutput`

GetExperience returns the Experience field if non-nil, zero value otherwise.

### GetExperienceOk

`func (o *VacanciesAdditionalVacancyFields) GetExperienceOk() (*VacancyExperienceOutput, bool)`

GetExperienceOk returns a tuple with the Experience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperience

`func (o *VacanciesAdditionalVacancyFields) SetExperience(v VacancyExperienceOutput)`

SetExperience sets Experience field to given value.

### HasExperience

`func (o *VacanciesAdditionalVacancyFields) HasExperience() bool`

HasExperience returns a boolean if a field has been set.

### SetExperienceNil

`func (o *VacanciesAdditionalVacancyFields) SetExperienceNil(b bool)`

 SetExperienceNil sets the value for Experience to be an explicit nil

### UnsetExperience
`func (o *VacanciesAdditionalVacancyFields) UnsetExperience()`

UnsetExperience ensures that no value is present for Experience, not even an explicit nil
### GetSnippet

`func (o *VacanciesAdditionalVacancyFields) GetSnippet() VacancySnippet`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *VacanciesAdditionalVacancyFields) GetSnippetOk() (*VacancySnippet, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *VacanciesAdditionalVacancyFields) SetSnippet(v VacancySnippet)`

SetSnippet sets Snippet field to given value.


### GetSortPointDistance

`func (o *VacanciesAdditionalVacancyFields) GetSortPointDistance() float32`

GetSortPointDistance returns the SortPointDistance field if non-nil, zero value otherwise.

### GetSortPointDistanceOk

`func (o *VacanciesAdditionalVacancyFields) GetSortPointDistanceOk() (*float32, bool)`

GetSortPointDistanceOk returns a tuple with the SortPointDistance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortPointDistance

`func (o *VacanciesAdditionalVacancyFields) SetSortPointDistance(v float32)`

SetSortPointDistance sets SortPointDistance field to given value.

### HasSortPointDistance

`func (o *VacanciesAdditionalVacancyFields) HasSortPointDistance() bool`

HasSortPointDistance returns a boolean if a field has been set.

### SetSortPointDistanceNil

`func (o *VacanciesAdditionalVacancyFields) SetSortPointDistanceNil(b bool)`

 SetSortPointDistanceNil sets the value for SortPointDistance to be an explicit nil

### UnsetSortPointDistance
`func (o *VacanciesAdditionalVacancyFields) UnsetSortPointDistance()`

UnsetSortPointDistance ensures that no value is present for SortPointDistance, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


