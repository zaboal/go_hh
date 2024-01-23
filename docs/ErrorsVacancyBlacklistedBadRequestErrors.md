# ErrorsVacancyBlacklistedBadRequestErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsVacancyBlacklistedBadRequestError**](ErrorsVacancyBlacklistedBadRequestError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsVacancyBlacklistedBadRequestErrors

`func NewErrorsVacancyBlacklistedBadRequestErrors(requestId string, errors []ErrorsVacancyBlacklistedBadRequestError, ) *ErrorsVacancyBlacklistedBadRequestErrors`

NewErrorsVacancyBlacklistedBadRequestErrors instantiates a new ErrorsVacancyBlacklistedBadRequestErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyBlacklistedBadRequestErrorsWithDefaults

`func NewErrorsVacancyBlacklistedBadRequestErrorsWithDefaults() *ErrorsVacancyBlacklistedBadRequestErrors`

NewErrorsVacancyBlacklistedBadRequestErrorsWithDefaults instantiates a new ErrorsVacancyBlacklistedBadRequestErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) GetErrors() []ErrorsVacancyBlacklistedBadRequestError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) GetErrorsOk() (*[]ErrorsVacancyBlacklistedBadRequestError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsVacancyBlacklistedBadRequestErrors) SetErrors(v []ErrorsVacancyBlacklistedBadRequestError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


