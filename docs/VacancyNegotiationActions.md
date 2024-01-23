# VacancyNegotiationActions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Arguments** | [**[]VacancyArguments**](VacancyArguments.md) | Обязательные и дополнительные аргументы для запроса | 
**Enabled** | **bool** | Возможно ли совершить действие | 
**Id** | **string** | Идентификатор действия | 
**Method** | **string** | HTTP метод, который необходимо выполнить | 
**Name** | **string** | Название действия | 
**ResultingEmployerState** | Pointer to [**NullableIncludesIdName**](IncludesIdName.md) |  | [optional] 
**Templates** | [**[]VacancyTemplates**](VacancyTemplates.md) | Шаблоны писем | 
**Url** | **string** | URL, на который необходимо выполнить запрос | 

## Methods

### NewVacancyNegotiationActions

`func NewVacancyNegotiationActions(arguments []VacancyArguments, enabled bool, id string, method string, name string, templates []VacancyTemplates, url string, ) *VacancyNegotiationActions`

NewVacancyNegotiationActions instantiates a new VacancyNegotiationActions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyNegotiationActionsWithDefaults

`func NewVacancyNegotiationActionsWithDefaults() *VacancyNegotiationActions`

NewVacancyNegotiationActionsWithDefaults instantiates a new VacancyNegotiationActions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArguments

`func (o *VacancyNegotiationActions) GetArguments() []VacancyArguments`

GetArguments returns the Arguments field if non-nil, zero value otherwise.

### GetArgumentsOk

`func (o *VacancyNegotiationActions) GetArgumentsOk() (*[]VacancyArguments, bool)`

GetArgumentsOk returns a tuple with the Arguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArguments

`func (o *VacancyNegotiationActions) SetArguments(v []VacancyArguments)`

SetArguments sets Arguments field to given value.


### GetEnabled

`func (o *VacancyNegotiationActions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VacancyNegotiationActions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VacancyNegotiationActions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetId

`func (o *VacancyNegotiationActions) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyNegotiationActions) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyNegotiationActions) SetId(v string)`

SetId sets Id field to given value.


### GetMethod

`func (o *VacancyNegotiationActions) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *VacancyNegotiationActions) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *VacancyNegotiationActions) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetName

`func (o *VacancyNegotiationActions) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyNegotiationActions) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyNegotiationActions) SetName(v string)`

SetName sets Name field to given value.


### GetResultingEmployerState

`func (o *VacancyNegotiationActions) GetResultingEmployerState() IncludesIdName`

GetResultingEmployerState returns the ResultingEmployerState field if non-nil, zero value otherwise.

### GetResultingEmployerStateOk

`func (o *VacancyNegotiationActions) GetResultingEmployerStateOk() (*IncludesIdName, bool)`

GetResultingEmployerStateOk returns a tuple with the ResultingEmployerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultingEmployerState

`func (o *VacancyNegotiationActions) SetResultingEmployerState(v IncludesIdName)`

SetResultingEmployerState sets ResultingEmployerState field to given value.

### HasResultingEmployerState

`func (o *VacancyNegotiationActions) HasResultingEmployerState() bool`

HasResultingEmployerState returns a boolean if a field has been set.

### SetResultingEmployerStateNil

`func (o *VacancyNegotiationActions) SetResultingEmployerStateNil(b bool)`

 SetResultingEmployerStateNil sets the value for ResultingEmployerState to be an explicit nil

### UnsetResultingEmployerState
`func (o *VacancyNegotiationActions) UnsetResultingEmployerState()`

UnsetResultingEmployerState ensures that no value is present for ResultingEmployerState, not even an explicit nil
### GetTemplates

`func (o *VacancyNegotiationActions) GetTemplates() []VacancyTemplates`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *VacancyNegotiationActions) GetTemplatesOk() (*[]VacancyTemplates, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *VacancyNegotiationActions) SetTemplates(v []VacancyTemplates)`

SetTemplates sets Templates field to given value.


### GetUrl

`func (o *VacancyNegotiationActions) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *VacancyNegotiationActions) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *VacancyNegotiationActions) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


