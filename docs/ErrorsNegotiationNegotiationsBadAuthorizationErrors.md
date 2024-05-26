# ErrorsNegotiationNegotiationsBadAuthorizationErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**BadArgument** | Pointer to **string** |  | [optional] 
**BadArguments** | Pointer to [**[]ErrorsVacancyApplyForbiddenErrorsAllOfBadArguments**](ErrorsVacancyApplyForbiddenErrorsAllOfBadArguments.md) |  | [optional] 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsCommonBadAuthorizationCommonAndEmployerError**](ErrorsCommonBadAuthorizationCommonAndEmployerError.md) | Массив с данными ошибок | 
**OauthError** | Pointer to **string** | Ошибки авторизации:   * &#x60;token-revoked&#x60; — Токен отозван пользователем, приложению необходимо [запросить новую авторизацию](#tag/Avtorizaciya-rabotodatelya/operation/authorize)   * &#x60;token-expired&#x60; — Время жизни &#x60;access_token&#x60; завершилось, необходимо [получить &#x60;refresh_token&#x60;](#tag/Avtorizaciya-rabotodatelya/operation/authorize)  | [optional] 

## Methods

### NewErrorsNegotiationNegotiationsBadAuthorizationErrors

`func NewErrorsNegotiationNegotiationsBadAuthorizationErrors(requestId string, errors []ErrorsCommonBadAuthorizationCommonAndEmployerError, ) *ErrorsNegotiationNegotiationsBadAuthorizationErrors`

NewErrorsNegotiationNegotiationsBadAuthorizationErrors instantiates a new ErrorsNegotiationNegotiationsBadAuthorizationErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsNegotiationNegotiationsBadAuthorizationErrorsWithDefaults

`func NewErrorsNegotiationNegotiationsBadAuthorizationErrorsWithDefaults() *ErrorsNegotiationNegotiationsBadAuthorizationErrors`

NewErrorsNegotiationNegotiationsBadAuthorizationErrorsWithDefaults instantiates a new ErrorsNegotiationNegotiationsBadAuthorizationErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetBadArgument

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetBadArgument() string`

GetBadArgument returns the BadArgument field if non-nil, zero value otherwise.

### GetBadArgumentOk

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetBadArgumentOk() (*string, bool)`

GetBadArgumentOk returns a tuple with the BadArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArgument

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) SetBadArgument(v string)`

SetBadArgument sets BadArgument field to given value.

### HasBadArgument

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) HasBadArgument() bool`

HasBadArgument returns a boolean if a field has been set.

### GetBadArguments

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetBadArguments() []ErrorsVacancyApplyForbiddenErrorsAllOfBadArguments`

GetBadArguments returns the BadArguments field if non-nil, zero value otherwise.

### GetBadArgumentsOk

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetBadArgumentsOk() (*[]ErrorsVacancyApplyForbiddenErrorsAllOfBadArguments, bool)`

GetBadArgumentsOk returns a tuple with the BadArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArguments

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) SetBadArguments(v []ErrorsVacancyApplyForbiddenErrorsAllOfBadArguments)`

SetBadArguments sets BadArguments field to given value.

### HasBadArguments

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) HasBadArguments() bool`

HasBadArguments returns a boolean if a field has been set.

### GetDescription

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetErrors() []ErrorsCommonBadAuthorizationCommonAndEmployerError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetErrorsOk() (*[]ErrorsCommonBadAuthorizationCommonAndEmployerError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) SetErrors(v []ErrorsCommonBadAuthorizationCommonAndEmployerError)`

SetErrors sets Errors field to given value.


### GetOauthError

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetOauthError() string`

GetOauthError returns the OauthError field if non-nil, zero value otherwise.

### GetOauthErrorOk

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) GetOauthErrorOk() (*string, bool)`

GetOauthErrorOk returns a tuple with the OauthError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthError

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) SetOauthError(v string)`

SetOauthError sets OauthError field to given value.

### HasOauthError

`func (o *ErrorsNegotiationNegotiationsBadAuthorizationErrors) HasOauthError() bool`

HasOauthError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


