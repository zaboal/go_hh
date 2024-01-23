# VacancySalary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **NullableString** | Код валюты из [справочника currency](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**From** | Pointer to **NullableInt32** | Нижняя граница зарплаты | [optional] 
**Gross** | Pointer to **NullableBool** | Признак что границы зарплаты указаны до вычета налогов | [optional] 
**To** | Pointer to **NullableInt32** | Верхняя граница зарплаты | [optional] 

## Methods

### NewVacancySalary

`func NewVacancySalary() *VacancySalary`

NewVacancySalary instantiates a new VacancySalary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancySalaryWithDefaults

`func NewVacancySalaryWithDefaults() *VacancySalary`

NewVacancySalaryWithDefaults instantiates a new VacancySalary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *VacancySalary) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *VacancySalary) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *VacancySalary) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *VacancySalary) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *VacancySalary) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *VacancySalary) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetFrom

`func (o *VacancySalary) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *VacancySalary) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *VacancySalary) SetFrom(v int32)`

SetFrom sets From field to given value.

### HasFrom

`func (o *VacancySalary) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *VacancySalary) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *VacancySalary) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetGross

`func (o *VacancySalary) GetGross() bool`

GetGross returns the Gross field if non-nil, zero value otherwise.

### GetGrossOk

`func (o *VacancySalary) GetGrossOk() (*bool, bool)`

GetGrossOk returns a tuple with the Gross field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGross

`func (o *VacancySalary) SetGross(v bool)`

SetGross sets Gross field to given value.

### HasGross

`func (o *VacancySalary) HasGross() bool`

HasGross returns a boolean if a field has been set.

### SetGrossNil

`func (o *VacancySalary) SetGrossNil(b bool)`

 SetGrossNil sets the value for Gross to be an explicit nil

### UnsetGross
`func (o *VacancySalary) UnsetGross()`

UnsetGross ensures that no value is present for Gross, not even an explicit nil
### GetTo

`func (o *VacancySalary) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *VacancySalary) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *VacancySalary) SetTo(v int32)`

SetTo sets To field to given value.

### HasTo

`func (o *VacancySalary) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *VacancySalary) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *VacancySalary) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


