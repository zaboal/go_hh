# VacancyDraftAssignedManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullName** | Pointer to **string** | ФИО | [optional] 
**Id** | Pointer to **string** | Идентификатор рабочего аккаунта менеджера, которому перейдет вакансия после публикации | [optional] 

## Methods

### NewVacancyDraftAssignedManager

`func NewVacancyDraftAssignedManager() *VacancyDraftAssignedManager`

NewVacancyDraftAssignedManager instantiates a new VacancyDraftAssignedManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftAssignedManagerWithDefaults

`func NewVacancyDraftAssignedManagerWithDefaults() *VacancyDraftAssignedManager`

NewVacancyDraftAssignedManagerWithDefaults instantiates a new VacancyDraftAssignedManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullName

`func (o *VacancyDraftAssignedManager) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *VacancyDraftAssignedManager) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *VacancyDraftAssignedManager) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *VacancyDraftAssignedManager) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetId

`func (o *VacancyDraftAssignedManager) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyDraftAssignedManager) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyDraftAssignedManager) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VacancyDraftAssignedManager) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


