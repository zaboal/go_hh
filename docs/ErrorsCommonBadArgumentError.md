# ErrorsCommonBadArgumentError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Причина ошибки. Возможные значения:   * &#x60;too_long_value&#x60; - слишком длинное значение   * &#x60;too_many_arguments&#x60; - слишком много аргументов   * &#x60;invalid&#x60; - значение введено не корректно  | [optional] 
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | Pointer to **string** | Название поля с ошибкой  | [optional] 

## Methods

### NewErrorsCommonBadArgumentError

`func NewErrorsCommonBadArgumentError(type_ string, ) *ErrorsCommonBadArgumentError`

NewErrorsCommonBadArgumentError instantiates a new ErrorsCommonBadArgumentError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadArgumentErrorWithDefaults

`func NewErrorsCommonBadArgumentErrorWithDefaults() *ErrorsCommonBadArgumentError`

NewErrorsCommonBadArgumentErrorWithDefaults instantiates a new ErrorsCommonBadArgumentError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ErrorsCommonBadArgumentError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorsCommonBadArgumentError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorsCommonBadArgumentError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorsCommonBadArgumentError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetType

`func (o *ErrorsCommonBadArgumentError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsCommonBadArgumentError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsCommonBadArgumentError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsCommonBadArgumentError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsCommonBadArgumentError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsCommonBadArgumentError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorsCommonBadArgumentError) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


