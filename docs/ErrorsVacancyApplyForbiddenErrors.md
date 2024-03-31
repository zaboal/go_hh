# ErrorsVacancyApplyForbiddenErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**BadArgument** | Pointer to **string** |  | [optional] 
**BadArguments** | Pointer to [**[]ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments**](ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments.md) |  | [optional] 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsVacancyApplyForbiddenError**](ErrorsVacancyApplyForbiddenError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsVacancyApplyForbiddenErrors

`func NewErrorsVacancyApplyForbiddenErrors(requestId string, errors []ErrorsVacancyApplyForbiddenError, ) *ErrorsVacancyApplyForbiddenErrors`

NewErrorsVacancyApplyForbiddenErrors instantiates a new ErrorsVacancyApplyForbiddenErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyApplyForbiddenErrorsWithDefaults

`func NewErrorsVacancyApplyForbiddenErrorsWithDefaults() *ErrorsVacancyApplyForbiddenErrors`

NewErrorsVacancyApplyForbiddenErrorsWithDefaults instantiates a new ErrorsVacancyApplyForbiddenErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsVacancyApplyForbiddenErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsVacancyApplyForbiddenErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsVacancyApplyForbiddenErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetBadArgument

`func (o *ErrorsVacancyApplyForbiddenErrors) GetBadArgument() string`

GetBadArgument returns the BadArgument field if non-nil, zero value otherwise.

### GetBadArgumentOk

`func (o *ErrorsVacancyApplyForbiddenErrors) GetBadArgumentOk() (*string, bool)`

GetBadArgumentOk returns a tuple with the BadArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArgument

`func (o *ErrorsVacancyApplyForbiddenErrors) SetBadArgument(v string)`

SetBadArgument sets BadArgument field to given value.

### HasBadArgument

`func (o *ErrorsVacancyApplyForbiddenErrors) HasBadArgument() bool`

HasBadArgument returns a boolean if a field has been set.

### GetBadArguments

`func (o *ErrorsVacancyApplyForbiddenErrors) GetBadArguments() []ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments`

GetBadArguments returns the BadArguments field if non-nil, zero value otherwise.

### GetBadArgumentsOk

`func (o *ErrorsVacancyApplyForbiddenErrors) GetBadArgumentsOk() (*[]ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments, bool)`

GetBadArgumentsOk returns a tuple with the BadArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArguments

`func (o *ErrorsVacancyApplyForbiddenErrors) SetBadArguments(v []ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments)`

SetBadArguments sets BadArguments field to given value.

### HasBadArguments

`func (o *ErrorsVacancyApplyForbiddenErrors) HasBadArguments() bool`

HasBadArguments returns a boolean if a field has been set.

### GetDescription

`func (o *ErrorsVacancyApplyForbiddenErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsVacancyApplyForbiddenErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsVacancyApplyForbiddenErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsVacancyApplyForbiddenErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsVacancyApplyForbiddenErrors) GetErrors() []ErrorsVacancyApplyForbiddenError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsVacancyApplyForbiddenErrors) GetErrorsOk() (*[]ErrorsVacancyApplyForbiddenError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsVacancyApplyForbiddenErrors) SetErrors(v []ErrorsVacancyApplyForbiddenError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


