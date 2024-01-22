# ErrorsSavedSearchForbiddenError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAccounts** | Pointer to [**[]ManagerAccount**](ManagerAccount.md) | Список доступных для токена аккаунтов менеджера в случае, если используемый рабочий аккаунт заблокирован. Актуально только в случае авторизации работодателя  | [optional] 
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | Pointer to **string** | Возможные ошибки: * &#x60;cant_send_to_yourself&#x60; — Нельзя передать сохраненный поиск самому себе * &#x60;user_auth_expected&#x60; — Передана авторизация приложения, метод требует [авторизации пользователя](#section/Avtorizaciya/Avtorizaciya-polzovatelya)  | [optional] 

## Methods

### NewErrorsSavedSearchForbiddenError

`func NewErrorsSavedSearchForbiddenError(type_ string, ) *ErrorsSavedSearchForbiddenError`

NewErrorsSavedSearchForbiddenError instantiates a new ErrorsSavedSearchForbiddenError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsSavedSearchForbiddenErrorWithDefaults

`func NewErrorsSavedSearchForbiddenErrorWithDefaults() *ErrorsSavedSearchForbiddenError`

NewErrorsSavedSearchForbiddenErrorWithDefaults instantiates a new ErrorsSavedSearchForbiddenError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAccounts

`func (o *ErrorsSavedSearchForbiddenError) GetAllowedAccounts() []ManagerAccount`

GetAllowedAccounts returns the AllowedAccounts field if non-nil, zero value otherwise.

### GetAllowedAccountsOk

`func (o *ErrorsSavedSearchForbiddenError) GetAllowedAccountsOk() (*[]ManagerAccount, bool)`

GetAllowedAccountsOk returns a tuple with the AllowedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAccounts

`func (o *ErrorsSavedSearchForbiddenError) SetAllowedAccounts(v []ManagerAccount)`

SetAllowedAccounts sets AllowedAccounts field to given value.

### HasAllowedAccounts

`func (o *ErrorsSavedSearchForbiddenError) HasAllowedAccounts() bool`

HasAllowedAccounts returns a boolean if a field has been set.

### GetType

`func (o *ErrorsSavedSearchForbiddenError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsSavedSearchForbiddenError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsSavedSearchForbiddenError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsSavedSearchForbiddenError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsSavedSearchForbiddenError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsSavedSearchForbiddenError) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ErrorsSavedSearchForbiddenError) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


