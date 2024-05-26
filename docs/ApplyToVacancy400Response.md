# ApplyToVacancy400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**BadArgument** | Pointer to **string** |  | [optional] 
**BadArguments** | Pointer to [**[]ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments**](ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments.md) |  | [optional] 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsCommonBadArgumentError**](ErrorsCommonBadArgumentError.md) | Массив с данными ошибок | 

## Methods

### NewApplyToVacancy400Response

`func NewApplyToVacancy400Response(requestId string, errors []ErrorsCommonBadArgumentError, ) *ApplyToVacancy400Response`

NewApplyToVacancy400Response instantiates a new ApplyToVacancy400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyToVacancy400ResponseWithDefaults

`func NewApplyToVacancy400ResponseWithDefaults() *ApplyToVacancy400Response`

NewApplyToVacancy400ResponseWithDefaults instantiates a new ApplyToVacancy400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ApplyToVacancy400Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ApplyToVacancy400Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ApplyToVacancy400Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetBadArgument

`func (o *ApplyToVacancy400Response) GetBadArgument() string`

GetBadArgument returns the BadArgument field if non-nil, zero value otherwise.

### GetBadArgumentOk

`func (o *ApplyToVacancy400Response) GetBadArgumentOk() (*string, bool)`

GetBadArgumentOk returns a tuple with the BadArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArgument

`func (o *ApplyToVacancy400Response) SetBadArgument(v string)`

SetBadArgument sets BadArgument field to given value.

### HasBadArgument

`func (o *ApplyToVacancy400Response) HasBadArgument() bool`

HasBadArgument returns a boolean if a field has been set.

### GetBadArguments

`func (o *ApplyToVacancy400Response) GetBadArguments() []ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments`

GetBadArguments returns the BadArguments field if non-nil, zero value otherwise.

### GetBadArgumentsOk

`func (o *ApplyToVacancy400Response) GetBadArgumentsOk() (*[]ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments, bool)`

GetBadArgumentsOk returns a tuple with the BadArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArguments

`func (o *ApplyToVacancy400Response) SetBadArguments(v []ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments)`

SetBadArguments sets BadArguments field to given value.

### HasBadArguments

`func (o *ApplyToVacancy400Response) HasBadArguments() bool`

HasBadArguments returns a boolean if a field has been set.

### GetDescription

`func (o *ApplyToVacancy400Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplyToVacancy400Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplyToVacancy400Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplyToVacancy400Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ApplyToVacancy400Response) GetErrors() []ErrorsCommonBadArgumentError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApplyToVacancy400Response) GetErrorsOk() (*[]ErrorsCommonBadArgumentError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApplyToVacancy400Response) SetErrors(v []ErrorsCommonBadArgumentError)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


