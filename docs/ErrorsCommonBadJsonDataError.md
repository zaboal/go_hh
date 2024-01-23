# ErrorsCommonBadJsonDataError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | Описание ошибки | [optional] 
**Pointer** | Pointer to **string** | Путь до параметра, в котором возникла ошибка.  Для указания параметра используется формат JsonPointer [RFC 6901](https://tools.ietf.org/html/rfc6901)  | [optional] 
**Reason** | Pointer to **string** | Причина ошибки. Возможные значения:   * &#x60;required&#x60; - отсутствует обязательное поле   * &#x60;invalid&#x60; - значение введено не корректно   * &#x60;size_less_than_min&#x60; - пустой массив   * &#x60;unexpected&#x60; - поле не ожидается  | [optional] 
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | Pointer to **string** | Название поля с ошибкой | [optional] 

## Methods

### NewErrorsCommonBadJsonDataError

`func NewErrorsCommonBadJsonDataError(type_ string, ) *ErrorsCommonBadJsonDataError`

NewErrorsCommonBadJsonDataError instantiates a new ErrorsCommonBadJsonDataError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadJsonDataErrorWithDefaults

`func NewErrorsCommonBadJsonDataErrorWithDefaults() *ErrorsCommonBadJsonDataError`

NewErrorsCommonBadJsonDataErrorWithDefaults instantiates a new ErrorsCommonBadJsonDataError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ErrorsCommonBadJsonDataError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsCommonBadJsonDataError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsCommonBadJsonDataError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsCommonBadJsonDataError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ErrorsCommonBadJsonDataError) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ErrorsCommonBadJsonDataError) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPointer

`func (o *ErrorsCommonBadJsonDataError) GetPointer() string`

GetPointer returns the Pointer field if non-nil, zero value otherwise.

### GetPointerOk

`func (o *ErrorsCommonBadJsonDataError) GetPointerOk() (*string, bool)`

GetPointerOk returns a tuple with the Pointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointer

`func (o *ErrorsCommonBadJsonDataError) SetPointer(v string)`

SetPointer sets Pointer field to given value.

### HasPointer

`func (o *ErrorsCommonBadJsonDataError) HasPointer() bool`

HasPointer returns a boolean if a field has been set.

### GetReason

`func (o *ErrorsCommonBadJsonDataError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorsCommonBadJsonDataError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorsCommonBadJsonDataError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorsCommonBadJsonDataError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetType

`func (o *ErrorsCommonBadJsonDataError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsCommonBadJsonDataError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsCommonBadJsonDataError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsCommonBadJsonDataError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsCommonBadJsonDataError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsCommonBadJsonDataError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorsCommonBadJsonDataError) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


