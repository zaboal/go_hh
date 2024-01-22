# ErrorsCommonBadAuthorizationInvalidRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Идентификатор типа ошибки, используются значения, описанные в [документе RFC 6749](http://tools.ietf.org/html/rfc6749#section-5.2)  | 
**ErrorDescription** | **string** | Дополнительное описание ошибки * &#x60;account not found&#x60; Ошибка может возникнуть, если передана неправильная пара &#x60;client_id&#x60; и &#x60;client_secret&#x60; * &#x60;account is locked&#x60; Пользовательский аккаунт заблокирован. Пользователь должен обратиться в службу поддержки сайта * &#x60;password invalidated&#x60; Пароль от пользовательского аккаунта устарел. Пользователь должен восстановить пароль на сайте &#x60;https://hh.ru&#x60; * &#x60;login not verified&#x60; Пользовательский аккаунт не подтвержден. Пользователь должен обратиться в службу поддержки сайта * &#x60;bad redirect url&#x60; передан неправильный &#x60;redirect_url&#x60; * &#x60;token is empty&#x60; Не передан &#x60;refresh_token&#x60; * &#x60;token not found&#x60; передан не правильный &#x60;refresh_token&#x60; * &#x60;code not found &#x60; переданный &#x60;authorization_code&#x60; не найден  | 

## Methods

### NewErrorsCommonBadAuthorizationInvalidRequestError

`func NewErrorsCommonBadAuthorizationInvalidRequestError(error_ string, errorDescription string, ) *ErrorsCommonBadAuthorizationInvalidRequestError`

NewErrorsCommonBadAuthorizationInvalidRequestError instantiates a new ErrorsCommonBadAuthorizationInvalidRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadAuthorizationInvalidRequestErrorWithDefaults

`func NewErrorsCommonBadAuthorizationInvalidRequestErrorWithDefaults() *ErrorsCommonBadAuthorizationInvalidRequestError`

NewErrorsCommonBadAuthorizationInvalidRequestErrorWithDefaults instantiates a new ErrorsCommonBadAuthorizationInvalidRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorsCommonBadAuthorizationInvalidRequestError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorsCommonBadAuthorizationInvalidRequestError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorsCommonBadAuthorizationInvalidRequestError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *ErrorsCommonBadAuthorizationInvalidRequestError) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ErrorsCommonBadAuthorizationInvalidRequestError) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ErrorsCommonBadAuthorizationInvalidRequestError) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


