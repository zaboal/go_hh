# EmployerManagersPhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | Код города | 
**Comment** | Pointer to **string** | Комментарий | [optional] 
**Country** | **string** | Код страны | 
**Formatted** | Pointer to **string** | Номер телефона отформатированный | [optional] 
**Number** | **string** | Номер | 

## Methods

### NewEmployerManagersPhone

`func NewEmployerManagersPhone(city string, country string, number string, ) *EmployerManagersPhone`

NewEmployerManagersPhone instantiates a new EmployerManagersPhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployerManagersPhoneWithDefaults

`func NewEmployerManagersPhoneWithDefaults() *EmployerManagersPhone`

NewEmployerManagersPhoneWithDefaults instantiates a new EmployerManagersPhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *EmployerManagersPhone) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *EmployerManagersPhone) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *EmployerManagersPhone) SetCity(v string)`

SetCity sets City field to given value.


### GetComment

`func (o *EmployerManagersPhone) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *EmployerManagersPhone) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *EmployerManagersPhone) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *EmployerManagersPhone) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCountry

`func (o *EmployerManagersPhone) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *EmployerManagersPhone) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *EmployerManagersPhone) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetFormatted

`func (o *EmployerManagersPhone) GetFormatted() string`

GetFormatted returns the Formatted field if non-nil, zero value otherwise.

### GetFormattedOk

`func (o *EmployerManagersPhone) GetFormattedOk() (*string, bool)`

GetFormattedOk returns a tuple with the Formatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatted

`func (o *EmployerManagersPhone) SetFormatted(v string)`

SetFormatted sets Formatted field to given value.

### HasFormatted

`func (o *EmployerManagersPhone) HasFormatted() bool`

HasFormatted returns a boolean if a field has been set.

### GetNumber

`func (o *EmployerManagersPhone) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *EmployerManagersPhone) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *EmployerManagersPhone) SetNumber(v string)`

SetNumber sets Number field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


