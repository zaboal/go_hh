# ErrorsVacancyFavoritedCombinedBadAuthErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsVacancyFavoritedBadAuthError**](ErrorsVacancyFavoritedBadAuthError.md) | Массив с данными ошибок | 
**OauthError** | Pointer to **string** | Ошибки авторизации:   * &#x60;token-revoked&#x60; — Токен отозван пользователем или сервером, приложению необходимо [запросить новую авторизацию](#tag/Avtorizaciya-rabotodatelya/operation/authorize)   * &#x60;token-expired&#x60; — Время жизни &#x60;access_token&#x60; завершилось, необходимо [получить &#x60;refresh_token&#x60;](#tag/Avtorizaciya-rabotodatelya/operation/authorize)  | [optional] 

## Methods

### NewErrorsVacancyFavoritedCombinedBadAuthErrors

`func NewErrorsVacancyFavoritedCombinedBadAuthErrors(requestId string, errors []ErrorsVacancyFavoritedBadAuthError, ) *ErrorsVacancyFavoritedCombinedBadAuthErrors`

NewErrorsVacancyFavoritedCombinedBadAuthErrors instantiates a new ErrorsVacancyFavoritedCombinedBadAuthErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsVacancyFavoritedCombinedBadAuthErrorsWithDefaults

`func NewErrorsVacancyFavoritedCombinedBadAuthErrorsWithDefaults() *ErrorsVacancyFavoritedCombinedBadAuthErrors`

NewErrorsVacancyFavoritedCombinedBadAuthErrorsWithDefaults instantiates a new ErrorsVacancyFavoritedCombinedBadAuthErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetDescription

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) GetErrors() []ErrorsVacancyFavoritedBadAuthError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) GetErrorsOk() (*[]ErrorsVacancyFavoritedBadAuthError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) SetErrors(v []ErrorsVacancyFavoritedBadAuthError)`

SetErrors sets Errors field to given value.


### GetOauthError

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) GetOauthError() string`

GetOauthError returns the OauthError field if non-nil, zero value otherwise.

### GetOauthErrorOk

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) GetOauthErrorOk() (*string, bool)`

GetOauthErrorOk returns a tuple with the OauthError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthError

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) SetOauthError(v string)`

SetOauthError sets OauthError field to given value.

### HasOauthError

`func (o *ErrorsVacancyFavoritedCombinedBadAuthErrors) HasOauthError() bool`

HasOauthError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


