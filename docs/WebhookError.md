# WebhookError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Текстовый идентификатор типа ошибки | 
**Value** | **string** | Описание ошибки:   * already_exist - У текущего пользователя/компании уже есть подписка для данного приложения (включая подписку на резюме)  | 

## Methods

### NewWebhookError

`func NewWebhookError(type_ string, value string, ) *WebhookError`

NewWebhookError instantiates a new WebhookError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookErrorWithDefaults

`func NewWebhookErrorWithDefaults() *WebhookError`

NewWebhookErrorWithDefaults instantiates a new WebhookError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WebhookError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookError) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *WebhookError) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *WebhookError) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *WebhookError) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


