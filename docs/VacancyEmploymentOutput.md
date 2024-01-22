# VacancyEmploymentOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Тип занятости из [справочника employment](#tag/Obshie-spravochniki/operation/get-dictionaries) | [optional] 
**Name** | Pointer to **string** | Название типа занятости | [optional] 

## Methods

### NewVacancyEmploymentOutput

`func NewVacancyEmploymentOutput() *VacancyEmploymentOutput`

NewVacancyEmploymentOutput instantiates a new VacancyEmploymentOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyEmploymentOutputWithDefaults

`func NewVacancyEmploymentOutputWithDefaults() *VacancyEmploymentOutput`

NewVacancyEmploymentOutputWithDefaults instantiates a new VacancyEmploymentOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyEmploymentOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyEmploymentOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyEmploymentOutput) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacancyEmploymentOutput) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VacancyEmploymentOutput) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VacancyEmploymentOutput) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *VacancyEmploymentOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyEmploymentOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyEmploymentOutput) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VacancyEmploymentOutput) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


