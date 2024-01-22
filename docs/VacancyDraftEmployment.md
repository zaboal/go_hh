# VacancyDraftEmployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Тип занятости из [справочника employment](#tag/Obshie-spravochniki/operation/get-dictionaries). По умолчанию или ошибке в поле ставится full | [optional] 

## Methods

### NewVacancyDraftEmployment

`func NewVacancyDraftEmployment() *VacancyDraftEmployment`

NewVacancyDraftEmployment instantiates a new VacancyDraftEmployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftEmploymentWithDefaults

`func NewVacancyDraftEmploymentWithDefaults() *VacancyDraftEmployment`

NewVacancyDraftEmploymentWithDefaults instantiates a new VacancyDraftEmployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyDraftEmployment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyDraftEmployment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyDraftEmployment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacancyDraftEmployment) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VacancyDraftEmployment) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VacancyDraftEmployment) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


