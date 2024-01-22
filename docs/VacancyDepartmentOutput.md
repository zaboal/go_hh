# VacancyDepartmentOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Департамент [из справочника](https://api.hh.ru/openapi/redoc#tag/Informaciya-o-rabotodatele/operation/get-employer-departments), от имени которого размещается вакансия (если данная возможность доступна для компании) | [optional] 
**Name** | Pointer to **string** | Название департамента работодателя | [optional] 

## Methods

### NewVacancyDepartmentOutput

`func NewVacancyDepartmentOutput() *VacancyDepartmentOutput`

NewVacancyDepartmentOutput instantiates a new VacancyDepartmentOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDepartmentOutputWithDefaults

`func NewVacancyDepartmentOutputWithDefaults() *VacancyDepartmentOutput`

NewVacancyDepartmentOutputWithDefaults instantiates a new VacancyDepartmentOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyDepartmentOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyDepartmentOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyDepartmentOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacancyDepartmentOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VacancyDepartmentOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyDepartmentOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyDepartmentOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacancyDepartmentOutput) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


