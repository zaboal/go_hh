# ErrorsEmployerManagerBadAuthorizationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAccounts** | Pointer to [**[]ManagerAccount**](ManagerAccount.md) | Список доступных для токена аккаунтов менеджера в случае, если используемый рабочий аккаунт заблокирован. Актуально только в случае авторизации работодателя  | [optional] 
**Reason** | Pointer to **string** | Ошибки при создании или редактировании менеджера работодателя:   * &#x60;already_exist&#x60; — Менеджер с такой почтой уже существует   * &#x60;creation_limit_exceeded&#x60; — Достигнут лимит на создание менеджеров   * &#x60;not_editable&#x60; — Поле *field_name* недоступно для редактирования  | [optional] 
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | Pointer to **string** | Общие ошибки:   * &#x60;bad_authorization&#x60; — Токен авторизации не существует или не валидный   * &#x60;token_expired&#x60; — Время жизни access_token завершилось, необходимо [выполнить обновление access_token](#section/Avtorizaciya/Obnovlenie-pary-access-i-refresh-tokenov)   * &#x60;token_revoked&#x60; — Токен отозван пользователем или сервером, приложению необходимо [запросить новую авторизацию](#section/Tipy-avtorizacij)   * &#x60;application_not_found&#x60; — Ваше приложение было удалено   * &#x60;used_manager_account_forbidden&#x60; — [Рабочий аккаунт заблокирован](https://github.com/hhru/api/blob/master/docs/errors.md#manager-accounts-blocked)   * &#x60;manager_extra_account_not_found&#x60; — В заголовке передан некорректный id аккаунта   * &#x60;email&#x60; — Менеджер с такой почтой уже существует   * &#x60;user_auth_expected&#x60; — Ожидается авторизация пользователя, передана авторизация приложения  | [optional] 

## Methods

### NewErrorsEmployerManagerBadAuthorizationError

`func NewErrorsEmployerManagerBadAuthorizationError(type_ string, ) *ErrorsEmployerManagerBadAuthorizationError`

NewErrorsEmployerManagerBadAuthorizationError instantiates a new ErrorsEmployerManagerBadAuthorizationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsEmployerManagerBadAuthorizationErrorWithDefaults

`func NewErrorsEmployerManagerBadAuthorizationErrorWithDefaults() *ErrorsEmployerManagerBadAuthorizationError`

NewErrorsEmployerManagerBadAuthorizationErrorWithDefaults instantiates a new ErrorsEmployerManagerBadAuthorizationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAccounts

`func (o *ErrorsEmployerManagerBadAuthorizationError) GetAllowedAccounts() []ManagerAccount`

GetAllowedAccounts returns the AllowedAccounts field if non-nil, zero value otherwise.

### GetAllowedAccountsOk

`func (o *ErrorsEmployerManagerBadAuthorizationError) GetAllowedAccountsOk() (*[]ManagerAccount, bool)`

GetAllowedAccountsOk returns a tuple with the AllowedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAccounts

`func (o *ErrorsEmployerManagerBadAuthorizationError) SetAllowedAccounts(v []ManagerAccount)`

SetAllowedAccounts sets AllowedAccounts field to given value.

### HasAllowedAccounts

`func (o *ErrorsEmployerManagerBadAuthorizationError) HasAllowedAccounts() bool`

HasAllowedAccounts returns a boolean if a field has been set.

### GetReason

`func (o *ErrorsEmployerManagerBadAuthorizationError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ErrorsEmployerManagerBadAuthorizationError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ErrorsEmployerManagerBadAuthorizationError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ErrorsEmployerManagerBadAuthorizationError) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetType

`func (o *ErrorsEmployerManagerBadAuthorizationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsEmployerManagerBadAuthorizationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsEmployerManagerBadAuthorizationError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsEmployerManagerBadAuthorizationError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsEmployerManagerBadAuthorizationError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsEmployerManagerBadAuthorizationError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorsEmployerManagerBadAuthorizationError) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


