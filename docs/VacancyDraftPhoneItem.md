# VacancyDraftPhoneItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | Код города | 
**Comment** | Pointer to **NullableString** | Комментарий (удобное время для звонка по этому номеру) | [optional] 
**Country** | **string** | Код страны | 
**Number** | **string** | Абонентская часть телефонного номера | 
**Formatted** | **string** | Телефонный номер целиком | 

## Methods

### NewVacancyDraftPhoneItem

`func NewVacancyDraftPhoneItem(city string, country string, number string, formatted string, ) *VacancyDraftPhoneItem`

NewVacancyDraftPhoneItem instantiates a new VacancyDraftPhoneItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVacancyDraftPhoneItemWithDefaults

`func NewVacancyDraftPhoneItemWithDefaults() *VacancyDraftPhoneItem`

NewVacancyDraftPhoneItemWithDefaults instantiates a new VacancyDraftPhoneItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *VacancyDraftPhoneItem) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *VacancyDraftPhoneItem) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *VacancyDraftPhoneItem) SetCity(v string)`

SetCity sets City field to given value.


### GetComment

`func (o *VacancyDraftPhoneItem) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *VacancyDraftPhoneItem) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *VacancyDraftPhoneItem) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *VacancyDraftPhoneItem) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *VacancyDraftPhoneItem) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *VacancyDraftPhoneItem) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetCountry

`func (o *VacancyDraftPhoneItem) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *VacancyDraftPhoneItem) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *VacancyDraftPhoneItem) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetNumber

`func (o *VacancyDraftPhoneItem) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *VacancyDraftPhoneItem) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *VacancyDraftPhoneItem) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetFormatted

`func (o *VacancyDraftPhoneItem) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *VacancyDraftPhoneItem) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *VacancyDraftPhoneItem) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


