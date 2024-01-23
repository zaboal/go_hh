# VacanciesArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | Pointer to [**[]VacanciesArgumentItem**](VacanciesArgumentItem.md) | Массив параметров поиска, переданных в запросе.  Возвращается только если в запросе передан параметр &#x60;describe_arguments&#x3D;true&#x60;. В массиве выдаются только те параметры, которые влияют на поиск вакансий. Неизвестные параметры игнорируются. Элемент списка с одним значением &#x60;argument&#x60; может повторяться несколько раз, если параметр имеет несколько значений  | [optional] 

## Methods

### NewVacanciesArguments

`func NewVacanciesArguments() *VacanciesArguments`

NewVacanciesArguments instantiates a new VacanciesArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesArgumentsWithDefaults

`func NewVacanciesArgumentsWithDefaults() *VacanciesArguments`

NewVacanciesArgumentsWithDefaults instantiates a new VacanciesArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *VacanciesArguments) GetArguments() []VacanciesArgumentItem`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *VacanciesArguments) GetArgumentsOk() (*[]VacanciesArgumentItem, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *VacanciesArguments) SetArguments(v []VacanciesArgumentItem)`

SetArguments sets Arguments field to given value.

### HasArguments

`func (o *VacanciesArguments) HasArguments() bool`

HasArguments returns a boolean if a field has been set.

### SetArgumentsNil

`func (o *VacanciesArguments) SetArgumentsNil(b bool)`

 SetArgumentsNil sets the value for Arguments to be an explicit nil

### UnsetArguments
`func (o *VacanciesArguments) UnsetArguments()`

UnsetArguments ensures that no value is present for Arguments, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


