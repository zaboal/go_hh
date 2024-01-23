# VacancyDraftTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Тест, который будет добавлен в вакансию | [optional] 
**Required** | Pointer to **NullableBool** | Флаг обязательности прохождения теста при отклике на вакансию | [optional] 

## Methods

### NewVacancyDraftTest

`func NewVacancyDraftTest() *VacancyDraftTest`

NewVacancyDraftTest instantiates a new VacancyDraftTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftTestWithDefaults

`func NewVacancyDraftTestWithDefaults() *VacancyDraftTest`

NewVacancyDraftTestWithDefaults instantiates a new VacancyDraftTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyDraftTest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyDraftTest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyDraftTest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacancyDraftTest) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *VacancyDraftTest) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *VacancyDraftTest) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetRequired

`func (o *VacancyDraftTest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *VacancyDraftTest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *VacancyDraftTest) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *VacancyDraftTest) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### SetRequiredNil

`func (o *VacancyDraftTest) SetRequiredNil(b bool)`

 SetRequiredNil sets the value for Required to be an explicit nil

### UnsetRequired
`func (o *VacancyDraftTest) UnsetRequired()`

UnsetRequired ensures that no value is present for Required, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


