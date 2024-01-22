# VacancyTypeOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Тип из [справочника vacancy_type](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Name** | Pointer to **string** | Название типа вакансии | [optional] 

## Methods

### NewVacancyTypeOutput

`func NewVacancyTypeOutput(id string, ) *VacancyTypeOutput`

NewVacancyTypeOutput instantiates a new VacancyTypeOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyTypeOutputWithDefaults

`func NewVacancyTypeOutputWithDefaults() *VacancyTypeOutput`

NewVacancyTypeOutputWithDefaults instantiates a new VacancyTypeOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyTypeOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyTypeOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyTypeOutput) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VacancyTypeOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyTypeOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyTypeOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacancyTypeOutput) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


