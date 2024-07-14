# ErrorsCommonBadAuthorizationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | Pointer to **string** | Общие ошибки:   * &#x60;bad_authorization&#x60; — Токен авторизации не существует или не валидный   * &#x60;token_expired&#x60; — Время жизни access_token завершилось, необходимо [выполнить обновление access_token](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov)   * &#x60;token_revoked&#x60; — Токен отозван пользователем или сервером, приложению необходимо [запросить новую авторизацию](#section/Tipy-avtorizacij)   * &#x60;application_not_found&#x60; — Ваше приложение было удалено   * &#x60;user_auth_expected&#x60; — Передана авторизация приложения, метод требует [авторизации пользователя](#get-auth)   * &#x60;application_auth_expected&#x60; — Передана авторизация пользователя, метод требует [авторизации приложения](#get-client-auth)   * &#x60;used_manager_account_forbidden&#x60; — [Рабочий аккаунт заблокирован](https://github.com/hhru/api/blob/master/docs/errors.md#manager-accounts-blocked)   * &#x60;manager_extra_account_not_found&#x60; — В заголовке передан некорректный id аккаунта   * &#x60;user_auth_expected&#x60; — Передана авторизация приложения, метод требует [авторизации пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)   * &#x60;application_auth_expected&#x60; — Передана авторизация пользователя, метод требует [авторизации приложения](##section/Avtorizaciya/Avtorizaciya-prilozheniya)  | [optional] 

## Methods

### NewErrorsCommonBadAuthorizationError

`func NewErrorsCommonBadAuthorizationError(type_ string, ) *ErrorsCommonBadAuthorizationError`

NewErrorsCommonBadAuthorizationError instantiates a new ErrorsCommonBadAuthorizationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadAuthorizationErrorWithDefaults

`func NewErrorsCommonBadAuthorizationErrorWithDefaults() *ErrorsCommonBadAuthorizationError`

NewErrorsCommonBadAuthorizationErrorWithDefaults instantiates a new ErrorsCommonBadAuthorizationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsCommonBadAuthorizationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsCommonBadAuthorizationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsCommonBadAuthorizationError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsCommonBadAuthorizationError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsCommonBadAuthorizationError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsCommonBadAuthorizationError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorsCommonBadAuthorizationError) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


