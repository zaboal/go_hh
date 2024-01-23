# VacanciesVacancyProlongate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**[]VacanciesVacancyProlongateActionsInner**](VacanciesVacancyProlongateActionsInner.md) | Список действий, которые можно предпринять для продления вакансии | 
**ExpiresAt** | **string** | Дата окончания публикации вакансии | 
**Id** | **string** | Идентификатор вакансии | 

## Methods

### NewVacanciesVacancyProlongate

`func NewVacanciesVacancyProlongate(actions []VacanciesVacancyProlongateActionsInner, expiresAt string, id string, ) *VacanciesVacancyProlongate`

NewVacanciesVacancyProlongate instantiates a new VacanciesVacancyProlongate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacanciesVacancyProlongateWithDefaults

`func NewVacanciesVacancyProlongateWithDefaults() *VacanciesVacancyProlongate`

NewVacanciesVacancyProlongateWithDefaults instantiates a new VacanciesVacancyProlongate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *VacanciesVacancyProlongate) GetActions() []VacanciesVacancyProlongateActionsInner`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *VacanciesVacancyProlongate) GetActionsOk() (*[]VacanciesVacancyProlongateActionsInner, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *VacanciesVacancyProlongate) SetActions(v []VacanciesVacancyProlongateActionsInner)`

SetActions sets Actions field to given value.


### GetExpiresAt

`func (o *VacanciesVacancyProlongate) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *VacanciesVacancyProlongate) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *VacanciesVacancyProlongate) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.


### GetId

`func (o *VacanciesVacancyProlongate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacanciesVacancyProlongate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacanciesVacancyProlongate) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


