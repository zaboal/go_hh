# ErrorsResumeVisibilityError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Название поля с ошибкой:  - &#x60;per_page&#x60; — передано невалидное количество элементов на странице (максимум 100). - &#x60;unknown_employer&#x60; — передан неизвестный идентификатор работодателя. - &#x60;limit_exceeded&#x60; — превышен лимит списка видимости. - &#x60;too_many_employers&#x60; — передано слишком много работодателей. - &#x60;id&#x60; — передан невалидный идентификатор работодателя  | 

## Methods

### NewErrorsResumeVisibilityError

`func NewErrorsResumeVisibilityError(type_ string, value string, ) *ErrorsResumeVisibilityError`

NewErrorsResumeVisibilityError instantiates a new ErrorsResumeVisibilityError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsResumeVisibilityErrorWithDefaults

`func NewErrorsResumeVisibilityErrorWithDefaults() *ErrorsResumeVisibilityError`

NewErrorsResumeVisibilityErrorWithDefaults instantiates a new ErrorsResumeVisibilityError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsResumeVisibilityError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsResumeVisibilityError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsResumeVisibilityError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsResumeVisibilityError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsResumeVisibilityError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsResumeVisibilityError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


