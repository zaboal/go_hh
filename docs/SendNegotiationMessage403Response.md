# SendNegotiationMessage403Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | Идентификатор запроса | 
**BadArgument** | Pointer to **string** |  | [optional] 
**BadArguments** | Pointer to [**[]ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments**](ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments.md) |  | [optional] 
**Description** | Pointer to **string** | Описание ошибки | [optional] 
**Errors** | [**[]ErrorsCommonBadAuthorizationCommonAndEmployerError**](ErrorsCommonBadAuthorizationCommonAndEmployerError.md) | Массив с данными ошибок | 
**OauthError** | Pointer to **string** | Ошибки авторизации:   * &#x60;token-revoked&#x60; — Токен отозван пользователем, приложению необходимо [запросить новую авторизацию](#tag/Avtorizaciya-rabotodatelya/operation/authorize)   * &#x60;token-expired&#x60; — Время жизни &#x60;access_token&#x60; завершилось, необходимо [получить &#x60;refresh_token&#x60;](#tag/Avtorizaciya-rabotodatelya/operation/authorize)  | [optional] 

## Methods

### NewSendNegotiationMessage403Response

`func NewSendNegotiationMessage403Response(requestId string, errors []ErrorsCommonBadAuthorizationCommonAndEmployerError, ) *SendNegotiationMessage403Response`

NewSendNegotiationMessage403Response instantiates a new SendNegotiationMessage403Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendNegotiationMessage403ResponseWithDefaults

`func NewSendNegotiationMessage403ResponseWithDefaults() *SendNegotiationMessage403Response`

NewSendNegotiationMessage403ResponseWithDefaults instantiates a new SendNegotiationMessage403Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *SendNegotiationMessage403Response) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *SendNegotiationMessage403Response) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *SendNegotiationMessage403Response) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetBadArgument

`func (o *SendNegotiationMessage403Response) GetBadArgument() string`

GetBadArgument returns the BadArgument field if non-nil, zero value otherwise.

### GetBadArgumentOk

`func (o *SendNegotiationMessage403Response) GetBadArgumentOk() (*string, bool)`

GetBadArgumentOk returns a tuple with the BadArgument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArgument

`func (o *SendNegotiationMessage403Response) SetBadArgument(v string)`

SetBadArgument sets BadArgument field to given value.

### HasBadArgument

`func (o *SendNegotiationMessage403Response) HasBadArgument() bool`

HasBadArgument returns a boolean if a field has been set.

### GetBadArguments

`func (o *SendNegotiationMessage403Response) GetBadArguments() []ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments`

GetBadArguments returns the BadArguments field if non-nil, zero value otherwise.

### GetBadArgumentsOk

`func (o *SendNegotiationMessage403Response) GetBadArgumentsOk() (*[]ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments, bool)`

GetBadArgumentsOk returns a tuple with the BadArguments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadArguments

`func (o *SendNegotiationMessage403Response) SetBadArguments(v []ErrorsVacancyApplyBadRequestErrorsAllOfBadArguments)`

SetBadArguments sets BadArguments field to given value.

### HasBadArguments

`func (o *SendNegotiationMessage403Response) HasBadArguments() bool`

HasBadArguments returns a boolean if a field has been set.

### GetDescription

`func (o *SendNegotiationMessage403Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SendNegotiationMessage403Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SendNegotiationMessage403Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SendNegotiationMessage403Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrors

`func (o *SendNegotiationMessage403Response) GetErrors() []ErrorsCommonBadAuthorizationCommonAndEmployerError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SendNegotiationMessage403Response) GetErrorsOk() (*[]ErrorsCommonBadAuthorizationCommonAndEmployerError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SendNegotiationMessage403Response) SetErrors(v []ErrorsCommonBadAuthorizationCommonAndEmployerError)`

SetErrors sets Errors field to given value.


### GetOauthError

`func (o *SendNegotiationMessage403Response) GetOauthError() string`

GetOauthError returns the OauthError field if non-nil, zero value otherwise.

### GetOauthErrorOk

`func (o *SendNegotiationMessage403Response) GetOauthErrorOk() (*string, bool)`

GetOauthErrorOk returns a tuple with the OauthError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthError

`func (o *SendNegotiationMessage403Response) SetOauthError(v string)`

SetOauthError sets OauthError field to given value.

### HasOauthError

`func (o *SendNegotiationMessage403Response) HasOauthError() bool`

HasOauthError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


