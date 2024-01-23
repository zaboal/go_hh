# VacancyLanguage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Идентификатор языка. Значения из справочника [/languages](#tag/Obshie-spravochniki/operation/get-dictionaries) | 
**Level** | [**VacancyLanguageLevel**](VacancyLanguageLevel.md) |  | 

## Methods

### NewVacancyLanguage

`func NewVacancyLanguage(id string, level VacancyLanguageLevel, ) *VacancyLanguage`

NewVacancyLanguage instantiates a new VacancyLanguage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyLanguageWithDefaults

`func NewVacancyLanguageWithDefaults() *VacancyLanguage`

NewVacancyLanguageWithDefaults instantiates a new VacancyLanguage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VacancyLanguage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VacancyLanguage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VacancyLanguage) SetId(v string)`

SetId sets Id field to given value.


### GetLevel

`func (o *VacancyLanguage) GetLevel() VacancyLanguageLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *VacancyLanguage) GetLevelOk() (*VacancyLanguageLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *VacancyLanguage) SetLevel(v VacancyLanguageLevel)`

SetLevel sets Level field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


