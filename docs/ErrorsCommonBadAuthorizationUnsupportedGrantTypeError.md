# ErrorsCommonBadAuthorizationUnsupportedGrantTypeError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Идентификатор типа ошибки, используются значения, описанные в [документе RFC 6749](http://tools.ietf.org/html/rfc6749#section-5.2)  | 
**ErrorDescription** | **string** | Дополнительное описание ошибки * &#x60;unsupported grant_type&#x60; Возникает, если передать неправильное значение в поле &#x60;grant_type&#x60;  | 

## Methods

### NewErrorsCommonBadAuthorizationUnsupportedGrantTypeError

`func NewErrorsCommonBadAuthorizationUnsupportedGrantTypeError(error_ string, errorDescription string, ) *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError`

NewErrorsCommonBadAuthorizationUnsupportedGrantTypeError instantiates a new ErrorsCommonBadAuthorizationUnsupportedGrantTypeError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadAuthorizationUnsupportedGrantTypeErrorWithDefaults

`func NewErrorsCommonBadAuthorizationUnsupportedGrantTypeErrorWithDefaults() *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError`

NewErrorsCommonBadAuthorizationUnsupportedGrantTypeErrorWithDefaults instantiates a new ErrorsCommonBadAuthorizationUnsupportedGrantTypeError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ErrorsCommonBadAuthorizationUnsupportedGrantTypeError) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


