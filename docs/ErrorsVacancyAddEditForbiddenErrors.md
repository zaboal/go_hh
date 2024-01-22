# ErrorsVacancyAddEditForbiddenErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsVacancyAddEditForbiddenError**](ErrorsVacancyAddEditForbiddenError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsVacancyAddEditForbiddenErrors

`func NewErrorsVacancyAddEditForbiddenErrors(requestId string, errors []ErrorsVacancyAddEditForbiddenError, ) *ErrorsVacancyAddEditForbiddenErrors`

NewErrorsVacancyAddEditForbiddenErrors instantiates a new ErrorsVacancyAddEditForbiddenErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyAddEditForbiddenErrorsWithDefaults

`func NewErrorsVacancyAddEditForbiddenErrorsWithDefaults() *ErrorsVacancyAddEditForbiddenErrors`

NewErrorsVacancyAddEditForbiddenErrorsWithDefaults instantiates a new ErrorsVacancyAddEditForbiddenErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsVacancyAddEditForbiddenErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsVacancyAddEditForbiddenErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsVacancyAddEditForbiddenErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsVacancyAddEditForbiddenErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsVacancyAddEditForbiddenErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsVacancyAddEditForbiddenErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsVacancyAddEditForbiddenErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsVacancyAddEditForbiddenErrors) GetErrors() []ErrorsVacancyAddEditForbiddenError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsVacancyAddEditForbiddenErrors) GetErrorsOk() (*[]ErrorsVacancyAddEditForbiddenError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsVacancyAddEditForbiddenErrors) SetErrors(v []ErrorsVacancyAddEditForbiddenError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


