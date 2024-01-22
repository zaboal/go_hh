# ErrorsVacancyApplyBadRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Название поля с ошибкой:  * &#x60;vacancy_not_found&#x60; — вакансия не найдена. * &#x60;resume_not_found&#x60; — резюме из отклика/приглашения скрыто, удалено или не найдено. * &#x60;limit_exceeded&#x60; — превышен лимит количества откликов/приглашений * &#x60;resource_policy_violation&#x60; — отклик нарушает правила пользования сервисом * &#x60;inappropriate_language&#x60; — отклик содержит нецензурную лексику  | 

## Methods

### NewErrorsVacancyApplyBadRequestError

`func NewErrorsVacancyApplyBadRequestError(type_ string, value string, ) *ErrorsVacancyApplyBadRequestError`

NewErrorsVacancyApplyBadRequestError instantiates a new ErrorsVacancyApplyBadRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyApplyBadRequestErrorWithDefaults

`func NewErrorsVacancyApplyBadRequestErrorWithDefaults() *ErrorsVacancyApplyBadRequestError`

NewErrorsVacancyApplyBadRequestErrorWithDefaults instantiates a new ErrorsVacancyApplyBadRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsVacancyApplyBadRequestError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsVacancyApplyBadRequestError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsVacancyApplyBadRequestError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsVacancyApplyBadRequestError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsVacancyApplyBadRequestError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsVacancyApplyBadRequestError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


