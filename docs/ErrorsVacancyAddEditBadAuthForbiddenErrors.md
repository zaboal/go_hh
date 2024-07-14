# ErrorsVacancyAddEditBadAuthForbiddenErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsVacancyAddEditForbiddenError**](ErrorsVacancyAddEditForbiddenError.md) | Массив с данными ошибок | 
**OauthError** | Pointer to **string** | Ошибки авторизации:   * &#x60;token-revoked&#x60; — Токен отозван пользователем или сервером, приложению необходимо [запросить новую авторизацию](#tag/Avtorizaciya-rabotodatelya/operation/authorize)   * &#x60;token-expired&#x60; — Время жизни &#x60;access_token&#x60; завершилось, необходимо [получить &#x60;refresh_token&#x60;](#tag/Avtorizaciya-rabotodatelya/operation/authorize)  | [optional] 

## Methods

### NewErrorsVacancyAddEditBadAuthForbiddenErrors

`func NewErrorsVacancyAddEditBadAuthForbiddenErrors(requestId string, errors []ErrorsVacancyAddEditForbiddenError, ) *ErrorsVacancyAddEditBadAuthForbiddenErrors`

NewErrorsVacancyAddEditBadAuthForbiddenErrors instantiates a new ErrorsVacancyAddEditBadAuthForbiddenErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyAddEditBadAuthForbiddenErrorsWithDefaults

`func NewErrorsVacancyAddEditBadAuthForbiddenErrorsWithDefaults() *ErrorsVacancyAddEditBadAuthForbiddenErrors`

NewErrorsVacancyAddEditBadAuthForbiddenErrorsWithDefaults instantiates a new ErrorsVacancyAddEditBadAuthForbiddenErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) GetErrors() []ErrorsVacancyAddEditForbiddenError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) GetErrorsOk() (*[]ErrorsVacancyAddEditForbiddenError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) SetErrors(v []ErrorsVacancyAddEditForbiddenError)`

SetErrors sets Errors field to given value.


### GetOauthError

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) GetOauthError() string`

GetOauthError returns the OauthError field if non-nil, zero value otherwise.

### GetOauthErrorOk

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) GetOauthErrorOk() (*string, bool)`

GetOauthErrorOk returns a tuple with the OauthError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthError

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) SetOauthError(v string)`

SetOauthError sets OauthError field to given value.

### HasOauthError

`func (o *ErrorsVacancyAddEditBadAuthForbiddenErrors) HasOauthError() bool`

HasOauthError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


