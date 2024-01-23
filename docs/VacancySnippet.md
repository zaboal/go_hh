# VacancySnippet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requirement** | Pointer to **NullableString** | Отрывок из требований по вакансии, если они найдены в тексте описания | [optional] 
**Responsibility** | Pointer to **NullableString** | Отрывок из обязанностей по вакансии, если они найдены в тексте описания | [optional] 

## Methods

### NewVacancySnippet

`func NewVacancySnippet() *VacancySnippet`

NewVacancySnippet instantiates a new VacancySnippet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancySnippetWithDefaults

`func NewVacancySnippetWithDefaults() *VacancySnippet`

NewVacancySnippetWithDefaults instantiates a new VacancySnippet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirement

`func (o *VacancySnippet) GetRequirement() string`

GetRequirement returns the Requirement field if non-nil, zero value otherwise.

### GetRequirementOk

`func (o *VacancySnippet) GetRequirementOk() (*string, bool)`

GetRequirementOk returns a tuple with the Requirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirement

`func (o *VacancySnippet) SetRequirement(v string)`

SetRequirement sets Requirement field to given value.

### HasRequirement

`func (o *VacancySnippet) HasRequirement() bool`

HasRequirement returns a boolean if a field has been set.

### SetRequirementNil

`func (o *VacancySnippet) SetRequirementNil(b bool)`

 SetRequirementNil sets the value for Requirement to be an explicit nil

### UnsetRequirement
`func (o *VacancySnippet) UnsetRequirement()`

UnsetRequirement ensures that no value is present for Requirement, not even an explicit nil
### GetResponsibility

`func (o *VacancySnippet) GetResponsibility() string`

GetResponsibility returns the Responsibility field if non-nil, zero value otherwise.

### GetResponsibilityOk

`func (o *VacancySnippet) GetResponsibilityOk() (*string, bool)`

GetResponsibilityOk returns a tuple with the Responsibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsibility

`func (o *VacancySnippet) SetResponsibility(v string)`

SetResponsibility sets Responsibility field to given value.

### HasResponsibility

`func (o *VacancySnippet) HasResponsibility() bool`

HasResponsibility returns a boolean if a field has been set.

### SetResponsibilityNil

`func (o *VacancySnippet) SetResponsibilityNil(b bool)`

 SetResponsibilityNil sets the value for Responsibility to be an explicit nil

### UnsetResponsibility
`func (o *VacancySnippet) UnsetResponsibility()`

UnsetResponsibility ensures that no value is present for Responsibility, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


