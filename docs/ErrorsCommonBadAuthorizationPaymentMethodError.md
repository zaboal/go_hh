# ErrorsCommonBadAuthorizationPaymentMethodError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Описание ошибки.  Причина ошибки в том, что вы запрашиваете один из [платных методов](#tag/Uslugi-rabotodatelya/operation/get-payable-api-method-access) без купленного доступа  | 

## Methods

### NewErrorsCommonBadAuthorizationPaymentMethodError

`func NewErrorsCommonBadAuthorizationPaymentMethodError(type_ string, value string, ) *ErrorsCommonBadAuthorizationPaymentMethodError`

NewErrorsCommonBadAuthorizationPaymentMethodError instantiates a new ErrorsCommonBadAuthorizationPaymentMethodError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadAuthorizationPaymentMethodErrorWithDefaults

`func NewErrorsCommonBadAuthorizationPaymentMethodErrorWithDefaults() *ErrorsCommonBadAuthorizationPaymentMethodError`

NewErrorsCommonBadAuthorizationPaymentMethodErrorWithDefaults instantiates a new ErrorsCommonBadAuthorizationPaymentMethodError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsCommonBadAuthorizationPaymentMethodError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsCommonBadAuthorizationPaymentMethodError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsCommonBadAuthorizationPaymentMethodError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsCommonBadAuthorizationPaymentMethodError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsCommonBadAuthorizationPaymentMethodError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsCommonBadAuthorizationPaymentMethodError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


