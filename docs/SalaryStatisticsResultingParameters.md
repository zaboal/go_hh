# SalaryStatisticsResultingParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Areas** | [**[]IncludesIdName**](IncludesIdName.md) | Коды регионов | 
**EmployeeLevels** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Уровни специалистов | [optional] 
**EmployersCount** | **float32** | Количество работодателей, позиции которых участвуют в выборке | 
**ExcludedAreas** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Исключенные коды регионов | [optional] 
**IndirectCalculation** | Pointer to [**NullableSalaryStatisticsIndirectCalculation**](SalaryStatisticsIndirectCalculation.md) |  | [optional] 
**Industries** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Отрасли | [optional] 
**PositionsCount** | **float32** | Количество позиций, по которым построена выборка | 
**Sources** | **[]string** | Источники данных. Возможные значения:  * &#x60;SALARIES&#x60; — данные из банка зарплат; * &#x60;RESUMES&#x60; — данные из резюме; * &#x60;VACANCIES&#x60; — данные из вакансий  | 
**Specialities** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Профессиональные области и специализаций | [optional] 

## Methods

### NewSalaryStatisticsResultingParameters

`func NewSalaryStatisticsResultingParameters(areas []IncludesIdName, employersCount float32, positionsCount float32, sources []string, ) *SalaryStatisticsResultingParameters`

NewSalaryStatisticsResultingParameters instantiates a new SalaryStatisticsResultingParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalaryStatisticsResultingParametersWithDefaults

`func NewSalaryStatisticsResultingParametersWithDefaults() *SalaryStatisticsResultingParameters`

NewSalaryStatisticsResultingParametersWithDefaults instantiates a new SalaryStatisticsResultingParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreas

`func (o *SalaryStatisticsResultingParameters) GetAreas() []IncludesIdName`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *SalaryStatisticsResultingParameters) GetAreasOk() (*[]IncludesIdName, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *SalaryStatisticsResultingParameters) SetAreas(v []IncludesIdName)`

SetAreas sets Areas field to given value.


### GetEmployeeLevels

`func (o *SalaryStatisticsResultingParameters) GetEmployeeLevels() []IncludesIdName`

GetEmployeeLevels returns the EmployeeLevels field if non-nil, zero value otherwise.

### GetEmployeeLevelsOk

`func (o *SalaryStatisticsResultingParameters) GetEmployeeLevelsOk() (*[]IncludesIdName, bool)`

GetEmployeeLevelsOk returns a tuple with the EmployeeLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployeeLevels

`func (o *SalaryStatisticsResultingParameters) SetEmployeeLevels(v []IncludesIdName)`

SetEmployeeLevels sets EmployeeLevels field to given value.

### HasEmployeeLevels

`func (o *SalaryStatisticsResultingParameters) HasEmployeeLevels() bool`

HasEmployeeLevels returns a boolean if a field has been set.

### SetEmployeeLevelsNil

`func (o *SalaryStatisticsResultingParameters) SetEmployeeLevelsNil(b bool)`

 SetEmployeeLevelsNil sets the value for EmployeeLevels to be an explicit nil

### UnsetEmployeeLevels
`func (o *SalaryStatisticsResultingParameters) UnsetEmployeeLevels()`

UnsetEmployeeLevels ensures that no value is present for EmployeeLevels, not even an explicit nil
### GetEmployersCount

`func (o *SalaryStatisticsResultingParameters) GetEmployersCount() float32`

GetEmployersCount returns the EmployersCount field if non-nil, zero value otherwise.

### GetEmployersCountOk

`func (o *SalaryStatisticsResultingParameters) GetEmployersCountOk() (*float32, bool)`

GetEmployersCountOk returns a tuple with the EmployersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmployersCount

`func (o *SalaryStatisticsResultingParameters) SetEmployersCount(v float32)`

SetEmployersCount sets EmployersCount field to given value.


### GetExcludedAreas

`func (o *SalaryStatisticsResultingParameters) GetExcludedAreas() []IncludesIdName`

GetExcludedAreas returns the ExcludedAreas field if non-nil, zero value otherwise.

### GetExcludedAreasOk

`func (o *SalaryStatisticsResultingParameters) GetExcludedAreasOk() (*[]IncludesIdName, bool)`

GetExcludedAreasOk returns a tuple with the ExcludedAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedAreas

`func (o *SalaryStatisticsResultingParameters) SetExcludedAreas(v []IncludesIdName)`

SetExcludedAreas sets ExcludedAreas field to given value.

### HasExcludedAreas

`func (o *SalaryStatisticsResultingParameters) HasExcludedAreas() bool`

HasExcludedAreas returns a boolean if a field has been set.

### SetExcludedAreasNil

`func (o *SalaryStatisticsResultingParameters) SetExcludedAreasNil(b bool)`

 SetExcludedAreasNil sets the value for ExcludedAreas to be an explicit nil

### UnsetExcludedAreas
`func (o *SalaryStatisticsResultingParameters) UnsetExcludedAreas()`

UnsetExcludedAreas ensures that no value is present for ExcludedAreas, not even an explicit nil
### GetIndirectCalculation

`func (o *SalaryStatisticsResultingParameters) GetIndirectCalculation() SalaryStatisticsIndirectCalculation`

GetIndirectCalculation returns the IndirectCalculation field if non-nil, zero value otherwise.

### GetIndirectCalculationOk

`func (o *SalaryStatisticsResultingParameters) GetIndirectCalculationOk() (*SalaryStatisticsIndirectCalculation, bool)`

GetIndirectCalculationOk returns a tuple with the IndirectCalculation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectCalculation

`func (o *SalaryStatisticsResultingParameters) SetIndirectCalculation(v SalaryStatisticsIndirectCalculation)`

SetIndirectCalculation sets IndirectCalculation field to given value.

### HasIndirectCalculation

`func (o *SalaryStatisticsResultingParameters) HasIndirectCalculation() bool`

HasIndirectCalculation returns a boolean if a field has been set.

### SetIndirectCalculationNil

`func (o *SalaryStatisticsResultingParameters) SetIndirectCalculationNil(b bool)`

 SetIndirectCalculationNil sets the value for IndirectCalculation to be an explicit nil

### UnsetIndirectCalculation
`func (o *SalaryStatisticsResultingParameters) UnsetIndirectCalculation()`

UnsetIndirectCalculation ensures that no value is present for IndirectCalculation, not even an explicit nil
### GetIndustries

`func (o *SalaryStatisticsResultingParameters) GetIndustries() []IncludesIdName`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *SalaryStatisticsResultingParameters) GetIndustriesOk() (*[]IncludesIdName, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *SalaryStatisticsResultingParameters) SetIndustries(v []IncludesIdName)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *SalaryStatisticsResultingParameters) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### SetIndustriesNil

`func (o *SalaryStatisticsResultingParameters) SetIndustriesNil(b bool)`

 SetIndustriesNil sets the value for Industries to be an explicit nil

### UnsetIndustries
`func (o *SalaryStatisticsResultingParameters) UnsetIndustries()`

UnsetIndustries ensures that no value is present for Industries, not even an explicit nil
### GetPositionsCount

`func (o *SalaryStatisticsResultingParameters) GetPositionsCount() float32`

GetPositionsCount returns the PositionsCount field if non-nil, zero value otherwise.

### GetPositionsCountOk

`func (o *SalaryStatisticsResultingParameters) GetPositionsCountOk() (*float32, bool)`

GetPositionsCountOk returns a tuple with the PositionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionsCount

`func (o *SalaryStatisticsResultingParameters) SetPositionsCount(v float32)`

SetPositionsCount sets PositionsCount field to given value.


### GetSources

`func (o *SalaryStatisticsResultingParameters) GetSources() []string`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SalaryStatisticsResultingParameters) GetSourcesOk() (*[]string, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SalaryStatisticsResultingParameters) SetSources(v []string)`

SetSources sets Sources field to given value.


### GetSpecialities

`func (o *SalaryStatisticsResultingParameters) GetSpecialities() []IncludesIdName`

GetSpecialities returns the Specialities field if non-nil, zero value otherwise.

### GetSpecialitiesOk

`func (o *SalaryStatisticsResultingParameters) GetSpecialitiesOk() (*[]IncludesIdName, bool)`

GetSpecialitiesOk returns a tuple with the Specialities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialities

`func (o *SalaryStatisticsResultingParameters) SetSpecialities(v []IncludesIdName)`

SetSpecialities sets Specialities field to given value.

### HasSpecialities

`func (o *SalaryStatisticsResultingParameters) HasSpecialities() bool`

HasSpecialities returns a boolean if a field has been set.

### SetSpecialitiesNil

`func (o *SalaryStatisticsResultingParameters) SetSpecialitiesNil(b bool)`

 SetSpecialitiesNil sets the value for Specialities to be an explicit nil

### UnsetSpecialities
`func (o *SalaryStatisticsResultingParameters) UnsetSpecialities()`

UnsetSpecialities ensures that no value is present for Specialities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


