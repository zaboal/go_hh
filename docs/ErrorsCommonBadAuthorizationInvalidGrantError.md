# ErrorsCommonBadAuthorizationInvalidGrantError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | Идентификатор типа ошибки, используются значения, описанные в [документе RFC 6749](http://tools.ietf.org/html/rfc6749#section-5.2)  | 
**ErrorDescription** | **string** | Дополнительное описание ошибки * &#x60;token has already been refreshed&#x60; Ошибка возникает при попытке воспользоваться refresh-токеном второй раз * &#x60;token not expired&#x60; Ошибка возникает при попытке обновить действующий access-токен. access-токен можно обновлять только после того, как он истек * &#x60;token was revoked&#x60; Токен был отозван. Например, токен отзывается, если время действия пароля истекло * &#x60;bad token &#x60; Передано неправильное значение токена * &#x60;code has already been used&#x60; &#x60;authorization_code&#x60; уже был использован (его можно использовать только один раз) * &#x60;code expired&#x60; &#x60;authorization_code&#x60; истек * &#x60;code was revoke&#x60; &#x60;authorization_code&#x60; был отозван (это происходит, если время действия пароля истекло) * &#x60;token deactivated &#x60; Токен был деактивирован. Токен деактивируется после того, как пользователь сменил пароль  | 

## Methods

### NewErrorsCommonBadAuthorizationInvalidGrantError

`func NewErrorsCommonBadAuthorizationInvalidGrantError(error_ string, errorDescription string, ) *ErrorsCommonBadAuthorizationInvalidGrantError`

NewErrorsCommonBadAuthorizationInvalidGrantError instantiates a new ErrorsCommonBadAuthorizationInvalidGrantError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadAuthorizationInvalidGrantErrorWithDefaults

`func NewErrorsCommonBadAuthorizationInvalidGrantErrorWithDefaults() *ErrorsCommonBadAuthorizationInvalidGrantError`

NewErrorsCommonBadAuthorizationInvalidGrantErrorWithDefaults instantiates a new ErrorsCommonBadAuthorizationInvalidGrantError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorsCommonBadAuthorizationInvalidGrantError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorsCommonBadAuthorizationInvalidGrantError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorsCommonBadAuthorizationInvalidGrantError) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *ErrorsCommonBadAuthorizationInvalidGrantError) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ErrorsCommonBadAuthorizationInvalidGrantError) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ErrorsCommonBadAuthorizationInvalidGrantError) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


