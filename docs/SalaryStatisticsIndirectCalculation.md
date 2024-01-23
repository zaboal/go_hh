# SalaryStatisticsIndirectCalculation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IndirectAreas** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Регионы, использованные при получении косвенной оценки | [optional] 
**IndirectEmployeeLevels** | Pointer to [**[]IncludesIdName**](IncludesIdName.md) | Уровни специалистов, включенные в выборку в регионе, использованном при получении косвенной оценки | [optional] 
**IndirectRegionalRatio** | **float32** | Региональный коэффициент, который был использован для получения косвенной оценки зарплат | 

## Methods

### NewSalaryStatisticsIndirectCalculation

`func NewSalaryStatisticsIndirectCalculation(indirectRegionalRatio float32, ) *SalaryStatisticsIndirectCalculation`

NewSalaryStatisticsIndirectCalculation instantiates a new SalaryStatisticsIndirectCalculation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSalaryStatisticsIndirectCalculationWithDefaults

`func NewSalaryStatisticsIndirectCalculationWithDefaults() *SalaryStatisticsIndirectCalculation`

NewSalaryStatisticsIndirectCalculationWithDefaults instantiates a new SalaryStatisticsIndirectCalculation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndirectAreas

`func (o *SalaryStatisticsIndirectCalculation) GetIndirectAreas() []IncludesIdName`

GetIndirectAreas returns the IndirectAreas field if non-nil, zero value otherwise.

### GetIndirectAreasOk

`func (o *SalaryStatisticsIndirectCalculation) GetIndirectAreasOk() (*[]IncludesIdName, bool)`

GetIndirectAreasOk returns a tuple with the IndirectAreas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectAreas

`func (o *SalaryStatisticsIndirectCalculation) SetIndirectAreas(v []IncludesIdName)`

SetIndirectAreas sets IndirectAreas field to given value.

### HasIndirectAreas

`func (o *SalaryStatisticsIndirectCalculation) HasIndirectAreas() bool`

HasIndirectAreas returns a boolean if a field has been set.

### SetIndirectAreasNil

`func (o *SalaryStatisticsIndirectCalculation) SetIndirectAreasNil(b bool)`

 SetIndirectAreasNil sets the value for IndirectAreas to be an explicit nil

### UnsetIndirectAreas
`func (o *SalaryStatisticsIndirectCalculation) UnsetIndirectAreas()`

UnsetIndirectAreas ensures that no value is present for IndirectAreas, not even an explicit nil
### GetIndirectEmployeeLevels

`func (o *SalaryStatisticsIndirectCalculation) GetIndirectEmployeeLevels() []IncludesIdName`

GetIndirectEmployeeLevels returns the IndirectEmployeeLevels field if non-nil, zero value otherwise.

### GetIndirectEmployeeLevelsOk

`func (o *SalaryStatisticsIndirectCalculation) GetIndirectEmployeeLevelsOk() (*[]IncludesIdName, bool)`

GetIndirectEmployeeLevelsOk returns a tuple with the IndirectEmployeeLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectEmployeeLevels

`func (o *SalaryStatisticsIndirectCalculation) SetIndirectEmployeeLevels(v []IncludesIdName)`

SetIndirectEmployeeLevels sets IndirectEmployeeLevels field to given value.

### HasIndirectEmployeeLevels

`func (o *SalaryStatisticsIndirectCalculation) HasIndirectEmployeeLevels() bool`

HasIndirectEmployeeLevels returns a boolean if a field has been set.

### SetIndirectEmployeeLevelsNil

`func (o *SalaryStatisticsIndirectCalculation) SetIndirectEmployeeLevelsNil(b bool)`

 SetIndirectEmployeeLevelsNil sets the value for IndirectEmployeeLevels to be an explicit nil

### UnsetIndirectEmployeeLevels
`func (o *SalaryStatisticsIndirectCalculation) UnsetIndirectEmployeeLevels()`

UnsetIndirectEmployeeLevels ensures that no value is present for IndirectEmployeeLevels, not even an explicit nil
### GetIndirectRegionalRatio

`func (o *SalaryStatisticsIndirectCalculation) GetIndirectRegionalRatio() float32`

GetIndirectRegionalRatio returns the IndirectRegionalRatio field if non-nil, zero value otherwise.

### GetIndirectRegionalRatioOk

`func (o *SalaryStatisticsIndirectCalculation) GetIndirectRegionalRatioOk() (*float32, bool)`

GetIndirectRegionalRatioOk returns a tuple with the IndirectRegionalRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndirectRegionalRatio

`func (o *SalaryStatisticsIndirectCalculation) SetIndirectRegionalRatio(v float32)`

SetIndirectRegionalRatio sets IndirectRegionalRatio field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


