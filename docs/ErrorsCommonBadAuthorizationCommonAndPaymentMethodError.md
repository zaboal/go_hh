# ErrorsCommonBadAuthorizationCommonAndPaymentMethodError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsCommonBadAuthorizationPaymentMethodError**](ErrorsCommonBadAuthorizationPaymentMethodError.md) | Массив с данными ошибок | 
**OauthError** | Pointer to **string** | Ошибки авторизации:   * &#x60;token-revoked&#x60; — Токен отозван пользователем или сервером, приложению необходимо [запросить новую авторизацию](#tag/Avtorizaciya-rabotodatelya/operation/authorize)   * &#x60;token-expired&#x60; — Время жизни &#x60;access_token&#x60; завершилось, необходимо [получить &#x60;refresh_token&#x60;](#tag/Avtorizaciya-rabotodatelya/operation/authorize)  | [optional] 

## Methods

### NewErrorsCommonBadAuthorizationCommonAndPaymentMethodError

`func NewErrorsCommonBadAuthorizationCommonAndPaymentMethodError(requestId string, errors []ErrorsCommonBadAuthorizationPaymentMethodError, ) *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError`

NewErrorsCommonBadAuthorizationCommonAndPaymentMethodError instantiates a new ErrorsCommonBadAuthorizationCommonAndPaymentMethodError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadAuthorizationCommonAndPaymentMethodErrorWithDefaults

`func NewErrorsCommonBadAuthorizationCommonAndPaymentMethodErrorWithDefaults() *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError`

NewErrorsCommonBadAuthorizationCommonAndPaymentMethodErrorWithDefaults instantiates a new ErrorsCommonBadAuthorizationCommonAndPaymentMethodError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) GetErrors() []ErrorsCommonBadAuthorizationPaymentMethodError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) GetErrorsOk() (*[]ErrorsCommonBadAuthorizationPaymentMethodError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) SetErrors(v []ErrorsCommonBadAuthorizationPaymentMethodError)`

SetErrors sets Errors field to given value.


### GetOauthError

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) GetOauthError() string`

GetOauthError returns the OauthError field if non-nil, zero value otherwise.

### GetOauthErrorOk

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) GetOauthErrorOk() (*string, bool)`

GetOauthErrorOk returns a tuple with the OauthError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthError

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) SetOauthError(v string)`

SetOauthError sets OauthError field to given value.

### HasOauthError

`func (o *ErrorsCommonBadAuthorizationCommonAndPaymentMethodError) HasOauthError() bool`

HasOauthError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


