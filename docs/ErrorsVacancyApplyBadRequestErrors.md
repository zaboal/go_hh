# ErrorsVacancyApplyBadRequestErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**BadArgument** | Pointer to **string** |  | [optional] 
**BadArguments** | Pointer to [**[]ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments**](ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments.md) |  | [optional] 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsVacancyApplyBadRequestError**](ErrorsVacancyApplyBadRequestError.md) | Массив с данными ошибок | 

## Methods

### NewErrorsVacancyApplyBadRequestErrors

`func NewErrorsVacancyApplyBadRequestErrors(requestId string, errors []ErrorsVacancyApplyBadRequestError, ) *ErrorsVacancyApplyBadRequestErrors`

NewErrorsVacancyApplyBadRequestErrors instantiates a new ErrorsVacancyApplyBadRequestErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyApplyBadRequestErrorsWithDefaults

`func NewErrorsVacancyApplyBadRequestErrorsWithDefaults() *ErrorsVacancyApplyBadRequestErrors`

NewErrorsVacancyApplyBadRequestErrorsWithDefaults instantiates a new ErrorsVacancyApplyBadRequestErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsVacancyApplyBadRequestErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsVacancyApplyBadRequestErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsVacancyApplyBadRequestErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetBadArgument

`func (o *ErrorsVacancyApplyBadRequestErrors) GetBadArgument() string`

GetBadArgument returns the BadArgument field if non-nil, zero value otherwise.

### GetBadArgumentOk

`func (o *ErrorsVacancyApplyBadRequestErrors) GetBadArgumentOk() (*string, bool)`

GetBadArgumentOk returns a tuple with the BadArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArgument

`func (o *ErrorsVacancyApplyBadRequestErrors) SetBadArgument(v string)`

SetBadArgument sets BadArgument field to given value.

### HasBadArgument

`func (o *ErrorsVacancyApplyBadRequestErrors) HasBadArgument() bool`

HasBadArgument returns a boolean if a field has been set.

### GetBadArguments

`func (o *ErrorsVacancyApplyBadRequestErrors) GetBadArguments() []ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments`

GetBadArguments returns the BadArguments field if non-nil, zero value otherwise.

### GetBadArgumentsOk

`func (o *ErrorsVacancyApplyBadRequestErrors) GetBadArgumentsOk() (*[]ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments, bool)`

GetBadArgumentsOk returns a tuple with the BadArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArguments

`func (o *ErrorsVacancyApplyBadRequestErrors) SetBadArguments(v []ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments)`

SetBadArguments sets BadArguments field to given value.

### HasBadArguments

`func (o *ErrorsVacancyApplyBadRequestErrors) HasBadArguments() bool`

HasBadArguments returns a boolean if a field has been set.

### GetDescription

`func (o *ErrorsVacancyApplyBadRequestErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsVacancyApplyBadRequestErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsVacancyApplyBadRequestErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsVacancyApplyBadRequestErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsVacancyApplyBadRequestErrors) GetErrors() []ErrorsVacancyApplyBadRequestError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsVacancyApplyBadRequestErrors) GetErrorsOk() (*[]ErrorsVacancyApplyBadRequestError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsVacancyApplyBadRequestErrors) SetErrors(v []ErrorsVacancyApplyBadRequestError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


