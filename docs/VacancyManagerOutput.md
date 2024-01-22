# VacancyManagerOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | Имя менеджера | 
**Id** | **string** | Контактное лицо (менеджер) по размещаемой вакансии, по умолчанию текущий пользователь | 
**LastName** | **string** | Фамилия менеджера | 
**MiddleName** | **NullableString** | Второе имя менеджера | 

## Methods

### NewVacancyManagerOutput

`func NewVacancyManagerOutput(firstName string, id string, lastName string, middleName NullableString, ) *VacancyManagerOutput`

NewVacancyManagerOutput instantiates a new VacancyManagerOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyManagerOutputWithDefaults

`func NewVacancyManagerOutputWithDefaults() *VacancyManagerOutput`

NewVacancyManagerOutputWithDefaults instantiates a new VacancyManagerOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *VacancyManagerOutput) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *VacancyManagerOutput) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *VacancyManagerOutput) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetId

`func (o *VacancyManagerOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyManagerOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyManagerOutput) SetId(v string)`

SetId sets Id field to given value.


### GetLastName

`func (o *VacancyManagerOutput) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *VacancyManagerOutput) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *VacancyManagerOutput) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetMiddleName

`func (o *VacancyManagerOutput) GetMiddleName() string`

GetMiddleName returns the MiddleName field if non-nil, zero value otherwise.

### GetMiddleNameOk

`func (o *VacancyManagerOutput) GetMiddleNameOk() (*string, bool)`

GetMiddleNameOk returns a tuple with the MiddleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddleName

`func (o *VacancyManagerOutput) SetMiddleName(v string)`

SetMiddleName sets MiddleName field to given value.


### SetMiddleNameNil

`func (o *VacancyManagerOutput) SetMiddleNameNil(b bool)`

 SetMiddleNameNil sets the value for MiddleName to be an explicit nil

### UnsetMiddleName
`func (o *VacancyManagerOutput) UnsetMiddleName()`

UnsetMiddleName ensures that no value is present for MiddleName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


