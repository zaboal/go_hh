# ErrorsVacancyFavoritedBadAuthErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsVacancyFavoritedBadAuthError**](ErrorsVacancyFavoritedBadAuthError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsVacancyFavoritedBadAuthErrors

`func NewErrorsVacancyFavoritedBadAuthErrors(requestId string, errors []ErrorsVacancyFavoritedBadAuthError, ) *ErrorsVacancyFavoritedBadAuthErrors`

NewErrorsVacancyFavoritedBadAuthErrors instantiates a new ErrorsVacancyFavoritedBadAuthErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyFavoritedBadAuthErrorsWithDefaults

`func NewErrorsVacancyFavoritedBadAuthErrorsWithDefaults() *ErrorsVacancyFavoritedBadAuthErrors`

NewErrorsVacancyFavoritedBadAuthErrorsWithDefaults instantiates a new ErrorsVacancyFavoritedBadAuthErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsVacancyFavoritedBadAuthErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsVacancyFavoritedBadAuthErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsVacancyFavoritedBadAuthErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsVacancyFavoritedBadAuthErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsVacancyFavoritedBadAuthErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsVacancyFavoritedBadAuthErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsVacancyFavoritedBadAuthErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsVacancyFavoritedBadAuthErrors) GetErrors() []ErrorsVacancyFavoritedBadAuthError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsVacancyFavoritedBadAuthErrors) GetErrorsOk() (*[]ErrorsVacancyFavoritedBadAuthError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsVacancyFavoritedBadAuthErrors) SetErrors(v []ErrorsVacancyFavoritedBadAuthError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


