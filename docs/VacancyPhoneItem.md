# VacancyPhoneItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | Код города | [optional] 
**Comment** | Pointer to **NullableString** | Комментарий (удобное время для звонка по этому номеру) | [optional] 
**Country** | Pointer to **string** | Код страны | [optional] 
**Formatted** | Pointer to **string** | Телефонный номер | [optional] 
**Number** | Pointer to **string** | Телефон | [optional] 

## Methods

### NewVacancyPhoneItem

`func NewVacancyPhoneItem() *VacancyPhoneItem`

NewVacancyPhoneItem instantiates a new VacancyPhoneItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyPhoneItemWithDefaults

`func NewVacancyPhoneItemWithDefaults() *VacancyPhoneItem`

NewVacancyPhoneItemWithDefaults instantiates a new VacancyPhoneItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *VacancyPhoneItem) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *VacancyPhoneItem) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *VacancyPhoneItem) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *VacancyPhoneItem) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetComment

`func (o *VacancyPhoneItem) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *VacancyPhoneItem) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *VacancyPhoneItem) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *VacancyPhoneItem) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *VacancyPhoneItem) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *VacancyPhoneItem) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCountry

`func (o *VacancyPhoneItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *VacancyPhoneItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *VacancyPhoneItem) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *VacancyPhoneItem) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetFormatted

`func (o *VacancyPhoneItem) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *VacancyPhoneItem) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *VacancyPhoneItem) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *VacancyPhoneItem) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.

### GetNumber

`func (o *VacancyPhoneItem) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *VacancyPhoneItem) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *VacancyPhoneItem) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *VacancyPhoneItem) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


