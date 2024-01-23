# ErrorsCommonBadAuthorizationCommonAndEmployerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Необходимо пройти капчу - &#x60;captcha_required&#x60;  | 
**AllowedAccounts** | Pointer to [**[]ManagerAccount**](ManagerAccount.md) | Список доступных для токена аккаунтов менеджера в случае, если используемый рабочий аккаунт заблокирован. Актуально только в случае авторизации работодателя  | [optional] 
**CaptchaUrl** | Pointer to **string** | Адрес веб-страницы, на которой можно пройти капчу.  После прохождения капчи аналогичный запрос в API должен выполниться успешно.  Приложение должно добавить в captcha_url обязательный параметр backurl,на который произойдет редирект   после прохождения капчи.  Backurl должен обязательно содержать схему, например, https:// или схему приложения  | [optional] 
**FallbackUrl** | Pointer to **string** | Адрес веб-страницы, на котором можно капчу. Аналогично параметру captcha_url | [optional] 

## Methods

### NewErrorsCommonBadAuthorizationCommonAndEmployerError

`func NewErrorsCommonBadAuthorizationCommonAndEmployerError(type_ string, value string, ) *ErrorsCommonBadAuthorizationCommonAndEmployerError`

NewErrorsCommonBadAuthorizationCommonAndEmployerError instantiates a new ErrorsCommonBadAuthorizationCommonAndEmployerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorsCommonBadAuthorizationCommonAndEmployerErrorWithDefaults

`func NewErrorsCommonBadAuthorizationCommonAndEmployerErrorWithDefaults() *ErrorsCommonBadAuthorizationCommonAndEmployerError`

NewErrorsCommonBadAuthorizationCommonAndEmployerErrorWithDefaults instantiates a new ErrorsCommonBadAuthorizationCommonAndEmployerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) SetValue(v string)`

SetValue sets Value field to given value.


### GetAllowedAccounts

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetAllowedAccounts() []ManagerAccount`

GetAllowedAccounts returns the AllowedAccounts field if non-nil, zero value otherwise.

### GetAllowedAccountsOk

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetAllowedAccountsOk() (*[]ManagerAccount, bool)`

GetAllowedAccountsOk returns a tuple with the AllowedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAccounts

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) SetAllowedAccounts(v []ManagerAccount)`

SetAllowedAccounts sets AllowedAccounts field to given value.

### HasAllowedAccounts

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) HasAllowedAccounts() bool`

HasAllowedAccounts returns a boolean if a field has been set.

### GetCaptchaUrl

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetCaptchaUrl() string`

GetCaptchaUrl returns the CaptchaUrl field if non-nil, zero value otherwise.

### GetCaptchaUrlOk

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetCaptchaUrlOk() (*string, bool)`

GetCaptchaUrlOk returns a tuple with the CaptchaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaUrl

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) SetCaptchaUrl(v string)`

SetCaptchaUrl sets CaptchaUrl field to given value.

### HasCaptchaUrl

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) HasCaptchaUrl() bool`

HasCaptchaUrl returns a boolean if a field has been set.

### GetFallbackUrl

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *ErrorsCommonBadAuthorizationCommonAndEmployerError) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


