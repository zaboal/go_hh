# VacancyLanguageOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор языка. Значения из справочника [/languages](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Level** | [**VacancyLanguageOutputAllOfLevel**](VacancyLanguageOutputAllOfLevel.md) |  | 
**Name** | **string** | Название языка | 

## Methods

### NewVacancyLanguageOutput

`func NewVacancyLanguageOutput(id string, level VacancyLanguageOutputAllOfLevel, name string, ) *VacancyLanguageOutput`

NewVacancyLanguageOutput instantiates a new VacancyLanguageOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyLanguageOutputWithDefaults

`func NewVacancyLanguageOutputWithDefaults() *VacancyLanguageOutput`

NewVacancyLanguageOutputWithDefaults instantiates a new VacancyLanguageOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyLanguageOutput) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyLanguageOutput) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyLanguageOutput) SetId(v string)`

SetId sets Id field to given value.


### GetLevel

`func (o *VacancyLanguageOutput) GetLevel() VacancyLanguageOutputAllOfLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *VacancyLanguageOutput) GetLevelOk() (*VacancyLanguageOutputAllOfLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *VacancyLanguageOutput) SetLevel(v VacancyLanguageOutputAllOfLevel)`

SetLevel sets Level field to given value.


### GetName

`func (o *VacancyLanguageOutput) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VacancyLanguageOutput) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VacancyLanguageOutput) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


