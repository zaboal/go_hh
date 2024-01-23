# VacancyDraftPhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | Код города | 
**Comment** | Pointer to **NullableString** | Комментарий (удобное время для звонка по этому номеру) | [optional] 
**Country** | **string** | Код страны | 
**Number** | **string** | Абонентская часть телефонного номера | 

## Methods

### NewVacancyDraftPhone

`func NewVacancyDraftPhone(city string, country string, number string, ) *VacancyDraftPhone`

NewVacancyDraftPhone instantiates a new VacancyDraftPhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftPhoneWithDefaults

`func NewVacancyDraftPhoneWithDefaults() *VacancyDraftPhone`

NewVacancyDraftPhoneWithDefaults instantiates a new VacancyDraftPhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *VacancyDraftPhone) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *VacancyDraftPhone) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *VacancyDraftPhone) SetCity(v string)`

SetCity sets City field to given value.


### GetComment

`func (o *VacancyDraftPhone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *VacancyDraftPhone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *VacancyDraftPhone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *VacancyDraftPhone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *VacancyDraftPhone) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *VacancyDraftPhone) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCountry

`func (o *VacancyDraftPhone) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *VacancyDraftPhone) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *VacancyDraftPhone) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetNumber

`func (o *VacancyDraftPhone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *VacancyDraftPhone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *VacancyDraftPhone) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


