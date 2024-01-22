# ErrorsVacancyProlongateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Тип ошибки | 
**Value** | **string** | Причина ошибки:  * &#x60;not_enough_purchased_services&#x60; — купленных услуг недостаточно для продления данного типа вакансии. * &#x60;quota_exceeded&#x60; — квота менеджера на публикацию данного типа вакансии закончилась. * &#x60;prolongation_forbidden&#x60; — продление вакансий недоступно текущему менеджеру. * &#x60;unavailable_for_archived&#x60; — продление недоступно для архивной вакансии. * &#x60;too_early&#x60; — продление раньше времени  | 

## Methods

### NewErrorsVacancyProlongateError

`func NewErrorsVacancyProlongateError(type_ string, value string, ) *ErrorsVacancyProlongateError`

NewErrorsVacancyProlongateError instantiates a new ErrorsVacancyProlongateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyProlongateErrorWithDefaults

`func NewErrorsVacancyProlongateErrorWithDefaults() *ErrorsVacancyProlongateError`

NewErrorsVacancyProlongateErrorWithDefaults instantiates a new ErrorsVacancyProlongateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsVacancyProlongateError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsVacancyProlongateError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsVacancyProlongateError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsVacancyProlongateError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsVacancyProlongateError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsVacancyProlongateError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


