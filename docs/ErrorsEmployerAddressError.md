# ErrorsEmployerAddressError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** | Причина ошибки. Описание: * &#x60;too_early&#x60; - Отступ по времени от текущей даты слишком большой. * &#x60;must_be_a_valid_ISO_8601_date-time_string&#x60; - Дата-время указаны в неверном формате  | [optional] 
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | Pointer to **string** | Название поля с ошибкой | [optional] 

## Methods

### NewErrorsEmployerAddressError

`func NewErrorsEmployerAddressError(type_ string, ) *ErrorsEmployerAddressError`

NewErrorsEmployerAddressError instantiates a new ErrorsEmployerAddressError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsEmployerAddressErrorWithDefaults

`func NewErrorsEmployerAddressErrorWithDefaults() *ErrorsEmployerAddressError`

NewErrorsEmployerAddressErrorWithDefaults instantiates a new ErrorsEmployerAddressError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *ErrorsEmployerAddressError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorsEmployerAddressError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorsEmployerAddressError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorsEmployerAddressError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetType

`func (o *ErrorsEmployerAddressError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsEmployerAddressError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsEmployerAddressError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsEmployerAddressError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsEmployerAddressError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsEmployerAddressError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorsEmployerAddressError) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


